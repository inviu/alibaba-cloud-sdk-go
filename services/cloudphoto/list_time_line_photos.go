package cloudphoto

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

// ListTimeLinePhotos invokes the cloudphoto.ListTimeLinePhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/listtimelinephotos.html
func (client *Client) ListTimeLinePhotos(request *ListTimeLinePhotosRequest) (response *ListTimeLinePhotosResponse, err error) {
	response = CreateListTimeLinePhotosResponse()
	err = client.DoAction(request, response)
	return
}

// ListTimeLinePhotosWithChan invokes the cloudphoto.ListTimeLinePhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/listtimelinephotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTimeLinePhotosWithChan(request *ListTimeLinePhotosRequest) (<-chan *ListTimeLinePhotosResponse, <-chan error) {
	responseChan := make(chan *ListTimeLinePhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTimeLinePhotos(request)
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

// ListTimeLinePhotosWithCallback invokes the cloudphoto.ListTimeLinePhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/listtimelinephotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTimeLinePhotosWithCallback(request *ListTimeLinePhotosRequest, callback func(response *ListTimeLinePhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTimeLinePhotosResponse
		var err error
		defer close(result)
		response, err = client.ListTimeLinePhotos(request)
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

// ListTimeLinePhotosRequest is the request struct for api ListTimeLinePhotos
type ListTimeLinePhotosRequest struct {
	*requests.RpcRequest
	Direction string           `position:"Query" name:"Direction"`
	Page      requests.Integer `position:"Query" name:"Page"`
	Size      requests.Integer `position:"Query" name:"Size"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	FilterBy  string           `position:"Query" name:"FilterBy"`
	Order     string           `position:"Query" name:"Order"`
	StoreName string           `position:"Query" name:"StoreName"`
	LibraryId string           `position:"Query" name:"LibraryId"`
}

// ListTimeLinePhotosResponse is the response struct for api ListTimeLinePhotos
type ListTimeLinePhotosResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	Action     string  `json:"Action" xml:"Action"`
	Photos     []Photo `json:"Photos" xml:"Photos"`
}

// CreateListTimeLinePhotosRequest creates a request to invoke ListTimeLinePhotos API
func CreateListTimeLinePhotosRequest() (request *ListTimeLinePhotosRequest) {
	request = &ListTimeLinePhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTimeLinePhotos", "cloudphoto", "openAPI")
	return
}

// CreateListTimeLinePhotosResponse creates a response to parse from ListTimeLinePhotos response
func CreateListTimeLinePhotosResponse() (response *ListTimeLinePhotosResponse) {
	response = &ListTimeLinePhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
