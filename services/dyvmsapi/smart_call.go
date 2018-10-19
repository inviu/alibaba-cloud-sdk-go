package dyvmsapi

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

// SmartCall invokes the dyvmsapi.SmartCall API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/smartcall.html
func (client *Client) SmartCall(request *SmartCallRequest) (response *SmartCallResponse, err error) {
	response = CreateSmartCallResponse()
	err = client.DoAction(request, response)
	return
}

// SmartCallWithChan invokes the dyvmsapi.SmartCall API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/smartcall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SmartCallWithChan(request *SmartCallRequest) (<-chan *SmartCallResponse, <-chan error) {
	responseChan := make(chan *SmartCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SmartCall(request)
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

// SmartCallWithCallback invokes the dyvmsapi.SmartCall API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/smartcall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SmartCallWithCallback(request *SmartCallRequest, callback func(response *SmartCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SmartCallResponse
		var err error
		defer close(result)
		response, err = client.SmartCall(request)
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

// SmartCallRequest is the request struct for api SmartCall
type SmartCallRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CalledShowNumber     string           `position:"Query" name:"CalledShowNumber"`
	CalledNumber         string           `position:"Query" name:"CalledNumber"`
	VoiceCode            string           `position:"Query" name:"VoiceCode"`
	RecordFlag           requests.Boolean `position:"Query" name:"RecordFlag"`
	Volume               requests.Integer `position:"Query" name:"Volume"`
	Speed                requests.Integer `position:"Query" name:"Speed"`
	AsrModelId           string           `position:"Query" name:"AsrModelId"`
	PauseTime            requests.Integer `position:"Query" name:"PauseTime"`
	MuteTime             requests.Integer `position:"Query" name:"MuteTime"`
	ActionCodeBreak      requests.Boolean `position:"Query" name:"ActionCodeBreak"`
	OutId                string           `position:"Query" name:"OutId"`
	DynamicId            string           `position:"Query" name:"DynamicId"`
}

// SmartCallResponse is the response struct for api SmartCall
type SmartCallResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CallId    string `json:"CallId" xml:"CallId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSmartCallRequest creates a request to invoke SmartCall API
func CreateSmartCallRequest() (request *SmartCallRequest) {
	request = &SmartCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "SmartCall", "", "")
	return
}

// CreateSmartCallResponse creates a response to parse from SmartCall response
func CreateSmartCallResponse() (response *SmartCallResponse) {
	response = &SmartCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
