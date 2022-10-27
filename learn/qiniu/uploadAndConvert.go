package main

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"strings"
	"time"
)

// 对七牛云上的已有文件进行处理
// 可以同时批量执行多种处理，比如同时压缩成多种大小的图片并保存在七牛云上
func process(key string) (err error) {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone=&storage.ZoneHuabei
	operationManager := storage.NewOperationManager(mac, &cfg)

	//key := "test.webp"
	name := strings.Split(key, ".")[0]
	saveBucket := bucket
	// 处理指令集合
	// 转成webp格式 ===  格式转换没法和缩放一起执行！！！！
	//fopFormat := fmt.Sprintf("imageView2/0/format/webp|saveas/%s",
	//	base64.URLEncoding.EncodeToString([]byte(bucket+":"+fmt.Sprintf("%s.webp", name))))
	//fopFormat := fmt.Sprintf("imageMogr2/format/webp|saveas/%s",
	//	storage.EncodedEntry(saveBucket, fmt.Sprintf("%s.webp", name)))

	// 缩放，w100/q/50
	fopScale100 := fmt.Sprintf("imageView2/2/w/100/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, fmt.Sprintf("%s_100.webp", name)))
	// 缩放，w170/q/50
	fopScale170 := fmt.Sprintf("imageView2/2/w/170/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, fmt.Sprintf("%s_170.webp", name)))
	// 缩放，w250/q/50
	fopScale250 := fmt.Sprintf("imageView2/2/w/250/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, fmt.Sprintf("%s_250.webp", name)))
	// 缩放，w375/q/50
	fopScale375 := fmt.Sprintf("imageView2/2/w/375/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, fmt.Sprintf("%s_375.webp", name)))
	// 缩放，w:200 h:160/50
	fopScale200160 := fmt.Sprintf("imageView2/1/w/200/h/160/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, fmt.Sprintf("%s_200_160.webp", name)))

	fopBatch := []string{fopScale100, fopScale170, fopScale250, fopScale375, fopScale200160}
	//fopBatch := []string{fopScale100}
	fops := strings.Join(fopBatch, ";")
	// 强制重新执行数据处理任务
	force := true
	// 数据处理指令全部完成之后，通知该地址
	notifyURL := ""
	// 数据处理的私有队列，必须指定以保障处理速度
	pipeline := "jxdxmx_img_queue"
	persistentId, err := operationManager.Pfop(bucket, key, fops, pipeline, notifyURL, force)
	for i := 0; err != nil && i < 3; i++ {
		persistentId, err = operationManager.Pfop(bucket, key, fops, pipeline, notifyURL, force)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("process qiniu's image success. persistentId:%s \n", persistentId)

	// 使用持久化ID查询状态
	// http://api.qiniu.com/status/get/prefop?id=z0.01z012cntx7vfwwdbp00mv9wxk0005vd
	// 状态码：0成功，1等待处理，2正在处理，3处理失败。
	// persistentId := "z0.597f28b445a2650c994bb208"
	start := time.Now()
	for {
		ret, err := operationManager.Prefop(persistentId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var m = map[int]string{0: "成功", 1: "等待处理", 2: "正在处理", 3: "处理失败"}
		//log.Println("处理状态:", m[ret.Items[0].Code])
		log.Println("处理状态:", m[ret.Code])
		if ret.Code == 0 {
			fmt.Println("已成功", time.Since(start))
			break
		}
		//fmt.Println(ret.String())
	}
	return
}
