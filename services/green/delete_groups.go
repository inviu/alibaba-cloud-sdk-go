package green

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

// DeleteGroups invokes the green.DeleteGroups API synchronously
// api document: https://help.aliyun.com/api/green/deletegroups.html
func (client *Client) DeleteGroups(request *DeleteGroupsRequest) (response *DeleteGroupsResponse, err error) {
	response = CreateDeleteGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGroupsWithChan invokes the green.DeleteGroups API asynchronously
// api document: https://help.aliyun.com/api/green/deletegroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteGroupsWithChan(request *DeleteGroupsRequest) (<-chan *DeleteGroupsResponse, <-chan error) {
	responseChan := make(chan *DeleteGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGroups(request)
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

// DeleteGroupsWithCallback invokes the green.DeleteGroups API asynchronously
// api document: https://help.aliyun.com/api/green/deletegroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteGroupsWithCallback(request *DeleteGroupsRequest, callback func(response *DeleteGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGroupsResponse
		var err error
		defer close(result)
		response, err = client.DeleteGroups(request)
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

// DeleteGroupsRequest is the request struct for api DeleteGroups
type DeleteGroupsRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// DeleteGroupsResponse is the response struct for api DeleteGroups
type DeleteGroupsResponse struct {
	*responses.BaseResponse
}

// CreateDeleteGroupsRequest creates a request to invoke DeleteGroups API
func CreateDeleteGroupsRequest() (request *DeleteGroupsRequest) {
	request = &DeleteGroupsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "DeleteGroups", "/green/sface/person/groups/delete", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGroupsResponse creates a response to parse from DeleteGroups response
func CreateDeleteGroupsResponse() (response *DeleteGroupsResponse) {
	response = &DeleteGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
