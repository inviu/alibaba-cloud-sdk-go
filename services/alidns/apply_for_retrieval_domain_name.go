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

// ApplyForRetrievalDomainName invokes the alidns.ApplyForRetrievalDomainName API synchronously
// api document: https://help.aliyun.com/api/alidns/applyforretrievaldomainname.html
func (client *Client) ApplyForRetrievalDomainName(request *ApplyForRetrievalDomainNameRequest) (response *ApplyForRetrievalDomainNameResponse, err error) {
	response = CreateApplyForRetrievalDomainNameResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyForRetrievalDomainNameWithChan invokes the alidns.ApplyForRetrievalDomainName API asynchronously
// api document: https://help.aliyun.com/api/alidns/applyforretrievaldomainname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyForRetrievalDomainNameWithChan(request *ApplyForRetrievalDomainNameRequest) (<-chan *ApplyForRetrievalDomainNameResponse, <-chan error) {
	responseChan := make(chan *ApplyForRetrievalDomainNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyForRetrievalDomainName(request)
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

// ApplyForRetrievalDomainNameWithCallback invokes the alidns.ApplyForRetrievalDomainName API asynchronously
// api document: https://help.aliyun.com/api/alidns/applyforretrievaldomainname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyForRetrievalDomainNameWithCallback(request *ApplyForRetrievalDomainNameRequest, callback func(response *ApplyForRetrievalDomainNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyForRetrievalDomainNameResponse
		var err error
		defer close(result)
		response, err = client.ApplyForRetrievalDomainName(request)
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

// ApplyForRetrievalDomainNameRequest is the request struct for api ApplyForRetrievalDomainName
type ApplyForRetrievalDomainNameRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

// ApplyForRetrievalDomainNameResponse is the response struct for api ApplyForRetrievalDomainName
type ApplyForRetrievalDomainNameResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainName string `json:"DomainName" xml:"DomainName"`
}

// CreateApplyForRetrievalDomainNameRequest creates a request to invoke ApplyForRetrievalDomainName API
func CreateApplyForRetrievalDomainNameRequest() (request *ApplyForRetrievalDomainNameRequest) {
	request = &ApplyForRetrievalDomainNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "ApplyForRetrievalDomainName", "", "")
	return
}

// CreateApplyForRetrievalDomainNameResponse creates a response to parse from ApplyForRetrievalDomainName response
func CreateApplyForRetrievalDomainNameResponse() (response *ApplyForRetrievalDomainNameResponse) {
	response = &ApplyForRetrievalDomainNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
