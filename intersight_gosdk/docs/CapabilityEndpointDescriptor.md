# CapabilityEndpointDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Description** | Pointer to **string** | Detailed information about the endpoint. | [optional] 
**Model** | Pointer to **string** | The model of the endpoint, for which this capability information is applicable. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the endpoint, for which this capability information is applicable. | [optional] 
**Version** | Pointer to **string** | The firmware or software version of the endpoint, for which this capability information is applicable. | [optional] 
**Capabilities** | Pointer to [**[]CapabilityCapabilityRelationship**](CapabilityCapabilityRelationship.md) | An array of relationships to capabilityCapability resources. | [optional] 

## Methods

### NewCapabilityEndpointDescriptor

`func NewCapabilityEndpointDescriptor(classId string, objectType string, ) *CapabilityEndpointDescriptor`

NewCapabilityEndpointDescriptor instantiates a new CapabilityEndpointDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityEndpointDescriptorWithDefaults

`func NewCapabilityEndpointDescriptorWithDefaults() *CapabilityEndpointDescriptor`

NewCapabilityEndpointDescriptorWithDefaults instantiates a new CapabilityEndpointDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityEndpointDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityEndpointDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityEndpointDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityEndpointDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityEndpointDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityEndpointDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityEndpointDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityEndpointDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityEndpointDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityEndpointDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityEndpointDescriptor) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityEndpointDescriptor) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityEndpointDescriptor) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityEndpointDescriptor) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *CapabilityEndpointDescriptor) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityEndpointDescriptor) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityEndpointDescriptor) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityEndpointDescriptor) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *CapabilityEndpointDescriptor) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapabilityEndpointDescriptor) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapabilityEndpointDescriptor) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapabilityEndpointDescriptor) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCapabilities

`func (o *CapabilityEndpointDescriptor) GetCapabilities() []CapabilityCapabilityRelationship`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CapabilityEndpointDescriptor) GetCapabilitiesOk() (*[]CapabilityCapabilityRelationship, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CapabilityEndpointDescriptor) SetCapabilities(v []CapabilityCapabilityRelationship)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CapabilityEndpointDescriptor) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *CapabilityEndpointDescriptor) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *CapabilityEndpointDescriptor) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


