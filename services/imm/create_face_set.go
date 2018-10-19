package imm

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

// CreateFaceSet invokes the imm.CreateFaceSet API synchronously
// api document: https://help.aliyun.com/api/imm/createfaceset.html
func (client *Client) CreateFaceSet(request *CreateFaceSetRequest) (response *CreateFaceSetResponse, err error) {
	response = CreateCreateFaceSetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFaceSetWithChan invokes the imm.CreateFaceSet API asynchronously
// api document: https://help.aliyun.com/api/imm/createfaceset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFaceSetWithChan(request *CreateFaceSetRequest) (<-chan *CreateFaceSetResponse, <-chan error) {
	responseChan := make(chan *CreateFaceSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFaceSet(request)
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

// CreateFaceSetWithCallback invokes the imm.CreateFaceSet API asynchronously
// api document: https://help.aliyun.com/api/imm/createfaceset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFaceSetWithCallback(request *CreateFaceSetRequest, callback func(response *CreateFaceSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFaceSetResponse
		var err error
		defer close(result)
		response, err = client.CreateFaceSet(request)
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

// CreateFaceSetRequest is the request struct for api CreateFaceSet
type CreateFaceSetRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
}

// CreateFaceSetResponse is the response struct for api CreateFaceSet
type CreateFaceSetResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SetId      string `json:"SetId" xml:"SetId"`
	Status     string `json:"Status" xml:"Status"`
	Photos     int    `json:"Photos" xml:"Photos"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime string `json:"ModifyTime" xml:"ModifyTime"`
	Faces      int    `json:"Faces" xml:"Faces"`
}

// CreateCreateFaceSetRequest creates a request to invoke CreateFaceSet API
func CreateCreateFaceSetRequest() (request *CreateFaceSetRequest) {
	request = &CreateFaceSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateFaceSet", "imm", "openAPI")
	return
}

// CreateCreateFaceSetResponse creates a response to parse from CreateFaceSet response
func CreateCreateFaceSetResponse() (response *CreateFaceSetResponse) {
	response = &CreateFaceSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
