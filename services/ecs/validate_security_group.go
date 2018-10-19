package ecs

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

// ValidateSecurityGroup invokes the ecs.ValidateSecurityGroup API synchronously
// api document: https://help.aliyun.com/api/ecs/validatesecuritygroup.html
func (client *Client) ValidateSecurityGroup(request *ValidateSecurityGroupRequest) (response *ValidateSecurityGroupResponse, err error) {
	response = CreateValidateSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ValidateSecurityGroupWithChan invokes the ecs.ValidateSecurityGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/validatesecuritygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ValidateSecurityGroupWithChan(request *ValidateSecurityGroupRequest) (<-chan *ValidateSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *ValidateSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValidateSecurityGroup(request)
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

// ValidateSecurityGroupWithCallback invokes the ecs.ValidateSecurityGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/validatesecuritygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ValidateSecurityGroupWithCallback(request *ValidateSecurityGroupRequest, callback func(response *ValidateSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValidateSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.ValidateSecurityGroup(request)
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

// ValidateSecurityGroupRequest is the request struct for api ValidateSecurityGroup
type ValidateSecurityGroupRequest struct {
	*requests.RpcRequest
	NicType              string           `position:"Query" name:"NicType"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourcePort           requests.Integer `position:"Query" name:"SourcePort"`
	SourceIp             string           `position:"Query" name:"SourceIp"`
	Direction            string           `position:"Query" name:"Direction"`
	DestIp               string           `position:"Query" name:"DestIp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	DestPort             requests.Integer `position:"Query" name:"DestPort"`
}

// ValidateSecurityGroupResponse is the response struct for api ValidateSecurityGroup
type ValidateSecurityGroupResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Policy             string             `json:"Policy" xml:"Policy"`
	TriggeredGroupRule TriggeredGroupRule `json:"TriggeredGroupRule" xml:"TriggeredGroupRule"`
}

// CreateValidateSecurityGroupRequest creates a request to invoke ValidateSecurityGroup API
func CreateValidateSecurityGroupRequest() (request *ValidateSecurityGroupRequest) {
	request = &ValidateSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ValidateSecurityGroup", "ecs", "openAPI")
	return
}

// CreateValidateSecurityGroupResponse creates a response to parse from ValidateSecurityGroup response
func CreateValidateSecurityGroupResponse() (response *ValidateSecurityGroupResponse) {
	response = &ValidateSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
