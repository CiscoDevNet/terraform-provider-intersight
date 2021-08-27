# HyperflexReplicationClusterReferenceToSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReplicationClusterReferenceToSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReplicationClusterReferenceToSchedule"]
**Schedule** | Pointer to [**NullableHyperflexReplicationSchedule**](HyperflexReplicationSchedule.md) |  | [optional] 
**TargetClusterEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 

## Methods

### NewHyperflexReplicationClusterReferenceToSchedule

`func NewHyperflexReplicationClusterReferenceToSchedule(classId string, objectType string, ) *HyperflexReplicationClusterReferenceToSchedule`

NewHyperflexReplicationClusterReferenceToSchedule instantiates a new HyperflexReplicationClusterReferenceToSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReplicationClusterReferenceToScheduleWithDefaults

`func NewHyperflexReplicationClusterReferenceToScheduleWithDefaults() *HyperflexReplicationClusterReferenceToSchedule`

NewHyperflexReplicationClusterReferenceToScheduleWithDefaults instantiates a new HyperflexReplicationClusterReferenceToSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReplicationClusterReferenceToSchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReplicationClusterReferenceToSchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSchedule

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetSchedule() HyperflexReplicationSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetScheduleOk() (*HyperflexReplicationSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HyperflexReplicationClusterReferenceToSchedule) SetSchedule(v HyperflexReplicationSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HyperflexReplicationClusterReferenceToSchedule) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *HyperflexReplicationClusterReferenceToSchedule) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *HyperflexReplicationClusterReferenceToSchedule) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTargetClusterEntityReference

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetTargetClusterEntityReference() HyperflexEntityReference`

GetTargetClusterEntityReference returns the TargetClusterEntityReference field if non-nil, zero value otherwise.

### GetTargetClusterEntityReferenceOk

`func (o *HyperflexReplicationClusterReferenceToSchedule) GetTargetClusterEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetTargetClusterEntityReferenceOk returns a tuple with the TargetClusterEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterEntityReference

`func (o *HyperflexReplicationClusterReferenceToSchedule) SetTargetClusterEntityReference(v HyperflexEntityReference)`

SetTargetClusterEntityReference sets TargetClusterEntityReference field to given value.

### HasTargetClusterEntityReference

`func (o *HyperflexReplicationClusterReferenceToSchedule) HasTargetClusterEntityReference() bool`

HasTargetClusterEntityReference returns a boolean if a field has been set.

### SetTargetClusterEntityReferenceNil

`func (o *HyperflexReplicationClusterReferenceToSchedule) SetTargetClusterEntityReferenceNil(b bool)`

 SetTargetClusterEntityReferenceNil sets the value for TargetClusterEntityReference to be an explicit nil

### UnsetTargetClusterEntityReference
`func (o *HyperflexReplicationClusterReferenceToSchedule) UnsetTargetClusterEntityReference()`

UnsetTargetClusterEntityReference ensures that no value is present for TargetClusterEntityReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


