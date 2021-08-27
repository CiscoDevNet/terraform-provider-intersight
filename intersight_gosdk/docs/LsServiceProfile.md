# LsServiceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ls.ServiceProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ls.ServiceProfile"]
**AssignState** | Pointer to **string** | Assignment state of the service profile. | [optional] [readonly] 
**AssocState** | Pointer to **string** | Association state of the service profile. | [optional] [readonly] 
**AssociatedServer** | Pointer to **string** | Server to which the UCS Manager service profile is associated to. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | Configuration state of the service profile. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the UCS Manager service profile. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the service profile. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewLsServiceProfile

`func NewLsServiceProfile(classId string, objectType string, ) *LsServiceProfile`

NewLsServiceProfile instantiates a new LsServiceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLsServiceProfileWithDefaults

`func NewLsServiceProfileWithDefaults() *LsServiceProfile`

NewLsServiceProfileWithDefaults instantiates a new LsServiceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LsServiceProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LsServiceProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LsServiceProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LsServiceProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LsServiceProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LsServiceProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssignState

`func (o *LsServiceProfile) GetAssignState() string`

GetAssignState returns the AssignState field if non-nil, zero value otherwise.

### GetAssignStateOk

`func (o *LsServiceProfile) GetAssignStateOk() (*string, bool)`

GetAssignStateOk returns a tuple with the AssignState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignState

`func (o *LsServiceProfile) SetAssignState(v string)`

SetAssignState sets AssignState field to given value.

### HasAssignState

`func (o *LsServiceProfile) HasAssignState() bool`

HasAssignState returns a boolean if a field has been set.

### GetAssocState

`func (o *LsServiceProfile) GetAssocState() string`

GetAssocState returns the AssocState field if non-nil, zero value otherwise.

### GetAssocStateOk

`func (o *LsServiceProfile) GetAssocStateOk() (*string, bool)`

GetAssocStateOk returns a tuple with the AssocState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocState

`func (o *LsServiceProfile) SetAssocState(v string)`

SetAssocState sets AssocState field to given value.

### HasAssocState

`func (o *LsServiceProfile) HasAssocState() bool`

HasAssocState returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *LsServiceProfile) GetAssociatedServer() string`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *LsServiceProfile) GetAssociatedServerOk() (*string, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *LsServiceProfile) SetAssociatedServer(v string)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *LsServiceProfile) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetConfigState

`func (o *LsServiceProfile) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *LsServiceProfile) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *LsServiceProfile) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *LsServiceProfile) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetName

`func (o *LsServiceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LsServiceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LsServiceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LsServiceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *LsServiceProfile) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *LsServiceProfile) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *LsServiceProfile) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *LsServiceProfile) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *LsServiceProfile) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *LsServiceProfile) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *LsServiceProfile) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *LsServiceProfile) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *LsServiceProfile) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *LsServiceProfile) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *LsServiceProfile) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *LsServiceProfile) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


