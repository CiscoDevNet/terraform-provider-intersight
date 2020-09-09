# SyslogPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalClients** | Pointer to [**[]SyslogLocalClientBase**](syslog.LocalClientBase.md) |  | [optional] 
**RemoteClients** | Pointer to [**[]SyslogRemoteClientBase**](syslog.RemoteClientBase.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSyslogPolicy

`func NewSyslogPolicy() *SyslogPolicy`

NewSyslogPolicy instantiates a new SyslogPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogPolicyWithDefaults

`func NewSyslogPolicyWithDefaults() *SyslogPolicy`

NewSyslogPolicyWithDefaults instantiates a new SyslogPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalClients

`func (o *SyslogPolicy) GetLocalClients() []SyslogLocalClientBase`

GetLocalClients returns the LocalClients field if non-nil, zero value otherwise.

### GetLocalClientsOk

`func (o *SyslogPolicy) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool)`

GetLocalClientsOk returns a tuple with the LocalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalClients

`func (o *SyslogPolicy) SetLocalClients(v []SyslogLocalClientBase)`

SetLocalClients sets LocalClients field to given value.

### HasLocalClients

`func (o *SyslogPolicy) HasLocalClients() bool`

HasLocalClients returns a boolean if a field has been set.

### GetRemoteClients

`func (o *SyslogPolicy) GetRemoteClients() []SyslogRemoteClientBase`

GetRemoteClients returns the RemoteClients field if non-nil, zero value otherwise.

### GetRemoteClientsOk

`func (o *SyslogPolicy) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool)`

GetRemoteClientsOk returns a tuple with the RemoteClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClients

`func (o *SyslogPolicy) SetRemoteClients(v []SyslogRemoteClientBase)`

SetRemoteClients sets RemoteClients field to given value.

### HasRemoteClients

`func (o *SyslogPolicy) HasRemoteClients() bool`

HasRemoteClients returns a boolean if a field has been set.

### GetOrganization

`func (o *SyslogPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SyslogPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SyslogPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SyslogPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SyslogPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SyslogPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SyslogPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SyslogPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SyslogPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SyslogPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


