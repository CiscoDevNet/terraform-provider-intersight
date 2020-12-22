# ContentComplexTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "content.ComplexType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "content.ComplexType"]
**Name** | Pointer to **string** | The unique name of this complex type within the grammar specification. | [optional] 
**Parameters** | Pointer to [**[]ContentBaseParameter**](ContentBaseParameter.md) |  | [optional] 

## Methods

### NewContentComplexTypeAllOf

`func NewContentComplexTypeAllOf(classId string, objectType string, ) *ContentComplexTypeAllOf`

NewContentComplexTypeAllOf instantiates a new ContentComplexTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentComplexTypeAllOfWithDefaults

`func NewContentComplexTypeAllOfWithDefaults() *ContentComplexTypeAllOf`

NewContentComplexTypeAllOfWithDefaults instantiates a new ContentComplexTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ContentComplexTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ContentComplexTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ContentComplexTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ContentComplexTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ContentComplexTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ContentComplexTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *ContentComplexTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentComplexTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentComplexTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentComplexTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ContentComplexTypeAllOf) GetParameters() []ContentBaseParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentComplexTypeAllOf) GetParametersOk() (*[]ContentBaseParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentComplexTypeAllOf) SetParameters(v []ContentBaseParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentComplexTypeAllOf) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ContentComplexTypeAllOf) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ContentComplexTypeAllOf) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


