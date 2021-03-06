package cloudapi

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

// DescribeApiErrorData invokes the cloudapi.DescribeApiErrorData API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapierrordata.html
func (client *Client) DescribeApiErrorData(request *DescribeApiErrorDataRequest) (response *DescribeApiErrorDataResponse, err error) {
	response = CreateDescribeApiErrorDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiErrorDataWithChan invokes the cloudapi.DescribeApiErrorData API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapierrordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiErrorDataWithChan(request *DescribeApiErrorDataRequest) (<-chan *DescribeApiErrorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeApiErrorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiErrorData(request)
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

// DescribeApiErrorDataWithCallback invokes the cloudapi.DescribeApiErrorData API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapierrordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiErrorDataWithCallback(request *DescribeApiErrorDataRequest, callback func(response *DescribeApiErrorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiErrorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiErrorData(request)
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

// DescribeApiErrorDataRequest is the request struct for api DescribeApiErrorData
type DescribeApiErrorDataRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribeApiErrorDataResponse is the response struct for api DescribeApiErrorData
type DescribeApiErrorDataResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ClientErrors ClientErrors `json:"ClientErrors" xml:"ClientErrors"`
	ServerErrors ServerErrors `json:"ServerErrors" xml:"ServerErrors"`
}

// CreateDescribeApiErrorDataRequest creates a request to invoke DescribeApiErrorData API
func CreateDescribeApiErrorDataRequest() (request *DescribeApiErrorDataRequest) {
	request = &DescribeApiErrorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiErrorData", "apigateway", "openAPI")
	return
}

// CreateDescribeApiErrorDataResponse creates a response to parse from DescribeApiErrorData response
func CreateDescribeApiErrorDataResponse() (response *DescribeApiErrorDataResponse) {
	response = &DescribeApiErrorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
