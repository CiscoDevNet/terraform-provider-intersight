# ChassisBaseProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**TargetPlatform** | Pointer to **string** | The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;FIAttached&#x60; - Chassis which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "FIAttached"]
**ConfigResult** | Pointer to [**ChassisConfigResultRelationship**](ChassisConfigResultRelationship.md) |  | [optional] 
**IomProfiles** | Pointer to [**[]ChassisIomProfileRelationship**](ChassisIomProfileRelationship.md) | An array of relationships to chassisIomProfile resources. | [optional] [readonly] 

## Methods

### NewChassisBaseProfileAllOf

`func NewChassisBaseProfileAllOf(classId string, objectType string, ) *ChassisBaseProfileAllOf`

NewChassisBaseProfileAllOf instantiates a new ChassisBaseProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisBaseProfileAllOfWithDefaults

`func NewChassisBaseProfileAllOfWithDefaults() *ChassisBaseProfileAllOf`

NewChassisBaseProfileAllOfWithDefaults instantiates a new ChassisBaseProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisBaseProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisBaseProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisBaseProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisBaseProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisBaseProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisBaseProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetPlatform

`func (o *ChassisBaseProfileAllOf) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *ChassisBaseProfileAllOf) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *ChassisBaseProfileAllOf) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *ChassisBaseProfileAllOf) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetConfigResult

`func (o *ChassisBaseProfileAllOf) GetConfigResult() ChassisConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ChassisBaseProfileAllOf) GetConfigResultOk() (*ChassisConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ChassisBaseProfileAllOf) SetConfigResult(v ChassisConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ChassisBaseProfileAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetIomProfiles

`func (o *ChassisBaseProfileAllOf) GetIomProfiles() []ChassisIomProfileRelationship`

GetIomProfiles returns the IomProfiles field if non-nil, zero value otherwise.

### GetIomProfilesOk

`func (o *ChassisBaseProfileAllOf) GetIomProfilesOk() (*[]ChassisIomProfileRelationship, bool)`

GetIomProfilesOk returns a tuple with the IomProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomProfiles

`func (o *ChassisBaseProfileAllOf) SetIomProfiles(v []ChassisIomProfileRelationship)`

SetIomProfiles sets IomProfiles field to given value.

### HasIomProfiles

`func (o *ChassisBaseProfileAllOf) HasIomProfiles() bool`

HasIomProfiles returns a boolean if a field has been set.

### SetIomProfilesNil

`func (o *ChassisBaseProfileAllOf) SetIomProfilesNil(b bool)`

 SetIomProfilesNil sets the value for IomProfiles to be an explicit nil

### UnsetIomProfiles
`func (o *ChassisBaseProfileAllOf) UnsetIomProfiles()`

UnsetIomProfiles ensures that no value is present for IomProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


