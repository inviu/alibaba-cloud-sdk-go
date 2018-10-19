package scdn

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

// DescribeScdnCertificateList invokes the scdn.DescribeScdnCertificateList API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdncertificatelist.html
func (client *Client) DescribeScdnCertificateList(request *DescribeScdnCertificateListRequest) (response *DescribeScdnCertificateListResponse, err error) {
	response = CreateDescribeScdnCertificateListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnCertificateListWithChan invokes the scdn.DescribeScdnCertificateList API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdncertificatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnCertificateListWithChan(request *DescribeScdnCertificateListRequest) (<-chan *DescribeScdnCertificateListResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnCertificateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnCertificateList(request)
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

// DescribeScdnCertificateListWithCallback invokes the scdn.DescribeScdnCertificateList API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdncertificatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnCertificateListWithCallback(request *DescribeScdnCertificateListRequest, callback func(response *DescribeScdnCertificateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnCertificateListResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnCertificateList(request)
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

// DescribeScdnCertificateListRequest is the request struct for api DescribeScdnCertificateList
type DescribeScdnCertificateListRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnCertificateListResponse is the response struct for api DescribeScdnCertificateList
type DescribeScdnCertificateListResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	CertificateListModel CertificateListModel `json:"CertificateListModel" xml:"CertificateListModel"`
}

// CreateDescribeScdnCertificateListRequest creates a request to invoke DescribeScdnCertificateList API
func CreateDescribeScdnCertificateListRequest() (request *DescribeScdnCertificateListRequest) {
	request = &DescribeScdnCertificateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnCertificateList", "", "")
	return
}

// CreateDescribeScdnCertificateListResponse creates a response to parse from DescribeScdnCertificateList response
func CreateDescribeScdnCertificateListResponse() (response *DescribeScdnCertificateListResponse) {
	response = &DescribeScdnCertificateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
