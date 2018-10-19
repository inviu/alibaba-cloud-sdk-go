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

// SubmitAIVideoSummaryJob invokes the vod.SubmitAIVideoSummaryJob API synchronously
// api document: https://help.aliyun.com/api/vod/submitaivideosummaryjob.html
func (client *Client) SubmitAIVideoSummaryJob(request *SubmitAIVideoSummaryJobRequest) (response *SubmitAIVideoSummaryJobResponse, err error) {
	response = CreateSubmitAIVideoSummaryJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAIVideoSummaryJobWithChan invokes the vod.SubmitAIVideoSummaryJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitaivideosummaryjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitAIVideoSummaryJobWithChan(request *SubmitAIVideoSummaryJobRequest) (<-chan *SubmitAIVideoSummaryJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAIVideoSummaryJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAIVideoSummaryJob(request)
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

// SubmitAIVideoSummaryJobWithCallback invokes the vod.SubmitAIVideoSummaryJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitaivideosummaryjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitAIVideoSummaryJobWithCallback(request *SubmitAIVideoSummaryJobRequest, callback func(response *SubmitAIVideoSummaryJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAIVideoSummaryJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAIVideoSummaryJob(request)
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

// SubmitAIVideoSummaryJobRequest is the request struct for api SubmitAIVideoSummaryJob
type SubmitAIVideoSummaryJobRequest struct {
	*requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	AIVideoSummaryConfig string `position:"Query" name:"AIVideoSummaryConfig"`
}

// SubmitAIVideoSummaryJobResponse is the response struct for api SubmitAIVideoSummaryJob
type SubmitAIVideoSummaryJobResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	AIVideoSummaryJob AIVideoSummaryJob `json:"AIVideoSummaryJob" xml:"AIVideoSummaryJob"`
}

// CreateSubmitAIVideoSummaryJobRequest creates a request to invoke SubmitAIVideoSummaryJob API
func CreateSubmitAIVideoSummaryJobRequest() (request *SubmitAIVideoSummaryJobRequest) {
	request = &SubmitAIVideoSummaryJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoSummaryJob", "vod", "openAPI")
	return
}

// CreateSubmitAIVideoSummaryJobResponse creates a response to parse from SubmitAIVideoSummaryJob response
func CreateSubmitAIVideoSummaryJobResponse() (response *SubmitAIVideoSummaryJobResponse) {
	response = &SubmitAIVideoSummaryJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
