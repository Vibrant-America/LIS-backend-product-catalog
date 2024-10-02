package util

func Union(a []int, b []int) []int {
	mp := make(map[int]int)
	for _, aa := range a {
		_, ok := mp[aa]
		if ok {
			mp[aa] += 1
		} else {
			mp[aa] = 1
		}
	}

	var res []int
	for _, bb := range b {
		val, ok := mp[bb]
		if ok {
			if val > 0 {
				res = append(res, bb)
				mp[bb] -= 1
			}
		}
	}
	return res
}

func Deduplicate(a []int, b []int) []int {
	mp := make(map[int]int)
	for _, bb := range b {
		_, ok := mp[bb]
		if ok {
			mp[bb] += 1
		} else {
			mp[bb] = 1
		}
	}

	var res []int
	for _, aa := range a {
		val, ok := mp[aa]
		if ok {
			if val > 0 {
				mp[aa] -= 1
			} else {
				res = append(res, aa)
			}
		} else {
			res = append(res, aa)
		}
	}
	return res
}

func LightEqual(a []int, b []int) bool {
	mp := make(map[int]int)
	for _, bb := range b {
		_, ok := mp[bb]
		if ok {
			mp[bb] += 1
		} else {
			mp[bb] = 1
		}
	}

	for _, aa := range a {
		val, ok := mp[aa]
		if ok {
			if val > 0 {
				mp[aa] -= 1
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
