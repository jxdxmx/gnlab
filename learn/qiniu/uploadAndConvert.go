package main

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"strings"
)

// 对七牛云上的已有文件进行处理
// 可以同时批量执行多种处理，比如同时压缩成多种大小的图片并保存在七牛云上
func uploadConvert() {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone=&storage.ZoneHuabei
	operationManager := storage.NewOperationManager(mac, &cfg)

	bucket := "gnlab"
	key := "test.webp"
	saveBucket := bucket
	// 处理指令集合
	// 转成webp格式
	// fopFormat := fmt.Sprintf("imageMogr2/format/webp|saveas/%s",
	//	storage.EncodedEntry(saveBucket, "test.webp"))

	// 缩放，w100/q/50
	fopScale100 := fmt.Sprintf("imageView2/2/w/100/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, "test_100.webp"))
	// 缩放，w170/q/50
	fopScale170 := fmt.Sprintf("imageView2/2/w/170/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, "test_170.webp"))
	// 缩放，w250/q/50
	fopScale250 := fmt.Sprintf("imageView2/2/w/250/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, "test_250.webp"))
	// 缩放，w375/q/50
	fopScale375 := fmt.Sprintf("imageView2/2/w/375/q/50|saveas/%s",
		storage.EncodedEntry(saveBucket, "test_375.webp"))

	fopBatch := []string{fopScale100, fopScale170, fopScale250, fopScale375}
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
}
