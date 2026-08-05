# CapabilityBiosTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.BiosTokens"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.BiosTokens"]
**BiosTokensMap** | Pointer to **interface{}** | Place holder for Bios Attribute Registry. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the Bios Attribute Registry. | [optional] [readonly] 

## Methods

### NewCapabilityBiosTokens

`func NewCapabilityBiosTokens(classId string, objectType string, ) *CapabilityBiosTokens`

NewCapabilityBiosTokens instantiates a new CapabilityBiosTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityBiosTokensWithDefaults

`func NewCapabilityBiosTokensWithDefaults() *CapabilityBiosTokens`

NewCapabilityBiosTokensWithDefaults instantiates a new CapabilityBiosTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityBiosTokens) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityBiosTokens) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityBiosTokens) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityBiosTokens) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityBiosTokens) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityBiosTokens) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBiosTokensMap

`func (o *CapabilityBiosTokens) GetBiosTokensMap() interface{}`

GetBiosTokensMap returns the BiosTokensMap field if non-nil, zero value otherwise.

### GetBiosTokensMapOk

`func (o *CapabilityBiosTokens) GetBiosTokensMapOk() (*interface{}, bool)`

GetBiosTokensMapOk returns a tuple with the BiosTokensMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosTokensMap

`func (o *CapabilityBiosTokens) SetBiosTokensMap(v interface{})`

SetBiosTokensMap sets BiosTokensMap field to given value.

### HasBiosTokensMap

`func (o *CapabilityBiosTokens) HasBiosTokensMap() bool`

HasBiosTokensMap returns a boolean if a field has been set.

### SetBiosTokensMapNil

`func (o *CapabilityBiosTokens) SetBiosTokensMapNil(b bool)`

 SetBiosTokensMapNil sets the value for BiosTokensMap to be an explicit nil

### UnsetBiosTokensMap
`func (o *CapabilityBiosTokens) UnsetBiosTokensMap()`

UnsetBiosTokensMap ensures that no value is present for BiosTokensMap, not even an explicit nil
### GetVersion

`func (o *CapabilityBiosTokens) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapabilityBiosTokens) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapabilityBiosTokens) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapabilityBiosTokens) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


