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

// QueryFailReasonForRegistrantProfileRealNameVerification invokes the domain.QueryFailReasonForRegistrantProfileRealNameVerification API synchronously
// api document: https://help.aliyun.com/api/domain/queryfailreasonforregistrantprofilerealnameverification.html
func (client *Client) QueryFailReasonForRegistrantProfileRealNameVerification(request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) (response *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, err error) {
	response = CreateQueryFailReasonForRegistrantProfileRealNameVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFailReasonForRegistrantProfileRealNameVerificationWithChan invokes the domain.QueryFailReasonForRegistrantProfileRealNameVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/queryfailreasonforregistrantprofilerealnameverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFailReasonForRegistrantProfileRealNameVerificationWithChan(request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) (<-chan *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, <-chan error) {
	responseChan := make(chan *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFailReasonForRegistrantProfileRealNameVerification(request)
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

// QueryFailReasonForRegistrantProfileRealNameVerificationWithCallback invokes the domain.QueryFailReasonForRegistrantProfileRealNameVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/queryfailreasonforregistrantprofilerealnameverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFailReasonForRegistrantProfileRealNameVerificationWithCallback(request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest, callback func(response *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFailReasonForRegistrantProfileRealNameVerificationResponse
		var err error
		defer close(result)
		response, err = client.QueryFailReasonForRegistrantProfileRealNameVerification(request)
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

// QueryFailReasonForRegistrantProfileRealNameVerificationRequest is the request struct for api QueryFailReasonForRegistrantProfileRealNameVerification
type QueryFailReasonForRegistrantProfileRealNameVerificationRequest struct {
	*requests.RpcRequest
	UserClientIp        string           `position:"Query" name:"UserClientIp"`
	RegistrantProfileID requests.Integer `position:"Query" name:"RegistrantProfileID"`
	Lang                string           `position:"Query" name:"Lang"`
}

// QueryFailReasonForRegistrantProfileRealNameVerificationResponse is the response struct for api QueryFailReasonForRegistrantProfileRealNameVerification
type QueryFailReasonForRegistrantProfileRealNameVerificationResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Data      []FailRecord `json:"Data" xml:"Data"`
}

// CreateQueryFailReasonForRegistrantProfileRealNameVerificationRequest creates a request to invoke QueryFailReasonForRegistrantProfileRealNameVerification API
func CreateQueryFailReasonForRegistrantProfileRealNameVerificationRequest() (request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) {
	request = &QueryFailReasonForRegistrantProfileRealNameVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryFailReasonForRegistrantProfileRealNameVerification", "", "")
	return
}

// CreateQueryFailReasonForRegistrantProfileRealNameVerificationResponse creates a response to parse from QueryFailReasonForRegistrantProfileRealNameVerification response
func CreateQueryFailReasonForRegistrantProfileRealNameVerificationResponse() (response *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) {
	response = &QueryFailReasonForRegistrantProfileRealNameVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
