package tools

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

type User struct {
	UserName string `json:username`
	Password string `json:password`
	Md5      string `-`
	Ip       string `-`
}
type Platform struct {
	Name string `json:name`
	Val  string `json:val`
}
type Users struct {
	User     []User     `json:"users"`
	Platform []Platform `json: "platform"`
}

var Config *Users

func init() {
	if Config == nil {
		LoadConfiguration("./conf.json")
	}
}
func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), &Config)
	if err != nil {
		return err
	}
	return err
}
func GetPlatform(pl string) string {
	for _, v := range Config.Platform {
		if v.Name == pl {
			return v.Val
		}
	}
	return ""
}

func GetMd5ByName(name string, pass string) string {
	for k, v := range Config.User {
		if (v.UserName == name) && (v.Password == pass) {
			times := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			Config.User[k].Md5 = Md5(name + pass + times)
			Config.User[k].Ip = "127.0.0.1"
			return Config.User[k].Md5
		}
	}
	return ""
}

func CheckSession(uname interface{}, md5 interface{}) bool {
	for _, v := range Config.User {
		if v.UserName == uname && v.Md5 == md5 {
			return true
		}
	}
	return false
}

func SetFiled(name string, data string) bool{
	for k, v := range Config.Platform {
		if v.Name == name {
			Config.Platform[k].Val = data
			return true
		}
	}
	return false

}
