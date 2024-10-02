package util

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Level struct {
	raw   string
	key   string
	index int
}

var levels []Level

func GetValueFromJson(key string, jsonStr string) (string, error) {
	levels = nil

	//prepare key levels
	raws := strings.Split(key, ".")
	for _, raw := range raws {
		levels = append(levels, Level{
			raw: raw,
		})
	}

	for _, lev := range levels {
		ok, _ := regexp.MatchString("[a-z]+", lev.raw)
		if ok {
			lev.key = lev.raw
			lev.index = -1
			continue
		} else {
			okk, _ := regexp.MatchString("[a-z]+\\[[0-9]+\\]", lev.raw)
			if okk {
				re, err := regexp.Compile("[a-z]")
				if err != nil {
					return "", errors.New("compiling regexp fail 1")
				}
				ke := re.FindString(lev.raw)
				if ke == "" {
					return "", errors.New("found empty key on " + lev.raw + " - " + ke)
				}
				lev.key = ke

				re, err = regexp.Compile("\\[[0-9]+\\]")
				if err != nil {
					return "", errors.New("compiling regexp fail 2")
				}
				ind := re.FindString(lev.raw)
				if ind == "" || ind[0] != '[' || ind[len(ind)-1] != ']' {
					return "", errors.New("found invalid index area " + lev.raw + " - " + ind)
				}
				ind = ind[1 : len(ind)-1]
				inde, err := strconv.Atoi(ind)
				if err != nil {
					return "", err
				}
				lev.index = inde
				continue
			} else {
				return "", errors.New("invalid json key: " + lev.raw)
			}
		}

	}

	//iterate json
	var x map[string]interface{}
	for ind, lev := range levels {
		if err := json.Unmarshal([]byte(jsonStr), &x); err != nil {
			return "", err
		}
		if val, ok := x[lev.key]; ok {
			// if this is the end of val
			if ind == len(levels)-1 {
				switch v := val.(type) {
				case bool:
					return strconv.FormatBool(v), nil
				case int:
					return strconv.Itoa(v), nil
				//case float64:
				//return strconv.FormatFloat(v), nil
				case string:
					return v, nil
				default:
					return "", errors.New("unknow final type")
				}
			}
			//if future iteration required
		} else {
			return "", errors.New(lev.key + " does not exist in level " + strconv.Itoa(ind))
		}
	}
	return "", nil
}
