# CapabilityUpdateConstraintMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**IsSecureBootSupported** | Pointer to **bool** | Flag to indicate support for secure boot. | [optional] [readonly] 
**MinSupportedVersion** | Pointer to **string** | Firmware version below which firmware update is not supported for this inventory unit. | [optional] [readonly] 
**Model** | Pointer to **string** | Model of the inventory unit which will be supported in firmware operation. | [optional] [readonly] 
**SupportedPlatforms** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityUpdateConstraintMetaAllOf

`func NewCapabilityUpdateConstraintMetaAllOf(classId string, objectType string, ) *CapabilityUpdateConstraintMetaAllOf`

NewCapabilityUpdateConstraintMetaAllOf instantiates a new CapabilityUpdateConstraintMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityUpdateConstraintMetaAllOfWithDefaults

`func NewCapabilityUpdateConstraintMetaAllOfWithDefaults() *CapabilityUpdateConstraintMetaAllOf`

NewCapabilityUpdateConstraintMetaAllOfWithDefaults instantiates a new CapabilityUpdateConstraintMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityUpdateConstraintMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityUpdateConstraintMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityUpdateConstraintMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityUpdateConstraintMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityUpdateConstraintMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityUpdateConstraintMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsSecureBootSupported

`func (o *CapabilityUpdateConstraintMetaAllOf) GetIsSecureBootSupported() bool`

GetIsSecureBootSupported returns the IsSecureBootSupported field if non-nil, zero value otherwise.

### GetIsSecureBootSupportedOk

`func (o *CapabilityUpdateConstraintMetaAllOf) GetIsSecureBootSupportedOk() (*bool, bool)`

GetIsSecureBootSupportedOk returns a tuple with the IsSecureBootSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBootSupported

`func (o *CapabilityUpdateConstraintMetaAllOf) SetIsSecureBootSupported(v bool)`

SetIsSecureBootSupported sets IsSecureBootSupported field to given value.

### HasIsSecureBootSupported

`func (o *CapabilityUpdateConstraintMetaAllOf) HasIsSecureBootSupported() bool`

HasIsSecureBootSupported returns a boolean if a field has been set.

### GetMinSupportedVersion

`func (o *CapabilityUpdateConstraintMetaAllOf) GetMinSupportedVersion() string`

GetMinSupportedVersion returns the MinSupportedVersion field if non-nil, zero value otherwise.

### GetMinSupportedVersionOk

`func (o *CapabilityUpdateConstraintMetaAllOf) GetMinSupportedVersionOk() (*string, bool)`

GetMinSupportedVersionOk returns a tuple with the MinSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSupportedVersion

`func (o *CapabilityUpdateConstraintMetaAllOf) SetMinSupportedVersion(v string)`

SetMinSupportedVersion sets MinSupportedVersion field to given value.

### HasMinSupportedVersion

`func (o *CapabilityUpdateConstraintMetaAllOf) HasMinSupportedVersion() bool`

HasMinSupportedVersion returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityUpdateConstraintMetaAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityUpdateConstraintMetaAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityUpdateConstraintMetaAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityUpdateConstraintMetaAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSupportedPlatforms

`func (o *CapabilityUpdateConstraintMetaAllOf) GetSupportedPlatforms() []string`

GetSupportedPlatforms returns the SupportedPlatforms field if non-nil, zero value otherwise.

### GetSupportedPlatformsOk

`func (o *CapabilityUpdateConstraintMetaAllOf) GetSupportedPlatformsOk() (*[]string, bool)`

GetSupportedPlatformsOk returns a tuple with the SupportedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatforms

`func (o *CapabilityUpdateConstraintMetaAllOf) SetSupportedPlatforms(v []string)`

SetSupportedPlatforms sets SupportedPlatforms field to given value.

### HasSupportedPlatforms

`func (o *CapabilityUpdateConstraintMetaAllOf) HasSupportedPlatforms() bool`

HasSupportedPlatforms returns a boolean if a field has been set.

### SetSupportedPlatformsNil

`func (o *CapabilityUpdateConstraintMetaAllOf) SetSupportedPlatformsNil(b bool)`

 SetSupportedPlatformsNil sets the value for SupportedPlatforms to be an explicit nil

### UnsetSupportedPlatforms
`func (o *CapabilityUpdateConstraintMetaAllOf) UnsetSupportedPlatforms()`

UnsetSupportedPlatforms ensures that no value is present for SupportedPlatforms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


