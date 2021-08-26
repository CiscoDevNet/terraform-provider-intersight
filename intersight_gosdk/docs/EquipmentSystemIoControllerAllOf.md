# EquipmentSystemIoControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SystemIoController"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SystemIoController"]
**ChassisId** | Pointer to **string** | The assigned identifier for a chassis. | [optional] [readonly] 
**ConnectionPath** | Pointer to **string** | Connection Path identifies the data path available between IOModule and FI. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | Connection status identifies the status of data path. | [optional] [readonly] 
**Description** | Pointer to **string** | This field gives a brief information on systemIOController. | [optional] [readonly] 
**ManagingInstance** | Pointer to **string** | This field identifies the CIMC that manages the controller. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the SIOC operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part Number identifier for the IO module. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for systemIOController. | [optional] [readonly] 
**SystemIoControllerId** | Pointer to **int64** | This represents system I/O Controller identifier. | [optional] [readonly] 
**Cmc** | Pointer to [**ManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SharedIoModule** | Pointer to [**EquipmentSharedIoModuleRelationship**](EquipmentSharedIoModuleRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSystemIoControllerAllOf

`func NewEquipmentSystemIoControllerAllOf(classId string, objectType string, ) *EquipmentSystemIoControllerAllOf`

NewEquipmentSystemIoControllerAllOf instantiates a new EquipmentSystemIoControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSystemIoControllerAllOfWithDefaults

`func NewEquipmentSystemIoControllerAllOfWithDefaults() *EquipmentSystemIoControllerAllOf`

NewEquipmentSystemIoControllerAllOfWithDefaults instantiates a new EquipmentSystemIoControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSystemIoControllerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSystemIoControllerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSystemIoControllerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSystemIoControllerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSystemIoControllerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSystemIoControllerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisId

`func (o *EquipmentSystemIoControllerAllOf) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentSystemIoControllerAllOf) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentSystemIoControllerAllOf) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentSystemIoControllerAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetConnectionPath

`func (o *EquipmentSystemIoControllerAllOf) GetConnectionPath() string`

GetConnectionPath returns the ConnectionPath field if non-nil, zero value otherwise.

### GetConnectionPathOk

`func (o *EquipmentSystemIoControllerAllOf) GetConnectionPathOk() (*string, bool)`

GetConnectionPathOk returns a tuple with the ConnectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPath

`func (o *EquipmentSystemIoControllerAllOf) SetConnectionPath(v string)`

SetConnectionPath sets ConnectionPath field to given value.

### HasConnectionPath

`func (o *EquipmentSystemIoControllerAllOf) HasConnectionPath() bool`

HasConnectionPath returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *EquipmentSystemIoControllerAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *EquipmentSystemIoControllerAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *EquipmentSystemIoControllerAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *EquipmentSystemIoControllerAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentSystemIoControllerAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentSystemIoControllerAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentSystemIoControllerAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentSystemIoControllerAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManagingInstance

`func (o *EquipmentSystemIoControllerAllOf) GetManagingInstance() string`

GetManagingInstance returns the ManagingInstance field if non-nil, zero value otherwise.

### GetManagingInstanceOk

`func (o *EquipmentSystemIoControllerAllOf) GetManagingInstanceOk() (*string, bool)`

GetManagingInstanceOk returns a tuple with the ManagingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingInstance

`func (o *EquipmentSystemIoControllerAllOf) SetManagingInstance(v string)`

SetManagingInstance sets ManagingInstance field to given value.

### HasManagingInstance

`func (o *EquipmentSystemIoControllerAllOf) HasManagingInstance() bool`

HasManagingInstance returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentSystemIoControllerAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSystemIoControllerAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSystemIoControllerAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSystemIoControllerAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSystemIoControllerAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSystemIoControllerAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSystemIoControllerAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSystemIoControllerAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentSystemIoControllerAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentSystemIoControllerAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentSystemIoControllerAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentSystemIoControllerAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSystemIoControllerId

`func (o *EquipmentSystemIoControllerAllOf) GetSystemIoControllerId() int64`

GetSystemIoControllerId returns the SystemIoControllerId field if non-nil, zero value otherwise.

### GetSystemIoControllerIdOk

`func (o *EquipmentSystemIoControllerAllOf) GetSystemIoControllerIdOk() (*int64, bool)`

GetSystemIoControllerIdOk returns a tuple with the SystemIoControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIoControllerId

`func (o *EquipmentSystemIoControllerAllOf) SetSystemIoControllerId(v int64)`

SetSystemIoControllerId sets SystemIoControllerId field to given value.

### HasSystemIoControllerId

`func (o *EquipmentSystemIoControllerAllOf) HasSystemIoControllerId() bool`

HasSystemIoControllerId returns a boolean if a field has been set.

### GetCmc

`func (o *EquipmentSystemIoControllerAllOf) GetCmc() ManagementControllerRelationship`

GetCmc returns the Cmc field if non-nil, zero value otherwise.

### GetCmcOk

`func (o *EquipmentSystemIoControllerAllOf) GetCmcOk() (*ManagementControllerRelationship, bool)`

GetCmcOk returns a tuple with the Cmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmc

`func (o *EquipmentSystemIoControllerAllOf) SetCmc(v ManagementControllerRelationship)`

SetCmc sets Cmc field to given value.

### HasCmc

`func (o *EquipmentSystemIoControllerAllOf) HasCmc() bool`

HasCmc returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentSystemIoControllerAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentSystemIoControllerAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentSystemIoControllerAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentSystemIoControllerAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentSystemIoControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSystemIoControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSystemIoControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSystemIoControllerAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentSystemIoControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSystemIoControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSystemIoControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSystemIoControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSharedIoModule

`func (o *EquipmentSystemIoControllerAllOf) GetSharedIoModule() EquipmentSharedIoModuleRelationship`

GetSharedIoModule returns the SharedIoModule field if non-nil, zero value otherwise.

### GetSharedIoModuleOk

`func (o *EquipmentSystemIoControllerAllOf) GetSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool)`

GetSharedIoModuleOk returns a tuple with the SharedIoModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedIoModule

`func (o *EquipmentSystemIoControllerAllOf) SetSharedIoModule(v EquipmentSharedIoModuleRelationship)`

SetSharedIoModule sets SharedIoModule field to given value.

### HasSharedIoModule

`func (o *EquipmentSystemIoControllerAllOf) HasSharedIoModule() bool`

HasSharedIoModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


