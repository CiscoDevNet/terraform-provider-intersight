# FabricSpanSourceVnicEthIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SpanSourceVnicEthIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SpanSourceVnicEthIf"]
**Vnic** | Pointer to [**NullableVnicEthIfRelationship**](VnicEthIfRelationship.md) |  | [optional] 

## Methods

### NewFabricSpanSourceVnicEthIf

`func NewFabricSpanSourceVnicEthIf(classId string, objectType string, ) *FabricSpanSourceVnicEthIf`

NewFabricSpanSourceVnicEthIf instantiates a new FabricSpanSourceVnicEthIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSpanSourceVnicEthIfWithDefaults

`func NewFabricSpanSourceVnicEthIfWithDefaults() *FabricSpanSourceVnicEthIf`

NewFabricSpanSourceVnicEthIfWithDefaults instantiates a new FabricSpanSourceVnicEthIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSpanSourceVnicEthIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSpanSourceVnicEthIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSpanSourceVnicEthIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSpanSourceVnicEthIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSpanSourceVnicEthIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSpanSourceVnicEthIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVnic

`func (o *FabricSpanSourceVnicEthIf) GetVnic() VnicEthIfRelationship`

GetVnic returns the Vnic field if non-nil, zero value otherwise.

### GetVnicOk

`func (o *FabricSpanSourceVnicEthIf) GetVnicOk() (*VnicEthIfRelationship, bool)`

GetVnicOk returns a tuple with the Vnic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnic

`func (o *FabricSpanSourceVnicEthIf) SetVnic(v VnicEthIfRelationship)`

SetVnic sets Vnic field to given value.

### HasVnic

`func (o *FabricSpanSourceVnicEthIf) HasVnic() bool`

HasVnic returns a boolean if a field has been set.

### SetVnicNil

`func (o *FabricSpanSourceVnicEthIf) SetVnicNil(b bool)`

 SetVnicNil sets the value for Vnic to be an explicit nil

### UnsetVnic
`func (o *FabricSpanSourceVnicEthIf) UnsetVnic()`

UnsetVnic ensures that no value is present for Vnic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


