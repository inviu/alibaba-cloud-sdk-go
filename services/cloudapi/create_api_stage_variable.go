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

// CreateApiStageVariable invokes the cloudapi.CreateApiStageVariable API synchronously
// api document: https://help.aliyun.com/api/cloudapi/createapistagevariable.html
func (client *Client) CreateApiStageVariable(request *CreateApiStageVariableRequest) (response *CreateApiStageVariableResponse, err error) {
	response = CreateCreateApiStageVariableResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApiStageVariableWithChan invokes the cloudapi.CreateApiStageVariable API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createapistagevariable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateApiStageVariableWithChan(request *CreateApiStageVariableRequest) (<-chan *CreateApiStageVariableResponse, <-chan error) {
	responseChan := make(chan *CreateApiStageVariableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApiStageVariable(request)
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

// CreateApiStageVariableWithCallback invokes the cloudapi.CreateApiStageVariable API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createapistagevariable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateApiStageVariableWithCallback(request *CreateApiStageVariableRequest, callback func(response *CreateApiStageVariableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApiStageVariableResponse
		var err error
		defer close(result)
		response, err = client.CreateApiStageVariable(request)
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

// CreateApiStageVariableRequest is the request struct for api CreateApiStageVariable
type CreateApiStageVariableRequest struct {
	*requests.RpcRequest
	SupportRoute    requests.Boolean `position:"Query" name:"SupportRoute"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	VariableName    string           `position:"Query" name:"VariableName"`
	VariableValue   string           `position:"Query" name:"VariableValue"`
	GroupId         string           `position:"Query" name:"GroupId"`
	StageRouteModel string           `position:"Query" name:"StageRouteModel"`
	StageId         string           `position:"Query" name:"StageId"`
}

// CreateApiStageVariableResponse is the response struct for api CreateApiStageVariable
type CreateApiStageVariableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateApiStageVariableRequest creates a request to invoke CreateApiStageVariable API
func CreateCreateApiStageVariableRequest() (request *CreateApiStageVariableRequest) {
	request = &CreateApiStageVariableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApiStageVariable", "apigateway", "openAPI")
	return
}

// CreateCreateApiStageVariableResponse creates a response to parse from CreateApiStageVariable response
func CreateCreateApiStageVariableResponse() (response *CreateApiStageVariableResponse) {
	response = &CreateApiStageVariableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
