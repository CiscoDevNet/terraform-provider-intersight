# EquipmentIoCardOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.IoCardOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.IoCardOperation"]
**AdminPeerPowerState** | Pointer to **string** | User configured power state of the peer IO module. * &#x60;None&#x60; - Placeholder default value for iom power state property. * &#x60;Reboot&#x60; - IO Module reboot state property value. | [optional] [default to "None"]
**AdminPowerState** | Pointer to **string** | User configured power state of the IO module. * &#x60;None&#x60; - Placeholder default value for iom power state property. * &#x60;Reboot&#x60; - IO Module reboot state property value. | [optional] [default to "None"]
**AffectedObjName** | Pointer to **string** | Placeholder for affected object name which is a combination of chassis and IOM ID. Used to display affected object in audit log. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | The configured state of these settings in the target IO module. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target IO module. Applying - This state denotes that the settings are being applied in the target IO module. Failed - This state denotes that the settings could not be applied in the target IO module. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**IoCardOperationStatus** | Pointer to [**[]EquipmentIoCardOperationStatus**](EquipmentIoCardOperationStatus.md) |  | [optional] 
**DeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**IoCard** | Pointer to [**NullableEquipmentIoCardRelationship**](EquipmentIoCardRelationship.md) |  | [optional] 

## Methods

### NewEquipmentIoCardOperation

`func NewEquipmentIoCardOperation(classId string, objectType string, ) *EquipmentIoCardOperation`

NewEquipmentIoCardOperation instantiates a new EquipmentIoCardOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardOperationWithDefaults

`func NewEquipmentIoCardOperationWithDefaults() *EquipmentIoCardOperation`

NewEquipmentIoCardOperationWithDefaults instantiates a new EquipmentIoCardOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIoCardOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIoCardOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIoCardOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIoCardOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIoCardOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIoCardOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminPeerPowerState

`func (o *EquipmentIoCardOperation) GetAdminPeerPowerState() string`

GetAdminPeerPowerState returns the AdminPeerPowerState field if non-nil, zero value otherwise.

### GetAdminPeerPowerStateOk

`func (o *EquipmentIoCardOperation) GetAdminPeerPowerStateOk() (*string, bool)`

GetAdminPeerPowerStateOk returns a tuple with the AdminPeerPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPeerPowerState

`func (o *EquipmentIoCardOperation) SetAdminPeerPowerState(v string)`

SetAdminPeerPowerState sets AdminPeerPowerState field to given value.

### HasAdminPeerPowerState

`func (o *EquipmentIoCardOperation) HasAdminPeerPowerState() bool`

HasAdminPeerPowerState returns a boolean if a field has been set.

### GetAdminPowerState

`func (o *EquipmentIoCardOperation) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *EquipmentIoCardOperation) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *EquipmentIoCardOperation) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *EquipmentIoCardOperation) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAffectedObjName

`func (o *EquipmentIoCardOperation) GetAffectedObjName() string`

GetAffectedObjName returns the AffectedObjName field if non-nil, zero value otherwise.

### GetAffectedObjNameOk

`func (o *EquipmentIoCardOperation) GetAffectedObjNameOk() (*string, bool)`

GetAffectedObjNameOk returns a tuple with the AffectedObjName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObjName

`func (o *EquipmentIoCardOperation) SetAffectedObjName(v string)`

SetAffectedObjName sets AffectedObjName field to given value.

### HasAffectedObjName

`func (o *EquipmentIoCardOperation) HasAffectedObjName() bool`

HasAffectedObjName returns a boolean if a field has been set.

### GetConfigState

`func (o *EquipmentIoCardOperation) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentIoCardOperation) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentIoCardOperation) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentIoCardOperation) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetIoCardOperationStatus

`func (o *EquipmentIoCardOperation) GetIoCardOperationStatus() []EquipmentIoCardOperationStatus`

GetIoCardOperationStatus returns the IoCardOperationStatus field if non-nil, zero value otherwise.

### GetIoCardOperationStatusOk

`func (o *EquipmentIoCardOperation) GetIoCardOperationStatusOk() (*[]EquipmentIoCardOperationStatus, bool)`

GetIoCardOperationStatusOk returns a tuple with the IoCardOperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardOperationStatus

`func (o *EquipmentIoCardOperation) SetIoCardOperationStatus(v []EquipmentIoCardOperationStatus)`

SetIoCardOperationStatus sets IoCardOperationStatus field to given value.

### HasIoCardOperationStatus

`func (o *EquipmentIoCardOperation) HasIoCardOperationStatus() bool`

HasIoCardOperationStatus returns a boolean if a field has been set.

### SetIoCardOperationStatusNil

`func (o *EquipmentIoCardOperation) SetIoCardOperationStatusNil(b bool)`

 SetIoCardOperationStatusNil sets the value for IoCardOperationStatus to be an explicit nil

### UnsetIoCardOperationStatus
`func (o *EquipmentIoCardOperation) UnsetIoCardOperationStatus()`

UnsetIoCardOperationStatus ensures that no value is present for IoCardOperationStatus, not even an explicit nil
### GetDeviceRegistration

`func (o *EquipmentIoCardOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentIoCardOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentIoCardOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentIoCardOperation) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### SetDeviceRegistrationNil

`func (o *EquipmentIoCardOperation) SetDeviceRegistrationNil(b bool)`

 SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil

### UnsetDeviceRegistration
`func (o *EquipmentIoCardOperation) UnsetDeviceRegistration()`

UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil
### GetIoCard

`func (o *EquipmentIoCardOperation) GetIoCard() EquipmentIoCardRelationship`

GetIoCard returns the IoCard field if non-nil, zero value otherwise.

### GetIoCardOk

`func (o *EquipmentIoCardOperation) GetIoCardOk() (*EquipmentIoCardRelationship, bool)`

GetIoCardOk returns a tuple with the IoCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCard

`func (o *EquipmentIoCardOperation) SetIoCard(v EquipmentIoCardRelationship)`

SetIoCard sets IoCard field to given value.

### HasIoCard

`func (o *EquipmentIoCardOperation) HasIoCard() bool`

HasIoCard returns a boolean if a field has been set.

### SetIoCardNil

`func (o *EquipmentIoCardOperation) SetIoCardNil(b bool)`

 SetIoCardNil sets the value for IoCard to be an explicit nil

### UnsetIoCard
`func (o *EquipmentIoCardOperation) UnsetIoCard()`

UnsetIoCard ensures that no value is present for IoCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


