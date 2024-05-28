# ManagementController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "management.Controller"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "management.Controller"]
**Certificate** | Pointer to [**CertificatemanagementImc**](CertificatemanagementImc.md) |  | [optional] 
**KmipClientCertificate** | Pointer to [**CertificatemanagementImc**](CertificatemanagementImc.md) |  | [optional] 
**Model** | Pointer to **string** | Model of the endpoint that houses the management controller. | [optional] [readonly] 
**RootCaCertificates** | Pointer to [**[]CertificatemanagementRootCaCertificate**](CertificatemanagementRootCaCertificate.md) |  | [optional] 
**UemStreamAdminState** | Pointer to **string** | Desired state of the UEM stream. * &#x60;Disabled&#x60; - The UEM event channel is disabled. * &#x60;Enabled&#x60; - The UEM event channel is enabled. | [optional] [default to "Disabled"]
**AdapterUnit** | Pointer to [**NullableAdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**EquipmentIoCardBase** | Pointer to [**NullableEquipmentIoCardBaseRelationship**](EquipmentIoCardBaseRelationship.md) |  | [optional] 
**EquipmentSharedIoModule** | Pointer to [**NullableEquipmentSharedIoModuleRelationship**](EquipmentSharedIoModuleRelationship.md) |  | [optional] 
**EquipmentSystemIoController** | Pointer to [**NullableEquipmentSystemIoControllerRelationship**](EquipmentSystemIoControllerRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**ManagementInterfaces** | Pointer to [**[]ManagementInterfaceRelationship**](ManagementInterfaceRelationship.md) | An array of relationships to managementInterface resources. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**StorageSasExpander** | Pointer to [**NullableStorageSasExpanderRelationship**](StorageSasExpanderRelationship.md) |  | [optional] 
**TopSystem** | Pointer to [**NullableTopSystemRelationship**](TopSystemRelationship.md) |  | [optional] 

## Methods

### NewManagementController

`func NewManagementController(classId string, objectType string, ) *ManagementController`

NewManagementController instantiates a new ManagementController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementControllerWithDefaults

`func NewManagementControllerWithDefaults() *ManagementController`

NewManagementControllerWithDefaults instantiates a new ManagementController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ManagementController) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ManagementController) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ManagementController) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ManagementController) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ManagementController) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ManagementController) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificate

`func (o *ManagementController) GetCertificate() CertificatemanagementImc`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ManagementController) GetCertificateOk() (*CertificatemanagementImc, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ManagementController) SetCertificate(v CertificatemanagementImc)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ManagementController) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetKmipClientCertificate

`func (o *ManagementController) GetKmipClientCertificate() CertificatemanagementImc`

GetKmipClientCertificate returns the KmipClientCertificate field if non-nil, zero value otherwise.

### GetKmipClientCertificateOk

`func (o *ManagementController) GetKmipClientCertificateOk() (*CertificatemanagementImc, bool)`

GetKmipClientCertificateOk returns a tuple with the KmipClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipClientCertificate

`func (o *ManagementController) SetKmipClientCertificate(v CertificatemanagementImc)`

SetKmipClientCertificate sets KmipClientCertificate field to given value.

### HasKmipClientCertificate

`func (o *ManagementController) HasKmipClientCertificate() bool`

HasKmipClientCertificate returns a boolean if a field has been set.

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

### GetRootCaCertificates

`func (o *ManagementController) GetRootCaCertificates() []CertificatemanagementRootCaCertificate`

GetRootCaCertificates returns the RootCaCertificates field if non-nil, zero value otherwise.

### GetRootCaCertificatesOk

`func (o *ManagementController) GetRootCaCertificatesOk() (*[]CertificatemanagementRootCaCertificate, bool)`

GetRootCaCertificatesOk returns a tuple with the RootCaCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCaCertificates

`func (o *ManagementController) SetRootCaCertificates(v []CertificatemanagementRootCaCertificate)`

SetRootCaCertificates sets RootCaCertificates field to given value.

### HasRootCaCertificates

`func (o *ManagementController) HasRootCaCertificates() bool`

HasRootCaCertificates returns a boolean if a field has been set.

### SetRootCaCertificatesNil

`func (o *ManagementController) SetRootCaCertificatesNil(b bool)`

 SetRootCaCertificatesNil sets the value for RootCaCertificates to be an explicit nil

### UnsetRootCaCertificates
`func (o *ManagementController) UnsetRootCaCertificates()`

UnsetRootCaCertificates ensures that no value is present for RootCaCertificates, not even an explicit nil
### GetUemStreamAdminState

`func (o *ManagementController) GetUemStreamAdminState() string`

GetUemStreamAdminState returns the UemStreamAdminState field if non-nil, zero value otherwise.

### GetUemStreamAdminStateOk

`func (o *ManagementController) GetUemStreamAdminStateOk() (*string, bool)`

GetUemStreamAdminStateOk returns a tuple with the UemStreamAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUemStreamAdminState

`func (o *ManagementController) SetUemStreamAdminState(v string)`

SetUemStreamAdminState sets UemStreamAdminState field to given value.

### HasUemStreamAdminState

`func (o *ManagementController) HasUemStreamAdminState() bool`

HasUemStreamAdminState returns a boolean if a field has been set.

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

### SetAdapterUnitNil

`func (o *ManagementController) SetAdapterUnitNil(b bool)`

 SetAdapterUnitNil sets the value for AdapterUnit to be an explicit nil

### UnsetAdapterUnit
`func (o *ManagementController) UnsetAdapterUnit()`

UnsetAdapterUnit ensures that no value is present for AdapterUnit, not even an explicit nil
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

### SetComputeBladeNil

`func (o *ManagementController) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *ManagementController) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
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

### SetComputeRackUnitNil

`func (o *ManagementController) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *ManagementController) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
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

### SetEquipmentIoCardBaseNil

`func (o *ManagementController) SetEquipmentIoCardBaseNil(b bool)`

 SetEquipmentIoCardBaseNil sets the value for EquipmentIoCardBase to be an explicit nil

### UnsetEquipmentIoCardBase
`func (o *ManagementController) UnsetEquipmentIoCardBase()`

UnsetEquipmentIoCardBase ensures that no value is present for EquipmentIoCardBase, not even an explicit nil
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

### SetEquipmentSharedIoModuleNil

`func (o *ManagementController) SetEquipmentSharedIoModuleNil(b bool)`

 SetEquipmentSharedIoModuleNil sets the value for EquipmentSharedIoModule to be an explicit nil

### UnsetEquipmentSharedIoModule
`func (o *ManagementController) UnsetEquipmentSharedIoModule()`

UnsetEquipmentSharedIoModule ensures that no value is present for EquipmentSharedIoModule, not even an explicit nil
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

### SetEquipmentSystemIoControllerNil

`func (o *ManagementController) SetEquipmentSystemIoControllerNil(b bool)`

 SetEquipmentSystemIoControllerNil sets the value for EquipmentSystemIoController to be an explicit nil

### UnsetEquipmentSystemIoController
`func (o *ManagementController) UnsetEquipmentSystemIoController()`

UnsetEquipmentSystemIoController ensures that no value is present for EquipmentSystemIoController, not even an explicit nil
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

### SetInventoryDeviceInfoNil

`func (o *ManagementController) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *ManagementController) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
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

### SetNetworkElementNil

`func (o *ManagementController) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *ManagementController) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
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

### SetRegisteredDeviceNil

`func (o *ManagementController) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ManagementController) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
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

### SetStorageSasExpanderNil

`func (o *ManagementController) SetStorageSasExpanderNil(b bool)`

 SetStorageSasExpanderNil sets the value for StorageSasExpander to be an explicit nil

### UnsetStorageSasExpander
`func (o *ManagementController) UnsetStorageSasExpander()`

UnsetStorageSasExpander ensures that no value is present for StorageSasExpander, not even an explicit nil
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

### SetTopSystemNil

`func (o *ManagementController) SetTopSystemNil(b bool)`

 SetTopSystemNil sets the value for TopSystem to be an explicit nil

### UnsetTopSystem
`func (o *ManagementController) UnsetTopSystem()`

UnsetTopSystem ensures that no value is present for TopSystem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


