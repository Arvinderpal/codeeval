package codeeval

func RemoveAllOccurances(s, o []int) []int {
	var newS []int
	oMap := make(map[int]bool, len(o))
	for _, e := range o {
		oMap[e] = true
	}
	for _, e := range s {
		if oMap[e] == true {
			continue
		}
		newS = append(newS, e)
	}
	return newS
}
