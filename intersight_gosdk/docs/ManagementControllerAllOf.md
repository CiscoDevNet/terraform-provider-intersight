# ManagementControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "management.Controller"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "management.Controller"]
**Model** | Pointer to **string** | Model of the endpoint that houses the management controller. | [optional] [readonly] 
**UemStreamAdminState** | Pointer to **string** | Desired state of the UEM stream. * &#x60;Disabled&#x60; - The UEM event channel is disabled. * &#x60;Enabled&#x60; - The UEM event channel is enabled. | [optional] [default to "Disabled"]
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**EquipmentIoCardBase** | Pointer to [**EquipmentIoCardBaseRelationship**](EquipmentIoCardBaseRelationship.md) |  | [optional] 
**EquipmentSharedIoModule** | Pointer to [**EquipmentSharedIoModuleRelationship**](EquipmentSharedIoModuleRelationship.md) |  | [optional] 
**EquipmentSystemIoController** | Pointer to [**EquipmentSystemIoControllerRelationship**](EquipmentSystemIoControllerRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**ManagementInterfaces** | Pointer to [**[]ManagementInterfaceRelationship**](ManagementInterfaceRelationship.md) | An array of relationships to managementInterface resources. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**StorageSasExpander** | Pointer to [**StorageSasExpanderRelationship**](StorageSasExpanderRelationship.md) |  | [optional] 
**TopSystem** | Pointer to [**TopSystemRelationship**](TopSystemRelationship.md) |  | [optional] 

## Methods

### NewManagementControllerAllOf

`func NewManagementControllerAllOf(classId string, objectType string, ) *ManagementControllerAllOf`

NewManagementControllerAllOf instantiates a new ManagementControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementControllerAllOfWithDefaults

`func NewManagementControllerAllOfWithDefaults() *ManagementControllerAllOf`

NewManagementControllerAllOfWithDefaults instantiates a new ManagementControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ManagementControllerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ManagementControllerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ManagementControllerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ManagementControllerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ManagementControllerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ManagementControllerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModel

`func (o *ManagementControllerAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ManagementControllerAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ManagementControllerAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ManagementControllerAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUemStreamAdminState

`func (o *ManagementControllerAllOf) GetUemStreamAdminState() string`

GetUemStreamAdminState returns the UemStreamAdminState field if non-nil, zero value otherwise.

### GetUemStreamAdminStateOk

`func (o *ManagementControllerAllOf) GetUemStreamAdminStateOk() (*string, bool)`

GetUemStreamAdminStateOk returns a tuple with the UemStreamAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUemStreamAdminState

`func (o *ManagementControllerAllOf) SetUemStreamAdminState(v string)`

SetUemStreamAdminState sets UemStreamAdminState field to given value.

### HasUemStreamAdminState

`func (o *ManagementControllerAllOf) HasUemStreamAdminState() bool`

HasUemStreamAdminState returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *ManagementControllerAllOf) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *ManagementControllerAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *ManagementControllerAllOf) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *ManagementControllerAllOf) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ManagementControllerAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ManagementControllerAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ManagementControllerAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ManagementControllerAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ManagementControllerAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ManagementControllerAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ManagementControllerAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ManagementControllerAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentIoCardBase

`func (o *ManagementControllerAllOf) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship`

GetEquipmentIoCardBase returns the EquipmentIoCardBase field if non-nil, zero value otherwise.

### GetEquipmentIoCardBaseOk

`func (o *ManagementControllerAllOf) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool)`

GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoCardBase

`func (o *ManagementControllerAllOf) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship)`

SetEquipmentIoCardBase sets EquipmentIoCardBase field to given value.

### HasEquipmentIoCardBase

`func (o *ManagementControllerAllOf) HasEquipmentIoCardBase() bool`

HasEquipmentIoCardBase returns a boolean if a field has been set.

### GetEquipmentSharedIoModule

`func (o *ManagementControllerAllOf) GetEquipmentSharedIoModule() EquipmentSharedIoModuleRelationship`

GetEquipmentSharedIoModule returns the EquipmentSharedIoModule field if non-nil, zero value otherwise.

### GetEquipmentSharedIoModuleOk

`func (o *ManagementControllerAllOf) GetEquipmentSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool)`

GetEquipmentSharedIoModuleOk returns a tuple with the EquipmentSharedIoModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSharedIoModule

`func (o *ManagementControllerAllOf) SetEquipmentSharedIoModule(v EquipmentSharedIoModuleRelationship)`

SetEquipmentSharedIoModule sets EquipmentSharedIoModule field to given value.

### HasEquipmentSharedIoModule

`func (o *ManagementControllerAllOf) HasEquipmentSharedIoModule() bool`

HasEquipmentSharedIoModule returns a boolean if a field has been set.

### GetEquipmentSystemIoController

`func (o *ManagementControllerAllOf) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship`

GetEquipmentSystemIoController returns the EquipmentSystemIoController field if non-nil, zero value otherwise.

### GetEquipmentSystemIoControllerOk

`func (o *ManagementControllerAllOf) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool)`

GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSystemIoController

`func (o *ManagementControllerAllOf) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship)`

SetEquipmentSystemIoController sets EquipmentSystemIoController field to given value.

### HasEquipmentSystemIoController

`func (o *ManagementControllerAllOf) HasEquipmentSystemIoController() bool`

HasEquipmentSystemIoController returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ManagementControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ManagementControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ManagementControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ManagementControllerAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementInterfaces

`func (o *ManagementControllerAllOf) GetManagementInterfaces() []ManagementInterfaceRelationship`

GetManagementInterfaces returns the ManagementInterfaces field if non-nil, zero value otherwise.

### GetManagementInterfacesOk

`func (o *ManagementControllerAllOf) GetManagementInterfacesOk() (*[]ManagementInterfaceRelationship, bool)`

GetManagementInterfacesOk returns a tuple with the ManagementInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementInterfaces

`func (o *ManagementControllerAllOf) SetManagementInterfaces(v []ManagementInterfaceRelationship)`

SetManagementInterfaces sets ManagementInterfaces field to given value.

### HasManagementInterfaces

`func (o *ManagementControllerAllOf) HasManagementInterfaces() bool`

HasManagementInterfaces returns a boolean if a field has been set.

### SetManagementInterfacesNil

`func (o *ManagementControllerAllOf) SetManagementInterfacesNil(b bool)`

 SetManagementInterfacesNil sets the value for ManagementInterfaces to be an explicit nil

### UnsetManagementInterfaces
`func (o *ManagementControllerAllOf) UnsetManagementInterfaces()`

UnsetManagementInterfaces ensures that no value is present for ManagementInterfaces, not even an explicit nil
### GetNetworkElement

`func (o *ManagementControllerAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *ManagementControllerAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *ManagementControllerAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *ManagementControllerAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ManagementControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ManagementControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ManagementControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ManagementControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *ManagementControllerAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *ManagementControllerAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *ManagementControllerAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *ManagementControllerAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *ManagementControllerAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *ManagementControllerAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetStorageSasExpander

`func (o *ManagementControllerAllOf) GetStorageSasExpander() StorageSasExpanderRelationship`

GetStorageSasExpander returns the StorageSasExpander field if non-nil, zero value otherwise.

### GetStorageSasExpanderOk

`func (o *ManagementControllerAllOf) GetStorageSasExpanderOk() (*StorageSasExpanderRelationship, bool)`

GetStorageSasExpanderOk returns a tuple with the StorageSasExpander field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSasExpander

`func (o *ManagementControllerAllOf) SetStorageSasExpander(v StorageSasExpanderRelationship)`

SetStorageSasExpander sets StorageSasExpander field to given value.

### HasStorageSasExpander

`func (o *ManagementControllerAllOf) HasStorageSasExpander() bool`

HasStorageSasExpander returns a boolean if a field has been set.

### GetTopSystem

`func (o *ManagementControllerAllOf) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ManagementControllerAllOf) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ManagementControllerAllOf) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ManagementControllerAllOf) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


