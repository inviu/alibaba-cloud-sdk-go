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

// CreateSimulatedSystemEvents invokes the ecs.CreateSimulatedSystemEvents API synchronously
// api document: https://help.aliyun.com/api/ecs/createsimulatedsystemevents.html
func (client *Client) CreateSimulatedSystemEvents(request *CreateSimulatedSystemEventsRequest) (response *CreateSimulatedSystemEventsResponse, err error) {
	response = CreateCreateSimulatedSystemEventsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSimulatedSystemEventsWithChan invokes the ecs.CreateSimulatedSystemEvents API asynchronously
// api document: https://help.aliyun.com/api/ecs/createsimulatedsystemevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSimulatedSystemEventsWithChan(request *CreateSimulatedSystemEventsRequest) (<-chan *CreateSimulatedSystemEventsResponse, <-chan error) {
	responseChan := make(chan *CreateSimulatedSystemEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSimulatedSystemEvents(request)
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

// CreateSimulatedSystemEventsWithCallback invokes the ecs.CreateSimulatedSystemEvents API asynchronously
// api document: https://help.aliyun.com/api/ecs/createsimulatedsystemevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSimulatedSystemEventsWithCallback(request *CreateSimulatedSystemEventsRequest, callback func(response *CreateSimulatedSystemEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSimulatedSystemEventsResponse
		var err error
		defer close(result)
		response, err = client.CreateSimulatedSystemEvents(request)
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

// CreateSimulatedSystemEventsRequest is the request struct for api CreateSimulatedSystemEvents
type CreateSimulatedSystemEventsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NotBefore            string           `position:"Query" name:"NotBefore"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
	EventType            string           `position:"Query" name:"EventType"`
}

// CreateSimulatedSystemEventsResponse is the response struct for api CreateSimulatedSystemEvents
type CreateSimulatedSystemEventsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	EventIdSet EventIdSet `json:"EventIdSet" xml:"EventIdSet"`
}

// CreateCreateSimulatedSystemEventsRequest creates a request to invoke CreateSimulatedSystemEvents API
func CreateCreateSimulatedSystemEventsRequest() (request *CreateSimulatedSystemEventsRequest) {
	request = &CreateSimulatedSystemEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateSimulatedSystemEvents", "ecs", "openAPI")
	return
}

// CreateCreateSimulatedSystemEventsResponse creates a response to parse from CreateSimulatedSystemEvents response
func CreateCreateSimulatedSystemEventsResponse() (response *CreateSimulatedSystemEventsResponse) {
	response = &CreateSimulatedSystemEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
