package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"mytool/02-craw/t66y/cmd"
	"os"
)

func init() {
	log.SetReportCaller(false)
	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.0000",
	})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	flag.Parse()
}

func main() {
	cmd.Execute()
}
