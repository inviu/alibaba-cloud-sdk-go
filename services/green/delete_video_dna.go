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

// DeleteVideoDna invokes the green.DeleteVideoDna API synchronously
// api document: https://help.aliyun.com/api/green/deletevideodna.html
func (client *Client) DeleteVideoDna(request *DeleteVideoDnaRequest) (response *DeleteVideoDnaResponse, err error) {
	response = CreateDeleteVideoDnaResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVideoDnaWithChan invokes the green.DeleteVideoDna API asynchronously
// api document: https://help.aliyun.com/api/green/deletevideodna.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVideoDnaWithChan(request *DeleteVideoDnaRequest) (<-chan *DeleteVideoDnaResponse, <-chan error) {
	responseChan := make(chan *DeleteVideoDnaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVideoDna(request)
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

// DeleteVideoDnaWithCallback invokes the green.DeleteVideoDna API asynchronously
// api document: https://help.aliyun.com/api/green/deletevideodna.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVideoDnaWithCallback(request *DeleteVideoDnaRequest, callback func(response *DeleteVideoDnaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVideoDnaResponse
		var err error
		defer close(result)
		response, err = client.DeleteVideoDna(request)
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

// DeleteVideoDnaRequest is the request struct for api DeleteVideoDna
type DeleteVideoDnaRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// DeleteVideoDnaResponse is the response struct for api DeleteVideoDna
type DeleteVideoDnaResponse struct {
	*responses.BaseResponse
}

// CreateDeleteVideoDnaRequest creates a request to invoke DeleteVideoDna API
func CreateDeleteVideoDnaRequest() (request *DeleteVideoDnaRequest) {
	request = &DeleteVideoDnaRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "DeleteVideoDna", "/green/video/dna/delete", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVideoDnaResponse creates a response to parse from DeleteVideoDna response
func CreateDeleteVideoDnaResponse() (response *DeleteVideoDnaResponse) {
	response = &DeleteVideoDnaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
