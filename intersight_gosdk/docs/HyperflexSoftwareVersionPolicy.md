# HyperflexSoftwareVersionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HxdpVersion** | Pointer to **string** | Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster. | [optional] 
**HypervisorVersion** | Pointer to **string** | Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster. | [optional] 
**ServerFirmwareVersion** | Pointer to **string** | Desired server firmware version to apply on the HyperFlex Cluster. | [optional] 
**UpgradeTypes** | Pointer to **[]string** |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**HxdpVersionInfo** | Pointer to [**SoftwareHyperflexDistributableRelationship**](software.HyperflexDistributable.Relationship.md) |  | [optional] 
**HypervisorVersionInfo** | Pointer to [**SoftwareHyperflexDistributableRelationship**](software.HyperflexDistributable.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**ServerFirmwareVersionInfo** | Pointer to [**FirmwareDistributableRelationship**](firmware.Distributable.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexSoftwareVersionPolicy

`func NewHyperflexSoftwareVersionPolicy() *HyperflexSoftwareVersionPolicy`

NewHyperflexSoftwareVersionPolicy instantiates a new HyperflexSoftwareVersionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareVersionPolicyWithDefaults

`func NewHyperflexSoftwareVersionPolicyWithDefaults() *HyperflexSoftwareVersionPolicy`

NewHyperflexSoftwareVersionPolicyWithDefaults instantiates a new HyperflexSoftwareVersionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHxdpVersion

`func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersion() string`

GetHxdpVersion returns the HxdpVersion field if non-nil, zero value otherwise.

### GetHxdpVersionOk

`func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersionOk() (*string, bool)`

GetHxdpVersionOk returns a tuple with the HxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersion

`func (o *HyperflexSoftwareVersionPolicy) SetHxdpVersion(v string)`

SetHxdpVersion sets HxdpVersion field to given value.

### HasHxdpVersion

`func (o *HyperflexSoftwareVersionPolicy) HasHxdpVersion() bool`

HasHxdpVersion returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexSoftwareVersionPolicy) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexSoftwareVersionPolicy) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetServerFirmwareVersion

`func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersion() string`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersionOk() (*string, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexSoftwareVersionPolicy) SetServerFirmwareVersion(v string)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexSoftwareVersionPolicy) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.

### GetUpgradeTypes

`func (o *HyperflexSoftwareVersionPolicy) GetUpgradeTypes() []string`

GetUpgradeTypes returns the UpgradeTypes field if non-nil, zero value otherwise.

### GetUpgradeTypesOk

`func (o *HyperflexSoftwareVersionPolicy) GetUpgradeTypesOk() (*[]string, bool)`

GetUpgradeTypesOk returns a tuple with the UpgradeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTypes

`func (o *HyperflexSoftwareVersionPolicy) SetUpgradeTypes(v []string)`

SetUpgradeTypes sets UpgradeTypes field to given value.

### HasUpgradeTypes

`func (o *HyperflexSoftwareVersionPolicy) HasUpgradeTypes() bool`

HasUpgradeTypes returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexSoftwareVersionPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexSoftwareVersionPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexSoftwareVersionPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexSoftwareVersionPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexSoftwareVersionPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexSoftwareVersionPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetHxdpVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersionInfo() SoftwareHyperflexDistributableRelationship`

GetHxdpVersionInfo returns the HxdpVersionInfo field if non-nil, zero value otherwise.

### GetHxdpVersionInfoOk

`func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool)`

GetHxdpVersionInfoOk returns a tuple with the HxdpVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) SetHxdpVersionInfo(v SoftwareHyperflexDistributableRelationship)`

SetHxdpVersionInfo sets HxdpVersionInfo field to given value.

### HasHxdpVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) HasHxdpVersionInfo() bool`

HasHxdpVersionInfo returns a boolean if a field has been set.

### GetHypervisorVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersionInfo() SoftwareHyperflexDistributableRelationship`

GetHypervisorVersionInfo returns the HypervisorVersionInfo field if non-nil, zero value otherwise.

### GetHypervisorVersionInfoOk

`func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool)`

GetHypervisorVersionInfoOk returns a tuple with the HypervisorVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) SetHypervisorVersionInfo(v SoftwareHyperflexDistributableRelationship)`

SetHypervisorVersionInfo sets HypervisorVersionInfo field to given value.

### HasHypervisorVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) HasHypervisorVersionInfo() bool`

HasHypervisorVersionInfo returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexSoftwareVersionPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexSoftwareVersionPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexSoftwareVersionPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexSoftwareVersionPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServerFirmwareVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersionInfo() FirmwareDistributableRelationship`

GetServerFirmwareVersionInfo returns the ServerFirmwareVersionInfo field if non-nil, zero value otherwise.

### GetServerFirmwareVersionInfoOk

`func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersionInfoOk() (*FirmwareDistributableRelationship, bool)`

GetServerFirmwareVersionInfoOk returns a tuple with the ServerFirmwareVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) SetServerFirmwareVersionInfo(v FirmwareDistributableRelationship)`

SetServerFirmwareVersionInfo sets ServerFirmwareVersionInfo field to given value.

### HasServerFirmwareVersionInfo

`func (o *HyperflexSoftwareVersionPolicy) HasServerFirmwareVersionInfo() bool`

HasServerFirmwareVersionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


