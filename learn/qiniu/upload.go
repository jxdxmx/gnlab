package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	accessKey = "k2QkKuqP-4BVZhXH66W8_QygWvz1pL9bPY82tii-"
	secretKey = "ujsZRKhnElqkKIgxtFRrTTQQ2GJdvpULQA1QeVoM"
	mac       = qbox.NewMac(accessKey, secretKey)
	bucket    = "gnlab"
	domain    = "http://image1.test.hrai.online/"
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
func uploadNetImg(url string, key string) (err error) {
	//putPolicy := storage.PutPolicy{
	//	Scope: bucket,
	//}
	//key := "test3.webp"

	//设置转码后保留的文件名称，传入相同的key，表示转码并覆盖，空间只保留转码后的文件名称
	saveEntry := base64.URLEncoding.EncodeToString([]byte(bucket + ":" + key))
	fops := "imageView2/0/format/webp|saveas/" + saveEntry
	putPolicy := storage.PutPolicy{
		Scope:         bucket,
		PersistentOps: fops,
	}
	putPolicy.Expires = 7200 // 2小时有效期
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

	//formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	rPutExtra := storage.RputExtra{
		Params: map[string]string{
			//"x:name": "github logo",
		},
	}
	dataLen := int64(len(bs))

	resumeUploader := storage.NewResumeUploaderEx(&cfg, &storage.Client{Client: &http.Client{}})
	err = resumeUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(bs), dataLen, &rPutExtra)
	for i := 0; err != nil && i < 3; i++ {
		err = resumeUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(bs), dataLen, &rPutExtra)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	//err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(bs), dataLen, &putExtra)
	//for i := 0; err != nil && i < 3; i++ {
	//	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(bs), dataLen, &putExtra)
	//}
	//if err != nil {
	//	fmt.Println(url, "uploadNetImg error", err)
	//	return
	//}

	//fmt.Printf("load network image from tx to qiniu success. %+v\n", ret)
	fmt.Printf("load network image from tx to qiniu success. %s\n", url)
	return
}

func uploadLocalImg() {
	//putPolicy := storage.PutPolicy{
	//	Scope: bucket,
	//}
	key := "test3.webp"
	putPolicy := storage.PutPolicy{
		//Scope: fmt.Sprintf("%s:%s", bucket, key),
		Scope: fmt.Sprintf("%s", bucket),
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
	fmt.Println(domain+ret.Key, ret.Hash)
}

// 缩放是可以，但是转换格式失败，还是jpg格式。。
func uploadAndProcess() {
	saveBucket := bucket
	//putPolicy := storage.PutPolicy{
	//	Scope: bucket,
	//}

	localFile := "D:\\01.工作目录\\20221020-图片加载速度慢\\diqiu.jpg"

	name := "diqiu2"
	key := fmt.Sprintf("%s.jpg", name)

	// 缩放，w100/q/50
	fopFormat := fmt.Sprintf("imageMogr2/format/webp|saveas/%s",
		storage.EncodedEntry(saveBucket, fmt.Sprintf("%s.webp", name)))
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

	fopBatch := []string{fopFormat, fopScale100, fopScale170, fopScale250, fopScale375}
	//fopBatch := []string{fopScale100}
	fops := strings.Join(fopBatch, ";")

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, key),
		//Scope:         fmt.Sprintf("%s", bucket),
		PersistentOps:      fops,
		PersistentPipeline: "jxdxmx_img_queue",
	}
	putPolicy.Expires = 7200 // 2小时有效期

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

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

func DeleteQiNiuFile(key string) (err error) {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	err = bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func DeleteQiNiuFileByTTL(key string, days int) (err error) {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	err = bucketManager.DeleteAfterDays(bucket, key, days)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
