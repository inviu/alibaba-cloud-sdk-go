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

// RegistFace invokes the imm.RegistFace API synchronously
// api document: https://help.aliyun.com/api/imm/registface.html
func (client *Client) RegistFace(request *RegistFaceRequest) (response *RegistFaceResponse, err error) {
	response = CreateRegistFaceResponse()
	err = client.DoAction(request, response)
	return
}

// RegistFaceWithChan invokes the imm.RegistFace API asynchronously
// api document: https://help.aliyun.com/api/imm/registface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegistFaceWithChan(request *RegistFaceRequest) (<-chan *RegistFaceResponse, <-chan error) {
	responseChan := make(chan *RegistFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegistFace(request)
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

// RegistFaceWithCallback invokes the imm.RegistFace API asynchronously
// api document: https://help.aliyun.com/api/imm/registface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegistFaceWithCallback(request *RegistFaceRequest, callback func(response *RegistFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegistFaceResponse
		var err error
		defer close(result)
		response, err = client.RegistFace(request)
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

// RegistFaceRequest is the request struct for api RegistFace
type RegistFaceRequest struct {
	*requests.RpcRequest
	ChooseBiggestFace  requests.Boolean `position:"Query" name:"ChooseBiggestFace"`
	IsQualityLimit     requests.Boolean `position:"Query" name:"IsQualityLimit"`
	Project            string           `position:"Query" name:"Project"`
	SrcUri             string           `position:"Query" name:"SrcUri"`
	RegisterCheckLevel string           `position:"Query" name:"RegisterCheckLevel"`
	GroupName          string           `position:"Query" name:"GroupName"`
	User               string           `position:"Query" name:"User"`
}

// RegistFaceResponse is the response struct for api RegistFace
type RegistFaceResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	GroupId   string   `json:"GroupId" xml:"GroupId"`
	ImageUrl  string   `json:"ImageUrl" xml:"ImageUrl"`
	ImageMd5  string   `json:"ImageMd5" xml:"ImageMd5"`
	ImageId   string   `json:"ImageId" xml:"ImageId"`
	GroupName string   `json:"GroupName" xml:"GroupName"`
	User      string   `json:"User" xml:"User"`
	Roll      float64  `json:"Roll" xml:"Roll"`
	FaceId    string   `json:"FaceId" xml:"FaceId"`
	Yaw       string   `json:"Yaw" xml:"Yaw"`
	Quality   float64  `json:"Quality" xml:"Quality"`
	Glasses   int      `json:"Glasses" xml:"Glasses"`
	Hat       int      `json:"Hat" xml:"Hat"`
	Pitch     float64  `json:"Pitch" xml:"Pitch"`
	Age       int      `json:"Age" xml:"Age"`
	Gender    string   `json:"Gender" xml:"Gender"`
	Axis      []string `json:"Axis" xml:"Axis"`
}

// CreateRegistFaceRequest creates a request to invoke RegistFace API
func CreateRegistFaceRequest() (request *RegistFaceRequest) {
	request = &RegistFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "RegistFace", "imm", "openAPI")
	return
}

// CreateRegistFaceResponse creates a response to parse from RegistFace response
func CreateRegistFaceResponse() (response *RegistFaceResponse) {
	response = &RegistFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
