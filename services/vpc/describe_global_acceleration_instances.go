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

// DescribeGlobalAccelerationInstances invokes the vpc.DescribeGlobalAccelerationInstances API synchronously
// api document: https://help.aliyun.com/api/vpc/describeglobalaccelerationinstances.html
func (client *Client) DescribeGlobalAccelerationInstances(request *DescribeGlobalAccelerationInstancesRequest) (response *DescribeGlobalAccelerationInstancesResponse, err error) {
	response = CreateDescribeGlobalAccelerationInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGlobalAccelerationInstancesWithChan invokes the vpc.DescribeGlobalAccelerationInstances API asynchronously
// api document: https://help.aliyun.com/api/vpc/describeglobalaccelerationinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGlobalAccelerationInstancesWithChan(request *DescribeGlobalAccelerationInstancesRequest) (<-chan *DescribeGlobalAccelerationInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeGlobalAccelerationInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGlobalAccelerationInstances(request)
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

// DescribeGlobalAccelerationInstancesWithCallback invokes the vpc.DescribeGlobalAccelerationInstances API asynchronously
// api document: https://help.aliyun.com/api/vpc/describeglobalaccelerationinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGlobalAccelerationInstancesWithCallback(request *DescribeGlobalAccelerationInstancesRequest, callback func(response *DescribeGlobalAccelerationInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGlobalAccelerationInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGlobalAccelerationInstances(request)
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

// DescribeGlobalAccelerationInstancesRequest is the request struct for api DescribeGlobalAccelerationInstances
type DescribeGlobalAccelerationInstancesRequest struct {
	*requests.RpcRequest
	IpAddress                    string           `position:"Query" name:"IpAddress"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BandwidthType                string           `position:"Query" name:"BandwidthType"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation              string           `position:"Query" name:"ServiceLocation"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	IncludeReservationData       requests.Boolean `position:"Query" name:"IncludeReservationData"`
	GlobalAccelerationInstanceId string           `position:"Query" name:"GlobalAccelerationInstanceId"`
	ServerId                     string           `position:"Query" name:"ServerId"`
	PageNumber                   requests.Integer `position:"Query" name:"PageNumber"`
	Name                         string           `position:"Query" name:"Name"`
	PageSize                     requests.Integer `position:"Query" name:"PageSize"`
	Status                       string           `position:"Query" name:"Status"`
}

// DescribeGlobalAccelerationInstancesResponse is the response struct for api DescribeGlobalAccelerationInstances
type DescribeGlobalAccelerationInstancesResponse struct {
	*responses.BaseResponse
	RequestId                   string                                                           `json:"RequestId" xml:"RequestId"`
	TotalCount                  int                                                              `json:"TotalCount" xml:"TotalCount"`
	PageNumber                  int                                                              `json:"PageNumber" xml:"PageNumber"`
	PageSize                    int                                                              `json:"PageSize" xml:"PageSize"`
	GlobalAccelerationInstances GlobalAccelerationInstancesInDescribeGlobalAccelerationInstances `json:"GlobalAccelerationInstances" xml:"GlobalAccelerationInstances"`
}

// CreateDescribeGlobalAccelerationInstancesRequest creates a request to invoke DescribeGlobalAccelerationInstances API
func CreateDescribeGlobalAccelerationInstancesRequest() (request *DescribeGlobalAccelerationInstancesRequest) {
	request = &DescribeGlobalAccelerationInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeGlobalAccelerationInstances", "vpc", "openAPI")
	return
}

// CreateDescribeGlobalAccelerationInstancesResponse creates a response to parse from DescribeGlobalAccelerationInstances response
func CreateDescribeGlobalAccelerationInstancesResponse() (response *DescribeGlobalAccelerationInstancesResponse) {
	response = &DescribeGlobalAccelerationInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
