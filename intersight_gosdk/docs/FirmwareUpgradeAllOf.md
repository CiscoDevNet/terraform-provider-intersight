# FirmwareUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.Upgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.Upgrade"]
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**ExcludeComponentPidList** | Pointer to [**NullableFirmwareExcludeComponentPidListType**](FirmwareExcludeComponentPidListType.md) |  | [optional] 
**UpgradeTriggerMethod** | Pointer to **string** | The source that triggered the upgrade. Either via profile or traditional way. * &#x60;none&#x60; - Upgrade is invoked within the service. * &#x60;profileTrigger&#x60; - Upgrade is invoked from a profile deployment. | [optional] [default to "none"]
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeAllOf

`func NewFirmwareUpgradeAllOf(classId string, objectType string, ) *FirmwareUpgradeAllOf`

NewFirmwareUpgradeAllOf instantiates a new FirmwareUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeAllOfWithDefaults

`func NewFirmwareUpgradeAllOfWithDefaults() *FirmwareUpgradeAllOf`

NewFirmwareUpgradeAllOfWithDefaults instantiates a new FirmwareUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgradeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgradeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeComponentList

`func (o *FirmwareUpgradeAllOf) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwareUpgradeAllOf) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwareUpgradeAllOf) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwareUpgradeAllOf) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### SetExcludeComponentListNil

`func (o *FirmwareUpgradeAllOf) SetExcludeComponentListNil(b bool)`

 SetExcludeComponentListNil sets the value for ExcludeComponentList to be an explicit nil

### UnsetExcludeComponentList
`func (o *FirmwareUpgradeAllOf) UnsetExcludeComponentList()`

UnsetExcludeComponentList ensures that no value is present for ExcludeComponentList, not even an explicit nil
### GetExcludeComponentPidList

`func (o *FirmwareUpgradeAllOf) GetExcludeComponentPidList() FirmwareExcludeComponentPidListType`

GetExcludeComponentPidList returns the ExcludeComponentPidList field if non-nil, zero value otherwise.

### GetExcludeComponentPidListOk

`func (o *FirmwareUpgradeAllOf) GetExcludeComponentPidListOk() (*FirmwareExcludeComponentPidListType, bool)`

GetExcludeComponentPidListOk returns a tuple with the ExcludeComponentPidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentPidList

`func (o *FirmwareUpgradeAllOf) SetExcludeComponentPidList(v FirmwareExcludeComponentPidListType)`

SetExcludeComponentPidList sets ExcludeComponentPidList field to given value.

### HasExcludeComponentPidList

`func (o *FirmwareUpgradeAllOf) HasExcludeComponentPidList() bool`

HasExcludeComponentPidList returns a boolean if a field has been set.

### SetExcludeComponentPidListNil

`func (o *FirmwareUpgradeAllOf) SetExcludeComponentPidListNil(b bool)`

 SetExcludeComponentPidListNil sets the value for ExcludeComponentPidList to be an explicit nil

### UnsetExcludeComponentPidList
`func (o *FirmwareUpgradeAllOf) UnsetExcludeComponentPidList()`

UnsetExcludeComponentPidList ensures that no value is present for ExcludeComponentPidList, not even an explicit nil
### GetUpgradeTriggerMethod

`func (o *FirmwareUpgradeAllOf) GetUpgradeTriggerMethod() string`

GetUpgradeTriggerMethod returns the UpgradeTriggerMethod field if non-nil, zero value otherwise.

### GetUpgradeTriggerMethodOk

`func (o *FirmwareUpgradeAllOf) GetUpgradeTriggerMethodOk() (*string, bool)`

GetUpgradeTriggerMethodOk returns a tuple with the UpgradeTriggerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTriggerMethod

`func (o *FirmwareUpgradeAllOf) SetUpgradeTriggerMethod(v string)`

SetUpgradeTriggerMethod sets UpgradeTriggerMethod field to given value.

### HasUpgradeTriggerMethod

`func (o *FirmwareUpgradeAllOf) HasUpgradeTriggerMethod() bool`

HasUpgradeTriggerMethod returns a boolean if a field has been set.

### GetDevice

`func (o *FirmwareUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUpgradeAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareUpgradeAllOf) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgradeAllOf) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgradeAllOf) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgradeAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


