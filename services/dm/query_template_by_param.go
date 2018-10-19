package dm

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

// QueryTemplateByParam invokes the dm.QueryTemplateByParam API synchronously
// api document: https://help.aliyun.com/api/dm/querytemplatebyparam.html
func (client *Client) QueryTemplateByParam(request *QueryTemplateByParamRequest) (response *QueryTemplateByParamResponse, err error) {
	response = CreateQueryTemplateByParamResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTemplateByParamWithChan invokes the dm.QueryTemplateByParam API asynchronously
// api document: https://help.aliyun.com/api/dm/querytemplatebyparam.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTemplateByParamWithChan(request *QueryTemplateByParamRequest) (<-chan *QueryTemplateByParamResponse, <-chan error) {
	responseChan := make(chan *QueryTemplateByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTemplateByParam(request)
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

// QueryTemplateByParamWithCallback invokes the dm.QueryTemplateByParam API asynchronously
// api document: https://help.aliyun.com/api/dm/querytemplatebyparam.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTemplateByParamWithCallback(request *QueryTemplateByParamRequest, callback func(response *QueryTemplateByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTemplateByParamResponse
		var err error
		defer close(result)
		response, err = client.QueryTemplateByParam(request)
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

// QueryTemplateByParamRequest is the request struct for api QueryTemplateByParam
type QueryTemplateByParamRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	Status               requests.Integer `position:"Query" name:"Status"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

// QueryTemplateByParamResponse is the response struct for api QueryTemplateByParam
type QueryTemplateByParamResponse struct {
	*responses.BaseResponse
	RequestId  string                     `json:"RequestId" xml:"RequestId"`
	TotalCount int                        `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                        `json:"PageSize" xml:"PageSize"`
	Data       DataInQueryTemplateByParam `json:"data" xml:"data"`
}

// CreateQueryTemplateByParamRequest creates a request to invoke QueryTemplateByParam API
func CreateQueryTemplateByParamRequest() (request *QueryTemplateByParamRequest) {
	request = &QueryTemplateByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryTemplateByParam", "", "")
	return
}

// CreateQueryTemplateByParamResponse creates a response to parse from QueryTemplateByParam response
func CreateQueryTemplateByParamResponse() (response *QueryTemplateByParamResponse) {
	response = &QueryTemplateByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
