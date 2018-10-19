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

// ListClusters invokes the ehpc.ListClusters API synchronously
// api document: https://help.aliyun.com/api/ehpc/listclusters.html
func (client *Client) ListClusters(request *ListClustersRequest) (response *ListClustersResponse, err error) {
	response = CreateListClustersResponse()
	err = client.DoAction(request, response)
	return
}

// ListClustersWithChan invokes the ehpc.ListClusters API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClustersWithChan(request *ListClustersRequest) (<-chan *ListClustersResponse, <-chan error) {
	responseChan := make(chan *ListClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusters(request)
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

// ListClustersWithCallback invokes the ehpc.ListClusters API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClustersWithCallback(request *ListClustersRequest, callback func(response *ListClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClustersResponse
		var err error
		defer close(result)
		response, err = client.ListClusters(request)
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

// ListClustersRequest is the request struct for api ListClusters
type ListClustersRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListClustersResponse is the response struct for api ListClusters
type ListClustersResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	Clusters   Clusters `json:"Clusters" xml:"Clusters"`
}

// CreateListClustersRequest creates a request to invoke ListClusters API
func CreateListClustersRequest() (request *ListClustersRequest) {
	request = &ListClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListClusters", "ehs", "openAPI")
	return
}

// CreateListClustersResponse creates a response to parse from ListClusters response
func CreateListClustersResponse() (response *ListClustersResponse) {
	response = &ListClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
