# CapabilitySwitchStorageLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchStorageLimits"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchStorageLimits"]
**MaximumUserZoneCount** | Pointer to **int64** | Maximum user zones per Switch/Fabric-Interconnect. | [optional] 
**MaximumVirtualFcInterfaces** | Pointer to **int64** | Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect. | [optional] 
**MaximumVirtualFcInterfacesPerBladeServer** | Pointer to **int64** | Maximum configurable Virtual Fibre Channel interfaces per blade. | [optional] 
**MaximumVsans** | Pointer to **int64** | Maximum configurable VSANs on Switch/Fabric-Interconnect. | [optional] 
**MaximumZoneCount** | Pointer to **int64** | Zone limit per Switch/Fabric-Interconnect. | [optional] 

## Methods

### NewCapabilitySwitchStorageLimits

`func NewCapabilitySwitchStorageLimits(classId string, objectType string, ) *CapabilitySwitchStorageLimits`

NewCapabilitySwitchStorageLimits instantiates a new CapabilitySwitchStorageLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchStorageLimitsWithDefaults

`func NewCapabilitySwitchStorageLimitsWithDefaults() *CapabilitySwitchStorageLimits`

NewCapabilitySwitchStorageLimitsWithDefaults instantiates a new CapabilitySwitchStorageLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchStorageLimits) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchStorageLimits) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchStorageLimits) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchStorageLimits) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchStorageLimits) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchStorageLimits) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaximumUserZoneCount

`func (o *CapabilitySwitchStorageLimits) GetMaximumUserZoneCount() int64`

GetMaximumUserZoneCount returns the MaximumUserZoneCount field if non-nil, zero value otherwise.

### GetMaximumUserZoneCountOk

`func (o *CapabilitySwitchStorageLimits) GetMaximumUserZoneCountOk() (*int64, bool)`

GetMaximumUserZoneCountOk returns a tuple with the MaximumUserZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumUserZoneCount

`func (o *CapabilitySwitchStorageLimits) SetMaximumUserZoneCount(v int64)`

SetMaximumUserZoneCount sets MaximumUserZoneCount field to given value.

### HasMaximumUserZoneCount

`func (o *CapabilitySwitchStorageLimits) HasMaximumUserZoneCount() bool`

HasMaximumUserZoneCount returns a boolean if a field has been set.

### GetMaximumVirtualFcInterfaces

`func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfaces() int64`

GetMaximumVirtualFcInterfaces returns the MaximumVirtualFcInterfaces field if non-nil, zero value otherwise.

### GetMaximumVirtualFcInterfacesOk

`func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfacesOk() (*int64, bool)`

GetMaximumVirtualFcInterfacesOk returns a tuple with the MaximumVirtualFcInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualFcInterfaces

`func (o *CapabilitySwitchStorageLimits) SetMaximumVirtualFcInterfaces(v int64)`

SetMaximumVirtualFcInterfaces sets MaximumVirtualFcInterfaces field to given value.

### HasMaximumVirtualFcInterfaces

`func (o *CapabilitySwitchStorageLimits) HasMaximumVirtualFcInterfaces() bool`

HasMaximumVirtualFcInterfaces returns a boolean if a field has been set.

### GetMaximumVirtualFcInterfacesPerBladeServer

`func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfacesPerBladeServer() int64`

GetMaximumVirtualFcInterfacesPerBladeServer returns the MaximumVirtualFcInterfacesPerBladeServer field if non-nil, zero value otherwise.

### GetMaximumVirtualFcInterfacesPerBladeServerOk

`func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfacesPerBladeServerOk() (*int64, bool)`

GetMaximumVirtualFcInterfacesPerBladeServerOk returns a tuple with the MaximumVirtualFcInterfacesPerBladeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVirtualFcInterfacesPerBladeServer

`func (o *CapabilitySwitchStorageLimits) SetMaximumVirtualFcInterfacesPerBladeServer(v int64)`

SetMaximumVirtualFcInterfacesPerBladeServer sets MaximumVirtualFcInterfacesPerBladeServer field to given value.

### HasMaximumVirtualFcInterfacesPerBladeServer

`func (o *CapabilitySwitchStorageLimits) HasMaximumVirtualFcInterfacesPerBladeServer() bool`

HasMaximumVirtualFcInterfacesPerBladeServer returns a boolean if a field has been set.

### GetMaximumVsans

`func (o *CapabilitySwitchStorageLimits) GetMaximumVsans() int64`

GetMaximumVsans returns the MaximumVsans field if non-nil, zero value otherwise.

### GetMaximumVsansOk

`func (o *CapabilitySwitchStorageLimits) GetMaximumVsansOk() (*int64, bool)`

GetMaximumVsansOk returns a tuple with the MaximumVsans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVsans

`func (o *CapabilitySwitchStorageLimits) SetMaximumVsans(v int64)`

SetMaximumVsans sets MaximumVsans field to given value.

### HasMaximumVsans

`func (o *CapabilitySwitchStorageLimits) HasMaximumVsans() bool`

HasMaximumVsans returns a boolean if a field has been set.

### GetMaximumZoneCount

`func (o *CapabilitySwitchStorageLimits) GetMaximumZoneCount() int64`

GetMaximumZoneCount returns the MaximumZoneCount field if non-nil, zero value otherwise.

### GetMaximumZoneCountOk

`func (o *CapabilitySwitchStorageLimits) GetMaximumZoneCountOk() (*int64, bool)`

GetMaximumZoneCountOk returns a tuple with the MaximumZoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumZoneCount

`func (o *CapabilitySwitchStorageLimits) SetMaximumZoneCount(v int64)`

SetMaximumZoneCount sets MaximumZoneCount field to given value.

### HasMaximumZoneCount

`func (o *CapabilitySwitchStorageLimits) HasMaximumZoneCount() bool`

HasMaximumZoneCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


