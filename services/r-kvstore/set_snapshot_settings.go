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

// SetSnapshotSettings invokes the r_kvstore.SetSnapshotSettings API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/setsnapshotsettings.html
func (client *Client) SetSnapshotSettings(request *SetSnapshotSettingsRequest) (response *SetSnapshotSettingsResponse, err error) {
	response = CreateSetSnapshotSettingsResponse()
	err = client.DoAction(request, response)
	return
}

// SetSnapshotSettingsWithChan invokes the r_kvstore.SetSnapshotSettings API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/setsnapshotsettings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetSnapshotSettingsWithChan(request *SetSnapshotSettingsRequest) (<-chan *SetSnapshotSettingsResponse, <-chan error) {
	responseChan := make(chan *SetSnapshotSettingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetSnapshotSettings(request)
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

// SetSnapshotSettingsWithCallback invokes the r_kvstore.SetSnapshotSettings API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/setsnapshotsettings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetSnapshotSettingsWithCallback(request *SetSnapshotSettingsRequest, callback func(response *SetSnapshotSettingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetSnapshotSettingsResponse
		var err error
		defer close(result)
		response, err = client.SetSnapshotSettings(request)
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

// SetSnapshotSettingsRequest is the request struct for api SetSnapshotSettings
type SetSnapshotSettingsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// SetSnapshotSettingsResponse is the response struct for api SetSnapshotSettings
type SetSnapshotSettingsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetSnapshotSettingsRequest creates a request to invoke SetSnapshotSettings API
func CreateSetSnapshotSettingsRequest() (request *SetSnapshotSettingsRequest) {
	request = &SetSnapshotSettingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "SetSnapshotSettings", "redisa", "openAPI")
	return
}

// CreateSetSnapshotSettingsResponse creates a response to parse from SetSnapshotSettings response
func CreateSetSnapshotSettingsResponse() (response *SetSnapshotSettingsResponse) {
	response = &SetSnapshotSettingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
