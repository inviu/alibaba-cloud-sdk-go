package live

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

// DeleteLiveMixNotifyConfig invokes the live.DeleteLiveMixNotifyConfig API synchronously
// api document: https://help.aliyun.com/api/live/deletelivemixnotifyconfig.html
func (client *Client) DeleteLiveMixNotifyConfig(request *DeleteLiveMixNotifyConfigRequest) (response *DeleteLiveMixNotifyConfigResponse, err error) {
	response = CreateDeleteLiveMixNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveMixNotifyConfigWithChan invokes the live.DeleteLiveMixNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deletelivemixnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveMixNotifyConfigWithChan(request *DeleteLiveMixNotifyConfigRequest) (<-chan *DeleteLiveMixNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveMixNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveMixNotifyConfig(request)
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

// DeleteLiveMixNotifyConfigWithCallback invokes the live.DeleteLiveMixNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deletelivemixnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveMixNotifyConfigWithCallback(request *DeleteLiveMixNotifyConfigRequest, callback func(response *DeleteLiveMixNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveMixNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveMixNotifyConfig(request)
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

// DeleteLiveMixNotifyConfigRequest is the request struct for api DeleteLiveMixNotifyConfig
type DeleteLiveMixNotifyConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveMixNotifyConfigResponse is the response struct for api DeleteLiveMixNotifyConfig
type DeleteLiveMixNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveMixNotifyConfigRequest creates a request to invoke DeleteLiveMixNotifyConfig API
func CreateDeleteLiveMixNotifyConfigRequest() (request *DeleteLiveMixNotifyConfigRequest) {
	request = &DeleteLiveMixNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveMixNotifyConfig", "live", "openAPI")
	return
}

// CreateDeleteLiveMixNotifyConfigResponse creates a response to parse from DeleteLiveMixNotifyConfig response
func CreateDeleteLiveMixNotifyConfigResponse() (response *DeleteLiveMixNotifyConfigResponse) {
	response = &DeleteLiveMixNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
