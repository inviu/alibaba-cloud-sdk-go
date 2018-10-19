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

// DescribeReadOnlyAccount invokes the drds.DescribeReadOnlyAccount API synchronously
// api document: https://help.aliyun.com/api/drds/describereadonlyaccount.html
func (client *Client) DescribeReadOnlyAccount(request *DescribeReadOnlyAccountRequest) (response *DescribeReadOnlyAccountResponse, err error) {
	response = CreateDescribeReadOnlyAccountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReadOnlyAccountWithChan invokes the drds.DescribeReadOnlyAccount API asynchronously
// api document: https://help.aliyun.com/api/drds/describereadonlyaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReadOnlyAccountWithChan(request *DescribeReadOnlyAccountRequest) (<-chan *DescribeReadOnlyAccountResponse, <-chan error) {
	responseChan := make(chan *DescribeReadOnlyAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReadOnlyAccount(request)
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

// DescribeReadOnlyAccountWithCallback invokes the drds.DescribeReadOnlyAccount API asynchronously
// api document: https://help.aliyun.com/api/drds/describereadonlyaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReadOnlyAccountWithCallback(request *DescribeReadOnlyAccountRequest, callback func(response *DescribeReadOnlyAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReadOnlyAccountResponse
		var err error
		defer close(result)
		response, err = client.DescribeReadOnlyAccount(request)
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

// DescribeReadOnlyAccountRequest is the request struct for api DescribeReadOnlyAccount
type DescribeReadOnlyAccountRequest struct {
	*requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeReadOnlyAccountResponse is the response struct for api DescribeReadOnlyAccount
type DescribeReadOnlyAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeReadOnlyAccountRequest creates a request to invoke DescribeReadOnlyAccount API
func CreateDescribeReadOnlyAccountRequest() (request *DescribeReadOnlyAccountRequest) {
	request = &DescribeReadOnlyAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "DescribeReadOnlyAccount", "", "")
	return
}

// CreateDescribeReadOnlyAccountResponse creates a response to parse from DescribeReadOnlyAccount response
func CreateDescribeReadOnlyAccountResponse() (response *DescribeReadOnlyAccountResponse) {
	response = &DescribeReadOnlyAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
