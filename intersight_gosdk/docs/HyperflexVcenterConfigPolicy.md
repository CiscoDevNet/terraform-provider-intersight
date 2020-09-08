# HyperflexVcenterConfigPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCenter** | Pointer to **string** | The vCenter datacenter name. | [optional] 
**Hostname** | Pointer to **string** | The vCenter server FQDN or IP. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter. | [optional] 
**SsoUrl** | Pointer to **string** | Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC. | [optional] 
**Username** | Pointer to **string** | The vCenter username (e.g. administrator@vsphere.local). | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexVcenterConfigPolicy

`func NewHyperflexVcenterConfigPolicy() *HyperflexVcenterConfigPolicy`

NewHyperflexVcenterConfigPolicy instantiates a new HyperflexVcenterConfigPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVcenterConfigPolicyWithDefaults

`func NewHyperflexVcenterConfigPolicyWithDefaults() *HyperflexVcenterConfigPolicy`

NewHyperflexVcenterConfigPolicyWithDefaults instantiates a new HyperflexVcenterConfigPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataCenter

`func (o *HyperflexVcenterConfigPolicy) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *HyperflexVcenterConfigPolicy) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *HyperflexVcenterConfigPolicy) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.

### HasDataCenter

`func (o *HyperflexVcenterConfigPolicy) HasDataCenter() bool`

HasDataCenter returns a boolean if a field has been set.

### GetHostname

`func (o *HyperflexVcenterConfigPolicy) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HyperflexVcenterConfigPolicy) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HyperflexVcenterConfigPolicy) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HyperflexVcenterConfigPolicy) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *HyperflexVcenterConfigPolicy) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *HyperflexVcenterConfigPolicy) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *HyperflexVcenterConfigPolicy) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *HyperflexVcenterConfigPolicy) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *HyperflexVcenterConfigPolicy) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HyperflexVcenterConfigPolicy) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HyperflexVcenterConfigPolicy) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HyperflexVcenterConfigPolicy) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSsoUrl

`func (o *HyperflexVcenterConfigPolicy) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *HyperflexVcenterConfigPolicy) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *HyperflexVcenterConfigPolicy) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *HyperflexVcenterConfigPolicy) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetUsername

`func (o *HyperflexVcenterConfigPolicy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HyperflexVcenterConfigPolicy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HyperflexVcenterConfigPolicy) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HyperflexVcenterConfigPolicy) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexVcenterConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexVcenterConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexVcenterConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexVcenterConfigPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexVcenterConfigPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexVcenterConfigPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexVcenterConfigPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexVcenterConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexVcenterConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexVcenterConfigPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


