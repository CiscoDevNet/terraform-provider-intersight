# AcmeDrawingBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "acme.DrawingBoard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "acme.DrawingBoard"]
**Descr** | Pointer to **string** | The description of the drawing board. | [optional] 
**Name** | Pointer to **string** | The name of the drawing board. | [optional] 
**ArrayOfAssociatedObjects** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] 
**AssociatedObject** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewAcmeDrawingBoard

`func NewAcmeDrawingBoard(classId string, objectType string, ) *AcmeDrawingBoard`

NewAcmeDrawingBoard instantiates a new AcmeDrawingBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeDrawingBoardWithDefaults

`func NewAcmeDrawingBoardWithDefaults() *AcmeDrawingBoard`

NewAcmeDrawingBoardWithDefaults instantiates a new AcmeDrawingBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AcmeDrawingBoard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AcmeDrawingBoard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AcmeDrawingBoard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AcmeDrawingBoard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AcmeDrawingBoard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AcmeDrawingBoard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescr

`func (o *AcmeDrawingBoard) GetDescr() string`

GetDescr returns the Descr field if non-nil, zero value otherwise.

### GetDescrOk

`func (o *AcmeDrawingBoard) GetDescrOk() (*string, bool)`

GetDescrOk returns a tuple with the Descr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescr

`func (o *AcmeDrawingBoard) SetDescr(v string)`

SetDescr sets Descr field to given value.

### HasDescr

`func (o *AcmeDrawingBoard) HasDescr() bool`

HasDescr returns a boolean if a field has been set.

### GetName

`func (o *AcmeDrawingBoard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeDrawingBoard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeDrawingBoard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcmeDrawingBoard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArrayOfAssociatedObjects

`func (o *AcmeDrawingBoard) GetArrayOfAssociatedObjects() []MoBaseMoRelationship`

GetArrayOfAssociatedObjects returns the ArrayOfAssociatedObjects field if non-nil, zero value otherwise.

### GetArrayOfAssociatedObjectsOk

`func (o *AcmeDrawingBoard) GetArrayOfAssociatedObjectsOk() (*[]MoBaseMoRelationship, bool)`

GetArrayOfAssociatedObjectsOk returns a tuple with the ArrayOfAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayOfAssociatedObjects

`func (o *AcmeDrawingBoard) SetArrayOfAssociatedObjects(v []MoBaseMoRelationship)`

SetArrayOfAssociatedObjects sets ArrayOfAssociatedObjects field to given value.

### HasArrayOfAssociatedObjects

`func (o *AcmeDrawingBoard) HasArrayOfAssociatedObjects() bool`

HasArrayOfAssociatedObjects returns a boolean if a field has been set.

### SetArrayOfAssociatedObjectsNil

`func (o *AcmeDrawingBoard) SetArrayOfAssociatedObjectsNil(b bool)`

 SetArrayOfAssociatedObjectsNil sets the value for ArrayOfAssociatedObjects to be an explicit nil

### UnsetArrayOfAssociatedObjects
`func (o *AcmeDrawingBoard) UnsetArrayOfAssociatedObjects()`

UnsetArrayOfAssociatedObjects ensures that no value is present for ArrayOfAssociatedObjects, not even an explicit nil
### GetAssociatedObject

`func (o *AcmeDrawingBoard) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *AcmeDrawingBoard) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *AcmeDrawingBoard) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *AcmeDrawingBoard) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


