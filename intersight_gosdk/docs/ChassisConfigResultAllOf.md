# ChassisConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.ConfigResult"]
**IomProfile** | Pointer to [**ChassisIomProfileRelationship**](ChassisIomProfileRelationship.md) |  | [optional] 
**Profile** | Pointer to [**ChassisBaseProfileRelationship**](ChassisBaseProfileRelationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]ChassisConfigResultEntryRelationship**](ChassisConfigResultEntryRelationship.md) | An array of relationships to chassisConfigResultEntry resources. | [optional] 

## Methods

### NewChassisConfigResultAllOf

`func NewChassisConfigResultAllOf(classId string, objectType string, ) *ChassisConfigResultAllOf`

NewChassisConfigResultAllOf instantiates a new ChassisConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisConfigResultAllOfWithDefaults

`func NewChassisConfigResultAllOfWithDefaults() *ChassisConfigResultAllOf`

NewChassisConfigResultAllOfWithDefaults instantiates a new ChassisConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisConfigResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisConfigResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisConfigResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisConfigResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisConfigResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisConfigResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIomProfile

`func (o *ChassisConfigResultAllOf) GetIomProfile() ChassisIomProfileRelationship`

GetIomProfile returns the IomProfile field if non-nil, zero value otherwise.

### GetIomProfileOk

`func (o *ChassisConfigResultAllOf) GetIomProfileOk() (*ChassisIomProfileRelationship, bool)`

GetIomProfileOk returns a tuple with the IomProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomProfile

`func (o *ChassisConfigResultAllOf) SetIomProfile(v ChassisIomProfileRelationship)`

SetIomProfile sets IomProfile field to given value.

### HasIomProfile

`func (o *ChassisConfigResultAllOf) HasIomProfile() bool`

HasIomProfile returns a boolean if a field has been set.

### GetProfile

`func (o *ChassisConfigResultAllOf) GetProfile() ChassisBaseProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChassisConfigResultAllOf) GetProfileOk() (*ChassisBaseProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChassisConfigResultAllOf) SetProfile(v ChassisBaseProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ChassisConfigResultAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *ChassisConfigResultAllOf) GetResultEntries() []ChassisConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *ChassisConfigResultAllOf) GetResultEntriesOk() (*[]ChassisConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *ChassisConfigResultAllOf) SetResultEntries(v []ChassisConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *ChassisConfigResultAllOf) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *ChassisConfigResultAllOf) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *ChassisConfigResultAllOf) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


