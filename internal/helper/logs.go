package helper

import (
	"fmt"
	"time"
)

func Log(m string, path string) {
	t := time.Now()
	log := fmt.Sprintf("[%s] %s %s", t.Format("2006-01-02 15:04:05"), m, path)

	fmt.Println(log)
}

func LogWorker(m string) {
	t := time.Now()
	log := fmt.Sprintf("[%s] %s", t.Format("2006-01-02 15:04:05"), m)

	fmt.Println(log)
}
