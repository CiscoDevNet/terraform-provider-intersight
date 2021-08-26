# CloudBaseBillingUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsBillingUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsBillingUnit"]
**Identity** | Pointer to **string** | The ID of the paying account. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the paying account. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the account, for example Active, Suspended, etc. | [optional] [readonly] 

## Methods

### NewCloudBaseBillingUnit

`func NewCloudBaseBillingUnit(classId string, objectType string, ) *CloudBaseBillingUnit`

NewCloudBaseBillingUnit instantiates a new CloudBaseBillingUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseBillingUnitWithDefaults

`func NewCloudBaseBillingUnitWithDefaults() *CloudBaseBillingUnit`

NewCloudBaseBillingUnitWithDefaults instantiates a new CloudBaseBillingUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseBillingUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseBillingUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseBillingUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseBillingUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseBillingUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseBillingUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *CloudBaseBillingUnit) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudBaseBillingUnit) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudBaseBillingUnit) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudBaseBillingUnit) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudBaseBillingUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBaseBillingUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBaseBillingUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBaseBillingUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBaseBillingUnit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBaseBillingUnit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBaseBillingUnit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBaseBillingUnit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


