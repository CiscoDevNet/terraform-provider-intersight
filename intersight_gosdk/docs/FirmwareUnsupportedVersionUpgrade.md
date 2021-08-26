# FirmwareUnsupportedVersionUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.UnsupportedVersionUpgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.UnsupportedVersionUpgrade"]
**UpgradeStatus** | Pointer to **string** | Workflow status of firmware upgrade. * &#x60;None&#x60; - Upgrade status is none when upgrade is in progress. * &#x60;Completed&#x60; - Upgrade completed successfully. * &#x60;Failed&#x60; - Upgrade status is failed when upgrade has failed. | [optional] [default to "None"]
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Distributable** | Pointer to [**FirmwareDistributableRelationship**](FirmwareDistributableRelationship.md) |  | [optional] 
**PhysicalIdentity** | Pointer to [**EquipmentPhysicalIdentityRelationship**](EquipmentPhysicalIdentityRelationship.md) |  | [optional] 

## Methods

### NewFirmwareUnsupportedVersionUpgrade

`func NewFirmwareUnsupportedVersionUpgrade(classId string, objectType string, ) *FirmwareUnsupportedVersionUpgrade`

NewFirmwareUnsupportedVersionUpgrade instantiates a new FirmwareUnsupportedVersionUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUnsupportedVersionUpgradeWithDefaults

`func NewFirmwareUnsupportedVersionUpgradeWithDefaults() *FirmwareUnsupportedVersionUpgrade`

NewFirmwareUnsupportedVersionUpgradeWithDefaults instantiates a new FirmwareUnsupportedVersionUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUnsupportedVersionUpgrade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUnsupportedVersionUpgrade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUnsupportedVersionUpgrade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUnsupportedVersionUpgrade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUnsupportedVersionUpgrade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUnsupportedVersionUpgrade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUpgradeStatus

`func (o *FirmwareUnsupportedVersionUpgrade) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *FirmwareUnsupportedVersionUpgrade) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *FirmwareUnsupportedVersionUpgrade) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *FirmwareUnsupportedVersionUpgrade) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetDevice

`func (o *FirmwareUnsupportedVersionUpgrade) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUnsupportedVersionUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUnsupportedVersionUpgrade) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUnsupportedVersionUpgrade) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDistributable

`func (o *FirmwareUnsupportedVersionUpgrade) GetDistributable() FirmwareDistributableRelationship`

GetDistributable returns the Distributable field if non-nil, zero value otherwise.

### GetDistributableOk

`func (o *FirmwareUnsupportedVersionUpgrade) GetDistributableOk() (*FirmwareDistributableRelationship, bool)`

GetDistributableOk returns a tuple with the Distributable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributable

`func (o *FirmwareUnsupportedVersionUpgrade) SetDistributable(v FirmwareDistributableRelationship)`

SetDistributable sets Distributable field to given value.

### HasDistributable

`func (o *FirmwareUnsupportedVersionUpgrade) HasDistributable() bool`

HasDistributable returns a boolean if a field has been set.

### GetPhysicalIdentity

`func (o *FirmwareUnsupportedVersionUpgrade) GetPhysicalIdentity() EquipmentPhysicalIdentityRelationship`

GetPhysicalIdentity returns the PhysicalIdentity field if non-nil, zero value otherwise.

### GetPhysicalIdentityOk

`func (o *FirmwareUnsupportedVersionUpgrade) GetPhysicalIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool)`

GetPhysicalIdentityOk returns a tuple with the PhysicalIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalIdentity

`func (o *FirmwareUnsupportedVersionUpgrade) SetPhysicalIdentity(v EquipmentPhysicalIdentityRelationship)`

SetPhysicalIdentity sets PhysicalIdentity field to given value.

### HasPhysicalIdentity

`func (o *FirmwareUnsupportedVersionUpgrade) HasPhysicalIdentity() bool`

HasPhysicalIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


