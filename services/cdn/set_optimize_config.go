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

// SetOptimizeConfig invokes the cdn.SetOptimizeConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setoptimizeconfig.html
func (client *Client) SetOptimizeConfig(request *SetOptimizeConfigRequest) (response *SetOptimizeConfigResponse, err error) {
	response = CreateSetOptimizeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetOptimizeConfigWithChan invokes the cdn.SetOptimizeConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setoptimizeconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetOptimizeConfigWithChan(request *SetOptimizeConfigRequest) (<-chan *SetOptimizeConfigResponse, <-chan error) {
	responseChan := make(chan *SetOptimizeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetOptimizeConfig(request)
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

// SetOptimizeConfigWithCallback invokes the cdn.SetOptimizeConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setoptimizeconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetOptimizeConfigWithCallback(request *SetOptimizeConfigRequest, callback func(response *SetOptimizeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetOptimizeConfigResponse
		var err error
		defer close(result)
		response, err = client.SetOptimizeConfig(request)
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

// SetOptimizeConfigRequest is the request struct for api SetOptimizeConfig
type SetOptimizeConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	Enable        string           `position:"Query" name:"Enable"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// SetOptimizeConfigResponse is the response struct for api SetOptimizeConfig
type SetOptimizeConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetOptimizeConfigRequest creates a request to invoke SetOptimizeConfig API
func CreateSetOptimizeConfigRequest() (request *SetOptimizeConfigRequest) {
	request = &SetOptimizeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetOptimizeConfig", "", "")
	return
}

// CreateSetOptimizeConfigResponse creates a response to parse from SetOptimizeConfig response
func CreateSetOptimizeConfigResponse() (response *SetOptimizeConfigResponse) {
	response = &SetOptimizeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
