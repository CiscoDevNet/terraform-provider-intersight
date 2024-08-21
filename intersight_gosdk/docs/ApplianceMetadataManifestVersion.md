# ApplianceMetadataManifestVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.MetadataManifestVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.MetadataManifestVersion"]
**FileSha** | Pointer to **string** | Metamanifest file checksum. | [optional] 
**FileTime** | Pointer to **time.Time** | The timestamp when the metamanifest was touched. | [optional] 
**MetadataType** | Pointer to **string** | Name of the bucket that is being monitored. * &#x60;hcl-meta&#x60; - Hcl bucket, metadata update will be automatically enabled. * &#x60;advisories&#x60; - Advisory bucket, metadata update will be automatically enabled. * &#x60;onprem-images&#x60; - Onprem images bucket, metadata update will be automatically enableds. | [optional] [default to "hcl-meta"]

## Methods

### NewApplianceMetadataManifestVersion

`func NewApplianceMetadataManifestVersion(classId string, objectType string, ) *ApplianceMetadataManifestVersion`

NewApplianceMetadataManifestVersion instantiates a new ApplianceMetadataManifestVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceMetadataManifestVersionWithDefaults

`func NewApplianceMetadataManifestVersionWithDefaults() *ApplianceMetadataManifestVersion`

NewApplianceMetadataManifestVersionWithDefaults instantiates a new ApplianceMetadataManifestVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceMetadataManifestVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceMetadataManifestVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceMetadataManifestVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceMetadataManifestVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceMetadataManifestVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceMetadataManifestVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileSha

`func (o *ApplianceMetadataManifestVersion) GetFileSha() string`

GetFileSha returns the FileSha field if non-nil, zero value otherwise.

### GetFileShaOk

`func (o *ApplianceMetadataManifestVersion) GetFileShaOk() (*string, bool)`

GetFileShaOk returns a tuple with the FileSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSha

`func (o *ApplianceMetadataManifestVersion) SetFileSha(v string)`

SetFileSha sets FileSha field to given value.

### HasFileSha

`func (o *ApplianceMetadataManifestVersion) HasFileSha() bool`

HasFileSha returns a boolean if a field has been set.

### GetFileTime

`func (o *ApplianceMetadataManifestVersion) GetFileTime() time.Time`

GetFileTime returns the FileTime field if non-nil, zero value otherwise.

### GetFileTimeOk

`func (o *ApplianceMetadataManifestVersion) GetFileTimeOk() (*time.Time, bool)`

GetFileTimeOk returns a tuple with the FileTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTime

`func (o *ApplianceMetadataManifestVersion) SetFileTime(v time.Time)`

SetFileTime sets FileTime field to given value.

### HasFileTime

`func (o *ApplianceMetadataManifestVersion) HasFileTime() bool`

HasFileTime returns a boolean if a field has been set.

### GetMetadataType

`func (o *ApplianceMetadataManifestVersion) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *ApplianceMetadataManifestVersion) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *ApplianceMetadataManifestVersion) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *ApplianceMetadataManifestVersion) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


