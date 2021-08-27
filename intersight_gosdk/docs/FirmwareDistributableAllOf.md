# FirmwareDistributableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.Distributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.Distributable"]
**FileLocation** | Pointer to **string** | The file location of the distributable. | [optional] 
**ImageCategory** | Pointer to **string** | The category into which the distributable falls into according to the supported platform series. For e.g.; C-Series/B-Series/Infrastructure. | [optional] 
**Origin** | Pointer to **string** | The source of the distributable. If it has been created by the user or system. * &#x60;System&#x60; - The distributable has been created by the System. * &#x60;User&#x60; - The distributable has been created by the User. | [optional] [default to "System"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewFirmwareDistributableAllOf

`func NewFirmwareDistributableAllOf(classId string, objectType string, ) *FirmwareDistributableAllOf`

NewFirmwareDistributableAllOf instantiates a new FirmwareDistributableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDistributableAllOfWithDefaults

`func NewFirmwareDistributableAllOfWithDefaults() *FirmwareDistributableAllOf`

NewFirmwareDistributableAllOfWithDefaults instantiates a new FirmwareDistributableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareDistributableAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareDistributableAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareDistributableAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareDistributableAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareDistributableAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareDistributableAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileLocation

`func (o *FirmwareDistributableAllOf) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *FirmwareDistributableAllOf) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *FirmwareDistributableAllOf) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *FirmwareDistributableAllOf) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetImageCategory

`func (o *FirmwareDistributableAllOf) GetImageCategory() string`

GetImageCategory returns the ImageCategory field if non-nil, zero value otherwise.

### GetImageCategoryOk

`func (o *FirmwareDistributableAllOf) GetImageCategoryOk() (*string, bool)`

GetImageCategoryOk returns a tuple with the ImageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCategory

`func (o *FirmwareDistributableAllOf) SetImageCategory(v string)`

SetImageCategory sets ImageCategory field to given value.

### HasImageCategory

`func (o *FirmwareDistributableAllOf) HasImageCategory() bool`

HasImageCategory returns a boolean if a field has been set.

### GetOrigin

`func (o *FirmwareDistributableAllOf) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *FirmwareDistributableAllOf) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *FirmwareDistributableAllOf) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *FirmwareDistributableAllOf) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCatalog

`func (o *FirmwareDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareDistributableAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


