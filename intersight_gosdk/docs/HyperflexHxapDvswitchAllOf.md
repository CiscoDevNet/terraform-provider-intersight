# HyperflexHxapDvswitchAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapDvswitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapDvswitch"]
**DvUplink** | Pointer to **string** | The name of the dvUplink referenced by this dvswitch. | [optional] 
**LastHostname** | Pointer to **string** | The last host that update this object. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**MemberHosts** | Pointer to [**[]HyperflexHxapHostRelationship**](HyperflexHxapHostRelationship.md) | An array of relationships to hyperflexHxapHost resources. | [optional] [readonly] 
**MemberVswitches** | Pointer to [**[]HyperflexHxapHostVswitchRelationship**](HyperflexHxapHostVswitchRelationship.md) | An array of relationships to hyperflexHxapHostVswitch resources. | [optional] [readonly] 

## Methods

### NewHyperflexHxapDvswitchAllOf

`func NewHyperflexHxapDvswitchAllOf(classId string, objectType string, ) *HyperflexHxapDvswitchAllOf`

NewHyperflexHxapDvswitchAllOf instantiates a new HyperflexHxapDvswitchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapDvswitchAllOfWithDefaults

`func NewHyperflexHxapDvswitchAllOfWithDefaults() *HyperflexHxapDvswitchAllOf`

NewHyperflexHxapDvswitchAllOfWithDefaults instantiates a new HyperflexHxapDvswitchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapDvswitchAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapDvswitchAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapDvswitchAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapDvswitchAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapDvswitchAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapDvswitchAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDvUplink

`func (o *HyperflexHxapDvswitchAllOf) GetDvUplink() string`

GetDvUplink returns the DvUplink field if non-nil, zero value otherwise.

### GetDvUplinkOk

`func (o *HyperflexHxapDvswitchAllOf) GetDvUplinkOk() (*string, bool)`

GetDvUplinkOk returns a tuple with the DvUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvUplink

`func (o *HyperflexHxapDvswitchAllOf) SetDvUplink(v string)`

SetDvUplink sets DvUplink field to given value.

### HasDvUplink

`func (o *HyperflexHxapDvswitchAllOf) HasDvUplink() bool`

HasDvUplink returns a boolean if a field has been set.

### GetLastHostname

`func (o *HyperflexHxapDvswitchAllOf) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *HyperflexHxapDvswitchAllOf) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *HyperflexHxapDvswitchAllOf) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *HyperflexHxapDvswitchAllOf) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapDvswitchAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapDvswitchAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapDvswitchAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapDvswitchAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMemberHosts

`func (o *HyperflexHxapDvswitchAllOf) GetMemberHosts() []HyperflexHxapHostRelationship`

GetMemberHosts returns the MemberHosts field if non-nil, zero value otherwise.

### GetMemberHostsOk

`func (o *HyperflexHxapDvswitchAllOf) GetMemberHostsOk() (*[]HyperflexHxapHostRelationship, bool)`

GetMemberHostsOk returns a tuple with the MemberHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberHosts

`func (o *HyperflexHxapDvswitchAllOf) SetMemberHosts(v []HyperflexHxapHostRelationship)`

SetMemberHosts sets MemberHosts field to given value.

### HasMemberHosts

`func (o *HyperflexHxapDvswitchAllOf) HasMemberHosts() bool`

HasMemberHosts returns a boolean if a field has been set.

### SetMemberHostsNil

`func (o *HyperflexHxapDvswitchAllOf) SetMemberHostsNil(b bool)`

 SetMemberHostsNil sets the value for MemberHosts to be an explicit nil

### UnsetMemberHosts
`func (o *HyperflexHxapDvswitchAllOf) UnsetMemberHosts()`

UnsetMemberHosts ensures that no value is present for MemberHosts, not even an explicit nil
### GetMemberVswitches

`func (o *HyperflexHxapDvswitchAllOf) GetMemberVswitches() []HyperflexHxapHostVswitchRelationship`

GetMemberVswitches returns the MemberVswitches field if non-nil, zero value otherwise.

### GetMemberVswitchesOk

`func (o *HyperflexHxapDvswitchAllOf) GetMemberVswitchesOk() (*[]HyperflexHxapHostVswitchRelationship, bool)`

GetMemberVswitchesOk returns a tuple with the MemberVswitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVswitches

`func (o *HyperflexHxapDvswitchAllOf) SetMemberVswitches(v []HyperflexHxapHostVswitchRelationship)`

SetMemberVswitches sets MemberVswitches field to given value.

### HasMemberVswitches

`func (o *HyperflexHxapDvswitchAllOf) HasMemberVswitches() bool`

HasMemberVswitches returns a boolean if a field has been set.

### SetMemberVswitchesNil

`func (o *HyperflexHxapDvswitchAllOf) SetMemberVswitchesNil(b bool)`

 SetMemberVswitchesNil sets the value for MemberVswitches to be an explicit nil

### UnsetMemberVswitches
`func (o *HyperflexHxapDvswitchAllOf) UnsetMemberVswitches()`

UnsetMemberVswitches ensures that no value is present for MemberVswitches, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


