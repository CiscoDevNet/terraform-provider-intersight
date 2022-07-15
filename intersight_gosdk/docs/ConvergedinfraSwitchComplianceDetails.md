# ConvergedinfraSwitchComplianceDetails

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

### NewConvergedinfraSwitchComplianceDetails

`func NewConvergedinfraSwitchComplianceDetails(classId string, objectType string, ) *ConvergedinfraSwitchComplianceDetails`

NewConvergedinfraSwitchComplianceDetails instantiates a new ConvergedinfraSwitchComplianceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraSwitchComplianceDetailsWithDefaults

`func NewConvergedinfraSwitchComplianceDetailsWithDefaults() *ConvergedinfraSwitchComplianceDetails`

NewConvergedinfraSwitchComplianceDetailsWithDefaults instantiates a new ConvergedinfraSwitchComplianceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraSwitchComplianceDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraSwitchComplianceDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraSwitchComplianceDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraSwitchComplianceDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFirmware

`func (o *ConvergedinfraSwitchComplianceDetails) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *ConvergedinfraSwitchComplianceDetails) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *ConvergedinfraSwitchComplianceDetails) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetModel

`func (o *ConvergedinfraSwitchComplianceDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConvergedinfraSwitchComplianceDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConvergedinfraSwitchComplianceDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *ConvergedinfraSwitchComplianceDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConvergedinfraSwitchComplianceDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConvergedinfraSwitchComplianceDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPodCompliance

`func (o *ConvergedinfraSwitchComplianceDetails) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship`

GetPodCompliance returns the PodCompliance field if non-nil, zero value otherwise.

### GetPodComplianceOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool)`

GetPodComplianceOk returns a tuple with the PodCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCompliance

`func (o *ConvergedinfraSwitchComplianceDetails) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship)`

SetPodCompliance sets PodCompliance field to given value.

### HasPodCompliance

`func (o *ConvergedinfraSwitchComplianceDetails) HasPodCompliance() bool`

HasPodCompliance returns a boolean if a field has been set.

### GetStorageCompliances

`func (o *ConvergedinfraSwitchComplianceDetails) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship`

GetStorageCompliances returns the StorageCompliances field if non-nil, zero value otherwise.

### GetStorageCompliancesOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetStorageCompliancesOk() (*[]ConvergedinfraStorageComplianceDetailsRelationship, bool)`

GetStorageCompliancesOk returns a tuple with the StorageCompliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCompliances

`func (o *ConvergedinfraSwitchComplianceDetails) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship)`

SetStorageCompliances sets StorageCompliances field to given value.

### HasStorageCompliances

`func (o *ConvergedinfraSwitchComplianceDetails) HasStorageCompliances() bool`

HasStorageCompliances returns a boolean if a field has been set.

### SetStorageCompliancesNil

`func (o *ConvergedinfraSwitchComplianceDetails) SetStorageCompliancesNil(b bool)`

 SetStorageCompliancesNil sets the value for StorageCompliances to be an explicit nil

### UnsetStorageCompliances
`func (o *ConvergedinfraSwitchComplianceDetails) UnsetStorageCompliances()`

UnsetStorageCompliances ensures that no value is present for StorageCompliances, not even an explicit nil
### GetSwitch

`func (o *ConvergedinfraSwitchComplianceDetails) GetSwitch() NetworkElementSummaryRelationship`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *ConvergedinfraSwitchComplianceDetails) GetSwitchOk() (*NetworkElementSummaryRelationship, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *ConvergedinfraSwitchComplianceDetails) SetSwitch(v NetworkElementSummaryRelationship)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *ConvergedinfraSwitchComplianceDetails) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


