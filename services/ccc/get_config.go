package ccc

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

// GetConfig invokes the ccc.GetConfig API synchronously
// api document: https://help.aliyun.com/api/ccc/getconfig.html
func (client *Client) GetConfig(request *GetConfigRequest) (response *GetConfigResponse, err error) {
	response = CreateGetConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetConfigWithChan invokes the ccc.GetConfig API asynchronously
// api document: https://help.aliyun.com/api/ccc/getconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConfigWithChan(request *GetConfigRequest) (<-chan *GetConfigResponse, <-chan error) {
	responseChan := make(chan *GetConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetConfig(request)
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

// GetConfigWithCallback invokes the ccc.GetConfig API asynchronously
// api document: https://help.aliyun.com/api/ccc/getconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConfigWithCallback(request *GetConfigRequest, callback func(response *GetConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetConfigResponse
		var err error
		defer close(result)
		response, err = client.GetConfig(request)
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

// GetConfigRequest is the request struct for api GetConfig
type GetConfigRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	ObjectId   string `position:"Query" name:"ObjectId"`
	ObjectType string `position:"Query" name:"ObjectType"`
}

// GetConfigResponse is the response struct for api GetConfig
type GetConfigResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ConfigItem     ConfigItem `json:"ConfigItem" xml:"ConfigItem"`
}

// CreateGetConfigRequest creates a request to invoke GetConfig API
func CreateGetConfigRequest() (request *GetConfigRequest) {
	request = &GetConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetConfig", "", "")
	return
}

// CreateGetConfigResponse creates a response to parse from GetConfig response
func CreateGetConfigResponse() (response *GetConfigResponse) {
	response = &GetConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
