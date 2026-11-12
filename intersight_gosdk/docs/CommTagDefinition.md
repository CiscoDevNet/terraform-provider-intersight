# CommTagDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "comm.TagDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "comm.TagDefinition"]
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**EnablePropagation** | Pointer to **bool** | If this flag is enabled, the tag will be propagated to related managed objects. Propagation is supported in a limited manner for path tags and it is not supported for key value. Rules for propagation are configured by Intersight and cannot be configured by user. | [optional] 
**Key** | Pointer to **string** | The string representation of the tag key. For PathTag definitions, \&quot;/\&quot; is interpreted as a path delimiter. The combination of key and tag type must be unique within an account. The key is case sensitive and cannot be empty. | [optional] 
**RestrictAllowedValues** | Pointer to **bool** | If this flag is enabled, then values of the KeyValue tag can only be values present in the allowedValues list. RestrictAllowedValues is not applicable to path tags. | [optional] 
**SysTag** | Pointer to **bool** | Specifies whether the tag is user-defined or owned by the system. | [optional] [readonly] 
**Type** | Pointer to **string** | An enum that defines the type of tag definition. KeyValue tags use a key with an optional value. PathTag tags use a hierarchical key with &#39;/&#39; as the delimiter and no value. Defaults to KeyValue. * &#x60;KeyValue&#x60; - KeyValue type of tag. Key is required for these tags. Value is optional. * &#x60;PathTag&#x60; - Key contain path information. Value is not present for these tags. The path is created by using the &#39;/&#39; character as a delimiter.For example, if the tag is \&quot;A/B/C\&quot;, then \&quot;A\&quot; is the parent tag, \&quot;B\&quot; is the child tag of \&quot;A\&quot; and \&quot;C\&quot; is the child tag of \&quot;B\&quot;. | [optional] [default to "KeyValue"]
**Usage** | Pointer to [**CommTagUsage**](CommTagUsage.md) |  | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**ParentTag** | Pointer to [**NullableCommTagDefinitionRelationship**](CommTagDefinitionRelationship.md) |  | [optional] 

## Methods

### NewCommTagDefinition

`func NewCommTagDefinition(classId string, objectType string, ) *CommTagDefinition`

NewCommTagDefinition instantiates a new CommTagDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommTagDefinitionWithDefaults

`func NewCommTagDefinitionWithDefaults() *CommTagDefinition`

NewCommTagDefinitionWithDefaults instantiates a new CommTagDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommTagDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommTagDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommTagDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommTagDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommTagDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommTagDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedValues

`func (o *CommTagDefinition) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *CommTagDefinition) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *CommTagDefinition) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *CommTagDefinition) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *CommTagDefinition) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *CommTagDefinition) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetEnablePropagation

`func (o *CommTagDefinition) GetEnablePropagation() bool`

GetEnablePropagation returns the EnablePropagation field if non-nil, zero value otherwise.

### GetEnablePropagationOk

`func (o *CommTagDefinition) GetEnablePropagationOk() (*bool, bool)`

GetEnablePropagationOk returns a tuple with the EnablePropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePropagation

`func (o *CommTagDefinition) SetEnablePropagation(v bool)`

SetEnablePropagation sets EnablePropagation field to given value.

### HasEnablePropagation

`func (o *CommTagDefinition) HasEnablePropagation() bool`

HasEnablePropagation returns a boolean if a field has been set.

### GetKey

`func (o *CommTagDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CommTagDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CommTagDefinition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CommTagDefinition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRestrictAllowedValues

`func (o *CommTagDefinition) GetRestrictAllowedValues() bool`

GetRestrictAllowedValues returns the RestrictAllowedValues field if non-nil, zero value otherwise.

### GetRestrictAllowedValuesOk

`func (o *CommTagDefinition) GetRestrictAllowedValuesOk() (*bool, bool)`

GetRestrictAllowedValuesOk returns a tuple with the RestrictAllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAllowedValues

`func (o *CommTagDefinition) SetRestrictAllowedValues(v bool)`

SetRestrictAllowedValues sets RestrictAllowedValues field to given value.

### HasRestrictAllowedValues

`func (o *CommTagDefinition) HasRestrictAllowedValues() bool`

HasRestrictAllowedValues returns a boolean if a field has been set.

### GetSysTag

`func (o *CommTagDefinition) GetSysTag() bool`

GetSysTag returns the SysTag field if non-nil, zero value otherwise.

### GetSysTagOk

`func (o *CommTagDefinition) GetSysTagOk() (*bool, bool)`

GetSysTagOk returns a tuple with the SysTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysTag

`func (o *CommTagDefinition) SetSysTag(v bool)`

SetSysTag sets SysTag field to given value.

### HasSysTag

`func (o *CommTagDefinition) HasSysTag() bool`

HasSysTag returns a boolean if a field has been set.

### GetType

`func (o *CommTagDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommTagDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommTagDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommTagDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *CommTagDefinition) GetUsage() CommTagUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CommTagDefinition) GetUsageOk() (*CommTagUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CommTagDefinition) SetUsage(v CommTagUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CommTagDefinition) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAccount

`func (o *CommTagDefinition) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CommTagDefinition) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CommTagDefinition) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CommTagDefinition) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CommTagDefinition) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CommTagDefinition) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetParentTag

`func (o *CommTagDefinition) GetParentTag() CommTagDefinitionRelationship`

GetParentTag returns the ParentTag field if non-nil, zero value otherwise.

### GetParentTagOk

`func (o *CommTagDefinition) GetParentTagOk() (*CommTagDefinitionRelationship, bool)`

GetParentTagOk returns a tuple with the ParentTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTag

`func (o *CommTagDefinition) SetParentTag(v CommTagDefinitionRelationship)`

SetParentTag sets ParentTag field to given value.

### HasParentTag

`func (o *CommTagDefinition) HasParentTag() bool`

HasParentTag returns a boolean if a field has been set.

### SetParentTagNil

`func (o *CommTagDefinition) SetParentTagNil(b bool)`

 SetParentTagNil sets the value for ParentTag to be an explicit nil

### UnsetParentTag
`func (o *CommTagDefinition) UnsetParentTag()`

UnsetParentTag ensures that no value is present for ParentTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


