# StorageNetAppCifsService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCifsService"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCifsService"]
**ActiveDirectoryFqdn** | Pointer to **string** | The fully qualified domain name of the Windows Active Directory to which this CIFS server belongs. | [optional] [readonly] 
**AdOrganizationalUnit** | Pointer to **string** | Identifies the organizational unit within the Active Directory domain to associate with the CIFS server. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive text comment for the CIFS server. | [optional] [readonly] 
**Enabled** | Pointer to **string** | Indicates that the CIFS service is administratively enabled. | [optional] [readonly] 
**ServerName** | Pointer to **string** | Name of the NetApp CIFS server. | [optional] [readonly] 
**SvmUuid** | Pointer to **string** | Unique identifier for the NetApp Storage Virtual Machine. | [optional] [readonly] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppCifsService

`func NewStorageNetAppCifsService(classId string, objectType string, ) *StorageNetAppCifsService`

NewStorageNetAppCifsService instantiates a new StorageNetAppCifsService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppCifsServiceWithDefaults

`func NewStorageNetAppCifsServiceWithDefaults() *StorageNetAppCifsService`

NewStorageNetAppCifsServiceWithDefaults instantiates a new StorageNetAppCifsService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCifsService) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCifsService) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCifsService) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCifsService) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCifsService) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCifsService) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveDirectoryFqdn

`func (o *StorageNetAppCifsService) GetActiveDirectoryFqdn() string`

GetActiveDirectoryFqdn returns the ActiveDirectoryFqdn field if non-nil, zero value otherwise.

### GetActiveDirectoryFqdnOk

`func (o *StorageNetAppCifsService) GetActiveDirectoryFqdnOk() (*string, bool)`

GetActiveDirectoryFqdnOk returns a tuple with the ActiveDirectoryFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryFqdn

`func (o *StorageNetAppCifsService) SetActiveDirectoryFqdn(v string)`

SetActiveDirectoryFqdn sets ActiveDirectoryFqdn field to given value.

### HasActiveDirectoryFqdn

`func (o *StorageNetAppCifsService) HasActiveDirectoryFqdn() bool`

HasActiveDirectoryFqdn returns a boolean if a field has been set.

### GetAdOrganizationalUnit

`func (o *StorageNetAppCifsService) GetAdOrganizationalUnit() string`

GetAdOrganizationalUnit returns the AdOrganizationalUnit field if non-nil, zero value otherwise.

### GetAdOrganizationalUnitOk

`func (o *StorageNetAppCifsService) GetAdOrganizationalUnitOk() (*string, bool)`

GetAdOrganizationalUnitOk returns a tuple with the AdOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdOrganizationalUnit

`func (o *StorageNetAppCifsService) SetAdOrganizationalUnit(v string)`

SetAdOrganizationalUnit sets AdOrganizationalUnit field to given value.

### HasAdOrganizationalUnit

`func (o *StorageNetAppCifsService) HasAdOrganizationalUnit() bool`

HasAdOrganizationalUnit returns a boolean if a field has been set.

### GetComment

`func (o *StorageNetAppCifsService) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StorageNetAppCifsService) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StorageNetAppCifsService) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StorageNetAppCifsService) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageNetAppCifsService) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppCifsService) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppCifsService) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppCifsService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerName

`func (o *StorageNetAppCifsService) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *StorageNetAppCifsService) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *StorageNetAppCifsService) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *StorageNetAppCifsService) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppCifsService) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppCifsService) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppCifsService) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppCifsService) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppCifsService) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppCifsService) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppCifsService) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppCifsService) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


