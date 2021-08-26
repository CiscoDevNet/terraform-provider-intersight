# KubernetesAddonPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonPolicy"]
**AddonConfiguration** | Pointer to [**NullableKubernetesAddonConfiguration**](KubernetesAddonConfiguration.md) |  | [optional] 
**AddonDefinition** | Pointer to [**KubernetesAddonDefinitionRelationship**](KubernetesAddonDefinitionRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesAddonPolicyAllOf

`func NewKubernetesAddonPolicyAllOf(classId string, objectType string, ) *KubernetesAddonPolicyAllOf`

NewKubernetesAddonPolicyAllOf instantiates a new KubernetesAddonPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonPolicyAllOfWithDefaults

`func NewKubernetesAddonPolicyAllOfWithDefaults() *KubernetesAddonPolicyAllOf`

NewKubernetesAddonPolicyAllOfWithDefaults instantiates a new KubernetesAddonPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddonConfiguration

`func (o *KubernetesAddonPolicyAllOf) GetAddonConfiguration() KubernetesAddonConfiguration`

GetAddonConfiguration returns the AddonConfiguration field if non-nil, zero value otherwise.

### GetAddonConfigurationOk

`func (o *KubernetesAddonPolicyAllOf) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool)`

GetAddonConfigurationOk returns a tuple with the AddonConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonConfiguration

`func (o *KubernetesAddonPolicyAllOf) SetAddonConfiguration(v KubernetesAddonConfiguration)`

SetAddonConfiguration sets AddonConfiguration field to given value.

### HasAddonConfiguration

`func (o *KubernetesAddonPolicyAllOf) HasAddonConfiguration() bool`

HasAddonConfiguration returns a boolean if a field has been set.

### SetAddonConfigurationNil

`func (o *KubernetesAddonPolicyAllOf) SetAddonConfigurationNil(b bool)`

 SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil

### UnsetAddonConfiguration
`func (o *KubernetesAddonPolicyAllOf) UnsetAddonConfiguration()`

UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
### GetAddonDefinition

`func (o *KubernetesAddonPolicyAllOf) GetAddonDefinition() KubernetesAddonDefinitionRelationship`

GetAddonDefinition returns the AddonDefinition field if non-nil, zero value otherwise.

### GetAddonDefinitionOk

`func (o *KubernetesAddonPolicyAllOf) GetAddonDefinitionOk() (*KubernetesAddonDefinitionRelationship, bool)`

GetAddonDefinitionOk returns a tuple with the AddonDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonDefinition

`func (o *KubernetesAddonPolicyAllOf) SetAddonDefinition(v KubernetesAddonDefinitionRelationship)`

SetAddonDefinition sets AddonDefinition field to given value.

### HasAddonDefinition

`func (o *KubernetesAddonPolicyAllOf) HasAddonDefinition() bool`

HasAddonDefinition returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesAddonPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddonPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddonPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddonPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


