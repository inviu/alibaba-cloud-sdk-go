package mts

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

// QueryJobList invokes the mts.QueryJobList API synchronously
// api document: https://help.aliyun.com/api/mts/queryjoblist.html
func (client *Client) QueryJobList(request *QueryJobListRequest) (response *QueryJobListResponse, err error) {
	response = CreateQueryJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryJobListWithChan invokes the mts.QueryJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryJobListWithChan(request *QueryJobListRequest) (<-chan *QueryJobListResponse, <-chan error) {
	responseChan := make(chan *QueryJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryJobList(request)
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

// QueryJobListWithCallback invokes the mts.QueryJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryJobListWithCallback(request *QueryJobListRequest, callback func(response *QueryJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryJobList(request)
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

// QueryJobListRequest is the request struct for api QueryJobList
type QueryJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryJobListResponse is the response struct for api QueryJobList
type QueryJobListResponse struct {
	*responses.BaseResponse
	RequestId      string                       `json:"RequestId" xml:"RequestId"`
	NonExistJobIds NonExistJobIdsInQueryJobList `json:"NonExistJobIds" xml:"NonExistJobIds"`
	JobList        JobListInQueryJobList        `json:"JobList" xml:"JobList"`
}

// CreateQueryJobListRequest creates a request to invoke QueryJobList API
func CreateQueryJobListRequest() (request *QueryJobListRequest) {
	request = &QueryJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryJobList", "mts", "openAPI")
	return
}

// CreateQueryJobListResponse creates a response to parse from QueryJobList response
func CreateQueryJobListResponse() (response *QueryJobListResponse) {
	response = &QueryJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
