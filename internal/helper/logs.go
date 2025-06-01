package helper

import (
	"fmt"
	"time"
)

func Log(m string, path string) string {
	t := time.Now()
	log := fmt.Sprintf("[%s] %s %s", t.Format("2006-01-02 15:04:05"), m, path)

	return log
}

func LogWorker(m string) string {
	t := time.Now()
	log := fmt.Sprintf("[%s] %s", t.Format("2006-01-02 15:04:05"), m)

	return log
}
