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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QuerySystemEventCount invokes the cms.QuerySystemEventCount API synchronously
// api document: https://help.aliyun.com/api/cms/querysystemeventcount.html
func (client *Client) QuerySystemEventCount(request *QuerySystemEventCountRequest) (response *QuerySystemEventCountResponse, err error) {
	response = CreateQuerySystemEventCountResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySystemEventCountWithChan invokes the cms.QuerySystemEventCount API asynchronously
// api document: https://help.aliyun.com/api/cms/querysystemeventcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySystemEventCountWithChan(request *QuerySystemEventCountRequest) (<-chan *QuerySystemEventCountResponse, <-chan error) {
	responseChan := make(chan *QuerySystemEventCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySystemEventCount(request)
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

// QuerySystemEventCountWithCallback invokes the cms.QuerySystemEventCount API asynchronously
// api document: https://help.aliyun.com/api/cms/querysystemeventcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySystemEventCountWithCallback(request *QuerySystemEventCountRequest, callback func(response *QuerySystemEventCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySystemEventCountResponse
		var err error
		defer close(result)
		response, err = client.QuerySystemEventCount(request)
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

// QuerySystemEventCountRequest is the request struct for api QuerySystemEventCount
type QuerySystemEventCountRequest struct {
	*requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

// QuerySystemEventCountResponse is the response struct for api QuerySystemEventCount
type QuerySystemEventCountResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateQuerySystemEventCountRequest creates a request to invoke QuerySystemEventCount API
func CreateQuerySystemEventCountRequest() (request *QuerySystemEventCountRequest) {
	request = &QuerySystemEventCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "QuerySystemEventCount", "cms", "openAPI")
	return
}

// CreateQuerySystemEventCountResponse creates a response to parse from QuerySystemEventCount response
func CreateQuerySystemEventCountResponse() (response *QuerySystemEventCountResponse) {
	response = &QuerySystemEventCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
