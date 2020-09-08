# HyperflexSoftwareVersionPolicyAllOf

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

### NewHyperflexSoftwareVersionPolicyAllOf

`func NewHyperflexSoftwareVersionPolicyAllOf() *HyperflexSoftwareVersionPolicyAllOf`

NewHyperflexSoftwareVersionPolicyAllOf instantiates a new HyperflexSoftwareVersionPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareVersionPolicyAllOfWithDefaults

`func NewHyperflexSoftwareVersionPolicyAllOfWithDefaults() *HyperflexSoftwareVersionPolicyAllOf`

NewHyperflexSoftwareVersionPolicyAllOfWithDefaults instantiates a new HyperflexSoftwareVersionPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHxdpVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersion() string`

GetHxdpVersion returns the HxdpVersion field if non-nil, zero value otherwise.

### GetHxdpVersionOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersionOk() (*string, bool)`

GetHxdpVersionOk returns a tuple with the HxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetHxdpVersion(v string)`

SetHxdpVersion sets HxdpVersion field to given value.

### HasHxdpVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasHxdpVersion() bool`

HasHxdpVersion returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetServerFirmwareVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersion() string`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionOk() (*string, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetServerFirmwareVersion(v string)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.

### GetUpgradeTypes

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetUpgradeTypes() []string`

GetUpgradeTypes returns the UpgradeTypes field if non-nil, zero value otherwise.

### GetUpgradeTypesOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetUpgradeTypesOk() (*[]string, bool)`

GetUpgradeTypesOk returns a tuple with the UpgradeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTypes

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetUpgradeTypes(v []string)`

SetUpgradeTypes sets UpgradeTypes field to given value.

### HasUpgradeTypes

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasUpgradeTypes() bool`

HasUpgradeTypes returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexSoftwareVersionPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetHxdpVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersionInfo() SoftwareHyperflexDistributableRelationship`

GetHxdpVersionInfo returns the HxdpVersionInfo field if non-nil, zero value otherwise.

### GetHxdpVersionInfoOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHxdpVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool)`

GetHxdpVersionInfoOk returns a tuple with the HxdpVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetHxdpVersionInfo(v SoftwareHyperflexDistributableRelationship)`

SetHxdpVersionInfo sets HxdpVersionInfo field to given value.

### HasHxdpVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasHxdpVersionInfo() bool`

HasHxdpVersionInfo returns a boolean if a field has been set.

### GetHypervisorVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersionInfo() SoftwareHyperflexDistributableRelationship`

GetHypervisorVersionInfo returns the HypervisorVersionInfo field if non-nil, zero value otherwise.

### GetHypervisorVersionInfoOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetHypervisorVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool)`

GetHypervisorVersionInfoOk returns a tuple with the HypervisorVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetHypervisorVersionInfo(v SoftwareHyperflexDistributableRelationship)`

SetHypervisorVersionInfo sets HypervisorVersionInfo field to given value.

### HasHypervisorVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasHypervisorVersionInfo() bool`

HasHypervisorVersionInfo returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServerFirmwareVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionInfo() FirmwareDistributableRelationship`

GetServerFirmwareVersionInfo returns the ServerFirmwareVersionInfo field if non-nil, zero value otherwise.

### GetServerFirmwareVersionInfoOk

`func (o *HyperflexSoftwareVersionPolicyAllOf) GetServerFirmwareVersionInfoOk() (*FirmwareDistributableRelationship, bool)`

GetServerFirmwareVersionInfoOk returns a tuple with the ServerFirmwareVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) SetServerFirmwareVersionInfo(v FirmwareDistributableRelationship)`

SetServerFirmwareVersionInfo sets ServerFirmwareVersionInfo field to given value.

### HasServerFirmwareVersionInfo

`func (o *HyperflexSoftwareVersionPolicyAllOf) HasServerFirmwareVersionInfo() bool`

HasServerFirmwareVersionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


