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

// ModifyCdnService invokes the cdn.ModifyCdnService API synchronously
// api document: https://help.aliyun.com/api/cdn/modifycdnservice.html
func (client *Client) ModifyCdnService(request *ModifyCdnServiceRequest) (response *ModifyCdnServiceResponse, err error) {
	response = CreateModifyCdnServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCdnServiceWithChan invokes the cdn.ModifyCdnService API asynchronously
// api document: https://help.aliyun.com/api/cdn/modifycdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCdnServiceWithChan(request *ModifyCdnServiceRequest) (<-chan *ModifyCdnServiceResponse, <-chan error) {
	responseChan := make(chan *ModifyCdnServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCdnService(request)
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

// ModifyCdnServiceWithCallback invokes the cdn.ModifyCdnService API asynchronously
// api document: https://help.aliyun.com/api/cdn/modifycdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCdnServiceWithCallback(request *ModifyCdnServiceRequest, callback func(response *ModifyCdnServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCdnServiceResponse
		var err error
		defer close(result)
		response, err = client.ModifyCdnService(request)
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

// ModifyCdnServiceRequest is the request struct for api ModifyCdnService
type ModifyCdnServiceRequest struct {
	*requests.RpcRequest
	SecurityToken      string           `position:"Query" name:"SecurityToken"`
	InternetChargeType string           `position:"Query" name:"InternetChargeType"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyCdnServiceResponse is the response struct for api ModifyCdnService
type ModifyCdnServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCdnServiceRequest creates a request to invoke ModifyCdnService API
func CreateModifyCdnServiceRequest() (request *ModifyCdnServiceRequest) {
	request = &ModifyCdnServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ModifyCdnService", "", "")
	return
}

// CreateModifyCdnServiceResponse creates a response to parse from ModifyCdnService response
func CreateModifyCdnServiceResponse() (response *ModifyCdnServiceResponse) {
	response = &ModifyCdnServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
