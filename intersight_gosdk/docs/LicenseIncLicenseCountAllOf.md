# LicenseIncLicenseCountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.IncLicenseCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.IncLicenseCount"]
**Premier100GfxCount** | Pointer to **int64** | The total number of devices claimed in the premier 100G fixed tier Intersight Nexus Cloud. | [optional] [readonly] 
**Premier10GfxCount** | Pointer to **int64** | The total number of devices claimed in the premier 10G fixed tier Intersight Nexus Cloud. | [optional] [readonly] 
**Premier1GfxCount** | Pointer to **int64** | The total number of devices claimed in the premier 1G fixed tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierD2OpsFixedCount** | Pointer to **int64** | The total number of devices claimed in the D2Ops Fixed premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierD2OpsModCount** | Pointer to **int64** | The total number of devices claimed in the D2Ops modular premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierMod4SlotCount** | Pointer to **int64** | The total number of devices claimed in the modular 4 slot premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**PremierMod8SlotCount** | Pointer to **int64** | The total number of devices claimed in the modular 8 slot premier tier Intersight Nexus Cloud. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](LicenseAccountLicenseDataRelationship.md) |  | [optional] 

## Methods

### NewLicenseIncLicenseCountAllOf

`func NewLicenseIncLicenseCountAllOf(classId string, objectType string, ) *LicenseIncLicenseCountAllOf`

NewLicenseIncLicenseCountAllOf instantiates a new LicenseIncLicenseCountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIncLicenseCountAllOfWithDefaults

`func NewLicenseIncLicenseCountAllOfWithDefaults() *LicenseIncLicenseCountAllOf`

NewLicenseIncLicenseCountAllOfWithDefaults instantiates a new LicenseIncLicenseCountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseIncLicenseCountAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseIncLicenseCountAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseIncLicenseCountAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseIncLicenseCountAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseIncLicenseCountAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseIncLicenseCountAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPremier100GfxCount

`func (o *LicenseIncLicenseCountAllOf) GetPremier100GfxCount() int64`

GetPremier100GfxCount returns the Premier100GfxCount field if non-nil, zero value otherwise.

### GetPremier100GfxCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremier100GfxCountOk() (*int64, bool)`

GetPremier100GfxCountOk returns a tuple with the Premier100GfxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremier100GfxCount

`func (o *LicenseIncLicenseCountAllOf) SetPremier100GfxCount(v int64)`

SetPremier100GfxCount sets Premier100GfxCount field to given value.

### HasPremier100GfxCount

`func (o *LicenseIncLicenseCountAllOf) HasPremier100GfxCount() bool`

HasPremier100GfxCount returns a boolean if a field has been set.

### GetPremier10GfxCount

`func (o *LicenseIncLicenseCountAllOf) GetPremier10GfxCount() int64`

GetPremier10GfxCount returns the Premier10GfxCount field if non-nil, zero value otherwise.

### GetPremier10GfxCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremier10GfxCountOk() (*int64, bool)`

GetPremier10GfxCountOk returns a tuple with the Premier10GfxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremier10GfxCount

`func (o *LicenseIncLicenseCountAllOf) SetPremier10GfxCount(v int64)`

SetPremier10GfxCount sets Premier10GfxCount field to given value.

### HasPremier10GfxCount

`func (o *LicenseIncLicenseCountAllOf) HasPremier10GfxCount() bool`

HasPremier10GfxCount returns a boolean if a field has been set.

### GetPremier1GfxCount

`func (o *LicenseIncLicenseCountAllOf) GetPremier1GfxCount() int64`

GetPremier1GfxCount returns the Premier1GfxCount field if non-nil, zero value otherwise.

### GetPremier1GfxCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremier1GfxCountOk() (*int64, bool)`

GetPremier1GfxCountOk returns a tuple with the Premier1GfxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremier1GfxCount

`func (o *LicenseIncLicenseCountAllOf) SetPremier1GfxCount(v int64)`

SetPremier1GfxCount sets Premier1GfxCount field to given value.

### HasPremier1GfxCount

`func (o *LicenseIncLicenseCountAllOf) HasPremier1GfxCount() bool`

HasPremier1GfxCount returns a boolean if a field has been set.

### GetPremierD2OpsFixedCount

`func (o *LicenseIncLicenseCountAllOf) GetPremierD2OpsFixedCount() int64`

GetPremierD2OpsFixedCount returns the PremierD2OpsFixedCount field if non-nil, zero value otherwise.

### GetPremierD2OpsFixedCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremierD2OpsFixedCountOk() (*int64, bool)`

GetPremierD2OpsFixedCountOk returns a tuple with the PremierD2OpsFixedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierD2OpsFixedCount

`func (o *LicenseIncLicenseCountAllOf) SetPremierD2OpsFixedCount(v int64)`

SetPremierD2OpsFixedCount sets PremierD2OpsFixedCount field to given value.

### HasPremierD2OpsFixedCount

`func (o *LicenseIncLicenseCountAllOf) HasPremierD2OpsFixedCount() bool`

HasPremierD2OpsFixedCount returns a boolean if a field has been set.

### GetPremierD2OpsModCount

`func (o *LicenseIncLicenseCountAllOf) GetPremierD2OpsModCount() int64`

GetPremierD2OpsModCount returns the PremierD2OpsModCount field if non-nil, zero value otherwise.

### GetPremierD2OpsModCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremierD2OpsModCountOk() (*int64, bool)`

GetPremierD2OpsModCountOk returns a tuple with the PremierD2OpsModCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierD2OpsModCount

`func (o *LicenseIncLicenseCountAllOf) SetPremierD2OpsModCount(v int64)`

SetPremierD2OpsModCount sets PremierD2OpsModCount field to given value.

### HasPremierD2OpsModCount

`func (o *LicenseIncLicenseCountAllOf) HasPremierD2OpsModCount() bool`

HasPremierD2OpsModCount returns a boolean if a field has been set.

### GetPremierMod4SlotCount

`func (o *LicenseIncLicenseCountAllOf) GetPremierMod4SlotCount() int64`

GetPremierMod4SlotCount returns the PremierMod4SlotCount field if non-nil, zero value otherwise.

### GetPremierMod4SlotCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremierMod4SlotCountOk() (*int64, bool)`

GetPremierMod4SlotCountOk returns a tuple with the PremierMod4SlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierMod4SlotCount

`func (o *LicenseIncLicenseCountAllOf) SetPremierMod4SlotCount(v int64)`

SetPremierMod4SlotCount sets PremierMod4SlotCount field to given value.

### HasPremierMod4SlotCount

`func (o *LicenseIncLicenseCountAllOf) HasPremierMod4SlotCount() bool`

HasPremierMod4SlotCount returns a boolean if a field has been set.

### GetPremierMod8SlotCount

`func (o *LicenseIncLicenseCountAllOf) GetPremierMod8SlotCount() int64`

GetPremierMod8SlotCount returns the PremierMod8SlotCount field if non-nil, zero value otherwise.

### GetPremierMod8SlotCountOk

`func (o *LicenseIncLicenseCountAllOf) GetPremierMod8SlotCountOk() (*int64, bool)`

GetPremierMod8SlotCountOk returns a tuple with the PremierMod8SlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremierMod8SlotCount

`func (o *LicenseIncLicenseCountAllOf) SetPremierMod8SlotCount(v int64)`

SetPremierMod8SlotCount sets PremierMod8SlotCount field to given value.

### HasPremierMod8SlotCount

`func (o *LicenseIncLicenseCountAllOf) HasPremierMod8SlotCount() bool`

HasPremierMod8SlotCount returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseIncLicenseCountAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseIncLicenseCountAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseIncLicenseCountAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseIncLicenseCountAllOf) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


