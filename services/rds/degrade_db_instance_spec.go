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

// DegradeDBInstanceSpec invokes the rds.DegradeDBInstanceSpec API synchronously
// api document: https://help.aliyun.com/api/rds/degradedbinstancespec.html
func (client *Client) DegradeDBInstanceSpec(request *DegradeDBInstanceSpecRequest) (response *DegradeDBInstanceSpecResponse, err error) {
	response = CreateDegradeDBInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// DegradeDBInstanceSpecWithChan invokes the rds.DegradeDBInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/rds/degradedbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DegradeDBInstanceSpecWithChan(request *DegradeDBInstanceSpecRequest) (<-chan *DegradeDBInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *DegradeDBInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DegradeDBInstanceSpec(request)
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

// DegradeDBInstanceSpecWithCallback invokes the rds.DegradeDBInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/rds/degradedbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DegradeDBInstanceSpecWithCallback(request *DegradeDBInstanceSpecRequest, callback func(response *DegradeDBInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DegradeDBInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.DegradeDBInstanceSpec(request)
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

// DegradeDBInstanceSpecRequest is the request struct for api DegradeDBInstanceSpec
type DegradeDBInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    requests.Integer `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
}

// DegradeDBInstanceSpecResponse is the response struct for api DegradeDBInstanceSpec
type DegradeDBInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDegradeDBInstanceSpecRequest creates a request to invoke DegradeDBInstanceSpec API
func CreateDegradeDBInstanceSpecRequest() (request *DegradeDBInstanceSpecRequest) {
	request = &DegradeDBInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DegradeDBInstanceSpec", "rds", "openAPI")
	return
}

// CreateDegradeDBInstanceSpecResponse creates a response to parse from DegradeDBInstanceSpec response
func CreateDegradeDBInstanceSpecResponse() (response *DegradeDBInstanceSpecResponse) {
	response = &DegradeDBInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
