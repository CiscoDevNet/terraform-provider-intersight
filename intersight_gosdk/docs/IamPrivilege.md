# IamPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Privilege"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Privilege"]
**AttributeNames** | Pointer to **[]string** |  | [optional] 
**Condition** | Pointer to **string** | The condition expression or criteria that determines when this conditional privilege can grant access to update the specified attribute names. | [optional] [readonly] 
**Description** | Pointer to **string** | Informative description about each privilege. | [optional] [readonly] 
**HostnamePrefix** | Pointer to **string** | The hostname prefix of the resource corresponding to this privilege. For example \\&#39;sentry\\&#39; in https://sentry.intersight.com . | [optional] [readonly] 
**IsConditionalPrivilege** | Pointer to **bool** | Flag that indicates whether this privilege is conditional. | [optional] [readonly] 
**Method** | Pointer to **string** | The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the privilege reported by microservice. | [optional] [readonly] 
**RestPath** | Pointer to **string** | The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions. | [optional] [readonly] 
**UrlPrefix** | Pointer to **string** | The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**System** | Pointer to [**NullableIamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewIamPrivilege

`func NewIamPrivilege(classId string, objectType string, ) *IamPrivilege`

NewIamPrivilege instantiates a new IamPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivilegeWithDefaults

`func NewIamPrivilegeWithDefaults() *IamPrivilege`

NewIamPrivilegeWithDefaults instantiates a new IamPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPrivilege) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPrivilege) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPrivilege) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPrivilege) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPrivilege) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPrivilege) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributeNames

`func (o *IamPrivilege) GetAttributeNames() []string`

GetAttributeNames returns the AttributeNames field if non-nil, zero value otherwise.

### GetAttributeNamesOk

`func (o *IamPrivilege) GetAttributeNamesOk() (*[]string, bool)`

GetAttributeNamesOk returns a tuple with the AttributeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeNames

`func (o *IamPrivilege) SetAttributeNames(v []string)`

SetAttributeNames sets AttributeNames field to given value.

### HasAttributeNames

`func (o *IamPrivilege) HasAttributeNames() bool`

HasAttributeNames returns a boolean if a field has been set.

### SetAttributeNamesNil

`func (o *IamPrivilege) SetAttributeNamesNil(b bool)`

 SetAttributeNamesNil sets the value for AttributeNames to be an explicit nil

### UnsetAttributeNames
`func (o *IamPrivilege) UnsetAttributeNames()`

UnsetAttributeNames ensures that no value is present for AttributeNames, not even an explicit nil
### GetCondition

`func (o *IamPrivilege) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *IamPrivilege) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *IamPrivilege) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *IamPrivilege) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDescription

`func (o *IamPrivilege) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamPrivilege) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamPrivilege) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamPrivilege) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### GetIsConditionalPrivilege

`func (o *IamPrivilege) GetIsConditionalPrivilege() bool`

GetIsConditionalPrivilege returns the IsConditionalPrivilege field if non-nil, zero value otherwise.

### GetIsConditionalPrivilegeOk

`func (o *IamPrivilege) GetIsConditionalPrivilegeOk() (*bool, bool)`

GetIsConditionalPrivilegeOk returns a tuple with the IsConditionalPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConditionalPrivilege

`func (o *IamPrivilege) SetIsConditionalPrivilege(v bool)`

SetIsConditionalPrivilege sets IsConditionalPrivilege field to given value.

### HasIsConditionalPrivilege

`func (o *IamPrivilege) HasIsConditionalPrivilege() bool`

HasIsConditionalPrivilege returns a boolean if a field has been set.

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

### SetAccountNil

`func (o *IamPrivilege) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamPrivilege) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
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

### SetSystemNil

`func (o *IamPrivilege) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *IamPrivilege) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


