# FabricMacLearningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.MacLearningSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.MacLearningSettings"]
**DisabledVlans** | Pointer to **string** | List of VLANs for which MAC learning is disabled. Applicable only when Fabric Interconnect is in Ethernet Switching Mode. | [optional] 

## Methods

### NewFabricMacLearningSettings

`func NewFabricMacLearningSettings(classId string, objectType string, ) *FabricMacLearningSettings`

NewFabricMacLearningSettings instantiates a new FabricMacLearningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricMacLearningSettingsWithDefaults

`func NewFabricMacLearningSettingsWithDefaults() *FabricMacLearningSettings`

NewFabricMacLearningSettingsWithDefaults instantiates a new FabricMacLearningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricMacLearningSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricMacLearningSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricMacLearningSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricMacLearningSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricMacLearningSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricMacLearningSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisabledVlans

`func (o *FabricMacLearningSettings) GetDisabledVlans() string`

GetDisabledVlans returns the DisabledVlans field if non-nil, zero value otherwise.

### GetDisabledVlansOk

`func (o *FabricMacLearningSettings) GetDisabledVlansOk() (*string, bool)`

GetDisabledVlansOk returns a tuple with the DisabledVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledVlans

`func (o *FabricMacLearningSettings) SetDisabledVlans(v string)`

SetDisabledVlans sets DisabledVlans field to given value.

### HasDisabledVlans

`func (o *FabricMacLearningSettings) HasDisabledVlans() bool`

HasDisabledVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


