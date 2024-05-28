# EquipmentEndPointLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.EndPointLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.EndPointLog"]
**CollectionTime** | Pointer to **time.Time** | The time at which the log was collected. | [optional] [readonly] 
**DownloadUrl** | Pointer to **string** | The Url to download the end point log. | [optional] [readonly] 
**FileName** | Pointer to **string** | The end point log file name. | [optional] [readonly] 
**LogType** | Pointer to **string** | The end point log file type. * &#x60;None&#x60; - End point log file type None. * &#x60;SEL&#x60; - End point log file type SEL. | [optional] [readonly] [default to "None"]
**Status** | Pointer to **string** | The end point log collection status. * &#x60;None&#x60; - Log collection not started. * &#x60;CollectionInProgress&#x60; - Log file collection is in progress. * &#x60;CollectionCompleted&#x60; - Log file collection completed. * &#x60;CollectionFailed&#x60; - Log file collection failed. * &#x60;UploadInProgress&#x60; - Log file upload is in progress. * &#x60;UploadCompleted&#x60; - Log file upload completed. * &#x60;UploadFailed&#x60; - Log file upload failed to complete. * &#x60;DownloadUrlCreationFailed&#x60; - Download Url creation failed. * &#x60;Completed&#x60; - Log collection and upload completed. | [optional] [readonly] [default to "None"]
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewEquipmentEndPointLog

`func NewEquipmentEndPointLog(classId string, objectType string, ) *EquipmentEndPointLog`

NewEquipmentEndPointLog instantiates a new EquipmentEndPointLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentEndPointLogWithDefaults

`func NewEquipmentEndPointLogWithDefaults() *EquipmentEndPointLog`

NewEquipmentEndPointLogWithDefaults instantiates a new EquipmentEndPointLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentEndPointLog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentEndPointLog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentEndPointLog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentEndPointLog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentEndPointLog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentEndPointLog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectionTime

`func (o *EquipmentEndPointLog) GetCollectionTime() time.Time`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *EquipmentEndPointLog) GetCollectionTimeOk() (*time.Time, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *EquipmentEndPointLog) SetCollectionTime(v time.Time)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *EquipmentEndPointLog) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *EquipmentEndPointLog) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *EquipmentEndPointLog) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *EquipmentEndPointLog) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *EquipmentEndPointLog) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetFileName

`func (o *EquipmentEndPointLog) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *EquipmentEndPointLog) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *EquipmentEndPointLog) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *EquipmentEndPointLog) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetLogType

`func (o *EquipmentEndPointLog) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *EquipmentEndPointLog) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *EquipmentEndPointLog) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *EquipmentEndPointLog) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetStatus

`func (o *EquipmentEndPointLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EquipmentEndPointLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EquipmentEndPointLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EquipmentEndPointLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentEndPointLog) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentEndPointLog) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentEndPointLog) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentEndPointLog) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentEndPointLog) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentEndPointLog) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetServer

`func (o *EquipmentEndPointLog) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *EquipmentEndPointLog) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *EquipmentEndPointLog) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *EquipmentEndPointLog) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *EquipmentEndPointLog) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *EquipmentEndPointLog) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


