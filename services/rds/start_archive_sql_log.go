package rds

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

// StartArchiveSQLLog invokes the rds.StartArchiveSQLLog API synchronously
// api document: https://help.aliyun.com/api/rds/startarchivesqllog.html
func (client *Client) StartArchiveSQLLog(request *StartArchiveSQLLogRequest) (response *StartArchiveSQLLogResponse, err error) {
	response = CreateStartArchiveSQLLogResponse()
	err = client.DoAction(request, response)
	return
}

// StartArchiveSQLLogWithChan invokes the rds.StartArchiveSQLLog API asynchronously
// api document: https://help.aliyun.com/api/rds/startarchivesqllog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartArchiveSQLLogWithChan(request *StartArchiveSQLLogRequest) (<-chan *StartArchiveSQLLogResponse, <-chan error) {
	responseChan := make(chan *StartArchiveSQLLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartArchiveSQLLog(request)
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

// StartArchiveSQLLogWithCallback invokes the rds.StartArchiveSQLLog API asynchronously
// api document: https://help.aliyun.com/api/rds/startarchivesqllog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartArchiveSQLLogWithCallback(request *StartArchiveSQLLogRequest, callback func(response *StartArchiveSQLLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartArchiveSQLLogResponse
		var err error
		defer close(result)
		response, err = client.StartArchiveSQLLog(request)
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

// StartArchiveSQLLogRequest is the request struct for api StartArchiveSQLLog
type StartArchiveSQLLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Database             string           `position:"Query" name:"Database"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	User                 string           `position:"Query" name:"User"`
	QueryKeywords        string           `position:"Query" name:"QueryKeywords"`
}

// StartArchiveSQLLogResponse is the response struct for api StartArchiveSQLLog
type StartArchiveSQLLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartArchiveSQLLogRequest creates a request to invoke StartArchiveSQLLog API
func CreateStartArchiveSQLLogRequest() (request *StartArchiveSQLLogRequest) {
	request = &StartArchiveSQLLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "StartArchiveSQLLog", "rds", "openAPI")
	return
}

// CreateStartArchiveSQLLogResponse creates a response to parse from StartArchiveSQLLog response
func CreateStartArchiveSQLLogResponse() (response *StartArchiveSQLLogResponse) {
	response = &StartArchiveSQLLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
