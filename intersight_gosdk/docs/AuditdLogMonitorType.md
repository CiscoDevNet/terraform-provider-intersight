# AuditdLogMonitorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "auditd.LogMonitorType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "auditd.LogMonitorType"]
**All** | Pointer to **string** | It can be configured to monitor all the log events. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**AuthLogFiles** | Pointer to **string** | It can be configured to monitor log events only w.r.t auth log files changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**CronFiles** | Pointer to **string** | It can be configured to monitor log events only w.r.t cron files changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**DnsClientFiles** | Pointer to **string** | It can be configured to monitor log events only w.r.t dns client files changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**Docker** | Pointer to **string** | It can be configured to monitor log events only w.r.t Docker executions and file changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**KernelModuleMgmt** | Pointer to **string** | It can be configured to monitor log events only w.r.t kernel module files changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**ProcessAudit** | Pointer to **string** | It can be configured to monitor log events only w.r.t process execution audit. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**SystemLogFiles** | Pointer to **string** | It can be configured to monitor log events only w.r.t system log files changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**SystemLoginReboot** | Pointer to **string** | It can be configured to monitor log events only w.r.t system login reboot file changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**SystemSoftware** | Pointer to **string** | It can be configured to monitor log events only w.r.t system software&#39;s binaries changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**SystemTimeChange** | Pointer to **string** | It can be configured to monitor log events only w.r.t system time file changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**UserGroupConfigFiles** | Pointer to **string** | It can be configured to monitor log events only w.r.t User Group Config Files changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]
**UserPrivilegeMgmt** | Pointer to **string** | It can be configured to monitor log events only w.r.t User Privilege management file changes. * &#x60;no&#x60; - Value to disable the specific monitoring rule. * &#x60;yes&#x60; - Value to enable the specific monitoring rule. | [optional] [default to "no"]

## Methods

### NewAuditdLogMonitorType

`func NewAuditdLogMonitorType(classId string, objectType string, ) *AuditdLogMonitorType`

NewAuditdLogMonitorType instantiates a new AuditdLogMonitorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditdLogMonitorTypeWithDefaults

`func NewAuditdLogMonitorTypeWithDefaults() *AuditdLogMonitorType`

NewAuditdLogMonitorTypeWithDefaults instantiates a new AuditdLogMonitorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AuditdLogMonitorType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AuditdLogMonitorType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AuditdLogMonitorType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AuditdLogMonitorType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AuditdLogMonitorType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AuditdLogMonitorType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAll

`func (o *AuditdLogMonitorType) GetAll() string`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *AuditdLogMonitorType) GetAllOk() (*string, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *AuditdLogMonitorType) SetAll(v string)`

SetAll sets All field to given value.

### HasAll

`func (o *AuditdLogMonitorType) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAuthLogFiles

`func (o *AuditdLogMonitorType) GetAuthLogFiles() string`

GetAuthLogFiles returns the AuthLogFiles field if non-nil, zero value otherwise.

### GetAuthLogFilesOk

`func (o *AuditdLogMonitorType) GetAuthLogFilesOk() (*string, bool)`

GetAuthLogFilesOk returns a tuple with the AuthLogFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLogFiles

`func (o *AuditdLogMonitorType) SetAuthLogFiles(v string)`

SetAuthLogFiles sets AuthLogFiles field to given value.

### HasAuthLogFiles

`func (o *AuditdLogMonitorType) HasAuthLogFiles() bool`

HasAuthLogFiles returns a boolean if a field has been set.

### GetCronFiles

`func (o *AuditdLogMonitorType) GetCronFiles() string`

GetCronFiles returns the CronFiles field if non-nil, zero value otherwise.

### GetCronFilesOk

`func (o *AuditdLogMonitorType) GetCronFilesOk() (*string, bool)`

GetCronFilesOk returns a tuple with the CronFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronFiles

`func (o *AuditdLogMonitorType) SetCronFiles(v string)`

SetCronFiles sets CronFiles field to given value.

### HasCronFiles

`func (o *AuditdLogMonitorType) HasCronFiles() bool`

HasCronFiles returns a boolean if a field has been set.

### GetDnsClientFiles

`func (o *AuditdLogMonitorType) GetDnsClientFiles() string`

GetDnsClientFiles returns the DnsClientFiles field if non-nil, zero value otherwise.

### GetDnsClientFilesOk

`func (o *AuditdLogMonitorType) GetDnsClientFilesOk() (*string, bool)`

GetDnsClientFilesOk returns a tuple with the DnsClientFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsClientFiles

`func (o *AuditdLogMonitorType) SetDnsClientFiles(v string)`

SetDnsClientFiles sets DnsClientFiles field to given value.

### HasDnsClientFiles

`func (o *AuditdLogMonitorType) HasDnsClientFiles() bool`

HasDnsClientFiles returns a boolean if a field has been set.

### GetDocker

`func (o *AuditdLogMonitorType) GetDocker() string`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *AuditdLogMonitorType) GetDockerOk() (*string, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *AuditdLogMonitorType) SetDocker(v string)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *AuditdLogMonitorType) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetKernelModuleMgmt

`func (o *AuditdLogMonitorType) GetKernelModuleMgmt() string`

GetKernelModuleMgmt returns the KernelModuleMgmt field if non-nil, zero value otherwise.

### GetKernelModuleMgmtOk

`func (o *AuditdLogMonitorType) GetKernelModuleMgmtOk() (*string, bool)`

GetKernelModuleMgmtOk returns a tuple with the KernelModuleMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelModuleMgmt

`func (o *AuditdLogMonitorType) SetKernelModuleMgmt(v string)`

SetKernelModuleMgmt sets KernelModuleMgmt field to given value.

### HasKernelModuleMgmt

`func (o *AuditdLogMonitorType) HasKernelModuleMgmt() bool`

HasKernelModuleMgmt returns a boolean if a field has been set.

### GetProcessAudit

`func (o *AuditdLogMonitorType) GetProcessAudit() string`

GetProcessAudit returns the ProcessAudit field if non-nil, zero value otherwise.

### GetProcessAuditOk

`func (o *AuditdLogMonitorType) GetProcessAuditOk() (*string, bool)`

GetProcessAuditOk returns a tuple with the ProcessAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessAudit

`func (o *AuditdLogMonitorType) SetProcessAudit(v string)`

SetProcessAudit sets ProcessAudit field to given value.

### HasProcessAudit

`func (o *AuditdLogMonitorType) HasProcessAudit() bool`

HasProcessAudit returns a boolean if a field has been set.

### GetSystemLogFiles

`func (o *AuditdLogMonitorType) GetSystemLogFiles() string`

GetSystemLogFiles returns the SystemLogFiles field if non-nil, zero value otherwise.

### GetSystemLogFilesOk

`func (o *AuditdLogMonitorType) GetSystemLogFilesOk() (*string, bool)`

GetSystemLogFilesOk returns a tuple with the SystemLogFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLogFiles

`func (o *AuditdLogMonitorType) SetSystemLogFiles(v string)`

SetSystemLogFiles sets SystemLogFiles field to given value.

### HasSystemLogFiles

`func (o *AuditdLogMonitorType) HasSystemLogFiles() bool`

HasSystemLogFiles returns a boolean if a field has been set.

### GetSystemLoginReboot

`func (o *AuditdLogMonitorType) GetSystemLoginReboot() string`

GetSystemLoginReboot returns the SystemLoginReboot field if non-nil, zero value otherwise.

### GetSystemLoginRebootOk

`func (o *AuditdLogMonitorType) GetSystemLoginRebootOk() (*string, bool)`

GetSystemLoginRebootOk returns a tuple with the SystemLoginReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLoginReboot

`func (o *AuditdLogMonitorType) SetSystemLoginReboot(v string)`

SetSystemLoginReboot sets SystemLoginReboot field to given value.

### HasSystemLoginReboot

`func (o *AuditdLogMonitorType) HasSystemLoginReboot() bool`

HasSystemLoginReboot returns a boolean if a field has been set.

### GetSystemSoftware

`func (o *AuditdLogMonitorType) GetSystemSoftware() string`

GetSystemSoftware returns the SystemSoftware field if non-nil, zero value otherwise.

### GetSystemSoftwareOk

`func (o *AuditdLogMonitorType) GetSystemSoftwareOk() (*string, bool)`

GetSystemSoftwareOk returns a tuple with the SystemSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSoftware

`func (o *AuditdLogMonitorType) SetSystemSoftware(v string)`

SetSystemSoftware sets SystemSoftware field to given value.

### HasSystemSoftware

`func (o *AuditdLogMonitorType) HasSystemSoftware() bool`

HasSystemSoftware returns a boolean if a field has been set.

### GetSystemTimeChange

`func (o *AuditdLogMonitorType) GetSystemTimeChange() string`

GetSystemTimeChange returns the SystemTimeChange field if non-nil, zero value otherwise.

### GetSystemTimeChangeOk

`func (o *AuditdLogMonitorType) GetSystemTimeChangeOk() (*string, bool)`

GetSystemTimeChangeOk returns a tuple with the SystemTimeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTimeChange

`func (o *AuditdLogMonitorType) SetSystemTimeChange(v string)`

SetSystemTimeChange sets SystemTimeChange field to given value.

### HasSystemTimeChange

`func (o *AuditdLogMonitorType) HasSystemTimeChange() bool`

HasSystemTimeChange returns a boolean if a field has been set.

### GetUserGroupConfigFiles

`func (o *AuditdLogMonitorType) GetUserGroupConfigFiles() string`

GetUserGroupConfigFiles returns the UserGroupConfigFiles field if non-nil, zero value otherwise.

### GetUserGroupConfigFilesOk

`func (o *AuditdLogMonitorType) GetUserGroupConfigFilesOk() (*string, bool)`

GetUserGroupConfigFilesOk returns a tuple with the UserGroupConfigFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupConfigFiles

`func (o *AuditdLogMonitorType) SetUserGroupConfigFiles(v string)`

SetUserGroupConfigFiles sets UserGroupConfigFiles field to given value.

### HasUserGroupConfigFiles

`func (o *AuditdLogMonitorType) HasUserGroupConfigFiles() bool`

HasUserGroupConfigFiles returns a boolean if a field has been set.

### GetUserPrivilegeMgmt

`func (o *AuditdLogMonitorType) GetUserPrivilegeMgmt() string`

GetUserPrivilegeMgmt returns the UserPrivilegeMgmt field if non-nil, zero value otherwise.

### GetUserPrivilegeMgmtOk

`func (o *AuditdLogMonitorType) GetUserPrivilegeMgmtOk() (*string, bool)`

GetUserPrivilegeMgmtOk returns a tuple with the UserPrivilegeMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrivilegeMgmt

`func (o *AuditdLogMonitorType) SetUserPrivilegeMgmt(v string)`

SetUserPrivilegeMgmt sets UserPrivilegeMgmt field to given value.

### HasUserPrivilegeMgmt

`func (o *AuditdLogMonitorType) HasUserPrivilegeMgmt() bool`

HasUserPrivilegeMgmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


