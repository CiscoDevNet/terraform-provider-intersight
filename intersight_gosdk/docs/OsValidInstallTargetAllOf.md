# OsValidInstallTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.ValidInstallTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.ValidInstallTarget"]
**Error** | Pointer to **string** | Error message if any errors are encountered while fetching and validating Install targets for the server. | [optional] [readonly] 
**M2Jbod** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**M2VirtualDrives** | Pointer to [**[]OsVirtualDriveResponse**](OsVirtualDriveResponse.md) |  | [optional] 
**MraidJbod** | Pointer to [**[]OsPhysicalDiskResponse**](OsPhysicalDiskResponse.md) |  | [optional] 
**MraidVirtualDrives** | Pointer to [**[]OsVirtualDriveResponse**](OsVirtualDriveResponse.md) |  | [optional] 
**Servers** | Pointer to [**[]ComputePhysicalRelationship**](ComputePhysicalRelationship.md) | An array of relationships to computePhysical resources. | [optional] 

## Methods

### NewOsValidInstallTargetAllOf

`func NewOsValidInstallTargetAllOf(classId string, objectType string, ) *OsValidInstallTargetAllOf`

NewOsValidInstallTargetAllOf instantiates a new OsValidInstallTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsValidInstallTargetAllOfWithDefaults

`func NewOsValidInstallTargetAllOfWithDefaults() *OsValidInstallTargetAllOf`

NewOsValidInstallTargetAllOfWithDefaults instantiates a new OsValidInstallTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsValidInstallTargetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsValidInstallTargetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsValidInstallTargetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsValidInstallTargetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsValidInstallTargetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsValidInstallTargetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetError

`func (o *OsValidInstallTargetAllOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OsValidInstallTargetAllOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OsValidInstallTargetAllOf) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OsValidInstallTargetAllOf) HasError() bool`

HasError returns a boolean if a field has been set.

### GetM2Jbod

`func (o *OsValidInstallTargetAllOf) GetM2Jbod() []OsPhysicalDiskResponse`

GetM2Jbod returns the M2Jbod field if non-nil, zero value otherwise.

### GetM2JbodOk

`func (o *OsValidInstallTargetAllOf) GetM2JbodOk() (*[]OsPhysicalDiskResponse, bool)`

GetM2JbodOk returns a tuple with the M2Jbod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2Jbod

`func (o *OsValidInstallTargetAllOf) SetM2Jbod(v []OsPhysicalDiskResponse)`

SetM2Jbod sets M2Jbod field to given value.

### HasM2Jbod

`func (o *OsValidInstallTargetAllOf) HasM2Jbod() bool`

HasM2Jbod returns a boolean if a field has been set.

### SetM2JbodNil

`func (o *OsValidInstallTargetAllOf) SetM2JbodNil(b bool)`

 SetM2JbodNil sets the value for M2Jbod to be an explicit nil

### UnsetM2Jbod
`func (o *OsValidInstallTargetAllOf) UnsetM2Jbod()`

UnsetM2Jbod ensures that no value is present for M2Jbod, not even an explicit nil
### GetM2VirtualDrives

`func (o *OsValidInstallTargetAllOf) GetM2VirtualDrives() []OsVirtualDriveResponse`

GetM2VirtualDrives returns the M2VirtualDrives field if non-nil, zero value otherwise.

### GetM2VirtualDrivesOk

`func (o *OsValidInstallTargetAllOf) GetM2VirtualDrivesOk() (*[]OsVirtualDriveResponse, bool)`

GetM2VirtualDrivesOk returns a tuple with the M2VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2VirtualDrives

`func (o *OsValidInstallTargetAllOf) SetM2VirtualDrives(v []OsVirtualDriveResponse)`

SetM2VirtualDrives sets M2VirtualDrives field to given value.

### HasM2VirtualDrives

`func (o *OsValidInstallTargetAllOf) HasM2VirtualDrives() bool`

HasM2VirtualDrives returns a boolean if a field has been set.

### SetM2VirtualDrivesNil

`func (o *OsValidInstallTargetAllOf) SetM2VirtualDrivesNil(b bool)`

 SetM2VirtualDrivesNil sets the value for M2VirtualDrives to be an explicit nil

### UnsetM2VirtualDrives
`func (o *OsValidInstallTargetAllOf) UnsetM2VirtualDrives()`

UnsetM2VirtualDrives ensures that no value is present for M2VirtualDrives, not even an explicit nil
### GetMraidJbod

`func (o *OsValidInstallTargetAllOf) GetMraidJbod() []OsPhysicalDiskResponse`

GetMraidJbod returns the MraidJbod field if non-nil, zero value otherwise.

### GetMraidJbodOk

`func (o *OsValidInstallTargetAllOf) GetMraidJbodOk() (*[]OsPhysicalDiskResponse, bool)`

GetMraidJbodOk returns a tuple with the MraidJbod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMraidJbod

`func (o *OsValidInstallTargetAllOf) SetMraidJbod(v []OsPhysicalDiskResponse)`

SetMraidJbod sets MraidJbod field to given value.

### HasMraidJbod

`func (o *OsValidInstallTargetAllOf) HasMraidJbod() bool`

HasMraidJbod returns a boolean if a field has been set.

### SetMraidJbodNil

`func (o *OsValidInstallTargetAllOf) SetMraidJbodNil(b bool)`

 SetMraidJbodNil sets the value for MraidJbod to be an explicit nil

### UnsetMraidJbod
`func (o *OsValidInstallTargetAllOf) UnsetMraidJbod()`

UnsetMraidJbod ensures that no value is present for MraidJbod, not even an explicit nil
### GetMraidVirtualDrives

`func (o *OsValidInstallTargetAllOf) GetMraidVirtualDrives() []OsVirtualDriveResponse`

GetMraidVirtualDrives returns the MraidVirtualDrives field if non-nil, zero value otherwise.

### GetMraidVirtualDrivesOk

`func (o *OsValidInstallTargetAllOf) GetMraidVirtualDrivesOk() (*[]OsVirtualDriveResponse, bool)`

GetMraidVirtualDrivesOk returns a tuple with the MraidVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMraidVirtualDrives

`func (o *OsValidInstallTargetAllOf) SetMraidVirtualDrives(v []OsVirtualDriveResponse)`

SetMraidVirtualDrives sets MraidVirtualDrives field to given value.

### HasMraidVirtualDrives

`func (o *OsValidInstallTargetAllOf) HasMraidVirtualDrives() bool`

HasMraidVirtualDrives returns a boolean if a field has been set.

### SetMraidVirtualDrivesNil

`func (o *OsValidInstallTargetAllOf) SetMraidVirtualDrivesNil(b bool)`

 SetMraidVirtualDrivesNil sets the value for MraidVirtualDrives to be an explicit nil

### UnsetMraidVirtualDrives
`func (o *OsValidInstallTargetAllOf) UnsetMraidVirtualDrives()`

UnsetMraidVirtualDrives ensures that no value is present for MraidVirtualDrives, not even an explicit nil
### GetServers

`func (o *OsValidInstallTargetAllOf) GetServers() []ComputePhysicalRelationship`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OsValidInstallTargetAllOf) GetServersOk() (*[]ComputePhysicalRelationship, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OsValidInstallTargetAllOf) SetServers(v []ComputePhysicalRelationship)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OsValidInstallTargetAllOf) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *OsValidInstallTargetAllOf) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *OsValidInstallTargetAllOf) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


