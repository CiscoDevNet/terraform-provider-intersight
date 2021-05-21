# VirtualizationVmwareKernelNetworkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareKernelNetwork"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareKernelNetwork"]
**FaultToleranceLogging** | Pointer to **bool** | Indicates that fault tolerance logging is enabled on this kernel network. | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **string** | Standard MAC address assigned to this kernel network. | [optional] 
**Management** | Pointer to **bool** | Indicates that management traffic is enabled on this kernel network. | [optional] 
**Mtu** | Pointer to **int64** | Maximum transmission unit configured on a kernel network. | [optional] 
**Vmotion** | Pointer to **bool** | Indicates that vmotion is enabled on this kernel network. | [optional] 
**Vsan** | Pointer to **bool** | Indicates that vsan traffic is enabled on this kernel network. | [optional] 
**VsphereProvisioning** | Pointer to **bool** | Indicates that vsphere provisioning traffic is enabled on this kernel network. | [optional] 
**VsphereReplication** | Pointer to **bool** | Indicates that vsphere replication is enabled on this kernel network. | [optional] 
**VsphereReplicationNfc** | Pointer to **bool** | Indicates that vsphere replication nfc is enabled on this kernel network. | [optional] 
**DistributedNetwork** | Pointer to [**VirtualizationVmwareDistributedNetworkRelationship**](virtualization.VmwareDistributedNetwork.Relationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 
**Network** | Pointer to [**VirtualizationVmwareNetworkRelationship**](virtualization.VmwareNetwork.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareKernelNetworkAllOf

`func NewVirtualizationVmwareKernelNetworkAllOf(classId string, objectType string, ) *VirtualizationVmwareKernelNetworkAllOf`

NewVirtualizationVmwareKernelNetworkAllOf instantiates a new VirtualizationVmwareKernelNetworkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareKernelNetworkAllOfWithDefaults

`func NewVirtualizationVmwareKernelNetworkAllOfWithDefaults() *VirtualizationVmwareKernelNetworkAllOf`

NewVirtualizationVmwareKernelNetworkAllOfWithDefaults instantiates a new VirtualizationVmwareKernelNetworkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFaultToleranceLogging

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetFaultToleranceLogging() bool`

GetFaultToleranceLogging returns the FaultToleranceLogging field if non-nil, zero value otherwise.

### GetFaultToleranceLoggingOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetFaultToleranceLoggingOk() (*bool, bool)`

GetFaultToleranceLoggingOk returns a tuple with the FaultToleranceLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultToleranceLogging

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetFaultToleranceLogging(v bool)`

SetFaultToleranceLogging sets FaultToleranceLogging field to given value.

### HasFaultToleranceLogging

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasFaultToleranceLogging() bool`

HasFaultToleranceLogging returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationVmwareKernelNetworkAllOf) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMacAddress

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetManagement

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetVmotion

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVmotion() bool`

GetVmotion returns the Vmotion field if non-nil, zero value otherwise.

### GetVmotionOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVmotionOk() (*bool, bool)`

GetVmotionOk returns a tuple with the Vmotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmotion

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetVmotion(v bool)`

SetVmotion sets Vmotion field to given value.

### HasVmotion

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasVmotion() bool`

HasVmotion returns a boolean if a field has been set.

### GetVsan

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsan() bool`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsanOk() (*bool, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetVsan(v bool)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetVsphereProvisioning

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsphereProvisioning() bool`

GetVsphereProvisioning returns the VsphereProvisioning field if non-nil, zero value otherwise.

### GetVsphereProvisioningOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsphereProvisioningOk() (*bool, bool)`

GetVsphereProvisioningOk returns a tuple with the VsphereProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereProvisioning

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetVsphereProvisioning(v bool)`

SetVsphereProvisioning sets VsphereProvisioning field to given value.

### HasVsphereProvisioning

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasVsphereProvisioning() bool`

HasVsphereProvisioning returns a boolean if a field has been set.

### GetVsphereReplication

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsphereReplication() bool`

GetVsphereReplication returns the VsphereReplication field if non-nil, zero value otherwise.

### GetVsphereReplicationOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsphereReplicationOk() (*bool, bool)`

GetVsphereReplicationOk returns a tuple with the VsphereReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereReplication

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetVsphereReplication(v bool)`

SetVsphereReplication sets VsphereReplication field to given value.

### HasVsphereReplication

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasVsphereReplication() bool`

HasVsphereReplication returns a boolean if a field has been set.

### GetVsphereReplicationNfc

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsphereReplicationNfc() bool`

GetVsphereReplicationNfc returns the VsphereReplicationNfc field if non-nil, zero value otherwise.

### GetVsphereReplicationNfcOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetVsphereReplicationNfcOk() (*bool, bool)`

GetVsphereReplicationNfcOk returns a tuple with the VsphereReplicationNfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereReplicationNfc

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetVsphereReplicationNfc(v bool)`

SetVsphereReplicationNfc sets VsphereReplicationNfc field to given value.

### HasVsphereReplicationNfc

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasVsphereReplicationNfc() bool`

HasVsphereReplicationNfc returns a boolean if a field has been set.

### GetDistributedNetwork

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetDistributedNetwork() VirtualizationVmwareDistributedNetworkRelationship`

GetDistributedNetwork returns the DistributedNetwork field if non-nil, zero value otherwise.

### GetDistributedNetworkOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetDistributedNetworkOk() (*VirtualizationVmwareDistributedNetworkRelationship, bool)`

GetDistributedNetworkOk returns a tuple with the DistributedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedNetwork

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetDistributedNetwork(v VirtualizationVmwareDistributedNetworkRelationship)`

SetDistributedNetwork sets DistributedNetwork field to given value.

### HasDistributedNetwork

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasDistributedNetwork() bool`

HasDistributedNetwork returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetNetwork() VirtualizationVmwareNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationVmwareKernelNetworkAllOf) GetNetworkOk() (*VirtualizationVmwareNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationVmwareKernelNetworkAllOf) SetNetwork(v VirtualizationVmwareNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationVmwareKernelNetworkAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


