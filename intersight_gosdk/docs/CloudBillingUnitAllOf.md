# CloudBillingUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.BillingUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.BillingUnit"]
**BillingId** | Pointer to **string** | The ID of the paying account. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the paying account. | [optional] [readonly] 

## Methods

### NewCloudBillingUnitAllOf

`func NewCloudBillingUnitAllOf(classId string, objectType string, ) *CloudBillingUnitAllOf`

NewCloudBillingUnitAllOf instantiates a new CloudBillingUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBillingUnitAllOfWithDefaults

`func NewCloudBillingUnitAllOfWithDefaults() *CloudBillingUnitAllOf`

NewCloudBillingUnitAllOfWithDefaults instantiates a new CloudBillingUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBillingUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBillingUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBillingUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBillingUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBillingUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBillingUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingId

`func (o *CloudBillingUnitAllOf) GetBillingId() string`

GetBillingId returns the BillingId field if non-nil, zero value otherwise.

### GetBillingIdOk

`func (o *CloudBillingUnitAllOf) GetBillingIdOk() (*string, bool)`

GetBillingIdOk returns a tuple with the BillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingId

`func (o *CloudBillingUnitAllOf) SetBillingId(v string)`

SetBillingId sets BillingId field to given value.

### HasBillingId

`func (o *CloudBillingUnitAllOf) HasBillingId() bool`

HasBillingId returns a boolean if a field has been set.

### GetName

`func (o *CloudBillingUnitAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBillingUnitAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBillingUnitAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBillingUnitAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


