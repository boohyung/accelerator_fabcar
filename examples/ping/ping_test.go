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

package ping

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"google.golang.org/grpc"

	pbbatch "github.com/nexledger/accelerator/protos"
)

const (
	channelId     = "accelerator"
	chaincodeName = "ping"
	numOfPings    = 50
	address       = "127.0.0.1:5050"
)

func TestAccelerator(t *testing.T) {
	ping(t)
	pong(t)
}

func ping(t *testing.T) {
	// grpc 클라이언트 생성
	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	// 50개의 string 타입 채널 배열? 생성 
	notifiers := make([]chan string, numOfPings)
	// 0~49까지 50번 반복
	for i := 0; i < numOfPings; i++ {
		// string 타입의 채널 notifier 생성
		notifier := make(chan string)
		notifiers[i] = notifier
		// 익명 함수에 대한 goroutine 실행(비동기)
		go func(i int, notifier chan string) {
			// TxRequest 메시지 
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "ping",
				// Args는 2차원 슬라이스 [[i를 바이트로 변환한 값], ["value of"+i를 바이트로 변환한 값]]
				Args:          [][]byte{[]byte(strconv.Itoa(i)), []byte("value of " + strconv.Itoa(i))},
			}
			// fmt.Println(i, req.Args)
			// accelerator/pkg/server/server.go에 선언
			// context.Background(): goroutine의 생애 주기를 관리하기 위한 컨텍스트 생성
			// TxRequest 메시지에는 {채널ID, 체인코드이름, 함수, 페이로드}가 포함
			resp, err := client.Execute(context.Background(), req)
			// txId와 Validation 출력
			// fmt.Println(resp)

			if err != nil {
				notifier <- "Failed to execute" + err.Error()
			} else {
				// notifier 채널에 "i:Txid" 데이터를 전달
				notifier <- strconv.Itoa(i) + ":" + resp.TxId
				// notifier <- strconv.Itoa(i) + ":" + string(resp.Payload)
			}
		}(i, notifier) // i, notifier 파라미터 입력
	}
	// 동기적으로 출력 수행(ping goroutine 완료 표시)
	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifiers[i])
	}
}

func pong(t *testing.T) {
	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifiers := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifiers[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "pong",
				Args:          [][]byte{[]byte(strconv.Itoa(i))},
			}
			resp, err := client.Query(context.Background(), req)
			fmt.Println(resp)
			if err != nil {
				notifier <- "Failed to query" + err.Error()
			} else {
				// notifier 채널에 "i:Payload" 값 전달
				notifier <- strconv.Itoa(i) + ":" + string(resp.Payload)
			}
		}(i, notifier)
	}

	for i := 0; i < numOfPings; i++ {
		// notifier 채널 안의 값 출력
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
