# \VirtualizationApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualizationVmwareClusterByMoid**](VirtualizationApi.md#GetVirtualizationVmwareClusterByMoid) | **Get** /api/v1/virtualization/VmwareClusters/{Moid} | Read a &#39;virtualization.VmwareCluster&#39; resource.
[**GetVirtualizationVmwareClusterList**](VirtualizationApi.md#GetVirtualizationVmwareClusterList) | **Get** /api/v1/virtualization/VmwareClusters | Read a &#39;virtualization.VmwareCluster&#39; resource.
[**GetVirtualizationVmwareDatacenterByMoid**](VirtualizationApi.md#GetVirtualizationVmwareDatacenterByMoid) | **Get** /api/v1/virtualization/VmwareDatacenters/{Moid} | Read a &#39;virtualization.VmwareDatacenter&#39; resource.
[**GetVirtualizationVmwareDatacenterList**](VirtualizationApi.md#GetVirtualizationVmwareDatacenterList) | **Get** /api/v1/virtualization/VmwareDatacenters | Read a &#39;virtualization.VmwareDatacenter&#39; resource.
[**GetVirtualizationVmwareDatastoreByMoid**](VirtualizationApi.md#GetVirtualizationVmwareDatastoreByMoid) | **Get** /api/v1/virtualization/VmwareDatastores/{Moid} | Read a &#39;virtualization.VmwareDatastore&#39; resource.
[**GetVirtualizationVmwareDatastoreList**](VirtualizationApi.md#GetVirtualizationVmwareDatastoreList) | **Get** /api/v1/virtualization/VmwareDatastores | Read a &#39;virtualization.VmwareDatastore&#39; resource.
[**GetVirtualizationVmwareHostByMoid**](VirtualizationApi.md#GetVirtualizationVmwareHostByMoid) | **Get** /api/v1/virtualization/VmwareHosts/{Moid} | Read a &#39;virtualization.VmwareHost&#39; resource.
[**GetVirtualizationVmwareHostList**](VirtualizationApi.md#GetVirtualizationVmwareHostList) | **Get** /api/v1/virtualization/VmwareHosts | Read a &#39;virtualization.VmwareHost&#39; resource.
[**GetVirtualizationVmwareVcenterByMoid**](VirtualizationApi.md#GetVirtualizationVmwareVcenterByMoid) | **Get** /api/v1/virtualization/VmwareVcenters/{Moid} | Read a &#39;virtualization.VmwareVcenter&#39; resource.
[**GetVirtualizationVmwareVcenterList**](VirtualizationApi.md#GetVirtualizationVmwareVcenterList) | **Get** /api/v1/virtualization/VmwareVcenters | Read a &#39;virtualization.VmwareVcenter&#39; resource.
[**GetVirtualizationVmwareVirtualMachineByMoid**](VirtualizationApi.md#GetVirtualizationVmwareVirtualMachineByMoid) | **Get** /api/v1/virtualization/VmwareVirtualMachines/{Moid} | Read a &#39;virtualization.VmwareVirtualMachine&#39; resource.
[**GetVirtualizationVmwareVirtualMachineList**](VirtualizationApi.md#GetVirtualizationVmwareVirtualMachineList) | **Get** /api/v1/virtualization/VmwareVirtualMachines | Read a &#39;virtualization.VmwareVirtualMachine&#39; resource.
[**PatchVirtualizationVmwareCluster**](VirtualizationApi.md#PatchVirtualizationVmwareCluster) | **Patch** /api/v1/virtualization/VmwareClusters/{Moid} | Update a &#39;virtualization.VmwareCluster&#39; resource.
[**PatchVirtualizationVmwareDatacenter**](VirtualizationApi.md#PatchVirtualizationVmwareDatacenter) | **Patch** /api/v1/virtualization/VmwareDatacenters/{Moid} | Update a &#39;virtualization.VmwareDatacenter&#39; resource.
[**PatchVirtualizationVmwareDatastore**](VirtualizationApi.md#PatchVirtualizationVmwareDatastore) | **Patch** /api/v1/virtualization/VmwareDatastores/{Moid} | Update a &#39;virtualization.VmwareDatastore&#39; resource.
[**PatchVirtualizationVmwareHost**](VirtualizationApi.md#PatchVirtualizationVmwareHost) | **Patch** /api/v1/virtualization/VmwareHosts/{Moid} | Update a &#39;virtualization.VmwareHost&#39; resource.
[**PatchVirtualizationVmwareVirtualMachine**](VirtualizationApi.md#PatchVirtualizationVmwareVirtualMachine) | **Patch** /api/v1/virtualization/VmwareVirtualMachines/{Moid} | Update a &#39;virtualization.VmwareVirtualMachine&#39; resource.
[**UpdateVirtualizationVmwareCluster**](VirtualizationApi.md#UpdateVirtualizationVmwareCluster) | **Post** /api/v1/virtualization/VmwareClusters/{Moid} | Update a &#39;virtualization.VmwareCluster&#39; resource.
[**UpdateVirtualizationVmwareDatacenter**](VirtualizationApi.md#UpdateVirtualizationVmwareDatacenter) | **Post** /api/v1/virtualization/VmwareDatacenters/{Moid} | Update a &#39;virtualization.VmwareDatacenter&#39; resource.
[**UpdateVirtualizationVmwareDatastore**](VirtualizationApi.md#UpdateVirtualizationVmwareDatastore) | **Post** /api/v1/virtualization/VmwareDatastores/{Moid} | Update a &#39;virtualization.VmwareDatastore&#39; resource.
[**UpdateVirtualizationVmwareHost**](VirtualizationApi.md#UpdateVirtualizationVmwareHost) | **Post** /api/v1/virtualization/VmwareHosts/{Moid} | Update a &#39;virtualization.VmwareHost&#39; resource.
[**UpdateVirtualizationVmwareVirtualMachine**](VirtualizationApi.md#UpdateVirtualizationVmwareVirtualMachine) | **Post** /api/v1/virtualization/VmwareVirtualMachines/{Moid} | Update a &#39;virtualization.VmwareVirtualMachine&#39; resource.



## GetVirtualizationVmwareClusterByMoid

> VirtualizationVmwareCluster GetVirtualizationVmwareClusterByMoid(ctx, moid).Execute()

Read a 'virtualization.VmwareCluster' resource.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareClusterByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareClusterByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareClusterByMoid`: VirtualizationVmwareCluster
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareClusterByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareClusterByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualizationVmwareCluster**](virtualization.VmwareCluster.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareClusterList

> VirtualizationVmwareClusterResponse GetVirtualizationVmwareClusterList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'virtualization.VmwareCluster' resource.

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
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareClusterList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareClusterList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareClusterList`: VirtualizationVmwareClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareClusterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareClusterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**VirtualizationVmwareClusterResponse**](virtualization.VmwareCluster.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareDatacenterByMoid

> VirtualizationVmwareDatacenter GetVirtualizationVmwareDatacenterByMoid(ctx, moid).Execute()

Read a 'virtualization.VmwareDatacenter' resource.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareDatacenterByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareDatacenterByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareDatacenterByMoid`: VirtualizationVmwareDatacenter
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareDatacenterByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareDatacenterByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualizationVmwareDatacenter**](virtualization.VmwareDatacenter.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareDatacenterList

> VirtualizationVmwareDatacenterResponse GetVirtualizationVmwareDatacenterList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'virtualization.VmwareDatacenter' resource.

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
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareDatacenterList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareDatacenterList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareDatacenterList`: VirtualizationVmwareDatacenterResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareDatacenterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareDatacenterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**VirtualizationVmwareDatacenterResponse**](virtualization.VmwareDatacenter.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareDatastoreByMoid

> VirtualizationVmwareDatastore GetVirtualizationVmwareDatastoreByMoid(ctx, moid).Execute()

Read a 'virtualization.VmwareDatastore' resource.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareDatastoreByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareDatastoreByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareDatastoreByMoid`: VirtualizationVmwareDatastore
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareDatastoreByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareDatastoreByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualizationVmwareDatastore**](virtualization.VmwareDatastore.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareDatastoreList

> VirtualizationVmwareDatastoreResponse GetVirtualizationVmwareDatastoreList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'virtualization.VmwareDatastore' resource.

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
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareDatastoreList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareDatastoreList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareDatastoreList`: VirtualizationVmwareDatastoreResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareDatastoreList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareDatastoreListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**VirtualizationVmwareDatastoreResponse**](virtualization.VmwareDatastore.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareHostByMoid

> VirtualizationVmwareHost GetVirtualizationVmwareHostByMoid(ctx, moid).Execute()

Read a 'virtualization.VmwareHost' resource.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareHostByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareHostByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareHostByMoid`: VirtualizationVmwareHost
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareHostByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareHostByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualizationVmwareHost**](virtualization.VmwareHost.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareHostList

> VirtualizationVmwareHostResponse GetVirtualizationVmwareHostList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'virtualization.VmwareHost' resource.

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
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareHostList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareHostList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareHostList`: VirtualizationVmwareHostResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareHostList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareHostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**VirtualizationVmwareHostResponse**](virtualization.VmwareHost.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareVcenterByMoid

> VirtualizationVmwareVcenter GetVirtualizationVmwareVcenterByMoid(ctx, moid).Execute()

Read a 'virtualization.VmwareVcenter' resource.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareVcenterByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareVcenterByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareVcenterByMoid`: VirtualizationVmwareVcenter
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareVcenterByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareVcenterByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualizationVmwareVcenter**](virtualization.VmwareVcenter.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareVcenterList

> VirtualizationVmwareVcenterResponse GetVirtualizationVmwareVcenterList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'virtualization.VmwareVcenter' resource.

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
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareVcenterList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareVcenterList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareVcenterList`: VirtualizationVmwareVcenterResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareVcenterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareVcenterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**VirtualizationVmwareVcenterResponse**](virtualization.VmwareVcenter.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareVirtualMachineByMoid

> VirtualizationVmwareVirtualMachine GetVirtualizationVmwareVirtualMachineByMoid(ctx, moid).Execute()

Read a 'virtualization.VmwareVirtualMachine' resource.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareVirtualMachineByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareVirtualMachineByMoid`: VirtualizationVmwareVirtualMachine
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareVirtualMachineByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareVirtualMachineByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualizationVmwareVirtualMachine**](virtualization.VmwareVirtualMachine.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualizationVmwareVirtualMachineList

> VirtualizationVmwareVirtualMachineResponse GetVirtualizationVmwareVirtualMachineList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'virtualization.VmwareVirtualMachine' resource.

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
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.GetVirtualizationVmwareVirtualMachineList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.GetVirtualizationVmwareVirtualMachineList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualizationVmwareVirtualMachineList`: VirtualizationVmwareVirtualMachineResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.GetVirtualizationVmwareVirtualMachineList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualizationVmwareVirtualMachineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**VirtualizationVmwareVirtualMachineResponse**](virtualization.VmwareVirtualMachine.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVirtualizationVmwareCluster

> VirtualizationVmwareCluster PatchVirtualizationVmwareCluster(ctx, moid).VirtualizationVmwareCluster(virtualizationVmwareCluster).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareCluster' resource.

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
    virtualizationVmwareCluster := *openapiclient.Newvirtualization.VmwareCluster("ClassId_example", "ObjectType_example") // VirtualizationVmwareCluster | The 'virtualization.VmwareCluster' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.PatchVirtualizationVmwareCluster(context.Background(), moid).VirtualizationVmwareCluster(virtualizationVmwareCluster).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.PatchVirtualizationVmwareCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchVirtualizationVmwareCluster`: VirtualizationVmwareCluster
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.PatchVirtualizationVmwareCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVirtualizationVmwareClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareCluster** | [**VirtualizationVmwareCluster**](VirtualizationVmwareCluster.md) | The &#39;virtualization.VmwareCluster&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareCluster**](virtualization.VmwareCluster.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVirtualizationVmwareDatacenter

> VirtualizationVmwareDatacenter PatchVirtualizationVmwareDatacenter(ctx, moid).VirtualizationVmwareDatacenter(virtualizationVmwareDatacenter).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareDatacenter' resource.

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
    virtualizationVmwareDatacenter := *openapiclient.Newvirtualization.VmwareDatacenter("ClassId_example", "ObjectType_example") // VirtualizationVmwareDatacenter | The 'virtualization.VmwareDatacenter' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.PatchVirtualizationVmwareDatacenter(context.Background(), moid).VirtualizationVmwareDatacenter(virtualizationVmwareDatacenter).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.PatchVirtualizationVmwareDatacenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchVirtualizationVmwareDatacenter`: VirtualizationVmwareDatacenter
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.PatchVirtualizationVmwareDatacenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVirtualizationVmwareDatacenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareDatacenter** | [**VirtualizationVmwareDatacenter**](VirtualizationVmwareDatacenter.md) | The &#39;virtualization.VmwareDatacenter&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareDatacenter**](virtualization.VmwareDatacenter.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVirtualizationVmwareDatastore

> VirtualizationVmwareDatastore PatchVirtualizationVmwareDatastore(ctx, moid).VirtualizationVmwareDatastore(virtualizationVmwareDatastore).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareDatastore' resource.

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
    virtualizationVmwareDatastore := *openapiclient.Newvirtualization.VmwareDatastore("ClassId_example", "ObjectType_example") // VirtualizationVmwareDatastore | The 'virtualization.VmwareDatastore' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.PatchVirtualizationVmwareDatastore(context.Background(), moid).VirtualizationVmwareDatastore(virtualizationVmwareDatastore).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.PatchVirtualizationVmwareDatastore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchVirtualizationVmwareDatastore`: VirtualizationVmwareDatastore
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.PatchVirtualizationVmwareDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVirtualizationVmwareDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareDatastore** | [**VirtualizationVmwareDatastore**](VirtualizationVmwareDatastore.md) | The &#39;virtualization.VmwareDatastore&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareDatastore**](virtualization.VmwareDatastore.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVirtualizationVmwareHost

> VirtualizationVmwareHost PatchVirtualizationVmwareHost(ctx, moid).VirtualizationVmwareHost(virtualizationVmwareHost).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareHost' resource.

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
    virtualizationVmwareHost := *openapiclient.Newvirtualization.VmwareHost("ClassId_example", "ObjectType_example") // VirtualizationVmwareHost | The 'virtualization.VmwareHost' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.PatchVirtualizationVmwareHost(context.Background(), moid).VirtualizationVmwareHost(virtualizationVmwareHost).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.PatchVirtualizationVmwareHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchVirtualizationVmwareHost`: VirtualizationVmwareHost
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.PatchVirtualizationVmwareHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVirtualizationVmwareHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareHost** | [**VirtualizationVmwareHost**](VirtualizationVmwareHost.md) | The &#39;virtualization.VmwareHost&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareHost**](virtualization.VmwareHost.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchVirtualizationVmwareVirtualMachine

> VirtualizationVmwareVirtualMachine PatchVirtualizationVmwareVirtualMachine(ctx, moid).VirtualizationVmwareVirtualMachine(virtualizationVmwareVirtualMachine).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareVirtualMachine' resource.

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
    virtualizationVmwareVirtualMachine := *openapiclient.Newvirtualization.VmwareVirtualMachine("ClassId_example", "ObjectType_example") // VirtualizationVmwareVirtualMachine | The 'virtualization.VmwareVirtualMachine' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.PatchVirtualizationVmwareVirtualMachine(context.Background(), moid).VirtualizationVmwareVirtualMachine(virtualizationVmwareVirtualMachine).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.PatchVirtualizationVmwareVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchVirtualizationVmwareVirtualMachine`: VirtualizationVmwareVirtualMachine
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.PatchVirtualizationVmwareVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchVirtualizationVmwareVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareVirtualMachine** | [**VirtualizationVmwareVirtualMachine**](VirtualizationVmwareVirtualMachine.md) | The &#39;virtualization.VmwareVirtualMachine&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareVirtualMachine**](virtualization.VmwareVirtualMachine.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualizationVmwareCluster

> VirtualizationVmwareCluster UpdateVirtualizationVmwareCluster(ctx, moid).VirtualizationVmwareCluster(virtualizationVmwareCluster).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareCluster' resource.

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
    virtualizationVmwareCluster := *openapiclient.Newvirtualization.VmwareCluster("ClassId_example", "ObjectType_example") // VirtualizationVmwareCluster | The 'virtualization.VmwareCluster' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.UpdateVirtualizationVmwareCluster(context.Background(), moid).VirtualizationVmwareCluster(virtualizationVmwareCluster).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.UpdateVirtualizationVmwareCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationVmwareCluster`: VirtualizationVmwareCluster
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.UpdateVirtualizationVmwareCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationVmwareClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareCluster** | [**VirtualizationVmwareCluster**](VirtualizationVmwareCluster.md) | The &#39;virtualization.VmwareCluster&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareCluster**](virtualization.VmwareCluster.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualizationVmwareDatacenter

> VirtualizationVmwareDatacenter UpdateVirtualizationVmwareDatacenter(ctx, moid).VirtualizationVmwareDatacenter(virtualizationVmwareDatacenter).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareDatacenter' resource.

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
    virtualizationVmwareDatacenter := *openapiclient.Newvirtualization.VmwareDatacenter("ClassId_example", "ObjectType_example") // VirtualizationVmwareDatacenter | The 'virtualization.VmwareDatacenter' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.UpdateVirtualizationVmwareDatacenter(context.Background(), moid).VirtualizationVmwareDatacenter(virtualizationVmwareDatacenter).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.UpdateVirtualizationVmwareDatacenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationVmwareDatacenter`: VirtualizationVmwareDatacenter
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.UpdateVirtualizationVmwareDatacenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationVmwareDatacenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareDatacenter** | [**VirtualizationVmwareDatacenter**](VirtualizationVmwareDatacenter.md) | The &#39;virtualization.VmwareDatacenter&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareDatacenter**](virtualization.VmwareDatacenter.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualizationVmwareDatastore

> VirtualizationVmwareDatastore UpdateVirtualizationVmwareDatastore(ctx, moid).VirtualizationVmwareDatastore(virtualizationVmwareDatastore).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareDatastore' resource.

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
    virtualizationVmwareDatastore := *openapiclient.Newvirtualization.VmwareDatastore("ClassId_example", "ObjectType_example") // VirtualizationVmwareDatastore | The 'virtualization.VmwareDatastore' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.UpdateVirtualizationVmwareDatastore(context.Background(), moid).VirtualizationVmwareDatastore(virtualizationVmwareDatastore).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.UpdateVirtualizationVmwareDatastore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationVmwareDatastore`: VirtualizationVmwareDatastore
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.UpdateVirtualizationVmwareDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationVmwareDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareDatastore** | [**VirtualizationVmwareDatastore**](VirtualizationVmwareDatastore.md) | The &#39;virtualization.VmwareDatastore&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareDatastore**](virtualization.VmwareDatastore.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualizationVmwareHost

> VirtualizationVmwareHost UpdateVirtualizationVmwareHost(ctx, moid).VirtualizationVmwareHost(virtualizationVmwareHost).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareHost' resource.

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
    virtualizationVmwareHost := *openapiclient.Newvirtualization.VmwareHost("ClassId_example", "ObjectType_example") // VirtualizationVmwareHost | The 'virtualization.VmwareHost' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.UpdateVirtualizationVmwareHost(context.Background(), moid).VirtualizationVmwareHost(virtualizationVmwareHost).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.UpdateVirtualizationVmwareHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationVmwareHost`: VirtualizationVmwareHost
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.UpdateVirtualizationVmwareHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationVmwareHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareHost** | [**VirtualizationVmwareHost**](VirtualizationVmwareHost.md) | The &#39;virtualization.VmwareHost&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareHost**](virtualization.VmwareHost.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualizationVmwareVirtualMachine

> VirtualizationVmwareVirtualMachine UpdateVirtualizationVmwareVirtualMachine(ctx, moid).VirtualizationVmwareVirtualMachine(virtualizationVmwareVirtualMachine).IfMatch(ifMatch).Execute()

Update a 'virtualization.VmwareVirtualMachine' resource.

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
    virtualizationVmwareVirtualMachine := *openapiclient.Newvirtualization.VmwareVirtualMachine("ClassId_example", "ObjectType_example") // VirtualizationVmwareVirtualMachine | The 'virtualization.VmwareVirtualMachine' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualizationApi.UpdateVirtualizationVmwareVirtualMachine(context.Background(), moid).VirtualizationVmwareVirtualMachine(virtualizationVmwareVirtualMachine).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualizationApi.UpdateVirtualizationVmwareVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualizationVmwareVirtualMachine`: VirtualizationVmwareVirtualMachine
    fmt.Fprintf(os.Stdout, "Response from `VirtualizationApi.UpdateVirtualizationVmwareVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualizationVmwareVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualizationVmwareVirtualMachine** | [**VirtualizationVmwareVirtualMachine**](VirtualizationVmwareVirtualMachine.md) | The &#39;virtualization.VmwareVirtualMachine&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**VirtualizationVmwareVirtualMachine**](virtualization.VmwareVirtualMachine.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

