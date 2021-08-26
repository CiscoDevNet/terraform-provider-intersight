# FabricLinkAggregationPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.LinkAggregationPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.LinkAggregationPolicy"]
**LacpRate** | Pointer to **string** | Flag used to indicate whether LACP PDUs are to be sent &#39;fast&#39;, i.e., every 1 second. * &#x60;normal&#x60; - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * &#x60;fast&#x60; - The standard 4th generation UCS Fabric Interconnect with 54 ports. | [optional] [default to "normal"]
**SuspendIndividual** | Pointer to **bool** | Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricLinkAggregationPolicyAllOf

`func NewFabricLinkAggregationPolicyAllOf(classId string, objectType string, ) *FabricLinkAggregationPolicyAllOf`

NewFabricLinkAggregationPolicyAllOf instantiates a new FabricLinkAggregationPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricLinkAggregationPolicyAllOfWithDefaults

`func NewFabricLinkAggregationPolicyAllOfWithDefaults() *FabricLinkAggregationPolicyAllOf`

NewFabricLinkAggregationPolicyAllOfWithDefaults instantiates a new FabricLinkAggregationPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricLinkAggregationPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricLinkAggregationPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricLinkAggregationPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricLinkAggregationPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricLinkAggregationPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricLinkAggregationPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLacpRate

`func (o *FabricLinkAggregationPolicyAllOf) GetLacpRate() string`

GetLacpRate returns the LacpRate field if non-nil, zero value otherwise.

### GetLacpRateOk

`func (o *FabricLinkAggregationPolicyAllOf) GetLacpRateOk() (*string, bool)`

GetLacpRateOk returns a tuple with the LacpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpRate

`func (o *FabricLinkAggregationPolicyAllOf) SetLacpRate(v string)`

SetLacpRate sets LacpRate field to given value.

### HasLacpRate

`func (o *FabricLinkAggregationPolicyAllOf) HasLacpRate() bool`

HasLacpRate returns a boolean if a field has been set.

### GetSuspendIndividual

`func (o *FabricLinkAggregationPolicyAllOf) GetSuspendIndividual() bool`

GetSuspendIndividual returns the SuspendIndividual field if non-nil, zero value otherwise.

### GetSuspendIndividualOk

`func (o *FabricLinkAggregationPolicyAllOf) GetSuspendIndividualOk() (*bool, bool)`

GetSuspendIndividualOk returns a tuple with the SuspendIndividual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendIndividual

`func (o *FabricLinkAggregationPolicyAllOf) SetSuspendIndividual(v bool)`

SetSuspendIndividual sets SuspendIndividual field to given value.

### HasSuspendIndividual

`func (o *FabricLinkAggregationPolicyAllOf) HasSuspendIndividual() bool`

HasSuspendIndividual returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricLinkAggregationPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricLinkAggregationPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricLinkAggregationPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricLinkAggregationPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


