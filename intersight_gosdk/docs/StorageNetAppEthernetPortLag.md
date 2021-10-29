# StorageNetAppEthernetPortLag

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

### NewStorageNetAppEthernetPortLag

`func NewStorageNetAppEthernetPortLag(classId string, objectType string, ) *StorageNetAppEthernetPortLag`

NewStorageNetAppEthernetPortLag instantiates a new StorageNetAppEthernetPortLag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortLagWithDefaults

`func NewStorageNetAppEthernetPortLagWithDefaults() *StorageNetAppEthernetPortLag`

NewStorageNetAppEthernetPortLagWithDefaults instantiates a new StorageNetAppEthernetPortLag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPortLag) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPortLag) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPortLag) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPortLag) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPortLag) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPortLag) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActivePorts

`func (o *StorageNetAppEthernetPortLag) GetActivePorts() []StorageNetAppPort`

GetActivePorts returns the ActivePorts field if non-nil, zero value otherwise.

### GetActivePortsOk

`func (o *StorageNetAppEthernetPortLag) GetActivePortsOk() (*[]StorageNetAppPort, bool)`

GetActivePortsOk returns a tuple with the ActivePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePorts

`func (o *StorageNetAppEthernetPortLag) SetActivePorts(v []StorageNetAppPort)`

SetActivePorts sets ActivePorts field to given value.

### HasActivePorts

`func (o *StorageNetAppEthernetPortLag) HasActivePorts() bool`

HasActivePorts returns a boolean if a field has been set.

### SetActivePortsNil

`func (o *StorageNetAppEthernetPortLag) SetActivePortsNil(b bool)`

 SetActivePortsNil sets the value for ActivePorts to be an explicit nil

### UnsetActivePorts
`func (o *StorageNetAppEthernetPortLag) UnsetActivePorts()`

UnsetActivePorts ensures that no value is present for ActivePorts, not even an explicit nil
### GetDistributionPolicy

`func (o *StorageNetAppEthernetPortLag) GetDistributionPolicy() string`

GetDistributionPolicy returns the DistributionPolicy field if non-nil, zero value otherwise.

### GetDistributionPolicyOk

`func (o *StorageNetAppEthernetPortLag) GetDistributionPolicyOk() (*string, bool)`

GetDistributionPolicyOk returns a tuple with the DistributionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPolicy

`func (o *StorageNetAppEthernetPortLag) SetDistributionPolicy(v string)`

SetDistributionPolicy sets DistributionPolicy field to given value.

### HasDistributionPolicy

`func (o *StorageNetAppEthernetPortLag) HasDistributionPolicy() bool`

HasDistributionPolicy returns a boolean if a field has been set.

### GetMemberPorts

`func (o *StorageNetAppEthernetPortLag) GetMemberPorts() []StorageNetAppPort`

GetMemberPorts returns the MemberPorts field if non-nil, zero value otherwise.

### GetMemberPortsOk

`func (o *StorageNetAppEthernetPortLag) GetMemberPortsOk() (*[]StorageNetAppPort, bool)`

GetMemberPortsOk returns a tuple with the MemberPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPorts

`func (o *StorageNetAppEthernetPortLag) SetMemberPorts(v []StorageNetAppPort)`

SetMemberPorts sets MemberPorts field to given value.

### HasMemberPorts

`func (o *StorageNetAppEthernetPortLag) HasMemberPorts() bool`

HasMemberPorts returns a boolean if a field has been set.

### SetMemberPortsNil

`func (o *StorageNetAppEthernetPortLag) SetMemberPortsNil(b bool)`

 SetMemberPortsNil sets the value for MemberPorts to be an explicit nil

### UnsetMemberPorts
`func (o *StorageNetAppEthernetPortLag) UnsetMemberPorts()`

UnsetMemberPorts ensures that no value is present for MemberPorts, not even an explicit nil
### GetMode

`func (o *StorageNetAppEthernetPortLag) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StorageNetAppEthernetPortLag) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StorageNetAppEthernetPortLag) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StorageNetAppEthernetPortLag) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


