# WorkloadGeneratedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.GeneratedObject"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.GeneratedObject"]
**Name** | Pointer to **string** | The name of the generated object as defined in the generated object definition of a blueprint. | [optional] [readonly] 
**Objects** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkloadGeneratedObject

`func NewWorkloadGeneratedObject(classId string, objectType string, ) *WorkloadGeneratedObject`

NewWorkloadGeneratedObject instantiates a new WorkloadGeneratedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadGeneratedObjectWithDefaults

`func NewWorkloadGeneratedObjectWithDefaults() *WorkloadGeneratedObject`

NewWorkloadGeneratedObjectWithDefaults instantiates a new WorkloadGeneratedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadGeneratedObject) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadGeneratedObject) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadGeneratedObject) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadGeneratedObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadGeneratedObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadGeneratedObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *WorkloadGeneratedObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadGeneratedObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadGeneratedObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadGeneratedObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjects

`func (o *WorkloadGeneratedObject) GetObjects() []MoMoRef`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *WorkloadGeneratedObject) GetObjectsOk() (*[]MoMoRef, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *WorkloadGeneratedObject) SetObjects(v []MoMoRef)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *WorkloadGeneratedObject) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *WorkloadGeneratedObject) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *WorkloadGeneratedObject) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


