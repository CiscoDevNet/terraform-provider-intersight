# AssetScopedTargetConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ScopedTargetConnection"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ScopedTargetConnection"]
**FullValidation** | Pointer to **bool** | When this flag is set to true, every IWO entity in the scope targets will be checked and discovery of the scope target will be regarded as a failure when anyone of these entities cannot be connected and validated. | [optional] 
**Port** | Pointer to **int64** | The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection. | [optional] 
**Scope** | Pointer to **string** | The group id of IWO entities upon which the discover of a scoped target is performed. For example, a database target may be scoped to the group of virtual machines upon which the database application is running. Scope value is group id created for all those virtual machines in this scope. | [optional] 

## Methods

### NewAssetScopedTargetConnection

`func NewAssetScopedTargetConnection(classId string, objectType string, ) *AssetScopedTargetConnection`

NewAssetScopedTargetConnection instantiates a new AssetScopedTargetConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetScopedTargetConnectionWithDefaults

`func NewAssetScopedTargetConnectionWithDefaults() *AssetScopedTargetConnection`

NewAssetScopedTargetConnectionWithDefaults instantiates a new AssetScopedTargetConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetScopedTargetConnection) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetScopedTargetConnection) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetScopedTargetConnection) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetScopedTargetConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetScopedTargetConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetScopedTargetConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFullValidation

`func (o *AssetScopedTargetConnection) GetFullValidation() bool`

GetFullValidation returns the FullValidation field if non-nil, zero value otherwise.

### GetFullValidationOk

`func (o *AssetScopedTargetConnection) GetFullValidationOk() (*bool, bool)`

GetFullValidationOk returns a tuple with the FullValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullValidation

`func (o *AssetScopedTargetConnection) SetFullValidation(v bool)`

SetFullValidation sets FullValidation field to given value.

### HasFullValidation

`func (o *AssetScopedTargetConnection) HasFullValidation() bool`

HasFullValidation returns a boolean if a field has been set.

### GetPort

`func (o *AssetScopedTargetConnection) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssetScopedTargetConnection) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssetScopedTargetConnection) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AssetScopedTargetConnection) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetScope

`func (o *AssetScopedTargetConnection) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AssetScopedTargetConnection) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AssetScopedTargetConnection) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AssetScopedTargetConnection) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


