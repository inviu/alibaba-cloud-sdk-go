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

// ModifyTask invokes the cms.ModifyTask API synchronously
// api document: https://help.aliyun.com/api/cms/modifytask.html
func (client *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
	response = CreateModifyTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTaskWithChan invokes the cms.ModifyTask API asynchronously
// api document: https://help.aliyun.com/api/cms/modifytask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTaskWithChan(request *ModifyTaskRequest) (<-chan *ModifyTaskResponse, <-chan error) {
	responseChan := make(chan *ModifyTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTask(request)
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

// ModifyTaskWithCallback invokes the cms.ModifyTask API asynchronously
// api document: https://help.aliyun.com/api/cms/modifytask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTaskWithCallback(request *ModifyTaskRequest, callback func(response *ModifyTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTaskResponse
		var err error
		defer close(result)
		response, err = client.ModifyTask(request)
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

// ModifyTaskRequest is the request struct for api ModifyTask
type ModifyTaskRequest struct {
	*requests.RpcRequest
	Address   string `position:"Query" name:"Address"`
	IspCity   string `position:"Query" name:"IspCity"`
	AlertIds  string `position:"Query" name:"AlertIds"`
	Options   string `position:"Query" name:"Options"`
	TaskName  string `position:"Query" name:"TaskName"`
	Interval  string `position:"Query" name:"Interval"`
	AlertRule string `position:"Query" name:"AlertRule"`
	TaskId    string `position:"Query" name:"TaskId"`
}

// ModifyTaskResponse is the response struct for api ModifyTask
type ModifyTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateModifyTaskRequest creates a request to invoke ModifyTask API
func CreateModifyTaskRequest() (request *ModifyTaskRequest) {
	request = &ModifyTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "ModifyTask", "cms", "openAPI")
	return
}

// CreateModifyTaskResponse creates a response to parse from ModifyTask response
func CreateModifyTaskResponse() (response *ModifyTaskResponse) {
	response = &ModifyTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
