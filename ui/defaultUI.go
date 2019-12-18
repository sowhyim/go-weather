package ui

import (
	"fmt"
)

var defaultMap = map[string]WeatherCode{
	"晴":    CodeSunny,
	"小雨":   CodeLightRain,
	"小雪":   CodeLightSnow,
	"雨夹雪":  CodeLightSleet,
	"雾":    CodeFog,
	"多云":   CodeCloudy,
	"局部多云": CodePartlyCloudy,
	"雷雨":   CodeThunderyShowers,
}

func (ui *Default) Echo() {
	if ui.Status != 1000 {
		return
	}

	fmt.Println(ui.Data.City, "天气:")
	echo(ui.Data.Yesterday, nil)
	for _, forcast := range ui.Data.Forecast {
		echo(nil, forcast)
	}
}

func echo(yesterday *Yesterday, forecast *Forecast) {
	if yesterday != nil {
		fmt.Printf("日期：%s , 天气：%s ", yesterday.Date, yesterday.Type)
		fmt.Print(codes[defaultMap[yesterday.Type]])
		fmt.Printf(" , 最高温：%s , 最低温：%s\n", yesterday.High, yesterday.Low)
	} else {
		fmt.Printf("日期：%s , 天气：%s ", forecast.Date, forecast.Type)
		fmt.Print(codes[defaultMap[forecast.Type]])
		fmt.Printf(" , 最高温：%s , 最低温：%s\n", forecast.High, forecast.Low)
	}
}
