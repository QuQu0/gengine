// Code generated from /Users/renyunyi/go/src/gengine/internal/iantlr/gengine.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 444, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 
	4, 76, 9, 76, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 
	3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 
	23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 
	3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 
	33, 3, 34, 3, 34, 5, 34, 231, 10, 34, 3, 34, 6, 34, 234, 10, 34, 13, 34, 
	14, 34, 235, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 
	36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 
	3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 
	41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 
	45, 6, 45, 289, 10, 45, 13, 45, 14, 45, 290, 3, 45, 7, 45, 294, 10, 45, 
	12, 45, 14, 45, 297, 11, 45, 3, 46, 6, 46, 300, 10, 46, 13, 46, 14, 46, 
	301, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 
	51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 
	3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 59, 3, 
	59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 63, 
	3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 
	68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 
	3, 72, 3, 72, 3, 72, 7, 72, 369, 10, 72, 12, 72, 14, 72, 372, 11, 72, 3, 
	72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 6, 74, 381, 10, 74, 13, 74, 
	14, 74, 382, 5, 74, 385, 10, 74, 3, 74, 3, 74, 6, 74, 389, 10, 74, 13, 
	74, 14, 74, 390, 3, 74, 6, 74, 394, 10, 74, 13, 74, 14, 74, 395, 3, 74, 
	3, 74, 3, 74, 3, 74, 6, 74, 402, 10, 74, 13, 74, 14, 74, 403, 5, 74, 406, 
	10, 74, 3, 74, 3, 74, 6, 74, 410, 10, 74, 13, 74, 14, 74, 411, 3, 74, 3, 
	74, 3, 74, 6, 74, 417, 10, 74, 13, 74, 14, 74, 418, 3, 74, 3, 74, 5, 74, 
	423, 10, 74, 3, 75, 3, 75, 3, 75, 3, 75, 7, 75, 429, 10, 75, 12, 75, 14, 
	75, 432, 11, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 76, 6, 76, 439, 10, 76, 
	13, 76, 14, 76, 440, 3, 76, 3, 76, 3, 430, 2, 77, 3, 3, 5, 4, 7, 5, 9, 
	6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 
	31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 
	2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 8, 71, 9, 
	73, 10, 75, 11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87, 17, 89, 18, 
	91, 19, 93, 20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105, 26, 107, 
	27, 109, 28, 111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121, 34, 123, 
	35, 125, 36, 127, 37, 129, 38, 131, 39, 133, 40, 135, 41, 137, 42, 139, 
	43, 141, 44, 143, 45, 145, 46, 147, 47, 149, 48, 151, 49, 3, 2, 33, 3, 
	2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 
	101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 
	104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 
	107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 
	110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 
	113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 
	116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 
	119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 
	122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 5, 2, 67, 92, 
	97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 36, 36, 94, 
	94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 436, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 69, 3, 
	2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 
	3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 
	85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 
	2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 
	2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 
	3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 
	2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 
	2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 
	129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 
	2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 
	3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 
	2, 151, 3, 2, 2, 2, 3, 153, 3, 2, 2, 2, 5, 158, 3, 2, 2, 2, 7, 161, 3, 
	2, 2, 2, 9, 166, 3, 2, 2, 2, 11, 168, 3, 2, 2, 2, 13, 174, 3, 2, 2, 2, 
	15, 176, 3, 2, 2, 2, 17, 178, 3, 2, 2, 2, 19, 180, 3, 2, 2, 2, 21, 182, 
	3, 2, 2, 2, 23, 184, 3, 2, 2, 2, 25, 186, 3, 2, 2, 2, 27, 188, 3, 2, 2, 
	2, 29, 190, 3, 2, 2, 2, 31, 192, 3, 2, 2, 2, 33, 194, 3, 2, 2, 2, 35, 196, 
	3, 2, 2, 2, 37, 198, 3, 2, 2, 2, 39, 200, 3, 2, 2, 2, 41, 202, 3, 2, 2, 
	2, 43, 204, 3, 2, 2, 2, 45, 206, 3, 2, 2, 2, 47, 208, 3, 2, 2, 2, 49, 210, 
	3, 2, 2, 2, 51, 212, 3, 2, 2, 2, 53, 214, 3, 2, 2, 2, 55, 216, 3, 2, 2, 
	2, 57, 218, 3, 2, 2, 2, 59, 220, 3, 2, 2, 2, 61, 222, 3, 2, 2, 2, 63, 224, 
	3, 2, 2, 2, 65, 226, 3, 2, 2, 2, 67, 228, 3, 2, 2, 2, 69, 237, 3, 2, 2, 
	2, 71, 241, 3, 2, 2, 2, 73, 246, 3, 2, 2, 2, 75, 249, 3, 2, 2, 2, 77, 252, 
	3, 2, 2, 2, 79, 257, 3, 2, 2, 2, 81, 263, 3, 2, 2, 2, 83, 268, 3, 2, 2, 
	2, 85, 277, 3, 2, 2, 2, 87, 283, 3, 2, 2, 2, 89, 288, 3, 2, 2, 2, 91, 299, 
	3, 2, 2, 2, 93, 303, 3, 2, 2, 2, 95, 305, 3, 2, 2, 2, 97, 307, 3, 2, 2, 
	2, 99, 309, 3, 2, 2, 2, 101, 311, 3, 2, 2, 2, 103, 314, 3, 2, 2, 2, 105, 
	316, 3, 2, 2, 2, 107, 318, 3, 2, 2, 2, 109, 321, 3, 2, 2, 2, 111, 324, 
	3, 2, 2, 2, 113, 327, 3, 2, 2, 2, 115, 329, 3, 2, 2, 2, 117, 332, 3, 2, 
	2, 2, 119, 334, 3, 2, 2, 2, 121, 337, 3, 2, 2, 2, 123, 340, 3, 2, 2, 2, 
	125, 343, 3, 2, 2, 2, 127, 346, 3, 2, 2, 2, 129, 348, 3, 2, 2, 2, 131, 
	350, 3, 2, 2, 2, 133, 352, 3, 2, 2, 2, 135, 354, 3, 2, 2, 2, 137, 356, 
	3, 2, 2, 2, 139, 358, 3, 2, 2, 2, 141, 360, 3, 2, 2, 2, 143, 362, 3, 2, 
	2, 2, 145, 375, 3, 2, 2, 2, 147, 422, 3, 2, 2, 2, 149, 424, 3, 2, 2, 2, 
	151, 438, 3, 2, 2, 2, 153, 154, 7, 101, 2, 2, 154, 155, 7, 113, 2, 2, 155, 
	156, 7, 112, 2, 2, 156, 157, 7, 101, 2, 2, 157, 4, 3, 2, 2, 2, 158, 159, 
	7, 107, 2, 2, 159, 160, 7, 104, 2, 2, 160, 6, 3, 2, 2, 2, 161, 162, 7, 
	103, 2, 2, 162, 163, 7, 110, 2, 2, 163, 164, 7, 117, 2, 2, 164, 165, 7, 
	103, 2, 2, 165, 8, 3, 2, 2, 2, 166, 167, 7, 46, 2, 2, 167, 10, 3, 2, 2, 
	2, 168, 169, 7, 66, 2, 2, 169, 170, 7, 112, 2, 2, 170, 171, 7, 99, 2, 2, 
	171, 172, 7, 111, 2, 2, 172, 173, 7, 103, 2, 2, 173, 12, 3, 2, 2, 2, 174, 
	175, 9, 2, 2, 2, 175, 14, 3, 2, 2, 2, 176, 177, 9, 3, 2, 2, 177, 16, 3, 
	2, 2, 2, 178, 179, 9, 4, 2, 2, 179, 18, 3, 2, 2, 2, 180, 181, 9, 5, 2, 
	2, 181, 20, 3, 2, 2, 2, 182, 183, 9, 6, 2, 2, 183, 22, 3, 2, 2, 2, 184, 
	185, 9, 7, 2, 2, 185, 24, 3, 2, 2, 2, 186, 187, 9, 8, 2, 2, 187, 26, 3, 
	2, 2, 2, 188, 189, 9, 9, 2, 2, 189, 28, 3, 2, 2, 2, 190, 191, 9, 10, 2, 
	2, 191, 30, 3, 2, 2, 2, 192, 193, 9, 11, 2, 2, 193, 32, 3, 2, 2, 2, 194, 
	195, 9, 12, 2, 2, 195, 34, 3, 2, 2, 2, 196, 197, 9, 13, 2, 2, 197, 36, 
	3, 2, 2, 2, 198, 199, 9, 14, 2, 2, 199, 38, 3, 2, 2, 2, 200, 201, 9, 15, 
	2, 2, 201, 40, 3, 2, 2, 2, 202, 203, 9, 16, 2, 2, 203, 42, 3, 2, 2, 2, 
	204, 205, 9, 17, 2, 2, 205, 44, 3, 2, 2, 2, 206, 207, 9, 18, 2, 2, 207, 
	46, 3, 2, 2, 2, 208, 209, 9, 19, 2, 2, 209, 48, 3, 2, 2, 2, 210, 211, 9, 
	20, 2, 2, 211, 50, 3, 2, 2, 2, 212, 213, 9, 21, 2, 2, 213, 52, 3, 2, 2, 
	2, 214, 215, 9, 22, 2, 2, 215, 54, 3, 2, 2, 2, 216, 217, 9, 23, 2, 2, 217, 
	56, 3, 2, 2, 2, 218, 219, 9, 24, 2, 2, 219, 58, 3, 2, 2, 2, 220, 221, 9, 
	25, 2, 2, 221, 60, 3, 2, 2, 2, 222, 223, 9, 26, 2, 2, 223, 62, 3, 2, 2, 
	2, 224, 225, 9, 27, 2, 2, 225, 64, 3, 2, 2, 2, 226, 227, 9, 28, 2, 2, 227, 
	66, 3, 2, 2, 2, 228, 230, 7, 71, 2, 2, 229, 231, 7, 47, 2, 2, 230, 229, 
	3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232, 234, 5, 13, 
	7, 2, 233, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 
	235, 236, 3, 2, 2, 2, 236, 68, 3, 2, 2, 2, 237, 238, 5, 41, 21, 2, 238, 
	239, 5, 31, 16, 2, 239, 240, 5, 37, 19, 2, 240, 70, 3, 2, 2, 2, 241, 242, 
	5, 49, 25, 2, 242, 243, 5, 55, 28, 2, 243, 244, 5, 37, 19, 2, 244, 245, 
	5, 23, 12, 2, 245, 72, 3, 2, 2, 2, 246, 247, 7, 40, 2, 2, 247, 248, 7, 
	40, 2, 2, 248, 74, 3, 2, 2, 2, 249, 250, 7, 126, 2, 2, 250, 251, 7, 126, 
	2, 2, 251, 76, 3, 2, 2, 2, 252, 253, 5, 53, 27, 2, 253, 254, 5, 49, 25, 
	2, 254, 255, 5, 55, 28, 2, 255, 256, 5, 23, 12, 2, 256, 78, 3, 2, 2, 2, 
	257, 258, 5, 25, 13, 2, 258, 259, 5, 15, 8, 2, 259, 260, 5, 37, 19, 2, 
	260, 261, 5, 51, 26, 2, 261, 262, 5, 23, 12, 2, 262, 80, 3, 2, 2, 2, 263, 
	264, 5, 41, 21, 2, 264, 265, 5, 55, 28, 2, 265, 266, 5, 37, 19, 2, 266, 
	267, 5, 37, 19, 2, 267, 82, 3, 2, 2, 2, 268, 269, 5, 51, 26, 2, 269, 270, 
	5, 15, 8, 2, 270, 271, 5, 37, 19, 2, 271, 272, 5, 31, 16, 2, 272, 273, 
	5, 23, 12, 2, 273, 274, 5, 41, 21, 2, 274, 275, 5, 19, 10, 2, 275, 276, 
	5, 23, 12, 2, 276, 84, 3, 2, 2, 2, 277, 278, 5, 17, 9, 2, 278, 279, 5, 
	23, 12, 2, 279, 280, 5, 27, 14, 2, 280, 281, 5, 31, 16, 2, 281, 282, 5, 
	41, 21, 2, 282, 86, 3, 2, 2, 2, 283, 284, 5, 23, 12, 2, 284, 285, 5, 41, 
	21, 2, 285, 286, 5, 21, 11, 2, 286, 88, 3, 2, 2, 2, 287, 289, 9, 29, 2, 
	2, 288, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 
	291, 3, 2, 2, 2, 291, 295, 3, 2, 2, 2, 292, 294, 9, 30, 2, 2, 293, 292, 
	3, 2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 295, 296, 3, 2, 
	2, 2, 296, 90, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 298, 300, 4, 50, 59, 2, 
	299, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 
	302, 3, 2, 2, 2, 302, 92, 3, 2, 2, 2, 303, 304, 7, 45, 2, 2, 304, 94, 3, 
	2, 2, 2, 305, 306, 7, 47, 2, 2, 306, 96, 3, 2, 2, 2, 307, 308, 7, 49, 2, 
	2, 308, 98, 3, 2, 2, 2, 309, 310, 7, 44, 2, 2, 310, 100, 3, 2, 2, 2, 311, 
	312, 7, 63, 2, 2, 312, 313, 7, 63, 2, 2, 313, 102, 3, 2, 2, 2, 314, 315, 
	7, 64, 2, 2, 315, 104, 3, 2, 2, 2, 316, 317, 7, 62, 2, 2, 317, 106, 3, 
	2, 2, 2, 318, 319, 7, 64, 2, 2, 319, 320, 7, 63, 2, 2, 320, 108, 3, 2, 
	2, 2, 321, 322, 7, 62, 2, 2, 322, 323, 7, 63, 2, 2, 323, 110, 3, 2, 2, 
	2, 324, 325, 7, 35, 2, 2, 325, 326, 7, 63, 2, 2, 326, 112, 3, 2, 2, 2, 
	327, 328, 7, 35, 2, 2, 328, 114, 3, 2, 2, 2, 329, 330, 7, 60, 2, 2, 330, 
	331, 7, 63, 2, 2, 331, 116, 3, 2, 2, 2, 332, 333, 7, 63, 2, 2, 333, 118, 
	3, 2, 2, 2, 334, 335, 7, 45, 2, 2, 335, 336, 7, 63, 2, 2, 336, 120, 3, 
	2, 2, 2, 337, 338, 7, 47, 2, 2, 338, 339, 7, 63, 2, 2, 339, 122, 3, 2, 
	2, 2, 340, 341, 7, 44, 2, 2, 341, 342, 7, 63, 2, 2, 342, 124, 3, 2, 2, 
	2, 343, 344, 7, 49, 2, 2, 344, 345, 7, 63, 2, 2, 345, 126, 3, 2, 2, 2, 
	346, 347, 7, 93, 2, 2, 347, 128, 3, 2, 2, 2, 348, 349, 7, 95, 2, 2, 349, 
	130, 3, 2, 2, 2, 350, 351, 7, 61, 2, 2, 351, 132, 3, 2, 2, 2, 352, 353, 
	7, 125, 2, 2, 353, 134, 3, 2, 2, 2, 354, 355, 7, 127, 2, 2, 355, 136, 3, 
	2, 2, 2, 356, 357, 7, 42, 2, 2, 357, 138, 3, 2, 2, 2, 358, 359, 7, 43, 
	2, 2, 359, 140, 3, 2, 2, 2, 360, 361, 7, 48, 2, 2, 361, 142, 3, 2, 2, 2, 
	362, 370, 7, 36, 2, 2, 363, 364, 7, 94, 2, 2, 364, 369, 11, 2, 2, 2, 365, 
	366, 7, 36, 2, 2, 366, 369, 7, 36, 2, 2, 367, 369, 10, 31, 2, 2, 368, 363, 
	3, 2, 2, 2, 368, 365, 3, 2, 2, 2, 368, 367, 3, 2, 2, 2, 369, 372, 3, 2, 
	2, 2, 370, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 373, 3, 2, 2, 2, 
	372, 370, 3, 2, 2, 2, 373, 374, 7, 36, 2, 2, 374, 144, 3, 2, 2, 2, 375, 
	376, 5, 89, 45, 2, 376, 377, 5, 141, 71, 2, 377, 378, 5, 89, 45, 2, 378, 
	146, 3, 2, 2, 2, 379, 381, 5, 13, 7, 2, 380, 379, 3, 2, 2, 2, 381, 382, 
	3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 385, 3, 2, 
	2, 2, 384, 380, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 
	386, 388, 7, 48, 2, 2, 387, 389, 5, 13, 7, 2, 388, 387, 3, 2, 2, 2, 389, 
	390, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 423, 
	3, 2, 2, 2, 392, 394, 5, 13, 7, 2, 393, 392, 3, 2, 2, 2, 394, 395, 3, 2, 
	2, 2, 395, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 
	397, 398, 7, 48, 2, 2, 398, 399, 5, 67, 34, 2, 399, 423, 3, 2, 2, 2, 400, 
	402, 5, 13, 7, 2, 401, 400, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 401, 
	3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 406, 3, 2, 2, 2, 405, 401, 3, 2, 
	2, 2, 405, 406, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 409, 7, 48, 2, 2, 
	408, 410, 5, 13, 7, 2, 409, 408, 3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 
	409, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 414, 
	5, 67, 34, 2, 414, 423, 3, 2, 2, 2, 415, 417, 5, 13, 7, 2, 416, 415, 3, 
	2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 418, 419, 3, 2, 2, 
	2, 419, 420, 3, 2, 2, 2, 420, 421, 5, 67, 34, 2, 421, 423, 3, 2, 2, 2, 
	422, 384, 3, 2, 2, 2, 422, 393, 3, 2, 2, 2, 422, 405, 3, 2, 2, 2, 422, 
	416, 3, 2, 2, 2, 423, 148, 3, 2, 2, 2, 424, 425, 7, 49, 2, 2, 425, 426, 
	7, 49, 2, 2, 426, 430, 3, 2, 2, 2, 427, 429, 11, 2, 2, 2, 428, 427, 3, 
	2, 2, 2, 429, 432, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 430, 428, 3, 2, 2, 
	2, 431, 433, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 433, 434, 7, 12, 2, 2, 434, 
	435, 3, 2, 2, 2, 435, 436, 8, 75, 2, 2, 436, 150, 3, 2, 2, 2, 437, 439, 
	9, 32, 2, 2, 438, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 438, 3, 2, 
	2, 2, 440, 441, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 8, 76, 2, 2, 
	443, 152, 3, 2, 2, 2, 22, 2, 230, 235, 290, 293, 295, 301, 368, 370, 382, 
	384, 390, 395, 403, 405, 411, 418, 422, 430, 440, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'conc'", "'if'", "'else'", "','", "'@name'", "", "", "'&&'", "'||'", 
	"", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'", "'=='", "'>'", 
	"'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='", "'+='", "'-='", "'*='", 
	"'/='", "'['", "']'", "';'", "'{'", "'}'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL", 
	"SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", 
	"MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", 
	"SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", 
	"SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", 
	"DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "DEC_DIGIT", "A", "B", "C", "D", 
	"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", 
	"T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "NIL", "RULE", 
	"AND", "OR", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", 
	"SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", 
	"GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", 
	"MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", 
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", 
	"REAL_LITERAL", "SL_COMMENT", "WS",
}

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewgengineLexer(input antlr.CharStream) *gengineLexer {

	l := new(gengineLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0 = 1
	gengineLexerT__1 = 2
	gengineLexerT__2 = 3
	gengineLexerT__3 = 4
	gengineLexerT__4 = 5
	gengineLexerNIL = 6
	gengineLexerRULE = 7
	gengineLexerAND = 8
	gengineLexerOR = 9
	gengineLexerTRUE = 10
	gengineLexerFALSE = 11
	gengineLexerNULL_LITERAL = 12
	gengineLexerSALIENCE = 13
	gengineLexerBEGIN = 14
	gengineLexerEND = 15
	gengineLexerSIMPLENAME = 16
	gengineLexerINT = 17
	gengineLexerPLUS = 18
	gengineLexerMINUS = 19
	gengineLexerDIV = 20
	gengineLexerMUL = 21
	gengineLexerEQUALS = 22
	gengineLexerGT = 23
	gengineLexerLT = 24
	gengineLexerGTE = 25
	gengineLexerLTE = 26
	gengineLexerNOTEQUALS = 27
	gengineLexerNOT = 28
	gengineLexerASSIGN = 29
	gengineLexerSET = 30
	gengineLexerPLUSEQUAL = 31
	gengineLexerMINUSEQUAL = 32
	gengineLexerMULTIEQUAL = 33
	gengineLexerDIVEQUAL = 34
	gengineLexerLSQARE = 35
	gengineLexerRSQARE = 36
	gengineLexerSEMICOLON = 37
	gengineLexerLR_BRACE = 38
	gengineLexerRR_BRACE = 39
	gengineLexerLR_BRACKET = 40
	gengineLexerRR_BRACKET = 41
	gengineLexerDOT = 42
	gengineLexerDQUOTA_STRING = 43
	gengineLexerDOTTEDNAME = 44
	gengineLexerREAL_LITERAL = 45
	gengineLexerSL_COMMENT = 46
	gengineLexerWS = 47
)

