# StorageNetAppFcInterfaceAllOf

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

### NewStorageNetAppFcInterfaceAllOf

`func NewStorageNetAppFcInterfaceAllOf(classId string, objectType string, ) *StorageNetAppFcInterfaceAllOf`

NewStorageNetAppFcInterfaceAllOf instantiates a new StorageNetAppFcInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcInterfaceAllOfWithDefaults

`func NewStorageNetAppFcInterfaceAllOfWithDefaults() *StorageNetAppFcInterfaceAllOf`

NewStorageNetAppFcInterfaceAllOfWithDefaults instantiates a new StorageNetAppFcInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *StorageNetAppFcInterfaceAllOf) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppFcInterfaceAllOf) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppFcInterfaceAllOf) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppFcInterfaceAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppFcInterfaceAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppFcInterfaceAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppFcInterfaceAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppFcInterfaceAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppFcInterfaceAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppFcInterfaceAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppFcInterfaceAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppFcInterfaceAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppFcInterfaceAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppFcInterfaceAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppFcInterfaceAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppFcInterfaceAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetPhysicalPort

`func (o *StorageNetAppFcInterfaceAllOf) GetPhysicalPort() StorageNetAppFcPortRelationship`

GetPhysicalPort returns the PhysicalPort field if non-nil, zero value otherwise.

### GetPhysicalPortOk

`func (o *StorageNetAppFcInterfaceAllOf) GetPhysicalPortOk() (*StorageNetAppFcPortRelationship, bool)`

GetPhysicalPortOk returns a tuple with the PhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPort

`func (o *StorageNetAppFcInterfaceAllOf) SetPhysicalPort(v StorageNetAppFcPortRelationship)`

SetPhysicalPort sets PhysicalPort field to given value.

### HasPhysicalPort

`func (o *StorageNetAppFcInterfaceAllOf) HasPhysicalPort() bool`

HasPhysicalPort returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppFcInterfaceAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppFcInterfaceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppFcInterfaceAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppFcInterfaceAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


