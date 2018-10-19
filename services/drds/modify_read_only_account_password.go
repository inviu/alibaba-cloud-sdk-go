package drds

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

// ModifyReadOnlyAccountPassword invokes the drds.ModifyReadOnlyAccountPassword API synchronously
// api document: https://help.aliyun.com/api/drds/modifyreadonlyaccountpassword.html
func (client *Client) ModifyReadOnlyAccountPassword(request *ModifyReadOnlyAccountPasswordRequest) (response *ModifyReadOnlyAccountPasswordResponse, err error) {
	response = CreateModifyReadOnlyAccountPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReadOnlyAccountPasswordWithChan invokes the drds.ModifyReadOnlyAccountPassword API asynchronously
// api document: https://help.aliyun.com/api/drds/modifyreadonlyaccountpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReadOnlyAccountPasswordWithChan(request *ModifyReadOnlyAccountPasswordRequest) (<-chan *ModifyReadOnlyAccountPasswordResponse, <-chan error) {
	responseChan := make(chan *ModifyReadOnlyAccountPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReadOnlyAccountPassword(request)
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

// ModifyReadOnlyAccountPasswordWithCallback invokes the drds.ModifyReadOnlyAccountPassword API asynchronously
// api document: https://help.aliyun.com/api/drds/modifyreadonlyaccountpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReadOnlyAccountPasswordWithCallback(request *ModifyReadOnlyAccountPasswordRequest, callback func(response *ModifyReadOnlyAccountPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReadOnlyAccountPasswordResponse
		var err error
		defer close(result)
		response, err = client.ModifyReadOnlyAccountPassword(request)
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

// ModifyReadOnlyAccountPasswordRequest is the request struct for api ModifyReadOnlyAccountPassword
type ModifyReadOnlyAccountPasswordRequest struct {
	*requests.RpcRequest
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	OriginPassword string `position:"Query" name:"OriginPassword"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// ModifyReadOnlyAccountPasswordResponse is the response struct for api ModifyReadOnlyAccountPassword
type ModifyReadOnlyAccountPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyReadOnlyAccountPasswordRequest creates a request to invoke ModifyReadOnlyAccountPassword API
func CreateModifyReadOnlyAccountPasswordRequest() (request *ModifyReadOnlyAccountPasswordRequest) {
	request = &ModifyReadOnlyAccountPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "ModifyReadOnlyAccountPassword", "", "")
	return
}

// CreateModifyReadOnlyAccountPasswordResponse creates a response to parse from ModifyReadOnlyAccountPassword response
func CreateModifyReadOnlyAccountPasswordResponse() (response *ModifyReadOnlyAccountPasswordResponse) {
	response = &ModifyReadOnlyAccountPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
