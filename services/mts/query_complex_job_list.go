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

// QueryComplexJobList invokes the mts.QueryComplexJobList API synchronously
// api document: https://help.aliyun.com/api/mts/querycomplexjoblist.html
func (client *Client) QueryComplexJobList(request *QueryComplexJobListRequest) (response *QueryComplexJobListResponse, err error) {
	response = CreateQueryComplexJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryComplexJobListWithChan invokes the mts.QueryComplexJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/querycomplexjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryComplexJobListWithChan(request *QueryComplexJobListRequest) (<-chan *QueryComplexJobListResponse, <-chan error) {
	responseChan := make(chan *QueryComplexJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryComplexJobList(request)
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

// QueryComplexJobListWithCallback invokes the mts.QueryComplexJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/querycomplexjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryComplexJobListWithCallback(request *QueryComplexJobListRequest, callback func(response *QueryComplexJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryComplexJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryComplexJobList(request)
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

// QueryComplexJobListRequest is the request struct for api QueryComplexJobList
type QueryComplexJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryComplexJobListResponse is the response struct for api QueryComplexJobList
type QueryComplexJobListResponse struct {
	*responses.BaseResponse
	RequestId      string                              `json:"RequestId" xml:"RequestId"`
	NonExistJobIds NonExistJobIdsInQueryComplexJobList `json:"NonExistJobIds" xml:"NonExistJobIds"`
	JobList        JobListInQueryComplexJobList        `json:"JobList" xml:"JobList"`
}

// CreateQueryComplexJobListRequest creates a request to invoke QueryComplexJobList API
func CreateQueryComplexJobListRequest() (request *QueryComplexJobListRequest) {
	request = &QueryComplexJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryComplexJobList", "mts", "openAPI")
	return
}

// CreateQueryComplexJobListResponse creates a response to parse from QueryComplexJobList response
func CreateQueryComplexJobListResponse() (response *QueryComplexJobListResponse) {
	response = &QueryComplexJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
