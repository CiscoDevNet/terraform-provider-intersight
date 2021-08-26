# FabricConfigChangeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ConfigChangeDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ConfigChangeDetail"]
**Profile** | Pointer to [**FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) |  | [optional] 

## Methods

### NewFabricConfigChangeDetail

`func NewFabricConfigChangeDetail(classId string, objectType string, ) *FabricConfigChangeDetail`

NewFabricConfigChangeDetail instantiates a new FabricConfigChangeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigChangeDetailWithDefaults

`func NewFabricConfigChangeDetailWithDefaults() *FabricConfigChangeDetail`

NewFabricConfigChangeDetailWithDefaults instantiates a new FabricConfigChangeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricConfigChangeDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricConfigChangeDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricConfigChangeDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricConfigChangeDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricConfigChangeDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricConfigChangeDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProfile

`func (o *FabricConfigChangeDetail) GetProfile() FabricSwitchProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FabricConfigChangeDetail) GetProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FabricConfigChangeDetail) SetProfile(v FabricSwitchProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FabricConfigChangeDetail) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


