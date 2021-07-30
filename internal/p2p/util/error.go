package util

import "log"

func LogError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func PanicError(err error) {
	if err != nil {
		log.Panicf(err.Error())
	}
}