# HyperflexInitiatorGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.InitiatorGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.InitiatorGroup"]
**InitiatorCount** | Pointer to **int64** | Count of initiators in the iSCSI initiator group. | [optional] [readonly] 
**Initiators** | Pointer to [**[]StorageHyperFlexIscsiInitiator**](StorageHyperFlexIscsiInitiator.md) |  | [optional] 
**InventorySource** | Pointer to **string** | The source of the iSCSI initiator group inventory. * &#x60;NOT_APPLICABLE&#x60; - The source of the HyperFlex inventory is not applicable. * &#x60;ONLINE&#x60; - The source of the HyperFlex inventory is online. * &#x60;OFFLINE&#x60; - The source of the HyperFlex inventory is offline. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**TargetUuids** | Pointer to **[]string** |  | [optional] 
**Uuid** | Pointer to **string** | UUID of the HyperFlex iSCSI initiator group. | [optional] [readonly] 
**Version** | Pointer to **int64** | Version of the iSCSI initiator group. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**Targets** | Pointer to [**[]HyperflexTargetRelationship**](HyperflexTargetRelationship.md) | An array of relationships to hyperflexTarget resources. | [optional] [readonly] 

## Methods

### NewHyperflexInitiatorGroupAllOf

`func NewHyperflexInitiatorGroupAllOf(classId string, objectType string, ) *HyperflexInitiatorGroupAllOf`

NewHyperflexInitiatorGroupAllOf instantiates a new HyperflexInitiatorGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexInitiatorGroupAllOfWithDefaults

`func NewHyperflexInitiatorGroupAllOfWithDefaults() *HyperflexInitiatorGroupAllOf`

NewHyperflexInitiatorGroupAllOfWithDefaults instantiates a new HyperflexInitiatorGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexInitiatorGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexInitiatorGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexInitiatorGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexInitiatorGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexInitiatorGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexInitiatorGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInitiatorCount

`func (o *HyperflexInitiatorGroupAllOf) GetInitiatorCount() int64`

GetInitiatorCount returns the InitiatorCount field if non-nil, zero value otherwise.

### GetInitiatorCountOk

`func (o *HyperflexInitiatorGroupAllOf) GetInitiatorCountOk() (*int64, bool)`

GetInitiatorCountOk returns a tuple with the InitiatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCount

`func (o *HyperflexInitiatorGroupAllOf) SetInitiatorCount(v int64)`

SetInitiatorCount sets InitiatorCount field to given value.

### HasInitiatorCount

`func (o *HyperflexInitiatorGroupAllOf) HasInitiatorCount() bool`

HasInitiatorCount returns a boolean if a field has been set.

### GetInitiators

`func (o *HyperflexInitiatorGroupAllOf) GetInitiators() []StorageHyperFlexIscsiInitiator`

GetInitiators returns the Initiators field if non-nil, zero value otherwise.

### GetInitiatorsOk

`func (o *HyperflexInitiatorGroupAllOf) GetInitiatorsOk() (*[]StorageHyperFlexIscsiInitiator, bool)`

GetInitiatorsOk returns a tuple with the Initiators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiators

`func (o *HyperflexInitiatorGroupAllOf) SetInitiators(v []StorageHyperFlexIscsiInitiator)`

SetInitiators sets Initiators field to given value.

### HasInitiators

`func (o *HyperflexInitiatorGroupAllOf) HasInitiators() bool`

HasInitiators returns a boolean if a field has been set.

### SetInitiatorsNil

`func (o *HyperflexInitiatorGroupAllOf) SetInitiatorsNil(b bool)`

 SetInitiatorsNil sets the value for Initiators to be an explicit nil

### UnsetInitiators
`func (o *HyperflexInitiatorGroupAllOf) UnsetInitiators()`

UnsetInitiators ensures that no value is present for Initiators, not even an explicit nil
### GetInventorySource

`func (o *HyperflexInitiatorGroupAllOf) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *HyperflexInitiatorGroupAllOf) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *HyperflexInitiatorGroupAllOf) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *HyperflexInitiatorGroupAllOf) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetTargetUuids

`func (o *HyperflexInitiatorGroupAllOf) GetTargetUuids() []string`

GetTargetUuids returns the TargetUuids field if non-nil, zero value otherwise.

### GetTargetUuidsOk

`func (o *HyperflexInitiatorGroupAllOf) GetTargetUuidsOk() (*[]string, bool)`

GetTargetUuidsOk returns a tuple with the TargetUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuids

`func (o *HyperflexInitiatorGroupAllOf) SetTargetUuids(v []string)`

SetTargetUuids sets TargetUuids field to given value.

### HasTargetUuids

`func (o *HyperflexInitiatorGroupAllOf) HasTargetUuids() bool`

HasTargetUuids returns a boolean if a field has been set.

### SetTargetUuidsNil

`func (o *HyperflexInitiatorGroupAllOf) SetTargetUuidsNil(b bool)`

 SetTargetUuidsNil sets the value for TargetUuids to be an explicit nil

### UnsetTargetUuids
`func (o *HyperflexInitiatorGroupAllOf) UnsetTargetUuids()`

UnsetTargetUuids ensures that no value is present for TargetUuids, not even an explicit nil
### GetUuid

`func (o *HyperflexInitiatorGroupAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexInitiatorGroupAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexInitiatorGroupAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexInitiatorGroupAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexInitiatorGroupAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexInitiatorGroupAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexInitiatorGroupAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexInitiatorGroupAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexInitiatorGroupAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexInitiatorGroupAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexInitiatorGroupAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexInitiatorGroupAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTargets

`func (o *HyperflexInitiatorGroupAllOf) GetTargets() []HyperflexTargetRelationship`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *HyperflexInitiatorGroupAllOf) GetTargetsOk() (*[]HyperflexTargetRelationship, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *HyperflexInitiatorGroupAllOf) SetTargets(v []HyperflexTargetRelationship)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *HyperflexInitiatorGroupAllOf) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *HyperflexInitiatorGroupAllOf) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *HyperflexInitiatorGroupAllOf) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


