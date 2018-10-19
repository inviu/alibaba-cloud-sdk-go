package bssopenapi

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

// GetPayAsYouGoPrice invokes the bssopenapi.GetPayAsYouGoPrice API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/getpayasyougoprice.html
func (client *Client) GetPayAsYouGoPrice(request *GetPayAsYouGoPriceRequest) (response *GetPayAsYouGoPriceResponse, err error) {
	response = CreateGetPayAsYouGoPriceResponse()
	err = client.DoAction(request, response)
	return
}

// GetPayAsYouGoPriceWithChan invokes the bssopenapi.GetPayAsYouGoPrice API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/getpayasyougoprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPayAsYouGoPriceWithChan(request *GetPayAsYouGoPriceRequest) (<-chan *GetPayAsYouGoPriceResponse, <-chan error) {
	responseChan := make(chan *GetPayAsYouGoPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPayAsYouGoPrice(request)
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

// GetPayAsYouGoPriceWithCallback invokes the bssopenapi.GetPayAsYouGoPrice API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/getpayasyougoprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPayAsYouGoPriceWithCallback(request *GetPayAsYouGoPriceRequest, callback func(response *GetPayAsYouGoPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPayAsYouGoPriceResponse
		var err error
		defer close(result)
		response, err = client.GetPayAsYouGoPrice(request)
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

// GetPayAsYouGoPriceRequest is the request struct for api GetPayAsYouGoPrice
type GetPayAsYouGoPriceRequest struct {
	*requests.RpcRequest
	ProductCode      string                          `position:"Query" name:"ProductCode"`
	SubscriptionType string                          `position:"Query" name:"SubscriptionType"`
	ModuleList       *[]GetPayAsYouGoPriceModuleList `position:"Query" name:"ModuleList"  type:"Repeated"`
	OwnerId          requests.Integer                `position:"Query" name:"OwnerId"`
	Region           string                          `position:"Query" name:"Region"`
	ProductType      string                          `position:"Query" name:"ProductType"`
}

// GetPayAsYouGoPriceModuleList is a repeated param struct in GetPayAsYouGoPriceRequest
type GetPayAsYouGoPriceModuleList struct {
	ModuleCode string `name:"ModuleCode"`
	PriceType  string `name:"PriceType"`
	Config     string `name:"Config"`
}

// GetPayAsYouGoPriceResponse is the response struct for api GetPayAsYouGoPrice
type GetPayAsYouGoPriceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPayAsYouGoPriceRequest creates a request to invoke GetPayAsYouGoPrice API
func CreateGetPayAsYouGoPriceRequest() (request *GetPayAsYouGoPriceRequest) {
	request = &GetPayAsYouGoPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "GetPayAsYouGoPrice", "", "")
	return
}

// CreateGetPayAsYouGoPriceResponse creates a response to parse from GetPayAsYouGoPrice response
func CreateGetPayAsYouGoPriceResponse() (response *GetPayAsYouGoPriceResponse) {
	response = &GetPayAsYouGoPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
