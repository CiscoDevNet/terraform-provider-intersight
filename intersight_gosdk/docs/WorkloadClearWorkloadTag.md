# WorkloadClearWorkloadTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.ClearWorkloadTag"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.ClearWorkloadTag"]
**Objects** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkloadClearWorkloadTag

`func NewWorkloadClearWorkloadTag(classId string, objectType string, ) *WorkloadClearWorkloadTag`

NewWorkloadClearWorkloadTag instantiates a new WorkloadClearWorkloadTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadClearWorkloadTagWithDefaults

`func NewWorkloadClearWorkloadTagWithDefaults() *WorkloadClearWorkloadTag`

NewWorkloadClearWorkloadTagWithDefaults instantiates a new WorkloadClearWorkloadTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadClearWorkloadTag) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadClearWorkloadTag) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadClearWorkloadTag) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadClearWorkloadTag) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadClearWorkloadTag) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadClearWorkloadTag) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjects

`func (o *WorkloadClearWorkloadTag) GetObjects() []MoMoRef`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *WorkloadClearWorkloadTag) GetObjectsOk() (*[]MoMoRef, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *WorkloadClearWorkloadTag) SetObjects(v []MoMoRef)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *WorkloadClearWorkloadTag) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *WorkloadClearWorkloadTag) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *WorkloadClearWorkloadTag) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


