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

// DescribeAlarms invokes the cms.DescribeAlarms API synchronously
// api document: https://help.aliyun.com/api/cms/describealarms.html
func (client *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
	response = CreateDescribeAlarmsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlarmsWithChan invokes the cms.DescribeAlarms API asynchronously
// api document: https://help.aliyun.com/api/cms/describealarms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmsWithChan(request *DescribeAlarmsRequest) (<-chan *DescribeAlarmsResponse, <-chan error) {
	responseChan := make(chan *DescribeAlarmsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlarms(request)
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

// DescribeAlarmsWithCallback invokes the cms.DescribeAlarms API asynchronously
// api document: https://help.aliyun.com/api/cms/describealarms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmsWithCallback(request *DescribeAlarmsRequest, callback func(response *DescribeAlarmsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlarmsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlarms(request)
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

// DescribeAlarmsRequest is the request struct for api DescribeAlarms
type DescribeAlarmsRequest struct {
	*requests.RpcRequest
	EnableState requests.Boolean `position:"Query" name:"EnableState"`
	Names       string           `position:"Query" name:"Names"`
	DisplayName string           `position:"Query" name:"DisplayName"`
	GroupId     string           `position:"Query" name:"GroupId"`
	Namespace   string           `position:"Query" name:"Namespace"`
	PageSize    string           `position:"Query" name:"PageSize"`
	AlertState  string           `position:"Query" name:"AlertState"`
	NameKeyword string           `position:"Query" name:"NameKeyword"`
	GroupBy     string           `position:"Query" name:"GroupBy"`
	Page        string           `position:"Query" name:"Page"`
	MetricName  string           `position:"Query" name:"MetricName"`
}

// DescribeAlarmsResponse is the response struct for api DescribeAlarms
type DescribeAlarmsResponse struct {
	*responses.BaseResponse
	RequestId  string                     `json:"RequestId" xml:"RequestId"`
	TraceId    string                     `json:"TraceId" xml:"TraceId"`
	Success    bool                       `json:"Success" xml:"Success"`
	Code       int                        `json:"Code" xml:"Code"`
	Message    string                     `json:"Message" xml:"Message"`
	Total      string                     `json:"Total" xml:"Total"`
	Datapoints DatapointsInDescribeAlarms `json:"Datapoints" xml:"Datapoints"`
}

// CreateDescribeAlarmsRequest creates a request to invoke DescribeAlarms API
func CreateDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
	request = &DescribeAlarmsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "DescribeAlarms", "cms", "openAPI")
	return
}

// CreateDescribeAlarmsResponse creates a response to parse from DescribeAlarms response
func CreateDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
	response = &DescribeAlarmsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
