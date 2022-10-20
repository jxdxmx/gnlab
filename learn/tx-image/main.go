package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	width, height = 640, 427
	//url           = "https://duohao-prod-1306306153.file.myqcloud.com/aW7Jv2bd7B/tmp_1b1f5740269f487fd8f8ccd4047e45c0.jpg?imageView2/1/w/%d/h/%d/q/%d"
	//url = "https://duohao-test-1306306153.file.myqcloud.com/wMpbZL6Wp9/ErXz6z5naLto7b34f8ca1dbc692b67b59fd4f3dbf468.jpg?%d%d%d"

	//url           = "https://duohao-prod-1306306153.file.myqcloud.com/aW7Jv2bd7B/tmp_1b1f5740269f487fd8f8ccd4047e45c0.jpg?imageView2/1/w/%d/h/%d/q/%d"

	//url = "https://dn-odum9helk.qbox.me/resource/gogopher.jpg?imageView2/1/w/%d/h/%d/q/%d"
	url = "https://duohao-prod-1306306153.file.myqcloud.com/wMpbZL6Wp9/NnChjdLo0C9K6085d57b5ee06ac7f793fa5f0f052cb5.jpg?imageView2/1/w/%d/h/%d/q/%d"
)

func main() {
	startTime := time.Now()
	for j := 1; j <= 600; j++ {
		fmt.Printf("=================== %s 第 %d 轮,开始时间 %s  ===================\n", time.Now().Format(time.RFC3339), j, startTime.Format(time.RFC3339))
		var wg sync.WaitGroup
		threads := 10
		wg.Add(threads)
		for i := 1; i <= threads; i++ {
			go accessImage(&wg, i, 30+i*2)
		}
		wg.Wait()
		time.Sleep(time.Second * 1)
	}
}

func accessImage(wg *sync.WaitGroup, seq, percent int) {
	defer func() { wg.Done() }()
	w, h := width*percent/100, height*percent/100
	newUrl := fmt.Sprintf(url, w, h, percent)
	client := &http.Client{}
	request, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		fmt.Println(seq, "request error ", err.Error())
	}
	t1 := time.Now()
	resp, err := client.Do(request) //发送请求
	if err != nil {
		fmt.Println(seq, "request error ", err.Error())
	}
	defer func() { _ = resp.Body.Close() }() //一定要关闭resp.Body
	defer func() {
		fmt.Printf("%d,缩放比例: %d,耗时: %v \n", seq, percent, time.Since(t1))
	}()
}
