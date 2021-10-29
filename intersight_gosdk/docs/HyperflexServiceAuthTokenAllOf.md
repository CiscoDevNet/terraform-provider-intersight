# HyperflexServiceAuthTokenAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServiceAuthToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServiceAuthToken"]
**ClientId** | Pointer to **string** | Client Id or tenant Id of the entity that uses the service auth token. | [optional] 
**CsiVersion** | Pointer to **string** | Version of Container Storage Interface (CSI) that the tokenOwner is associated with. | [optional] 
**ServiceAuthToken** | Pointer to **string** | Service auth token that has been created by HyperFlex cluster. | [optional] [readonly] 
**Status** | Pointer to **string** | Represents status of ervice auth claim or revocation. * &#x60;Unknown&#x60; - Unknown claim state of the service auth token. * &#x60;Claiming&#x60; - The service auth token claim is in progress. * &#x60;Claimed&#x60; - The service auth token has been successfully claimed. * &#x60;FailedToClaim&#x60; - Cannot claim the service auth token on the underlying HyperFlex cluster. * &#x60;Revoking&#x60; - The service auth token revocation is in progress. * &#x60;Revoked&#x60; - The service auth token revocation has been successfully revoked. * &#x60;FailedToRevoke&#x60; - Cannot revoke the service auth token on the underlying HyperFlex cluster. | [optional] [readonly] [default to "Unknown"]
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**TokenOwner** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexServiceAuthTokenAllOf

`func NewHyperflexServiceAuthTokenAllOf(classId string, objectType string, ) *HyperflexServiceAuthTokenAllOf`

NewHyperflexServiceAuthTokenAllOf instantiates a new HyperflexServiceAuthTokenAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServiceAuthTokenAllOfWithDefaults

`func NewHyperflexServiceAuthTokenAllOfWithDefaults() *HyperflexServiceAuthTokenAllOf`

NewHyperflexServiceAuthTokenAllOfWithDefaults instantiates a new HyperflexServiceAuthTokenAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServiceAuthTokenAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServiceAuthTokenAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServiceAuthTokenAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServiceAuthTokenAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServiceAuthTokenAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServiceAuthTokenAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientId

`func (o *HyperflexServiceAuthTokenAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *HyperflexServiceAuthTokenAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *HyperflexServiceAuthTokenAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *HyperflexServiceAuthTokenAllOf) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCsiVersion

`func (o *HyperflexServiceAuthTokenAllOf) GetCsiVersion() string`

GetCsiVersion returns the CsiVersion field if non-nil, zero value otherwise.

### GetCsiVersionOk

`func (o *HyperflexServiceAuthTokenAllOf) GetCsiVersionOk() (*string, bool)`

GetCsiVersionOk returns a tuple with the CsiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsiVersion

`func (o *HyperflexServiceAuthTokenAllOf) SetCsiVersion(v string)`

SetCsiVersion sets CsiVersion field to given value.

### HasCsiVersion

`func (o *HyperflexServiceAuthTokenAllOf) HasCsiVersion() bool`

HasCsiVersion returns a boolean if a field has been set.

### GetServiceAuthToken

`func (o *HyperflexServiceAuthTokenAllOf) GetServiceAuthToken() string`

GetServiceAuthToken returns the ServiceAuthToken field if non-nil, zero value otherwise.

### GetServiceAuthTokenOk

`func (o *HyperflexServiceAuthTokenAllOf) GetServiceAuthTokenOk() (*string, bool)`

GetServiceAuthTokenOk returns a tuple with the ServiceAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAuthToken

`func (o *HyperflexServiceAuthTokenAllOf) SetServiceAuthToken(v string)`

SetServiceAuthToken sets ServiceAuthToken field to given value.

### HasServiceAuthToken

`func (o *HyperflexServiceAuthTokenAllOf) HasServiceAuthToken() bool`

HasServiceAuthToken returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexServiceAuthTokenAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexServiceAuthTokenAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexServiceAuthTokenAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexServiceAuthTokenAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexServiceAuthTokenAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexServiceAuthTokenAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexServiceAuthTokenAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexServiceAuthTokenAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexServiceAuthTokenAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexServiceAuthTokenAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexServiceAuthTokenAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexServiceAuthTokenAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTokenOwner

`func (o *HyperflexServiceAuthTokenAllOf) GetTokenOwner() MoBaseMoRelationship`

GetTokenOwner returns the TokenOwner field if non-nil, zero value otherwise.

### GetTokenOwnerOk

`func (o *HyperflexServiceAuthTokenAllOf) GetTokenOwnerOk() (*MoBaseMoRelationship, bool)`

GetTokenOwnerOk returns a tuple with the TokenOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOwner

`func (o *HyperflexServiceAuthTokenAllOf) SetTokenOwner(v MoBaseMoRelationship)`

SetTokenOwner sets TokenOwner field to given value.

### HasTokenOwner

`func (o *HyperflexServiceAuthTokenAllOf) HasTokenOwner() bool`

HasTokenOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


