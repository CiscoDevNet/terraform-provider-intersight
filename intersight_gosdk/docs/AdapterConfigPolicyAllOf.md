# AdapterConfigPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.ConfigPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.ConfigPolicy"]
**Settings** | Pointer to [**[]AdapterAdapterConfig**](AdapterAdapterConfig.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewAdapterConfigPolicyAllOf

`func NewAdapterConfigPolicyAllOf(classId string, objectType string, ) *AdapterConfigPolicyAllOf`

NewAdapterConfigPolicyAllOf instantiates a new AdapterConfigPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterConfigPolicyAllOfWithDefaults

`func NewAdapterConfigPolicyAllOfWithDefaults() *AdapterConfigPolicyAllOf`

NewAdapterConfigPolicyAllOfWithDefaults instantiates a new AdapterConfigPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterConfigPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterConfigPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterConfigPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterConfigPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterConfigPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterConfigPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSettings

`func (o *AdapterConfigPolicyAllOf) GetSettings() []AdapterAdapterConfig`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AdapterConfigPolicyAllOf) GetSettingsOk() (*[]AdapterAdapterConfig, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AdapterConfigPolicyAllOf) SetSettings(v []AdapterAdapterConfig)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AdapterConfigPolicyAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *AdapterConfigPolicyAllOf) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *AdapterConfigPolicyAllOf) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetOrganization

`func (o *AdapterConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AdapterConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AdapterConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AdapterConfigPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *AdapterConfigPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *AdapterConfigPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *AdapterConfigPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *AdapterConfigPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *AdapterConfigPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *AdapterConfigPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


