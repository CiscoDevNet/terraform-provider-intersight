# HyperflexHxapHostInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapHostInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapHostInterface"]
**BondState** | Pointer to [**NullableHyperflexBondState**](HyperflexBondState.md) |  | [optional] 
**HostName** | Pointer to **string** | The UUID of the host to which this interface belongs to. | [optional] 
**HostUuid** | Pointer to **string** | The UUID of the host to which this interface belongs to. | [optional] 
**IfType** | Pointer to **string** | A hint of the interface type, such as \&quot;ovs-bond\&quot;, \&quot;device\&quot;, \&quot;openvswitch\&quot;. | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**LinkState** | Pointer to **string** | Link state information such as \&quot;up\&quot;, \&quot;down\&quot;. * &#x60;unknown&#x60; - The interface line is unknown. * &#x60;up&#x60; - The interface line is up. * &#x60;down&#x60; - The interface line is down. * &#x60;degraded&#x60; - For a bond/team interface, not all member interface is up. | [optional] [default to "unknown"]
**Mac** | Pointer to **string** | The MAC address of the interface. | [optional] 
**Mtu** | Pointer to **int64** | The MTU size of the interface. | [optional] 
**Name** | Pointer to **string** | The name of the host to which this interface belongs to. | [optional] 
**Vlans** | Pointer to **string** | A list of vlans allowed on this interface. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**DvUplink** | Pointer to [**HyperflexHxapDvUplinkRelationship**](HyperflexHxapDvUplinkRelationship.md) |  | [optional] 
**Host** | Pointer to [**HyperflexHxapHostRelationship**](HyperflexHxapHostRelationship.md) |  | [optional] 
**Network** | Pointer to [**HyperflexHxapNetworkRelationship**](HyperflexHxapNetworkRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapHostInterface

`func NewHyperflexHxapHostInterface(classId string, objectType string, ) *HyperflexHxapHostInterface`

NewHyperflexHxapHostInterface instantiates a new HyperflexHxapHostInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapHostInterfaceWithDefaults

`func NewHyperflexHxapHostInterfaceWithDefaults() *HyperflexHxapHostInterface`

NewHyperflexHxapHostInterfaceWithDefaults instantiates a new HyperflexHxapHostInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapHostInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapHostInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapHostInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapHostInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapHostInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapHostInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *HyperflexHxapHostInterface) GetBondState() HyperflexBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *HyperflexHxapHostInterface) GetBondStateOk() (*HyperflexBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *HyperflexHxapHostInterface) SetBondState(v HyperflexBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *HyperflexHxapHostInterface) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *HyperflexHxapHostInterface) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *HyperflexHxapHostInterface) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetHostName

`func (o *HyperflexHxapHostInterface) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexHxapHostInterface) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexHxapHostInterface) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexHxapHostInterface) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostUuid

`func (o *HyperflexHxapHostInterface) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *HyperflexHxapHostInterface) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *HyperflexHxapHostInterface) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *HyperflexHxapHostInterface) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetIfType

`func (o *HyperflexHxapHostInterface) GetIfType() string`

GetIfType returns the IfType field if non-nil, zero value otherwise.

### GetIfTypeOk

`func (o *HyperflexHxapHostInterface) GetIfTypeOk() (*string, bool)`

GetIfTypeOk returns a tuple with the IfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfType

`func (o *HyperflexHxapHostInterface) SetIfType(v string)`

SetIfType sets IfType field to given value.

### HasIfType

`func (o *HyperflexHxapHostInterface) HasIfType() bool`

HasIfType returns a boolean if a field has been set.

### GetIpAddresses

`func (o *HyperflexHxapHostInterface) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *HyperflexHxapHostInterface) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *HyperflexHxapHostInterface) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *HyperflexHxapHostInterface) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *HyperflexHxapHostInterface) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *HyperflexHxapHostInterface) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetLinkState

`func (o *HyperflexHxapHostInterface) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *HyperflexHxapHostInterface) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *HyperflexHxapHostInterface) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *HyperflexHxapHostInterface) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetMac

`func (o *HyperflexHxapHostInterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *HyperflexHxapHostInterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *HyperflexHxapHostInterface) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *HyperflexHxapHostInterface) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMtu

`func (o *HyperflexHxapHostInterface) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *HyperflexHxapHostInterface) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *HyperflexHxapHostInterface) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *HyperflexHxapHostInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *HyperflexHxapHostInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHxapHostInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHxapHostInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHxapHostInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlans

`func (o *HyperflexHxapHostInterface) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *HyperflexHxapHostInterface) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *HyperflexHxapHostInterface) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *HyperflexHxapHostInterface) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapHostInterface) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapHostInterface) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapHostInterface) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapHostInterface) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDvUplink

`func (o *HyperflexHxapHostInterface) GetDvUplink() HyperflexHxapDvUplinkRelationship`

GetDvUplink returns the DvUplink field if non-nil, zero value otherwise.

### GetDvUplinkOk

`func (o *HyperflexHxapHostInterface) GetDvUplinkOk() (*HyperflexHxapDvUplinkRelationship, bool)`

GetDvUplinkOk returns a tuple with the DvUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvUplink

`func (o *HyperflexHxapHostInterface) SetDvUplink(v HyperflexHxapDvUplinkRelationship)`

SetDvUplink sets DvUplink field to given value.

### HasDvUplink

`func (o *HyperflexHxapHostInterface) HasDvUplink() bool`

HasDvUplink returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHxapHostInterface) GetHost() HyperflexHxapHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHxapHostInterface) GetHostOk() (*HyperflexHxapHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHxapHostInterface) SetHost(v HyperflexHxapHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHxapHostInterface) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetwork

`func (o *HyperflexHxapHostInterface) GetNetwork() HyperflexHxapNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HyperflexHxapHostInterface) GetNetworkOk() (*HyperflexHxapNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HyperflexHxapHostInterface) SetNetwork(v HyperflexHxapNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HyperflexHxapHostInterface) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


