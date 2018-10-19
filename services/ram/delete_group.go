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

// DeleteGroup invokes the ram.DeleteGroup API synchronously
// api document: https://help.aliyun.com/api/ram/deletegroup.html
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
	response = CreateDeleteGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGroupWithChan invokes the ram.DeleteGroup API asynchronously
// api document: https://help.aliyun.com/api/ram/deletegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteGroupWithChan(request *DeleteGroupRequest) (<-chan *DeleteGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGroup(request)
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

// DeleteGroupWithCallback invokes the ram.DeleteGroup API asynchronously
// api document: https://help.aliyun.com/api/ram/deletegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteGroupWithCallback(request *DeleteGroupRequest, callback func(response *DeleteGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteGroup(request)
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

// DeleteGroupRequest is the request struct for api DeleteGroup
type DeleteGroupRequest struct {
	*requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

// DeleteGroupResponse is the response struct for api DeleteGroup
type DeleteGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteGroupRequest creates a request to invoke DeleteGroup API
func CreateDeleteGroupRequest() (request *DeleteGroupRequest) {
	request = &DeleteGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeleteGroup", "", "")
	return
}

// CreateDeleteGroupResponse creates a response to parse from DeleteGroup response
func CreateDeleteGroupResponse() (response *DeleteGroupResponse) {
	response = &DeleteGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
