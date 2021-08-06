# ComputeServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ServerSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ServerSetting"]
**AdminLocatorLedState** | Pointer to **string** | User configured state of the locator LED for the server. * &#x60;None&#x60; - No operation property for locator led. * &#x60;On&#x60; - The Locator Led is turned on. * &#x60;Off&#x60; - The Locator Led is turned off. | [optional] [default to "None"]
**AdminPowerState** | Pointer to **string** | User configured power state of the server. * &#x60;Policy&#x60; - Power state is set to the default value in the policy. * &#x60;PowerOn&#x60; - Power state of the server is set to On. * &#x60;PowerOff&#x60; - Power state is the server set to Off. * &#x60;PowerCycle&#x60; - Power state the server is reset. * &#x60;HardReset&#x60; - Power state the server is hard reset. * &#x60;Shutdown&#x60; - Operating system on the server is shut down. * &#x60;Reboot&#x60; - Power state of IMC is rebooted. | [optional] [default to "Policy"]
**CertificatesAction** | Pointer to [**NullableCertificatemanagementCertificateBase**](certificatemanagement.CertificateBase.md) |  | [optional] 
**CmosReset** | Pointer to **string** | The allowed actions on the CMOS Reset. * &#x60;Ready&#x60; - CMOS Reset operation is allowed to be done on the server in this state. * &#x60;Pending&#x60; - This indicates that the previous CMOS Reset operation on this server has not completed due to a pending power cycle. CMOS Reset operation cannot be done on the server when in this state. * &#x60;Reset&#x60; - The value that the UI/API needs to provide to trigger a CMOS Reset operation on a server. | [optional] [default to "Ready"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "Applied"]
**FrontPanelLockState** | Pointer to **string** | The allowed actions on the Front Panel Lock. * &#x60;Unlock&#x60; - Front Panel of the server is set to Unlocked state. * &#x60;Lock&#x60; - Front Panel of the server is set to Locked state. | [optional] [default to "Unlock"]
**Name** | Pointer to **string** | The property used to identify the name of the server it is associated with. | [optional] [readonly] 
**OneTimeBootDevice** | Pointer to **string** | The name of the device chosen by user for configuring One-Time Boot device. | [optional] 
**PersistentMemoryOperation** | Pointer to [**NullableComputePersistentMemoryOperation**](compute.PersistentMemoryOperation.md) |  | [optional] 
**ServerConfig** | Pointer to [**NullableComputeServerConfig**](compute.ServerConfig.md) |  | [optional] 
**StorageControllerOperation** | Pointer to [**NullableComputeStorageControllerOperation**](compute.StorageControllerOperation.md) |  | [optional] 
**StoragePhysicalDriveOperation** | Pointer to [**NullableComputeStoragePhysicalDriveOperation**](compute.StoragePhysicalDriveOperation.md) |  | [optional] 
**StorageVirtualDriveOperation** | Pointer to [**NullableComputeStorageVirtualDriveOperation**](compute.StorageVirtualDriveOperation.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningWorkflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewComputeServerSetting

`func NewComputeServerSetting(classId string, objectType string, ) *ComputeServerSetting`

NewComputeServerSetting instantiates a new ComputeServerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeServerSettingWithDefaults

`func NewComputeServerSettingWithDefaults() *ComputeServerSetting`

NewComputeServerSettingWithDefaults instantiates a new ComputeServerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeServerSetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeServerSetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeServerSetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeServerSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeServerSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeServerSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetCertificatesAction

`func (o *ComputeServerSetting) GetCertificatesAction() CertificatemanagementCertificateBase`

GetCertificatesAction returns the CertificatesAction field if non-nil, zero value otherwise.

### GetCertificatesActionOk

`func (o *ComputeServerSetting) GetCertificatesActionOk() (*CertificatemanagementCertificateBase, bool)`

GetCertificatesActionOk returns a tuple with the CertificatesAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesAction

`func (o *ComputeServerSetting) SetCertificatesAction(v CertificatemanagementCertificateBase)`

SetCertificatesAction sets CertificatesAction field to given value.

### HasCertificatesAction

`func (o *ComputeServerSetting) HasCertificatesAction() bool`

HasCertificatesAction returns a boolean if a field has been set.

### SetCertificatesActionNil

`func (o *ComputeServerSetting) SetCertificatesActionNil(b bool)`

 SetCertificatesActionNil sets the value for CertificatesAction to be an explicit nil

### UnsetCertificatesAction
`func (o *ComputeServerSetting) UnsetCertificatesAction()`

UnsetCertificatesAction ensures that no value is present for CertificatesAction, not even an explicit nil
### GetCmosReset

`func (o *ComputeServerSetting) GetCmosReset() string`

GetCmosReset returns the CmosReset field if non-nil, zero value otherwise.

### GetCmosResetOk

`func (o *ComputeServerSetting) GetCmosResetOk() (*string, bool)`

GetCmosResetOk returns a tuple with the CmosReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmosReset

`func (o *ComputeServerSetting) SetCmosReset(v string)`

SetCmosReset sets CmosReset field to given value.

### HasCmosReset

`func (o *ComputeServerSetting) HasCmosReset() bool`

HasCmosReset returns a boolean if a field has been set.

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

### GetFrontPanelLockState

`func (o *ComputeServerSetting) GetFrontPanelLockState() string`

GetFrontPanelLockState returns the FrontPanelLockState field if non-nil, zero value otherwise.

### GetFrontPanelLockStateOk

`func (o *ComputeServerSetting) GetFrontPanelLockStateOk() (*string, bool)`

GetFrontPanelLockStateOk returns a tuple with the FrontPanelLockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPanelLockState

`func (o *ComputeServerSetting) SetFrontPanelLockState(v string)`

SetFrontPanelLockState sets FrontPanelLockState field to given value.

### HasFrontPanelLockState

`func (o *ComputeServerSetting) HasFrontPanelLockState() bool`

HasFrontPanelLockState returns a boolean if a field has been set.

### GetName

`func (o *ComputeServerSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeServerSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeServerSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeServerSetting) HasName() bool`

HasName returns a boolean if a field has been set.

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

### SetPersistentMemoryOperationNil

`func (o *ComputeServerSetting) SetPersistentMemoryOperationNil(b bool)`

 SetPersistentMemoryOperationNil sets the value for PersistentMemoryOperation to be an explicit nil

### UnsetPersistentMemoryOperation
`func (o *ComputeServerSetting) UnsetPersistentMemoryOperation()`

UnsetPersistentMemoryOperation ensures that no value is present for PersistentMemoryOperation, not even an explicit nil
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

### SetServerConfigNil

`func (o *ComputeServerSetting) SetServerConfigNil(b bool)`

 SetServerConfigNil sets the value for ServerConfig to be an explicit nil

### UnsetServerConfig
`func (o *ComputeServerSetting) UnsetServerConfig()`

UnsetServerConfig ensures that no value is present for ServerConfig, not even an explicit nil
### GetStorageControllerOperation

`func (o *ComputeServerSetting) GetStorageControllerOperation() ComputeStorageControllerOperation`

GetStorageControllerOperation returns the StorageControllerOperation field if non-nil, zero value otherwise.

### GetStorageControllerOperationOk

`func (o *ComputeServerSetting) GetStorageControllerOperationOk() (*ComputeStorageControllerOperation, bool)`

GetStorageControllerOperationOk returns a tuple with the StorageControllerOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerOperation

`func (o *ComputeServerSetting) SetStorageControllerOperation(v ComputeStorageControllerOperation)`

SetStorageControllerOperation sets StorageControllerOperation field to given value.

### HasStorageControllerOperation

`func (o *ComputeServerSetting) HasStorageControllerOperation() bool`

HasStorageControllerOperation returns a boolean if a field has been set.

### SetStorageControllerOperationNil

`func (o *ComputeServerSetting) SetStorageControllerOperationNil(b bool)`

 SetStorageControllerOperationNil sets the value for StorageControllerOperation to be an explicit nil

### UnsetStorageControllerOperation
`func (o *ComputeServerSetting) UnsetStorageControllerOperation()`

UnsetStorageControllerOperation ensures that no value is present for StorageControllerOperation, not even an explicit nil
### GetStoragePhysicalDriveOperation

`func (o *ComputeServerSetting) GetStoragePhysicalDriveOperation() ComputeStoragePhysicalDriveOperation`

GetStoragePhysicalDriveOperation returns the StoragePhysicalDriveOperation field if non-nil, zero value otherwise.

### GetStoragePhysicalDriveOperationOk

`func (o *ComputeServerSetting) GetStoragePhysicalDriveOperationOk() (*ComputeStoragePhysicalDriveOperation, bool)`

GetStoragePhysicalDriveOperationOk returns a tuple with the StoragePhysicalDriveOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalDriveOperation

`func (o *ComputeServerSetting) SetStoragePhysicalDriveOperation(v ComputeStoragePhysicalDriveOperation)`

SetStoragePhysicalDriveOperation sets StoragePhysicalDriveOperation field to given value.

### HasStoragePhysicalDriveOperation

`func (o *ComputeServerSetting) HasStoragePhysicalDriveOperation() bool`

HasStoragePhysicalDriveOperation returns a boolean if a field has been set.

### SetStoragePhysicalDriveOperationNil

`func (o *ComputeServerSetting) SetStoragePhysicalDriveOperationNil(b bool)`

 SetStoragePhysicalDriveOperationNil sets the value for StoragePhysicalDriveOperation to be an explicit nil

### UnsetStoragePhysicalDriveOperation
`func (o *ComputeServerSetting) UnsetStoragePhysicalDriveOperation()`

UnsetStoragePhysicalDriveOperation ensures that no value is present for StoragePhysicalDriveOperation, not even an explicit nil
### GetStorageVirtualDriveOperation

`func (o *ComputeServerSetting) GetStorageVirtualDriveOperation() ComputeStorageVirtualDriveOperation`

GetStorageVirtualDriveOperation returns the StorageVirtualDriveOperation field if non-nil, zero value otherwise.

### GetStorageVirtualDriveOperationOk

`func (o *ComputeServerSetting) GetStorageVirtualDriveOperationOk() (*ComputeStorageVirtualDriveOperation, bool)`

GetStorageVirtualDriveOperationOk returns a tuple with the StorageVirtualDriveOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualDriveOperation

`func (o *ComputeServerSetting) SetStorageVirtualDriveOperation(v ComputeStorageVirtualDriveOperation)`

SetStorageVirtualDriveOperation sets StorageVirtualDriveOperation field to given value.

### HasStorageVirtualDriveOperation

`func (o *ComputeServerSetting) HasStorageVirtualDriveOperation() bool`

HasStorageVirtualDriveOperation returns a boolean if a field has been set.

### SetStorageVirtualDriveOperationNil

`func (o *ComputeServerSetting) SetStorageVirtualDriveOperationNil(b bool)`

 SetStorageVirtualDriveOperationNil sets the value for StorageVirtualDriveOperation to be an explicit nil

### UnsetStorageVirtualDriveOperation
`func (o *ComputeServerSetting) UnsetStorageVirtualDriveOperation()`

UnsetStorageVirtualDriveOperation ensures that no value is present for StorageVirtualDriveOperation, not even an explicit nil
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


