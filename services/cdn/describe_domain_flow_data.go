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

// DescribeDomainFlowData invokes the cdn.DescribeDomainFlowData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainflowdata.html
func (client *Client) DescribeDomainFlowData(request *DescribeDomainFlowDataRequest) (response *DescribeDomainFlowDataResponse, err error) {
	response = CreateDescribeDomainFlowDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainFlowDataWithChan invokes the cdn.DescribeDomainFlowData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainflowdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainFlowDataWithChan(request *DescribeDomainFlowDataRequest) (<-chan *DescribeDomainFlowDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainFlowDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainFlowData(request)
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

// DescribeDomainFlowDataWithCallback invokes the cdn.DescribeDomainFlowData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainflowdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainFlowDataWithCallback(request *DescribeDomainFlowDataRequest, callback func(response *DescribeDomainFlowDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainFlowDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainFlowData(request)
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

// DescribeDomainFlowDataRequest is the request struct for api DescribeDomainFlowData
type DescribeDomainFlowDataRequest struct {
	*requests.RpcRequest
	FixTimeGap     string           `position:"Query" name:"FixTimeGap"`
	TimeMerge      string           `position:"Query" name:"TimeMerge"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	DomainType     string           `position:"Query" name:"DomainType"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDomainFlowDataResponse is the response struct for api DescribeDomainFlowData
type DescribeDomainFlowDataResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	DomainName          string              `json:"DomainName" xml:"DomainName"`
	DataInterval        string              `json:"DataInterval" xml:"DataInterval"`
	StartTime           string              `json:"StartTime" xml:"StartTime"`
	EndTime             string              `json:"EndTime" xml:"EndTime"`
	FlowDataPerInterval FlowDataPerInterval `json:"FlowDataPerInterval" xml:"FlowDataPerInterval"`
}

// CreateDescribeDomainFlowDataRequest creates a request to invoke DescribeDomainFlowData API
func CreateDescribeDomainFlowDataRequest() (request *DescribeDomainFlowDataRequest) {
	request = &DescribeDomainFlowDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainFlowData", "", "")
	return
}

// CreateDescribeDomainFlowDataResponse creates a response to parse from DescribeDomainFlowData response
func CreateDescribeDomainFlowDataResponse() (response *DescribeDomainFlowDataResponse) {
	response = &DescribeDomainFlowDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
