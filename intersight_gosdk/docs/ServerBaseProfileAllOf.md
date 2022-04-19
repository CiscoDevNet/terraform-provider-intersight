# ServerBaseProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**UuidAddressType** | Pointer to **string** | UUID address allocation type selected to assign an UUID address for the server. * &#x60;NONE&#x60; - The user did not assign any UUID address. * &#x60;STATIC&#x60; - The user assigns a static UUID address. * &#x60;POOL&#x60; - The user selects a pool from which the address will be leased. | [optional] [default to "NONE"]
**ConfigResult** | Pointer to [**ServerConfigResultRelationship**](ServerConfigResultRelationship.md) |  | [optional] 
**UuidPool** | Pointer to [**UuidpoolPoolRelationship**](UuidpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewServerBaseProfileAllOf

`func NewServerBaseProfileAllOf(classId string, objectType string, ) *ServerBaseProfileAllOf`

NewServerBaseProfileAllOf instantiates a new ServerBaseProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBaseProfileAllOfWithDefaults

`func NewServerBaseProfileAllOfWithDefaults() *ServerBaseProfileAllOf`

NewServerBaseProfileAllOfWithDefaults instantiates a new ServerBaseProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerBaseProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerBaseProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerBaseProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerBaseProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerBaseProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerBaseProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetPlatform

`func (o *ServerBaseProfileAllOf) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *ServerBaseProfileAllOf) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *ServerBaseProfileAllOf) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *ServerBaseProfileAllOf) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetUuidAddressType

`func (o *ServerBaseProfileAllOf) GetUuidAddressType() string`

GetUuidAddressType returns the UuidAddressType field if non-nil, zero value otherwise.

### GetUuidAddressTypeOk

`func (o *ServerBaseProfileAllOf) GetUuidAddressTypeOk() (*string, bool)`

GetUuidAddressTypeOk returns a tuple with the UuidAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidAddressType

`func (o *ServerBaseProfileAllOf) SetUuidAddressType(v string)`

SetUuidAddressType sets UuidAddressType field to given value.

### HasUuidAddressType

`func (o *ServerBaseProfileAllOf) HasUuidAddressType() bool`

HasUuidAddressType returns a boolean if a field has been set.

### GetConfigResult

`func (o *ServerBaseProfileAllOf) GetConfigResult() ServerConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ServerBaseProfileAllOf) GetConfigResultOk() (*ServerConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ServerBaseProfileAllOf) SetConfigResult(v ServerConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ServerBaseProfileAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetUuidPool

`func (o *ServerBaseProfileAllOf) GetUuidPool() UuidpoolPoolRelationship`

GetUuidPool returns the UuidPool field if non-nil, zero value otherwise.

### GetUuidPoolOk

`func (o *ServerBaseProfileAllOf) GetUuidPoolOk() (*UuidpoolPoolRelationship, bool)`

GetUuidPoolOk returns a tuple with the UuidPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidPool

`func (o *ServerBaseProfileAllOf) SetUuidPool(v UuidpoolPoolRelationship)`

SetUuidPool sets UuidPool field to given value.

### HasUuidPool

`func (o *ServerBaseProfileAllOf) HasUuidPool() bool`

HasUuidPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


