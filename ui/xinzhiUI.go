package ui

import "fmt"

var xinzhiMap = map[string]WeatherCode{
	"clear-day":           CodeSunny,
	"clear-night":         CodeSunny,
	"rain":                CodeLightRain,
	"snow":                CodeLightSnow,
	"sleet":               CodeLightSleet,
	"wind":                CodePartlyCloudy,
	"fog":                 CodeFog,
	"cloudy":              CodeCloudy,
	"partly-cloudy-day":   CodePartlyCloudy,
	"partly-cloudy-night": CodePartlyCloudy,
	"thunderstorm":        CodeThunderyShowers,
}

func (ui *XinZhi)Echo(){
	for i, _ := range ui.Results {
		fmt.Println(ui.Results[i].Location.Name)
		for _, val := range ui.Results[i].Daily {
			fmt.Printf("%+v\n", val)
		}
		fmt.Println(ui.Results[i].LastUpdate)
	}
}