# HyperflexAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **bool** |  | [optional] [readonly] 
**AcknowledgedBy** | Pointer to **string** |  | [optional] [readonly] 
**AcknowledgedTime** | Pointer to **int64** |  | [optional] [readonly] 
**AcknowledgedTimeAsUtc** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**EntityData** | Pointer to **string** |  | [optional] [readonly] 
**EntityName** | Pointer to **string** |  | [optional] [readonly] 
**EntityType** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**EntityUuId** | Pointer to **string** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**TriggeredTime** | Pointer to **int64** |  | [optional] [readonly] 
**TriggeredTimeAsUtc** | Pointer to **string** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexAlarm

`func NewHyperflexAlarm() *HyperflexAlarm`

NewHyperflexAlarm instantiates a new HyperflexAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAlarmWithDefaults

`func NewHyperflexAlarmWithDefaults() *HyperflexAlarm`

NewHyperflexAlarmWithDefaults instantiates a new HyperflexAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *HyperflexAlarm) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *HyperflexAlarm) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *HyperflexAlarm) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *HyperflexAlarm) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetAcknowledgedBy

`func (o *HyperflexAlarm) GetAcknowledgedBy() string`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *HyperflexAlarm) GetAcknowledgedByOk() (*string, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *HyperflexAlarm) SetAcknowledgedBy(v string)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *HyperflexAlarm) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### GetAcknowledgedTime

`func (o *HyperflexAlarm) GetAcknowledgedTime() int64`

GetAcknowledgedTime returns the AcknowledgedTime field if non-nil, zero value otherwise.

### GetAcknowledgedTimeOk

`func (o *HyperflexAlarm) GetAcknowledgedTimeOk() (*int64, bool)`

GetAcknowledgedTimeOk returns a tuple with the AcknowledgedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedTime

`func (o *HyperflexAlarm) SetAcknowledgedTime(v int64)`

SetAcknowledgedTime sets AcknowledgedTime field to given value.

### HasAcknowledgedTime

`func (o *HyperflexAlarm) HasAcknowledgedTime() bool`

HasAcknowledgedTime returns a boolean if a field has been set.

### GetAcknowledgedTimeAsUtc

`func (o *HyperflexAlarm) GetAcknowledgedTimeAsUtc() string`

GetAcknowledgedTimeAsUtc returns the AcknowledgedTimeAsUtc field if non-nil, zero value otherwise.

### GetAcknowledgedTimeAsUtcOk

`func (o *HyperflexAlarm) GetAcknowledgedTimeAsUtcOk() (*string, bool)`

GetAcknowledgedTimeAsUtcOk returns a tuple with the AcknowledgedTimeAsUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedTimeAsUtc

`func (o *HyperflexAlarm) SetAcknowledgedTimeAsUtc(v string)`

SetAcknowledgedTimeAsUtc sets AcknowledgedTimeAsUtc field to given value.

### HasAcknowledgedTimeAsUtc

`func (o *HyperflexAlarm) HasAcknowledgedTimeAsUtc() bool`

HasAcknowledgedTimeAsUtc returns a boolean if a field has been set.

### GetDescription

`func (o *HyperflexAlarm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexAlarm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexAlarm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexAlarm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityData

`func (o *HyperflexAlarm) GetEntityData() string`

GetEntityData returns the EntityData field if non-nil, zero value otherwise.

### GetEntityDataOk

`func (o *HyperflexAlarm) GetEntityDataOk() (*string, bool)`

GetEntityDataOk returns a tuple with the EntityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityData

`func (o *HyperflexAlarm) SetEntityData(v string)`

SetEntityData sets EntityData field to given value.

### HasEntityData

`func (o *HyperflexAlarm) HasEntityData() bool`

HasEntityData returns a boolean if a field has been set.

### GetEntityName

`func (o *HyperflexAlarm) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *HyperflexAlarm) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *HyperflexAlarm) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *HyperflexAlarm) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityType

`func (o *HyperflexAlarm) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HyperflexAlarm) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HyperflexAlarm) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *HyperflexAlarm) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityUuId

`func (o *HyperflexAlarm) GetEntityUuId() string`

GetEntityUuId returns the EntityUuId field if non-nil, zero value otherwise.

### GetEntityUuIdOk

`func (o *HyperflexAlarm) GetEntityUuIdOk() (*string, bool)`

GetEntityUuIdOk returns a tuple with the EntityUuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuId

`func (o *HyperflexAlarm) SetEntityUuId(v string)`

SetEntityUuId sets EntityUuId field to given value.

### HasEntityUuId

`func (o *HyperflexAlarm) HasEntityUuId() bool`

HasEntityUuId returns a boolean if a field has been set.

### GetMessage

`func (o *HyperflexAlarm) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HyperflexAlarm) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HyperflexAlarm) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HyperflexAlarm) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *HyperflexAlarm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexAlarm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexAlarm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexAlarm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexAlarm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexAlarm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexAlarm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexAlarm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggeredTime

`func (o *HyperflexAlarm) GetTriggeredTime() int64`

GetTriggeredTime returns the TriggeredTime field if non-nil, zero value otherwise.

### GetTriggeredTimeOk

`func (o *HyperflexAlarm) GetTriggeredTimeOk() (*int64, bool)`

GetTriggeredTimeOk returns a tuple with the TriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTime

`func (o *HyperflexAlarm) SetTriggeredTime(v int64)`

SetTriggeredTime sets TriggeredTime field to given value.

### HasTriggeredTime

`func (o *HyperflexAlarm) HasTriggeredTime() bool`

HasTriggeredTime returns a boolean if a field has been set.

### GetTriggeredTimeAsUtc

`func (o *HyperflexAlarm) GetTriggeredTimeAsUtc() string`

GetTriggeredTimeAsUtc returns the TriggeredTimeAsUtc field if non-nil, zero value otherwise.

### GetTriggeredTimeAsUtcOk

`func (o *HyperflexAlarm) GetTriggeredTimeAsUtcOk() (*string, bool)`

GetTriggeredTimeAsUtcOk returns a tuple with the TriggeredTimeAsUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTimeAsUtc

`func (o *HyperflexAlarm) SetTriggeredTimeAsUtc(v string)`

SetTriggeredTimeAsUtc sets TriggeredTimeAsUtc field to given value.

### HasTriggeredTimeAsUtc

`func (o *HyperflexAlarm) HasTriggeredTimeAsUtc() bool`

HasTriggeredTimeAsUtc returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexAlarm) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexAlarm) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexAlarm) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexAlarm) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexAlarm) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexAlarm) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexAlarm) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexAlarm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


