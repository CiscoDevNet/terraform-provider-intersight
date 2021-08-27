# IqnpoolBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.Block"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.Block"]
**IqnSuffixBlock** | Pointer to [**IqnpoolIqnSuffixBlock**](IqnpoolIqnSuffixBlock.md) |  | [optional] 
**Pool** | Pointer to [**IqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewIqnpoolBlock

`func NewIqnpoolBlock(classId string, objectType string, ) *IqnpoolBlock`

NewIqnpoolBlock instantiates a new IqnpoolBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolBlockWithDefaults

`func NewIqnpoolBlockWithDefaults() *IqnpoolBlock`

NewIqnpoolBlockWithDefaults instantiates a new IqnpoolBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolBlock) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolBlock) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolBlock) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIqnSuffixBlock

`func (o *IqnpoolBlock) GetIqnSuffixBlock() IqnpoolIqnSuffixBlock`

GetIqnSuffixBlock returns the IqnSuffixBlock field if non-nil, zero value otherwise.

### GetIqnSuffixBlockOk

`func (o *IqnpoolBlock) GetIqnSuffixBlockOk() (*IqnpoolIqnSuffixBlock, bool)`

GetIqnSuffixBlockOk returns a tuple with the IqnSuffixBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffixBlock

`func (o *IqnpoolBlock) SetIqnSuffixBlock(v IqnpoolIqnSuffixBlock)`

SetIqnSuffixBlock sets IqnSuffixBlock field to given value.

### HasIqnSuffixBlock

`func (o *IqnpoolBlock) HasIqnSuffixBlock() bool`

HasIqnSuffixBlock returns a boolean if a field has been set.

### GetPool

`func (o *IqnpoolBlock) GetPool() IqnpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IqnpoolBlock) GetPoolOk() (*IqnpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IqnpoolBlock) SetPool(v IqnpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IqnpoolBlock) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


