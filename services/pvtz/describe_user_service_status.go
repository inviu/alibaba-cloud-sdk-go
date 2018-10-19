package pvtz

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

// DescribeUserServiceStatus invokes the pvtz.DescribeUserServiceStatus API synchronously
// api document: https://help.aliyun.com/api/pvtz/describeuserservicestatus.html
func (client *Client) DescribeUserServiceStatus(request *DescribeUserServiceStatusRequest) (response *DescribeUserServiceStatusResponse, err error) {
	response = CreateDescribeUserServiceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserServiceStatusWithChan invokes the pvtz.DescribeUserServiceStatus API asynchronously
// api document: https://help.aliyun.com/api/pvtz/describeuserservicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserServiceStatusWithChan(request *DescribeUserServiceStatusRequest) (<-chan *DescribeUserServiceStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUserServiceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserServiceStatus(request)
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

// DescribeUserServiceStatusWithCallback invokes the pvtz.DescribeUserServiceStatus API asynchronously
// api document: https://help.aliyun.com/api/pvtz/describeuserservicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserServiceStatusWithCallback(request *DescribeUserServiceStatusRequest, callback func(response *DescribeUserServiceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserServiceStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserServiceStatus(request)
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

// DescribeUserServiceStatusRequest is the request struct for api DescribeUserServiceStatus
type DescribeUserServiceStatusRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
}

// DescribeUserServiceStatusResponse is the response struct for api DescribeUserServiceStatus
type DescribeUserServiceStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateDescribeUserServiceStatusRequest creates a request to invoke DescribeUserServiceStatus API
func CreateDescribeUserServiceStatusRequest() (request *DescribeUserServiceStatusRequest) {
	request = &DescribeUserServiceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "DescribeUserServiceStatus", "pvtz", "openAPI")
	return
}

// CreateDescribeUserServiceStatusResponse creates a response to parse from DescribeUserServiceStatus response
func CreateDescribeUserServiceStatusResponse() (response *DescribeUserServiceStatusResponse) {
	response = &DescribeUserServiceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
