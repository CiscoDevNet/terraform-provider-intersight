# FirmwareDriverDistributableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The device type on which the driver is installable. | [optional] 
**Directory** | Pointer to **string** | Indicates in which directory path this driver will be added. | [optional] 
**Osname** | Pointer to **string** | The operating system name to which this driver is compatible. | [optional] 
**Osversion** | Pointer to **string** | OS Version. It is populated as part of the image import operation. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareDriverDistributableAllOf

`func NewFirmwareDriverDistributableAllOf() *FirmwareDriverDistributableAllOf`

NewFirmwareDriverDistributableAllOf instantiates a new FirmwareDriverDistributableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDriverDistributableAllOfWithDefaults

`func NewFirmwareDriverDistributableAllOfWithDefaults() *FirmwareDriverDistributableAllOf`

NewFirmwareDriverDistributableAllOfWithDefaults instantiates a new FirmwareDriverDistributableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *FirmwareDriverDistributableAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FirmwareDriverDistributableAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FirmwareDriverDistributableAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FirmwareDriverDistributableAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDirectory

`func (o *FirmwareDriverDistributableAllOf) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *FirmwareDriverDistributableAllOf) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *FirmwareDriverDistributableAllOf) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *FirmwareDriverDistributableAllOf) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetOsname

`func (o *FirmwareDriverDistributableAllOf) GetOsname() string`

GetOsname returns the Osname field if non-nil, zero value otherwise.

### GetOsnameOk

`func (o *FirmwareDriverDistributableAllOf) GetOsnameOk() (*string, bool)`

GetOsnameOk returns a tuple with the Osname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsname

`func (o *FirmwareDriverDistributableAllOf) SetOsname(v string)`

SetOsname sets Osname field to given value.

### HasOsname

`func (o *FirmwareDriverDistributableAllOf) HasOsname() bool`

HasOsname returns a boolean if a field has been set.

### GetOsversion

`func (o *FirmwareDriverDistributableAllOf) GetOsversion() string`

GetOsversion returns the Osversion field if non-nil, zero value otherwise.

### GetOsversionOk

`func (o *FirmwareDriverDistributableAllOf) GetOsversionOk() (*string, bool)`

GetOsversionOk returns a tuple with the Osversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsversion

`func (o *FirmwareDriverDistributableAllOf) SetOsversion(v string)`

SetOsversion sets Osversion field to given value.

### HasOsversion

`func (o *FirmwareDriverDistributableAllOf) HasOsversion() bool`

HasOsversion returns a boolean if a field has been set.

### GetCatalog

`func (o *FirmwareDriverDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareDriverDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareDriverDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareDriverDistributableAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


