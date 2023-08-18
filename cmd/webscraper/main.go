package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/eviccari/console-stores-price-comparator/configs"
)

const (
	WithErrorStatus   = 1
	WithSuccessStatus = 0
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	configs.Load()
	xboxURI := configs.Params.XboxStore.StoreBaseURI
	res, err := http.Get(xboxURI)
	if err != nil {
		logger.Error(err.Error())
		finish(WithErrorStatus)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(err.Error())
		finish(WithErrorStatus)
	}
	fmt.Println(string(body))
	finish(WithSuccessStatus)
}

func finish(statusCode int) {
	os.Exit(statusCode)
}
