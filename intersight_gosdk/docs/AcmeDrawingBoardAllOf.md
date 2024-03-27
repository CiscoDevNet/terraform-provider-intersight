# AcmeDrawingBoardAllOf

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

### NewAcmeDrawingBoardAllOf

`func NewAcmeDrawingBoardAllOf(classId string, objectType string, ) *AcmeDrawingBoardAllOf`

NewAcmeDrawingBoardAllOf instantiates a new AcmeDrawingBoardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeDrawingBoardAllOfWithDefaults

`func NewAcmeDrawingBoardAllOfWithDefaults() *AcmeDrawingBoardAllOf`

NewAcmeDrawingBoardAllOfWithDefaults instantiates a new AcmeDrawingBoardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AcmeDrawingBoardAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AcmeDrawingBoardAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AcmeDrawingBoardAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AcmeDrawingBoardAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AcmeDrawingBoardAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AcmeDrawingBoardAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescr

`func (o *AcmeDrawingBoardAllOf) GetDescr() string`

GetDescr returns the Descr field if non-nil, zero value otherwise.

### GetDescrOk

`func (o *AcmeDrawingBoardAllOf) GetDescrOk() (*string, bool)`

GetDescrOk returns a tuple with the Descr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescr

`func (o *AcmeDrawingBoardAllOf) SetDescr(v string)`

SetDescr sets Descr field to given value.

### HasDescr

`func (o *AcmeDrawingBoardAllOf) HasDescr() bool`

HasDescr returns a boolean if a field has been set.

### GetName

`func (o *AcmeDrawingBoardAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeDrawingBoardAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeDrawingBoardAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcmeDrawingBoardAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArrayOfAssociatedObjects

`func (o *AcmeDrawingBoardAllOf) GetArrayOfAssociatedObjects() []MoBaseMoRelationship`

GetArrayOfAssociatedObjects returns the ArrayOfAssociatedObjects field if non-nil, zero value otherwise.

### GetArrayOfAssociatedObjectsOk

`func (o *AcmeDrawingBoardAllOf) GetArrayOfAssociatedObjectsOk() (*[]MoBaseMoRelationship, bool)`

GetArrayOfAssociatedObjectsOk returns a tuple with the ArrayOfAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayOfAssociatedObjects

`func (o *AcmeDrawingBoardAllOf) SetArrayOfAssociatedObjects(v []MoBaseMoRelationship)`

SetArrayOfAssociatedObjects sets ArrayOfAssociatedObjects field to given value.

### HasArrayOfAssociatedObjects

`func (o *AcmeDrawingBoardAllOf) HasArrayOfAssociatedObjects() bool`

HasArrayOfAssociatedObjects returns a boolean if a field has been set.

### SetArrayOfAssociatedObjectsNil

`func (o *AcmeDrawingBoardAllOf) SetArrayOfAssociatedObjectsNil(b bool)`

 SetArrayOfAssociatedObjectsNil sets the value for ArrayOfAssociatedObjects to be an explicit nil

### UnsetArrayOfAssociatedObjects
`func (o *AcmeDrawingBoardAllOf) UnsetArrayOfAssociatedObjects()`

UnsetArrayOfAssociatedObjects ensures that no value is present for ArrayOfAssociatedObjects, not even an explicit nil
### GetAssociatedObject

`func (o *AcmeDrawingBoardAllOf) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *AcmeDrawingBoardAllOf) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *AcmeDrawingBoardAllOf) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *AcmeDrawingBoardAllOf) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


