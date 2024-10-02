package util

import (
	"encoding/json"
	"math"
)

// SwapTo load any JSON into corresponding struct
func SwapTo(request, model interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, model)
	return
}

func Round2(f1 float64) float64 {
	return math.Round(f1*100) / 100
}
