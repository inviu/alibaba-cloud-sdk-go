package dcdn

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

// RefreshDcdnObjectCaches invokes the dcdn.RefreshDcdnObjectCaches API synchronously
// api document: https://help.aliyun.com/api/dcdn/refreshdcdnobjectcaches.html
func (client *Client) RefreshDcdnObjectCaches(request *RefreshDcdnObjectCachesRequest) (response *RefreshDcdnObjectCachesResponse, err error) {
	response = CreateRefreshDcdnObjectCachesResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshDcdnObjectCachesWithChan invokes the dcdn.RefreshDcdnObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/dcdn/refreshdcdnobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshDcdnObjectCachesWithChan(request *RefreshDcdnObjectCachesRequest) (<-chan *RefreshDcdnObjectCachesResponse, <-chan error) {
	responseChan := make(chan *RefreshDcdnObjectCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshDcdnObjectCaches(request)
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

// RefreshDcdnObjectCachesWithCallback invokes the dcdn.RefreshDcdnObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/dcdn/refreshdcdnobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshDcdnObjectCachesWithCallback(request *RefreshDcdnObjectCachesRequest, callback func(response *RefreshDcdnObjectCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshDcdnObjectCachesResponse
		var err error
		defer close(result)
		response, err = client.RefreshDcdnObjectCaches(request)
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

// RefreshDcdnObjectCachesRequest is the request struct for api RefreshDcdnObjectCaches
type RefreshDcdnObjectCachesRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	ObjectType    string           `position:"Query" name:"ObjectType"`
}

// RefreshDcdnObjectCachesResponse is the response struct for api RefreshDcdnObjectCaches
type RefreshDcdnObjectCachesResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	RefreshTaskId string `json:"RefreshTaskId" xml:"RefreshTaskId"`
}

// CreateRefreshDcdnObjectCachesRequest creates a request to invoke RefreshDcdnObjectCaches API
func CreateRefreshDcdnObjectCachesRequest() (request *RefreshDcdnObjectCachesRequest) {
	request = &RefreshDcdnObjectCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "RefreshDcdnObjectCaches", "dcdn", "openAPI")
	return
}

// CreateRefreshDcdnObjectCachesResponse creates a response to parse from RefreshDcdnObjectCaches response
func CreateRefreshDcdnObjectCachesResponse() (response *RefreshDcdnObjectCachesResponse) {
	response = &RefreshDcdnObjectCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
