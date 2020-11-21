# MetaAccessPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meta.AccessPrivilege"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meta.AccessPrivilege"]
**Method** | Pointer to **string** | The type of CRUD operation (create, read, update, delete) for which an access privilege is required. * &#x60;Update&#x60; - The &#39;update&#39; operation/state. * &#x60;Create&#x60; - The &#39;create&#39; operation/state. * &#x60;Read&#x60; - The &#39;read&#39; operation/state. * &#x60;Delete&#x60; - The &#39;delete&#39; operation/state. | [optional] [readonly] [default to "Update"]
**Privilege** | Pointer to **string** | The name of the privilege which is required to invoke the specified CRUD method. | [optional] [readonly] 

## Methods

### NewMetaAccessPrivilege

`func NewMetaAccessPrivilege(classId string, objectType string, ) *MetaAccessPrivilege`

NewMetaAccessPrivilege instantiates a new MetaAccessPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAccessPrivilegeWithDefaults

`func NewMetaAccessPrivilegeWithDefaults() *MetaAccessPrivilege`

NewMetaAccessPrivilegeWithDefaults instantiates a new MetaAccessPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetaAccessPrivilege) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetaAccessPrivilege) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetaAccessPrivilege) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetaAccessPrivilege) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetaAccessPrivilege) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetaAccessPrivilege) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMethod

`func (o *MetaAccessPrivilege) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MetaAccessPrivilege) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MetaAccessPrivilege) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MetaAccessPrivilege) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPrivilege

`func (o *MetaAccessPrivilege) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *MetaAccessPrivilege) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *MetaAccessPrivilege) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *MetaAccessPrivilege) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


