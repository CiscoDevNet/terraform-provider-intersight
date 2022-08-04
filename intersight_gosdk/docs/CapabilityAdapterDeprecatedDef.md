# CapabilityAdapterDeprecatedDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterDeprecatedDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterDeprecatedDef"]
**Model** | Pointer to **string** | Model of the unsupported adapter. | [optional] 
**Vendor** | Pointer to **string** | Vendor of the unsupported adapter. | [optional] 

## Methods

### NewCapabilityAdapterDeprecatedDef

`func NewCapabilityAdapterDeprecatedDef(classId string, objectType string, ) *CapabilityAdapterDeprecatedDef`

NewCapabilityAdapterDeprecatedDef instantiates a new CapabilityAdapterDeprecatedDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterDeprecatedDefWithDefaults

`func NewCapabilityAdapterDeprecatedDefWithDefaults() *CapabilityAdapterDeprecatedDef`

NewCapabilityAdapterDeprecatedDefWithDefaults instantiates a new CapabilityAdapterDeprecatedDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterDeprecatedDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterDeprecatedDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterDeprecatedDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterDeprecatedDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterDeprecatedDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterDeprecatedDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModel

`func (o *CapabilityAdapterDeprecatedDef) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityAdapterDeprecatedDef) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityAdapterDeprecatedDef) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityAdapterDeprecatedDef) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *CapabilityAdapterDeprecatedDef) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityAdapterDeprecatedDef) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityAdapterDeprecatedDef) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityAdapterDeprecatedDef) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


