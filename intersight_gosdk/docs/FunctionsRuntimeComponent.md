# FunctionsRuntimeComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.RuntimeComponent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.RuntimeComponent"]
**Name** | Pointer to **string** | Official name of the component. | [optional] 
**Version** | Pointer to **string** | Official version of the component. | [optional] 

## Methods

### NewFunctionsRuntimeComponent

`func NewFunctionsRuntimeComponent(classId string, objectType string, ) *FunctionsRuntimeComponent`

NewFunctionsRuntimeComponent instantiates a new FunctionsRuntimeComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsRuntimeComponentWithDefaults

`func NewFunctionsRuntimeComponentWithDefaults() *FunctionsRuntimeComponent`

NewFunctionsRuntimeComponentWithDefaults instantiates a new FunctionsRuntimeComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsRuntimeComponent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsRuntimeComponent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsRuntimeComponent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsRuntimeComponent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsRuntimeComponent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsRuntimeComponent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *FunctionsRuntimeComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsRuntimeComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsRuntimeComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsRuntimeComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *FunctionsRuntimeComponent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunctionsRuntimeComponent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunctionsRuntimeComponent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunctionsRuntimeComponent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


