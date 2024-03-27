# \FeedbackApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedbackFeedbackPost**](FeedbackApi.md#CreateFeedbackFeedbackPost) | **Post** /api/v1/feedback/FeedbackPosts | Create a &#39;feedback.FeedbackPost&#39; resource.
[**PatchFeedbackFeedbackPost**](FeedbackApi.md#PatchFeedbackFeedbackPost) | **Patch** /api/v1/feedback/FeedbackPosts/{Moid} | Update a &#39;feedback.FeedbackPost&#39; resource.
[**UpdateFeedbackFeedbackPost**](FeedbackApi.md#UpdateFeedbackFeedbackPost) | **Post** /api/v1/feedback/FeedbackPosts/{Moid} | Update a &#39;feedback.FeedbackPost&#39; resource.



## CreateFeedbackFeedbackPost

> FeedbackFeedbackPost CreateFeedbackFeedbackPost(ctx).FeedbackFeedbackPost(feedbackFeedbackPost).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'feedback.FeedbackPost' resource.

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
    feedbackFeedbackPost := *openapiclient.NewFeedbackFeedbackPost("ClassId_example", "ObjectType_example") // FeedbackFeedbackPost | The 'feedback.FeedbackPost' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedbackApi.CreateFeedbackFeedbackPost(context.Background()).FeedbackFeedbackPost(feedbackFeedbackPost).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbackApi.CreateFeedbackFeedbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeedbackFeedbackPost`: FeedbackFeedbackPost
    fmt.Fprintf(os.Stdout, "Response from `FeedbackApi.CreateFeedbackFeedbackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackFeedbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackFeedbackPost** | [**FeedbackFeedbackPost**](FeedbackFeedbackPost.md) | The &#39;feedback.FeedbackPost&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**FeedbackFeedbackPost**](FeedbackFeedbackPost.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFeedbackFeedbackPost

> FeedbackFeedbackPost PatchFeedbackFeedbackPost(ctx, moid).FeedbackFeedbackPost(feedbackFeedbackPost).IfMatch(ifMatch).Execute()

Update a 'feedback.FeedbackPost' resource.

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
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    feedbackFeedbackPost := *openapiclient.NewFeedbackFeedbackPost("ClassId_example", "ObjectType_example") // FeedbackFeedbackPost | The 'feedback.FeedbackPost' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedbackApi.PatchFeedbackFeedbackPost(context.Background(), moid).FeedbackFeedbackPost(feedbackFeedbackPost).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbackApi.PatchFeedbackFeedbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFeedbackFeedbackPost`: FeedbackFeedbackPost
    fmt.Fprintf(os.Stdout, "Response from `FeedbackApi.PatchFeedbackFeedbackPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFeedbackFeedbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feedbackFeedbackPost** | [**FeedbackFeedbackPost**](FeedbackFeedbackPost.md) | The &#39;feedback.FeedbackPost&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**FeedbackFeedbackPost**](FeedbackFeedbackPost.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeedbackFeedbackPost

> FeedbackFeedbackPost UpdateFeedbackFeedbackPost(ctx, moid).FeedbackFeedbackPost(feedbackFeedbackPost).IfMatch(ifMatch).Execute()

Update a 'feedback.FeedbackPost' resource.

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
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    feedbackFeedbackPost := *openapiclient.NewFeedbackFeedbackPost("ClassId_example", "ObjectType_example") // FeedbackFeedbackPost | The 'feedback.FeedbackPost' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeedbackApi.UpdateFeedbackFeedbackPost(context.Background(), moid).FeedbackFeedbackPost(feedbackFeedbackPost).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbackApi.UpdateFeedbackFeedbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeedbackFeedbackPost`: FeedbackFeedbackPost
    fmt.Fprintf(os.Stdout, "Response from `FeedbackApi.UpdateFeedbackFeedbackPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedbackFeedbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feedbackFeedbackPost** | [**FeedbackFeedbackPost**](FeedbackFeedbackPost.md) | The &#39;feedback.FeedbackPost&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**FeedbackFeedbackPost**](FeedbackFeedbackPost.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

