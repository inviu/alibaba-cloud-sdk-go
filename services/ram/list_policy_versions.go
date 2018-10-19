package ram

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

// ListPolicyVersions invokes the ram.ListPolicyVersions API synchronously
// api document: https://help.aliyun.com/api/ram/listpolicyversions.html
func (client *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (response *ListPolicyVersionsResponse, err error) {
	response = CreateListPolicyVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPolicyVersionsWithChan invokes the ram.ListPolicyVersions API asynchronously
// api document: https://help.aliyun.com/api/ram/listpolicyversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPolicyVersionsWithChan(request *ListPolicyVersionsRequest) (<-chan *ListPolicyVersionsResponse, <-chan error) {
	responseChan := make(chan *ListPolicyVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPolicyVersions(request)
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

// ListPolicyVersionsWithCallback invokes the ram.ListPolicyVersions API asynchronously
// api document: https://help.aliyun.com/api/ram/listpolicyversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPolicyVersionsWithCallback(request *ListPolicyVersionsRequest, callback func(response *ListPolicyVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPolicyVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListPolicyVersions(request)
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

// ListPolicyVersionsRequest is the request struct for api ListPolicyVersions
type ListPolicyVersionsRequest struct {
	*requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

// ListPolicyVersionsResponse is the response struct for api ListPolicyVersions
type ListPolicyVersionsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	PolicyVersions PolicyVersions `json:"PolicyVersions" xml:"PolicyVersions"`
}

// CreateListPolicyVersionsRequest creates a request to invoke ListPolicyVersions API
func CreateListPolicyVersionsRequest() (request *ListPolicyVersionsRequest) {
	request = &ListPolicyVersionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListPolicyVersions", "", "")
	return
}

// CreateListPolicyVersionsResponse creates a response to parse from ListPolicyVersions response
func CreateListPolicyVersionsResponse() (response *ListPolicyVersionsResponse) {
	response = &ListPolicyVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
