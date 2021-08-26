# HyperflexHxapVirtualMachineNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapVirtualMachineNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapVirtualMachineNetworkInterface"]
**InterfaceName** | Pointer to **string** | Operating system assigned name for network interface. | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**PrimaryIpAddress** | Pointer to **string** | Primary IP address of the network interface. | [optional] 
**Status** | Pointer to **string** | Current status of virtual network interface status. * &#x60;Up&#x60; - Virtual network interface is up and running. * &#x60;Down&#x60; - Virtual network interface is down and not running. | [optional] [default to "Up"]
**VirtualMachineName** | Pointer to **string** | A reference to the virtual machine where this network object is attached to. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**Network** | Pointer to [**HyperflexHxapNetworkRelationship**](HyperflexHxapNetworkRelationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**HyperflexHxapVirtualMachineRelationship**](HyperflexHxapVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapVirtualMachineNetworkInterface

`func NewHyperflexHxapVirtualMachineNetworkInterface(classId string, objectType string, ) *HyperflexHxapVirtualMachineNetworkInterface`

NewHyperflexHxapVirtualMachineNetworkInterface instantiates a new HyperflexHxapVirtualMachineNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapVirtualMachineNetworkInterfaceWithDefaults

`func NewHyperflexHxapVirtualMachineNetworkInterfaceWithDefaults() *HyperflexHxapVirtualMachineNetworkInterface`

NewHyperflexHxapVirtualMachineNetworkInterfaceWithDefaults instantiates a new HyperflexHxapVirtualMachineNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaceName

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpAddress

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *HyperflexHxapVirtualMachineNetworkInterface) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetPrimaryIpAddress

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetPrimaryIpAddress() string`

GetPrimaryIpAddress returns the PrimaryIpAddress field if non-nil, zero value otherwise.

### GetPrimaryIpAddressOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetPrimaryIpAddressOk() (*string, bool)`

GetPrimaryIpAddressOk returns a tuple with the PrimaryIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIpAddress

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetPrimaryIpAddress(v string)`

SetPrimaryIpAddress sets PrimaryIpAddress field to given value.

### HasPrimaryIpAddress

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasPrimaryIpAddress() bool`

HasPrimaryIpAddress returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualMachineName

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachineName() string`

GetVirtualMachineName returns the VirtualMachineName field if non-nil, zero value otherwise.

### GetVirtualMachineNameOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachineNameOk() (*string, bool)`

GetVirtualMachineNameOk returns a tuple with the VirtualMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineName

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetVirtualMachineName(v string)`

SetVirtualMachineName sets VirtualMachineName field to given value.

### HasVirtualMachineName

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasVirtualMachineName() bool`

HasVirtualMachineName returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNetwork

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetNetwork() HyperflexHxapNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetNetworkOk() (*HyperflexHxapNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetNetwork(v HyperflexHxapNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *HyperflexHxapVirtualMachineNetworkInterface) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *HyperflexHxapVirtualMachineNetworkInterface) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


