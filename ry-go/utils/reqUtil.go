package utils

import (
	"log"
	"net/http"
	"sync"

	"github.com/imroc/req/v3"
)

var (
	client     *req.Client
	clientOnce sync.Once
)

func getClient() *req.Client {
	clientOnce.Do(func() {
		client = req.C()
		client.EnableDebugLog()
		client.EnableDumpAll()
	})
	return client
}

// ReqGet req的GET请求
//
// 参数
//
//	-reqUrl: 请求URL
//	-queryParam: 请求参数
//	-headerMap: 请求头参数
//
// 返回
//
//	-string: 成功的返回结果
//	-error: 错误
//
//goland:noinspection GoUnusedExportedFunction
func ReqGet(reqUrl string, queryParam, headerMap map[string]string) (string, error) {
	resp, err := getClient().R().
		SetQueryParams(queryParam).
		SetHeaders(headerMap).
		EnableDump().
		Get(reqUrl)
	if err != nil {
		log.Printf("ReqGet Error ==%+v", err)
		return "", err
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf("ReqGet Response ===%+v", resp)
		return resp.String(), nil
	}
	return "", nil
}
