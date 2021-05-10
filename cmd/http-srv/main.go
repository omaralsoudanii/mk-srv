package main

import (
	"http-srv/v1/zlog"
	"os"
	"time"
)

func main() {
	_ = os.Setenv("TZ", "UTC")
	_, _ = time.LoadLocation("UTC")

	zlog.Info("Application starting....", zlog.String("timezone", time.Now().Location().String()))
}
