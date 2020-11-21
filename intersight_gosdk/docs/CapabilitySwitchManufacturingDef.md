# CapabilitySwitchManufacturingDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchManufacturingDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchManufacturingDef"]
**Caption** | Pointer to **string** | Caption for Switch/Fabric-Interconnect. | [optional] 
**Description** | Pointer to **string** | Description for Switch/Fabric-Interconnect. | [optional] 
**PartNumber** | Pointer to **string** | Part Number for Switch/Fabric-Interconnect. | [optional] 
**ProductName** | Pointer to **string** | Product Name for Switch/Fabric-Interconnect. | [optional] 

## Methods

### NewCapabilitySwitchManufacturingDef

`func NewCapabilitySwitchManufacturingDef(classId string, objectType string, ) *CapabilitySwitchManufacturingDef`

NewCapabilitySwitchManufacturingDef instantiates a new CapabilitySwitchManufacturingDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchManufacturingDefWithDefaults

`func NewCapabilitySwitchManufacturingDefWithDefaults() *CapabilitySwitchManufacturingDef`

NewCapabilitySwitchManufacturingDefWithDefaults instantiates a new CapabilitySwitchManufacturingDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchManufacturingDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchManufacturingDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchManufacturingDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchManufacturingDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchManufacturingDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchManufacturingDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaption

`func (o *CapabilitySwitchManufacturingDef) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CapabilitySwitchManufacturingDef) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CapabilitySwitchManufacturingDef) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CapabilitySwitchManufacturingDef) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilitySwitchManufacturingDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilitySwitchManufacturingDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilitySwitchManufacturingDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilitySwitchManufacturingDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPartNumber

`func (o *CapabilitySwitchManufacturingDef) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *CapabilitySwitchManufacturingDef) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *CapabilitySwitchManufacturingDef) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *CapabilitySwitchManufacturingDef) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProductName

`func (o *CapabilitySwitchManufacturingDef) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CapabilitySwitchManufacturingDef) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CapabilitySwitchManufacturingDef) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CapabilitySwitchManufacturingDef) HasProductName() bool`

HasProductName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


