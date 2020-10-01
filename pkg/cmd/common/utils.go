package common

import (
	"github.com/Benbentwo/utils/log"
	"strconv"
)

func MustStrConvToStr(num int) string {
	s := strconv.Itoa(num)
	return s
}

func MustStrConvToInt(str string) int {
	s, err := strconv.Atoi(str)
	if err != nil {
		log.Logger().Fatalf("Couldnt' convert %s to an int", str)
	}
	return s
}
