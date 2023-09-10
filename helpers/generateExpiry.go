package helpers

import "time"

func GenerateExpiry() time.Time{
	return time.Now().Add(10 * time.Minute)
}