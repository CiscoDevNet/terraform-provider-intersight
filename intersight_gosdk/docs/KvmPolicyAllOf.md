# KvmPolicyAllOf

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
**TunneledKvmEnabled** | Pointer to **bool** | Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM. | [optional] [default to false]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewKvmPolicyAllOf

`func NewKvmPolicyAllOf(classId string, objectType string, ) *KvmPolicyAllOf`

NewKvmPolicyAllOf instantiates a new KvmPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmPolicyAllOfWithDefaults

`func NewKvmPolicyAllOfWithDefaults() *KvmPolicyAllOf`

NewKvmPolicyAllOfWithDefaults instantiates a new KvmPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetTunneledKvmEnabled

`func (o *KvmPolicyAllOf) GetTunneledKvmEnabled() bool`

GetTunneledKvmEnabled returns the TunneledKvmEnabled field if non-nil, zero value otherwise.

### GetTunneledKvmEnabledOk

`func (o *KvmPolicyAllOf) GetTunneledKvmEnabledOk() (*bool, bool)`

GetTunneledKvmEnabledOk returns a tuple with the TunneledKvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmEnabled

`func (o *KvmPolicyAllOf) SetTunneledKvmEnabled(v bool)`

SetTunneledKvmEnabled sets TunneledKvmEnabled field to given value.

### HasTunneledKvmEnabled

`func (o *KvmPolicyAllOf) HasTunneledKvmEnabled() bool`

HasTunneledKvmEnabled returns a boolean if a field has been set.

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


