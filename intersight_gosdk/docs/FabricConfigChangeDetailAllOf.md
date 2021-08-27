# FabricConfigChangeDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ConfigChangeDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ConfigChangeDetail"]
**Profile** | Pointer to [**FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) |  | [optional] 

## Methods

### NewFabricConfigChangeDetailAllOf

`func NewFabricConfigChangeDetailAllOf(classId string, objectType string, ) *FabricConfigChangeDetailAllOf`

NewFabricConfigChangeDetailAllOf instantiates a new FabricConfigChangeDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigChangeDetailAllOfWithDefaults

`func NewFabricConfigChangeDetailAllOfWithDefaults() *FabricConfigChangeDetailAllOf`

NewFabricConfigChangeDetailAllOfWithDefaults instantiates a new FabricConfigChangeDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricConfigChangeDetailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricConfigChangeDetailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricConfigChangeDetailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricConfigChangeDetailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricConfigChangeDetailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricConfigChangeDetailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProfile

`func (o *FabricConfigChangeDetailAllOf) GetProfile() FabricSwitchProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FabricConfigChangeDetailAllOf) GetProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FabricConfigChangeDetailAllOf) SetProfile(v FabricSwitchProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FabricConfigChangeDetailAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


