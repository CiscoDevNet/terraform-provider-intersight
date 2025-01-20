# HciFaultToleranceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.FaultToleranceState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.FaultToleranceState"]
**CurrentMaxFaultTolerance** | Pointer to **int32** | The maximum fault tolerance level that is supported currently. | [optional] [readonly] 
**DesiredMaxFaultTolerance** | Pointer to **int32** | The maximum fault tolerance level that is desired. | [optional] [readonly] 
**DomainAwarenessLevel** | Pointer to **string** | Domain awareness level corresponds to unit of cluster group. Part of payload for both cluster create &amp; update operations. | [optional] [readonly] 

## Methods

### NewHciFaultToleranceState

`func NewHciFaultToleranceState(classId string, objectType string, ) *HciFaultToleranceState`

NewHciFaultToleranceState instantiates a new HciFaultToleranceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciFaultToleranceStateWithDefaults

`func NewHciFaultToleranceStateWithDefaults() *HciFaultToleranceState`

NewHciFaultToleranceStateWithDefaults instantiates a new HciFaultToleranceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciFaultToleranceState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciFaultToleranceState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciFaultToleranceState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciFaultToleranceState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciFaultToleranceState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciFaultToleranceState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentMaxFaultTolerance

`func (o *HciFaultToleranceState) GetCurrentMaxFaultTolerance() int32`

GetCurrentMaxFaultTolerance returns the CurrentMaxFaultTolerance field if non-nil, zero value otherwise.

### GetCurrentMaxFaultToleranceOk

`func (o *HciFaultToleranceState) GetCurrentMaxFaultToleranceOk() (*int32, bool)`

GetCurrentMaxFaultToleranceOk returns a tuple with the CurrentMaxFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMaxFaultTolerance

`func (o *HciFaultToleranceState) SetCurrentMaxFaultTolerance(v int32)`

SetCurrentMaxFaultTolerance sets CurrentMaxFaultTolerance field to given value.

### HasCurrentMaxFaultTolerance

`func (o *HciFaultToleranceState) HasCurrentMaxFaultTolerance() bool`

HasCurrentMaxFaultTolerance returns a boolean if a field has been set.

### GetDesiredMaxFaultTolerance

`func (o *HciFaultToleranceState) GetDesiredMaxFaultTolerance() int32`

GetDesiredMaxFaultTolerance returns the DesiredMaxFaultTolerance field if non-nil, zero value otherwise.

### GetDesiredMaxFaultToleranceOk

`func (o *HciFaultToleranceState) GetDesiredMaxFaultToleranceOk() (*int32, bool)`

GetDesiredMaxFaultToleranceOk returns a tuple with the DesiredMaxFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredMaxFaultTolerance

`func (o *HciFaultToleranceState) SetDesiredMaxFaultTolerance(v int32)`

SetDesiredMaxFaultTolerance sets DesiredMaxFaultTolerance field to given value.

### HasDesiredMaxFaultTolerance

`func (o *HciFaultToleranceState) HasDesiredMaxFaultTolerance() bool`

HasDesiredMaxFaultTolerance returns a boolean if a field has been set.

### GetDomainAwarenessLevel

`func (o *HciFaultToleranceState) GetDomainAwarenessLevel() string`

GetDomainAwarenessLevel returns the DomainAwarenessLevel field if non-nil, zero value otherwise.

### GetDomainAwarenessLevelOk

`func (o *HciFaultToleranceState) GetDomainAwarenessLevelOk() (*string, bool)`

GetDomainAwarenessLevelOk returns a tuple with the DomainAwarenessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAwarenessLevel

`func (o *HciFaultToleranceState) SetDomainAwarenessLevel(v string)`

SetDomainAwarenessLevel sets DomainAwarenessLevel field to given value.

### HasDomainAwarenessLevel

`func (o *HciFaultToleranceState) HasDomainAwarenessLevel() bool`

HasDomainAwarenessLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


