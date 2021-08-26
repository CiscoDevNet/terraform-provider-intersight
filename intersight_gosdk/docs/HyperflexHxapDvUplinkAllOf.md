# HyperflexHxapDvUplinkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapDvUplink"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapDvUplink"]
**BondState** | Pointer to [**NullableHyperflexBondState**](HyperflexBondState.md) |  | [optional] 
**NetInterfaces** | Pointer to **[]string** |  | [optional] 
**Vlans** | Pointer to **string** | The vlans associated with this this cluster wide uplink. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**MemberHosts** | Pointer to [**[]HyperflexHxapHostRelationship**](HyperflexHxapHostRelationship.md) | An array of relationships to hyperflexHxapHost resources. | [optional] [readonly] 
**MemberUplinks** | Pointer to [**[]HyperflexHxapHostInterfaceRelationship**](HyperflexHxapHostInterfaceRelationship.md) | An array of relationships to hyperflexHxapHostInterface resources. | [optional] [readonly] 

## Methods

### NewHyperflexHxapDvUplinkAllOf

`func NewHyperflexHxapDvUplinkAllOf(classId string, objectType string, ) *HyperflexHxapDvUplinkAllOf`

NewHyperflexHxapDvUplinkAllOf instantiates a new HyperflexHxapDvUplinkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapDvUplinkAllOfWithDefaults

`func NewHyperflexHxapDvUplinkAllOfWithDefaults() *HyperflexHxapDvUplinkAllOf`

NewHyperflexHxapDvUplinkAllOfWithDefaults instantiates a new HyperflexHxapDvUplinkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapDvUplinkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapDvUplinkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapDvUplinkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapDvUplinkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapDvUplinkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapDvUplinkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *HyperflexHxapDvUplinkAllOf) GetBondState() HyperflexBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *HyperflexHxapDvUplinkAllOf) GetBondStateOk() (*HyperflexBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *HyperflexHxapDvUplinkAllOf) SetBondState(v HyperflexBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *HyperflexHxapDvUplinkAllOf) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *HyperflexHxapDvUplinkAllOf) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *HyperflexHxapDvUplinkAllOf) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetNetInterfaces

`func (o *HyperflexHxapDvUplinkAllOf) GetNetInterfaces() []string`

GetNetInterfaces returns the NetInterfaces field if non-nil, zero value otherwise.

### GetNetInterfacesOk

`func (o *HyperflexHxapDvUplinkAllOf) GetNetInterfacesOk() (*[]string, bool)`

GetNetInterfacesOk returns a tuple with the NetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterfaces

`func (o *HyperflexHxapDvUplinkAllOf) SetNetInterfaces(v []string)`

SetNetInterfaces sets NetInterfaces field to given value.

### HasNetInterfaces

`func (o *HyperflexHxapDvUplinkAllOf) HasNetInterfaces() bool`

HasNetInterfaces returns a boolean if a field has been set.

### SetNetInterfacesNil

`func (o *HyperflexHxapDvUplinkAllOf) SetNetInterfacesNil(b bool)`

 SetNetInterfacesNil sets the value for NetInterfaces to be an explicit nil

### UnsetNetInterfaces
`func (o *HyperflexHxapDvUplinkAllOf) UnsetNetInterfaces()`

UnsetNetInterfaces ensures that no value is present for NetInterfaces, not even an explicit nil
### GetVlans

`func (o *HyperflexHxapDvUplinkAllOf) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *HyperflexHxapDvUplinkAllOf) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *HyperflexHxapDvUplinkAllOf) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *HyperflexHxapDvUplinkAllOf) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapDvUplinkAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapDvUplinkAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapDvUplinkAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapDvUplinkAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMemberHosts

`func (o *HyperflexHxapDvUplinkAllOf) GetMemberHosts() []HyperflexHxapHostRelationship`

GetMemberHosts returns the MemberHosts field if non-nil, zero value otherwise.

### GetMemberHostsOk

`func (o *HyperflexHxapDvUplinkAllOf) GetMemberHostsOk() (*[]HyperflexHxapHostRelationship, bool)`

GetMemberHostsOk returns a tuple with the MemberHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHosts

`func (o *HyperflexHxapDvUplinkAllOf) SetMemberHosts(v []HyperflexHxapHostRelationship)`

SetMemberHosts sets MemberHosts field to given value.

### HasMemberHosts

`func (o *HyperflexHxapDvUplinkAllOf) HasMemberHosts() bool`

HasMemberHosts returns a boolean if a field has been set.

### SetMemberHostsNil

`func (o *HyperflexHxapDvUplinkAllOf) SetMemberHostsNil(b bool)`

 SetMemberHostsNil sets the value for MemberHosts to be an explicit nil

### UnsetMemberHosts
`func (o *HyperflexHxapDvUplinkAllOf) UnsetMemberHosts()`

UnsetMemberHosts ensures that no value is present for MemberHosts, not even an explicit nil
### GetMemberUplinks

`func (o *HyperflexHxapDvUplinkAllOf) GetMemberUplinks() []HyperflexHxapHostInterfaceRelationship`

GetMemberUplinks returns the MemberUplinks field if non-nil, zero value otherwise.

### GetMemberUplinksOk

`func (o *HyperflexHxapDvUplinkAllOf) GetMemberUplinksOk() (*[]HyperflexHxapHostInterfaceRelationship, bool)`

GetMemberUplinksOk returns a tuple with the MemberUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUplinks

`func (o *HyperflexHxapDvUplinkAllOf) SetMemberUplinks(v []HyperflexHxapHostInterfaceRelationship)`

SetMemberUplinks sets MemberUplinks field to given value.

### HasMemberUplinks

`func (o *HyperflexHxapDvUplinkAllOf) HasMemberUplinks() bool`

HasMemberUplinks returns a boolean if a field has been set.

### SetMemberUplinksNil

`func (o *HyperflexHxapDvUplinkAllOf) SetMemberUplinksNil(b bool)`

 SetMemberUplinksNil sets the value for MemberUplinks to be an explicit nil

### UnsetMemberUplinks
`func (o *HyperflexHxapDvUplinkAllOf) UnsetMemberUplinks()`

UnsetMemberUplinks ensures that no value is present for MemberUplinks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


