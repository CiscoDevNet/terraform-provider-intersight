# VirtualizationIweDvUplinkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweDvUplink"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweDvUplink"]
**BondState** | Pointer to [**NullableVirtualizationBondState**](VirtualizationBondState.md) |  | [optional] 
**NetInterfaces** | Pointer to **[]string** |  | [optional] 
**Vlans** | Pointer to **string** | The vlans associated with this this cluster wide uplink. | [optional] [readonly] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**MemberHosts** | Pointer to [**[]VirtualizationIweHostRelationship**](VirtualizationIweHostRelationship.md) | An array of relationships to virtualizationIweHost resources. | [optional] [readonly] 
**MemberUplinks** | Pointer to [**[]VirtualizationIweHostInterfaceRelationship**](VirtualizationIweHostInterfaceRelationship.md) | An array of relationships to virtualizationIweHostInterface resources. | [optional] [readonly] 

## Methods

### NewVirtualizationIweDvUplinkAllOf

`func NewVirtualizationIweDvUplinkAllOf(classId string, objectType string, ) *VirtualizationIweDvUplinkAllOf`

NewVirtualizationIweDvUplinkAllOf instantiates a new VirtualizationIweDvUplinkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweDvUplinkAllOfWithDefaults

`func NewVirtualizationIweDvUplinkAllOfWithDefaults() *VirtualizationIweDvUplinkAllOf`

NewVirtualizationIweDvUplinkAllOfWithDefaults instantiates a new VirtualizationIweDvUplinkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweDvUplinkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweDvUplinkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweDvUplinkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweDvUplinkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweDvUplinkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweDvUplinkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *VirtualizationIweDvUplinkAllOf) GetBondState() VirtualizationBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *VirtualizationIweDvUplinkAllOf) GetBondStateOk() (*VirtualizationBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *VirtualizationIweDvUplinkAllOf) SetBondState(v VirtualizationBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *VirtualizationIweDvUplinkAllOf) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *VirtualizationIweDvUplinkAllOf) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *VirtualizationIweDvUplinkAllOf) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetNetInterfaces

`func (o *VirtualizationIweDvUplinkAllOf) GetNetInterfaces() []string`

GetNetInterfaces returns the NetInterfaces field if non-nil, zero value otherwise.

### GetNetInterfacesOk

`func (o *VirtualizationIweDvUplinkAllOf) GetNetInterfacesOk() (*[]string, bool)`

GetNetInterfacesOk returns a tuple with the NetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterfaces

`func (o *VirtualizationIweDvUplinkAllOf) SetNetInterfaces(v []string)`

SetNetInterfaces sets NetInterfaces field to given value.

### HasNetInterfaces

`func (o *VirtualizationIweDvUplinkAllOf) HasNetInterfaces() bool`

HasNetInterfaces returns a boolean if a field has been set.

### SetNetInterfacesNil

`func (o *VirtualizationIweDvUplinkAllOf) SetNetInterfacesNil(b bool)`

 SetNetInterfacesNil sets the value for NetInterfaces to be an explicit nil

### UnsetNetInterfaces
`func (o *VirtualizationIweDvUplinkAllOf) UnsetNetInterfaces()`

UnsetNetInterfaces ensures that no value is present for NetInterfaces, not even an explicit nil
### GetVlans

`func (o *VirtualizationIweDvUplinkAllOf) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *VirtualizationIweDvUplinkAllOf) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *VirtualizationIweDvUplinkAllOf) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *VirtualizationIweDvUplinkAllOf) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweDvUplinkAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweDvUplinkAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweDvUplinkAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweDvUplinkAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMemberHosts

`func (o *VirtualizationIweDvUplinkAllOf) GetMemberHosts() []VirtualizationIweHostRelationship`

GetMemberHosts returns the MemberHosts field if non-nil, zero value otherwise.

### GetMemberHostsOk

`func (o *VirtualizationIweDvUplinkAllOf) GetMemberHostsOk() (*[]VirtualizationIweHostRelationship, bool)`

GetMemberHostsOk returns a tuple with the MemberHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHosts

`func (o *VirtualizationIweDvUplinkAllOf) SetMemberHosts(v []VirtualizationIweHostRelationship)`

SetMemberHosts sets MemberHosts field to given value.

### HasMemberHosts

`func (o *VirtualizationIweDvUplinkAllOf) HasMemberHosts() bool`

HasMemberHosts returns a boolean if a field has been set.

### SetMemberHostsNil

`func (o *VirtualizationIweDvUplinkAllOf) SetMemberHostsNil(b bool)`

 SetMemberHostsNil sets the value for MemberHosts to be an explicit nil

### UnsetMemberHosts
`func (o *VirtualizationIweDvUplinkAllOf) UnsetMemberHosts()`

UnsetMemberHosts ensures that no value is present for MemberHosts, not even an explicit nil
### GetMemberUplinks

`func (o *VirtualizationIweDvUplinkAllOf) GetMemberUplinks() []VirtualizationIweHostInterfaceRelationship`

GetMemberUplinks returns the MemberUplinks field if non-nil, zero value otherwise.

### GetMemberUplinksOk

`func (o *VirtualizationIweDvUplinkAllOf) GetMemberUplinksOk() (*[]VirtualizationIweHostInterfaceRelationship, bool)`

GetMemberUplinksOk returns a tuple with the MemberUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUplinks

`func (o *VirtualizationIweDvUplinkAllOf) SetMemberUplinks(v []VirtualizationIweHostInterfaceRelationship)`

SetMemberUplinks sets MemberUplinks field to given value.

### HasMemberUplinks

`func (o *VirtualizationIweDvUplinkAllOf) HasMemberUplinks() bool`

HasMemberUplinks returns a boolean if a field has been set.

### SetMemberUplinksNil

`func (o *VirtualizationIweDvUplinkAllOf) SetMemberUplinksNil(b bool)`

 SetMemberUplinksNil sets the value for MemberUplinks to be an explicit nil

### UnsetMemberUplinks
`func (o *VirtualizationIweDvUplinkAllOf) UnsetMemberUplinks()`

UnsetMemberUplinks ensures that no value is present for MemberUplinks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


