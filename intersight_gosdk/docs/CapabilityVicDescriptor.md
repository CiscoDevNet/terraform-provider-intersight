# CapabilityVicDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.VicDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.VicDescriptor"]
**Model** | Pointer to **string** | The model of the endpoint, for which this capability information is applicable. | [optional] 
**VicId** | Pointer to **string** | Vic Id assigned for the adapter. | [optional] 

## Methods

### NewCapabilityVicDescriptor

`func NewCapabilityVicDescriptor(classId string, objectType string, ) *CapabilityVicDescriptor`

NewCapabilityVicDescriptor instantiates a new CapabilityVicDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityVicDescriptorWithDefaults

`func NewCapabilityVicDescriptorWithDefaults() *CapabilityVicDescriptor`

NewCapabilityVicDescriptorWithDefaults instantiates a new CapabilityVicDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityVicDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityVicDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityVicDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityVicDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityVicDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityVicDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModel

`func (o *CapabilityVicDescriptor) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityVicDescriptor) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityVicDescriptor) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityVicDescriptor) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVicId

`func (o *CapabilityVicDescriptor) GetVicId() string`

GetVicId returns the VicId field if non-nil, zero value otherwise.

### GetVicIdOk

`func (o *CapabilityVicDescriptor) GetVicIdOk() (*string, bool)`

GetVicIdOk returns a tuple with the VicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVicId

`func (o *CapabilityVicDescriptor) SetVicId(v string)`

SetVicId sets VicId field to given value.

### HasVicId

`func (o *CapabilityVicDescriptor) HasVicId() bool`

HasVicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


