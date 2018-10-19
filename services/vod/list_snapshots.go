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

// ListSnapshots invokes the vod.ListSnapshots API synchronously
// api document: https://help.aliyun.com/api/vod/listsnapshots.html
func (client *Client) ListSnapshots(request *ListSnapshotsRequest) (response *ListSnapshotsResponse, err error) {
	response = CreateListSnapshotsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSnapshotsWithChan invokes the vod.ListSnapshots API asynchronously
// api document: https://help.aliyun.com/api/vod/listsnapshots.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSnapshotsWithChan(request *ListSnapshotsRequest) (<-chan *ListSnapshotsResponse, <-chan error) {
	responseChan := make(chan *ListSnapshotsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSnapshots(request)
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

// ListSnapshotsWithCallback invokes the vod.ListSnapshots API asynchronously
// api document: https://help.aliyun.com/api/vod/listsnapshots.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSnapshotsWithCallback(request *ListSnapshotsRequest, callback func(response *ListSnapshotsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSnapshotsResponse
		var err error
		defer close(result)
		response, err = client.ListSnapshots(request)
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

// ListSnapshotsRequest is the request struct for api ListSnapshots
type ListSnapshotsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotType         string           `position:"Query" name:"SnapshotType"`
	PageNo               string           `position:"Query" name:"PageNo"`
	PageSize             string           `position:"Query" name:"PageSize"`
	VideoId              string           `position:"Query" name:"VideoId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AuthTimeout          string           `position:"Query" name:"AuthTimeout"`
}

// ListSnapshotsResponse is the response struct for api ListSnapshots
type ListSnapshotsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	MediaSnapshot MediaSnapshot `json:"MediaSnapshot" xml:"MediaSnapshot"`
}

// CreateListSnapshotsRequest creates a request to invoke ListSnapshots API
func CreateListSnapshotsRequest() (request *ListSnapshotsRequest) {
	request = &ListSnapshotsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListSnapshots", "vod", "openAPI")
	return
}

// CreateListSnapshotsResponse creates a response to parse from ListSnapshots response
func CreateListSnapshotsResponse() (response *ListSnapshotsResponse) {
	response = &ListSnapshotsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
