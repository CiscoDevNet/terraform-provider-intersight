# AssetProductInformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillTo** | Pointer to [**AssetAddressInformation**](asset.AddressInformation.md) |  | [optional] 
**Description** | Pointer to **string** | Short description of the Cisco product that helps identify the product easily. example \&quot;DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\&quot;. | [optional] [readonly] 
**Family** | Pointer to **string** | Family that the product belongs to. Example \&quot;UCSB\&quot;. | [optional] [readonly] 
**Group** | Pointer to **string** | Group that the product belongs to. It is one higher level categorization above family. example \&quot;Switch\&quot;. | [optional] [readonly] 
**Number** | Pointer to **string** | Product number that identifies the product. example PID. example \&quot;UCS-FI-6248UP-CH2\&quot;. | [optional] [readonly] 
**ShipTo** | Pointer to [**AssetAddressInformation**](asset.AddressInformation.md) |  | [optional] 
**SubType** | Pointer to **string** | Sub type of the product being specified. example \&quot;UCS 6200 SER\&quot;. | [optional] [readonly] 

## Methods

### NewAssetProductInformationAllOf

`func NewAssetProductInformationAllOf() *AssetProductInformationAllOf`

NewAssetProductInformationAllOf instantiates a new AssetProductInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetProductInformationAllOfWithDefaults

`func NewAssetProductInformationAllOfWithDefaults() *AssetProductInformationAllOf`

NewAssetProductInformationAllOfWithDefaults instantiates a new AssetProductInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillTo

`func (o *AssetProductInformationAllOf) GetBillTo() AssetAddressInformation`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *AssetProductInformationAllOf) GetBillToOk() (*AssetAddressInformation, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *AssetProductInformationAllOf) SetBillTo(v AssetAddressInformation)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *AssetProductInformationAllOf) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### GetDescription

`func (o *AssetProductInformationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetProductInformationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetProductInformationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetProductInformationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamily

`func (o *AssetProductInformationAllOf) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AssetProductInformationAllOf) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AssetProductInformationAllOf) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AssetProductInformationAllOf) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetGroup

`func (o *AssetProductInformationAllOf) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AssetProductInformationAllOf) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AssetProductInformationAllOf) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AssetProductInformationAllOf) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetNumber

`func (o *AssetProductInformationAllOf) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AssetProductInformationAllOf) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AssetProductInformationAllOf) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AssetProductInformationAllOf) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipTo

`func (o *AssetProductInformationAllOf) GetShipTo() AssetAddressInformation`

GetShipTo returns the ShipTo field if non-nil, zero value otherwise.

### GetShipToOk

`func (o *AssetProductInformationAllOf) GetShipToOk() (*AssetAddressInformation, bool)`

GetShipToOk returns a tuple with the ShipTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTo

`func (o *AssetProductInformationAllOf) SetShipTo(v AssetAddressInformation)`

SetShipTo sets ShipTo field to given value.

### HasShipTo

`func (o *AssetProductInformationAllOf) HasShipTo() bool`

HasShipTo returns a boolean if a field has been set.

### GetSubType

`func (o *AssetProductInformationAllOf) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *AssetProductInformationAllOf) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *AssetProductInformationAllOf) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *AssetProductInformationAllOf) HasSubType() bool`

HasSubType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


