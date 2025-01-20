# HciLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.License"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.License"]
**Category** | Pointer to **string** | The category of a license instance. | [optional] [readonly] 
**ConsumptionDetails** | Pointer to [**[]HciConsumption**](HciConsumption.md) |  | [optional] 
**ExpiryDate** | Pointer to **time.Time** | The expiry date of a license instance. | [optional] [readonly] 
**ExtId** | Pointer to **string** | The unique identifier of a license. | [optional] [readonly] 
**LicenseId** | Pointer to **string** | The identifier of a license, usually in LIC-xxxx format. | [optional] [readonly] 
**Meter** | Pointer to **string** | A license capacity is expressed in terms of meter. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of a license instance. | [optional] [readonly] 
**PcExtId** | Pointer to **string** | The unique identifier of the domain manager (Prism Central) instance which manages this cluster. | [optional] [readonly] 
**Quantity** | Pointer to **float64** | The scope defines where licenses can be applied. | [optional] [readonly] 
**Scope** | Pointer to **string** | The scope defines where licenses can be applied. | [optional] [readonly] 
**SubCategory** | Pointer to **string** | The subCategory of a license instance, such as if it is an add-on or with unlimited capacity. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of a license instance. | [optional] [readonly] 
**DomainManager** | Pointer to [**NullableHciDomainManagerRelationship**](HciDomainManagerRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciLicense

`func NewHciLicense(classId string, objectType string, ) *HciLicense`

NewHciLicense instantiates a new HciLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciLicenseWithDefaults

`func NewHciLicenseWithDefaults() *HciLicense`

NewHciLicenseWithDefaults instantiates a new HciLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciLicense) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciLicense) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciLicense) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciLicense) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciLicense) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciLicense) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HciLicense) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HciLicense) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HciLicense) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HciLicense) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConsumptionDetails

`func (o *HciLicense) GetConsumptionDetails() []HciConsumption`

GetConsumptionDetails returns the ConsumptionDetails field if non-nil, zero value otherwise.

### GetConsumptionDetailsOk

`func (o *HciLicense) GetConsumptionDetailsOk() (*[]HciConsumption, bool)`

GetConsumptionDetailsOk returns a tuple with the ConsumptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionDetails

`func (o *HciLicense) SetConsumptionDetails(v []HciConsumption)`

SetConsumptionDetails sets ConsumptionDetails field to given value.

### HasConsumptionDetails

`func (o *HciLicense) HasConsumptionDetails() bool`

HasConsumptionDetails returns a boolean if a field has been set.

### SetConsumptionDetailsNil

`func (o *HciLicense) SetConsumptionDetailsNil(b bool)`

 SetConsumptionDetailsNil sets the value for ConsumptionDetails to be an explicit nil

### UnsetConsumptionDetails
`func (o *HciLicense) UnsetConsumptionDetails()`

UnsetConsumptionDetails ensures that no value is present for ConsumptionDetails, not even an explicit nil
### GetExpiryDate

`func (o *HciLicense) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *HciLicense) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *HciLicense) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *HciLicense) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetExtId

`func (o *HciLicense) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *HciLicense) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *HciLicense) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *HciLicense) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetLicenseId

`func (o *HciLicense) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *HciLicense) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *HciLicense) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *HciLicense) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetMeter

`func (o *HciLicense) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *HciLicense) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *HciLicense) SetMeter(v string)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *HciLicense) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetName

`func (o *HciLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciLicense) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciLicense) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcExtId

`func (o *HciLicense) GetPcExtId() string`

GetPcExtId returns the PcExtId field if non-nil, zero value otherwise.

### GetPcExtIdOk

`func (o *HciLicense) GetPcExtIdOk() (*string, bool)`

GetPcExtIdOk returns a tuple with the PcExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcExtId

`func (o *HciLicense) SetPcExtId(v string)`

SetPcExtId sets PcExtId field to given value.

### HasPcExtId

`func (o *HciLicense) HasPcExtId() bool`

HasPcExtId returns a boolean if a field has been set.

### GetQuantity

`func (o *HciLicense) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *HciLicense) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *HciLicense) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *HciLicense) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetScope

`func (o *HciLicense) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HciLicense) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HciLicense) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HciLicense) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSubCategory

`func (o *HciLicense) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *HciLicense) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *HciLicense) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *HciLicense) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetType

`func (o *HciLicense) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciLicense) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciLicense) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciLicense) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomainManager

`func (o *HciLicense) GetDomainManager() HciDomainManagerRelationship`

GetDomainManager returns the DomainManager field if non-nil, zero value otherwise.

### GetDomainManagerOk

`func (o *HciLicense) GetDomainManagerOk() (*HciDomainManagerRelationship, bool)`

GetDomainManagerOk returns a tuple with the DomainManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainManager

`func (o *HciLicense) SetDomainManager(v HciDomainManagerRelationship)`

SetDomainManager sets DomainManager field to given value.

### HasDomainManager

`func (o *HciLicense) HasDomainManager() bool`

HasDomainManager returns a boolean if a field has been set.

### SetDomainManagerNil

`func (o *HciLicense) SetDomainManagerNil(b bool)`

 SetDomainManagerNil sets the value for DomainManager to be an explicit nil

### UnsetDomainManager
`func (o *HciLicense) UnsetDomainManager()`

UnsetDomainManager ensures that no value is present for DomainManager, not even an explicit nil
### GetRegisteredDevice

`func (o *HciLicense) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciLicense) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciLicense) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciLicense) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciLicense) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciLicense) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


