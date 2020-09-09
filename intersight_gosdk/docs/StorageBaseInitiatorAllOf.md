# StorageBaseInitiatorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iqn** | Pointer to **string** | IQN (iSCSI qualified name). Can be up to 255 characters long and has the format iqn.yyyy-mm.naming-authority:unique name. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the initiator represented in the storage array. | [optional] [readonly] 
**Type** | Pointer to **string** | Initiator type, it can be FC or iSCSI. * &#x60;FC&#x60; - Fibre channel initiator type which contains WWN of an HBA on the host. * &#x60;iSCSI&#x60; - An iSCSI initiator type which contains the IQN (iSCSI Qualified Name) used by the host. * &#x60;Mixed&#x60; - Initiator type for systems using both FC and iSCSI connections. | [optional] [readonly] [default to "FC"]
**Wwn** | Pointer to **string** | World wide name, 128 bit address represented in hexadecimal notation. For example, 51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01. | [optional] [readonly] 

## Methods

### NewStorageBaseInitiatorAllOf

`func NewStorageBaseInitiatorAllOf() *StorageBaseInitiatorAllOf`

NewStorageBaseInitiatorAllOf instantiates a new StorageBaseInitiatorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseInitiatorAllOfWithDefaults

`func NewStorageBaseInitiatorAllOfWithDefaults() *StorageBaseInitiatorAllOf`

NewStorageBaseInitiatorAllOfWithDefaults instantiates a new StorageBaseInitiatorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIqn

`func (o *StorageBaseInitiatorAllOf) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *StorageBaseInitiatorAllOf) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *StorageBaseInitiatorAllOf) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *StorageBaseInitiatorAllOf) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseInitiatorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseInitiatorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseInitiatorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseInitiatorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *StorageBaseInitiatorAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageBaseInitiatorAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageBaseInitiatorAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageBaseInitiatorAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWwn

`func (o *StorageBaseInitiatorAllOf) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *StorageBaseInitiatorAllOf) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *StorageBaseInitiatorAllOf) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *StorageBaseInitiatorAllOf) HasWwn() bool`

HasWwn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


