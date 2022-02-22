# KvmTunneledKvmPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.TunneledKvmPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.TunneledKvmPolicy"]
**TunneledKvmConfiguration** | Pointer to **bool** | Enable or Disable configuration of tunneled KVM for a specific account. | [optional] [default to false]
**TunneledKvmLaunch** | Pointer to **bool** | Enable or Disable launching tunneled KVM for a specific account. | [optional] [default to false]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewKvmTunneledKvmPolicyAllOf

`func NewKvmTunneledKvmPolicyAllOf(classId string, objectType string, ) *KvmTunneledKvmPolicyAllOf`

NewKvmTunneledKvmPolicyAllOf instantiates a new KvmTunneledKvmPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmTunneledKvmPolicyAllOfWithDefaults

`func NewKvmTunneledKvmPolicyAllOfWithDefaults() *KvmTunneledKvmPolicyAllOf`

NewKvmTunneledKvmPolicyAllOfWithDefaults instantiates a new KvmTunneledKvmPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmTunneledKvmPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmTunneledKvmPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmTunneledKvmPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmTunneledKvmPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmTunneledKvmPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmTunneledKvmPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTunneledKvmConfiguration

`func (o *KvmTunneledKvmPolicyAllOf) GetTunneledKvmConfiguration() bool`

GetTunneledKvmConfiguration returns the TunneledKvmConfiguration field if non-nil, zero value otherwise.

### GetTunneledKvmConfigurationOk

`func (o *KvmTunneledKvmPolicyAllOf) GetTunneledKvmConfigurationOk() (*bool, bool)`

GetTunneledKvmConfigurationOk returns a tuple with the TunneledKvmConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmConfiguration

`func (o *KvmTunneledKvmPolicyAllOf) SetTunneledKvmConfiguration(v bool)`

SetTunneledKvmConfiguration sets TunneledKvmConfiguration field to given value.

### HasTunneledKvmConfiguration

`func (o *KvmTunneledKvmPolicyAllOf) HasTunneledKvmConfiguration() bool`

HasTunneledKvmConfiguration returns a boolean if a field has been set.

### GetTunneledKvmLaunch

`func (o *KvmTunneledKvmPolicyAllOf) GetTunneledKvmLaunch() bool`

GetTunneledKvmLaunch returns the TunneledKvmLaunch field if non-nil, zero value otherwise.

### GetTunneledKvmLaunchOk

`func (o *KvmTunneledKvmPolicyAllOf) GetTunneledKvmLaunchOk() (*bool, bool)`

GetTunneledKvmLaunchOk returns a tuple with the TunneledKvmLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmLaunch

`func (o *KvmTunneledKvmPolicyAllOf) SetTunneledKvmLaunch(v bool)`

SetTunneledKvmLaunch sets TunneledKvmLaunch field to given value.

### HasTunneledKvmLaunch

`func (o *KvmTunneledKvmPolicyAllOf) HasTunneledKvmLaunch() bool`

HasTunneledKvmLaunch returns a boolean if a field has been set.

### GetAccount

`func (o *KvmTunneledKvmPolicyAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *KvmTunneledKvmPolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *KvmTunneledKvmPolicyAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *KvmTunneledKvmPolicyAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


