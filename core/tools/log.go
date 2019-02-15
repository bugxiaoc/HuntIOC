package tools

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogInfo *log.Logger
	LogErr  *log.Logger
)

func init() {
	checkDir()
	file := "./log/" + time.Now().Format("20060102") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	LogInfo = log.New(logFile, "[INFO] -- ", log.Ldate|log.Ltime|log.Lshortfile)
	LogErr = log.New(logFile, "[Error] -- ", log.Ldate|log.Ltime|log.Lshortfile)
	LogInfo.Print("Log Init Success.")
}

func checkDir() {
	path := "./log/"
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("Open Dir Err ,msg [%v]\n", err)
	}
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}
