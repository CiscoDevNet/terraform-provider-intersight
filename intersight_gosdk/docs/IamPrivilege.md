# IamPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostnamePrefix** | Pointer to **string** | The hostname prefix of the resource corresponding to this privilege. For example \\&#39;sentry\\&#39; in https://sentry.intersight.com . | [optional] [readonly] 
**Method** | Pointer to **string** | The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the privilege reported by microservice. | [optional] [readonly] 
**RestPath** | Pointer to **string** | The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions. | [optional] [readonly] 
**UrlPrefix** | Pointer to **string** | The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamPrivilege

`func NewIamPrivilege() *IamPrivilege`

NewIamPrivilege instantiates a new IamPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivilegeWithDefaults

`func NewIamPrivilegeWithDefaults() *IamPrivilege`

NewIamPrivilegeWithDefaults instantiates a new IamPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnamePrefix

`func (o *IamPrivilege) GetHostnamePrefix() string`

GetHostnamePrefix returns the HostnamePrefix field if non-nil, zero value otherwise.

### GetHostnamePrefixOk

`func (o *IamPrivilege) GetHostnamePrefixOk() (*string, bool)`

GetHostnamePrefixOk returns a tuple with the HostnamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamePrefix

`func (o *IamPrivilege) SetHostnamePrefix(v string)`

SetHostnamePrefix sets HostnamePrefix field to given value.

### HasHostnamePrefix

`func (o *IamPrivilege) HasHostnamePrefix() bool`

HasHostnamePrefix returns a boolean if a field has been set.

### GetMethod

`func (o *IamPrivilege) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IamPrivilege) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IamPrivilege) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IamPrivilege) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *IamPrivilege) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPrivilege) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPrivilege) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPrivilege) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestPath

`func (o *IamPrivilege) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *IamPrivilege) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *IamPrivilege) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *IamPrivilege) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.

### GetUrlPrefix

`func (o *IamPrivilege) GetUrlPrefix() string`

GetUrlPrefix returns the UrlPrefix field if non-nil, zero value otherwise.

### GetUrlPrefixOk

`func (o *IamPrivilege) GetUrlPrefixOk() (*string, bool)`

GetUrlPrefixOk returns a tuple with the UrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrefix

`func (o *IamPrivilege) SetUrlPrefix(v string)`

SetUrlPrefix sets UrlPrefix field to given value.

### HasUrlPrefix

`func (o *IamPrivilege) HasUrlPrefix() bool`

HasUrlPrefix returns a boolean if a field has been set.

### GetAccount

`func (o *IamPrivilege) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamPrivilege) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamPrivilege) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamPrivilege) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSystem

`func (o *IamPrivilege) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamPrivilege) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamPrivilege) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamPrivilege) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


