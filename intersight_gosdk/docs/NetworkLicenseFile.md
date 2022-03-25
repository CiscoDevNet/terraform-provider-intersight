# NetworkLicenseFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.LicenseFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.LicenseFile"]
**ExpiryDate** | Pointer to **string** | The expiry date of the license feature. | [optional] 
**FeatureName** | Pointer to **string** | The name of the license feature. | [optional] 
**FileId** | Pointer to **string** | The file Id of the license file. | [optional] 
**HostId** | Pointer to **string** | The identifier to identify license host Id. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the license. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkLicenseFile

`func NewNetworkLicenseFile(classId string, objectType string, ) *NetworkLicenseFile`

NewNetworkLicenseFile instantiates a new NetworkLicenseFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLicenseFileWithDefaults

`func NewNetworkLicenseFileWithDefaults() *NetworkLicenseFile`

NewNetworkLicenseFileWithDefaults instantiates a new NetworkLicenseFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkLicenseFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkLicenseFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkLicenseFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkLicenseFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkLicenseFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkLicenseFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpiryDate

`func (o *NetworkLicenseFile) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NetworkLicenseFile) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NetworkLicenseFile) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NetworkLicenseFile) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetFeatureName

`func (o *NetworkLicenseFile) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *NetworkLicenseFile) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *NetworkLicenseFile) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *NetworkLicenseFile) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetFileId

`func (o *NetworkLicenseFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *NetworkLicenseFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *NetworkLicenseFile) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *NetworkLicenseFile) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetHostId

`func (o *NetworkLicenseFile) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *NetworkLicenseFile) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *NetworkLicenseFile) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *NetworkLicenseFile) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetVendor

`func (o *NetworkLicenseFile) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *NetworkLicenseFile) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *NetworkLicenseFile) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *NetworkLicenseFile) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkLicenseFile) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkLicenseFile) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkLicenseFile) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkLicenseFile) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkLicenseFile) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkLicenseFile) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkLicenseFile) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkLicenseFile) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


