# IamTenantLinking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.TenantLinking"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.TenantLinking"]
**Product** | Pointer to **string** | The product name for the tenant linking (e.g., CUI, CCC). | [optional] 
**TenantId** | Pointer to **string** | The tenant ID for the linked product. | [optional] 

## Methods

### NewIamTenantLinking

`func NewIamTenantLinking(classId string, objectType string, ) *IamTenantLinking`

NewIamTenantLinking instantiates a new IamTenantLinking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTenantLinkingWithDefaults

`func NewIamTenantLinkingWithDefaults() *IamTenantLinking`

NewIamTenantLinkingWithDefaults instantiates a new IamTenantLinking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamTenantLinking) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamTenantLinking) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamTenantLinking) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamTenantLinking) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamTenantLinking) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamTenantLinking) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProduct

`func (o *IamTenantLinking) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *IamTenantLinking) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *IamTenantLinking) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *IamTenantLinking) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTenantId

`func (o *IamTenantLinking) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IamTenantLinking) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IamTenantLinking) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IamTenantLinking) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


