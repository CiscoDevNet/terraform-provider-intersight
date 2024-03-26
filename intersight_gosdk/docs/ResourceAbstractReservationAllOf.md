# ResourceAbstractReservationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "resourcepool.MembershipReservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "resourcepool.MembershipReservation"]
**Description** | Pointer to **string** | Details of the use case for which the reservation was created, such as decommissioning. | [optional] 
**Expiration** | Pointer to **time.Time** | The resource reservation includes an expiration date and a timestamp indicating when this management object will be cleared. The expiration date is set during the decommissioning process and is maintained for a period of 3 months. | [optional] 
**ReservationSelector** | Pointer to **string** | The unique identification of the resource is based on the resource OData string, which is mentioned as part of the ReservationSelector. For example, &#39;Serial eq &#39;EM6259AE6B&#39;. | [optional] 
**ResourceType** | Pointer to **string** | The type of resource that is placed into resource groups or pools. Resource Type can be either &#39;compute.Blade&#39; or &#39;compute.RackUnit for pools. | [optional] 
**Status** | Pointer to **string** | The reservation status can be in the &#39;Created&#39;, &#39;Processing&#39;, &#39;Failed&#39;, or &#39;Finished&#39; state. * &#x60;Created&#x60; - By default, a reservation is in Created status. * &#x60;Processing&#x60; - A reservation is changed to Processing status for appliance mode resource claim requests. * &#x60;Failed&#x60; - A reservation is changed to Failed status if the validations on resources, resource groups fails. * &#x60;Finished&#x60; - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters. | [optional] [readonly] [default to "Created"]
**Identity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewResourceAbstractReservationAllOf

`func NewResourceAbstractReservationAllOf(classId string, objectType string, ) *ResourceAbstractReservationAllOf`

NewResourceAbstractReservationAllOf instantiates a new ResourceAbstractReservationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAbstractReservationAllOfWithDefaults

`func NewResourceAbstractReservationAllOfWithDefaults() *ResourceAbstractReservationAllOf`

NewResourceAbstractReservationAllOfWithDefaults instantiates a new ResourceAbstractReservationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceAbstractReservationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceAbstractReservationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceAbstractReservationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceAbstractReservationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceAbstractReservationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceAbstractReservationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ResourceAbstractReservationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceAbstractReservationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceAbstractReservationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceAbstractReservationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *ResourceAbstractReservationAllOf) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ResourceAbstractReservationAllOf) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ResourceAbstractReservationAllOf) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ResourceAbstractReservationAllOf) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetReservationSelector

`func (o *ResourceAbstractReservationAllOf) GetReservationSelector() string`

GetReservationSelector returns the ReservationSelector field if non-nil, zero value otherwise.

### GetReservationSelectorOk

`func (o *ResourceAbstractReservationAllOf) GetReservationSelectorOk() (*string, bool)`

GetReservationSelectorOk returns a tuple with the ReservationSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationSelector

`func (o *ResourceAbstractReservationAllOf) SetReservationSelector(v string)`

SetReservationSelector sets ReservationSelector field to given value.

### HasReservationSelector

`func (o *ResourceAbstractReservationAllOf) HasReservationSelector() bool`

HasReservationSelector returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceAbstractReservationAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceAbstractReservationAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceAbstractReservationAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceAbstractReservationAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceAbstractReservationAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceAbstractReservationAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceAbstractReservationAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceAbstractReservationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIdentity

`func (o *ResourceAbstractReservationAllOf) GetIdentity() MoBaseMoRelationship`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ResourceAbstractReservationAllOf) GetIdentityOk() (*MoBaseMoRelationship, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ResourceAbstractReservationAllOf) SetIdentity(v MoBaseMoRelationship)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ResourceAbstractReservationAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


