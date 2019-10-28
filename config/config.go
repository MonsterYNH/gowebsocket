package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	ENV_HEART_CHECK = 10
)

func init() {
	if check := os.Getenv("ENV_HEART_CHECK"); len(check) != 0 {
		checkNum, err := strconv.Atoi(check)
		if err != nil {
			panic(fmt.Sprintf("ERROR: set heart check environment failed, error: %s", err))
		}
		ENV_HEART_CHECK = checkNum
	}
}
