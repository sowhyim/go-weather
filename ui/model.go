package ui

type WeatherCode int

const (
	CodeUnknown WeatherCode = iota
	CodeCloudy
	CodeFog
	CodeHeavyRain
	CodeHeavyShowers
	CodeHeavySnow
	CodeHeavySnowShowers
	CodeLightRain
	CodeLightShowers
	CodeLightSleet
	CodeLightSleetShowers
	CodeLightSnow
	CodeLightSnowShowers
	CodePartlyCloudy
	CodeSunny
	CodeThunderyHeavyRain
	CodeThunderyShowers
	CodeThunderySnowShowers
	CodeVeryCloudy
)

var codes = map[WeatherCode]string{
	CodeUnknown:             "✨",
	CodeCloudy:              "☁️",
	CodeFog:                 "🌫",
	CodeHeavyRain:           "🌧",
	CodeHeavyShowers:        "🌧",
	CodeHeavySnow:           "❄️",
	CodeHeavySnowShowers:    "❄️",
	CodeLightRain:           "🌦",
	CodeLightShowers:        "🌦",
	CodeLightSleet:          "🌧",
	CodeLightSleetShowers:   "🌧",
	CodeLightSnow:           "🌨",
	CodeLightSnowShowers:    "🌨",
	CodePartlyCloudy:        "⛅️",
	CodeSunny:               "☀️",
	CodeThunderyHeavyRain:   "🌩",
	CodeThunderyShowers:     "⛈",
	CodeThunderySnowShowers: "⛈",
	CodeVeryCloudy:          "☁️",
}

type XinZhi struct {
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

type Default struct {
	Data   Data   `json:"data"`
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

type Data struct {
	City      string    `json:"city"`
	Yesterday *Yesterday `json:"yesterday"`
	Forecast   []*Forecast   `json:"forecast"`
	Ganmao    string    `json:"ganmao"`
	Wendu     string    `json:"wendu"`
}

type Yesterday struct {
	Date string `json:"date"`
	High string `json:"high"`
	Fx   string `json:"fx"`
	Low  string `json:"low"`
	Fl   string `json:"fl"`
	Type string `json:"type"`
}

type Forecast struct {
	Date      string `json:"date"`
	High      string `json:"high"`
	Fengli    string `json:"fengli"`
	Low       string `json:"low"`
	Fengxiang string `json:"fengxiang"`
	Type      string `json:"type"`
}