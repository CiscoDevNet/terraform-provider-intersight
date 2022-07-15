# ConvergedinfraSwitchComplianceDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.SwitchComplianceDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.SwitchComplianceDetails"]
**Firmware** | Pointer to **string** | The firmware of the switch as received from inventory. | [optional] [readonly] 
**Model** | Pointer to **string** | The model information of the switch. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of switch component. It must be set to either Fabric Interconnect, Nexus or MDS. * &#x60;FabricInterconnect&#x60; - The default Switch type of UCSM and IMM mode devices. * &#x60;NexusDevice&#x60; - Switch type of Nexus devices. * &#x60;MDSDevice&#x60; - Switch type of Nexus MDS devices. | [optional] [readonly] [default to "FabricInterconnect"]
**PodCompliance** | Pointer to [**ConvergedinfraPodComplianceInfoRelationship**](ConvergedinfraPodComplianceInfoRelationship.md) |  | [optional] 
**StorageCompliances** | Pointer to [**[]ConvergedinfraStorageComplianceDetailsRelationship**](ConvergedinfraStorageComplianceDetailsRelationship.md) | An array of relationships to convergedinfraStorageComplianceDetails resources. | [optional] [readonly] 
**Switch** | Pointer to [**NetworkElementSummaryRelationship**](NetworkElementSummaryRelationship.md) |  | [optional] 

## Methods

### NewConvergedinfraSwitchComplianceDetailsAllOf

`func NewConvergedinfraSwitchComplianceDetailsAllOf(classId string, objectType string, ) *ConvergedinfraSwitchComplianceDetailsAllOf`

NewConvergedinfraSwitchComplianceDetailsAllOf instantiates a new ConvergedinfraSwitchComplianceDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraSwitchComplianceDetailsAllOfWithDefaults

`func NewConvergedinfraSwitchComplianceDetailsAllOfWithDefaults() *ConvergedinfraSwitchComplianceDetailsAllOf`

NewConvergedinfraSwitchComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraSwitchComplianceDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFirmware

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetModel

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPodCompliance

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### GetStorageCompliances

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship`

GetStorageCompliances returns the StorageCompliances field if non-nil, zero value otherwise.

### GetStorageCompliancesOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetStorageCompliancesOk() (*[]ConvergedinfraStorageComplianceDetailsRelationship, bool)`

GetStorageCompliancesOk returns a tuple with the StorageCompliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCompliances

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship)`

SetStorageCompliances sets StorageCompliances field to given value.

### HasStorageCompliances

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasStorageCompliances() bool`

HasStorageCompliances returns a boolean if a field has been set.

### SetStorageCompliancesNil

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetStorageCompliancesNil(b bool)`

 SetStorageCompliancesNil sets the value for StorageCompliances to be an explicit nil

### UnsetStorageCompliances
`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) UnsetStorageCompliances()`

UnsetStorageCompliances ensures that no value is present for StorageCompliances, not even an explicit nil
### GetSwitch

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetSwitch() NetworkElementSummaryRelationship`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetSwitchOk() (*NetworkElementSummaryRelationship, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetSwitch(v NetworkElementSummaryRelationship)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


