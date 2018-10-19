package domain

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

// QueryTransferInList invokes the domain.QueryTransferInList API synchronously
// api document: https://help.aliyun.com/api/domain/querytransferinlist.html
func (client *Client) QueryTransferInList(request *QueryTransferInListRequest) (response *QueryTransferInListResponse, err error) {
	response = CreateQueryTransferInListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTransferInListWithChan invokes the domain.QueryTransferInList API asynchronously
// api document: https://help.aliyun.com/api/domain/querytransferinlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTransferInListWithChan(request *QueryTransferInListRequest) (<-chan *QueryTransferInListResponse, <-chan error) {
	responseChan := make(chan *QueryTransferInListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTransferInList(request)
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

// QueryTransferInListWithCallback invokes the domain.QueryTransferInList API asynchronously
// api document: https://help.aliyun.com/api/domain/querytransferinlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTransferInListWithCallback(request *QueryTransferInListRequest, callback func(response *QueryTransferInListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTransferInListResponse
		var err error
		defer close(result)
		response, err = client.QueryTransferInList(request)
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

// QueryTransferInListRequest is the request struct for api QueryTransferInList
type QueryTransferInListRequest struct {
	*requests.RpcRequest
	SubmissionStartDate    requests.Integer `position:"Query" name:"SubmissionStartDate"`
	UserClientIp           string           `position:"Query" name:"UserClientIp"`
	SubmissionEndDate      requests.Integer `position:"Query" name:"SubmissionEndDate"`
	DomainName             string           `position:"Query" name:"DomainName"`
	SimpleTransferInStatus string           `position:"Query" name:"SimpleTransferInStatus"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	Lang                   string           `position:"Query" name:"Lang"`
	PageNum                requests.Integer `position:"Query" name:"PageNum"`
}

// QueryTransferInListResponse is the response struct for api QueryTransferInList
type QueryTransferInListResponse struct {
	*responses.BaseResponse
	RequestId      string                    `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                       `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                       `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   int                       `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       int                       `json:"PageSize" xml:"PageSize"`
	PrePage        bool                      `json:"PrePage" xml:"PrePage"`
	NextPage       bool                      `json:"NextPage" xml:"NextPage"`
	Data           DataInQueryTransferInList `json:"Data" xml:"Data"`
}

// CreateQueryTransferInListRequest creates a request to invoke QueryTransferInList API
func CreateQueryTransferInListRequest() (request *QueryTransferInListRequest) {
	request = &QueryTransferInListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryTransferInList", "", "")
	return
}

// CreateQueryTransferInListResponse creates a response to parse from QueryTransferInList response
func CreateQueryTransferInListResponse() (response *QueryTransferInListResponse) {
	response = &QueryTransferInListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
