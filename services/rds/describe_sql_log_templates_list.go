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

// DescribeSqlLogTemplatesList invokes the rds.DescribeSqlLogTemplatesList API synchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtemplateslist.html
func (client *Client) DescribeSqlLogTemplatesList(request *DescribeSqlLogTemplatesListRequest) (response *DescribeSqlLogTemplatesListResponse, err error) {
	response = CreateDescribeSqlLogTemplatesListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSqlLogTemplatesListWithChan invokes the rds.DescribeSqlLogTemplatesList API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtemplateslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSqlLogTemplatesListWithChan(request *DescribeSqlLogTemplatesListRequest) (<-chan *DescribeSqlLogTemplatesListResponse, <-chan error) {
	responseChan := make(chan *DescribeSqlLogTemplatesListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSqlLogTemplatesList(request)
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

// DescribeSqlLogTemplatesListWithCallback invokes the rds.DescribeSqlLogTemplatesList API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogtemplateslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSqlLogTemplatesListWithCallback(request *DescribeSqlLogTemplatesListRequest, callback func(response *DescribeSqlLogTemplatesListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSqlLogTemplatesListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSqlLogTemplatesList(request)
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

// DescribeSqlLogTemplatesListRequest is the request struct for api DescribeSqlLogTemplatesList
type DescribeSqlLogTemplatesListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	MinAvgConsume        requests.Integer `position:"Query" name:"MinAvgConsume"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaxRecordsPerPage    requests.Integer `position:"Query" name:"MaxRecordsPerPage"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaxAvgConsume        requests.Integer `position:"Query" name:"MaxAvgConsume"`
	SortKey              string           `position:"Query" name:"SortKey"`
	MinAvgScanRows       requests.Integer `position:"Query" name:"MinAvgScanRows"`
	SqType               string           `position:"Query" name:"SqType"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SortMethod           string           `position:"Query" name:"SortMethod"`
	PageNumbers          requests.Integer `position:"Query" name:"PageNumbers"`
	PagingId             string           `position:"Query" name:"PagingId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	MaxAvgScanRows       requests.Integer `position:"Query" name:"MaxAvgScanRows"`
}

// DescribeSqlLogTemplatesListResponse is the response struct for api DescribeSqlLogTemplatesList
type DescribeSqlLogTemplatesListResponse struct {
	*responses.BaseResponse
	RequestId         string                             `json:"RequestId" xml:"RequestId"`
	DBInstanceID      int                                `json:"DBInstanceID" xml:"DBInstanceID"`
	DBInstanceName    string                             `json:"DBInstanceName" xml:"DBInstanceName"`
	StartTime         string                             `json:"StartTime" xml:"StartTime"`
	EndTime           string                             `json:"EndTime" xml:"EndTime"`
	TotalRecords      int                                `json:"TotalRecords" xml:"TotalRecords"`
	PagingID          string                             `json:"PagingID" xml:"PagingID"`
	MaxRecordsPerPage int                                `json:"MaxRecordsPerPage" xml:"MaxRecordsPerPage"`
	PageNumbers       int                                `json:"PageNumbers" xml:"PageNumbers"`
	ItemsNumbers      int                                `json:"ItemsNumbers" xml:"ItemsNumbers"`
	Items             ItemsInDescribeSqlLogTemplatesList `json:"Items" xml:"Items"`
}

// CreateDescribeSqlLogTemplatesListRequest creates a request to invoke DescribeSqlLogTemplatesList API
func CreateDescribeSqlLogTemplatesListRequest() (request *DescribeSqlLogTemplatesListRequest) {
	request = &DescribeSqlLogTemplatesListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSqlLogTemplatesList", "rds", "openAPI")
	return
}

// CreateDescribeSqlLogTemplatesListResponse creates a response to parse from DescribeSqlLogTemplatesList response
func CreateDescribeSqlLogTemplatesListResponse() (response *DescribeSqlLogTemplatesListResponse) {
	response = &DescribeSqlLogTemplatesListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
