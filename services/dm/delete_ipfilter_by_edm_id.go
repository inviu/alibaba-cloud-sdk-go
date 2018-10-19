package dm

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

// DeleteIpfilterByEdmId invokes the dm.DeleteIpfilterByEdmId API synchronously
// api document: https://help.aliyun.com/api/dm/deleteipfilterbyedmid.html
func (client *Client) DeleteIpfilterByEdmId(request *DeleteIpfilterByEdmIdRequest) (response *DeleteIpfilterByEdmIdResponse, err error) {
	response = CreateDeleteIpfilterByEdmIdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIpfilterByEdmIdWithChan invokes the dm.DeleteIpfilterByEdmId API asynchronously
// api document: https://help.aliyun.com/api/dm/deleteipfilterbyedmid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteIpfilterByEdmIdWithChan(request *DeleteIpfilterByEdmIdRequest) (<-chan *DeleteIpfilterByEdmIdResponse, <-chan error) {
	responseChan := make(chan *DeleteIpfilterByEdmIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIpfilterByEdmId(request)
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

// DeleteIpfilterByEdmIdWithCallback invokes the dm.DeleteIpfilterByEdmId API asynchronously
// api document: https://help.aliyun.com/api/dm/deleteipfilterbyedmid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteIpfilterByEdmIdWithCallback(request *DeleteIpfilterByEdmIdRequest, callback func(response *DeleteIpfilterByEdmIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIpfilterByEdmIdResponse
		var err error
		defer close(result)
		response, err = client.DeleteIpfilterByEdmId(request)
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

// DeleteIpfilterByEdmIdRequest is the request struct for api DeleteIpfilterByEdmId
type DeleteIpfilterByEdmIdRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
	Id                   string           `position:"Query" name:"Id"`
}

// DeleteIpfilterByEdmIdResponse is the response struct for api DeleteIpfilterByEdmId
type DeleteIpfilterByEdmIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteIpfilterByEdmIdRequest creates a request to invoke DeleteIpfilterByEdmId API
func CreateDeleteIpfilterByEdmIdRequest() (request *DeleteIpfilterByEdmIdRequest) {
	request = &DeleteIpfilterByEdmIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DeleteIpfilterByEdmId", "", "")
	return
}

// CreateDeleteIpfilterByEdmIdResponse creates a response to parse from DeleteIpfilterByEdmId response
func CreateDeleteIpfilterByEdmIdResponse() (response *DeleteIpfilterByEdmIdResponse) {
	response = &DeleteIpfilterByEdmIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
