# DeviceconnectorPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "deviceconnector.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "deviceconnector.Policy"]
**LockoutEnabled** | Pointer to **bool** | Enables configuration lockout on the endpoint. | [optional] [default to true]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewDeviceconnectorPolicyAllOf

`func NewDeviceconnectorPolicyAllOf(classId string, objectType string, ) *DeviceconnectorPolicyAllOf`

NewDeviceconnectorPolicyAllOf instantiates a new DeviceconnectorPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceconnectorPolicyAllOfWithDefaults

`func NewDeviceconnectorPolicyAllOfWithDefaults() *DeviceconnectorPolicyAllOf`

NewDeviceconnectorPolicyAllOfWithDefaults instantiates a new DeviceconnectorPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DeviceconnectorPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DeviceconnectorPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DeviceconnectorPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DeviceconnectorPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DeviceconnectorPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DeviceconnectorPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLockoutEnabled

`func (o *DeviceconnectorPolicyAllOf) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *DeviceconnectorPolicyAllOf) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *DeviceconnectorPolicyAllOf) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *DeviceconnectorPolicyAllOf) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetOrganization

`func (o *DeviceconnectorPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DeviceconnectorPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DeviceconnectorPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DeviceconnectorPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *DeviceconnectorPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *DeviceconnectorPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *DeviceconnectorPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *DeviceconnectorPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *DeviceconnectorPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *DeviceconnectorPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


