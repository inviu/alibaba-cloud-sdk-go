package dcdn

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

// DescribeDcdnRefreshQuota invokes the dcdn.DescribeDcdnRefreshQuota API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnrefreshquota.html
func (client *Client) DescribeDcdnRefreshQuota(request *DescribeDcdnRefreshQuotaRequest) (response *DescribeDcdnRefreshQuotaResponse, err error) {
	response = CreateDescribeDcdnRefreshQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnRefreshQuotaWithChan invokes the dcdn.DescribeDcdnRefreshQuota API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnrefreshquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnRefreshQuotaWithChan(request *DescribeDcdnRefreshQuotaRequest) (<-chan *DescribeDcdnRefreshQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnRefreshQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnRefreshQuota(request)
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

// DescribeDcdnRefreshQuotaWithCallback invokes the dcdn.DescribeDcdnRefreshQuota API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnrefreshquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnRefreshQuotaWithCallback(request *DescribeDcdnRefreshQuotaRequest, callback func(response *DescribeDcdnRefreshQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnRefreshQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnRefreshQuota(request)
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

// DescribeDcdnRefreshQuotaRequest is the request struct for api DescribeDcdnRefreshQuota
type DescribeDcdnRefreshQuotaRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnRefreshQuotaResponse is the response struct for api DescribeDcdnRefreshQuota
type DescribeDcdnRefreshQuotaResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	UrlQuota      string `json:"UrlQuota" xml:"UrlQuota"`
	DirQuota      string `json:"DirQuota" xml:"DirQuota"`
	UrlRemain     string `json:"UrlRemain" xml:"UrlRemain"`
	DirRemain     string `json:"DirRemain" xml:"DirRemain"`
	PreloadQuota  string `json:"PreloadQuota" xml:"PreloadQuota"`
	BlockQuota    string `json:"BlockQuota" xml:"BlockQuota"`
	PreloadRemain string `json:"PreloadRemain" xml:"PreloadRemain"`
	BlockRemain   string `json:"blockRemain" xml:"blockRemain"`
}

// CreateDescribeDcdnRefreshQuotaRequest creates a request to invoke DescribeDcdnRefreshQuota API
func CreateDescribeDcdnRefreshQuotaRequest() (request *DescribeDcdnRefreshQuotaRequest) {
	request = &DescribeDcdnRefreshQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnRefreshQuota", "dcdn", "openAPI")
	return
}

// CreateDescribeDcdnRefreshQuotaResponse creates a response to parse from DescribeDcdnRefreshQuota response
func CreateDescribeDcdnRefreshQuotaResponse() (response *DescribeDcdnRefreshQuotaResponse) {
	response = &DescribeDcdnRefreshQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
