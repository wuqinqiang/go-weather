package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wuqinqiang/go-weather/entity"
)

const (
	 Key = "xx"
	Uri = "https://restapi.amap.com/v3/weather/weatherInfo"
)

func GetWeather(code int) (map[string]interface{}, error) {
	info, err := GetWeatherRequest(code)
	if err != nil {
		return nil, err
	}
	infoMap := make(map[string]interface{})
	infoMap["城市:"] = info.Lives[0].City
	infoMap["天气现象:"] = info.Lives[0].Weather
	infoMap["实时气温:"] = info.Lives[0].Temperature
	infoMap["数据发布:"] = info.Lives[0].Reporttime
	return infoMap, nil
}

func GetWeatherRequest(code int) (entity.ResponseInfo, error) {
	info := entity.ResponseInfo{}
	client := &http.Client{Timeout: 2 * time.Second}
	url := fmt.Sprintf(Uri+"?key=%s&city=%d", Key, code)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("查询错误:", err)
		return info, errors.New("查询失败")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return info, errors.New("查询失败")
	}
	if err := json.Unmarshal(body, &info); err != nil {
		fmt.Println("unmarshal response err:", err)
		return info, errors.New("查询失败,请输入正确的城市码")
	}

	if info.Status != "1" {
		return info, errors.New("api 密钥key错误，请检查")
	}
	return info, nil
}
