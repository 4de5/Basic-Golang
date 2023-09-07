//Break --> untuk mengHENTIKAN
//Continue --> untuk menJEDA

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i > 5 {
			break //BREAK
		}
		fmt.Println(i)
	}

	fmt.Println()
	for index := 0; index <= 20; index++ {
		if index%2 == 0 {
			continue //CONTINUE
		}
		fmt.Println(index)
	}
}
