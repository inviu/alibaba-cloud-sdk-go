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

// RemoveAppsAuthorities invokes the cloudapi.RemoveAppsAuthorities API synchronously
// api document: https://help.aliyun.com/api/cloudapi/removeappsauthorities.html
func (client *Client) RemoveAppsAuthorities(request *RemoveAppsAuthoritiesRequest) (response *RemoveAppsAuthoritiesResponse, err error) {
	response = CreateRemoveAppsAuthoritiesResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveAppsAuthoritiesWithChan invokes the cloudapi.RemoveAppsAuthorities API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/removeappsauthorities.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveAppsAuthoritiesWithChan(request *RemoveAppsAuthoritiesRequest) (<-chan *RemoveAppsAuthoritiesResponse, <-chan error) {
	responseChan := make(chan *RemoveAppsAuthoritiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveAppsAuthorities(request)
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

// RemoveAppsAuthoritiesWithCallback invokes the cloudapi.RemoveAppsAuthorities API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/removeappsauthorities.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveAppsAuthoritiesWithCallback(request *RemoveAppsAuthoritiesRequest, callback func(response *RemoveAppsAuthoritiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveAppsAuthoritiesResponse
		var err error
		defer close(result)
		response, err = client.RemoveAppsAuthorities(request)
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

// RemoveAppsAuthoritiesRequest is the request struct for api RemoveAppsAuthorities
type RemoveAppsAuthoritiesRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	AppIds        string `position:"Query" name:"AppIds"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// RemoveAppsAuthoritiesResponse is the response struct for api RemoveAppsAuthorities
type RemoveAppsAuthoritiesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveAppsAuthoritiesRequest creates a request to invoke RemoveAppsAuthorities API
func CreateRemoveAppsAuthoritiesRequest() (request *RemoveAppsAuthoritiesRequest) {
	request = &RemoveAppsAuthoritiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveAppsAuthorities", "apigateway", "openAPI")
	return
}

// CreateRemoveAppsAuthoritiesResponse creates a response to parse from RemoveAppsAuthorities response
func CreateRemoveAppsAuthoritiesResponse() (response *RemoveAppsAuthoritiesResponse) {
	response = &RemoveAppsAuthoritiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
