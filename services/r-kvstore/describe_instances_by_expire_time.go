package r_kvstore

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

// DescribeInstancesByExpireTime invokes the r_kvstore.DescribeInstancesByExpireTime API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancesbyexpiretime.html
func (client *Client) DescribeInstancesByExpireTime(request *DescribeInstancesByExpireTimeRequest) (response *DescribeInstancesByExpireTimeResponse, err error) {
	response = CreateDescribeInstancesByExpireTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstancesByExpireTimeWithChan invokes the r_kvstore.DescribeInstancesByExpireTime API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancesbyexpiretime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstancesByExpireTimeWithChan(request *DescribeInstancesByExpireTimeRequest) (<-chan *DescribeInstancesByExpireTimeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstancesByExpireTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstancesByExpireTime(request)
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

// DescribeInstancesByExpireTimeWithCallback invokes the r_kvstore.DescribeInstancesByExpireTime API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancesbyexpiretime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstancesByExpireTimeWithCallback(request *DescribeInstancesByExpireTimeRequest, callback func(response *DescribeInstancesByExpireTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstancesByExpireTimeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstancesByExpireTime(request)
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

// DescribeInstancesByExpireTimeRequest is the request struct for api DescribeInstancesByExpireTime
type DescribeInstancesByExpireTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	HasExpiredRes        requests.Boolean `position:"Query" name:"HasExpiredRes"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ExpirePeriod         requests.Integer `position:"Query" name:"ExpirePeriod"`
}

// DescribeInstancesByExpireTimeResponse is the response struct for api DescribeInstancesByExpireTime
type DescribeInstancesByExpireTimeResponse struct {
	*responses.BaseResponse
	RequestId  string                                   `json:"RequestId" xml:"RequestId"`
	PageNumber int                                      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                                      `json:"PageSize" xml:"PageSize"`
	TotalCount int                                      `json:"TotalCount" xml:"TotalCount"`
	Instances  InstancesInDescribeInstancesByExpireTime `json:"Instances" xml:"Instances"`
}

// CreateDescribeInstancesByExpireTimeRequest creates a request to invoke DescribeInstancesByExpireTime API
func CreateDescribeInstancesByExpireTimeRequest() (request *DescribeInstancesByExpireTimeRequest) {
	request = &DescribeInstancesByExpireTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstancesByExpireTime", "redisa", "openAPI")
	return
}

// CreateDescribeInstancesByExpireTimeResponse creates a response to parse from DescribeInstancesByExpireTime response
func CreateDescribeInstancesByExpireTimeResponse() (response *DescribeInstancesByExpireTimeResponse) {
	response = &DescribeInstancesByExpireTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
