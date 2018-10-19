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

// DescribeApps invokes the cloudapi.DescribeApps API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapps.html
func (client *Client) DescribeApps(request *DescribeAppsRequest) (response *DescribeAppsResponse, err error) {
	response = CreateDescribeAppsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppsWithChan invokes the cloudapi.DescribeApps API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppsWithChan(request *DescribeAppsRequest) (<-chan *DescribeAppsResponse, <-chan error) {
	responseChan := make(chan *DescribeAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApps(request)
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

// DescribeAppsWithCallback invokes the cloudapi.DescribeApps API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppsWithCallback(request *DescribeAppsRequest, callback func(response *DescribeAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppsResponse
		var err error
		defer close(result)
		response, err = client.DescribeApps(request)
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

// DescribeAppsRequest is the request struct for api DescribeApps
type DescribeAppsRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	AppOwner      requests.Integer `position:"Query" name:"AppOwner"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeAppsResponse is the response struct for api DescribeApps
type DescribeAppsResponse struct {
	*responses.BaseResponse
	RequestId  string             `json:"RequestId" xml:"RequestId"`
	TotalCount int                `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                `json:"PageSize" xml:"PageSize"`
	PageNumber int                `json:"PageNumber" xml:"PageNumber"`
	Apps       AppsInDescribeApps `json:"Apps" xml:"Apps"`
}

// CreateDescribeAppsRequest creates a request to invoke DescribeApps API
func CreateDescribeAppsRequest() (request *DescribeAppsRequest) {
	request = &DescribeAppsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApps", "apigateway", "openAPI")
	return
}

// CreateDescribeAppsResponse creates a response to parse from DescribeApps response
func CreateDescribeAppsResponse() (response *DescribeAppsResponse) {
	response = &DescribeAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
