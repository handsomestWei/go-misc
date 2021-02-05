package transfor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpClient struct {
	c *http.Client
}

// 创建http客户端连接
func NewHttpClient(ReqTimeOut int64) *httpClient {
	return &httpClient{
		c: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					// tcp连接超时
					Timeout: time.Duration(30) * time.Second,
					// 为了维持http keepalive状态 每隔多长时间发送Keep-Alive报文
					KeepAlive: time.Duration(30) * time.Second,
				}).DialContext,
				// 连接池对所有host的最大连接数量
				MaxIdleConns: 300,
				// 连接池对每个host的最大连接数量
				MaxIdleConnsPerHost: 10,
				// 空闲timeout设置，也即socket在该时间内没有交互则自动关闭连接
				IdleConnTimeout: time.Duration(90) * time.Second,
			},
			// 客户端超时设置
			Timeout: time.Duration(ReqTimeOut) * time.Second,
		},
	}

}

// get
func (h *httpClient) DoGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// post form
func (h *httpClient) DoPostForm(remoteUrl string, parameters map[string]string) ([]byte, error) {
	paramMap := url.Values{}
	for k, v := range parameters {
		paramMap.Add(k, v)
	}

	request, err := http.NewRequest(http.MethodPost, remoteUrl, strings.NewReader(paramMap.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;utf-8")
	resp, err := h.c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// post json
func (h *httpClient) DoPostJson(remoteUrl string, parameters interface{}) ([]byte, error) {
	bytesData, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest(http.MethodPost, remoteUrl, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := h.c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
