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

// ModifyUserPasswords invokes the ehpc.ModifyUserPasswords API synchronously
// api document: https://help.aliyun.com/api/ehpc/modifyuserpasswords.html
func (client *Client) ModifyUserPasswords(request *ModifyUserPasswordsRequest) (response *ModifyUserPasswordsResponse, err error) {
	response = CreateModifyUserPasswordsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUserPasswordsWithChan invokes the ehpc.ModifyUserPasswords API asynchronously
// api document: https://help.aliyun.com/api/ehpc/modifyuserpasswords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUserPasswordsWithChan(request *ModifyUserPasswordsRequest) (<-chan *ModifyUserPasswordsResponse, <-chan error) {
	responseChan := make(chan *ModifyUserPasswordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUserPasswords(request)
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

// ModifyUserPasswordsWithCallback invokes the ehpc.ModifyUserPasswords API asynchronously
// api document: https://help.aliyun.com/api/ehpc/modifyuserpasswords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUserPasswordsWithCallback(request *ModifyUserPasswordsRequest, callback func(response *ModifyUserPasswordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUserPasswordsResponse
		var err error
		defer close(result)
		response, err = client.ModifyUserPasswords(request)
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

// ModifyUserPasswordsRequest is the request struct for api ModifyUserPasswords
type ModifyUserPasswordsRequest struct {
	*requests.RpcRequest
	ClusterId string                     `position:"Query" name:"ClusterId"`
	User      *[]ModifyUserPasswordsUser `position:"Query" name:"User"  type:"Repeated"`
}

// ModifyUserPasswordsUser is a repeated param struct in ModifyUserPasswordsRequest
type ModifyUserPasswordsUser struct {
	Password string `name:"Password"`
	Name     string `name:"Name"`
}

// ModifyUserPasswordsResponse is the response struct for api ModifyUserPasswords
type ModifyUserPasswordsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyUserPasswordsRequest creates a request to invoke ModifyUserPasswords API
func CreateModifyUserPasswordsRequest() (request *ModifyUserPasswordsRequest) {
	request = &ModifyUserPasswordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ModifyUserPasswords", "ehs", "openAPI")
	return
}

// CreateModifyUserPasswordsResponse creates a response to parse from ModifyUserPasswords response
func CreateModifyUserPasswordsResponse() (response *ModifyUserPasswordsResponse) {
	response = &ModifyUserPasswordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
