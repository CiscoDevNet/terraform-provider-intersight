# FabricLinkAggregationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.LinkAggregationPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.LinkAggregationPolicy"]
**LacpRate** | Pointer to **string** | Flag used to indicate whether LACP PDUs are to be sent &#39;fast&#39;, i.e., every 1 second. * &#x60;normal&#x60; - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * &#x60;fast&#x60; - The standard 4th generation UCS Fabric Interconnect with 54 ports. | [optional] [default to "normal"]
**SuspendIndividual** | Pointer to **bool** | Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricLinkAggregationPolicy

`func NewFabricLinkAggregationPolicy(classId string, objectType string, ) *FabricLinkAggregationPolicy`

NewFabricLinkAggregationPolicy instantiates a new FabricLinkAggregationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricLinkAggregationPolicyWithDefaults

`func NewFabricLinkAggregationPolicyWithDefaults() *FabricLinkAggregationPolicy`

NewFabricLinkAggregationPolicyWithDefaults instantiates a new FabricLinkAggregationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricLinkAggregationPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricLinkAggregationPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricLinkAggregationPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricLinkAggregationPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricLinkAggregationPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricLinkAggregationPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLacpRate

`func (o *FabricLinkAggregationPolicy) GetLacpRate() string`

GetLacpRate returns the LacpRate field if non-nil, zero value otherwise.

### GetLacpRateOk

`func (o *FabricLinkAggregationPolicy) GetLacpRateOk() (*string, bool)`

GetLacpRateOk returns a tuple with the LacpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpRate

`func (o *FabricLinkAggregationPolicy) SetLacpRate(v string)`

SetLacpRate sets LacpRate field to given value.

### HasLacpRate

`func (o *FabricLinkAggregationPolicy) HasLacpRate() bool`

HasLacpRate returns a boolean if a field has been set.

### GetSuspendIndividual

`func (o *FabricLinkAggregationPolicy) GetSuspendIndividual() bool`

GetSuspendIndividual returns the SuspendIndividual field if non-nil, zero value otherwise.

### GetSuspendIndividualOk

`func (o *FabricLinkAggregationPolicy) GetSuspendIndividualOk() (*bool, bool)`

GetSuspendIndividualOk returns a tuple with the SuspendIndividual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendIndividual

`func (o *FabricLinkAggregationPolicy) SetSuspendIndividual(v bool)`

SetSuspendIndividual sets SuspendIndividual field to given value.

### HasSuspendIndividual

`func (o *FabricLinkAggregationPolicy) HasSuspendIndividual() bool`

HasSuspendIndividual returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricLinkAggregationPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricLinkAggregationPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricLinkAggregationPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricLinkAggregationPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


