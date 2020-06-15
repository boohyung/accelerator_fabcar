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

package batch

import (
	"time"

	"github.com/pkg/errors"

	"github.com/nexledger/accelerator/pkg/batch/queue"
	"github.com/nexledger/accelerator/pkg/batch/queue/cutter"
	"github.com/nexledger/accelerator/pkg/batch/route"
	"github.com/nexledger/accelerator/pkg/batch/route/encoding"
	"github.com/nexledger/accelerator/pkg/batch/route/fab"
	"github.com/nexledger/accelerator/pkg/batch/tx"
	"github.com/nexledger/accelerator/pkg/fabwrap"
)

type Acceleration struct {
	Type               string
	ChannelId          string
	ChaincodeName      string
	Fcn                string
	QueueSize          int
	MaxBatchItems      int
	MaxWaitTimeSeconds int64
	MaxBatchBytes      int
	ReadKeyIndices     []int
	WriteKeyIndices    []int
	Encoding           string
	Recovery           bool
}

type Client struct {
	ctx               fabwrap.Context
	executeSchedulers map[string]*queue.Scheduler
	// key타입: string
	// value타입: queue.Scheduler 구조체 포인터
	querySchedulers   map[string]*queue.Scheduler
}
// Client 구조체에 Execute 함수를 연결. s를 포인터 리시버로 사용
// 입력: 채널ID, 체인코드이름, 함수명, 2차원 배열 인수(byte = uint8)
// 리턴: 결과, 에러
func (s *Client) Execute(channelId, chaincodeName, fcn string, args [][]byte) (*tx.Result, error) {
	name := nameOf(channelId, chaincodeName, fcn)
	fmt.Println(name)
	fmt.Println(args)
	// executeSchedulers 맵에 같은 key가 있는지 확인
	if scheduler, ok := s.executeSchedulers[name]; !ok {
		return nil, errors.New("Execute Scheduler not found: " + name)
	} else {
		// 없다면 process() 실행
		fmt.Println("here!!")
		return process(scheduler, args)
	}
}
// Client 구조체에 Query 함수를 연결. s를 포인터 리시버로 사용
// 입력: 채널ID, 체인코드이름, 함수명, 2차원 배열 인수(byte = uint8)
func (s *Client) Query(channelId, chaincodeName, fcn string, args [][]byte) (*tx.Result, error) {
	// 큐 이름 생성
	name := nameOf(channelId, chaincodeName, fcn)
	if scheduler, ok := s.querySchedulers[name]; !ok {
		return nil, errors.New("Query Scheduler not found: " + name)
	} else {
		return process(scheduler, args)
	}
}

func (s *Client) Register(acc *Acceleration) error {
	var schedulers map[string]*queue.Scheduler
	if acc.Type == "execute" {
		schedulers = s.executeSchedulers
	} else if acc.Type == "query" {
		schedulers = s.querySchedulers
	} else {
		return errors.New("Unsupported type: " + acc.Type)
	}

	name := nameOf(acc.ChannelId, acc.ChaincodeName, acc.Fcn)
	if _, ok := schedulers[name]; ok {
		return errors.New("Scheduler already registered: " + name)
	}

	encoder, err := encoding.New(acc.Encoding)
	if err != nil {
		return err
	}

	invoker, err := fab.New(s.ctx, acc.ChannelId, acc.ChaincodeName, acc.Fcn, acc.Type, encoder)
	if err != nil {
		return err
	}

	sender, err := route.NewSender(invoker, route.NewResponder(encoder), acc.Recovery)
	if err != nil {
		return err
	}

	cutterOpts := make([]cutter.Composition, 0)
	if acc.MaxBatchItems > 0 {
		cutterOpts = append(cutterOpts, cutter.WithItemCountCutter(acc.MaxBatchItems))
	}
	if acc.MaxBatchBytes > 0 {
		cutterOpts = append(cutterOpts, cutter.WithByteLenCutter(acc.MaxBatchBytes))
	}
	if len(acc.ReadKeyIndices) > 0 {
		cutterOpts = append(cutterOpts, cutter.WithMVCCCutter(acc.ReadKeyIndices, acc.WriteKeyIndices))
	}

	scheduler := queue.NewScheduler(
		queue.NewProcessor(sender, cutterOpts),
		time.Duration(acc.MaxWaitTimeSeconds)*time.Second,
		acc.QueueSize,
	)
	scheduler.Start()
	schedulers[name] = scheduler

	return nil
}
// 큐 이름을 생성 "채널ID:체인코드이름:함수명"
func nameOf(channelId, chaincodeName, fcn string) string {
	return channelId + ":" + chaincodeName + ":" + fcn
}
// 입력: map 타입의 스케줄러, char형 2차원 배열 데이터
// 출력: tx.Result 구조체 포인터, 에러
func process(s *queue.Scheduler, args [][]byte) (*tx.Result, error) {
	// tx.Result 형 채널 notify 생성
	// 채널: goroutine끼리 데이터를 주고 받고, 실행 흐름을 제어하기 위해 사용. 동기식이라 채널에 값이 들어올 때까지 대기, 값이 들어오면 대기를 끝내고 다음 코드를 실행
	notify := make(chan *tx.Result)
	// Scheduler 구조체의 멤버 scheduledItems 채널에 char형 2차원 배열 데이터와 notify 채널의 데이터를 전달
	s.Schedule(&tx.Item{Args: args, Notifier: notify})
	// notify 채널에서 값을 꺼내 result에 대입
	result := <-notify
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func New(ctx fabwrap.Context) *Client {
	return &Client{
		ctx,
		make(map[string]*queue.Scheduler),
		make(map[string]*queue.Scheduler),
	}
}
