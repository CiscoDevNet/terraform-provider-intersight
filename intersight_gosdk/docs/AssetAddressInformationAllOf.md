# AssetAddressInformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.AddressInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.AddressInformation"]
**Address1** | Pointer to **string** | Address Line one of the address information. example \&quot;PO BOX 641570\&quot;. | [optional] [readonly] 
**Address2** | Pointer to **string** | Address Line two of the address information. example \&quot;Cisco Systems\&quot;. | [optional] [readonly] 
**Address3** | Pointer to **string** | Address Line three of the address information. example \&quot;Cisco Systems\&quot;. | [optional] [readonly] 
**City** | Pointer to **string** | City in which the address resides. example \&quot;San Jose\&quot;. | [optional] [readonly] 
**Country** | Pointer to **string** | Country in which the address resides. example \&quot;US\&quot;. | [optional] [readonly] 
**County** | Pointer to **string** | County in which the address resides. example \&quot;Washington County\&quot;. | [optional] [readonly] 
**Location** | Pointer to **string** | Location in which the address resides. example \&quot;14852\&quot;. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the user whose address is being populated. | [optional] [readonly] 
**PostalCode** | Pointer to **string** | Postal Code in which the address resides. example \&quot;95164-1570\&quot;. | [optional] [readonly] 
**Province** | Pointer to **string** | Province in which the address resides. example \&quot;AB\&quot;. | [optional] [readonly] 
**State** | Pointer to **string** | State in which the address resides. example \&quot;CA\&quot;. | [optional] [readonly] 

## Methods

### NewAssetAddressInformationAllOf

`func NewAssetAddressInformationAllOf(classId string, objectType string, ) *AssetAddressInformationAllOf`

NewAssetAddressInformationAllOf instantiates a new AssetAddressInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetAddressInformationAllOfWithDefaults

`func NewAssetAddressInformationAllOfWithDefaults() *AssetAddressInformationAllOf`

NewAssetAddressInformationAllOfWithDefaults instantiates a new AssetAddressInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetAddressInformationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetAddressInformationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetAddressInformationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetAddressInformationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetAddressInformationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetAddressInformationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress1

`func (o *AssetAddressInformationAllOf) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *AssetAddressInformationAllOf) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *AssetAddressInformationAllOf) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *AssetAddressInformationAllOf) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *AssetAddressInformationAllOf) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *AssetAddressInformationAllOf) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *AssetAddressInformationAllOf) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *AssetAddressInformationAllOf) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *AssetAddressInformationAllOf) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *AssetAddressInformationAllOf) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *AssetAddressInformationAllOf) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *AssetAddressInformationAllOf) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetCity

`func (o *AssetAddressInformationAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AssetAddressInformationAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AssetAddressInformationAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AssetAddressInformationAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AssetAddressInformationAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AssetAddressInformationAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AssetAddressInformationAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AssetAddressInformationAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *AssetAddressInformationAllOf) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *AssetAddressInformationAllOf) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *AssetAddressInformationAllOf) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *AssetAddressInformationAllOf) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetLocation

`func (o *AssetAddressInformationAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AssetAddressInformationAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AssetAddressInformationAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AssetAddressInformationAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *AssetAddressInformationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetAddressInformationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetAddressInformationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetAddressInformationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostalCode

`func (o *AssetAddressInformationAllOf) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AssetAddressInformationAllOf) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AssetAddressInformationAllOf) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AssetAddressInformationAllOf) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProvince

`func (o *AssetAddressInformationAllOf) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *AssetAddressInformationAllOf) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *AssetAddressInformationAllOf) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *AssetAddressInformationAllOf) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetState

`func (o *AssetAddressInformationAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AssetAddressInformationAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AssetAddressInformationAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AssetAddressInformationAllOf) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


