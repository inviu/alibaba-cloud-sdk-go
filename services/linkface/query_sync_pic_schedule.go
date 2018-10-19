package linkface

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

// QuerySyncPicSchedule invokes the linkface.QuerySyncPicSchedule API synchronously
// api document: https://help.aliyun.com/api/linkface/querysyncpicschedule.html
func (client *Client) QuerySyncPicSchedule(request *QuerySyncPicScheduleRequest) (response *QuerySyncPicScheduleResponse, err error) {
	response = CreateQuerySyncPicScheduleResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySyncPicScheduleWithChan invokes the linkface.QuerySyncPicSchedule API asynchronously
// api document: https://help.aliyun.com/api/linkface/querysyncpicschedule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySyncPicScheduleWithChan(request *QuerySyncPicScheduleRequest) (<-chan *QuerySyncPicScheduleResponse, <-chan error) {
	responseChan := make(chan *QuerySyncPicScheduleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySyncPicSchedule(request)
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

// QuerySyncPicScheduleWithCallback invokes the linkface.QuerySyncPicSchedule API asynchronously
// api document: https://help.aliyun.com/api/linkface/querysyncpicschedule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySyncPicScheduleWithCallback(request *QuerySyncPicScheduleRequest, callback func(response *QuerySyncPicScheduleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySyncPicScheduleResponse
		var err error
		defer close(result)
		response, err = client.QuerySyncPicSchedule(request)
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

// QuerySyncPicScheduleRequest is the request struct for api QuerySyncPicSchedule
type QuerySyncPicScheduleRequest struct {
	*requests.RpcRequest
	IotId string `position:"Body" name:"IotId"`
}

// QuerySyncPicScheduleResponse is the response struct for api QuerySyncPicSchedule
type QuerySyncPicScheduleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQuerySyncPicScheduleRequest creates a request to invoke QuerySyncPicSchedule API
func CreateQuerySyncPicScheduleRequest() (request *QuerySyncPicScheduleRequest) {
	request = &QuerySyncPicScheduleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkFace", "2018-07-20", "QuerySyncPicSchedule", "", "")
	return
}

// CreateQuerySyncPicScheduleResponse creates a response to parse from QuerySyncPicSchedule response
func CreateQuerySyncPicScheduleResponse() (response *QuerySyncPicScheduleResponse) {
	response = &QuerySyncPicScheduleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
