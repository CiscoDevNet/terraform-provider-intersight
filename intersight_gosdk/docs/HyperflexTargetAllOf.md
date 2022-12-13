# HyperflexTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Target"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Target"]
**AuthMethod** | Pointer to **string** | Auth method of the target inventory. * &#x60;NOT_APPLICABLE&#x60; - Authorization method of the HyperFlex iSCSI target is not applicable. * &#x60;CHAP&#x60; - Authorization method of the HyperFlex iSCSI target is CHAP. * &#x60;NONE&#x60; - Authorization method of the HyperFlex iSCSI target is none. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**InitiatorGroupUuids** | Pointer to **[]string** |  | [optional] 
**InventorySource** | Pointer to **string** | Source of the target inventory. * &#x60;NOT_APPLICABLE&#x60; - The source of the HyperFlex inventory is not applicable. * &#x60;ONLINE&#x60; - The source of the HyperFlex inventory is online. * &#x60;OFFLINE&#x60; - The source of the HyperFlex inventory is offline. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**Iqn** | Pointer to **string** | The iSCSI qualified name (IQN) of target. | [optional] [readonly] 
**LunUuids** | Pointer to **[]string** |  | [optional] 
**NumActiveInitiators** | Pointer to **int64** | Number of active initiators in the initiator group. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the HyperFlex iSCSI target. | [optional] [readonly] 
**Version** | Pointer to **int64** | Version of the Initiator Group. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**InitiatorGroups** | Pointer to [**[]HyperflexInitiatorGroupRelationship**](HyperflexInitiatorGroupRelationship.md) | An array of relationships to hyperflexInitiatorGroup resources. | [optional] [readonly] 
**Luns** | Pointer to [**[]HyperflexLunRelationship**](HyperflexLunRelationship.md) | An array of relationships to hyperflexLun resources. | [optional] [readonly] 

## Methods

### NewHyperflexTargetAllOf

`func NewHyperflexTargetAllOf(classId string, objectType string, ) *HyperflexTargetAllOf`

NewHyperflexTargetAllOf instantiates a new HyperflexTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexTargetAllOfWithDefaults

`func NewHyperflexTargetAllOfWithDefaults() *HyperflexTargetAllOf`

NewHyperflexTargetAllOfWithDefaults instantiates a new HyperflexTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexTargetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexTargetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexTargetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexTargetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexTargetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexTargetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthMethod

`func (o *HyperflexTargetAllOf) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *HyperflexTargetAllOf) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *HyperflexTargetAllOf) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *HyperflexTargetAllOf) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetInitiatorGroupUuids

`func (o *HyperflexTargetAllOf) GetInitiatorGroupUuids() []string`

GetInitiatorGroupUuids returns the InitiatorGroupUuids field if non-nil, zero value otherwise.

### GetInitiatorGroupUuidsOk

`func (o *HyperflexTargetAllOf) GetInitiatorGroupUuidsOk() (*[]string, bool)`

GetInitiatorGroupUuidsOk returns a tuple with the InitiatorGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorGroupUuids

`func (o *HyperflexTargetAllOf) SetInitiatorGroupUuids(v []string)`

SetInitiatorGroupUuids sets InitiatorGroupUuids field to given value.

### HasInitiatorGroupUuids

`func (o *HyperflexTargetAllOf) HasInitiatorGroupUuids() bool`

HasInitiatorGroupUuids returns a boolean if a field has been set.

### SetInitiatorGroupUuidsNil

`func (o *HyperflexTargetAllOf) SetInitiatorGroupUuidsNil(b bool)`

 SetInitiatorGroupUuidsNil sets the value for InitiatorGroupUuids to be an explicit nil

### UnsetInitiatorGroupUuids
`func (o *HyperflexTargetAllOf) UnsetInitiatorGroupUuids()`

UnsetInitiatorGroupUuids ensures that no value is present for InitiatorGroupUuids, not even an explicit nil
### GetInventorySource

`func (o *HyperflexTargetAllOf) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *HyperflexTargetAllOf) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *HyperflexTargetAllOf) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *HyperflexTargetAllOf) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetIqn

`func (o *HyperflexTargetAllOf) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *HyperflexTargetAllOf) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *HyperflexTargetAllOf) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *HyperflexTargetAllOf) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetLunUuids

`func (o *HyperflexTargetAllOf) GetLunUuids() []string`

GetLunUuids returns the LunUuids field if non-nil, zero value otherwise.

### GetLunUuidsOk

`func (o *HyperflexTargetAllOf) GetLunUuidsOk() (*[]string, bool)`

GetLunUuidsOk returns a tuple with the LunUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunUuids

`func (o *HyperflexTargetAllOf) SetLunUuids(v []string)`

SetLunUuids sets LunUuids field to given value.

### HasLunUuids

`func (o *HyperflexTargetAllOf) HasLunUuids() bool`

HasLunUuids returns a boolean if a field has been set.

### SetLunUuidsNil

`func (o *HyperflexTargetAllOf) SetLunUuidsNil(b bool)`

 SetLunUuidsNil sets the value for LunUuids to be an explicit nil

### UnsetLunUuids
`func (o *HyperflexTargetAllOf) UnsetLunUuids()`

UnsetLunUuids ensures that no value is present for LunUuids, not even an explicit nil
### GetNumActiveInitiators

`func (o *HyperflexTargetAllOf) GetNumActiveInitiators() int64`

GetNumActiveInitiators returns the NumActiveInitiators field if non-nil, zero value otherwise.

### GetNumActiveInitiatorsOk

`func (o *HyperflexTargetAllOf) GetNumActiveInitiatorsOk() (*int64, bool)`

GetNumActiveInitiatorsOk returns a tuple with the NumActiveInitiators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveInitiators

`func (o *HyperflexTargetAllOf) SetNumActiveInitiators(v int64)`

SetNumActiveInitiators sets NumActiveInitiators field to given value.

### HasNumActiveInitiators

`func (o *HyperflexTargetAllOf) HasNumActiveInitiators() bool`

HasNumActiveInitiators returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexTargetAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexTargetAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexTargetAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexTargetAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexTargetAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexTargetAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexTargetAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexTargetAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexTargetAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexTargetAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexTargetAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexTargetAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetInitiatorGroups

`func (o *HyperflexTargetAllOf) GetInitiatorGroups() []HyperflexInitiatorGroupRelationship`

GetInitiatorGroups returns the InitiatorGroups field if non-nil, zero value otherwise.

### GetInitiatorGroupsOk

`func (o *HyperflexTargetAllOf) GetInitiatorGroupsOk() (*[]HyperflexInitiatorGroupRelationship, bool)`

GetInitiatorGroupsOk returns a tuple with the InitiatorGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorGroups

`func (o *HyperflexTargetAllOf) SetInitiatorGroups(v []HyperflexInitiatorGroupRelationship)`

SetInitiatorGroups sets InitiatorGroups field to given value.

### HasInitiatorGroups

`func (o *HyperflexTargetAllOf) HasInitiatorGroups() bool`

HasInitiatorGroups returns a boolean if a field has been set.

### SetInitiatorGroupsNil

`func (o *HyperflexTargetAllOf) SetInitiatorGroupsNil(b bool)`

 SetInitiatorGroupsNil sets the value for InitiatorGroups to be an explicit nil

### UnsetInitiatorGroups
`func (o *HyperflexTargetAllOf) UnsetInitiatorGroups()`

UnsetInitiatorGroups ensures that no value is present for InitiatorGroups, not even an explicit nil
### GetLuns

`func (o *HyperflexTargetAllOf) GetLuns() []HyperflexLunRelationship`

GetLuns returns the Luns field if non-nil, zero value otherwise.

### GetLunsOk

`func (o *HyperflexTargetAllOf) GetLunsOk() (*[]HyperflexLunRelationship, bool)`

GetLunsOk returns a tuple with the Luns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLuns

`func (o *HyperflexTargetAllOf) SetLuns(v []HyperflexLunRelationship)`

SetLuns sets Luns field to given value.

### HasLuns

`func (o *HyperflexTargetAllOf) HasLuns() bool`

HasLuns returns a boolean if a field has been set.

### SetLunsNil

`func (o *HyperflexTargetAllOf) SetLunsNil(b bool)`

 SetLunsNil sets the value for Luns to be an explicit nil

### UnsetLuns
`func (o *HyperflexTargetAllOf) UnsetLuns()`

UnsetLuns ensures that no value is present for Luns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


