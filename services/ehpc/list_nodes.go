package ehpc

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

// ListNodes invokes the ehpc.ListNodes API synchronously
// api document: https://help.aliyun.com/api/ehpc/listnodes.html
func (client *Client) ListNodes(request *ListNodesRequest) (response *ListNodesResponse, err error) {
	response = CreateListNodesResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodesWithChan invokes the ehpc.ListNodes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listnodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNodesWithChan(request *ListNodesRequest) (<-chan *ListNodesResponse, <-chan error) {
	responseChan := make(chan *ListNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodes(request)
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

// ListNodesWithCallback invokes the ehpc.ListNodes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listnodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNodesWithCallback(request *ListNodesRequest, callback func(response *ListNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodesResponse
		var err error
		defer close(result)
		response, err = client.ListNodes(request)
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

// ListNodesRequest is the request struct for api ListNodes
type ListNodesRequest struct {
	*requests.RpcRequest
	HostName   string           `position:"Query" name:"HostName"`
	Role       string           `position:"Query" name:"Role"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListNodesResponse is the response struct for api ListNodes
type ListNodesResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	TotalCount int              `json:"TotalCount" xml:"TotalCount"`
	PageNumber int              `json:"PageNumber" xml:"PageNumber"`
	PageSize   int              `json:"PageSize" xml:"PageSize"`
	Nodes      NodesInListNodes `json:"Nodes" xml:"Nodes"`
}

// CreateListNodesRequest creates a request to invoke ListNodes API
func CreateListNodesRequest() (request *ListNodesRequest) {
	request = &ListNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListNodes", "ehs", "openAPI")
	return
}

// CreateListNodesResponse creates a response to parse from ListNodes response
func CreateListNodesResponse() (response *ListNodesResponse) {
	response = &ListNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
