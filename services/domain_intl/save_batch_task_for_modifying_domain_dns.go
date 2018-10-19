package domain_intl

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

// SaveBatchTaskForModifyingDomainDns invokes the domain_intl.SaveBatchTaskForModifyingDomainDns API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskformodifyingdomaindns.html
func (client *Client) SaveBatchTaskForModifyingDomainDns(request *SaveBatchTaskForModifyingDomainDnsRequest) (response *SaveBatchTaskForModifyingDomainDnsResponse, err error) {
	response = CreateSaveBatchTaskForModifyingDomainDnsResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForModifyingDomainDnsWithChan invokes the domain_intl.SaveBatchTaskForModifyingDomainDns API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskformodifyingdomaindns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForModifyingDomainDnsWithChan(request *SaveBatchTaskForModifyingDomainDnsRequest) (<-chan *SaveBatchTaskForModifyingDomainDnsResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForModifyingDomainDnsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForModifyingDomainDns(request)
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

// SaveBatchTaskForModifyingDomainDnsWithCallback invokes the domain_intl.SaveBatchTaskForModifyingDomainDns API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskformodifyingdomaindns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForModifyingDomainDnsWithCallback(request *SaveBatchTaskForModifyingDomainDnsRequest, callback func(response *SaveBatchTaskForModifyingDomainDnsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForModifyingDomainDnsResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForModifyingDomainDns(request)
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

// SaveBatchTaskForModifyingDomainDnsRequest is the request struct for api SaveBatchTaskForModifyingDomainDns
type SaveBatchTaskForModifyingDomainDnsRequest struct {
	*requests.RpcRequest
	UserClientIp     string           `position:"Query" name:"UserClientIp"`
	Lang             string           `position:"Query" name:"Lang"`
	DomainName       *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	AliyunDns        requests.Boolean `position:"Query" name:"AliyunDns"`
	DomainNameServer *[]string        `position:"Query" name:"DomainNameServer"  type:"Repeated"`
}

// SaveBatchTaskForModifyingDomainDnsResponse is the response struct for api SaveBatchTaskForModifyingDomainDns
type SaveBatchTaskForModifyingDomainDnsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForModifyingDomainDnsRequest creates a request to invoke SaveBatchTaskForModifyingDomainDns API
func CreateSaveBatchTaskForModifyingDomainDnsRequest() (request *SaveBatchTaskForModifyingDomainDnsRequest) {
	request = &SaveBatchTaskForModifyingDomainDnsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForModifyingDomainDns", "", "")
	return
}

// CreateSaveBatchTaskForModifyingDomainDnsResponse creates a response to parse from SaveBatchTaskForModifyingDomainDns response
func CreateSaveBatchTaskForModifyingDomainDnsResponse() (response *SaveBatchTaskForModifyingDomainDnsResponse) {
	response = &SaveBatchTaskForModifyingDomainDnsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
