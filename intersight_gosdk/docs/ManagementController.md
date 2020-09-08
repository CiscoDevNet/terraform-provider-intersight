# ManagementController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model of the endpoint that houses the management controller. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EquipmentIoCardBase** | Pointer to [**EquipmentIoCardBaseRelationship**](equipment.IoCardBase.Relationship.md) |  | [optional] 
**EquipmentSharedIoModule** | Pointer to [**EquipmentSharedIoModuleRelationship**](equipment.SharedIoModule.Relationship.md) |  | [optional] 
**EquipmentSystemIoController** | Pointer to [**EquipmentSystemIoControllerRelationship**](equipment.SystemIoController.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**ManagementInterfaces** | Pointer to [**[]ManagementInterfaceRelationship**](management.Interface.Relationship.md) | An array of relationships to managementInterface resources. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**StorageSasExpander** | Pointer to [**StorageSasExpanderRelationship**](storage.SasExpander.Relationship.md) |  | [optional] 
**TopSystem** | Pointer to [**TopSystemRelationship**](top.System.Relationship.md) |  | [optional] 

## Methods

### NewManagementController

`func NewManagementController() *ManagementController`

NewManagementController instantiates a new ManagementController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementControllerWithDefaults

`func NewManagementControllerWithDefaults() *ManagementController`

NewManagementControllerWithDefaults instantiates a new ManagementController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ManagementController) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ManagementController) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ManagementController) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ManagementController) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *ManagementController) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *ManagementController) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *ManagementController) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *ManagementController) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ManagementController) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ManagementController) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ManagementController) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ManagementController) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ManagementController) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ManagementController) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ManagementController) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ManagementController) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentIoCardBase

`func (o *ManagementController) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship`

GetEquipmentIoCardBase returns the EquipmentIoCardBase field if non-nil, zero value otherwise.

### GetEquipmentIoCardBaseOk

`func (o *ManagementController) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool)`

GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoCardBase

`func (o *ManagementController) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship)`

SetEquipmentIoCardBase sets EquipmentIoCardBase field to given value.

### HasEquipmentIoCardBase

`func (o *ManagementController) HasEquipmentIoCardBase() bool`

HasEquipmentIoCardBase returns a boolean if a field has been set.

### GetEquipmentSharedIoModule

`func (o *ManagementController) GetEquipmentSharedIoModule() EquipmentSharedIoModuleRelationship`

GetEquipmentSharedIoModule returns the EquipmentSharedIoModule field if non-nil, zero value otherwise.

### GetEquipmentSharedIoModuleOk

`func (o *ManagementController) GetEquipmentSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool)`

GetEquipmentSharedIoModuleOk returns a tuple with the EquipmentSharedIoModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSharedIoModule

`func (o *ManagementController) SetEquipmentSharedIoModule(v EquipmentSharedIoModuleRelationship)`

SetEquipmentSharedIoModule sets EquipmentSharedIoModule field to given value.

### HasEquipmentSharedIoModule

`func (o *ManagementController) HasEquipmentSharedIoModule() bool`

HasEquipmentSharedIoModule returns a boolean if a field has been set.

### GetEquipmentSystemIoController

`func (o *ManagementController) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship`

GetEquipmentSystemIoController returns the EquipmentSystemIoController field if non-nil, zero value otherwise.

### GetEquipmentSystemIoControllerOk

`func (o *ManagementController) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool)`

GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSystemIoController

`func (o *ManagementController) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship)`

SetEquipmentSystemIoController sets EquipmentSystemIoController field to given value.

### HasEquipmentSystemIoController

`func (o *ManagementController) HasEquipmentSystemIoController() bool`

HasEquipmentSystemIoController returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ManagementController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ManagementController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ManagementController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ManagementController) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementInterfaces

`func (o *ManagementController) GetManagementInterfaces() []ManagementInterfaceRelationship`

GetManagementInterfaces returns the ManagementInterfaces field if non-nil, zero value otherwise.

### GetManagementInterfacesOk

`func (o *ManagementController) GetManagementInterfacesOk() (*[]ManagementInterfaceRelationship, bool)`

GetManagementInterfacesOk returns a tuple with the ManagementInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementInterfaces

`func (o *ManagementController) SetManagementInterfaces(v []ManagementInterfaceRelationship)`

SetManagementInterfaces sets ManagementInterfaces field to given value.

### HasManagementInterfaces

`func (o *ManagementController) HasManagementInterfaces() bool`

HasManagementInterfaces returns a boolean if a field has been set.

### SetManagementInterfacesNil

`func (o *ManagementController) SetManagementInterfacesNil(b bool)`

 SetManagementInterfacesNil sets the value for ManagementInterfaces to be an explicit nil

### UnsetManagementInterfaces
`func (o *ManagementController) UnsetManagementInterfaces()`

UnsetManagementInterfaces ensures that no value is present for ManagementInterfaces, not even an explicit nil
### GetNetworkElement

`func (o *ManagementController) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *ManagementController) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *ManagementController) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *ManagementController) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ManagementController) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ManagementController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ManagementController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ManagementController) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *ManagementController) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *ManagementController) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *ManagementController) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *ManagementController) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *ManagementController) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *ManagementController) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetStorageSasExpander

`func (o *ManagementController) GetStorageSasExpander() StorageSasExpanderRelationship`

GetStorageSasExpander returns the StorageSasExpander field if non-nil, zero value otherwise.

### GetStorageSasExpanderOk

`func (o *ManagementController) GetStorageSasExpanderOk() (*StorageSasExpanderRelationship, bool)`

GetStorageSasExpanderOk returns a tuple with the StorageSasExpander field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSasExpander

`func (o *ManagementController) SetStorageSasExpander(v StorageSasExpanderRelationship)`

SetStorageSasExpander sets StorageSasExpander field to given value.

### HasStorageSasExpander

`func (o *ManagementController) HasStorageSasExpander() bool`

HasStorageSasExpander returns a boolean if a field has been set.

### GetTopSystem

`func (o *ManagementController) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ManagementController) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ManagementController) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ManagementController) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


