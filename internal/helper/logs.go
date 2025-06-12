package helper

import (
	"fmt"
	"strconv"
	"time"
)

func Log(m string, id int, path string) {
	t := time.Now()
	var ty string

	if id == 0 {
		ty = "[System]"
	} else {
		ty = "[User : " + strconv.Itoa(id) + "]"
	}

	log := fmt.Sprintf("[%s] %s %s %s", t.Format("02 January 2006, 15:04:05"), ty, m, path)

	fmt.Println(log)
}

func LogWorker(m string) {
	t := time.Now()
	log := fmt.Sprintf("[%s] %s", t.Format("02 January 2006, 15:04:05"), m)

	fmt.Println(log)
}
