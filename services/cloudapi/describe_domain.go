package cloudapi

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

// DescribeDomain invokes the cloudapi.DescribeDomain API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describedomain.html
func (client *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
	response = CreateDescribeDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainWithChan invokes the cloudapi.DescribeDomain API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describedomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainWithChan(request *DescribeDomainRequest) (<-chan *DescribeDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomain(request)
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

// DescribeDomainWithCallback invokes the cloudapi.DescribeDomain API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describedomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainWithCallback(request *DescribeDomainRequest, callback func(response *DescribeDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomain(request)
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

// DescribeDomainRequest is the request struct for api DescribeDomain
type DescribeDomainRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	DomainName    string `position:"Query" name:"DomainName"`
}

// DescribeDomainResponse is the response struct for api DescribeDomain
type DescribeDomainResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	GroupId               string `json:"GroupId" xml:"GroupId"`
	DomainName            string `json:"DomainName" xml:"DomainName"`
	SubDomain             string `json:"SubDomain" xml:"SubDomain"`
	CertificateId         string `json:"CertificateId" xml:"CertificateId"`
	CertificateName       string `json:"CertificateName" xml:"CertificateName"`
	CertificateBody       string `json:"CertificateBody" xml:"CertificateBody"`
	CertificatePrivateKey string `json:"CertificatePrivateKey" xml:"CertificatePrivateKey"`
	DomainBindingStatus   string `json:"DomainBindingStatus" xml:"DomainBindingStatus"`
	DomainCNAMEStatus     string `json:"DomainCNAMEStatus" xml:"DomainCNAMEStatus"`
	DomainLegalStatus     string `json:"DomainLegalStatus" xml:"DomainLegalStatus"`
	DomainWebSocketStatus string `json:"DomainWebSocketStatus" xml:"DomainWebSocketStatus"`
	DomainRemark          string `json:"DomainRemark" xml:"DomainRemark"`
}

// CreateDescribeDomainRequest creates a request to invoke DescribeDomain API
func CreateDescribeDomainRequest() (request *DescribeDomainRequest) {
	request = &DescribeDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDomain", "apigateway", "openAPI")
	return
}

// CreateDescribeDomainResponse creates a response to parse from DescribeDomain response
func CreateDescribeDomainResponse() (response *DescribeDomainResponse) {
	response = &DescribeDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
