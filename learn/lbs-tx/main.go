package main

import "fmt"

const (
	apiServer = "https://apis.map.qq.com"
	apiKey    = "GTYBZ-CFVRD-HLJ4X-HH673-W63GK-KWB75"
	apiSk     = "SWBaAG1xsVAICn5nVFXeuGHSFiRo7oF2"
)

func main() {
	// 这个接口调用可能会失败【500.接口调用超时】，必须要有重试机制！！！
	api := "/ws/place/v1/search"
	params := map[string]string{
		"key": apiKey,
		//"keyword": "西湖", // 如果没有关键词，那么就是推荐了
		"boundary": "nearby(30.291606341928432,120.06976537001471,1000,1)",
		//"boundary":    "region(杭州市)",
		"page_index":  "1",
		"page_size":   "20",
		"get_subpois": "0",
		//"orderby":     "_distance", // 默认排序会综合考虑距离、权重等多方面因素
	}
	paramsString := GenerateParamsString(api, params)
	fmt.Println(apiServer + api + "?" + paramsString)

	//api2 := "/ws/geocoder/v1"
	//params2 := map[string]string{
	//	"location":    "30.29690,120.07528",
	//	"get_poi":     "1",
	//	"poi_options": "radius=1000;policy=4",
	//	"key":         apiKey,
	//}
	//paramsString2 := GenerateParamsString(api2, params2)
	//fmt.Println(apiServer + api2 + "?" + paramsString2)

	//// https://apis.map.qq.com/ws/place/v1/suggestion?key=GTYBZ-CFVRD-HLJ4X-HH673-W63GK-KWB75&keyword=%E8%A5%BF%E6%B9%96&page_index=1&page_size=20&sig=6d3e89dd32fb17e3c81690ea095a9a36
	//api := "/ws/place/v1/suggestion"
	//params := map[string]string{
	//	"key":        apiKey,
	//	"keyword":    "西湖",
	//	"page_index": "1",
	//	"page_size":  "20",
	//}
	//paramsString := GenerateParamsString(api, params)
	//fmt.Println(apiServer + api + "?" + paramsString)

	//// https://apis.map.qq.com/ws/place/v1/explore?boundary=nearby(30.291606341928432,120.06976537001471,1000,1)&key=GTYBZ-CFVRD-HLJ4X-HH673-W63GK-KWB75&page_index=1&page_size=20&sig=61eb6aecc637f2afaf8f5e1df32490cb
	//api := "/ws/place/v1/explore"
	//params := map[string]string{
	//	"key":        apiKey,
	//	"boundary":   "nearby(30.291606341928432,120.06976537001471,1000,1)",
	//	"page_index": "1",
	//	"page_size":  "20",
	//}
	//paramsString := GenerateParamsString(api, params)
	//fmt.Println(apiServer + api + "?" + paramsString)
}
