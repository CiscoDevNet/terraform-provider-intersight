# HyperflexProxySettingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | HTTP Proxy server FQDN or IP. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | The password for the HTTP Proxy. | [optional] 
**Port** | Pointer to **int64** | The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive. | [optional] 
**Username** | Pointer to **string** | The username for the HTTP Proxy. | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexProxySettingPolicy

`func NewHyperflexProxySettingPolicy() *HyperflexProxySettingPolicy`

NewHyperflexProxySettingPolicy instantiates a new HyperflexProxySettingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexProxySettingPolicyWithDefaults

`func NewHyperflexProxySettingPolicyWithDefaults() *HyperflexProxySettingPolicy`

NewHyperflexProxySettingPolicyWithDefaults instantiates a new HyperflexProxySettingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *HyperflexProxySettingPolicy) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HyperflexProxySettingPolicy) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HyperflexProxySettingPolicy) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HyperflexProxySettingPolicy) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *HyperflexProxySettingPolicy) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *HyperflexProxySettingPolicy) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *HyperflexProxySettingPolicy) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *HyperflexProxySettingPolicy) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *HyperflexProxySettingPolicy) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HyperflexProxySettingPolicy) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HyperflexProxySettingPolicy) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HyperflexProxySettingPolicy) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *HyperflexProxySettingPolicy) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HyperflexProxySettingPolicy) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HyperflexProxySettingPolicy) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *HyperflexProxySettingPolicy) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *HyperflexProxySettingPolicy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HyperflexProxySettingPolicy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HyperflexProxySettingPolicy) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HyperflexProxySettingPolicy) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexProxySettingPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexProxySettingPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexProxySettingPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexProxySettingPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexProxySettingPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexProxySettingPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexProxySettingPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexProxySettingPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexProxySettingPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexProxySettingPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


