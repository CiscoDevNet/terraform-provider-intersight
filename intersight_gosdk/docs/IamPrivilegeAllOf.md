# IamPrivilegeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Privilege"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Privilege"]
**HostnamePrefix** | Pointer to **string** | The hostname prefix of the resource corresponding to this privilege. For example \\&#39;sentry\\&#39; in https://sentry.intersight.com . | [optional] [readonly] 
**Method** | Pointer to **string** | The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the privilege reported by microservice. | [optional] [readonly] 
**RestPath** | Pointer to **string** | The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions. | [optional] [readonly] 
**UrlPrefix** | Pointer to **string** | The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**System** | Pointer to [**IamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewIamPrivilegeAllOf

`func NewIamPrivilegeAllOf(classId string, objectType string, ) *IamPrivilegeAllOf`

NewIamPrivilegeAllOf instantiates a new IamPrivilegeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivilegeAllOfWithDefaults

`func NewIamPrivilegeAllOfWithDefaults() *IamPrivilegeAllOf`

NewIamPrivilegeAllOfWithDefaults instantiates a new IamPrivilegeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPrivilegeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPrivilegeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPrivilegeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPrivilegeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPrivilegeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPrivilegeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostnamePrefix

`func (o *IamPrivilegeAllOf) GetHostnamePrefix() string`

GetHostnamePrefix returns the HostnamePrefix field if non-nil, zero value otherwise.

### GetHostnamePrefixOk

`func (o *IamPrivilegeAllOf) GetHostnamePrefixOk() (*string, bool)`

GetHostnamePrefixOk returns a tuple with the HostnamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamePrefix

`func (o *IamPrivilegeAllOf) SetHostnamePrefix(v string)`

SetHostnamePrefix sets HostnamePrefix field to given value.

### HasHostnamePrefix

`func (o *IamPrivilegeAllOf) HasHostnamePrefix() bool`

HasHostnamePrefix returns a boolean if a field has been set.

### GetMethod

`func (o *IamPrivilegeAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IamPrivilegeAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IamPrivilegeAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IamPrivilegeAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *IamPrivilegeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPrivilegeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPrivilegeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPrivilegeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestPath

`func (o *IamPrivilegeAllOf) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *IamPrivilegeAllOf) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *IamPrivilegeAllOf) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *IamPrivilegeAllOf) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.

### GetUrlPrefix

`func (o *IamPrivilegeAllOf) GetUrlPrefix() string`

GetUrlPrefix returns the UrlPrefix field if non-nil, zero value otherwise.

### GetUrlPrefixOk

`func (o *IamPrivilegeAllOf) GetUrlPrefixOk() (*string, bool)`

GetUrlPrefixOk returns a tuple with the UrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrefix

`func (o *IamPrivilegeAllOf) SetUrlPrefix(v string)`

SetUrlPrefix sets UrlPrefix field to given value.

### HasUrlPrefix

`func (o *IamPrivilegeAllOf) HasUrlPrefix() bool`

HasUrlPrefix returns a boolean if a field has been set.

### GetAccount

`func (o *IamPrivilegeAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamPrivilegeAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamPrivilegeAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamPrivilegeAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSystem

`func (o *IamPrivilegeAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamPrivilegeAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamPrivilegeAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamPrivilegeAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


