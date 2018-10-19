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

// ListAIVideoCoverJob invokes the vod.ListAIVideoCoverJob API synchronously
// api document: https://help.aliyun.com/api/vod/listaivideocoverjob.html
func (client *Client) ListAIVideoCoverJob(request *ListAIVideoCoverJobRequest) (response *ListAIVideoCoverJobResponse, err error) {
	response = CreateListAIVideoCoverJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListAIVideoCoverJobWithChan invokes the vod.ListAIVideoCoverJob API asynchronously
// api document: https://help.aliyun.com/api/vod/listaivideocoverjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAIVideoCoverJobWithChan(request *ListAIVideoCoverJobRequest) (<-chan *ListAIVideoCoverJobResponse, <-chan error) {
	responseChan := make(chan *ListAIVideoCoverJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAIVideoCoverJob(request)
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

// ListAIVideoCoverJobWithCallback invokes the vod.ListAIVideoCoverJob API asynchronously
// api document: https://help.aliyun.com/api/vod/listaivideocoverjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAIVideoCoverJobWithCallback(request *ListAIVideoCoverJobRequest, callback func(response *ListAIVideoCoverJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAIVideoCoverJobResponse
		var err error
		defer close(result)
		response, err = client.ListAIVideoCoverJob(request)
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

// ListAIVideoCoverJobRequest is the request struct for api ListAIVideoCoverJob
type ListAIVideoCoverJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AIVideoCoverJobIds   string `position:"Query" name:"AIVideoCoverJobIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

// ListAIVideoCoverJobResponse is the response struct for api ListAIVideoCoverJob
type ListAIVideoCoverJobResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	NonExistAIVideoCoverJobIds NonExistAIVideoCoverJobIds `json:"NonExistAIVideoCoverJobIds" xml:"NonExistAIVideoCoverJobIds"`
	AIVideoCoverJobList        AIVideoCoverJobList        `json:"AIVideoCoverJobList" xml:"AIVideoCoverJobList"`
}

// CreateListAIVideoCoverJobRequest creates a request to invoke ListAIVideoCoverJob API
func CreateListAIVideoCoverJobRequest() (request *ListAIVideoCoverJobRequest) {
	request = &ListAIVideoCoverJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoCoverJob", "vod", "openAPI")
	return
}

// CreateListAIVideoCoverJobResponse creates a response to parse from ListAIVideoCoverJob response
func CreateListAIVideoCoverJobResponse() (response *ListAIVideoCoverJobResponse) {
	response = &ListAIVideoCoverJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
