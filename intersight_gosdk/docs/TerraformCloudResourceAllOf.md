# TerraformCloudResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terraform.CloudResource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terraform.CloudResource"]
**CurrentStatus** | Pointer to **string** | Currentstatus of the resource if applicable on the cloud. | [optional] 
**DesiredStatus** | Pointer to **string** | Desiredstatus of the resource if applicable on the cloud. | [optional] 
**ResourceId** | Pointer to **string** | Unique id of the resource from the cloud provider. | [optional] 

## Methods

### NewTerraformCloudResourceAllOf

`func NewTerraformCloudResourceAllOf(classId string, objectType string, ) *TerraformCloudResourceAllOf`

NewTerraformCloudResourceAllOf instantiates a new TerraformCloudResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformCloudResourceAllOfWithDefaults

`func NewTerraformCloudResourceAllOfWithDefaults() *TerraformCloudResourceAllOf`

NewTerraformCloudResourceAllOfWithDefaults instantiates a new TerraformCloudResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerraformCloudResourceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerraformCloudResourceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerraformCloudResourceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerraformCloudResourceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerraformCloudResourceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerraformCloudResourceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentStatus

`func (o *TerraformCloudResourceAllOf) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *TerraformCloudResourceAllOf) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *TerraformCloudResourceAllOf) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *TerraformCloudResourceAllOf) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetDesiredStatus

`func (o *TerraformCloudResourceAllOf) GetDesiredStatus() string`

GetDesiredStatus returns the DesiredStatus field if non-nil, zero value otherwise.

### GetDesiredStatusOk

`func (o *TerraformCloudResourceAllOf) GetDesiredStatusOk() (*string, bool)`

GetDesiredStatusOk returns a tuple with the DesiredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredStatus

`func (o *TerraformCloudResourceAllOf) SetDesiredStatus(v string)`

SetDesiredStatus sets DesiredStatus field to given value.

### HasDesiredStatus

`func (o *TerraformCloudResourceAllOf) HasDesiredStatus() bool`

HasDesiredStatus returns a boolean if a field has been set.

### GetResourceId

`func (o *TerraformCloudResourceAllOf) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TerraformCloudResourceAllOf) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TerraformCloudResourceAllOf) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *TerraformCloudResourceAllOf) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


