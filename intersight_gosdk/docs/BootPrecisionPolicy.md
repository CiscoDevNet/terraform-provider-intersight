# BootPrecisionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.PrecisionPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.PrecisionPolicy"]
**BootDevices** | Pointer to [**[]BootDeviceBase**](BootDeviceBase.md) |  | [optional] 
**ConfiguredBootMode** | Pointer to **string** | Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. To apply this setting, Please reboot the server. * &#x60;Legacy&#x60; - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader. * &#x60;Uefi&#x60; - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. | [optional] [default to "Legacy"]
**EnforceUefiSecureBoot** | Pointer to **bool** | If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). | [optional] [default to false]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewBootPrecisionPolicy

`func NewBootPrecisionPolicy(classId string, objectType string, ) *BootPrecisionPolicy`

NewBootPrecisionPolicy instantiates a new BootPrecisionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPrecisionPolicyWithDefaults

`func NewBootPrecisionPolicyWithDefaults() *BootPrecisionPolicy`

NewBootPrecisionPolicyWithDefaults instantiates a new BootPrecisionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootPrecisionPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootPrecisionPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootPrecisionPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootPrecisionPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootPrecisionPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootPrecisionPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootDevices

`func (o *BootPrecisionPolicy) GetBootDevices() []BootDeviceBase`

GetBootDevices returns the BootDevices field if non-nil, zero value otherwise.

### GetBootDevicesOk

`func (o *BootPrecisionPolicy) GetBootDevicesOk() (*[]BootDeviceBase, bool)`

GetBootDevicesOk returns a tuple with the BootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevices

`func (o *BootPrecisionPolicy) SetBootDevices(v []BootDeviceBase)`

SetBootDevices sets BootDevices field to given value.

### HasBootDevices

`func (o *BootPrecisionPolicy) HasBootDevices() bool`

HasBootDevices returns a boolean if a field has been set.

### SetBootDevicesNil

`func (o *BootPrecisionPolicy) SetBootDevicesNil(b bool)`

 SetBootDevicesNil sets the value for BootDevices to be an explicit nil

### UnsetBootDevices
`func (o *BootPrecisionPolicy) UnsetBootDevices()`

UnsetBootDevices ensures that no value is present for BootDevices, not even an explicit nil
### GetConfiguredBootMode

`func (o *BootPrecisionPolicy) GetConfiguredBootMode() string`

GetConfiguredBootMode returns the ConfiguredBootMode field if non-nil, zero value otherwise.

### GetConfiguredBootModeOk

`func (o *BootPrecisionPolicy) GetConfiguredBootModeOk() (*string, bool)`

GetConfiguredBootModeOk returns a tuple with the ConfiguredBootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredBootMode

`func (o *BootPrecisionPolicy) SetConfiguredBootMode(v string)`

SetConfiguredBootMode sets ConfiguredBootMode field to given value.

### HasConfiguredBootMode

`func (o *BootPrecisionPolicy) HasConfiguredBootMode() bool`

HasConfiguredBootMode returns a boolean if a field has been set.

### GetEnforceUefiSecureBoot

`func (o *BootPrecisionPolicy) GetEnforceUefiSecureBoot() bool`

GetEnforceUefiSecureBoot returns the EnforceUefiSecureBoot field if non-nil, zero value otherwise.

### GetEnforceUefiSecureBootOk

`func (o *BootPrecisionPolicy) GetEnforceUefiSecureBootOk() (*bool, bool)`

GetEnforceUefiSecureBootOk returns a tuple with the EnforceUefiSecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUefiSecureBoot

`func (o *BootPrecisionPolicy) SetEnforceUefiSecureBoot(v bool)`

SetEnforceUefiSecureBoot sets EnforceUefiSecureBoot field to given value.

### HasEnforceUefiSecureBoot

`func (o *BootPrecisionPolicy) HasEnforceUefiSecureBoot() bool`

HasEnforceUefiSecureBoot returns a boolean if a field has been set.

### GetOrganization

`func (o *BootPrecisionPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BootPrecisionPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BootPrecisionPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BootPrecisionPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *BootPrecisionPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BootPrecisionPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BootPrecisionPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BootPrecisionPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *BootPrecisionPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *BootPrecisionPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


