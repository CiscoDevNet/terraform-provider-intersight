# HyperflexHxapEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapEvent"]
**FirstTime** | Pointer to **string** | First timestamp of the event occurrence. | [optional] 
**Identity** | Pointer to **string** | Internally generated identity (UUID) of this event. | [optional] 
**LastTime** | Pointer to **string** | Last timestamp of the event occurrence. | [optional] 
**Message** | Pointer to **string** | Full description of the event. | [optional] 
**OwnerName** | Pointer to **string** | Name of the owner with which event is associated. | [optional] 
**OwnerType** | Pointer to **string** | Type of the object with which event is associated (Host, Cluster, VM). * &#x60;Unknown&#x60; - Value is Unknown if the target type is unidentified. * &#x60;Cluster&#x60; - Cluster refers to HyperFlex AP Cluster. * &#x60;Host&#x60; - Host refers to server node which is part of HyperFlex AP Cluster. * &#x60;VM&#x60; - VM refers to Virtual machine available on a HyperFlex AP Node. * &#x60;Disk&#x60; - Disk refers to Virtual Disk available on a HyperFlex AP Cluster. | [optional] [default to "Unknown"]
**OwnerUuid** | Pointer to **string** | UUID of the owner with which event is associated. | [optional] 
**Severity** | Pointer to **string** | Severity level of the event (Info/Warning/Critical). * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [default to "None"]
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](hyperflex.HxapCluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapEventAllOf

`func NewHyperflexHxapEventAllOf(classId string, objectType string, ) *HyperflexHxapEventAllOf`

NewHyperflexHxapEventAllOf instantiates a new HyperflexHxapEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapEventAllOfWithDefaults

`func NewHyperflexHxapEventAllOfWithDefaults() *HyperflexHxapEventAllOf`

NewHyperflexHxapEventAllOfWithDefaults instantiates a new HyperflexHxapEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFirstTime

`func (o *HyperflexHxapEventAllOf) GetFirstTime() string`

GetFirstTime returns the FirstTime field if non-nil, zero value otherwise.

### GetFirstTimeOk

`func (o *HyperflexHxapEventAllOf) GetFirstTimeOk() (*string, bool)`

GetFirstTimeOk returns a tuple with the FirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTime

`func (o *HyperflexHxapEventAllOf) SetFirstTime(v string)`

SetFirstTime sets FirstTime field to given value.

### HasFirstTime

`func (o *HyperflexHxapEventAllOf) HasFirstTime() bool`

HasFirstTime returns a boolean if a field has been set.

### GetIdentity

`func (o *HyperflexHxapEventAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HyperflexHxapEventAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HyperflexHxapEventAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HyperflexHxapEventAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLastTime

`func (o *HyperflexHxapEventAllOf) GetLastTime() string`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *HyperflexHxapEventAllOf) GetLastTimeOk() (*string, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *HyperflexHxapEventAllOf) SetLastTime(v string)`

SetLastTime sets LastTime field to given value.

### HasLastTime

`func (o *HyperflexHxapEventAllOf) HasLastTime() bool`

HasLastTime returns a boolean if a field has been set.

### GetMessage

`func (o *HyperflexHxapEventAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HyperflexHxapEventAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HyperflexHxapEventAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HyperflexHxapEventAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOwnerName

`func (o *HyperflexHxapEventAllOf) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *HyperflexHxapEventAllOf) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *HyperflexHxapEventAllOf) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *HyperflexHxapEventAllOf) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerType

`func (o *HyperflexHxapEventAllOf) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *HyperflexHxapEventAllOf) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *HyperflexHxapEventAllOf) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *HyperflexHxapEventAllOf) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetOwnerUuid

`func (o *HyperflexHxapEventAllOf) GetOwnerUuid() string`

GetOwnerUuid returns the OwnerUuid field if non-nil, zero value otherwise.

### GetOwnerUuidOk

`func (o *HyperflexHxapEventAllOf) GetOwnerUuidOk() (*string, bool)`

GetOwnerUuidOk returns a tuple with the OwnerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUuid

`func (o *HyperflexHxapEventAllOf) SetOwnerUuid(v string)`

SetOwnerUuid sets OwnerUuid field to given value.

### HasOwnerUuid

`func (o *HyperflexHxapEventAllOf) HasOwnerUuid() bool`

HasOwnerUuid returns a boolean if a field has been set.

### GetSeverity

`func (o *HyperflexHxapEventAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *HyperflexHxapEventAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *HyperflexHxapEventAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *HyperflexHxapEventAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapEventAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapEventAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapEventAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapEventAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


