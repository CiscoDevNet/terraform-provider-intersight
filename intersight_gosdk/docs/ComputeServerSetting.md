# ComputeServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ServerSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ServerSetting"]
**AdminLocatorLedState** | Pointer to **string** | User configured state of the locator LED for the server. * &#x60;None&#x60; - No operation property for locator led. * &#x60;On&#x60; - The Locator Led is turned on. * &#x60;Off&#x60; - The Locator Led is turned off. | [optional] [default to "None"]
**AdminPowerState** | Pointer to **string** | User configured power state of the server. * &#x60;Policy&#x60; - Power state is set to the default value in the policy. * &#x60;PowerOn&#x60; - Power state of the server is set to On. * &#x60;PowerOff&#x60; - Power state is the server set to Off. * &#x60;PowerCycle&#x60; - Power state the server is reset. * &#x60;HardReset&#x60; - Power state the server is hard reset. * &#x60;Shutdown&#x60; - Operating system on the server is shut down. * &#x60;Reboot&#x60; - Power state of IMC is rebooted. | [optional] [default to "Policy"]
**CertificatesAction** | Pointer to [**NullableCertificatemanagementCertificateBase**](CertificatemanagementCertificateBase.md) |  | [optional] 
**ClearSel** | Pointer to **string** | Clear system event log on a server. * &#x60;Ready&#x60; - Clear system event log operation is allowed on the server in this state. * &#x60;Clear&#x60; - Trigger a clear system event log operation on a server. | [optional] [default to "Ready"]
**CmosReset** | Pointer to **string** | The allowed actions on the CMOS Reset. * &#x60;Ready&#x60; - CMOS Reset operation is allowed to be done on the server in this state. * &#x60;Pending&#x60; - The identifier to state that the previous CMOS Reset operation on this server has not completed due to a pending power cycle. CMOS Reset operation cannot be done on the server when in this state. * &#x60;Reset&#x60; - The value that the UI/API needs to provide to trigger a CMOS Reset operation on a server. | [optional] [default to "Ready"]
**CollectSel** | Pointer to **string** | Collect system event log from a server. * &#x60;Ready&#x60; - Collect system event log operation is allowed on the server in this state. * &#x60;Collect&#x60; - Trigger a collect system event log operation on a server. | [optional] [default to "Ready"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Scheduled&#x60; - User configured settings are scheduled to be applied. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "Applied"]
**FrontPanelLockState** | Pointer to **string** | The allowed actions on the Front Panel Lock. * &#x60;None&#x60; - Front Panel of the server is set to None state. It is required so that the next frontPanelLockState operation can be triggered. * &#x60;Lock&#x60; - Front Panel of the server is set to Locked state. * &#x60;Unlock&#x60; - Front Panel of the server is set to Unlocked state. | [optional] [default to "None"]
**HostInitConfiguration** | Pointer to **string** | The JSON formatted host initialization configuration containing the basic information for doing an initial boot. The information will be sent to CIMC and stored in host-init.json file on the server. The stored file can only be access using IPMI tool on the host OS. | [optional] 
**KvmReset** | Pointer to **string** | The allowed actions on the vKVM Reset. * &#x60;Ready&#x60; - Reset vKVM operation is allowed to be done on the server in this state. * &#x60;Reset&#x60; - The value that the UI/API needs to provide to trigger a Reset vKVM operation on a server. | [optional] [default to "Ready"]
**Name** | Pointer to **string** | The property used to identify the name of the server it is associated with. | [optional] [readonly] 
**OneTimeBootDevice** | Pointer to **string** | The name of the device chosen by user for configuring One-Time Boot device. | [optional] 
**PersistentMemoryOperation** | Pointer to [**NullableComputePersistentMemoryOperation**](ComputePersistentMemoryOperation.md) |  | [optional] 
**PersonalitySetting** | Pointer to [**NullableComputePersonalitySetting**](ComputePersonalitySetting.md) |  | [optional] 
**ServerConfig** | Pointer to [**NullableComputeServerConfig**](ComputeServerConfig.md) |  | [optional] 
**ServerOpStatus** | Pointer to [**[]ComputeServerOpStatus**](ComputeServerOpStatus.md) |  | [optional] 
**StorageControllerOperation** | Pointer to [**NullableComputeStorageControllerOperation**](ComputeStorageControllerOperation.md) |  | [optional] 
**StoragePhysicalDriveOperation** | Pointer to [**NullableComputeStoragePhysicalDriveOperation**](ComputeStoragePhysicalDriveOperation.md) |  | [optional] 
**StorageUtilityImageOperation** | Pointer to [**NullableComputeStorageUtilityImageOperation**](ComputeStorageUtilityImageOperation.md) |  | [optional] 
**StorageVirtualDriveOperation** | Pointer to [**NullableComputeStorageVirtualDriveOperation**](ComputeStorageVirtualDriveOperation.md) |  | [optional] 
**TpmReset** | Pointer to **string** | Clear the configuration of TPM chip in the server. * &#x60;None&#x60; - Perform no action on the TPM. * &#x60;ClearTpm&#x60; - Clear the configuration and restore factory defaults of TPM chip in the server. | [optional] [default to "None"]
**TunneledKvmState** | Pointer to **string** | By default, the tunneled vKVM property appears in Ready state. The property can be configured by performing allowed actions. Once the property is configured, it reverts to Ready state. * &#x60;Ready&#x60; - Tunneled vKVM is ready to be configured on the server. * &#x60;Enable&#x60; - Tunneled vKVM is enabled for the server. * &#x60;Disable&#x60; - Tunneled vKVM is disabled for the server. | [optional] [default to "Ready"]
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningWorkflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

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
### GetClearSel

`func (o *ComputeServerSetting) GetClearSel() string`

GetClearSel returns the ClearSel field if non-nil, zero value otherwise.

### GetClearSelOk

`func (o *ComputeServerSetting) GetClearSelOk() (*string, bool)`

GetClearSelOk returns a tuple with the ClearSel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSel

`func (o *ComputeServerSetting) SetClearSel(v string)`

SetClearSel sets ClearSel field to given value.

### HasClearSel

`func (o *ComputeServerSetting) HasClearSel() bool`

HasClearSel returns a boolean if a field has been set.

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

### GetCollectSel

`func (o *ComputeServerSetting) GetCollectSel() string`

GetCollectSel returns the CollectSel field if non-nil, zero value otherwise.

### GetCollectSelOk

`func (o *ComputeServerSetting) GetCollectSelOk() (*string, bool)`

GetCollectSelOk returns a tuple with the CollectSel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectSel

`func (o *ComputeServerSetting) SetCollectSel(v string)`

SetCollectSel sets CollectSel field to given value.

### HasCollectSel

`func (o *ComputeServerSetting) HasCollectSel() bool`

HasCollectSel returns a boolean if a field has been set.

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

### GetHostInitConfiguration

`func (o *ComputeServerSetting) GetHostInitConfiguration() string`

GetHostInitConfiguration returns the HostInitConfiguration field if non-nil, zero value otherwise.

### GetHostInitConfigurationOk

`func (o *ComputeServerSetting) GetHostInitConfigurationOk() (*string, bool)`

GetHostInitConfigurationOk returns a tuple with the HostInitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInitConfiguration

`func (o *ComputeServerSetting) SetHostInitConfiguration(v string)`

SetHostInitConfiguration sets HostInitConfiguration field to given value.

### HasHostInitConfiguration

`func (o *ComputeServerSetting) HasHostInitConfiguration() bool`

HasHostInitConfiguration returns a boolean if a field has been set.

### GetKvmReset

`func (o *ComputeServerSetting) GetKvmReset() string`

GetKvmReset returns the KvmReset field if non-nil, zero value otherwise.

### GetKvmResetOk

`func (o *ComputeServerSetting) GetKvmResetOk() (*string, bool)`

GetKvmResetOk returns a tuple with the KvmReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmReset

`func (o *ComputeServerSetting) SetKvmReset(v string)`

SetKvmReset sets KvmReset field to given value.

### HasKvmReset

`func (o *ComputeServerSetting) HasKvmReset() bool`

HasKvmReset returns a boolean if a field has been set.

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
### GetPersonalitySetting

`func (o *ComputeServerSetting) GetPersonalitySetting() ComputePersonalitySetting`

GetPersonalitySetting returns the PersonalitySetting field if non-nil, zero value otherwise.

### GetPersonalitySettingOk

`func (o *ComputeServerSetting) GetPersonalitySettingOk() (*ComputePersonalitySetting, bool)`

GetPersonalitySettingOk returns a tuple with the PersonalitySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalitySetting

`func (o *ComputeServerSetting) SetPersonalitySetting(v ComputePersonalitySetting)`

SetPersonalitySetting sets PersonalitySetting field to given value.

### HasPersonalitySetting

`func (o *ComputeServerSetting) HasPersonalitySetting() bool`

HasPersonalitySetting returns a boolean if a field has been set.

### SetPersonalitySettingNil

`func (o *ComputeServerSetting) SetPersonalitySettingNil(b bool)`

 SetPersonalitySettingNil sets the value for PersonalitySetting to be an explicit nil

### UnsetPersonalitySetting
`func (o *ComputeServerSetting) UnsetPersonalitySetting()`

UnsetPersonalitySetting ensures that no value is present for PersonalitySetting, not even an explicit nil
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
### GetServerOpStatus

`func (o *ComputeServerSetting) GetServerOpStatus() []ComputeServerOpStatus`

GetServerOpStatus returns the ServerOpStatus field if non-nil, zero value otherwise.

### GetServerOpStatusOk

`func (o *ComputeServerSetting) GetServerOpStatusOk() (*[]ComputeServerOpStatus, bool)`

GetServerOpStatusOk returns a tuple with the ServerOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOpStatus

`func (o *ComputeServerSetting) SetServerOpStatus(v []ComputeServerOpStatus)`

SetServerOpStatus sets ServerOpStatus field to given value.

### HasServerOpStatus

`func (o *ComputeServerSetting) HasServerOpStatus() bool`

HasServerOpStatus returns a boolean if a field has been set.

### SetServerOpStatusNil

`func (o *ComputeServerSetting) SetServerOpStatusNil(b bool)`

 SetServerOpStatusNil sets the value for ServerOpStatus to be an explicit nil

### UnsetServerOpStatus
`func (o *ComputeServerSetting) UnsetServerOpStatus()`

UnsetServerOpStatus ensures that no value is present for ServerOpStatus, not even an explicit nil
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
### GetStorageUtilityImageOperation

`func (o *ComputeServerSetting) GetStorageUtilityImageOperation() ComputeStorageUtilityImageOperation`

GetStorageUtilityImageOperation returns the StorageUtilityImageOperation field if non-nil, zero value otherwise.

### GetStorageUtilityImageOperationOk

`func (o *ComputeServerSetting) GetStorageUtilityImageOperationOk() (*ComputeStorageUtilityImageOperation, bool)`

GetStorageUtilityImageOperationOk returns a tuple with the StorageUtilityImageOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilityImageOperation

`func (o *ComputeServerSetting) SetStorageUtilityImageOperation(v ComputeStorageUtilityImageOperation)`

SetStorageUtilityImageOperation sets StorageUtilityImageOperation field to given value.

### HasStorageUtilityImageOperation

`func (o *ComputeServerSetting) HasStorageUtilityImageOperation() bool`

HasStorageUtilityImageOperation returns a boolean if a field has been set.

### SetStorageUtilityImageOperationNil

`func (o *ComputeServerSetting) SetStorageUtilityImageOperationNil(b bool)`

 SetStorageUtilityImageOperationNil sets the value for StorageUtilityImageOperation to be an explicit nil

### UnsetStorageUtilityImageOperation
`func (o *ComputeServerSetting) UnsetStorageUtilityImageOperation()`

UnsetStorageUtilityImageOperation ensures that no value is present for StorageUtilityImageOperation, not even an explicit nil
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
### GetTpmReset

`func (o *ComputeServerSetting) GetTpmReset() string`

GetTpmReset returns the TpmReset field if non-nil, zero value otherwise.

### GetTpmResetOk

`func (o *ComputeServerSetting) GetTpmResetOk() (*string, bool)`

GetTpmResetOk returns a tuple with the TpmReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmReset

`func (o *ComputeServerSetting) SetTpmReset(v string)`

SetTpmReset sets TpmReset field to given value.

### HasTpmReset

`func (o *ComputeServerSetting) HasTpmReset() bool`

HasTpmReset returns a boolean if a field has been set.

### GetTunneledKvmState

`func (o *ComputeServerSetting) GetTunneledKvmState() string`

GetTunneledKvmState returns the TunneledKvmState field if non-nil, zero value otherwise.

### GetTunneledKvmStateOk

`func (o *ComputeServerSetting) GetTunneledKvmStateOk() (*string, bool)`

GetTunneledKvmStateOk returns a tuple with the TunneledKvmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmState

`func (o *ComputeServerSetting) SetTunneledKvmState(v string)`

SetTunneledKvmState sets TunneledKvmState field to given value.

### HasTunneledKvmState

`func (o *ComputeServerSetting) HasTunneledKvmState() bool`

HasTunneledKvmState returns a boolean if a field has been set.

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


