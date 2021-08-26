# HyperflexReplicationClusterReferenceToScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationClusterReferenceToSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationClusterReferenceToSchedule"]
**Schedule** | Pointer to [**NullableHyperflexReplicationSchedule**](HyperflexReplicationSchedule.md) |  | [optional] 
**TargetClusterEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 

## Methods

### NewHyperflexReplicationClusterReferenceToScheduleAllOf

`func NewHyperflexReplicationClusterReferenceToScheduleAllOf(classId string, objectType string, ) *HyperflexReplicationClusterReferenceToScheduleAllOf`

NewHyperflexReplicationClusterReferenceToScheduleAllOf instantiates a new HyperflexReplicationClusterReferenceToScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationClusterReferenceToScheduleAllOfWithDefaults

`func NewHyperflexReplicationClusterReferenceToScheduleAllOfWithDefaults() *HyperflexReplicationClusterReferenceToScheduleAllOf`

NewHyperflexReplicationClusterReferenceToScheduleAllOfWithDefaults instantiates a new HyperflexReplicationClusterReferenceToScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSchedule

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetSchedule() HyperflexReplicationSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetScheduleOk() (*HyperflexReplicationSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetSchedule(v HyperflexReplicationSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTargetClusterEntityReference

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetTargetClusterEntityReference() HyperflexEntityReference`

GetTargetClusterEntityReference returns the TargetClusterEntityReference field if non-nil, zero value otherwise.

### GetTargetClusterEntityReferenceOk

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetTargetClusterEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetTargetClusterEntityReferenceOk returns a tuple with the TargetClusterEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterEntityReference

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetTargetClusterEntityReference(v HyperflexEntityReference)`

SetTargetClusterEntityReference sets TargetClusterEntityReference field to given value.

### HasTargetClusterEntityReference

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) HasTargetClusterEntityReference() bool`

HasTargetClusterEntityReference returns a boolean if a field has been set.

### SetTargetClusterEntityReferenceNil

`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetTargetClusterEntityReferenceNil(b bool)`

 SetTargetClusterEntityReferenceNil sets the value for TargetClusterEntityReference to be an explicit nil

### UnsetTargetClusterEntityReference
`func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) UnsetTargetClusterEntityReference()`

UnsetTargetClusterEntityReference ensures that no value is present for TargetClusterEntityReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


