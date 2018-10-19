package vod

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

// OpenVodService invokes the vod.OpenVodService API synchronously
// api document: https://help.aliyun.com/api/vod/openvodservice.html
func (client *Client) OpenVodService(request *OpenVodServiceRequest) (response *OpenVodServiceResponse, err error) {
	response = CreateOpenVodServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenVodServiceWithChan invokes the vod.OpenVodService API asynchronously
// api document: https://help.aliyun.com/api/vod/openvodservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OpenVodServiceWithChan(request *OpenVodServiceRequest) (<-chan *OpenVodServiceResponse, <-chan error) {
	responseChan := make(chan *OpenVodServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenVodService(request)
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

// OpenVodServiceWithCallback invokes the vod.OpenVodService API asynchronously
// api document: https://help.aliyun.com/api/vod/openvodservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OpenVodServiceWithCallback(request *OpenVodServiceRequest, callback func(response *OpenVodServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenVodServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenVodService(request)
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

// OpenVodServiceRequest is the request struct for api OpenVodService
type OpenVodServiceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

// OpenVodServiceResponse is the response struct for api OpenVodService
type OpenVodServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Success   bool   `json:"success" xml:"success"`
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
}

// CreateOpenVodServiceRequest creates a request to invoke OpenVodService API
func CreateOpenVodServiceRequest() (request *OpenVodServiceRequest) {
	request = &OpenVodServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "OpenVodService", "vod", "openAPI")
	return
}

// CreateOpenVodServiceResponse creates a response to parse from OpenVodService response
func CreateOpenVodServiceResponse() (response *OpenVodServiceResponse) {
	response = &OpenVodServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
