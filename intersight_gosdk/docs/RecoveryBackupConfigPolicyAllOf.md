# RecoveryBackupConfigPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.BackupConfigPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.BackupConfigPolicy"]
**BackupProfiles** | Pointer to [**[]RecoveryBackupProfileRelationship**](RecoveryBackupProfileRelationship.md) | An array of relationships to recoveryBackupProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewRecoveryBackupConfigPolicyAllOf

`func NewRecoveryBackupConfigPolicyAllOf(classId string, objectType string, ) *RecoveryBackupConfigPolicyAllOf`

NewRecoveryBackupConfigPolicyAllOf instantiates a new RecoveryBackupConfigPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupConfigPolicyAllOfWithDefaults

`func NewRecoveryBackupConfigPolicyAllOfWithDefaults() *RecoveryBackupConfigPolicyAllOf`

NewRecoveryBackupConfigPolicyAllOfWithDefaults instantiates a new RecoveryBackupConfigPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryBackupConfigPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryBackupConfigPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryBackupConfigPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryBackupConfigPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryBackupConfigPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryBackupConfigPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupProfiles

`func (o *RecoveryBackupConfigPolicyAllOf) GetBackupProfiles() []RecoveryBackupProfileRelationship`

GetBackupProfiles returns the BackupProfiles field if non-nil, zero value otherwise.

### GetBackupProfilesOk

`func (o *RecoveryBackupConfigPolicyAllOf) GetBackupProfilesOk() (*[]RecoveryBackupProfileRelationship, bool)`

GetBackupProfilesOk returns a tuple with the BackupProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfiles

`func (o *RecoveryBackupConfigPolicyAllOf) SetBackupProfiles(v []RecoveryBackupProfileRelationship)`

SetBackupProfiles sets BackupProfiles field to given value.

### HasBackupProfiles

`func (o *RecoveryBackupConfigPolicyAllOf) HasBackupProfiles() bool`

HasBackupProfiles returns a boolean if a field has been set.

### SetBackupProfilesNil

`func (o *RecoveryBackupConfigPolicyAllOf) SetBackupProfilesNil(b bool)`

 SetBackupProfilesNil sets the value for BackupProfiles to be an explicit nil

### UnsetBackupProfiles
`func (o *RecoveryBackupConfigPolicyAllOf) UnsetBackupProfiles()`

UnsetBackupProfiles ensures that no value is present for BackupProfiles, not even an explicit nil
### GetOrganization

`func (o *RecoveryBackupConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryBackupConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryBackupConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryBackupConfigPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


