# HyperflexIscsiNetworkAllOf

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
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexIscsiNetworkAllOf

`func NewHyperflexIscsiNetworkAllOf(classId string, objectType string, ) *HyperflexIscsiNetworkAllOf`

NewHyperflexIscsiNetworkAllOf instantiates a new HyperflexIscsiNetworkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexIscsiNetworkAllOfWithDefaults

`func NewHyperflexIscsiNetworkAllOfWithDefaults() *HyperflexIscsiNetworkAllOf`

NewHyperflexIscsiNetworkAllOfWithDefaults instantiates a new HyperflexIscsiNetworkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexIscsiNetworkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexIscsiNetworkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexIscsiNetworkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexIscsiNetworkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexIscsiNetworkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexIscsiNetworkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *HyperflexIscsiNetworkAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *HyperflexIscsiNetworkAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *HyperflexIscsiNetworkAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *HyperflexIscsiNetworkAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInventorySource

`func (o *HyperflexIscsiNetworkAllOf) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *HyperflexIscsiNetworkAllOf) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *HyperflexIscsiNetworkAllOf) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *HyperflexIscsiNetworkAllOf) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetIpRanges

`func (o *HyperflexIscsiNetworkAllOf) GetIpRanges() []NetworkHyperFlexInternetProtocolAddressRange`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *HyperflexIscsiNetworkAllOf) GetIpRangesOk() (*[]NetworkHyperFlexInternetProtocolAddressRange, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *HyperflexIscsiNetworkAllOf) SetIpRanges(v []NetworkHyperFlexInternetProtocolAddressRange)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *HyperflexIscsiNetworkAllOf) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *HyperflexIscsiNetworkAllOf) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *HyperflexIscsiNetworkAllOf) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetIscsiCip

`func (o *HyperflexIscsiNetworkAllOf) GetIscsiCip() string`

GetIscsiCip returns the IscsiCip field if non-nil, zero value otherwise.

### GetIscsiCipOk

`func (o *HyperflexIscsiNetworkAllOf) GetIscsiCipOk() (*string, bool)`

GetIscsiCipOk returns a tuple with the IscsiCip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiCip

`func (o *HyperflexIscsiNetworkAllOf) SetIscsiCip(v string)`

SetIscsiCip sets IscsiCip field to given value.

### HasIscsiCip

`func (o *HyperflexIscsiNetworkAllOf) HasIscsiCip() bool`

HasIscsiCip returns a boolean if a field has been set.

### GetMtu

`func (o *HyperflexIscsiNetworkAllOf) GetMtu() string`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *HyperflexIscsiNetworkAllOf) GetMtuOk() (*string, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *HyperflexIscsiNetworkAllOf) SetMtu(v string)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *HyperflexIscsiNetworkAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *HyperflexIscsiNetworkAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexIscsiNetworkAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexIscsiNetworkAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexIscsiNetworkAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *HyperflexIscsiNetworkAllOf) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *HyperflexIscsiNetworkAllOf) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *HyperflexIscsiNetworkAllOf) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *HyperflexIscsiNetworkAllOf) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetUcsHost

`func (o *HyperflexIscsiNetworkAllOf) GetUcsHost() string`

GetUcsHost returns the UcsHost field if non-nil, zero value otherwise.

### GetUcsHostOk

`func (o *HyperflexIscsiNetworkAllOf) GetUcsHostOk() (*string, bool)`

GetUcsHostOk returns a tuple with the UcsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsHost

`func (o *HyperflexIscsiNetworkAllOf) SetUcsHost(v string)`

SetUcsHost sets UcsHost field to given value.

### HasUcsHost

`func (o *HyperflexIscsiNetworkAllOf) HasUcsHost() bool`

HasUcsHost returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexIscsiNetworkAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexIscsiNetworkAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexIscsiNetworkAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexIscsiNetworkAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexIscsiNetworkAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexIscsiNetworkAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexIscsiNetworkAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexIscsiNetworkAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVlanName

`func (o *HyperflexIscsiNetworkAllOf) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *HyperflexIscsiNetworkAllOf) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *HyperflexIscsiNetworkAllOf) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *HyperflexIscsiNetworkAllOf) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanid

`func (o *HyperflexIscsiNetworkAllOf) GetVlanid() int64`

GetVlanid returns the Vlanid field if non-nil, zero value otherwise.

### GetVlanidOk

`func (o *HyperflexIscsiNetworkAllOf) GetVlanidOk() (*int64, bool)`

GetVlanidOk returns a tuple with the Vlanid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanid

`func (o *HyperflexIscsiNetworkAllOf) SetVlanid(v int64)`

SetVlanid sets Vlanid field to given value.

### HasVlanid

`func (o *HyperflexIscsiNetworkAllOf) HasVlanid() bool`

HasVlanid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexIscsiNetworkAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexIscsiNetworkAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexIscsiNetworkAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexIscsiNetworkAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


