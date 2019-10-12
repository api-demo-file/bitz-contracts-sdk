package service

import (
	"net/url"
	"io/ioutil"
	"strings"
	"net/http"
	"sort"
	"crypto/md5"
	"encoding/hex"
)

// Http Get请求基础函数, 通过封装Go语言Http请求, 支持火币网REST API的HTTP Get请求
// strUrl: 请求的URL
// strParams: string类型的请求参数, user=lxz&pwd=lxz
// return: 请求结果
func HttpGetRequest(strUrl string, mapParams map[string]string) string {
	httpClient := &http.Client{}

	var strRequestUrl string
	if nil == mapParams {
		strRequestUrl = strUrl
	} else {
		strParams := Map2UrlQuery(mapParams)
		strRequestUrl = strUrl + "?" + strParams
	}

	// 构建Request, 并且按官方要求添加Http Header
	request, err := http.NewRequest("GET", strRequestUrl, nil)
	if nil != err {
		return err.Error()
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	// 发出请求
	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if nil != err {
		return err.Error()
	}

	// 解析响应内容
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

// Http POST请求基础函数, 通过封装Go语言Http请求, 支持火币网REST API的HTTP POST请求
// strUrl: 请求的URL
// mapParams: map类型的请求参数
// return: 请求结果
func HttpPostRequest(strUrl string, mapParams map[string]string) string {
	httpClient := &http.Client{}

	var queryStr string

	values := url.Values{}

	for key, val := range mapParams {
		values.Add(key, val)
	}
	queryStr = values.Encode()

	request, err := http.NewRequest("POST", strUrl, strings.NewReader(queryStr))

	if nil != err {
		return err.Error()
	}
	request.Header.Add("User-Agent", "'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if nil != err {
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

// 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]string) string {
	var strParams string
	values := url.Values{}

	for key, value := range mapParams {
		values.Add(key, value)
	}
	strParams = values.Encode()

	return strParams
}

func BitzSign(mapParams map[string]string, secretKey string) string{
	var keys []string
	for key := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var queryStr string
	for _,k := range keys {
		queryStr += k + "=" +mapParams[k] + "&"
	}

	queryStr = queryStr[:len(queryStr) -1]

	queryStr += secretKey

	h := md5.New()
	h.Write([]byte(queryStr))
	sign := hex.EncodeToString(h.Sum(nil))
	return sign
}