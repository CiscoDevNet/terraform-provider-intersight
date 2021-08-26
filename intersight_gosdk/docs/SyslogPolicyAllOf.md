# SyslogPolicyAllOf

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

### NewSyslogPolicyAllOf

`func NewSyslogPolicyAllOf(classId string, objectType string, ) *SyslogPolicyAllOf`

NewSyslogPolicyAllOf instantiates a new SyslogPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogPolicyAllOfWithDefaults

`func NewSyslogPolicyAllOfWithDefaults() *SyslogPolicyAllOf`

NewSyslogPolicyAllOfWithDefaults instantiates a new SyslogPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SyslogPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SyslogPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SyslogPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SyslogPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SyslogPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SyslogPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetLocalClientsNil

`func (o *SyslogPolicyAllOf) SetLocalClientsNil(b bool)`

 SetLocalClientsNil sets the value for LocalClients to be an explicit nil

### UnsetLocalClients
`func (o *SyslogPolicyAllOf) UnsetLocalClients()`

UnsetLocalClients ensures that no value is present for LocalClients, not even an explicit nil
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

### SetRemoteClientsNil

`func (o *SyslogPolicyAllOf) SetRemoteClientsNil(b bool)`

 SetRemoteClientsNil sets the value for RemoteClients to be an explicit nil

### UnsetRemoteClients
`func (o *SyslogPolicyAllOf) UnsetRemoteClients()`

UnsetRemoteClients ensures that no value is present for RemoteClients, not even an explicit nil
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


