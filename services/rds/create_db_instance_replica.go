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

// CreateDBInstanceReplica invokes the rds.CreateDBInstanceReplica API synchronously
// api document: https://help.aliyun.com/api/rds/createdbinstancereplica.html
func (client *Client) CreateDBInstanceReplica(request *CreateDBInstanceReplicaRequest) (response *CreateDBInstanceReplicaResponse, err error) {
	response = CreateCreateDBInstanceReplicaResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBInstanceReplicaWithChan invokes the rds.CreateDBInstanceReplica API asynchronously
// api document: https://help.aliyun.com/api/rds/createdbinstancereplica.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceReplicaWithChan(request *CreateDBInstanceReplicaRequest) (<-chan *CreateDBInstanceReplicaResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceReplicaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstanceReplica(request)
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

// CreateDBInstanceReplicaWithCallback invokes the rds.CreateDBInstanceReplica API asynchronously
// api document: https://help.aliyun.com/api/rds/createdbinstancereplica.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceReplicaWithCallback(request *CreateDBInstanceReplicaRequest, callback func(response *CreateDBInstanceReplicaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceReplicaResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstanceReplica(request)
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

// CreateDBInstanceReplicaRequest is the request struct for api CreateDBInstanceReplica
type CreateDBInstanceReplicaRequest struct {
	*requests.RpcRequest
	ConnectionMode        string           `position:"Query" name:"ConnectionMode"`
	DomainMode            string           `position:"Query" name:"DomainMode"`
	ReplicaDescription    string           `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset       string           `position:"Query" name:"SystemDBCharset"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	Engine                string           `position:"Query" name:"Engine"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	DBInstanceNetType     string           `position:"Query" name:"DBInstanceNetType"`
	Period                string           `position:"Query" name:"Period"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime              string           `position:"Query" name:"UsedTime"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string           `position:"Query" name:"SecurityIPList"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	SourceDBInstanceId    string           `position:"Query" name:"SourceDBInstanceId"`
	ReplicaMode           string           `position:"Query" name:"ReplicaMode"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	PayType               string           `position:"Query" name:"PayType"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

// CreateDBInstanceReplicaResponse is the response struct for api CreateDBInstanceReplica
type CreateDBInstanceReplicaResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId      int    `json:"OrderId" xml:"OrderId"`
	ReplicaId    string `json:"ReplicaId" xml:"ReplicaId"`
	WorkflowId   string `json:"WorkflowId" xml:"WorkflowId"`
}

// CreateCreateDBInstanceReplicaRequest creates a request to invoke CreateDBInstanceReplica API
func CreateCreateDBInstanceReplicaRequest() (request *CreateDBInstanceReplicaRequest) {
	request = &CreateDBInstanceReplicaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstanceReplica", "rds", "openAPI")
	return
}

// CreateCreateDBInstanceReplicaResponse creates a response to parse from CreateDBInstanceReplica response
func CreateCreateDBInstanceReplicaResponse() (response *CreateDBInstanceReplicaResponse) {
	response = &CreateDBInstanceReplicaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
