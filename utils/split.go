package utils

import (
	"strconv"
	"strings"
)

func GetStartAndEndOfRange(val string) (int, int) {
	idRange := strings.Split(val, "-")

	start, err := strconv.Atoi(idRange[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(idRange[1])
	if err != nil {
		panic(err)
	}

	return start, end

}
