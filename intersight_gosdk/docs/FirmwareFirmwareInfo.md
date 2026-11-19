# FirmwareFirmwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.FirmwareInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.FirmwareInfo"]
**RunningVersion** | Pointer to **string** | The firmware version currently running on the component. | [optional] [readonly] 

## Methods

### NewFirmwareFirmwareInfo

`func NewFirmwareFirmwareInfo(classId string, objectType string, ) *FirmwareFirmwareInfo`

NewFirmwareFirmwareInfo instantiates a new FirmwareFirmwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFirmwareInfoWithDefaults

`func NewFirmwareFirmwareInfoWithDefaults() *FirmwareFirmwareInfo`

NewFirmwareFirmwareInfoWithDefaults instantiates a new FirmwareFirmwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareFirmwareInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareFirmwareInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareFirmwareInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareFirmwareInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareFirmwareInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareFirmwareInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRunningVersion

`func (o *FirmwareFirmwareInfo) GetRunningVersion() string`

GetRunningVersion returns the RunningVersion field if non-nil, zero value otherwise.

### GetRunningVersionOk

`func (o *FirmwareFirmwareInfo) GetRunningVersionOk() (*string, bool)`

GetRunningVersionOk returns a tuple with the RunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningVersion

`func (o *FirmwareFirmwareInfo) SetRunningVersion(v string)`

SetRunningVersion sets RunningVersion field to given value.

### HasRunningVersion

`func (o *FirmwareFirmwareInfo) HasRunningVersion() bool`

HasRunningVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


