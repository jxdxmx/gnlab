package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"net/http"
)

var (
	accessKey = "k2QkKuqP-4BVZhXH66W8_QygWvz1pL9bPY82tii-"
	secretKey = "ujsZRKhnElqkKIgxtFRrTTQQ2GJdvpULQA1QeVoM"
	mac       = qbox.NewMac(accessKey, secretKey)
)

func accessImage(url string) (bs []byte, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(url, "access network image error.", err.Error())
		return
	}
	resp, err := client.Do(request) //发送请求
	if err != nil {
		fmt.Println(url, "access network image error.", err.Error())
	}
	defer func() { _ = resp.Body.Close() }() //一定要关闭resp.Body
	bs = make([]byte, 0, resp.ContentLength)
	bs, err = ioutil.ReadAll(resp.Body)
	return
}

// 上传腾讯云上面现有的图片到七牛云
func uploadTXImg(url string, filename string) (err error) {
	bucket := "gnlab"
	//putPolicy := storage.PutPolicy{
	//	Scope: bucket,
	//}
	//key := "test3.webp"
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, filename),
	}
	putPolicy.Expires = 7200 // 2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	bs, err := accessImage(url)
	for i := 0; err != nil && i < 3; i++ {
		fmt.Println(i+1, "retry access image:", url)
		bs, err = accessImage(url)
	}
	if err != nil {
		fmt.Println("无法获取图片")
		return
	}

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			//"x:name": "github logo",
		},
	}
	dataLen := int64(len(bs))

	err = formUploader.Put(context.Background(), &ret, upToken, filename, bytes.NewReader(bs), dataLen, &putExtra)
	for i := 0; err != nil && i < 3; i++ {
		err = formUploader.Put(context.Background(), &ret, upToken, filename, bytes.NewReader(bs), dataLen, &putExtra)
	}
	if err != nil {
		fmt.Println(url, "uploadTXImg error", err)
		return
	}

	fmt.Printf("load network image from tx to qiniu success. %+v\n", ret)
	return
}

func uploadLocalImg() {
	bucket := "gnlab"
	//putPolicy := storage.PutPolicy{
	//	Scope: bucket,
	//}
	key := "test3.webp"
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, key),
	}
	putPolicy.Expires = 7200 // 2小时有效期

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	localFile := "D:\\01.工作目录\\20221020-图片加载速度慢\\test.webp"

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	for i := 0; err != nil && i < 3; i++ {
		err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// http://image1.test.hrai.online/自救设备.jpg?imageView2/1/w/200/h/200
	fmt.Println("http://image1.test.hrai.online/"+ret.Key, ret.Hash)
}
