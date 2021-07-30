package main

import (
	"net/http"

	"github.com/joho/godotenv"

	"github.com/EugeneTseitlin/dash-go-code-challange/internal/p2p/server"
	"github.com/EugeneTseitlin/dash-go-code-challange/internal/p2p/util"
)

func main() {
	var err error
	err = godotenv.Load()
	util.PanicError(err)

	router := server.CreateRouter()
	err = http.ListenAndServe(":8090", router)
	util.LogError(err)
}
