# HyperflexHxapHostInterfaceAllOf

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

### NewHyperflexHxapHostInterfaceAllOf

`func NewHyperflexHxapHostInterfaceAllOf(classId string, objectType string, ) *HyperflexHxapHostInterfaceAllOf`

NewHyperflexHxapHostInterfaceAllOf instantiates a new HyperflexHxapHostInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapHostInterfaceAllOfWithDefaults

`func NewHyperflexHxapHostInterfaceAllOfWithDefaults() *HyperflexHxapHostInterfaceAllOf`

NewHyperflexHxapHostInterfaceAllOfWithDefaults instantiates a new HyperflexHxapHostInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapHostInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapHostInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapHostInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapHostInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *HyperflexHxapHostInterfaceAllOf) GetBondState() HyperflexBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetBondStateOk() (*HyperflexBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *HyperflexHxapHostInterfaceAllOf) SetBondState(v HyperflexBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *HyperflexHxapHostInterfaceAllOf) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *HyperflexHxapHostInterfaceAllOf) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *HyperflexHxapHostInterfaceAllOf) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetHostName

`func (o *HyperflexHxapHostInterfaceAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexHxapHostInterfaceAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexHxapHostInterfaceAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostUuid

`func (o *HyperflexHxapHostInterfaceAllOf) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *HyperflexHxapHostInterfaceAllOf) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *HyperflexHxapHostInterfaceAllOf) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetIfType

`func (o *HyperflexHxapHostInterfaceAllOf) GetIfType() string`

GetIfType returns the IfType field if non-nil, zero value otherwise.

### GetIfTypeOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetIfTypeOk() (*string, bool)`

GetIfTypeOk returns a tuple with the IfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfType

`func (o *HyperflexHxapHostInterfaceAllOf) SetIfType(v string)`

SetIfType sets IfType field to given value.

### HasIfType

`func (o *HyperflexHxapHostInterfaceAllOf) HasIfType() bool`

HasIfType returns a boolean if a field has been set.

### GetIpAddresses

`func (o *HyperflexHxapHostInterfaceAllOf) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *HyperflexHxapHostInterfaceAllOf) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *HyperflexHxapHostInterfaceAllOf) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *HyperflexHxapHostInterfaceAllOf) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *HyperflexHxapHostInterfaceAllOf) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetLinkState

`func (o *HyperflexHxapHostInterfaceAllOf) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *HyperflexHxapHostInterfaceAllOf) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *HyperflexHxapHostInterfaceAllOf) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetMac

`func (o *HyperflexHxapHostInterfaceAllOf) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *HyperflexHxapHostInterfaceAllOf) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *HyperflexHxapHostInterfaceAllOf) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMtu

`func (o *HyperflexHxapHostInterfaceAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *HyperflexHxapHostInterfaceAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *HyperflexHxapHostInterfaceAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *HyperflexHxapHostInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHxapHostInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHxapHostInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlans

`func (o *HyperflexHxapHostInterfaceAllOf) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *HyperflexHxapHostInterfaceAllOf) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *HyperflexHxapHostInterfaceAllOf) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapHostInterfaceAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapHostInterfaceAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapHostInterfaceAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDvUplink

`func (o *HyperflexHxapHostInterfaceAllOf) GetDvUplink() HyperflexHxapDvUplinkRelationship`

GetDvUplink returns the DvUplink field if non-nil, zero value otherwise.

### GetDvUplinkOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetDvUplinkOk() (*HyperflexHxapDvUplinkRelationship, bool)`

GetDvUplinkOk returns a tuple with the DvUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvUplink

`func (o *HyperflexHxapHostInterfaceAllOf) SetDvUplink(v HyperflexHxapDvUplinkRelationship)`

SetDvUplink sets DvUplink field to given value.

### HasDvUplink

`func (o *HyperflexHxapHostInterfaceAllOf) HasDvUplink() bool`

HasDvUplink returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHxapHostInterfaceAllOf) GetHost() HyperflexHxapHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetHostOk() (*HyperflexHxapHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHxapHostInterfaceAllOf) SetHost(v HyperflexHxapHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHxapHostInterfaceAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetwork

`func (o *HyperflexHxapHostInterfaceAllOf) GetNetwork() HyperflexHxapNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HyperflexHxapHostInterfaceAllOf) GetNetworkOk() (*HyperflexHxapNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HyperflexHxapHostInterfaceAllOf) SetNetwork(v HyperflexHxapNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HyperflexHxapHostInterfaceAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


