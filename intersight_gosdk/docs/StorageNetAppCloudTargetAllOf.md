# StorageNetAppCloudTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCloudTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCloudTarget"]
**AccessKey** | Pointer to **string** | Access key ID for AWS_S3 and other S3 compatible provider types. | [optional] [readonly] 
**AuthenticationType** | Pointer to **string** | Authentication used to access the target. | [optional] [readonly] 
**AzureAccount** | Pointer to **string** | Azure cloud account name. | [optional] [readonly] 
**CapUrl** | Pointer to **string** | Specifies a full URL of the request to a CAP server for retrieving temporary credentials (access-key, secret-password, and session token) for accessing the object store. | [optional] [readonly] 
**CertificateValidationEnabled** | Pointer to **bool** | Is SSL/TLS certificate validation enabled? The default value is true. This can only be modified for SGWS, IBM_COS, and ONTAP_S3 provider types. | [optional] [readonly] 
**CloudStorage** | Pointer to **[]string** |  | [optional] 
**Container** | Pointer to **string** | Data bucket/container name. | [optional] [readonly] 
**Ipspace** | Pointer to **string** | IPspace to use in order to reach the cloud target. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the cloud target. | [optional] [readonly] 
**Owner** | Pointer to **string** | Owner of the target. Allowed values are FabricPool or SnapMirror. * &#x60;FabricPool&#x60; - NetApp FabricPool Cloud Target owner. * &#x60;SnapMirror&#x60; - NetApp SnapMirror Cloud Target owner. | [optional] [readonly] [default to "FabricPool"]
**Port** | Pointer to **int64** | Port number of the object store that ONTAP uses when establishing a connection. | [optional] [readonly] 
**ProviderType** | Pointer to **string** | Type of cloud provider, e.g., AWS_S3 or ONTAP_S3. * &#x60;ONTAP_S3&#x60; - Cloud target provider type ONTAP_S3. * &#x60;AliCloud&#x60; - Cloud target provider type AliCloud. * &#x60;AWS_S3&#x60; - Cloud target provider type AWS S3. * &#x60;Azure_Cloud&#x60; - Cloud target provider type Azure_Cloud. * &#x60;GoogleCloud&#x60; - Cloud target provider type GoogleCloud. * &#x60;IBM_COS&#x60; - Cloud target provider type IBM_COS. * &#x60;SGWS&#x60; - Cloud target provider type SGWS. | [optional] [readonly] [default to "ONTAP_S3"]
**Server** | Pointer to **string** | Fully qualified domain name of the object store server. | [optional] [readonly] 
**SnapMirrorUse** | Pointer to **string** | Use of the cloud target by SnapMirror. * &#x60;data&#x60; - Data is stored in the SnapMirror target. * &#x60;metadata&#x60; - Metadata is stored in the SnapMirror target. | [optional] [readonly] [default to "data"]
**SslEnabled** | Pointer to **bool** | SSL/HTTPS enabled or not. | [optional] [readonly] 
**Used** | Pointer to **int64** | The amount of cloud space used by all the aggregates attached to the target, in bytes. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of the cloud target. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppCloudTargetAllOf

`func NewStorageNetAppCloudTargetAllOf(classId string, objectType string, ) *StorageNetAppCloudTargetAllOf`

NewStorageNetAppCloudTargetAllOf instantiates a new StorageNetAppCloudTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppCloudTargetAllOfWithDefaults

`func NewStorageNetAppCloudTargetAllOfWithDefaults() *StorageNetAppCloudTargetAllOf`

NewStorageNetAppCloudTargetAllOfWithDefaults instantiates a new StorageNetAppCloudTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCloudTargetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCloudTargetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCloudTargetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCloudTargetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCloudTargetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCloudTargetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessKey

`func (o *StorageNetAppCloudTargetAllOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *StorageNetAppCloudTargetAllOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *StorageNetAppCloudTargetAllOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *StorageNetAppCloudTargetAllOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *StorageNetAppCloudTargetAllOf) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *StorageNetAppCloudTargetAllOf) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *StorageNetAppCloudTargetAllOf) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *StorageNetAppCloudTargetAllOf) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAzureAccount

`func (o *StorageNetAppCloudTargetAllOf) GetAzureAccount() string`

GetAzureAccount returns the AzureAccount field if non-nil, zero value otherwise.

### GetAzureAccountOk

`func (o *StorageNetAppCloudTargetAllOf) GetAzureAccountOk() (*string, bool)`

GetAzureAccountOk returns a tuple with the AzureAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAccount

`func (o *StorageNetAppCloudTargetAllOf) SetAzureAccount(v string)`

SetAzureAccount sets AzureAccount field to given value.

### HasAzureAccount

`func (o *StorageNetAppCloudTargetAllOf) HasAzureAccount() bool`

HasAzureAccount returns a boolean if a field has been set.

### GetCapUrl

`func (o *StorageNetAppCloudTargetAllOf) GetCapUrl() string`

GetCapUrl returns the CapUrl field if non-nil, zero value otherwise.

### GetCapUrlOk

`func (o *StorageNetAppCloudTargetAllOf) GetCapUrlOk() (*string, bool)`

GetCapUrlOk returns a tuple with the CapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapUrl

`func (o *StorageNetAppCloudTargetAllOf) SetCapUrl(v string)`

SetCapUrl sets CapUrl field to given value.

### HasCapUrl

`func (o *StorageNetAppCloudTargetAllOf) HasCapUrl() bool`

HasCapUrl returns a boolean if a field has been set.

### GetCertificateValidationEnabled

`func (o *StorageNetAppCloudTargetAllOf) GetCertificateValidationEnabled() bool`

GetCertificateValidationEnabled returns the CertificateValidationEnabled field if non-nil, zero value otherwise.

### GetCertificateValidationEnabledOk

`func (o *StorageNetAppCloudTargetAllOf) GetCertificateValidationEnabledOk() (*bool, bool)`

GetCertificateValidationEnabledOk returns a tuple with the CertificateValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidationEnabled

`func (o *StorageNetAppCloudTargetAllOf) SetCertificateValidationEnabled(v bool)`

SetCertificateValidationEnabled sets CertificateValidationEnabled field to given value.

### HasCertificateValidationEnabled

`func (o *StorageNetAppCloudTargetAllOf) HasCertificateValidationEnabled() bool`

HasCertificateValidationEnabled returns a boolean if a field has been set.

### GetCloudStorage

`func (o *StorageNetAppCloudTargetAllOf) GetCloudStorage() []string`

GetCloudStorage returns the CloudStorage field if non-nil, zero value otherwise.

### GetCloudStorageOk

`func (o *StorageNetAppCloudTargetAllOf) GetCloudStorageOk() (*[]string, bool)`

GetCloudStorageOk returns a tuple with the CloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorage

`func (o *StorageNetAppCloudTargetAllOf) SetCloudStorage(v []string)`

SetCloudStorage sets CloudStorage field to given value.

### HasCloudStorage

`func (o *StorageNetAppCloudTargetAllOf) HasCloudStorage() bool`

HasCloudStorage returns a boolean if a field has been set.

### SetCloudStorageNil

`func (o *StorageNetAppCloudTargetAllOf) SetCloudStorageNil(b bool)`

 SetCloudStorageNil sets the value for CloudStorage to be an explicit nil

### UnsetCloudStorage
`func (o *StorageNetAppCloudTargetAllOf) UnsetCloudStorage()`

UnsetCloudStorage ensures that no value is present for CloudStorage, not even an explicit nil
### GetContainer

`func (o *StorageNetAppCloudTargetAllOf) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *StorageNetAppCloudTargetAllOf) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *StorageNetAppCloudTargetAllOf) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *StorageNetAppCloudTargetAllOf) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetIpspace

`func (o *StorageNetAppCloudTargetAllOf) GetIpspace() string`

GetIpspace returns the Ipspace field if non-nil, zero value otherwise.

### GetIpspaceOk

`func (o *StorageNetAppCloudTargetAllOf) GetIpspaceOk() (*string, bool)`

GetIpspaceOk returns a tuple with the Ipspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpspace

`func (o *StorageNetAppCloudTargetAllOf) SetIpspace(v string)`

SetIpspace sets Ipspace field to given value.

### HasIpspace

`func (o *StorageNetAppCloudTargetAllOf) HasIpspace() bool`

HasIpspace returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppCloudTargetAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppCloudTargetAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppCloudTargetAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppCloudTargetAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *StorageNetAppCloudTargetAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageNetAppCloudTargetAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageNetAppCloudTargetAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StorageNetAppCloudTargetAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPort

`func (o *StorageNetAppCloudTargetAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageNetAppCloudTargetAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageNetAppCloudTargetAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageNetAppCloudTargetAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProviderType

`func (o *StorageNetAppCloudTargetAllOf) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *StorageNetAppCloudTargetAllOf) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *StorageNetAppCloudTargetAllOf) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *StorageNetAppCloudTargetAllOf) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetServer

`func (o *StorageNetAppCloudTargetAllOf) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *StorageNetAppCloudTargetAllOf) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *StorageNetAppCloudTargetAllOf) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *StorageNetAppCloudTargetAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSnapMirrorUse

`func (o *StorageNetAppCloudTargetAllOf) GetSnapMirrorUse() string`

GetSnapMirrorUse returns the SnapMirrorUse field if non-nil, zero value otherwise.

### GetSnapMirrorUseOk

`func (o *StorageNetAppCloudTargetAllOf) GetSnapMirrorUseOk() (*string, bool)`

GetSnapMirrorUseOk returns a tuple with the SnapMirrorUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapMirrorUse

`func (o *StorageNetAppCloudTargetAllOf) SetSnapMirrorUse(v string)`

SetSnapMirrorUse sets SnapMirrorUse field to given value.

### HasSnapMirrorUse

`func (o *StorageNetAppCloudTargetAllOf) HasSnapMirrorUse() bool`

HasSnapMirrorUse returns a boolean if a field has been set.

### GetSslEnabled

`func (o *StorageNetAppCloudTargetAllOf) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *StorageNetAppCloudTargetAllOf) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *StorageNetAppCloudTargetAllOf) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *StorageNetAppCloudTargetAllOf) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetUsed

`func (o *StorageNetAppCloudTargetAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageNetAppCloudTargetAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageNetAppCloudTargetAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageNetAppCloudTargetAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppCloudTargetAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppCloudTargetAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppCloudTargetAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppCloudTargetAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppCloudTargetAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppCloudTargetAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppCloudTargetAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppCloudTargetAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


