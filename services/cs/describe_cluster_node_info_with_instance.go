package cs

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

// DescribeClusterNodeInfoWithInstance invokes the cs.DescribeClusterNodeInfoWithInstance API synchronously
// api document: https://help.aliyun.com/api/cs/describeclusternodeinfowithinstance.html
func (client *Client) DescribeClusterNodeInfoWithInstance(request *DescribeClusterNodeInfoWithInstanceRequest) (response *DescribeClusterNodeInfoWithInstanceResponse, err error) {
	response = CreateDescribeClusterNodeInfoWithInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterNodeInfoWithInstanceWithChan invokes the cs.DescribeClusterNodeInfoWithInstance API asynchronously
// api document: https://help.aliyun.com/api/cs/describeclusternodeinfowithinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterNodeInfoWithInstanceWithChan(request *DescribeClusterNodeInfoWithInstanceRequest) (<-chan *DescribeClusterNodeInfoWithInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterNodeInfoWithInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterNodeInfoWithInstance(request)
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

// DescribeClusterNodeInfoWithInstanceWithCallback invokes the cs.DescribeClusterNodeInfoWithInstance API asynchronously
// api document: https://help.aliyun.com/api/cs/describeclusternodeinfowithinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterNodeInfoWithInstanceWithCallback(request *DescribeClusterNodeInfoWithInstanceRequest, callback func(response *DescribeClusterNodeInfoWithInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterNodeInfoWithInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterNodeInfoWithInstance(request)
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

// DescribeClusterNodeInfoWithInstanceRequest is the request struct for api DescribeClusterNodeInfoWithInstance
type DescribeClusterNodeInfoWithInstanceRequest struct {
	*requests.RoaRequest
	Token      string `position:"Path" name:"Token"`
	InstanceId string `position:"Path" name:"InstanceId"`
}

// DescribeClusterNodeInfoWithInstanceResponse is the response struct for api DescribeClusterNodeInfoWithInstance
type DescribeClusterNodeInfoWithInstanceResponse struct {
	*responses.BaseResponse
}

// CreateDescribeClusterNodeInfoWithInstanceRequest creates a request to invoke DescribeClusterNodeInfoWithInstance API
func CreateDescribeClusterNodeInfoWithInstanceRequest() (request *DescribeClusterNodeInfoWithInstanceRequest) {
	request = &DescribeClusterNodeInfoWithInstanceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodeInfoWithInstance", "/token/[Token]/instance/[InstanceId]/node_info", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeClusterNodeInfoWithInstanceResponse creates a response to parse from DescribeClusterNodeInfoWithInstance response
func CreateDescribeClusterNodeInfoWithInstanceResponse() (response *DescribeClusterNodeInfoWithInstanceResponse) {
	response = &DescribeClusterNodeInfoWithInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
