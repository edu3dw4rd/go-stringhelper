package stringhelper

import (
	"strconv"
	"strings"
)

// StringToInt64Array converts string separated by ',' to []int64
func StringToInt64Array(input string) (nums []int64, err error) {
	splitStr := strings.Split(input, ",")

	for _, str := range splitStr {
		num, err := strconv.ParseInt(string(str), 10, 64)

		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return
}
