package cbn

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

// CreateCen invokes the cbn.CreateCen API synchronously
// api document: https://help.aliyun.com/api/cbn/createcen.html
func (client *Client) CreateCen(request *CreateCenRequest) (response *CreateCenResponse, err error) {
	response = CreateCreateCenResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCenWithChan invokes the cbn.CreateCen API asynchronously
// api document: https://help.aliyun.com/api/cbn/createcen.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCenWithChan(request *CreateCenRequest) (<-chan *CreateCenResponse, <-chan error) {
	responseChan := make(chan *CreateCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCen(request)
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

// CreateCenWithCallback invokes the cbn.CreateCen API asynchronously
// api document: https://help.aliyun.com/api/cbn/createcen.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCenWithCallback(request *CreateCenRequest, callback func(response *CreateCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCenResponse
		var err error
		defer close(result)
		response, err = client.CreateCen(request)
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

// CreateCenRequest is the request struct for api CreateCen
type CreateCenRequest struct {
	*requests.RpcRequest
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
}

// CreateCenResponse is the response struct for api CreateCen
type CreateCenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CenId     string `json:"CenId" xml:"CenId"`
}

// CreateCreateCenRequest creates a request to invoke CreateCen API
func CreateCreateCenRequest() (request *CreateCenRequest) {
	request = &CreateCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateCen", "cbn", "openAPI")
	return
}

// CreateCreateCenResponse creates a response to parse from CreateCen response
func CreateCreateCenResponse() (response *CreateCenResponse) {
	response = &CreateCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
