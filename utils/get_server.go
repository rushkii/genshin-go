package utils

import (
	"strconv"
)

func GetServer(uid int) string {
	var res string

	server := map[string]string{
		"1": "cn_gf01",
		"2": "cn_gf01",
		"5": "cn_qd01",
		"6": "os_usa",
		"7": "os_euro",
		"8": "os_asia",
		"9": "os_cht",
	}

	for k, v := range server {
		if k == strconv.Itoa(uid)[0:1] {
			res = v
		}
	}
	return res
}
