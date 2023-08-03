# SoftwarerepositoryAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.Authorization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.Authorization"]
**IsAsdv4AlarmDismissed** | Pointer to **bool** | The state of the alarm dismissal for the &#39;IsAsdDialogDismissed&#39; alarm. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account&#39;s behalf. | [optional] 
**RepositoryType** | Pointer to **string** | The external repository for which this authorization has been provided. The only supported repository today is cisco.com. * &#x60;Cisco&#x60; - External repository hosted on cisco.com. * &#x60;IntersightCloud&#x60; - Repository hosted by the Intersight Cloud. * &#x60;LocalMachine&#x60; - The file is available on the local client machine. Used as an upload source type. * &#x60;NetworkShare&#x60; - External repository in the customer datacenter. This will typically be a file server. | [optional] [default to "Cisco"]
**UserId** | Pointer to **string** | The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account&#39;s behalf. | [optional] 
**Version** | Pointer to **string** | The Automated Software Distribution version of the authorization MO. * &#x60;V3&#x60; - The client is running Automated Software Distribution V3. * &#x60;V4&#x60; - The client is running Automated Software Distribution V4. | [optional] [readonly] [default to "V3"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryAuthorization

`func NewSoftwarerepositoryAuthorization(classId string, objectType string, ) *SoftwarerepositoryAuthorization`

NewSoftwarerepositoryAuthorization instantiates a new SoftwarerepositoryAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryAuthorizationWithDefaults

`func NewSoftwarerepositoryAuthorizationWithDefaults() *SoftwarerepositoryAuthorization`

NewSoftwarerepositoryAuthorizationWithDefaults instantiates a new SoftwarerepositoryAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryAuthorization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryAuthorization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryAuthorization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryAuthorization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryAuthorization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryAuthorization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsAsdv4AlarmDismissed

`func (o *SoftwarerepositoryAuthorization) GetIsAsdv4AlarmDismissed() bool`

GetIsAsdv4AlarmDismissed returns the IsAsdv4AlarmDismissed field if non-nil, zero value otherwise.

### GetIsAsdv4AlarmDismissedOk

`func (o *SoftwarerepositoryAuthorization) GetIsAsdv4AlarmDismissedOk() (*bool, bool)`

GetIsAsdv4AlarmDismissedOk returns a tuple with the IsAsdv4AlarmDismissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsdv4AlarmDismissed

`func (o *SoftwarerepositoryAuthorization) SetIsAsdv4AlarmDismissed(v bool)`

SetIsAsdv4AlarmDismissed sets IsAsdv4AlarmDismissed field to given value.

### HasIsAsdv4AlarmDismissed

`func (o *SoftwarerepositoryAuthorization) HasIsAsdv4AlarmDismissed() bool`

HasIsAsdv4AlarmDismissed returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *SoftwarerepositoryAuthorization) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *SoftwarerepositoryAuthorization) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *SoftwarerepositoryAuthorization) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *SoftwarerepositoryAuthorization) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *SoftwarerepositoryAuthorization) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SoftwarerepositoryAuthorization) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SoftwarerepositoryAuthorization) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SoftwarerepositoryAuthorization) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRepositoryType

`func (o *SoftwarerepositoryAuthorization) GetRepositoryType() string`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *SoftwarerepositoryAuthorization) GetRepositoryTypeOk() (*string, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *SoftwarerepositoryAuthorization) SetRepositoryType(v string)`

SetRepositoryType sets RepositoryType field to given value.

### HasRepositoryType

`func (o *SoftwarerepositoryAuthorization) HasRepositoryType() bool`

HasRepositoryType returns a boolean if a field has been set.

### GetUserId

`func (o *SoftwarerepositoryAuthorization) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SoftwarerepositoryAuthorization) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SoftwarerepositoryAuthorization) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SoftwarerepositoryAuthorization) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwarerepositoryAuthorization) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryAuthorization) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryAuthorization) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryAuthorization) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *SoftwarerepositoryAuthorization) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SoftwarerepositoryAuthorization) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SoftwarerepositoryAuthorization) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SoftwarerepositoryAuthorization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


