package r_kvstore

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

// RestoreSnapshot invokes the r_kvstore.RestoreSnapshot API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/restoresnapshot.html
func (client *Client) RestoreSnapshot(request *RestoreSnapshotRequest) (response *RestoreSnapshotResponse, err error) {
	response = CreateRestoreSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// RestoreSnapshotWithChan invokes the r_kvstore.RestoreSnapshot API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/restoresnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestoreSnapshotWithChan(request *RestoreSnapshotRequest) (<-chan *RestoreSnapshotResponse, <-chan error) {
	responseChan := make(chan *RestoreSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestoreSnapshot(request)
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

// RestoreSnapshotWithCallback invokes the r_kvstore.RestoreSnapshot API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/restoresnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestoreSnapshotWithCallback(request *RestoreSnapshotRequest, callback func(response *RestoreSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestoreSnapshotResponse
		var err error
		defer close(result)
		response, err = client.RestoreSnapshot(request)
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

// RestoreSnapshotRequest is the request struct for api RestoreSnapshot
type RestoreSnapshotRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RestoreSnapshotResponse is the response struct for api RestoreSnapshot
type RestoreSnapshotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestoreSnapshotRequest creates a request to invoke RestoreSnapshot API
func CreateRestoreSnapshotRequest() (request *RestoreSnapshotRequest) {
	request = &RestoreSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "RestoreSnapshot", "redisa", "openAPI")
	return
}

// CreateRestoreSnapshotResponse creates a response to parse from RestoreSnapshot response
func CreateRestoreSnapshotResponse() (response *RestoreSnapshotResponse) {
	response = &RestoreSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
