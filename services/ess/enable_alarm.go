package ess

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

// EnableAlarm invokes the ess.EnableAlarm API synchronously
// api document: https://help.aliyun.com/api/ess/enablealarm.html
func (client *Client) EnableAlarm(request *EnableAlarmRequest) (response *EnableAlarmResponse, err error) {
	response = CreateEnableAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// EnableAlarmWithChan invokes the ess.EnableAlarm API asynchronously
// api document: https://help.aliyun.com/api/ess/enablealarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableAlarmWithChan(request *EnableAlarmRequest) (<-chan *EnableAlarmResponse, <-chan error) {
	responseChan := make(chan *EnableAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAlarm(request)
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

// EnableAlarmWithCallback invokes the ess.EnableAlarm API asynchronously
// api document: https://help.aliyun.com/api/ess/enablealarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableAlarmWithCallback(request *EnableAlarmRequest, callback func(response *EnableAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAlarmResponse
		var err error
		defer close(result)
		response, err = client.EnableAlarm(request)
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

// EnableAlarmRequest is the request struct for api EnableAlarm
type EnableAlarmRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AlarmTaskId          string           `position:"Query" name:"AlarmTaskId"`
}

// EnableAlarmResponse is the response struct for api EnableAlarm
type EnableAlarmResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableAlarmRequest creates a request to invoke EnableAlarm API
func CreateEnableAlarmRequest() (request *EnableAlarmRequest) {
	request = &EnableAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "EnableAlarm", "ess", "openAPI")
	return
}

// CreateEnableAlarmResponse creates a response to parse from EnableAlarm response
func CreateEnableAlarmResponse() (response *EnableAlarmResponse) {
	response = &EnableAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
