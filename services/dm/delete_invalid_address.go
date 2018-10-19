package dm

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

// DeleteInvalidAddress invokes the dm.DeleteInvalidAddress API synchronously
// api document: https://help.aliyun.com/api/dm/deleteinvalidaddress.html
func (client *Client) DeleteInvalidAddress(request *DeleteInvalidAddressRequest) (response *DeleteInvalidAddressResponse, err error) {
	response = CreateDeleteInvalidAddressResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteInvalidAddressWithChan invokes the dm.DeleteInvalidAddress API asynchronously
// api document: https://help.aliyun.com/api/dm/deleteinvalidaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteInvalidAddressWithChan(request *DeleteInvalidAddressRequest) (<-chan *DeleteInvalidAddressResponse, <-chan error) {
	responseChan := make(chan *DeleteInvalidAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteInvalidAddress(request)
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

// DeleteInvalidAddressWithCallback invokes the dm.DeleteInvalidAddress API asynchronously
// api document: https://help.aliyun.com/api/dm/deleteinvalidaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteInvalidAddressWithCallback(request *DeleteInvalidAddressRequest, callback func(response *DeleteInvalidAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteInvalidAddressResponse
		var err error
		defer close(result)
		response, err = client.DeleteInvalidAddress(request)
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

// DeleteInvalidAddressRequest is the request struct for api DeleteInvalidAddress
type DeleteInvalidAddressRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ToAddress            string           `position:"Query" name:"ToAddress"`
}

// DeleteInvalidAddressResponse is the response struct for api DeleteInvalidAddress
type DeleteInvalidAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteInvalidAddressRequest creates a request to invoke DeleteInvalidAddress API
func CreateDeleteInvalidAddressRequest() (request *DeleteInvalidAddressRequest) {
	request = &DeleteInvalidAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DeleteInvalidAddress", "", "")
	return
}

// CreateDeleteInvalidAddressResponse creates a response to parse from DeleteInvalidAddress response
func CreateDeleteInvalidAddressResponse() (response *DeleteInvalidAddressResponse) {
	response = &DeleteInvalidAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
