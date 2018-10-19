package push

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

// QueryTags invokes the push.QueryTags API synchronously
// api document: https://help.aliyun.com/api/push/querytags.html
func (client *Client) QueryTags(request *QueryTagsRequest) (response *QueryTagsResponse, err error) {
	response = CreateQueryTagsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTagsWithChan invokes the push.QueryTags API asynchronously
// api document: https://help.aliyun.com/api/push/querytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTagsWithChan(request *QueryTagsRequest) (<-chan *QueryTagsResponse, <-chan error) {
	responseChan := make(chan *QueryTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTags(request)
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

// QueryTagsWithCallback invokes the push.QueryTags API asynchronously
// api document: https://help.aliyun.com/api/push/querytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTagsWithCallback(request *QueryTagsRequest, callback func(response *QueryTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTagsResponse
		var err error
		defer close(result)
		response, err = client.QueryTags(request)
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

// QueryTagsRequest is the request struct for api QueryTags
type QueryTagsRequest struct {
	*requests.RpcRequest
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	ClientKey string           `position:"Query" name:"ClientKey"`
	KeyType   string           `position:"Query" name:"KeyType"`
}

// QueryTagsResponse is the response struct for api QueryTags
type QueryTagsResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	TagInfos  TagInfosInQueryTags `json:"TagInfos" xml:"TagInfos"`
}

// CreateQueryTagsRequest creates a request to invoke QueryTags API
func CreateQueryTagsRequest() (request *QueryTagsRequest) {
	request = &QueryTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryTags", "", "")
	return
}

// CreateQueryTagsResponse creates a response to parse from QueryTags response
func CreateQueryTagsResponse() (response *QueryTagsResponse) {
	response = &QueryTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
