# FirmwareComponentMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentLabel** | Pointer to **string** | The name of the component in the compressed HSU bundle. | [optional] 
**ComponentType** | Pointer to **string** | The type of component image within the distributable. * &#x60;ALL&#x60; - This represents all the components. * &#x60;ALL,HDD&#x60; - This represents all the components plus the HDDs. * &#x60;None&#x60; - This represents none of the components. * &#x60;NXOS&#x60; - This represents NXOS components. * &#x60;IOM&#x60; - This represents IOM components. * &#x60;PSU&#x60; - This represents PSU components. * &#x60;CIMC&#x60; - This represents CIMC components. * &#x60;BIOS&#x60; - This represents BIOS components. * &#x60;PCIE&#x60; - This represents PCIE components. * &#x60;Drive&#x60; - This represents Storage components. * &#x60;DIMM&#x60; - This represents DIMM components. * &#x60;BoardController&#x60; - This represents Board Controller components. * &#x60;StorageController&#x60; - This represents Storage Controller components. * &#x60;HBA&#x60; - This represents HBA components. * &#x60;GPU&#x60; - This represents GPU components. * &#x60;SasExpander&#x60; - This represents SasExpander components. * &#x60;MSwitch&#x60; - This represents mSwitch components. * &#x60;CMC&#x60; - This represents CMC components. | [optional] [default to "ALL"]
**Description** | Pointer to **string** | This shows the description of component image within the distributable. | [optional] 
**Disruption** | Pointer to **string** | The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle. * &#x60;None&#x60; - Indicates that the component did not receive a disruption request. * &#x60;HostReboot&#x60; - Indicates that the component received a host reboot request. * &#x60;EndpointReboot&#x60; - Indicates that the component received an end point reboot request. * &#x60;ManualPowerCycle&#x60; - Indicates that the component received a manual power cycle request. * &#x60;AutomaticPowerCycle&#x60; - Indicates that the component received an automatic power cycle request. | [optional] [default to "None"]
**IsOobSupported** | Pointer to **bool** | If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. | [optional] 
**Model** | Pointer to **string** | The model of the component image in the distributable. | [optional] 
**OobManageability** | Pointer to **[]string** |  | [optional] 
**PackedVersion** | Pointer to **string** | The image version of components packaged in the distributable. | [optional] 
**RedfishUrl** | Pointer to **string** | The redfish target for each component. | [optional] 
**Vendor** | Pointer to **string** | The version of component image in the distributable. | [optional] 

## Methods

### NewFirmwareComponentMeta

`func NewFirmwareComponentMeta() *FirmwareComponentMeta`

NewFirmwareComponentMeta instantiates a new FirmwareComponentMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareComponentMetaWithDefaults

`func NewFirmwareComponentMetaWithDefaults() *FirmwareComponentMeta`

NewFirmwareComponentMetaWithDefaults instantiates a new FirmwareComponentMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentLabel

`func (o *FirmwareComponentMeta) GetComponentLabel() string`

GetComponentLabel returns the ComponentLabel field if non-nil, zero value otherwise.

### GetComponentLabelOk

`func (o *FirmwareComponentMeta) GetComponentLabelOk() (*string, bool)`

GetComponentLabelOk returns a tuple with the ComponentLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLabel

`func (o *FirmwareComponentMeta) SetComponentLabel(v string)`

SetComponentLabel sets ComponentLabel field to given value.

### HasComponentLabel

`func (o *FirmwareComponentMeta) HasComponentLabel() bool`

HasComponentLabel returns a boolean if a field has been set.

### GetComponentType

`func (o *FirmwareComponentMeta) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *FirmwareComponentMeta) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *FirmwareComponentMeta) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *FirmwareComponentMeta) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetDescription

`func (o *FirmwareComponentMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirmwareComponentMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirmwareComponentMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirmwareComponentMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisruption

`func (o *FirmwareComponentMeta) GetDisruption() string`

GetDisruption returns the Disruption field if non-nil, zero value otherwise.

### GetDisruptionOk

`func (o *FirmwareComponentMeta) GetDisruptionOk() (*string, bool)`

GetDisruptionOk returns a tuple with the Disruption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruption

`func (o *FirmwareComponentMeta) SetDisruption(v string)`

SetDisruption sets Disruption field to given value.

### HasDisruption

`func (o *FirmwareComponentMeta) HasDisruption() bool`

HasDisruption returns a boolean if a field has been set.

### GetIsOobSupported

`func (o *FirmwareComponentMeta) GetIsOobSupported() bool`

GetIsOobSupported returns the IsOobSupported field if non-nil, zero value otherwise.

### GetIsOobSupportedOk

`func (o *FirmwareComponentMeta) GetIsOobSupportedOk() (*bool, bool)`

GetIsOobSupportedOk returns a tuple with the IsOobSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOobSupported

`func (o *FirmwareComponentMeta) SetIsOobSupported(v bool)`

SetIsOobSupported sets IsOobSupported field to given value.

### HasIsOobSupported

`func (o *FirmwareComponentMeta) HasIsOobSupported() bool`

HasIsOobSupported returns a boolean if a field has been set.

### GetModel

`func (o *FirmwareComponentMeta) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FirmwareComponentMeta) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FirmwareComponentMeta) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FirmwareComponentMeta) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOobManageability

`func (o *FirmwareComponentMeta) GetOobManageability() []string`

GetOobManageability returns the OobManageability field if non-nil, zero value otherwise.

### GetOobManageabilityOk

`func (o *FirmwareComponentMeta) GetOobManageabilityOk() (*[]string, bool)`

GetOobManageabilityOk returns a tuple with the OobManageability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobManageability

`func (o *FirmwareComponentMeta) SetOobManageability(v []string)`

SetOobManageability sets OobManageability field to given value.

### HasOobManageability

`func (o *FirmwareComponentMeta) HasOobManageability() bool`

HasOobManageability returns a boolean if a field has been set.

### GetPackedVersion

`func (o *FirmwareComponentMeta) GetPackedVersion() string`

GetPackedVersion returns the PackedVersion field if non-nil, zero value otherwise.

### GetPackedVersionOk

`func (o *FirmwareComponentMeta) GetPackedVersionOk() (*string, bool)`

GetPackedVersionOk returns a tuple with the PackedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackedVersion

`func (o *FirmwareComponentMeta) SetPackedVersion(v string)`

SetPackedVersion sets PackedVersion field to given value.

### HasPackedVersion

`func (o *FirmwareComponentMeta) HasPackedVersion() bool`

HasPackedVersion returns a boolean if a field has been set.

### GetRedfishUrl

`func (o *FirmwareComponentMeta) GetRedfishUrl() string`

GetRedfishUrl returns the RedfishUrl field if non-nil, zero value otherwise.

### GetRedfishUrlOk

`func (o *FirmwareComponentMeta) GetRedfishUrlOk() (*string, bool)`

GetRedfishUrlOk returns a tuple with the RedfishUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishUrl

`func (o *FirmwareComponentMeta) SetRedfishUrl(v string)`

SetRedfishUrl sets RedfishUrl field to given value.

### HasRedfishUrl

`func (o *FirmwareComponentMeta) HasRedfishUrl() bool`

HasRedfishUrl returns a boolean if a field has been set.

### GetVendor

`func (o *FirmwareComponentMeta) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareComponentMeta) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareComponentMeta) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FirmwareComponentMeta) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


