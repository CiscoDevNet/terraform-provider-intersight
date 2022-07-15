# KubernetesNvidiaGpuProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NvidiaGpuProduct"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NvidiaGpuProduct"]
**MigCapable** | Pointer to **bool** | True if this Nvidia GPU supports MIG. | [optional] 
**MigProfiles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKubernetesNvidiaGpuProduct

`func NewKubernetesNvidiaGpuProduct(classId string, objectType string, ) *KubernetesNvidiaGpuProduct`

NewKubernetesNvidiaGpuProduct instantiates a new KubernetesNvidiaGpuProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNvidiaGpuProductWithDefaults

`func NewKubernetesNvidiaGpuProductWithDefaults() *KubernetesNvidiaGpuProduct`

NewKubernetesNvidiaGpuProductWithDefaults instantiates a new KubernetesNvidiaGpuProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNvidiaGpuProduct) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNvidiaGpuProduct) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNvidiaGpuProduct) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNvidiaGpuProduct) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNvidiaGpuProduct) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNvidiaGpuProduct) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMigCapable

`func (o *KubernetesNvidiaGpuProduct) GetMigCapable() bool`

GetMigCapable returns the MigCapable field if non-nil, zero value otherwise.

### GetMigCapableOk

`func (o *KubernetesNvidiaGpuProduct) GetMigCapableOk() (*bool, bool)`

GetMigCapableOk returns a tuple with the MigCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigCapable

`func (o *KubernetesNvidiaGpuProduct) SetMigCapable(v bool)`

SetMigCapable sets MigCapable field to given value.

### HasMigCapable

`func (o *KubernetesNvidiaGpuProduct) HasMigCapable() bool`

HasMigCapable returns a boolean if a field has been set.

### GetMigProfiles

`func (o *KubernetesNvidiaGpuProduct) GetMigProfiles() []string`

GetMigProfiles returns the MigProfiles field if non-nil, zero value otherwise.

### GetMigProfilesOk

`func (o *KubernetesNvidiaGpuProduct) GetMigProfilesOk() (*[]string, bool)`

GetMigProfilesOk returns a tuple with the MigProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigProfiles

`func (o *KubernetesNvidiaGpuProduct) SetMigProfiles(v []string)`

SetMigProfiles sets MigProfiles field to given value.

### HasMigProfiles

`func (o *KubernetesNvidiaGpuProduct) HasMigProfiles() bool`

HasMigProfiles returns a boolean if a field has been set.

### SetMigProfilesNil

`func (o *KubernetesNvidiaGpuProduct) SetMigProfilesNil(b bool)`

 SetMigProfilesNil sets the value for MigProfiles to be an explicit nil

### UnsetMigProfiles
`func (o *KubernetesNvidiaGpuProduct) UnsetMigProfiles()`

UnsetMigProfiles ensures that no value is present for MigProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


