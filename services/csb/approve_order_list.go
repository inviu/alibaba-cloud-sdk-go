package csb

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

// ApproveOrderList invokes the csb.ApproveOrderList API synchronously
// api document: https://help.aliyun.com/api/csb/approveorderlist.html
func (client *Client) ApproveOrderList(request *ApproveOrderListRequest) (response *ApproveOrderListResponse, err error) {
	response = CreateApproveOrderListResponse()
	err = client.DoAction(request, response)
	return
}

// ApproveOrderListWithChan invokes the csb.ApproveOrderList API asynchronously
// api document: https://help.aliyun.com/api/csb/approveorderlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApproveOrderListWithChan(request *ApproveOrderListRequest) (<-chan *ApproveOrderListResponse, <-chan error) {
	responseChan := make(chan *ApproveOrderListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApproveOrderList(request)
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

// ApproveOrderListWithCallback invokes the csb.ApproveOrderList API asynchronously
// api document: https://help.aliyun.com/api/csb/approveorderlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApproveOrderListWithCallback(request *ApproveOrderListRequest, callback func(response *ApproveOrderListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApproveOrderListResponse
		var err error
		defer close(result)
		response, err = client.ApproveOrderList(request)
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

// ApproveOrderListRequest is the request struct for api ApproveOrderList
type ApproveOrderListRequest struct {
	*requests.RpcRequest
	Data string `position:"Body" name:"Data"`
}

// ApproveOrderListResponse is the response struct for api ApproveOrderList
type ApproveOrderListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateApproveOrderListRequest creates a request to invoke ApproveOrderList API
func CreateApproveOrderListRequest() (request *ApproveOrderListRequest) {
	request = &ApproveOrderListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "ApproveOrderList", "", "")
	return
}

// CreateApproveOrderListResponse creates a response to parse from ApproveOrderList response
func CreateApproveOrderListResponse() (response *ApproveOrderListResponse) {
	response = &ApproveOrderListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
