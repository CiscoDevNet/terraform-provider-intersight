# SyslogPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "syslog.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "syslog.Policy"]
**LocalClients** | Pointer to [**[]SyslogLocalClientBase**](SyslogLocalClientBase.md) |  | [optional] 
**RemoteClients** | Pointer to [**[]SyslogRemoteClientBase**](SyslogRemoteClientBase.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSyslogPolicy

`func NewSyslogPolicy(classId string, objectType string, ) *SyslogPolicy`

NewSyslogPolicy instantiates a new SyslogPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogPolicyWithDefaults

`func NewSyslogPolicyWithDefaults() *SyslogPolicy`

NewSyslogPolicyWithDefaults instantiates a new SyslogPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SyslogPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SyslogPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SyslogPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SyslogPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SyslogPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SyslogPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetLocalClientsNil

`func (o *SyslogPolicy) SetLocalClientsNil(b bool)`

 SetLocalClientsNil sets the value for LocalClients to be an explicit nil

### UnsetLocalClients
`func (o *SyslogPolicy) UnsetLocalClients()`

UnsetLocalClients ensures that no value is present for LocalClients, not even an explicit nil
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

### SetRemoteClientsNil

`func (o *SyslogPolicy) SetRemoteClientsNil(b bool)`

 SetRemoteClientsNil sets the value for RemoteClients to be an explicit nil

### UnsetRemoteClients
`func (o *SyslogPolicy) UnsetRemoteClients()`

UnsetRemoteClients ensures that no value is present for RemoteClients, not even an explicit nil
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


