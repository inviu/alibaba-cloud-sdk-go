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

// DescribeDomainCustomLogConfig invokes the cdn.DescribeDomainCustomLogConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomaincustomlogconfig.html
func (client *Client) DescribeDomainCustomLogConfig(request *DescribeDomainCustomLogConfigRequest) (response *DescribeDomainCustomLogConfigResponse, err error) {
	response = CreateDescribeDomainCustomLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainCustomLogConfigWithChan invokes the cdn.DescribeDomainCustomLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaincustomlogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCustomLogConfigWithChan(request *DescribeDomainCustomLogConfigRequest) (<-chan *DescribeDomainCustomLogConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainCustomLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainCustomLogConfig(request)
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

// DescribeDomainCustomLogConfigWithCallback invokes the cdn.DescribeDomainCustomLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaincustomlogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCustomLogConfigWithCallback(request *DescribeDomainCustomLogConfigRequest, callback func(response *DescribeDomainCustomLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainCustomLogConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainCustomLogConfig(request)
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

// DescribeDomainCustomLogConfigRequest is the request struct for api DescribeDomainCustomLogConfig
type DescribeDomainCustomLogConfigRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainCustomLogConfigResponse is the response struct for api DescribeDomainCustomLogConfig
type DescribeDomainCustomLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ConfigId  string `json:"ConfigId" xml:"ConfigId"`
	Remark    string `json:"Remark" xml:"Remark"`
	Sample    string `json:"Sample" xml:"Sample"`
	Tag       string `json:"Tag" xml:"Tag"`
}

// CreateDescribeDomainCustomLogConfigRequest creates a request to invoke DescribeDomainCustomLogConfig API
func CreateDescribeDomainCustomLogConfigRequest() (request *DescribeDomainCustomLogConfigRequest) {
	request = &DescribeDomainCustomLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCustomLogConfig", "", "")
	return
}

// CreateDescribeDomainCustomLogConfigResponse creates a response to parse from DescribeDomainCustomLogConfig response
func CreateDescribeDomainCustomLogConfigResponse() (response *DescribeDomainCustomLogConfigResponse) {
	response = &DescribeDomainCustomLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
