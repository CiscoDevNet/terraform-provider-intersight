# StorageNetAppCifsServiceAllOf

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

### NewStorageNetAppCifsServiceAllOf

`func NewStorageNetAppCifsServiceAllOf(classId string, objectType string, ) *StorageNetAppCifsServiceAllOf`

NewStorageNetAppCifsServiceAllOf instantiates a new StorageNetAppCifsServiceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppCifsServiceAllOfWithDefaults

`func NewStorageNetAppCifsServiceAllOfWithDefaults() *StorageNetAppCifsServiceAllOf`

NewStorageNetAppCifsServiceAllOfWithDefaults instantiates a new StorageNetAppCifsServiceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCifsServiceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCifsServiceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCifsServiceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCifsServiceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCifsServiceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCifsServiceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveDirectoryFqdn

`func (o *StorageNetAppCifsServiceAllOf) GetActiveDirectoryFqdn() string`

GetActiveDirectoryFqdn returns the ActiveDirectoryFqdn field if non-nil, zero value otherwise.

### GetActiveDirectoryFqdnOk

`func (o *StorageNetAppCifsServiceAllOf) GetActiveDirectoryFqdnOk() (*string, bool)`

GetActiveDirectoryFqdnOk returns a tuple with the ActiveDirectoryFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryFqdn

`func (o *StorageNetAppCifsServiceAllOf) SetActiveDirectoryFqdn(v string)`

SetActiveDirectoryFqdn sets ActiveDirectoryFqdn field to given value.

### HasActiveDirectoryFqdn

`func (o *StorageNetAppCifsServiceAllOf) HasActiveDirectoryFqdn() bool`

HasActiveDirectoryFqdn returns a boolean if a field has been set.

### GetAdOrganizationalUnit

`func (o *StorageNetAppCifsServiceAllOf) GetAdOrganizationalUnit() string`

GetAdOrganizationalUnit returns the AdOrganizationalUnit field if non-nil, zero value otherwise.

### GetAdOrganizationalUnitOk

`func (o *StorageNetAppCifsServiceAllOf) GetAdOrganizationalUnitOk() (*string, bool)`

GetAdOrganizationalUnitOk returns a tuple with the AdOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdOrganizationalUnit

`func (o *StorageNetAppCifsServiceAllOf) SetAdOrganizationalUnit(v string)`

SetAdOrganizationalUnit sets AdOrganizationalUnit field to given value.

### HasAdOrganizationalUnit

`func (o *StorageNetAppCifsServiceAllOf) HasAdOrganizationalUnit() bool`

HasAdOrganizationalUnit returns a boolean if a field has been set.

### GetComment

`func (o *StorageNetAppCifsServiceAllOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StorageNetAppCifsServiceAllOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StorageNetAppCifsServiceAllOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StorageNetAppCifsServiceAllOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageNetAppCifsServiceAllOf) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppCifsServiceAllOf) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppCifsServiceAllOf) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppCifsServiceAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerName

`func (o *StorageNetAppCifsServiceAllOf) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *StorageNetAppCifsServiceAllOf) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *StorageNetAppCifsServiceAllOf) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *StorageNetAppCifsServiceAllOf) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppCifsServiceAllOf) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppCifsServiceAllOf) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppCifsServiceAllOf) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppCifsServiceAllOf) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppCifsServiceAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppCifsServiceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppCifsServiceAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppCifsServiceAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


