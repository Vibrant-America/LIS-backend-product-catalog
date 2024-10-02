package util

func MapMergeStringFloat(a, b map[string]float64) {
	for k, v := range b {
		a[k] = v
	}
}
