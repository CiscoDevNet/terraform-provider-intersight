# ChassisBaseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**TargetPlatform** | Pointer to **string** | The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;FIAttached&#x60; - Chassis which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "FIAttached"]
**ConfigResult** | Pointer to [**NullableChassisConfigResultRelationship**](ChassisConfigResultRelationship.md) |  | [optional] 
**IomProfiles** | Pointer to [**[]ChassisIomProfileRelationship**](ChassisIomProfileRelationship.md) | An array of relationships to chassisIomProfile resources. | [optional] [readonly] 

## Methods

### NewChassisBaseProfile

`func NewChassisBaseProfile(classId string, objectType string, ) *ChassisBaseProfile`

NewChassisBaseProfile instantiates a new ChassisBaseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisBaseProfileWithDefaults

`func NewChassisBaseProfileWithDefaults() *ChassisBaseProfile`

NewChassisBaseProfileWithDefaults instantiates a new ChassisBaseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisBaseProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisBaseProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisBaseProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisBaseProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisBaseProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisBaseProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetPlatform

`func (o *ChassisBaseProfile) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *ChassisBaseProfile) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *ChassisBaseProfile) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *ChassisBaseProfile) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetConfigResult

`func (o *ChassisBaseProfile) GetConfigResult() ChassisConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ChassisBaseProfile) GetConfigResultOk() (*ChassisConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ChassisBaseProfile) SetConfigResult(v ChassisConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ChassisBaseProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### SetConfigResultNil

`func (o *ChassisBaseProfile) SetConfigResultNil(b bool)`

 SetConfigResultNil sets the value for ConfigResult to be an explicit nil

### UnsetConfigResult
`func (o *ChassisBaseProfile) UnsetConfigResult()`

UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil
### GetIomProfiles

`func (o *ChassisBaseProfile) GetIomProfiles() []ChassisIomProfileRelationship`

GetIomProfiles returns the IomProfiles field if non-nil, zero value otherwise.

### GetIomProfilesOk

`func (o *ChassisBaseProfile) GetIomProfilesOk() (*[]ChassisIomProfileRelationship, bool)`

GetIomProfilesOk returns a tuple with the IomProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomProfiles

`func (o *ChassisBaseProfile) SetIomProfiles(v []ChassisIomProfileRelationship)`

SetIomProfiles sets IomProfiles field to given value.

### HasIomProfiles

`func (o *ChassisBaseProfile) HasIomProfiles() bool`

HasIomProfiles returns a boolean if a field has been set.

### SetIomProfilesNil

`func (o *ChassisBaseProfile) SetIomProfilesNil(b bool)`

 SetIomProfilesNil sets the value for IomProfiles to be an explicit nil

### UnsetIomProfiles
`func (o *ChassisBaseProfile) UnsetIomProfiles()`

UnsetIomProfiles ensures that no value is present for IomProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


