# PartnerintegrationDeviceConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.DeviceConnector"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.DeviceConnector"]
**Action** | Pointer to **string** | Action to be performed on the device connector recipe. * &#x60;None&#x60; - Default Value of the action, i.e. do nothing. * &#x60;Build&#x60; - Build the device connector service image. * &#x60;Deploy&#x60; - Deploy the device connector service on the appliance. * &#x60;Upload&#x60; - Upload a file to the Partner Integration Appliance bucket. | [optional] [default to "None"]
**BuildStartTime** | Pointer to **string** | Time when build was triggered. | [optional] [readonly] 
**BuildStatus** | Pointer to **string** | Status of build for device connector recipe. * &#x60;None&#x60; - Default value of the status. i.e. done nothing. * &#x60;BackendInProgress&#x60; - The backend build is in progress. * &#x60;BackendFailed&#x60; - The backend build has failed. * &#x60;DockerInProgress&#x60; - The docker build is in progress. * &#x60;DockerFailed&#x60; - The docker build has failed. * &#x60;Completed&#x60; - The operation completed successfully. | [optional] [readonly] [default to "None"]
**ImageName** | Pointer to **string** | Name of the docker image that is built. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the device connector recipe. | [optional] 
**SrcBucket** | Pointer to **string** | Name of the bucket to pick up the file from. | [optional] [readonly] 
**SrcFileName** | Pointer to **string** | Name of source file to upload. | [optional] [readonly] 
**Logs** | Pointer to [**[]PartnerintegrationDcLogsRelationship**](PartnerintegrationDcLogsRelationship.md) | An array of relationships to partnerintegrationDcLogs resources. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationDeviceConnector

`func NewPartnerintegrationDeviceConnector(classId string, objectType string, ) *PartnerintegrationDeviceConnector`

NewPartnerintegrationDeviceConnector instantiates a new PartnerintegrationDeviceConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationDeviceConnectorWithDefaults

`func NewPartnerintegrationDeviceConnectorWithDefaults() *PartnerintegrationDeviceConnector`

NewPartnerintegrationDeviceConnectorWithDefaults instantiates a new PartnerintegrationDeviceConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationDeviceConnector) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationDeviceConnector) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationDeviceConnector) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationDeviceConnector) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationDeviceConnector) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationDeviceConnector) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *PartnerintegrationDeviceConnector) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PartnerintegrationDeviceConnector) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PartnerintegrationDeviceConnector) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PartnerintegrationDeviceConnector) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBuildStartTime

`func (o *PartnerintegrationDeviceConnector) GetBuildStartTime() string`

GetBuildStartTime returns the BuildStartTime field if non-nil, zero value otherwise.

### GetBuildStartTimeOk

`func (o *PartnerintegrationDeviceConnector) GetBuildStartTimeOk() (*string, bool)`

GetBuildStartTimeOk returns a tuple with the BuildStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStartTime

`func (o *PartnerintegrationDeviceConnector) SetBuildStartTime(v string)`

SetBuildStartTime sets BuildStartTime field to given value.

### HasBuildStartTime

`func (o *PartnerintegrationDeviceConnector) HasBuildStartTime() bool`

HasBuildStartTime returns a boolean if a field has been set.

### GetBuildStatus

`func (o *PartnerintegrationDeviceConnector) GetBuildStatus() string`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *PartnerintegrationDeviceConnector) GetBuildStatusOk() (*string, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *PartnerintegrationDeviceConnector) SetBuildStatus(v string)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *PartnerintegrationDeviceConnector) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### GetImageName

`func (o *PartnerintegrationDeviceConnector) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PartnerintegrationDeviceConnector) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PartnerintegrationDeviceConnector) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PartnerintegrationDeviceConnector) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *PartnerintegrationDeviceConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerintegrationDeviceConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerintegrationDeviceConnector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerintegrationDeviceConnector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrcBucket

`func (o *PartnerintegrationDeviceConnector) GetSrcBucket() string`

GetSrcBucket returns the SrcBucket field if non-nil, zero value otherwise.

### GetSrcBucketOk

`func (o *PartnerintegrationDeviceConnector) GetSrcBucketOk() (*string, bool)`

GetSrcBucketOk returns a tuple with the SrcBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBucket

`func (o *PartnerintegrationDeviceConnector) SetSrcBucket(v string)`

SetSrcBucket sets SrcBucket field to given value.

### HasSrcBucket

`func (o *PartnerintegrationDeviceConnector) HasSrcBucket() bool`

HasSrcBucket returns a boolean if a field has been set.

### GetSrcFileName

`func (o *PartnerintegrationDeviceConnector) GetSrcFileName() string`

GetSrcFileName returns the SrcFileName field if non-nil, zero value otherwise.

### GetSrcFileNameOk

`func (o *PartnerintegrationDeviceConnector) GetSrcFileNameOk() (*string, bool)`

GetSrcFileNameOk returns a tuple with the SrcFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcFileName

`func (o *PartnerintegrationDeviceConnector) SetSrcFileName(v string)`

SetSrcFileName sets SrcFileName field to given value.

### HasSrcFileName

`func (o *PartnerintegrationDeviceConnector) HasSrcFileName() bool`

HasSrcFileName returns a boolean if a field has been set.

### GetLogs

`func (o *PartnerintegrationDeviceConnector) GetLogs() []PartnerintegrationDcLogsRelationship`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PartnerintegrationDeviceConnector) GetLogsOk() (*[]PartnerintegrationDcLogsRelationship, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PartnerintegrationDeviceConnector) SetLogs(v []PartnerintegrationDcLogsRelationship)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PartnerintegrationDeviceConnector) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *PartnerintegrationDeviceConnector) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *PartnerintegrationDeviceConnector) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetOrganization

`func (o *PartnerintegrationDeviceConnector) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerintegrationDeviceConnector) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerintegrationDeviceConnector) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerintegrationDeviceConnector) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *PartnerintegrationDeviceConnector) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *PartnerintegrationDeviceConnector) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


