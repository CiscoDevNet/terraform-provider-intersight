# StorageNetAppBaseEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Cause** | Pointer to **string** | A message describing the cause for the event. | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | Unique identifier of the cluster across the datacenter. | [optional] [readonly] 
**CurrentState** | Pointer to **string** | The current state of the event. * &#x60;unknown&#x60; - Default unknown current state. * &#x60;new&#x60; - The current state of the event is new. * &#x60;acknowledged&#x60; - The current state of the event is acknowledged. * &#x60;resolved&#x60; - The current state of the event is resolved. * &#x60;obsolete&#x60; - The current state of the event is obsolete. | [optional] [readonly] [default to "unknown"]
**Duration** | Pointer to **string** | Time since the event was created, in ISO8601 standard. | [optional] [readonly] 
**ImpactArea** | Pointer to **string** | Impact area of the event (availability, capacity, configuration, performance, protection, or security). * &#x60;unknown&#x60; - Default unknown impact area. * &#x60;availability&#x60; - The impact area of the event is availability. * &#x60;capacity&#x60; - The impact area of the event is capacity. * &#x60;configuration&#x60; - The impact area of the event is configuration. * &#x60;performance&#x60; - The impact area of the event is performance. * &#x60;protection&#x60; - The impact area of the event is protection. * &#x60;security&#x60; - The impact area of the event is security. | [optional] [readonly] [default to "unknown"]
**ImpactLevel** | Pointer to **string** | Impact level of the event (event, risk, incident, or upgrade). * &#x60;unknown&#x60; - Default unknown impact level. * &#x60;event&#x60; - The impact level of the event is event. * &#x60;risk&#x60; - The impact level of the event is risk. * &#x60;incident&#x60; - The impact level of the event is incident. * &#x60;upgrade&#x60; - The impact level of the event is upgrade. | [optional] [readonly] [default to "unknown"]
**ImpactResourceName** | Pointer to **string** | The full name of the source of the event. | [optional] [readonly] 
**ImpactResourceType** | Pointer to **string** | Type of resource with which the event is associated. * &#x60;unknown&#x60; - Default unknown resource type. * &#x60;aggregate&#x60; - The type of resource impacted by the event is an aggregate. * &#x60;cluster&#x60; - The type of resource impacted by the event is a cluster. * &#x60;cluster_node&#x60; - The type of resource impacted by the event is a node. * &#x60;disk&#x60; - The type of resource impacted by the event is a disk. * &#x60;fcp_lif&#x60; - The type of resource impacted by the event is a FC interface. * &#x60;fcp_port&#x60; - The type of resource impacted by the event is a FC port. * &#x60;lun&#x60; - The type of resource impacted by the event is a lun. * &#x60;network_lif&#x60; - The type of resource impacted by the event is an ethernet interface. * &#x60;network_port&#x60; - The type of resource impacted by the event is an ethernet port. * &#x60;volume&#x60; - The type of resource impacted by the event is a volume. * &#x60;vserver&#x60; - The type of resource impacted by the event is a storage VM. | [optional] [readonly] [default to "unknown"]
**ImpactResourceUuid** | Pointer to **string** | The unique identifier of the impacted resource. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the event that occurred. | [optional] [readonly] 
**NodeUuid** | Pointer to **string** | Unique identifier of the node across the cluster. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the event. * &#x60;unknown&#x60; - Default unknown severity. * &#x60;normal&#x60; - The severity of the event is normal. * &#x60;information&#x60; - The severity of the event is information. * &#x60;warning&#x60; - The severity of the event is warning. * &#x60;error&#x60; - The severity of the event is error. * &#x60;critical&#x60; - The severity of the event is critical. | [optional] [readonly] [default to "unknown"]
**SvmUuid** | Pointer to **string** | Unique identifier of the storage VM. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Unique identifier of the event. | [optional] [readonly] 

## Methods

### NewStorageNetAppBaseEventAllOf

`func NewStorageNetAppBaseEventAllOf(classId string, objectType string, ) *StorageNetAppBaseEventAllOf`

NewStorageNetAppBaseEventAllOf instantiates a new StorageNetAppBaseEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppBaseEventAllOfWithDefaults

`func NewStorageNetAppBaseEventAllOfWithDefaults() *StorageNetAppBaseEventAllOf`

NewStorageNetAppBaseEventAllOfWithDefaults instantiates a new StorageNetAppBaseEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppBaseEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppBaseEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppBaseEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppBaseEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppBaseEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppBaseEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCause

`func (o *StorageNetAppBaseEventAllOf) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *StorageNetAppBaseEventAllOf) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *StorageNetAppBaseEventAllOf) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *StorageNetAppBaseEventAllOf) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetClusterUuid

`func (o *StorageNetAppBaseEventAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *StorageNetAppBaseEventAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *StorageNetAppBaseEventAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *StorageNetAppBaseEventAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetCurrentState

`func (o *StorageNetAppBaseEventAllOf) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *StorageNetAppBaseEventAllOf) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *StorageNetAppBaseEventAllOf) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *StorageNetAppBaseEventAllOf) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetDuration

`func (o *StorageNetAppBaseEventAllOf) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *StorageNetAppBaseEventAllOf) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *StorageNetAppBaseEventAllOf) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *StorageNetAppBaseEventAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetImpactArea

`func (o *StorageNetAppBaseEventAllOf) GetImpactArea() string`

GetImpactArea returns the ImpactArea field if non-nil, zero value otherwise.

### GetImpactAreaOk

`func (o *StorageNetAppBaseEventAllOf) GetImpactAreaOk() (*string, bool)`

GetImpactAreaOk returns a tuple with the ImpactArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactArea

`func (o *StorageNetAppBaseEventAllOf) SetImpactArea(v string)`

SetImpactArea sets ImpactArea field to given value.

### HasImpactArea

`func (o *StorageNetAppBaseEventAllOf) HasImpactArea() bool`

HasImpactArea returns a boolean if a field has been set.

### GetImpactLevel

`func (o *StorageNetAppBaseEventAllOf) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *StorageNetAppBaseEventAllOf) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *StorageNetAppBaseEventAllOf) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *StorageNetAppBaseEventAllOf) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetImpactResourceName

`func (o *StorageNetAppBaseEventAllOf) GetImpactResourceName() string`

GetImpactResourceName returns the ImpactResourceName field if non-nil, zero value otherwise.

### GetImpactResourceNameOk

`func (o *StorageNetAppBaseEventAllOf) GetImpactResourceNameOk() (*string, bool)`

GetImpactResourceNameOk returns a tuple with the ImpactResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactResourceName

`func (o *StorageNetAppBaseEventAllOf) SetImpactResourceName(v string)`

SetImpactResourceName sets ImpactResourceName field to given value.

### HasImpactResourceName

`func (o *StorageNetAppBaseEventAllOf) HasImpactResourceName() bool`

HasImpactResourceName returns a boolean if a field has been set.

### GetImpactResourceType

`func (o *StorageNetAppBaseEventAllOf) GetImpactResourceType() string`

GetImpactResourceType returns the ImpactResourceType field if non-nil, zero value otherwise.

### GetImpactResourceTypeOk

`func (o *StorageNetAppBaseEventAllOf) GetImpactResourceTypeOk() (*string, bool)`

GetImpactResourceTypeOk returns a tuple with the ImpactResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactResourceType

`func (o *StorageNetAppBaseEventAllOf) SetImpactResourceType(v string)`

SetImpactResourceType sets ImpactResourceType field to given value.

### HasImpactResourceType

`func (o *StorageNetAppBaseEventAllOf) HasImpactResourceType() bool`

HasImpactResourceType returns a boolean if a field has been set.

### GetImpactResourceUuid

`func (o *StorageNetAppBaseEventAllOf) GetImpactResourceUuid() string`

GetImpactResourceUuid returns the ImpactResourceUuid field if non-nil, zero value otherwise.

### GetImpactResourceUuidOk

`func (o *StorageNetAppBaseEventAllOf) GetImpactResourceUuidOk() (*string, bool)`

GetImpactResourceUuidOk returns a tuple with the ImpactResourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactResourceUuid

`func (o *StorageNetAppBaseEventAllOf) SetImpactResourceUuid(v string)`

SetImpactResourceUuid sets ImpactResourceUuid field to given value.

### HasImpactResourceUuid

`func (o *StorageNetAppBaseEventAllOf) HasImpactResourceUuid() bool`

HasImpactResourceUuid returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppBaseEventAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppBaseEventAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppBaseEventAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppBaseEventAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeUuid

`func (o *StorageNetAppBaseEventAllOf) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *StorageNetAppBaseEventAllOf) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *StorageNetAppBaseEventAllOf) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *StorageNetAppBaseEventAllOf) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetSeverity

`func (o *StorageNetAppBaseEventAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StorageNetAppBaseEventAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StorageNetAppBaseEventAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *StorageNetAppBaseEventAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppBaseEventAllOf) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppBaseEventAllOf) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppBaseEventAllOf) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppBaseEventAllOf) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppBaseEventAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppBaseEventAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppBaseEventAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppBaseEventAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


