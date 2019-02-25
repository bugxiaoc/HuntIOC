package tools

import (
	"log"
	"os"
	"time"
)

var (
	LogInfo, LogErr *log.Logger
)

func init() {
	CheckDir("./log/")
	CheckDir("./src/")
	CheckDir("./src/json")
	CheckDir("./src/csv")
	file := "./log/" + time.Now().Format("20060102") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	LogInfo = log.New(logFile, "[INFO] -- ", log.Ldate|log.Ltime|log.Lshortfile)
	LogErr = log.New(logFile, "[Error] -- ", log.Ldate|log.Ltime|log.Lshortfile)
	LogInfo.Print("Log Init Success.")
}
