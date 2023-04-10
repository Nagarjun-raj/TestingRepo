/* HackerLand University has the following grading policy
>Any  less than 40 is a failing grade
>If the difference between the grade and the next multiple of 5 is less than 3, round grade up to the next multiple of 5.
>Example
> grade=84 round to 85 (85 - 84 is less than 3)
> grade=29 do not round (result is less than 40)
>grade=57 do not round (60 - 57 is 3 or higher)
*/

package main

import "fmt"

func final2(temp map[string]int) map[string]int {
	temp1 := make(map[string]int)
	for k, v := range temp {
		if v <= 40 || v%5 == 0 {
			temp1[k] = v
		} else {
			for i := 0; i <= 10; i++ {
				w := v + i
				if w%5 == 0 && i <= 2 {
					v = w
					temp1[k] = v
				} else if w%5 == 0 && i >= 3 {
					temp1[k] = v
				}
			}
		}
	}

	return temp1
}

func main() {
	temp := map[string]int{
		"Student1": 83,
		"Student2": 59,
		"Student3": 57,
		"Student4": 45,
		"Student5": 25,
	}
	j := final2(temp)
	fmt.Println(j)
}
