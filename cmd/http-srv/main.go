package main

import (
	"os"
	"time"

	"github.com/omaralsoudanii/http-srv/zlog"
)

func main() {
	_ = os.Setenv("TZ", "UTC")
	_, _ = time.LoadLocation("UTC")

	zlog.Info("Application starting....")
	zlog.Info(time.Now().Format("01-02-2006 15:04:05"))
}
