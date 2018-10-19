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

// DescribeResourceDiagnosis invokes the rds.DescribeResourceDiagnosis API synchronously
// api document: https://help.aliyun.com/api/rds/describeresourcediagnosis.html
func (client *Client) DescribeResourceDiagnosis(request *DescribeResourceDiagnosisRequest) (response *DescribeResourceDiagnosisResponse, err error) {
	response = CreateDescribeResourceDiagnosisResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceDiagnosisWithChan invokes the rds.DescribeResourceDiagnosis API asynchronously
// api document: https://help.aliyun.com/api/rds/describeresourcediagnosis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceDiagnosisWithChan(request *DescribeResourceDiagnosisRequest) (<-chan *DescribeResourceDiagnosisResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceDiagnosisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceDiagnosis(request)
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

// DescribeResourceDiagnosisWithCallback invokes the rds.DescribeResourceDiagnosis API asynchronously
// api document: https://help.aliyun.com/api/rds/describeresourcediagnosis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceDiagnosisWithCallback(request *DescribeResourceDiagnosisRequest, callback func(response *DescribeResourceDiagnosisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceDiagnosisResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceDiagnosis(request)
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

// DescribeResourceDiagnosisRequest is the request struct for api DescribeResourceDiagnosis
type DescribeResourceDiagnosisRequest struct {
	*requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

// DescribeResourceDiagnosisResponse is the response struct for api DescribeResourceDiagnosis
type DescribeResourceDiagnosisResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	StartTime  string   `json:"StartTime" xml:"StartTime"`
	EndTime    string   `json:"EndTime" xml:"EndTime"`
	CPU        []string `json:"CPU" xml:"CPU"`
	Memory     []string `json:"Memory" xml:"Memory"`
	Storage    []string `json:"Storage" xml:"Storage"`
	IOPS       []string `json:"IOPS" xml:"IOPS"`
	Connection []string `json:"Connection" xml:"Connection"`
}

// CreateDescribeResourceDiagnosisRequest creates a request to invoke DescribeResourceDiagnosis API
func CreateDescribeResourceDiagnosisRequest() (request *DescribeResourceDiagnosisRequest) {
	request = &DescribeResourceDiagnosisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeResourceDiagnosis", "rds", "openAPI")
	return
}

// CreateDescribeResourceDiagnosisResponse creates a response to parse from DescribeResourceDiagnosis response
func CreateDescribeResourceDiagnosisResponse() (response *DescribeResourceDiagnosisResponse) {
	response = &DescribeResourceDiagnosisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
