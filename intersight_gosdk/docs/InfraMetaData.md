# InfraMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "infra.MetaData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "infra.MetaData"]
**Name** | Pointer to **string** | Name of the meta property which identifies a specific resource. | [optional] [readonly] 
**Value** | Pointer to **string** | Value of the meta property which identifies a specific resource. | [optional] [readonly] 

## Methods

### NewInfraMetaData

`func NewInfraMetaData(classId string, objectType string, ) *InfraMetaData`

NewInfraMetaData instantiates a new InfraMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraMetaDataWithDefaults

`func NewInfraMetaDataWithDefaults() *InfraMetaData`

NewInfraMetaDataWithDefaults instantiates a new InfraMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InfraMetaData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InfraMetaData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InfraMetaData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InfraMetaData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InfraMetaData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InfraMetaData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *InfraMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfraMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfraMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfraMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *InfraMetaData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InfraMetaData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InfraMetaData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InfraMetaData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


