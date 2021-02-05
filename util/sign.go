package util

import (
	"fmt"
	"github.com/handsomestWei/go-misc/security"
	"sort"
	"strings"
)

// 参数加签，常用于接口参数校验
func SignParam(param map[string]interface{}, skipKey, tailKey string) string {

	// 1、参数key升序排列
	keyList := make([]string, 0)
	for k, _ := range param {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)

	// 2、key=value键值对用&连接起来，略过空值和指定key
	var signStr string
	for _, k := range keyList {
		if strings.EqualFold(k, skipKey) {
			continue
		}
		if val := fmt.Sprintf("%v", param[k]); val != "" {
			signStr += fmt.Sprintf("%s=%s&", k, val)
		}
	}

	// 3、在键值对末尾加上key=tailKey
	if tailKey != "" {
		signStr += fmt.Sprintf("key=%s&", tailKey)
	}

	// 4、计算md5
	return strings.ToUpper(security.Md5ToHexString(signStr))
}

// 参数验签，常用于接口参数校验
func ValidateParam(param map[string]interface{}, signKey string, extraParam string) bool {
	var signVal string // 签名串

	// 1、参数key升序排列
	keyList := make([]string, 0)
	for k, _ := range param {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)

	// 2、key=value键值对用&连接起来，略过空值和签名串
	var signStr string
	for _, k := range keyList {
		if strings.EqualFold(k, signKey) {
			if v, ok := param[signKey].(string); ok {
				// 获取签名串，验签用
				signVal = v
			}
			continue
		}
		if val := fmt.Sprintf("%v", param[k]); val != "" {
			signStr += fmt.Sprintf("%s=%s&", k, val)
		}
	}

	// 3、在键值对末尾加上额外内容
	if extraParam != "" {
		signStr += extraParam
	} else {
		// 去掉末尾的符号&
		extraParam = extraParam[:len(extraParam)-1]
	}

	// 4、md5验签
	return strings.ToUpper(signVal) == strings.ToUpper(security.Md5ToHexString(signStr))
}
