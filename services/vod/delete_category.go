package vod

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

// DeleteCategory invokes the vod.DeleteCategory API synchronously
// api document: https://help.aliyun.com/api/vod/deletecategory.html
func (client *Client) DeleteCategory(request *DeleteCategoryRequest) (response *DeleteCategoryResponse, err error) {
	response = CreateDeleteCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCategoryWithChan invokes the vod.DeleteCategory API asynchronously
// api document: https://help.aliyun.com/api/vod/deletecategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCategoryWithChan(request *DeleteCategoryRequest) (<-chan *DeleteCategoryResponse, <-chan error) {
	responseChan := make(chan *DeleteCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCategory(request)
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

// DeleteCategoryWithCallback invokes the vod.DeleteCategory API asynchronously
// api document: https://help.aliyun.com/api/vod/deletecategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCategoryWithCallback(request *DeleteCategoryRequest, callback func(response *DeleteCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCategoryResponse
		var err error
		defer close(result)
		response, err = client.DeleteCategory(request)
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

// DeleteCategoryRequest is the request struct for api DeleteCategory
type DeleteCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CateId               requests.Integer `position:"Query" name:"CateId"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
}

// DeleteCategoryResponse is the response struct for api DeleteCategory
type DeleteCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCategoryRequest creates a request to invoke DeleteCategory API
func CreateDeleteCategoryRequest() (request *DeleteCategoryRequest) {
	request = &DeleteCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteCategory", "vod", "openAPI")
	return
}

// CreateDeleteCategoryResponse creates a response to parse from DeleteCategory response
func CreateDeleteCategoryResponse() (response *DeleteCategoryResponse) {
	response = &DeleteCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
