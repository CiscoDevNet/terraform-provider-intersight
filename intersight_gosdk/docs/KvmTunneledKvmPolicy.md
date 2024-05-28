# KvmTunneledKvmPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.TunneledKvmPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.TunneledKvmPolicy"]
**TunneledKvmConfiguration** | Pointer to **bool** | Enable or Disable configuration of tunneled KVM for a specific account. | [optional] [default to false]
**TunneledKvmLaunch** | Pointer to **bool** | Enable or Disable launching tunneled KVM for a specific account. | [optional] [default to false]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewKvmTunneledKvmPolicy

`func NewKvmTunneledKvmPolicy(classId string, objectType string, ) *KvmTunneledKvmPolicy`

NewKvmTunneledKvmPolicy instantiates a new KvmTunneledKvmPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmTunneledKvmPolicyWithDefaults

`func NewKvmTunneledKvmPolicyWithDefaults() *KvmTunneledKvmPolicy`

NewKvmTunneledKvmPolicyWithDefaults instantiates a new KvmTunneledKvmPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmTunneledKvmPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmTunneledKvmPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmTunneledKvmPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmTunneledKvmPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmTunneledKvmPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmTunneledKvmPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTunneledKvmConfiguration

`func (o *KvmTunneledKvmPolicy) GetTunneledKvmConfiguration() bool`

GetTunneledKvmConfiguration returns the TunneledKvmConfiguration field if non-nil, zero value otherwise.

### GetTunneledKvmConfigurationOk

`func (o *KvmTunneledKvmPolicy) GetTunneledKvmConfigurationOk() (*bool, bool)`

GetTunneledKvmConfigurationOk returns a tuple with the TunneledKvmConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmConfiguration

`func (o *KvmTunneledKvmPolicy) SetTunneledKvmConfiguration(v bool)`

SetTunneledKvmConfiguration sets TunneledKvmConfiguration field to given value.

### HasTunneledKvmConfiguration

`func (o *KvmTunneledKvmPolicy) HasTunneledKvmConfiguration() bool`

HasTunneledKvmConfiguration returns a boolean if a field has been set.

### GetTunneledKvmLaunch

`func (o *KvmTunneledKvmPolicy) GetTunneledKvmLaunch() bool`

GetTunneledKvmLaunch returns the TunneledKvmLaunch field if non-nil, zero value otherwise.

### GetTunneledKvmLaunchOk

`func (o *KvmTunneledKvmPolicy) GetTunneledKvmLaunchOk() (*bool, bool)`

GetTunneledKvmLaunchOk returns a tuple with the TunneledKvmLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvmLaunch

`func (o *KvmTunneledKvmPolicy) SetTunneledKvmLaunch(v bool)`

SetTunneledKvmLaunch sets TunneledKvmLaunch field to given value.

### HasTunneledKvmLaunch

`func (o *KvmTunneledKvmPolicy) HasTunneledKvmLaunch() bool`

HasTunneledKvmLaunch returns a boolean if a field has been set.

### GetAccount

`func (o *KvmTunneledKvmPolicy) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *KvmTunneledKvmPolicy) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *KvmTunneledKvmPolicy) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *KvmTunneledKvmPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *KvmTunneledKvmPolicy) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *KvmTunneledKvmPolicy) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


