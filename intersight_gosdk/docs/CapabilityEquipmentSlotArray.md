# CapabilityEquipmentSlotArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.EquipmentSlotArray"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.EquipmentSlotArray"]
**FirstIndex** | Pointer to **float32** | First Index information for a Switch/Fabric-Interconnect hardware. | [optional] 
**Height** | Pointer to **float32** | Height information for a Switch/Fabric-Interconnect hardware. | [optional] 
**HorizontalStartOffset** | Pointer to **float32** | Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware. | [optional] 
**InlineGroupSeparation** | Pointer to **float32** | Inline Group Separation information for a Switch/Fabric-Interconnect hardware. | [optional] 
**InlineGroupSize** | Pointer to **float32** | Inline Group Size information for a Switch/Fabric-Interconnect hardware. | [optional] 
**InlineOffset** | Pointer to **float32** | Inline Offset information for a Switch/Fabric-Interconnect hardware. | [optional] 
**Location** | Pointer to **string** | Location information for a Switch/Fabric-Interconnect hardware. | [optional] 
**NumberOfSlots** | Pointer to **int64** | Number of Slots information for a Switch/Fabric-Interconnect hardware. | [optional] 
**Orientation** | Pointer to **string** | Orientation information for a Switch/Fabric-Interconnect hardware. | [optional] 
**Selector** | Pointer to **string** | Selector information for a Switch/Fabric-Interconnect hardware. | [optional] 
**SlotsPerLine** | Pointer to **int64** | Slots per line information for a Switch/Fabric-Interconnect hardware. | [optional] 
**TransverseGroupSeparation** | Pointer to **float32** | Transverse Group Separation information for a Switch/Fabric-Interconnect hardware. | [optional] 
**TransverseGroupSize** | Pointer to **float32** | Transverse Group Size information for a Switch/Fabric-Interconnect hardware. | [optional] 
**TransverseOffset** | Pointer to **float32** | Transverse Offset information for a Switch/Fabric-Interconnect hardware. | [optional] 
**VerticalStartOffset** | Pointer to **float32** | Vertical Start Offset information for a Switch/Fabric-Interconnect hardware. | [optional] 
**Width** | Pointer to **float32** | Width information for a Switch/Fabric-Interconnect hardware. | [optional] 

## Methods

### NewCapabilityEquipmentSlotArray

`func NewCapabilityEquipmentSlotArray(classId string, objectType string, ) *CapabilityEquipmentSlotArray`

NewCapabilityEquipmentSlotArray instantiates a new CapabilityEquipmentSlotArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityEquipmentSlotArrayWithDefaults

`func NewCapabilityEquipmentSlotArrayWithDefaults() *CapabilityEquipmentSlotArray`

NewCapabilityEquipmentSlotArrayWithDefaults instantiates a new CapabilityEquipmentSlotArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityEquipmentSlotArray) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityEquipmentSlotArray) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityEquipmentSlotArray) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityEquipmentSlotArray) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityEquipmentSlotArray) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityEquipmentSlotArray) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFirstIndex

`func (o *CapabilityEquipmentSlotArray) GetFirstIndex() float32`

GetFirstIndex returns the FirstIndex field if non-nil, zero value otherwise.

### GetFirstIndexOk

`func (o *CapabilityEquipmentSlotArray) GetFirstIndexOk() (*float32, bool)`

GetFirstIndexOk returns a tuple with the FirstIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstIndex

`func (o *CapabilityEquipmentSlotArray) SetFirstIndex(v float32)`

SetFirstIndex sets FirstIndex field to given value.

### HasFirstIndex

`func (o *CapabilityEquipmentSlotArray) HasFirstIndex() bool`

HasFirstIndex returns a boolean if a field has been set.

### GetHeight

`func (o *CapabilityEquipmentSlotArray) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CapabilityEquipmentSlotArray) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CapabilityEquipmentSlotArray) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *CapabilityEquipmentSlotArray) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetHorizontalStartOffset

`func (o *CapabilityEquipmentSlotArray) GetHorizontalStartOffset() float32`

GetHorizontalStartOffset returns the HorizontalStartOffset field if non-nil, zero value otherwise.

### GetHorizontalStartOffsetOk

`func (o *CapabilityEquipmentSlotArray) GetHorizontalStartOffsetOk() (*float32, bool)`

GetHorizontalStartOffsetOk returns a tuple with the HorizontalStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalStartOffset

`func (o *CapabilityEquipmentSlotArray) SetHorizontalStartOffset(v float32)`

SetHorizontalStartOffset sets HorizontalStartOffset field to given value.

### HasHorizontalStartOffset

`func (o *CapabilityEquipmentSlotArray) HasHorizontalStartOffset() bool`

HasHorizontalStartOffset returns a boolean if a field has been set.

### GetInlineGroupSeparation

`func (o *CapabilityEquipmentSlotArray) GetInlineGroupSeparation() float32`

GetInlineGroupSeparation returns the InlineGroupSeparation field if non-nil, zero value otherwise.

### GetInlineGroupSeparationOk

`func (o *CapabilityEquipmentSlotArray) GetInlineGroupSeparationOk() (*float32, bool)`

GetInlineGroupSeparationOk returns a tuple with the InlineGroupSeparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineGroupSeparation

`func (o *CapabilityEquipmentSlotArray) SetInlineGroupSeparation(v float32)`

SetInlineGroupSeparation sets InlineGroupSeparation field to given value.

### HasInlineGroupSeparation

`func (o *CapabilityEquipmentSlotArray) HasInlineGroupSeparation() bool`

HasInlineGroupSeparation returns a boolean if a field has been set.

### GetInlineGroupSize

`func (o *CapabilityEquipmentSlotArray) GetInlineGroupSize() float32`

GetInlineGroupSize returns the InlineGroupSize field if non-nil, zero value otherwise.

### GetInlineGroupSizeOk

`func (o *CapabilityEquipmentSlotArray) GetInlineGroupSizeOk() (*float32, bool)`

GetInlineGroupSizeOk returns a tuple with the InlineGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineGroupSize

`func (o *CapabilityEquipmentSlotArray) SetInlineGroupSize(v float32)`

SetInlineGroupSize sets InlineGroupSize field to given value.

### HasInlineGroupSize

`func (o *CapabilityEquipmentSlotArray) HasInlineGroupSize() bool`

HasInlineGroupSize returns a boolean if a field has been set.

### GetInlineOffset

`func (o *CapabilityEquipmentSlotArray) GetInlineOffset() float32`

GetInlineOffset returns the InlineOffset field if non-nil, zero value otherwise.

### GetInlineOffsetOk

`func (o *CapabilityEquipmentSlotArray) GetInlineOffsetOk() (*float32, bool)`

GetInlineOffsetOk returns a tuple with the InlineOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineOffset

`func (o *CapabilityEquipmentSlotArray) SetInlineOffset(v float32)`

SetInlineOffset sets InlineOffset field to given value.

### HasInlineOffset

`func (o *CapabilityEquipmentSlotArray) HasInlineOffset() bool`

HasInlineOffset returns a boolean if a field has been set.

### GetLocation

`func (o *CapabilityEquipmentSlotArray) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CapabilityEquipmentSlotArray) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CapabilityEquipmentSlotArray) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CapabilityEquipmentSlotArray) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNumberOfSlots

`func (o *CapabilityEquipmentSlotArray) GetNumberOfSlots() int64`

GetNumberOfSlots returns the NumberOfSlots field if non-nil, zero value otherwise.

### GetNumberOfSlotsOk

`func (o *CapabilityEquipmentSlotArray) GetNumberOfSlotsOk() (*int64, bool)`

GetNumberOfSlotsOk returns a tuple with the NumberOfSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSlots

`func (o *CapabilityEquipmentSlotArray) SetNumberOfSlots(v int64)`

SetNumberOfSlots sets NumberOfSlots field to given value.

### HasNumberOfSlots

`func (o *CapabilityEquipmentSlotArray) HasNumberOfSlots() bool`

HasNumberOfSlots returns a boolean if a field has been set.

### GetOrientation

`func (o *CapabilityEquipmentSlotArray) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CapabilityEquipmentSlotArray) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CapabilityEquipmentSlotArray) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CapabilityEquipmentSlotArray) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetSelector

`func (o *CapabilityEquipmentSlotArray) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CapabilityEquipmentSlotArray) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CapabilityEquipmentSlotArray) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *CapabilityEquipmentSlotArray) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSlotsPerLine

`func (o *CapabilityEquipmentSlotArray) GetSlotsPerLine() int64`

GetSlotsPerLine returns the SlotsPerLine field if non-nil, zero value otherwise.

### GetSlotsPerLineOk

`func (o *CapabilityEquipmentSlotArray) GetSlotsPerLineOk() (*int64, bool)`

GetSlotsPerLineOk returns a tuple with the SlotsPerLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotsPerLine

`func (o *CapabilityEquipmentSlotArray) SetSlotsPerLine(v int64)`

SetSlotsPerLine sets SlotsPerLine field to given value.

### HasSlotsPerLine

`func (o *CapabilityEquipmentSlotArray) HasSlotsPerLine() bool`

HasSlotsPerLine returns a boolean if a field has been set.

### GetTransverseGroupSeparation

`func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSeparation() float32`

GetTransverseGroupSeparation returns the TransverseGroupSeparation field if non-nil, zero value otherwise.

### GetTransverseGroupSeparationOk

`func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSeparationOk() (*float32, bool)`

GetTransverseGroupSeparationOk returns a tuple with the TransverseGroupSeparation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransverseGroupSeparation

`func (o *CapabilityEquipmentSlotArray) SetTransverseGroupSeparation(v float32)`

SetTransverseGroupSeparation sets TransverseGroupSeparation field to given value.

### HasTransverseGroupSeparation

`func (o *CapabilityEquipmentSlotArray) HasTransverseGroupSeparation() bool`

HasTransverseGroupSeparation returns a boolean if a field has been set.

### GetTransverseGroupSize

`func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSize() float32`

GetTransverseGroupSize returns the TransverseGroupSize field if non-nil, zero value otherwise.

### GetTransverseGroupSizeOk

`func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSizeOk() (*float32, bool)`

GetTransverseGroupSizeOk returns a tuple with the TransverseGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransverseGroupSize

`func (o *CapabilityEquipmentSlotArray) SetTransverseGroupSize(v float32)`

SetTransverseGroupSize sets TransverseGroupSize field to given value.

### HasTransverseGroupSize

`func (o *CapabilityEquipmentSlotArray) HasTransverseGroupSize() bool`

HasTransverseGroupSize returns a boolean if a field has been set.

### GetTransverseOffset

`func (o *CapabilityEquipmentSlotArray) GetTransverseOffset() float32`

GetTransverseOffset returns the TransverseOffset field if non-nil, zero value otherwise.

### GetTransverseOffsetOk

`func (o *CapabilityEquipmentSlotArray) GetTransverseOffsetOk() (*float32, bool)`

GetTransverseOffsetOk returns a tuple with the TransverseOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransverseOffset

`func (o *CapabilityEquipmentSlotArray) SetTransverseOffset(v float32)`

SetTransverseOffset sets TransverseOffset field to given value.

### HasTransverseOffset

`func (o *CapabilityEquipmentSlotArray) HasTransverseOffset() bool`

HasTransverseOffset returns a boolean if a field has been set.

### GetVerticalStartOffset

`func (o *CapabilityEquipmentSlotArray) GetVerticalStartOffset() float32`

GetVerticalStartOffset returns the VerticalStartOffset field if non-nil, zero value otherwise.

### GetVerticalStartOffsetOk

`func (o *CapabilityEquipmentSlotArray) GetVerticalStartOffsetOk() (*float32, bool)`

GetVerticalStartOffsetOk returns a tuple with the VerticalStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalStartOffset

`func (o *CapabilityEquipmentSlotArray) SetVerticalStartOffset(v float32)`

SetVerticalStartOffset sets VerticalStartOffset field to given value.

### HasVerticalStartOffset

`func (o *CapabilityEquipmentSlotArray) HasVerticalStartOffset() bool`

HasVerticalStartOffset returns a boolean if a field has been set.

### GetWidth

`func (o *CapabilityEquipmentSlotArray) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CapabilityEquipmentSlotArray) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CapabilityEquipmentSlotArray) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CapabilityEquipmentSlotArray) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


