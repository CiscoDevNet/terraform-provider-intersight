# StorageNetAppEthernetPortLagAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppEthernetPortLag"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppEthernetPortLag"]
**ActivePorts** | Pointer to [**[]StorageNetAppPort**](StorageNetAppPort.md) |  | [optional] 
**DistributionPolicy** | Pointer to **string** | Policy for mapping flows to ports for outbound packets through a LAG (ifgrp). * &#x60;none&#x60; - Default unknown distribution policy type. * &#x60;port&#x60; - Network traffic is distributed based on the transport layer (TCP/UDP) ports. * &#x60;ip&#x60; - Network traffic is distributed based on IP addresses. * &#x60;mac&#x60; - Network traffic is distributed based on MAC addresses. * &#x60;sequential&#x60; - Network traffic is distributed in round-robin fashion from the list of configured, available ports. | [optional] [default to "none"]
**MemberPorts** | Pointer to [**[]StorageNetAppPort**](StorageNetAppPort.md) |  | [optional] 
**Mode** | Pointer to **string** | Determines how the ports interact with the switch. * &#x60;none&#x60; - Default unknown lag mode type. * &#x60;multimode_lacp&#x60; - Bundle multiple member ports of the interface group using Link Aggregation Control Protocol. * &#x60;multimode&#x60; - Bundle multiple member ports of the interface group to act as a single trunked port. * &#x60;singlemode&#x60; - Provide port redundancy using member ports of the interface group for failover. | [optional] [default to "none"]

## Methods

### NewStorageNetAppEthernetPortLagAllOf

`func NewStorageNetAppEthernetPortLagAllOf(classId string, objectType string, ) *StorageNetAppEthernetPortLagAllOf`

NewStorageNetAppEthernetPortLagAllOf instantiates a new StorageNetAppEthernetPortLagAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortLagAllOfWithDefaults

`func NewStorageNetAppEthernetPortLagAllOfWithDefaults() *StorageNetAppEthernetPortLagAllOf`

NewStorageNetAppEthernetPortLagAllOfWithDefaults instantiates a new StorageNetAppEthernetPortLagAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPortLagAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPortLagAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPortLagAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPortLagAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPortLagAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPortLagAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActivePorts

`func (o *StorageNetAppEthernetPortLagAllOf) GetActivePorts() []StorageNetAppPort`

GetActivePorts returns the ActivePorts field if non-nil, zero value otherwise.

### GetActivePortsOk

`func (o *StorageNetAppEthernetPortLagAllOf) GetActivePortsOk() (*[]StorageNetAppPort, bool)`

GetActivePortsOk returns a tuple with the ActivePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePorts

`func (o *StorageNetAppEthernetPortLagAllOf) SetActivePorts(v []StorageNetAppPort)`

SetActivePorts sets ActivePorts field to given value.

### HasActivePorts

`func (o *StorageNetAppEthernetPortLagAllOf) HasActivePorts() bool`

HasActivePorts returns a boolean if a field has been set.

### SetActivePortsNil

`func (o *StorageNetAppEthernetPortLagAllOf) SetActivePortsNil(b bool)`

 SetActivePortsNil sets the value for ActivePorts to be an explicit nil

### UnsetActivePorts
`func (o *StorageNetAppEthernetPortLagAllOf) UnsetActivePorts()`

UnsetActivePorts ensures that no value is present for ActivePorts, not even an explicit nil
### GetDistributionPolicy

`func (o *StorageNetAppEthernetPortLagAllOf) GetDistributionPolicy() string`

GetDistributionPolicy returns the DistributionPolicy field if non-nil, zero value otherwise.

### GetDistributionPolicyOk

`func (o *StorageNetAppEthernetPortLagAllOf) GetDistributionPolicyOk() (*string, bool)`

GetDistributionPolicyOk returns a tuple with the DistributionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPolicy

`func (o *StorageNetAppEthernetPortLagAllOf) SetDistributionPolicy(v string)`

SetDistributionPolicy sets DistributionPolicy field to given value.

### HasDistributionPolicy

`func (o *StorageNetAppEthernetPortLagAllOf) HasDistributionPolicy() bool`

HasDistributionPolicy returns a boolean if a field has been set.

### GetMemberPorts

`func (o *StorageNetAppEthernetPortLagAllOf) GetMemberPorts() []StorageNetAppPort`

GetMemberPorts returns the MemberPorts field if non-nil, zero value otherwise.

### GetMemberPortsOk

`func (o *StorageNetAppEthernetPortLagAllOf) GetMemberPortsOk() (*[]StorageNetAppPort, bool)`

GetMemberPortsOk returns a tuple with the MemberPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPorts

`func (o *StorageNetAppEthernetPortLagAllOf) SetMemberPorts(v []StorageNetAppPort)`

SetMemberPorts sets MemberPorts field to given value.

### HasMemberPorts

`func (o *StorageNetAppEthernetPortLagAllOf) HasMemberPorts() bool`

HasMemberPorts returns a boolean if a field has been set.

### SetMemberPortsNil

`func (o *StorageNetAppEthernetPortLagAllOf) SetMemberPortsNil(b bool)`

 SetMemberPortsNil sets the value for MemberPorts to be an explicit nil

### UnsetMemberPorts
`func (o *StorageNetAppEthernetPortLagAllOf) UnsetMemberPorts()`

UnsetMemberPorts ensures that no value is present for MemberPorts, not even an explicit nil
### GetMode

`func (o *StorageNetAppEthernetPortLagAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StorageNetAppEthernetPortLagAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StorageNetAppEthernetPortLagAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StorageNetAppEthernetPortLagAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


