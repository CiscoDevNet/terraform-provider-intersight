# BlueprintCleanupOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.CleanupOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.CleanupOperation"]
**GeneratedObjects** | Pointer to [**[]BlueprintGeneratedObjectOperationTarget**](BlueprintGeneratedObjectOperationTarget.md) |  | [optional] 
**Operation** | Pointer to **string** | The operation to be performed after the service item deployment completes. Currently, only the &#39;Delete&#39; operation is supported. * &#x60;Update&#x60; - The operation is to update the object via an HTTP PATCH operation. * &#x60;Delete&#x60; - The operation is to immediately delete the object via an HTTP DELETE operation. * &#x60;MarkDeleted&#x60; - The object is marked as deleted. This adds a tag to the object indicating that it should be deleted as part of the object cleanup task during workload deployment. | [optional] [default to "Update"]

## Methods

### NewBlueprintCleanupOperation

`func NewBlueprintCleanupOperation(classId string, objectType string, ) *BlueprintCleanupOperation`

NewBlueprintCleanupOperation instantiates a new BlueprintCleanupOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCleanupOperationWithDefaults

`func NewBlueprintCleanupOperationWithDefaults() *BlueprintCleanupOperation`

NewBlueprintCleanupOperationWithDefaults instantiates a new BlueprintCleanupOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintCleanupOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintCleanupOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintCleanupOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintCleanupOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintCleanupOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintCleanupOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGeneratedObjects

`func (o *BlueprintCleanupOperation) GetGeneratedObjects() []BlueprintGeneratedObjectOperationTarget`

GetGeneratedObjects returns the GeneratedObjects field if non-nil, zero value otherwise.

### GetGeneratedObjectsOk

`func (o *BlueprintCleanupOperation) GetGeneratedObjectsOk() (*[]BlueprintGeneratedObjectOperationTarget, bool)`

GetGeneratedObjectsOk returns a tuple with the GeneratedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedObjects

`func (o *BlueprintCleanupOperation) SetGeneratedObjects(v []BlueprintGeneratedObjectOperationTarget)`

SetGeneratedObjects sets GeneratedObjects field to given value.

### HasGeneratedObjects

`func (o *BlueprintCleanupOperation) HasGeneratedObjects() bool`

HasGeneratedObjects returns a boolean if a field has been set.

### SetGeneratedObjectsNil

`func (o *BlueprintCleanupOperation) SetGeneratedObjectsNil(b bool)`

 SetGeneratedObjectsNil sets the value for GeneratedObjects to be an explicit nil

### UnsetGeneratedObjects
`func (o *BlueprintCleanupOperation) UnsetGeneratedObjects()`

UnsetGeneratedObjects ensures that no value is present for GeneratedObjects, not even an explicit nil
### GetOperation

`func (o *BlueprintCleanupOperation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BlueprintCleanupOperation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BlueprintCleanupOperation) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BlueprintCleanupOperation) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


