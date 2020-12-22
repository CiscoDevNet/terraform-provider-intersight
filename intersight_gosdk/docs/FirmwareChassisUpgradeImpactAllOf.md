# FirmwareChassisUpgradeImpactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ChassisUpgradeImpact"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ChassisUpgradeImpact"]
**ImpactDetail** | Pointer to [**[]FirmwareComponentImpact**](FirmwareComponentImpact.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the chassis that will be affected by the upgrade. | [optional] 
**UserLabel** | Pointer to **string** | Details for the chassis that will be impacted by the upgrade. | [optional] 

## Methods

### NewFirmwareChassisUpgradeImpactAllOf

`func NewFirmwareChassisUpgradeImpactAllOf(classId string, objectType string, ) *FirmwareChassisUpgradeImpactAllOf`

NewFirmwareChassisUpgradeImpactAllOf instantiates a new FirmwareChassisUpgradeImpactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareChassisUpgradeImpactAllOfWithDefaults

`func NewFirmwareChassisUpgradeImpactAllOfWithDefaults() *FirmwareChassisUpgradeImpactAllOf`

NewFirmwareChassisUpgradeImpactAllOfWithDefaults instantiates a new FirmwareChassisUpgradeImpactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareChassisUpgradeImpactAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareChassisUpgradeImpactAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareChassisUpgradeImpactAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareChassisUpgradeImpactAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareChassisUpgradeImpactAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareChassisUpgradeImpactAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImpactDetail

`func (o *FirmwareChassisUpgradeImpactAllOf) GetImpactDetail() []FirmwareComponentImpact`

GetImpactDetail returns the ImpactDetail field if non-nil, zero value otherwise.

### GetImpactDetailOk

`func (o *FirmwareChassisUpgradeImpactAllOf) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool)`

GetImpactDetailOk returns a tuple with the ImpactDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDetail

`func (o *FirmwareChassisUpgradeImpactAllOf) SetImpactDetail(v []FirmwareComponentImpact)`

SetImpactDetail sets ImpactDetail field to given value.

### HasImpactDetail

`func (o *FirmwareChassisUpgradeImpactAllOf) HasImpactDetail() bool`

HasImpactDetail returns a boolean if a field has been set.

### SetImpactDetailNil

`func (o *FirmwareChassisUpgradeImpactAllOf) SetImpactDetailNil(b bool)`

 SetImpactDetailNil sets the value for ImpactDetail to be an explicit nil

### UnsetImpactDetail
`func (o *FirmwareChassisUpgradeImpactAllOf) UnsetImpactDetail()`

UnsetImpactDetail ensures that no value is present for ImpactDetail, not even an explicit nil
### GetName

`func (o *FirmwareChassisUpgradeImpactAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareChassisUpgradeImpactAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareChassisUpgradeImpactAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirmwareChassisUpgradeImpactAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLabel

`func (o *FirmwareChassisUpgradeImpactAllOf) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FirmwareChassisUpgradeImpactAllOf) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FirmwareChassisUpgradeImpactAllOf) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FirmwareChassisUpgradeImpactAllOf) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


