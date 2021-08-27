# FirmwareDriverDistributable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.DriverDistributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.DriverDistributable"]
**Category** | Pointer to **string** | The device type on which the driver is installable. | [optional] 
**Directory** | Pointer to **string** | Indicates in which directory path this driver will be added. | [optional] 
**Osname** | Pointer to **string** | The operating system name to which this driver is compatible. | [optional] 
**Osversion** | Pointer to **string** | OS Version. It is populated as part of the image import operation. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewFirmwareDriverDistributable

`func NewFirmwareDriverDistributable(classId string, objectType string, ) *FirmwareDriverDistributable`

NewFirmwareDriverDistributable instantiates a new FirmwareDriverDistributable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDriverDistributableWithDefaults

`func NewFirmwareDriverDistributableWithDefaults() *FirmwareDriverDistributable`

NewFirmwareDriverDistributableWithDefaults instantiates a new FirmwareDriverDistributable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareDriverDistributable) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareDriverDistributable) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareDriverDistributable) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareDriverDistributable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareDriverDistributable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareDriverDistributable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *FirmwareDriverDistributable) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FirmwareDriverDistributable) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FirmwareDriverDistributable) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FirmwareDriverDistributable) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDirectory

`func (o *FirmwareDriverDistributable) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *FirmwareDriverDistributable) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *FirmwareDriverDistributable) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *FirmwareDriverDistributable) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetOsname

`func (o *FirmwareDriverDistributable) GetOsname() string`

GetOsname returns the Osname field if non-nil, zero value otherwise.

### GetOsnameOk

`func (o *FirmwareDriverDistributable) GetOsnameOk() (*string, bool)`

GetOsnameOk returns a tuple with the Osname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsname

`func (o *FirmwareDriverDistributable) SetOsname(v string)`

SetOsname sets Osname field to given value.

### HasOsname

`func (o *FirmwareDriverDistributable) HasOsname() bool`

HasOsname returns a boolean if a field has been set.

### GetOsversion

`func (o *FirmwareDriverDistributable) GetOsversion() string`

GetOsversion returns the Osversion field if non-nil, zero value otherwise.

### GetOsversionOk

`func (o *FirmwareDriverDistributable) GetOsversionOk() (*string, bool)`

GetOsversionOk returns a tuple with the Osversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsversion

`func (o *FirmwareDriverDistributable) SetOsversion(v string)`

SetOsversion sets Osversion field to given value.

### HasOsversion

`func (o *FirmwareDriverDistributable) HasOsversion() bool`

HasOsversion returns a boolean if a field has been set.

### GetCatalog

`func (o *FirmwareDriverDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareDriverDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareDriverDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareDriverDistributable) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


