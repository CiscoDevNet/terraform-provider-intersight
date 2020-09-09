# EquipmentSystemIoController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | Pointer to **string** | The assigned identifier for a chassis. | [optional] [readonly] 
**ConnectionPath** | Pointer to **string** | Connection Path identifies the data path available between IOModule and FI. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | Connection status identifies the status of data path. | [optional] [readonly] 
**Description** | Pointer to **string** | This field gives a brief information on systemIOController. | [optional] [readonly] 
**ManagingInstance** | Pointer to **string** | This field identifies the CIMC that manages the controller. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the SIOC operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part Number identifier for the IO module. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for systemIOController. | [optional] [readonly] 
**SystemIoControllerId** | Pointer to **int64** | This represents system I/O Controller identifier. | [optional] [readonly] 
**Cmc** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**SharedIoModule** | Pointer to [**EquipmentSharedIoModuleRelationship**](equipment.SharedIoModule.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentSystemIoController

`func NewEquipmentSystemIoController() *EquipmentSystemIoController`

NewEquipmentSystemIoController instantiates a new EquipmentSystemIoController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSystemIoControllerWithDefaults

`func NewEquipmentSystemIoControllerWithDefaults() *EquipmentSystemIoController`

NewEquipmentSystemIoControllerWithDefaults instantiates a new EquipmentSystemIoController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisId

`func (o *EquipmentSystemIoController) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentSystemIoController) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentSystemIoController) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentSystemIoController) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetConnectionPath

`func (o *EquipmentSystemIoController) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentSystemIoController) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentSystemIoController) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentSystemIoController) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *EquipmentSystemIoController) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *EquipmentSystemIoController) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *EquipmentSystemIoController) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *EquipmentSystemIoController) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentSystemIoController) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentSystemIoController) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentSystemIoController) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentSystemIoController) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManagingInstance

`func (o *EquipmentSystemIoController) GetManagingInstance() string`

GetManagingInstance returns the ManagingInstance field if non-nil, zero value otherwise.

### GetManagingInstanceOk

`func (o *EquipmentSystemIoController) GetManagingInstanceOk() (*string, bool)`

GetManagingInstanceOk returns a tuple with the ManagingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingInstance

`func (o *EquipmentSystemIoController) SetManagingInstance(v string)`

SetManagingInstance sets ManagingInstance field to given value.

### HasManagingInstance

`func (o *EquipmentSystemIoController) HasManagingInstance() bool`

HasManagingInstance returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentSystemIoController) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSystemIoController) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSystemIoController) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSystemIoController) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSystemIoController) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSystemIoController) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSystemIoController) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSystemIoController) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentSystemIoController) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentSystemIoController) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentSystemIoController) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentSystemIoController) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSystemIoControllerId

`func (o *EquipmentSystemIoController) GetSystemIoControllerId() int64`

GetSystemIoControllerId returns the SystemIoControllerId field if non-nil, zero value otherwise.

### GetSystemIoControllerIdOk

`func (o *EquipmentSystemIoController) GetSystemIoControllerIdOk() (*int64, bool)`

GetSystemIoControllerIdOk returns a tuple with the SystemIoControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIoControllerId

`func (o *EquipmentSystemIoController) SetSystemIoControllerId(v int64)`

SetSystemIoControllerId sets SystemIoControllerId field to given value.

### HasSystemIoControllerId

`func (o *EquipmentSystemIoController) HasSystemIoControllerId() bool`

HasSystemIoControllerId returns a boolean if a field has been set.

### GetCmc

`func (o *EquipmentSystemIoController) GetCmc() ManagementControllerRelationship`

GetCmc returns the Cmc field if non-nil, zero value otherwise.

### GetCmcOk

`func (o *EquipmentSystemIoController) GetCmcOk() (*ManagementControllerRelationship, bool)`

GetCmcOk returns a tuple with the Cmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmc

`func (o *EquipmentSystemIoController) SetCmc(v ManagementControllerRelationship)`

SetCmc sets Cmc field to given value.

### HasCmc

`func (o *EquipmentSystemIoController) HasCmc() bool`

HasCmc returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentSystemIoController) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentSystemIoController) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentSystemIoController) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentSystemIoController) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentSystemIoController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSystemIoController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSystemIoController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSystemIoController) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentSystemIoController) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSystemIoController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSystemIoController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSystemIoController) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSharedIoModule

`func (o *EquipmentSystemIoController) GetSharedIoModule() EquipmentSharedIoModuleRelationship`

GetSharedIoModule returns the SharedIoModule field if non-nil, zero value otherwise.

### GetSharedIoModuleOk

`func (o *EquipmentSystemIoController) GetSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool)`

GetSharedIoModuleOk returns a tuple with the SharedIoModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedIoModule

`func (o *EquipmentSystemIoController) SetSharedIoModule(v EquipmentSharedIoModuleRelationship)`

SetSharedIoModule sets SharedIoModule field to given value.

### HasSharedIoModule

`func (o *EquipmentSystemIoController) HasSharedIoModule() bool`

HasSharedIoModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


