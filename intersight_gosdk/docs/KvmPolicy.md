# KvmPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.Policy"]
**EnableLocalServerVideo** | Pointer to **bool** | If enabled, displays KVM session on any monitor attached to the server. | [optional] [default to true]
**EnableVideoEncryption** | Pointer to **bool** | If enabled, encrypts all video information sent through KVM. Please note that this is no longer applicable for servers running versions 4.2 and above. | [optional] [default to true]
**Enabled** | Pointer to **bool** | State of the vKVM service on the endpoint. | [optional] [default to true]
**MaximumSessions** | Pointer to **int64** | The maximum number of concurrent KVM sessions allowed. | [optional] [default to 4]
**RemotePort** | Pointer to **int64** | The port used for KVM communication. | [optional] [default to 2068]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewKvmPolicy

`func NewKvmPolicy(classId string, objectType string, ) *KvmPolicy`

NewKvmPolicy instantiates a new KvmPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmPolicyWithDefaults

`func NewKvmPolicyWithDefaults() *KvmPolicy`

NewKvmPolicyWithDefaults instantiates a new KvmPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableLocalServerVideo

`func (o *KvmPolicy) GetEnableLocalServerVideo() bool`

GetEnableLocalServerVideo returns the EnableLocalServerVideo field if non-nil, zero value otherwise.

### GetEnableLocalServerVideoOk

`func (o *KvmPolicy) GetEnableLocalServerVideoOk() (*bool, bool)`

GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalServerVideo

`func (o *KvmPolicy) SetEnableLocalServerVideo(v bool)`

SetEnableLocalServerVideo sets EnableLocalServerVideo field to given value.

### HasEnableLocalServerVideo

`func (o *KvmPolicy) HasEnableLocalServerVideo() bool`

HasEnableLocalServerVideo returns a boolean if a field has been set.

### GetEnableVideoEncryption

`func (o *KvmPolicy) GetEnableVideoEncryption() bool`

GetEnableVideoEncryption returns the EnableVideoEncryption field if non-nil, zero value otherwise.

### GetEnableVideoEncryptionOk

`func (o *KvmPolicy) GetEnableVideoEncryptionOk() (*bool, bool)`

GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoEncryption

`func (o *KvmPolicy) SetEnableVideoEncryption(v bool)`

SetEnableVideoEncryption sets EnableVideoEncryption field to given value.

### HasEnableVideoEncryption

`func (o *KvmPolicy) HasEnableVideoEncryption() bool`

HasEnableVideoEncryption returns a boolean if a field has been set.

### GetEnabled

`func (o *KvmPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KvmPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KvmPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KvmPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumSessions

`func (o *KvmPolicy) GetMaximumSessions() int64`

GetMaximumSessions returns the MaximumSessions field if non-nil, zero value otherwise.

### GetMaximumSessionsOk

`func (o *KvmPolicy) GetMaximumSessionsOk() (*int64, bool)`

GetMaximumSessionsOk returns a tuple with the MaximumSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSessions

`func (o *KvmPolicy) SetMaximumSessions(v int64)`

SetMaximumSessions sets MaximumSessions field to given value.

### HasMaximumSessions

`func (o *KvmPolicy) HasMaximumSessions() bool`

HasMaximumSessions returns a boolean if a field has been set.

### GetRemotePort

`func (o *KvmPolicy) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *KvmPolicy) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *KvmPolicy) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *KvmPolicy) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetOrganization

`func (o *KvmPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KvmPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KvmPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KvmPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *KvmPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *KvmPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *KvmPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *KvmPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *KvmPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *KvmPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


