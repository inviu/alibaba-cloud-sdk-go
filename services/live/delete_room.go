package live

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

// DeleteRoom invokes the live.DeleteRoom API synchronously
// api document: https://help.aliyun.com/api/live/deleteroom.html
func (client *Client) DeleteRoom(request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
	response = CreateDeleteRoomResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRoomWithChan invokes the live.DeleteRoom API asynchronously
// api document: https://help.aliyun.com/api/live/deleteroom.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRoomWithChan(request *DeleteRoomRequest) (<-chan *DeleteRoomResponse, <-chan error) {
	responseChan := make(chan *DeleteRoomResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRoom(request)
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

// DeleteRoomWithCallback invokes the live.DeleteRoom API asynchronously
// api document: https://help.aliyun.com/api/live/deleteroom.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRoomWithCallback(request *DeleteRoomRequest, callback func(response *DeleteRoomResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRoomResponse
		var err error
		defer close(result)
		response, err = client.DeleteRoom(request)
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

// DeleteRoomRequest is the request struct for api DeleteRoom
type DeleteRoomRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	RoomId  string           `position:"Query" name:"RoomId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// DeleteRoomResponse is the response struct for api DeleteRoom
type DeleteRoomResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRoomRequest creates a request to invoke DeleteRoom API
func CreateDeleteRoomRequest() (request *DeleteRoomRequest) {
	request = &DeleteRoomRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteRoom", "live", "openAPI")
	return
}

// CreateDeleteRoomResponse creates a response to parse from DeleteRoom response
func CreateDeleteRoomResponse() (response *DeleteRoomResponse) {
	response = &DeleteRoomResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
