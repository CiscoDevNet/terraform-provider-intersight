# FabricBaseClusterProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**SwitchProfilesCount** | Pointer to **int64** | Number of switch profiles that are part of this cluster profile. | [optional] [readonly] 
**TargetPlatform** | Pointer to **string** | Type of the profile. &#39;UcsDomain&#39; profile for network and management configuration on UCS Fabric Interconnect. &#39;UnifiedEdge&#39; profile for network, management and chassis configuration on Unified Edge. * &#x60;UCS Domain&#x60; - Profile/policy type for network and management configuration on UCS Fabric Interconnect. * &#x60;Unified Edge&#x60; - Profile/policy type for network, management and chassis configuration on Unified Edge. | [optional] [default to "UCS Domain"]

## Methods

### NewFabricBaseClusterProfile

`func NewFabricBaseClusterProfile(classId string, objectType string, ) *FabricBaseClusterProfile`

NewFabricBaseClusterProfile instantiates a new FabricBaseClusterProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricBaseClusterProfileWithDefaults

`func NewFabricBaseClusterProfileWithDefaults() *FabricBaseClusterProfile`

NewFabricBaseClusterProfileWithDefaults instantiates a new FabricBaseClusterProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricBaseClusterProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricBaseClusterProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricBaseClusterProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricBaseClusterProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricBaseClusterProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricBaseClusterProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchProfilesCount

`func (o *FabricBaseClusterProfile) GetSwitchProfilesCount() int64`

GetSwitchProfilesCount returns the SwitchProfilesCount field if non-nil, zero value otherwise.

### GetSwitchProfilesCountOk

`func (o *FabricBaseClusterProfile) GetSwitchProfilesCountOk() (*int64, bool)`

GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilesCount

`func (o *FabricBaseClusterProfile) SetSwitchProfilesCount(v int64)`

SetSwitchProfilesCount sets SwitchProfilesCount field to given value.

### HasSwitchProfilesCount

`func (o *FabricBaseClusterProfile) HasSwitchProfilesCount() bool`

HasSwitchProfilesCount returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *FabricBaseClusterProfile) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *FabricBaseClusterProfile) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *FabricBaseClusterProfile) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *FabricBaseClusterProfile) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


