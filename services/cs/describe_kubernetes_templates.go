package cs

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

// DescribeKubernetesTemplates invokes the cs.DescribeKubernetesTemplates API synchronously
// api document: https://help.aliyun.com/api/cs/describekubernetestemplates.html
func (client *Client) DescribeKubernetesTemplates(request *DescribeKubernetesTemplatesRequest) (response *DescribeKubernetesTemplatesResponse, err error) {
	response = CreateDescribeKubernetesTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeKubernetesTemplatesWithChan invokes the cs.DescribeKubernetesTemplates API asynchronously
// api document: https://help.aliyun.com/api/cs/describekubernetestemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKubernetesTemplatesWithChan(request *DescribeKubernetesTemplatesRequest) (<-chan *DescribeKubernetesTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeKubernetesTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeKubernetesTemplates(request)
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

// DescribeKubernetesTemplatesWithCallback invokes the cs.DescribeKubernetesTemplates API asynchronously
// api document: https://help.aliyun.com/api/cs/describekubernetestemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKubernetesTemplatesWithCallback(request *DescribeKubernetesTemplatesRequest, callback func(response *DescribeKubernetesTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeKubernetesTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeKubernetesTemplates(request)
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

// DescribeKubernetesTemplatesRequest is the request struct for api DescribeKubernetesTemplates
type DescribeKubernetesTemplatesRequest struct {
	*requests.RoaRequest
	KubernetesVersion string `position:"Query" name:"KubernetesVersion"`
	Region            string `position:"Query" name:"Region"`
}

// DescribeKubernetesTemplatesResponse is the response struct for api DescribeKubernetesTemplates
type DescribeKubernetesTemplatesResponse struct {
	*responses.BaseResponse
}

// CreateDescribeKubernetesTemplatesRequest creates a request to invoke DescribeKubernetesTemplates API
func CreateDescribeKubernetesTemplatesRequest() (request *DescribeKubernetesTemplatesRequest) {
	request = &DescribeKubernetesTemplatesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeKubernetesTemplates", "/k8s/templates", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeKubernetesTemplatesResponse creates a response to parse from DescribeKubernetesTemplates response
func CreateDescribeKubernetesTemplatesResponse() (response *DescribeKubernetesTemplatesResponse) {
	response = &DescribeKubernetesTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
