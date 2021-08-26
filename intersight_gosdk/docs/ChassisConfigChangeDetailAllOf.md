# ChassisConfigChangeDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.ConfigChangeDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.ConfigChangeDetail"]
**Profile** | Pointer to [**ChassisProfileRelationship**](ChassisProfileRelationship.md) |  | [optional] 

## Methods

### NewChassisConfigChangeDetailAllOf

`func NewChassisConfigChangeDetailAllOf(classId string, objectType string, ) *ChassisConfigChangeDetailAllOf`

NewChassisConfigChangeDetailAllOf instantiates a new ChassisConfigChangeDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisConfigChangeDetailAllOfWithDefaults

`func NewChassisConfigChangeDetailAllOfWithDefaults() *ChassisConfigChangeDetailAllOf`

NewChassisConfigChangeDetailAllOfWithDefaults instantiates a new ChassisConfigChangeDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisConfigChangeDetailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisConfigChangeDetailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisConfigChangeDetailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisConfigChangeDetailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisConfigChangeDetailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisConfigChangeDetailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProfile

`func (o *ChassisConfigChangeDetailAllOf) GetProfile() ChassisProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChassisConfigChangeDetailAllOf) GetProfileOk() (*ChassisProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChassisConfigChangeDetailAllOf) SetProfile(v ChassisProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ChassisConfigChangeDetailAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


