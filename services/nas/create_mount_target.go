package nas

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

// CreateMountTarget invokes the nas.CreateMountTarget API synchronously
// api document: https://help.aliyun.com/api/nas/createmounttarget.html
func (client *Client) CreateMountTarget(request *CreateMountTargetRequest) (response *CreateMountTargetResponse, err error) {
	response = CreateCreateMountTargetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMountTargetWithChan invokes the nas.CreateMountTarget API asynchronously
// api document: https://help.aliyun.com/api/nas/createmounttarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMountTargetWithChan(request *CreateMountTargetRequest) (<-chan *CreateMountTargetResponse, <-chan error) {
	responseChan := make(chan *CreateMountTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMountTarget(request)
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

// CreateMountTargetWithCallback invokes the nas.CreateMountTarget API asynchronously
// api document: https://help.aliyun.com/api/nas/createmounttarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMountTargetWithCallback(request *CreateMountTargetRequest, callback func(response *CreateMountTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMountTargetResponse
		var err error
		defer close(result)
		response, err = client.CreateMountTarget(request)
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

// CreateMountTargetRequest is the request struct for api CreateMountTarget
type CreateMountTargetRequest struct {
	*requests.RpcRequest
	FileSystemId    string `position:"Query" name:"FileSystemId"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	NetworkType     string `position:"Query" name:"NetworkType"`
	VpcId           string `position:"Query" name:"VpcId"`
	VSwitchId       string `position:"Query" name:"VSwitchId"`
}

// CreateMountTargetResponse is the response struct for api CreateMountTarget
type CreateMountTargetResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	MountTargetDomain string `json:"MountTargetDomain" xml:"MountTargetDomain"`
}

// CreateCreateMountTargetRequest creates a request to invoke CreateMountTarget API
func CreateCreateMountTargetRequest() (request *CreateMountTargetRequest) {
	request = &CreateMountTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateMountTarget", "nas", "openAPI")
	return
}

// CreateCreateMountTargetResponse creates a response to parse from CreateMountTarget response
func CreateCreateMountTargetResponse() (response *CreateMountTargetResponse) {
	response = &CreateMountTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
