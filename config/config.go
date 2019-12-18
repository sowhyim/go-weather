package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func GetConfig() string{
	user, _ := user.Current()
	bmsg, err := ioutil.ReadFile(user.HomeDir + "/.weatherc")
	if err != nil {
		log.Fatalln(err)
	}
	return string(bmsg)
}

func SetConfig(key string) {
	user, _ := user.Current()
	file, err := os.OpenFile(user.HomeDir+"/.weatherc", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	file.WriteString(key)
}

func Check()bool {
	user, _ := user.Current()
	bmsg, err := ioutil.ReadFile(user.HomeDir + "/.weatherc")
	if err != nil {
		log.Fatalln(err)
	}
	if string(bmsg) == ""{
		return false
	}
	return true
}