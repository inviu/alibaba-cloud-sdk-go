package r_kvstore

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

// SwitchTempInstance invokes the r_kvstore.SwitchTempInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/switchtempinstance.html
func (client *Client) SwitchTempInstance(request *SwitchTempInstanceRequest) (response *SwitchTempInstanceResponse, err error) {
	response = CreateSwitchTempInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchTempInstanceWithChan invokes the r_kvstore.SwitchTempInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/switchtempinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchTempInstanceWithChan(request *SwitchTempInstanceRequest) (<-chan *SwitchTempInstanceResponse, <-chan error) {
	responseChan := make(chan *SwitchTempInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchTempInstance(request)
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

// SwitchTempInstanceWithCallback invokes the r_kvstore.SwitchTempInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/switchtempinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchTempInstanceWithCallback(request *SwitchTempInstanceRequest, callback func(response *SwitchTempInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchTempInstanceResponse
		var err error
		defer close(result)
		response, err = client.SwitchTempInstance(request)
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

// SwitchTempInstanceRequest is the request struct for api SwitchTempInstance
type SwitchTempInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// SwitchTempInstanceResponse is the response struct for api SwitchTempInstance
type SwitchTempInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchTempInstanceRequest creates a request to invoke SwitchTempInstance API
func CreateSwitchTempInstanceRequest() (request *SwitchTempInstanceRequest) {
	request = &SwitchTempInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchTempInstance", "redisa", "openAPI")
	return
}

// CreateSwitchTempInstanceResponse creates a response to parse from SwitchTempInstance response
func CreateSwitchTempInstanceResponse() (response *SwitchTempInstanceResponse) {
	response = &SwitchTempInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
