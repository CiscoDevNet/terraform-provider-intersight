# EquipmentExpanderModuleOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ExpanderModuleOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ExpanderModuleOperation"]
**AdminOperation** | Pointer to **string** | User configured operation on the X-Fabric Module. * &#x60;None&#x60; - Placeholder default value for iom power state property. * &#x60;Reboot&#x60; - IO Module reboot state property value. | [optional] [default to "None"]
**AffectedObjName** | Pointer to **string** | Placeholder for affected object name which is a combination of chassis and X-Fabric Module ID. Used to display affected object in audit log. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | The configured state of these settings in the target X-Fabric Module. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target X-Fabric Module. Applying - This state denotes that the settings are being applied in the target X-Fabric Module. Failed - This state denotes that the settings could not be applied in the target X-Fabric Module. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. * &#x60;Scheduled&#x60; - User configured settings are scheduled to be applied. | [optional] [readonly] [default to "None"]
**ExpanderModuleOperationStatus** | Pointer to [**[]EquipmentIoCardOperationStatus**](EquipmentIoCardOperationStatus.md) |  | [optional] 
**Serial** | Pointer to **string** | Serial ID of the X-Fabric Module. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**ExpanderModule** | Pointer to [**NullableEquipmentExpanderModuleRelationship**](EquipmentExpanderModuleRelationship.md) |  | [optional] 

## Methods

### NewEquipmentExpanderModuleOperation

`func NewEquipmentExpanderModuleOperation(classId string, objectType string, ) *EquipmentExpanderModuleOperation`

NewEquipmentExpanderModuleOperation instantiates a new EquipmentExpanderModuleOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentExpanderModuleOperationWithDefaults

`func NewEquipmentExpanderModuleOperationWithDefaults() *EquipmentExpanderModuleOperation`

NewEquipmentExpanderModuleOperationWithDefaults instantiates a new EquipmentExpanderModuleOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentExpanderModuleOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentExpanderModuleOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentExpanderModuleOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentExpanderModuleOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentExpanderModuleOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentExpanderModuleOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminOperation

`func (o *EquipmentExpanderModuleOperation) GetAdminOperation() string`

GetAdminOperation returns the AdminOperation field if non-nil, zero value otherwise.

### GetAdminOperationOk

`func (o *EquipmentExpanderModuleOperation) GetAdminOperationOk() (*string, bool)`

GetAdminOperationOk returns a tuple with the AdminOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOperation

`func (o *EquipmentExpanderModuleOperation) SetAdminOperation(v string)`

SetAdminOperation sets AdminOperation field to given value.

### HasAdminOperation

`func (o *EquipmentExpanderModuleOperation) HasAdminOperation() bool`

HasAdminOperation returns a boolean if a field has been set.

### GetAffectedObjName

`func (o *EquipmentExpanderModuleOperation) GetAffectedObjName() string`

GetAffectedObjName returns the AffectedObjName field if non-nil, zero value otherwise.

### GetAffectedObjNameOk

`func (o *EquipmentExpanderModuleOperation) GetAffectedObjNameOk() (*string, bool)`

GetAffectedObjNameOk returns a tuple with the AffectedObjName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObjName

`func (o *EquipmentExpanderModuleOperation) SetAffectedObjName(v string)`

SetAffectedObjName sets AffectedObjName field to given value.

### HasAffectedObjName

`func (o *EquipmentExpanderModuleOperation) HasAffectedObjName() bool`

HasAffectedObjName returns a boolean if a field has been set.

### GetConfigState

`func (o *EquipmentExpanderModuleOperation) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentExpanderModuleOperation) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentExpanderModuleOperation) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentExpanderModuleOperation) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetExpanderModuleOperationStatus

`func (o *EquipmentExpanderModuleOperation) GetExpanderModuleOperationStatus() []EquipmentIoCardOperationStatus`

GetExpanderModuleOperationStatus returns the ExpanderModuleOperationStatus field if non-nil, zero value otherwise.

### GetExpanderModuleOperationStatusOk

`func (o *EquipmentExpanderModuleOperation) GetExpanderModuleOperationStatusOk() (*[]EquipmentIoCardOperationStatus, bool)`

GetExpanderModuleOperationStatusOk returns a tuple with the ExpanderModuleOperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderModuleOperationStatus

`func (o *EquipmentExpanderModuleOperation) SetExpanderModuleOperationStatus(v []EquipmentIoCardOperationStatus)`

SetExpanderModuleOperationStatus sets ExpanderModuleOperationStatus field to given value.

### HasExpanderModuleOperationStatus

`func (o *EquipmentExpanderModuleOperation) HasExpanderModuleOperationStatus() bool`

HasExpanderModuleOperationStatus returns a boolean if a field has been set.

### SetExpanderModuleOperationStatusNil

`func (o *EquipmentExpanderModuleOperation) SetExpanderModuleOperationStatusNil(b bool)`

 SetExpanderModuleOperationStatusNil sets the value for ExpanderModuleOperationStatus to be an explicit nil

### UnsetExpanderModuleOperationStatus
`func (o *EquipmentExpanderModuleOperation) UnsetExpanderModuleOperationStatus()`

UnsetExpanderModuleOperationStatus ensures that no value is present for ExpanderModuleOperationStatus, not even an explicit nil
### GetSerial

`func (o *EquipmentExpanderModuleOperation) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentExpanderModuleOperation) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentExpanderModuleOperation) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentExpanderModuleOperation) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentExpanderModuleOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentExpanderModuleOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentExpanderModuleOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentExpanderModuleOperation) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### SetDeviceRegistrationNil

`func (o *EquipmentExpanderModuleOperation) SetDeviceRegistrationNil(b bool)`

 SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil

### UnsetDeviceRegistration
`func (o *EquipmentExpanderModuleOperation) UnsetDeviceRegistration()`

UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil
### GetExpanderModule

`func (o *EquipmentExpanderModuleOperation) GetExpanderModule() EquipmentExpanderModuleRelationship`

GetExpanderModule returns the ExpanderModule field if non-nil, zero value otherwise.

### GetExpanderModuleOk

`func (o *EquipmentExpanderModuleOperation) GetExpanderModuleOk() (*EquipmentExpanderModuleRelationship, bool)`

GetExpanderModuleOk returns a tuple with the ExpanderModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderModule

`func (o *EquipmentExpanderModuleOperation) SetExpanderModule(v EquipmentExpanderModuleRelationship)`

SetExpanderModule sets ExpanderModule field to given value.

### HasExpanderModule

`func (o *EquipmentExpanderModuleOperation) HasExpanderModule() bool`

HasExpanderModule returns a boolean if a field has been set.

### SetExpanderModuleNil

`func (o *EquipmentExpanderModuleOperation) SetExpanderModuleNil(b bool)`

 SetExpanderModuleNil sets the value for ExpanderModule to be an explicit nil

### UnsetExpanderModule
`func (o *EquipmentExpanderModuleOperation) UnsetExpanderModule()`

UnsetExpanderModule ensures that no value is present for ExpanderModule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


