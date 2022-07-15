# ConvergedinfraStorageComplianceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.StorageComplianceDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.StorageComplianceDetails"]
**Os** | Pointer to **string** | The operating system name and version (e.g. NetApp ONTAP 9.10) running on the storage array for which the compliance is getting evaluated. | [optional] [readonly] 
**Protocol** | Pointer to **string** | The protocol configured for the communication between the switch and the storage array. | [optional] [readonly] 
**RefDevice** | Pointer to **string** | The reference device (e.g. adapter, fabric interconnect) against which the storage compliance is getting evaluated. * &#x60;Server&#x60; - The component type for a server in a converged infrastructure pod. * &#x60;Adapter&#x60; - The component type for an adapter on a server in a converged infrastructure pod. * &#x60;FabricInterconnect&#x60; - The component type for a fabric interconnect in a converged infrastructure pod. * &#x60;Nexus&#x60; - The component type for a nexus switch in a converged infrastructure pod. * &#x60;Storage&#x60; - The component type for a storage array in a converged infrastructure pod. | [optional] [readonly] [default to "Server"]
**AdapterCompliance** | Pointer to [**ConvergedinfraAdapterComplianceDetailsRelationship**](ConvergedinfraAdapterComplianceDetailsRelationship.md) |  | [optional] 
**PodCompliance** | Pointer to [**ConvergedinfraPodComplianceInfoRelationship**](ConvergedinfraPodComplianceInfoRelationship.md) |  | [optional] 
**StorageArray** | Pointer to [**StorageBaseArrayRelationship**](StorageBaseArrayRelationship.md) |  | [optional] 
**SwitchCompliance** | Pointer to [**ConvergedinfraSwitchComplianceDetailsRelationship**](ConvergedinfraSwitchComplianceDetailsRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraStorageComplianceDetails

`func NewConvergedinfraStorageComplianceDetails(classId string, objectType string, ) *ConvergedinfraStorageComplianceDetails`

NewConvergedinfraStorageComplianceDetails instantiates a new ConvergedinfraStorageComplianceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraStorageComplianceDetailsWithDefaults

`func NewConvergedinfraStorageComplianceDetailsWithDefaults() *ConvergedinfraStorageComplianceDetails`

NewConvergedinfraStorageComplianceDetailsWithDefaults instantiates a new ConvergedinfraStorageComplianceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraStorageComplianceDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraStorageComplianceDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraStorageComplianceDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraStorageComplianceDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraStorageComplianceDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraStorageComplianceDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOs

`func (o *ConvergedinfraStorageComplianceDetails) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ConvergedinfraStorageComplianceDetails) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ConvergedinfraStorageComplianceDetails) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ConvergedinfraStorageComplianceDetails) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetProtocol

`func (o *ConvergedinfraStorageComplianceDetails) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ConvergedinfraStorageComplianceDetails) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ConvergedinfraStorageComplianceDetails) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ConvergedinfraStorageComplianceDetails) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRefDevice

`func (o *ConvergedinfraStorageComplianceDetails) GetRefDevice() string`

GetRefDevice returns the RefDevice field if non-nil, zero value otherwise.

### GetRefDeviceOk

`func (o *ConvergedinfraStorageComplianceDetails) GetRefDeviceOk() (*string, bool)`

GetRefDeviceOk returns a tuple with the RefDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefDevice

`func (o *ConvergedinfraStorageComplianceDetails) SetRefDevice(v string)`

SetRefDevice sets RefDevice field to given value.

### HasRefDevice

`func (o *ConvergedinfraStorageComplianceDetails) HasRefDevice() bool`

HasRefDevice returns a boolean if a field has been set.

### GetAdapterCompliance

`func (o *ConvergedinfraStorageComplianceDetails) GetAdapterCompliance() ConvergedinfraAdapterComplianceDetailsRelationship`

GetAdapterCompliance returns the AdapterCompliance field if non-nil, zero value otherwise.

### GetAdapterComplianceOk

`func (o *ConvergedinfraStorageComplianceDetails) GetAdapterComplianceOk() (*ConvergedinfraAdapterComplianceDetailsRelationship, bool)`

GetAdapterComplianceOk returns a tuple with the AdapterCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterCompliance

`func (o *ConvergedinfraStorageComplianceDetails) SetAdapterCompliance(v ConvergedinfraAdapterComplianceDetailsRelationship)`

SetAdapterCompliance sets AdapterCompliance field to given value.

### HasAdapterCompliance

`func (o *ConvergedinfraStorageComplianceDetails) HasAdapterCompliance() bool`

HasAdapterCompliance returns a boolean if a field has been set.

### GetPodCompliance

`func (o *ConvergedinfraStorageComplianceDetails) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraStorageComplianceDetails) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraStorageComplianceDetails) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraStorageComplianceDetails) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### GetStorageArray

`func (o *ConvergedinfraStorageComplianceDetails) GetStorageArray() StorageBaseArrayRelationship`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *ConvergedinfraStorageComplianceDetails) GetStorageArrayOk() (*StorageBaseArrayRelationship, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *ConvergedinfraStorageComplianceDetails) SetStorageArray(v StorageBaseArrayRelationship)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *ConvergedinfraStorageComplianceDetails) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetSwitchCompliance

`func (o *ConvergedinfraStorageComplianceDetails) GetSwitchCompliance() ConvergedinfraSwitchComplianceDetailsRelationship`

GetSwitchCompliance returns the SwitchCompliance field if non-nil, zero value otherwise.

### GetSwitchComplianceOk

`func (o *ConvergedinfraStorageComplianceDetails) GetSwitchComplianceOk() (*ConvergedinfraSwitchComplianceDetailsRelationship, bool)`

GetSwitchComplianceOk returns a tuple with the SwitchCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchCompliance

`func (o *ConvergedinfraStorageComplianceDetails) SetSwitchCompliance(v ConvergedinfraSwitchComplianceDetailsRelationship)`

SetSwitchCompliance sets SwitchCompliance field to given value.

### HasSwitchCompliance

`func (o *ConvergedinfraStorageComplianceDetails) HasSwitchCompliance() bool`

HasSwitchCompliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


