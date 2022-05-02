# ServiceitemHealthCheckErrorElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "serviceitem.HealthCheckErrorElement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "serviceitem.HealthCheckErrorElement"]
**ElementType** | Pointer to **string** | Type of the element of the service item. Examples are Server, Storage, UCS Fabric Interconnect etc. | [optional] 
**Elements** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewServiceitemHealthCheckErrorElement

`func NewServiceitemHealthCheckErrorElement(classId string, objectType string, ) *ServiceitemHealthCheckErrorElement`

NewServiceitemHealthCheckErrorElement instantiates a new ServiceitemHealthCheckErrorElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceitemHealthCheckErrorElementWithDefaults

`func NewServiceitemHealthCheckErrorElementWithDefaults() *ServiceitemHealthCheckErrorElement`

NewServiceitemHealthCheckErrorElementWithDefaults instantiates a new ServiceitemHealthCheckErrorElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServiceitemHealthCheckErrorElement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServiceitemHealthCheckErrorElement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServiceitemHealthCheckErrorElement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServiceitemHealthCheckErrorElement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServiceitemHealthCheckErrorElement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServiceitemHealthCheckErrorElement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetElementType

`func (o *ServiceitemHealthCheckErrorElement) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ServiceitemHealthCheckErrorElement) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ServiceitemHealthCheckErrorElement) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *ServiceitemHealthCheckErrorElement) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetElements

`func (o *ServiceitemHealthCheckErrorElement) GetElements() []MoMoRef`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ServiceitemHealthCheckErrorElement) GetElementsOk() (*[]MoMoRef, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ServiceitemHealthCheckErrorElement) SetElements(v []MoMoRef)`

SetElements sets Elements field to given value.

### HasElements

`func (o *ServiceitemHealthCheckErrorElement) HasElements() bool`

HasElements returns a boolean if a field has been set.

### SetElementsNil

`func (o *ServiceitemHealthCheckErrorElement) SetElementsNil(b bool)`

 SetElementsNil sets the value for Elements to be an explicit nil

### UnsetElements
`func (o *ServiceitemHealthCheckErrorElement) UnsetElements()`

UnsetElements ensures that no value is present for Elements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


