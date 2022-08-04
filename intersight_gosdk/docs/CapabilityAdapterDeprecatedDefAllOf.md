# CapabilityAdapterDeprecatedDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterDeprecatedDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterDeprecatedDef"]
**Model** | Pointer to **string** | Model of the unsupported adapter. | [optional] 
**Vendor** | Pointer to **string** | Vendor of the unsupported adapter. | [optional] 

## Methods

### NewCapabilityAdapterDeprecatedDefAllOf

`func NewCapabilityAdapterDeprecatedDefAllOf(classId string, objectType string, ) *CapabilityAdapterDeprecatedDefAllOf`

NewCapabilityAdapterDeprecatedDefAllOf instantiates a new CapabilityAdapterDeprecatedDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterDeprecatedDefAllOfWithDefaults

`func NewCapabilityAdapterDeprecatedDefAllOfWithDefaults() *CapabilityAdapterDeprecatedDefAllOf`

NewCapabilityAdapterDeprecatedDefAllOfWithDefaults instantiates a new CapabilityAdapterDeprecatedDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterDeprecatedDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterDeprecatedDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModel

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityAdapterDeprecatedDefAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityAdapterDeprecatedDefAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityAdapterDeprecatedDefAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityAdapterDeprecatedDefAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityAdapterDeprecatedDefAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


