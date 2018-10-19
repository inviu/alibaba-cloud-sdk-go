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

// QueryTagJobList invokes the mts.QueryTagJobList API synchronously
// api document: https://help.aliyun.com/api/mts/querytagjoblist.html
func (client *Client) QueryTagJobList(request *QueryTagJobListRequest) (response *QueryTagJobListResponse, err error) {
	response = CreateQueryTagJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTagJobListWithChan invokes the mts.QueryTagJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/querytagjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTagJobListWithChan(request *QueryTagJobListRequest) (<-chan *QueryTagJobListResponse, <-chan error) {
	responseChan := make(chan *QueryTagJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTagJobList(request)
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

// QueryTagJobListWithCallback invokes the mts.QueryTagJobList API asynchronously
// api document: https://help.aliyun.com/api/mts/querytagjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTagJobListWithCallback(request *QueryTagJobListRequest, callback func(response *QueryTagJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTagJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryTagJobList(request)
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

// QueryTagJobListRequest is the request struct for api QueryTagJobList
type QueryTagJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TagJobIds            string           `position:"Query" name:"TagJobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryTagJobListResponse is the response struct for api QueryTagJobList
type QueryTagJobListResponse struct {
	*responses.BaseResponse
	RequestId   string                       `json:"RequestId" xml:"RequestId"`
	NonExistIds NonExistIdsInQueryTagJobList `json:"NonExistIds" xml:"NonExistIds"`
	TagJobList  TagJobList                   `json:"TagJobList" xml:"TagJobList"`
}

// CreateQueryTagJobListRequest creates a request to invoke QueryTagJobList API
func CreateQueryTagJobListRequest() (request *QueryTagJobListRequest) {
	request = &QueryTagJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryTagJobList", "mts", "openAPI")
	return
}

// CreateQueryTagJobListResponse creates a response to parse from QueryTagJobList response
func CreateQueryTagJobListResponse() (response *QueryTagJobListResponse) {
	response = &QueryTagJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
