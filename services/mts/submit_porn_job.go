package mts

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

// SubmitPornJob invokes the mts.SubmitPornJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitpornjob.html
func (client *Client) SubmitPornJob(request *SubmitPornJobRequest) (response *SubmitPornJobResponse, err error) {
	response = CreateSubmitPornJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitPornJobWithChan invokes the mts.SubmitPornJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitpornjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitPornJobWithChan(request *SubmitPornJobRequest) (<-chan *SubmitPornJobResponse, <-chan error) {
	responseChan := make(chan *SubmitPornJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitPornJob(request)
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

// SubmitPornJobWithCallback invokes the mts.SubmitPornJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitpornjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitPornJobWithCallback(request *SubmitPornJobRequest, callback func(response *SubmitPornJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitPornJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitPornJob(request)
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

// SubmitPornJobRequest is the request struct for api SubmitPornJob
type SubmitPornJobRequest struct {
	*requests.RpcRequest
	Input                string           `position:"Query" name:"Input"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PornConfig           string           `position:"Query" name:"PornConfig"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
}

// SubmitPornJobResponse is the response struct for api SubmitPornJob
type SubmitPornJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitPornJobRequest creates a request to invoke SubmitPornJob API
func CreateSubmitPornJobRequest() (request *SubmitPornJobRequest) {
	request = &SubmitPornJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitPornJob", "mts", "openAPI")
	return
}

// CreateSubmitPornJobResponse creates a response to parse from SubmitPornJob response
func CreateSubmitPornJobResponse() (response *SubmitPornJobResponse) {
	response = &SubmitPornJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
