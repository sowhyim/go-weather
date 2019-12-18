// 这是用于心知天气API的天气获取
// 本项目暂只支持心知天气，其他平台暂于TODO中
package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather/ui"
)

var xinzhiurl = "https://api.seniverse.com/v3/weather/daily.json?key=%v&location=shenzhen&language=zh-Hans&unit=c&start=-1&days=4"

// 这个私人key，做测试使用
var xinzhikey = "SyWfXUWAGN-8I2y6q"

func XinZhiWeather(key string) error {
	rsp, err := http.Get(fmt.Sprintf(xinzhiurl, key))
	if err != nil {
		return err
	}
	if rsp.StatusCode != 200 {
		return errors.New("err request")
	}

	msg := rsp.Body
	by, _ := ioutil.ReadAll(msg)
	var out *ui.XinZhi
	err = json.Unmarshal(by, &out)
	if err != nil {
		return err
	}
	out.Echo()
	return nil
}
