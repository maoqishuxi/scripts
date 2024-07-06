package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// 目标 URL
	targetURL := "https://subabase-crawler.vercel.app/api/weibo?uid="

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
		for _, uid := range uids {
			// 创建 HTTP 请求
			req, err := http.NewRequest("GET", targetURL+uid, nil)
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
			_, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read response body: %v", err)
			}
			resp.Body.Close()

			// 打印响应
			// fmt.Printf("Response: %s\n", body)

			// 每隔一段时间发送一次请求
			sleepDuration := time.Duration(rand.Intn(31) + 20)
			fmt.Printf("run %s finish, sleep %ss \n", uid, sleepDuration)
			time.Sleep(sleepDuration * time.Second)
		}

		// 每隔一段时间发送一次请求
		sleepDurationALL := time.Duration(rand.Intn(240) + 360)
		fmt.Printf("run all finish, sleep %ss \n", sleepDurationALL)
		time.Sleep(sleepDurationALL * time.Second)
	}
}
