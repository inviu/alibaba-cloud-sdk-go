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

// ListPhotoProcessTasks invokes the imm.ListPhotoProcessTasks API synchronously
// api document: https://help.aliyun.com/api/imm/listphotoprocesstasks.html
func (client *Client) ListPhotoProcessTasks(request *ListPhotoProcessTasksRequest) (response *ListPhotoProcessTasksResponse, err error) {
	response = CreateListPhotoProcessTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ListPhotoProcessTasksWithChan invokes the imm.ListPhotoProcessTasks API asynchronously
// api document: https://help.aliyun.com/api/imm/listphotoprocesstasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPhotoProcessTasksWithChan(request *ListPhotoProcessTasksRequest) (<-chan *ListPhotoProcessTasksResponse, <-chan error) {
	responseChan := make(chan *ListPhotoProcessTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhotoProcessTasks(request)
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

// ListPhotoProcessTasksWithCallback invokes the imm.ListPhotoProcessTasks API asynchronously
// api document: https://help.aliyun.com/api/imm/listphotoprocesstasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPhotoProcessTasksWithCallback(request *ListPhotoProcessTasksRequest, callback func(response *ListPhotoProcessTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhotoProcessTasksResponse
		var err error
		defer close(result)
		response, err = client.ListPhotoProcessTasks(request)
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

// ListPhotoProcessTasksRequest is the request struct for api ListPhotoProcessTasks
type ListPhotoProcessTasksRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Marker  string           `position:"Query" name:"Marker"`
	Project string           `position:"Query" name:"Project"`
}

// ListPhotoProcessTasksResponse is the response struct for api ListPhotoProcessTasks
type ListPhotoProcessTasksResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	NextMarker string      `json:"NextMarker" xml:"NextMarker"`
	Tasks      []TasksItem `json:"Tasks" xml:"Tasks"`
}

// CreateListPhotoProcessTasksRequest creates a request to invoke ListPhotoProcessTasks API
func CreateListPhotoProcessTasksRequest() (request *ListPhotoProcessTasksRequest) {
	request = &ListPhotoProcessTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListPhotoProcessTasks", "imm", "openAPI")
	return
}

// CreateListPhotoProcessTasksResponse creates a response to parse from ListPhotoProcessTasks response
func CreateListPhotoProcessTasksResponse() (response *ListPhotoProcessTasksResponse) {
	response = &ListPhotoProcessTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
