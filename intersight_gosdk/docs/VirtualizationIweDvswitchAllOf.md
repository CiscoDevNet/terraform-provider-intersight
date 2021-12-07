# VirtualizationIweDvswitchAllOf

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

### NewVirtualizationIweDvswitchAllOf

`func NewVirtualizationIweDvswitchAllOf(classId string, objectType string, ) *VirtualizationIweDvswitchAllOf`

NewVirtualizationIweDvswitchAllOf instantiates a new VirtualizationIweDvswitchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweDvswitchAllOfWithDefaults

`func NewVirtualizationIweDvswitchAllOfWithDefaults() *VirtualizationIweDvswitchAllOf`

NewVirtualizationIweDvswitchAllOfWithDefaults instantiates a new VirtualizationIweDvswitchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweDvswitchAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweDvswitchAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweDvswitchAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweDvswitchAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweDvswitchAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweDvswitchAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDvUplink

`func (o *VirtualizationIweDvswitchAllOf) GetDvUplink() string`

GetDvUplink returns the DvUplink field if non-nil, zero value otherwise.

### GetDvUplinkOk

`func (o *VirtualizationIweDvswitchAllOf) GetDvUplinkOk() (*string, bool)`

GetDvUplinkOk returns a tuple with the DvUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvUplink

`func (o *VirtualizationIweDvswitchAllOf) SetDvUplink(v string)`

SetDvUplink sets DvUplink field to given value.

### HasDvUplink

`func (o *VirtualizationIweDvswitchAllOf) HasDvUplink() bool`

HasDvUplink returns a boolean if a field has been set.

### GetLastHostname

`func (o *VirtualizationIweDvswitchAllOf) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *VirtualizationIweDvswitchAllOf) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *VirtualizationIweDvswitchAllOf) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *VirtualizationIweDvswitchAllOf) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweDvswitchAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweDvswitchAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweDvswitchAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweDvswitchAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMemberHosts

`func (o *VirtualizationIweDvswitchAllOf) GetMemberHosts() []VirtualizationIweHostRelationship`

GetMemberHosts returns the MemberHosts field if non-nil, zero value otherwise.

### GetMemberHostsOk

`func (o *VirtualizationIweDvswitchAllOf) GetMemberHostsOk() (*[]VirtualizationIweHostRelationship, bool)`

GetMemberHostsOk returns a tuple with the MemberHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHosts

`func (o *VirtualizationIweDvswitchAllOf) SetMemberHosts(v []VirtualizationIweHostRelationship)`

SetMemberHosts sets MemberHosts field to given value.

### HasMemberHosts

`func (o *VirtualizationIweDvswitchAllOf) HasMemberHosts() bool`

HasMemberHosts returns a boolean if a field has been set.

### SetMemberHostsNil

`func (o *VirtualizationIweDvswitchAllOf) SetMemberHostsNil(b bool)`

 SetMemberHostsNil sets the value for MemberHosts to be an explicit nil

### UnsetMemberHosts
`func (o *VirtualizationIweDvswitchAllOf) UnsetMemberHosts()`

UnsetMemberHosts ensures that no value is present for MemberHosts, not even an explicit nil
### GetMemberVswitches

`func (o *VirtualizationIweDvswitchAllOf) GetMemberVswitches() []VirtualizationIweHostVswitchRelationship`

GetMemberVswitches returns the MemberVswitches field if non-nil, zero value otherwise.

### GetMemberVswitchesOk

`func (o *VirtualizationIweDvswitchAllOf) GetMemberVswitchesOk() (*[]VirtualizationIweHostVswitchRelationship, bool)`

GetMemberVswitchesOk returns a tuple with the MemberVswitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVswitches

`func (o *VirtualizationIweDvswitchAllOf) SetMemberVswitches(v []VirtualizationIweHostVswitchRelationship)`

SetMemberVswitches sets MemberVswitches field to given value.

### HasMemberVswitches

`func (o *VirtualizationIweDvswitchAllOf) HasMemberVswitches() bool`

HasMemberVswitches returns a boolean if a field has been set.

### SetMemberVswitchesNil

`func (o *VirtualizationIweDvswitchAllOf) SetMemberVswitchesNil(b bool)`

 SetMemberVswitchesNil sets the value for MemberVswitches to be an explicit nil

### UnsetMemberVswitches
`func (o *VirtualizationIweDvswitchAllOf) UnsetMemberVswitches()`

UnsetMemberVswitches ensures that no value is present for MemberVswitches, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


