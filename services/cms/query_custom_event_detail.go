package cms

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

// QueryCustomEventDetail invokes the cms.QueryCustomEventDetail API synchronously
// api document: https://help.aliyun.com/api/cms/querycustomeventdetail.html
func (client *Client) QueryCustomEventDetail(request *QueryCustomEventDetailRequest) (response *QueryCustomEventDetailResponse, err error) {
	response = CreateQueryCustomEventDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCustomEventDetailWithChan invokes the cms.QueryCustomEventDetail API asynchronously
// api document: https://help.aliyun.com/api/cms/querycustomeventdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCustomEventDetailWithChan(request *QueryCustomEventDetailRequest) (<-chan *QueryCustomEventDetailResponse, <-chan error) {
	responseChan := make(chan *QueryCustomEventDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCustomEventDetail(request)
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

// QueryCustomEventDetailWithCallback invokes the cms.QueryCustomEventDetail API asynchronously
// api document: https://help.aliyun.com/api/cms/querycustomeventdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCustomEventDetailWithCallback(request *QueryCustomEventDetailRequest, callback func(response *QueryCustomEventDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCustomEventDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryCustomEventDetail(request)
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

// QueryCustomEventDetailRequest is the request struct for api QueryCustomEventDetail
type QueryCustomEventDetailRequest struct {
	*requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

// QueryCustomEventDetailResponse is the response struct for api QueryCustomEventDetail
type QueryCustomEventDetailResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateQueryCustomEventDetailRequest creates a request to invoke QueryCustomEventDetail API
func CreateQueryCustomEventDetailRequest() (request *QueryCustomEventDetailRequest) {
	request = &QueryCustomEventDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "QueryCustomEventDetail", "cms", "openAPI")
	return
}

// CreateQueryCustomEventDetailResponse creates a response to parse from QueryCustomEventDetail response
func CreateQueryCustomEventDetailResponse() (response *QueryCustomEventDetailResponse) {
	response = &QueryCustomEventDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
