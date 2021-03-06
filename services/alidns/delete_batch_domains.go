package alidns

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

// DeleteBatchDomains invokes the alidns.DeleteBatchDomains API synchronously
// api document: https://help.aliyun.com/api/alidns/deletebatchdomains.html
func (client *Client) DeleteBatchDomains(request *DeleteBatchDomainsRequest) (response *DeleteBatchDomainsResponse, err error) {
	response = CreateDeleteBatchDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBatchDomainsWithChan invokes the alidns.DeleteBatchDomains API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletebatchdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBatchDomainsWithChan(request *DeleteBatchDomainsRequest) (<-chan *DeleteBatchDomainsResponse, <-chan error) {
	responseChan := make(chan *DeleteBatchDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBatchDomains(request)
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

// DeleteBatchDomainsWithCallback invokes the alidns.DeleteBatchDomains API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletebatchdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBatchDomainsWithCallback(request *DeleteBatchDomainsRequest, callback func(response *DeleteBatchDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBatchDomainsResponse
		var err error
		defer close(result)
		response, err = client.DeleteBatchDomains(request)
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

// DeleteBatchDomainsRequest is the request struct for api DeleteBatchDomains
type DeleteBatchDomainsRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Domains      string `position:"Query" name:"Domains"`
}

// DeleteBatchDomainsResponse is the response struct for api DeleteBatchDomains
type DeleteBatchDomainsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
}

// CreateDeleteBatchDomainsRequest creates a request to invoke DeleteBatchDomains API
func CreateDeleteBatchDomainsRequest() (request *DeleteBatchDomainsRequest) {
	request = &DeleteBatchDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteBatchDomains", "", "")
	return
}

// CreateDeleteBatchDomainsResponse creates a response to parse from DeleteBatchDomains response
func CreateDeleteBatchDomainsResponse() (response *DeleteBatchDomainsResponse) {
	response = &DeleteBatchDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
