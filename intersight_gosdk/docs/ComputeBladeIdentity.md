# ComputeBladeIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.BladeIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.BladeIdentity"]
**ChassisId** | Pointer to **int64** | Chassis Identifier of a blade server. | [optional] [readonly] 
**FirmwareSupportability** | Pointer to **string** | Describes whether the running CIMC version supports Intersight managed mode. * &#x60;Unknown&#x60; - The running firmware version is unknown. * &#x60;Supported&#x60; - The running firmware version is known and supports IMM mode. * &#x60;NotSupported&#x60; - The running firmware version is known and does not support IMM mode. | [optional] [readonly] [default to "Unknown"]
**Presence** | Pointer to **string** | The presence state of the blade server. * &#x60;Unknown&#x60; - The default presence state. * &#x60;Equipped&#x60; - The server is equipped in the slot. * &#x60;EquippedMismatch&#x60; - The slot is equipped, but there is another server currently inventoried in the slot. * &#x60;Missing&#x60; - The server is not present in the given slot. | [optional] [readonly] [default to "Unknown"]
**SlotId** | Pointer to **int64** | Chassis slot number of a blade server. | [optional] [readonly] 

## Methods

### NewComputeBladeIdentity

`func NewComputeBladeIdentity(classId string, objectType string, ) *ComputeBladeIdentity`

NewComputeBladeIdentity instantiates a new ComputeBladeIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBladeIdentityWithDefaults

`func NewComputeBladeIdentityWithDefaults() *ComputeBladeIdentity`

NewComputeBladeIdentityWithDefaults instantiates a new ComputeBladeIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeBladeIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeBladeIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeBladeIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeBladeIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeBladeIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeBladeIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisId

`func (o *ComputeBladeIdentity) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ComputeBladeIdentity) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ComputeBladeIdentity) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ComputeBladeIdentity) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetFirmwareSupportability

`func (o *ComputeBladeIdentity) GetFirmwareSupportability() string`

GetFirmwareSupportability returns the FirmwareSupportability field if non-nil, zero value otherwise.

### GetFirmwareSupportabilityOk

`func (o *ComputeBladeIdentity) GetFirmwareSupportabilityOk() (*string, bool)`

GetFirmwareSupportabilityOk returns a tuple with the FirmwareSupportability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareSupportability

`func (o *ComputeBladeIdentity) SetFirmwareSupportability(v string)`

SetFirmwareSupportability sets FirmwareSupportability field to given value.

### HasFirmwareSupportability

`func (o *ComputeBladeIdentity) HasFirmwareSupportability() bool`

HasFirmwareSupportability returns a boolean if a field has been set.

### GetPresence

`func (o *ComputeBladeIdentity) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputeBladeIdentity) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputeBladeIdentity) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputeBladeIdentity) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSlotId

`func (o *ComputeBladeIdentity) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *ComputeBladeIdentity) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *ComputeBladeIdentity) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *ComputeBladeIdentity) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


