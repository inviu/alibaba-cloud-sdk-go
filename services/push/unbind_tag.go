package push

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

// UnbindTag invokes the push.UnbindTag API synchronously
// api document: https://help.aliyun.com/api/push/unbindtag.html
func (client *Client) UnbindTag(request *UnbindTagRequest) (response *UnbindTagResponse, err error) {
	response = CreateUnbindTagResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindTagWithChan invokes the push.UnbindTag API asynchronously
// api document: https://help.aliyun.com/api/push/unbindtag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindTagWithChan(request *UnbindTagRequest) (<-chan *UnbindTagResponse, <-chan error) {
	responseChan := make(chan *UnbindTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindTag(request)
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

// UnbindTagWithCallback invokes the push.UnbindTag API asynchronously
// api document: https://help.aliyun.com/api/push/unbindtag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindTagWithCallback(request *UnbindTagRequest, callback func(response *UnbindTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindTagResponse
		var err error
		defer close(result)
		response, err = client.UnbindTag(request)
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

// UnbindTagRequest is the request struct for api UnbindTag
type UnbindTagRequest struct {
	*requests.RpcRequest
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	ClientKey string           `position:"Query" name:"ClientKey"`
	KeyType   string           `position:"Query" name:"KeyType"`
	TagName   string           `position:"Query" name:"TagName"`
}

// UnbindTagResponse is the response struct for api UnbindTag
type UnbindTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindTagRequest creates a request to invoke UnbindTag API
func CreateUnbindTagRequest() (request *UnbindTagRequest) {
	request = &UnbindTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "UnbindTag", "", "")
	return
}

// CreateUnbindTagResponse creates a response to parse from UnbindTag response
func CreateUnbindTagResponse() (response *UnbindTagResponse) {
	response = &UnbindTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
