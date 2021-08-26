# AssetProductInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ProductInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ProductInformation"]
**BillTo** | Pointer to [**NullableAssetAddressInformation**](AssetAddressInformation.md) |  | [optional] 
**Description** | Pointer to **string** | Short description of the Cisco product that helps identify the product easily. example \&quot;DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\&quot;. | [optional] [readonly] 
**Family** | Pointer to **string** | Family that the product belongs to. Example \&quot;UCSB\&quot;. | [optional] [readonly] 
**Group** | Pointer to **string** | Group that the product belongs to. It is one higher level categorization above family. example \&quot;Switch\&quot;. | [optional] [readonly] 
**Number** | Pointer to **string** | Product number that identifies the product. example PID. example \&quot;UCS-FI-6248UP-CH2\&quot;. | [optional] [readonly] 
**ShipTo** | Pointer to [**NullableAssetAddressInformation**](AssetAddressInformation.md) |  | [optional] 
**SubType** | Pointer to **string** | Sub type of the product being specified. example \&quot;UCS 6200 SER\&quot;. | [optional] [readonly] 

## Methods

### NewAssetProductInformation

`func NewAssetProductInformation(classId string, objectType string, ) *AssetProductInformation`

NewAssetProductInformation instantiates a new AssetProductInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetProductInformationWithDefaults

`func NewAssetProductInformationWithDefaults() *AssetProductInformation`

NewAssetProductInformationWithDefaults instantiates a new AssetProductInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetProductInformation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetProductInformation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetProductInformation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetProductInformation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetProductInformation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetProductInformation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillTo

`func (o *AssetProductInformation) GetBillTo() AssetAddressInformation`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *AssetProductInformation) GetBillToOk() (*AssetAddressInformation, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *AssetProductInformation) SetBillTo(v AssetAddressInformation)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *AssetProductInformation) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### SetBillToNil

`func (o *AssetProductInformation) SetBillToNil(b bool)`

 SetBillToNil sets the value for BillTo to be an explicit nil

### UnsetBillTo
`func (o *AssetProductInformation) UnsetBillTo()`

UnsetBillTo ensures that no value is present for BillTo, not even an explicit nil
### GetDescription

`func (o *AssetProductInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetProductInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetProductInformation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetProductInformation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamily

`func (o *AssetProductInformation) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AssetProductInformation) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AssetProductInformation) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AssetProductInformation) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetGroup

`func (o *AssetProductInformation) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AssetProductInformation) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AssetProductInformation) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AssetProductInformation) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetNumber

`func (o *AssetProductInformation) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AssetProductInformation) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AssetProductInformation) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AssetProductInformation) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipTo

`func (o *AssetProductInformation) GetShipTo() AssetAddressInformation`

GetShipTo returns the ShipTo field if non-nil, zero value otherwise.

### GetShipToOk

`func (o *AssetProductInformation) GetShipToOk() (*AssetAddressInformation, bool)`

GetShipToOk returns a tuple with the ShipTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTo

`func (o *AssetProductInformation) SetShipTo(v AssetAddressInformation)`

SetShipTo sets ShipTo field to given value.

### HasShipTo

`func (o *AssetProductInformation) HasShipTo() bool`

HasShipTo returns a boolean if a field has been set.

### SetShipToNil

`func (o *AssetProductInformation) SetShipToNil(b bool)`

 SetShipToNil sets the value for ShipTo to be an explicit nil

### UnsetShipTo
`func (o *AssetProductInformation) UnsetShipTo()`

UnsetShipTo ensures that no value is present for ShipTo, not even an explicit nil
### GetSubType

`func (o *AssetProductInformation) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *AssetProductInformation) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *AssetProductInformation) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *AssetProductInformation) HasSubType() bool`

HasSubType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


