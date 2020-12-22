# FirmwareFabricUpgradeImpactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.FabricUpgradeImpact"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.FabricUpgradeImpact"]
**ImpactDetail** | Pointer to [**[]FirmwareComponentImpact**](FirmwareComponentImpact.md) |  | [optional] 
**Serial** | Pointer to **string** | Details for the Fabric Interconnect that will be impacted by the upgrade. | [optional] 

## Methods

### NewFirmwareFabricUpgradeImpactAllOf

`func NewFirmwareFabricUpgradeImpactAllOf(classId string, objectType string, ) *FirmwareFabricUpgradeImpactAllOf`

NewFirmwareFabricUpgradeImpactAllOf instantiates a new FirmwareFabricUpgradeImpactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFabricUpgradeImpactAllOfWithDefaults

`func NewFirmwareFabricUpgradeImpactAllOfWithDefaults() *FirmwareFabricUpgradeImpactAllOf`

NewFirmwareFabricUpgradeImpactAllOfWithDefaults instantiates a new FirmwareFabricUpgradeImpactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareFabricUpgradeImpactAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareFabricUpgradeImpactAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareFabricUpgradeImpactAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareFabricUpgradeImpactAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareFabricUpgradeImpactAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareFabricUpgradeImpactAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImpactDetail

`func (o *FirmwareFabricUpgradeImpactAllOf) GetImpactDetail() []FirmwareComponentImpact`

GetImpactDetail returns the ImpactDetail field if non-nil, zero value otherwise.

### GetImpactDetailOk

`func (o *FirmwareFabricUpgradeImpactAllOf) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool)`

GetImpactDetailOk returns a tuple with the ImpactDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDetail

`func (o *FirmwareFabricUpgradeImpactAllOf) SetImpactDetail(v []FirmwareComponentImpact)`

SetImpactDetail sets ImpactDetail field to given value.

### HasImpactDetail

`func (o *FirmwareFabricUpgradeImpactAllOf) HasImpactDetail() bool`

HasImpactDetail returns a boolean if a field has been set.

### SetImpactDetailNil

`func (o *FirmwareFabricUpgradeImpactAllOf) SetImpactDetailNil(b bool)`

 SetImpactDetailNil sets the value for ImpactDetail to be an explicit nil

### UnsetImpactDetail
`func (o *FirmwareFabricUpgradeImpactAllOf) UnsetImpactDetail()`

UnsetImpactDetail ensures that no value is present for ImpactDetail, not even an explicit nil
### GetSerial

`func (o *FirmwareFabricUpgradeImpactAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *FirmwareFabricUpgradeImpactAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *FirmwareFabricUpgradeImpactAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *FirmwareFabricUpgradeImpactAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


