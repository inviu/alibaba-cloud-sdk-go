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

// DeleteVideo invokes the vod.DeleteVideo API synchronously
// api document: https://help.aliyun.com/api/vod/deletevideo.html
func (client *Client) DeleteVideo(request *DeleteVideoRequest) (response *DeleteVideoResponse, err error) {
	response = CreateDeleteVideoResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVideoWithChan invokes the vod.DeleteVideo API asynchronously
// api document: https://help.aliyun.com/api/vod/deletevideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVideoWithChan(request *DeleteVideoRequest) (<-chan *DeleteVideoResponse, <-chan error) {
	responseChan := make(chan *DeleteVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVideo(request)
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

// DeleteVideoWithCallback invokes the vod.DeleteVideo API asynchronously
// api document: https://help.aliyun.com/api/vod/deletevideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVideoWithCallback(request *DeleteVideoRequest, callback func(response *DeleteVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVideoResponse
		var err error
		defer close(result)
		response, err = client.DeleteVideo(request)
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

// DeleteVideoRequest is the request struct for api DeleteVideo
type DeleteVideoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VideoIds             string           `position:"Query" name:"VideoIds"`
}

// DeleteVideoResponse is the response struct for api DeleteVideo
type DeleteVideoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVideoRequest creates a request to invoke DeleteVideo API
func CreateDeleteVideoRequest() (request *DeleteVideoRequest) {
	request = &DeleteVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteVideo", "vod", "openAPI")
	return
}

// CreateDeleteVideoResponse creates a response to parse from DeleteVideo response
func CreateDeleteVideoResponse() (response *DeleteVideoResponse) {
	response = &DeleteVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
