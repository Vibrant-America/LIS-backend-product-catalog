package util

import (
	"strconv"
	"strings"
)

func ReadCharge(uri string) (string, int64, error) {
	res := strings.Split(uri, "::")

	charge_type_id, err := strconv.Atoi(res[1])
	if err != nil {
		return "", 0, err
	}
	charge_type := res[0]

	return charge_type, int64(charge_type_id), nil
}
