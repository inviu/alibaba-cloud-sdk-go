package imm

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/inviu/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/inviu/alibaba-cloud-sdk-go/sdk/responses"
)

// DetectFace invokes the imm.DetectFace API synchronously
// api document: https://help.aliyun.com/api/imm/detectface.html
func (client *Client) DetectFace(request *DetectFaceRequest) (response *DetectFaceResponse, err error) {
	response = CreateDetectFaceResponse()
	err = client.DoAction(request, response)
	return
}

// DetectFaceWithChan invokes the imm.DetectFace API asynchronously
// api document: https://help.aliyun.com/api/imm/detectface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectFaceWithChan(request *DetectFaceRequest) (<-chan *DetectFaceResponse, <-chan error) {
	responseChan := make(chan *DetectFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectFace(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DetectFaceWithCallback invokes the imm.DetectFace API asynchronously
// api document: https://help.aliyun.com/api/imm/detectface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectFaceWithCallback(request *DetectFaceRequest, callback func(response *DetectFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectFaceResponse
		var err error
		defer close(result)
		response, err = client.DetectFace(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DetectFaceRequest is the request struct for api DetectFace
type DetectFaceRequest struct {
	*requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
}

// DetectFaceResponse is the response struct for api DetectFace
type DetectFaceResponse struct {
	*responses.BaseResponse
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	SrcUris        []string             `json:"SrcUris" xml:"SrcUris"`
	SuccessDetails []SuccessDetailsItem `json:"SuccessDetails" xml:"SuccessDetails"`
	FailDetails    []FailDetailsItem    `json:"FailDetails" xml:"FailDetails"`
}

// CreateDetectFaceRequest creates a request to invoke DetectFace API
func CreateDetectFaceRequest() (request *DetectFaceRequest) {
	request = &DetectFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DetectFace", "imm", "openAPI")
	return
}

// CreateDetectFaceResponse creates a response to parse from DetectFace response
func CreateDetectFaceResponse() (response *DetectFaceResponse) {
	response = &DetectFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
