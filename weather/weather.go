// 这是用于心知天气API的天气获取
// 本项目暂只支持心知天气，其他平台暂于TODO中
package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://api.seniverse.com/v3/weather/daily.json?key=%v&location=shenzhen&language=zh-Hans&unit=c&start=0&days=3"
// 这个私人key，做测试使用
var Key = "SyWfXUWAGN-8I2y6q"

func Weather(key string){
	fmt.Println(fmt.Sprintf(url, key))
	rsp, err := http.Get(fmt.Sprintf(url, key))
	if err != nil {
		panic(err)
	}
	msg := rsp.Body
	by, _ := ioutil.ReadAll(msg)
	var out *Rsp
	err = json.Unmarshal(by, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out.Results[0].Location.Name)
	for _, val := range out.Results[0].Daily {
		fmt.Printf("%+v\n", val)
	}
	fmt.Println(out.Results[0].LastUpdate)
}

type Rsp struct {
	Results []*Results `json:"results"`
}

type Results struct {
	Daily      []*Daily  `json:"daily"`
	Location   *Location `json:"location"`
	LastUpdate string    `json:"last_update"`
}

type Location struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	Path           string `json:"path"`
	Timezone       string `json:"timezone"`
	TimezoneOffset string `json:"timezone_offset"`
}

type Daily struct {
	Date                string `json:"date"`
	TextDay             string `json:"text_day"`
	CodeDay             string `json:"code_day"`
	TextNight           string `json:"text_night"`
	CodeNight           string `json:"code_night"`
	High                string `json:"high"`
	Low                 string `json:"low"`
	Precip              string `json:"precip"`
	WindDirection       string `json:"wind_direction"`
	WindDirectionDegree string `json:"wind_direction_degree"`
	WindSpeed           string `json:"wind_speed"`
	WindScale           string `json:"wind_scale"`
}
