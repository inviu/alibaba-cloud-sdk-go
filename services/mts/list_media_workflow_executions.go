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

// ListMediaWorkflowExecutions invokes the mts.ListMediaWorkflowExecutions API synchronously
// api document: https://help.aliyun.com/api/mts/listmediaworkflowexecutions.html
func (client *Client) ListMediaWorkflowExecutions(request *ListMediaWorkflowExecutionsRequest) (response *ListMediaWorkflowExecutionsResponse, err error) {
	response = CreateListMediaWorkflowExecutionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMediaWorkflowExecutionsWithChan invokes the mts.ListMediaWorkflowExecutions API asynchronously
// api document: https://help.aliyun.com/api/mts/listmediaworkflowexecutions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListMediaWorkflowExecutionsWithChan(request *ListMediaWorkflowExecutionsRequest) (<-chan *ListMediaWorkflowExecutionsResponse, <-chan error) {
	responseChan := make(chan *ListMediaWorkflowExecutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMediaWorkflowExecutions(request)
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

// ListMediaWorkflowExecutionsWithCallback invokes the mts.ListMediaWorkflowExecutions API asynchronously
// api document: https://help.aliyun.com/api/mts/listmediaworkflowexecutions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListMediaWorkflowExecutionsWithCallback(request *ListMediaWorkflowExecutionsRequest, callback func(response *ListMediaWorkflowExecutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMediaWorkflowExecutionsResponse
		var err error
		defer close(result)
		response, err = client.ListMediaWorkflowExecutions(request)
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

// ListMediaWorkflowExecutionsRequest is the request struct for api ListMediaWorkflowExecutions
type ListMediaWorkflowExecutionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	InputFileURL         string           `position:"Query" name:"InputFileURL"`
	NextPageToken        string           `position:"Query" name:"NextPageToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      requests.Integer `position:"Query" name:"MaximumPageSize"`
	MediaWorkflowId      string           `position:"Query" name:"MediaWorkflowId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaWorkflowName    string           `position:"Query" name:"MediaWorkflowName"`
}

// ListMediaWorkflowExecutionsResponse is the response struct for api ListMediaWorkflowExecutions
type ListMediaWorkflowExecutionsResponse struct {
	*responses.BaseResponse
	RequestId                  string                                                  `json:"RequestId" xml:"RequestId"`
	NextPageToken              string                                                  `json:"NextPageToken" xml:"NextPageToken"`
	MediaWorkflowExecutionList MediaWorkflowExecutionListInListMediaWorkflowExecutions `json:"MediaWorkflowExecutionList" xml:"MediaWorkflowExecutionList"`
}

// CreateListMediaWorkflowExecutionsRequest creates a request to invoke ListMediaWorkflowExecutions API
func CreateListMediaWorkflowExecutionsRequest() (request *ListMediaWorkflowExecutionsRequest) {
	request = &ListMediaWorkflowExecutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListMediaWorkflowExecutions", "mts", "openAPI")
	return
}

// CreateListMediaWorkflowExecutionsResponse creates a response to parse from ListMediaWorkflowExecutions response
func CreateListMediaWorkflowExecutionsResponse() (response *ListMediaWorkflowExecutionsResponse) {
	response = &ListMediaWorkflowExecutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
