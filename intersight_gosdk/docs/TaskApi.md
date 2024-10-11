# \TaskApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskCatalystSdwanScopedInventory**](TaskApi.md#CreateTaskCatalystSdwanScopedInventory) | **Post** /api/v1/task/CatalystSdwanScopedInventories | Create a &#39;task.CatalystSdwanScopedInventory&#39; resource.
[**CreateTaskFabricMosScopedInventory**](TaskApi.md#CreateTaskFabricMosScopedInventory) | **Post** /api/v1/task/FabricMosScopedInventories | Create a &#39;task.FabricMosScopedInventory&#39; resource.
[**CreateTaskHitachiScopedInventory**](TaskApi.md#CreateTaskHitachiScopedInventory) | **Post** /api/v1/task/HitachiScopedInventories | Create a &#39;task.HitachiScopedInventory&#39; resource.
[**CreateTaskHyperFlexManagementScopedInventory**](TaskApi.md#CreateTaskHyperFlexManagementScopedInventory) | **Post** /api/v1/task/HyperFlexManagementScopedInventories | Create a &#39;task.HyperFlexManagementScopedInventory&#39; resource.
[**CreateTaskHyperflexScopedInventory**](TaskApi.md#CreateTaskHyperflexScopedInventory) | **Post** /api/v1/task/HyperflexScopedInventories | Create a &#39;task.HyperflexScopedInventory&#39; resource.
[**CreateTaskMdsScopedInventory**](TaskApi.md#CreateTaskMdsScopedInventory) | **Post** /api/v1/task/MdsScopedInventories | Create a &#39;task.MdsScopedInventory&#39; resource.
[**CreateTaskMdsSystemScopedInventory**](TaskApi.md#CreateTaskMdsSystemScopedInventory) | **Post** /api/v1/task/MdsSystemScopedInventories | Create a &#39;task.MdsSystemScopedInventory&#39; resource.
[**CreateTaskMerakiScopedInventory**](TaskApi.md#CreateTaskMerakiScopedInventory) | **Post** /api/v1/task/MerakiScopedInventories | Create a &#39;task.MerakiScopedInventory&#39; resource.
[**CreateTaskNetAppScopedInventory**](TaskApi.md#CreateTaskNetAppScopedInventory) | **Post** /api/v1/task/NetAppScopedInventories | Create a &#39;task.NetAppScopedInventory&#39; resource.
[**CreateTaskNexusScopedInventory**](TaskApi.md#CreateTaskNexusScopedInventory) | **Post** /api/v1/task/NexusScopedInventories | Create a &#39;task.NexusScopedInventory&#39; resource.
[**CreateTaskNexusSystemScopedInventory**](TaskApi.md#CreateTaskNexusSystemScopedInventory) | **Post** /api/v1/task/NexusSystemScopedInventories | Create a &#39;task.NexusSystemScopedInventory&#39; resource.
[**CreateTaskNexusVlanScopedInventory**](TaskApi.md#CreateTaskNexusVlanScopedInventory) | **Post** /api/v1/task/NexusVlanScopedInventories | Create a &#39;task.NexusVlanScopedInventory&#39; resource.
[**CreateTaskPublicCloudScopedInventory**](TaskApi.md#CreateTaskPublicCloudScopedInventory) | **Post** /api/v1/task/PublicCloudScopedInventories | Create a &#39;task.PublicCloudScopedInventory&#39; resource.
[**CreateTaskPureScopedInventory**](TaskApi.md#CreateTaskPureScopedInventory) | **Post** /api/v1/task/PureScopedInventories | Create a &#39;task.PureScopedInventory&#39; resource.
[**CreateTaskServerScopedInventory**](TaskApi.md#CreateTaskServerScopedInventory) | **Post** /api/v1/task/ServerScopedInventories | Create a &#39;task.ServerScopedInventory&#39; resource.
[**CreateTaskWorkflowAction**](TaskApi.md#CreateTaskWorkflowAction) | **Post** /api/v1/task/WorkflowActions | Create a &#39;task.WorkflowAction&#39; resource.
[**GetTaskWorkflowActionByMoid**](TaskApi.md#GetTaskWorkflowActionByMoid) | **Get** /api/v1/task/WorkflowActions/{Moid} | Read a &#39;task.WorkflowAction&#39; resource.
[**GetTaskWorkflowActionList**](TaskApi.md#GetTaskWorkflowActionList) | **Get** /api/v1/task/WorkflowActions | Read a &#39;task.WorkflowAction&#39; resource.
[**PatchTaskWorkflowAction**](TaskApi.md#PatchTaskWorkflowAction) | **Patch** /api/v1/task/WorkflowActions/{Moid} | Update a &#39;task.WorkflowAction&#39; resource.
[**UpdateTaskWorkflowAction**](TaskApi.md#UpdateTaskWorkflowAction) | **Post** /api/v1/task/WorkflowActions/{Moid} | Update a &#39;task.WorkflowAction&#39; resource.



## CreateTaskCatalystSdwanScopedInventory

> TaskCatalystSdwanScopedInventory CreateTaskCatalystSdwanScopedInventory(ctx).TaskCatalystSdwanScopedInventory(taskCatalystSdwanScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.CatalystSdwanScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskCatalystSdwanScopedInventory := *openapiclient.NewTaskCatalystSdwanScopedInventory("ClassId_example", "ObjectType_example") // TaskCatalystSdwanScopedInventory | The 'task.CatalystSdwanScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskCatalystSdwanScopedInventory(context.Background()).TaskCatalystSdwanScopedInventory(taskCatalystSdwanScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskCatalystSdwanScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskCatalystSdwanScopedInventory`: TaskCatalystSdwanScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskCatalystSdwanScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskCatalystSdwanScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskCatalystSdwanScopedInventory** | [**TaskCatalystSdwanScopedInventory**](TaskCatalystSdwanScopedInventory.md) | The &#39;task.CatalystSdwanScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskCatalystSdwanScopedInventory**](TaskCatalystSdwanScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskFabricMosScopedInventory

> TaskFabricMosScopedInventory CreateTaskFabricMosScopedInventory(ctx).TaskFabricMosScopedInventory(taskFabricMosScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.FabricMosScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskFabricMosScopedInventory := *openapiclient.NewTaskFabricMosScopedInventory("ClassId_example", "ObjectType_example") // TaskFabricMosScopedInventory | The 'task.FabricMosScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskFabricMosScopedInventory(context.Background()).TaskFabricMosScopedInventory(taskFabricMosScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskFabricMosScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskFabricMosScopedInventory`: TaskFabricMosScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskFabricMosScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskFabricMosScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskFabricMosScopedInventory** | [**TaskFabricMosScopedInventory**](TaskFabricMosScopedInventory.md) | The &#39;task.FabricMosScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskFabricMosScopedInventory**](TaskFabricMosScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskHitachiScopedInventory

> TaskHitachiScopedInventory CreateTaskHitachiScopedInventory(ctx).TaskHitachiScopedInventory(taskHitachiScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.HitachiScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskHitachiScopedInventory := *openapiclient.NewTaskHitachiScopedInventory("ClassId_example", "ObjectType_example") // TaskHitachiScopedInventory | The 'task.HitachiScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskHitachiScopedInventory(context.Background()).TaskHitachiScopedInventory(taskHitachiScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskHitachiScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskHitachiScopedInventory`: TaskHitachiScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskHitachiScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskHitachiScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskHitachiScopedInventory** | [**TaskHitachiScopedInventory**](TaskHitachiScopedInventory.md) | The &#39;task.HitachiScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskHitachiScopedInventory**](TaskHitachiScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskHyperFlexManagementScopedInventory

> TaskHyperFlexManagementScopedInventory CreateTaskHyperFlexManagementScopedInventory(ctx).TaskHyperFlexManagementScopedInventory(taskHyperFlexManagementScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.HyperFlexManagementScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskHyperFlexManagementScopedInventory := *openapiclient.NewTaskHyperFlexManagementScopedInventory("ClassId_example", "ObjectType_example") // TaskHyperFlexManagementScopedInventory | The 'task.HyperFlexManagementScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskHyperFlexManagementScopedInventory(context.Background()).TaskHyperFlexManagementScopedInventory(taskHyperFlexManagementScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskHyperFlexManagementScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskHyperFlexManagementScopedInventory`: TaskHyperFlexManagementScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskHyperFlexManagementScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskHyperFlexManagementScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskHyperFlexManagementScopedInventory** | [**TaskHyperFlexManagementScopedInventory**](TaskHyperFlexManagementScopedInventory.md) | The &#39;task.HyperFlexManagementScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskHyperFlexManagementScopedInventory**](TaskHyperFlexManagementScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskHyperflexScopedInventory

> TaskHyperflexScopedInventory CreateTaskHyperflexScopedInventory(ctx).TaskHyperflexScopedInventory(taskHyperflexScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.HyperflexScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskHyperflexScopedInventory := *openapiclient.NewTaskHyperflexScopedInventory("ClassId_example", "ObjectType_example") // TaskHyperflexScopedInventory | The 'task.HyperflexScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskHyperflexScopedInventory(context.Background()).TaskHyperflexScopedInventory(taskHyperflexScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskHyperflexScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskHyperflexScopedInventory`: TaskHyperflexScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskHyperflexScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskHyperflexScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskHyperflexScopedInventory** | [**TaskHyperflexScopedInventory**](TaskHyperflexScopedInventory.md) | The &#39;task.HyperflexScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskHyperflexScopedInventory**](TaskHyperflexScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskMdsScopedInventory

> TaskMdsScopedInventory CreateTaskMdsScopedInventory(ctx).TaskMdsScopedInventory(taskMdsScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.MdsScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskMdsScopedInventory := *openapiclient.NewTaskMdsScopedInventory("ClassId_example", "ObjectType_example") // TaskMdsScopedInventory | The 'task.MdsScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskMdsScopedInventory(context.Background()).TaskMdsScopedInventory(taskMdsScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskMdsScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskMdsScopedInventory`: TaskMdsScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskMdsScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskMdsScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskMdsScopedInventory** | [**TaskMdsScopedInventory**](TaskMdsScopedInventory.md) | The &#39;task.MdsScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskMdsScopedInventory**](TaskMdsScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskMdsSystemScopedInventory

> TaskMdsSystemScopedInventory CreateTaskMdsSystemScopedInventory(ctx).TaskMdsSystemScopedInventory(taskMdsSystemScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.MdsSystemScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskMdsSystemScopedInventory := *openapiclient.NewTaskMdsSystemScopedInventory("ClassId_example", "ObjectType_example") // TaskMdsSystemScopedInventory | The 'task.MdsSystemScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskMdsSystemScopedInventory(context.Background()).TaskMdsSystemScopedInventory(taskMdsSystemScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskMdsSystemScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskMdsSystemScopedInventory`: TaskMdsSystemScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskMdsSystemScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskMdsSystemScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskMdsSystemScopedInventory** | [**TaskMdsSystemScopedInventory**](TaskMdsSystemScopedInventory.md) | The &#39;task.MdsSystemScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskMdsSystemScopedInventory**](TaskMdsSystemScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskMerakiScopedInventory

> TaskMerakiScopedInventory CreateTaskMerakiScopedInventory(ctx).TaskMerakiScopedInventory(taskMerakiScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.MerakiScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskMerakiScopedInventory := *openapiclient.NewTaskMerakiScopedInventory("ClassId_example", "ObjectType_example") // TaskMerakiScopedInventory | The 'task.MerakiScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskMerakiScopedInventory(context.Background()).TaskMerakiScopedInventory(taskMerakiScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskMerakiScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskMerakiScopedInventory`: TaskMerakiScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskMerakiScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskMerakiScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskMerakiScopedInventory** | [**TaskMerakiScopedInventory**](TaskMerakiScopedInventory.md) | The &#39;task.MerakiScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskMerakiScopedInventory**](TaskMerakiScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskNetAppScopedInventory

> TaskNetAppScopedInventory CreateTaskNetAppScopedInventory(ctx).TaskNetAppScopedInventory(taskNetAppScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.NetAppScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskNetAppScopedInventory := *openapiclient.NewTaskNetAppScopedInventory("ClassId_example", "ObjectType_example") // TaskNetAppScopedInventory | The 'task.NetAppScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskNetAppScopedInventory(context.Background()).TaskNetAppScopedInventory(taskNetAppScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskNetAppScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskNetAppScopedInventory`: TaskNetAppScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskNetAppScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskNetAppScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskNetAppScopedInventory** | [**TaskNetAppScopedInventory**](TaskNetAppScopedInventory.md) | The &#39;task.NetAppScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskNetAppScopedInventory**](TaskNetAppScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskNexusScopedInventory

> TaskNexusScopedInventory CreateTaskNexusScopedInventory(ctx).TaskNexusScopedInventory(taskNexusScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.NexusScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskNexusScopedInventory := *openapiclient.NewTaskNexusScopedInventory("ClassId_example", "ObjectType_example") // TaskNexusScopedInventory | The 'task.NexusScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskNexusScopedInventory(context.Background()).TaskNexusScopedInventory(taskNexusScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskNexusScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskNexusScopedInventory`: TaskNexusScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskNexusScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskNexusScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskNexusScopedInventory** | [**TaskNexusScopedInventory**](TaskNexusScopedInventory.md) | The &#39;task.NexusScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskNexusScopedInventory**](TaskNexusScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskNexusSystemScopedInventory

> TaskNexusSystemScopedInventory CreateTaskNexusSystemScopedInventory(ctx).TaskNexusSystemScopedInventory(taskNexusSystemScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.NexusSystemScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskNexusSystemScopedInventory := *openapiclient.NewTaskNexusSystemScopedInventory("ClassId_example", "ObjectType_example") // TaskNexusSystemScopedInventory | The 'task.NexusSystemScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskNexusSystemScopedInventory(context.Background()).TaskNexusSystemScopedInventory(taskNexusSystemScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskNexusSystemScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskNexusSystemScopedInventory`: TaskNexusSystemScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskNexusSystemScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskNexusSystemScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskNexusSystemScopedInventory** | [**TaskNexusSystemScopedInventory**](TaskNexusSystemScopedInventory.md) | The &#39;task.NexusSystemScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskNexusSystemScopedInventory**](TaskNexusSystemScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskNexusVlanScopedInventory

> TaskNexusVlanScopedInventory CreateTaskNexusVlanScopedInventory(ctx).TaskNexusVlanScopedInventory(taskNexusVlanScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.NexusVlanScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskNexusVlanScopedInventory := *openapiclient.NewTaskNexusVlanScopedInventory("ClassId_example", "ObjectType_example") // TaskNexusVlanScopedInventory | The 'task.NexusVlanScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskNexusVlanScopedInventory(context.Background()).TaskNexusVlanScopedInventory(taskNexusVlanScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskNexusVlanScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskNexusVlanScopedInventory`: TaskNexusVlanScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskNexusVlanScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskNexusVlanScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskNexusVlanScopedInventory** | [**TaskNexusVlanScopedInventory**](TaskNexusVlanScopedInventory.md) | The &#39;task.NexusVlanScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskNexusVlanScopedInventory**](TaskNexusVlanScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskPublicCloudScopedInventory

> TaskPublicCloudScopedInventory CreateTaskPublicCloudScopedInventory(ctx).TaskPublicCloudScopedInventory(taskPublicCloudScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.PublicCloudScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskPublicCloudScopedInventory := *openapiclient.NewTaskPublicCloudScopedInventory("ClassId_example", "ObjectType_example") // TaskPublicCloudScopedInventory | The 'task.PublicCloudScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskPublicCloudScopedInventory(context.Background()).TaskPublicCloudScopedInventory(taskPublicCloudScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskPublicCloudScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskPublicCloudScopedInventory`: TaskPublicCloudScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskPublicCloudScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskPublicCloudScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskPublicCloudScopedInventory** | [**TaskPublicCloudScopedInventory**](TaskPublicCloudScopedInventory.md) | The &#39;task.PublicCloudScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskPublicCloudScopedInventory**](TaskPublicCloudScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskPureScopedInventory

> TaskPureScopedInventory CreateTaskPureScopedInventory(ctx).TaskPureScopedInventory(taskPureScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.PureScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskPureScopedInventory := *openapiclient.NewTaskPureScopedInventory("ClassId_example", "ObjectType_example") // TaskPureScopedInventory | The 'task.PureScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskPureScopedInventory(context.Background()).TaskPureScopedInventory(taskPureScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskPureScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskPureScopedInventory`: TaskPureScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskPureScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskPureScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskPureScopedInventory** | [**TaskPureScopedInventory**](TaskPureScopedInventory.md) | The &#39;task.PureScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskPureScopedInventory**](TaskPureScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskServerScopedInventory

> TaskServerScopedInventory CreateTaskServerScopedInventory(ctx).TaskServerScopedInventory(taskServerScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.ServerScopedInventory' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskServerScopedInventory := *openapiclient.NewTaskServerScopedInventory("ClassId_example", "ObjectType_example") // TaskServerScopedInventory | The 'task.ServerScopedInventory' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskServerScopedInventory(context.Background()).TaskServerScopedInventory(taskServerScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskServerScopedInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskServerScopedInventory`: TaskServerScopedInventory
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskServerScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskServerScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskServerScopedInventory** | [**TaskServerScopedInventory**](TaskServerScopedInventory.md) | The &#39;task.ServerScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskServerScopedInventory**](TaskServerScopedInventory.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskWorkflowAction

> TaskWorkflowAction CreateTaskWorkflowAction(ctx).TaskWorkflowAction(taskWorkflowAction).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.WorkflowAction' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	taskWorkflowAction := *openapiclient.NewTaskWorkflowAction("ClassId_example", "ObjectType_example") // TaskWorkflowAction | The 'task.WorkflowAction' resource to create.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.CreateTaskWorkflowAction(context.Background()).TaskWorkflowAction(taskWorkflowAction).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskWorkflowAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskWorkflowAction`: TaskWorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskWorkflowAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskWorkflowActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskWorkflowAction** | [**TaskWorkflowAction**](TaskWorkflowAction.md) | The &#39;task.WorkflowAction&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskWorkflowAction**](TaskWorkflowAction.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskWorkflowActionByMoid

> TaskWorkflowAction GetTaskWorkflowActionByMoid(ctx, moid).Execute()

Read a 'task.WorkflowAction' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	moid := "moid_example" // string | The unique Moid identifier of a resource instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.GetTaskWorkflowActionByMoid(context.Background(), moid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskWorkflowActionByMoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskWorkflowActionByMoid`: TaskWorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskWorkflowActionByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskWorkflowActionByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskWorkflowAction**](TaskWorkflowAction.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskWorkflowActionList

> TaskWorkflowActionResponse GetTaskWorkflowActionList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'task.WorkflowAction' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	filter := "$filter=CreateTime gt 2012-08-29T21:58:33Z" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
	orderby := "$orderby=CreationTime" // string | Determines what properties are used to sort the collection of resources. (optional)
	top := int32($top=10) // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
	skip := int32($skip=100) // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
	select_ := "$select=CreateTime,ModTime" // string | Specifies a subset of properties to return. (optional) (default to "")
	expand := "$expand=DisplayNames" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
	apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e., the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
	count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
	inlinecount := "$inlinecount=true" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
	at := "at=VersionType eq 'Configured'" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
	tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.GetTaskWorkflowActionList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskWorkflowActionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskWorkflowActionList`: TaskWorkflowActionResponse
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskWorkflowActionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskWorkflowActionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e., the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**TaskWorkflowActionResponse**](TaskWorkflowActionResponse.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTaskWorkflowAction

> TaskWorkflowAction PatchTaskWorkflowAction(ctx, moid).TaskWorkflowAction(taskWorkflowAction).IfMatch(ifMatch).Execute()

Update a 'task.WorkflowAction' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	moid := "moid_example" // string | The unique Moid identifier of a resource instance.
	taskWorkflowAction := *openapiclient.NewTaskWorkflowAction("ClassId_example", "ObjectType_example") // TaskWorkflowAction | The 'task.WorkflowAction' resource to update.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.PatchTaskWorkflowAction(context.Background(), moid).TaskWorkflowAction(taskWorkflowAction).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.PatchTaskWorkflowAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTaskWorkflowAction`: TaskWorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.PatchTaskWorkflowAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTaskWorkflowActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskWorkflowAction** | [**TaskWorkflowAction**](TaskWorkflowAction.md) | The &#39;task.WorkflowAction&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**TaskWorkflowAction**](TaskWorkflowAction.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskWorkflowAction

> TaskWorkflowAction UpdateTaskWorkflowAction(ctx, moid).TaskWorkflowAction(taskWorkflowAction).IfMatch(ifMatch).Execute()

Update a 'task.WorkflowAction' resource.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CiscoDevNet/intersight-go"
)

func main() {
	moid := "moid_example" // string | The unique Moid identifier of a resource instance.
	taskWorkflowAction := *openapiclient.NewTaskWorkflowAction("ClassId_example", "ObjectType_example") // TaskWorkflowAction | The 'task.WorkflowAction' resource to update.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskApi.UpdateTaskWorkflowAction(context.Background(), moid).TaskWorkflowAction(taskWorkflowAction).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.UpdateTaskWorkflowAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTaskWorkflowAction`: TaskWorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `TaskApi.UpdateTaskWorkflowAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskWorkflowActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskWorkflowAction** | [**TaskWorkflowAction**](TaskWorkflowAction.md) | The &#39;task.WorkflowAction&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**TaskWorkflowAction**](TaskWorkflowAction.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

