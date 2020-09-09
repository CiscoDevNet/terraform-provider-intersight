# SyslogPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalClients** | Pointer to [**[]SyslogLocalClientBase**](syslog.LocalClientBase.md) |  | [optional] 
**RemoteClients** | Pointer to [**[]SyslogRemoteClientBase**](syslog.RemoteClientBase.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSyslogPolicyAllOf

`func NewSyslogPolicyAllOf() *SyslogPolicyAllOf`

NewSyslogPolicyAllOf instantiates a new SyslogPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogPolicyAllOfWithDefaults

`func NewSyslogPolicyAllOfWithDefaults() *SyslogPolicyAllOf`

NewSyslogPolicyAllOfWithDefaults instantiates a new SyslogPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalClients

`func (o *SyslogPolicyAllOf) GetLocalClients() []SyslogLocalClientBase`

GetLocalClients returns the LocalClients field if non-nil, zero value otherwise.

### GetLocalClientsOk

`func (o *SyslogPolicyAllOf) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool)`

GetLocalClientsOk returns a tuple with the LocalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalClients

`func (o *SyslogPolicyAllOf) SetLocalClients(v []SyslogLocalClientBase)`

SetLocalClients sets LocalClients field to given value.

### HasLocalClients

`func (o *SyslogPolicyAllOf) HasLocalClients() bool`

HasLocalClients returns a boolean if a field has been set.

### GetRemoteClients

`func (o *SyslogPolicyAllOf) GetRemoteClients() []SyslogRemoteClientBase`

GetRemoteClients returns the RemoteClients field if non-nil, zero value otherwise.

### GetRemoteClientsOk

`func (o *SyslogPolicyAllOf) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool)`

GetRemoteClientsOk returns a tuple with the RemoteClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClients

`func (o *SyslogPolicyAllOf) SetRemoteClients(v []SyslogRemoteClientBase)`

SetRemoteClients sets RemoteClients field to given value.

### HasRemoteClients

`func (o *SyslogPolicyAllOf) HasRemoteClients() bool`

HasRemoteClients returns a boolean if a field has been set.

### GetOrganization

`func (o *SyslogPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SyslogPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SyslogPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SyslogPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SyslogPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SyslogPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SyslogPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SyslogPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SyslogPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SyslogPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


