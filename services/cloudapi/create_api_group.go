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

// CreateApiGroup invokes the cloudapi.CreateApiGroup API synchronously
// api document: https://help.aliyun.com/api/cloudapi/createapigroup.html
func (client *Client) CreateApiGroup(request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
	response = CreateCreateApiGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApiGroupWithChan invokes the cloudapi.CreateApiGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createapigroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateApiGroupWithChan(request *CreateApiGroupRequest) (<-chan *CreateApiGroupResponse, <-chan error) {
	responseChan := make(chan *CreateApiGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApiGroup(request)
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

// CreateApiGroupWithCallback invokes the cloudapi.CreateApiGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createapigroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateApiGroupWithCallback(request *CreateApiGroupRequest, callback func(response *CreateApiGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApiGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateApiGroup(request)
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

// CreateApiGroupRequest is the request struct for api CreateApiGroup
type CreateApiGroupRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Description   string `position:"Query" name:"Description"`
	Source        string `position:"Query" name:"Source"`
	GroupName     string `position:"Query" name:"GroupName"`
}

// CreateApiGroupResponse is the response struct for api CreateApiGroup
type CreateApiGroupResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	GroupId     string `json:"GroupId" xml:"GroupId"`
	GroupName   string `json:"GroupName" xml:"GroupName"`
	SubDomain   string `json:"SubDomain" xml:"SubDomain"`
	Description string `json:"Description" xml:"Description"`
}

// CreateCreateApiGroupRequest creates a request to invoke CreateApiGroup API
func CreateCreateApiGroupRequest() (request *CreateApiGroupRequest) {
	request = &CreateApiGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApiGroup", "apigateway", "openAPI")
	return
}

// CreateCreateApiGroupResponse creates a response to parse from CreateApiGroup response
func CreateCreateApiGroupResponse() (response *CreateApiGroupResponse) {
	response = &CreateApiGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
