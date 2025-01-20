# ResourceChassisAndSlotQualification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.ChassisAndSlotQualification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.ChassisAndSlotQualification"]
**ChassisIdRange** | Pointer to [**ResourceChassisIdRangeFilter**](ResourceChassisIdRangeFilter.md) |  | [optional] 
**SlotIdRanges** | Pointer to [**[]ResourceSlotIdRangeFilter**](ResourceSlotIdRangeFilter.md) |  | [optional] 

## Methods

### NewResourceChassisAndSlotQualification

`func NewResourceChassisAndSlotQualification(classId string, objectType string, ) *ResourceChassisAndSlotQualification`

NewResourceChassisAndSlotQualification instantiates a new ResourceChassisAndSlotQualification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceChassisAndSlotQualificationWithDefaults

`func NewResourceChassisAndSlotQualificationWithDefaults() *ResourceChassisAndSlotQualification`

NewResourceChassisAndSlotQualificationWithDefaults instantiates a new ResourceChassisAndSlotQualification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceChassisAndSlotQualification) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceChassisAndSlotQualification) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceChassisAndSlotQualification) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceChassisAndSlotQualification) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceChassisAndSlotQualification) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceChassisAndSlotQualification) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisIdRange

`func (o *ResourceChassisAndSlotQualification) GetChassisIdRange() ResourceChassisIdRangeFilter`

GetChassisIdRange returns the ChassisIdRange field if non-nil, zero value otherwise.

### GetChassisIdRangeOk

`func (o *ResourceChassisAndSlotQualification) GetChassisIdRangeOk() (*ResourceChassisIdRangeFilter, bool)`

GetChassisIdRangeOk returns a tuple with the ChassisIdRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIdRange

`func (o *ResourceChassisAndSlotQualification) SetChassisIdRange(v ResourceChassisIdRangeFilter)`

SetChassisIdRange sets ChassisIdRange field to given value.

### HasChassisIdRange

`func (o *ResourceChassisAndSlotQualification) HasChassisIdRange() bool`

HasChassisIdRange returns a boolean if a field has been set.

### GetSlotIdRanges

`func (o *ResourceChassisAndSlotQualification) GetSlotIdRanges() []ResourceSlotIdRangeFilter`

GetSlotIdRanges returns the SlotIdRanges field if non-nil, zero value otherwise.

### GetSlotIdRangesOk

`func (o *ResourceChassisAndSlotQualification) GetSlotIdRangesOk() (*[]ResourceSlotIdRangeFilter, bool)`

GetSlotIdRangesOk returns a tuple with the SlotIdRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotIdRanges

`func (o *ResourceChassisAndSlotQualification) SetSlotIdRanges(v []ResourceSlotIdRangeFilter)`

SetSlotIdRanges sets SlotIdRanges field to given value.

### HasSlotIdRanges

`func (o *ResourceChassisAndSlotQualification) HasSlotIdRanges() bool`

HasSlotIdRanges returns a boolean if a field has been set.

### SetSlotIdRangesNil

`func (o *ResourceChassisAndSlotQualification) SetSlotIdRangesNil(b bool)`

 SetSlotIdRangesNil sets the value for SlotIdRanges to be an explicit nil

### UnsetSlotIdRanges
`func (o *ResourceChassisAndSlotQualification) UnsetSlotIdRanges()`

UnsetSlotIdRanges ensures that no value is present for SlotIdRanges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


