package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/http"
	"strings"
)

func main() {
	//uploadNetImg("https://duohao-test-1306306153.file.myqcloud.com/wMpbZL6Wp9/tmp_b3b94258b29f12d11666e1bb14d493ef.jpg", "b3b94258b29f12d11666e1bb14d493ef.webp")
	//uploadLocalImg() // 上传本地图片
	//uploadConvert() // 对图片进行变换，一次可以执行多个任务 ==> 没法正确的转换格式

	//uploadAndProcess() // 上传本地图片
	//DeleteQiNiuFile() // 删一个七牛文件

	// 上传网上图片到七牛云
	url := "https://t7.baidu.com/it/u=2708276792,219516514&fm=193&f=GIF"
	key := "building19.webp"
	var err error
	if err = uploadNetImg(url, key); err != nil {
		fmt.Printf("**********抓取图片 %s 并上传到七牛失败\n", url)
	}
	name := strings.Split(key, ".")[0]
	fmt.Printf("%s%s.webp\n", domain, name)
	fmt.Printf("%s%s_100.webp\n", domain, name)
	fmt.Printf("%s%s_170.webp\n", domain, name)
	fmt.Printf("%s%s_250.webp\n", domain, name)
	fmt.Printf("%s%s_375.webp\n", domain, name)
	fmt.Printf("%s%s_200_160.webp\n", domain, name)
	if err = process(key); err != nil {
		fmt.Printf("************对七牛文件 %s%s 进行格式转换或缩放失败\n", domain, key)
	}
	////else {
	//name := strings.Split(key, ".")[0]
	//fmt.Printf("%s%s.webp\n", domain, name)
	//fmt.Printf("%s%s_100.webp\n", domain, name)
	//fmt.Printf("%s%s_170.webp\n", domain, name)
	//fmt.Printf("%s%s_250.webp\n", domain, name)
	//fmt.Printf("%s%s_375.webp\n", domain, name)
	//fmt.Printf("%s%s_200_160.webp\n", domain, name)

	//	var done = make(chan struct{})
	//	go func() {
	//		defer func() { done <- struct{}{} }()
	//		// 确保转码成功后再删除原文件
	//		var err error
	//		for i := 0; i < 60; i++ {
	//			time.Sleep(time.Second)
	//			var bs []byte
	//			if bs, err = accessImage(fmt.Sprintf("%s%s_375.webp", domain, name)); err == nil && !strings.Contains(string(bs), "Document not found") {
	//				break
	//			}
	//		}
	//		if err != nil {
	//			if err = DeleteQiNiuFileByTTL(key, 30); err != nil {
	//				fmt.Printf("************60s内未能完成转换格式和缩放的功能,老格式的七牛文件将在30天后自动删除，%s%s\n", domain, key)
	//			}
	//			return
	//		}
	//		time.Sleep(time.Second * 5)
	//		if err = DeleteQiNiuFile(key); err != nil {
	//			fmt.Printf("************删除老格式的七牛文件失败\n")
	//		} else {
	//			fmt.Printf("************原jpg文件已从七牛中删除，%s%s\n", domain, key)
	//		}
	//	}()
	//	<-done
	//}

	//sampleUp()
}

// 可以做到上传的同时、将图片转换格式为webp
func sampleUp() {
	localFile := "D:\\01.工作目录\\20221020-图片加载速度慢\\equipment-4.jpg"

	//空间保留的文件名称
	key := "equipment-4.webp"
	//设置转码后保留的文件名称，传入相同的key，表示转码并覆盖，空间只保留转码后的文件名称
	saveEntry := base64.URLEncoding.EncodeToString([]byte(bucket + ":" + key))

	fops := "imageView2/0/format/webp|saveas/" + saveEntry
	putPolicy := storage.PutPolicy{
		//Scope:         bucket,
		Scope:         fmt.Sprintf("%s:%s", bucket, key),
		PersistentOps: fops,
	}
	putPolicy.Expires = 7200 // 2小时有效期
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploaderEx(&cfg, &storage.Client{Client: &http.Client{}})
	ret := storage.PutRet{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
