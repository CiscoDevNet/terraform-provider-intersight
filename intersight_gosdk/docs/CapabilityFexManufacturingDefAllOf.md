# CapabilityFexManufacturingDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FexManufacturingDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FexManufacturingDef"]
**Caption** | Pointer to **string** | Caption for Fabric extender. | [optional] 
**Description** | Pointer to **string** | Description for Fabric extender. | [optional] 
**FexCodeName** | Pointer to **string** | Code Name for Fabric extender. | [optional] 
**Pid** | Pointer to **string** | Product Identifier for a Fabric extender. | [optional] 
**ProductName** | Pointer to **string** | Product Name for Fabric extender. | [optional] 
**Sku** | Pointer to **string** | SKU information for Fabric extender. | [optional] 
**Vid** | Pointer to **string** | VID information for Fabric extender. | [optional] 

## Methods

### NewCapabilityFexManufacturingDefAllOf

`func NewCapabilityFexManufacturingDefAllOf(classId string, objectType string, ) *CapabilityFexManufacturingDefAllOf`

NewCapabilityFexManufacturingDefAllOf instantiates a new CapabilityFexManufacturingDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFexManufacturingDefAllOfWithDefaults

`func NewCapabilityFexManufacturingDefAllOfWithDefaults() *CapabilityFexManufacturingDefAllOf`

NewCapabilityFexManufacturingDefAllOfWithDefaults instantiates a new CapabilityFexManufacturingDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFexManufacturingDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFexManufacturingDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFexManufacturingDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFexManufacturingDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFexManufacturingDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFexManufacturingDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilityFexManufacturingDefAllOf) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilityFexManufacturingDefAllOf) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilityFexManufacturingDefAllOf) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilityFexManufacturingDefAllOf) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityFexManufacturingDefAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityFexManufacturingDefAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityFexManufacturingDefAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityFexManufacturingDefAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFexCodeName

`func (o *CapabilityFexManufacturingDefAllOf) GetFexCodeName() string`

GetFexCodeName returns the FexCodeName field if non-nil, zero value otherwise.

### GetFexCodeNameOk

`func (o *CapabilityFexManufacturingDefAllOf) GetFexCodeNameOk() (*string, bool)`

GetFexCodeNameOk returns a tuple with the FexCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFexCodeName

`func (o *CapabilityFexManufacturingDefAllOf) SetFexCodeName(v string)`

SetFexCodeName sets FexCodeName field to given value.

### HasFexCodeName

`func (o *CapabilityFexManufacturingDefAllOf) HasFexCodeName() bool`

HasFexCodeName returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityFexManufacturingDefAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityFexManufacturingDefAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityFexManufacturingDefAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityFexManufacturingDefAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilityFexManufacturingDefAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilityFexManufacturingDefAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilityFexManufacturingDefAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilityFexManufacturingDefAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *CapabilityFexManufacturingDefAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CapabilityFexManufacturingDefAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CapabilityFexManufacturingDefAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CapabilityFexManufacturingDefAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *CapabilityFexManufacturingDefAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *CapabilityFexManufacturingDefAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *CapabilityFexManufacturingDefAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *CapabilityFexManufacturingDefAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


