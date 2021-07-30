package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/EugeneTseitlin/dash-go-code-challenge/internal/p2p/server"
	"github.com/EugeneTseitlin/dash-go-code-challenge/internal/p2p/util"
)

func main() {
	var err error
	err = godotenv.Load()
	util.LogError(err)

	router := server.CreateRouter()
	err = http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), router)
	util.LogError(err)
}
