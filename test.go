package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	num:=300000
	group1 := sync.WaitGroup{}
	group2 := sync.WaitGroup{}
	group1.Add(num)
	group2.Add(num)
	startTime := time.Now().Unix()

	for i:=0;i<num ;i++  {
		go func(group11,group22  *sync.WaitGroup) {
			group11.Done()
			//a:= rand.Intn(100000000000)
			startTime1 := time.Now().Unix()
			//http.Get("http://localhost:3001/user/fastRegister/" + strconv.Itoa(a))
			http.Get("http://192.168.99.100:3003/user/test/")
			fmt.Println("一条请求发送成功:--->",time.Now().Unix()-startTime1)
			group22.Done()
		}(&group1,&group2)
	}
		group1.Wait()
		startEndTime:=time.Now().Unix()
		fmt.Println("所有请求发送成功:--->",startEndTime-startTime)
		fmt.Println("startEndTime-->",startEndTime)
		group2.Wait()
		fmt.Println("所有请求发送响应:--->",time.Now().Unix()-startEndTime)



}