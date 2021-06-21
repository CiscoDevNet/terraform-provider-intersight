# AssetSubscriptionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.Subscription"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.Subscription"]
**ApplicationName** | Pointer to **string** | Application name reported by Cisco Install Base. | [optional] [readonly] 
**SubscriptionRefId** | Pointer to **string** | Identifies the consumption-based subscription. | [optional] [readonly] 
**Deployments** | Pointer to [**[]AssetDeploymentRelationship**](AssetDeploymentRelationship.md) | An array of relationships to assetDeployment resources. | [optional] [readonly] 
**SubscriptionAccount** | Pointer to [**AssetSubscriptionAccountRelationship**](asset.SubscriptionAccount.Relationship.md) |  | [optional] 

## Methods

### NewAssetSubscriptionAllOf

`func NewAssetSubscriptionAllOf(classId string, objectType string, ) *AssetSubscriptionAllOf`

NewAssetSubscriptionAllOf instantiates a new AssetSubscriptionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSubscriptionAllOfWithDefaults

`func NewAssetSubscriptionAllOfWithDefaults() *AssetSubscriptionAllOf`

NewAssetSubscriptionAllOfWithDefaults instantiates a new AssetSubscriptionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetSubscriptionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetSubscriptionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetSubscriptionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetSubscriptionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetSubscriptionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetSubscriptionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApplicationName

`func (o *AssetSubscriptionAllOf) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AssetSubscriptionAllOf) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AssetSubscriptionAllOf) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AssetSubscriptionAllOf) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetSubscriptionRefId

`func (o *AssetSubscriptionAllOf) GetSubscriptionRefId() string`

GetSubscriptionRefId returns the SubscriptionRefId field if non-nil, zero value otherwise.

### GetSubscriptionRefIdOk

`func (o *AssetSubscriptionAllOf) GetSubscriptionRefIdOk() (*string, bool)`

GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRefId

`func (o *AssetSubscriptionAllOf) SetSubscriptionRefId(v string)`

SetSubscriptionRefId sets SubscriptionRefId field to given value.

### HasSubscriptionRefId

`func (o *AssetSubscriptionAllOf) HasSubscriptionRefId() bool`

HasSubscriptionRefId returns a boolean if a field has been set.

### GetDeployments

`func (o *AssetSubscriptionAllOf) GetDeployments() []AssetDeploymentRelationship`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *AssetSubscriptionAllOf) GetDeploymentsOk() (*[]AssetDeploymentRelationship, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *AssetSubscriptionAllOf) SetDeployments(v []AssetDeploymentRelationship)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *AssetSubscriptionAllOf) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### SetDeploymentsNil

`func (o *AssetSubscriptionAllOf) SetDeploymentsNil(b bool)`

 SetDeploymentsNil sets the value for Deployments to be an explicit nil

### UnsetDeployments
`func (o *AssetSubscriptionAllOf) UnsetDeployments()`

UnsetDeployments ensures that no value is present for Deployments, not even an explicit nil
### GetSubscriptionAccount

`func (o *AssetSubscriptionAllOf) GetSubscriptionAccount() AssetSubscriptionAccountRelationship`

GetSubscriptionAccount returns the SubscriptionAccount field if non-nil, zero value otherwise.

### GetSubscriptionAccountOk

`func (o *AssetSubscriptionAllOf) GetSubscriptionAccountOk() (*AssetSubscriptionAccountRelationship, bool)`

GetSubscriptionAccountOk returns a tuple with the SubscriptionAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAccount

`func (o *AssetSubscriptionAllOf) SetSubscriptionAccount(v AssetSubscriptionAccountRelationship)`

SetSubscriptionAccount sets SubscriptionAccount field to given value.

### HasSubscriptionAccount

`func (o *AssetSubscriptionAllOf) HasSubscriptionAccount() bool`

HasSubscriptionAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


