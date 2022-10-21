package main

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	accessKey = "k2QkKuqP-4BVZhXH66W8_QygWvz1pL9bPY82tii-"
	secretKey = "ujsZRKhnElqkKIgxtFRrTTQQ2GJdvpULQA1QeVoM"
	mac       = qbox.NewMac(accessKey, secretKey)
)

func upload() {
	bucket := "gnlab"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 // 2小时有效期

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	localFile := "D:\\01.工作目录\\20221020-图片加载速度慢\\自救设备.jpg"

	key := "自救设备.jpg"
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
	if err != nil {
		fmt.Println(err)
		return
	}
	// http://rk1x9pnvh.hd-bkt.clouddn.com/自救设备.jpg?imageView2/1/w/200/h/200
	fmt.Println(ret.Key, ret.Hash)

}
