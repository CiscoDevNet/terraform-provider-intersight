# FirmwareUpgradeImpactStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.UpgradeImpactStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.UpgradeImpactStatus"]
**Upgrade** | Pointer to [**FirmwareUpgradeBaseRelationship**](FirmwareUpgradeBaseRelationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeImpactStatus

`func NewFirmwareUpgradeImpactStatus(classId string, objectType string, ) *FirmwareUpgradeImpactStatus`

NewFirmwareUpgradeImpactStatus instantiates a new FirmwareUpgradeImpactStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeImpactStatusWithDefaults

`func NewFirmwareUpgradeImpactStatusWithDefaults() *FirmwareUpgradeImpactStatus`

NewFirmwareUpgradeImpactStatusWithDefaults instantiates a new FirmwareUpgradeImpactStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgradeImpactStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeImpactStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeImpactStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgradeImpactStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeImpactStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeImpactStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUpgrade

`func (o *FirmwareUpgradeImpactStatus) GetUpgrade() FirmwareUpgradeBaseRelationship`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *FirmwareUpgradeImpactStatus) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *FirmwareUpgradeImpactStatus) SetUpgrade(v FirmwareUpgradeBaseRelationship)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *FirmwareUpgradeImpactStatus) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


