package live

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

// DeleteLiveRecordVodConfig invokes the live.DeleteLiveRecordVodConfig API synchronously
// api document: https://help.aliyun.com/api/live/deleteliverecordvodconfig.html
func (client *Client) DeleteLiveRecordVodConfig(request *DeleteLiveRecordVodConfigRequest) (response *DeleteLiveRecordVodConfigResponse, err error) {
	response = CreateDeleteLiveRecordVodConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveRecordVodConfigWithChan invokes the live.DeleteLiveRecordVodConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deleteliverecordvodconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveRecordVodConfigWithChan(request *DeleteLiveRecordVodConfigRequest) (<-chan *DeleteLiveRecordVodConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveRecordVodConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveRecordVodConfig(request)
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

// DeleteLiveRecordVodConfigWithCallback invokes the live.DeleteLiveRecordVodConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deleteliverecordvodconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveRecordVodConfigWithCallback(request *DeleteLiveRecordVodConfigRequest, callback func(response *DeleteLiveRecordVodConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveRecordVodConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveRecordVodConfig(request)
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

// DeleteLiveRecordVodConfigRequest is the request struct for api DeleteLiveRecordVodConfig
type DeleteLiveRecordVodConfigRequest struct {
	*requests.RpcRequest
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	StreamName    string           `position:"Query" name:"StreamName"`
}

// DeleteLiveRecordVodConfigResponse is the response struct for api DeleteLiveRecordVodConfig
type DeleteLiveRecordVodConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveRecordVodConfigRequest creates a request to invoke DeleteLiveRecordVodConfig API
func CreateDeleteLiveRecordVodConfigRequest() (request *DeleteLiveRecordVodConfigRequest) {
	request = &DeleteLiveRecordVodConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveRecordVodConfig", "live", "openAPI")
	return
}

// CreateDeleteLiveRecordVodConfigResponse creates a response to parse from DeleteLiveRecordVodConfig response
func CreateDeleteLiveRecordVodConfigResponse() (response *DeleteLiveRecordVodConfigResponse) {
	response = &DeleteLiveRecordVodConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
