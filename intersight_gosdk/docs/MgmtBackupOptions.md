# MgmtBackupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.BackupOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.BackupOptions"]
**AutoRollover** | Pointer to **bool** | If enabled and the limit on the number of backups per account has been reached, the system would attempt to clear space by removing the oldest backup instance as part of the backup operation. If disabled and the limit has been reached, the request to create a backup operation would fail. | [optional] 
**PreserveIdentities** | Pointer to **bool** | Backup all identities derived from pools. | [optional] 

## Methods

### NewMgmtBackupOptions

`func NewMgmtBackupOptions(classId string, objectType string, ) *MgmtBackupOptions`

NewMgmtBackupOptions instantiates a new MgmtBackupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtBackupOptionsWithDefaults

`func NewMgmtBackupOptionsWithDefaults() *MgmtBackupOptions`

NewMgmtBackupOptionsWithDefaults instantiates a new MgmtBackupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtBackupOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtBackupOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtBackupOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtBackupOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtBackupOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtBackupOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoRollover

`func (o *MgmtBackupOptions) GetAutoRollover() bool`

GetAutoRollover returns the AutoRollover field if non-nil, zero value otherwise.

### GetAutoRolloverOk

`func (o *MgmtBackupOptions) GetAutoRolloverOk() (*bool, bool)`

GetAutoRolloverOk returns a tuple with the AutoRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRollover

`func (o *MgmtBackupOptions) SetAutoRollover(v bool)`

SetAutoRollover sets AutoRollover field to given value.

### HasAutoRollover

`func (o *MgmtBackupOptions) HasAutoRollover() bool`

HasAutoRollover returns a boolean if a field has been set.

### GetPreserveIdentities

`func (o *MgmtBackupOptions) GetPreserveIdentities() bool`

GetPreserveIdentities returns the PreserveIdentities field if non-nil, zero value otherwise.

### GetPreserveIdentitiesOk

`func (o *MgmtBackupOptions) GetPreserveIdentitiesOk() (*bool, bool)`

GetPreserveIdentitiesOk returns a tuple with the PreserveIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveIdentities

`func (o *MgmtBackupOptions) SetPreserveIdentities(v bool)`

SetPreserveIdentities sets PreserveIdentities field to given value.

### HasPreserveIdentities

`func (o *MgmtBackupOptions) HasPreserveIdentities() bool`

HasPreserveIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


