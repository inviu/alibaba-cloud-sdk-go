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

// GetUser invokes the ccc.GetUser API synchronously
// api document: https://help.aliyun.com/api/ccc/getuser.html
func (client *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
	response = CreateGetUserResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserWithChan invokes the ccc.GetUser API asynchronously
// api document: https://help.aliyun.com/api/ccc/getuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserWithChan(request *GetUserRequest) (<-chan *GetUserResponse, <-chan error) {
	responseChan := make(chan *GetUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUser(request)
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

// GetUserWithCallback invokes the ccc.GetUser API asynchronously
// api document: https://help.aliyun.com/api/ccc/getuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserWithCallback(request *GetUserRequest, callback func(response *GetUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserResponse
		var err error
		defer close(result)
		response, err = client.GetUser(request)
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

// GetUserRequest is the request struct for api GetUser
type GetUserRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

// GetUserResponse is the response struct for api GetUser
type GetUserResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	User           User   `json:"User" xml:"User"`
}

// CreateGetUserRequest creates a request to invoke GetUser API
func CreateGetUserRequest() (request *GetUserRequest) {
	request = &GetUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetUser", "", "")
	return
}

// CreateGetUserResponse creates a response to parse from GetUser response
func CreateGetUserResponse() (response *GetUserResponse) {
	response = &GetUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
