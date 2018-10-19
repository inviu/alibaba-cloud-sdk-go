package imm

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

// ListTagPhotos invokes the imm.ListTagPhotos API synchronously
// api document: https://help.aliyun.com/api/imm/listtagphotos.html
func (client *Client) ListTagPhotos(request *ListTagPhotosRequest) (response *ListTagPhotosResponse, err error) {
	response = CreateListTagPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagPhotosWithChan invokes the imm.ListTagPhotos API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagPhotosWithChan(request *ListTagPhotosRequest) (<-chan *ListTagPhotosResponse, <-chan error) {
	responseChan := make(chan *ListTagPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagPhotos(request)
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

// ListTagPhotosWithCallback invokes the imm.ListTagPhotos API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagPhotosWithCallback(request *ListTagPhotosRequest, callback func(response *ListTagPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagPhotosResponse
		var err error
		defer close(result)
		response, err = client.ListTagPhotos(request)
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

// ListTagPhotosRequest is the request struct for api ListTagPhotos
type ListTagPhotosRequest struct {
	*requests.RpcRequest
	TagName string `position:"Query" name:"TagName"`
	MaxKeys string `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

// ListTagPhotosResponse is the response struct for api ListTagPhotos
type ListTagPhotosResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	NextMarker string       `json:"NextMarker" xml:"NextMarker"`
	Photos     []PhotosItem `json:"Photos" xml:"Photos"`
}

// CreateListTagPhotosRequest creates a request to invoke ListTagPhotos API
func CreateListTagPhotosRequest() (request *ListTagPhotosRequest) {
	request = &ListTagPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListTagPhotos", "imm", "openAPI")
	return
}

// CreateListTagPhotosResponse creates a response to parse from ListTagPhotos response
func CreateListTagPhotosResponse() (response *ListTagPhotosResponse) {
	response = &ListTagPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
