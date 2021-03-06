package imm

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

// ListFaceGroups invokes the imm.ListFaceGroups API synchronously
// api document: https://help.aliyun.com/api/imm/listfacegroups.html
func (client *Client) ListFaceGroups(request *ListFaceGroupsRequest) (response *ListFaceGroupsResponse, err error) {
	response = CreateListFaceGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFaceGroupsWithChan invokes the imm.ListFaceGroups API asynchronously
// api document: https://help.aliyun.com/api/imm/listfacegroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFaceGroupsWithChan(request *ListFaceGroupsRequest) (<-chan *ListFaceGroupsResponse, <-chan error) {
	responseChan := make(chan *ListFaceGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaceGroups(request)
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

// ListFaceGroupsWithCallback invokes the imm.ListFaceGroups API asynchronously
// api document: https://help.aliyun.com/api/imm/listfacegroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFaceGroupsWithCallback(request *ListFaceGroupsRequest, callback func(response *ListFaceGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFaceGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListFaceGroups(request)
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

// ListFaceGroupsRequest is the request struct for api ListFaceGroups
type ListFaceGroupsRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Marker  requests.Integer `position:"Query" name:"Marker"`
	Project string           `position:"Query" name:"Project"`
	SetId   string           `position:"Query" name:"SetId"`
}

// ListFaceGroupsResponse is the response struct for api ListFaceGroups
type ListFaceGroupsResponse struct {
	*responses.BaseResponse
	RequestId  string                       `json:"RequestId" xml:"RequestId"`
	NextMarker int                          `json:"NextMarker" xml:"NextMarker"`
	Groups     []GroupsItemInListFaceGroups `json:"Groups" xml:"Groups"`
}

// CreateListFaceGroupsRequest creates a request to invoke ListFaceGroups API
func CreateListFaceGroupsRequest() (request *ListFaceGroupsRequest) {
	request = &ListFaceGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListFaceGroups", "imm", "openAPI")
	return
}

// CreateListFaceGroupsResponse creates a response to parse from ListFaceGroups response
func CreateListFaceGroupsResponse() (response *ListFaceGroupsResponse) {
	response = &ListFaceGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
