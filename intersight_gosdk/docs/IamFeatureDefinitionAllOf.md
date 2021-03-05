# IamFeatureDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.FeatureDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.FeatureDefinition"]
**Feature** | Pointer to **string** | The beta feature that will be enabled for specific account. * &#x60;IWO&#x60; - Intersight Workflow Optimizer. * &#x60;Hitachi&#x60; - Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework. * &#x60;Kubernetes&#x60; - Enables ability to create and manage Kubernetes clusters. * &#x60;NetAppIO&#x60; - Support to claim NetApp Storage arrays as IO targets. * &#x60;IvsPublicCloud&#x60; - Enables virtualization service for public clouds. * &#x60;TerraformCloud&#x60; - Enables an ability to create Terraform Cloud. | [optional] [default to "IWO"]

## Methods

### NewIamFeatureDefinitionAllOf

`func NewIamFeatureDefinitionAllOf(classId string, objectType string, ) *IamFeatureDefinitionAllOf`

NewIamFeatureDefinitionAllOf instantiates a new IamFeatureDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamFeatureDefinitionAllOfWithDefaults

`func NewIamFeatureDefinitionAllOfWithDefaults() *IamFeatureDefinitionAllOf`

NewIamFeatureDefinitionAllOfWithDefaults instantiates a new IamFeatureDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamFeatureDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamFeatureDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamFeatureDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamFeatureDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamFeatureDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamFeatureDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeature

`func (o *IamFeatureDefinitionAllOf) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *IamFeatureDefinitionAllOf) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *IamFeatureDefinitionAllOf) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *IamFeatureDefinitionAllOf) HasFeature() bool`

HasFeature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


