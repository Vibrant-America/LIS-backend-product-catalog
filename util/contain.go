package util

func StringContain(a []string, b string) bool {
	for _, aa := range a {
		if aa == b {
			return true
		}
	}
	return false
}

func RemoveDuplicateInt(intSlice []int32) []int32 {
	allKeys := make(map[int32]bool)
	list := []int32{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func IntOverLap(a []int, b []int) bool {
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				return true
			}
		}
	}
	return false
}

func StringOverLap(a []string, b []string) bool {
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				return true
			}
		}
	}
	return false
}
