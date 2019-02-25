package tools

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 计算字符串的md5值
func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}

//
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		LogErr.Printf("Get Root Path Error: %v", err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func CheckDir(path string){
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err.Error())
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}
func OpenFile(path string, flag bool) *os.File{
	_, err := os.Stat(path)
	if err == nil {
		if flag{
			err = os.Remove(path)
			if err != nil{
				LogErr.Println("Remove File Err . Path: ", path)
			}
		}
	}
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error(), "name:", path)
		return nil
	}else {
		return fileObj
	}
}

func WriteCsv(path string, data []string) string {
	fileObj := OpenFile(path, false)
	defer fileObj.Close()
	fileObj.WriteString("\xEF\xBB\xBF")
	csvFile := csv.NewWriter(fileObj)
	csvFile.Write(data)
	csvFile.Flush()
	return path
}

func WriteFile(path string, data []byte) string {
	fileObj := OpenFile(path, true)
	defer fileObj.Close()
	fileObj.Write(data)
	return path
}