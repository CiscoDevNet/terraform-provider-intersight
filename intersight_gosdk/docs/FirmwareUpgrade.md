# FirmwareUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.Upgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.Upgrade"]
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**ExcludeComponentPidList** | Pointer to [**NullableFirmwareExcludeComponentPidListType**](FirmwareExcludeComponentPidListType.md) |  | [optional] 
**ServerName** | Pointer to **string** | Name of the server on which the firmware upgrade operation is performed. | [optional] [readonly] 
**UpgradeTriggerMethod** | Pointer to **string** | The source that triggered the upgrade. Either via profile or traditional way. * &#x60;none&#x60; - Upgrade is invoked within the service. * &#x60;profileTrigger&#x60; - Upgrade is invoked from a profile deployment. | [optional] [default to "none"]
**Device** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgrade

`func NewFirmwareUpgrade(classId string, objectType string, ) *FirmwareUpgrade`

NewFirmwareUpgrade instantiates a new FirmwareUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeWithDefaults

`func NewFirmwareUpgradeWithDefaults() *FirmwareUpgrade`

NewFirmwareUpgradeWithDefaults instantiates a new FirmwareUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgrade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgrade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgrade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgrade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgrade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgrade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeComponentList

`func (o *FirmwareUpgrade) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwareUpgrade) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwareUpgrade) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwareUpgrade) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### SetExcludeComponentListNil

`func (o *FirmwareUpgrade) SetExcludeComponentListNil(b bool)`

 SetExcludeComponentListNil sets the value for ExcludeComponentList to be an explicit nil

### UnsetExcludeComponentList
`func (o *FirmwareUpgrade) UnsetExcludeComponentList()`

UnsetExcludeComponentList ensures that no value is present for ExcludeComponentList, not even an explicit nil
### GetExcludeComponentPidList

`func (o *FirmwareUpgrade) GetExcludeComponentPidList() FirmwareExcludeComponentPidListType`

GetExcludeComponentPidList returns the ExcludeComponentPidList field if non-nil, zero value otherwise.

### GetExcludeComponentPidListOk

`func (o *FirmwareUpgrade) GetExcludeComponentPidListOk() (*FirmwareExcludeComponentPidListType, bool)`

GetExcludeComponentPidListOk returns a tuple with the ExcludeComponentPidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentPidList

`func (o *FirmwareUpgrade) SetExcludeComponentPidList(v FirmwareExcludeComponentPidListType)`

SetExcludeComponentPidList sets ExcludeComponentPidList field to given value.

### HasExcludeComponentPidList

`func (o *FirmwareUpgrade) HasExcludeComponentPidList() bool`

HasExcludeComponentPidList returns a boolean if a field has been set.

### SetExcludeComponentPidListNil

`func (o *FirmwareUpgrade) SetExcludeComponentPidListNil(b bool)`

 SetExcludeComponentPidListNil sets the value for ExcludeComponentPidList to be an explicit nil

### UnsetExcludeComponentPidList
`func (o *FirmwareUpgrade) UnsetExcludeComponentPidList()`

UnsetExcludeComponentPidList ensures that no value is present for ExcludeComponentPidList, not even an explicit nil
### GetServerName

`func (o *FirmwareUpgrade) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *FirmwareUpgrade) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *FirmwareUpgrade) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *FirmwareUpgrade) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetUpgradeTriggerMethod

`func (o *FirmwareUpgrade) GetUpgradeTriggerMethod() string`

GetUpgradeTriggerMethod returns the UpgradeTriggerMethod field if non-nil, zero value otherwise.

### GetUpgradeTriggerMethodOk

`func (o *FirmwareUpgrade) GetUpgradeTriggerMethodOk() (*string, bool)`

GetUpgradeTriggerMethodOk returns a tuple with the UpgradeTriggerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTriggerMethod

`func (o *FirmwareUpgrade) SetUpgradeTriggerMethod(v string)`

SetUpgradeTriggerMethod sets UpgradeTriggerMethod field to given value.

### HasUpgradeTriggerMethod

`func (o *FirmwareUpgrade) HasUpgradeTriggerMethod() bool`

HasUpgradeTriggerMethod returns a boolean if a field has been set.

### GetDevice

`func (o *FirmwareUpgrade) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUpgrade) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUpgrade) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *FirmwareUpgrade) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *FirmwareUpgrade) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetServer

`func (o *FirmwareUpgrade) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgrade) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgrade) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgrade) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *FirmwareUpgrade) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *FirmwareUpgrade) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


