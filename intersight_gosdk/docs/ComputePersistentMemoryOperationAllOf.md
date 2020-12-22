# ComputePersistentMemoryOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PersistentMemoryOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PersistentMemoryOperation"]
**AdminAction** | Pointer to **string** | Administrative actions that can be performed on the Persistent Memory Modules. * &#x60;None&#x60; - No action on the selected Persistent Memory Modules. * &#x60;SecureErase&#x60; - Secure Erase action on the selected Persistent Memory Modules. * &#x60;Unlock&#x60; - Unlock action on the selected Persistent Memory Modules. | [optional] [default to "None"]
**IsSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;securePassphrase&#39; property has been set. | [optional] [readonly] [default to false]
**Modules** | Pointer to [**[]ComputePersistentMemoryModule**](ComputePersistentMemoryModule.md) |  | [optional] 
**SecurePassphrase** | Pointer to **string** | Secure passphrase of the Persistent Memory Modules of the server. | [optional] 

## Methods

### NewComputePersistentMemoryOperationAllOf

`func NewComputePersistentMemoryOperationAllOf(classId string, objectType string, ) *ComputePersistentMemoryOperationAllOf`

NewComputePersistentMemoryOperationAllOf instantiates a new ComputePersistentMemoryOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePersistentMemoryOperationAllOfWithDefaults

`func NewComputePersistentMemoryOperationAllOfWithDefaults() *ComputePersistentMemoryOperationAllOf`

NewComputePersistentMemoryOperationAllOfWithDefaults instantiates a new ComputePersistentMemoryOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePersistentMemoryOperationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePersistentMemoryOperationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePersistentMemoryOperationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePersistentMemoryOperationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePersistentMemoryOperationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePersistentMemoryOperationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *ComputePersistentMemoryOperationAllOf) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *ComputePersistentMemoryOperationAllOf) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *ComputePersistentMemoryOperationAllOf) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *ComputePersistentMemoryOperationAllOf) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetIsSecurePassphraseSet

`func (o *ComputePersistentMemoryOperationAllOf) GetIsSecurePassphraseSet() bool`

GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field if non-nil, zero value otherwise.

### GetIsSecurePassphraseSetOk

`func (o *ComputePersistentMemoryOperationAllOf) GetIsSecurePassphraseSetOk() (*bool, bool)`

GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurePassphraseSet

`func (o *ComputePersistentMemoryOperationAllOf) SetIsSecurePassphraseSet(v bool)`

SetIsSecurePassphraseSet sets IsSecurePassphraseSet field to given value.

### HasIsSecurePassphraseSet

`func (o *ComputePersistentMemoryOperationAllOf) HasIsSecurePassphraseSet() bool`

HasIsSecurePassphraseSet returns a boolean if a field has been set.

### GetModules

`func (o *ComputePersistentMemoryOperationAllOf) GetModules() []ComputePersistentMemoryModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ComputePersistentMemoryOperationAllOf) GetModulesOk() (*[]ComputePersistentMemoryModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ComputePersistentMemoryOperationAllOf) SetModules(v []ComputePersistentMemoryModule)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ComputePersistentMemoryOperationAllOf) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *ComputePersistentMemoryOperationAllOf) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *ComputePersistentMemoryOperationAllOf) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil
### GetSecurePassphrase

`func (o *ComputePersistentMemoryOperationAllOf) GetSecurePassphrase() string`

GetSecurePassphrase returns the SecurePassphrase field if non-nil, zero value otherwise.

### GetSecurePassphraseOk

`func (o *ComputePersistentMemoryOperationAllOf) GetSecurePassphraseOk() (*string, bool)`

GetSecurePassphraseOk returns a tuple with the SecurePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePassphrase

`func (o *ComputePersistentMemoryOperationAllOf) SetSecurePassphrase(v string)`

SetSecurePassphrase sets SecurePassphrase field to given value.

### HasSecurePassphrase

`func (o *ComputePersistentMemoryOperationAllOf) HasSecurePassphrase() bool`

HasSecurePassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


