package string

import "fmt"

/*
Its a long problem , but its about valid parenthesis

() - valid
()( - invalid

	you have to find the maximum depth of the string that is how deeep
	the parenthesis are before closing
	"(1+(2*3)+((8)/4))+1"
	here depth is 3

	logic count ( as 1 and ) as -1 , if at the end it is 0 , then its valid and return
	the max count of ( before )

	https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/description/

*/

func MaxDepth() {

	s := "(1+(2*3)+((8)/4))+1"

	max := 0
	counter := 0

	for _, val := range s {
		if string(val) == "(" {
			counter++
			if max < counter {
				max = counter
			}
		}

		if string(val) == ")" {
			counter--
		}
	}

	fmt.Println("Max depth is", max)
}
