# ServerMigrationKeyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.MigrationKeyDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.MigrationKeyDetails"]
**MigrationKeyInfo** | Pointer to [**NullableComputeMigrationKeyInfo**](ComputeMigrationKeyInfo.md) |  | [optional] 
**DownloadStatus** | Pointer to [**NullableComputeDownloadStatusRelationship**](ComputeDownloadStatusRelationship.md) |  | [optional] 
**Profile** | Pointer to [**NullableServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewServerMigrationKeyDetails

`func NewServerMigrationKeyDetails(classId string, objectType string, ) *ServerMigrationKeyDetails`

NewServerMigrationKeyDetails instantiates a new ServerMigrationKeyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMigrationKeyDetailsWithDefaults

`func NewServerMigrationKeyDetailsWithDefaults() *ServerMigrationKeyDetails`

NewServerMigrationKeyDetailsWithDefaults instantiates a new ServerMigrationKeyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerMigrationKeyDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerMigrationKeyDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerMigrationKeyDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerMigrationKeyDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerMigrationKeyDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerMigrationKeyDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMigrationKeyInfo

`func (o *ServerMigrationKeyDetails) GetMigrationKeyInfo() ComputeMigrationKeyInfo`

GetMigrationKeyInfo returns the MigrationKeyInfo field if non-nil, zero value otherwise.

### GetMigrationKeyInfoOk

`func (o *ServerMigrationKeyDetails) GetMigrationKeyInfoOk() (*ComputeMigrationKeyInfo, bool)`

GetMigrationKeyInfoOk returns a tuple with the MigrationKeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationKeyInfo

`func (o *ServerMigrationKeyDetails) SetMigrationKeyInfo(v ComputeMigrationKeyInfo)`

SetMigrationKeyInfo sets MigrationKeyInfo field to given value.

### HasMigrationKeyInfo

`func (o *ServerMigrationKeyDetails) HasMigrationKeyInfo() bool`

HasMigrationKeyInfo returns a boolean if a field has been set.

### SetMigrationKeyInfoNil

`func (o *ServerMigrationKeyDetails) SetMigrationKeyInfoNil(b bool)`

 SetMigrationKeyInfoNil sets the value for MigrationKeyInfo to be an explicit nil

### UnsetMigrationKeyInfo
`func (o *ServerMigrationKeyDetails) UnsetMigrationKeyInfo()`

UnsetMigrationKeyInfo ensures that no value is present for MigrationKeyInfo, not even an explicit nil
### GetDownloadStatus

`func (o *ServerMigrationKeyDetails) GetDownloadStatus() ComputeDownloadStatusRelationship`

GetDownloadStatus returns the DownloadStatus field if non-nil, zero value otherwise.

### GetDownloadStatusOk

`func (o *ServerMigrationKeyDetails) GetDownloadStatusOk() (*ComputeDownloadStatusRelationship, bool)`

GetDownloadStatusOk returns a tuple with the DownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStatus

`func (o *ServerMigrationKeyDetails) SetDownloadStatus(v ComputeDownloadStatusRelationship)`

SetDownloadStatus sets DownloadStatus field to given value.

### HasDownloadStatus

`func (o *ServerMigrationKeyDetails) HasDownloadStatus() bool`

HasDownloadStatus returns a boolean if a field has been set.

### SetDownloadStatusNil

`func (o *ServerMigrationKeyDetails) SetDownloadStatusNil(b bool)`

 SetDownloadStatusNil sets the value for DownloadStatus to be an explicit nil

### UnsetDownloadStatus
`func (o *ServerMigrationKeyDetails) UnsetDownloadStatus()`

UnsetDownloadStatus ensures that no value is present for DownloadStatus, not even an explicit nil
### GetProfile

`func (o *ServerMigrationKeyDetails) GetProfile() ServerProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerMigrationKeyDetails) GetProfileOk() (*ServerProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerMigrationKeyDetails) SetProfile(v ServerProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServerMigrationKeyDetails) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *ServerMigrationKeyDetails) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ServerMigrationKeyDetails) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetServer

`func (o *ServerMigrationKeyDetails) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerMigrationKeyDetails) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerMigrationKeyDetails) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerMigrationKeyDetails) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ServerMigrationKeyDetails) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ServerMigrationKeyDetails) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


