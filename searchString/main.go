package main

import (
	"fmt"
	"sort"
)

func main() {
	myArr := []string{"aa", "bb", "cc", "dd"}

	//Version O(N*log(N))
	sort.Strings(myArr)
	if sort.SearchStrings(myArr, "bb") < len(myArr) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}

	//Version O(N)
	for _, str := range myArr {
		if str == "bb" {
			fmt.Println("Found")
			return
		}
	}
	fmt.Println("Not Found")
}
