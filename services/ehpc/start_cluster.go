package ehpc

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

// StartCluster invokes the ehpc.StartCluster API synchronously
// api document: https://help.aliyun.com/api/ehpc/startcluster.html
func (client *Client) StartCluster(request *StartClusterRequest) (response *StartClusterResponse, err error) {
	response = CreateStartClusterResponse()
	err = client.DoAction(request, response)
	return
}

// StartClusterWithChan invokes the ehpc.StartCluster API asynchronously
// api document: https://help.aliyun.com/api/ehpc/startcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartClusterWithChan(request *StartClusterRequest) (<-chan *StartClusterResponse, <-chan error) {
	responseChan := make(chan *StartClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartCluster(request)
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

// StartClusterWithCallback invokes the ehpc.StartCluster API asynchronously
// api document: https://help.aliyun.com/api/ehpc/startcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartClusterWithCallback(request *StartClusterRequest, callback func(response *StartClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartClusterResponse
		var err error
		defer close(result)
		response, err = client.StartCluster(request)
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

// StartClusterRequest is the request struct for api StartCluster
type StartClusterRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// StartClusterResponse is the response struct for api StartCluster
type StartClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartClusterRequest creates a request to invoke StartCluster API
func CreateStartClusterRequest() (request *StartClusterRequest) {
	request = &StartClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StartCluster", "ehs", "openAPI")
	return
}

// CreateStartClusterResponse creates a response to parse from StartCluster response
func CreateStartClusterResponse() (response *StartClusterResponse) {
	response = &StartClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
