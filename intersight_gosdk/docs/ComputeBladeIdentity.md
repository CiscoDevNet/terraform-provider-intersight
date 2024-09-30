# ComputeBladeIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.BladeIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.BladeIdentity"]
**ChassisId** | Pointer to **int64** | Chassis Identifier of a blade server. | [optional] [readonly] 
**CurrentChassisId** | Pointer to **int64** | The id of the chassis that the blade is currently located in. | [optional] [readonly] 
**CurrentSlotId** | Pointer to **int64** | The slot number in the chassis that the blade is currently located in. | [optional] [readonly] 
**FirmwareSupportability** | Pointer to **string** | Describes whether the running CIMC version supports Intersight managed mode. * &#x60;Unknown&#x60; - The running firmware version is unknown. * &#x60;Supported&#x60; - The running firmware version is known and supports IMM mode. * &#x60;NotSupported&#x60; - The running firmware version is known and does not support IMM mode. | [optional] [readonly] [default to "Unknown"]
**ManagerSlotId** | Pointer to **int64** | Chassis slot number of the manager compute server. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence state of the blade server. * &#x60;Unknown&#x60; - The default presence state. * &#x60;Equipped&#x60; - The server is equipped in the slot. * &#x60;EquippedMismatch&#x60; - The slot is equipped, but there is another server currently inventoried in the slot. * &#x60;Missing&#x60; - The server is not present in the given slot. | [optional] [readonly] [default to "Unknown"]
**SlotId** | Pointer to **int64** | Chassis slot number of a blade server. | [optional] [readonly] 
**DiscoveredBladeIdInCurrLocation** | Pointer to [**[]ComputeBladeIdentityRelationship**](ComputeBladeIdentityRelationship.md) | An array of relationships to computeBladeIdentity resources. | [optional] [readonly] 
**ManagedNodes** | Pointer to [**[]ComputeBladeIdentityRelationship**](ComputeBladeIdentityRelationship.md) | An array of relationships to computeBladeIdentity resources. | [optional] [readonly] 
**NewBladeIdInDiscoveredLocation** | Pointer to [**[]ComputeBladeIdentityRelationship**](ComputeBladeIdentityRelationship.md) | An array of relationships to computeBladeIdentity resources. | [optional] [readonly] 

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

### GetCurrentChassisId

`func (o *ComputeBladeIdentity) GetCurrentChassisId() int64`

GetCurrentChassisId returns the CurrentChassisId field if non-nil, zero value otherwise.

### GetCurrentChassisIdOk

`func (o *ComputeBladeIdentity) GetCurrentChassisIdOk() (*int64, bool)`

GetCurrentChassisIdOk returns a tuple with the CurrentChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChassisId

`func (o *ComputeBladeIdentity) SetCurrentChassisId(v int64)`

SetCurrentChassisId sets CurrentChassisId field to given value.

### HasCurrentChassisId

`func (o *ComputeBladeIdentity) HasCurrentChassisId() bool`

HasCurrentChassisId returns a boolean if a field has been set.

### GetCurrentSlotId

`func (o *ComputeBladeIdentity) GetCurrentSlotId() int64`

GetCurrentSlotId returns the CurrentSlotId field if non-nil, zero value otherwise.

### GetCurrentSlotIdOk

`func (o *ComputeBladeIdentity) GetCurrentSlotIdOk() (*int64, bool)`

GetCurrentSlotIdOk returns a tuple with the CurrentSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSlotId

`func (o *ComputeBladeIdentity) SetCurrentSlotId(v int64)`

SetCurrentSlotId sets CurrentSlotId field to given value.

### HasCurrentSlotId

`func (o *ComputeBladeIdentity) HasCurrentSlotId() bool`

HasCurrentSlotId returns a boolean if a field has been set.

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

### GetManagerSlotId

`func (o *ComputeBladeIdentity) GetManagerSlotId() int64`

GetManagerSlotId returns the ManagerSlotId field if non-nil, zero value otherwise.

### GetManagerSlotIdOk

`func (o *ComputeBladeIdentity) GetManagerSlotIdOk() (*int64, bool)`

GetManagerSlotIdOk returns a tuple with the ManagerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerSlotId

`func (o *ComputeBladeIdentity) SetManagerSlotId(v int64)`

SetManagerSlotId sets ManagerSlotId field to given value.

### HasManagerSlotId

`func (o *ComputeBladeIdentity) HasManagerSlotId() bool`

HasManagerSlotId returns a boolean if a field has been set.

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

### GetDiscoveredBladeIdInCurrLocation

`func (o *ComputeBladeIdentity) GetDiscoveredBladeIdInCurrLocation() []ComputeBladeIdentityRelationship`

GetDiscoveredBladeIdInCurrLocation returns the DiscoveredBladeIdInCurrLocation field if non-nil, zero value otherwise.

### GetDiscoveredBladeIdInCurrLocationOk

`func (o *ComputeBladeIdentity) GetDiscoveredBladeIdInCurrLocationOk() (*[]ComputeBladeIdentityRelationship, bool)`

GetDiscoveredBladeIdInCurrLocationOk returns a tuple with the DiscoveredBladeIdInCurrLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredBladeIdInCurrLocation

`func (o *ComputeBladeIdentity) SetDiscoveredBladeIdInCurrLocation(v []ComputeBladeIdentityRelationship)`

SetDiscoveredBladeIdInCurrLocation sets DiscoveredBladeIdInCurrLocation field to given value.

### HasDiscoveredBladeIdInCurrLocation

`func (o *ComputeBladeIdentity) HasDiscoveredBladeIdInCurrLocation() bool`

HasDiscoveredBladeIdInCurrLocation returns a boolean if a field has been set.

### SetDiscoveredBladeIdInCurrLocationNil

`func (o *ComputeBladeIdentity) SetDiscoveredBladeIdInCurrLocationNil(b bool)`

 SetDiscoveredBladeIdInCurrLocationNil sets the value for DiscoveredBladeIdInCurrLocation to be an explicit nil

### UnsetDiscoveredBladeIdInCurrLocation
`func (o *ComputeBladeIdentity) UnsetDiscoveredBladeIdInCurrLocation()`

UnsetDiscoveredBladeIdInCurrLocation ensures that no value is present for DiscoveredBladeIdInCurrLocation, not even an explicit nil
### GetManagedNodes

`func (o *ComputeBladeIdentity) GetManagedNodes() []ComputeBladeIdentityRelationship`

GetManagedNodes returns the ManagedNodes field if non-nil, zero value otherwise.

### GetManagedNodesOk

`func (o *ComputeBladeIdentity) GetManagedNodesOk() (*[]ComputeBladeIdentityRelationship, bool)`

GetManagedNodesOk returns a tuple with the ManagedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNodes

`func (o *ComputeBladeIdentity) SetManagedNodes(v []ComputeBladeIdentityRelationship)`

SetManagedNodes sets ManagedNodes field to given value.

### HasManagedNodes

`func (o *ComputeBladeIdentity) HasManagedNodes() bool`

HasManagedNodes returns a boolean if a field has been set.

### SetManagedNodesNil

`func (o *ComputeBladeIdentity) SetManagedNodesNil(b bool)`

 SetManagedNodesNil sets the value for ManagedNodes to be an explicit nil

### UnsetManagedNodes
`func (o *ComputeBladeIdentity) UnsetManagedNodes()`

UnsetManagedNodes ensures that no value is present for ManagedNodes, not even an explicit nil
### GetNewBladeIdInDiscoveredLocation

`func (o *ComputeBladeIdentity) GetNewBladeIdInDiscoveredLocation() []ComputeBladeIdentityRelationship`

GetNewBladeIdInDiscoveredLocation returns the NewBladeIdInDiscoveredLocation field if non-nil, zero value otherwise.

### GetNewBladeIdInDiscoveredLocationOk

`func (o *ComputeBladeIdentity) GetNewBladeIdInDiscoveredLocationOk() (*[]ComputeBladeIdentityRelationship, bool)`

GetNewBladeIdInDiscoveredLocationOk returns a tuple with the NewBladeIdInDiscoveredLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewBladeIdInDiscoveredLocation

`func (o *ComputeBladeIdentity) SetNewBladeIdInDiscoveredLocation(v []ComputeBladeIdentityRelationship)`

SetNewBladeIdInDiscoveredLocation sets NewBladeIdInDiscoveredLocation field to given value.

### HasNewBladeIdInDiscoveredLocation

`func (o *ComputeBladeIdentity) HasNewBladeIdInDiscoveredLocation() bool`

HasNewBladeIdInDiscoveredLocation returns a boolean if a field has been set.

### SetNewBladeIdInDiscoveredLocationNil

`func (o *ComputeBladeIdentity) SetNewBladeIdInDiscoveredLocationNil(b bool)`

 SetNewBladeIdInDiscoveredLocationNil sets the value for NewBladeIdInDiscoveredLocation to be an explicit nil

### UnsetNewBladeIdInDiscoveredLocation
`func (o *ComputeBladeIdentity) UnsetNewBladeIdInDiscoveredLocation()`

UnsetNewBladeIdInDiscoveredLocation ensures that no value is present for NewBladeIdInDiscoveredLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


