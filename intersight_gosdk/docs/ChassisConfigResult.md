# ChassisConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.ConfigResult"]
**IomProfile** | Pointer to [**NullableChassisIomProfileRelationship**](ChassisIomProfileRelationship.md) |  | [optional] 
**Profile** | Pointer to [**NullableChassisBaseProfileRelationship**](ChassisBaseProfileRelationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]ChassisConfigResultEntryRelationship**](ChassisConfigResultEntryRelationship.md) | An array of relationships to chassisConfigResultEntry resources. | [optional] 

## Methods

### NewChassisConfigResult

`func NewChassisConfigResult(classId string, objectType string, ) *ChassisConfigResult`

NewChassisConfigResult instantiates a new ChassisConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisConfigResultWithDefaults

`func NewChassisConfigResultWithDefaults() *ChassisConfigResult`

NewChassisConfigResultWithDefaults instantiates a new ChassisConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisConfigResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisConfigResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisConfigResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisConfigResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisConfigResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisConfigResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIomProfile

`func (o *ChassisConfigResult) GetIomProfile() ChassisIomProfileRelationship`

GetIomProfile returns the IomProfile field if non-nil, zero value otherwise.

### GetIomProfileOk

`func (o *ChassisConfigResult) GetIomProfileOk() (*ChassisIomProfileRelationship, bool)`

GetIomProfileOk returns a tuple with the IomProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomProfile

`func (o *ChassisConfigResult) SetIomProfile(v ChassisIomProfileRelationship)`

SetIomProfile sets IomProfile field to given value.

### HasIomProfile

`func (o *ChassisConfigResult) HasIomProfile() bool`

HasIomProfile returns a boolean if a field has been set.

### SetIomProfileNil

`func (o *ChassisConfigResult) SetIomProfileNil(b bool)`

 SetIomProfileNil sets the value for IomProfile to be an explicit nil

### UnsetIomProfile
`func (o *ChassisConfigResult) UnsetIomProfile()`

UnsetIomProfile ensures that no value is present for IomProfile, not even an explicit nil
### GetProfile

`func (o *ChassisConfigResult) GetProfile() ChassisBaseProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChassisConfigResult) GetProfileOk() (*ChassisBaseProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChassisConfigResult) SetProfile(v ChassisBaseProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ChassisConfigResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *ChassisConfigResult) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ChassisConfigResult) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetResultEntries

`func (o *ChassisConfigResult) GetResultEntries() []ChassisConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *ChassisConfigResult) GetResultEntriesOk() (*[]ChassisConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *ChassisConfigResult) SetResultEntries(v []ChassisConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *ChassisConfigResult) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *ChassisConfigResult) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *ChassisConfigResult) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


