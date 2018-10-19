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

// DescribeDomainCCData invokes the cdn.DescribeDomainCCData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainccdata.html
func (client *Client) DescribeDomainCCData(request *DescribeDomainCCDataRequest) (response *DescribeDomainCCDataResponse, err error) {
	response = CreateDescribeDomainCCDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainCCDataWithChan invokes the cdn.DescribeDomainCCData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainccdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCCDataWithChan(request *DescribeDomainCCDataRequest) (<-chan *DescribeDomainCCDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainCCDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainCCData(request)
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

// DescribeDomainCCDataWithCallback invokes the cdn.DescribeDomainCCData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainccdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCCDataWithCallback(request *DescribeDomainCCDataRequest, callback func(response *DescribeDomainCCDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainCCDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainCCData(request)
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

// DescribeDomainCCDataRequest is the request struct for api DescribeDomainCCData
type DescribeDomainCCDataRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainCCDataResponse is the response struct for api DescribeDomainCCData
type DescribeDomainCCDataResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	DomainName   string     `json:"DomainName" xml:"DomainName"`
	DataInterval string     `json:"DataInterval" xml:"DataInterval"`
	StartTime    string     `json:"StartTime" xml:"StartTime"`
	EndTime      string     `json:"EndTime" xml:"EndTime"`
	CCDataList   CCDataList `json:"CCDataList" xml:"CCDataList"`
}

// CreateDescribeDomainCCDataRequest creates a request to invoke DescribeDomainCCData API
func CreateDescribeDomainCCDataRequest() (request *DescribeDomainCCDataRequest) {
	request = &DescribeDomainCCDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCCData", "", "")
	return
}

// CreateDescribeDomainCCDataResponse creates a response to parse from DescribeDomainCCData response
func CreateDescribeDomainCCDataResponse() (response *DescribeDomainCCDataResponse) {
	response = &DescribeDomainCCDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
