# ApplianceRemoteFileImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.RemoteFileImport"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.RemoteFileImport"]
**Filename** | Pointer to **string** | The name of the file to be imported. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the machine where the file is located. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Password for remote requiest. | [optional] 
**Path** | Pointer to **string** | The port that should be used for the remote request. | [optional] 
**Port** | Pointer to **int64** | The port that should be used for the remote request. | [optional] 
**Protocol** | Pointer to **string** | Protocol for the remote request. * &#x60;scp&#x60; - Secure Copy Protocol (SCP) to access the file server. * &#x60;sftp&#x60; - SSH File Transfer Protocol (SFTP) to access file server. | [optional] [default to "scp"]
**Username** | Pointer to **string** | The username for the remote request. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceRemoteFileImport

`func NewApplianceRemoteFileImport(classId string, objectType string, ) *ApplianceRemoteFileImport`

NewApplianceRemoteFileImport instantiates a new ApplianceRemoteFileImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceRemoteFileImportWithDefaults

`func NewApplianceRemoteFileImportWithDefaults() *ApplianceRemoteFileImport`

NewApplianceRemoteFileImportWithDefaults instantiates a new ApplianceRemoteFileImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceRemoteFileImport) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceRemoteFileImport) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceRemoteFileImport) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceRemoteFileImport) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceRemoteFileImport) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceRemoteFileImport) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilename

`func (o *ApplianceRemoteFileImport) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApplianceRemoteFileImport) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApplianceRemoteFileImport) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApplianceRemoteFileImport) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceRemoteFileImport) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceRemoteFileImport) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceRemoteFileImport) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceRemoteFileImport) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *ApplianceRemoteFileImport) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ApplianceRemoteFileImport) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ApplianceRemoteFileImport) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ApplianceRemoteFileImport) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *ApplianceRemoteFileImport) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApplianceRemoteFileImport) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApplianceRemoteFileImport) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApplianceRemoteFileImport) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *ApplianceRemoteFileImport) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApplianceRemoteFileImport) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApplianceRemoteFileImport) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApplianceRemoteFileImport) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *ApplianceRemoteFileImport) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApplianceRemoteFileImport) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApplianceRemoteFileImport) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApplianceRemoteFileImport) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplianceRemoteFileImport) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplianceRemoteFileImport) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplianceRemoteFileImport) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplianceRemoteFileImport) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *ApplianceRemoteFileImport) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApplianceRemoteFileImport) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApplianceRemoteFileImport) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApplianceRemoteFileImport) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceRemoteFileImport) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceRemoteFileImport) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceRemoteFileImport) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceRemoteFileImport) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


