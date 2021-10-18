# TerraformCloudResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terraform.CloudResource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terraform.CloudResource"]
**CurrentStatus** | Pointer to **string** | Currentstatus of the resource if applicable on the cloud. | [optional] 
**DesiredStatus** | Pointer to **string** | Desiredstatus of the resource if applicable on the cloud. | [optional] 
**ResourceId** | Pointer to **string** | Unique id of the resource from the cloud provider. | [optional] 

## Methods

### NewTerraformCloudResource

`func NewTerraformCloudResource(classId string, objectType string, ) *TerraformCloudResource`

NewTerraformCloudResource instantiates a new TerraformCloudResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformCloudResourceWithDefaults

`func NewTerraformCloudResourceWithDefaults() *TerraformCloudResource`

NewTerraformCloudResourceWithDefaults instantiates a new TerraformCloudResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerraformCloudResource) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerraformCloudResource) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerraformCloudResource) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerraformCloudResource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerraformCloudResource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerraformCloudResource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentStatus

`func (o *TerraformCloudResource) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *TerraformCloudResource) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *TerraformCloudResource) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *TerraformCloudResource) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetDesiredStatus

`func (o *TerraformCloudResource) GetDesiredStatus() string`

GetDesiredStatus returns the DesiredStatus field if non-nil, zero value otherwise.

### GetDesiredStatusOk

`func (o *TerraformCloudResource) GetDesiredStatusOk() (*string, bool)`

GetDesiredStatusOk returns a tuple with the DesiredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredStatus

`func (o *TerraformCloudResource) SetDesiredStatus(v string)`

SetDesiredStatus sets DesiredStatus field to given value.

### HasDesiredStatus

`func (o *TerraformCloudResource) HasDesiredStatus() bool`

HasDesiredStatus returns a boolean if a field has been set.

### GetResourceId

`func (o *TerraformCloudResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TerraformCloudResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TerraformCloudResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *TerraformCloudResource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


