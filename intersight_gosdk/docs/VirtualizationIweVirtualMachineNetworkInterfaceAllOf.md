# VirtualizationIweVirtualMachineNetworkInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweVirtualMachineNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweVirtualMachineNetworkInterface"]
**InterfaceName** | Pointer to **string** | Operating system assigned name for network interface. | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**PrimaryIpAddress** | Pointer to **string** | Primary IP address of the network interface. | [optional] 
**Status** | Pointer to **string** | Current status of virtual network interface status. * &#x60;Up&#x60; - Virtual network interface is up and running. * &#x60;Down&#x60; - Virtual network interface is down and not running. | [optional] [default to "Up"]
**VirtualMachineName** | Pointer to **string** | A reference to the virtual machine where this network object is attached to. | [optional] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**Network** | Pointer to [**VirtualizationIweNetworkRelationship**](VirtualizationIweNetworkRelationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationIweVirtualMachineRelationship**](VirtualizationIweVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweVirtualMachineNetworkInterfaceAllOf

`func NewVirtualizationIweVirtualMachineNetworkInterfaceAllOf(classId string, objectType string, ) *VirtualizationIweVirtualMachineNetworkInterfaceAllOf`

NewVirtualizationIweVirtualMachineNetworkInterfaceAllOf instantiates a new VirtualizationIweVirtualMachineNetworkInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweVirtualMachineNetworkInterfaceAllOfWithDefaults

`func NewVirtualizationIweVirtualMachineNetworkInterfaceAllOfWithDefaults() *VirtualizationIweVirtualMachineNetworkInterfaceAllOf`

NewVirtualizationIweVirtualMachineNetworkInterfaceAllOfWithDefaults instantiates a new VirtualizationIweVirtualMachineNetworkInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaceName

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetPrimaryIpAddress

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetPrimaryIpAddress() string`

GetPrimaryIpAddress returns the PrimaryIpAddress field if non-nil, zero value otherwise.

### GetPrimaryIpAddressOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetPrimaryIpAddressOk() (*string, bool)`

GetPrimaryIpAddressOk returns a tuple with the PrimaryIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIpAddress

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetPrimaryIpAddress(v string)`

SetPrimaryIpAddress sets PrimaryIpAddress field to given value.

### HasPrimaryIpAddress

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasPrimaryIpAddress() bool`

HasPrimaryIpAddress returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualMachineName

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetVirtualMachineName() string`

GetVirtualMachineName returns the VirtualMachineName field if non-nil, zero value otherwise.

### GetVirtualMachineNameOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetVirtualMachineNameOk() (*string, bool)`

GetVirtualMachineNameOk returns a tuple with the VirtualMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineName

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetVirtualMachineName(v string)`

SetVirtualMachineName sets VirtualMachineName field to given value.

### HasVirtualMachineName

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasVirtualMachineName() bool`

HasVirtualMachineName returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetNetwork() VirtualizationIweNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetNetworkOk() (*VirtualizationIweNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetNetwork(v VirtualizationIweNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetVirtualMachine() VirtualizationIweVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) GetVirtualMachineOk() (*VirtualizationIweVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) SetVirtualMachine(v VirtualizationIweVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationIweVirtualMachineNetworkInterfaceAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


