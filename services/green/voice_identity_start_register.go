package green

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

// VoiceIdentityStartRegister invokes the green.VoiceIdentityStartRegister API synchronously
// api document: https://help.aliyun.com/api/green/voiceidentitystartregister.html
func (client *Client) VoiceIdentityStartRegister(request *VoiceIdentityStartRegisterRequest) (response *VoiceIdentityStartRegisterResponse, err error) {
	response = CreateVoiceIdentityStartRegisterResponse()
	err = client.DoAction(request, response)
	return
}

// VoiceIdentityStartRegisterWithChan invokes the green.VoiceIdentityStartRegister API asynchronously
// api document: https://help.aliyun.com/api/green/voiceidentitystartregister.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VoiceIdentityStartRegisterWithChan(request *VoiceIdentityStartRegisterRequest) (<-chan *VoiceIdentityStartRegisterResponse, <-chan error) {
	responseChan := make(chan *VoiceIdentityStartRegisterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VoiceIdentityStartRegister(request)
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

// VoiceIdentityStartRegisterWithCallback invokes the green.VoiceIdentityStartRegister API asynchronously
// api document: https://help.aliyun.com/api/green/voiceidentitystartregister.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VoiceIdentityStartRegisterWithCallback(request *VoiceIdentityStartRegisterRequest, callback func(response *VoiceIdentityStartRegisterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VoiceIdentityStartRegisterResponse
		var err error
		defer close(result)
		response, err = client.VoiceIdentityStartRegister(request)
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

// VoiceIdentityStartRegisterRequest is the request struct for api VoiceIdentityStartRegister
type VoiceIdentityStartRegisterRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// VoiceIdentityStartRegisterResponse is the response struct for api VoiceIdentityStartRegister
type VoiceIdentityStartRegisterResponse struct {
	*responses.BaseResponse
}

// CreateVoiceIdentityStartRegisterRequest creates a request to invoke VoiceIdentityStartRegister API
func CreateVoiceIdentityStartRegisterRequest() (request *VoiceIdentityStartRegisterRequest) {
	request = &VoiceIdentityStartRegisterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "VoiceIdentityStartRegister", "/green/voice/auth/start/register", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVoiceIdentityStartRegisterResponse creates a response to parse from VoiceIdentityStartRegister response
func CreateVoiceIdentityStartRegisterResponse() (response *VoiceIdentityStartRegisterResponse) {
	response = &VoiceIdentityStartRegisterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
