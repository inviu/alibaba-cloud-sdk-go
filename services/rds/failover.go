package rds

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

// Failover invokes the rds.Failover API synchronously
// api document: https://help.aliyun.com/api/rds/failover.html
func (client *Client) Failover(request *FailoverRequest) (response *FailoverResponse, err error) {
	response = CreateFailoverResponse()
	err = client.DoAction(request, response)
	return
}

// FailoverWithChan invokes the rds.Failover API asynchronously
// api document: https://help.aliyun.com/api/rds/failover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FailoverWithChan(request *FailoverRequest) (<-chan *FailoverResponse, <-chan error) {
	responseChan := make(chan *FailoverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Failover(request)
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

// FailoverWithCallback invokes the rds.Failover API asynchronously
// api document: https://help.aliyun.com/api/rds/failover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FailoverWithCallback(request *FailoverRequest, callback func(response *FailoverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FailoverResponse
		var err error
		defer close(result)
		response, err = client.Failover(request)
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

// FailoverRequest is the request struct for api Failover
type FailoverRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PrimaryInstanceId    string           `position:"Query" name:"PrimaryInstanceId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// FailoverResponse is the response struct for api Failover
type FailoverResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateFailoverRequest creates a request to invoke Failover API
func CreateFailoverRequest() (request *FailoverRequest) {
	request = &FailoverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "Failover", "rds", "openAPI")
	return
}

// CreateFailoverResponse creates a response to parse from Failover response
func CreateFailoverResponse() (response *FailoverResponse) {
	response = &FailoverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
