# BootPrecisionPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootDevices** | Pointer to [**[]BootDeviceBase**](boot.DeviceBase.md) |  | [optional] 
**ConfiguredBootMode** | Pointer to **string** | Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. * &#x60;Legacy&#x60; - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader. * &#x60;Uefi&#x60; - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. | [optional] [default to "Legacy"]
**EnforceUefiSecureBoot** | Pointer to **bool** | If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewBootPrecisionPolicyAllOf

`func NewBootPrecisionPolicyAllOf() *BootPrecisionPolicyAllOf`

NewBootPrecisionPolicyAllOf instantiates a new BootPrecisionPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPrecisionPolicyAllOfWithDefaults

`func NewBootPrecisionPolicyAllOfWithDefaults() *BootPrecisionPolicyAllOf`

NewBootPrecisionPolicyAllOfWithDefaults instantiates a new BootPrecisionPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootDevices

`func (o *BootPrecisionPolicyAllOf) GetBootDevices() []BootDeviceBase`

GetBootDevices returns the BootDevices field if non-nil, zero value otherwise.

### GetBootDevicesOk

`func (o *BootPrecisionPolicyAllOf) GetBootDevicesOk() (*[]BootDeviceBase, bool)`

GetBootDevicesOk returns a tuple with the BootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevices

`func (o *BootPrecisionPolicyAllOf) SetBootDevices(v []BootDeviceBase)`

SetBootDevices sets BootDevices field to given value.

### HasBootDevices

`func (o *BootPrecisionPolicyAllOf) HasBootDevices() bool`

HasBootDevices returns a boolean if a field has been set.

### GetConfiguredBootMode

`func (o *BootPrecisionPolicyAllOf) GetConfiguredBootMode() string`

GetConfiguredBootMode returns the ConfiguredBootMode field if non-nil, zero value otherwise.

### GetConfiguredBootModeOk

`func (o *BootPrecisionPolicyAllOf) GetConfiguredBootModeOk() (*string, bool)`

GetConfiguredBootModeOk returns a tuple with the ConfiguredBootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredBootMode

`func (o *BootPrecisionPolicyAllOf) SetConfiguredBootMode(v string)`

SetConfiguredBootMode sets ConfiguredBootMode field to given value.

### HasConfiguredBootMode

`func (o *BootPrecisionPolicyAllOf) HasConfiguredBootMode() bool`

HasConfiguredBootMode returns a boolean if a field has been set.

### GetEnforceUefiSecureBoot

`func (o *BootPrecisionPolicyAllOf) GetEnforceUefiSecureBoot() bool`

GetEnforceUefiSecureBoot returns the EnforceUefiSecureBoot field if non-nil, zero value otherwise.

### GetEnforceUefiSecureBootOk

`func (o *BootPrecisionPolicyAllOf) GetEnforceUefiSecureBootOk() (*bool, bool)`

GetEnforceUefiSecureBootOk returns a tuple with the EnforceUefiSecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUefiSecureBoot

`func (o *BootPrecisionPolicyAllOf) SetEnforceUefiSecureBoot(v bool)`

SetEnforceUefiSecureBoot sets EnforceUefiSecureBoot field to given value.

### HasEnforceUefiSecureBoot

`func (o *BootPrecisionPolicyAllOf) HasEnforceUefiSecureBoot() bool`

HasEnforceUefiSecureBoot returns a boolean if a field has been set.

### GetOrganization

`func (o *BootPrecisionPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BootPrecisionPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BootPrecisionPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BootPrecisionPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *BootPrecisionPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BootPrecisionPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BootPrecisionPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BootPrecisionPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *BootPrecisionPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *BootPrecisionPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


