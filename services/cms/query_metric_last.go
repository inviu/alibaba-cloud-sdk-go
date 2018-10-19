package cms

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

// QueryMetricLast invokes the cms.QueryMetricLast API synchronously
// api document: https://help.aliyun.com/api/cms/querymetriclast.html
func (client *Client) QueryMetricLast(request *QueryMetricLastRequest) (response *QueryMetricLastResponse, err error) {
	response = CreateQueryMetricLastResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMetricLastWithChan invokes the cms.QueryMetricLast API asynchronously
// api document: https://help.aliyun.com/api/cms/querymetriclast.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMetricLastWithChan(request *QueryMetricLastRequest) (<-chan *QueryMetricLastResponse, <-chan error) {
	responseChan := make(chan *QueryMetricLastResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMetricLast(request)
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

// QueryMetricLastWithCallback invokes the cms.QueryMetricLast API asynchronously
// api document: https://help.aliyun.com/api/cms/querymetriclast.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMetricLastWithCallback(request *QueryMetricLastRequest, callback func(response *QueryMetricLastResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMetricLastResponse
		var err error
		defer close(result)
		response, err = client.QueryMetricLast(request)
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

// QueryMetricLastRequest is the request struct for api QueryMetricLast
type QueryMetricLastRequest struct {
	*requests.RpcRequest
	Cursor          string           `position:"Query" name:"Cursor"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Period          string           `position:"Query" name:"Period"`
	Metric          string           `position:"Query" name:"Metric"`
	Length          string           `position:"Query" name:"Length"`
	Project         string           `position:"Query" name:"Project"`
	EndTime         string           `position:"Query" name:"EndTime"`
	Express         string           `position:"Query" name:"Express"`
	StartTime       string           `position:"Query" name:"StartTime"`
	Page            string           `position:"Query" name:"Page"`
	Dimensions      string           `position:"Query" name:"Dimensions"`
}

// QueryMetricLastResponse is the response struct for api QueryMetricLast
type QueryMetricLastResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Cursor     string `json:"Cursor" xml:"Cursor"`
	Datapoints string `json:"Datapoints" xml:"Datapoints"`
	Period     string `json:"Period" xml:"Period"`
}

// CreateQueryMetricLastRequest creates a request to invoke QueryMetricLast API
func CreateQueryMetricLastRequest() (request *QueryMetricLastRequest) {
	request = &QueryMetricLastRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "QueryMetricLast", "cms", "openAPI")
	return
}

// CreateQueryMetricLastResponse creates a response to parse from QueryMetricLast response
func CreateQueryMetricLastResponse() (response *QueryMetricLastResponse) {
	response = &QueryMetricLastResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
