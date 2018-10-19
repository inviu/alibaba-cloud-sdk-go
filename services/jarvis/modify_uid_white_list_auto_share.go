package jarvis

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

// ModifyUidWhiteListAutoShare invokes the jarvis.ModifyUidWhiteListAutoShare API synchronously
// api document: https://help.aliyun.com/api/jarvis/modifyuidwhitelistautoshare.html
func (client *Client) ModifyUidWhiteListAutoShare(request *ModifyUidWhiteListAutoShareRequest) (response *ModifyUidWhiteListAutoShareResponse, err error) {
	response = CreateModifyUidWhiteListAutoShareResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUidWhiteListAutoShareWithChan invokes the jarvis.ModifyUidWhiteListAutoShare API asynchronously
// api document: https://help.aliyun.com/api/jarvis/modifyuidwhitelistautoshare.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUidWhiteListAutoShareWithChan(request *ModifyUidWhiteListAutoShareRequest) (<-chan *ModifyUidWhiteListAutoShareResponse, <-chan error) {
	responseChan := make(chan *ModifyUidWhiteListAutoShareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUidWhiteListAutoShare(request)
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

// ModifyUidWhiteListAutoShareWithCallback invokes the jarvis.ModifyUidWhiteListAutoShare API asynchronously
// api document: https://help.aliyun.com/api/jarvis/modifyuidwhitelistautoshare.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUidWhiteListAutoShareWithCallback(request *ModifyUidWhiteListAutoShareRequest, callback func(response *ModifyUidWhiteListAutoShareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUidWhiteListAutoShareResponse
		var err error
		defer close(result)
		response, err = client.ModifyUidWhiteListAutoShare(request)
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

// ModifyUidWhiteListAutoShareRequest is the request struct for api ModifyUidWhiteListAutoShare
type ModifyUidWhiteListAutoShareRequest struct {
	*requests.RpcRequest
	SourceIp      string           `position:"Query" name:"SourceIp"`
	AutoConfig    requests.Integer `position:"Query" name:"AutoConfig"`
	ProductName   string           `position:"Query" name:"ProductName"`
	WhiteListType requests.Integer `position:"Query" name:"WhiteListType"`
	Lang          string           `position:"Query" name:"Lang"`
	SrcUid        string           `position:"Query" name:"SrcUid"`
	SourceCode    string           `position:"Query" name:"SourceCode"`
}

// ModifyUidWhiteListAutoShareResponse is the response struct for api ModifyUidWhiteListAutoShare
type ModifyUidWhiteListAutoShareResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateModifyUidWhiteListAutoShareRequest creates a request to invoke ModifyUidWhiteListAutoShare API
func CreateModifyUidWhiteListAutoShareRequest() (request *ModifyUidWhiteListAutoShareRequest) {
	request = &ModifyUidWhiteListAutoShareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "ModifyUidWhiteListAutoShare", "", "")
	return
}

// CreateModifyUidWhiteListAutoShareResponse creates a response to parse from ModifyUidWhiteListAutoShare response
func CreateModifyUidWhiteListAutoShareResponse() (response *ModifyUidWhiteListAutoShareResponse) {
	response = &ModifyUidWhiteListAutoShareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
