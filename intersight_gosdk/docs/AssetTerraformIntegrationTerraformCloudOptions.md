# AssetTerraformIntegrationTerraformCloudOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.TerraformIntegrationTerraformCloudOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.TerraformIntegrationTerraformCloudOptions"]
**DefaultTerraformOrganization** | Pointer to **string** | Default organization for Terraform Cloud platform type. | [optional] 

## Methods

### NewAssetTerraformIntegrationTerraformCloudOptions

`func NewAssetTerraformIntegrationTerraformCloudOptions(classId string, objectType string, ) *AssetTerraformIntegrationTerraformCloudOptions`

NewAssetTerraformIntegrationTerraformCloudOptions instantiates a new AssetTerraformIntegrationTerraformCloudOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTerraformIntegrationTerraformCloudOptionsWithDefaults

`func NewAssetTerraformIntegrationTerraformCloudOptionsWithDefaults() *AssetTerraformIntegrationTerraformCloudOptions`

NewAssetTerraformIntegrationTerraformCloudOptionsWithDefaults instantiates a new AssetTerraformIntegrationTerraformCloudOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetTerraformIntegrationTerraformCloudOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetTerraformIntegrationTerraformCloudOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetTerraformIntegrationTerraformCloudOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetTerraformIntegrationTerraformCloudOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetTerraformIntegrationTerraformCloudOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetTerraformIntegrationTerraformCloudOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformCloudOptions) GetDefaultTerraformOrganization() string`

GetDefaultTerraformOrganization returns the DefaultTerraformOrganization field if non-nil, zero value otherwise.

### GetDefaultTerraformOrganizationOk

`func (o *AssetTerraformIntegrationTerraformCloudOptions) GetDefaultTerraformOrganizationOk() (*string, bool)`

GetDefaultTerraformOrganizationOk returns a tuple with the DefaultTerraformOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformCloudOptions) SetDefaultTerraformOrganization(v string)`

SetDefaultTerraformOrganization sets DefaultTerraformOrganization field to given value.

### HasDefaultTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformCloudOptions) HasDefaultTerraformOrganization() bool`

HasDefaultTerraformOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


