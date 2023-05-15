# StorageNetAppSnapMirrorRelationshipAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppSnapMirrorRelationship"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppSnapMirrorRelationship"]
**DestinationPath** | Pointer to **string** | Path to the destination endpoint of a SnapMirror relationship. Examples: ONTAP FlexVol/FlexGroup - svm1:volume1; ONTAP SVM - svm1: ; ONTAP Consistency Group - svm1:/cg/cg_name. | [optional] [readonly] 
**Healthy** | Pointer to **string** | Whether the relationship is healthy or not. | [optional] [readonly] 
**LagTime** | Pointer to **string** | Time since the exported Snapshot copy was created. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Name of the NetApp SnapMirror policy. | [optional] [readonly] 
**PolicyType** | Pointer to **string** | SnapMirror policy type can be async, sync, or continuous. | [optional] [readonly] 
**PolicyUuid** | Pointer to **string** | Uuid of the NetApp SnapMirror policy. | [optional] [readonly] 
**SourcePath** | Pointer to **string** | Path to the source endpoint of a SnapMirror relationship. Examples: ONTAP FlexVol/FlexGroup - svm1:volume1; ONTAP SVM - svm1: ; ONTAP Consistency Group - svm1:/cg/cg_name. | [optional] [readonly] 
**State** | Pointer to **string** | State of the relationship. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of the NetApp SnapMirror relationship. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**DestinationTenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 
**Policy** | Pointer to [**StorageNetAppBaseSnapMirrorPolicyRelationship**](StorageNetAppBaseSnapMirrorPolicyRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppSnapMirrorRelationshipAllOf

`func NewStorageNetAppSnapMirrorRelationshipAllOf(classId string, objectType string, ) *StorageNetAppSnapMirrorRelationshipAllOf`

NewStorageNetAppSnapMirrorRelationshipAllOf instantiates a new StorageNetAppSnapMirrorRelationshipAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppSnapMirrorRelationshipAllOfWithDefaults

`func NewStorageNetAppSnapMirrorRelationshipAllOfWithDefaults() *StorageNetAppSnapMirrorRelationshipAllOf`

NewStorageNetAppSnapMirrorRelationshipAllOfWithDefaults instantiates a new StorageNetAppSnapMirrorRelationshipAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationPath

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetHealthy

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetHealthy() string`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetHealthyOk() (*string, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetHealthy(v string)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetLagTime

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetLagTime() string`

GetLagTime returns the LagTime field if non-nil, zero value otherwise.

### GetLagTimeOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetLagTimeOk() (*string, bool)`

GetLagTimeOk returns a tuple with the LagTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagTime

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetLagTime(v string)`

SetLagTime sets LagTime field to given value.

### HasLagTime

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasLagTime() bool`

HasLagTime returns a boolean if a field has been set.

### GetPolicyName

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyType

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPolicyUuid

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyUuid() string`

GetPolicyUuid returns the PolicyUuid field if non-nil, zero value otherwise.

### GetPolicyUuidOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyUuidOk() (*string, bool)`

GetPolicyUuidOk returns a tuple with the PolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUuid

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetPolicyUuid(v string)`

SetPolicyUuid sets PolicyUuid field to given value.

### HasPolicyUuid

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasPolicyUuid() bool`

HasPolicyUuid returns a boolean if a field has been set.

### GetSourcePath

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetDestinationTenant

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetDestinationTenant() StorageNetAppStorageVmRelationship`

GetDestinationTenant returns the DestinationTenant field if non-nil, zero value otherwise.

### GetDestinationTenantOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetDestinationTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetDestinationTenantOk returns a tuple with the DestinationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTenant

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetDestinationTenant(v StorageNetAppStorageVmRelationship)`

SetDestinationTenant sets DestinationTenant field to given value.

### HasDestinationTenant

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasDestinationTenant() bool`

HasDestinationTenant returns a boolean if a field has been set.

### GetPolicy

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicy() StorageNetAppBaseSnapMirrorPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) GetPolicyOk() (*StorageNetAppBaseSnapMirrorPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) SetPolicy(v StorageNetAppBaseSnapMirrorPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StorageNetAppSnapMirrorRelationshipAllOf) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


