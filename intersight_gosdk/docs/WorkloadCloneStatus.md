# WorkloadCloneStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.CloneStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.CloneStatus"]
**BlueprintRefName** | Pointer to **string** | The reference name of the blueprint for which the clone operation is being tracked. | [optional] [readonly] 
**CloneStatusCounter** | Pointer to **int64** | The count of clone operations performed for the input. | [optional] [readonly] 
**CloneStatusEntry** | Pointer to [**[]WorkloadCloneStatusEntry**](WorkloadCloneStatusEntry.md) |  | [optional] 
**MoType** | Pointer to **string** | The managed object type of the input for which the clone operation is being tracked. It is used to identify the type of input data that is being cloned. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the input for which the clone operation is being tracked. | [optional] [readonly] 

## Methods

### NewWorkloadCloneStatus

`func NewWorkloadCloneStatus(classId string, objectType string, ) *WorkloadCloneStatus`

NewWorkloadCloneStatus instantiates a new WorkloadCloneStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadCloneStatusWithDefaults

`func NewWorkloadCloneStatusWithDefaults() *WorkloadCloneStatus`

NewWorkloadCloneStatusWithDefaults instantiates a new WorkloadCloneStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadCloneStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadCloneStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadCloneStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadCloneStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadCloneStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadCloneStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintRefName

`func (o *WorkloadCloneStatus) GetBlueprintRefName() string`

GetBlueprintRefName returns the BlueprintRefName field if non-nil, zero value otherwise.

### GetBlueprintRefNameOk

`func (o *WorkloadCloneStatus) GetBlueprintRefNameOk() (*string, bool)`

GetBlueprintRefNameOk returns a tuple with the BlueprintRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintRefName

`func (o *WorkloadCloneStatus) SetBlueprintRefName(v string)`

SetBlueprintRefName sets BlueprintRefName field to given value.

### HasBlueprintRefName

`func (o *WorkloadCloneStatus) HasBlueprintRefName() bool`

HasBlueprintRefName returns a boolean if a field has been set.

### GetCloneStatusCounter

`func (o *WorkloadCloneStatus) GetCloneStatusCounter() int64`

GetCloneStatusCounter returns the CloneStatusCounter field if non-nil, zero value otherwise.

### GetCloneStatusCounterOk

`func (o *WorkloadCloneStatus) GetCloneStatusCounterOk() (*int64, bool)`

GetCloneStatusCounterOk returns a tuple with the CloneStatusCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneStatusCounter

`func (o *WorkloadCloneStatus) SetCloneStatusCounter(v int64)`

SetCloneStatusCounter sets CloneStatusCounter field to given value.

### HasCloneStatusCounter

`func (o *WorkloadCloneStatus) HasCloneStatusCounter() bool`

HasCloneStatusCounter returns a boolean if a field has been set.

### GetCloneStatusEntry

`func (o *WorkloadCloneStatus) GetCloneStatusEntry() []WorkloadCloneStatusEntry`

GetCloneStatusEntry returns the CloneStatusEntry field if non-nil, zero value otherwise.

### GetCloneStatusEntryOk

`func (o *WorkloadCloneStatus) GetCloneStatusEntryOk() (*[]WorkloadCloneStatusEntry, bool)`

GetCloneStatusEntryOk returns a tuple with the CloneStatusEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneStatusEntry

`func (o *WorkloadCloneStatus) SetCloneStatusEntry(v []WorkloadCloneStatusEntry)`

SetCloneStatusEntry sets CloneStatusEntry field to given value.

### HasCloneStatusEntry

`func (o *WorkloadCloneStatus) HasCloneStatusEntry() bool`

HasCloneStatusEntry returns a boolean if a field has been set.

### SetCloneStatusEntryNil

`func (o *WorkloadCloneStatus) SetCloneStatusEntryNil(b bool)`

 SetCloneStatusEntryNil sets the value for CloneStatusEntry to be an explicit nil

### UnsetCloneStatusEntry
`func (o *WorkloadCloneStatus) UnsetCloneStatusEntry()`

UnsetCloneStatusEntry ensures that no value is present for CloneStatusEntry, not even an explicit nil
### GetMoType

`func (o *WorkloadCloneStatus) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *WorkloadCloneStatus) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *WorkloadCloneStatus) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *WorkloadCloneStatus) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetName

`func (o *WorkloadCloneStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadCloneStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadCloneStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadCloneStatus) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


