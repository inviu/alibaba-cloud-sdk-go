package aegis

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

// GetAccountStatistics invokes the aegis.GetAccountStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/getaccountstatistics.html
func (client *Client) GetAccountStatistics(request *GetAccountStatisticsRequest) (response *GetAccountStatisticsResponse, err error) {
	response = CreateGetAccountStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccountStatisticsWithChan invokes the aegis.GetAccountStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/getaccountstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccountStatisticsWithChan(request *GetAccountStatisticsRequest) (<-chan *GetAccountStatisticsResponse, <-chan error) {
	responseChan := make(chan *GetAccountStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountStatistics(request)
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

// GetAccountStatisticsWithCallback invokes the aegis.GetAccountStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/getaccountstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccountStatisticsWithCallback(request *GetAccountStatisticsRequest, callback func(response *GetAccountStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccountStatisticsResponse
		var err error
		defer close(result)
		response, err = client.GetAccountStatistics(request)
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

// GetAccountStatisticsRequest is the request struct for api GetAccountStatistics
type GetAccountStatisticsRequest struct {
	*requests.RpcRequest
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

// GetAccountStatisticsResponse is the response struct for api GetAccountStatistics
type GetAccountStatisticsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAccountStatisticsRequest creates a request to invoke GetAccountStatistics API
func CreateGetAccountStatisticsRequest() (request *GetAccountStatisticsRequest) {
	request = &GetAccountStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "GetAccountStatistics", "vipaegis", "openAPI")
	return
}

// CreateGetAccountStatisticsResponse creates a response to parse from GetAccountStatistics response
func CreateGetAccountStatisticsResponse() (response *GetAccountStatisticsResponse) {
	response = &GetAccountStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
