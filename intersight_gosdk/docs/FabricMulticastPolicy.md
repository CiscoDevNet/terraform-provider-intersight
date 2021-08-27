# FabricMulticastPolicy

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

### NewFabricMulticastPolicy

`func NewFabricMulticastPolicy(classId string, objectType string, ) *FabricMulticastPolicy`

NewFabricMulticastPolicy instantiates a new FabricMulticastPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricMulticastPolicyWithDefaults

`func NewFabricMulticastPolicyWithDefaults() *FabricMulticastPolicy`

NewFabricMulticastPolicyWithDefaults instantiates a new FabricMulticastPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricMulticastPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricMulticastPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricMulticastPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricMulticastPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricMulticastPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricMulticastPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetQuerierIpAddress

`func (o *FabricMulticastPolicy) GetQuerierIpAddress() string`

GetQuerierIpAddress returns the QuerierIpAddress field if non-nil, zero value otherwise.

### GetQuerierIpAddressOk

`func (o *FabricMulticastPolicy) GetQuerierIpAddressOk() (*string, bool)`

GetQuerierIpAddressOk returns a tuple with the QuerierIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierIpAddress

`func (o *FabricMulticastPolicy) SetQuerierIpAddress(v string)`

SetQuerierIpAddress sets QuerierIpAddress field to given value.

### HasQuerierIpAddress

`func (o *FabricMulticastPolicy) HasQuerierIpAddress() bool`

HasQuerierIpAddress returns a boolean if a field has been set.

### GetQuerierIpAddressPeer

`func (o *FabricMulticastPolicy) GetQuerierIpAddressPeer() string`

GetQuerierIpAddressPeer returns the QuerierIpAddressPeer field if non-nil, zero value otherwise.

### GetQuerierIpAddressPeerOk

`func (o *FabricMulticastPolicy) GetQuerierIpAddressPeerOk() (*string, bool)`

GetQuerierIpAddressPeerOk returns a tuple with the QuerierIpAddressPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierIpAddressPeer

`func (o *FabricMulticastPolicy) SetQuerierIpAddressPeer(v string)`

SetQuerierIpAddressPeer sets QuerierIpAddressPeer field to given value.

### HasQuerierIpAddressPeer

`func (o *FabricMulticastPolicy) HasQuerierIpAddressPeer() bool`

HasQuerierIpAddressPeer returns a boolean if a field has been set.

### GetQuerierState

`func (o *FabricMulticastPolicy) GetQuerierState() string`

GetQuerierState returns the QuerierState field if non-nil, zero value otherwise.

### GetQuerierStateOk

`func (o *FabricMulticastPolicy) GetQuerierStateOk() (*string, bool)`

GetQuerierStateOk returns a tuple with the QuerierState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierState

`func (o *FabricMulticastPolicy) SetQuerierState(v string)`

SetQuerierState sets QuerierState field to given value.

### HasQuerierState

`func (o *FabricMulticastPolicy) HasQuerierState() bool`

HasQuerierState returns a boolean if a field has been set.

### GetSnoopingState

`func (o *FabricMulticastPolicy) GetSnoopingState() string`

GetSnoopingState returns the SnoopingState field if non-nil, zero value otherwise.

### GetSnoopingStateOk

`func (o *FabricMulticastPolicy) GetSnoopingStateOk() (*string, bool)`

GetSnoopingStateOk returns a tuple with the SnoopingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopingState

`func (o *FabricMulticastPolicy) SetSnoopingState(v string)`

SetSnoopingState sets SnoopingState field to given value.

### HasSnoopingState

`func (o *FabricMulticastPolicy) HasSnoopingState() bool`

HasSnoopingState returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricMulticastPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricMulticastPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricMulticastPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricMulticastPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


