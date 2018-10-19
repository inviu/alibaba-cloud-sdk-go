package mts

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

// AddPornPipeline invokes the mts.AddPornPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/addpornpipeline.html
func (client *Client) AddPornPipeline(request *AddPornPipelineRequest) (response *AddPornPipelineResponse, err error) {
	response = CreateAddPornPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// AddPornPipelineWithChan invokes the mts.AddPornPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addpornpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddPornPipelineWithChan(request *AddPornPipelineRequest) (<-chan *AddPornPipelineResponse, <-chan error) {
	responseChan := make(chan *AddPornPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPornPipeline(request)
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

// AddPornPipelineWithCallback invokes the mts.AddPornPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addpornpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddPornPipelineWithCallback(request *AddPornPipelineRequest, callback func(response *AddPornPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPornPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddPornPipeline(request)
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

// AddPornPipelineRequest is the request struct for api AddPornPipeline
type AddPornPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
}

// AddPornPipelineResponse is the response struct for api AddPornPipeline
type AddPornPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateAddPornPipelineRequest creates a request to invoke AddPornPipeline API
func CreateAddPornPipelineRequest() (request *AddPornPipelineRequest) {
	request = &AddPornPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddPornPipeline", "mts", "openAPI")
	return
}

// CreateAddPornPipelineResponse creates a response to parse from AddPornPipeline response
func CreateAddPornPipelineResponse() (response *AddPornPipelineResponse) {
	response = &AddPornPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
