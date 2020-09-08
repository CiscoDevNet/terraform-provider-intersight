# HyperflexLocalCredentialPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactoryHypervisorPassword** | Pointer to **bool** | Indicates if Hypervisor password is the factory set default password. For HyperFlex Data Platform versions 3.0 or higher, enable this if the default password was not changed on the Hypervisor. It is required to supply a new custom Hypervisor password that will be applied to the Hypervisor during deployment. For HyperFlex Data Platform versions prior to 3.0 release, this setting has no effect and the default password will be used for initial install. The Hypervisor password should be changed after deployment. | [optional] 
**HxdpRootPwd** | Pointer to **string** | HyperFlex storage controller VM password must contain a minimum of 10 characters, with at least 1 lowercase, 1 uppercase, 1 numeric, and 1 of these -_@#$%^&amp;*! special characters. | [optional] 
**HypervisorAdmin** | Pointer to **string** | Hypervisor administrator username must contain only alphanumeric characters. Use the root account for ESXi deployments. | [optional] 
**HypervisorAdminPwd** | Pointer to **string** | The ESXi root password. For HyperFlex Data Platform 3.0 or later, if the factory default password was not manually changed, you must set a new custom password. If the password was manually changed, you must not enable the factory default password property and provide the current hypervisor password. Note - All HyperFlex nodes require the same hypervisor password for installation. For HyperFlex Data Platform prior to 3.0, use the default password \&quot;Cisco123\&quot; for newly manufactured HyperFlex servers. A custom password should only be entered if hypervisor credentials were manually changed prior to deployment. | [optional] 
**IsHxdpRootPwdSet** | Pointer to **bool** | Indicates whether the value of the &#39;hxdpRootPwd&#39; property has been set. | [optional] [readonly] 
**IsHypervisorAdminPwdSet** | Pointer to **bool** | Indicates whether the value of the &#39;hypervisorAdminPwd&#39; property has been set. | [optional] [readonly] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexLocalCredentialPolicy

`func NewHyperflexLocalCredentialPolicy() *HyperflexLocalCredentialPolicy`

NewHyperflexLocalCredentialPolicy instantiates a new HyperflexLocalCredentialPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLocalCredentialPolicyWithDefaults

`func NewHyperflexLocalCredentialPolicyWithDefaults() *HyperflexLocalCredentialPolicy`

NewHyperflexLocalCredentialPolicyWithDefaults instantiates a new HyperflexLocalCredentialPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactoryHypervisorPassword

`func (o *HyperflexLocalCredentialPolicy) GetFactoryHypervisorPassword() bool`

GetFactoryHypervisorPassword returns the FactoryHypervisorPassword field if non-nil, zero value otherwise.

### GetFactoryHypervisorPasswordOk

`func (o *HyperflexLocalCredentialPolicy) GetFactoryHypervisorPasswordOk() (*bool, bool)`

GetFactoryHypervisorPasswordOk returns a tuple with the FactoryHypervisorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryHypervisorPassword

`func (o *HyperflexLocalCredentialPolicy) SetFactoryHypervisorPassword(v bool)`

SetFactoryHypervisorPassword sets FactoryHypervisorPassword field to given value.

### HasFactoryHypervisorPassword

`func (o *HyperflexLocalCredentialPolicy) HasFactoryHypervisorPassword() bool`

HasFactoryHypervisorPassword returns a boolean if a field has been set.

### GetHxdpRootPwd

`func (o *HyperflexLocalCredentialPolicy) GetHxdpRootPwd() string`

GetHxdpRootPwd returns the HxdpRootPwd field if non-nil, zero value otherwise.

### GetHxdpRootPwdOk

`func (o *HyperflexLocalCredentialPolicy) GetHxdpRootPwdOk() (*string, bool)`

GetHxdpRootPwdOk returns a tuple with the HxdpRootPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpRootPwd

`func (o *HyperflexLocalCredentialPolicy) SetHxdpRootPwd(v string)`

SetHxdpRootPwd sets HxdpRootPwd field to given value.

### HasHxdpRootPwd

`func (o *HyperflexLocalCredentialPolicy) HasHxdpRootPwd() bool`

HasHxdpRootPwd returns a boolean if a field has been set.

### GetHypervisorAdmin

`func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdmin() string`

GetHypervisorAdmin returns the HypervisorAdmin field if non-nil, zero value otherwise.

### GetHypervisorAdminOk

`func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdminOk() (*string, bool)`

GetHypervisorAdminOk returns a tuple with the HypervisorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorAdmin

`func (o *HyperflexLocalCredentialPolicy) SetHypervisorAdmin(v string)`

SetHypervisorAdmin sets HypervisorAdmin field to given value.

### HasHypervisorAdmin

`func (o *HyperflexLocalCredentialPolicy) HasHypervisorAdmin() bool`

HasHypervisorAdmin returns a boolean if a field has been set.

### GetHypervisorAdminPwd

`func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdminPwd() string`

GetHypervisorAdminPwd returns the HypervisorAdminPwd field if non-nil, zero value otherwise.

### GetHypervisorAdminPwdOk

`func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdminPwdOk() (*string, bool)`

GetHypervisorAdminPwdOk returns a tuple with the HypervisorAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorAdminPwd

`func (o *HyperflexLocalCredentialPolicy) SetHypervisorAdminPwd(v string)`

SetHypervisorAdminPwd sets HypervisorAdminPwd field to given value.

### HasHypervisorAdminPwd

`func (o *HyperflexLocalCredentialPolicy) HasHypervisorAdminPwd() bool`

HasHypervisorAdminPwd returns a boolean if a field has been set.

### GetIsHxdpRootPwdSet

`func (o *HyperflexLocalCredentialPolicy) GetIsHxdpRootPwdSet() bool`

GetIsHxdpRootPwdSet returns the IsHxdpRootPwdSet field if non-nil, zero value otherwise.

### GetIsHxdpRootPwdSetOk

`func (o *HyperflexLocalCredentialPolicy) GetIsHxdpRootPwdSetOk() (*bool, bool)`

GetIsHxdpRootPwdSetOk returns a tuple with the IsHxdpRootPwdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHxdpRootPwdSet

`func (o *HyperflexLocalCredentialPolicy) SetIsHxdpRootPwdSet(v bool)`

SetIsHxdpRootPwdSet sets IsHxdpRootPwdSet field to given value.

### HasIsHxdpRootPwdSet

`func (o *HyperflexLocalCredentialPolicy) HasIsHxdpRootPwdSet() bool`

HasIsHxdpRootPwdSet returns a boolean if a field has been set.

### GetIsHypervisorAdminPwdSet

`func (o *HyperflexLocalCredentialPolicy) GetIsHypervisorAdminPwdSet() bool`

GetIsHypervisorAdminPwdSet returns the IsHypervisorAdminPwdSet field if non-nil, zero value otherwise.

### GetIsHypervisorAdminPwdSetOk

`func (o *HyperflexLocalCredentialPolicy) GetIsHypervisorAdminPwdSetOk() (*bool, bool)`

GetIsHypervisorAdminPwdSetOk returns a tuple with the IsHypervisorAdminPwdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHypervisorAdminPwdSet

`func (o *HyperflexLocalCredentialPolicy) SetIsHypervisorAdminPwdSet(v bool)`

SetIsHypervisorAdminPwdSet sets IsHypervisorAdminPwdSet field to given value.

### HasIsHypervisorAdminPwdSet

`func (o *HyperflexLocalCredentialPolicy) HasIsHypervisorAdminPwdSet() bool`

HasIsHypervisorAdminPwdSet returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexLocalCredentialPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexLocalCredentialPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexLocalCredentialPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexLocalCredentialPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexLocalCredentialPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexLocalCredentialPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexLocalCredentialPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexLocalCredentialPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexLocalCredentialPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexLocalCredentialPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


