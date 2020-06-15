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

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
//param stub 체인코드인터페이스, target 체인코드의 매개변수(함수명, 함수의 인수)
func Invoke(stub shim.ChaincodeStubInterface, target func(shim.ChaincodeStubInterface, []string) pb.Response) pb.Response {
	// 3중 배열 선언 및 0으로 초기화
	items := make([][][]byte, 0)
	if err := decode(stub.GetArgs()[1], &items); err != nil {
		return shim.Error("Failed to unmarshal request")
	}
	// items 배열 길이
	itemSize := len(items)
	// 2차원 배열 생성(byte 타입, 길이, 용량)
	payloads := make([][]byte, itemSize, itemSize)
	// 2차원 배열 items 요소 탐색 "for 인덱스, 값 := range 배열"
	for i, item := range items {
		// 요소 길이
		argsSize := len(item)
		// 1차원 배열 생성(string 타입, 길이, 용량)
		args := make([]string, argsSize, argsSize)
		// 배열 item 내 요소 탐색
		for j, arg := range item {
			// item 내 요소를 string 타입으로 변환 후 차례로 삽입
			args[j] = string(arg)
		}
		// fabcar 체인코드 내 함수 실행 결과
		result := target(stub, args)
		// 체인코드의 print를 확인하려면?
		fmt.Println(result)
		if result.Status == shim.ERROR {
			return shim.Error("Failed to invoke: " + result.Message)
		}
		//payloads 입력
		payloads[i] = result.Payload
	}
	// 인코딩
	response, err := encode(payloads)
	if err != nil {
		return shim.Error("Failed to marshal response")
	}
	// 체인코드 실행 결과 리턴
	return shim.Success(response)
}

func Invoke_no_arg(stub shim.ChaincodeStubInterface, target func(shim.ChaincodeStubInterface) pb.Response) pb.Response {
	// 3중 배열 선언 및 0으로 초기화
	items := make([][][]byte, 0)
	if err := decode(stub.GetArgs()[1], &items); err != nil {
		return shim.Error("Failed to unmarshal request")
	}
	// items 배열 길이
	itemSize := len(items)
	// 2차원 배열 생성(byte 타입, 길이, 용량)
	payloads := make([][]byte, itemSize, itemSize)
	// 2차원 배열 items 요소 탐색 "for 인덱스, 값 := range 배열"
	for i, item := range items {
		// 요소 길이
		argsSize := len(item)
		// 1차원 배열 생성(string 타입, 길이, 용량)
		args := make([]string, argsSize, argsSize)
		// 배열 item 내 요소 탐색
		for j, arg := range item {
			// item 내 요소를 string 타입으로 변환 후 차례로 삽입
			args[j] = string(arg)
		}
		// fabcar 체인코드 내 함수 실행 결과
		result := target(stub)
		// 체인코드의 print를 확인하려면?
		fmt.Println(result)
		if result.Status == shim.ERROR {
			return shim.Error("Failed to invoke: " + result.Message)
		}
		//payloads 입력
		payloads[i] = result.Payload
	}
	// 인코딩
	response, err := encode(payloads)
	if err != nil {
		return shim.Error("Failed to marshal response")
	}
	// 체인코드 실행 결과 리턴
	return shim.Success(response)
}


func encode(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(d []byte, v interface{}) error {
	buf := bytes.NewBuffer(d)
	if err := gob.NewDecoder(buf).Decode(v); err != nil {
		return err
	}
	return nil
}
