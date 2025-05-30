# CoremanagementCoreFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "coremanagement.CoreFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "coremanagement.CoreFile"]
**AdminState** | Pointer to **string** | Admin state prop to trigger file upload. * &#x60;None&#x60; - Admin configured None State. * &#x60;Upload&#x60; - Admin configured Upload State. | [optional] [default to "None"]
**CoreFileDownloadUrl** | Pointer to **string** | The Url to download the core file. It will be set only after successful completion of core file upload to storage service. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | The device object type for the end point. | [optional] [readonly] 
**EndpointIdentifier** | Pointer to **string** | Endpoint device identifier. In IMM devices, it will be Fabric Interconnect hostname. | [optional] [readonly] 
**FileName** | Pointer to **string** | The name of core file from endpoint device. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | File size of core file in bytes. | [optional] [readonly] 
**NodeId** | Pointer to **string** | Node id within cluster where core file is present. | [optional] [readonly] 
**Pid** | Pointer to **string** | Product identification of the device. | [optional] [readonly] 
**Reason** | Pointer to **string** | Reason for upload failure, if any. In successful upload case, it will be empty. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number of the device. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of core file upload. Valid values are InventoryComplete (default), UploadInProgress, Completed, UploadFailed, FileDownloadUrlCreationFailed, CoreRemovedDownloadOnly. * &#x60;InventoryComplete&#x60; - Default status for all mos before file upload is requested. * &#x60;UploadInProgress&#x60; - File upload is in progress. * &#x60;UploadFailed&#x60; - File upload to storage service failed. * &#x60;Completed&#x60; - File upload to storage service completed successfully. * &#x60;FileDownloadUrlCreationFailed&#x60; - File upload to storage service completed successfully but download url creation failed. * &#x60;CoreRemovedDownloadOnly&#x60; - File upload to storage service completed successfully but file removed from endpoint device. | [optional] [readonly] [default to "InventoryComplete"]
**InventoryParent** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewCoremanagementCoreFile

`func NewCoremanagementCoreFile(classId string, objectType string, ) *CoremanagementCoreFile`

NewCoremanagementCoreFile instantiates a new CoremanagementCoreFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoremanagementCoreFileWithDefaults

`func NewCoremanagementCoreFileWithDefaults() *CoremanagementCoreFile`

NewCoremanagementCoreFileWithDefaults instantiates a new CoremanagementCoreFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CoremanagementCoreFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CoremanagementCoreFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CoremanagementCoreFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CoremanagementCoreFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CoremanagementCoreFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CoremanagementCoreFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *CoremanagementCoreFile) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *CoremanagementCoreFile) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *CoremanagementCoreFile) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *CoremanagementCoreFile) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetCoreFileDownloadUrl

`func (o *CoremanagementCoreFile) GetCoreFileDownloadUrl() string`

GetCoreFileDownloadUrl returns the CoreFileDownloadUrl field if non-nil, zero value otherwise.

### GetCoreFileDownloadUrlOk

`func (o *CoremanagementCoreFile) GetCoreFileDownloadUrlOk() (*string, bool)`

GetCoreFileDownloadUrlOk returns a tuple with the CoreFileDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreFileDownloadUrl

`func (o *CoremanagementCoreFile) SetCoreFileDownloadUrl(v string)`

SetCoreFileDownloadUrl sets CoreFileDownloadUrl field to given value.

### HasCoreFileDownloadUrl

`func (o *CoremanagementCoreFile) HasCoreFileDownloadUrl() bool`

HasCoreFileDownloadUrl returns a boolean if a field has been set.

### GetDeviceType

`func (o *CoremanagementCoreFile) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CoremanagementCoreFile) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CoremanagementCoreFile) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CoremanagementCoreFile) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEndpointIdentifier

`func (o *CoremanagementCoreFile) GetEndpointIdentifier() string`

GetEndpointIdentifier returns the EndpointIdentifier field if non-nil, zero value otherwise.

### GetEndpointIdentifierOk

`func (o *CoremanagementCoreFile) GetEndpointIdentifierOk() (*string, bool)`

GetEndpointIdentifierOk returns a tuple with the EndpointIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointIdentifier

`func (o *CoremanagementCoreFile) SetEndpointIdentifier(v string)`

SetEndpointIdentifier sets EndpointIdentifier field to given value.

### HasEndpointIdentifier

`func (o *CoremanagementCoreFile) HasEndpointIdentifier() bool`

HasEndpointIdentifier returns a boolean if a field has been set.

### GetFileName

`func (o *CoremanagementCoreFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CoremanagementCoreFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CoremanagementCoreFile) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CoremanagementCoreFile) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *CoremanagementCoreFile) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *CoremanagementCoreFile) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *CoremanagementCoreFile) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *CoremanagementCoreFile) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetNodeId

`func (o *CoremanagementCoreFile) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CoremanagementCoreFile) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CoremanagementCoreFile) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *CoremanagementCoreFile) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetPid

`func (o *CoremanagementCoreFile) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CoremanagementCoreFile) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CoremanagementCoreFile) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CoremanagementCoreFile) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetReason

`func (o *CoremanagementCoreFile) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CoremanagementCoreFile) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CoremanagementCoreFile) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CoremanagementCoreFile) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSerial

`func (o *CoremanagementCoreFile) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CoremanagementCoreFile) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CoremanagementCoreFile) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CoremanagementCoreFile) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *CoremanagementCoreFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoremanagementCoreFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoremanagementCoreFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoremanagementCoreFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInventoryParent

`func (o *CoremanagementCoreFile) GetInventoryParent() MoBaseMoRelationship`

GetInventoryParent returns the InventoryParent field if non-nil, zero value otherwise.

### GetInventoryParentOk

`func (o *CoremanagementCoreFile) GetInventoryParentOk() (*MoBaseMoRelationship, bool)`

GetInventoryParentOk returns a tuple with the InventoryParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryParent

`func (o *CoremanagementCoreFile) SetInventoryParent(v MoBaseMoRelationship)`

SetInventoryParent sets InventoryParent field to given value.

### HasInventoryParent

`func (o *CoremanagementCoreFile) HasInventoryParent() bool`

HasInventoryParent returns a boolean if a field has been set.

### SetInventoryParentNil

`func (o *CoremanagementCoreFile) SetInventoryParentNil(b bool)`

 SetInventoryParentNil sets the value for InventoryParent to be an explicit nil

### UnsetInventoryParent
`func (o *CoremanagementCoreFile) UnsetInventoryParent()`

UnsetInventoryParent ensures that no value is present for InventoryParent, not even an explicit nil
### GetRegisteredDevice

`func (o *CoremanagementCoreFile) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CoremanagementCoreFile) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CoremanagementCoreFile) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CoremanagementCoreFile) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *CoremanagementCoreFile) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *CoremanagementCoreFile) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


