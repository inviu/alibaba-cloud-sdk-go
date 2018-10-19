package smartag

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

// CreateCloudConnectNetwork invokes the smartag.CreateCloudConnectNetwork API synchronously
// api document: https://help.aliyun.com/api/smartag/createcloudconnectnetwork.html
func (client *Client) CreateCloudConnectNetwork(request *CreateCloudConnectNetworkRequest) (response *CreateCloudConnectNetworkResponse, err error) {
	response = CreateCreateCloudConnectNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCloudConnectNetworkWithChan invokes the smartag.CreateCloudConnectNetwork API asynchronously
// api document: https://help.aliyun.com/api/smartag/createcloudconnectnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCloudConnectNetworkWithChan(request *CreateCloudConnectNetworkRequest) (<-chan *CreateCloudConnectNetworkResponse, <-chan error) {
	responseChan := make(chan *CreateCloudConnectNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCloudConnectNetwork(request)
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

// CreateCloudConnectNetworkWithCallback invokes the smartag.CreateCloudConnectNetwork API asynchronously
// api document: https://help.aliyun.com/api/smartag/createcloudconnectnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCloudConnectNetworkWithCallback(request *CreateCloudConnectNetworkRequest, callback func(response *CreateCloudConnectNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCloudConnectNetworkResponse
		var err error
		defer close(result)
		response, err = client.CreateCloudConnectNetwork(request)
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

// CreateCloudConnectNetworkRequest is the request struct for api CreateCloudConnectNetwork
type CreateCloudConnectNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateCloudConnectNetworkResponse is the response struct for api CreateCloudConnectNetwork
type CreateCloudConnectNetworkResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	CcnId       string `json:"CcnId" xml:"CcnId"`
	Name        string `json:"Name" xml:"Name"`
	Status      string `json:"Status" xml:"Status"`
	Description string `json:"Description" xml:"Description"`
}

// CreateCreateCloudConnectNetworkRequest creates a request to invoke CreateCloudConnectNetwork API
func CreateCreateCloudConnectNetworkRequest() (request *CreateCloudConnectNetworkRequest) {
	request = &CreateCloudConnectNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "CreateCloudConnectNetwork", "smartag", "openAPI")
	return
}

// CreateCreateCloudConnectNetworkResponse creates a response to parse from CreateCloudConnectNetwork response
func CreateCreateCloudConnectNetworkResponse() (response *CreateCloudConnectNetworkResponse) {
	response = &CreateCloudConnectNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
