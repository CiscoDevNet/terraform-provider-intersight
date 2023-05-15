# StorageDriveSecurityPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.DriveSecurityPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.DriveSecurityPolicy"]
**KeySetting** | Pointer to [**NullableStorageKeySetting**](StorageKeySetting.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewStorageDriveSecurityPolicyAllOf

`func NewStorageDriveSecurityPolicyAllOf(classId string, objectType string, ) *StorageDriveSecurityPolicyAllOf`

NewStorageDriveSecurityPolicyAllOf instantiates a new StorageDriveSecurityPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDriveSecurityPolicyAllOfWithDefaults

`func NewStorageDriveSecurityPolicyAllOfWithDefaults() *StorageDriveSecurityPolicyAllOf`

NewStorageDriveSecurityPolicyAllOfWithDefaults instantiates a new StorageDriveSecurityPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageDriveSecurityPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageDriveSecurityPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageDriveSecurityPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageDriveSecurityPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageDriveSecurityPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageDriveSecurityPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKeySetting

`func (o *StorageDriveSecurityPolicyAllOf) GetKeySetting() StorageKeySetting`

GetKeySetting returns the KeySetting field if non-nil, zero value otherwise.

### GetKeySettingOk

`func (o *StorageDriveSecurityPolicyAllOf) GetKeySettingOk() (*StorageKeySetting, bool)`

GetKeySettingOk returns a tuple with the KeySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySetting

`func (o *StorageDriveSecurityPolicyAllOf) SetKeySetting(v StorageKeySetting)`

SetKeySetting sets KeySetting field to given value.

### HasKeySetting

`func (o *StorageDriveSecurityPolicyAllOf) HasKeySetting() bool`

HasKeySetting returns a boolean if a field has been set.

### SetKeySettingNil

`func (o *StorageDriveSecurityPolicyAllOf) SetKeySettingNil(b bool)`

 SetKeySettingNil sets the value for KeySetting to be an explicit nil

### UnsetKeySetting
`func (o *StorageDriveSecurityPolicyAllOf) UnsetKeySetting()`

UnsetKeySetting ensures that no value is present for KeySetting, not even an explicit nil
### GetOrganization

`func (o *StorageDriveSecurityPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *StorageDriveSecurityPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *StorageDriveSecurityPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *StorageDriveSecurityPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *StorageDriveSecurityPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *StorageDriveSecurityPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *StorageDriveSecurityPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *StorageDriveSecurityPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *StorageDriveSecurityPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *StorageDriveSecurityPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


