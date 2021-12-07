# VirtualizationIweDvUplink

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

### NewVirtualizationIweDvUplink

`func NewVirtualizationIweDvUplink(classId string, objectType string, ) *VirtualizationIweDvUplink`

NewVirtualizationIweDvUplink instantiates a new VirtualizationIweDvUplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweDvUplinkWithDefaults

`func NewVirtualizationIweDvUplinkWithDefaults() *VirtualizationIweDvUplink`

NewVirtualizationIweDvUplinkWithDefaults instantiates a new VirtualizationIweDvUplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweDvUplink) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweDvUplink) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweDvUplink) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweDvUplink) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweDvUplink) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweDvUplink) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBondState

`func (o *VirtualizationIweDvUplink) GetBondState() VirtualizationBondState`

GetBondState returns the BondState field if non-nil, zero value otherwise.

### GetBondStateOk

`func (o *VirtualizationIweDvUplink) GetBondStateOk() (*VirtualizationBondState, bool)`

GetBondStateOk returns a tuple with the BondState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondState

`func (o *VirtualizationIweDvUplink) SetBondState(v VirtualizationBondState)`

SetBondState sets BondState field to given value.

### HasBondState

`func (o *VirtualizationIweDvUplink) HasBondState() bool`

HasBondState returns a boolean if a field has been set.

### SetBondStateNil

`func (o *VirtualizationIweDvUplink) SetBondStateNil(b bool)`

 SetBondStateNil sets the value for BondState to be an explicit nil

### UnsetBondState
`func (o *VirtualizationIweDvUplink) UnsetBondState()`

UnsetBondState ensures that no value is present for BondState, not even an explicit nil
### GetNetInterfaces

`func (o *VirtualizationIweDvUplink) GetNetInterfaces() []string`

GetNetInterfaces returns the NetInterfaces field if non-nil, zero value otherwise.

### GetNetInterfacesOk

`func (o *VirtualizationIweDvUplink) GetNetInterfacesOk() (*[]string, bool)`

GetNetInterfacesOk returns a tuple with the NetInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterfaces

`func (o *VirtualizationIweDvUplink) SetNetInterfaces(v []string)`

SetNetInterfaces sets NetInterfaces field to given value.

### HasNetInterfaces

`func (o *VirtualizationIweDvUplink) HasNetInterfaces() bool`

HasNetInterfaces returns a boolean if a field has been set.

### SetNetInterfacesNil

`func (o *VirtualizationIweDvUplink) SetNetInterfacesNil(b bool)`

 SetNetInterfacesNil sets the value for NetInterfaces to be an explicit nil

### UnsetNetInterfaces
`func (o *VirtualizationIweDvUplink) UnsetNetInterfaces()`

UnsetNetInterfaces ensures that no value is present for NetInterfaces, not even an explicit nil
### GetVlans

`func (o *VirtualizationIweDvUplink) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *VirtualizationIweDvUplink) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *VirtualizationIweDvUplink) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *VirtualizationIweDvUplink) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweDvUplink) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweDvUplink) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweDvUplink) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweDvUplink) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMemberHosts

`func (o *VirtualizationIweDvUplink) GetMemberHosts() []VirtualizationIweHostRelationship`

GetMemberHosts returns the MemberHosts field if non-nil, zero value otherwise.

### GetMemberHostsOk

`func (o *VirtualizationIweDvUplink) GetMemberHostsOk() (*[]VirtualizationIweHostRelationship, bool)`

GetMemberHostsOk returns a tuple with the MemberHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHosts

`func (o *VirtualizationIweDvUplink) SetMemberHosts(v []VirtualizationIweHostRelationship)`

SetMemberHosts sets MemberHosts field to given value.

### HasMemberHosts

`func (o *VirtualizationIweDvUplink) HasMemberHosts() bool`

HasMemberHosts returns a boolean if a field has been set.

### SetMemberHostsNil

`func (o *VirtualizationIweDvUplink) SetMemberHostsNil(b bool)`

 SetMemberHostsNil sets the value for MemberHosts to be an explicit nil

### UnsetMemberHosts
`func (o *VirtualizationIweDvUplink) UnsetMemberHosts()`

UnsetMemberHosts ensures that no value is present for MemberHosts, not even an explicit nil
### GetMemberUplinks

`func (o *VirtualizationIweDvUplink) GetMemberUplinks() []VirtualizationIweHostInterfaceRelationship`

GetMemberUplinks returns the MemberUplinks field if non-nil, zero value otherwise.

### GetMemberUplinksOk

`func (o *VirtualizationIweDvUplink) GetMemberUplinksOk() (*[]VirtualizationIweHostInterfaceRelationship, bool)`

GetMemberUplinksOk returns a tuple with the MemberUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUplinks

`func (o *VirtualizationIweDvUplink) SetMemberUplinks(v []VirtualizationIweHostInterfaceRelationship)`

SetMemberUplinks sets MemberUplinks field to given value.

### HasMemberUplinks

`func (o *VirtualizationIweDvUplink) HasMemberUplinks() bool`

HasMemberUplinks returns a boolean if a field has been set.

### SetMemberUplinksNil

`func (o *VirtualizationIweDvUplink) SetMemberUplinksNil(b bool)`

 SetMemberUplinksNil sets the value for MemberUplinks to be an explicit nil

### UnsetMemberUplinks
`func (o *VirtualizationIweDvUplink) UnsetMemberUplinks()`

UnsetMemberUplinks ensures that no value is present for MemberUplinks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


