# BulkRestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.RestResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.RestResult"]
**Body** | Pointer to [**MoBaseMo**](MoBaseMo.md) |  | [optional] 

## Methods

### NewBulkRestResult

`func NewBulkRestResult(classId string, objectType string, ) *BulkRestResult`

NewBulkRestResult instantiates a new BulkRestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRestResultWithDefaults

`func NewBulkRestResultWithDefaults() *BulkRestResult`

NewBulkRestResultWithDefaults instantiates a new BulkRestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkRestResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkRestResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkRestResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkRestResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkRestResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkRestResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *BulkRestResult) GetBody() MoBaseMo`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BulkRestResult) GetBodyOk() (*MoBaseMo, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BulkRestResult) SetBody(v MoBaseMo)`

SetBody sets Body field to given value.

### HasBody

`func (o *BulkRestResult) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


