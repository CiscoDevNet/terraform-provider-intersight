# ServicenowIncidentDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "servicenow.IncidentDoc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "servicenow.IncidentDoc"]
**ExampleValue** | Pointer to **string** | Example value for Incident property. | [optional] 
**InternalType** | Pointer to **string** | Internal type for Incident property. | [optional] 
**Label** | Pointer to **string** | Label for Incident property. | [optional] 
**PropertyName** | Pointer to **string** | Name for Incident property. | [optional] 

## Methods

### NewServicenowIncidentDoc

`func NewServicenowIncidentDoc(classId string, objectType string, ) *ServicenowIncidentDoc`

NewServicenowIncidentDoc instantiates a new ServicenowIncidentDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicenowIncidentDocWithDefaults

`func NewServicenowIncidentDocWithDefaults() *ServicenowIncidentDoc`

NewServicenowIncidentDocWithDefaults instantiates a new ServicenowIncidentDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServicenowIncidentDoc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServicenowIncidentDoc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServicenowIncidentDoc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServicenowIncidentDoc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServicenowIncidentDoc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServicenowIncidentDoc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExampleValue

`func (o *ServicenowIncidentDoc) GetExampleValue() string`

GetExampleValue returns the ExampleValue field if non-nil, zero value otherwise.

### GetExampleValueOk

`func (o *ServicenowIncidentDoc) GetExampleValueOk() (*string, bool)`

GetExampleValueOk returns a tuple with the ExampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleValue

`func (o *ServicenowIncidentDoc) SetExampleValue(v string)`

SetExampleValue sets ExampleValue field to given value.

### HasExampleValue

`func (o *ServicenowIncidentDoc) HasExampleValue() bool`

HasExampleValue returns a boolean if a field has been set.

### GetInternalType

`func (o *ServicenowIncidentDoc) GetInternalType() string`

GetInternalType returns the InternalType field if non-nil, zero value otherwise.

### GetInternalTypeOk

`func (o *ServicenowIncidentDoc) GetInternalTypeOk() (*string, bool)`

GetInternalTypeOk returns a tuple with the InternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalType

`func (o *ServicenowIncidentDoc) SetInternalType(v string)`

SetInternalType sets InternalType field to given value.

### HasInternalType

`func (o *ServicenowIncidentDoc) HasInternalType() bool`

HasInternalType returns a boolean if a field has been set.

### GetLabel

`func (o *ServicenowIncidentDoc) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServicenowIncidentDoc) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServicenowIncidentDoc) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServicenowIncidentDoc) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPropertyName

`func (o *ServicenowIncidentDoc) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ServicenowIncidentDoc) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ServicenowIncidentDoc) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *ServicenowIncidentDoc) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


