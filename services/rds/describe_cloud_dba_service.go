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

// DescribeCloudDBAService invokes the rds.DescribeCloudDBAService API synchronously
// api document: https://help.aliyun.com/api/rds/describeclouddbaservice.html
func (client *Client) DescribeCloudDBAService(request *DescribeCloudDBAServiceRequest) (response *DescribeCloudDBAServiceResponse, err error) {
	response = CreateDescribeCloudDBAServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudDBAServiceWithChan invokes the rds.DescribeCloudDBAService API asynchronously
// api document: https://help.aliyun.com/api/rds/describeclouddbaservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCloudDBAServiceWithChan(request *DescribeCloudDBAServiceRequest) (<-chan *DescribeCloudDBAServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudDBAServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudDBAService(request)
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

// DescribeCloudDBAServiceWithCallback invokes the rds.DescribeCloudDBAService API asynchronously
// api document: https://help.aliyun.com/api/rds/describeclouddbaservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCloudDBAServiceWithCallback(request *DescribeCloudDBAServiceRequest, callback func(response *DescribeCloudDBAServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudDBAServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudDBAService(request)
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

// DescribeCloudDBAServiceRequest is the request struct for api DescribeCloudDBAService
type DescribeCloudDBAServiceRequest struct {
	*requests.RpcRequest
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

// DescribeCloudDBAServiceResponse is the response struct for api DescribeCloudDBAService
type DescribeCloudDBAServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ListData  string `json:"ListData" xml:"ListData"`
	AttrData  string `json:"AttrData" xml:"AttrData"`
}

// CreateDescribeCloudDBAServiceRequest creates a request to invoke DescribeCloudDBAService API
func CreateDescribeCloudDBAServiceRequest() (request *DescribeCloudDBAServiceRequest) {
	request = &DescribeCloudDBAServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeCloudDBAService", "rds", "openAPI")
	return
}

// CreateDescribeCloudDBAServiceResponse creates a response to parse from DescribeCloudDBAService response
func CreateDescribeCloudDBAServiceResponse() (response *DescribeCloudDBAServiceResponse) {
	response = &DescribeCloudDBAServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
