# SdwanVmanageAccountPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointAddress** | Pointer to **string** | VManage server hostname (FQDN) that the acccount holds information for. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | Local password for authenticating with the vManage server. | [optional] 
**Port** | Pointer to **int64** | VManage Port number on which the application is running. | [optional] 
**Username** | Pointer to **string** | Local username for authenticating with the vManage server. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]SdwanProfileRelationship**](sdwan.Profile.Relationship.md) | An array of relationships to sdwanProfile resources. | [optional] 

## Methods

### NewSdwanVmanageAccountPolicy

`func NewSdwanVmanageAccountPolicy() *SdwanVmanageAccountPolicy`

NewSdwanVmanageAccountPolicy instantiates a new SdwanVmanageAccountPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanVmanageAccountPolicyWithDefaults

`func NewSdwanVmanageAccountPolicyWithDefaults() *SdwanVmanageAccountPolicy`

NewSdwanVmanageAccountPolicyWithDefaults instantiates a new SdwanVmanageAccountPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointAddress

`func (o *SdwanVmanageAccountPolicy) GetEndpointAddress() string`

GetEndpointAddress returns the EndpointAddress field if non-nil, zero value otherwise.

### GetEndpointAddressOk

`func (o *SdwanVmanageAccountPolicy) GetEndpointAddressOk() (*string, bool)`

GetEndpointAddressOk returns a tuple with the EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAddress

`func (o *SdwanVmanageAccountPolicy) SetEndpointAddress(v string)`

SetEndpointAddress sets EndpointAddress field to given value.

### HasEndpointAddress

`func (o *SdwanVmanageAccountPolicy) HasEndpointAddress() bool`

HasEndpointAddress returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *SdwanVmanageAccountPolicy) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *SdwanVmanageAccountPolicy) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *SdwanVmanageAccountPolicy) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *SdwanVmanageAccountPolicy) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *SdwanVmanageAccountPolicy) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SdwanVmanageAccountPolicy) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SdwanVmanageAccountPolicy) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SdwanVmanageAccountPolicy) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *SdwanVmanageAccountPolicy) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SdwanVmanageAccountPolicy) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SdwanVmanageAccountPolicy) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SdwanVmanageAccountPolicy) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *SdwanVmanageAccountPolicy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SdwanVmanageAccountPolicy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SdwanVmanageAccountPolicy) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SdwanVmanageAccountPolicy) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanVmanageAccountPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanVmanageAccountPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanVmanageAccountPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanVmanageAccountPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SdwanVmanageAccountPolicy) GetProfiles() []SdwanProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SdwanVmanageAccountPolicy) GetProfilesOk() (*[]SdwanProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SdwanVmanageAccountPolicy) SetProfiles(v []SdwanProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SdwanVmanageAccountPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SdwanVmanageAccountPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SdwanVmanageAccountPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


