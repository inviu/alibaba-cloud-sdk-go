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

// EipFillProduct invokes the ecs.EipFillProduct API synchronously
// api document: https://help.aliyun.com/api/ecs/eipfillproduct.html
func (client *Client) EipFillProduct(request *EipFillProductRequest) (response *EipFillProductResponse, err error) {
	response = CreateEipFillProductResponse()
	err = client.DoAction(request, response)
	return
}

// EipFillProductWithChan invokes the ecs.EipFillProduct API asynchronously
// api document: https://help.aliyun.com/api/ecs/eipfillproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EipFillProductWithChan(request *EipFillProductRequest) (<-chan *EipFillProductResponse, <-chan error) {
	responseChan := make(chan *EipFillProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EipFillProduct(request)
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

// EipFillProductWithCallback invokes the ecs.EipFillProduct API asynchronously
// api document: https://help.aliyun.com/api/ecs/eipfillproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EipFillProductWithCallback(request *EipFillProductRequest, callback func(response *EipFillProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EipFillProductResponse
		var err error
		defer close(result)
		response, err = client.EipFillProduct(request)
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

// EipFillProductRequest is the request struct for api EipFillProduct
type EipFillProductRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Data                 string           `position:"Query" name:"data"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// EipFillProductResponse is the response struct for api EipFillProduct
type EipFillProductResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
	Code      string `json:"code" xml:"code"`
	Success   bool   `json:"success" xml:"success"`
	Message   string `json:"message" xml:"message"`
}

// CreateEipFillProductRequest creates a request to invoke EipFillProduct API
func CreateEipFillProductRequest() (request *EipFillProductRequest) {
	request = &EipFillProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "EipFillProduct", "ecs", "openAPI")
	return
}

// CreateEipFillProductResponse creates a response to parse from EipFillProduct response
func CreateEipFillProductResponse() (response *EipFillProductResponse) {
	response = &EipFillProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
