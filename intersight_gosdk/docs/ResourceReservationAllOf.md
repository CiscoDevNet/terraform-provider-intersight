# ResourceReservationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.Reservation"]
**Expiration** | Pointer to **time.Time** | Expiration of the resource Reservation. | [optional] [readonly] 
**MarkFail** | Pointer to **bool** | MarkFail is used to set the reservation status to Failed. | [optional] 
**ResourceMoids** | Pointer to **[]string** |  | [optional] 
**ResourceType** | Pointer to **string** | Type of resources which will get filled into the resource groups. | [optional] 
**Status** | Pointer to **string** | Status of the Reservation. * &#x60;Created&#x60; - By default, a reservation is in Created status. * &#x60;Processing&#x60; - A reservation is changed to Processing status for appliance mode resource claim requests. * &#x60;Failed&#x60; - A reservation is changed to Failed status if the validations on resources, resource groups fails. * &#x60;Finished&#x60; - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters. | [optional] [readonly] [default to "Created"]
**UserMoid** | Pointer to **string** | Moid of the user who created the reservation. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Groups** | Pointer to [**[]ResourceGroupRelationship**](ResourceGroupRelationship.md) | An array of relationships to resourceGroup resources. | [optional] 

## Methods

### NewResourceReservationAllOf

`func NewResourceReservationAllOf(classId string, objectType string, ) *ResourceReservationAllOf`

NewResourceReservationAllOf instantiates a new ResourceReservationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReservationAllOfWithDefaults

`func NewResourceReservationAllOfWithDefaults() *ResourceReservationAllOf`

NewResourceReservationAllOfWithDefaults instantiates a new ResourceReservationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceReservationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceReservationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceReservationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceReservationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceReservationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceReservationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpiration

`func (o *ResourceReservationAllOf) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ResourceReservationAllOf) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ResourceReservationAllOf) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ResourceReservationAllOf) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetMarkFail

`func (o *ResourceReservationAllOf) GetMarkFail() bool`

GetMarkFail returns the MarkFail field if non-nil, zero value otherwise.

### GetMarkFailOk

`func (o *ResourceReservationAllOf) GetMarkFailOk() (*bool, bool)`

GetMarkFailOk returns a tuple with the MarkFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkFail

`func (o *ResourceReservationAllOf) SetMarkFail(v bool)`

SetMarkFail sets MarkFail field to given value.

### HasMarkFail

`func (o *ResourceReservationAllOf) HasMarkFail() bool`

HasMarkFail returns a boolean if a field has been set.

### GetResourceMoids

`func (o *ResourceReservationAllOf) GetResourceMoids() []string`

GetResourceMoids returns the ResourceMoids field if non-nil, zero value otherwise.

### GetResourceMoidsOk

`func (o *ResourceReservationAllOf) GetResourceMoidsOk() (*[]string, bool)`

GetResourceMoidsOk returns a tuple with the ResourceMoids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMoids

`func (o *ResourceReservationAllOf) SetResourceMoids(v []string)`

SetResourceMoids sets ResourceMoids field to given value.

### HasResourceMoids

`func (o *ResourceReservationAllOf) HasResourceMoids() bool`

HasResourceMoids returns a boolean if a field has been set.

### SetResourceMoidsNil

`func (o *ResourceReservationAllOf) SetResourceMoidsNil(b bool)`

 SetResourceMoidsNil sets the value for ResourceMoids to be an explicit nil

### UnsetResourceMoids
`func (o *ResourceReservationAllOf) UnsetResourceMoids()`

UnsetResourceMoids ensures that no value is present for ResourceMoids, not even an explicit nil
### GetResourceType

`func (o *ResourceReservationAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceReservationAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceReservationAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceReservationAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceReservationAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceReservationAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceReservationAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceReservationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserMoid

`func (o *ResourceReservationAllOf) GetUserMoid() string`

GetUserMoid returns the UserMoid field if non-nil, zero value otherwise.

### GetUserMoidOk

`func (o *ResourceReservationAllOf) GetUserMoidOk() (*string, bool)`

GetUserMoidOk returns a tuple with the UserMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMoid

`func (o *ResourceReservationAllOf) SetUserMoid(v string)`

SetUserMoid sets UserMoid field to given value.

### HasUserMoid

`func (o *ResourceReservationAllOf) HasUserMoid() bool`

HasUserMoid returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceReservationAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceReservationAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceReservationAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceReservationAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetGroups

`func (o *ResourceReservationAllOf) GetGroups() []ResourceGroupRelationship`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ResourceReservationAllOf) GetGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ResourceReservationAllOf) SetGroups(v []ResourceGroupRelationship)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ResourceReservationAllOf) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *ResourceReservationAllOf) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ResourceReservationAllOf) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


