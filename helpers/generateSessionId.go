package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateSessionId() string {
	rand.Seed(time.Now().Unix())
	min := 1000000000
	max := 9999999999
	return fmt.Sprintf("%v", rand.Intn(max-min+1)+min)
}
