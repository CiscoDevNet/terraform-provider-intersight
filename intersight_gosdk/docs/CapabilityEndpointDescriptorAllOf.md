# CapabilityEndpointDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Detailed information about the endpoint. | [optional] 
**Model** | Pointer to **string** | The model of the endpoint, for which this capability information is applicable. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the endpoint, for which this capability information is applicable. | [optional] 
**Version** | Pointer to **string** | The firmware or software version of the endpoint, for which this capability information is applicable. | [optional] 
**Capabilities** | Pointer to [**[]CapabilityCapabilityRelationship**](capability.Capability.Relationship.md) | An array of relationships to capabilityCapability resources. | [optional] 
**Section** | Pointer to [**CapabilitySectionRelationship**](capability.Section.Relationship.md) |  | [optional] 

## Methods

### NewCapabilityEndpointDescriptorAllOf

`func NewCapabilityEndpointDescriptorAllOf() *CapabilityEndpointDescriptorAllOf`

NewCapabilityEndpointDescriptorAllOf instantiates a new CapabilityEndpointDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityEndpointDescriptorAllOfWithDefaults

`func NewCapabilityEndpointDescriptorAllOfWithDefaults() *CapabilityEndpointDescriptorAllOf`

NewCapabilityEndpointDescriptorAllOfWithDefaults instantiates a new CapabilityEndpointDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CapabilityEndpointDescriptorAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityEndpointDescriptorAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityEndpointDescriptorAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityEndpointDescriptorAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityEndpointDescriptorAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityEndpointDescriptorAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityEndpointDescriptorAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityEndpointDescriptorAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *CapabilityEndpointDescriptorAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityEndpointDescriptorAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityEndpointDescriptorAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityEndpointDescriptorAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *CapabilityEndpointDescriptorAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapabilityEndpointDescriptorAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapabilityEndpointDescriptorAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapabilityEndpointDescriptorAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCapabilities

`func (o *CapabilityEndpointDescriptorAllOf) GetCapabilities() []CapabilityCapabilityRelationship`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CapabilityEndpointDescriptorAllOf) GetCapabilitiesOk() (*[]CapabilityCapabilityRelationship, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CapabilityEndpointDescriptorAllOf) SetCapabilities(v []CapabilityCapabilityRelationship)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CapabilityEndpointDescriptorAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *CapabilityEndpointDescriptorAllOf) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *CapabilityEndpointDescriptorAllOf) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetSection

`func (o *CapabilityEndpointDescriptorAllOf) GetSection() CapabilitySectionRelationship`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *CapabilityEndpointDescriptorAllOf) GetSectionOk() (*CapabilitySectionRelationship, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *CapabilityEndpointDescriptorAllOf) SetSection(v CapabilitySectionRelationship)`

SetSection sets Section field to given value.

### HasSection

`func (o *CapabilityEndpointDescriptorAllOf) HasSection() bool`

HasSection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


