# SdwanVmanageAccountPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdwan.VmanageAccountPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdwan.VmanageAccountPolicy"]
**EndpointAddress** | Pointer to **string** | VManage server hostname (FQDN) that the acccount holds information for. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Local password for authenticating with the vManage server. | [optional] 
**Port** | Pointer to **int64** | VManage Port number on which the application is running. | [optional] [default to 8443]
**Username** | Pointer to **string** | Local username for authenticating with the vManage server. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]SdwanProfileRelationship**](SdwanProfileRelationship.md) | An array of relationships to sdwanProfile resources. | [optional] 

## Methods

### NewSdwanVmanageAccountPolicyAllOf

`func NewSdwanVmanageAccountPolicyAllOf(classId string, objectType string, ) *SdwanVmanageAccountPolicyAllOf`

NewSdwanVmanageAccountPolicyAllOf instantiates a new SdwanVmanageAccountPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanVmanageAccountPolicyAllOfWithDefaults

`func NewSdwanVmanageAccountPolicyAllOfWithDefaults() *SdwanVmanageAccountPolicyAllOf`

NewSdwanVmanageAccountPolicyAllOfWithDefaults instantiates a new SdwanVmanageAccountPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanVmanageAccountPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanVmanageAccountPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanVmanageAccountPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanVmanageAccountPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndpointAddress

`func (o *SdwanVmanageAccountPolicyAllOf) GetEndpointAddress() string`

GetEndpointAddress returns the EndpointAddress field if non-nil, zero value otherwise.

### GetEndpointAddressOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetEndpointAddressOk() (*string, bool)`

GetEndpointAddressOk returns a tuple with the EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAddress

`func (o *SdwanVmanageAccountPolicyAllOf) SetEndpointAddress(v string)`

SetEndpointAddress sets EndpointAddress field to given value.

### HasEndpointAddress

`func (o *SdwanVmanageAccountPolicyAllOf) HasEndpointAddress() bool`

HasEndpointAddress returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *SdwanVmanageAccountPolicyAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *SdwanVmanageAccountPolicyAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *SdwanVmanageAccountPolicyAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *SdwanVmanageAccountPolicyAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SdwanVmanageAccountPolicyAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SdwanVmanageAccountPolicyAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *SdwanVmanageAccountPolicyAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SdwanVmanageAccountPolicyAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SdwanVmanageAccountPolicyAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *SdwanVmanageAccountPolicyAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SdwanVmanageAccountPolicyAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SdwanVmanageAccountPolicyAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanVmanageAccountPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanVmanageAccountPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanVmanageAccountPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SdwanVmanageAccountPolicyAllOf) GetProfiles() []SdwanProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SdwanVmanageAccountPolicyAllOf) GetProfilesOk() (*[]SdwanProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SdwanVmanageAccountPolicyAllOf) SetProfiles(v []SdwanProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SdwanVmanageAccountPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SdwanVmanageAccountPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SdwanVmanageAccountPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


