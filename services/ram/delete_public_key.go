package ram

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

// DeletePublicKey invokes the ram.DeletePublicKey API synchronously
// api document: https://help.aliyun.com/api/ram/deletepublickey.html
func (client *Client) DeletePublicKey(request *DeletePublicKeyRequest) (response *DeletePublicKeyResponse, err error) {
	response = CreateDeletePublicKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePublicKeyWithChan invokes the ram.DeletePublicKey API asynchronously
// api document: https://help.aliyun.com/api/ram/deletepublickey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePublicKeyWithChan(request *DeletePublicKeyRequest) (<-chan *DeletePublicKeyResponse, <-chan error) {
	responseChan := make(chan *DeletePublicKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePublicKey(request)
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

// DeletePublicKeyWithCallback invokes the ram.DeletePublicKey API asynchronously
// api document: https://help.aliyun.com/api/ram/deletepublickey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePublicKeyWithCallback(request *DeletePublicKeyRequest, callback func(response *DeletePublicKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePublicKeyResponse
		var err error
		defer close(result)
		response, err = client.DeletePublicKey(request)
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

// DeletePublicKeyRequest is the request struct for api DeletePublicKey
type DeletePublicKeyRequest struct {
	*requests.RpcRequest
	UserName        string `position:"Query" name:"UserName"`
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
}

// DeletePublicKeyResponse is the response struct for api DeletePublicKey
type DeletePublicKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePublicKeyRequest creates a request to invoke DeletePublicKey API
func CreateDeletePublicKeyRequest() (request *DeletePublicKeyRequest) {
	request = &DeletePublicKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeletePublicKey", "", "")
	return
}

// CreateDeletePublicKeyResponse creates a response to parse from DeletePublicKey response
func CreateDeletePublicKeyResponse() (response *DeletePublicKeyResponse) {
	response = &DeletePublicKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
