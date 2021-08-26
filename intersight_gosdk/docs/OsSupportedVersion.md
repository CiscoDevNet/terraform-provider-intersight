# OsSupportedVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.SupportedVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.SupportedVersion"]
**VersionName** | Pointer to **string** | The OsInstall Supported Operating System Version Name. | [optional] [readonly] 
**Vendor** | Pointer to [**HclOperatingSystemVendorRelationship**](HclOperatingSystemVendorRelationship.md) |  | [optional] 
**Version** | Pointer to [**HclOperatingSystemRelationship**](HclOperatingSystemRelationship.md) |  | [optional] 

## Methods

### NewOsSupportedVersion

`func NewOsSupportedVersion(classId string, objectType string, ) *OsSupportedVersion`

NewOsSupportedVersion instantiates a new OsSupportedVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsSupportedVersionWithDefaults

`func NewOsSupportedVersionWithDefaults() *OsSupportedVersion`

NewOsSupportedVersionWithDefaults instantiates a new OsSupportedVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsSupportedVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsSupportedVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsSupportedVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsSupportedVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsSupportedVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsSupportedVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersionName

`func (o *OsSupportedVersion) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *OsSupportedVersion) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *OsSupportedVersion) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *OsSupportedVersion) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetVendor

`func (o *OsSupportedVersion) GetVendor() HclOperatingSystemVendorRelationship`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OsSupportedVersion) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OsSupportedVersion) SetVendor(v HclOperatingSystemVendorRelationship)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OsSupportedVersion) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *OsSupportedVersion) GetVersion() HclOperatingSystemRelationship`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OsSupportedVersion) GetVersionOk() (*HclOperatingSystemRelationship, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OsSupportedVersion) SetVersion(v HclOperatingSystemRelationship)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OsSupportedVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


