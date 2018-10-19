package cdn

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

// DescribeL2VipsByDynamicDomain invokes the cdn.DescribeL2VipsByDynamicDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/describel2vipsbydynamicdomain.html
func (client *Client) DescribeL2VipsByDynamicDomain(request *DescribeL2VipsByDynamicDomainRequest) (response *DescribeL2VipsByDynamicDomainResponse, err error) {
	response = CreateDescribeL2VipsByDynamicDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeL2VipsByDynamicDomainWithChan invokes the cdn.DescribeL2VipsByDynamicDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/describel2vipsbydynamicdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeL2VipsByDynamicDomainWithChan(request *DescribeL2VipsByDynamicDomainRequest) (<-chan *DescribeL2VipsByDynamicDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeL2VipsByDynamicDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeL2VipsByDynamicDomain(request)
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

// DescribeL2VipsByDynamicDomainWithCallback invokes the cdn.DescribeL2VipsByDynamicDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/describel2vipsbydynamicdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeL2VipsByDynamicDomainWithCallback(request *DescribeL2VipsByDynamicDomainRequest, callback func(response *DescribeL2VipsByDynamicDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeL2VipsByDynamicDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeL2VipsByDynamicDomain(request)
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

// DescribeL2VipsByDynamicDomainRequest is the request struct for api DescribeL2VipsByDynamicDomain
type DescribeL2VipsByDynamicDomainRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeL2VipsByDynamicDomainResponse is the response struct for api DescribeL2VipsByDynamicDomain
type DescribeL2VipsByDynamicDomainResponse struct {
	*responses.BaseResponse
	RequestId  string                              `json:"RequestId" xml:"RequestId"`
	DomainName string                              `json:"DomainName" xml:"DomainName"`
	Vips       VipsInDescribeL2VipsByDynamicDomain `json:"Vips" xml:"Vips"`
}

// CreateDescribeL2VipsByDynamicDomainRequest creates a request to invoke DescribeL2VipsByDynamicDomain API
func CreateDescribeL2VipsByDynamicDomainRequest() (request *DescribeL2VipsByDynamicDomainRequest) {
	request = &DescribeL2VipsByDynamicDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeL2VipsByDynamicDomain", "", "")
	return
}

// CreateDescribeL2VipsByDynamicDomainResponse creates a response to parse from DescribeL2VipsByDynamicDomain response
func CreateDescribeL2VipsByDynamicDomainResponse() (response *DescribeL2VipsByDynamicDomainResponse) {
	response = &DescribeL2VipsByDynamicDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
