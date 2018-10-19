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

// SdkGenerateByGroup invokes the cloudapi.SdkGenerateByGroup API synchronously
// api document: https://help.aliyun.com/api/cloudapi/sdkgeneratebygroup.html
func (client *Client) SdkGenerateByGroup(request *SdkGenerateByGroupRequest) (response *SdkGenerateByGroupResponse, err error) {
	response = CreateSdkGenerateByGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SdkGenerateByGroupWithChan invokes the cloudapi.SdkGenerateByGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/sdkgeneratebygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SdkGenerateByGroupWithChan(request *SdkGenerateByGroupRequest) (<-chan *SdkGenerateByGroupResponse, <-chan error) {
	responseChan := make(chan *SdkGenerateByGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SdkGenerateByGroup(request)
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

// SdkGenerateByGroupWithCallback invokes the cloudapi.SdkGenerateByGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/sdkgeneratebygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SdkGenerateByGroupWithCallback(request *SdkGenerateByGroupRequest, callback func(response *SdkGenerateByGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SdkGenerateByGroupResponse
		var err error
		defer close(result)
		response, err = client.SdkGenerateByGroup(request)
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

// SdkGenerateByGroupRequest is the request struct for api SdkGenerateByGroup
type SdkGenerateByGroupRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	Language      string `position:"Query" name:"Language"`
}

// SdkGenerateByGroupResponse is the response struct for api SdkGenerateByGroup
type SdkGenerateByGroupResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DownloadLink string `json:"DownloadLink" xml:"DownloadLink"`
}

// CreateSdkGenerateByGroupRequest creates a request to invoke SdkGenerateByGroup API
func CreateSdkGenerateByGroupRequest() (request *SdkGenerateByGroupRequest) {
	request = &SdkGenerateByGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SdkGenerateByGroup", "apigateway", "openAPI")
	return
}

// CreateSdkGenerateByGroupResponse creates a response to parse from SdkGenerateByGroup response
func CreateSdkGenerateByGroupResponse() (response *SdkGenerateByGroupResponse) {
	response = &SdkGenerateByGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
