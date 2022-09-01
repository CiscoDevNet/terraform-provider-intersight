# StorageNetAppLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppLicense"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppLicense"]
**ClusterUuid** | Pointer to **string** | Unique identity of the device. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the license package. | [optional] [readonly] 
**State** | Pointer to **string** | Summary state of license package based on all installed licenses. * &#x60;Unknown&#x60; - The summary state of the license package is unknown. * &#x60;Compliant&#x60; - The summary state of the license package is compliant. * &#x60;Noncompliant&#x60; - The summary state of the license package is noncompliant. * &#x60;Unlicensed&#x60; - The summary state of the license package is unlicensed. | [optional] [readonly] [default to "Unknown"]
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppLicense

`func NewStorageNetAppLicense(classId string, objectType string, ) *StorageNetAppLicense`

NewStorageNetAppLicense instantiates a new StorageNetAppLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppLicenseWithDefaults

`func NewStorageNetAppLicenseWithDefaults() *StorageNetAppLicense`

NewStorageNetAppLicenseWithDefaults instantiates a new StorageNetAppLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppLicense) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppLicense) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppLicense) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppLicense) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppLicense) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppLicense) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterUuid

`func (o *StorageNetAppLicense) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *StorageNetAppLicense) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *StorageNetAppLicense) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *StorageNetAppLicense) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppLicense) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppLicense) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppLicense) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppLicense) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppLicense) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppLicense) HasState() bool`

HasState returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppLicense) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppLicense) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppLicense) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppLicense) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


