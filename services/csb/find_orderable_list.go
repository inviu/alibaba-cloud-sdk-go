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

// FindOrderableList invokes the csb.FindOrderableList API synchronously
// api document: https://help.aliyun.com/api/csb/findorderablelist.html
func (client *Client) FindOrderableList(request *FindOrderableListRequest) (response *FindOrderableListResponse, err error) {
	response = CreateFindOrderableListResponse()
	err = client.DoAction(request, response)
	return
}

// FindOrderableListWithChan invokes the csb.FindOrderableList API asynchronously
// api document: https://help.aliyun.com/api/csb/findorderablelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindOrderableListWithChan(request *FindOrderableListRequest) (<-chan *FindOrderableListResponse, <-chan error) {
	responseChan := make(chan *FindOrderableListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindOrderableList(request)
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

// FindOrderableListWithCallback invokes the csb.FindOrderableList API asynchronously
// api document: https://help.aliyun.com/api/csb/findorderablelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindOrderableListWithCallback(request *FindOrderableListRequest, callback func(response *FindOrderableListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindOrderableListResponse
		var err error
		defer close(result)
		response, err = client.FindOrderableList(request)
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

// FindOrderableListRequest is the request struct for api FindOrderableList
type FindOrderableListRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Query" name:"ProjectName"`
	CsbId       requests.Integer `position:"Query" name:"CsbId"`
	Alias       string           `position:"Query" name:"Alias"`
	ServiceName string           `position:"Query" name:"ServiceName"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
}

// FindOrderableListResponse is the response struct for api FindOrderableList
type FindOrderableListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindOrderableListRequest creates a request to invoke FindOrderableList API
func CreateFindOrderableListRequest() (request *FindOrderableListRequest) {
	request = &FindOrderableListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindOrderableList", "", "")
	return
}

// CreateFindOrderableListResponse creates a response to parse from FindOrderableList response
func CreateFindOrderableListResponse() (response *FindOrderableListResponse) {
	response = &FindOrderableListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
