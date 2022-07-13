# InfraMigGpuConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "infra.MigGpuConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "infra.MigGpuConfiguration"]
**MigProfileName** | Pointer to **string** | The predefined MIG profile name, e.g. 1g.5gb, 2g.10gb, etc. | [optional] 

## Methods

### NewInfraMigGpuConfigurationAllOf

`func NewInfraMigGpuConfigurationAllOf(classId string, objectType string, ) *InfraMigGpuConfigurationAllOf`

NewInfraMigGpuConfigurationAllOf instantiates a new InfraMigGpuConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraMigGpuConfigurationAllOfWithDefaults

`func NewInfraMigGpuConfigurationAllOfWithDefaults() *InfraMigGpuConfigurationAllOf`

NewInfraMigGpuConfigurationAllOfWithDefaults instantiates a new InfraMigGpuConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InfraMigGpuConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InfraMigGpuConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InfraMigGpuConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InfraMigGpuConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InfraMigGpuConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InfraMigGpuConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMigProfileName

`func (o *InfraMigGpuConfigurationAllOf) GetMigProfileName() string`

GetMigProfileName returns the MigProfileName field if non-nil, zero value otherwise.

### GetMigProfileNameOk

`func (o *InfraMigGpuConfigurationAllOf) GetMigProfileNameOk() (*string, bool)`

GetMigProfileNameOk returns a tuple with the MigProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigProfileName

`func (o *InfraMigGpuConfigurationAllOf) SetMigProfileName(v string)`

SetMigProfileName sets MigProfileName field to given value.

### HasMigProfileName

`func (o *InfraMigGpuConfigurationAllOf) HasMigProfileName() bool`

HasMigProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


