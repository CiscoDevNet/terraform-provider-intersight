# ResourceMemoryQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.MemoryQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.MemoryQualifier"]
**MemoryCapacityRange** | Pointer to [**ResourceMemoryCapacityRangeFilter**](ResourceMemoryCapacityRangeFilter.md) |  | [optional] 
**MemoryUnitsRange** | Pointer to [**ResourceMemoryUnitsRangeFilter**](ResourceMemoryUnitsRangeFilter.md) |  | [optional] 

## Methods

### NewResourceMemoryQualifier

`func NewResourceMemoryQualifier(classId string, objectType string, ) *ResourceMemoryQualifier`

NewResourceMemoryQualifier instantiates a new ResourceMemoryQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMemoryQualifierWithDefaults

`func NewResourceMemoryQualifierWithDefaults() *ResourceMemoryQualifier`

NewResourceMemoryQualifierWithDefaults instantiates a new ResourceMemoryQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceMemoryQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceMemoryQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceMemoryQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceMemoryQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceMemoryQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceMemoryQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMemoryCapacityRange

`func (o *ResourceMemoryQualifier) GetMemoryCapacityRange() ResourceMemoryCapacityRangeFilter`

GetMemoryCapacityRange returns the MemoryCapacityRange field if non-nil, zero value otherwise.

### GetMemoryCapacityRangeOk

`func (o *ResourceMemoryQualifier) GetMemoryCapacityRangeOk() (*ResourceMemoryCapacityRangeFilter, bool)`

GetMemoryCapacityRangeOk returns a tuple with the MemoryCapacityRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacityRange

`func (o *ResourceMemoryQualifier) SetMemoryCapacityRange(v ResourceMemoryCapacityRangeFilter)`

SetMemoryCapacityRange sets MemoryCapacityRange field to given value.

### HasMemoryCapacityRange

`func (o *ResourceMemoryQualifier) HasMemoryCapacityRange() bool`

HasMemoryCapacityRange returns a boolean if a field has been set.

### GetMemoryUnitsRange

`func (o *ResourceMemoryQualifier) GetMemoryUnitsRange() ResourceMemoryUnitsRangeFilter`

GetMemoryUnitsRange returns the MemoryUnitsRange field if non-nil, zero value otherwise.

### GetMemoryUnitsRangeOk

`func (o *ResourceMemoryQualifier) GetMemoryUnitsRangeOk() (*ResourceMemoryUnitsRangeFilter, bool)`

GetMemoryUnitsRangeOk returns a tuple with the MemoryUnitsRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUnitsRange

`func (o *ResourceMemoryQualifier) SetMemoryUnitsRange(v ResourceMemoryUnitsRangeFilter)`

SetMemoryUnitsRange sets MemoryUnitsRange field to given value.

### HasMemoryUnitsRange

`func (o *ResourceMemoryQualifier) HasMemoryUnitsRange() bool`

HasMemoryUnitsRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


