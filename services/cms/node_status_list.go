package cms

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

// NodeStatusList invokes the cms.NodeStatusList API synchronously
// api document: https://help.aliyun.com/api/cms/nodestatuslist.html
func (client *Client) NodeStatusList(request *NodeStatusListRequest) (response *NodeStatusListResponse, err error) {
	response = CreateNodeStatusListResponse()
	err = client.DoAction(request, response)
	return
}

// NodeStatusListWithChan invokes the cms.NodeStatusList API asynchronously
// api document: https://help.aliyun.com/api/cms/nodestatuslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) NodeStatusListWithChan(request *NodeStatusListRequest) (<-chan *NodeStatusListResponse, <-chan error) {
	responseChan := make(chan *NodeStatusListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeStatusList(request)
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

// NodeStatusListWithCallback invokes the cms.NodeStatusList API asynchronously
// api document: https://help.aliyun.com/api/cms/nodestatuslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) NodeStatusListWithCallback(request *NodeStatusListRequest, callback func(response *NodeStatusListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeStatusListResponse
		var err error
		defer close(result)
		response, err = client.NodeStatusList(request)
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

// NodeStatusListRequest is the request struct for api NodeStatusList
type NodeStatusListRequest struct {
	*requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

// NodeStatusListResponse is the response struct for api NodeStatusList
type NodeStatusListResponse struct {
	*responses.BaseResponse
	ErrorCode      int            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool           `json:"Success" xml:"Success"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	NodeStatusList NodeStatusList `json:"NodeStatusList" xml:"NodeStatusList"`
}

// CreateNodeStatusListRequest creates a request to invoke NodeStatusList API
func CreateNodeStatusListRequest() (request *NodeStatusListRequest) {
	request = &NodeStatusListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "NodeStatusList", "cms", "openAPI")
	return
}

// CreateNodeStatusListResponse creates a response to parse from NodeStatusList response
func CreateNodeStatusListResponse() (response *NodeStatusListResponse) {
	response = &NodeStatusListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
