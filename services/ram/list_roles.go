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

// ListRoles invokes the ram.ListRoles API synchronously
// api document: https://help.aliyun.com/api/ram/listroles.html
func (client *Client) ListRoles(request *ListRolesRequest) (response *ListRolesResponse, err error) {
	response = CreateListRolesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRolesWithChan invokes the ram.ListRoles API asynchronously
// api document: https://help.aliyun.com/api/ram/listroles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRolesWithChan(request *ListRolesRequest) (<-chan *ListRolesResponse, <-chan error) {
	responseChan := make(chan *ListRolesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRoles(request)
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

// ListRolesWithCallback invokes the ram.ListRoles API asynchronously
// api document: https://help.aliyun.com/api/ram/listroles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRolesWithCallback(request *ListRolesRequest, callback func(response *ListRolesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRolesResponse
		var err error
		defer close(result)
		response, err = client.ListRoles(request)
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

// ListRolesRequest is the request struct for api ListRoles
type ListRolesRequest struct {
	*requests.RpcRequest
	Marker   string           `position:"Query" name:"Marker"`
	MaxItems requests.Integer `position:"Query" name:"MaxItems"`
}

// ListRolesResponse is the response struct for api ListRoles
type ListRolesResponse struct {
	*responses.BaseResponse
	RequestId   string           `json:"RequestId" xml:"RequestId"`
	IsTruncated bool             `json:"IsTruncated" xml:"IsTruncated"`
	Marker      string           `json:"Marker" xml:"Marker"`
	Roles       RolesInListRoles `json:"Roles" xml:"Roles"`
}

// CreateListRolesRequest creates a request to invoke ListRoles API
func CreateListRolesRequest() (request *ListRolesRequest) {
	request = &ListRolesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListRoles", "", "")
	return
}

// CreateListRolesResponse creates a response to parse from ListRoles response
func CreateListRolesResponse() (response *ListRolesResponse) {
	response = &ListRolesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
