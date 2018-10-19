package ehpc

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

// ModifyImageGatewayConfig invokes the ehpc.ModifyImageGatewayConfig API synchronously
// api document: https://help.aliyun.com/api/ehpc/modifyimagegatewayconfig.html
func (client *Client) ModifyImageGatewayConfig(request *ModifyImageGatewayConfigRequest) (response *ModifyImageGatewayConfigResponse, err error) {
	response = CreateModifyImageGatewayConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyImageGatewayConfigWithChan invokes the ehpc.ModifyImageGatewayConfig API asynchronously
// api document: https://help.aliyun.com/api/ehpc/modifyimagegatewayconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyImageGatewayConfigWithChan(request *ModifyImageGatewayConfigRequest) (<-chan *ModifyImageGatewayConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyImageGatewayConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyImageGatewayConfig(request)
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

// ModifyImageGatewayConfigWithCallback invokes the ehpc.ModifyImageGatewayConfig API asynchronously
// api document: https://help.aliyun.com/api/ehpc/modifyimagegatewayconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyImageGatewayConfigWithCallback(request *ModifyImageGatewayConfigRequest, callback func(response *ModifyImageGatewayConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyImageGatewayConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyImageGatewayConfig(request)
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

// ModifyImageGatewayConfigRequest is the request struct for api ModifyImageGatewayConfig
type ModifyImageGatewayConfigRequest struct {
	*requests.RpcRequest
	DefaultRepoLocation    string                          `position:"Query" name:"DefaultRepoLocation"`
	DBPassword             string                          `position:"Query" name:"DBPassword"`
	Repo                   *[]ModifyImageGatewayConfigRepo `position:"Query" name:"Repo"  type:"Repeated"`
	DBType                 string                          `position:"Query" name:"DBType"`
	DBUsername             string                          `position:"Query" name:"DBUsername"`
	DBServerInfo           string                          `position:"Query" name:"DBServerInfo"`
	PullUpdateTimeout      requests.Integer                `position:"Query" name:"PullUpdateTimeout"`
	ClusterId              string                          `position:"Query" name:"ClusterId"`
	ImageExpirationTimeout string                          `position:"Query" name:"ImageExpirationTimeout"`
}

// ModifyImageGatewayConfigRepo is a repeated param struct in ModifyImageGatewayConfigRequest
type ModifyImageGatewayConfigRepo struct {
	Auth     string `name:"Auth"`
	Location string `name:"Location"`
	URL      string `name:"URL"`
}

// ModifyImageGatewayConfigResponse is the response struct for api ModifyImageGatewayConfig
type ModifyImageGatewayConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyImageGatewayConfigRequest creates a request to invoke ModifyImageGatewayConfig API
func CreateModifyImageGatewayConfigRequest() (request *ModifyImageGatewayConfigRequest) {
	request = &ModifyImageGatewayConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ModifyImageGatewayConfig", "ehs", "openAPI")
	return
}

// CreateModifyImageGatewayConfigResponse creates a response to parse from ModifyImageGatewayConfig response
func CreateModifyImageGatewayConfigResponse() (response *ModifyImageGatewayConfigResponse) {
	response = &ModifyImageGatewayConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
