package live

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

// AddCasterProgram invokes the live.AddCasterProgram API synchronously
// api document: https://help.aliyun.com/api/live/addcasterprogram.html
func (client *Client) AddCasterProgram(request *AddCasterProgramRequest) (response *AddCasterProgramResponse, err error) {
	response = CreateAddCasterProgramResponse()
	err = client.DoAction(request, response)
	return
}

// AddCasterProgramWithChan invokes the live.AddCasterProgram API asynchronously
// api document: https://help.aliyun.com/api/live/addcasterprogram.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCasterProgramWithChan(request *AddCasterProgramRequest) (<-chan *AddCasterProgramResponse, <-chan error) {
	responseChan := make(chan *AddCasterProgramResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCasterProgram(request)
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

// AddCasterProgramWithCallback invokes the live.AddCasterProgram API asynchronously
// api document: https://help.aliyun.com/api/live/addcasterprogram.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCasterProgramWithCallback(request *AddCasterProgramRequest, callback func(response *AddCasterProgramResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCasterProgramResponse
		var err error
		defer close(result)
		response, err = client.AddCasterProgram(request)
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

// AddCasterProgramRequest is the request struct for api AddCasterProgram
type AddCasterProgramRequest struct {
	*requests.RpcRequest
	CasterId string                     `position:"Query" name:"CasterId"`
	Episode  *[]AddCasterProgramEpisode `position:"Query" name:"Episode"  type:"Repeated"`
	OwnerId  requests.Integer           `position:"Query" name:"OwnerId"`
}

// AddCasterProgramEpisode is a repeated param struct in AddCasterProgramRequest
type AddCasterProgramEpisode struct {
	ResourceId  string    `name:"ResourceId"`
	ComponentId *[]string `name:"ComponentId" type:"Repeated"`
	SwitchType  string    `name:"SwitchType"`
	EpisodeType string    `name:"EpisodeType"`
	EpisodeName string    `name:"EpisodeName"`
	EndTime     string    `name:"EndTime"`
	StartTime   string    `name:"StartTime"`
}

// AddCasterProgramResponse is the response struct for api AddCasterProgram
type AddCasterProgramResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	EpisodeIds EpisodeIds `json:"EpisodeIds" xml:"EpisodeIds"`
}

// CreateAddCasterProgramRequest creates a request to invoke AddCasterProgram API
func CreateAddCasterProgramRequest() (request *AddCasterProgramRequest) {
	request = &AddCasterProgramRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddCasterProgram", "live", "openAPI")
	return
}

// CreateAddCasterProgramResponse creates a response to parse from AddCasterProgram response
func CreateAddCasterProgramResponse() (response *AddCasterProgramResponse) {
	response = &AddCasterProgramResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
