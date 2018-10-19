package vpc

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

// DescribeCustomerGateway invokes the vpc.DescribeCustomerGateway API synchronously
// api document: https://help.aliyun.com/api/vpc/describecustomergateway.html
func (client *Client) DescribeCustomerGateway(request *DescribeCustomerGatewayRequest) (response *DescribeCustomerGatewayResponse, err error) {
	response = CreateDescribeCustomerGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomerGatewayWithChan invokes the vpc.DescribeCustomerGateway API asynchronously
// api document: https://help.aliyun.com/api/vpc/describecustomergateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomerGatewayWithChan(request *DescribeCustomerGatewayRequest) (<-chan *DescribeCustomerGatewayResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomerGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomerGateway(request)
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

// DescribeCustomerGatewayWithCallback invokes the vpc.DescribeCustomerGateway API asynchronously
// api document: https://help.aliyun.com/api/vpc/describecustomergateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomerGatewayWithCallback(request *DescribeCustomerGatewayRequest, callback func(response *DescribeCustomerGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomerGatewayResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomerGateway(request)
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

// DescribeCustomerGatewayRequest is the request struct for api DescribeCustomerGateway
type DescribeCustomerGatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string           `position:"Query" name:"CustomerGatewayId"`
}

// DescribeCustomerGatewayResponse is the response struct for api DescribeCustomerGateway
type DescribeCustomerGatewayResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	CustomerGatewayId string `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	IpAddress         string `json:"IpAddress" xml:"IpAddress"`
	Name              string `json:"Name" xml:"Name"`
	Description       string `json:"Description" xml:"Description"`
	CreateTime        int    `json:"CreateTime" xml:"CreateTime"`
}

// CreateDescribeCustomerGatewayRequest creates a request to invoke DescribeCustomerGateway API
func CreateDescribeCustomerGatewayRequest() (request *DescribeCustomerGatewayRequest) {
	request = &DescribeCustomerGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCustomerGateway", "vpc", "openAPI")
	return
}

// CreateDescribeCustomerGatewayResponse creates a response to parse from DescribeCustomerGateway response
func CreateDescribeCustomerGatewayResponse() (response *DescribeCustomerGatewayResponse) {
	response = &DescribeCustomerGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
