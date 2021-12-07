# VirtualizationIweDvswitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweDvswitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweDvswitch"]
**DvUplink** | Pointer to **string** | The name of the dvUplink referenced by this dvswitch. | [optional] 
**LastHostname** | Pointer to **string** | The last host that update this object. | [optional] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**MemberHosts** | Pointer to [**[]VirtualizationIweHostRelationship**](VirtualizationIweHostRelationship.md) | An array of relationships to virtualizationIweHost resources. | [optional] [readonly] 
**MemberVswitches** | Pointer to [**[]VirtualizationIweHostVswitchRelationship**](VirtualizationIweHostVswitchRelationship.md) | An array of relationships to virtualizationIweHostVswitch resources. | [optional] [readonly] 

## Methods

### NewVirtualizationIweDvswitch

`func NewVirtualizationIweDvswitch(classId string, objectType string, ) *VirtualizationIweDvswitch`

NewVirtualizationIweDvswitch instantiates a new VirtualizationIweDvswitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweDvswitchWithDefaults

`func NewVirtualizationIweDvswitchWithDefaults() *VirtualizationIweDvswitch`

NewVirtualizationIweDvswitchWithDefaults instantiates a new VirtualizationIweDvswitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweDvswitch) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweDvswitch) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweDvswitch) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweDvswitch) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweDvswitch) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweDvswitch) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDvUplink

`func (o *VirtualizationIweDvswitch) GetDvUplink() string`

GetDvUplink returns the DvUplink field if non-nil, zero value otherwise.

### GetDvUplinkOk

`func (o *VirtualizationIweDvswitch) GetDvUplinkOk() (*string, bool)`

GetDvUplinkOk returns a tuple with the DvUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvUplink

`func (o *VirtualizationIweDvswitch) SetDvUplink(v string)`

SetDvUplink sets DvUplink field to given value.

### HasDvUplink

`func (o *VirtualizationIweDvswitch) HasDvUplink() bool`

HasDvUplink returns a boolean if a field has been set.

### GetLastHostname

`func (o *VirtualizationIweDvswitch) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *VirtualizationIweDvswitch) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *VirtualizationIweDvswitch) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *VirtualizationIweDvswitch) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweDvswitch) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweDvswitch) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweDvswitch) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweDvswitch) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMemberHosts

`func (o *VirtualizationIweDvswitch) GetMemberHosts() []VirtualizationIweHostRelationship`

GetMemberHosts returns the MemberHosts field if non-nil, zero value otherwise.

### GetMemberHostsOk

`func (o *VirtualizationIweDvswitch) GetMemberHostsOk() (*[]VirtualizationIweHostRelationship, bool)`

GetMemberHostsOk returns a tuple with the MemberHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHosts

`func (o *VirtualizationIweDvswitch) SetMemberHosts(v []VirtualizationIweHostRelationship)`

SetMemberHosts sets MemberHosts field to given value.

### HasMemberHosts

`func (o *VirtualizationIweDvswitch) HasMemberHosts() bool`

HasMemberHosts returns a boolean if a field has been set.

### SetMemberHostsNil

`func (o *VirtualizationIweDvswitch) SetMemberHostsNil(b bool)`

 SetMemberHostsNil sets the value for MemberHosts to be an explicit nil

### UnsetMemberHosts
`func (o *VirtualizationIweDvswitch) UnsetMemberHosts()`

UnsetMemberHosts ensures that no value is present for MemberHosts, not even an explicit nil
### GetMemberVswitches

`func (o *VirtualizationIweDvswitch) GetMemberVswitches() []VirtualizationIweHostVswitchRelationship`

GetMemberVswitches returns the MemberVswitches field if non-nil, zero value otherwise.

### GetMemberVswitchesOk

`func (o *VirtualizationIweDvswitch) GetMemberVswitchesOk() (*[]VirtualizationIweHostVswitchRelationship, bool)`

GetMemberVswitchesOk returns a tuple with the MemberVswitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVswitches

`func (o *VirtualizationIweDvswitch) SetMemberVswitches(v []VirtualizationIweHostVswitchRelationship)`

SetMemberVswitches sets MemberVswitches field to given value.

### HasMemberVswitches

`func (o *VirtualizationIweDvswitch) HasMemberVswitches() bool`

HasMemberVswitches returns a boolean if a field has been set.

### SetMemberVswitchesNil

`func (o *VirtualizationIweDvswitch) SetMemberVswitchesNil(b bool)`

 SetMemberVswitchesNil sets the value for MemberVswitches to be an explicit nil

### UnsetMemberVswitches
`func (o *VirtualizationIweDvswitch) UnsetMemberVswitches()`

UnsetMemberVswitches ensures that no value is present for MemberVswitches, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


