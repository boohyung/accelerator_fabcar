/*
 *    Copyright 2019 Samsung SDS
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// 현재 파일이 fabcar 패키지에 포함됨
package fabcar

import (
	"context"
	"fmt"
	"strconv"

	// 테스트에서 사용하는 표준 패키지
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"

	pbbatch "github.com/nexledger/accelerator/protos"
)

const (
	channelId     = "accelerator"
	chaincodeName = "fabcar"
	numOfPings    = 1000             // tx 개수
	address       = "127.0.0.1:5050" // 서버 주소
)

/**
 * @dev 메서드의 실행 시간을 측정한다.
 * @param tag 메서드 이름, msg 메시지
 * @return 익명 함수를 통해 경과 시간을 리턴
 */
func ElapsedTime(tag string, msg string) func() {
	if msg != "" {
		log.Printf("[%s] %s", tag, msg)
	}
	// 현재 = 시작 시간
	start := time.Now()
	// 경과 시간
	return func() { log.Printf("[%s] Elapsed Time: %s", tag, time.Since(start)) }
}

func TestAccelerator(t *testing.T) {

	initLedger(t)     // 초기화 목적으로 10개의 key:value pair를 원장에 write
	createCar(t)      // 1개의 key:value pair를 원장에 write
	queryCar(t)       // 원장에서 1개의 key를 기준으로 value를 read
	changeCarOwner(t) // 1개의 key를 기준으로 특정 value를 원장에 write(modify)
	queryAllCars(t)   // 원장에 저장된 모든 key:value를 read(최대 1000개)
}

func initLedger(t *testing.T) {

	defer ElapsedTime("initLedger", "start")()

	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifiers := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifiers[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "initLedger",
			}
			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				notifier <- "Failed to execute" + err.Error()
			} else {
				notifier <- "TxId: " + resp.TxId
			}
		}(i, notifier)
	}

	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifiers[i])
	}
	// fmt.Println(<-notifier)

}

func queryAllCars(t *testing.T) {

	defer ElapsedTime("queryAllCars", "start")()

	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifier := make(chan string)
	go func(notifier chan string) {
		req := &pbbatch.TxRequest{
			ChannelId:     channelId,
			ChaincodeName: chaincodeName,
			Fcn:           "queryAllCars",
		}
		resp, err := client.Query(context.Background(), req)

		if err != nil {
			notifier <- "Failed to query" + err.Error()
		} else {
			notifier <- string(resp.Payload)
		}
	}(notifier)
	fmt.Println(<-notifier)

}

func createCar(t *testing.T) {

	defer ElapsedTime("createCar", "start")()

	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifiers := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifiers[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "createCar",
				Args:          [][]byte{[]byte("CAR" + strconv.Itoa(i)), []byte("Make" + strconv.Itoa(i)), []byte("Model" + strconv.Itoa(i)), []byte("Colour" + strconv.Itoa(i)), []byte("Owner" + strconv.Itoa(i))},
			}
			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				notifier <- "Failed to execute" + err.Error()
			} else {
				notifier <- "TxId of CAR" + strconv.Itoa(i) + ":" + resp.TxId
			}
		}(i, notifier)
	}
	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifiers[i])
	}
}

func queryCar(t *testing.T) {

	defer ElapsedTime("queryCar", "start")()

	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifiers := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifiers[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "queryCar",
				Args:          [][]byte{[]byte("CAR" + strconv.Itoa(i))},
			}
			resp, err := client.Query(context.Background(), req)
			// fmt.Println(resp.Payload)
			if err != nil {
				notifier <- "Failed to query" + err.Error()
			} else {
				notifier <- "CAR" + strconv.Itoa(i) + ":" + string(resp.Payload)
			}
		}(i, notifier)
	}

	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifiers[i])
	}
}

func changeCarOwner(t *testing.T) {

	defer ElapsedTime("changeCarOwner", "start")()

	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifiers := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifiers[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "changeCarOwner",
				Args:          [][]byte{[]byte("CAR" + strconv.Itoa(i)), []byte("Owner" + strconv.Itoa(i+1))},
			}
			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				notifier <- "Failed to execute" + err.Error()
			} else {
				notifier <- "TxId of CAR" + strconv.Itoa(i) + ":" + resp.TxId
			}
		}(i, notifier)
	}

	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifiers[i])
	}
}

func connect(t *testing.T) *grpc.ClientConn {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect server.", err)
	}
	return cc
}
