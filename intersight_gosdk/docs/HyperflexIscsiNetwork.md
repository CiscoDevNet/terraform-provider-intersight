# HyperflexIscsiNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.IscsiNetwork"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.IscsiNetwork"]
**Gateway** | Pointer to **string** | The gateway of the HyperFlex iSCSI network. | [optional] [readonly] 
**InventorySource** | Pointer to **string** | Source of this inventory object. * &#x60;NOT_APPLICABLE&#x60; - The source of the HyperFlex inventory is not applicable. * &#x60;ONLINE&#x60; - The source of the HyperFlex inventory is online. * &#x60;OFFLINE&#x60; - The source of the HyperFlex inventory is offline. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**IpRanges** | Pointer to [**[]NetworkHyperFlexInternetProtocolAddressRange**](NetworkHyperFlexInternetProtocolAddressRange.md) |  | [optional] 
**IscsiCip** | Pointer to **string** | An IP within the iSCSI IP Address Range which is CIP for iSCSI network. | [optional] [readonly] 
**Mtu** | Pointer to **string** | The maximum transmission unit of the HyperFlex iSCSI network. * &#x60;UNKNOWN&#x60; - The maximum transmission unit of the HyperFlex iSCSI network is unknown. * &#x60;MTU_1500&#x60; - The maximum transmission unit of the HyperFlex iSCSI network is 1500 bytes. * &#x60;MTU_9000&#x60; - The maximum transmission unit of the HyperFlex iSCSI network is 9000 bytes. | [optional] [readonly] [default to "UNKNOWN"]
**Name** | Pointer to **string** | Name of the HyperFlex iSCSI network configuration. | [optional] [readonly] 
**Subnet** | Pointer to **string** | Subnet of the HyperFlex iSCSI network. Subnet is in a.b.c.d/e notation. | [optional] [readonly] 
**UcsHost** | Pointer to **string** | UCS Manager Host IP or FQDN. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the HyperFlex iSCSI network configuration. | [optional] [readonly] 
**Version** | Pointer to **int64** | Version of Network configuration in Inventory. | [optional] [readonly] 
**VlanName** | Pointer to **string** | The Virtual local area network (VLAN) name of the HyperFlex iSCSI network. | [optional] [readonly] 
**Vlanid** | Pointer to **int64** | The VLAN ID of the HyperFlex iSCSI network. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexIscsiNetwork

`func NewHyperflexIscsiNetwork(classId string, objectType string, ) *HyperflexIscsiNetwork`

NewHyperflexIscsiNetwork instantiates a new HyperflexIscsiNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexIscsiNetworkWithDefaults

`func NewHyperflexIscsiNetworkWithDefaults() *HyperflexIscsiNetwork`

NewHyperflexIscsiNetworkWithDefaults instantiates a new HyperflexIscsiNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexIscsiNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexIscsiNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexIscsiNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexIscsiNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexIscsiNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexIscsiNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *HyperflexIscsiNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *HyperflexIscsiNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *HyperflexIscsiNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *HyperflexIscsiNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInventorySource

`func (o *HyperflexIscsiNetwork) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *HyperflexIscsiNetwork) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *HyperflexIscsiNetwork) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *HyperflexIscsiNetwork) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetIpRanges

`func (o *HyperflexIscsiNetwork) GetIpRanges() []NetworkHyperFlexInternetProtocolAddressRange`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *HyperflexIscsiNetwork) GetIpRangesOk() (*[]NetworkHyperFlexInternetProtocolAddressRange, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *HyperflexIscsiNetwork) SetIpRanges(v []NetworkHyperFlexInternetProtocolAddressRange)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *HyperflexIscsiNetwork) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *HyperflexIscsiNetwork) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *HyperflexIscsiNetwork) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetIscsiCip

`func (o *HyperflexIscsiNetwork) GetIscsiCip() string`

GetIscsiCip returns the IscsiCip field if non-nil, zero value otherwise.

### GetIscsiCipOk

`func (o *HyperflexIscsiNetwork) GetIscsiCipOk() (*string, bool)`

GetIscsiCipOk returns a tuple with the IscsiCip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiCip

`func (o *HyperflexIscsiNetwork) SetIscsiCip(v string)`

SetIscsiCip sets IscsiCip field to given value.

### HasIscsiCip

`func (o *HyperflexIscsiNetwork) HasIscsiCip() bool`

HasIscsiCip returns a boolean if a field has been set.

### GetMtu

`func (o *HyperflexIscsiNetwork) GetMtu() string`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *HyperflexIscsiNetwork) GetMtuOk() (*string, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *HyperflexIscsiNetwork) SetMtu(v string)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *HyperflexIscsiNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *HyperflexIscsiNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexIscsiNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexIscsiNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexIscsiNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *HyperflexIscsiNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *HyperflexIscsiNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *HyperflexIscsiNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *HyperflexIscsiNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetUcsHost

`func (o *HyperflexIscsiNetwork) GetUcsHost() string`

GetUcsHost returns the UcsHost field if non-nil, zero value otherwise.

### GetUcsHostOk

`func (o *HyperflexIscsiNetwork) GetUcsHostOk() (*string, bool)`

GetUcsHostOk returns a tuple with the UcsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsHost

`func (o *HyperflexIscsiNetwork) SetUcsHost(v string)`

SetUcsHost sets UcsHost field to given value.

### HasUcsHost

`func (o *HyperflexIscsiNetwork) HasUcsHost() bool`

HasUcsHost returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexIscsiNetwork) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexIscsiNetwork) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexIscsiNetwork) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexIscsiNetwork) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexIscsiNetwork) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexIscsiNetwork) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexIscsiNetwork) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexIscsiNetwork) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVlanName

`func (o *HyperflexIscsiNetwork) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *HyperflexIscsiNetwork) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *HyperflexIscsiNetwork) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *HyperflexIscsiNetwork) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanid

`func (o *HyperflexIscsiNetwork) GetVlanid() int64`

GetVlanid returns the Vlanid field if non-nil, zero value otherwise.

### GetVlanidOk

`func (o *HyperflexIscsiNetwork) GetVlanidOk() (*int64, bool)`

GetVlanidOk returns a tuple with the Vlanid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanid

`func (o *HyperflexIscsiNetwork) SetVlanid(v int64)`

SetVlanid sets Vlanid field to given value.

### HasVlanid

`func (o *HyperflexIscsiNetwork) HasVlanid() bool`

HasVlanid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexIscsiNetwork) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexIscsiNetwork) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexIscsiNetwork) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexIscsiNetwork) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HyperflexIscsiNetwork) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HyperflexIscsiNetwork) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


