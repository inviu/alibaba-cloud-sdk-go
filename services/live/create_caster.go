package live

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

// CreateCaster invokes the live.CreateCaster API synchronously
// api document: https://help.aliyun.com/api/live/createcaster.html
func (client *Client) CreateCaster(request *CreateCasterRequest) (response *CreateCasterResponse, err error) {
	response = CreateCreateCasterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCasterWithChan invokes the live.CreateCaster API asynchronously
// api document: https://help.aliyun.com/api/live/createcaster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCasterWithChan(request *CreateCasterRequest) (<-chan *CreateCasterResponse, <-chan error) {
	responseChan := make(chan *CreateCasterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCaster(request)
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

// CreateCasterWithCallback invokes the live.CreateCaster API asynchronously
// api document: https://help.aliyun.com/api/live/createcaster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCasterWithCallback(request *CreateCasterRequest, callback func(response *CreateCasterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCasterResponse
		var err error
		defer close(result)
		response, err = client.CreateCaster(request)
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

// CreateCasterRequest is the request struct for api CreateCaster
type CreateCasterRequest struct {
	*requests.RpcRequest
	CasterTemplate string           `position:"Query" name:"CasterTemplate"`
	ExpireTime     string           `position:"Query" name:"ExpireTime"`
	NormType       requests.Integer `position:"Query" name:"NormType"`
	CasterName     string           `position:"Query" name:"CasterName"`
	ClientToken    string           `position:"Query" name:"ClientToken"`
	ChargeType     string           `position:"Query" name:"ChargeType"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	PurchaseTime   string           `position:"Query" name:"PurchaseTime"`
}

// CreateCasterResponse is the response struct for api CreateCaster
type CreateCasterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CasterId  string `json:"CasterId" xml:"CasterId"`
}

// CreateCreateCasterRequest creates a request to invoke CreateCaster API
func CreateCreateCasterRequest() (request *CreateCasterRequest) {
	request = &CreateCasterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateCaster", "live", "openAPI")
	return
}

// CreateCreateCasterResponse creates a response to parse from CreateCaster response
func CreateCreateCasterResponse() (response *CreateCasterResponse) {
	response = &CreateCasterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
