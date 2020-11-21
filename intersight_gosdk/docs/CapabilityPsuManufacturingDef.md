# CapabilityPsuManufacturingDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.PsuManufacturingDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.PsuManufacturingDef"]
**Caption** | Pointer to **string** | Caption for a power supply unit. | [optional] 
**Description** | Pointer to **string** | Description for a power supply unit. | [optional] 
**Pid** | Pointer to **string** | Product Identifier for a power supply unit. | [optional] 
**ProductName** | Pointer to **string** | Product Name for Power Supplu Unit. | [optional] 
**Sku** | Pointer to **string** | SKU information for a power supply unit. | [optional] 
**Vid** | Pointer to **string** | VID information for a power supply unit. | [optional] 

## Methods

### NewCapabilityPsuManufacturingDef

`func NewCapabilityPsuManufacturingDef(classId string, objectType string, ) *CapabilityPsuManufacturingDef`

NewCapabilityPsuManufacturingDef instantiates a new CapabilityPsuManufacturingDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPsuManufacturingDefWithDefaults

`func NewCapabilityPsuManufacturingDefWithDefaults() *CapabilityPsuManufacturingDef`

NewCapabilityPsuManufacturingDefWithDefaults instantiates a new CapabilityPsuManufacturingDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityPsuManufacturingDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityPsuManufacturingDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityPsuManufacturingDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityPsuManufacturingDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityPsuManufacturingDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityPsuManufacturingDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilityPsuManufacturingDef) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilityPsuManufacturingDef) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilityPsuManufacturingDef) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilityPsuManufacturingDef) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityPsuManufacturingDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityPsuManufacturingDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityPsuManufacturingDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityPsuManufacturingDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityPsuManufacturingDef) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityPsuManufacturingDef) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityPsuManufacturingDef) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityPsuManufacturingDef) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilityPsuManufacturingDef) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilityPsuManufacturingDef) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilityPsuManufacturingDef) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilityPsuManufacturingDef) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *CapabilityPsuManufacturingDef) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilityPsuManufacturingDef) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilityPsuManufacturingDef) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilityPsuManufacturingDef) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilityPsuManufacturingDef) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilityPsuManufacturingDef) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilityPsuManufacturingDef) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilityPsuManufacturingDef) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


