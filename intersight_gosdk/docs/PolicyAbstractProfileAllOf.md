# PolicyAbstractProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the profile. | [optional] 
**Name** | Pointer to **string** | Name of the concrete profile. | [optional] 
**Type** | Pointer to **string** | Defines the type of the profile. Accepted value is instance. * &#x60;instance&#x60; - The profile defines the configuration for a specific instance of a target. | [optional] [default to "instance"]
**SrcTemplate** | Pointer to [**PolicyAbstractProfileRelationship**](policy.AbstractProfile.Relationship.md) |  | [optional] 

## Methods

### NewPolicyAbstractProfileAllOf

`func NewPolicyAbstractProfileAllOf() *PolicyAbstractProfileAllOf`

NewPolicyAbstractProfileAllOf instantiates a new PolicyAbstractProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractProfileAllOfWithDefaults

`func NewPolicyAbstractProfileAllOfWithDefaults() *PolicyAbstractProfileAllOf`

NewPolicyAbstractProfileAllOfWithDefaults instantiates a new PolicyAbstractProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PolicyAbstractProfileAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyAbstractProfileAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyAbstractProfileAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyAbstractProfileAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyAbstractProfileAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAbstractProfileAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAbstractProfileAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAbstractProfileAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PolicyAbstractProfileAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyAbstractProfileAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyAbstractProfileAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyAbstractProfileAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSrcTemplate

`func (o *PolicyAbstractProfileAllOf) GetSrcTemplate() PolicyAbstractProfileRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *PolicyAbstractProfileAllOf) GetSrcTemplateOk() (*PolicyAbstractProfileRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *PolicyAbstractProfileAllOf) SetSrcTemplate(v PolicyAbstractProfileRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *PolicyAbstractProfileAllOf) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


