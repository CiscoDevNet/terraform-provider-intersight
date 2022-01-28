# HyperflexKeyEncryptionKeyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.KeyEncryptionKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.KeyEncryptionKey"]
**IsAccountRecovery** | Pointer to **bool** | This defines whether we need to operate in an account recovery scenario or not. If yes, then most of the parameters will be populated from an internal MO. So, some of the input parameters MAY be ignored, if this value is set to true. | [optional] [default to false]
**IsKekSet** | Pointer to **bool** | Indicates whether the value of the &#39;kek&#39; property has been set. | [optional] [readonly] [default to false]
**IsPassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;passphrase&#39; property has been set. | [optional] [readonly] [default to false]
**Iteration** | Pointer to **int64** | Number of iterations we want the hash to be run. | [optional] 
**Kek** | Pointer to **string** | Key encryption key used to encrypt the DEK&#39;s on the HyperFlex cluster. | [optional] 
**KeyId** | Pointer to **string** | Resource id + time of creation used for retrieving the KEK. | [optional] 
**KeyState** | Pointer to **string** | Last known Key encryption key state for this Key. * &#x60;NEW&#x60; - Key Encryption key is newly created. * &#x60;ACTIVE&#x60; - Key Encryption key is deployed on active resource. * &#x60;INACTIVE&#x60; - Key Encryption key is inactive and not used. * &#x60;INPROGRESS&#x60; - Key Encryption key is in a state where it was used on Intersight but did not receive confirmation from platform of success/failure. | [optional] [default to "NEW"]
**Passphrase** | Pointer to **string** | Initial passphrase for the encryption policy, password must contain a minimum of 12 characters, with at least 1 lowercase, 1 uppercase, 1 numeric. | [optional] 
**ResourceType** | Pointer to **string** | Resource type on which this key will be applied. * &#x60;CLUSTER&#x60; - Encryption is per HyperFlex cluster. * &#x60;DATASTORE&#x60; - Encryption is per dataStore on the HyperFlex cluster. * &#x60;DRIVE&#x60; - Encryption is per drive on the HyperFlex cluster. | [optional] [default to "CLUSTER"]
**TransitKek** | Pointer to **string** | Copy of Key encryption key, which is used for sending the key over to the remote device endpoint. It is not persisited anywhere. | [optional] 
**ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) |  | [optional] 
**ResourceMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexKeyEncryptionKeyAllOf

`func NewHyperflexKeyEncryptionKeyAllOf(classId string, objectType string, ) *HyperflexKeyEncryptionKeyAllOf`

NewHyperflexKeyEncryptionKeyAllOf instantiates a new HyperflexKeyEncryptionKeyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexKeyEncryptionKeyAllOfWithDefaults

`func NewHyperflexKeyEncryptionKeyAllOfWithDefaults() *HyperflexKeyEncryptionKeyAllOf`

NewHyperflexKeyEncryptionKeyAllOfWithDefaults instantiates a new HyperflexKeyEncryptionKeyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexKeyEncryptionKeyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexKeyEncryptionKeyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexKeyEncryptionKeyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexKeyEncryptionKeyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsAccountRecovery

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIsAccountRecovery() bool`

GetIsAccountRecovery returns the IsAccountRecovery field if non-nil, zero value otherwise.

### GetIsAccountRecoveryOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIsAccountRecoveryOk() (*bool, bool)`

GetIsAccountRecoveryOk returns a tuple with the IsAccountRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountRecovery

`func (o *HyperflexKeyEncryptionKeyAllOf) SetIsAccountRecovery(v bool)`

SetIsAccountRecovery sets IsAccountRecovery field to given value.

### HasIsAccountRecovery

`func (o *HyperflexKeyEncryptionKeyAllOf) HasIsAccountRecovery() bool`

HasIsAccountRecovery returns a boolean if a field has been set.

### GetIsKekSet

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIsKekSet() bool`

GetIsKekSet returns the IsKekSet field if non-nil, zero value otherwise.

### GetIsKekSetOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIsKekSetOk() (*bool, bool)`

GetIsKekSetOk returns a tuple with the IsKekSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKekSet

`func (o *HyperflexKeyEncryptionKeyAllOf) SetIsKekSet(v bool)`

SetIsKekSet sets IsKekSet field to given value.

### HasIsKekSet

`func (o *HyperflexKeyEncryptionKeyAllOf) HasIsKekSet() bool`

HasIsKekSet returns a boolean if a field has been set.

### GetIsPassphraseSet

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *HyperflexKeyEncryptionKeyAllOf) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *HyperflexKeyEncryptionKeyAllOf) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetIteration

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIteration() int64`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetIterationOk() (*int64, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *HyperflexKeyEncryptionKeyAllOf) SetIteration(v int64)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *HyperflexKeyEncryptionKeyAllOf) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetKek

`func (o *HyperflexKeyEncryptionKeyAllOf) GetKek() string`

GetKek returns the Kek field if non-nil, zero value otherwise.

### GetKekOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetKekOk() (*string, bool)`

GetKekOk returns a tuple with the Kek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKek

`func (o *HyperflexKeyEncryptionKeyAllOf) SetKek(v string)`

SetKek sets Kek field to given value.

### HasKek

`func (o *HyperflexKeyEncryptionKeyAllOf) HasKek() bool`

HasKek returns a boolean if a field has been set.

### GetKeyId

`func (o *HyperflexKeyEncryptionKeyAllOf) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *HyperflexKeyEncryptionKeyAllOf) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *HyperflexKeyEncryptionKeyAllOf) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeyState

`func (o *HyperflexKeyEncryptionKeyAllOf) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *HyperflexKeyEncryptionKeyAllOf) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.

### HasKeyState

`func (o *HyperflexKeyEncryptionKeyAllOf) HasKeyState() bool`

HasKeyState returns a boolean if a field has been set.

### GetPassphrase

`func (o *HyperflexKeyEncryptionKeyAllOf) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *HyperflexKeyEncryptionKeyAllOf) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *HyperflexKeyEncryptionKeyAllOf) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetResourceType

`func (o *HyperflexKeyEncryptionKeyAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HyperflexKeyEncryptionKeyAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *HyperflexKeyEncryptionKeyAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTransitKek

`func (o *HyperflexKeyEncryptionKeyAllOf) GetTransitKek() string`

GetTransitKek returns the TransitKek field if non-nil, zero value otherwise.

### GetTransitKekOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetTransitKekOk() (*string, bool)`

GetTransitKekOk returns a tuple with the TransitKek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitKek

`func (o *HyperflexKeyEncryptionKeyAllOf) SetTransitKek(v string)`

SetTransitKek sets TransitKek field to given value.

### HasTransitKek

`func (o *HyperflexKeyEncryptionKeyAllOf) HasTransitKek() bool`

HasTransitKek returns a boolean if a field has been set.

### GetClusterProfile

`func (o *HyperflexKeyEncryptionKeyAllOf) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexKeyEncryptionKeyAllOf) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexKeyEncryptionKeyAllOf) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.

### GetResourceMo

`func (o *HyperflexKeyEncryptionKeyAllOf) GetResourceMo() MoBaseMoRelationship`

GetResourceMo returns the ResourceMo field if non-nil, zero value otherwise.

### GetResourceMoOk

`func (o *HyperflexKeyEncryptionKeyAllOf) GetResourceMoOk() (*MoBaseMoRelationship, bool)`

GetResourceMoOk returns a tuple with the ResourceMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMo

`func (o *HyperflexKeyEncryptionKeyAllOf) SetResourceMo(v MoBaseMoRelationship)`

SetResourceMo sets ResourceMo field to given value.

### HasResourceMo

`func (o *HyperflexKeyEncryptionKeyAllOf) HasResourceMo() bool`

HasResourceMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


