# CapabilityTimeZoneMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.TimeZoneMetaData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.TimeZoneMetaData"]
**SupportedTimeZones** | Pointer to **[]string** |  | [optional] 
**TargetType** | Pointer to **string** | Server name (e.g., Mustang, UCS). | [optional] 

## Methods

### NewCapabilityTimeZoneMetaData

`func NewCapabilityTimeZoneMetaData(classId string, objectType string, ) *CapabilityTimeZoneMetaData`

NewCapabilityTimeZoneMetaData instantiates a new CapabilityTimeZoneMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityTimeZoneMetaDataWithDefaults

`func NewCapabilityTimeZoneMetaDataWithDefaults() *CapabilityTimeZoneMetaData`

NewCapabilityTimeZoneMetaDataWithDefaults instantiates a new CapabilityTimeZoneMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityTimeZoneMetaData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityTimeZoneMetaData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityTimeZoneMetaData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityTimeZoneMetaData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityTimeZoneMetaData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityTimeZoneMetaData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSupportedTimeZones

`func (o *CapabilityTimeZoneMetaData) GetSupportedTimeZones() []string`

GetSupportedTimeZones returns the SupportedTimeZones field if non-nil, zero value otherwise.

### GetSupportedTimeZonesOk

`func (o *CapabilityTimeZoneMetaData) GetSupportedTimeZonesOk() (*[]string, bool)`

GetSupportedTimeZonesOk returns a tuple with the SupportedTimeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTimeZones

`func (o *CapabilityTimeZoneMetaData) SetSupportedTimeZones(v []string)`

SetSupportedTimeZones sets SupportedTimeZones field to given value.

### HasSupportedTimeZones

`func (o *CapabilityTimeZoneMetaData) HasSupportedTimeZones() bool`

HasSupportedTimeZones returns a boolean if a field has been set.

### SetSupportedTimeZonesNil

`func (o *CapabilityTimeZoneMetaData) SetSupportedTimeZonesNil(b bool)`

 SetSupportedTimeZonesNil sets the value for SupportedTimeZones to be an explicit nil

### UnsetSupportedTimeZones
`func (o *CapabilityTimeZoneMetaData) UnsetSupportedTimeZones()`

UnsetSupportedTimeZones ensures that no value is present for SupportedTimeZones, not even an explicit nil
### GetTargetType

`func (o *CapabilityTimeZoneMetaData) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CapabilityTimeZoneMetaData) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CapabilityTimeZoneMetaData) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CapabilityTimeZoneMetaData) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


