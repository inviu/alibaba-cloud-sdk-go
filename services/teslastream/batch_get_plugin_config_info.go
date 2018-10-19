package teslastream

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

// BatchGetPluginConfigInfo invokes the teslastream.BatchGetPluginConfigInfo API synchronously
// api document: https://help.aliyun.com/api/teslastream/batchgetpluginconfiginfo.html
func (client *Client) BatchGetPluginConfigInfo(request *BatchGetPluginConfigInfoRequest) (response *BatchGetPluginConfigInfoResponse, err error) {
	response = CreateBatchGetPluginConfigInfoResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetPluginConfigInfoWithChan invokes the teslastream.BatchGetPluginConfigInfo API asynchronously
// api document: https://help.aliyun.com/api/teslastream/batchgetpluginconfiginfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetPluginConfigInfoWithChan(request *BatchGetPluginConfigInfoRequest) (<-chan *BatchGetPluginConfigInfoResponse, <-chan error) {
	responseChan := make(chan *BatchGetPluginConfigInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetPluginConfigInfo(request)
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

// BatchGetPluginConfigInfoWithCallback invokes the teslastream.BatchGetPluginConfigInfo API asynchronously
// api document: https://help.aliyun.com/api/teslastream/batchgetpluginconfiginfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetPluginConfigInfoWithCallback(request *BatchGetPluginConfigInfoRequest, callback func(response *BatchGetPluginConfigInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetPluginConfigInfoResponse
		var err error
		defer close(result)
		response, err = client.BatchGetPluginConfigInfo(request)
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

// BatchGetPluginConfigInfoRequest is the request struct for api BatchGetPluginConfigInfo
type BatchGetPluginConfigInfoRequest struct {
	*requests.RpcRequest
	PluginInfos string `position:"Query" name:"PluginInfos"`
}

// BatchGetPluginConfigInfoResponse is the response struct for api BatchGetPluginConfigInfo
type BatchGetPluginConfigInfoResponse struct {
	*responses.BaseResponse
	Code      int      `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Data      []Plugin `json:"Data" xml:"Data"`
}

// CreateBatchGetPluginConfigInfoRequest creates a request to invoke BatchGetPluginConfigInfo API
func CreateBatchGetPluginConfigInfoRequest() (request *BatchGetPluginConfigInfoRequest) {
	request = &BatchGetPluginConfigInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaStream", "2018-01-15", "BatchGetPluginConfigInfo", "", "")
	return
}

// CreateBatchGetPluginConfigInfoResponse creates a response to parse from BatchGetPluginConfigInfo response
func CreateBatchGetPluginConfigInfoResponse() (response *BatchGetPluginConfigInfoResponse) {
	response = &BatchGetPluginConfigInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
