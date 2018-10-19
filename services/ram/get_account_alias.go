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

// GetAccountAlias invokes the ram.GetAccountAlias API synchronously
// api document: https://help.aliyun.com/api/ram/getaccountalias.html
func (client *Client) GetAccountAlias(request *GetAccountAliasRequest) (response *GetAccountAliasResponse, err error) {
	response = CreateGetAccountAliasResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccountAliasWithChan invokes the ram.GetAccountAlias API asynchronously
// api document: https://help.aliyun.com/api/ram/getaccountalias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccountAliasWithChan(request *GetAccountAliasRequest) (<-chan *GetAccountAliasResponse, <-chan error) {
	responseChan := make(chan *GetAccountAliasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountAlias(request)
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

// GetAccountAliasWithCallback invokes the ram.GetAccountAlias API asynchronously
// api document: https://help.aliyun.com/api/ram/getaccountalias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccountAliasWithCallback(request *GetAccountAliasRequest, callback func(response *GetAccountAliasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccountAliasResponse
		var err error
		defer close(result)
		response, err = client.GetAccountAlias(request)
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

// GetAccountAliasRequest is the request struct for api GetAccountAlias
type GetAccountAliasRequest struct {
	*requests.RpcRequest
}

// GetAccountAliasResponse is the response struct for api GetAccountAlias
type GetAccountAliasResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	AccountAlias string `json:"AccountAlias" xml:"AccountAlias"`
}

// CreateGetAccountAliasRequest creates a request to invoke GetAccountAlias API
func CreateGetAccountAliasRequest() (request *GetAccountAliasRequest) {
	request = &GetAccountAliasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "GetAccountAlias", "", "")
	return
}

// CreateGetAccountAliasResponse creates a response to parse from GetAccountAlias response
func CreateGetAccountAliasResponse() (response *GetAccountAliasResponse) {
	response = &GetAccountAliasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
