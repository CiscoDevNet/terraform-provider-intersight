# CloudAwsKeyPairAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsKeyPair"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsKeyPair"]
**FingerPrint** | Pointer to **string** | Either the SHA-1 digest of the DER encoded private key or  MD5 public key fingerprint. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Used in authenticating to the virtual machine . | [optional] [readonly] 
**AwsBillingUnit** | Pointer to [**CloudAwsBillingUnitRelationship**](cloud.AwsBillingUnit.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsKeyPairAllOf

`func NewCloudAwsKeyPairAllOf(classId string, objectType string, ) *CloudAwsKeyPairAllOf`

NewCloudAwsKeyPairAllOf instantiates a new CloudAwsKeyPairAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsKeyPairAllOfWithDefaults

`func NewCloudAwsKeyPairAllOfWithDefaults() *CloudAwsKeyPairAllOf`

NewCloudAwsKeyPairAllOfWithDefaults instantiates a new CloudAwsKeyPairAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsKeyPairAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsKeyPairAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsKeyPairAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsKeyPairAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsKeyPairAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsKeyPairAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFingerPrint

`func (o *CloudAwsKeyPairAllOf) GetFingerPrint() string`

GetFingerPrint returns the FingerPrint field if non-nil, zero value otherwise.

### GetFingerPrintOk

`func (o *CloudAwsKeyPairAllOf) GetFingerPrintOk() (*string, bool)`

GetFingerPrintOk returns a tuple with the FingerPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerPrint

`func (o *CloudAwsKeyPairAllOf) SetFingerPrint(v string)`

SetFingerPrint sets FingerPrint field to given value.

### HasFingerPrint

`func (o *CloudAwsKeyPairAllOf) HasFingerPrint() bool`

HasFingerPrint returns a boolean if a field has been set.

### GetPublicKey

`func (o *CloudAwsKeyPairAllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CloudAwsKeyPairAllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CloudAwsKeyPairAllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CloudAwsKeyPairAllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetAwsBillingUnit

`func (o *CloudAwsKeyPairAllOf) GetAwsBillingUnit() CloudAwsBillingUnitRelationship`

GetAwsBillingUnit returns the AwsBillingUnit field if non-nil, zero value otherwise.

### GetAwsBillingUnitOk

`func (o *CloudAwsKeyPairAllOf) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool)`

GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBillingUnit

`func (o *CloudAwsKeyPairAllOf) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship)`

SetAwsBillingUnit sets AwsBillingUnit field to given value.

### HasAwsBillingUnit

`func (o *CloudAwsKeyPairAllOf) HasAwsBillingUnit() bool`

HasAwsBillingUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


