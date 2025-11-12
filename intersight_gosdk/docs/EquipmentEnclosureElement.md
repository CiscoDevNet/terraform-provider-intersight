# EquipmentEnclosureElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.EnclosureElement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.EnclosureElement"]
**Description** | Pointer to **string** | This field is to provide description for the Enclosure module. | [optional] [readonly] 
**EnclosureId** | Pointer to **string** | This field acts as the identifier for this particular Enclosure Module, within the Server or Fabric Interconnect. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the pluggable Enclosure Module. | [optional] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | This field is used to indicate this equipment enclosure&#39;s operational state. | [optional] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Enclosure Module. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the Enclosure module. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Enclosure Module. | [optional] [readonly] 
**Status** | Pointer to **string** | This field is to abstract the status of the Enclosure module. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the Enclosure module. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Enclosure Module. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**FirmwareRunningFirmwares** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentEnclosureElement

`func NewEquipmentEnclosureElement(classId string, objectType string, ) *EquipmentEnclosureElement`

NewEquipmentEnclosureElement instantiates a new EquipmentEnclosureElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentEnclosureElementWithDefaults

`func NewEquipmentEnclosureElementWithDefaults() *EquipmentEnclosureElement`

NewEquipmentEnclosureElementWithDefaults instantiates a new EquipmentEnclosureElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentEnclosureElement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentEnclosureElement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentEnclosureElement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentEnclosureElement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentEnclosureElement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentEnclosureElement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *EquipmentEnclosureElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentEnclosureElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentEnclosureElement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentEnclosureElement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnclosureId

`func (o *EquipmentEnclosureElement) GetEnclosureId() string`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *EquipmentEnclosureElement) GetEnclosureIdOk() (*string, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *EquipmentEnclosureElement) SetEnclosureId(v string)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *EquipmentEnclosureElement) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetName

`func (o *EquipmentEnclosureElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentEnclosureElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentEnclosureElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentEnclosureElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentEnclosureElement) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentEnclosureElement) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentEnclosureElement) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentEnclosureElement) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentEnclosureElement) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentEnclosureElement) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentEnclosureElement) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentEnclosureElement) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentEnclosureElement) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentEnclosureElement) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentEnclosureElement) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentEnclosureElement) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentEnclosureElement) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentEnclosureElement) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentEnclosureElement) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentEnclosureElement) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentEnclosureElement) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentEnclosureElement) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentEnclosureElement) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentEnclosureElement) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentEnclosureElement) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentEnclosureElement) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetStatus

`func (o *EquipmentEnclosureElement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EquipmentEnclosureElement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EquipmentEnclosureElement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EquipmentEnclosureElement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EquipmentEnclosureElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EquipmentEnclosureElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EquipmentEnclosureElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EquipmentEnclosureElement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentEnclosureElement) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentEnclosureElement) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentEnclosureElement) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentEnclosureElement) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentEnclosureElement) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentEnclosureElement) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentEnclosureElement) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentEnclosureElement) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### SetComputeRackUnitNil

`func (o *EquipmentEnclosureElement) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *EquipmentEnclosureElement) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
### GetFirmwareRunningFirmwares

`func (o *EquipmentEnclosureElement) GetFirmwareRunningFirmwares() []FirmwareRunningFirmwareRelationship`

GetFirmwareRunningFirmwares returns the FirmwareRunningFirmwares field if non-nil, zero value otherwise.

### GetFirmwareRunningFirmwaresOk

`func (o *EquipmentEnclosureElement) GetFirmwareRunningFirmwaresOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetFirmwareRunningFirmwaresOk returns a tuple with the FirmwareRunningFirmwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareRunningFirmwares

`func (o *EquipmentEnclosureElement) SetFirmwareRunningFirmwares(v []FirmwareRunningFirmwareRelationship)`

SetFirmwareRunningFirmwares sets FirmwareRunningFirmwares field to given value.

### HasFirmwareRunningFirmwares

`func (o *EquipmentEnclosureElement) HasFirmwareRunningFirmwares() bool`

HasFirmwareRunningFirmwares returns a boolean if a field has been set.

### SetFirmwareRunningFirmwaresNil

`func (o *EquipmentEnclosureElement) SetFirmwareRunningFirmwaresNil(b bool)`

 SetFirmwareRunningFirmwaresNil sets the value for FirmwareRunningFirmwares to be an explicit nil

### UnsetFirmwareRunningFirmwares
`func (o *EquipmentEnclosureElement) UnsetFirmwareRunningFirmwares()`

UnsetFirmwareRunningFirmwares ensures that no value is present for FirmwareRunningFirmwares, not even an explicit nil
### GetGraphicsCards

`func (o *EquipmentEnclosureElement) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *EquipmentEnclosureElement) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *EquipmentEnclosureElement) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *EquipmentEnclosureElement) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *EquipmentEnclosureElement) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *EquipmentEnclosureElement) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentEnclosureElement) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentEnclosureElement) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentEnclosureElement) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentEnclosureElement) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentEnclosureElement) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentEnclosureElement) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


