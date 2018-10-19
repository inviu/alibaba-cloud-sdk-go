package slb

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

// DescribeServerCertificates invokes the slb.DescribeServerCertificates API synchronously
// api document: https://help.aliyun.com/api/slb/describeservercertificates.html
func (client *Client) DescribeServerCertificates(request *DescribeServerCertificatesRequest) (response *DescribeServerCertificatesResponse, err error) {
	response = CreateDescribeServerCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServerCertificatesWithChan invokes the slb.DescribeServerCertificates API asynchronously
// api document: https://help.aliyun.com/api/slb/describeservercertificates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServerCertificatesWithChan(request *DescribeServerCertificatesRequest) (<-chan *DescribeServerCertificatesResponse, <-chan error) {
	responseChan := make(chan *DescribeServerCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServerCertificates(request)
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

// DescribeServerCertificatesWithCallback invokes the slb.DescribeServerCertificates API asynchronously
// api document: https://help.aliyun.com/api/slb/describeservercertificates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServerCertificatesWithCallback(request *DescribeServerCertificatesRequest, callback func(response *DescribeServerCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServerCertificatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeServerCertificates(request)
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

// DescribeServerCertificatesRequest is the request struct for api DescribeServerCertificates
type DescribeServerCertificatesRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ServerCertificateId  string           `position:"Query" name:"ServerCertificateId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// DescribeServerCertificatesResponse is the response struct for api DescribeServerCertificates
type DescribeServerCertificatesResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	ServerCertificates ServerCertificates `json:"ServerCertificates" xml:"ServerCertificates"`
}

// CreateDescribeServerCertificatesRequest creates a request to invoke DescribeServerCertificates API
func CreateDescribeServerCertificatesRequest() (request *DescribeServerCertificatesRequest) {
	request = &DescribeServerCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeServerCertificates", "slb", "openAPI")
	return
}

// CreateDescribeServerCertificatesResponse creates a response to parse from DescribeServerCertificates response
func CreateDescribeServerCertificatesResponse() (response *DescribeServerCertificatesResponse) {
	response = &DescribeServerCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
