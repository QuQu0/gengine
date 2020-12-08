package map_slice_array

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/google/martian/v3/log"

	"testing"
	"time"
)

type AS struct {
	MI *[3]int
	MM [4]int
}

const a_1 = `
rule "array test" "m dec"
begin

a=1
AS.MI[1] = 22/3 + 1 -10*2 
println("AS.MI[1]--->",AS.MI[1])
println("AS.MI[a]--->",AS.MI[1])

a = 3
AS.MM[3] = 22333
println("AS.MM[3]--->",AS.MM[3])
println("AS.MM[a]--->",AS.MM[a])

a = 1
AA[a] = 10001
println("AA[1]--->",AA[1])
println("AA[a]--->",AA[a])

//can't set value, but can get value
//AAA[1] = 30000
println("AAA[1]---->",AAA[1])
end
`

func Test_Array(t *testing.T) {
	AS := &AS{
		MI: &[3]int{},
		MM: [4]int{},
	}

	var AA [2]int
	AA = [2]int{1, 2}

	var AAA [2]int
	AAA = [2]int{1, 2}

	dataContext := context.NewDataContext()
	dataContext.Add("PrintName", fmt.Println)
	dataContext.Add("AS", AS)
	//single array inject, must be ptr
	dataContext.Add("AA", &AA)
	dataContext.Add("AAA", AAA)
	dataContext.Add("println", fmt.Println)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(a_1)
	end1 := time.Now().UnixNano()

	log.Infof("rules num:%d, load rules cost time:%d ns", len(ruleBuilder.Kc.RuleEntities), end1-start1)

	if err != nil {
		log.Errorf("err:%s ", err)
	} else {
		eng := engine.NewGengine()
		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil {
			log.Errorf("execute rule error: %v", err)
		}
		log.Infof("execute rule cost %d ns", end-start)
	}
}
