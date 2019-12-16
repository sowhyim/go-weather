package parse

import (
	"log"
	"os"
	"strings"
	"weather/config"
	"weather/weather"
)

func Parse() {
	if len(os.Args) == 1 {
		log.Println("this is help")
	}

	switch strings.ToLower(os.Args[1]) {
	case "set":
		if len(os.Args) == 3 {
			config.SetConfig(os.Args[2])
		}
	case "gw":
		weather.Weather(config.GetConfig())
	}
}
