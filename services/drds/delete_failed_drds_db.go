package drds

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

// DeleteFailedDrdsDB invokes the drds.DeleteFailedDrdsDB API synchronously
// api document: https://help.aliyun.com/api/drds/deletefaileddrdsdb.html
func (client *Client) DeleteFailedDrdsDB(request *DeleteFailedDrdsDBRequest) (response *DeleteFailedDrdsDBResponse, err error) {
	response = CreateDeleteFailedDrdsDBResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFailedDrdsDBWithChan invokes the drds.DeleteFailedDrdsDB API asynchronously
// api document: https://help.aliyun.com/api/drds/deletefaileddrdsdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFailedDrdsDBWithChan(request *DeleteFailedDrdsDBRequest) (<-chan *DeleteFailedDrdsDBResponse, <-chan error) {
	responseChan := make(chan *DeleteFailedDrdsDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFailedDrdsDB(request)
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

// DeleteFailedDrdsDBWithCallback invokes the drds.DeleteFailedDrdsDB API asynchronously
// api document: https://help.aliyun.com/api/drds/deletefaileddrdsdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFailedDrdsDBWithCallback(request *DeleteFailedDrdsDBRequest, callback func(response *DeleteFailedDrdsDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFailedDrdsDBResponse
		var err error
		defer close(result)
		response, err = client.DeleteFailedDrdsDB(request)
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

// DeleteFailedDrdsDBRequest is the request struct for api DeleteFailedDrdsDB
type DeleteFailedDrdsDBRequest struct {
	*requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DeleteFailedDrdsDBResponse is the response struct for api DeleteFailedDrdsDB
type DeleteFailedDrdsDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteFailedDrdsDBRequest creates a request to invoke DeleteFailedDrdsDB API
func CreateDeleteFailedDrdsDBRequest() (request *DeleteFailedDrdsDBRequest) {
	request = &DeleteFailedDrdsDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "DeleteFailedDrdsDB", "", "")
	return
}

// CreateDeleteFailedDrdsDBResponse creates a response to parse from DeleteFailedDrdsDB response
func CreateDeleteFailedDrdsDBResponse() (response *DeleteFailedDrdsDBResponse) {
	response = &DeleteFailedDrdsDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
