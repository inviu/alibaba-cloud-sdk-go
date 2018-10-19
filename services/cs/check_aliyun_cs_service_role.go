package cs

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

// CheckAliyunCSServiceRole invokes the cs.CheckAliyunCSServiceRole API synchronously
// api document: https://help.aliyun.com/api/cs/checkaliyuncsservicerole.html
func (client *Client) CheckAliyunCSServiceRole(request *CheckAliyunCSServiceRoleRequest) (response *CheckAliyunCSServiceRoleResponse, err error) {
	response = CreateCheckAliyunCSServiceRoleResponse()
	err = client.DoAction(request, response)
	return
}

// CheckAliyunCSServiceRoleWithChan invokes the cs.CheckAliyunCSServiceRole API asynchronously
// api document: https://help.aliyun.com/api/cs/checkaliyuncsservicerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAliyunCSServiceRoleWithChan(request *CheckAliyunCSServiceRoleRequest) (<-chan *CheckAliyunCSServiceRoleResponse, <-chan error) {
	responseChan := make(chan *CheckAliyunCSServiceRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAliyunCSServiceRole(request)
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

// CheckAliyunCSServiceRoleWithCallback invokes the cs.CheckAliyunCSServiceRole API asynchronously
// api document: https://help.aliyun.com/api/cs/checkaliyuncsservicerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAliyunCSServiceRoleWithCallback(request *CheckAliyunCSServiceRoleRequest, callback func(response *CheckAliyunCSServiceRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAliyunCSServiceRoleResponse
		var err error
		defer close(result)
		response, err = client.CheckAliyunCSServiceRole(request)
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

// CheckAliyunCSServiceRoleRequest is the request struct for api CheckAliyunCSServiceRole
type CheckAliyunCSServiceRoleRequest struct {
	*requests.RoaRequest
}

// CheckAliyunCSServiceRoleResponse is the response struct for api CheckAliyunCSServiceRole
type CheckAliyunCSServiceRoleResponse struct {
	*responses.BaseResponse
}

// CreateCheckAliyunCSServiceRoleRequest creates a request to invoke CheckAliyunCSServiceRole API
func CreateCheckAliyunCSServiceRoleRequest() (request *CheckAliyunCSServiceRoleRequest) {
	request = &CheckAliyunCSServiceRoleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "CheckAliyunCSServiceRole", "/aliyuncsrole/status", "", "")
	request.Method = requests.GET
	return
}

// CreateCheckAliyunCSServiceRoleResponse creates a response to parse from CheckAliyunCSServiceRole response
func CreateCheckAliyunCSServiceRoleResponse() (response *CheckAliyunCSServiceRoleResponse) {
	response = &CheckAliyunCSServiceRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
