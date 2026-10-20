# ServicenowChangeRequestDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "servicenow.ChangeRequestDoc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "servicenow.ChangeRequestDoc"]
**ExampleValue** | Pointer to **string** | Example value for Change request property. | [optional] 
**InternalType** | Pointer to **string** | Internal type for Change request property. | [optional] 
**Label** | Pointer to **string** | Label for Change request property. | [optional] 
**PropertyName** | Pointer to **string** | Name for Change request property. | [optional] 

## Methods

### NewServicenowChangeRequestDoc

`func NewServicenowChangeRequestDoc(classId string, objectType string, ) *ServicenowChangeRequestDoc`

NewServicenowChangeRequestDoc instantiates a new ServicenowChangeRequestDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicenowChangeRequestDocWithDefaults

`func NewServicenowChangeRequestDocWithDefaults() *ServicenowChangeRequestDoc`

NewServicenowChangeRequestDocWithDefaults instantiates a new ServicenowChangeRequestDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServicenowChangeRequestDoc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServicenowChangeRequestDoc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServicenowChangeRequestDoc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServicenowChangeRequestDoc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServicenowChangeRequestDoc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServicenowChangeRequestDoc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExampleValue

`func (o *ServicenowChangeRequestDoc) GetExampleValue() string`

GetExampleValue returns the ExampleValue field if non-nil, zero value otherwise.

### GetExampleValueOk

`func (o *ServicenowChangeRequestDoc) GetExampleValueOk() (*string, bool)`

GetExampleValueOk returns a tuple with the ExampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleValue

`func (o *ServicenowChangeRequestDoc) SetExampleValue(v string)`

SetExampleValue sets ExampleValue field to given value.

### HasExampleValue

`func (o *ServicenowChangeRequestDoc) HasExampleValue() bool`

HasExampleValue returns a boolean if a field has been set.

### GetInternalType

`func (o *ServicenowChangeRequestDoc) GetInternalType() string`

GetInternalType returns the InternalType field if non-nil, zero value otherwise.

### GetInternalTypeOk

`func (o *ServicenowChangeRequestDoc) GetInternalTypeOk() (*string, bool)`

GetInternalTypeOk returns a tuple with the InternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalType

`func (o *ServicenowChangeRequestDoc) SetInternalType(v string)`

SetInternalType sets InternalType field to given value.

### HasInternalType

`func (o *ServicenowChangeRequestDoc) HasInternalType() bool`

HasInternalType returns a boolean if a field has been set.

### GetLabel

`func (o *ServicenowChangeRequestDoc) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServicenowChangeRequestDoc) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServicenowChangeRequestDoc) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServicenowChangeRequestDoc) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPropertyName

`func (o *ServicenowChangeRequestDoc) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ServicenowChangeRequestDoc) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ServicenowChangeRequestDoc) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *ServicenowChangeRequestDoc) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


