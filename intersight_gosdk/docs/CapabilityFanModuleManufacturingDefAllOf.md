# CapabilityFanModuleManufacturingDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FanModuleManufacturingDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FanModuleManufacturingDef"]
**Caption** | Pointer to **string** | Caption for a fan module. | [optional] 
**Description** | Pointer to **string** | Description for a fan module. | [optional] 
**Pid** | Pointer to **string** | Product Identifier for a fan module. | [optional] 
**ProductName** | Pointer to **string** | Product Name for Fan Module Unit. | [optional] 
**Sku** | Pointer to **string** | SKU information for a fan module. | [optional] 
**Vid** | Pointer to **string** | VID information for a fan module. | [optional] 

## Methods

### NewCapabilityFanModuleManufacturingDefAllOf

`func NewCapabilityFanModuleManufacturingDefAllOf(classId string, objectType string, ) *CapabilityFanModuleManufacturingDefAllOf`

NewCapabilityFanModuleManufacturingDefAllOf instantiates a new CapabilityFanModuleManufacturingDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFanModuleManufacturingDefAllOfWithDefaults

`func NewCapabilityFanModuleManufacturingDefAllOfWithDefaults() *CapabilityFanModuleManufacturingDefAllOf`

NewCapabilityFanModuleManufacturingDefAllOfWithDefaults instantiates a new CapabilityFanModuleManufacturingDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilityFanModuleManufacturingDefAllOf) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityFanModuleManufacturingDefAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityFanModuleManufacturingDefAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilityFanModuleManufacturingDefAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilityFanModuleManufacturingDefAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilityFanModuleManufacturingDefAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilityFanModuleManufacturingDefAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilityFanModuleManufacturingDefAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


