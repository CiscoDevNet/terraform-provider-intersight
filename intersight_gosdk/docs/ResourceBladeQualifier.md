# ResourceBladeQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.BladeQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.BladeQualifier"]
**ChassisAndSlotIdRange** | Pointer to [**[]ResourceChassisAndSlotQualification**](ResourceChassisAndSlotQualification.md) |  | [optional] 
**ChassisPids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResourceBladeQualifier

`func NewResourceBladeQualifier(classId string, objectType string, ) *ResourceBladeQualifier`

NewResourceBladeQualifier instantiates a new ResourceBladeQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceBladeQualifierWithDefaults

`func NewResourceBladeQualifierWithDefaults() *ResourceBladeQualifier`

NewResourceBladeQualifierWithDefaults instantiates a new ResourceBladeQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceBladeQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceBladeQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceBladeQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceBladeQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceBladeQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceBladeQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisAndSlotIdRange

`func (o *ResourceBladeQualifier) GetChassisAndSlotIdRange() []ResourceChassisAndSlotQualification`

GetChassisAndSlotIdRange returns the ChassisAndSlotIdRange field if non-nil, zero value otherwise.

### GetChassisAndSlotIdRangeOk

`func (o *ResourceBladeQualifier) GetChassisAndSlotIdRangeOk() (*[]ResourceChassisAndSlotQualification, bool)`

GetChassisAndSlotIdRangeOk returns a tuple with the ChassisAndSlotIdRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisAndSlotIdRange

`func (o *ResourceBladeQualifier) SetChassisAndSlotIdRange(v []ResourceChassisAndSlotQualification)`

SetChassisAndSlotIdRange sets ChassisAndSlotIdRange field to given value.

### HasChassisAndSlotIdRange

`func (o *ResourceBladeQualifier) HasChassisAndSlotIdRange() bool`

HasChassisAndSlotIdRange returns a boolean if a field has been set.

### SetChassisAndSlotIdRangeNil

`func (o *ResourceBladeQualifier) SetChassisAndSlotIdRangeNil(b bool)`

 SetChassisAndSlotIdRangeNil sets the value for ChassisAndSlotIdRange to be an explicit nil

### UnsetChassisAndSlotIdRange
`func (o *ResourceBladeQualifier) UnsetChassisAndSlotIdRange()`

UnsetChassisAndSlotIdRange ensures that no value is present for ChassisAndSlotIdRange, not even an explicit nil
### GetChassisPids

`func (o *ResourceBladeQualifier) GetChassisPids() []string`

GetChassisPids returns the ChassisPids field if non-nil, zero value otherwise.

### GetChassisPidsOk

`func (o *ResourceBladeQualifier) GetChassisPidsOk() (*[]string, bool)`

GetChassisPidsOk returns a tuple with the ChassisPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisPids

`func (o *ResourceBladeQualifier) SetChassisPids(v []string)`

SetChassisPids sets ChassisPids field to given value.

### HasChassisPids

`func (o *ResourceBladeQualifier) HasChassisPids() bool`

HasChassisPids returns a boolean if a field has been set.

### SetChassisPidsNil

`func (o *ResourceBladeQualifier) SetChassisPidsNil(b bool)`

 SetChassisPidsNil sets the value for ChassisPids to be an explicit nil

### UnsetChassisPids
`func (o *ResourceBladeQualifier) UnsetChassisPids()`

UnsetChassisPids ensures that no value is present for ChassisPids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


