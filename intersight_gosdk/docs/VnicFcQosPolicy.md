# VnicFcQosPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcQosPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcQosPolicy"]
**Burst** | Pointer to **int64** | The burst traffic, in bytes, allowed on the vNIC. | [optional] [default to 1024]
**Cos** | Pointer to **int64** | Class of Service to be associated to the traffic on the virtual interface. | [optional] [default to 3]
**MaxDataFieldSize** | Pointer to **int64** | The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports. | [optional] [default to 2112]
**Priority** | Pointer to **string** | The priortity matching the System QoS specified in the fabric profile. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [readonly] [default to "Best Effort"]
**RateLimit** | Pointer to **int64** | The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. | [optional] [default to 0]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicFcQosPolicy

`func NewVnicFcQosPolicy(classId string, objectType string, ) *VnicFcQosPolicy`

NewVnicFcQosPolicy instantiates a new VnicFcQosPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcQosPolicyWithDefaults

`func NewVnicFcQosPolicyWithDefaults() *VnicFcQosPolicy`

NewVnicFcQosPolicyWithDefaults instantiates a new VnicFcQosPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcQosPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcQosPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcQosPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcQosPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcQosPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcQosPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBurst

`func (o *VnicFcQosPolicy) GetBurst() int64`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *VnicFcQosPolicy) GetBurstOk() (*int64, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *VnicFcQosPolicy) SetBurst(v int64)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *VnicFcQosPolicy) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetCos

`func (o *VnicFcQosPolicy) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicFcQosPolicy) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicFcQosPolicy) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicFcQosPolicy) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMaxDataFieldSize

`func (o *VnicFcQosPolicy) GetMaxDataFieldSize() int64`

GetMaxDataFieldSize returns the MaxDataFieldSize field if non-nil, zero value otherwise.

### GetMaxDataFieldSizeOk

`func (o *VnicFcQosPolicy) GetMaxDataFieldSizeOk() (*int64, bool)`

GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataFieldSize

`func (o *VnicFcQosPolicy) SetMaxDataFieldSize(v int64)`

SetMaxDataFieldSize sets MaxDataFieldSize field to given value.

### HasMaxDataFieldSize

`func (o *VnicFcQosPolicy) HasMaxDataFieldSize() bool`

HasMaxDataFieldSize returns a boolean if a field has been set.

### GetPriority

`func (o *VnicFcQosPolicy) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VnicFcQosPolicy) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VnicFcQosPolicy) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VnicFcQosPolicy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRateLimit

`func (o *VnicFcQosPolicy) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VnicFcQosPolicy) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VnicFcQosPolicy) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VnicFcQosPolicy) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicFcQosPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicFcQosPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicFcQosPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicFcQosPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


