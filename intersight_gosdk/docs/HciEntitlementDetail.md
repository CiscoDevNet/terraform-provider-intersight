# HciEntitlementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.EntitlementDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.EntitlementDetail"]
**Category** | Pointer to **string** | The category of the entitlement. | [optional] [readonly] 
**EarliestExpiryDate** | Pointer to **time.Time** | The earliest expiry date of the entitlement. | [optional] [readonly] 
**Meter** | Pointer to **string** | The meter of the entitlement. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the entitlement. | [optional] [readonly] 
**Quantity** | Pointer to **float64** | The quantity of the entitlement. | [optional] [readonly] 
**Scope** | Pointer to **string** | The scope of the entitlement. | [optional] [readonly] 
**SubCategory** | Pointer to **string** | The subCategory of the entitlement. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the license entitlement. | [optional] [readonly] 

## Methods

### NewHciEntitlementDetail

`func NewHciEntitlementDetail(classId string, objectType string, ) *HciEntitlementDetail`

NewHciEntitlementDetail instantiates a new HciEntitlementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciEntitlementDetailWithDefaults

`func NewHciEntitlementDetailWithDefaults() *HciEntitlementDetail`

NewHciEntitlementDetailWithDefaults instantiates a new HciEntitlementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciEntitlementDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciEntitlementDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciEntitlementDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciEntitlementDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciEntitlementDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciEntitlementDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *HciEntitlementDetail) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HciEntitlementDetail) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HciEntitlementDetail) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HciEntitlementDetail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEarliestExpiryDate

`func (o *HciEntitlementDetail) GetEarliestExpiryDate() time.Time`

GetEarliestExpiryDate returns the EarliestExpiryDate field if non-nil, zero value otherwise.

### GetEarliestExpiryDateOk

`func (o *HciEntitlementDetail) GetEarliestExpiryDateOk() (*time.Time, bool)`

GetEarliestExpiryDateOk returns a tuple with the EarliestExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestExpiryDate

`func (o *HciEntitlementDetail) SetEarliestExpiryDate(v time.Time)`

SetEarliestExpiryDate sets EarliestExpiryDate field to given value.

### HasEarliestExpiryDate

`func (o *HciEntitlementDetail) HasEarliestExpiryDate() bool`

HasEarliestExpiryDate returns a boolean if a field has been set.

### GetMeter

`func (o *HciEntitlementDetail) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *HciEntitlementDetail) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *HciEntitlementDetail) SetMeter(v string)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *HciEntitlementDetail) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetName

`func (o *HciEntitlementDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciEntitlementDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciEntitlementDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciEntitlementDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *HciEntitlementDetail) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *HciEntitlementDetail) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *HciEntitlementDetail) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *HciEntitlementDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetScope

`func (o *HciEntitlementDetail) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HciEntitlementDetail) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HciEntitlementDetail) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HciEntitlementDetail) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSubCategory

`func (o *HciEntitlementDetail) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *HciEntitlementDetail) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *HciEntitlementDetail) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *HciEntitlementDetail) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetType

`func (o *HciEntitlementDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciEntitlementDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciEntitlementDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciEntitlementDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


