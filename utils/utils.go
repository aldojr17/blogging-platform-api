package utils

import (
	log "blogging-platform-api/logger"
	"os"
	"strconv"
	"time"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func ConvertToInteger(value string) int {
	converted, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	return converted
}

func GenerateCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func ConvertTimestampToFormattedDate(timestamp int64) string {
	t := time.Unix(timestamp, 0).UTC()

	return t.Format(time.RFC3339)
}
