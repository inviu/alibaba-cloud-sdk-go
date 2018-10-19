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

// PutEvent invokes the cms.PutEvent API synchronously
// api document: https://help.aliyun.com/api/cms/putevent.html
func (client *Client) PutEvent(request *PutEventRequest) (response *PutEventResponse, err error) {
	response = CreatePutEventResponse()
	err = client.DoAction(request, response)
	return
}

// PutEventWithChan invokes the cms.PutEvent API asynchronously
// api document: https://help.aliyun.com/api/cms/putevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutEventWithChan(request *PutEventRequest) (<-chan *PutEventResponse, <-chan error) {
	responseChan := make(chan *PutEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutEvent(request)
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

// PutEventWithCallback invokes the cms.PutEvent API asynchronously
// api document: https://help.aliyun.com/api/cms/putevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutEventWithCallback(request *PutEventRequest, callback func(response *PutEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutEventResponse
		var err error
		defer close(result)
		response, err = client.PutEvent(request)
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

// PutEventRequest is the request struct for api PutEvent
type PutEventRequest struct {
	*requests.RpcRequest
	EventInfo string `position:"Query" name:"EventInfo"`
}

// PutEventResponse is the response struct for api PutEvent
type PutEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreatePutEventRequest creates a request to invoke PutEvent API
func CreatePutEventRequest() (request *PutEventRequest) {
	request = &PutEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "PutEvent", "cms", "openAPI")
	return
}

// CreatePutEventResponse creates a response to parse from PutEvent response
func CreatePutEventResponse() (response *PutEventResponse) {
	response = &PutEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
