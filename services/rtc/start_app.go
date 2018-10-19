package rtc

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

// StartApp invokes the rtc.StartApp API synchronously
// api document: https://help.aliyun.com/api/rtc/startapp.html
func (client *Client) StartApp(request *StartAppRequest) (response *StartAppResponse, err error) {
	response = CreateStartAppResponse()
	err = client.DoAction(request, response)
	return
}

// StartAppWithChan invokes the rtc.StartApp API asynchronously
// api document: https://help.aliyun.com/api/rtc/startapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartAppWithChan(request *StartAppRequest) (<-chan *StartAppResponse, <-chan error) {
	responseChan := make(chan *StartAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartApp(request)
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

// StartAppWithCallback invokes the rtc.StartApp API asynchronously
// api document: https://help.aliyun.com/api/rtc/startapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartAppWithCallback(request *StartAppRequest, callback func(response *StartAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartAppResponse
		var err error
		defer close(result)
		response, err = client.StartApp(request)
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

// StartAppRequest is the request struct for api StartApp
type StartAppRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// StartAppResponse is the response struct for api StartApp
type StartAppResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartAppRequest creates a request to invoke StartApp API
func CreateStartAppRequest() (request *StartAppRequest) {
	request = &StartAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "StartApp", "rtc", "openAPI")
	return
}

// CreateStartAppResponse creates a response to parse from StartApp response
func CreateStartAppResponse() (response *StartAppResponse) {
	response = &StartAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
