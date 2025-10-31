# FirmwareFeatureVersionMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.FeatureVersionMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.FeatureVersionMap"]
**FeatureName** | Pointer to **string** | Name of the feature for which the version map is applicable. * &#x60;Unknown&#x60; - Unknown or Invalid feature in the equipment. * &#x60;ServerRole&#x60; - Server Role support for Fabric Interconnect Direct Hardware. * &#x60;FIAuditd&#x60; - AuditD feature for Fabric Interconnect. | [optional] [readonly] [default to "Unknown"]
**VersionMap** | Pointer to [**NullableFirmwareVersionMap**](FirmwareVersionMap.md) |  | [optional] 

## Methods

### NewFirmwareFeatureVersionMap

`func NewFirmwareFeatureVersionMap(classId string, objectType string, ) *FirmwareFeatureVersionMap`

NewFirmwareFeatureVersionMap instantiates a new FirmwareFeatureVersionMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFeatureVersionMapWithDefaults

`func NewFirmwareFeatureVersionMapWithDefaults() *FirmwareFeatureVersionMap`

NewFirmwareFeatureVersionMapWithDefaults instantiates a new FirmwareFeatureVersionMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareFeatureVersionMap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareFeatureVersionMap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareFeatureVersionMap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareFeatureVersionMap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareFeatureVersionMap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareFeatureVersionMap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureName

`func (o *FirmwareFeatureVersionMap) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FirmwareFeatureVersionMap) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FirmwareFeatureVersionMap) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *FirmwareFeatureVersionMap) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetVersionMap

`func (o *FirmwareFeatureVersionMap) GetVersionMap() FirmwareVersionMap`

GetVersionMap returns the VersionMap field if non-nil, zero value otherwise.

### GetVersionMapOk

`func (o *FirmwareFeatureVersionMap) GetVersionMapOk() (*FirmwareVersionMap, bool)`

GetVersionMapOk returns a tuple with the VersionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMap

`func (o *FirmwareFeatureVersionMap) SetVersionMap(v FirmwareVersionMap)`

SetVersionMap sets VersionMap field to given value.

### HasVersionMap

`func (o *FirmwareFeatureVersionMap) HasVersionMap() bool`

HasVersionMap returns a boolean if a field has been set.

### SetVersionMapNil

`func (o *FirmwareFeatureVersionMap) SetVersionMapNil(b bool)`

 SetVersionMapNil sets the value for VersionMap to be an explicit nil

### UnsetVersionMap
`func (o *FirmwareFeatureVersionMap) UnsetVersionMap()`

UnsetVersionMap ensures that no value is present for VersionMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


