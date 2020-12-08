package engine

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/internal/base"
	"gengine/internal/core/errors"
	parser "gengine/internal/iantlr/alr"
	"gengine/internal/iparser"
	"gengine/internal/tool"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"sync"

	"github.com/google/martian/v3/log"
)

const (
	SORT_MODEL        = 1
	CONCOURRENT_MODEL = 2
	MIX_MODEL         = 3
	INVERSE_MIX_MODEL = 4
)

// when you use NewGenginePool, you just think of it as the connection pool of mysql, the higher QPS you want to support, the more resource you need to give
type GenginePool struct {
	runningLock  sync.Mutex
	freeGengines []*gengineWrapper

	//just for check whether a rule exist
	ruleBuilder *builder.RuleBuilder

	execModel int
	apis      map[string]interface{}

	additionLock     sync.Mutex
	additionGengines []*gengineWrapper
	additionNum      int64

	updateLock sync.Mutex
	clear      bool //whether rules has been cleared ，if true it means there is no rules in gengine

	rbSlice []*builder.RuleBuilder
	//total gengine instance number
	max int64

	getEngineLock sync.RWMutex //just one can get this lock
}

type gengineWrapper struct {
	tag         int64 // one to one between the ruleBuilder slice
	rulebuilder *builder.RuleBuilder
	gengine     *Gengine

	addition bool // when gengine resource is not enough and poollength >  minPool  and  poollength < maxPool, new gengine will be create, and it will be tagged addition=true; when poollength <  minPool it will be tagged addition=false
}

func (gw *gengineWrapper) clearInjected(keys ...string) {
	if gw == nil || gw.rulebuilder == nil || gw.rulebuilder.Dc == nil {
		return
	}
	gw.rulebuilder.Dc.Del(keys...)
}

//poolLen  -> gengine pool length to init
//em       -> rule execute model: 1 sort model; 2 :concurrent model; 3 mix model
//rStr  -> rules string
//apiOuter -> which user want to add rule engine to use
// just init once!!!

// best practise：
// when there has cost-time operate in your rule or you want to support high concurrency(> 200000QPS) , please set poolMinLen bigger Appropriately
// when you use NewGenginePool,you just think of it as the connection pool of mysql, the higher QPS you want to support, the more resource you need to give
func NewGenginePool(poolMinLen, poolMaxLen int64, em int, rulesStr string, apiOuter map[string]interface{}) (*GenginePool, error) {

	if !(0 < poolMinLen && poolMinLen < poolMaxLen) {
		return nil, errors.New("pool length must be bigger than 0, and poolMaxLen must bigger than poolMinLen")
	}

	if em != SORT_MODEL && em != CONCOURRENT_MODEL && em != MIX_MODEL && em != INVERSE_MIX_MODEL {
		return nil, errors.New(fmt.Sprintf("exec model must be SORT_MODEL(1) or CONCOURRENT_MODEL(2) or MIX_MODEL(3) or INVERSE_MIX_MODEL(4), now it is %d", em))
	}

	fg := make([]*gengineWrapper, poolMinLen)
	for i := int64(0); i < poolMinLen; i++ {
		fg[i] = &gengineWrapper{
			tag:      i,
			gengine:  NewGengine(),
			addition: false,
		}
	}

	ag := make([]*gengineWrapper, poolMaxLen-poolMinLen)
	for j := int64(0); j < poolMaxLen-poolMinLen; j++ {
		ag[j] = &gengineWrapper{
			tag:      j + poolMinLen,
			gengine:  NewGengine(),
			addition: true,
		}
	}

	srcRb, e := makeRuleBuilder(rulesStr, apiOuter)
	if e != nil {
		return nil, e
	}

	rbs := make([]*builder.RuleBuilder, poolMaxLen)
	for i := 0; i < int(poolMaxLen); i++ {
		rb, e := makeRuleBuilder(rulesStr, apiOuter)
		if e != nil {
			return nil, e
		}
		rbs[i] = rb
	}

	p := &GenginePool{
		ruleBuilder:      srcRb,
		freeGengines:     fg,
		apis:             apiOuter,
		execModel:        em,
		additionNum:      poolMaxLen - poolMinLen,
		additionGengines: ag,
		clear:            false,
		rbSlice:          rbs,
		max:              poolMaxLen,
	}
	return p, nil
}

//this could ensure make thread safety!
func makeRuleBuilder(ruleStr string, apiOuter map[string]interface{}) (*builder.RuleBuilder, error) {
	dataContext := context.NewDataContext()
	if apiOuter != nil {
		for k, v := range apiOuter {
			dataContext.Add(k, v)
		}
	}

	rb := builder.NewRuleBuilder(dataContext)
	if ruleStr != "" {
		if e := rb.BuildRuleFromString(ruleStr); e != nil {
			return nil, errors.New(fmt.Sprintf("build rule from string err: %+v", e))
		}
	} else {
		log.Infof("the ruleStr is \"\"")
	}
	return rb, nil
}

// if there is no enough gengine source, no request will take a lock
func (gp *GenginePool) getGengine() (*gengineWrapper, error) {

	for {
		gp.getEngineLock.Lock()
		//check if there has enough resource in pool
		numFree := len(gp.freeGengines)
		if numFree > 0 {
			gp.runningLock.Lock()
			gw := gp.freeGengines[0]
			gp.freeGengines = gp.freeGengines[1:]
			gp.runningLock.Unlock()
			gp.getEngineLock.Unlock()
			return gw, nil
		}

		//check if there has addition resource
		numAddition := len(gp.additionGengines)
		if numAddition > 0 {
			gp.additionLock.Lock()
			gw := gp.additionGengines[0]
			gp.additionGengines = gp.additionGengines[1:]
			gp.additionLock.Unlock()
			gp.getEngineLock.Unlock()
			return gw, nil
		}

		gp.getEngineLock.Unlock()
	}
}

// async return gengine resource to pool,and update the rules
func (gp *GenginePool) putGengineLocked(gw *gengineWrapper) {
	//addition resource
	go func() {
		if gw.addition {
			gp.additionLock.Lock()
			gp.additionGengines = append(gp.additionGengines, gw)
			gp.additionLock.Unlock()
		} else {
			gp.runningLock.Lock()
			gp.freeGengines = append(gp.freeGengines, gw)
			gp.runningLock.Unlock()
		}
	}()
}

//sync method
//update the all rules in all engine in the pool
//update success: return nil
//update failed: return error
// this is very different from connection pool,
//connection pool just need to init once
//while gengine pool need to update the rules whenever the user want to init
func (gp *GenginePool) UpdatePooledRules(ruleStr string) error {
	//check the rules
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()

	rbi, e := makeRuleBuilder(ruleStr, gp.apis)
	if e != nil {
		return e
	} else {
		if len(rbi.Kc.RuleEntities) == 0 {
			return errors.New(fmt.Sprintf("if you want to clear all rules, use method \"pool.ClearPoolRules()\""))
		}

		//new instance array
		rbs := make([]*builder.RuleBuilder, gp.max)
		for i := 0; i < int(gp.max); i++ {
			rb, e := makeRuleBuilder(ruleStr, gp.apis)
			if e != nil {
				return e
			}
			rbs[i] = rb
		}

		//update instance
		gp.rbSlice = rbs
		//update core
		gp.ruleBuilder = rbi
		gp.clear = false
		return nil
	}
}

func getKc(ruleString string) (*base.KnowledgeContext, error) {

	in := antlr.NewInputStream(ruleString)
	lexer := parser.NewgengineLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	kc := base.NewKnowledgeContext()

	listener := iparser.NewGengineParserListener(kc)
	psr := parser.NewgengineParser(stream)
	psr.BuildParseTrees = true

	errListener := iparser.NewGengineErrorListener()
	psr.AddErrorListener(errListener)
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Primary())

	if len(errListener.GrammarErrors) > 0 {
		return nil, errors.New(fmt.Sprintf("%+v", errListener.GrammarErrors))
	}

	if len(listener.ParseErrors) > 0 {
		return nil, errors.New(fmt.Sprintf("%+v", listener.ParseErrors))
	}

	if len(kc.RuleEntities) == 0 {
		return nil, errors.New(fmt.Sprintf("no rules to update or add."))
	}

	return kc, nil
}

func updateIncremental(kc *base.KnowledgeContext, rb *builder.RuleBuilder) {
	//copy
	newRuleEntities := make(map[string]*base.RuleEntity, len(rb.Kc.RuleEntities))
	for mk, mv := range rb.Kc.RuleEntities {
		newRuleEntities[mk] = mv
	}

	//copy
	newSortRules := make([]*base.RuleEntity, len(rb.Kc.SortRules))
	for sk, sv := range rb.Kc.SortRules {
		newSortRules[sk] = sv
	}

	//kc store the new rules
	for k, v := range kc.RuleEntities {
		//init
		v.Initialize(rb.Dc)

		if vm, ok := newRuleEntities[k]; ok {
			//repalce update
			//search
			index := rb.Kc.SortRulesIndexMap[v.RuleName]
			if v.Salience == vm.Salience {
				//replace
				newSortRules[index] = v
			} else {
				newSortRules := append(newSortRules[:index], newSortRules[index+1:]...)
				low, mid := tool.BinarySearch(newSortRules, v.Salience)

				ire := []*base.RuleEntity{v}
				if mid == 0 {
					newRe := append(ire, newSortRules[low:]...)
					newSortRules = append(newSortRules[:low], newRe...)
				} else {
					newRe := append(ire, newSortRules[mid:]...)
					newSortRules = append(newSortRules[:mid], newRe...)
				}

				//update the sort index
				indexMap := make(map[string]int)
				for k, v := range newSortRules {
					indexMap[v.RuleName] = k
				}
				rb.Kc.SortRulesIndexMap = indexMap
			}

			newRuleEntities[k] = v
		} else {
			//add update
			low, mid := tool.BinarySearch(newSortRules, v.Salience)

			ire := []*base.RuleEntity{v}
			if mid == 0 {
				newRe := append(ire, newSortRules[low:]...)
				newSortRules = append(newSortRules[:low], newRe...)
			} else {
				newRe := append(ire, newSortRules[mid:]...)
				newSortRules = append(newSortRules[:mid], newRe...)
			}

			//update the sort index
			indexMap := make(map[string]int)
			for k, v := range newSortRules {
				indexMap[v.RuleName] = k
			}
			rb.Kc.SortRulesIndexMap = indexMap

			newRuleEntities[k] = v
		}
	}

	rb.Kc.RuleEntities = newRuleEntities
	rb.Kc.SortRules = newSortRules
}

//sync method
//incremental update the rules in all engine in the pool
//incremental update success: return nil
//incremental update failed: return error
// if a rule already exists, this method will use the new rule to replace the old one
// if a rule doesn't exist, this method will add the new rule to the existed rules list
//see: func (builder *RuleBuilder)BuildRuleWithIncremental(ruleString string) in rule_builder.go
func (gp *GenginePool) UpdatePooledRulesIncremental(ruleStr string) error {
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()

	//new main
	kci, e := getKc(ruleStr)
	if e != nil {
		return e
	}

	//new instance
	kcs := make([]*base.KnowledgeContext, gp.max)
	for i := 0; i < int(gp.max); i++ {
		kc, e := getKc(ruleStr)
		if e != nil {
			return e
		}
		kcs[i] = kc
	}

	//update main
	updateIncremental(kci, gp.ruleBuilder)

	//update instance
	for i := 0; i < int(gp.max); i++ {
		updateIncremental(kcs[i], gp.rbSlice[i])
	}

	gp.clear = false
	return nil
}

//clear all rules in engine in pool
func (gp *GenginePool) ClearPoolRules() {
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()
	gp.ruleBuilder = nil
	gp.clear = true
	for i := 0; i < int(gp.max); i++ {
		gp.rbSlice[i].Kc.ClearRules()
	}
}

/*
1 sort model
2 concurrent model
3 mix model
4 inverse mix model
*/
func (gp *GenginePool) SetExecModel(execModel int) error {
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()
	if execModel != SORT_MODEL && execModel != CONCOURRENT_MODEL && execModel != MIX_MODEL && execModel != INVERSE_MIX_MODEL {
		return errors.New(fmt.Sprintf("exec model must be SORT_MODEL(1) or CONCOURRENT_MODEL(2) or MIX_MODEL(3) or INVERSE_MIX_MODEL(4), now it is %d", execModel))
	} else {
		gp.execModel = execModel
	}
	return nil
}

func (gp *GenginePool) GetExecModel() int {
	return gp.execModel
}

//check the rule whether exist
func (gp *GenginePool) IsExist(ruleName string) bool {
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()

	if gp.clear || gp.ruleBuilder == nil {
		return false
	}
	_, ok := gp.ruleBuilder.Kc.RuleEntities[ruleName]
	return ok
}

func (gp *GenginePool) GetRulesNumber() int {
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()

	if gp.clear || gp.ruleBuilder == nil {
		return 0
	}
	return len(gp.ruleBuilder.Kc.RuleEntities)
}

func (gp *GenginePool) prepare(reqName string, req interface{}, respName string, resp interface{}) (*gengineWrapper, error) {
	//get gengine resource
	gw, e := gp.getGengine()
	if e != nil {
		return nil, e
	}

	gw.rulebuilder = gp.rbSlice[gw.tag]

	if reqName != "" && req != nil {
		gw.rulebuilder.Dc.Add(reqName, req)
	}

	if respName != "" && resp != nil {
		gw.rulebuilder.Dc.Add(respName, resp)
	}
	return gw, nil
}

func (gp *GenginePool) prepareWithMultiInput(data map[string]interface{}) (*gengineWrapper, error) {
	//get gengine resource
	gw, e := gp.getGengine()
	if e != nil {
		return nil, e
	}

	gw.rulebuilder = gp.rbSlice[gw.tag]

	for k, v := range data {
		//user should not inject "" string or nil value
		if k != "" && v != nil {
			gw.rulebuilder.Dc.Add(k, v)
		} else {
			log.Errorf("you should not input null string or nil value")
		}
	}

	return gw, nil
}

//execute rules as the user set execute model when init or update
//req, it is better to be ptr, or you will not get changed data
//resp, it is better to be ptr, or you will not get changed dat
func (gp *GenginePool) ExecuteRules(reqName string, req interface{}, respName string, resp interface{}) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepare(reqName, req, respName, resp)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.rulebuilder.Dc.Del(reqName, respName)
		gp.putGengineLocked(gw)
	}()

	if gp.execModel == SORT_MODEL { //sort
		// when some rule execute error ,it will continue to execute last
		return gw.gengine.Execute(gw.rulebuilder, true)
	}

	if gp.execModel == CONCOURRENT_MODEL { //concurrent
		return gw.gengine.ExecuteConcurrent(gw.rulebuilder)
	}

	if gp.execModel == MIX_MODEL { //mix
		return gw.gengine.ExecuteMixModel(gw.rulebuilder)
	}

	if gp.execModel == INVERSE_MIX_MODEL { // inverse mix model
		return gw.gengine.ExecuteInverseMixModel(gw.rulebuilder)
	}

	return nil
}

/**
user can input more data to use in engine

it is no difference with ExecuteRules, you just can inject more data use this api
*/
func (gp *GenginePool) ExecuteRulesWithMultiInput(data map[string]interface{}) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	if gp.execModel == SORT_MODEL { //sort
		// when some rule execute error ,it will continue to execute last
		return gw.gengine.Execute(gw.rulebuilder, true)
	}

	if gp.execModel == CONCOURRENT_MODEL { //concurrent
		return gw.gengine.ExecuteConcurrent(gw.rulebuilder)
	}

	if gp.execModel == MIX_MODEL { //mix
		return gw.gengine.ExecuteMixModel(gw.rulebuilder)
	}

	if gp.execModel == INVERSE_MIX_MODEL { // inverse mix model
		return gw.gengine.ExecuteInverseMixModel(gw.rulebuilder)
	}

	return nil

}

/**
you can use stag to control the gengine execute rules behavior out of pool
if you know what you are doing, it will improve your rules execute performance

if you want to know more about stag, to see the note above every method in Gengine.go
*/
//req, it is better to be ptr, or you will not get changed data
//resp, it is better to be ptr, or you will not get changed data
func (gp *GenginePool) ExecuteRulesWithStopTag(reqName string, req interface{}, respName string, resp interface{}, stag *Stag) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepare(reqName, req, respName, resp)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(reqName, reqName)
		gp.putGengineLocked(gw)
	}()

	if gp.execModel == SORT_MODEL { //sort
		// when some rule execute error ,it will continue to execute last
		return gw.gengine.ExecuteWithStopTagDirect(gw.rulebuilder, true, stag)
	}

	if gp.execModel == CONCOURRENT_MODEL { //concurrent
		return gw.gengine.ExecuteConcurrent(gw.rulebuilder)
	}

	if gp.execModel == MIX_MODEL { //mix
		return gw.gengine.ExecuteMixModelWithStopTagDirect(gw.rulebuilder, stag)
	}

	if gp.execModel == INVERSE_MIX_MODEL { // inverse mix model
		return gw.gengine.ExecuteInverseMixModel(gw.rulebuilder)
	}

	return nil
}

func (gp *GenginePool) ExecuteRulesWithMultiInputAndStopTag(data map[string]interface{}, stag *Stag) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	if gp.execModel == SORT_MODEL { //sort
		// when some rule execute error ,it will continue to execute last
		return gw.gengine.ExecuteWithStopTagDirect(gw.rulebuilder, true, stag)
	}

	if gp.execModel == CONCOURRENT_MODEL { //concurrent
		return gw.gengine.ExecuteConcurrent(gw.rulebuilder)
	}

	if gp.execModel == MIX_MODEL { //mix
		return gw.gengine.ExecuteMixModelWithStopTagDirect(gw.rulebuilder, stag)
	}

	if gp.execModel == INVERSE_MIX_MODEL { // inverse mix model
		return gw.gengine.ExecuteInverseMixModel(gw.rulebuilder)
	}

	return nil
}

/**
see ExecuteSelectedRules in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesWithMultiInput(data map[string]interface{}, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteSelectedRules(gw.rulebuilder, names)
}

/**
see ExecuteSelectedRulesWithControl in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesWithControlWithMultiInput(data map[string]interface{}, b bool, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteSelectedRulesWithControl(gw.rulebuilder, b, names)
}

/**
see ExecuteSelectedRulesWithControlAndStopTag in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesWithControlAndStopTagWithMultiInput(data map[string]interface{}, b bool, stag *Stag, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteSelectedRulesWithControlAndStopTag(gw.rulebuilder, b, stag, names)
}

/**
see ExecuteSelectedRulesConcurrent in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesConcurrentWithMultiInput(data map[string]interface{}, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteSelectedRulesConcurrent(gw.rulebuilder, names)
}

/**
see ExecuteSelectedRulesMixModel in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesMixModelWithMultiInput(data map[string]interface{}, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteSelectedRulesMixModel(gw.rulebuilder, names)
}

// see ExecuteInverseMixModel in gengine.go
func (gp *GenginePool) ExecuteInverseMixModelWithMultiInput(data map[string]interface{}) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteInverseMixModel(gw.rulebuilder)
}

//see ExecuteInverseMixModelWithSelected in gengine.go
func (gp *GenginePool) ExecuteSelectedRulesInverseMixModelWithMultiInput(data map[string]interface{}, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	return gw.gengine.ExecuteSelectedRulesInverseMixModel(gw.rulebuilder, names)
}

/***
this make user could use exemodel to control the select-exemodel
*/
func (gp *GenginePool) ExecuteSelected(data map[string]interface{}, names []string) error {
	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer func() {
		gw.clearInjected(getKeys(data)...)
		gp.putGengineLocked(gw)
	}()

	if gp.execModel == SORT_MODEL {
		return gw.gengine.ExecuteSelectedRules(gw.rulebuilder, names)
	}

	if gp.execModel == CONCOURRENT_MODEL {
		return gw.gengine.ExecuteSelectedRulesConcurrent(gw.rulebuilder, names)
	}

	if gp.execModel == MIX_MODEL {
		return gw.gengine.ExecuteSelectedRulesMixModel(gw.rulebuilder, names)
	}

	if gp.execModel == INVERSE_MIX_MODEL {
		return gw.gengine.ExecuteSelectedRulesInverseMixModel(gw.rulebuilder, names)
	}

	return nil
}

func getKeys(data map[string]interface{}) []string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}
