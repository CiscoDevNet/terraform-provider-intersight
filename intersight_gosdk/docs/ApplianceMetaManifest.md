# ApplianceMetaManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.MetaManifest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.MetaManifest"]
**CreationDate** | Pointer to **string** | Record creation date for explicitly tracking all collections for purging the old records. | [optional] [readonly] 
**FileDescription** | Pointer to **string** | Metadata Manifest ImageBundle file description. | [optional] [readonly] 
**FileName** | Pointer to **string** | File names of the Metadata Manifest ImageBundle as reported by the pvapp portal. | [optional] [readonly] 
**FilechkSum** | Pointer to **string** | The md5 checksum of the Metadata Manifest ImageBundle file. | [optional] [readonly] 
**InstallDate** | Pointer to **time.Time** | Install date of the Metadata Manifest ImageBundle. | [optional] [readonly] 
**Status** | Pointer to **string** | Status reported for the successful installation of the meta data files. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceMetaManifest

`func NewApplianceMetaManifest(classId string, objectType string, ) *ApplianceMetaManifest`

NewApplianceMetaManifest instantiates a new ApplianceMetaManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceMetaManifestWithDefaults

`func NewApplianceMetaManifestWithDefaults() *ApplianceMetaManifest`

NewApplianceMetaManifestWithDefaults instantiates a new ApplianceMetaManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceMetaManifest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceMetaManifest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceMetaManifest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceMetaManifest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceMetaManifest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceMetaManifest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreationDate

`func (o *ApplianceMetaManifest) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApplianceMetaManifest) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApplianceMetaManifest) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ApplianceMetaManifest) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetFileDescription

`func (o *ApplianceMetaManifest) GetFileDescription() string`

GetFileDescription returns the FileDescription field if non-nil, zero value otherwise.

### GetFileDescriptionOk

`func (o *ApplianceMetaManifest) GetFileDescriptionOk() (*string, bool)`

GetFileDescriptionOk returns a tuple with the FileDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDescription

`func (o *ApplianceMetaManifest) SetFileDescription(v string)`

SetFileDescription sets FileDescription field to given value.

### HasFileDescription

`func (o *ApplianceMetaManifest) HasFileDescription() bool`

HasFileDescription returns a boolean if a field has been set.

### GetFileName

`func (o *ApplianceMetaManifest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ApplianceMetaManifest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ApplianceMetaManifest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ApplianceMetaManifest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilechkSum

`func (o *ApplianceMetaManifest) GetFilechkSum() string`

GetFilechkSum returns the FilechkSum field if non-nil, zero value otherwise.

### GetFilechkSumOk

`func (o *ApplianceMetaManifest) GetFilechkSumOk() (*string, bool)`

GetFilechkSumOk returns a tuple with the FilechkSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilechkSum

`func (o *ApplianceMetaManifest) SetFilechkSum(v string)`

SetFilechkSum sets FilechkSum field to given value.

### HasFilechkSum

`func (o *ApplianceMetaManifest) HasFilechkSum() bool`

HasFilechkSum returns a boolean if a field has been set.

### GetInstallDate

`func (o *ApplianceMetaManifest) GetInstallDate() time.Time`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *ApplianceMetaManifest) GetInstallDateOk() (*time.Time, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *ApplianceMetaManifest) SetInstallDate(v time.Time)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *ApplianceMetaManifest) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceMetaManifest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceMetaManifest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceMetaManifest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceMetaManifest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceMetaManifest) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceMetaManifest) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceMetaManifest) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceMetaManifest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


