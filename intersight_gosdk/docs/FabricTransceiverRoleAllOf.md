# FabricTransceiverRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). * &#x60;1Gbps&#x60; - Admin configurable speed 1Gbps. * &#x60;10Gbps&#x60; - Admin configurable speed 10Gbps. * &#x60;25Gbps&#x60; - Admin configurable speed 25Gbps. * &#x60;40Gbps&#x60; - Admin configurable speed 40Gbps. * &#x60;100Gbps&#x60; - Admin configurable speed 100Gbps. | [optional] [default to "Auto"]
**Fec** | Pointer to **string** | Forward error correction configuration for the port. * &#x60;Auto&#x60; - Forward error correction option &#39;Auto&#39;. * &#x60;Cl91&#x60; - Forward error correction option &#39;cl91&#39;. * &#x60;Cl74&#x60; - Forward error correction option &#39;cl74&#39;. | [optional] [default to "Auto"]

## Methods

### NewFabricTransceiverRoleAllOf

`func NewFabricTransceiverRoleAllOf(classId string, objectType string, ) *FabricTransceiverRoleAllOf`

NewFabricTransceiverRoleAllOf instantiates a new FabricTransceiverRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricTransceiverRoleAllOfWithDefaults

`func NewFabricTransceiverRoleAllOfWithDefaults() *FabricTransceiverRoleAllOf`

NewFabricTransceiverRoleAllOfWithDefaults instantiates a new FabricTransceiverRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricTransceiverRoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricTransceiverRoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricTransceiverRoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricTransceiverRoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricTransceiverRoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricTransceiverRoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricTransceiverRoleAllOf) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricTransceiverRoleAllOf) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricTransceiverRoleAllOf) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricTransceiverRoleAllOf) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFec

`func (o *FabricTransceiverRoleAllOf) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricTransceiverRoleAllOf) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricTransceiverRoleAllOf) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricTransceiverRoleAllOf) HasFec() bool`

HasFec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


