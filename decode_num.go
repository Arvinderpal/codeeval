package codeeval

import (
	"fmt"
	"os"
	"strconv"
)

func DecodeNum(str string) int {
	// fmt.Printf("%v.", str)
	if len(str) <= 1 {
		return 1
	}
	num := 0

	for i := 0; i < len(str)-1; i++ {
		val, err := strconv.Atoi(str[i : i+2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		if val > 0 && val < 27 {
			if i+2 == len(str) {
				num += DecodeNum("")
			} else {
				num += DecodeNum(str[i+2:])
			}
		}
	}
	return num + 1 // we add one for the case where is letter is separate. e.g. 2, 3, 6
}
