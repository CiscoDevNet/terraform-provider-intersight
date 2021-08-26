# SoftwarerepositoryOperatingSystemFileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.OperatingSystemFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.OperatingSystemFile"]
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryOperatingSystemFileAllOf

`func NewSoftwarerepositoryOperatingSystemFileAllOf(classId string, objectType string, ) *SoftwarerepositoryOperatingSystemFileAllOf`

NewSoftwarerepositoryOperatingSystemFileAllOf instantiates a new SoftwarerepositoryOperatingSystemFileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryOperatingSystemFileAllOfWithDefaults

`func NewSoftwarerepositoryOperatingSystemFileAllOfWithDefaults() *SoftwarerepositoryOperatingSystemFileAllOf`

NewSoftwarerepositoryOperatingSystemFileAllOfWithDefaults instantiates a new SoftwarerepositoryOperatingSystemFileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVendor

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwarerepositoryOperatingSystemFileAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


