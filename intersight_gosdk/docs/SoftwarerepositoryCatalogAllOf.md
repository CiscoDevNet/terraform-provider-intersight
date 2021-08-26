# SoftwarerepositoryCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.Catalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.Catalog"]
**IsImagePullFailure** | Pointer to **bool** | The status of the image catalog synchronization operation. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the catalog. The names are populated and predefined during MO creation. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**System** | Pointer to [**IamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryCatalogAllOf

`func NewSoftwarerepositoryCatalogAllOf(classId string, objectType string, ) *SoftwarerepositoryCatalogAllOf`

NewSoftwarerepositoryCatalogAllOf instantiates a new SoftwarerepositoryCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCatalogAllOfWithDefaults

`func NewSoftwarerepositoryCatalogAllOfWithDefaults() *SoftwarerepositoryCatalogAllOf`

NewSoftwarerepositoryCatalogAllOfWithDefaults instantiates a new SoftwarerepositoryCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCatalogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCatalogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCatalogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCatalogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCatalogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCatalogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsImagePullFailure

`func (o *SoftwarerepositoryCatalogAllOf) GetIsImagePullFailure() bool`

GetIsImagePullFailure returns the IsImagePullFailure field if non-nil, zero value otherwise.

### GetIsImagePullFailureOk

`func (o *SoftwarerepositoryCatalogAllOf) GetIsImagePullFailureOk() (*bool, bool)`

GetIsImagePullFailureOk returns a tuple with the IsImagePullFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImagePullFailure

`func (o *SoftwarerepositoryCatalogAllOf) SetIsImagePullFailure(v bool)`

SetIsImagePullFailure sets IsImagePullFailure field to given value.

### HasIsImagePullFailure

`func (o *SoftwarerepositoryCatalogAllOf) HasIsImagePullFailure() bool`

HasIsImagePullFailure returns a boolean if a field has been set.

### GetName

`func (o *SoftwarerepositoryCatalogAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryCatalogAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryCatalogAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryCatalogAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *SoftwarerepositoryCatalogAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SoftwarerepositoryCatalogAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SoftwarerepositoryCatalogAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SoftwarerepositoryCatalogAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSystem

`func (o *SoftwarerepositoryCatalogAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SoftwarerepositoryCatalogAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SoftwarerepositoryCatalogAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *SoftwarerepositoryCatalogAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


