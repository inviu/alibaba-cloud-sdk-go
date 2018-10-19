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

// DescribeDnsProductInstances invokes the alidns.DescribeDnsProductInstances API synchronously
// api document: https://help.aliyun.com/api/alidns/describednsproductinstances.html
func (client *Client) DescribeDnsProductInstances(request *DescribeDnsProductInstancesRequest) (response *DescribeDnsProductInstancesResponse, err error) {
	response = CreateDescribeDnsProductInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsProductInstancesWithChan invokes the alidns.DescribeDnsProductInstances API asynchronously
// api document: https://help.aliyun.com/api/alidns/describednsproductinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDnsProductInstancesWithChan(request *DescribeDnsProductInstancesRequest) (<-chan *DescribeDnsProductInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsProductInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsProductInstances(request)
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

// DescribeDnsProductInstancesWithCallback invokes the alidns.DescribeDnsProductInstances API asynchronously
// api document: https://help.aliyun.com/api/alidns/describednsproductinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDnsProductInstancesWithCallback(request *DescribeDnsProductInstancesRequest, callback func(response *DescribeDnsProductInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsProductInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsProductInstances(request)
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

// DescribeDnsProductInstancesRequest is the request struct for api DescribeDnsProductInstances
type DescribeDnsProductInstancesRequest struct {
	*requests.RpcRequest
	Lang         string           `position:"Query" name:"Lang"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	VersionCode  string           `position:"Query" name:"VersionCode"`
}

// DescribeDnsProductInstancesResponse is the response struct for api DescribeDnsProductInstances
type DescribeDnsProductInstancesResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	DnsProducts DnsProducts `json:"DnsProducts" xml:"DnsProducts"`
}

// CreateDescribeDnsProductInstancesRequest creates a request to invoke DescribeDnsProductInstances API
func CreateDescribeDnsProductInstancesRequest() (request *DescribeDnsProductInstancesRequest) {
	request = &DescribeDnsProductInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductInstances", "", "")
	return
}

// CreateDescribeDnsProductInstancesResponse creates a response to parse from DescribeDnsProductInstances response
func CreateDescribeDnsProductInstancesResponse() (response *DescribeDnsProductInstancesResponse) {
	response = &DescribeDnsProductInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
