package linkface

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

// QueryLicenses invokes the linkface.QueryLicenses API synchronously
// api document: https://help.aliyun.com/api/linkface/querylicenses.html
func (client *Client) QueryLicenses(request *QueryLicensesRequest) (response *QueryLicensesResponse, err error) {
	response = CreateQueryLicensesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryLicensesWithChan invokes the linkface.QueryLicenses API asynchronously
// api document: https://help.aliyun.com/api/linkface/querylicenses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryLicensesWithChan(request *QueryLicensesRequest) (<-chan *QueryLicensesResponse, <-chan error) {
	responseChan := make(chan *QueryLicensesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryLicenses(request)
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

// QueryLicensesWithCallback invokes the linkface.QueryLicenses API asynchronously
// api document: https://help.aliyun.com/api/linkface/querylicenses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryLicensesWithCallback(request *QueryLicensesRequest, callback func(response *QueryLicensesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryLicensesResponse
		var err error
		defer close(result)
		response, err = client.QueryLicenses(request)
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

// QueryLicensesRequest is the request struct for api QueryLicenses
type QueryLicensesRequest struct {
	*requests.RpcRequest
	LicenseType requests.Integer `position:"Body" name:"LicenseType"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	CurrentPage requests.Integer `position:"Body" name:"CurrentPage"`
}

// QueryLicensesResponse is the response struct for api QueryLicenses
type QueryLicensesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	PageCount int        `json:"PageCount" xml:"PageCount"`
	PageSize  int        `json:"PageSize" xml:"PageSize"`
	Page      int        `json:"Page" xml:"Page"`
	Total     int        `json:"Total" xml:"Total"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateQueryLicensesRequest creates a request to invoke QueryLicenses API
func CreateQueryLicensesRequest() (request *QueryLicensesRequest) {
	request = &QueryLicensesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkFace", "2018-07-20", "QueryLicenses", "", "")
	return
}

// CreateQueryLicensesResponse creates a response to parse from QueryLicenses response
func CreateQueryLicensesResponse() (response *QueryLicensesResponse) {
	response = &QueryLicensesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
