# WorkloadDefinitionMapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.DefinitionMapper"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.DefinitionMapper"]
**DefinitionName** | Pointer to **string** | The name of the definition being referenced. | [optional] 
**Organization** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**UsePreferredVersion** | Pointer to **bool** | Indicates whether this version is the default version of the referenced definition. If set to true, the version should be the default. | [optional] 
**Version** | Pointer to **int64** | The version number of the referenced definition. | [optional] 

## Methods

### NewWorkloadDefinitionMapper

`func NewWorkloadDefinitionMapper(classId string, objectType string, ) *WorkloadDefinitionMapper`

NewWorkloadDefinitionMapper instantiates a new WorkloadDefinitionMapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDefinitionMapperWithDefaults

`func NewWorkloadDefinitionMapperWithDefaults() *WorkloadDefinitionMapper`

NewWorkloadDefinitionMapperWithDefaults instantiates a new WorkloadDefinitionMapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadDefinitionMapper) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadDefinitionMapper) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadDefinitionMapper) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadDefinitionMapper) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadDefinitionMapper) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadDefinitionMapper) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefinitionName

`func (o *WorkloadDefinitionMapper) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *WorkloadDefinitionMapper) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *WorkloadDefinitionMapper) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.

### HasDefinitionName

`func (o *WorkloadDefinitionMapper) HasDefinitionName() bool`

HasDefinitionName returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkloadDefinitionMapper) GetOrganization() MoMoRef`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkloadDefinitionMapper) GetOrganizationOk() (*MoMoRef, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkloadDefinitionMapper) SetOrganization(v MoMoRef)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkloadDefinitionMapper) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetUsePreferredVersion

`func (o *WorkloadDefinitionMapper) GetUsePreferredVersion() bool`

GetUsePreferredVersion returns the UsePreferredVersion field if non-nil, zero value otherwise.

### GetUsePreferredVersionOk

`func (o *WorkloadDefinitionMapper) GetUsePreferredVersionOk() (*bool, bool)`

GetUsePreferredVersionOk returns a tuple with the UsePreferredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredVersion

`func (o *WorkloadDefinitionMapper) SetUsePreferredVersion(v bool)`

SetUsePreferredVersion sets UsePreferredVersion field to given value.

### HasUsePreferredVersion

`func (o *WorkloadDefinitionMapper) HasUsePreferredVersion() bool`

HasUsePreferredVersion returns a boolean if a field has been set.

### GetVersion

`func (o *WorkloadDefinitionMapper) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkloadDefinitionMapper) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkloadDefinitionMapper) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkloadDefinitionMapper) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


