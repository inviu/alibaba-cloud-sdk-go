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

// DeleteApiStageVariable invokes the cloudapi.DeleteApiStageVariable API synchronously
// api document: https://help.aliyun.com/api/cloudapi/deleteapistagevariable.html
func (client *Client) DeleteApiStageVariable(request *DeleteApiStageVariableRequest) (response *DeleteApiStageVariableResponse, err error) {
	response = CreateDeleteApiStageVariableResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApiStageVariableWithChan invokes the cloudapi.DeleteApiStageVariable API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deleteapistagevariable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteApiStageVariableWithChan(request *DeleteApiStageVariableRequest) (<-chan *DeleteApiStageVariableResponse, <-chan error) {
	responseChan := make(chan *DeleteApiStageVariableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApiStageVariable(request)
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

// DeleteApiStageVariableWithCallback invokes the cloudapi.DeleteApiStageVariable API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deleteapistagevariable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteApiStageVariableWithCallback(request *DeleteApiStageVariableRequest, callback func(response *DeleteApiStageVariableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApiStageVariableResponse
		var err error
		defer close(result)
		response, err = client.DeleteApiStageVariable(request)
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

// DeleteApiStageVariableRequest is the request struct for api DeleteApiStageVariable
type DeleteApiStageVariableRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	VariableName  string `position:"Query" name:"VariableName"`
	GroupId       string `position:"Query" name:"GroupId"`
	StageId       string `position:"Query" name:"StageId"`
}

// DeleteApiStageVariableResponse is the response struct for api DeleteApiStageVariable
type DeleteApiStageVariableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteApiStageVariableRequest creates a request to invoke DeleteApiStageVariable API
func CreateDeleteApiStageVariableRequest() (request *DeleteApiStageVariableRequest) {
	request = &DeleteApiStageVariableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApiStageVariable", "apigateway", "openAPI")
	return
}

// CreateDeleteApiStageVariableResponse creates a response to parse from DeleteApiStageVariable response
func CreateDeleteApiStageVariableResponse() (response *DeleteApiStageVariableResponse) {
	response = &DeleteApiStageVariableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
