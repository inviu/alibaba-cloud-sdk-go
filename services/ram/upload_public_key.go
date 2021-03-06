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

// UploadPublicKey invokes the ram.UploadPublicKey API synchronously
// api document: https://help.aliyun.com/api/ram/uploadpublickey.html
func (client *Client) UploadPublicKey(request *UploadPublicKeyRequest) (response *UploadPublicKeyResponse, err error) {
	response = CreateUploadPublicKeyResponse()
	err = client.DoAction(request, response)
	return
}

// UploadPublicKeyWithChan invokes the ram.UploadPublicKey API asynchronously
// api document: https://help.aliyun.com/api/ram/uploadpublickey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadPublicKeyWithChan(request *UploadPublicKeyRequest) (<-chan *UploadPublicKeyResponse, <-chan error) {
	responseChan := make(chan *UploadPublicKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadPublicKey(request)
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

// UploadPublicKeyWithCallback invokes the ram.UploadPublicKey API asynchronously
// api document: https://help.aliyun.com/api/ram/uploadpublickey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadPublicKeyWithCallback(request *UploadPublicKeyRequest, callback func(response *UploadPublicKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadPublicKeyResponse
		var err error
		defer close(result)
		response, err = client.UploadPublicKey(request)
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

// UploadPublicKeyRequest is the request struct for api UploadPublicKey
type UploadPublicKeyRequest struct {
	*requests.RpcRequest
	UserName      string `position:"Query" name:"UserName"`
	PublicKeySpec string `position:"Query" name:"PublicKeySpec"`
}

// UploadPublicKeyResponse is the response struct for api UploadPublicKey
type UploadPublicKeyResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	PublicKey PublicKey `json:"PublicKey" xml:"PublicKey"`
}

// CreateUploadPublicKeyRequest creates a request to invoke UploadPublicKey API
func CreateUploadPublicKeyRequest() (request *UploadPublicKeyRequest) {
	request = &UploadPublicKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "UploadPublicKey", "", "")
	return
}

// CreateUploadPublicKeyResponse creates a response to parse from UploadPublicKey response
func CreateUploadPublicKeyResponse() (response *UploadPublicKeyResponse) {
	response = &UploadPublicKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
