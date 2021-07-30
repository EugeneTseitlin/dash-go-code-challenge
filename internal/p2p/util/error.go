package util

import "log"

func LogError(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}

func PanicError(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
