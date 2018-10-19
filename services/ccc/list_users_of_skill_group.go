package ccc

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

// ListUsersOfSkillGroup invokes the ccc.ListUsersOfSkillGroup API synchronously
// api document: https://help.aliyun.com/api/ccc/listusersofskillgroup.html
func (client *Client) ListUsersOfSkillGroup(request *ListUsersOfSkillGroupRequest) (response *ListUsersOfSkillGroupResponse, err error) {
	response = CreateListUsersOfSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListUsersOfSkillGroupWithChan invokes the ccc.ListUsersOfSkillGroup API asynchronously
// api document: https://help.aliyun.com/api/ccc/listusersofskillgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUsersOfSkillGroupWithChan(request *ListUsersOfSkillGroupRequest) (<-chan *ListUsersOfSkillGroupResponse, <-chan error) {
	responseChan := make(chan *ListUsersOfSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUsersOfSkillGroup(request)
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

// ListUsersOfSkillGroupWithCallback invokes the ccc.ListUsersOfSkillGroup API asynchronously
// api document: https://help.aliyun.com/api/ccc/listusersofskillgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUsersOfSkillGroupWithCallback(request *ListUsersOfSkillGroupRequest, callback func(response *ListUsersOfSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUsersOfSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.ListUsersOfSkillGroup(request)
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

// ListUsersOfSkillGroupRequest is the request struct for api ListUsersOfSkillGroup
type ListUsersOfSkillGroupRequest struct {
	*requests.RpcRequest
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListUsersOfSkillGroupResponse is the response struct for api ListUsersOfSkillGroup
type ListUsersOfSkillGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Users          Users  `json:"Users" xml:"Users"`
}

// CreateListUsersOfSkillGroupRequest creates a request to invoke ListUsersOfSkillGroup API
func CreateListUsersOfSkillGroupRequest() (request *ListUsersOfSkillGroupRequest) {
	request = &ListUsersOfSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListUsersOfSkillGroup", "", "")
	return
}

// CreateListUsersOfSkillGroupResponse creates a response to parse from ListUsersOfSkillGroup response
func CreateListUsersOfSkillGroupResponse() (response *ListUsersOfSkillGroupResponse) {
	response = &ListUsersOfSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
