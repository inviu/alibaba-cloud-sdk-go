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

// UpdateLiveMixNotifyConfig invokes the live.UpdateLiveMixNotifyConfig API synchronously
// api document: https://help.aliyun.com/api/live/updatelivemixnotifyconfig.html
func (client *Client) UpdateLiveMixNotifyConfig(request *UpdateLiveMixNotifyConfigRequest) (response *UpdateLiveMixNotifyConfigResponse, err error) {
	response = CreateUpdateLiveMixNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveMixNotifyConfigWithChan invokes the live.UpdateLiveMixNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/updatelivemixnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLiveMixNotifyConfigWithChan(request *UpdateLiveMixNotifyConfigRequest) (<-chan *UpdateLiveMixNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveMixNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveMixNotifyConfig(request)
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

// UpdateLiveMixNotifyConfigWithCallback invokes the live.UpdateLiveMixNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/updatelivemixnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLiveMixNotifyConfigWithCallback(request *UpdateLiveMixNotifyConfigRequest, callback func(response *UpdateLiveMixNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveMixNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveMixNotifyConfig(request)
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

// UpdateLiveMixNotifyConfigRequest is the request struct for api UpdateLiveMixNotifyConfig
type UpdateLiveMixNotifyConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	NotifyUrl     string           `position:"Query" name:"NotifyUrl"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateLiveMixNotifyConfigResponse is the response struct for api UpdateLiveMixNotifyConfig
type UpdateLiveMixNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveMixNotifyConfigRequest creates a request to invoke UpdateLiveMixNotifyConfig API
func CreateUpdateLiveMixNotifyConfigRequest() (request *UpdateLiveMixNotifyConfigRequest) {
	request = &UpdateLiveMixNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLiveMixNotifyConfig", "live", "openAPI")
	return
}

// CreateUpdateLiveMixNotifyConfigResponse creates a response to parse from UpdateLiveMixNotifyConfig response
func CreateUpdateLiveMixNotifyConfigResponse() (response *UpdateLiveMixNotifyConfigResponse) {
	response = &UpdateLiveMixNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
