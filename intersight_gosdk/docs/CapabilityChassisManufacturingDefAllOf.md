# CapabilityChassisManufacturingDefAllOf

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

### NewCapabilityChassisManufacturingDefAllOf

`func NewCapabilityChassisManufacturingDefAllOf(classId string, objectType string, ) *CapabilityChassisManufacturingDefAllOf`

NewCapabilityChassisManufacturingDefAllOf instantiates a new CapabilityChassisManufacturingDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityChassisManufacturingDefAllOfWithDefaults

`func NewCapabilityChassisManufacturingDefAllOfWithDefaults() *CapabilityChassisManufacturingDefAllOf`

NewCapabilityChassisManufacturingDefAllOfWithDefaults instantiates a new CapabilityChassisManufacturingDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityChassisManufacturingDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityChassisManufacturingDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityChassisManufacturingDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityChassisManufacturingDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilityChassisManufacturingDefAllOf) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilityChassisManufacturingDefAllOf) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilityChassisManufacturingDefAllOf) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetChassisCodeName

`func (o *CapabilityChassisManufacturingDefAllOf) GetChassisCodeName() string`

GetChassisCodeName returns the ChassisCodeName field if non-nil, zero value otherwise.

### GetChassisCodeNameOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetChassisCodeNameOk() (*string, bool)`

GetChassisCodeNameOk returns a tuple with the ChassisCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisCodeName

`func (o *CapabilityChassisManufacturingDefAllOf) SetChassisCodeName(v string)`

SetChassisCodeName sets ChassisCodeName field to given value.

### HasChassisCodeName

`func (o *CapabilityChassisManufacturingDefAllOf) HasChassisCodeName() bool`

HasChassisCodeName returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityChassisManufacturingDefAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityChassisManufacturingDefAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityChassisManufacturingDefAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityChassisManufacturingDefAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityChassisManufacturingDefAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityChassisManufacturingDefAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilityChassisManufacturingDefAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilityChassisManufacturingDefAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilityChassisManufacturingDefAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *CapabilityChassisManufacturingDefAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilityChassisManufacturingDefAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilityChassisManufacturingDefAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilityChassisManufacturingDefAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilityChassisManufacturingDefAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilityChassisManufacturingDefAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilityChassisManufacturingDefAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


