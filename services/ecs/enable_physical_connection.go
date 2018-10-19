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

// EnablePhysicalConnection invokes the ecs.EnablePhysicalConnection API synchronously
// api document: https://help.aliyun.com/api/ecs/enablephysicalconnection.html
func (client *Client) EnablePhysicalConnection(request *EnablePhysicalConnectionRequest) (response *EnablePhysicalConnectionResponse, err error) {
	response = CreateEnablePhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// EnablePhysicalConnectionWithChan invokes the ecs.EnablePhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/enablephysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnablePhysicalConnectionWithChan(request *EnablePhysicalConnectionRequest) (<-chan *EnablePhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *EnablePhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnablePhysicalConnection(request)
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

// EnablePhysicalConnectionWithCallback invokes the ecs.EnablePhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/enablephysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnablePhysicalConnectionWithCallback(request *EnablePhysicalConnectionRequest, callback func(response *EnablePhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnablePhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.EnablePhysicalConnection(request)
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

// EnablePhysicalConnectionRequest is the request struct for api EnablePhysicalConnection
type EnablePhysicalConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// EnablePhysicalConnectionResponse is the response struct for api EnablePhysicalConnection
type EnablePhysicalConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnablePhysicalConnectionRequest creates a request to invoke EnablePhysicalConnection API
func CreateEnablePhysicalConnectionRequest() (request *EnablePhysicalConnectionRequest) {
	request = &EnablePhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "EnablePhysicalConnection", "ecs", "openAPI")
	return
}

// CreateEnablePhysicalConnectionResponse creates a response to parse from EnablePhysicalConnection response
func CreateEnablePhysicalConnectionResponse() (response *EnablePhysicalConnectionResponse) {
	response = &EnablePhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
