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

// DeleteConference invokes the rtc.DeleteConference API synchronously
// api document: https://help.aliyun.com/api/rtc/deleteconference.html
func (client *Client) DeleteConference(request *DeleteConferenceRequest) (response *DeleteConferenceResponse, err error) {
	response = CreateDeleteConferenceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteConferenceWithChan invokes the rtc.DeleteConference API asynchronously
// api document: https://help.aliyun.com/api/rtc/deleteconference.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConferenceWithChan(request *DeleteConferenceRequest) (<-chan *DeleteConferenceResponse, <-chan error) {
	responseChan := make(chan *DeleteConferenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteConference(request)
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

// DeleteConferenceWithCallback invokes the rtc.DeleteConference API asynchronously
// api document: https://help.aliyun.com/api/rtc/deleteconference.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConferenceWithCallback(request *DeleteConferenceRequest, callback func(response *DeleteConferenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteConferenceResponse
		var err error
		defer close(result)
		response, err = client.DeleteConference(request)
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

// DeleteConferenceRequest is the request struct for api DeleteConference
type DeleteConferenceRequest struct {
	*requests.RpcRequest
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	ConferenceId string           `position:"Query" name:"ConferenceId"`
	AppId        string           `position:"Query" name:"AppId"`
}

// DeleteConferenceResponse is the response struct for api DeleteConference
type DeleteConferenceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ConferenceId string `json:"ConferenceId" xml:"ConferenceId"`
}

// CreateDeleteConferenceRequest creates a request to invoke DeleteConference API
func CreateDeleteConferenceRequest() (request *DeleteConferenceRequest) {
	request = &DeleteConferenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DeleteConference", "rtc", "openAPI")
	return
}

// CreateDeleteConferenceResponse creates a response to parse from DeleteConference response
func CreateDeleteConferenceResponse() (response *DeleteConferenceResponse) {
	response = &DeleteConferenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
