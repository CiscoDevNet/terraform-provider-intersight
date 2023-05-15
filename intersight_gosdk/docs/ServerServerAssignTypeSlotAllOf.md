# ServerServerAssignTypeSlotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.ServerAssignTypeSlot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.ServerAssignTypeSlot"]
**ChassisId** | Pointer to **int64** | Chassis-id of the slot that would be assigned to this pre-assigned server profile. | [optional] [default to 0]
**DomainName** | Pointer to **string** | Domain name of the Fabric Interconnect to which the chassis is or to be connected. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters. | [optional] 
**SlotId** | Pointer to **int64** | Slot-id of the server that would be assigned to this pre-assigned server profile. | [optional] [default to 0]

## Methods

### NewServerServerAssignTypeSlotAllOf

`func NewServerServerAssignTypeSlotAllOf(classId string, objectType string, ) *ServerServerAssignTypeSlotAllOf`

NewServerServerAssignTypeSlotAllOf instantiates a new ServerServerAssignTypeSlotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerServerAssignTypeSlotAllOfWithDefaults

`func NewServerServerAssignTypeSlotAllOfWithDefaults() *ServerServerAssignTypeSlotAllOf`

NewServerServerAssignTypeSlotAllOfWithDefaults instantiates a new ServerServerAssignTypeSlotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerServerAssignTypeSlotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerServerAssignTypeSlotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerServerAssignTypeSlotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerServerAssignTypeSlotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerServerAssignTypeSlotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerServerAssignTypeSlotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisId

`func (o *ServerServerAssignTypeSlotAllOf) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ServerServerAssignTypeSlotAllOf) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ServerServerAssignTypeSlotAllOf) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ServerServerAssignTypeSlotAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetDomainName

`func (o *ServerServerAssignTypeSlotAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ServerServerAssignTypeSlotAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ServerServerAssignTypeSlotAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ServerServerAssignTypeSlotAllOf) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetSlotId

`func (o *ServerServerAssignTypeSlotAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *ServerServerAssignTypeSlotAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *ServerServerAssignTypeSlotAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *ServerServerAssignTypeSlotAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


