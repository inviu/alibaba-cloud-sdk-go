package drds

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

// ModifyRdsReadWeight invokes the drds.ModifyRdsReadWeight API synchronously
// api document: https://help.aliyun.com/api/drds/modifyrdsreadweight.html
func (client *Client) ModifyRdsReadWeight(request *ModifyRdsReadWeightRequest) (response *ModifyRdsReadWeightResponse, err error) {
	response = CreateModifyRdsReadWeightResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRdsReadWeightWithChan invokes the drds.ModifyRdsReadWeight API asynchronously
// api document: https://help.aliyun.com/api/drds/modifyrdsreadweight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRdsReadWeightWithChan(request *ModifyRdsReadWeightRequest) (<-chan *ModifyRdsReadWeightResponse, <-chan error) {
	responseChan := make(chan *ModifyRdsReadWeightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRdsReadWeight(request)
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

// ModifyRdsReadWeightWithCallback invokes the drds.ModifyRdsReadWeight API asynchronously
// api document: https://help.aliyun.com/api/drds/modifyrdsreadweight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRdsReadWeightWithCallback(request *ModifyRdsReadWeightRequest, callback func(response *ModifyRdsReadWeightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRdsReadWeightResponse
		var err error
		defer close(result)
		response, err = client.ModifyRdsReadWeight(request)
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

// ModifyRdsReadWeightRequest is the request struct for api ModifyRdsReadWeight
type ModifyRdsReadWeightRequest struct {
	*requests.RpcRequest
	InstanceNames  string `position:"Query" name:"InstanceNames"`
	DbName         string `position:"Query" name:"DbName"`
	Weights        string `position:"Query" name:"Weights"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// ModifyRdsReadWeightResponse is the response struct for api ModifyRdsReadWeight
type ModifyRdsReadWeightResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyRdsReadWeightRequest creates a request to invoke ModifyRdsReadWeight API
func CreateModifyRdsReadWeightRequest() (request *ModifyRdsReadWeightRequest) {
	request = &ModifyRdsReadWeightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "ModifyRdsReadWeight", "", "")
	return
}

// CreateModifyRdsReadWeightResponse creates a response to parse from ModifyRdsReadWeight response
func CreateModifyRdsReadWeightResponse() (response *ModifyRdsReadWeightResponse) {
	response = &ModifyRdsReadWeightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
