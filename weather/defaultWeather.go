package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"weather/ui"
)

var deurl = "http://wthrcdn.etouch.cn/weather_mini?citykey=101280601"

func DefaultWeather() {
	rsp, err := http.Get(deurl)
	if err != nil {
		log.Fatalln(err)
	}

	if rsp.StatusCode != 200{
		fmt.Println("can't get any weather info!! please try again later.")
	}

	msg := rsp.Body
	by, _ := ioutil.ReadAll(msg)
	var out *ui.Default
	err = json.Unmarshal(by, &out)
	if err != nil {
		panic(err)
	}

	out.Echo()
}
