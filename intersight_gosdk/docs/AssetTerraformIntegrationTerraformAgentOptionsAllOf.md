# AssetTerraformIntegrationTerraformAgentOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.TerraformIntegrationTerraformAgentOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.TerraformIntegrationTerraformAgentOptions"]
**ManagedHosts** | Pointer to **[]string** |  | [optional] 
**TerraformAgentPoolName** | Pointer to **string** | Agent pool name for Terraform Agent platform type. | [optional] 
**TerraformOrganization** | Pointer to **string** | Organization for Terraform Agent platform type. | [optional] 

## Methods

### NewAssetTerraformIntegrationTerraformAgentOptionsAllOf

`func NewAssetTerraformIntegrationTerraformAgentOptionsAllOf(classId string, objectType string, ) *AssetTerraformIntegrationTerraformAgentOptionsAllOf`

NewAssetTerraformIntegrationTerraformAgentOptionsAllOf instantiates a new AssetTerraformIntegrationTerraformAgentOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTerraformIntegrationTerraformAgentOptionsAllOfWithDefaults

`func NewAssetTerraformIntegrationTerraformAgentOptionsAllOfWithDefaults() *AssetTerraformIntegrationTerraformAgentOptionsAllOf`

NewAssetTerraformIntegrationTerraformAgentOptionsAllOfWithDefaults instantiates a new AssetTerraformIntegrationTerraformAgentOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetManagedHosts

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetManagedHosts() []string`

GetManagedHosts returns the ManagedHosts field if non-nil, zero value otherwise.

### GetManagedHostsOk

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetManagedHostsOk() (*[]string, bool)`

GetManagedHostsOk returns a tuple with the ManagedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedHosts

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetManagedHosts(v []string)`

SetManagedHosts sets ManagedHosts field to given value.

### HasManagedHosts

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) HasManagedHosts() bool`

HasManagedHosts returns a boolean if a field has been set.

### SetManagedHostsNil

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetManagedHostsNil(b bool)`

 SetManagedHostsNil sets the value for ManagedHosts to be an explicit nil

### UnsetManagedHosts
`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) UnsetManagedHosts()`

UnsetManagedHosts ensures that no value is present for ManagedHosts, not even an explicit nil
### GetTerraformAgentPoolName

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformAgentPoolName() string`

GetTerraformAgentPoolName returns the TerraformAgentPoolName field if non-nil, zero value otherwise.

### GetTerraformAgentPoolNameOk

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformAgentPoolNameOk() (*string, bool)`

GetTerraformAgentPoolNameOk returns a tuple with the TerraformAgentPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformAgentPoolName

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetTerraformAgentPoolName(v string)`

SetTerraformAgentPoolName sets TerraformAgentPoolName field to given value.

### HasTerraformAgentPoolName

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) HasTerraformAgentPoolName() bool`

HasTerraformAgentPoolName returns a boolean if a field has been set.

### GetTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformOrganization() string`

GetTerraformOrganization returns the TerraformOrganization field if non-nil, zero value otherwise.

### GetTerraformOrganizationOk

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) GetTerraformOrganizationOk() (*string, bool)`

GetTerraformOrganizationOk returns a tuple with the TerraformOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) SetTerraformOrganization(v string)`

SetTerraformOrganization sets TerraformOrganization field to given value.

### HasTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformAgentOptionsAllOf) HasTerraformOrganization() bool`

HasTerraformOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


