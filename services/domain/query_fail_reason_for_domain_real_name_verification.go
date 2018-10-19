package domain

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

// QueryFailReasonForDomainRealNameVerification invokes the domain.QueryFailReasonForDomainRealNameVerification API synchronously
// api document: https://help.aliyun.com/api/domain/queryfailreasonfordomainrealnameverification.html
func (client *Client) QueryFailReasonForDomainRealNameVerification(request *QueryFailReasonForDomainRealNameVerificationRequest) (response *QueryFailReasonForDomainRealNameVerificationResponse, err error) {
	response = CreateQueryFailReasonForDomainRealNameVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFailReasonForDomainRealNameVerificationWithChan invokes the domain.QueryFailReasonForDomainRealNameVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/queryfailreasonfordomainrealnameverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFailReasonForDomainRealNameVerificationWithChan(request *QueryFailReasonForDomainRealNameVerificationRequest) (<-chan *QueryFailReasonForDomainRealNameVerificationResponse, <-chan error) {
	responseChan := make(chan *QueryFailReasonForDomainRealNameVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFailReasonForDomainRealNameVerification(request)
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

// QueryFailReasonForDomainRealNameVerificationWithCallback invokes the domain.QueryFailReasonForDomainRealNameVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/queryfailreasonfordomainrealnameverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFailReasonForDomainRealNameVerificationWithCallback(request *QueryFailReasonForDomainRealNameVerificationRequest, callback func(response *QueryFailReasonForDomainRealNameVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFailReasonForDomainRealNameVerificationResponse
		var err error
		defer close(result)
		response, err = client.QueryFailReasonForDomainRealNameVerification(request)
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

// QueryFailReasonForDomainRealNameVerificationRequest is the request struct for api QueryFailReasonForDomainRealNameVerification
type QueryFailReasonForDomainRealNameVerificationRequest struct {
	*requests.RpcRequest
	RealNameVerificationAction string `position:"Query" name:"RealNameVerificationAction"`
	UserClientIp               string `position:"Query" name:"UserClientIp"`
	DomainName                 string `position:"Query" name:"DomainName"`
	Lang                       string `position:"Query" name:"Lang"`
}

// QueryFailReasonForDomainRealNameVerificationResponse is the response struct for api QueryFailReasonForDomainRealNameVerification
type QueryFailReasonForDomainRealNameVerificationResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Data      []FailRecord `json:"Data" xml:"Data"`
}

// CreateQueryFailReasonForDomainRealNameVerificationRequest creates a request to invoke QueryFailReasonForDomainRealNameVerification API
func CreateQueryFailReasonForDomainRealNameVerificationRequest() (request *QueryFailReasonForDomainRealNameVerificationRequest) {
	request = &QueryFailReasonForDomainRealNameVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryFailReasonForDomainRealNameVerification", "", "")
	return
}

// CreateQueryFailReasonForDomainRealNameVerificationResponse creates a response to parse from QueryFailReasonForDomainRealNameVerification response
func CreateQueryFailReasonForDomainRealNameVerificationResponse() (response *QueryFailReasonForDomainRealNameVerificationResponse) {
	response = &QueryFailReasonForDomainRealNameVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
