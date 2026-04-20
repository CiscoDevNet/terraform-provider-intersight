# ComputeMigrationKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.MigrationKeyInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.MigrationKeyInfo"]
**BootLunType** | Pointer to **string** | Type of Boot LUN used to boot the operating system. * &#x60;SAN&#x60; - Refers to the SAN boot type. * &#x60;Local&#x60; - Refers to the Local boot type. | [optional] [readonly] [default to "SAN"]
**InitiatorAddresses** | Pointer to **[]string** |  | [optional] 
**InstanceId** | Pointer to **string** | Unique Identifier that represents the boot lun instance. | [optional] [readonly] 
**IsKeyValueSet** | Pointer to **bool** | Indicates whether the value of the &#39;keyValue&#39; property has been set. | [optional] [readonly] [default to false]
**LunId** | Pointer to **int64** | Identifies the LUN ID associated with the migration key. | [optional] [readonly] 
**TargetAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewComputeMigrationKeyInfo

`func NewComputeMigrationKeyInfo(classId string, objectType string, ) *ComputeMigrationKeyInfo`

NewComputeMigrationKeyInfo instantiates a new ComputeMigrationKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeMigrationKeyInfoWithDefaults

`func NewComputeMigrationKeyInfoWithDefaults() *ComputeMigrationKeyInfo`

NewComputeMigrationKeyInfoWithDefaults instantiates a new ComputeMigrationKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeMigrationKeyInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeMigrationKeyInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeMigrationKeyInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeMigrationKeyInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeMigrationKeyInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeMigrationKeyInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootLunType

`func (o *ComputeMigrationKeyInfo) GetBootLunType() string`

GetBootLunType returns the BootLunType field if non-nil, zero value otherwise.

### GetBootLunTypeOk

`func (o *ComputeMigrationKeyInfo) GetBootLunTypeOk() (*string, bool)`

GetBootLunTypeOk returns a tuple with the BootLunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootLunType

`func (o *ComputeMigrationKeyInfo) SetBootLunType(v string)`

SetBootLunType sets BootLunType field to given value.

### HasBootLunType

`func (o *ComputeMigrationKeyInfo) HasBootLunType() bool`

HasBootLunType returns a boolean if a field has been set.

### GetInitiatorAddresses

`func (o *ComputeMigrationKeyInfo) GetInitiatorAddresses() []string`

GetInitiatorAddresses returns the InitiatorAddresses field if non-nil, zero value otherwise.

### GetInitiatorAddressesOk

`func (o *ComputeMigrationKeyInfo) GetInitiatorAddressesOk() (*[]string, bool)`

GetInitiatorAddressesOk returns a tuple with the InitiatorAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorAddresses

`func (o *ComputeMigrationKeyInfo) SetInitiatorAddresses(v []string)`

SetInitiatorAddresses sets InitiatorAddresses field to given value.

### HasInitiatorAddresses

`func (o *ComputeMigrationKeyInfo) HasInitiatorAddresses() bool`

HasInitiatorAddresses returns a boolean if a field has been set.

### SetInitiatorAddressesNil

`func (o *ComputeMigrationKeyInfo) SetInitiatorAddressesNil(b bool)`

 SetInitiatorAddressesNil sets the value for InitiatorAddresses to be an explicit nil

### UnsetInitiatorAddresses
`func (o *ComputeMigrationKeyInfo) UnsetInitiatorAddresses()`

UnsetInitiatorAddresses ensures that no value is present for InitiatorAddresses, not even an explicit nil
### GetInstanceId

`func (o *ComputeMigrationKeyInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ComputeMigrationKeyInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ComputeMigrationKeyInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ComputeMigrationKeyInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetIsKeyValueSet

`func (o *ComputeMigrationKeyInfo) GetIsKeyValueSet() bool`

GetIsKeyValueSet returns the IsKeyValueSet field if non-nil, zero value otherwise.

### GetIsKeyValueSetOk

`func (o *ComputeMigrationKeyInfo) GetIsKeyValueSetOk() (*bool, bool)`

GetIsKeyValueSetOk returns a tuple with the IsKeyValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeyValueSet

`func (o *ComputeMigrationKeyInfo) SetIsKeyValueSet(v bool)`

SetIsKeyValueSet sets IsKeyValueSet field to given value.

### HasIsKeyValueSet

`func (o *ComputeMigrationKeyInfo) HasIsKeyValueSet() bool`

HasIsKeyValueSet returns a boolean if a field has been set.

### GetLunId

`func (o *ComputeMigrationKeyInfo) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *ComputeMigrationKeyInfo) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *ComputeMigrationKeyInfo) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *ComputeMigrationKeyInfo) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetTargetAddresses

`func (o *ComputeMigrationKeyInfo) GetTargetAddresses() []string`

GetTargetAddresses returns the TargetAddresses field if non-nil, zero value otherwise.

### GetTargetAddressesOk

`func (o *ComputeMigrationKeyInfo) GetTargetAddressesOk() (*[]string, bool)`

GetTargetAddressesOk returns a tuple with the TargetAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAddresses

`func (o *ComputeMigrationKeyInfo) SetTargetAddresses(v []string)`

SetTargetAddresses sets TargetAddresses field to given value.

### HasTargetAddresses

`func (o *ComputeMigrationKeyInfo) HasTargetAddresses() bool`

HasTargetAddresses returns a boolean if a field has been set.

### SetTargetAddressesNil

`func (o *ComputeMigrationKeyInfo) SetTargetAddressesNil(b bool)`

 SetTargetAddressesNil sets the value for TargetAddresses to be an explicit nil

### UnsetTargetAddresses
`func (o *ComputeMigrationKeyInfo) UnsetTargetAddresses()`

UnsetTargetAddresses ensures that no value is present for TargetAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


