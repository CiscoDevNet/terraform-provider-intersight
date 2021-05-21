# VirtualizationVmwareKernelNetwork

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

### NewVirtualizationVmwareKernelNetwork

`func NewVirtualizationVmwareKernelNetwork(classId string, objectType string, ) *VirtualizationVmwareKernelNetwork`

NewVirtualizationVmwareKernelNetwork instantiates a new VirtualizationVmwareKernelNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareKernelNetworkWithDefaults

`func NewVirtualizationVmwareKernelNetworkWithDefaults() *VirtualizationVmwareKernelNetwork`

NewVirtualizationVmwareKernelNetworkWithDefaults instantiates a new VirtualizationVmwareKernelNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareKernelNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareKernelNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareKernelNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareKernelNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareKernelNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareKernelNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFaultToleranceLogging

`func (o *VirtualizationVmwareKernelNetwork) GetFaultToleranceLogging() bool`

GetFaultToleranceLogging returns the FaultToleranceLogging field if non-nil, zero value otherwise.

### GetFaultToleranceLoggingOk

`func (o *VirtualizationVmwareKernelNetwork) GetFaultToleranceLoggingOk() (*bool, bool)`

GetFaultToleranceLoggingOk returns a tuple with the FaultToleranceLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultToleranceLogging

`func (o *VirtualizationVmwareKernelNetwork) SetFaultToleranceLogging(v bool)`

SetFaultToleranceLogging sets FaultToleranceLogging field to given value.

### HasFaultToleranceLogging

`func (o *VirtualizationVmwareKernelNetwork) HasFaultToleranceLogging() bool`

HasFaultToleranceLogging returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationVmwareKernelNetwork) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationVmwareKernelNetwork) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationVmwareKernelNetwork) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationVmwareKernelNetwork) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationVmwareKernelNetwork) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationVmwareKernelNetwork) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMacAddress

`func (o *VirtualizationVmwareKernelNetwork) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwareKernelNetwork) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwareKernelNetwork) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwareKernelNetwork) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetManagement

`func (o *VirtualizationVmwareKernelNetwork) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *VirtualizationVmwareKernelNetwork) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *VirtualizationVmwareKernelNetwork) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *VirtualizationVmwareKernelNetwork) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualizationVmwareKernelNetwork) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualizationVmwareKernelNetwork) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualizationVmwareKernelNetwork) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualizationVmwareKernelNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetVmotion

`func (o *VirtualizationVmwareKernelNetwork) GetVmotion() bool`

GetVmotion returns the Vmotion field if non-nil, zero value otherwise.

### GetVmotionOk

`func (o *VirtualizationVmwareKernelNetwork) GetVmotionOk() (*bool, bool)`

GetVmotionOk returns a tuple with the Vmotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmotion

`func (o *VirtualizationVmwareKernelNetwork) SetVmotion(v bool)`

SetVmotion sets Vmotion field to given value.

### HasVmotion

`func (o *VirtualizationVmwareKernelNetwork) HasVmotion() bool`

HasVmotion returns a boolean if a field has been set.

### GetVsan

`func (o *VirtualizationVmwareKernelNetwork) GetVsan() bool`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *VirtualizationVmwareKernelNetwork) GetVsanOk() (*bool, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *VirtualizationVmwareKernelNetwork) SetVsan(v bool)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *VirtualizationVmwareKernelNetwork) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetVsphereProvisioning

`func (o *VirtualizationVmwareKernelNetwork) GetVsphereProvisioning() bool`

GetVsphereProvisioning returns the VsphereProvisioning field if non-nil, zero value otherwise.

### GetVsphereProvisioningOk

`func (o *VirtualizationVmwareKernelNetwork) GetVsphereProvisioningOk() (*bool, bool)`

GetVsphereProvisioningOk returns a tuple with the VsphereProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereProvisioning

`func (o *VirtualizationVmwareKernelNetwork) SetVsphereProvisioning(v bool)`

SetVsphereProvisioning sets VsphereProvisioning field to given value.

### HasVsphereProvisioning

`func (o *VirtualizationVmwareKernelNetwork) HasVsphereProvisioning() bool`

HasVsphereProvisioning returns a boolean if a field has been set.

### GetVsphereReplication

`func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplication() bool`

GetVsphereReplication returns the VsphereReplication field if non-nil, zero value otherwise.

### GetVsphereReplicationOk

`func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplicationOk() (*bool, bool)`

GetVsphereReplicationOk returns a tuple with the VsphereReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereReplication

`func (o *VirtualizationVmwareKernelNetwork) SetVsphereReplication(v bool)`

SetVsphereReplication sets VsphereReplication field to given value.

### HasVsphereReplication

`func (o *VirtualizationVmwareKernelNetwork) HasVsphereReplication() bool`

HasVsphereReplication returns a boolean if a field has been set.

### GetVsphereReplicationNfc

`func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplicationNfc() bool`

GetVsphereReplicationNfc returns the VsphereReplicationNfc field if non-nil, zero value otherwise.

### GetVsphereReplicationNfcOk

`func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplicationNfcOk() (*bool, bool)`

GetVsphereReplicationNfcOk returns a tuple with the VsphereReplicationNfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereReplicationNfc

`func (o *VirtualizationVmwareKernelNetwork) SetVsphereReplicationNfc(v bool)`

SetVsphereReplicationNfc sets VsphereReplicationNfc field to given value.

### HasVsphereReplicationNfc

`func (o *VirtualizationVmwareKernelNetwork) HasVsphereReplicationNfc() bool`

HasVsphereReplicationNfc returns a boolean if a field has been set.

### GetDistributedNetwork

`func (o *VirtualizationVmwareKernelNetwork) GetDistributedNetwork() VirtualizationVmwareDistributedNetworkRelationship`

GetDistributedNetwork returns the DistributedNetwork field if non-nil, zero value otherwise.

### GetDistributedNetworkOk

`func (o *VirtualizationVmwareKernelNetwork) GetDistributedNetworkOk() (*VirtualizationVmwareDistributedNetworkRelationship, bool)`

GetDistributedNetworkOk returns a tuple with the DistributedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedNetwork

`func (o *VirtualizationVmwareKernelNetwork) SetDistributedNetwork(v VirtualizationVmwareDistributedNetworkRelationship)`

SetDistributedNetwork sets DistributedNetwork field to given value.

### HasDistributedNetwork

`func (o *VirtualizationVmwareKernelNetwork) HasDistributedNetwork() bool`

HasDistributedNetwork returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwareKernelNetwork) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareKernelNetwork) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareKernelNetwork) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareKernelNetwork) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationVmwareKernelNetwork) GetNetwork() VirtualizationVmwareNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationVmwareKernelNetwork) GetNetworkOk() (*VirtualizationVmwareNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationVmwareKernelNetwork) SetNetwork(v VirtualizationVmwareNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationVmwareKernelNetwork) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


