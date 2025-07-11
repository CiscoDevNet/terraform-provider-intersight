# HciApiLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ApiLimit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ApiLimit"]
**ActualSize** | Pointer to **int32** | The actual number of records returned by the API. | [optional] [readonly] 
**ApiName** | Pointer to **string** | The name of the Nutanix API that has reach the limit. | [optional] [readonly] 
**IsLimitReached** | Pointer to **bool** | Indicates if the limit is reached. | [optional] [readonly] 
**Limit** | Pointer to **int32** | The number of records the API is allowed to return. | [optional] [readonly] 

## Methods

### NewHciApiLimit

`func NewHciApiLimit(classId string, objectType string, ) *HciApiLimit`

NewHciApiLimit instantiates a new HciApiLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciApiLimitWithDefaults

`func NewHciApiLimitWithDefaults() *HciApiLimit`

NewHciApiLimitWithDefaults instantiates a new HciApiLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciApiLimit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciApiLimit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciApiLimit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciApiLimit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciApiLimit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciApiLimit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActualSize

`func (o *HciApiLimit) GetActualSize() int32`

GetActualSize returns the ActualSize field if non-nil, zero value otherwise.

### GetActualSizeOk

`func (o *HciApiLimit) GetActualSizeOk() (*int32, bool)`

GetActualSizeOk returns a tuple with the ActualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSize

`func (o *HciApiLimit) SetActualSize(v int32)`

SetActualSize sets ActualSize field to given value.

### HasActualSize

`func (o *HciApiLimit) HasActualSize() bool`

HasActualSize returns a boolean if a field has been set.

### GetApiName

`func (o *HciApiLimit) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *HciApiLimit) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *HciApiLimit) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *HciApiLimit) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetIsLimitReached

`func (o *HciApiLimit) GetIsLimitReached() bool`

GetIsLimitReached returns the IsLimitReached field if non-nil, zero value otherwise.

### GetIsLimitReachedOk

`func (o *HciApiLimit) GetIsLimitReachedOk() (*bool, bool)`

GetIsLimitReachedOk returns a tuple with the IsLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimitReached

`func (o *HciApiLimit) SetIsLimitReached(v bool)`

SetIsLimitReached sets IsLimitReached field to given value.

### HasIsLimitReached

`func (o *HciApiLimit) HasIsLimitReached() bool`

HasIsLimitReached returns a boolean if a field has been set.

### GetLimit

`func (o *HciApiLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *HciApiLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *HciApiLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *HciApiLimit) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


