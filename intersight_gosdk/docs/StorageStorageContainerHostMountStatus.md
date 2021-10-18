# StorageStorageContainerHostMountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.StorageContainerHostMountStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.StorageContainerHostMountStatus"]
**Accessibility** | Pointer to **string** | Host specific storage container accessibility status. * &#x60;NOT_APPLICABLE&#x60; - The HyperFlex storage container accessibility on host is not applicable. * &#x60;ACCESSIBLE&#x60; - The HyperFlex storage container is accessible on the host. * &#x60;NOT_ACCESSIBLE&#x60; - The HyperFlex storage container is not accessible on the host. * &#x60;PARTIALLY_ACCESSIBLE&#x60; - The HyperFlex storage container is partially accessible on the host. * &#x60;READONLY&#x60; - The HyperFlex storage container accessibility is readonly on the host. * &#x60;HOST_NOT_REACHABLE&#x60; - The storage container is not accessible via this host because it is not accessible. * &#x60;UNKNOWN&#x60; - The storage container accessibility via this host is unknown. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**HostName** | Pointer to **string** | The name of the host corresponding to the mount and accessibility status of the storage container. | [optional] [readonly] 
**Mounted** | Pointer to **bool** | Host specific storage container mount status. | [optional] [readonly] 
**Reason** | Pointer to **string** | Host specific storage container mount status reason. | [optional] [readonly] 

## Methods

### NewStorageStorageContainerHostMountStatus

`func NewStorageStorageContainerHostMountStatus(classId string, objectType string, ) *StorageStorageContainerHostMountStatus`

NewStorageStorageContainerHostMountStatus instantiates a new StorageStorageContainerHostMountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStorageContainerHostMountStatusWithDefaults

`func NewStorageStorageContainerHostMountStatusWithDefaults() *StorageStorageContainerHostMountStatus`

NewStorageStorageContainerHostMountStatusWithDefaults instantiates a new StorageStorageContainerHostMountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageStorageContainerHostMountStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageStorageContainerHostMountStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageStorageContainerHostMountStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageStorageContainerHostMountStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageStorageContainerHostMountStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageStorageContainerHostMountStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessibility

`func (o *StorageStorageContainerHostMountStatus) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *StorageStorageContainerHostMountStatus) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *StorageStorageContainerHostMountStatus) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *StorageStorageContainerHostMountStatus) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetHostName

`func (o *StorageStorageContainerHostMountStatus) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *StorageStorageContainerHostMountStatus) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *StorageStorageContainerHostMountStatus) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *StorageStorageContainerHostMountStatus) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetMounted

`func (o *StorageStorageContainerHostMountStatus) GetMounted() bool`

GetMounted returns the Mounted field if non-nil, zero value otherwise.

### GetMountedOk

`func (o *StorageStorageContainerHostMountStatus) GetMountedOk() (*bool, bool)`

GetMountedOk returns a tuple with the Mounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounted

`func (o *StorageStorageContainerHostMountStatus) SetMounted(v bool)`

SetMounted sets Mounted field to given value.

### HasMounted

`func (o *StorageStorageContainerHostMountStatus) HasMounted() bool`

HasMounted returns a boolean if a field has been set.

### GetReason

`func (o *StorageStorageContainerHostMountStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StorageStorageContainerHostMountStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StorageStorageContainerHostMountStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *StorageStorageContainerHostMountStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


