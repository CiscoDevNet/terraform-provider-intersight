# OsSupportedVersionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.SupportedVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.SupportedVersion"]
**VersionName** | Pointer to **string** | The OsInstall Supported Operating System Version Name. | [optional] [readonly] 
**Vendor** | Pointer to [**HclOperatingSystemVendorRelationship**](HclOperatingSystemVendorRelationship.md) |  | [optional] 
**Version** | Pointer to [**HclOperatingSystemRelationship**](HclOperatingSystemRelationship.md) |  | [optional] 

## Methods

### NewOsSupportedVersionAllOf

`func NewOsSupportedVersionAllOf(classId string, objectType string, ) *OsSupportedVersionAllOf`

NewOsSupportedVersionAllOf instantiates a new OsSupportedVersionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsSupportedVersionAllOfWithDefaults

`func NewOsSupportedVersionAllOfWithDefaults() *OsSupportedVersionAllOf`

NewOsSupportedVersionAllOfWithDefaults instantiates a new OsSupportedVersionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsSupportedVersionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsSupportedVersionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsSupportedVersionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsSupportedVersionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsSupportedVersionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsSupportedVersionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersionName

`func (o *OsSupportedVersionAllOf) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *OsSupportedVersionAllOf) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *OsSupportedVersionAllOf) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *OsSupportedVersionAllOf) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetVendor

`func (o *OsSupportedVersionAllOf) GetVendor() HclOperatingSystemVendorRelationship`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OsSupportedVersionAllOf) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OsSupportedVersionAllOf) SetVendor(v HclOperatingSystemVendorRelationship)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OsSupportedVersionAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *OsSupportedVersionAllOf) GetVersion() HclOperatingSystemRelationship`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OsSupportedVersionAllOf) GetVersionOk() (*HclOperatingSystemRelationship, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OsSupportedVersionAllOf) SetVersion(v HclOperatingSystemRelationship)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OsSupportedVersionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


