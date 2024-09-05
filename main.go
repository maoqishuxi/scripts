package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	uids2 "scripts/uids"
	"time"
)

func main() {
	// 循环发送请求
	for {
		uids := uids2.GetUids()

		for _, uid := range uids.Data {
			if !uid.Status {
				continue
			}

			// 目标 URL
			targetURL := fmt.Sprintf("http://supabase.julai.fun:9000/api/weibo?uid=%s&page=%d&count=%d", uid.UID, 1, 10)

			fmt.Printf("%s start run %s %s \n", time.Now().In(time.FixedZone("CST", 8*60*60)).Format("2006-01-02 15:04:05"), uid.ScreenName, uid.UID)
			// 创建 HTTP 请求
			req, err := http.NewRequest("GET", targetURL, nil)
			if err != nil {
				log.Fatalf("Failed to create HTTP request: %v", err)
			}

			// 发送 HTTP 请求
			resp, err := uids2.Proxy().Do(req)
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
			sleepDuration := time.Duration(rand.Intn(21) + 10)
			fmt.Printf("%s run %s %s finish, sleep %s \n", time.Now().In(time.FixedZone("CST", 8*60*60)).Format("2006-01-02 15:04:05"), uid.ScreenName, uid.UID, sleepDuration)
			time.Sleep(sleepDuration * time.Second)
		}

		// 每隔一段时间发送一次请求
		sleepDurationALL := time.Duration(rand.Intn(240) + 240)
		fmt.Printf("%s run all finish, sleep %s \n", time.Now().In(time.FixedZone("CST", 8*60*60)).Format("2006-01-02 15:04:05"), sleepDurationALL)
		time.Sleep(sleepDurationALL * time.Second)
	}
}
