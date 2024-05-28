# LicenseIncLicenseCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IncLicenseCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IncLicenseCount"]
**Premier100GfxCount** | Pointer to **int64** | The total number of devices claimed in the premier 100G fixed tier Intersight Nexus Cloud. | [optional] [readonly] 
**Premier10GfxCount** | Pointer to **int64** | The total number of devices claimed in the premier 10G fixed tier Intersight Nexus Cloud. | [optional] [readonly] 
**Premier1GfxCount** | Pointer to **int64** | The total number of devices claimed in the premier 1G fixed tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierCentralizedMod8SlotCount** | Pointer to **int64** | The total number of devices claimed in the CentralizedMod8Slot premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierD2OpsFixedCount** | Pointer to **int64** | The total number of devices claimed in the D2Ops Fixed premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierD2OpsModCount** | Pointer to **int64** | The total number of devices claimed in the D2Ops modular premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierDistributedMod8SlotCount** | Pointer to **int64** | The total number of devices claimed in the DistributedMod8Slot premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierMod4SlotCount** | Pointer to **int64** | The total number of devices claimed in the modular 4 slot premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierMod8SlotCount** | Pointer to **int64** | The total number of devices claimed in the modular 8 slot premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**NullableLicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIncLicenseCount

`func NewLicenseIncLicenseCount(classId string, objectType string, ) *LicenseIncLicenseCount`

NewLicenseIncLicenseCount instantiates a new LicenseIncLicenseCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIncLicenseCountWithDefaults

`func NewLicenseIncLicenseCountWithDefaults() *LicenseIncLicenseCount`

NewLicenseIncLicenseCountWithDefaults instantiates a new LicenseIncLicenseCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIncLicenseCount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIncLicenseCount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIncLicenseCount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIncLicenseCount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIncLicenseCount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIncLicenseCount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPremier100GfxCount

`func (o *LicenseIncLicenseCount) GetPremier100GfxCount() int64`

GetPremier100GfxCount returns the Premier100GfxCount field if non-nil, zero value otherwise.

### GetPremier100GfxCountOk

`func (o *LicenseIncLicenseCount) GetPremier100GfxCountOk() (*int64, bool)`

GetPremier100GfxCountOk returns a tuple with the Premier100GfxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremier100GfxCount

`func (o *LicenseIncLicenseCount) SetPremier100GfxCount(v int64)`

SetPremier100GfxCount sets Premier100GfxCount field to given value.

### HasPremier100GfxCount

`func (o *LicenseIncLicenseCount) HasPremier100GfxCount() bool`

HasPremier100GfxCount returns a boolean if a field has been set.

### GetPremier10GfxCount

`func (o *LicenseIncLicenseCount) GetPremier10GfxCount() int64`

GetPremier10GfxCount returns the Premier10GfxCount field if non-nil, zero value otherwise.

### GetPremier10GfxCountOk

`func (o *LicenseIncLicenseCount) GetPremier10GfxCountOk() (*int64, bool)`

GetPremier10GfxCountOk returns a tuple with the Premier10GfxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremier10GfxCount

`func (o *LicenseIncLicenseCount) SetPremier10GfxCount(v int64)`

SetPremier10GfxCount sets Premier10GfxCount field to given value.

### HasPremier10GfxCount

`func (o *LicenseIncLicenseCount) HasPremier10GfxCount() bool`

HasPremier10GfxCount returns a boolean if a field has been set.

### GetPremier1GfxCount

`func (o *LicenseIncLicenseCount) GetPremier1GfxCount() int64`

GetPremier1GfxCount returns the Premier1GfxCount field if non-nil, zero value otherwise.

### GetPremier1GfxCountOk

`func (o *LicenseIncLicenseCount) GetPremier1GfxCountOk() (*int64, bool)`

GetPremier1GfxCountOk returns a tuple with the Premier1GfxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremier1GfxCount

`func (o *LicenseIncLicenseCount) SetPremier1GfxCount(v int64)`

SetPremier1GfxCount sets Premier1GfxCount field to given value.

### HasPremier1GfxCount

`func (o *LicenseIncLicenseCount) HasPremier1GfxCount() bool`

HasPremier1GfxCount returns a boolean if a field has been set.

### GetPremierCentralizedMod8SlotCount

`func (o *LicenseIncLicenseCount) GetPremierCentralizedMod8SlotCount() int64`

GetPremierCentralizedMod8SlotCount returns the PremierCentralizedMod8SlotCount field if non-nil, zero value otherwise.

### GetPremierCentralizedMod8SlotCountOk

`func (o *LicenseIncLicenseCount) GetPremierCentralizedMod8SlotCountOk() (*int64, bool)`

GetPremierCentralizedMod8SlotCountOk returns a tuple with the PremierCentralizedMod8SlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierCentralizedMod8SlotCount

`func (o *LicenseIncLicenseCount) SetPremierCentralizedMod8SlotCount(v int64)`

SetPremierCentralizedMod8SlotCount sets PremierCentralizedMod8SlotCount field to given value.

### HasPremierCentralizedMod8SlotCount

`func (o *LicenseIncLicenseCount) HasPremierCentralizedMod8SlotCount() bool`

HasPremierCentralizedMod8SlotCount returns a boolean if a field has been set.

### GetPremierD2OpsFixedCount

`func (o *LicenseIncLicenseCount) GetPremierD2OpsFixedCount() int64`

GetPremierD2OpsFixedCount returns the PremierD2OpsFixedCount field if non-nil, zero value otherwise.

### GetPremierD2OpsFixedCountOk

`func (o *LicenseIncLicenseCount) GetPremierD2OpsFixedCountOk() (*int64, bool)`

GetPremierD2OpsFixedCountOk returns a tuple with the PremierD2OpsFixedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierD2OpsFixedCount

`func (o *LicenseIncLicenseCount) SetPremierD2OpsFixedCount(v int64)`

SetPremierD2OpsFixedCount sets PremierD2OpsFixedCount field to given value.

### HasPremierD2OpsFixedCount

`func (o *LicenseIncLicenseCount) HasPremierD2OpsFixedCount() bool`

HasPremierD2OpsFixedCount returns a boolean if a field has been set.

### GetPremierD2OpsModCount

`func (o *LicenseIncLicenseCount) GetPremierD2OpsModCount() int64`

GetPremierD2OpsModCount returns the PremierD2OpsModCount field if non-nil, zero value otherwise.

### GetPremierD2OpsModCountOk

`func (o *LicenseIncLicenseCount) GetPremierD2OpsModCountOk() (*int64, bool)`

GetPremierD2OpsModCountOk returns a tuple with the PremierD2OpsModCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierD2OpsModCount

`func (o *LicenseIncLicenseCount) SetPremierD2OpsModCount(v int64)`

SetPremierD2OpsModCount sets PremierD2OpsModCount field to given value.

### HasPremierD2OpsModCount

`func (o *LicenseIncLicenseCount) HasPremierD2OpsModCount() bool`

HasPremierD2OpsModCount returns a boolean if a field has been set.

### GetPremierDistributedMod8SlotCount

`func (o *LicenseIncLicenseCount) GetPremierDistributedMod8SlotCount() int64`

GetPremierDistributedMod8SlotCount returns the PremierDistributedMod8SlotCount field if non-nil, zero value otherwise.

### GetPremierDistributedMod8SlotCountOk

`func (o *LicenseIncLicenseCount) GetPremierDistributedMod8SlotCountOk() (*int64, bool)`

GetPremierDistributedMod8SlotCountOk returns a tuple with the PremierDistributedMod8SlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierDistributedMod8SlotCount

`func (o *LicenseIncLicenseCount) SetPremierDistributedMod8SlotCount(v int64)`

SetPremierDistributedMod8SlotCount sets PremierDistributedMod8SlotCount field to given value.

### HasPremierDistributedMod8SlotCount

`func (o *LicenseIncLicenseCount) HasPremierDistributedMod8SlotCount() bool`

HasPremierDistributedMod8SlotCount returns a boolean if a field has been set.

### GetPremierMod4SlotCount

`func (o *LicenseIncLicenseCount) GetPremierMod4SlotCount() int64`

GetPremierMod4SlotCount returns the PremierMod4SlotCount field if non-nil, zero value otherwise.

### GetPremierMod4SlotCountOk

`func (o *LicenseIncLicenseCount) GetPremierMod4SlotCountOk() (*int64, bool)`

GetPremierMod4SlotCountOk returns a tuple with the PremierMod4SlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierMod4SlotCount

`func (o *LicenseIncLicenseCount) SetPremierMod4SlotCount(v int64)`

SetPremierMod4SlotCount sets PremierMod4SlotCount field to given value.

### HasPremierMod4SlotCount

`func (o *LicenseIncLicenseCount) HasPremierMod4SlotCount() bool`

HasPremierMod4SlotCount returns a boolean if a field has been set.

### GetPremierMod8SlotCount

`func (o *LicenseIncLicenseCount) GetPremierMod8SlotCount() int64`

GetPremierMod8SlotCount returns the PremierMod8SlotCount field if non-nil, zero value otherwise.

### GetPremierMod8SlotCountOk

`func (o *LicenseIncLicenseCount) GetPremierMod8SlotCountOk() (*int64, bool)`

GetPremierMod8SlotCountOk returns a tuple with the PremierMod8SlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierMod8SlotCount

`func (o *LicenseIncLicenseCount) SetPremierMod8SlotCount(v int64)`

SetPremierMod8SlotCount sets PremierMod8SlotCount field to given value.

### HasPremierMod8SlotCount

`func (o *LicenseIncLicenseCount) HasPremierMod8SlotCount() bool`

HasPremierMod8SlotCount returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIncLicenseCount) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIncLicenseCount) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIncLicenseCount) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIncLicenseCount) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.

### SetAccountLicenseDataNil

`func (o *LicenseIncLicenseCount) SetAccountLicenseDataNil(b bool)`

 SetAccountLicenseDataNil sets the value for AccountLicenseData to be an explicit nil

### UnsetAccountLicenseData
`func (o *LicenseIncLicenseCount) UnsetAccountLicenseData()`

UnsetAccountLicenseData ensures that no value is present for AccountLicenseData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


