# \OnpremApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOnpremApplianceSystemInfoByMoid**](OnpremApi.md#GetOnpremApplianceSystemInfoByMoid) | **Get** /api/v1/onprem/ApplianceSystemInfos/{Moid} | Read a &#39;onprem.ApplianceSystemInfo&#39; resource.
[**GetOnpremApplianceSystemInfoList**](OnpremApi.md#GetOnpremApplianceSystemInfoList) | **Get** /api/v1/onprem/ApplianceSystemInfos | Read a &#39;onprem.ApplianceSystemInfo&#39; resource.
[**GetOnpremAuditRecordByMoid**](OnpremApi.md#GetOnpremAuditRecordByMoid) | **Get** /api/v1/onprem/AuditRecords/{Moid} | Read a &#39;onprem.AuditRecord&#39; resource.
[**GetOnpremAuditRecordList**](OnpremApi.md#GetOnpremAuditRecordList) | **Get** /api/v1/onprem/AuditRecords | Read a &#39;onprem.AuditRecord&#39; resource.
[**GetOnpremUpgradeByMoid**](OnpremApi.md#GetOnpremUpgradeByMoid) | **Get** /api/v1/onprem/Upgrades/{Moid} | Read a &#39;onprem.Upgrade&#39; resource.
[**GetOnpremUpgradeList**](OnpremApi.md#GetOnpremUpgradeList) | **Get** /api/v1/onprem/Upgrades | Read a &#39;onprem.Upgrade&#39; resource.
[**GetOnpremUserPreferenceByMoid**](OnpremApi.md#GetOnpremUserPreferenceByMoid) | **Get** /api/v1/onprem/UserPreferences/{Moid} | Read a &#39;onprem.UserPreference&#39; resource.
[**GetOnpremUserPreferenceList**](OnpremApi.md#GetOnpremUserPreferenceList) | **Get** /api/v1/onprem/UserPreferences | Read a &#39;onprem.UserPreference&#39; resource.
[**PatchOnpremUpgrade**](OnpremApi.md#PatchOnpremUpgrade) | **Patch** /api/v1/onprem/Upgrades/{Moid} | Update a &#39;onprem.Upgrade&#39; resource.
[**PatchOnpremUserPreference**](OnpremApi.md#PatchOnpremUserPreference) | **Patch** /api/v1/onprem/UserPreferences/{Moid} | Update a &#39;onprem.UserPreference&#39; resource.
[**UpdateOnpremUpgrade**](OnpremApi.md#UpdateOnpremUpgrade) | **Post** /api/v1/onprem/Upgrades/{Moid} | Update a &#39;onprem.Upgrade&#39; resource.
[**UpdateOnpremUserPreference**](OnpremApi.md#UpdateOnpremUserPreference) | **Post** /api/v1/onprem/UserPreferences/{Moid} | Update a &#39;onprem.UserPreference&#39; resource.



## GetOnpremApplianceSystemInfoByMoid

> OnpremApplianceSystemInfo GetOnpremApplianceSystemInfoByMoid(ctx, moid).Execute()

Read a 'onprem.ApplianceSystemInfo' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremApplianceSystemInfoByMoid(context.Background(), moid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremApplianceSystemInfoByMoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremApplianceSystemInfoByMoid`: OnpremApplianceSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremApplianceSystemInfoByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremApplianceSystemInfoByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnpremApplianceSystemInfo**](OnpremApplianceSystemInfo.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremApplianceSystemInfoList

> OnpremApplianceSystemInfoResponse GetOnpremApplianceSystemInfoList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'onprem.ApplianceSystemInfo' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremApplianceSystemInfoList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremApplianceSystemInfoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremApplianceSystemInfoList`: OnpremApplianceSystemInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremApplianceSystemInfoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremApplianceSystemInfoListRequest struct via the builder pattern


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

[**OnpremApplianceSystemInfoResponse**](OnpremApplianceSystemInfoResponse.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremAuditRecordByMoid

> OnpremAuditRecord GetOnpremAuditRecordByMoid(ctx, moid).Execute()

Read a 'onprem.AuditRecord' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremAuditRecordByMoid(context.Background(), moid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremAuditRecordByMoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremAuditRecordByMoid`: OnpremAuditRecord
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremAuditRecordByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremAuditRecordByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnpremAuditRecord**](OnpremAuditRecord.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremAuditRecordList

> OnpremAuditRecordResponse GetOnpremAuditRecordList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'onprem.AuditRecord' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremAuditRecordList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremAuditRecordList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremAuditRecordList`: OnpremAuditRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremAuditRecordList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremAuditRecordListRequest struct via the builder pattern


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

[**OnpremAuditRecordResponse**](OnpremAuditRecordResponse.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremUpgradeByMoid

> OnpremUpgrade GetOnpremUpgradeByMoid(ctx, moid).Execute()

Read a 'onprem.Upgrade' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremUpgradeByMoid(context.Background(), moid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremUpgradeByMoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremUpgradeByMoid`: OnpremUpgrade
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremUpgradeByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremUpgradeByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnpremUpgrade**](OnpremUpgrade.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremUpgradeList

> OnpremUpgradeResponse GetOnpremUpgradeList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'onprem.Upgrade' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremUpgradeList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremUpgradeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremUpgradeList`: OnpremUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremUpgradeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremUpgradeListRequest struct via the builder pattern


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

[**OnpremUpgradeResponse**](OnpremUpgradeResponse.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremUserPreferenceByMoid

> OnpremUserPreference GetOnpremUserPreferenceByMoid(ctx, moid).Execute()

Read a 'onprem.UserPreference' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremUserPreferenceByMoid(context.Background(), moid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremUserPreferenceByMoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremUserPreferenceByMoid`: OnpremUserPreference
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremUserPreferenceByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremUserPreferenceByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnpremUserPreference**](OnpremUserPreference.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnpremUserPreferenceList

> OnpremUserPreferenceResponse GetOnpremUserPreferenceList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'onprem.UserPreference' resource.

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
	resp, r, err := apiClient.OnpremApi.GetOnpremUserPreferenceList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.GetOnpremUserPreferenceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnpremUserPreferenceList`: OnpremUserPreferenceResponse
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.GetOnpremUserPreferenceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOnpremUserPreferenceListRequest struct via the builder pattern


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

[**OnpremUserPreferenceResponse**](OnpremUserPreferenceResponse.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOnpremUpgrade

> OnpremUpgrade PatchOnpremUpgrade(ctx, moid).OnpremUpgrade(onpremUpgrade).IfMatch(ifMatch).Execute()

Update a 'onprem.Upgrade' resource.

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
	onpremUpgrade := *openapiclient.NewOnpremUpgrade("ClassId_example", "ObjectType_example") // OnpremUpgrade | The 'onprem.Upgrade' resource to update.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnpremApi.PatchOnpremUpgrade(context.Background(), moid).OnpremUpgrade(onpremUpgrade).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.PatchOnpremUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOnpremUpgrade`: OnpremUpgrade
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.PatchOnpremUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOnpremUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onpremUpgrade** | [**OnpremUpgrade**](OnpremUpgrade.md) | The &#39;onprem.Upgrade&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**OnpremUpgrade**](OnpremUpgrade.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOnpremUserPreference

> OnpremUserPreference PatchOnpremUserPreference(ctx, moid).OnpremUserPreference(onpremUserPreference).IfMatch(ifMatch).Execute()

Update a 'onprem.UserPreference' resource.

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
	onpremUserPreference := *openapiclient.NewOnpremUserPreference("ClassId_example", "ObjectType_example") // OnpremUserPreference | The 'onprem.UserPreference' resource to update.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnpremApi.PatchOnpremUserPreference(context.Background(), moid).OnpremUserPreference(onpremUserPreference).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.PatchOnpremUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOnpremUserPreference`: OnpremUserPreference
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.PatchOnpremUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOnpremUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onpremUserPreference** | [**OnpremUserPreference**](OnpremUserPreference.md) | The &#39;onprem.UserPreference&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**OnpremUserPreference**](OnpremUserPreference.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnpremUpgrade

> OnpremUpgrade UpdateOnpremUpgrade(ctx, moid).OnpremUpgrade(onpremUpgrade).IfMatch(ifMatch).Execute()

Update a 'onprem.Upgrade' resource.

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
	onpremUpgrade := *openapiclient.NewOnpremUpgrade("ClassId_example", "ObjectType_example") // OnpremUpgrade | The 'onprem.Upgrade' resource to update.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnpremApi.UpdateOnpremUpgrade(context.Background(), moid).OnpremUpgrade(onpremUpgrade).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.UpdateOnpremUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOnpremUpgrade`: OnpremUpgrade
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.UpdateOnpremUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnpremUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onpremUpgrade** | [**OnpremUpgrade**](OnpremUpgrade.md) | The &#39;onprem.Upgrade&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**OnpremUpgrade**](OnpremUpgrade.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnpremUserPreference

> OnpremUserPreference UpdateOnpremUserPreference(ctx, moid).OnpremUserPreference(onpremUserPreference).IfMatch(ifMatch).Execute()

Update a 'onprem.UserPreference' resource.

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
	onpremUserPreference := *openapiclient.NewOnpremUserPreference("ClassId_example", "ObjectType_example") // OnpremUserPreference | The 'onprem.UserPreference' resource to update.
	ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnpremApi.UpdateOnpremUserPreference(context.Background(), moid).OnpremUserPreference(onpremUserPreference).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnpremApi.UpdateOnpremUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOnpremUserPreference`: OnpremUserPreference
	fmt.Fprintf(os.Stdout, "Response from `OnpremApi.UpdateOnpremUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnpremUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onpremUserPreference** | [**OnpremUserPreference**](OnpremUserPreference.md) | The &#39;onprem.UserPreference&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**OnpremUserPreference**](OnpremUserPreference.md)

### Authorization

[http_signature](../README.md#http_signature), [cookieAuth](../README.md#cookieAuth), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

