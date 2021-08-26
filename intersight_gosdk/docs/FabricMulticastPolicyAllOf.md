# FabricMulticastPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.MulticastPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.MulticastPolicy"]
**QuerierIpAddress** | Pointer to **string** | Used to define the IGMP Querier IP address. | [optional] 
**QuerierIpAddressPeer** | Pointer to **string** | Used to define the IGMP Querier IP address of the peer switch. | [optional] 
**QuerierState** | Pointer to **string** | Administrative state of the IGMP Querier for this VLAN. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**SnoopingState** | Pointer to **string** | Administrative state of the IGMP Snooping for this VLAN. * &#x60;Enabled&#x60; - Admin configured Enabled State. * &#x60;Disabled&#x60; - Admin configured Disabled State. | [optional] [default to "Enabled"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricMulticastPolicyAllOf

`func NewFabricMulticastPolicyAllOf(classId string, objectType string, ) *FabricMulticastPolicyAllOf`

NewFabricMulticastPolicyAllOf instantiates a new FabricMulticastPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricMulticastPolicyAllOfWithDefaults

`func NewFabricMulticastPolicyAllOfWithDefaults() *FabricMulticastPolicyAllOf`

NewFabricMulticastPolicyAllOfWithDefaults instantiates a new FabricMulticastPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricMulticastPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricMulticastPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricMulticastPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricMulticastPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricMulticastPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricMulticastPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetQuerierIpAddress

`func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddress() string`

GetQuerierIpAddress returns the QuerierIpAddress field if non-nil, zero value otherwise.

### GetQuerierIpAddressOk

`func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddressOk() (*string, bool)`

GetQuerierIpAddressOk returns a tuple with the QuerierIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierIpAddress

`func (o *FabricMulticastPolicyAllOf) SetQuerierIpAddress(v string)`

SetQuerierIpAddress sets QuerierIpAddress field to given value.

### HasQuerierIpAddress

`func (o *FabricMulticastPolicyAllOf) HasQuerierIpAddress() bool`

HasQuerierIpAddress returns a boolean if a field has been set.

### GetQuerierIpAddressPeer

`func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddressPeer() string`

GetQuerierIpAddressPeer returns the QuerierIpAddressPeer field if non-nil, zero value otherwise.

### GetQuerierIpAddressPeerOk

`func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddressPeerOk() (*string, bool)`

GetQuerierIpAddressPeerOk returns a tuple with the QuerierIpAddressPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierIpAddressPeer

`func (o *FabricMulticastPolicyAllOf) SetQuerierIpAddressPeer(v string)`

SetQuerierIpAddressPeer sets QuerierIpAddressPeer field to given value.

### HasQuerierIpAddressPeer

`func (o *FabricMulticastPolicyAllOf) HasQuerierIpAddressPeer() bool`

HasQuerierIpAddressPeer returns a boolean if a field has been set.

### GetQuerierState

`func (o *FabricMulticastPolicyAllOf) GetQuerierState() string`

GetQuerierState returns the QuerierState field if non-nil, zero value otherwise.

### GetQuerierStateOk

`func (o *FabricMulticastPolicyAllOf) GetQuerierStateOk() (*string, bool)`

GetQuerierStateOk returns a tuple with the QuerierState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierState

`func (o *FabricMulticastPolicyAllOf) SetQuerierState(v string)`

SetQuerierState sets QuerierState field to given value.

### HasQuerierState

`func (o *FabricMulticastPolicyAllOf) HasQuerierState() bool`

HasQuerierState returns a boolean if a field has been set.

### GetSnoopingState

`func (o *FabricMulticastPolicyAllOf) GetSnoopingState() string`

GetSnoopingState returns the SnoopingState field if non-nil, zero value otherwise.

### GetSnoopingStateOk

`func (o *FabricMulticastPolicyAllOf) GetSnoopingStateOk() (*string, bool)`

GetSnoopingStateOk returns a tuple with the SnoopingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopingState

`func (o *FabricMulticastPolicyAllOf) SetSnoopingState(v string)`

SetSnoopingState sets SnoopingState field to given value.

### HasSnoopingState

`func (o *FabricMulticastPolicyAllOf) HasSnoopingState() bool`

HasSnoopingState returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricMulticastPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricMulticastPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricMulticastPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricMulticastPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


