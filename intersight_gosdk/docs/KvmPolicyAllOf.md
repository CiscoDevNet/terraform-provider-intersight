# KvmPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableLocalServerVideo** | Pointer to **bool** | If enabled, displays KVM session on any monitor attached to the server. | [optional] 
**EnableVideoEncryption** | Pointer to **bool** | If enabled, encrypts all video information sent through KVM. | [optional] 
**Enabled** | Pointer to **bool** | State of the vKVM service on the endpoint. | [optional] 
**MaximumSessions** | Pointer to **int64** | The maximum number of concurrent KVM sessions allowed. | [optional] 
**RemotePort** | Pointer to **int64** | The port used for KVM communication. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewKvmPolicyAllOf

`func NewKvmPolicyAllOf() *KvmPolicyAllOf`

NewKvmPolicyAllOf instantiates a new KvmPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmPolicyAllOfWithDefaults

`func NewKvmPolicyAllOfWithDefaults() *KvmPolicyAllOf`

NewKvmPolicyAllOfWithDefaults instantiates a new KvmPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableLocalServerVideo

`func (o *KvmPolicyAllOf) GetEnableLocalServerVideo() bool`

GetEnableLocalServerVideo returns the EnableLocalServerVideo field if non-nil, zero value otherwise.

### GetEnableLocalServerVideoOk

`func (o *KvmPolicyAllOf) GetEnableLocalServerVideoOk() (*bool, bool)`

GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalServerVideo

`func (o *KvmPolicyAllOf) SetEnableLocalServerVideo(v bool)`

SetEnableLocalServerVideo sets EnableLocalServerVideo field to given value.

### HasEnableLocalServerVideo

`func (o *KvmPolicyAllOf) HasEnableLocalServerVideo() bool`

HasEnableLocalServerVideo returns a boolean if a field has been set.

### GetEnableVideoEncryption

`func (o *KvmPolicyAllOf) GetEnableVideoEncryption() bool`

GetEnableVideoEncryption returns the EnableVideoEncryption field if non-nil, zero value otherwise.

### GetEnableVideoEncryptionOk

`func (o *KvmPolicyAllOf) GetEnableVideoEncryptionOk() (*bool, bool)`

GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoEncryption

`func (o *KvmPolicyAllOf) SetEnableVideoEncryption(v bool)`

SetEnableVideoEncryption sets EnableVideoEncryption field to given value.

### HasEnableVideoEncryption

`func (o *KvmPolicyAllOf) HasEnableVideoEncryption() bool`

HasEnableVideoEncryption returns a boolean if a field has been set.

### GetEnabled

`func (o *KvmPolicyAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KvmPolicyAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KvmPolicyAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KvmPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumSessions

`func (o *KvmPolicyAllOf) GetMaximumSessions() int64`

GetMaximumSessions returns the MaximumSessions field if non-nil, zero value otherwise.

### GetMaximumSessionsOk

`func (o *KvmPolicyAllOf) GetMaximumSessionsOk() (*int64, bool)`

GetMaximumSessionsOk returns a tuple with the MaximumSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSessions

`func (o *KvmPolicyAllOf) SetMaximumSessions(v int64)`

SetMaximumSessions sets MaximumSessions field to given value.

### HasMaximumSessions

`func (o *KvmPolicyAllOf) HasMaximumSessions() bool`

HasMaximumSessions returns a boolean if a field has been set.

### GetRemotePort

`func (o *KvmPolicyAllOf) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *KvmPolicyAllOf) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *KvmPolicyAllOf) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *KvmPolicyAllOf) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetOrganization

`func (o *KvmPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KvmPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KvmPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KvmPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *KvmPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *KvmPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *KvmPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *KvmPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *KvmPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *KvmPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


