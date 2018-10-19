package csb

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

// FindOrderedList invokes the csb.FindOrderedList API synchronously
// api document: https://help.aliyun.com/api/csb/findorderedlist.html
func (client *Client) FindOrderedList(request *FindOrderedListRequest) (response *FindOrderedListResponse, err error) {
	response = CreateFindOrderedListResponse()
	err = client.DoAction(request, response)
	return
}

// FindOrderedListWithChan invokes the csb.FindOrderedList API asynchronously
// api document: https://help.aliyun.com/api/csb/findorderedlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindOrderedListWithChan(request *FindOrderedListRequest) (<-chan *FindOrderedListResponse, <-chan error) {
	responseChan := make(chan *FindOrderedListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindOrderedList(request)
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

// FindOrderedListWithCallback invokes the csb.FindOrderedList API asynchronously
// api document: https://help.aliyun.com/api/csb/findorderedlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindOrderedListWithCallback(request *FindOrderedListRequest, callback func(response *FindOrderedListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindOrderedListResponse
		var err error
		defer close(result)
		response, err = client.FindOrderedList(request)
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

// FindOrderedListRequest is the request struct for api FindOrderedList
type FindOrderedListRequest struct {
	*requests.RpcRequest
	ProjectName  string           `position:"Query" name:"ProjectName"`
	ShowDelOrder requests.Boolean `position:"Query" name:"ShowDelOrder"`
	CsbId        requests.Integer `position:"Query" name:"CsbId"`
	Alias        string           `position:"Query" name:"Alias"`
	ServiceName  string           `position:"Query" name:"ServiceName"`
	PageNum      requests.Integer `position:"Query" name:"PageNum"`
	ServiceId    requests.Integer `position:"Query" name:"ServiceId"`
	Status       string           `position:"Query" name:"Status"`
}

// FindOrderedListResponse is the response struct for api FindOrderedList
type FindOrderedListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindOrderedListRequest creates a request to invoke FindOrderedList API
func CreateFindOrderedListRequest() (request *FindOrderedListRequest) {
	request = &FindOrderedListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindOrderedList", "", "")
	return
}

// CreateFindOrderedListResponse creates a response to parse from FindOrderedList response
func CreateFindOrderedListResponse() (response *FindOrderedListResponse) {
	response = &FindOrderedListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
