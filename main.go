package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// 目标 URL
	targetURL := "https://subabase-crawler.vercel.app/api/weibo?uid=1656918431"

	// 代理 URL
	proxyURL, err := url.Parse("http://192.168.0.208:1080")
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v", err)
	}

	// uids
	uids := []string{"1281503010", "1784977674", "1922411040", "2423296524", "1751162512", "1656918431"}
	// 创建一个自定义的 HTTP 客户端，使用代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: transport,
	}

	// 循环发送请求
	for {
		// 创建 HTTP 请求
		req, err := http.NewRequest("GET", targetURL, nil)
		if err != nil {
			log.Fatalf("Failed to create HTTP request: %v", err)
		}

		// 发送 HTTP 请求
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send HTTP request: %v", err)
			continue
		}

		// 读取响应
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
		}
		resp.Body.Close()

		// 打印响应
		fmt.Printf("Response: %s\n", body)

		// 每隔一段时间发送一次请求
		time.Sleep(10 * time.Second)
	}
}
