package cms

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

// StopTasks invokes the cms.StopTasks API synchronously
// api document: https://help.aliyun.com/api/cms/stoptasks.html
func (client *Client) StopTasks(request *StopTasksRequest) (response *StopTasksResponse, err error) {
	response = CreateStopTasksResponse()
	err = client.DoAction(request, response)
	return
}

// StopTasksWithChan invokes the cms.StopTasks API asynchronously
// api document: https://help.aliyun.com/api/cms/stoptasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopTasksWithChan(request *StopTasksRequest) (<-chan *StopTasksResponse, <-chan error) {
	responseChan := make(chan *StopTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopTasks(request)
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

// StopTasksWithCallback invokes the cms.StopTasks API asynchronously
// api document: https://help.aliyun.com/api/cms/stoptasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopTasksWithCallback(request *StopTasksRequest, callback func(response *StopTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopTasksResponse
		var err error
		defer close(result)
		response, err = client.StopTasks(request)
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

// StopTasksRequest is the request struct for api StopTasks
type StopTasksRequest struct {
	*requests.RpcRequest
	TaskIds string `position:"Query" name:"TaskIds"`
}

// StopTasksResponse is the response struct for api StopTasks
type StopTasksResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateStopTasksRequest creates a request to invoke StopTasks API
func CreateStopTasksRequest() (request *StopTasksRequest) {
	request = &StopTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "StopTasks", "cms", "openAPI")
	return
}

// CreateStopTasksResponse creates a response to parse from StopTasks response
func CreateStopTasksResponse() (response *StopTasksResponse) {
	response = &StopTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
