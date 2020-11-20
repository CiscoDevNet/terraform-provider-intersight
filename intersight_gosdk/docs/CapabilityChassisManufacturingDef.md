# CapabilityChassisManufacturingDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ChassisManufacturingDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ChassisManufacturingDef"]
**Caption** | Pointer to **string** | Caption for Chassis enclosure. | [optional] 
**ChassisCodeName** | Pointer to **string** | Chassis Code Name for Chassis enclosure. | [optional] 
**Description** | Pointer to **string** | Description for Chassis enclosure. | [optional] 
**Pid** | Pointer to **string** | Product Identifier for a Chassis enclosure. | [optional] 
**ProductName** | Pointer to **string** | Product Name for Chassis enclosure. | [optional] 
**Sku** | Pointer to **string** | SKU information for Chassis enclosure. | [optional] 
**Vid** | Pointer to **string** | VID information for Chassis enclosure. | [optional] 

## Methods

### NewCapabilityChassisManufacturingDef

`func NewCapabilityChassisManufacturingDef(classId string, objectType string, ) *CapabilityChassisManufacturingDef`

NewCapabilityChassisManufacturingDef instantiates a new CapabilityChassisManufacturingDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityChassisManufacturingDefWithDefaults

`func NewCapabilityChassisManufacturingDefWithDefaults() *CapabilityChassisManufacturingDef`

NewCapabilityChassisManufacturingDefWithDefaults instantiates a new CapabilityChassisManufacturingDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityChassisManufacturingDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityChassisManufacturingDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityChassisManufacturingDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityChassisManufacturingDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityChassisManufacturingDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityChassisManufacturingDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilityChassisManufacturingDef) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilityChassisManufacturingDef) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilityChassisManufacturingDef) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilityChassisManufacturingDef) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetChassisCodeName

`func (o *CapabilityChassisManufacturingDef) GetChassisCodeName() string`

GetChassisCodeName returns the ChassisCodeName field if non-nil, zero value otherwise.

### GetChassisCodeNameOk

`func (o *CapabilityChassisManufacturingDef) GetChassisCodeNameOk() (*string, bool)`

GetChassisCodeNameOk returns a tuple with the ChassisCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisCodeName

`func (o *CapabilityChassisManufacturingDef) SetChassisCodeName(v string)`

SetChassisCodeName sets ChassisCodeName field to given value.

### HasChassisCodeName

`func (o *CapabilityChassisManufacturingDef) HasChassisCodeName() bool`

HasChassisCodeName returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityChassisManufacturingDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityChassisManufacturingDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityChassisManufacturingDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityChassisManufacturingDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityChassisManufacturingDef) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityChassisManufacturingDef) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityChassisManufacturingDef) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityChassisManufacturingDef) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilityChassisManufacturingDef) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilityChassisManufacturingDef) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilityChassisManufacturingDef) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilityChassisManufacturingDef) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *CapabilityChassisManufacturingDef) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilityChassisManufacturingDef) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilityChassisManufacturingDef) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilityChassisManufacturingDef) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilityChassisManufacturingDef) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilityChassisManufacturingDef) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilityChassisManufacturingDef) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilityChassisManufacturingDef) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


