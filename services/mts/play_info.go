package mts

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

// PlayInfo invokes the mts.PlayInfo API synchronously
// api document: https://help.aliyun.com/api/mts/playinfo.html
func (client *Client) PlayInfo(request *PlayInfoRequest) (response *PlayInfoResponse, err error) {
	response = CreatePlayInfoResponse()
	err = client.DoAction(request, response)
	return
}

// PlayInfoWithChan invokes the mts.PlayInfo API asynchronously
// api document: https://help.aliyun.com/api/mts/playinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PlayInfoWithChan(request *PlayInfoRequest) (<-chan *PlayInfoResponse, <-chan error) {
	responseChan := make(chan *PlayInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PlayInfo(request)
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

// PlayInfoWithCallback invokes the mts.PlayInfo API asynchronously
// api document: https://help.aliyun.com/api/mts/playinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PlayInfoWithCallback(request *PlayInfoRequest, callback func(response *PlayInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PlayInfoResponse
		var err error
		defer close(result)
		response, err = client.PlayInfo(request)
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

// PlayInfoRequest is the request struct for api PlayInfo
type PlayInfoRequest struct {
	*requests.RpcRequest
	PlayDomain           string           `position:"Query" name:"PlayDomain"`
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	Formats              string           `position:"Query" name:"Formats"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	HlsUriToken          string           `position:"Query" name:"HlsUriToken"`
	Terminal             string           `position:"Query" name:"Terminal"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
	Rand                 string           `position:"Query" name:"Rand"`
	AuthTimeout          requests.Integer `position:"Query" name:"AuthTimeout"`
	AuthInfo             string           `position:"Query" name:"AuthInfo"`
}

// PlayInfoResponse is the response struct for api PlayInfo
type PlayInfoResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	NotFoundCDNDomain NotFoundCDNDomain `json:"NotFoundCDNDomain" xml:"NotFoundCDNDomain"`
	PlayInfoList      PlayInfoList      `json:"PlayInfoList" xml:"PlayInfoList"`
}

// CreatePlayInfoRequest creates a request to invoke PlayInfo API
func CreatePlayInfoRequest() (request *PlayInfoRequest) {
	request = &PlayInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "PlayInfo", "mts", "openAPI")
	return
}

// CreatePlayInfoResponse creates a response to parse from PlayInfo response
func CreatePlayInfoResponse() (response *PlayInfoResponse) {
	response = &PlayInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
