# FirmwareFirmwareInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Component category. For example, MRAID is under storage controller, CIMC is under management controller. | [optional] 
**Label** | Pointer to **string** | The name of the component to reflect on UI. | [optional] 
**Model** | Pointer to **string** | Model deatils of component. | [optional] 
**UpdateUri** | Pointer to **string** | The redfish URI to get the firmware inventory of server components. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the component. | [optional] 
**Version** | Pointer to **string** | The firmware running version on the component. | [optional] 

## Methods

### NewFirmwareFirmwareInventoryAllOf

`func NewFirmwareFirmwareInventoryAllOf() *FirmwareFirmwareInventoryAllOf`

NewFirmwareFirmwareInventoryAllOf instantiates a new FirmwareFirmwareInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFirmwareInventoryAllOfWithDefaults

`func NewFirmwareFirmwareInventoryAllOfWithDefaults() *FirmwareFirmwareInventoryAllOf`

NewFirmwareFirmwareInventoryAllOfWithDefaults instantiates a new FirmwareFirmwareInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *FirmwareFirmwareInventoryAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FirmwareFirmwareInventoryAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FirmwareFirmwareInventoryAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FirmwareFirmwareInventoryAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLabel

`func (o *FirmwareFirmwareInventoryAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FirmwareFirmwareInventoryAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FirmwareFirmwareInventoryAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FirmwareFirmwareInventoryAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetModel

`func (o *FirmwareFirmwareInventoryAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FirmwareFirmwareInventoryAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FirmwareFirmwareInventoryAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FirmwareFirmwareInventoryAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUpdateUri

`func (o *FirmwareFirmwareInventoryAllOf) GetUpdateUri() string`

GetUpdateUri returns the UpdateUri field if non-nil, zero value otherwise.

### GetUpdateUriOk

`func (o *FirmwareFirmwareInventoryAllOf) GetUpdateUriOk() (*string, bool)`

GetUpdateUriOk returns a tuple with the UpdateUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUri

`func (o *FirmwareFirmwareInventoryAllOf) SetUpdateUri(v string)`

SetUpdateUri sets UpdateUri field to given value.

### HasUpdateUri

`func (o *FirmwareFirmwareInventoryAllOf) HasUpdateUri() bool`

HasUpdateUri returns a boolean if a field has been set.

### GetVendor

`func (o *FirmwareFirmwareInventoryAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareFirmwareInventoryAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareFirmwareInventoryAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FirmwareFirmwareInventoryAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *FirmwareFirmwareInventoryAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FirmwareFirmwareInventoryAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FirmwareFirmwareInventoryAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FirmwareFirmwareInventoryAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


