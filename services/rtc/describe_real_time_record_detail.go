package rtc

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

// DescribeRealTimeRecordDetail invokes the rtc.DescribeRealTimeRecordDetail API synchronously
// api document: https://help.aliyun.com/api/rtc/describerealtimerecorddetail.html
func (client *Client) DescribeRealTimeRecordDetail(request *DescribeRealTimeRecordDetailRequest) (response *DescribeRealTimeRecordDetailResponse, err error) {
	response = CreateDescribeRealTimeRecordDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRealTimeRecordDetailWithChan invokes the rtc.DescribeRealTimeRecordDetail API asynchronously
// api document: https://help.aliyun.com/api/rtc/describerealtimerecorddetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRealTimeRecordDetailWithChan(request *DescribeRealTimeRecordDetailRequest) (<-chan *DescribeRealTimeRecordDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeRealTimeRecordDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRealTimeRecordDetail(request)
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

// DescribeRealTimeRecordDetailWithCallback invokes the rtc.DescribeRealTimeRecordDetail API asynchronously
// api document: https://help.aliyun.com/api/rtc/describerealtimerecorddetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRealTimeRecordDetailWithCallback(request *DescribeRealTimeRecordDetailRequest, callback func(response *DescribeRealTimeRecordDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRealTimeRecordDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeRealTimeRecordDetail(request)
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

// DescribeRealTimeRecordDetailRequest is the request struct for api DescribeRealTimeRecordDetail
type DescribeRealTimeRecordDetailRequest struct {
	*requests.RpcRequest
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	RecordId  string           `position:"Query" name:"RecordId"`
	AppId     string           `position:"Query" name:"AppId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
}

// DescribeRealTimeRecordDetailResponse is the response struct for api DescribeRealTimeRecordDetail
type DescribeRealTimeRecordDetailResponse struct {
	*responses.BaseResponse
	RequestId         string                                          `json:"RequestId" xml:"RequestId"`
	RecordDetailInfos RecordDetailInfosInDescribeRealTimeRecordDetail `json:"RecordDetailInfos" xml:"RecordDetailInfos"`
}

// CreateDescribeRealTimeRecordDetailRequest creates a request to invoke DescribeRealTimeRecordDetail API
func CreateDescribeRealTimeRecordDetailRequest() (request *DescribeRealTimeRecordDetailRequest) {
	request = &DescribeRealTimeRecordDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRealTimeRecordDetail", "rtc", "openAPI")
	return
}

// CreateDescribeRealTimeRecordDetailResponse creates a response to parse from DescribeRealTimeRecordDetail response
func CreateDescribeRealTimeRecordDetailResponse() (response *DescribeRealTimeRecordDetailResponse) {
	response = &DescribeRealTimeRecordDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
