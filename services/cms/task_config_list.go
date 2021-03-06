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

// TaskConfigList invokes the cms.TaskConfigList API synchronously
// api document: https://help.aliyun.com/api/cms/taskconfiglist.html
func (client *Client) TaskConfigList(request *TaskConfigListRequest) (response *TaskConfigListResponse, err error) {
	response = CreateTaskConfigListResponse()
	err = client.DoAction(request, response)
	return
}

// TaskConfigListWithChan invokes the cms.TaskConfigList API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfiglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigListWithChan(request *TaskConfigListRequest) (<-chan *TaskConfigListResponse, <-chan error) {
	responseChan := make(chan *TaskConfigListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaskConfigList(request)
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

// TaskConfigListWithCallback invokes the cms.TaskConfigList API asynchronously
// api document: https://help.aliyun.com/api/cms/taskconfiglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskConfigListWithCallback(request *TaskConfigListRequest, callback func(response *TaskConfigListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaskConfigListResponse
		var err error
		defer close(result)
		response, err = client.TaskConfigList(request)
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

// TaskConfigListRequest is the request struct for api TaskConfigList
type TaskConfigListRequest struct {
	*requests.RpcRequest
	GroupId    requests.Integer `position:"Query" name:"GroupId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	TaskName   string           `position:"Query" name:"TaskName"`
	Id         requests.Integer `position:"Query" name:"Id"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// TaskConfigListResponse is the response struct for api TaskConfigList
type TaskConfigListResponse struct {
	*responses.BaseResponse
	ErrorCode    int      `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool     `json:"Success" xml:"Success"`
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	PageNumber   int      `json:"PageNumber" xml:"PageNumber"`
	PageSize     int      `json:"PageSize" xml:"PageSize"`
	PageTotal    int      `json:"PageTotal" xml:"PageTotal"`
	Total        int      `json:"Total" xml:"Total"`
	TaskList     TaskList `json:"TaskList" xml:"TaskList"`
}

// CreateTaskConfigListRequest creates a request to invoke TaskConfigList API
func CreateTaskConfigListRequest() (request *TaskConfigListRequest) {
	request = &TaskConfigListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigList", "cms", "openAPI")
	return
}

// CreateTaskConfigListResponse creates a response to parse from TaskConfigList response
func CreateTaskConfigListResponse() (response *TaskConfigListResponse) {
	response = &TaskConfigListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
