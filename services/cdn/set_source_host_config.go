package cdn

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

// SetSourceHostConfig invokes the cdn.SetSourceHostConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setsourcehostconfig.html
func (client *Client) SetSourceHostConfig(request *SetSourceHostConfigRequest) (response *SetSourceHostConfigResponse, err error) {
	response = CreateSetSourceHostConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetSourceHostConfigWithChan invokes the cdn.SetSourceHostConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setsourcehostconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetSourceHostConfigWithChan(request *SetSourceHostConfigRequest) (<-chan *SetSourceHostConfigResponse, <-chan error) {
	responseChan := make(chan *SetSourceHostConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetSourceHostConfig(request)
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

// SetSourceHostConfigWithCallback invokes the cdn.SetSourceHostConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setsourcehostconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetSourceHostConfigWithCallback(request *SetSourceHostConfigRequest, callback func(response *SetSourceHostConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetSourceHostConfigResponse
		var err error
		defer close(result)
		response, err = client.SetSourceHostConfig(request)
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

// SetSourceHostConfigRequest is the request struct for api SetSourceHostConfig
type SetSourceHostConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	Enable        string           `position:"Query" name:"Enable"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	BackSrcDomain string           `position:"Query" name:"BackSrcDomain"`
}

// SetSourceHostConfigResponse is the response struct for api SetSourceHostConfig
type SetSourceHostConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetSourceHostConfigRequest creates a request to invoke SetSourceHostConfig API
func CreateSetSourceHostConfigRequest() (request *SetSourceHostConfigRequest) {
	request = &SetSourceHostConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetSourceHostConfig", "", "")
	return
}

// CreateSetSourceHostConfigResponse creates a response to parse from SetSourceHostConfig response
func CreateSetSourceHostConfigResponse() (response *SetSourceHostConfigResponse) {
	response = &SetSourceHostConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
