# VirtualizationVmwareUplinkPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareUplinkPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareUplinkPort"]
**Identity** | Pointer to **string** | The VMware managed object reference as a string. | [optional] 
**Key** | Pointer to **string** | The internally assigned key of this uplink port object. This entity is not manipulated by users. | [optional] 
**Name** | Pointer to **string** | User-provided name to identify the uplink port object. | [optional] 
**DistributedNetwork** | Pointer to [**VirtualizationVmwareDistributedNetworkRelationship**](virtualization.VmwareDistributedNetwork.Relationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 
**PhysicalNetworkInterface** | Pointer to [**VirtualizationVmwarePhysicalNetworkInterfaceRelationship**](virtualization.VmwarePhysicalNetworkInterface.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareUplinkPortAllOf

`func NewVirtualizationVmwareUplinkPortAllOf(classId string, objectType string, ) *VirtualizationVmwareUplinkPortAllOf`

NewVirtualizationVmwareUplinkPortAllOf instantiates a new VirtualizationVmwareUplinkPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareUplinkPortAllOfWithDefaults

`func NewVirtualizationVmwareUplinkPortAllOfWithDefaults() *VirtualizationVmwareUplinkPortAllOf`

NewVirtualizationVmwareUplinkPortAllOfWithDefaults instantiates a new VirtualizationVmwareUplinkPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareUplinkPortAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareUplinkPortAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareUplinkPortAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareUplinkPortAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *VirtualizationVmwareUplinkPortAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationVmwareUplinkPortAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationVmwareUplinkPortAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetKey

`func (o *VirtualizationVmwareUplinkPortAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VirtualizationVmwareUplinkPortAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VirtualizationVmwareUplinkPortAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmwareUplinkPortAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmwareUplinkPortAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmwareUplinkPortAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistributedNetwork

`func (o *VirtualizationVmwareUplinkPortAllOf) GetDistributedNetwork() VirtualizationVmwareDistributedNetworkRelationship`

GetDistributedNetwork returns the DistributedNetwork field if non-nil, zero value otherwise.

### GetDistributedNetworkOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetDistributedNetworkOk() (*VirtualizationVmwareDistributedNetworkRelationship, bool)`

GetDistributedNetworkOk returns a tuple with the DistributedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedNetwork

`func (o *VirtualizationVmwareUplinkPortAllOf) SetDistributedNetwork(v VirtualizationVmwareDistributedNetworkRelationship)`

SetDistributedNetwork sets DistributedNetwork field to given value.

### HasDistributedNetwork

`func (o *VirtualizationVmwareUplinkPortAllOf) HasDistributedNetwork() bool`

HasDistributedNetwork returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwareUplinkPortAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareUplinkPortAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareUplinkPortAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPhysicalNetworkInterface

`func (o *VirtualizationVmwareUplinkPortAllOf) GetPhysicalNetworkInterface() VirtualizationVmwarePhysicalNetworkInterfaceRelationship`

GetPhysicalNetworkInterface returns the PhysicalNetworkInterface field if non-nil, zero value otherwise.

### GetPhysicalNetworkInterfaceOk

`func (o *VirtualizationVmwareUplinkPortAllOf) GetPhysicalNetworkInterfaceOk() (*VirtualizationVmwarePhysicalNetworkInterfaceRelationship, bool)`

GetPhysicalNetworkInterfaceOk returns a tuple with the PhysicalNetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalNetworkInterface

`func (o *VirtualizationVmwareUplinkPortAllOf) SetPhysicalNetworkInterface(v VirtualizationVmwarePhysicalNetworkInterfaceRelationship)`

SetPhysicalNetworkInterface sets PhysicalNetworkInterface field to given value.

### HasPhysicalNetworkInterface

`func (o *VirtualizationVmwareUplinkPortAllOf) HasPhysicalNetworkInterface() bool`

HasPhysicalNetworkInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


