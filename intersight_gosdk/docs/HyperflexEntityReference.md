# HyperflexEntityReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.EntityReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.EntityReference"]
**Confignum** | Pointer to **int64** | Configuration number for this reference. | [optional] [readonly] 
**Id** | Pointer to **string** | Uuid of the entity for this reference. | [optional] [readonly] 
**Idtype** | Pointer to **string** | Type of entity id for this reference. * &#x60;VCMOID&#x60; - The entity reference ID type is VC MOID. * &#x60;VMBIOSUUID&#x60; - The entity reference ID type is VM Bios UUID. * &#x60;VMDSPATH&#x60; - The entity reference ID type is VM Datastore Path. * &#x60;VMINSTANCEUUID&#x60; - The entity reference ID type is VM Instance UUID. * &#x60;VMNAME&#x60; - The entity reference ID type is VM Name. | [optional] [readonly] [default to "VCMOID"]
**Name** | Pointer to **string** | Name of the entity for this entity reference. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the entity for this entity reference. * &#x60;DISK&#x60; - This entity type is a disk. * &#x60;PNODE&#x60; - This entity type is a P node. * &#x60;NODE&#x60; - This entity type is a node. * &#x60;CLUSTER&#x60; - This entity type is a cluster. * &#x60;DATASTORE&#x60; - This entity is a datastore. * &#x60;VIRTNODE&#x60; - This entity is a HyperFlex virtual node. * &#x60;VIRTCLUSTER&#x60; - This entity type is a virtual cluster. * &#x60;VIRTDATASTORE&#x60; - This entity type is a virtual data store. * &#x60;VIRTMACHINE&#x60; - This entity type is a virtual machine. * &#x60;PDISK&#x60; - This entity type is a P disk. * &#x60;PDATASTORE&#x60; - This entity type is a P Datastore. * &#x60;VIRTMACHINESNAPSHOT&#x60; - This entity is a virtual machine snapshot. * &#x60;FOLDER&#x60; - This entity type is a folder. * &#x60;RESOURCEPOOL&#x60; - This entity type is a resource pool. * &#x60;FILE&#x60; - This entity type is a file. * &#x60;VIRTDATACENTER&#x60; - This entity type is a virtual data center. * &#x60;REPLICATION_APPLIANCE&#x60; - This entity type is a replication appliance. * &#x60;REPLICATION_JOB&#x60; - This entity type is a replication job. * &#x60;IP_POOL&#x60; - This entity type is an IP Pool. * &#x60;REPLICATION_INFO&#x60; - This entity type is a replication information. * &#x60;DP_VM_SNAPSHOT&#x60; - This entity type is a DP VM Snapshot. * &#x60;DP_VMGROUP_SNAPSHOT&#x60; - This entity type is a DP VM Group Snapshot. * &#x60;DP_VM_CONFIG&#x60; - This entity type is a DP VM Configuration. * &#x60;DP_VM&#x60; - This entity type is a DP VM. * &#x60;DP_VMGROUP&#x60; - This entity type is a DP VM Group. * &#x60;DP_VM_SNAPSHOT_POINT&#x60; - This entity type is a DP VM Snapshot Point. * &#x60;CLUSTER_PAIR&#x60; - This entity is a cluster pair. * &#x60;HX_TASK&#x60; - This entity type is a HyperFlex task. * &#x60;ZONE&#x60; - This entity type is a zone. | [optional] [readonly] [default to "DISK"]

## Methods

### NewHyperflexEntityReference

`func NewHyperflexEntityReference(classId string, objectType string, ) *HyperflexEntityReference`

NewHyperflexEntityReference instantiates a new HyperflexEntityReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexEntityReferenceWithDefaults

`func NewHyperflexEntityReferenceWithDefaults() *HyperflexEntityReference`

NewHyperflexEntityReferenceWithDefaults instantiates a new HyperflexEntityReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexEntityReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexEntityReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexEntityReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexEntityReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexEntityReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexEntityReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfignum

`func (o *HyperflexEntityReference) GetConfignum() int64`

GetConfignum returns the Confignum field if non-nil, zero value otherwise.

### GetConfignumOk

`func (o *HyperflexEntityReference) GetConfignumOk() (*int64, bool)`

GetConfignumOk returns a tuple with the Confignum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfignum

`func (o *HyperflexEntityReference) SetConfignum(v int64)`

SetConfignum sets Confignum field to given value.

### HasConfignum

`func (o *HyperflexEntityReference) HasConfignum() bool`

HasConfignum returns a boolean if a field has been set.

### GetId

`func (o *HyperflexEntityReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperflexEntityReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperflexEntityReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HyperflexEntityReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdtype

`func (o *HyperflexEntityReference) GetIdtype() string`

GetIdtype returns the Idtype field if non-nil, zero value otherwise.

### GetIdtypeOk

`func (o *HyperflexEntityReference) GetIdtypeOk() (*string, bool)`

GetIdtypeOk returns a tuple with the Idtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdtype

`func (o *HyperflexEntityReference) SetIdtype(v string)`

SetIdtype sets Idtype field to given value.

### HasIdtype

`func (o *HyperflexEntityReference) HasIdtype() bool`

HasIdtype returns a boolean if a field has been set.

### GetName

`func (o *HyperflexEntityReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexEntityReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexEntityReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexEntityReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *HyperflexEntityReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexEntityReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexEntityReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexEntityReference) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


