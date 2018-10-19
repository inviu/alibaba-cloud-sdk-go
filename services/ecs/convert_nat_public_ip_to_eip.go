package ecs

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

// ConvertNatPublicIpToEip invokes the ecs.ConvertNatPublicIpToEip API synchronously
// api document: https://help.aliyun.com/api/ecs/convertnatpubliciptoeip.html
func (client *Client) ConvertNatPublicIpToEip(request *ConvertNatPublicIpToEipRequest) (response *ConvertNatPublicIpToEipResponse, err error) {
	response = CreateConvertNatPublicIpToEipResponse()
	err = client.DoAction(request, response)
	return
}

// ConvertNatPublicIpToEipWithChan invokes the ecs.ConvertNatPublicIpToEip API asynchronously
// api document: https://help.aliyun.com/api/ecs/convertnatpubliciptoeip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConvertNatPublicIpToEipWithChan(request *ConvertNatPublicIpToEipRequest) (<-chan *ConvertNatPublicIpToEipResponse, <-chan error) {
	responseChan := make(chan *ConvertNatPublicIpToEipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConvertNatPublicIpToEip(request)
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

// ConvertNatPublicIpToEipWithCallback invokes the ecs.ConvertNatPublicIpToEip API asynchronously
// api document: https://help.aliyun.com/api/ecs/convertnatpubliciptoeip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConvertNatPublicIpToEipWithCallback(request *ConvertNatPublicIpToEipRequest, callback func(response *ConvertNatPublicIpToEipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConvertNatPublicIpToEipResponse
		var err error
		defer close(result)
		response, err = client.ConvertNatPublicIpToEip(request)
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

// ConvertNatPublicIpToEipRequest is the request struct for api ConvertNatPublicIpToEip
type ConvertNatPublicIpToEipRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ConvertNatPublicIpToEipResponse is the response struct for api ConvertNatPublicIpToEip
type ConvertNatPublicIpToEipResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConvertNatPublicIpToEipRequest creates a request to invoke ConvertNatPublicIpToEip API
func CreateConvertNatPublicIpToEipRequest() (request *ConvertNatPublicIpToEipRequest) {
	request = &ConvertNatPublicIpToEipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ConvertNatPublicIpToEip", "ecs", "openAPI")
	return
}

// CreateConvertNatPublicIpToEipResponse creates a response to parse from ConvertNatPublicIpToEip response
func CreateConvertNatPublicIpToEipResponse() (response *ConvertNatPublicIpToEipResponse) {
	response = &ConvertNatPublicIpToEipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
