# CapabilityFanModuleManufacturingDef

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

### NewCapabilityFanModuleManufacturingDef

`func NewCapabilityFanModuleManufacturingDef(classId string, objectType string, ) *CapabilityFanModuleManufacturingDef`

NewCapabilityFanModuleManufacturingDef instantiates a new CapabilityFanModuleManufacturingDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFanModuleManufacturingDefWithDefaults

`func NewCapabilityFanModuleManufacturingDefWithDefaults() *CapabilityFanModuleManufacturingDef`

NewCapabilityFanModuleManufacturingDefWithDefaults instantiates a new CapabilityFanModuleManufacturingDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFanModuleManufacturingDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFanModuleManufacturingDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFanModuleManufacturingDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFanModuleManufacturingDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFanModuleManufacturingDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFanModuleManufacturingDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilityFanModuleManufacturingDef) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilityFanModuleManufacturingDef) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilityFanModuleManufacturingDef) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilityFanModuleManufacturingDef) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityFanModuleManufacturingDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityFanModuleManufacturingDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityFanModuleManufacturingDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityFanModuleManufacturingDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityFanModuleManufacturingDef) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityFanModuleManufacturingDef) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityFanModuleManufacturingDef) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityFanModuleManufacturingDef) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilityFanModuleManufacturingDef) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilityFanModuleManufacturingDef) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilityFanModuleManufacturingDef) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilityFanModuleManufacturingDef) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *CapabilityFanModuleManufacturingDef) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilityFanModuleManufacturingDef) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilityFanModuleManufacturingDef) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilityFanModuleManufacturingDef) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilityFanModuleManufacturingDef) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilityFanModuleManufacturingDef) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilityFanModuleManufacturingDef) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilityFanModuleManufacturingDef) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


