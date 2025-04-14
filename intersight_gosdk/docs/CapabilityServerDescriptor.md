# CapabilityServerDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerDescriptor"]
**IsNcsiEnabled** | Pointer to **bool** | Indicates whether the CIMC to VIC side-band interface is enabled on the server. | [optional] 
**IsPplEnabled** | Pointer to **bool** | Indicates Processor Package Power Limit for the server. | [optional] [default to false]
**MlomAdapterPcieSlotNumber** | Pointer to **int64** | Indicates PCIe Slot numerical value for each Server model MLOM slot. | [optional] 
**ServerFormFactor** | Pointer to **string** | The form factor (blade/rack/etc) of the server. * &#x60;unknown&#x60; - The form factor of the server is unknown. * &#x60;blade&#x60; - Blade server form factor. * &#x60;rack&#x60; - Rack unit server form factor. | [optional] [readonly] [default to "unknown"]

## Methods

### NewCapabilityServerDescriptor

`func NewCapabilityServerDescriptor(classId string, objectType string, ) *CapabilityServerDescriptor`

NewCapabilityServerDescriptor instantiates a new CapabilityServerDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerDescriptorWithDefaults

`func NewCapabilityServerDescriptorWithDefaults() *CapabilityServerDescriptor`

NewCapabilityServerDescriptorWithDefaults instantiates a new CapabilityServerDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsNcsiEnabled

`func (o *CapabilityServerDescriptor) GetIsNcsiEnabled() bool`

GetIsNcsiEnabled returns the IsNcsiEnabled field if non-nil, zero value otherwise.

### GetIsNcsiEnabledOk

`func (o *CapabilityServerDescriptor) GetIsNcsiEnabledOk() (*bool, bool)`

GetIsNcsiEnabledOk returns a tuple with the IsNcsiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNcsiEnabled

`func (o *CapabilityServerDescriptor) SetIsNcsiEnabled(v bool)`

SetIsNcsiEnabled sets IsNcsiEnabled field to given value.

### HasIsNcsiEnabled

`func (o *CapabilityServerDescriptor) HasIsNcsiEnabled() bool`

HasIsNcsiEnabled returns a boolean if a field has been set.

### GetIsPplEnabled

`func (o *CapabilityServerDescriptor) GetIsPplEnabled() bool`

GetIsPplEnabled returns the IsPplEnabled field if non-nil, zero value otherwise.

### GetIsPplEnabledOk

`func (o *CapabilityServerDescriptor) GetIsPplEnabledOk() (*bool, bool)`

GetIsPplEnabledOk returns a tuple with the IsPplEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPplEnabled

`func (o *CapabilityServerDescriptor) SetIsPplEnabled(v bool)`

SetIsPplEnabled sets IsPplEnabled field to given value.

### HasIsPplEnabled

`func (o *CapabilityServerDescriptor) HasIsPplEnabled() bool`

HasIsPplEnabled returns a boolean if a field has been set.

### GetMlomAdapterPcieSlotNumber

`func (o *CapabilityServerDescriptor) GetMlomAdapterPcieSlotNumber() int64`

GetMlomAdapterPcieSlotNumber returns the MlomAdapterPcieSlotNumber field if non-nil, zero value otherwise.

### GetMlomAdapterPcieSlotNumberOk

`func (o *CapabilityServerDescriptor) GetMlomAdapterPcieSlotNumberOk() (*int64, bool)`

GetMlomAdapterPcieSlotNumberOk returns a tuple with the MlomAdapterPcieSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlomAdapterPcieSlotNumber

`func (o *CapabilityServerDescriptor) SetMlomAdapterPcieSlotNumber(v int64)`

SetMlomAdapterPcieSlotNumber sets MlomAdapterPcieSlotNumber field to given value.

### HasMlomAdapterPcieSlotNumber

`func (o *CapabilityServerDescriptor) HasMlomAdapterPcieSlotNumber() bool`

HasMlomAdapterPcieSlotNumber returns a boolean if a field has been set.

### GetServerFormFactor

`func (o *CapabilityServerDescriptor) GetServerFormFactor() string`

GetServerFormFactor returns the ServerFormFactor field if non-nil, zero value otherwise.

### GetServerFormFactorOk

`func (o *CapabilityServerDescriptor) GetServerFormFactorOk() (*string, bool)`

GetServerFormFactorOk returns a tuple with the ServerFormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFormFactor

`func (o *CapabilityServerDescriptor) SetServerFormFactor(v string)`

SetServerFormFactor sets ServerFormFactor field to given value.

### HasServerFormFactor

`func (o *CapabilityServerDescriptor) HasServerFormFactor() bool`

HasServerFormFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


