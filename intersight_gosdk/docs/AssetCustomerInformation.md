# AssetCustomerInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**AssetAddressInformation**](asset.AddressInformation.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier for an end customer. This identifier is allocated by Cisco. | [optional] [readonly] 
**Name** | Pointer to **string** | Name as per the information provided by the user. | [optional] [readonly] 

## Methods

### NewAssetCustomerInformation

`func NewAssetCustomerInformation() *AssetCustomerInformation`

NewAssetCustomerInformation instantiates a new AssetCustomerInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetCustomerInformationWithDefaults

`func NewAssetCustomerInformationWithDefaults() *AssetCustomerInformation`

NewAssetCustomerInformationWithDefaults instantiates a new AssetCustomerInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AssetCustomerInformation) GetAddress() AssetAddressInformation`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AssetCustomerInformation) GetAddressOk() (*AssetAddressInformation, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AssetCustomerInformation) SetAddress(v AssetAddressInformation)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AssetCustomerInformation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *AssetCustomerInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetCustomerInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetCustomerInformation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetCustomerInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AssetCustomerInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetCustomerInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetCustomerInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetCustomerInformation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


