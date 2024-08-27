package utils

import (
	"os"
	"time"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GenerateCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func ConvertTimestampToFormattedDate(timestamp int64) string {
	t := time.Unix(timestamp, 0).UTC()

	return t.Format(time.RFC3339)
}
