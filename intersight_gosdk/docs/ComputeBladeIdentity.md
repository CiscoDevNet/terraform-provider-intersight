# ComputeBladeIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | Pointer to **int64** | Chassis Identifier of a blade server. | [optional] 
**DeviceMoId** | Pointer to **string** | FI Device registration Mo ID. | [optional] 
**PendingDiscovery** | Pointer to **string** | Indicates pending discovery flag. | [optional] 
**SlotId** | Pointer to **int64** | Chassis slot number of a blade server. | [optional] 

## Methods

### NewComputeBladeIdentity

`func NewComputeBladeIdentity() *ComputeBladeIdentity`

NewComputeBladeIdentity instantiates a new ComputeBladeIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBladeIdentityWithDefaults

`func NewComputeBladeIdentityWithDefaults() *ComputeBladeIdentity`

NewComputeBladeIdentityWithDefaults instantiates a new ComputeBladeIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetDeviceMoId

`func (o *ComputeBladeIdentity) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ComputeBladeIdentity) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ComputeBladeIdentity) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ComputeBladeIdentity) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetPendingDiscovery

`func (o *ComputeBladeIdentity) GetPendingDiscovery() string`

GetPendingDiscovery returns the PendingDiscovery field if non-nil, zero value otherwise.

### GetPendingDiscoveryOk

`func (o *ComputeBladeIdentity) GetPendingDiscoveryOk() (*string, bool)`

GetPendingDiscoveryOk returns a tuple with the PendingDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDiscovery

`func (o *ComputeBladeIdentity) SetPendingDiscovery(v string)`

SetPendingDiscovery sets PendingDiscovery field to given value.

### HasPendingDiscovery

`func (o *ComputeBladeIdentity) HasPendingDiscovery() bool`

HasPendingDiscovery returns a boolean if a field has been set.

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


