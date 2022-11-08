package main

import (
	"crypto/md5"
	"encoding/hex"
	"git.gnlab.com/duohao/backend_content.git/conf"
	"net/url"
	"sort"
)

var (
	// 需要编码的参数
	needEncodeKeys = map[string]bool{"keyword": true}
)

func ParamsSort(params map[string]string) (string, string) {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paramsString string
	var paramsStringSign string // 原始参数，不进行urlEncode，后续生成签名时用
	for i, v := range keys {
		if i > 0 {
			paramsString += "&"
			paramsStringSign += "&"
		}
		paramsStringSign += v + "=" + params[v]
		if needEncodeKeys[v] {
			paramsString += v + "=" + url.QueryEscape(params[v])
		} else {
			paramsString += v + "=" + params[v]
		}
	}
	return paramsString, paramsStringSign
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sig(apiURI string, paramsString string) string {
	sig := apiURI + "?" + paramsString + conf.LBS_TENCENT_API_SK
	return MD5(sig)
}

func GenerateParamsString(apiURI string, params map[string]string) string {
	paramsString, paramsStringSign := ParamsSort(params)
	sig := Sig(apiURI, paramsStringSign)
	return paramsString + "&sig=" + sig
}
