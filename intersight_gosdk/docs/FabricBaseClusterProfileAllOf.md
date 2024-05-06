# FabricBaseClusterProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**SwitchProfilesCount** | Pointer to **int64** | Number of switch profiles that are part of this cluster profile. | [optional] [readonly] 

## Methods

### NewFabricBaseClusterProfileAllOf

`func NewFabricBaseClusterProfileAllOf(classId string, objectType string, ) *FabricBaseClusterProfileAllOf`

NewFabricBaseClusterProfileAllOf instantiates a new FabricBaseClusterProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricBaseClusterProfileAllOfWithDefaults

`func NewFabricBaseClusterProfileAllOfWithDefaults() *FabricBaseClusterProfileAllOf`

NewFabricBaseClusterProfileAllOfWithDefaults instantiates a new FabricBaseClusterProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricBaseClusterProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricBaseClusterProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricBaseClusterProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricBaseClusterProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricBaseClusterProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricBaseClusterProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchProfilesCount

`func (o *FabricBaseClusterProfileAllOf) GetSwitchProfilesCount() int64`

GetSwitchProfilesCount returns the SwitchProfilesCount field if non-nil, zero value otherwise.

### GetSwitchProfilesCountOk

`func (o *FabricBaseClusterProfileAllOf) GetSwitchProfilesCountOk() (*int64, bool)`

GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilesCount

`func (o *FabricBaseClusterProfileAllOf) SetSwitchProfilesCount(v int64)`

SetSwitchProfilesCount sets SwitchProfilesCount field to given value.

### HasSwitchProfilesCount

`func (o *FabricBaseClusterProfileAllOf) HasSwitchProfilesCount() bool`

HasSwitchProfilesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


