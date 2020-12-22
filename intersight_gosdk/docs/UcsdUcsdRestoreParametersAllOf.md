# UcsdUcsdRestoreParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ucsd.UcsdRestoreParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ucsd.UcsdRestoreParameters"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Location** | Pointer to **string** | The complete location of the path on the server. The location should be specified in the following format- hostname-or-ipv4address&lt;:port&gt;/absolute-file-path. | [optional] 
**Password** | Pointer to **string** | The password of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol. | [optional] 
**Protocol** | Pointer to **string** | The protocol used to backup the UCS Director. | [optional] 
**RestoreConfigurationFiles** | Pointer to **bool** | Decides whether UCS Director property files should also be restored. | [optional] 
**RestoreLicense** | Pointer to **bool** | Decides whether license should also be restored. | [optional] 
**Username** | Pointer to **string** | The username of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol. | [optional] 

## Methods

### NewUcsdUcsdRestoreParametersAllOf

`func NewUcsdUcsdRestoreParametersAllOf(classId string, objectType string, ) *UcsdUcsdRestoreParametersAllOf`

NewUcsdUcsdRestoreParametersAllOf instantiates a new UcsdUcsdRestoreParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcsdUcsdRestoreParametersAllOfWithDefaults

`func NewUcsdUcsdRestoreParametersAllOfWithDefaults() *UcsdUcsdRestoreParametersAllOf`

NewUcsdUcsdRestoreParametersAllOfWithDefaults instantiates a new UcsdUcsdRestoreParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UcsdUcsdRestoreParametersAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UcsdUcsdRestoreParametersAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UcsdUcsdRestoreParametersAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UcsdUcsdRestoreParametersAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *UcsdUcsdRestoreParametersAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *UcsdUcsdRestoreParametersAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *UcsdUcsdRestoreParametersAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetLocation

`func (o *UcsdUcsdRestoreParametersAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UcsdUcsdRestoreParametersAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UcsdUcsdRestoreParametersAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPassword

`func (o *UcsdUcsdRestoreParametersAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UcsdUcsdRestoreParametersAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UcsdUcsdRestoreParametersAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtocol

`func (o *UcsdUcsdRestoreParametersAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UcsdUcsdRestoreParametersAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UcsdUcsdRestoreParametersAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRestoreConfigurationFiles

`func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreConfigurationFiles() bool`

GetRestoreConfigurationFiles returns the RestoreConfigurationFiles field if non-nil, zero value otherwise.

### GetRestoreConfigurationFilesOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreConfigurationFilesOk() (*bool, bool)`

GetRestoreConfigurationFilesOk returns a tuple with the RestoreConfigurationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreConfigurationFiles

`func (o *UcsdUcsdRestoreParametersAllOf) SetRestoreConfigurationFiles(v bool)`

SetRestoreConfigurationFiles sets RestoreConfigurationFiles field to given value.

### HasRestoreConfigurationFiles

`func (o *UcsdUcsdRestoreParametersAllOf) HasRestoreConfigurationFiles() bool`

HasRestoreConfigurationFiles returns a boolean if a field has been set.

### GetRestoreLicense

`func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreLicense() bool`

GetRestoreLicense returns the RestoreLicense field if non-nil, zero value otherwise.

### GetRestoreLicenseOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreLicenseOk() (*bool, bool)`

GetRestoreLicenseOk returns a tuple with the RestoreLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreLicense

`func (o *UcsdUcsdRestoreParametersAllOf) SetRestoreLicense(v bool)`

SetRestoreLicense sets RestoreLicense field to given value.

### HasRestoreLicense

`func (o *UcsdUcsdRestoreParametersAllOf) HasRestoreLicense() bool`

HasRestoreLicense returns a boolean if a field has been set.

### GetUsername

`func (o *UcsdUcsdRestoreParametersAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UcsdUcsdRestoreParametersAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UcsdUcsdRestoreParametersAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UcsdUcsdRestoreParametersAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


