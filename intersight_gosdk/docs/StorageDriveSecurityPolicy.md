# StorageDriveSecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.DriveSecurityPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.DriveSecurityPolicy"]
**KeySetting** | Pointer to [**NullableStorageKeySetting**](StorageKeySetting.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewStorageDriveSecurityPolicy

`func NewStorageDriveSecurityPolicy(classId string, objectType string, ) *StorageDriveSecurityPolicy`

NewStorageDriveSecurityPolicy instantiates a new StorageDriveSecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDriveSecurityPolicyWithDefaults

`func NewStorageDriveSecurityPolicyWithDefaults() *StorageDriveSecurityPolicy`

NewStorageDriveSecurityPolicyWithDefaults instantiates a new StorageDriveSecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageDriveSecurityPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageDriveSecurityPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageDriveSecurityPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageDriveSecurityPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageDriveSecurityPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageDriveSecurityPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKeySetting

`func (o *StorageDriveSecurityPolicy) GetKeySetting() StorageKeySetting`

GetKeySetting returns the KeySetting field if non-nil, zero value otherwise.

### GetKeySettingOk

`func (o *StorageDriveSecurityPolicy) GetKeySettingOk() (*StorageKeySetting, bool)`

GetKeySettingOk returns a tuple with the KeySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySetting

`func (o *StorageDriveSecurityPolicy) SetKeySetting(v StorageKeySetting)`

SetKeySetting sets KeySetting field to given value.

### HasKeySetting

`func (o *StorageDriveSecurityPolicy) HasKeySetting() bool`

HasKeySetting returns a boolean if a field has been set.

### SetKeySettingNil

`func (o *StorageDriveSecurityPolicy) SetKeySettingNil(b bool)`

 SetKeySettingNil sets the value for KeySetting to be an explicit nil

### UnsetKeySetting
`func (o *StorageDriveSecurityPolicy) UnsetKeySetting()`

UnsetKeySetting ensures that no value is present for KeySetting, not even an explicit nil
### GetOrganization

`func (o *StorageDriveSecurityPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *StorageDriveSecurityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *StorageDriveSecurityPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *StorageDriveSecurityPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *StorageDriveSecurityPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *StorageDriveSecurityPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *StorageDriveSecurityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *StorageDriveSecurityPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *StorageDriveSecurityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *StorageDriveSecurityPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *StorageDriveSecurityPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *StorageDriveSecurityPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


