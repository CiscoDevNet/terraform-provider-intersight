# IamFeatureDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | Pointer to **string** | The beta feature that will be enabled for specific account. * &#x60;IWO&#x60; - Intersight Workflow Optimizer. | [optional] [default to "IWO"]

## Methods

### NewIamFeatureDefinition

`func NewIamFeatureDefinition() *IamFeatureDefinition`

NewIamFeatureDefinition instantiates a new IamFeatureDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamFeatureDefinitionWithDefaults

`func NewIamFeatureDefinitionWithDefaults() *IamFeatureDefinition`

NewIamFeatureDefinitionWithDefaults instantiates a new IamFeatureDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *IamFeatureDefinition) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *IamFeatureDefinition) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *IamFeatureDefinition) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *IamFeatureDefinition) HasFeature() bool`

HasFeature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


