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

// CreateOfficeConversionTask invokes the imm.CreateOfficeConversionTask API synchronously
// api document: https://help.aliyun.com/api/imm/createofficeconversiontask.html
func (client *Client) CreateOfficeConversionTask(request *CreateOfficeConversionTaskRequest) (response *CreateOfficeConversionTaskResponse, err error) {
	response = CreateCreateOfficeConversionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOfficeConversionTaskWithChan invokes the imm.CreateOfficeConversionTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createofficeconversiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOfficeConversionTaskWithChan(request *CreateOfficeConversionTaskRequest) (<-chan *CreateOfficeConversionTaskResponse, <-chan error) {
	responseChan := make(chan *CreateOfficeConversionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOfficeConversionTask(request)
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

// CreateOfficeConversionTaskWithCallback invokes the imm.CreateOfficeConversionTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createofficeconversiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOfficeConversionTaskWithCallback(request *CreateOfficeConversionTaskRequest, callback func(response *CreateOfficeConversionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOfficeConversionTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateOfficeConversionTask(request)
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

// CreateOfficeConversionTaskRequest is the request struct for api CreateOfficeConversionTask
type CreateOfficeConversionTaskRequest struct {
	*requests.RpcRequest
	ImageSpec       string           `position:"Query" name:"ImageSpec"`
	SrcType         string           `position:"Query" name:"SrcType"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	ModelId         string           `position:"Query" name:"ModelId"`
	Project         string           `position:"Query" name:"Project"`
	ExternalID      string           `position:"Query" name:"ExternalID"`
	MaxSheetRow     requests.Integer `position:"Query" name:"MaxSheetRow"`
	MaxSheetCount   requests.Integer `position:"Query" name:"MaxSheetCount"`
	EndPage         requests.Integer `position:"Query" name:"EndPage"`
	TgtFileSuffix   string           `position:"Query" name:"TgtFileSuffix"`
	SheetOnePage    requests.Boolean `position:"Query" name:"SheetOnePage"`
	Password        string           `position:"Query" name:"Password"`
	StartPage       requests.Integer `position:"Query" name:"StartPage"`
	MaxSheetCol     requests.Integer `position:"Query" name:"MaxSheetCol"`
	TgtType         string           `position:"Query" name:"TgtType"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	TgtFilePrefix   string           `position:"Query" name:"TgtFilePrefix"`
	SrcUri          string           `position:"Query" name:"SrcUri"`
	TgtFilePages    string           `position:"Query" name:"TgtFilePages"`
	TgtUri          string           `position:"Query" name:"TgtUri"`
}

// CreateOfficeConversionTaskResponse is the response struct for api CreateOfficeConversionTask
type CreateOfficeConversionTaskResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
	TgtLoc     string `json:"TgtLoc" xml:"TgtLoc"`
	Status     string `json:"Status" xml:"Status"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	Percent    int    `json:"Percent" xml:"Percent"`
}

// CreateCreateOfficeConversionTaskRequest creates a request to invoke CreateOfficeConversionTask API
func CreateCreateOfficeConversionTaskRequest() (request *CreateOfficeConversionTaskRequest) {
	request = &CreateOfficeConversionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateOfficeConversionTask", "imm", "openAPI")
	return
}

// CreateCreateOfficeConversionTaskResponse creates a response to parse from CreateOfficeConversionTask response
func CreateCreateOfficeConversionTaskResponse() (response *CreateOfficeConversionTaskResponse) {
	response = &CreateOfficeConversionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
