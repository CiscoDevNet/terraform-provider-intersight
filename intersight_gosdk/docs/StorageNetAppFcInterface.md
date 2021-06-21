# StorageNetAppFcInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppFcInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppFcInterface"]
**Enabled** | Pointer to **string** | FC interface is enabled or not. | [optional] [readonly] 
**State** | Pointer to **string** | State of the FC interface. * &#x60;down&#x60; - An inactive port is listed as Down. * &#x60;up&#x60; - An active port is listed as Up. * &#x60;present&#x60; - An active port is listed as present. | [optional] [readonly] [default to "down"]
**Uuid** | Pointer to **string** | Uuid of  NetApp FC Interface. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](storage.NetAppNode.Relationship.md) |  | [optional] 
**PhysicalPort** | Pointer to [**StorageNetAppFcPortRelationship**](storage.NetAppFcPort.Relationship.md) |  | [optional] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](storage.NetAppStorageVm.Relationship.md) |  | [optional] 

## Methods

### NewStorageNetAppFcInterface

`func NewStorageNetAppFcInterface(classId string, objectType string, ) *StorageNetAppFcInterface`

NewStorageNetAppFcInterface instantiates a new StorageNetAppFcInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcInterfaceWithDefaults

`func NewStorageNetAppFcInterfaceWithDefaults() *StorageNetAppFcInterface`

NewStorageNetAppFcInterfaceWithDefaults instantiates a new StorageNetAppFcInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *StorageNetAppFcInterface) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppFcInterface) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppFcInterface) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppFcInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppFcInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppFcInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppFcInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppFcInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppFcInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppFcInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppFcInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppFcInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppFcInterface) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppFcInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppFcInterface) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppFcInterface) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetPhysicalPort

`func (o *StorageNetAppFcInterface) GetPhysicalPort() StorageNetAppFcPortRelationship`

GetPhysicalPort returns the PhysicalPort field if non-nil, zero value otherwise.

### GetPhysicalPortOk

`func (o *StorageNetAppFcInterface) GetPhysicalPortOk() (*StorageNetAppFcPortRelationship, bool)`

GetPhysicalPortOk returns a tuple with the PhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPort

`func (o *StorageNetAppFcInterface) SetPhysicalPort(v StorageNetAppFcPortRelationship)`

SetPhysicalPort sets PhysicalPort field to given value.

### HasPhysicalPort

`func (o *StorageNetAppFcInterface) HasPhysicalPort() bool`

HasPhysicalPort returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppFcInterface) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppFcInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppFcInterface) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppFcInterface) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


