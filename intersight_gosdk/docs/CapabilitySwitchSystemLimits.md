# CapabilitySwitchSystemLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchSystemLimits"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchSystemLimits"]
**MaximumChassisCount** | Pointer to **int64** | Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect. | [optional] 
**MaximumFexPerDomain** | Pointer to **int64** | Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect. | [optional] 
**MaximumServersPerDomain** | Pointer to **int64** | Maximum UCS servers per Switch/Fabric-Interconnect. | [optional] 

## Methods

### NewCapabilitySwitchSystemLimits

`func NewCapabilitySwitchSystemLimits(classId string, objectType string, ) *CapabilitySwitchSystemLimits`

NewCapabilitySwitchSystemLimits instantiates a new CapabilitySwitchSystemLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchSystemLimitsWithDefaults

`func NewCapabilitySwitchSystemLimitsWithDefaults() *CapabilitySwitchSystemLimits`

NewCapabilitySwitchSystemLimitsWithDefaults instantiates a new CapabilitySwitchSystemLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchSystemLimits) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchSystemLimits) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchSystemLimits) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchSystemLimits) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchSystemLimits) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchSystemLimits) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaximumChassisCount

`func (o *CapabilitySwitchSystemLimits) GetMaximumChassisCount() int64`

GetMaximumChassisCount returns the MaximumChassisCount field if non-nil, zero value otherwise.

### GetMaximumChassisCountOk

`func (o *CapabilitySwitchSystemLimits) GetMaximumChassisCountOk() (*int64, bool)`

GetMaximumChassisCountOk returns a tuple with the MaximumChassisCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumChassisCount

`func (o *CapabilitySwitchSystemLimits) SetMaximumChassisCount(v int64)`

SetMaximumChassisCount sets MaximumChassisCount field to given value.

### HasMaximumChassisCount

`func (o *CapabilitySwitchSystemLimits) HasMaximumChassisCount() bool`

HasMaximumChassisCount returns a boolean if a field has been set.

### GetMaximumFexPerDomain

`func (o *CapabilitySwitchSystemLimits) GetMaximumFexPerDomain() int64`

GetMaximumFexPerDomain returns the MaximumFexPerDomain field if non-nil, zero value otherwise.

### GetMaximumFexPerDomainOk

`func (o *CapabilitySwitchSystemLimits) GetMaximumFexPerDomainOk() (*int64, bool)`

GetMaximumFexPerDomainOk returns a tuple with the MaximumFexPerDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFexPerDomain

`func (o *CapabilitySwitchSystemLimits) SetMaximumFexPerDomain(v int64)`

SetMaximumFexPerDomain sets MaximumFexPerDomain field to given value.

### HasMaximumFexPerDomain

`func (o *CapabilitySwitchSystemLimits) HasMaximumFexPerDomain() bool`

HasMaximumFexPerDomain returns a boolean if a field has been set.

### GetMaximumServersPerDomain

`func (o *CapabilitySwitchSystemLimits) GetMaximumServersPerDomain() int64`

GetMaximumServersPerDomain returns the MaximumServersPerDomain field if non-nil, zero value otherwise.

### GetMaximumServersPerDomainOk

`func (o *CapabilitySwitchSystemLimits) GetMaximumServersPerDomainOk() (*int64, bool)`

GetMaximumServersPerDomainOk returns a tuple with the MaximumServersPerDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumServersPerDomain

`func (o *CapabilitySwitchSystemLimits) SetMaximumServersPerDomain(v int64)`

SetMaximumServersPerDomain sets MaximumServersPerDomain field to given value.

### HasMaximumServersPerDomain

`func (o *CapabilitySwitchSystemLimits) HasMaximumServersPerDomain() bool`

HasMaximumServersPerDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


