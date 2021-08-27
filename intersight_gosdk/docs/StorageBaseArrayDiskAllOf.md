# StorageBaseArrayDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | Disk name available in storage array. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Storage disk part number. | [optional] [readonly] 
**Protocol** | Pointer to **string** | Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe. * &#x60;Unknown&#x60; - Disk protocol is unknown. * &#x60;SAS&#x60; - Serial Attached SCSI protocol (SAS) used in disk. * &#x60;NVMe&#x60; - Non-volatile memory express (NVMe) protocol used in disk. * &#x60;SATA&#x60; - Serial Advanced Technology Attachment (SATA) used in disk. | [optional] [readonly] [default to "Unknown"]
**Speed** | Pointer to **int64** | Disk speed for read or write operation measured in rpm. | [optional] [readonly] 
**Status** | Pointer to **string** | Storage disk health status. * &#x60;Unknown&#x60; - Component status is not available. * &#x60;Ok&#x60; - Component is healthy and no issues found. * &#x60;Degraded&#x60; - Functioning, but not at full capability due to a non-fatal failure. * &#x60;Critical&#x60; - Not functioning or requiring immediate attention. * &#x60;Offline&#x60; - Component is installed, but powered off. * &#x60;Identifying&#x60; - Component is in initialization process. * &#x60;NotAvailable&#x60; - Component is not installed or not available. * &#x60;Updating&#x60; - Software update is in progress. * &#x60;Unrecognized&#x60; - Component is not recognized. It may be because the component is not installed properly or it is not supported. | [optional] [readonly] [default to "Unknown"]
**StorageUtilization** | Pointer to [**NullableStorageBaseCapacity**](StorageBaseCapacity.md) |  | [optional] 
**Type** | Pointer to **string** | Storage disk type - it can be SSD, HDD, NVRAM. * &#x60;Unknown&#x60; - Default unknown disk type. * &#x60;SSD&#x60; - Storage disk with Solid state disk. * &#x60;HDD&#x60; - Storage disk with Hard disk drive. * &#x60;NVRAM&#x60; - Storage disk with Non-volatile random-access memory type. * &#x60;SATA&#x60; - Disk drive implementation with Serial Advanced Technology Attachment (SATA). * &#x60;BSAS&#x60; - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * &#x60;FC&#x60; - Storage disk with Fiber Channel. * &#x60;FSAS&#x60; - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives. * &#x60;LUN&#x60; - Logical Unit Number refers to a logical disk. * &#x60;MSATA&#x60; - SATA disk in multi-disk carrier storage shelf. * &#x60;SAS&#x60; - Storage disk with serial attached SCSI. * &#x60;VMDISK&#x60; - Virtual machine Data Disk. | [optional] [readonly] [default to "Unknown"]
**Version** | Pointer to **string** | Storage disk version number. | [optional] [readonly] 

## Methods

### NewStorageBaseArrayDiskAllOf

`func NewStorageBaseArrayDiskAllOf(classId string, objectType string, ) *StorageBaseArrayDiskAllOf`

NewStorageBaseArrayDiskAllOf instantiates a new StorageBaseArrayDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseArrayDiskAllOfWithDefaults

`func NewStorageBaseArrayDiskAllOfWithDefaults() *StorageBaseArrayDiskAllOf`

NewStorageBaseArrayDiskAllOfWithDefaults instantiates a new StorageBaseArrayDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseArrayDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseArrayDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseArrayDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseArrayDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseArrayDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseArrayDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageBaseArrayDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseArrayDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseArrayDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseArrayDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartNumber

`func (o *StorageBaseArrayDiskAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *StorageBaseArrayDiskAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *StorageBaseArrayDiskAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *StorageBaseArrayDiskAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageBaseArrayDiskAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageBaseArrayDiskAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageBaseArrayDiskAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageBaseArrayDiskAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpeed

`func (o *StorageBaseArrayDiskAllOf) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageBaseArrayDiskAllOf) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageBaseArrayDiskAllOf) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageBaseArrayDiskAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBaseArrayDiskAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBaseArrayDiskAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBaseArrayDiskAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBaseArrayDiskAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseArrayDiskAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseArrayDiskAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseArrayDiskAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseArrayDiskAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StorageBaseArrayDiskAllOf) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StorageBaseArrayDiskAllOf) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
### GetType

`func (o *StorageBaseArrayDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageBaseArrayDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageBaseArrayDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageBaseArrayDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *StorageBaseArrayDiskAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageBaseArrayDiskAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageBaseArrayDiskAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageBaseArrayDiskAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


