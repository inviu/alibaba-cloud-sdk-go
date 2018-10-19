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

// DeactivateMediaWorkflow invokes the mts.DeactivateMediaWorkflow API synchronously
// api document: https://help.aliyun.com/api/mts/deactivatemediaworkflow.html
func (client *Client) DeactivateMediaWorkflow(request *DeactivateMediaWorkflowRequest) (response *DeactivateMediaWorkflowResponse, err error) {
	response = CreateDeactivateMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// DeactivateMediaWorkflowWithChan invokes the mts.DeactivateMediaWorkflow API asynchronously
// api document: https://help.aliyun.com/api/mts/deactivatemediaworkflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeactivateMediaWorkflowWithChan(request *DeactivateMediaWorkflowRequest) (<-chan *DeactivateMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *DeactivateMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeactivateMediaWorkflow(request)
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

// DeactivateMediaWorkflowWithCallback invokes the mts.DeactivateMediaWorkflow API asynchronously
// api document: https://help.aliyun.com/api/mts/deactivatemediaworkflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeactivateMediaWorkflowWithCallback(request *DeactivateMediaWorkflowRequest, callback func(response *DeactivateMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeactivateMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.DeactivateMediaWorkflow(request)
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

// DeactivateMediaWorkflowRequest is the request struct for api DeactivateMediaWorkflow
type DeactivateMediaWorkflowRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string           `position:"Query" name:"MediaWorkflowId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeactivateMediaWorkflowResponse is the response struct for api DeactivateMediaWorkflow
type DeactivateMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	MediaWorkflow MediaWorkflow `json:"MediaWorkflow" xml:"MediaWorkflow"`
}

// CreateDeactivateMediaWorkflowRequest creates a request to invoke DeactivateMediaWorkflow API
func CreateDeactivateMediaWorkflowRequest() (request *DeactivateMediaWorkflowRequest) {
	request = &DeactivateMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "DeactivateMediaWorkflow", "mts", "openAPI")
	return
}

// CreateDeactivateMediaWorkflowResponse creates a response to parse from DeactivateMediaWorkflow response
func CreateDeactivateMediaWorkflowResponse() (response *DeactivateMediaWorkflowResponse) {
	response = &DeactivateMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
