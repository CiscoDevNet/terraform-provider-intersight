# CloudBaseBillingUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsBillingUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsBillingUnit"]
**Identity** | Pointer to **string** | The ID of the paying account. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the paying account. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the account, for example Active, Suspended, etc. | [optional] [readonly] 

## Methods

### NewCloudBaseBillingUnitAllOf

`func NewCloudBaseBillingUnitAllOf(classId string, objectType string, ) *CloudBaseBillingUnitAllOf`

NewCloudBaseBillingUnitAllOf instantiates a new CloudBaseBillingUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseBillingUnitAllOfWithDefaults

`func NewCloudBaseBillingUnitAllOfWithDefaults() *CloudBaseBillingUnitAllOf`

NewCloudBaseBillingUnitAllOfWithDefaults instantiates a new CloudBaseBillingUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseBillingUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseBillingUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseBillingUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseBillingUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseBillingUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseBillingUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *CloudBaseBillingUnitAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudBaseBillingUnitAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudBaseBillingUnitAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudBaseBillingUnitAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudBaseBillingUnitAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBaseBillingUnitAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBaseBillingUnitAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBaseBillingUnitAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBaseBillingUnitAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBaseBillingUnitAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBaseBillingUnitAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBaseBillingUnitAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


