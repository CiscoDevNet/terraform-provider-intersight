# AssetContractInformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillTo** | Pointer to [**AssetAddressInformation**](asset.AddressInformation.md) |  | [optional] 
**BillToGlobalUltimate** | Pointer to [**AssetGlobalUltimate**](asset.GlobalUltimate.md) |  | [optional] 
**ContractNumber** | Pointer to **string** | Contract number for the Cisco support contract purchased for the Cisco device. | [optional] [readonly] 
**LineStatus** | Pointer to **string** | Contract status as per the Cisco Contract APIx. | [optional] [readonly] 

## Methods

### NewAssetContractInformationAllOf

`func NewAssetContractInformationAllOf() *AssetContractInformationAllOf`

NewAssetContractInformationAllOf instantiates a new AssetContractInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetContractInformationAllOfWithDefaults

`func NewAssetContractInformationAllOfWithDefaults() *AssetContractInformationAllOf`

NewAssetContractInformationAllOfWithDefaults instantiates a new AssetContractInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillTo

`func (o *AssetContractInformationAllOf) GetBillTo() AssetAddressInformation`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *AssetContractInformationAllOf) GetBillToOk() (*AssetAddressInformation, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *AssetContractInformationAllOf) SetBillTo(v AssetAddressInformation)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *AssetContractInformationAllOf) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### GetBillToGlobalUltimate

`func (o *AssetContractInformationAllOf) GetBillToGlobalUltimate() AssetGlobalUltimate`

GetBillToGlobalUltimate returns the BillToGlobalUltimate field if non-nil, zero value otherwise.

### GetBillToGlobalUltimateOk

`func (o *AssetContractInformationAllOf) GetBillToGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetBillToGlobalUltimateOk returns a tuple with the BillToGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToGlobalUltimate

`func (o *AssetContractInformationAllOf) SetBillToGlobalUltimate(v AssetGlobalUltimate)`

SetBillToGlobalUltimate sets BillToGlobalUltimate field to given value.

### HasBillToGlobalUltimate

`func (o *AssetContractInformationAllOf) HasBillToGlobalUltimate() bool`

HasBillToGlobalUltimate returns a boolean if a field has been set.

### GetContractNumber

`func (o *AssetContractInformationAllOf) GetContractNumber() string`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *AssetContractInformationAllOf) GetContractNumberOk() (*string, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *AssetContractInformationAllOf) SetContractNumber(v string)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *AssetContractInformationAllOf) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetLineStatus

`func (o *AssetContractInformationAllOf) GetLineStatus() string`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *AssetContractInformationAllOf) GetLineStatusOk() (*string, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *AssetContractInformationAllOf) SetLineStatus(v string)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *AssetContractInformationAllOf) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


