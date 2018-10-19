package csb

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

// CommitSuccessedServices invokes the csb.CommitSuccessedServices API synchronously
// api document: https://help.aliyun.com/api/csb/commitsuccessedservices.html
func (client *Client) CommitSuccessedServices(request *CommitSuccessedServicesRequest) (response *CommitSuccessedServicesResponse, err error) {
	response = CreateCommitSuccessedServicesResponse()
	err = client.DoAction(request, response)
	return
}

// CommitSuccessedServicesWithChan invokes the csb.CommitSuccessedServices API asynchronously
// api document: https://help.aliyun.com/api/csb/commitsuccessedservices.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CommitSuccessedServicesWithChan(request *CommitSuccessedServicesRequest) (<-chan *CommitSuccessedServicesResponse, <-chan error) {
	responseChan := make(chan *CommitSuccessedServicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CommitSuccessedServices(request)
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

// CommitSuccessedServicesWithCallback invokes the csb.CommitSuccessedServices API asynchronously
// api document: https://help.aliyun.com/api/csb/commitsuccessedservices.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CommitSuccessedServicesWithCallback(request *CommitSuccessedServicesRequest, callback func(response *CommitSuccessedServicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CommitSuccessedServicesResponse
		var err error
		defer close(result)
		response, err = client.CommitSuccessedServices(request)
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

// CommitSuccessedServicesRequest is the request struct for api CommitSuccessedServices
type CommitSuccessedServicesRequest struct {
	*requests.RpcRequest
	CsbName  string `position:"Query" name:"CsbName"`
	Services string `position:"Body" name:"Services"`
}

// CommitSuccessedServicesResponse is the response struct for api CommitSuccessedServices
type CommitSuccessedServicesResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCommitSuccessedServicesRequest creates a request to invoke CommitSuccessedServices API
func CreateCommitSuccessedServicesRequest() (request *CommitSuccessedServicesRequest) {
	request = &CommitSuccessedServicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "CommitSuccessedServices", "", "")
	return
}

// CreateCommitSuccessedServicesResponse creates a response to parse from CommitSuccessedServices response
func CreateCommitSuccessedServicesResponse() (response *CommitSuccessedServicesResponse) {
	response = &CommitSuccessedServicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
