# CapabilityCimcFirmwareDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.CimcFirmwareDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.CimcFirmwareDescriptor"]
**AdapterEpProxyEnabled** | Pointer to **bool** | Indicates whether the server uses ep proxy to communicate with the adapter. | [optional] [readonly] 
**Revision** | Pointer to **string** | Revision information for the server. | [optional] 
**UuidSupportedVer** | Pointer to **string** | Minimum server firmware version for UUID feature support. | [optional] [readonly] 

## Methods

### NewCapabilityCimcFirmwareDescriptor

`func NewCapabilityCimcFirmwareDescriptor(classId string, objectType string, ) *CapabilityCimcFirmwareDescriptor`

NewCapabilityCimcFirmwareDescriptor instantiates a new CapabilityCimcFirmwareDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityCimcFirmwareDescriptorWithDefaults

`func NewCapabilityCimcFirmwareDescriptorWithDefaults() *CapabilityCimcFirmwareDescriptor`

NewCapabilityCimcFirmwareDescriptorWithDefaults instantiates a new CapabilityCimcFirmwareDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityCimcFirmwareDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityCimcFirmwareDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityCimcFirmwareDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityCimcFirmwareDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityCimcFirmwareDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityCimcFirmwareDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterEpProxyEnabled

`func (o *CapabilityCimcFirmwareDescriptor) GetAdapterEpProxyEnabled() bool`

GetAdapterEpProxyEnabled returns the AdapterEpProxyEnabled field if non-nil, zero value otherwise.

### GetAdapterEpProxyEnabledOk

`func (o *CapabilityCimcFirmwareDescriptor) GetAdapterEpProxyEnabledOk() (*bool, bool)`

GetAdapterEpProxyEnabledOk returns a tuple with the AdapterEpProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterEpProxyEnabled

`func (o *CapabilityCimcFirmwareDescriptor) SetAdapterEpProxyEnabled(v bool)`

SetAdapterEpProxyEnabled sets AdapterEpProxyEnabled field to given value.

### HasAdapterEpProxyEnabled

`func (o *CapabilityCimcFirmwareDescriptor) HasAdapterEpProxyEnabled() bool`

HasAdapterEpProxyEnabled returns a boolean if a field has been set.

### GetRevision

`func (o *CapabilityCimcFirmwareDescriptor) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CapabilityCimcFirmwareDescriptor) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CapabilityCimcFirmwareDescriptor) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CapabilityCimcFirmwareDescriptor) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetUuidSupportedVer

`func (o *CapabilityCimcFirmwareDescriptor) GetUuidSupportedVer() string`

GetUuidSupportedVer returns the UuidSupportedVer field if non-nil, zero value otherwise.

### GetUuidSupportedVerOk

`func (o *CapabilityCimcFirmwareDescriptor) GetUuidSupportedVerOk() (*string, bool)`

GetUuidSupportedVerOk returns a tuple with the UuidSupportedVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSupportedVer

`func (o *CapabilityCimcFirmwareDescriptor) SetUuidSupportedVer(v string)`

SetUuidSupportedVer sets UuidSupportedVer field to given value.

### HasUuidSupportedVer

`func (o *CapabilityCimcFirmwareDescriptor) HasUuidSupportedVer() bool`

HasUuidSupportedVer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


