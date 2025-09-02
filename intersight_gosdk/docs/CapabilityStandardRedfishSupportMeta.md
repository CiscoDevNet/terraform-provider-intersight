# CapabilityStandardRedfishSupportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.StandardRedfishSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.StandardRedfishSupportMeta"]
**Description** | Pointer to **string** | Verbose description regarding this group of platform. | [optional] [readonly] 
**PlatformType** | Pointer to **[]string** |  | [optional] 
**SeriesId** | Pointer to **string** | Classification of a set of server platform type. | [optional] [readonly] 

## Methods

### NewCapabilityStandardRedfishSupportMeta

`func NewCapabilityStandardRedfishSupportMeta(classId string, objectType string, ) *CapabilityStandardRedfishSupportMeta`

NewCapabilityStandardRedfishSupportMeta instantiates a new CapabilityStandardRedfishSupportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityStandardRedfishSupportMetaWithDefaults

`func NewCapabilityStandardRedfishSupportMetaWithDefaults() *CapabilityStandardRedfishSupportMeta`

NewCapabilityStandardRedfishSupportMetaWithDefaults instantiates a new CapabilityStandardRedfishSupportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityStandardRedfishSupportMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityStandardRedfishSupportMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityStandardRedfishSupportMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityStandardRedfishSupportMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityStandardRedfishSupportMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityStandardRedfishSupportMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityStandardRedfishSupportMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityStandardRedfishSupportMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityStandardRedfishSupportMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityStandardRedfishSupportMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatformType

`func (o *CapabilityStandardRedfishSupportMeta) GetPlatformType() []string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *CapabilityStandardRedfishSupportMeta) GetPlatformTypeOk() (*[]string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *CapabilityStandardRedfishSupportMeta) SetPlatformType(v []string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *CapabilityStandardRedfishSupportMeta) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformTypeNil

`func (o *CapabilityStandardRedfishSupportMeta) SetPlatformTypeNil(b bool)`

 SetPlatformTypeNil sets the value for PlatformType to be an explicit nil

### UnsetPlatformType
`func (o *CapabilityStandardRedfishSupportMeta) UnsetPlatformType()`

UnsetPlatformType ensures that no value is present for PlatformType, not even an explicit nil
### GetSeriesId

`func (o *CapabilityStandardRedfishSupportMeta) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *CapabilityStandardRedfishSupportMeta) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *CapabilityStandardRedfishSupportMeta) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *CapabilityStandardRedfishSupportMeta) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


