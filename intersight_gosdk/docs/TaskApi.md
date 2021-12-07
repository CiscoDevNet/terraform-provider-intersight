# \TaskApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskHitachiScopedInventory**](TaskApi.md#CreateTaskHitachiScopedInventory) | **Post** /api/v1/task/HitachiScopedInventories | Create a &#39;task.HitachiScopedInventory&#39; resource.
[**CreateTaskIweScopedInventory**](TaskApi.md#CreateTaskIweScopedInventory) | **Post** /api/v1/task/IweScopedInventories | Create a &#39;task.IweScopedInventory&#39; resource.
[**CreateTaskNetAppScopedInventory**](TaskApi.md#CreateTaskNetAppScopedInventory) | **Post** /api/v1/task/NetAppScopedInventories | Create a &#39;task.NetAppScopedInventory&#39; resource.
[**CreateTaskPublicCloudScopedInventory**](TaskApi.md#CreateTaskPublicCloudScopedInventory) | **Post** /api/v1/task/PublicCloudScopedInventories | Create a &#39;task.PublicCloudScopedInventory&#39; resource.
[**CreateTaskPureScopedInventory**](TaskApi.md#CreateTaskPureScopedInventory) | **Post** /api/v1/task/PureScopedInventories | Create a &#39;task.PureScopedInventory&#39; resource.
[**CreateTaskServerScopedInventory**](TaskApi.md#CreateTaskServerScopedInventory) | **Post** /api/v1/task/ServerScopedInventories | Create a &#39;task.ServerScopedInventory&#39; resource.



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
    openapiclient "./openapi"
)

func main() {
    taskHitachiScopedInventory := *openapiclient.NewTaskHitachiScopedInventory("ClassId_example", "ObjectType_example") // TaskHitachiScopedInventory | The 'task.HitachiScopedInventory' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTaskHitachiScopedInventory(context.Background()).TaskHitachiScopedInventory(taskHitachiScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
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

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskIweScopedInventory

> TaskIweScopedInventory CreateTaskIweScopedInventory(ctx).TaskIweScopedInventory(taskIweScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'task.IweScopedInventory' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskIweScopedInventory := *openapiclient.NewTaskIweScopedInventory("ClassId_example", "ObjectType_example") // TaskIweScopedInventory | The 'task.IweScopedInventory' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTaskIweScopedInventory(context.Background()).TaskIweScopedInventory(taskIweScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTaskIweScopedInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTaskIweScopedInventory`: TaskIweScopedInventory
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTaskIweScopedInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskIweScopedInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskIweScopedInventory** | [**TaskIweScopedInventory**](TaskIweScopedInventory.md) | The &#39;task.IweScopedInventory&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**TaskIweScopedInventory**](TaskIweScopedInventory.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

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
    openapiclient "./openapi"
)

func main() {
    taskNetAppScopedInventory := *openapiclient.NewTaskNetAppScopedInventory("ClassId_example", "ObjectType_example") // TaskNetAppScopedInventory | The 'task.NetAppScopedInventory' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTaskNetAppScopedInventory(context.Background()).TaskNetAppScopedInventory(taskNetAppScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
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

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

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
    openapiclient "./openapi"
)

func main() {
    taskPublicCloudScopedInventory := *openapiclient.NewTaskPublicCloudScopedInventory("ClassId_example", "ObjectType_example") // TaskPublicCloudScopedInventory | The 'task.PublicCloudScopedInventory' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTaskPublicCloudScopedInventory(context.Background()).TaskPublicCloudScopedInventory(taskPublicCloudScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
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

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

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
    openapiclient "./openapi"
)

func main() {
    taskPureScopedInventory := *openapiclient.NewTaskPureScopedInventory("ClassId_example", "ObjectType_example") // TaskPureScopedInventory | The 'task.PureScopedInventory' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTaskPureScopedInventory(context.Background()).TaskPureScopedInventory(taskPureScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
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

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

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
    openapiclient "./openapi"
)

func main() {
    taskServerScopedInventory := *openapiclient.NewTaskServerScopedInventory("ClassId_example", "ObjectType_example") // TaskServerScopedInventory | The 'task.ServerScopedInventory' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTaskServerScopedInventory(context.Background()).TaskServerScopedInventory(taskServerScopedInventory).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
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

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

