# HyperflexProxySettingPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ProxySettingPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ProxySettingPolicy"]
**Hostname** | Pointer to **string** | HTTP Proxy server FQDN or IP. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password for the HTTP Proxy. | [optional] 
**Port** | Pointer to **int64** | The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive. | [optional] 
**Username** | Pointer to **string** | The username for the HTTP Proxy. | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexProxySettingPolicyAllOf

`func NewHyperflexProxySettingPolicyAllOf(classId string, objectType string, ) *HyperflexProxySettingPolicyAllOf`

NewHyperflexProxySettingPolicyAllOf instantiates a new HyperflexProxySettingPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexProxySettingPolicyAllOfWithDefaults

`func NewHyperflexProxySettingPolicyAllOfWithDefaults() *HyperflexProxySettingPolicyAllOf`

NewHyperflexProxySettingPolicyAllOfWithDefaults instantiates a new HyperflexProxySettingPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexProxySettingPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexProxySettingPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexProxySettingPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexProxySettingPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexProxySettingPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexProxySettingPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *HyperflexProxySettingPolicyAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HyperflexProxySettingPolicyAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HyperflexProxySettingPolicyAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HyperflexProxySettingPolicyAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *HyperflexProxySettingPolicyAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *HyperflexProxySettingPolicyAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *HyperflexProxySettingPolicyAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *HyperflexProxySettingPolicyAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *HyperflexProxySettingPolicyAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HyperflexProxySettingPolicyAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HyperflexProxySettingPolicyAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HyperflexProxySettingPolicyAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *HyperflexProxySettingPolicyAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HyperflexProxySettingPolicyAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HyperflexProxySettingPolicyAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *HyperflexProxySettingPolicyAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *HyperflexProxySettingPolicyAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HyperflexProxySettingPolicyAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HyperflexProxySettingPolicyAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HyperflexProxySettingPolicyAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexProxySettingPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexProxySettingPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexProxySettingPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexProxySettingPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexProxySettingPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexProxySettingPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexProxySettingPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexProxySettingPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexProxySettingPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexProxySettingPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


