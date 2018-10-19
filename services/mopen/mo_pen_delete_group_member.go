package mopen

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

// MoPenDeleteGroupMember invokes the mopen.MoPenDeleteGroupMember API synchronously
// api document: https://help.aliyun.com/api/mopen/mopendeletegroupmember.html
func (client *Client) MoPenDeleteGroupMember(request *MoPenDeleteGroupMemberRequest) (response *MoPenDeleteGroupMemberResponse, err error) {
	response = CreateMoPenDeleteGroupMemberResponse()
	err = client.DoAction(request, response)
	return
}

// MoPenDeleteGroupMemberWithChan invokes the mopen.MoPenDeleteGroupMember API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopendeletegroupmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenDeleteGroupMemberWithChan(request *MoPenDeleteGroupMemberRequest) (<-chan *MoPenDeleteGroupMemberResponse, <-chan error) {
	responseChan := make(chan *MoPenDeleteGroupMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoPenDeleteGroupMember(request)
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

// MoPenDeleteGroupMemberWithCallback invokes the mopen.MoPenDeleteGroupMember API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopendeletegroupmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenDeleteGroupMemberWithCallback(request *MoPenDeleteGroupMemberRequest, callback func(response *MoPenDeleteGroupMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoPenDeleteGroupMemberResponse
		var err error
		defer close(result)
		response, err = client.MoPenDeleteGroupMember(request)
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

// MoPenDeleteGroupMemberRequest is the request struct for api MoPenDeleteGroupMember
type MoPenDeleteGroupMemberRequest struct {
	*requests.RpcRequest
	GroupId    requests.Integer `position:"Body" name:"GroupId"`
	DeviceName string           `position:"Body" name:"DeviceName"`
}

// MoPenDeleteGroupMemberResponse is the response struct for api MoPenDeleteGroupMember
type MoPenDeleteGroupMemberResponse struct {
	*responses.BaseResponse
	Code        string `json:"Code" xml:"Code"`
	Success     bool   `json:"Success" xml:"Success"`
	Description string `json:"Description" xml:"Description"`
	Message     string `json:"Message" xml:"Message"`
}

// CreateMoPenDeleteGroupMemberRequest creates a request to invoke MoPenDeleteGroupMember API
func CreateMoPenDeleteGroupMemberRequest() (request *MoPenDeleteGroupMemberRequest) {
	request = &MoPenDeleteGroupMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("MoPen", "2018-02-11", "MoPenDeleteGroupMember", "", "")
	return
}

// CreateMoPenDeleteGroupMemberResponse creates a response to parse from MoPenDeleteGroupMember response
func CreateMoPenDeleteGroupMemberResponse() (response *MoPenDeleteGroupMemberResponse) {
	response = &MoPenDeleteGroupMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
