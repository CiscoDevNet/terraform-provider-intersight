# FabricUdldGlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.UdldGlobalSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.UdldGlobalSettings"]
**MessageInterval** | Pointer to **int64** | Configures the time between UDLD probe messages on ports that are in advertisement mode and are currently determined to be bidirectional. Valid values are from 7 to 90 seconds. | [optional] [default to 15]
**RecoveryAction** | Pointer to **string** | UDLD recovery when enabled, attempts to bring an UDLD error-disabled port out of reset. * &#x60;none&#x60; - The standard 4th generation UCS Fabric Interconnect with 54 ports. * &#x60;reset&#x60; - The expanded 4th generation UCS Fabric Interconnect with 108 ports. | [optional] [default to "none"]

## Methods

### NewFabricUdldGlobalSettings

`func NewFabricUdldGlobalSettings(classId string, objectType string, ) *FabricUdldGlobalSettings`

NewFabricUdldGlobalSettings instantiates a new FabricUdldGlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricUdldGlobalSettingsWithDefaults

`func NewFabricUdldGlobalSettingsWithDefaults() *FabricUdldGlobalSettings`

NewFabricUdldGlobalSettingsWithDefaults instantiates a new FabricUdldGlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricUdldGlobalSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricUdldGlobalSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricUdldGlobalSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricUdldGlobalSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricUdldGlobalSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricUdldGlobalSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessageInterval

`func (o *FabricUdldGlobalSettings) GetMessageInterval() int64`

GetMessageInterval returns the MessageInterval field if non-nil, zero value otherwise.

### GetMessageIntervalOk

`func (o *FabricUdldGlobalSettings) GetMessageIntervalOk() (*int64, bool)`

GetMessageIntervalOk returns a tuple with the MessageInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageInterval

`func (o *FabricUdldGlobalSettings) SetMessageInterval(v int64)`

SetMessageInterval sets MessageInterval field to given value.

### HasMessageInterval

`func (o *FabricUdldGlobalSettings) HasMessageInterval() bool`

HasMessageInterval returns a boolean if a field has been set.

### GetRecoveryAction

`func (o *FabricUdldGlobalSettings) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *FabricUdldGlobalSettings) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *FabricUdldGlobalSettings) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.

### HasRecoveryAction

`func (o *FabricUdldGlobalSettings) HasRecoveryAction() bool`

HasRecoveryAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


