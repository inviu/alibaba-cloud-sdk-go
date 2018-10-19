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

// DescribeApiQpsData invokes the cloudapi.DescribeApiQpsData API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapiqpsdata.html
func (client *Client) DescribeApiQpsData(request *DescribeApiQpsDataRequest) (response *DescribeApiQpsDataResponse, err error) {
	response = CreateDescribeApiQpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiQpsDataWithChan invokes the cloudapi.DescribeApiQpsData API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapiqpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiQpsDataWithChan(request *DescribeApiQpsDataRequest) (<-chan *DescribeApiQpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeApiQpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiQpsData(request)
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

// DescribeApiQpsDataWithCallback invokes the cloudapi.DescribeApiQpsData API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapiqpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiQpsDataWithCallback(request *DescribeApiQpsDataRequest, callback func(response *DescribeApiQpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiQpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiQpsData(request)
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

// DescribeApiQpsDataRequest is the request struct for api DescribeApiQpsData
type DescribeApiQpsDataRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribeApiQpsDataResponse is the response struct for api DescribeApiQpsData
type DescribeApiQpsDataResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	CallSuccesses CallSuccesses `json:"CallSuccesses" xml:"CallSuccesses"`
	CallFails     CallFails     `json:"CallFails" xml:"CallFails"`
}

// CreateDescribeApiQpsDataRequest creates a request to invoke DescribeApiQpsData API
func CreateDescribeApiQpsDataRequest() (request *DescribeApiQpsDataRequest) {
	request = &DescribeApiQpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiQpsData", "apigateway", "openAPI")
	return
}

// CreateDescribeApiQpsDataResponse creates a response to parse from DescribeApiQpsData response
func CreateDescribeApiQpsDataResponse() (response *DescribeApiQpsDataResponse) {
	response = &DescribeApiQpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
