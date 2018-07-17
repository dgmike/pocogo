package routes

import (
	"log"
	"time"
)

// SyncLogs to log requests
func SyncLogs(before time.Time) {
	elapsedTime := time.Since(before) / 1000
	log.Printf("Elapsed: %dms", elapsedTime)
}
