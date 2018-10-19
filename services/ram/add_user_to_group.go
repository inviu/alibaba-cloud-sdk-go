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

// AddUserToGroup invokes the ram.AddUserToGroup API synchronously
// api document: https://help.aliyun.com/api/ram/addusertogroup.html
func (client *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
	response = CreateAddUserToGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserToGroupWithChan invokes the ram.AddUserToGroup API asynchronously
// api document: https://help.aliyun.com/api/ram/addusertogroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddUserToGroupWithChan(request *AddUserToGroupRequest) (<-chan *AddUserToGroupResponse, <-chan error) {
	responseChan := make(chan *AddUserToGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUserToGroup(request)
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

// AddUserToGroupWithCallback invokes the ram.AddUserToGroup API asynchronously
// api document: https://help.aliyun.com/api/ram/addusertogroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddUserToGroupWithCallback(request *AddUserToGroupRequest, callback func(response *AddUserToGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserToGroupResponse
		var err error
		defer close(result)
		response, err = client.AddUserToGroup(request)
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

// AddUserToGroupRequest is the request struct for api AddUserToGroup
type AddUserToGroupRequest struct {
	*requests.RpcRequest
	UserName  string `position:"Query" name:"UserName"`
	GroupName string `position:"Query" name:"GroupName"`
}

// AddUserToGroupResponse is the response struct for api AddUserToGroup
type AddUserToGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddUserToGroupRequest creates a request to invoke AddUserToGroup API
func CreateAddUserToGroupRequest() (request *AddUserToGroupRequest) {
	request = &AddUserToGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "AddUserToGroup", "", "")
	return
}

// CreateAddUserToGroupResponse creates a response to parse from AddUserToGroup response
func CreateAddUserToGroupResponse() (response *AddUserToGroupResponse) {
	response = &AddUserToGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
