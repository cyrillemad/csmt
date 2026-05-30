package debug

import "log"

func SLog(message string) {
	if Status {
		log.Println(message)
	}
}
