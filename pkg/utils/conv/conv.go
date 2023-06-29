package conv

import (
	"fmt"
	"strconv"
)

func ConvStrToUint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		panic(fmt.Sprintf("cannot convert string %s to uint", str))
	}
	return uint(num)
}
