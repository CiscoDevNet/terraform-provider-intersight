# ComputeServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminLocatorLedState** | Pointer to **string** | User configured state of the locator LED for the server. * &#x60;None&#x60; - No operation property for locator led. * &#x60;On&#x60; - The Locator Led is turned on. * &#x60;Off&#x60; - The Locator Led is turned off. | [optional] [default to "None"]
**AdminPowerState** | Pointer to **string** | User configured power state of the server. * &#x60;Policy&#x60; - Power state is set to the default value in the policy. * &#x60;PowerOn&#x60; - Power state of the server is set to On. * &#x60;PowerOff&#x60; - Power state is the server set to Off. * &#x60;PowerCycle&#x60; - Power state the server is reset. * &#x60;HardReset&#x60; - Power state the server is hard reset. * &#x60;Shutdown&#x60; - Operating system on the server is shut down. * &#x60;Reboot&#x60; - Power state of IMC is rebooted. | [optional] [default to "Policy"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "Applied"]
**OneTimeBootDevice** | Pointer to **string** | The name of the device chosen by user for configuring One-Time Boot device. | [optional] 
**PersistentMemoryOperation** | Pointer to [**ComputePersistentMemoryOperation**](compute.PersistentMemoryOperation.md) |  | [optional] 
**ServerConfig** | Pointer to [**ComputeServerConfig**](compute.ServerConfig.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningWorkflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewComputeServerSetting

`func NewComputeServerSetting() *ComputeServerSetting`

NewComputeServerSetting instantiates a new ComputeServerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeServerSettingWithDefaults

`func NewComputeServerSettingWithDefaults() *ComputeServerSetting`

NewComputeServerSettingWithDefaults instantiates a new ComputeServerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminLocatorLedState

`func (o *ComputeServerSetting) GetAdminLocatorLedState() string`

GetAdminLocatorLedState returns the AdminLocatorLedState field if non-nil, zero value otherwise.

### GetAdminLocatorLedStateOk

`func (o *ComputeServerSetting) GetAdminLocatorLedStateOk() (*string, bool)`

GetAdminLocatorLedStateOk returns a tuple with the AdminLocatorLedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedState

`func (o *ComputeServerSetting) SetAdminLocatorLedState(v string)`

SetAdminLocatorLedState sets AdminLocatorLedState field to given value.

### HasAdminLocatorLedState

`func (o *ComputeServerSetting) HasAdminLocatorLedState() bool`

HasAdminLocatorLedState returns a boolean if a field has been set.

### GetAdminPowerState

`func (o *ComputeServerSetting) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputeServerSetting) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputeServerSetting) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputeServerSetting) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetConfigState

`func (o *ComputeServerSetting) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *ComputeServerSetting) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *ComputeServerSetting) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *ComputeServerSetting) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetOneTimeBootDevice

`func (o *ComputeServerSetting) GetOneTimeBootDevice() string`

GetOneTimeBootDevice returns the OneTimeBootDevice field if non-nil, zero value otherwise.

### GetOneTimeBootDeviceOk

`func (o *ComputeServerSetting) GetOneTimeBootDeviceOk() (*string, bool)`

GetOneTimeBootDeviceOk returns a tuple with the OneTimeBootDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeBootDevice

`func (o *ComputeServerSetting) SetOneTimeBootDevice(v string)`

SetOneTimeBootDevice sets OneTimeBootDevice field to given value.

### HasOneTimeBootDevice

`func (o *ComputeServerSetting) HasOneTimeBootDevice() bool`

HasOneTimeBootDevice returns a boolean if a field has been set.

### GetPersistentMemoryOperation

`func (o *ComputeServerSetting) GetPersistentMemoryOperation() ComputePersistentMemoryOperation`

GetPersistentMemoryOperation returns the PersistentMemoryOperation field if non-nil, zero value otherwise.

### GetPersistentMemoryOperationOk

`func (o *ComputeServerSetting) GetPersistentMemoryOperationOk() (*ComputePersistentMemoryOperation, bool)`

GetPersistentMemoryOperationOk returns a tuple with the PersistentMemoryOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryOperation

`func (o *ComputeServerSetting) SetPersistentMemoryOperation(v ComputePersistentMemoryOperation)`

SetPersistentMemoryOperation sets PersistentMemoryOperation field to given value.

### HasPersistentMemoryOperation

`func (o *ComputeServerSetting) HasPersistentMemoryOperation() bool`

HasPersistentMemoryOperation returns a boolean if a field has been set.

### GetServerConfig

`func (o *ComputeServerSetting) GetServerConfig() ComputeServerConfig`

GetServerConfig returns the ServerConfig field if non-nil, zero value otherwise.

### GetServerConfigOk

`func (o *ComputeServerSetting) GetServerConfigOk() (*ComputeServerConfig, bool)`

GetServerConfigOk returns a tuple with the ServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfig

`func (o *ComputeServerSetting) SetServerConfig(v ComputeServerConfig)`

SetServerConfig sets ServerConfig field to given value.

### HasServerConfig

`func (o *ComputeServerSetting) HasServerConfig() bool`

HasServerConfig returns a boolean if a field has been set.

### GetLocatorLed

`func (o *ComputeServerSetting) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeServerSetting) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeServerSetting) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeServerSetting) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ComputeServerSetting) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeServerSetting) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeServerSetting) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeServerSetting) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningWorkflow

`func (o *ComputeServerSetting) GetRunningWorkflow() WorkflowWorkflowInfoRelationship`

GetRunningWorkflow returns the RunningWorkflow field if non-nil, zero value otherwise.

### GetRunningWorkflowOk

`func (o *ComputeServerSetting) GetRunningWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowOk returns a tuple with the RunningWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflow

`func (o *ComputeServerSetting) SetRunningWorkflow(v WorkflowWorkflowInfoRelationship)`

SetRunningWorkflow sets RunningWorkflow field to given value.

### HasRunningWorkflow

`func (o *ComputeServerSetting) HasRunningWorkflow() bool`

HasRunningWorkflow returns a boolean if a field has been set.

### GetServer

`func (o *ComputeServerSetting) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ComputeServerSetting) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ComputeServerSetting) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ComputeServerSetting) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


