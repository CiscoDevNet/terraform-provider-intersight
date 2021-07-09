package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenAdditionalProperties(m interface{}) string {
	var s string
	if m != nil && reflect.ValueOf(m).Len() > 0 {
		j, err := json.Marshal(m)
		if err != nil {
			log.Printf("Error occurred while flattening and json parsing: %s", err)
		} else {
			s = string(j)
		}
	}
	return s
}
func flattenMoMoRef(ref *models.MoMoRef) map[string]interface{} {
	x := make(map[string]interface{})
	x["additional_properties"] = flattenAdditionalProperties(ref.AdditionalProperties)
	x["class_id"] = ref.GetClassId()
	x["moid"] = ref.GetMoid()
	x["object_type"] = ref.GetObjectType()
	x["selector"] = ref.GetSelector()
	return x
}

func flattenListAdapterAdapterConfig(p []models.AdapterAdapterConfig, d *schema.ResourceData) []map[string]interface{} {
	var adapteradapterconfigs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		adapteradapterconfig := make(map[string]interface{})
		adapteradapterconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		adapteradapterconfig["class_id"] = item.GetClassId()
		adapteradapterconfig["dce_interface_settings"] = (func(p []models.AdapterDceInterfaceSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterdceinterfacesettingss []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				adapterdceinterfacesettings := make(map[string]interface{})
				adapterdceinterfacesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				adapterdceinterfacesettings["class_id"] = item.GetClassId()
				adapterdceinterfacesettings["fec_mode"] = item.GetFecMode()
				adapterdceinterfacesettings["interface_id"] = item.GetInterfaceId()
				adapterdceinterfacesettings["object_type"] = item.GetObjectType()
				adapterdceinterfacesettingss = append(adapterdceinterfacesettingss, adapterdceinterfacesettings)
			}
			return adapterdceinterfacesettingss
		})(item.GetDceInterfaceSettings(), d)
		adapteradapterconfig["eth_settings"] = (func(p models.AdapterEthSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterethsettingss []map[string]interface{}
			var ret models.AdapterEthSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterethsettings := make(map[string]interface{})
			adapterethsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterethsettings["class_id"] = item.GetClassId()
			adapterethsettings["lldp_enabled"] = item.GetLldpEnabled()
			adapterethsettings["object_type"] = item.GetObjectType()

			adapterethsettingss = append(adapterethsettingss, adapterethsettings)
			return adapterethsettingss
		})(item.GetEthSettings(), d)
		adapteradapterconfig["fc_settings"] = (func(p models.AdapterFcSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterfcsettingss []map[string]interface{}
			var ret models.AdapterFcSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterfcsettings := make(map[string]interface{})
			adapterfcsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterfcsettings["class_id"] = item.GetClassId()
			adapterfcsettings["fip_enabled"] = item.GetFipEnabled()
			adapterfcsettings["object_type"] = item.GetObjectType()

			adapterfcsettingss = append(adapterfcsettingss, adapterfcsettings)
			return adapterfcsettingss
		})(item.GetFcSettings(), d)
		adapteradapterconfig["object_type"] = item.GetObjectType()
		adapteradapterconfig["port_channel_settings"] = (func(p models.AdapterPortChannelSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterportchannelsettingss []map[string]interface{}
			var ret models.AdapterPortChannelSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterportchannelsettings := make(map[string]interface{})
			adapterportchannelsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterportchannelsettings["class_id"] = item.GetClassId()
			adapterportchannelsettings["enabled"] = item.GetEnabled()
			adapterportchannelsettings["object_type"] = item.GetObjectType()

			adapterportchannelsettingss = append(adapterportchannelsettingss, adapterportchannelsettings)
			return adapterportchannelsettingss
		})(item.GetPortChannelSettings(), d)
		adapteradapterconfig["slot_id"] = item.GetSlotId()
		adapteradapterconfigs = append(adapteradapterconfigs, adapteradapterconfig)
	}
	return adapteradapterconfigs
}
func flattenListAdapterExtEthInterfaceRelationship(p []models.AdapterExtEthInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterextethinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterextethinterfacerelationship := flattenMoMoRef(item)
		adapterextethinterfacerelationships = append(adapterextethinterfacerelationships, adapterextethinterfacerelationship)
	}
	return adapterextethinterfacerelationships
}
func flattenListAdapterHostEthInterfaceRelationship(p []models.AdapterHostEthInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostethinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterhostethinterfacerelationship := flattenMoMoRef(item)
		adapterhostethinterfacerelationships = append(adapterhostethinterfacerelationships, adapterhostethinterfacerelationship)
	}
	return adapterhostethinterfacerelationships
}
func flattenListAdapterHostFcInterfaceRelationship(p []models.AdapterHostFcInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostfcinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterhostfcinterfacerelationship := flattenMoMoRef(item)
		adapterhostfcinterfacerelationships = append(adapterhostfcinterfacerelationships, adapterhostfcinterfacerelationship)
	}
	return adapterhostfcinterfacerelationships
}
func flattenListAdapterHostIscsiInterfaceRelationship(p []models.AdapterHostIscsiInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostiscsiinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterhostiscsiinterfacerelationship := flattenMoMoRef(item)
		adapterhostiscsiinterfacerelationships = append(adapterhostiscsiinterfacerelationships, adapterhostiscsiinterfacerelationship)
	}
	return adapterhostiscsiinterfacerelationships
}
func flattenListAdapterUnitRelationship(p []models.AdapterUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterunitrelationship := flattenMoMoRef(item)
		adapterunitrelationships = append(adapterunitrelationships, adapterunitrelationship)
	}
	return adapterunitrelationships
}
func flattenListApplianceApiStatus(p []models.ApplianceApiStatus, d *schema.ResourceData) []map[string]interface{} {
	var applianceapistatuss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		applianceapistatus := make(map[string]interface{})
		applianceapistatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		applianceapistatus["class_id"] = item.GetClassId()
		applianceapistatus["elapsed_time"] = item.GetElapsedTime()
		applianceapistatus["object_name"] = item.GetObjectName()
		applianceapistatus["object_type"] = item.GetObjectType()
		applianceapistatus["reason"] = item.GetReason()
		applianceapistatus["response"] = item.GetResponse()
		applianceapistatuss = append(applianceapistatuss, applianceapistatus)
	}
	return applianceapistatuss
}
func flattenListApplianceAppStatusRelationship(p []models.ApplianceAppStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var applianceappstatusrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		applianceappstatusrelationship := flattenMoMoRef(item)
		applianceappstatusrelationships = append(applianceappstatusrelationships, applianceappstatusrelationship)
	}
	return applianceappstatusrelationships
}
func flattenListApplianceCertRenewalPhase(p []models.ApplianceCertRenewalPhase, d *schema.ResourceData) []map[string]interface{} {
	var appliancecertrenewalphases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		appliancecertrenewalphase := make(map[string]interface{})
		appliancecertrenewalphase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		appliancecertrenewalphase["class_id"] = item.GetClassId()
		appliancecertrenewalphase["end_time"] = item.GetEndTime()
		appliancecertrenewalphase["failed"] = item.GetFailed()
		appliancecertrenewalphase["message"] = item.GetMessage()
		appliancecertrenewalphase["name"] = item.GetName()
		appliancecertrenewalphase["object_type"] = item.GetObjectType()
		appliancecertrenewalphase["start_time"] = item.GetStartTime()
		appliancecertrenewalphases = append(appliancecertrenewalphases, appliancecertrenewalphase)
	}
	return appliancecertrenewalphases
}
func flattenListApplianceDataExportPolicyRelationship(p []models.ApplianceDataExportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		appliancedataexportpolicyrelationship := flattenMoMoRef(item)
		appliancedataexportpolicyrelationships = append(appliancedataexportpolicyrelationships, appliancedataexportpolicyrelationship)
	}
	return appliancedataexportpolicyrelationships
}
func flattenListApplianceFileSystemStatusRelationship(p []models.ApplianceFileSystemStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancefilesystemstatusrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		appliancefilesystemstatusrelationship := flattenMoMoRef(item)
		appliancefilesystemstatusrelationships = append(appliancefilesystemstatusrelationships, appliancefilesystemstatusrelationship)
	}
	return appliancefilesystemstatusrelationships
}
func flattenListApplianceGroupStatusRelationship(p []models.ApplianceGroupStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancegroupstatusrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		appliancegroupstatusrelationship := flattenMoMoRef(item)
		appliancegroupstatusrelationships = append(appliancegroupstatusrelationships, appliancegroupstatusrelationship)
	}
	return appliancegroupstatusrelationships
}
func flattenListApplianceKeyValuePair(p []models.ApplianceKeyValuePair, d *schema.ResourceData) []map[string]interface{} {
	var appliancekeyvaluepairs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		appliancekeyvaluepair := make(map[string]interface{})
		appliancekeyvaluepair["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		appliancekeyvaluepair["class_id"] = item.GetClassId()
		appliancekeyvaluepair["key"] = item.GetKey()
		appliancekeyvaluepair["object_type"] = item.GetObjectType()
		appliancekeyvaluepair["value"] = item.GetValue()
		appliancekeyvaluepairs = append(appliancekeyvaluepairs, appliancekeyvaluepair)
	}
	return appliancekeyvaluepairs
}
func flattenListApplianceStatusCheck(p []models.ApplianceStatusCheck, d *schema.ResourceData) []map[string]interface{} {
	var appliancestatuschecks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		appliancestatuscheck := make(map[string]interface{})
		appliancestatuscheck["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		appliancestatuscheck["class_id"] = item.GetClassId()
		appliancestatuscheck["code"] = item.GetCode()
		appliancestatuscheck["object_type"] = item.GetObjectType()
		appliancestatuscheck["result"] = item.GetResult()
		appliancestatuschecks = append(appliancestatuschecks, appliancestatuscheck)
	}
	return appliancestatuschecks
}
func flattenListAssetClusterMemberRelationship(p []models.AssetClusterMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		assetclustermemberrelationship := flattenMoMoRef(item)
		assetclustermemberrelationships = append(assetclustermemberrelationships, assetclustermemberrelationship)
	}
	return assetclustermemberrelationships
}
func flattenListAssetConnection(p []models.AssetConnection, d *schema.ResourceData) []map[string]interface{} {
	var assetconnections []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		assetconnection := make(map[string]interface{})
		assetconnection["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetconnection["class_id"] = item.GetClassId()
		assetconnection["credential"] = (func(p models.AssetCredential, d *schema.ResourceData) []map[string]interface{} {
			var assetcredentials []map[string]interface{}
			var ret models.AssetCredential
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			assetcredential := make(map[string]interface{})
			assetcredential["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetcredential["class_id"] = item.GetClassId()
			assetcredential["object_type"] = item.GetObjectType()

			assetcredentials = append(assetcredentials, assetcredential)
			return assetcredentials
		})(item.GetCredential(), d)
		assetconnection["object_type"] = item.GetObjectType()
		assetconnections = append(assetconnections, assetconnection)
	}
	return assetconnections
}
func flattenListAssetDeploymentRelationship(p []models.AssetDeploymentRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeploymentrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		assetdeploymentrelationship := flattenMoMoRef(item)
		assetdeploymentrelationships = append(assetdeploymentrelationships, assetdeploymentrelationship)
	}
	return assetdeploymentrelationships
}
func flattenListAssetDeploymentDeviceRelationship(p []models.AssetDeploymentDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeploymentdevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		assetdeploymentdevicerelationship := flattenMoMoRef(item)
		assetdeploymentdevicerelationships = append(assetdeploymentdevicerelationships, assetdeploymentdevicerelationship)
	}
	return assetdeploymentdevicerelationships
}
func flattenListAssetDeviceRegistrationRelationship(p []models.AssetDeviceRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceregistrationrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		assetdeviceregistrationrelationship := flattenMoMoRef(item)
		assetdeviceregistrationrelationships = append(assetdeviceregistrationrelationships, assetdeviceregistrationrelationship)
	}
	return assetdeviceregistrationrelationships
}
func flattenListAssetMeteringType(p []models.AssetMeteringType, d *schema.ResourceData) []map[string]interface{} {
	var assetmeteringtypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		assetmeteringtype := make(map[string]interface{})
		assetmeteringtype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetmeteringtype["class_id"] = item.GetClassId()
		assetmeteringtype["name"] = item.GetName()
		assetmeteringtype["object_type"] = item.GetObjectType()
		assetmeteringtype["unit"] = item.GetUnit()
		assetmeteringtypes = append(assetmeteringtypes, assetmeteringtype)
	}
	return assetmeteringtypes
}
func flattenListAssetService(p []models.AssetService, d *schema.ResourceData) []map[string]interface{} {
	var assetservices []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		assetservice := make(map[string]interface{})
		assetservice["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetservice["class_id"] = item.GetClassId()
		assetservice["object_type"] = item.GetObjectType()
		assetservice["options"] = (func(p models.AssetServiceOptions, d *schema.ResourceData) []map[string]interface{} {
			var assetserviceoptionss []map[string]interface{}
			var ret models.AssetServiceOptions
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			assetserviceoptions := make(map[string]interface{})
			assetserviceoptions["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetserviceoptions["class_id"] = item.GetClassId()
			assetserviceoptions["object_type"] = item.GetObjectType()

			assetserviceoptionss = append(assetserviceoptionss, assetserviceoptions)
			return assetserviceoptionss
		})(item.GetOptions(), d)
		assetservices = append(assetservices, assetservice)
	}
	return assetservices
}
func flattenListBiosBootDeviceRelationship(p []models.BiosBootDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosbootdevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		biosbootdevicerelationship := flattenMoMoRef(item)
		biosbootdevicerelationships = append(biosbootdevicerelationships, biosbootdevicerelationship)
	}
	return biosbootdevicerelationships
}
func flattenListBiosUnitRelationship(p []models.BiosUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		biosunitrelationship := flattenMoMoRef(item)
		biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	}
	return biosunitrelationships
}
func flattenListBootCddDeviceRelationship(p []models.BootCddDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootcdddevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootcdddevicerelationship := flattenMoMoRef(item)
		bootcdddevicerelationships = append(bootcdddevicerelationships, bootcdddevicerelationship)
	}
	return bootcdddevicerelationships
}
func flattenListBootDeviceBase(p []models.BootDeviceBase, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		bootdevicebase := make(map[string]interface{})
		bootdevicebase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		bootdevicebase["class_id"] = item.GetClassId()
		bootdevicebase["enabled"] = item.GetEnabled()
		bootdevicebase["name"] = item.GetName()
		bootdevicebase["object_type"] = item.GetObjectType()
		bootdevicebases = append(bootdevicebases, bootdevicebase)
	}
	return bootdevicebases
}
func flattenListBootHddDeviceRelationship(p []models.BootHddDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var boothdddevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		boothdddevicerelationship := flattenMoMoRef(item)
		boothdddevicerelationships = append(boothdddevicerelationships, boothdddevicerelationship)
	}
	return boothdddevicerelationships
}
func flattenListBootIscsiDeviceRelationship(p []models.BootIscsiDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootiscsidevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootiscsidevicerelationship := flattenMoMoRef(item)
		bootiscsidevicerelationships = append(bootiscsidevicerelationships, bootiscsidevicerelationship)
	}
	return bootiscsidevicerelationships
}
func flattenListBootNvmeDeviceRelationship(p []models.BootNvmeDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootnvmedevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootnvmedevicerelationship := flattenMoMoRef(item)
		bootnvmedevicerelationships = append(bootnvmedevicerelationships, bootnvmedevicerelationship)
	}
	return bootnvmedevicerelationships
}
func flattenListBootPchStorageDeviceRelationship(p []models.BootPchStorageDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootpchstoragedevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootpchstoragedevicerelationship := flattenMoMoRef(item)
		bootpchstoragedevicerelationships = append(bootpchstoragedevicerelationships, bootpchstoragedevicerelationship)
	}
	return bootpchstoragedevicerelationships
}
func flattenListBootPxeDeviceRelationship(p []models.BootPxeDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootpxedevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootpxedevicerelationship := flattenMoMoRef(item)
		bootpxedevicerelationships = append(bootpxedevicerelationships, bootpxedevicerelationship)
	}
	return bootpxedevicerelationships
}
func flattenListBootSanDeviceRelationship(p []models.BootSanDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootsandevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootsandevicerelationship := flattenMoMoRef(item)
		bootsandevicerelationships = append(bootsandevicerelationships, bootsandevicerelationship)
	}
	return bootsandevicerelationships
}
func flattenListBootSdDeviceRelationship(p []models.BootSdDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootsddevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootsddevicerelationship := flattenMoMoRef(item)
		bootsddevicerelationships = append(bootsddevicerelationships, bootsddevicerelationship)
	}
	return bootsddevicerelationships
}
func flattenListBootUefiShellDeviceRelationship(p []models.BootUefiShellDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootuefishelldevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootuefishelldevicerelationship := flattenMoMoRef(item)
		bootuefishelldevicerelationships = append(bootuefishelldevicerelationships, bootuefishelldevicerelationship)
	}
	return bootuefishelldevicerelationships
}
func flattenListBootUsbDeviceRelationship(p []models.BootUsbDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootusbdevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootusbdevicerelationship := flattenMoMoRef(item)
		bootusbdevicerelationships = append(bootusbdevicerelationships, bootusbdevicerelationship)
	}
	return bootusbdevicerelationships
}
func flattenListBootVmediaDeviceRelationship(p []models.BootVmediaDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootvmediadevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		bootvmediadevicerelationship := flattenMoMoRef(item)
		bootvmediadevicerelationships = append(bootvmediadevicerelationships, bootvmediadevicerelationship)
	}
	return bootvmediadevicerelationships
}
func flattenListCapabilityCapabilityRelationship(p []models.CapabilityCapabilityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var capabilitycapabilityrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		capabilitycapabilityrelationship := flattenMoMoRef(item)
		capabilitycapabilityrelationships = append(capabilitycapabilityrelationships, capabilitycapabilityrelationship)
	}
	return capabilitycapabilityrelationships
}
func flattenListCapabilityPortRange(p []models.CapabilityPortRange, d *schema.ResourceData) []map[string]interface{} {
	var capabilityportranges []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		capabilityportrange := make(map[string]interface{})
		capabilityportrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		capabilityportrange["class_id"] = item.GetClassId()
		capabilityportrange["end_port_id"] = item.GetEndPortId()
		capabilityportrange["end_slot_id"] = item.GetEndSlotId()
		capabilityportrange["object_type"] = item.GetObjectType()
		capabilityportrange["start_port_id"] = item.GetStartPortId()
		capabilityportrange["start_slot_id"] = item.GetStartSlotId()
		capabilityportranges = append(capabilityportranges, capabilityportrange)
	}
	return capabilityportranges
}
func flattenListCapabilitySwitchingModeCapability(p []models.CapabilitySwitchingModeCapability, d *schema.ResourceData) []map[string]interface{} {
	var capabilityswitchingmodecapabilitys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		capabilityswitchingmodecapability := make(map[string]interface{})
		capabilityswitchingmodecapability["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		capabilityswitchingmodecapability["class_id"] = item.GetClassId()
		capabilityswitchingmodecapability["object_type"] = item.GetObjectType()
		capabilityswitchingmodecapability["switching_mode"] = item.GetSwitchingMode()
		capabilityswitchingmodecapability["vp_compression_supported"] = item.GetVpCompressionSupported()
		capabilityswitchingmodecapabilitys = append(capabilityswitchingmodecapabilitys, capabilityswitchingmodecapability)
	}
	return capabilityswitchingmodecapabilitys
}
func flattenListCertificatemanagementCertificateBase(p []models.CertificatemanagementCertificateBase, d *schema.ResourceData) []map[string]interface{} {
	var certificatemanagementcertificatebases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		certificatemanagementcertificatebase := make(map[string]interface{})
		certificatemanagementcertificatebase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		certificatemanagementcertificatebase["certificate"] = (func(p models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
			var x509certificates []map[string]interface{}
			var ret models.X509Certificate
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			x509certificate := make(map[string]interface{})
			x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			x509certificate["class_id"] = item.GetClassId()
			x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
				var pkixdistinguishednames []map[string]interface{}
				var ret models.PkixDistinguishedName
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				pkixdistinguishedname := make(map[string]interface{})
				pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				pkixdistinguishedname["class_id"] = item.GetClassId()
				pkixdistinguishedname["common_name"] = item.GetCommonName()
				pkixdistinguishedname["country"] = item.GetCountry()
				pkixdistinguishedname["locality"] = item.GetLocality()
				pkixdistinguishedname["object_type"] = item.GetObjectType()
				pkixdistinguishedname["organization"] = item.GetOrganization()
				pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
				pkixdistinguishedname["state"] = item.GetState()

				pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
				return pkixdistinguishednames
			})(item.GetIssuer(), d)
			x509certificate["not_after"] = item.GetNotAfter()
			x509certificate["not_before"] = item.GetNotBefore()
			x509certificate["object_type"] = item.GetObjectType()
			x509certificate["pem_certificate"] = item.GetPemCertificate()
			x509certificate["sha256_fingerprint"] = item.GetSha256Fingerprint()
			x509certificate["signature_algorithm"] = item.GetSignatureAlgorithm()
			x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
				var pkixdistinguishednames []map[string]interface{}
				var ret models.PkixDistinguishedName
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				pkixdistinguishedname := make(map[string]interface{})
				pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				pkixdistinguishedname["class_id"] = item.GetClassId()
				pkixdistinguishedname["common_name"] = item.GetCommonName()
				pkixdistinguishedname["country"] = item.GetCountry()
				pkixdistinguishedname["locality"] = item.GetLocality()
				pkixdistinguishedname["object_type"] = item.GetObjectType()
				pkixdistinguishedname["organization"] = item.GetOrganization()
				pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
				pkixdistinguishedname["state"] = item.GetState()

				pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
				return pkixdistinguishednames
			})(item.GetSubject(), d)

			x509certificates = append(x509certificates, x509certificate)
			return x509certificates
		})(item.GetCertificate(), d)
		certificatemanagementcertificatebase["class_id"] = item.GetClassId()
		certificatemanagementcertificatebase["enabled"] = item.GetEnabled()
		certificatemanagementcertificatebase["is_privatekey_set"] = item.GetIsPrivatekeySet()
		certificatemanagementcertificatebase["object_type"] = item.GetObjectType()
		privatekey_x, exists := d.GetOk("certificates")
		if exists && privatekey_x != nil {
			certificatemanagementcertificatebase["privatekey"] = privatekey_x.([]interface{})[len(certificatemanagementcertificatebases)].(map[string]interface{})["privatekey"]
		}
		certificatemanagementcertificatebases = append(certificatemanagementcertificatebases, certificatemanagementcertificatebase)
	}
	return certificatemanagementcertificatebases
}
func flattenListChassisConfigChangeDetailRelationship(p []models.ChassisConfigChangeDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var chassisconfigchangedetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		chassisconfigchangedetailrelationship := flattenMoMoRef(item)
		chassisconfigchangedetailrelationships = append(chassisconfigchangedetailrelationships, chassisconfigchangedetailrelationship)
	}
	return chassisconfigchangedetailrelationships
}
func flattenListChassisConfigResultEntryRelationship(p []models.ChassisConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var chassisconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		chassisconfigresultentryrelationship := flattenMoMoRef(item)
		chassisconfigresultentryrelationships = append(chassisconfigresultentryrelationships, chassisconfigresultentryrelationship)
	}
	return chassisconfigresultentryrelationships
}
func flattenListChassisIomProfileRelationship(p []models.ChassisIomProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var chassisiomprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		chassisiomprofilerelationship := flattenMoMoRef(item)
		chassisiomprofilerelationships = append(chassisiomprofilerelationships, chassisiomprofilerelationship)
	}
	return chassisiomprofilerelationships
}
func flattenListCloudCustomAttributes(p []models.CloudCustomAttributes, d *schema.ResourceData) []map[string]interface{} {
	var cloudcustomattributess []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		cloudcustomattributes := make(map[string]interface{})
		cloudcustomattributes["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		cloudcustomattributes["attribute_name"] = item.GetAttributeName()
		cloudcustomattributes["attribute_type"] = item.GetAttributeType()
		cloudcustomattributes["attribute_value"] = item.GetAttributeValue()
		cloudcustomattributes["class_id"] = item.GetClassId()
		cloudcustomattributes["object_type"] = item.GetObjectType()
		cloudcustomattributess = append(cloudcustomattributess, cloudcustomattributes)
	}
	return cloudcustomattributess
}
func flattenListComputeBladeRelationship(p []models.ComputeBladeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computebladerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		computebladerelationship := flattenMoMoRef(item)
		computebladerelationships = append(computebladerelationships, computebladerelationship)
	}
	return computebladerelationships
}
func flattenListComputeIpAddress(p []models.ComputeIpAddress, d *schema.ResourceData) []map[string]interface{} {
	var computeipaddresss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		computeipaddress := make(map[string]interface{})
		computeipaddress["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		computeipaddress["address"] = item.GetAddress()
		computeipaddress["category"] = item.GetCategory()
		computeipaddress["class_id"] = item.GetClassId()
		computeipaddress["default_gateway"] = item.GetDefaultGateway()
		computeipaddress["dn"] = item.GetDn()
		computeipaddress["http_port"] = item.GetHttpPort()
		computeipaddress["https_port"] = item.GetHttpsPort()
		computeipaddress["kvm_port"] = item.GetKvmPort()
		computeipaddress["kvm_vlan"] = item.GetKvmVlan()
		computeipaddress["name"] = item.GetName()
		computeipaddress["object_type"] = item.GetObjectType()
		computeipaddress["subnet"] = item.GetSubnet()
		computeipaddress["type"] = item.GetType()
		computeipaddresss = append(computeipaddresss, computeipaddress)
	}
	return computeipaddresss
}
func flattenListComputeMappingRelationship(p []models.ComputeMappingRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computemappingrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		computemappingrelationship := flattenMoMoRef(item)
		computemappingrelationships = append(computemappingrelationships, computemappingrelationship)
	}
	return computemappingrelationships
}
func flattenListComputePhysicalRelationship(p []models.ComputePhysicalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computephysicalrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		computephysicalrelationship := flattenMoMoRef(item)
		computephysicalrelationships = append(computephysicalrelationships, computephysicalrelationship)
	}
	return computephysicalrelationships
}
func flattenListComputeRackUnitRelationship(p []models.ComputeRackUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		computerackunitrelationship := flattenMoMoRef(item)
		computerackunitrelationships = append(computerackunitrelationships, computerackunitrelationship)
	}
	return computerackunitrelationships
}
func flattenListCondHclStatusDetailRelationship(p []models.CondHclStatusDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusdetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		condhclstatusdetailrelationship := flattenMoMoRef(item)
		condhclstatusdetailrelationships = append(condhclstatusdetailrelationships, condhclstatusdetailrelationship)
	}
	return condhclstatusdetailrelationships
}
func flattenListConfigExportedItemRelationship(p []models.ConfigExportedItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configexporteditemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		configexporteditemrelationship := flattenMoMoRef(item)
		configexporteditemrelationships = append(configexporteditemrelationships, configexporteditemrelationship)
	}
	return configexporteditemrelationships
}
func flattenListConfigImportedItemRelationship(p []models.ConfigImportedItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configimporteditemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		configimporteditemrelationship := flattenMoMoRef(item)
		configimporteditemrelationships = append(configimporteditemrelationships, configimporteditemrelationship)
	}
	return configimporteditemrelationships
}
func flattenListConfigMoRef(p []models.ConfigMoRef, d *schema.ResourceData) []map[string]interface{} {
	var configmorefs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		configmoref := make(map[string]interface{})
		configmoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		configmoref["class_id"] = item.GetClassId()
		configmoref["moid"] = item.GetMoid()
		configmoref["object_type"] = item.GetObjectType()
		configmorefs = append(configmorefs, configmoref)
	}
	return configmorefs
}
func flattenListConnectorpackConnectorPackUpdate(p []models.ConnectorpackConnectorPackUpdate, d *schema.ResourceData) []map[string]interface{} {
	var connectorpackconnectorpackupdates []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		connectorpackconnectorpackupdate := make(map[string]interface{})
		connectorpackconnectorpackupdate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		connectorpackconnectorpackupdate["class_id"] = item.GetClassId()
		connectorpackconnectorpackupdate["current_version"] = item.GetCurrentVersion()
		connectorpackconnectorpackupdate["name"] = item.GetName()
		connectorpackconnectorpackupdate["new_version"] = item.GetNewVersion()
		connectorpackconnectorpackupdate["object_type"] = item.GetObjectType()
		connectorpackconnectorpackupdates = append(connectorpackconnectorpackupdates, connectorpackconnectorpackupdate)
	}
	return connectorpackconnectorpackupdates
}
func flattenListContentComplexType(p []models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
	var contentcomplextypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		contentcomplextype := make(map[string]interface{})
		contentcomplextype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		contentcomplextype["class_id"] = item.GetClassId()
		contentcomplextype["name"] = item.GetName()
		contentcomplextype["object_type"] = item.GetObjectType()
		contentcomplextype["parameters"] = (func(p []models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
			var contentbaseparameters []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				contentbaseparameter := make(map[string]interface{})
				contentbaseparameter["accept_single_value"] = item.GetAcceptSingleValue()
				contentbaseparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				contentbaseparameter["class_id"] = item.GetClassId()
				contentbaseparameter["complex_type"] = item.GetComplexType()
				contentbaseparameter["item_type"] = item.GetItemType()
				contentbaseparameter["name"] = item.GetName()
				contentbaseparameter["object_type"] = item.GetObjectType()
				contentbaseparameter["path"] = item.GetPath()
				contentbaseparameter["secure"] = item.GetSecure()
				contentbaseparameter["type"] = item.GetType()
				contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
			}
			return contentbaseparameters
		})(item.GetParameters(), d)
		contentcomplextypes = append(contentcomplextypes, contentcomplextype)
	}
	return contentcomplextypes
}
func flattenListContentParameter(p []models.ContentParameter, d *schema.ResourceData) []map[string]interface{} {
	var contentparameters []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		contentparameter := make(map[string]interface{})
		contentparameter["accept_single_value"] = item.GetAcceptSingleValue()
		contentparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		contentparameter["class_id"] = item.GetClassId()
		contentparameter["complex_type"] = item.GetComplexType()
		contentparameter["item_type"] = item.GetItemType()
		contentparameter["name"] = item.GetName()
		contentparameter["object_type"] = item.GetObjectType()
		contentparameter["path"] = item.GetPath()
		contentparameter["secure"] = item.GetSecure()
		contentparameter["type"] = item.GetType()
		contentparameters = append(contentparameters, contentparameter)
	}
	return contentparameters
}
func flattenListCrdCustomResourceConfigProperty(p []models.CrdCustomResourceConfigProperty, d *schema.ResourceData) []map[string]interface{} {
	var crdcustomresourceconfigpropertys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		crdcustomresourceconfigproperty := make(map[string]interface{})
		crdcustomresourceconfigproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		crdcustomresourceconfigproperty["class_id"] = item.GetClassId()
		crdcustomresourceconfigproperty["key"] = item.GetKey()
		crdcustomresourceconfigproperty["object_type"] = item.GetObjectType()
		crdcustomresourceconfigproperty["value"] = item.GetValue()
		crdcustomresourceconfigpropertys = append(crdcustomresourceconfigpropertys, crdcustomresourceconfigproperty)
	}
	return crdcustomresourceconfigpropertys
}
func flattenListEquipmentFanRelationship(p []models.EquipmentFanRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentfanrelationship := flattenMoMoRef(item)
		equipmentfanrelationships = append(equipmentfanrelationships, equipmentfanrelationship)
	}
	return equipmentfanrelationships
}
func flattenListEquipmentFanModuleRelationship(p []models.EquipmentFanModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentfanmodulerelationship := flattenMoMoRef(item)
		equipmentfanmodulerelationships = append(equipmentfanmodulerelationships, equipmentfanmodulerelationship)
	}
	return equipmentfanmodulerelationships
}
func flattenListEquipmentIoCardRelationship(p []models.EquipmentIoCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentiocardrelationship := flattenMoMoRef(item)
		equipmentiocardrelationships = append(equipmentiocardrelationships, equipmentiocardrelationship)
	}
	return equipmentiocardrelationships
}
func flattenListEquipmentIoCardIdentity(p []models.EquipmentIoCardIdentity, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardidentitys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		equipmentiocardidentity := make(map[string]interface{})
		equipmentiocardidentity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		equipmentiocardidentity["class_id"] = item.GetClassId()
		equipmentiocardidentity["io_card_moid"] = item.GetIoCardMoid()
		equipmentiocardidentity["module_id"] = item.GetModuleId()
		equipmentiocardidentity["network_element_moid"] = item.GetNetworkElementMoid()
		equipmentiocardidentity["object_type"] = item.GetObjectType()
		equipmentiocardidentity["switch_id"] = item.GetSwitchId()
		equipmentiocardidentitys = append(equipmentiocardidentitys, equipmentiocardidentity)
	}
	return equipmentiocardidentitys
}
func flattenListEquipmentIoExpanderRelationship(p []models.EquipmentIoExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentioexpanderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentioexpanderrelationship := flattenMoMoRef(item)
		equipmentioexpanderrelationships = append(equipmentioexpanderrelationships, equipmentioexpanderrelationship)
	}
	return equipmentioexpanderrelationships
}
func flattenListEquipmentPsuRelationship(p []models.EquipmentPsuRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentpsurelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentpsurelationship := flattenMoMoRef(item)
		equipmentpsurelationships = append(equipmentpsurelationships, equipmentpsurelationship)
	}
	return equipmentpsurelationships
}
func flattenListEquipmentRackEnclosureSlotRelationship(p []models.EquipmentRackEnclosureSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentrackenclosureslotrelationship := flattenMoMoRef(item)
		equipmentrackenclosureslotrelationships = append(equipmentrackenclosureslotrelationships, equipmentrackenclosureslotrelationship)
	}
	return equipmentrackenclosureslotrelationships
}
func flattenListEquipmentSwitchCardRelationship(p []models.EquipmentSwitchCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentswitchcardrelationship := flattenMoMoRef(item)
		equipmentswitchcardrelationships = append(equipmentswitchcardrelationships, equipmentswitchcardrelationship)
	}
	return equipmentswitchcardrelationships
}
func flattenListEquipmentSystemIoControllerRelationship(p []models.EquipmentSystemIoControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentsystemiocontrollerrelationship := flattenMoMoRef(item)
		equipmentsystemiocontrollerrelationships = append(equipmentsystemiocontrollerrelationships, equipmentsystemiocontrollerrelationship)
	}
	return equipmentsystemiocontrollerrelationships
}
func flattenListEquipmentTpmRelationship(p []models.EquipmentTpmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmenttpmrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmenttpmrelationship := flattenMoMoRef(item)
		equipmenttpmrelationships = append(equipmenttpmrelationships, equipmenttpmrelationship)
	}
	return equipmenttpmrelationships
}
func flattenListEtherHostPortRelationship(p []models.EtherHostPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherhostportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		etherhostportrelationship := flattenMoMoRef(item)
		etherhostportrelationships = append(etherhostportrelationships, etherhostportrelationship)
	}
	return etherhostportrelationships
}
func flattenListEtherNetworkPortRelationship(p []models.EtherNetworkPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ethernetworkportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ethernetworkportrelationship := flattenMoMoRef(item)
		ethernetworkportrelationships = append(ethernetworkportrelationships, ethernetworkportrelationship)
	}
	return ethernetworkportrelationships
}
func flattenListEtherPhysicalPortRelationship(p []models.EtherPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		etherphysicalportrelationship := flattenMoMoRef(item)
		etherphysicalportrelationships = append(etherphysicalportrelationships, etherphysicalportrelationship)
	}
	return etherphysicalportrelationships
}
func flattenListEtherPortChannelRelationship(p []models.EtherPortChannelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherportchannelrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		etherportchannelrelationship := flattenMoMoRef(item)
		etherportchannelrelationships = append(etherportchannelrelationships, etherportchannelrelationship)
	}
	return etherportchannelrelationships
}
func flattenListFabricConfigChangeDetailRelationship(p []models.FabricConfigChangeDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricconfigchangedetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricconfigchangedetailrelationship := flattenMoMoRef(item)
		fabricconfigchangedetailrelationships = append(fabricconfigchangedetailrelationships, fabricconfigchangedetailrelationship)
	}
	return fabricconfigchangedetailrelationships
}
func flattenListFabricConfigResultEntryRelationship(p []models.FabricConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricconfigresultentryrelationship := flattenMoMoRef(item)
		fabricconfigresultentryrelationships = append(fabricconfigresultentryrelationships, fabricconfigresultentryrelationship)
	}
	return fabricconfigresultentryrelationships
}
func flattenListFabricEthNetworkGroupPolicyRelationship(p []models.FabricEthNetworkGroupPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricethnetworkgrouppolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricethnetworkgrouppolicyrelationship := flattenMoMoRef(item)
		fabricethnetworkgrouppolicyrelationships = append(fabricethnetworkgrouppolicyrelationships, fabricethnetworkgrouppolicyrelationship)
	}
	return fabricethnetworkgrouppolicyrelationships
}
func flattenListFabricPortIdentifier(p []models.FabricPortIdentifier, d *schema.ResourceData) []map[string]interface{} {
	var fabricportidentifiers []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		fabricportidentifier := make(map[string]interface{})
		fabricportidentifier["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		fabricportidentifier["aggregate_port_id"] = item.GetAggregatePortId()
		fabricportidentifier["class_id"] = item.GetClassId()
		fabricportidentifier["object_type"] = item.GetObjectType()
		fabricportidentifier["port_id"] = item.GetPortId()
		fabricportidentifier["slot_id"] = item.GetSlotId()
		fabricportidentifiers = append(fabricportidentifiers, fabricportidentifier)
	}
	return fabricportidentifiers
}
func flattenListFabricQosClass(p []models.FabricQosClass, d *schema.ResourceData) []map[string]interface{} {
	var fabricqosclasss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		fabricqosclass := make(map[string]interface{})
		fabricqosclass["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		fabricqosclass["admin_state"] = item.GetAdminState()
		fabricqosclass["bandwidth_percent"] = item.GetBandwidthPercent()
		fabricqosclass["class_id"] = item.GetClassId()
		fabricqosclass["cos"] = item.GetCos()
		fabricqosclass["mtu"] = item.GetMtu()
		fabricqosclass["multicast_optimize"] = item.GetMulticastOptimize()
		fabricqosclass["name"] = item.GetName()
		fabricqosclass["object_type"] = item.GetObjectType()
		fabricqosclass["packet_drop"] = item.GetPacketDrop()
		fabricqosclass["weight"] = item.GetWeight()
		fabricqosclasss = append(fabricqosclasss, fabricqosclass)
	}
	return fabricqosclasss
}
func flattenListFabricSwitchProfileRelationship(p []models.FabricSwitchProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricswitchprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricswitchprofilerelationship := flattenMoMoRef(item)
		fabricswitchprofilerelationships = append(fabricswitchprofilerelationships, fabricswitchprofilerelationship)
	}
	return fabricswitchprofilerelationships
}
func flattenListFcPhysicalPortRelationship(p []models.FcPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcphysicalportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fcphysicalportrelationship := flattenMoMoRef(item)
		fcphysicalportrelationships = append(fcphysicalportrelationships, fcphysicalportrelationship)
	}
	return fcphysicalportrelationships
}
func flattenListFcPortChannelRelationship(p []models.FcPortChannelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcportchannelrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fcportchannelrelationship := flattenMoMoRef(item)
		fcportchannelrelationships = append(fcportchannelrelationships, fcportchannelrelationship)
	}
	return fcportchannelrelationships
}
func flattenListFcpoolBlock(p []models.FcpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		fcpoolblock := make(map[string]interface{})
		fcpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		fcpoolblock["class_id"] = item.GetClassId()
		fcpoolblock["from"] = item.GetFrom()
		fcpoolblock["object_type"] = item.GetObjectType()
		fcpoolblock["size"] = item.GetSize()
		fcpoolblock["to"] = item.GetTo()
		fcpoolblocks = append(fcpoolblocks, fcpoolblock)
	}
	return fcpoolblocks
}
func flattenListFcpoolFcBlockRelationship(p []models.FcpoolFcBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolfcblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fcpoolfcblockrelationship := flattenMoMoRef(item)
		fcpoolfcblockrelationships = append(fcpoolfcblockrelationships, fcpoolfcblockrelationship)
	}
	return fcpoolfcblockrelationships
}
func flattenListFirmwareBaseImpact(p []models.FirmwareBaseImpact, d *schema.ResourceData) []map[string]interface{} {
	var firmwarebaseimpacts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		firmwarebaseimpact := make(map[string]interface{})
		firmwarebaseimpact["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarebaseimpact["class_id"] = item.GetClassId()
		firmwarebaseimpact["computation_error"] = item.GetComputationError()
		firmwarebaseimpact["computation_status_detail"] = item.GetComputationStatusDetail()
		firmwarebaseimpact["domain_name"] = item.GetDomainName()
		firmwarebaseimpact["end_point"] = item.GetEndPoint()
		firmwarebaseimpact["firmware_version"] = item.GetFirmwareVersion()
		firmwarebaseimpact["impact_type"] = item.GetImpactType()
		firmwarebaseimpact["model"] = item.GetModel()
		firmwarebaseimpact["object_type"] = item.GetObjectType()
		firmwarebaseimpact["target_firmware_version"] = item.GetTargetFirmwareVersion()
		firmwarebaseimpact["version_impact"] = item.GetVersionImpact()
		firmwarebaseimpacts = append(firmwarebaseimpacts, firmwarebaseimpact)
	}
	return firmwarebaseimpacts
}
func flattenListFirmwareComponentMeta(p []models.FirmwareComponentMeta, d *schema.ResourceData) []map[string]interface{} {
	var firmwarecomponentmetas []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		firmwarecomponentmeta := make(map[string]interface{})
		firmwarecomponentmeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarecomponentmeta["class_id"] = item.GetClassId()
		firmwarecomponentmeta["component_label"] = item.GetComponentLabel()
		firmwarecomponentmeta["component_type"] = item.GetComponentType()
		firmwarecomponentmeta["description"] = item.GetDescription()
		firmwarecomponentmeta["disruption"] = item.GetDisruption()
		firmwarecomponentmeta["image_path"] = item.GetImagePath()
		firmwarecomponentmeta["is_oob_supported"] = item.GetIsOobSupported()
		firmwarecomponentmeta["model"] = item.GetModel()
		firmwarecomponentmeta["object_type"] = item.GetObjectType()
		firmwarecomponentmeta["oob_manageability"] = item.GetOobManageability()
		firmwarecomponentmeta["packed_version"] = item.GetPackedVersion()
		firmwarecomponentmeta["redfish_url"] = item.GetRedfishUrl()
		firmwarecomponentmeta["vendor"] = item.GetVendor()
		firmwarecomponentmetas = append(firmwarecomponentmetas, firmwarecomponentmeta)
	}
	return firmwarecomponentmetas
}
func flattenListFirmwareDistributableMetaRelationship(p []models.FirmwareDistributableMetaRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredistributablemetarelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		firmwaredistributablemetarelationship := flattenMoMoRef(item)
		firmwaredistributablemetarelationships = append(firmwaredistributablemetarelationships, firmwaredistributablemetarelationship)
	}
	return firmwaredistributablemetarelationships
}
func flattenListFirmwareFirmwareInventory(p []models.FirmwareFirmwareInventory, d *schema.ResourceData) []map[string]interface{} {
	var firmwarefirmwareinventorys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		firmwarefirmwareinventory := make(map[string]interface{})
		firmwarefirmwareinventory["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarefirmwareinventory["category"] = item.GetCategory()
		firmwarefirmwareinventory["class_id"] = item.GetClassId()
		firmwarefirmwareinventory["label"] = item.GetLabel()
		firmwarefirmwareinventory["model"] = item.GetModel()
		firmwarefirmwareinventory["object_type"] = item.GetObjectType()
		firmwarefirmwareinventory["update_uri"] = item.GetUpdateUri()
		firmwarefirmwareinventory["vendor"] = item.GetVendor()
		firmwarefirmwareinventory["nr_version"] = item.GetVersion()
		firmwarefirmwareinventorys = append(firmwarefirmwareinventorys, firmwarefirmwareinventory)
	}
	return firmwarefirmwareinventorys
}
func flattenListFirmwareRunningFirmwareRelationship(p []models.FirmwareRunningFirmwareRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		firmwarerunningfirmwarerelationship := flattenMoMoRef(item)
		firmwarerunningfirmwarerelationships = append(firmwarerunningfirmwarerelationships, firmwarerunningfirmwarerelationship)
	}
	return firmwarerunningfirmwarerelationships
}
func flattenListForecastDefinitionRelationship(p []models.ForecastDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		forecastdefinitionrelationship := flattenMoMoRef(item)
		forecastdefinitionrelationships = append(forecastdefinitionrelationships, forecastdefinitionrelationship)
	}
	return forecastdefinitionrelationships
}
func flattenListGraphicsCardRelationship(p []models.GraphicsCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		graphicscardrelationship := flattenMoMoRef(item)
		graphicscardrelationships = append(graphicscardrelationships, graphicscardrelationship)
	}
	return graphicscardrelationships
}
func flattenListGraphicsControllerRelationship(p []models.GraphicsControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		graphicscontrollerrelationship := flattenMoMoRef(item)
		graphicscontrollerrelationships = append(graphicscontrollerrelationships, graphicscontrollerrelationship)
	}
	return graphicscontrollerrelationships
}
func flattenListHclConstraint(p []models.HclConstraint, d *schema.ResourceData) []map[string]interface{} {
	var hclconstraints []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hclconstraint := make(map[string]interface{})
		hclconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hclconstraint["class_id"] = item.GetClassId()
		hclconstraint["constraint_name"] = item.GetConstraintName()
		hclconstraint["constraint_value"] = item.GetConstraintValue()
		hclconstraint["object_type"] = item.GetObjectType()
		hclconstraints = append(hclconstraints, hclconstraint)
	}
	return hclconstraints
}
func flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p []models.HclHyperflexSoftwareCompatibilityInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hclhyperflexsoftwarecompatibilityinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hclhyperflexsoftwarecompatibilityinforelationship := flattenMoMoRef(item)
		hclhyperflexsoftwarecompatibilityinforelationships = append(hclhyperflexsoftwarecompatibilityinforelationships, hclhyperflexsoftwarecompatibilityinforelationship)
	}
	return hclhyperflexsoftwarecompatibilityinforelationships
}
func flattenListHclOperatingSystemRelationship(p []models.HclOperatingSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hcloperatingsystemrelationship := flattenMoMoRef(item)
		hcloperatingsystemrelationships = append(hcloperatingsystemrelationships, hcloperatingsystemrelationship)
	}
	return hcloperatingsystemrelationships
}
func flattenListHyperflexAlarmRelationship(p []models.HyperflexAlarmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexalarmrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexalarmrelationship := flattenMoMoRef(item)
		hyperflexalarmrelationships = append(hyperflexalarmrelationships, hyperflexalarmrelationship)
	}
	return hyperflexalarmrelationships
}
func flattenListHyperflexBaseClusterRelationship(p []models.HyperflexBaseClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexbaseclusterrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexbaseclusterrelationship := flattenMoMoRef(item)
		hyperflexbaseclusterrelationships = append(hyperflexbaseclusterrelationships, hyperflexbaseclusterrelationship)
	}
	return hyperflexbaseclusterrelationships
}
func flattenListHyperflexCapabilityInfoRelationship(p []models.HyperflexCapabilityInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexcapabilityinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexcapabilityinforelationship := flattenMoMoRef(item)
		hyperflexcapabilityinforelationships = append(hyperflexcapabilityinforelationships, hyperflexcapabilityinforelationship)
	}
	return hyperflexcapabilityinforelationships
}
func flattenListHyperflexClusterProfileRelationship(p []models.HyperflexClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexclusterprofilerelationship := flattenMoMoRef(item)
		hyperflexclusterprofilerelationships = append(hyperflexclusterprofilerelationships, hyperflexclusterprofilerelationship)
	}
	return hyperflexclusterprofilerelationships
}
func flattenListHyperflexConfigResultEntryRelationship(p []models.HyperflexConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexconfigresultentryrelationship := flattenMoMoRef(item)
		hyperflexconfigresultentryrelationships = append(hyperflexconfigresultentryrelationships, hyperflexconfigresultentryrelationship)
	}
	return hyperflexconfigresultentryrelationships
}
func flattenListHyperflexFeatureLimitEntry(p []models.HyperflexFeatureLimitEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitentrys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexfeaturelimitentry := make(map[string]interface{})
		hyperflexfeaturelimitentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexfeaturelimitentry["class_id"] = item.GetClassId()
		hyperflexfeaturelimitentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.GetClassId()
			hyperflexappsettingconstraint["hxdp_version"] = item.GetHxdpVersion()
			hyperflexappsettingconstraint["hypervisor_type"] = item.GetHypervisorType()
			hyperflexappsettingconstraint["mgmt_platform"] = item.GetMgmtPlatform()
			hyperflexappsettingconstraint["object_type"] = item.GetObjectType()
			hyperflexappsettingconstraint["server_model"] = item.GetServerModel()

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexfeaturelimitentry["name"] = item.GetName()
		hyperflexfeaturelimitentry["object_type"] = item.GetObjectType()
		hyperflexfeaturelimitentry["value"] = item.GetValue()
		hyperflexfeaturelimitentrys = append(hyperflexfeaturelimitentrys, hyperflexfeaturelimitentry)
	}
	return hyperflexfeaturelimitentrys
}
func flattenListHyperflexHealthCheckScriptInfo(p []models.HyperflexHealthCheckScriptInfo, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthcheckscriptinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexhealthcheckscriptinfo := make(map[string]interface{})
		hyperflexhealthcheckscriptinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhealthcheckscriptinfo["aggregate_script_name"] = item.GetAggregateScriptName()
		hyperflexhealthcheckscriptinfo["class_id"] = item.GetClassId()
		hyperflexhealthcheckscriptinfo["hyperflex_version"] = item.GetHyperflexVersion()
		hyperflexhealthcheckscriptinfo["object_type"] = item.GetObjectType()
		hyperflexhealthcheckscriptinfo["script_execute_location"] = item.GetScriptExecuteLocation()
		hyperflexhealthcheckscriptinfo["script_input"] = item.GetScriptInput()
		hyperflexhealthcheckscriptinfo["script_name"] = item.GetScriptName()
		hyperflexhealthcheckscriptinfo["nr_version"] = item.GetVersion()
		hyperflexhealthcheckscriptinfos = append(hyperflexhealthcheckscriptinfos, hyperflexhealthcheckscriptinfo)
	}
	return hyperflexhealthcheckscriptinfos
}
func flattenListHyperflexHxHostMountStatusDt(p []models.HyperflexHxHostMountStatusDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxhostmountstatusdts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexhxhostmountstatusdt := make(map[string]interface{})
		hyperflexhxhostmountstatusdt["accessibility"] = item.GetAccessibility()
		hyperflexhxhostmountstatusdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhxhostmountstatusdt["class_id"] = item.GetClassId()
		hyperflexhxhostmountstatusdt["host_name"] = item.GetHostName()
		hyperflexhxhostmountstatusdt["mounted"] = item.GetMounted()
		hyperflexhxhostmountstatusdt["object_type"] = item.GetObjectType()
		hyperflexhxhostmountstatusdt["reason"] = item.GetReason()
		hyperflexhxhostmountstatusdts = append(hyperflexhxhostmountstatusdts, hyperflexhxhostmountstatusdt)
	}
	return hyperflexhxhostmountstatusdts
}
func flattenListHyperflexHxZoneResiliencyInfoDt(p []models.HyperflexHxZoneResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxzoneresiliencyinfodts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexhxzoneresiliencyinfodt := make(map[string]interface{})
		hyperflexhxzoneresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhxzoneresiliencyinfodt["class_id"] = item.GetClassId()
		hyperflexhxzoneresiliencyinfodt["name"] = item.GetName()
		hyperflexhxzoneresiliencyinfodt["object_type"] = item.GetObjectType()
		hyperflexhxzoneresiliencyinfodt["resiliency_info"] = (func(p models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexhxresiliencyinfodts []map[string]interface{}
			var ret models.HyperflexHxResiliencyInfoDt
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexhxresiliencyinfodt := make(map[string]interface{})
			hyperflexhxresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexhxresiliencyinfodt["class_id"] = item.GetClassId()
			hyperflexhxresiliencyinfodt["data_replication_factor"] = item.GetDataReplicationFactor()
			hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.GetHddFailuresTolerable()
			hyperflexhxresiliencyinfodt["messages"] = item.GetMessages()
			hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.GetNodeFailuresTolerable()
			hyperflexhxresiliencyinfodt["object_type"] = item.GetObjectType()
			hyperflexhxresiliencyinfodt["policy_compliance"] = item.GetPolicyCompliance()
			hyperflexhxresiliencyinfodt["resiliency_state"] = item.GetResiliencyState()
			hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.GetSsdFailuresTolerable()

			hyperflexhxresiliencyinfodts = append(hyperflexhxresiliencyinfodts, hyperflexhxresiliencyinfodt)
			return hyperflexhxresiliencyinfodts
		})(item.GetResiliencyInfo(), d)
		hyperflexhxzoneresiliencyinfodts = append(hyperflexhxzoneresiliencyinfodts, hyperflexhxzoneresiliencyinfodt)
	}
	return hyperflexhxzoneresiliencyinfodts
}
func flattenListHyperflexHxapHostRelationship(p []models.HyperflexHxapHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxaphostrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexhxaphostrelationship := flattenMoMoRef(item)
		hyperflexhxaphostrelationships = append(hyperflexhxaphostrelationships, hyperflexhxaphostrelationship)
	}
	return hyperflexhxaphostrelationships
}
func flattenListHyperflexHxapHostInterfaceRelationship(p []models.HyperflexHxapHostInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxaphostinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexhxaphostinterfacerelationship := flattenMoMoRef(item)
		hyperflexhxaphostinterfacerelationships = append(hyperflexhxaphostinterfacerelationships, hyperflexhxaphostinterfacerelationship)
	}
	return hyperflexhxaphostinterfacerelationships
}
func flattenListHyperflexHxapHostVswitchRelationship(p []models.HyperflexHxapHostVswitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxaphostvswitchrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexhxaphostvswitchrelationship := flattenMoMoRef(item)
		hyperflexhxaphostvswitchrelationships = append(hyperflexhxaphostvswitchrelationships, hyperflexhxaphostvswitchrelationship)
	}
	return hyperflexhxaphostvswitchrelationships
}
func flattenListHyperflexHxdpVersionRelationship(p []models.HyperflexHxdpVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxdpversionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexhxdpversionrelationship := flattenMoMoRef(item)
		hyperflexhxdpversionrelationships = append(hyperflexhxdpversionrelationships, hyperflexhxdpversionrelationship)
	}
	return hyperflexhxdpversionrelationships
}
func flattenListHyperflexIpAddrRange(p []models.HyperflexIpAddrRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexipaddrranges []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexipaddrrange := make(map[string]interface{})
		hyperflexipaddrrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexipaddrrange["class_id"] = item.GetClassId()
		hyperflexipaddrrange["end_addr"] = item.GetEndAddr()
		hyperflexipaddrrange["gateway"] = item.GetGateway()
		hyperflexipaddrrange["netmask"] = item.GetNetmask()
		hyperflexipaddrrange["object_type"] = item.GetObjectType()
		hyperflexipaddrrange["start_addr"] = item.GetStartAddr()
		hyperflexipaddrranges = append(hyperflexipaddrranges, hyperflexipaddrrange)
	}
	return hyperflexipaddrranges
}
func flattenListHyperflexMapClusterIdToProtectionInfo(p []models.HyperflexMapClusterIdToProtectionInfo, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexmapclusteridtoprotectioninfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexmapclusteridtoprotectioninfo := make(map[string]interface{})
		hyperflexmapclusteridtoprotectioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexmapclusteridtoprotectioninfo["class_id"] = item.GetClassId()
		hyperflexmapclusteridtoprotectioninfo["cluster_id"] = item.GetClusterId()
		hyperflexmapclusteridtoprotectioninfo["object_type"] = item.GetObjectType()
		hyperflexmapclusteridtoprotectioninfo["protection_info"] = (func(p models.HyperflexProtectionInfo, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexprotectioninfos []map[string]interface{}
			var ret models.HyperflexProtectionInfo
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexprotectioninfo := make(map[string]interface{})
			hyperflexprotectioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexprotectioninfo["class_id"] = item.GetClassId()
			hyperflexprotectioninfo["object_type"] = item.GetObjectType()
			hyperflexprotectioninfo["vm_current_protection_info"] = (func(p models.HyperflexSnapshotInfoBrief, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexsnapshotinfobriefs []map[string]interface{}
				var ret models.HyperflexSnapshotInfoBrief
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexsnapshotinfobrief := make(map[string]interface{})
				hyperflexsnapshotinfobrief["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexsnapshotinfobrief["class_id"] = item.GetClassId()
				hyperflexsnapshotinfobrief["object_type"] = item.GetObjectType()
				hyperflexsnapshotinfobrief["replication_status"] = (func(p models.HyperflexReplicationStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexreplicationstatuss []map[string]interface{}
					var ret models.HyperflexReplicationStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexreplicationstatus := make(map[string]interface{})
					hyperflexreplicationstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexreplicationstatus["bytes_replicated"] = item.GetBytesReplicated()
					hyperflexreplicationstatus["class_id"] = item.GetClassId()
					hyperflexreplicationstatus["end_time"] = item.GetEndTime()
					hyperflexreplicationstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.GetClassId()
						hyperflexerrorstack["message"] = item.GetMessage()
						hyperflexerrorstack["message_id"] = item.GetMessageId()
						hyperflexerrorstack["object_type"] = item.GetObjectType()

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexreplicationstatus["object_type"] = item.GetObjectType()
					hyperflexreplicationstatus["pack_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexentityreferences []map[string]interface{}
						var ret models.HyperflexEntityReference
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexentityreference := make(map[string]interface{})
						hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexentityreference["class_id"] = item.GetClassId()
						hyperflexentityreference["confignum"] = item.GetConfignum()
						hyperflexentityreference["id"] = item.GetId()
						hyperflexentityreference["idtype"] = item.GetIdtype()
						hyperflexentityreference["name"] = item.GetName()
						hyperflexentityreference["object_type"] = item.GetObjectType()
						hyperflexentityreference["type"] = item.GetType()

						hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
						return hyperflexentityreferences
					})(item.GetPackEntityReference(), d)
					hyperflexreplicationstatus["pct_complete"] = item.GetPctComplete()
					hyperflexreplicationstatus["rpo_status"] = (func(p models.HyperflexRpoStatus, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexrpostatuss []map[string]interface{}
						var ret models.HyperflexRpoStatus
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexrpostatus := make(map[string]interface{})
						hyperflexrpostatus["actual"] = item.GetActual()
						hyperflexrpostatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexrpostatus["class_id"] = item.GetClassId()
						hyperflexrpostatus["expected"] = item.GetExpected()
						hyperflexrpostatus["object_type"] = item.GetObjectType()
						hyperflexrpostatus["rpo_exceeded"] = item.GetRpoExceeded()

						hyperflexrpostatuss = append(hyperflexrpostatuss, hyperflexrpostatus)
						return hyperflexrpostatuss
					})(item.GetRpoStatus(), d)
					hyperflexreplicationstatus["start_time"] = item.GetStartTime()
					hyperflexreplicationstatus["status"] = item.GetStatus()

					hyperflexreplicationstatuss = append(hyperflexreplicationstatuss, hyperflexreplicationstatus)
					return hyperflexreplicationstatuss
				})(item.GetReplicationStatus(), d)
				hyperflexsnapshotinfobrief["site"] = item.GetSite()
				hyperflexsnapshotinfobrief["snapshot_status"] = (func(p models.HyperflexSnapshotStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexsnapshotstatuss []map[string]interface{}
					var ret models.HyperflexSnapshotStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexsnapshotstatus := make(map[string]interface{})
					hyperflexsnapshotstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexsnapshotstatus["class_id"] = item.GetClassId()
					hyperflexsnapshotstatus["description"] = item.GetDescription()
					hyperflexsnapshotstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.GetClassId()
						hyperflexerrorstack["message"] = item.GetMessage()
						hyperflexerrorstack["message_id"] = item.GetMessageId()
						hyperflexerrorstack["object_type"] = item.GetObjectType()

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexsnapshotstatus["object_type"] = item.GetObjectType()
					hyperflexsnapshotstatus["pct_complete"] = item.GetPctComplete()
					hyperflexsnapshotstatus["status"] = item.GetStatus()
					hyperflexsnapshotstatus["timestamp"] = item.GetTimestamp()
					hyperflexsnapshotstatus["used_space"] = item.GetUsedSpace()

					hyperflexsnapshotstatuss = append(hyperflexsnapshotstatuss, hyperflexsnapshotstatus)
					return hyperflexsnapshotstatuss
				})(item.GetSnapshotStatus(), d)
				hyperflexsnapshotinfobrief["vm_snapshot_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetVmSnapshotEntityReference(), d)

				hyperflexsnapshotinfobriefs = append(hyperflexsnapshotinfobriefs, hyperflexsnapshotinfobrief)
				return hyperflexsnapshotinfobriefs
			})(item.GetVmCurrentProtectionInfo(), d)
			hyperflexprotectioninfo["vm_last_successful_protection_info"] = (func(p models.HyperflexSnapshotInfoBrief, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexsnapshotinfobriefs []map[string]interface{}
				var ret models.HyperflexSnapshotInfoBrief
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexsnapshotinfobrief := make(map[string]interface{})
				hyperflexsnapshotinfobrief["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexsnapshotinfobrief["class_id"] = item.GetClassId()
				hyperflexsnapshotinfobrief["object_type"] = item.GetObjectType()
				hyperflexsnapshotinfobrief["replication_status"] = (func(p models.HyperflexReplicationStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexreplicationstatuss []map[string]interface{}
					var ret models.HyperflexReplicationStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexreplicationstatus := make(map[string]interface{})
					hyperflexreplicationstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexreplicationstatus["bytes_replicated"] = item.GetBytesReplicated()
					hyperflexreplicationstatus["class_id"] = item.GetClassId()
					hyperflexreplicationstatus["end_time"] = item.GetEndTime()
					hyperflexreplicationstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.GetClassId()
						hyperflexerrorstack["message"] = item.GetMessage()
						hyperflexerrorstack["message_id"] = item.GetMessageId()
						hyperflexerrorstack["object_type"] = item.GetObjectType()

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexreplicationstatus["object_type"] = item.GetObjectType()
					hyperflexreplicationstatus["pack_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexentityreferences []map[string]interface{}
						var ret models.HyperflexEntityReference
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexentityreference := make(map[string]interface{})
						hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexentityreference["class_id"] = item.GetClassId()
						hyperflexentityreference["confignum"] = item.GetConfignum()
						hyperflexentityreference["id"] = item.GetId()
						hyperflexentityreference["idtype"] = item.GetIdtype()
						hyperflexentityreference["name"] = item.GetName()
						hyperflexentityreference["object_type"] = item.GetObjectType()
						hyperflexentityreference["type"] = item.GetType()

						hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
						return hyperflexentityreferences
					})(item.GetPackEntityReference(), d)
					hyperflexreplicationstatus["pct_complete"] = item.GetPctComplete()
					hyperflexreplicationstatus["rpo_status"] = (func(p models.HyperflexRpoStatus, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexrpostatuss []map[string]interface{}
						var ret models.HyperflexRpoStatus
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexrpostatus := make(map[string]interface{})
						hyperflexrpostatus["actual"] = item.GetActual()
						hyperflexrpostatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexrpostatus["class_id"] = item.GetClassId()
						hyperflexrpostatus["expected"] = item.GetExpected()
						hyperflexrpostatus["object_type"] = item.GetObjectType()
						hyperflexrpostatus["rpo_exceeded"] = item.GetRpoExceeded()

						hyperflexrpostatuss = append(hyperflexrpostatuss, hyperflexrpostatus)
						return hyperflexrpostatuss
					})(item.GetRpoStatus(), d)
					hyperflexreplicationstatus["start_time"] = item.GetStartTime()
					hyperflexreplicationstatus["status"] = item.GetStatus()

					hyperflexreplicationstatuss = append(hyperflexreplicationstatuss, hyperflexreplicationstatus)
					return hyperflexreplicationstatuss
				})(item.GetReplicationStatus(), d)
				hyperflexsnapshotinfobrief["site"] = item.GetSite()
				hyperflexsnapshotinfobrief["snapshot_status"] = (func(p models.HyperflexSnapshotStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexsnapshotstatuss []map[string]interface{}
					var ret models.HyperflexSnapshotStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexsnapshotstatus := make(map[string]interface{})
					hyperflexsnapshotstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexsnapshotstatus["class_id"] = item.GetClassId()
					hyperflexsnapshotstatus["description"] = item.GetDescription()
					hyperflexsnapshotstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.GetClassId()
						hyperflexerrorstack["message"] = item.GetMessage()
						hyperflexerrorstack["message_id"] = item.GetMessageId()
						hyperflexerrorstack["object_type"] = item.GetObjectType()

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexsnapshotstatus["object_type"] = item.GetObjectType()
					hyperflexsnapshotstatus["pct_complete"] = item.GetPctComplete()
					hyperflexsnapshotstatus["status"] = item.GetStatus()
					hyperflexsnapshotstatus["timestamp"] = item.GetTimestamp()
					hyperflexsnapshotstatus["used_space"] = item.GetUsedSpace()

					hyperflexsnapshotstatuss = append(hyperflexsnapshotstatuss, hyperflexsnapshotstatus)
					return hyperflexsnapshotstatuss
				})(item.GetSnapshotStatus(), d)
				hyperflexsnapshotinfobrief["vm_snapshot_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetVmSnapshotEntityReference(), d)

				hyperflexsnapshotinfobriefs = append(hyperflexsnapshotinfobriefs, hyperflexsnapshotinfobrief)
				return hyperflexsnapshotinfobriefs
			})(item.GetVmLastSuccessfulProtectionInfo(), d)
			hyperflexprotectioninfo["vm_space_usage"] = (func(p models.HyperflexVmProtectionSpaceUsage, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexvmprotectionspaceusages []map[string]interface{}
				var ret models.HyperflexVmProtectionSpaceUsage
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexvmprotectionspaceusage := make(map[string]interface{})
				hyperflexvmprotectionspaceusage["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexvmprotectionspaceusage["class_id"] = item.GetClassId()
				hyperflexvmprotectionspaceusage["object_type"] = item.GetObjectType()
				hyperflexvmprotectionspaceusage["space_usage"] = item.GetSpaceUsage()

				hyperflexvmprotectionspaceusages = append(hyperflexvmprotectionspaceusages, hyperflexvmprotectionspaceusage)
				return hyperflexvmprotectionspaceusages
			})(item.GetVmSpaceUsage(), d)

			hyperflexprotectioninfos = append(hyperflexprotectioninfos, hyperflexprotectioninfo)
			return hyperflexprotectioninfos
		})(item.GetProtectionInfo(), d)
		hyperflexmapclusteridtoprotectioninfos = append(hyperflexmapclusteridtoprotectioninfos, hyperflexmapclusteridtoprotectioninfo)
	}
	return hyperflexmapclusteridtoprotectioninfos
}
func flattenListHyperflexMapClusterIdToStSnapshotPoint(p []models.HyperflexMapClusterIdToStSnapshotPoint, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexmapclusteridtostsnapshotpoints []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexmapclusteridtostsnapshotpoint := make(map[string]interface{})
		hyperflexmapclusteridtostsnapshotpoint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexmapclusteridtostsnapshotpoint["class_id"] = item.GetClassId()
		hyperflexmapclusteridtostsnapshotpoint["cluster_id"] = item.GetClusterId()
		hyperflexmapclusteridtostsnapshotpoint["object_type"] = item.GetObjectType()
		hyperflexmapclusteridtostsnapshotpoint["snapshot_point"] = (func(p models.HyperflexSnapshotPoint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexsnapshotpoints []map[string]interface{}
			var ret models.HyperflexSnapshotPoint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexsnapshotpoint := make(map[string]interface{})
			hyperflexsnapshotpoint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexsnapshotpoint["class_id"] = item.GetClassId()
			hyperflexsnapshotpoint["cluster_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexentityreferences []map[string]interface{}
				var ret models.HyperflexEntityReference
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexentityreference := make(map[string]interface{})
				hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexentityreference["class_id"] = item.GetClassId()
				hyperflexentityreference["confignum"] = item.GetConfignum()
				hyperflexentityreference["id"] = item.GetId()
				hyperflexentityreference["idtype"] = item.GetIdtype()
				hyperflexentityreference["name"] = item.GetName()
				hyperflexentityreference["object_type"] = item.GetObjectType()
				hyperflexentityreference["type"] = item.GetType()

				hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
				return hyperflexentityreferences
			})(item.GetClusterEntityReference(), d)
			hyperflexsnapshotpoint["object_type"] = item.GetObjectType()
			hyperflexsnapshotpoint["replication_status"] = (func(p models.HyperflexReplicationStatus, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationstatuss []map[string]interface{}
				var ret models.HyperflexReplicationStatus
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationstatus := make(map[string]interface{})
				hyperflexreplicationstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationstatus["bytes_replicated"] = item.GetBytesReplicated()
				hyperflexreplicationstatus["class_id"] = item.GetClassId()
				hyperflexreplicationstatus["end_time"] = item.GetEndTime()
				hyperflexreplicationstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexerrorstacks []map[string]interface{}
					var ret models.HyperflexErrorStack
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexerrorstack := make(map[string]interface{})
					hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexerrorstack["class_id"] = item.GetClassId()
					hyperflexerrorstack["message"] = item.GetMessage()
					hyperflexerrorstack["message_id"] = item.GetMessageId()
					hyperflexerrorstack["object_type"] = item.GetObjectType()

					hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
					return hyperflexerrorstacks
				})(item.GetError(), d)
				hyperflexreplicationstatus["object_type"] = item.GetObjectType()
				hyperflexreplicationstatus["pack_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetPackEntityReference(), d)
				hyperflexreplicationstatus["pct_complete"] = item.GetPctComplete()
				hyperflexreplicationstatus["rpo_status"] = (func(p models.HyperflexRpoStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexrpostatuss []map[string]interface{}
					var ret models.HyperflexRpoStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexrpostatus := make(map[string]interface{})
					hyperflexrpostatus["actual"] = item.GetActual()
					hyperflexrpostatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexrpostatus["class_id"] = item.GetClassId()
					hyperflexrpostatus["expected"] = item.GetExpected()
					hyperflexrpostatus["object_type"] = item.GetObjectType()
					hyperflexrpostatus["rpo_exceeded"] = item.GetRpoExceeded()

					hyperflexrpostatuss = append(hyperflexrpostatuss, hyperflexrpostatus)
					return hyperflexrpostatuss
				})(item.GetRpoStatus(), d)
				hyperflexreplicationstatus["start_time"] = item.GetStartTime()
				hyperflexreplicationstatus["status"] = item.GetStatus()

				hyperflexreplicationstatuss = append(hyperflexreplicationstatuss, hyperflexreplicationstatus)
				return hyperflexreplicationstatuss
			})(item.GetReplicationStatus(), d)
			hyperflexsnapshotpoint["snapshot_files"] = (func(p models.HyperflexSnapshotFiles, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexsnapshotfiless []map[string]interface{}
				var ret models.HyperflexSnapshotFiles
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexsnapshotfiles := make(map[string]interface{})
				hyperflexsnapshotfiles["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexsnapshotfiles["class_id"] = item.GetClassId()
				hyperflexsnapshotfiles["name_tracked_files"] = (func(p []models.HyperflexFilePath, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexfilepaths []map[string]interface{}
					if len(p) == 0 {
						return nil
					}
					for _, item := range p {
						hyperflexfilepath := make(map[string]interface{})
						hyperflexfilepath["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexfilepath["class_id"] = item.GetClassId()
						hyperflexfilepath["ds_info"] = (func(p models.HyperflexDatastoreInfo, d *schema.ResourceData) []map[string]interface{} {
							var hyperflexdatastoreinfos []map[string]interface{}
							var ret models.HyperflexDatastoreInfo
							if reflect.DeepEqual(ret, p) {
								return nil
							}
							item := p
							hyperflexdatastoreinfo := make(map[string]interface{})
							hyperflexdatastoreinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							hyperflexdatastoreinfo["class_id"] = item.GetClassId()
							hyperflexdatastoreinfo["ds_backend_id"] = item.GetDsBackendId()
							hyperflexdatastoreinfo["ds_frontend_id"] = item.GetDsFrontendId()
							hyperflexdatastoreinfo["object_type"] = item.GetObjectType()

							hyperflexdatastoreinfos = append(hyperflexdatastoreinfos, hyperflexdatastoreinfo)
							return hyperflexdatastoreinfos
						})(item.GetDsInfo(), d)
						hyperflexfilepath["object_type"] = item.GetObjectType()
						hyperflexfilepath["relative_file_path"] = item.GetRelativeFilePath()
						hyperflexfilepaths = append(hyperflexfilepaths, hyperflexfilepath)
					}
					return hyperflexfilepaths
				})(item.GetNameTrackedFiles(), d)
				hyperflexsnapshotfiles["object_type"] = item.GetObjectType()
				hyperflexsnapshotfiles["uuid_tracked_disks_map"] = (func(p []models.HyperflexMapUuidToTrackedDisk, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexmapuuidtotrackeddisks []map[string]interface{}
					if len(p) == 0 {
						return nil
					}
					for _, item := range p {
						hyperflexmapuuidtotrackeddisk := make(map[string]interface{})
						hyperflexmapuuidtotrackeddisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexmapuuidtotrackeddisk["class_id"] = item.GetClassId()
						hyperflexmapuuidtotrackeddisk["object_type"] = item.GetObjectType()
						hyperflexmapuuidtotrackeddisk["tracked_disk"] = (func(p models.HyperflexTrackedDisk, d *schema.ResourceData) []map[string]interface{} {
							var hyperflextrackeddisks []map[string]interface{}
							var ret models.HyperflexTrackedDisk
							if reflect.DeepEqual(ret, p) {
								return nil
							}
							item := p
							hyperflextrackeddisk := make(map[string]interface{})
							hyperflextrackeddisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							hyperflextrackeddisk["class_id"] = item.GetClassId()
							hyperflextrackeddisk["disk_files"] = (func(p []models.HyperflexTrackedFile, d *schema.ResourceData) []map[string]interface{} {
								var hyperflextrackedfiles []map[string]interface{}
								if len(p) == 0 {
									return nil
								}
								for _, item := range p {
									hyperflextrackedfile := make(map[string]interface{})
									hyperflextrackedfile["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
									hyperflextrackedfile["class_id"] = item.GetClassId()
									hyperflextrackedfile["file_path"] = (func(p models.HyperflexFilePath, d *schema.ResourceData) []map[string]interface{} {
										var hyperflexfilepaths []map[string]interface{}
										var ret models.HyperflexFilePath
										if reflect.DeepEqual(ret, p) {
											return nil
										}
										item := p
										hyperflexfilepath := make(map[string]interface{})
										hyperflexfilepath["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
										hyperflexfilepath["class_id"] = item.GetClassId()
										hyperflexfilepath["ds_info"] = (func(p models.HyperflexDatastoreInfo, d *schema.ResourceData) []map[string]interface{} {
											var hyperflexdatastoreinfos []map[string]interface{}
											var ret models.HyperflexDatastoreInfo
											if reflect.DeepEqual(ret, p) {
												return nil
											}
											item := p
											hyperflexdatastoreinfo := make(map[string]interface{})
											hyperflexdatastoreinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
											hyperflexdatastoreinfo["class_id"] = item.GetClassId()
											hyperflexdatastoreinfo["ds_backend_id"] = item.GetDsBackendId()
											hyperflexdatastoreinfo["ds_frontend_id"] = item.GetDsFrontendId()
											hyperflexdatastoreinfo["object_type"] = item.GetObjectType()

											hyperflexdatastoreinfos = append(hyperflexdatastoreinfos, hyperflexdatastoreinfo)
											return hyperflexdatastoreinfos
										})(item.GetDsInfo(), d)
										hyperflexfilepath["object_type"] = item.GetObjectType()
										hyperflexfilepath["relative_file_path"] = item.GetRelativeFilePath()

										hyperflexfilepaths = append(hyperflexfilepaths, hyperflexfilepath)
										return hyperflexfilepaths
									})(item.GetFilePath(), d)
									hyperflextrackedfile["file_type"] = item.GetFileType()
									hyperflextrackedfile["object_type"] = item.GetObjectType()
									hyperflextrackedfiles = append(hyperflextrackedfiles, hyperflextrackedfile)
								}
								return hyperflextrackedfiles
							})(item.GetDiskFiles(), d)
							hyperflextrackeddisk["disk_type"] = item.GetDiskType()
							hyperflextrackeddisk["object_type"] = item.GetObjectType()

							hyperflextrackeddisks = append(hyperflextrackeddisks, hyperflextrackeddisk)
							return hyperflextrackeddisks
						})(item.GetTrackedDisk(), d)
						hyperflexmapuuidtotrackeddisk["uuid"] = item.GetUuid()
						hyperflexmapuuidtotrackeddisks = append(hyperflexmapuuidtotrackeddisks, hyperflexmapuuidtotrackeddisk)
					}
					return hyperflexmapuuidtotrackeddisks
				})(item.GetUuidTrackedDisksMap(), d)

				hyperflexsnapshotfiless = append(hyperflexsnapshotfiless, hyperflexsnapshotfiles)
				return hyperflexsnapshotfiless
			})(item.GetSnapshotFiles(), d)
			hyperflexsnapshotpoint["snapshot_point_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexentityreferences []map[string]interface{}
				var ret models.HyperflexEntityReference
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexentityreference := make(map[string]interface{})
				hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexentityreference["class_id"] = item.GetClassId()
				hyperflexentityreference["confignum"] = item.GetConfignum()
				hyperflexentityreference["id"] = item.GetId()
				hyperflexentityreference["idtype"] = item.GetIdtype()
				hyperflexentityreference["name"] = item.GetName()
				hyperflexentityreference["object_type"] = item.GetObjectType()
				hyperflexentityreference["type"] = item.GetType()

				hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
				return hyperflexentityreferences
			})(item.GetSnapshotPointEntityReference(), d)
			hyperflexsnapshotpoint["snapshot_status"] = (func(p models.HyperflexSnapshotStatus, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexsnapshotstatuss []map[string]interface{}
				var ret models.HyperflexSnapshotStatus
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexsnapshotstatus := make(map[string]interface{})
				hyperflexsnapshotstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexsnapshotstatus["class_id"] = item.GetClassId()
				hyperflexsnapshotstatus["description"] = item.GetDescription()
				hyperflexsnapshotstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexerrorstacks []map[string]interface{}
					var ret models.HyperflexErrorStack
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexerrorstack := make(map[string]interface{})
					hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexerrorstack["class_id"] = item.GetClassId()
					hyperflexerrorstack["message"] = item.GetMessage()
					hyperflexerrorstack["message_id"] = item.GetMessageId()
					hyperflexerrorstack["object_type"] = item.GetObjectType()

					hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
					return hyperflexerrorstacks
				})(item.GetError(), d)
				hyperflexsnapshotstatus["object_type"] = item.GetObjectType()
				hyperflexsnapshotstatus["pct_complete"] = item.GetPctComplete()
				hyperflexsnapshotstatus["status"] = item.GetStatus()
				hyperflexsnapshotstatus["timestamp"] = item.GetTimestamp()
				hyperflexsnapshotstatus["used_space"] = item.GetUsedSpace()

				hyperflexsnapshotstatuss = append(hyperflexsnapshotstatuss, hyperflexsnapshotstatus)
				return hyperflexsnapshotstatuss
			})(item.GetSnapshotStatus(), d)
			hyperflexsnapshotpoint["vm_runtime_info"] = (func(p models.HyperflexVirtualMachine, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexvirtualmachines []map[string]interface{}
				var ret models.HyperflexVirtualMachine
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexvirtualmachine := make(map[string]interface{})
				hyperflexvirtualmachine["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexvirtualmachine["class_id"] = item.GetClassId()
				hyperflexvirtualmachine["object_type"] = item.GetObjectType()
				hyperflexvirtualmachine["run_time_info"] = (func(p models.HyperflexVirtualMachineRuntimeInfo, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexvirtualmachineruntimeinfos []map[string]interface{}
					var ret models.HyperflexVirtualMachineRuntimeInfo
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexvirtualmachineruntimeinfo := make(map[string]interface{})
					hyperflexvirtualmachineruntimeinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexvirtualmachineruntimeinfo["bios_uuid"] = item.GetBiosUuid()
					hyperflexvirtualmachineruntimeinfo["class_id"] = item.GetClassId()
					hyperflexvirtualmachineruntimeinfo["connection_state"] = item.GetConnectionState()
					hyperflexvirtualmachineruntimeinfo["cpu_usage"] = item.GetCpuUsage()
					hyperflexvirtualmachineruntimeinfo["folder"] = item.GetFolder()
					hyperflexvirtualmachineruntimeinfo["guest_family"] = item.GetGuestFamily()
					hyperflexvirtualmachineruntimeinfo["guest_full_name"] = item.GetGuestFullName()
					hyperflexvirtualmachineruntimeinfo["guest_id"] = item.GetGuestId()
					hyperflexvirtualmachineruntimeinfo["guest_state"] = item.GetGuestState()
					hyperflexvirtualmachineruntimeinfo["host_name"] = item.GetHostName()
					hyperflexvirtualmachineruntimeinfo["instance_uuid"] = item.GetInstanceUuid()
					hyperflexvirtualmachineruntimeinfo["memory_mb"] = item.GetMemoryMb()
					hyperflexvirtualmachineruntimeinfo["memory_usage"] = item.GetMemoryUsage()
					hyperflexvirtualmachineruntimeinfo["moid"] = item.GetMoid()
					hyperflexvirtualmachineruntimeinfo["name"] = item.GetName()
					hyperflexvirtualmachineruntimeinfo["networks"] = item.GetNetworks()
					hyperflexvirtualmachineruntimeinfo["num_cpu"] = item.GetNumCpu()
					hyperflexvirtualmachineruntimeinfo["object_type"] = item.GetObjectType()
					hyperflexvirtualmachineruntimeinfo["power_state"] = item.GetPowerState()
					hyperflexvirtualmachineruntimeinfo["provisioned_size"] = item.GetProvisionedSize()
					hyperflexvirtualmachineruntimeinfo["resource_pool"] = item.GetResourcePool()
					hyperflexvirtualmachineruntimeinfo["used_size"] = item.GetUsedSize()
					hyperflexvirtualmachineruntimeinfo["nr_version"] = item.GetVersion()
					hyperflexvirtualmachineruntimeinfo["vmx_path"] = item.GetVmxPath()

					hyperflexvirtualmachineruntimeinfos = append(hyperflexvirtualmachineruntimeinfos, hyperflexvirtualmachineruntimeinfo)
					return hyperflexvirtualmachineruntimeinfos
				})(item.GetRunTimeInfo(), d)
				hyperflexvirtualmachine["status_code"] = item.GetStatusCode()
				hyperflexvirtualmachine["uuid"] = item.GetUuid()

				hyperflexvirtualmachines = append(hyperflexvirtualmachines, hyperflexvirtualmachine)
				return hyperflexvirtualmachines
			})(item.GetVmRuntimeInfo(), d)

			hyperflexsnapshotpoints = append(hyperflexsnapshotpoints, hyperflexsnapshotpoint)
			return hyperflexsnapshotpoints
		})(item.GetSnapshotPoint(), d)
		hyperflexmapclusteridtostsnapshotpoints = append(hyperflexmapclusteridtostsnapshotpoints, hyperflexmapclusteridtostsnapshotpoint)
	}
	return hyperflexmapclusteridtostsnapshotpoints
}
func flattenListHyperflexNamedVlan(p []models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexnamedvlan := make(map[string]interface{})
		hyperflexnamedvlan["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexnamedvlan["class_id"] = item.GetClassId()
		hyperflexnamedvlan["name"] = item.GetName()
		hyperflexnamedvlan["object_type"] = item.GetObjectType()
		hyperflexnamedvlan["vlan_id"] = item.GetVlanId()
		hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	}
	return hyperflexnamedvlans
}
func flattenListHyperflexNetworkPort(p []models.HyperflexNetworkPort, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnetworkports []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexnetworkport := make(map[string]interface{})
		hyperflexnetworkport["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexnetworkport["bond_state"] = (func(p models.HyperflexBondState, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexbondstates []map[string]interface{}
			var ret models.HyperflexBondState
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexbondstate := make(map[string]interface{})
			hyperflexbondstate["active_slave"] = item.GetActiveSlave()
			hyperflexbondstate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexbondstate["class_id"] = item.GetClassId()
			hyperflexbondstate["mode"] = item.GetMode()
			hyperflexbondstate["object_type"] = item.GetObjectType()
			hyperflexbondstate["slaves"] = item.GetSlaves()

			hyperflexbondstates = append(hyperflexbondstates, hyperflexbondstate)
			return hyperflexbondstates
		})(item.GetBondState(), d)
		hyperflexnetworkport["class_id"] = item.GetClassId()
		hyperflexnetworkport["name"] = item.GetName()
		hyperflexnetworkport["net_interfaces"] = item.GetNetInterfaces()
		hyperflexnetworkport["object_type"] = item.GetObjectType()
		hyperflexnetworkport["port_type"] = item.GetPortType()
		hyperflexnetworkport["vlans"] = item.GetVlans()
		hyperflexnetworkports = append(hyperflexnetworkports, hyperflexnetworkport)
	}
	return hyperflexnetworkports
}
func flattenListHyperflexNodeRelationship(p []models.HyperflexNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnoderelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexnoderelationship := flattenMoMoRef(item)
		hyperflexnoderelationships = append(hyperflexnoderelationships, hyperflexnoderelationship)
	}
	return hyperflexnoderelationships
}
func flattenListHyperflexNodeProfileRelationship(p []models.HyperflexNodeProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexnodeprofilerelationship := flattenMoMoRef(item)
		hyperflexnodeprofilerelationships = append(hyperflexnodeprofilerelationships, hyperflexnodeprofilerelationship)
	}
	return hyperflexnodeprofilerelationships
}
func flattenListHyperflexReplicationClusterReferenceToSchedule(p []models.HyperflexReplicationClusterReferenceToSchedule, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexreplicationclusterreferencetoschedules []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexreplicationclusterreferencetoschedule := make(map[string]interface{})
		hyperflexreplicationclusterreferencetoschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexreplicationclusterreferencetoschedule["class_id"] = item.GetClassId()
		hyperflexreplicationclusterreferencetoschedule["object_type"] = item.GetObjectType()
		hyperflexreplicationclusterreferencetoschedule["schedule"] = (func(p models.HyperflexReplicationSchedule, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexreplicationschedules []map[string]interface{}
			var ret models.HyperflexReplicationSchedule
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexreplicationschedule := make(map[string]interface{})
			hyperflexreplicationschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexreplicationschedule["backup_interval"] = item.GetBackupInterval()
			hyperflexreplicationschedule["class_id"] = item.GetClassId()
			hyperflexreplicationschedule["object_type"] = item.GetObjectType()

			hyperflexreplicationschedules = append(hyperflexreplicationschedules, hyperflexreplicationschedule)
			return hyperflexreplicationschedules
		})(item.GetSchedule(), d)
		hyperflexreplicationclusterreferencetoschedule["target_cluster_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexentityreferences []map[string]interface{}
			var ret models.HyperflexEntityReference
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexentityreference := make(map[string]interface{})
			hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexentityreference["class_id"] = item.GetClassId()
			hyperflexentityreference["confignum"] = item.GetConfignum()
			hyperflexentityreference["id"] = item.GetId()
			hyperflexentityreference["idtype"] = item.GetIdtype()
			hyperflexentityreference["name"] = item.GetName()
			hyperflexentityreference["object_type"] = item.GetObjectType()
			hyperflexentityreference["type"] = item.GetType()

			hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
			return hyperflexentityreferences
		})(item.GetTargetClusterEntityReference(), d)
		hyperflexreplicationclusterreferencetoschedules = append(hyperflexreplicationclusterreferencetoschedules, hyperflexreplicationclusterreferencetoschedule)
	}
	return hyperflexreplicationclusterreferencetoschedules
}
func flattenListHyperflexServerFirmwareVersionEntryRelationship(p []models.HyperflexServerFirmwareVersionEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexserverfirmwareversionentryrelationship := flattenMoMoRef(item)
		hyperflexserverfirmwareversionentryrelationships = append(hyperflexserverfirmwareversionentryrelationships, hyperflexserverfirmwareversionentryrelationship)
	}
	return hyperflexserverfirmwareversionentryrelationships
}
func flattenListHyperflexServerFirmwareVersionInfo(p []models.HyperflexServerFirmwareVersionInfo, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversioninfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexserverfirmwareversioninfo := make(map[string]interface{})
		hyperflexserverfirmwareversioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexserverfirmwareversioninfo["class_id"] = item.GetClassId()
		hyperflexserverfirmwareversioninfo["object_type"] = item.GetObjectType()
		hyperflexserverfirmwareversioninfo["server_platform"] = item.GetServerPlatform()
		hyperflexserverfirmwareversioninfo["nr_version"] = item.GetVersion()
		hyperflexserverfirmwareversioninfos = append(hyperflexserverfirmwareversioninfos, hyperflexserverfirmwareversioninfo)
	}
	return hyperflexserverfirmwareversioninfos
}
func flattenListHyperflexServerModelEntry(p []models.HyperflexServerModelEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelentrys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexservermodelentry := make(map[string]interface{})
		hyperflexservermodelentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexservermodelentry["class_id"] = item.GetClassId()
		hyperflexservermodelentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.GetClassId()
			hyperflexappsettingconstraint["hxdp_version"] = item.GetHxdpVersion()
			hyperflexappsettingconstraint["hypervisor_type"] = item.GetHypervisorType()
			hyperflexappsettingconstraint["mgmt_platform"] = item.GetMgmtPlatform()
			hyperflexappsettingconstraint["object_type"] = item.GetObjectType()
			hyperflexappsettingconstraint["server_model"] = item.GetServerModel()

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexservermodelentry["name"] = item.GetName()
		hyperflexservermodelentry["object_type"] = item.GetObjectType()
		hyperflexservermodelentry["value"] = item.GetValue()
		hyperflexservermodelentrys = append(hyperflexservermodelentrys, hyperflexservermodelentry)
	}
	return hyperflexservermodelentrys
}
func flattenListHyperflexSoftwareDistributionComponentRelationship(p []models.HyperflexSoftwareDistributionComponentRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwaredistributioncomponentrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexsoftwaredistributioncomponentrelationship := flattenMoMoRef(item)
		hyperflexsoftwaredistributioncomponentrelationships = append(hyperflexsoftwaredistributioncomponentrelationships, hyperflexsoftwaredistributioncomponentrelationship)
	}
	return hyperflexsoftwaredistributioncomponentrelationships
}
func flattenListHyperflexSoftwareDistributionEntryRelationship(p []models.HyperflexSoftwareDistributionEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwaredistributionentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexsoftwaredistributionentryrelationship := flattenMoMoRef(item)
		hyperflexsoftwaredistributionentryrelationships = append(hyperflexsoftwaredistributionentryrelationships, hyperflexsoftwaredistributionentryrelationship)
	}
	return hyperflexsoftwaredistributionentryrelationships
}
func flattenListHyperflexSoftwareDistributionVersionRelationship(p []models.HyperflexSoftwareDistributionVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwaredistributionversionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexsoftwaredistributionversionrelationship := flattenMoMoRef(item)
		hyperflexsoftwaredistributionversionrelationships = append(hyperflexsoftwaredistributionversionrelationships, hyperflexsoftwaredistributionversionrelationship)
	}
	return hyperflexsoftwaredistributionversionrelationships
}
func flattenListHyperflexVmDisk(p []models.HyperflexVmDisk, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvmdisks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexvmdisk := make(map[string]interface{})
		hyperflexvmdisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexvmdisk["boot_order"] = item.GetBootOrder()
		hyperflexvmdisk["bus"] = item.GetBus()
		hyperflexvmdisk["class_id"] = item.GetClassId()
		hyperflexvmdisk["name"] = item.GetName()
		hyperflexvmdisk["object_type"] = item.GetObjectType()
		hyperflexvmdisk["type"] = item.GetType()
		hyperflexvmdisk["virtual_disk"] = (func(p models.HyperflexVdiskConfig, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexvdiskconfigs []map[string]interface{}
			var ret models.HyperflexVdiskConfig
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexvdiskconfig := make(map[string]interface{})
			hyperflexvdiskconfig["access_mode"] = item.GetAccessMode()
			hyperflexvdiskconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexvdiskconfig["capacity"] = item.GetCapacity()
			hyperflexvdiskconfig["class_id"] = item.GetClassId()
			hyperflexvdiskconfig["mode"] = item.GetMode()
			hyperflexvdiskconfig["name"] = item.GetName()
			hyperflexvdiskconfig["object_type"] = item.GetObjectType()
			hyperflexvdiskconfig["source_file_path"] = item.GetSourceFilePath()
			hyperflexvdiskconfig["source_virtual_disk"] = item.GetSourceVirtualDisk()
			hyperflexvdiskconfig["status"] = (func(p models.HyperflexDiskStatus, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexdiskstatuss []map[string]interface{}
				var ret models.HyperflexDiskStatus
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexdiskstatus := make(map[string]interface{})
				hyperflexdiskstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexdiskstatus["class_id"] = item.GetClassId()
				hyperflexdiskstatus["download_percentage"] = item.GetDownloadPercentage()
				hyperflexdiskstatus["object_type"] = item.GetObjectType()
				hyperflexdiskstatus["state"] = item.GetState()
				hyperflexdiskstatus["volume_handle"] = item.GetVolumeHandle()
				hyperflexdiskstatus["volume_name"] = item.GetVolumeName()

				hyperflexdiskstatuss = append(hyperflexdiskstatuss, hyperflexdiskstatus)
				return hyperflexdiskstatuss
			})(item.GetStatus(), d)
			hyperflexvdiskconfig["uuid"] = item.GetUuid()

			hyperflexvdiskconfigs = append(hyperflexvdiskconfigs, hyperflexvdiskconfig)
			return hyperflexvdiskconfigs
		})(item.GetVirtualDisk(), d)
		hyperflexvmdisk["virtual_disk_reference"] = item.GetVirtualDiskReference()
		hyperflexvmdisks = append(hyperflexvmdisks, hyperflexvmdisk)
	}
	return hyperflexvmdisks
}
func flattenListHyperflexVmInterface(p []models.HyperflexVmInterface, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvminterfaces []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexvminterface := make(map[string]interface{})
		hyperflexvminterface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexvminterface["bridge"] = item.GetBridge()
		hyperflexvminterface["class_id"] = item.GetClassId()
		hyperflexvminterface["ip_address"] = item.GetIpAddress()
		hyperflexvminterface["mac_address"] = item.GetMacAddress()
		hyperflexvminterface["model"] = item.GetModel()
		hyperflexvminterface["name"] = item.GetName()
		hyperflexvminterface["object_type"] = item.GetObjectType()
		hyperflexvminterfaces = append(hyperflexvminterfaces, hyperflexvminterface)
	}
	return hyperflexvminterfaces
}
func flattenListHyperflexVolumeRelationship(p []models.HyperflexVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvolumerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexvolumerelationship := flattenMoMoRef(item)
		hyperflexvolumerelationships = append(hyperflexvolumerelationships, hyperflexvolumerelationship)
	}
	return hyperflexvolumerelationships
}
func flattenListIaasConnectorPackRelationship(p []models.IaasConnectorPackRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasconnectorpackrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iaasconnectorpackrelationship := flattenMoMoRef(item)
		iaasconnectorpackrelationships = append(iaasconnectorpackrelationships, iaasconnectorpackrelationship)
	}
	return iaasconnectorpackrelationships
}
func flattenListIaasDeviceStatusRelationship(p []models.IaasDeviceStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasdevicestatusrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iaasdevicestatusrelationship := flattenMoMoRef(item)
		iaasdevicestatusrelationships = append(iaasdevicestatusrelationships, iaasdevicestatusrelationship)
	}
	return iaasdevicestatusrelationships
}
func flattenListIaasLicenseKeysInfo(p []models.IaasLicenseKeysInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicensekeysinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iaaslicensekeysinfo := make(map[string]interface{})
		iaaslicensekeysinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iaaslicensekeysinfo["class_id"] = item.GetClassId()
		iaaslicensekeysinfo["nr_count"] = item.GetCount()
		iaaslicensekeysinfo["expiration_date"] = item.GetExpirationDate()
		iaaslicensekeysinfo["license_id"] = item.GetLicenseId()
		iaaslicensekeysinfo["object_type"] = item.GetObjectType()
		iaaslicensekeysinfo["pid"] = item.GetPid()
		iaaslicensekeysinfos = append(iaaslicensekeysinfos, iaaslicensekeysinfo)
	}
	return iaaslicensekeysinfos
}
func flattenListIaasLicenseUtilizationInfo(p []models.IaasLicenseUtilizationInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseutilizationinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iaaslicenseutilizationinfo := make(map[string]interface{})
		iaaslicenseutilizationinfo["actual_used"] = item.GetActualUsed()
		iaaslicenseutilizationinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iaaslicenseutilizationinfo["class_id"] = item.GetClassId()
		iaaslicenseutilizationinfo["label"] = item.GetLabel()
		iaaslicenseutilizationinfo["licensed_limit"] = item.GetLicensedLimit()
		iaaslicenseutilizationinfo["object_type"] = item.GetObjectType()
		iaaslicenseutilizationinfo["sku"] = item.GetSku()
		iaaslicenseutilizationinfos = append(iaaslicenseutilizationinfos, iaaslicenseutilizationinfo)
	}
	return iaaslicenseutilizationinfos
}
func flattenListIaasMostRunTasksRelationship(p []models.IaasMostRunTasksRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasmostruntasksrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iaasmostruntasksrelationship := flattenMoMoRef(item)
		iaasmostruntasksrelationships = append(iaasmostruntasksrelationships, iaasmostruntasksrelationship)
	}
	return iaasmostruntasksrelationships
}
func flattenListIaasWorkflowSteps(p []models.IaasWorkflowSteps, d *schema.ResourceData) []map[string]interface{} {
	var iaasworkflowstepss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iaasworkflowsteps := make(map[string]interface{})
		iaasworkflowsteps["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iaasworkflowsteps["class_id"] = item.GetClassId()
		iaasworkflowsteps["completed_time"] = item.GetCompletedTime()
		iaasworkflowsteps["name"] = item.GetName()
		iaasworkflowsteps["object_type"] = item.GetObjectType()
		iaasworkflowsteps["status"] = item.GetStatus()
		iaasworkflowsteps["status_message"] = item.GetStatusMessage()
		iaasworkflowstepss = append(iaasworkflowstepss, iaasworkflowsteps)
	}
	return iaasworkflowstepss
}
func flattenListIamAccountPermissions(p []models.IamAccountPermissions, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountpermissionss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iamaccountpermissions := make(map[string]interface{})
		iamaccountpermissions["account_identifier"] = item.GetAccountIdentifier()
		iamaccountpermissions["account_name"] = item.GetAccountName()
		iamaccountpermissions["account_status"] = item.GetAccountStatus()
		iamaccountpermissions["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamaccountpermissions["class_id"] = item.GetClassId()
		iamaccountpermissions["object_type"] = item.GetObjectType()
		iamaccountpermissions["permissions"] = (func(p []models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var iampermissionreferences []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				iampermissionreference := make(map[string]interface{})
				iampermissionreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				iampermissionreference["class_id"] = item.GetClassId()
				iampermissionreference["object_type"] = item.GetObjectType()
				iampermissionreference["permission_identifier"] = item.GetPermissionIdentifier()
				iampermissionreference["permission_name"] = item.GetPermissionName()
				iampermissionreferences = append(iampermissionreferences, iampermissionreference)
			}
			return iampermissionreferences
		})(item.GetPermissions(), d)
		iamaccountpermissionss = append(iamaccountpermissionss, iamaccountpermissions)
	}
	return iamaccountpermissionss
}
func flattenListIamApiKeyRelationship(p []models.IamApiKeyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamapikeyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamapikeyrelationship := flattenMoMoRef(item)
		iamapikeyrelationships = append(iamapikeyrelationships, iamapikeyrelationship)
	}
	return iamapikeyrelationships
}
func flattenListIamAppRegistrationRelationship(p []models.IamAppRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamappregistrationrelationship := flattenMoMoRef(item)
		iamappregistrationrelationships = append(iamappregistrationrelationships, iamappregistrationrelationship)
	}
	return iamappregistrationrelationships
}
func flattenListIamDomainGroupRelationship(p []models.IamDomainGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamdomaingrouprelationship := flattenMoMoRef(item)
		iamdomaingrouprelationships = append(iamdomaingrouprelationships, iamdomaingrouprelationship)
	}
	return iamdomaingrouprelationships
}
func flattenListIamEndPointPrivilegeRelationship(p []models.IamEndPointPrivilegeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointprivilegerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamendpointprivilegerelationship := flattenMoMoRef(item)
		iamendpointprivilegerelationships = append(iamendpointprivilegerelationships, iamendpointprivilegerelationship)
	}
	return iamendpointprivilegerelationships
}
func flattenListIamEndPointRoleRelationship(p []models.IamEndPointRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointrolerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamendpointrolerelationship := flattenMoMoRef(item)
		iamendpointrolerelationships = append(iamendpointrolerelationships, iamendpointrolerelationship)
	}
	return iamendpointrolerelationships
}
func flattenListIamEndPointUserRoleRelationship(p []models.IamEndPointUserRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrolerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamendpointuserrolerelationship := flattenMoMoRef(item)
		iamendpointuserrolerelationships = append(iamendpointuserrolerelationships, iamendpointuserrolerelationship)
	}
	return iamendpointuserrolerelationships
}
func flattenListIamFeatureDefinition(p []models.IamFeatureDefinition, d *schema.ResourceData) []map[string]interface{} {
	var iamfeaturedefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iamfeaturedefinition := make(map[string]interface{})
		iamfeaturedefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamfeaturedefinition["class_id"] = item.GetClassId()
		iamfeaturedefinition["feature"] = item.GetFeature()
		iamfeaturedefinition["object_type"] = item.GetObjectType()
		iamfeaturedefinitions = append(iamfeaturedefinitions, iamfeaturedefinition)
	}
	return iamfeaturedefinitions
}
func flattenListIamGroupPermissionToRoles(p []models.IamGroupPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iamgrouppermissiontoroless []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iamgrouppermissiontoroles := make(map[string]interface{})
		iamgrouppermissiontoroles["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamgrouppermissiontoroles["class_id"] = item.GetClassId()
		iamgrouppermissiontoroles["group"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			var ret models.MoMoRef
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.GetClassId()
			momoref["moid"] = item.GetMoid()
			momoref["object_type"] = item.GetObjectType()
			momoref["selector"] = item.GetSelector()

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetGroup(), d)
		iamgrouppermissiontoroles["object_type"] = item.GetObjectType()
		iamgrouppermissiontoroles["orgs"] = (func(p []models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				momoref := make(map[string]interface{})
				momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				momoref["class_id"] = item.GetClassId()
				momoref["moid"] = item.GetMoid()
				momoref["object_type"] = item.GetObjectType()
				momoref["selector"] = item.GetSelector()
				momorefs = append(momorefs, momoref)
			}
			return momorefs
		})(item.GetOrgs(), d)
		iamgrouppermissiontoroless = append(iamgrouppermissiontoroless, iamgrouppermissiontoroles)
	}
	return iamgrouppermissiontoroless
}
func flattenListIamIdpRelationship(p []models.IamIdpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamidprelationship := flattenMoMoRef(item)
		iamidprelationships = append(iamidprelationships, iamidprelationship)
	}
	return iamidprelationships
}
func flattenListIamIdpReferenceRelationship(p []models.IamIdpReferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamidpreferencerelationship := flattenMoMoRef(item)
		iamidpreferencerelationships = append(iamidpreferencerelationships, iamidpreferencerelationship)
	}
	return iamidpreferencerelationships
}
func flattenListIamIpAddressRelationship(p []models.IamIpAddressRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamipaddressrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamipaddressrelationship := flattenMoMoRef(item)
		iamipaddressrelationships = append(iamipaddressrelationships, iamipaddressrelationship)
	}
	return iamipaddressrelationships
}
func flattenListIamLdapGroupRelationship(p []models.IamLdapGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldapgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamldapgrouprelationship := flattenMoMoRef(item)
		iamldapgrouprelationships = append(iamldapgrouprelationships, iamldapgrouprelationship)
	}
	return iamldapgrouprelationships
}
func flattenListIamLdapProviderRelationship(p []models.IamLdapProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldapproviderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamldapproviderrelationship := flattenMoMoRef(item)
		iamldapproviderrelationships = append(iamldapproviderrelationships, iamldapproviderrelationship)
	}
	return iamldapproviderrelationships
}
func flattenListIamOAuthTokenRelationship(p []models.IamOAuthTokenRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamoauthtokenrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamoauthtokenrelationship := flattenMoMoRef(item)
		iamoauthtokenrelationships = append(iamoauthtokenrelationships, iamoauthtokenrelationship)
	}
	return iamoauthtokenrelationships
}
func flattenListIamPermissionRelationship(p []models.IamPermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iampermissionrelationship := flattenMoMoRef(item)
		iampermissionrelationships = append(iampermissionrelationships, iampermissionrelationship)
	}
	return iampermissionrelationships
}
func flattenListIamPermissionToRoles(p []models.IamPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iampermissiontoroless []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iampermissiontoroles := make(map[string]interface{})
		iampermissiontoroles["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iampermissiontoroles["class_id"] = item.GetClassId()
		iampermissiontoroles["object_type"] = item.GetObjectType()
		iampermissiontoroles["permission"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			var ret models.MoMoRef
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.GetClassId()
			momoref["moid"] = item.GetMoid()
			momoref["object_type"] = item.GetObjectType()
			momoref["selector"] = item.GetSelector()

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetPermission(), d)
		iampermissiontoroles["roles"] = (func(p []models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				momoref := make(map[string]interface{})
				momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				momoref["class_id"] = item.GetClassId()
				momoref["moid"] = item.GetMoid()
				momoref["object_type"] = item.GetObjectType()
				momoref["selector"] = item.GetSelector()
				momorefs = append(momorefs, momoref)
			}
			return momorefs
		})(item.GetRoles(), d)
		iampermissiontoroless = append(iampermissiontoroless, iampermissiontoroles)
	}
	return iampermissiontoroless
}
func flattenListIamPrivilegeRelationship(p []models.IamPrivilegeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamprivilegerelationship := flattenMoMoRef(item)
		iamprivilegerelationships = append(iamprivilegerelationships, iamprivilegerelationship)
	}
	return iamprivilegerelationships
}
func flattenListIamPrivilegeSetRelationship(p []models.IamPrivilegeSetRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegesetrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamprivilegesetrelationship := flattenMoMoRef(item)
		iamprivilegesetrelationships = append(iamprivilegesetrelationships, iamprivilegesetrelationship)
	}
	return iamprivilegesetrelationships
}
func flattenListIamResourcePermissionRelationship(p []models.IamResourcePermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcepermissionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamresourcepermissionrelationship := flattenMoMoRef(item)
		iamresourcepermissionrelationships = append(iamresourcepermissionrelationships, iamresourcepermissionrelationship)
	}
	return iamresourcepermissionrelationships
}
func flattenListIamResourceRolesRelationship(p []models.IamResourceRolesRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcerolesrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamresourcerolesrelationship := flattenMoMoRef(item)
		iamresourcerolesrelationships = append(iamresourcerolesrelationships, iamresourcerolesrelationship)
	}
	return iamresourcerolesrelationships
}
func flattenListIamRoleRelationship(p []models.IamRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamrolerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamrolerelationship := flattenMoMoRef(item)
		iamrolerelationships = append(iamrolerelationships, iamrolerelationship)
	}
	return iamrolerelationships
}
func flattenListIamSessionRelationship(p []models.IamSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamsessionrelationship := flattenMoMoRef(item)
		iamsessionrelationships = append(iamsessionrelationships, iamsessionrelationship)
	}
	return iamsessionrelationships
}
func flattenListIamUserRelationship(p []models.IamUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamuserrelationship := flattenMoMoRef(item)
		iamuserrelationships = append(iamuserrelationships, iamuserrelationship)
	}
	return iamuserrelationships
}
func flattenListIamUserGroupRelationship(p []models.IamUserGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamusergrouprelationship := flattenMoMoRef(item)
		iamusergrouprelationships = append(iamusergrouprelationships, iamusergrouprelationship)
	}
	return iamusergrouprelationships
}
func flattenListIamUserPreferenceRelationship(p []models.IamUserPreferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserpreferencerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamuserpreferencerelationship := flattenMoMoRef(item)
		iamuserpreferencerelationships = append(iamuserpreferencerelationships, iamuserpreferencerelationship)
	}
	return iamuserpreferencerelationships
}
func flattenListInfraMetaData(p []models.InfraMetaData, d *schema.ResourceData) []map[string]interface{} {
	var inframetadatas []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		inframetadata := make(map[string]interface{})
		inframetadata["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		inframetadata["class_id"] = item.GetClassId()
		inframetadata["name"] = item.GetName()
		inframetadata["object_type"] = item.GetObjectType()
		inframetadata["value"] = item.GetValue()
		inframetadatas = append(inframetadatas, inframetadata)
	}
	return inframetadatas
}
func flattenListInventoryGenericInventoryRelationship(p []models.InventoryGenericInventoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		inventorygenericinventoryrelationship := flattenMoMoRef(item)
		inventorygenericinventoryrelationships = append(inventorygenericinventoryrelationships, inventorygenericinventoryrelationship)
	}
	return inventorygenericinventoryrelationships
}
func flattenListInventoryGenericInventoryHolderRelationship(p []models.InventoryGenericInventoryHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		inventorygenericinventoryholderrelationship := flattenMoMoRef(item)
		inventorygenericinventoryholderrelationships = append(inventorygenericinventoryholderrelationships, inventorygenericinventoryholderrelationship)
	}
	return inventorygenericinventoryholderrelationships
}
func flattenListIppoolIpLeaseRelationship(p []models.IppoolIpLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipleaserelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ippoolipleaserelationship := flattenMoMoRef(item)
		ippoolipleaserelationships = append(ippoolipleaserelationships, ippoolipleaserelationship)
	}
	return ippoolipleaserelationships
}
func flattenListIppoolIpV4Block(p []models.IppoolIpV4Block, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv4blocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ippoolipv4block := make(map[string]interface{})
		ippoolipv4block["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ippoolipv4block["class_id"] = item.GetClassId()
		ippoolipv4block["from"] = item.GetFrom()
		ippoolipv4block["object_type"] = item.GetObjectType()
		ippoolipv4block["size"] = item.GetSize()
		ippoolipv4block["to"] = item.GetTo()
		ippoolipv4blocks = append(ippoolipv4blocks, ippoolipv4block)
	}
	return ippoolipv4blocks
}
func flattenListIppoolIpV6Block(p []models.IppoolIpV6Block, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv6blocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ippoolipv6block := make(map[string]interface{})
		ippoolipv6block["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ippoolipv6block["class_id"] = item.GetClassId()
		ippoolipv6block["from"] = item.GetFrom()
		ippoolipv6block["object_type"] = item.GetObjectType()
		ippoolipv6block["size"] = item.GetSize()
		ippoolipv6block["to"] = item.GetTo()
		ippoolipv6blocks = append(ippoolipv6blocks, ippoolipv6block)
	}
	return ippoolipv6blocks
}
func flattenListIppoolPoolRelationship(p []models.IppoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolpoolrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ippoolpoolrelationship := flattenMoMoRef(item)
		ippoolpoolrelationships = append(ippoolpoolrelationships, ippoolpoolrelationship)
	}
	return ippoolpoolrelationships
}
func flattenListIppoolShadowBlockRelationship(p []models.IppoolShadowBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ippoolshadowblockrelationship := flattenMoMoRef(item)
		ippoolshadowblockrelationships = append(ippoolshadowblockrelationships, ippoolshadowblockrelationship)
	}
	return ippoolshadowblockrelationships
}
func flattenListIppoolShadowPoolRelationship(p []models.IppoolShadowPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowpoolrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ippoolshadowpoolrelationship := flattenMoMoRef(item)
		ippoolshadowpoolrelationships = append(ippoolshadowpoolrelationships, ippoolshadowpoolrelationship)
	}
	return ippoolshadowpoolrelationships
}
func flattenListIqnpoolBlockRelationship(p []models.IqnpoolBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iqnpoolblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iqnpoolblockrelationship := flattenMoMoRef(item)
		iqnpoolblockrelationships = append(iqnpoolblockrelationships, iqnpoolblockrelationship)
	}
	return iqnpoolblockrelationships
}
func flattenListIqnpoolIqnSuffixBlock(p []models.IqnpoolIqnSuffixBlock, d *schema.ResourceData) []map[string]interface{} {
	var iqnpooliqnsuffixblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iqnpooliqnsuffixblock := make(map[string]interface{})
		iqnpooliqnsuffixblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iqnpooliqnsuffixblock["class_id"] = item.GetClassId()
		iqnpooliqnsuffixblock["from"] = item.GetFrom()
		iqnpooliqnsuffixblock["object_type"] = item.GetObjectType()
		iqnpooliqnsuffixblock["size"] = item.GetSize()
		iqnpooliqnsuffixblock["suffix"] = item.GetSuffix()
		iqnpooliqnsuffixblock["to"] = item.GetTo()
		iqnpooliqnsuffixblocks = append(iqnpooliqnsuffixblocks, iqnpooliqnsuffixblock)
	}
	return iqnpooliqnsuffixblocks
}
func flattenListKubernetesAciCniProfileRelationship(p []models.KubernetesAciCniProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesacicniprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesacicniprofilerelationship := flattenMoMoRef(item)
		kubernetesacicniprofilerelationships = append(kubernetesacicniprofilerelationships, kubernetesacicniprofilerelationship)
	}
	return kubernetesacicniprofilerelationships
}
func flattenListKubernetesAciCniTenantClusterAllocationRelationship(p []models.KubernetesAciCniTenantClusterAllocationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesacicnitenantclusterallocationrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesacicnitenantclusterallocationrelationship := flattenMoMoRef(item)
		kubernetesacicnitenantclusterallocationrelationships = append(kubernetesacicnitenantclusterallocationrelationships, kubernetesacicnitenantclusterallocationrelationship)
	}
	return kubernetesacicnitenantclusterallocationrelationships
}
func flattenListKubernetesAddon(p []models.KubernetesAddon, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesaddons []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		kubernetesaddon := make(map[string]interface{})
		kubernetesaddon["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesaddon["addon_configuration"] = (func(p models.KubernetesAddonConfiguration, d *schema.ResourceData) []map[string]interface{} {
			var kubernetesaddonconfigurations []map[string]interface{}
			var ret models.KubernetesAddonConfiguration
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			kubernetesaddonconfiguration := make(map[string]interface{})
			kubernetesaddonconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			kubernetesaddonconfiguration["class_id"] = item.GetClassId()
			kubernetesaddonconfiguration["install_strategy"] = item.GetInstallStrategy()
			kubernetesaddonconfiguration["object_type"] = item.GetObjectType()
			kubernetesaddonconfiguration["override_sets"] = (func(p []models.KubernetesKeyValue, d *schema.ResourceData) []map[string]interface{} {
				var kuberneteskeyvalues []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					kuberneteskeyvalue := make(map[string]interface{})
					kuberneteskeyvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					kuberneteskeyvalue["class_id"] = item.GetClassId()
					kuberneteskeyvalue["key"] = item.GetKey()
					kuberneteskeyvalue["object_type"] = item.GetObjectType()
					kuberneteskeyvalue["value"] = item.GetValue()
					kuberneteskeyvalues = append(kuberneteskeyvalues, kuberneteskeyvalue)
				}
				return kuberneteskeyvalues
			})(item.GetOverrideSets(), d)
			kubernetesaddonconfiguration["overrides"] = item.GetOverrides()
			kubernetesaddonconfiguration["release_name"] = item.GetReleaseName()
			kubernetesaddonconfiguration["release_namespace"] = item.GetReleaseNamespace()
			kubernetesaddonconfiguration["upgrade_strategy"] = item.GetUpgradeStrategy()

			kubernetesaddonconfigurations = append(kubernetesaddonconfigurations, kubernetesaddonconfiguration)
			return kubernetesaddonconfigurations
		})(item.GetAddonConfiguration(), d)
		kubernetesaddon["addon_policy"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			var ret models.MoMoRef
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.GetClassId()
			momoref["moid"] = item.GetMoid()
			momoref["object_type"] = item.GetObjectType()
			momoref["selector"] = item.GetSelector()

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetAddonPolicy(), d)
		kubernetesaddon["class_id"] = item.GetClassId()
		kubernetesaddon["name"] = item.GetName()
		kubernetesaddon["object_type"] = item.GetObjectType()
		kubernetesaddons = append(kubernetesaddons, kubernetesaddon)
	}
	return kubernetesaddons
}
func flattenListKubernetesClusterProfileRelationship(p []models.KubernetesClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesclusterprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesclusterprofilerelationship := flattenMoMoRef(item)
		kubernetesclusterprofilerelationships = append(kubernetesclusterprofilerelationships, kubernetesclusterprofilerelationship)
	}
	return kubernetesclusterprofilerelationships
}
func flattenListKubernetesConfigResultEntryRelationship(p []models.KubernetesConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesconfigresultentryrelationship := flattenMoMoRef(item)
		kubernetesconfigresultentryrelationships = append(kubernetesconfigresultentryrelationships, kubernetesconfigresultentryrelationship)
	}
	return kubernetesconfigresultentryrelationships
}
func flattenListKubernetesEssentialAddon(p []models.KubernetesEssentialAddon, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesessentialaddons []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		kubernetesessentialaddon := make(map[string]interface{})
		kubernetesessentialaddon["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesessentialaddon["addon_configuration"] = (func(p models.KubernetesAddonConfiguration, d *schema.ResourceData) []map[string]interface{} {
			var kubernetesaddonconfigurations []map[string]interface{}
			var ret models.KubernetesAddonConfiguration
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			kubernetesaddonconfiguration := make(map[string]interface{})
			kubernetesaddonconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			kubernetesaddonconfiguration["class_id"] = item.GetClassId()
			kubernetesaddonconfiguration["install_strategy"] = item.GetInstallStrategy()
			kubernetesaddonconfiguration["object_type"] = item.GetObjectType()
			kubernetesaddonconfiguration["override_sets"] = (func(p []models.KubernetesKeyValue, d *schema.ResourceData) []map[string]interface{} {
				var kuberneteskeyvalues []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					kuberneteskeyvalue := make(map[string]interface{})
					kuberneteskeyvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					kuberneteskeyvalue["class_id"] = item.GetClassId()
					kuberneteskeyvalue["key"] = item.GetKey()
					kuberneteskeyvalue["object_type"] = item.GetObjectType()
					kuberneteskeyvalue["value"] = item.GetValue()
					kuberneteskeyvalues = append(kuberneteskeyvalues, kuberneteskeyvalue)
				}
				return kuberneteskeyvalues
			})(item.GetOverrideSets(), d)
			kubernetesaddonconfiguration["overrides"] = item.GetOverrides()
			kubernetesaddonconfiguration["release_name"] = item.GetReleaseName()
			kubernetesaddonconfiguration["release_namespace"] = item.GetReleaseNamespace()
			kubernetesaddonconfiguration["upgrade_strategy"] = item.GetUpgradeStrategy()

			kubernetesaddonconfigurations = append(kubernetesaddonconfigurations, kubernetesaddonconfiguration)
			return kubernetesaddonconfigurations
		})(item.GetAddonConfiguration(), d)
		kubernetesessentialaddon["addon_definition"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			var ret models.MoMoRef
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.GetClassId()
			momoref["moid"] = item.GetMoid()
			momoref["object_type"] = item.GetObjectType()
			momoref["selector"] = item.GetSelector()

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetAddonDefinition(), d)
		kubernetesessentialaddon["class_id"] = item.GetClassId()
		kubernetesessentialaddon["name"] = item.GetName()
		kubernetesessentialaddon["object_type"] = item.GetObjectType()
		kubernetesessentialaddons = append(kubernetesessentialaddons, kubernetesessentialaddon)
	}
	return kubernetesessentialaddons
}
func flattenListKubernetesNodeAddress(p []models.KubernetesNodeAddress, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodeaddresss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		kubernetesnodeaddress := make(map[string]interface{})
		kubernetesnodeaddress["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesnodeaddress["address"] = item.GetAddress()
		kubernetesnodeaddress["class_id"] = item.GetClassId()
		kubernetesnodeaddress["object_type"] = item.GetObjectType()
		kubernetesnodeaddress["type"] = item.GetType()
		kubernetesnodeaddresss = append(kubernetesnodeaddresss, kubernetesnodeaddress)
	}
	return kubernetesnodeaddresss
}
func flattenListKubernetesNodeGroupLabel(p []models.KubernetesNodeGroupLabel, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodegrouplabels []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		kubernetesnodegrouplabel := make(map[string]interface{})
		kubernetesnodegrouplabel["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesnodegrouplabel["class_id"] = item.GetClassId()
		kubernetesnodegrouplabel["key"] = item.GetKey()
		kubernetesnodegrouplabel["object_type"] = item.GetObjectType()
		kubernetesnodegrouplabel["value"] = item.GetValue()
		kubernetesnodegrouplabels = append(kubernetesnodegrouplabels, kubernetesnodegrouplabel)
	}
	return kubernetesnodegrouplabels
}
func flattenListKubernetesNodeGroupProfileRelationship(p []models.KubernetesNodeGroupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodegroupprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesnodegroupprofilerelationship := flattenMoMoRef(item)
		kubernetesnodegroupprofilerelationships = append(kubernetesnodegroupprofilerelationships, kubernetesnodegroupprofilerelationship)
	}
	return kubernetesnodegroupprofilerelationships
}
func flattenListKubernetesNodeGroupTaint(p []models.KubernetesNodeGroupTaint, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodegrouptaints []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		kubernetesnodegrouptaint := make(map[string]interface{})
		kubernetesnodegrouptaint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesnodegrouptaint["class_id"] = item.GetClassId()
		kubernetesnodegrouptaint["effect"] = item.GetEffect()
		kubernetesnodegrouptaint["key"] = item.GetKey()
		kubernetesnodegrouptaint["object_type"] = item.GetObjectType()
		kubernetesnodegrouptaint["value"] = item.GetValue()
		kubernetesnodegrouptaints = append(kubernetesnodegrouptaints, kubernetesnodegrouptaint)
	}
	return kubernetesnodegrouptaints
}
func flattenListKubernetesNodeProfileRelationship(p []models.KubernetesNodeProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodeprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesnodeprofilerelationship := flattenMoMoRef(item)
		kubernetesnodeprofilerelationships = append(kubernetesnodeprofilerelationships, kubernetesnodeprofilerelationship)
	}
	return kubernetesnodeprofilerelationships
}
func flattenListKubernetesNodeStatus(p []models.KubernetesNodeStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodestatuss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		kubernetesnodestatus := make(map[string]interface{})
		kubernetesnodestatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesnodestatus["class_id"] = item.GetClassId()
		kubernetesnodestatus["object_type"] = item.GetObjectType()
		kubernetesnodestatus["status"] = item.GetStatus()
		kubernetesnodestatus["type"] = item.GetType()
		kubernetesnodestatuss = append(kubernetesnodestatuss, kubernetesnodestatus)
	}
	return kubernetesnodestatuss
}
func flattenListKubernetesVirtualMachineInfrastructureProviderRelationship(p []models.KubernetesVirtualMachineInfrastructureProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesvirtualmachineinfrastructureproviderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		kubernetesvirtualmachineinfrastructureproviderrelationship := flattenMoMoRef(item)
		kubernetesvirtualmachineinfrastructureproviderrelationships = append(kubernetesvirtualmachineinfrastructureproviderrelationships, kubernetesvirtualmachineinfrastructureproviderrelationship)
	}
	return kubernetesvirtualmachineinfrastructureproviderrelationships
}
func flattenListLicenseLicenseInfoRelationship(p []models.LicenseLicenseInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenselicenseinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		licenselicenseinforelationship := flattenMoMoRef(item)
		licenselicenseinforelationships = append(licenselicenseinforelationships, licenselicenseinforelationship)
	}
	return licenselicenseinforelationships
}
func flattenListMacpoolBlock(p []models.MacpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var macpoolblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		macpoolblock := make(map[string]interface{})
		macpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		macpoolblock["class_id"] = item.GetClassId()
		macpoolblock["from"] = item.GetFrom()
		macpoolblock["object_type"] = item.GetObjectType()
		macpoolblock["size"] = item.GetSize()
		macpoolblock["to"] = item.GetTo()
		macpoolblocks = append(macpoolblocks, macpoolblock)
	}
	return macpoolblocks
}
func flattenListMacpoolIdBlockRelationship(p []models.MacpoolIdBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolidblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		macpoolidblockrelationship := flattenMoMoRef(item)
		macpoolidblockrelationships = append(macpoolidblockrelationships, macpoolidblockrelationship)
	}
	return macpoolidblockrelationships
}
func flattenListManagementInterfaceRelationship(p []models.ManagementInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managementinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		managementinterfacerelationship := flattenMoMoRef(item)
		managementinterfacerelationships = append(managementinterfacerelationships, managementinterfacerelationship)
	}
	return managementinterfacerelationships
}
func flattenListMemoryArrayRelationship(p []models.MemoryArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memoryarrayrelationship := flattenMoMoRef(item)
		memoryarrayrelationships = append(memoryarrayrelationships, memoryarrayrelationship)
	}
	return memoryarrayrelationships
}
func flattenListMemoryPersistentMemoryGoal(p []models.MemoryPersistentMemoryGoal, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorygoals []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		memorypersistentmemorygoal := make(map[string]interface{})
		memorypersistentmemorygoal["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		memorypersistentmemorygoal["class_id"] = item.GetClassId()
		memorypersistentmemorygoal["memory_mode_percentage"] = item.GetMemoryModePercentage()
		memorypersistentmemorygoal["object_type"] = item.GetObjectType()
		memorypersistentmemorygoal["persistent_memory_type"] = item.GetPersistentMemoryType()
		memorypersistentmemorygoal["socket_id"] = item.GetSocketId()
		memorypersistentmemorygoals = append(memorypersistentmemorygoals, memorypersistentmemorygoal)
	}
	return memorypersistentmemorygoals
}
func flattenListMemoryPersistentMemoryLogicalNamespace(p []models.MemoryPersistentMemoryLogicalNamespace, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylogicalnamespaces []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		memorypersistentmemorylogicalnamespace := make(map[string]interface{})
		memorypersistentmemorylogicalnamespace["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		memorypersistentmemorylogicalnamespace["capacity"] = item.GetCapacity()
		memorypersistentmemorylogicalnamespace["class_id"] = item.GetClassId()
		memorypersistentmemorylogicalnamespace["mode"] = item.GetMode()
		memorypersistentmemorylogicalnamespace["name"] = item.GetName()
		memorypersistentmemorylogicalnamespace["object_type"] = item.GetObjectType()
		memorypersistentmemorylogicalnamespace["socket_id"] = item.GetSocketId()
		memorypersistentmemorylogicalnamespace["socket_memory_id"] = item.GetSocketMemoryId()
		memorypersistentmemorylogicalnamespaces = append(memorypersistentmemorylogicalnamespaces, memorypersistentmemorylogicalnamespace)
	}
	return memorypersistentmemorylogicalnamespaces
}
func flattenListMemoryPersistentMemoryNamespaceRelationship(p []models.MemoryPersistentMemoryNamespaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemorynamespacerelationship := flattenMoMoRef(item)
		memorypersistentmemorynamespacerelationships = append(memorypersistentmemorynamespacerelationships, memorypersistentmemorynamespacerelationship)
	}
	return memorypersistentmemorynamespacerelationships
}
func flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p []models.MemoryPersistentMemoryNamespaceConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespaceconfigresultrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemorynamespaceconfigresultrelationship := flattenMoMoRef(item)
		memorypersistentmemorynamespaceconfigresultrelationships = append(memorypersistentmemorynamespaceconfigresultrelationships, memorypersistentmemorynamespaceconfigresultrelationship)
	}
	return memorypersistentmemorynamespaceconfigresultrelationships
}
func flattenListMemoryPersistentMemoryRegionRelationship(p []models.MemoryPersistentMemoryRegionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemoryregionrelationship := flattenMoMoRef(item)
		memorypersistentmemoryregionrelationships = append(memorypersistentmemoryregionrelationships, memorypersistentmemoryregionrelationship)
	}
	return memorypersistentmemoryregionrelationships
}
func flattenListMemoryPersistentMemoryUnitRelationship(p []models.MemoryPersistentMemoryUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemoryunitrelationship := flattenMoMoRef(item)
		memorypersistentmemoryunitrelationships = append(memorypersistentmemoryunitrelationships, memorypersistentmemoryunitrelationship)
	}
	return memorypersistentmemoryunitrelationships
}
func flattenListMemoryUnitRelationship(p []models.MemoryUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memoryunitrelationship := flattenMoMoRef(item)
		memoryunitrelationships = append(memoryunitrelationships, memoryunitrelationship)
	}
	return memoryunitrelationships
}
func flattenListMetaAccessPrivilege(p []models.MetaAccessPrivilege, d *schema.ResourceData) []map[string]interface{} {
	var metaaccessprivileges []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metaaccessprivilege := make(map[string]interface{})
		metaaccessprivilege["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metaaccessprivilege["class_id"] = item.GetClassId()
		metaaccessprivilege["method"] = item.GetMethod()
		metaaccessprivilege["object_type"] = item.GetObjectType()
		metaaccessprivilege["privilege"] = item.GetPrivilege()
		metaaccessprivileges = append(metaaccessprivileges, metaaccessprivilege)
	}
	return metaaccessprivileges
}
func flattenListMetaDisplayNameDefinition(p []models.MetaDisplayNameDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metadisplaynamedefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metadisplaynamedefinition := make(map[string]interface{})
		metadisplaynamedefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metadisplaynamedefinition["class_id"] = item.GetClassId()
		metadisplaynamedefinition["format"] = item.GetFormat()
		metadisplaynamedefinition["include_ancestor"] = item.GetIncludeAncestor()
		metadisplaynamedefinition["name"] = item.GetName()
		metadisplaynamedefinition["object_type"] = item.GetObjectType()
		metadisplaynamedefinitions = append(metadisplaynamedefinitions, metadisplaynamedefinition)
	}
	return metadisplaynamedefinitions
}
func flattenListMetaPropDefinition(p []models.MetaPropDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metapropdefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metapropdefinition := make(map[string]interface{})
		metapropdefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metapropdefinition["api_access"] = item.GetApiAccess()
		metapropdefinition["class_id"] = item.GetClassId()
		metapropdefinition["name"] = item.GetName()
		metapropdefinition["object_type"] = item.GetObjectType()
		metapropdefinition["op_security"] = item.GetOpSecurity()
		metapropdefinition["search_weight"] = item.GetSearchWeight()
		metapropdefinitions = append(metapropdefinitions, metapropdefinition)
	}
	return metapropdefinitions
}
func flattenListMetaRelationshipDefinition(p []models.MetaRelationshipDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metarelationshipdefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metarelationshipdefinition := make(map[string]interface{})
		metarelationshipdefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metarelationshipdefinition["api_access"] = item.GetApiAccess()
		metarelationshipdefinition["class_id"] = item.GetClassId()
		metarelationshipdefinition["collection"] = item.GetCollection()
		metarelationshipdefinition["export"] = item.GetExport()
		metarelationshipdefinition["export_with_peer"] = item.GetExportWithPeer()
		metarelationshipdefinition["name"] = item.GetName()
		metarelationshipdefinition["object_type"] = item.GetObjectType()
		metarelationshipdefinition["peer_rel_name"] = item.GetPeerRelName()
		metarelationshipdefinition["peer_sync"] = item.GetPeerSync()
		metarelationshipdefinition["type"] = item.GetType()
		metarelationshipdefinitions = append(metarelationshipdefinitions, metarelationshipdefinition)
	}
	return metarelationshipdefinitions
}
func flattenListMoBaseMoRelationship(p []models.MoBaseMoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		mobasemorelationship := flattenMoMoRef(item)
		mobasemorelationships = append(mobasemorelationships, mobasemorelationship)
	}
	return mobasemorelationships
}
func flattenListMoTag(p []models.MoTag, d *schema.ResourceData) []map[string]interface{} {
	var motags []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		motag := make(map[string]interface{})
		motag["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		motag["key"] = item.GetKey()
		motag["value"] = item.GetValue()
		motags = append(motags, motag)
	}
	return motags
}
func flattenListNetworkElementRelationship(p []models.NetworkElementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		networkelementrelationship := flattenMoMoRef(item)
		networkelementrelationships = append(networkelementrelationships, networkelementrelationship)
	}
	return networkelementrelationships
}
func flattenListNiaapiDetail(p []models.NiaapiDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapidetails []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		niaapidetail := make(map[string]interface{})
		niaapidetail["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapidetail["chksum"] = item.GetChksum()
		niaapidetail["class_id"] = item.GetClassId()
		niaapidetail["filename"] = item.GetFilename()
		niaapidetail["name"] = item.GetName()
		niaapidetail["object_type"] = item.GetObjectType()
		niaapidetails = append(niaapidetails, niaapidetail)
	}
	return niaapidetails
}
func flattenListNiaapiRevisionInfo(p []models.NiaapiRevisionInfo, d *schema.ResourceData) []map[string]interface{} {
	var niaapirevisioninfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		niaapirevisioninfo := make(map[string]interface{})
		niaapirevisioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapirevisioninfo["class_id"] = item.GetClassId()
		niaapirevisioninfo["date_published"] = item.GetDatePublished()
		niaapirevisioninfo["object_type"] = item.GetObjectType()
		niaapirevisioninfo["revision_comment"] = item.GetRevisionComment()
		niaapirevisioninfo["revision_no"] = item.GetRevisionNo()
		niaapirevisioninfos = append(niaapirevisioninfos, niaapirevisioninfo)
	}
	return niaapirevisioninfos
}
func flattenListNiatelemetryLogicalLink(p []models.NiatelemetryLogicalLink, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrylogicallinks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		niatelemetrylogicallink := make(map[string]interface{})
		niatelemetrylogicallink["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niatelemetrylogicallink["class_id"] = item.GetClassId()
		niatelemetrylogicallink["db_id"] = item.GetDbId()
		niatelemetrylogicallink["is_present"] = item.GetIsPresent()
		niatelemetrylogicallink["link_addr1"] = item.GetLinkAddr1()
		niatelemetrylogicallink["link_addr2"] = item.GetLinkAddr2()
		niatelemetrylogicallink["link_state"] = item.GetLinkState()
		niatelemetrylogicallink["link_type"] = item.GetLinkType()
		niatelemetrylogicallink["object_type"] = item.GetObjectType()
		niatelemetrylogicallink["uptime"] = item.GetUptime()
		niatelemetrylogicallinks = append(niatelemetrylogicallinks, niatelemetrylogicallink)
	}
	return niatelemetrylogicallinks
}
func flattenListNotificationAbstractCondition(p []models.NotificationAbstractCondition, d *schema.ResourceData) []map[string]interface{} {
	var notificationabstractconditions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		notificationabstractcondition := make(map[string]interface{})
		notificationabstractcondition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		notificationabstractcondition["class_id"] = item.GetClassId()
		notificationabstractcondition["object_type"] = item.GetObjectType()
		notificationabstractconditions = append(notificationabstractconditions, notificationabstractcondition)
	}
	return notificationabstractconditions
}
func flattenListNotificationAction(p []models.NotificationAction, d *schema.ResourceData) []map[string]interface{} {
	var notificationactions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		notificationaction := make(map[string]interface{})
		notificationaction["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		notificationaction["class_id"] = item.GetClassId()
		notificationaction["object_type"] = item.GetObjectType()
		notificationactions = append(notificationactions, notificationaction)
	}
	return notificationactions
}
func flattenListNtpAuthNtpServer(p []models.NtpAuthNtpServer, d *schema.ResourceData) []map[string]interface{} {
	var ntpauthntpservers []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ntpauthntpserver := make(map[string]interface{})
		ntpauthntpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ntpauthntpserver["class_id"] = item.GetClassId()
		ntpauthntpserver["key_type"] = item.GetKeyType()
		ntpauthntpserver["object_type"] = item.GetObjectType()
		ntpauthntpserver["server_name"] = item.GetServerName()
		ntpauthntpserver["sym_key_id"] = item.GetSymKeyId()
		ntpauthntpserver["sym_key_value"] = item.GetSymKeyValue()
		ntpauthntpservers = append(ntpauthntpservers, ntpauthntpserver)
	}
	return ntpauthntpservers
}
func flattenListOnpremImagePackage(p []models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var onpremimagepackages []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremimagepackage := make(map[string]interface{})
		onpremimagepackage["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremimagepackage["class_id"] = item.GetClassId()
		onpremimagepackage["file_path"] = item.GetFilePath()
		onpremimagepackage["file_sha"] = item.GetFileSha()
		onpremimagepackage["file_size"] = item.GetFileSize()
		onpremimagepackage["file_time"] = item.GetFileTime()
		onpremimagepackage["filename"] = item.GetFilename()
		onpremimagepackage["name"] = item.GetName()
		onpremimagepackage["object_type"] = item.GetObjectType()
		onpremimagepackage["package_type"] = item.GetPackageType()
		onpremimagepackage["nr_version"] = item.GetVersion()
		onpremimagepackages = append(onpremimagepackages, onpremimagepackage)
	}
	return onpremimagepackages
}
func flattenListOnpremUpgradeNote(p []models.OnpremUpgradeNote, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradenotes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremupgradenote := make(map[string]interface{})
		onpremupgradenote["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremupgradenote["class_id"] = item.GetClassId()
		onpremupgradenote["message"] = item.GetMessage()
		onpremupgradenote["object_type"] = item.GetObjectType()
		onpremupgradenotes = append(onpremupgradenotes, onpremupgradenote)
	}
	return onpremupgradenotes
}
func flattenListOnpremUpgradePhase(p []models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremupgradephase := make(map[string]interface{})
		onpremupgradephase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremupgradephase["class_id"] = item.GetClassId()
		onpremupgradephase["elapsed_time"] = item.GetElapsedTime()
		onpremupgradephase["end_time"] = item.GetEndTime()
		onpremupgradephase["failed"] = item.GetFailed()
		onpremupgradephase["message"] = item.GetMessage()
		onpremupgradephase["name"] = item.GetName()
		onpremupgradephase["object_type"] = item.GetObjectType()
		onpremupgradephase["start_time"] = item.GetStartTime()
		onpremupgradephases = append(onpremupgradephases, onpremupgradephase)
	}
	return onpremupgradephases
}
func flattenListOprsKvpair(p []models.OprsKvpair, d *schema.ResourceData) []map[string]interface{} {
	var oprskvpairs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		oprskvpair := make(map[string]interface{})
		oprskvpair["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		oprskvpair["class_id"] = item.GetClassId()
		oprskvpair["key"] = item.GetKey()
		oprskvpair["object_type"] = item.GetObjectType()
		oprskvpair["value"] = item.GetValue()
		oprskvpairs = append(oprskvpairs, oprskvpair)
	}
	return oprskvpairs
}
func flattenListOrganizationOrganizationRelationship(p []models.OrganizationOrganizationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		organizationorganizationrelationship := flattenMoMoRef(item)
		organizationorganizationrelationships = append(organizationorganizationrelationships, organizationorganizationrelationship)
	}
	return organizationorganizationrelationships
}
func flattenListOsConfigurationFileRelationship(p []models.OsConfigurationFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		osconfigurationfilerelationship := flattenMoMoRef(item)
		osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	}
	return osconfigurationfilerelationships
}
func flattenListOsDistributionRelationship(p []models.OsDistributionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osdistributionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		osdistributionrelationship := flattenMoMoRef(item)
		osdistributionrelationships = append(osdistributionrelationships, osdistributionrelationship)
	}
	return osdistributionrelationships
}
func flattenListOsPlaceHolder(p []models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
	var osplaceholders []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		osplaceholder := make(map[string]interface{})
		osplaceholder["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osplaceholder["class_id"] = item.GetClassId()
		osplaceholder["is_value_set"] = item.GetIsValueSet()
		osplaceholder["object_type"] = item.GetObjectType()
		osplaceholder["type"] = (func(p models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {
			var workflowprimitivedatatypes []map[string]interface{}
			var ret models.WorkflowPrimitiveDataType
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowprimitivedatatype := make(map[string]interface{})
			workflowprimitivedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowprimitivedatatype["class_id"] = item.GetClassId()
			workflowprimitivedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.GetClassId()
				workflowdefaultvalue["is_value_set"] = item.GetIsValueSet()
				workflowdefaultvalue["object_type"] = item.GetObjectType()
				workflowdefaultvalue["override"] = item.GetOverride()
				workflowdefaultvalue["value"] = flattenAdditionalProperties(item.Value)

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowprimitivedatatype["description"] = item.GetDescription()
			workflowprimitivedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
				var workflowdisplaymetas []map[string]interface{}
				var ret models.WorkflowDisplayMeta
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdisplaymeta := make(map[string]interface{})
				workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdisplaymeta["class_id"] = item.GetClassId()
				workflowdisplaymeta["inventory_selector"] = item.GetInventorySelector()
				workflowdisplaymeta["object_type"] = item.GetObjectType()
				workflowdisplaymeta["widget_type"] = item.GetWidgetType()

				workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
				return workflowdisplaymetas
			})(item.GetDisplayMeta(), d)
			workflowprimitivedatatype["input_parameters"] = flattenAdditionalProperties(item.InputParameters)
			workflowprimitivedatatype["label"] = item.GetLabel()
			workflowprimitivedatatype["name"] = item.GetName()
			workflowprimitivedatatype["object_type"] = item.GetObjectType()
			workflowprimitivedatatype["properties"] = (func(p models.WorkflowPrimitiveDataProperty, d *schema.ResourceData) []map[string]interface{} {
				var workflowprimitivedatapropertys []map[string]interface{}
				var ret models.WorkflowPrimitiveDataProperty
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowprimitivedataproperty := make(map[string]interface{})
				workflowprimitivedataproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowprimitivedataproperty["class_id"] = item.GetClassId()
				workflowprimitivedataproperty["constraints"] = (func(p models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {
					var workflowconstraintss []map[string]interface{}
					var ret models.WorkflowConstraints
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					workflowconstraints := make(map[string]interface{})
					workflowconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					workflowconstraints["class_id"] = item.GetClassId()
					workflowconstraints["enum_list"] = (func(p []models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var workflowenumentrys []map[string]interface{}
						if len(p) == 0 {
							return nil
						}
						for _, item := range p {
							workflowenumentry := make(map[string]interface{})
							workflowenumentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							workflowenumentry["class_id"] = item.GetClassId()
							workflowenumentry["label"] = item.GetLabel()
							workflowenumentry["object_type"] = item.GetObjectType()
							workflowenumentry["value"] = item.GetValue()
							workflowenumentrys = append(workflowenumentrys, workflowenumentry)
						}
						return workflowenumentrys
					})(item.GetEnumList(), d)
					workflowconstraints["max"] = item.GetMax()
					workflowconstraints["min"] = item.GetMin()
					workflowconstraints["object_type"] = item.GetObjectType()
					workflowconstraints["regex"] = item.GetRegex()

					workflowconstraintss = append(workflowconstraintss, workflowconstraints)
					return workflowconstraintss
				})(item.GetConstraints(), d)
				workflowprimitivedataproperty["inventory_selector"] = (func(p []models.WorkflowMoReferenceProperty, d *schema.ResourceData) []map[string]interface{} {
					var workflowmoreferencepropertys []map[string]interface{}
					if len(p) == 0 {
						return nil
					}
					for _, item := range p {
						workflowmoreferenceproperty := make(map[string]interface{})
						workflowmoreferenceproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						workflowmoreferenceproperty["class_id"] = item.GetClassId()
						workflowmoreferenceproperty["display_attributes"] = item.GetDisplayAttributes()
						workflowmoreferenceproperty["object_type"] = item.GetObjectType()
						workflowmoreferenceproperty["selector"] = item.GetSelector()
						workflowmoreferenceproperty["value_attribute"] = item.GetValueAttribute()
						workflowmoreferencepropertys = append(workflowmoreferencepropertys, workflowmoreferenceproperty)
					}
					return workflowmoreferencepropertys
				})(item.GetInventorySelector(), d)
				workflowprimitivedataproperty["object_type"] = item.GetObjectType()
				workflowprimitivedataproperty["secure"] = item.GetSecure()
				workflowprimitivedataproperty["type"] = item.GetType()

				workflowprimitivedatapropertys = append(workflowprimitivedatapropertys, workflowprimitivedataproperty)
				return workflowprimitivedatapropertys
			})(item.GetProperties(), d)
			workflowprimitivedatatype["required"] = item.GetRequired()

			workflowprimitivedatatypes = append(workflowprimitivedatatypes, workflowprimitivedatatype)
			return workflowprimitivedatatypes
		})(item.GetType(), d)
		osplaceholder["value"] = flattenAdditionalProperties(item.Value)
		osplaceholders = append(osplaceholders, osplaceholder)
	}
	return osplaceholders
}
func flattenListOsServerConfig(p []models.OsServerConfig, d *schema.ResourceData) []map[string]interface{} {
	var osserverconfigs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		osserverconfig := make(map[string]interface{})
		osserverconfig["additional_parameters"] = (func(p []models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
			var osplaceholders []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				osplaceholder := make(map[string]interface{})
				osplaceholder["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				osplaceholder["class_id"] = item.GetClassId()
				osplaceholder["is_value_set"] = item.GetIsValueSet()
				osplaceholder["object_type"] = item.GetObjectType()
				osplaceholder["type"] = (func(p models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {
					var workflowprimitivedatatypes []map[string]interface{}
					var ret models.WorkflowPrimitiveDataType
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					workflowprimitivedatatype := make(map[string]interface{})
					workflowprimitivedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					workflowprimitivedatatype["class_id"] = item.GetClassId()
					workflowprimitivedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
						var workflowdefaultvalues []map[string]interface{}
						var ret models.WorkflowDefaultValue
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						workflowdefaultvalue := make(map[string]interface{})
						workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						workflowdefaultvalue["class_id"] = item.GetClassId()
						workflowdefaultvalue["is_value_set"] = item.GetIsValueSet()
						workflowdefaultvalue["object_type"] = item.GetObjectType()
						workflowdefaultvalue["override"] = item.GetOverride()
						workflowdefaultvalue["value"] = flattenAdditionalProperties(item.Value)

						workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
						return workflowdefaultvalues
					})(item.GetDefault(), d)
					workflowprimitivedatatype["description"] = item.GetDescription()
					workflowprimitivedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
						var workflowdisplaymetas []map[string]interface{}
						var ret models.WorkflowDisplayMeta
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						workflowdisplaymeta := make(map[string]interface{})
						workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						workflowdisplaymeta["class_id"] = item.GetClassId()
						workflowdisplaymeta["inventory_selector"] = item.GetInventorySelector()
						workflowdisplaymeta["object_type"] = item.GetObjectType()
						workflowdisplaymeta["widget_type"] = item.GetWidgetType()

						workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
						return workflowdisplaymetas
					})(item.GetDisplayMeta(), d)
					workflowprimitivedatatype["input_parameters"] = flattenAdditionalProperties(item.InputParameters)
					workflowprimitivedatatype["label"] = item.GetLabel()
					workflowprimitivedatatype["name"] = item.GetName()
					workflowprimitivedatatype["object_type"] = item.GetObjectType()
					workflowprimitivedatatype["properties"] = (func(p models.WorkflowPrimitiveDataProperty, d *schema.ResourceData) []map[string]interface{} {
						var workflowprimitivedatapropertys []map[string]interface{}
						var ret models.WorkflowPrimitiveDataProperty
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						workflowprimitivedataproperty := make(map[string]interface{})
						workflowprimitivedataproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						workflowprimitivedataproperty["class_id"] = item.GetClassId()
						workflowprimitivedataproperty["constraints"] = (func(p models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {
							var workflowconstraintss []map[string]interface{}
							var ret models.WorkflowConstraints
							if reflect.DeepEqual(ret, p) {
								return nil
							}
							item := p
							workflowconstraints := make(map[string]interface{})
							workflowconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							workflowconstraints["class_id"] = item.GetClassId()
							workflowconstraints["enum_list"] = (func(p []models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
								var workflowenumentrys []map[string]interface{}
								if len(p) == 0 {
									return nil
								}
								for _, item := range p {
									workflowenumentry := make(map[string]interface{})
									workflowenumentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
									workflowenumentry["class_id"] = item.GetClassId()
									workflowenumentry["label"] = item.GetLabel()
									workflowenumentry["object_type"] = item.GetObjectType()
									workflowenumentry["value"] = item.GetValue()
									workflowenumentrys = append(workflowenumentrys, workflowenumentry)
								}
								return workflowenumentrys
							})(item.GetEnumList(), d)
							workflowconstraints["max"] = item.GetMax()
							workflowconstraints["min"] = item.GetMin()
							workflowconstraints["object_type"] = item.GetObjectType()
							workflowconstraints["regex"] = item.GetRegex()

							workflowconstraintss = append(workflowconstraintss, workflowconstraints)
							return workflowconstraintss
						})(item.GetConstraints(), d)
						workflowprimitivedataproperty["inventory_selector"] = (func(p []models.WorkflowMoReferenceProperty, d *schema.ResourceData) []map[string]interface{} {
							var workflowmoreferencepropertys []map[string]interface{}
							if len(p) == 0 {
								return nil
							}
							for _, item := range p {
								workflowmoreferenceproperty := make(map[string]interface{})
								workflowmoreferenceproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
								workflowmoreferenceproperty["class_id"] = item.GetClassId()
								workflowmoreferenceproperty["display_attributes"] = item.GetDisplayAttributes()
								workflowmoreferenceproperty["object_type"] = item.GetObjectType()
								workflowmoreferenceproperty["selector"] = item.GetSelector()
								workflowmoreferenceproperty["value_attribute"] = item.GetValueAttribute()
								workflowmoreferencepropertys = append(workflowmoreferencepropertys, workflowmoreferenceproperty)
							}
							return workflowmoreferencepropertys
						})(item.GetInventorySelector(), d)
						workflowprimitivedataproperty["object_type"] = item.GetObjectType()
						workflowprimitivedataproperty["secure"] = item.GetSecure()
						workflowprimitivedataproperty["type"] = item.GetType()

						workflowprimitivedatapropertys = append(workflowprimitivedatapropertys, workflowprimitivedataproperty)
						return workflowprimitivedatapropertys
					})(item.GetProperties(), d)
					workflowprimitivedatatype["required"] = item.GetRequired()

					workflowprimitivedatatypes = append(workflowprimitivedatatypes, workflowprimitivedatatype)
					return workflowprimitivedatatypes
				})(item.GetType(), d)
				osplaceholder["value"] = flattenAdditionalProperties(item.Value)
				osplaceholders = append(osplaceholders, osplaceholder)
			}
			return osplaceholders
		})(item.GetAdditionalParameters(), d)
		osserverconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osserverconfig["answers"] = (func(p models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {
			var osanswerss []map[string]interface{}
			var ret models.OsAnswers
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			osanswers := make(map[string]interface{})
			osanswers["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			answer_file_x, exists := d.GetOk("answers")
			if exists && answer_file_x != nil {
				answer_file_y := answer_file_x.([]interface{})[0].(map[string]interface{})
				osanswers["answer_file"] = answer_file_y["answer_file"]
			}
			osanswers["class_id"] = item.GetClassId()
			osanswers["hostname"] = item.GetHostname()
			osanswers["ip_config_type"] = item.GetIpConfigType()
			osanswers["ip_configuration"] = (func(p models.OsIpConfiguration, d *schema.ResourceData) []map[string]interface{} {
				var osipconfigurations []map[string]interface{}
				var ret models.OsIpConfiguration
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				osipconfiguration := make(map[string]interface{})
				osipconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				osipconfiguration["class_id"] = item.GetClassId()
				osipconfiguration["object_type"] = item.GetObjectType()

				osipconfigurations = append(osipconfigurations, osipconfiguration)
				return osipconfigurations
			})(item.GetIpConfiguration(), d)
			osanswers["is_answer_file_set"] = item.GetIsAnswerFileSet()
			osanswers["is_root_password_crypted"] = item.GetIsRootPasswordCrypted()
			osanswers["is_root_password_set"] = item.GetIsRootPasswordSet()
			osanswers["nameserver"] = item.GetNameserver()
			osanswers["object_type"] = item.GetObjectType()
			osanswers["product_key"] = item.GetProductKey()
			root_password_x, exists := d.GetOk("answers")
			if exists && root_password_x != nil {
				root_password_y := root_password_x.([]interface{})[0].(map[string]interface{})
				osanswers["root_password"] = root_password_y["root_password"]
			}
			osanswers["nr_source"] = item.GetSource()

			osanswerss = append(osanswerss, osanswers)
			return osanswerss
		})(item.GetAnswers(), d)
		osserverconfig["class_id"] = item.GetClassId()
		osserverconfig["error_msgs"] = item.GetErrorMsgs()
		osserverconfig["install_target"] = item.GetInstallTarget()
		osserverconfig["object_type"] = item.GetObjectType()
		osserverconfig["processed_install_target"] = (func(p models.OsInstallTarget, d *schema.ResourceData) []map[string]interface{} {
			var osinstalltargets []map[string]interface{}
			var ret models.OsInstallTarget
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			osinstalltarget := make(map[string]interface{})
			osinstalltarget["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			osinstalltarget["class_id"] = item.GetClassId()
			osinstalltarget["object_type"] = item.GetObjectType()

			osinstalltargets = append(osinstalltargets, osinstalltarget)
			return osinstalltargets
		})(item.GetProcessedInstallTarget(), d)
		osserverconfig["serial_number"] = item.GetSerialNumber()
		osserverconfigs = append(osserverconfigs, osserverconfig)
	}
	return osserverconfigs
}
func flattenListOsValidationInformation(p []models.OsValidationInformation, d *schema.ResourceData) []map[string]interface{} {
	var osvalidationinformations []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		osvalidationinformation := make(map[string]interface{})
		osvalidationinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osvalidationinformation["class_id"] = item.GetClassId()
		osvalidationinformation["error_msg"] = item.GetErrorMsg()
		osvalidationinformation["object_type"] = item.GetObjectType()
		osvalidationinformation["status"] = item.GetStatus()
		osvalidationinformation["step_name"] = item.GetStepName()
		osvalidationinformations = append(osvalidationinformations, osvalidationinformation)
	}
	return osvalidationinformations
}
func flattenListPciCoprocessorCardRelationship(p []models.PciCoprocessorCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcicoprocessorcardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pcicoprocessorcardrelationship := flattenMoMoRef(item)
		pcicoprocessorcardrelationships = append(pcicoprocessorcardrelationships, pcicoprocessorcardrelationship)
	}
	return pcicoprocessorcardrelationships
}
func flattenListPciDeviceRelationship(p []models.PciDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcidevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pcidevicerelationship := flattenMoMoRef(item)
		pcidevicerelationships = append(pcidevicerelationships, pcidevicerelationship)
	}
	return pcidevicerelationships
}
func flattenListPciLinkRelationship(p []models.PciLinkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcilinkrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pcilinkrelationship := flattenMoMoRef(item)
		pcilinkrelationships = append(pcilinkrelationships, pcilinkrelationship)
	}
	return pcilinkrelationships
}
func flattenListPciSwitchRelationship(p []models.PciSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pciswitchrelationship := flattenMoMoRef(item)
		pciswitchrelationships = append(pciswitchrelationships, pciswitchrelationship)
	}
	return pciswitchrelationships
}
func flattenListPolicyAbstractConfigProfileRelationship(p []models.PolicyAbstractConfigProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractconfigprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		policyabstractconfigprofilerelationship := flattenMoMoRef(item)
		policyabstractconfigprofilerelationships = append(policyabstractconfigprofilerelationships, policyabstractconfigprofilerelationship)
	}
	return policyabstractconfigprofilerelationships
}
func flattenListPolicyAbstractPolicyRelationship(p []models.PolicyAbstractPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractpolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		policyabstractpolicyrelationship := flattenMoMoRef(item)
		policyabstractpolicyrelationships = append(policyabstractpolicyrelationships, policyabstractpolicyrelationship)
	}
	return policyabstractpolicyrelationships
}
func flattenListPolicyinventoryJobInfo(p []models.PolicyinventoryJobInfo, d *schema.ResourceData) []map[string]interface{} {
	var policyinventoryjobinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		policyinventoryjobinfo := make(map[string]interface{})
		policyinventoryjobinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		policyinventoryjobinfo["class_id"] = item.GetClassId()
		policyinventoryjobinfo["execution_status"] = item.GetExecutionStatus()
		policyinventoryjobinfo["last_scheduled_time"] = item.GetLastScheduledTime()
		policyinventoryjobinfo["object_type"] = item.GetObjectType()
		policyinventoryjobinfo["policy_id"] = item.GetPolicyId()
		policyinventoryjobinfo["policy_name"] = item.GetPolicyName()
		policyinventoryjobinfos = append(policyinventoryjobinfos, policyinventoryjobinfo)
	}
	return policyinventoryjobinfos
}
func flattenListPortGroupRelationship(p []models.PortGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		portgrouprelationship := flattenMoMoRef(item)
		portgrouprelationships = append(portgrouprelationships, portgrouprelationship)
	}
	return portgrouprelationships
}
func flattenListPortMacBindingRelationship(p []models.PortMacBindingRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portmacbindingrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		portmacbindingrelationship := flattenMoMoRef(item)
		portmacbindingrelationships = append(portmacbindingrelationships, portmacbindingrelationship)
	}
	return portmacbindingrelationships
}
func flattenListPortSubGroupRelationship(p []models.PortSubGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		portsubgrouprelationship := flattenMoMoRef(item)
		portsubgrouprelationships = append(portsubgrouprelationships, portsubgrouprelationship)
	}
	return portsubgrouprelationships
}
func flattenListProcessorUnitRelationship(p []models.ProcessorUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var processorunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		processorunitrelationship := flattenMoMoRef(item)
		processorunitrelationships = append(processorunitrelationships, processorunitrelationship)
	}
	return processorunitrelationships
}
func flattenListRecommendationPhysicalItemRelationship(p []models.RecommendationPhysicalItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recommendationphysicalitemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		recommendationphysicalitemrelationship := flattenMoMoRef(item)
		recommendationphysicalitemrelationships = append(recommendationphysicalitemrelationships, recommendationphysicalitemrelationship)
	}
	return recommendationphysicalitemrelationships
}
func flattenListRecoveryBackupProfileRelationship(p []models.RecoveryBackupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		recoverybackupprofilerelationship := flattenMoMoRef(item)
		recoverybackupprofilerelationships = append(recoverybackupprofilerelationships, recoverybackupprofilerelationship)
	}
	return recoverybackupprofilerelationships
}
func flattenListRecoveryConfigResultEntryRelationship(p []models.RecoveryConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		recoveryconfigresultentryrelationship := flattenMoMoRef(item)
		recoveryconfigresultentryrelationships = append(recoveryconfigresultentryrelationships, recoveryconfigresultentryrelationship)
	}
	return recoveryconfigresultentryrelationships
}
func flattenListResourceGroupRelationship(p []models.ResourceGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		resourcegrouprelationship := flattenMoMoRef(item)
		resourcegrouprelationships = append(resourcegrouprelationships, resourcegrouprelationship)
	}
	return resourcegrouprelationships
}
func flattenListResourcePerTypeCombinedSelector(p []models.ResourcePerTypeCombinedSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourcepertypecombinedselectors []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		resourcepertypecombinedselector := make(map[string]interface{})
		resourcepertypecombinedselector["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		resourcepertypecombinedselector["class_id"] = item.GetClassId()
		resourcepertypecombinedselector["combined_selector"] = item.GetCombinedSelector()
		resourcepertypecombinedselector["empty_filter"] = item.GetEmptyFilter()
		resourcepertypecombinedselector["object_type"] = item.GetObjectType()
		resourcepertypecombinedselector["selector_object_type"] = item.GetSelectorObjectType()
		resourcepertypecombinedselectors = append(resourcepertypecombinedselectors, resourcepertypecombinedselector)
	}
	return resourcepertypecombinedselectors
}
func flattenListResourceSelector(p []models.ResourceSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourceselectors []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		resourceselector := make(map[string]interface{})
		resourceselector["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		resourceselector["class_id"] = item.GetClassId()
		resourceselector["object_type"] = item.GetObjectType()
		resourceselector["selector"] = item.GetSelector()
		resourceselectors = append(resourceselectors, resourceselector)
	}
	return resourceselectors
}
func flattenListSdcardPartition(p []models.SdcardPartition, d *schema.ResourceData) []map[string]interface{} {
	var sdcardpartitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sdcardpartition := make(map[string]interface{})
		sdcardpartition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sdcardpartition["class_id"] = item.GetClassId()
		sdcardpartition["object_type"] = item.GetObjectType()
		sdcardpartition["type"] = item.GetType()
		sdcardpartition["virtual_drives"] = (func(p []models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var sdcardvirtualdrives []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				sdcardvirtualdrive := make(map[string]interface{})
				sdcardvirtualdrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				sdcardvirtualdrive["class_id"] = item.GetClassId()
				sdcardvirtualdrive["enable"] = item.GetEnable()
				sdcardvirtualdrive["object_type"] = item.GetObjectType()
				sdcardvirtualdrives = append(sdcardvirtualdrives, sdcardvirtualdrive)
			}
			return sdcardvirtualdrives
		})(item.GetVirtualDrives(), d)
		sdcardpartitions = append(sdcardpartitions, sdcardpartition)
	}
	return sdcardpartitions
}
func flattenListSdwanNetworkConfigurationType(p []models.SdwanNetworkConfigurationType, d *schema.ResourceData) []map[string]interface{} {
	var sdwannetworkconfigurationtypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sdwannetworkconfigurationtype := make(map[string]interface{})
		sdwannetworkconfigurationtype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sdwannetworkconfigurationtype["class_id"] = item.GetClassId()
		sdwannetworkconfigurationtype["network_type"] = item.GetNetworkType()
		sdwannetworkconfigurationtype["object_type"] = item.GetObjectType()
		sdwannetworkconfigurationtype["port_group"] = item.GetPortGroup()
		sdwannetworkconfigurationtype["vlan"] = item.GetVlan()
		sdwannetworkconfigurationtypes = append(sdwannetworkconfigurationtypes, sdwannetworkconfigurationtype)
	}
	return sdwannetworkconfigurationtypes
}
func flattenListSdwanProfileRelationship(p []models.SdwanProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		sdwanprofilerelationship := flattenMoMoRef(item)
		sdwanprofilerelationships = append(sdwanprofilerelationships, sdwanprofilerelationship)
	}
	return sdwanprofilerelationships
}
func flattenListSdwanRouterNodeRelationship(p []models.SdwanRouterNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouternoderelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		sdwanrouternoderelationship := flattenMoMoRef(item)
		sdwanrouternoderelationships = append(sdwanrouternoderelationships, sdwanrouternoderelationship)
	}
	return sdwanrouternoderelationships
}
func flattenListSdwanTemplateInputsType(p []models.SdwanTemplateInputsType, d *schema.ResourceData) []map[string]interface{} {
	var sdwantemplateinputstypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sdwantemplateinputstype := make(map[string]interface{})
		sdwantemplateinputstype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sdwantemplateinputstype["class_id"] = item.GetClassId()
		sdwantemplateinputstype["editable"] = item.GetEditable()
		sdwantemplateinputstype["key"] = item.GetKey()
		sdwantemplateinputstype["object_type"] = item.GetObjectType()
		sdwantemplateinputstype["required"] = item.GetRequired()
		sdwantemplateinputstype["template"] = item.GetTemplate()
		sdwantemplateinputstype["title"] = item.GetTitle()
		sdwantemplateinputstype["type"] = item.GetType()
		sdwantemplateinputstype["value"] = item.GetValue()
		sdwantemplateinputstypes = append(sdwantemplateinputstypes, sdwantemplateinputstype)
	}
	return sdwantemplateinputstypes
}
func flattenListSecurityUnitRelationship(p []models.SecurityUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var securityunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		securityunitrelationship := flattenMoMoRef(item)
		securityunitrelationships = append(securityunitrelationships, securityunitrelationship)
	}
	return securityunitrelationships
}
func flattenListServerConfigChangeDetailRelationship(p []models.ServerConfigChangeDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigchangedetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		serverconfigchangedetailrelationship := flattenMoMoRef(item)
		serverconfigchangedetailrelationships = append(serverconfigchangedetailrelationships, serverconfigchangedetailrelationship)
	}
	return serverconfigchangedetailrelationships
}
func flattenListServerConfigResultEntryRelationship(p []models.ServerConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		serverconfigresultentryrelationship := flattenMoMoRef(item)
		serverconfigresultentryrelationships = append(serverconfigresultentryrelationships, serverconfigresultentryrelationship)
	}
	return serverconfigresultentryrelationships
}
func flattenListSnmpTrap(p []models.SnmpTrap, d *schema.ResourceData) []map[string]interface{} {
	var snmptraps []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		snmptrap := make(map[string]interface{})
		snmptrap["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		snmptrap["class_id"] = item.GetClassId()
		snmptrap["community"] = item.GetCommunity()
		snmptrap["destination"] = item.GetDestination()
		snmptrap["enabled"] = item.GetEnabled()
		snmptrap["object_type"] = item.GetObjectType()
		snmptrap["port"] = item.GetPort()
		snmptrap["type"] = item.GetType()
		snmptrap["user"] = item.GetUser()
		snmptrap["nr_version"] = item.GetVersion()
		snmptraps = append(snmptraps, snmptrap)
	}
	return snmptraps
}
func flattenListSnmpUser(p []models.SnmpUser, d *schema.ResourceData) []map[string]interface{} {
	var snmpusers []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		snmpuser := make(map[string]interface{})
		snmpuser["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		auth_password_x, exists := d.GetOk("snmp_users")
		if exists && auth_password_x != nil {
			snmpuser["auth_password"] = auth_password_x.([]interface{})[len(snmpusers)].(map[string]interface{})["auth_password"]
		}
		snmpuser["auth_type"] = item.GetAuthType()
		snmpuser["class_id"] = item.GetClassId()
		snmpuser["is_auth_password_set"] = item.GetIsAuthPasswordSet()
		snmpuser["is_privacy_password_set"] = item.GetIsPrivacyPasswordSet()
		snmpuser["name"] = item.GetName()
		snmpuser["object_type"] = item.GetObjectType()
		privacy_password_x, exists := d.GetOk("snmp_users")
		if exists && privacy_password_x != nil {
			snmpuser["privacy_password"] = privacy_password_x.([]interface{})[len(snmpusers)].(map[string]interface{})["privacy_password"]
		}
		snmpuser["privacy_type"] = item.GetPrivacyType()
		snmpuser["security_level"] = item.GetSecurityLevel()
		snmpusers = append(snmpusers, snmpuser)
	}
	return snmpusers
}
func flattenListSoftwareHyperflexDistributableRelationship(p []models.SoftwareHyperflexDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarehyperflexdistributablerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		softwarehyperflexdistributablerelationship := flattenMoMoRef(item)
		softwarehyperflexdistributablerelationships = append(softwarehyperflexdistributablerelationships, softwarehyperflexdistributablerelationship)
	}
	return softwarehyperflexdistributablerelationships
}
func flattenListSoftwareUcsdDistributableRelationship(p []models.SoftwareUcsdDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwareucsddistributablerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		softwareucsddistributablerelationship := flattenMoMoRef(item)
		softwareucsddistributablerelationships = append(softwareucsddistributablerelationships, softwareucsddistributablerelationship)
	}
	return softwareucsddistributablerelationships
}
func flattenListSoftwarerepositoryConstraintModels(p []models.SoftwarerepositoryConstraintModels, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryconstraintmodelss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		softwarerepositoryconstraintmodels := make(map[string]interface{})
		softwarerepositoryconstraintmodels["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		softwarerepositoryconstraintmodels["class_id"] = item.GetClassId()
		softwarerepositoryconstraintmodels["min_version"] = item.GetMinVersion()
		softwarerepositoryconstraintmodels["name"] = item.GetName()
		softwarerepositoryconstraintmodels["object_type"] = item.GetObjectType()
		softwarerepositoryconstraintmodels["platform_regex"] = item.GetPlatformRegex()
		softwarerepositoryconstraintmodels["supported_models"] = item.GetSupportedModels()
		softwarerepositoryconstraintmodelss = append(softwarerepositoryconstraintmodelss, softwarerepositoryconstraintmodels)
	}
	return softwarerepositoryconstraintmodelss
}
func flattenListStorageBaseInitiator(p []models.StorageBaseInitiator, d *schema.ResourceData) []map[string]interface{} {
	var storagebaseinitiators []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagebaseinitiator := make(map[string]interface{})
		storagebaseinitiator["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagebaseinitiator["class_id"] = item.GetClassId()
		storagebaseinitiator["iqn"] = item.GetIqn()
		storagebaseinitiator["name"] = item.GetName()
		storagebaseinitiator["object_type"] = item.GetObjectType()
		storagebaseinitiator["type"] = item.GetType()
		storagebaseinitiator["wwn"] = item.GetWwn()
		storagebaseinitiators = append(storagebaseinitiators, storagebaseinitiator)
	}
	return storagebaseinitiators
}
func flattenListStorageControllerRelationship(p []models.StorageControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagecontrollerrelationship := flattenMoMoRef(item)
		storagecontrollerrelationships = append(storagecontrollerrelationships, storagecontrollerrelationship)
	}
	return storagecontrollerrelationships
}
func flattenListStorageDiskGroupRelationship(p []models.StorageDiskGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagediskgrouprelationship := flattenMoMoRef(item)
		storagediskgrouprelationships = append(storagediskgrouprelationships, storagediskgrouprelationship)
	}
	return storagediskgrouprelationships
}
func flattenListStorageDiskGroupPolicyRelationship(p []models.StorageDiskGroupPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouppolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagediskgrouppolicyrelationship := flattenMoMoRef(item)
		storagediskgrouppolicyrelationships = append(storagediskgrouppolicyrelationships, storagediskgrouppolicyrelationship)
	}
	return storagediskgrouppolicyrelationships
}
func flattenListStorageEnclosureRelationship(p []models.StorageEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageenclosurerelationship := flattenMoMoRef(item)
		storageenclosurerelationships = append(storageenclosurerelationships, storageenclosurerelationship)
	}
	return storageenclosurerelationships
}
func flattenListStorageEnclosureDiskRelationship(p []models.StorageEnclosureDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurediskrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageenclosurediskrelationship := flattenMoMoRef(item)
		storageenclosurediskrelationships = append(storageenclosurediskrelationships, storageenclosurediskrelationship)
	}
	return storageenclosurediskrelationships
}
func flattenListStorageEnclosureDiskSlotEpRelationship(p []models.StorageEnclosureDiskSlotEpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosuredisksloteprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageenclosuredisksloteprelationship := flattenMoMoRef(item)
		storageenclosuredisksloteprelationships = append(storageenclosuredisksloteprelationships, storageenclosuredisksloteprelationship)
	}
	return storageenclosuredisksloteprelationships
}
func flattenListStorageFlexFlashControllerRelationship(p []models.StorageFlexFlashControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashcontrollerrelationship := flattenMoMoRef(item)
		storageflexflashcontrollerrelationships = append(storageflexflashcontrollerrelationships, storageflexflashcontrollerrelationship)
	}
	return storageflexflashcontrollerrelationships
}
func flattenListStorageFlexFlashControllerPropsRelationship(p []models.StorageFlexFlashControllerPropsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerpropsrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashcontrollerpropsrelationship := flattenMoMoRef(item)
		storageflexflashcontrollerpropsrelationships = append(storageflexflashcontrollerpropsrelationships, storageflexflashcontrollerpropsrelationship)
	}
	return storageflexflashcontrollerpropsrelationships
}
func flattenListStorageFlexFlashPhysicalDriveRelationship(p []models.StorageFlexFlashPhysicalDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashphysicaldriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashphysicaldriverelationship := flattenMoMoRef(item)
		storageflexflashphysicaldriverelationships = append(storageflexflashphysicaldriverelationships, storageflexflashphysicaldriverelationship)
	}
	return storageflexflashphysicaldriverelationships
}
func flattenListStorageFlexFlashVirtualDriveRelationship(p []models.StorageFlexFlashVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashvirtualdriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashvirtualdriverelationship := flattenMoMoRef(item)
		storageflexflashvirtualdriverelationships = append(storageflexflashvirtualdriverelationships, storageflexflashvirtualdriverelationship)
	}
	return storageflexflashvirtualdriverelationships
}
func flattenListStorageFlexUtilControllerRelationship(p []models.StorageFlexUtilControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexutilcontrollerrelationship := flattenMoMoRef(item)
		storageflexutilcontrollerrelationships = append(storageflexutilcontrollerrelationships, storageflexutilcontrollerrelationship)
	}
	return storageflexutilcontrollerrelationships
}
func flattenListStorageFlexUtilPhysicalDriveRelationship(p []models.StorageFlexUtilPhysicalDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilphysicaldriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexutilphysicaldriverelationship := flattenMoMoRef(item)
		storageflexutilphysicaldriverelationships = append(storageflexutilphysicaldriverelationships, storageflexutilphysicaldriverelationship)
	}
	return storageflexutilphysicaldriverelationships
}
func flattenListStorageFlexUtilVirtualDriveRelationship(p []models.StorageFlexUtilVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilvirtualdriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexutilvirtualdriverelationship := flattenMoMoRef(item)
		storageflexutilvirtualdriverelationships = append(storageflexutilvirtualdriverelationships, storageflexutilvirtualdriverelationship)
	}
	return storageflexutilvirtualdriverelationships
}
func flattenListStorageHitachiParityGroupRelationship(p []models.StorageHitachiParityGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehitachiparitygrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagehitachiparitygrouprelationship := flattenMoMoRef(item)
		storagehitachiparitygrouprelationships = append(storagehitachiparitygrouprelationships, storagehitachiparitygrouprelationship)
	}
	return storagehitachiparitygrouprelationships
}
func flattenListStorageHyperFlexStorageContainerRelationship(p []models.StorageHyperFlexStorageContainerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehyperflexstoragecontainerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagehyperflexstoragecontainerrelationship := flattenMoMoRef(item)
		storagehyperflexstoragecontainerrelationships = append(storagehyperflexstoragecontainerrelationships, storagehyperflexstoragecontainerrelationship)
	}
	return storagehyperflexstoragecontainerrelationships
}
func flattenListStorageHyperFlexVolumeRelationship(p []models.StorageHyperFlexVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehyperflexvolumerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagehyperflexvolumerelationship := flattenMoMoRef(item)
		storagehyperflexvolumerelationships = append(storagehyperflexvolumerelationships, storagehyperflexvolumerelationship)
	}
	return storagehyperflexvolumerelationships
}
func flattenListStorageItemRelationship(p []models.StorageItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageitemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageitemrelationship := flattenMoMoRef(item)
		storageitemrelationships = append(storageitemrelationships, storageitemrelationship)
	}
	return storageitemrelationships
}
func flattenListStorageLocalDisk(p []models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
	var storagelocaldisks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagelocaldisk := make(map[string]interface{})
		storagelocaldisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagelocaldisk["class_id"] = item.GetClassId()
		storagelocaldisk["object_type"] = item.GetObjectType()
		storagelocaldisk["slot_number"] = item.GetSlotNumber()
		storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
	}
	return storagelocaldisks
}
func flattenListStorageNetAppAggregateRelationship(p []models.StorageNetAppAggregateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappaggregaterelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagenetappaggregaterelationship := flattenMoMoRef(item)
		storagenetappaggregaterelationships = append(storagenetappaggregaterelationships, storagenetappaggregaterelationship)
	}
	return storagenetappaggregaterelationships
}
func flattenListStorageNetAppExportPolicyRule(p []models.StorageNetAppExportPolicyRule, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappexportpolicyrules []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagenetappexportpolicyrule := make(map[string]interface{})
		storagenetappexportpolicyrule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagenetappexportpolicyrule["class_id"] = item.GetClassId()
		storagenetappexportpolicyrule["client_match"] = item.GetClientMatch()
		storagenetappexportpolicyrule["index"] = item.GetIndex()
		storagenetappexportpolicyrule["object_type"] = item.GetObjectType()
		storagenetappexportpolicyrule["ro_rule"] = item.GetRoRule()
		storagenetappexportpolicyrule["rw_rule"] = item.GetRwRule()
		storagenetappexportpolicyrule["super_user"] = item.GetSuperUser()
		storagenetappexportpolicyrule["user"] = item.GetUser()
		storagenetappexportpolicyrules = append(storagenetappexportpolicyrules, storagenetappexportpolicyrule)
	}
	return storagenetappexportpolicyrules
}
func flattenListStorageNetAppInitiatorGroupRelationship(p []models.StorageNetAppInitiatorGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappinitiatorgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagenetappinitiatorgrouprelationship := flattenMoMoRef(item)
		storagenetappinitiatorgrouprelationships = append(storagenetappinitiatorgrouprelationships, storagenetappinitiatorgrouprelationship)
	}
	return storagenetappinitiatorgrouprelationships
}
func flattenListStoragePhysicalDiskRelationship(p []models.StoragePhysicalDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagephysicaldiskrelationship := flattenMoMoRef(item)
		storagephysicaldiskrelationships = append(storagephysicaldiskrelationships, storagephysicaldiskrelationship)
	}
	return storagephysicaldiskrelationships
}
func flattenListStoragePhysicalDiskExtensionRelationship(p []models.StoragePhysicalDiskExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskextensionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagephysicaldiskextensionrelationship := flattenMoMoRef(item)
		storagephysicaldiskextensionrelationships = append(storagephysicaldiskextensionrelationships, storagephysicaldiskextensionrelationship)
	}
	return storagephysicaldiskextensionrelationships
}
func flattenListStoragePhysicalDiskUsageRelationship(p []models.StoragePhysicalDiskUsageRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskusagerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagephysicaldiskusagerelationship := flattenMoMoRef(item)
		storagephysicaldiskusagerelationships = append(storagephysicaldiskusagerelationships, storagephysicaldiskusagerelationship)
	}
	return storagephysicaldiskusagerelationships
}
func flattenListStoragePureHostRelationship(p []models.StoragePureHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagepurehostrelationship := flattenMoMoRef(item)
		storagepurehostrelationships = append(storagepurehostrelationships, storagepurehostrelationship)
	}
	return storagepurehostrelationships
}
func flattenListStoragePureHostGroupRelationship(p []models.StoragePureHostGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagepurehostgrouprelationship := flattenMoMoRef(item)
		storagepurehostgrouprelationships = append(storagepurehostgrouprelationships, storagepurehostgrouprelationship)
	}
	return storagepurehostgrouprelationships
}
func flattenListStoragePureReplicationBlackout(p []models.StoragePureReplicationBlackout, d *schema.ResourceData) []map[string]interface{} {
	var storagepurereplicationblackouts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagepurereplicationblackout := make(map[string]interface{})
		storagepurereplicationblackout["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagepurereplicationblackout["class_id"] = item.GetClassId()
		storagepurereplicationblackout["end"] = item.GetEnd()
		storagepurereplicationblackout["object_type"] = item.GetObjectType()
		storagepurereplicationblackout["start"] = item.GetStart()
		storagepurereplicationblackouts = append(storagepurereplicationblackouts, storagepurereplicationblackout)
	}
	return storagepurereplicationblackouts
}
func flattenListStoragePureVolumeRelationship(p []models.StoragePureVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagepurevolumerelationship := flattenMoMoRef(item)
		storagepurevolumerelationships = append(storagepurevolumerelationships, storagepurevolumerelationship)
	}
	return storagepurevolumerelationships
}
func flattenListStorageSasExpanderRelationship(p []models.StorageSasExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagesasexpanderrelationship := flattenMoMoRef(item)
		storagesasexpanderrelationships = append(storagesasexpanderrelationships, storagesasexpanderrelationship)
	}
	return storagesasexpanderrelationships
}
func flattenListStorageSasPortRelationship(p []models.StorageSasPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagesasportrelationship := flattenMoMoRef(item)
		storagesasportrelationships = append(storagesasportrelationships, storagesasportrelationship)
	}
	return storagesasportrelationships
}
func flattenListStorageSpanRelationship(p []models.StorageSpanRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagespanrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagespanrelationship := flattenMoMoRef(item)
		storagespanrelationships = append(storagespanrelationships, storagespanrelationship)
	}
	return storagespanrelationships
}
func flattenListStorageSpanGroup(p []models.StorageSpanGroup, d *schema.ResourceData) []map[string]interface{} {
	var storagespangroups []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagespangroup := make(map[string]interface{})
		storagespangroup["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagespangroup["class_id"] = item.GetClassId()
		storagespangroup["disks"] = (func(p []models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var storagelocaldisks []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				storagelocaldisk := make(map[string]interface{})
				storagelocaldisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				storagelocaldisk["class_id"] = item.GetClassId()
				storagelocaldisk["object_type"] = item.GetObjectType()
				storagelocaldisk["slot_number"] = item.GetSlotNumber()
				storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
			}
			return storagelocaldisks
		})(item.GetDisks(), d)
		storagespangroup["object_type"] = item.GetObjectType()
		storagespangroups = append(storagespangroups, storagespangroup)
	}
	return storagespangroups
}
func flattenListStorageStoragePolicyRelationship(p []models.StorageStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagestoragepolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagestoragepolicyrelationship := flattenMoMoRef(item)
		storagestoragepolicyrelationships = append(storagestoragepolicyrelationships, storagestoragepolicyrelationship)
	}
	return storagestoragepolicyrelationships
}
func flattenListStorageVdMemberEpRelationship(p []models.StorageVdMemberEpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevdmembereprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevdmembereprelationship := flattenMoMoRef(item)
		storagevdmembereprelationships = append(storagevdmembereprelationships, storagevdmembereprelationship)
	}
	return storagevdmembereprelationships
}
func flattenListStorageVirtualDriveRelationship(p []models.StorageVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevirtualdriverelationship := flattenMoMoRef(item)
		storagevirtualdriverelationships = append(storagevirtualdriverelationships, storagevirtualdriverelationship)
	}
	return storagevirtualdriverelationships
}
func flattenListStorageVirtualDriveConfig(p []models.StorageVirtualDriveConfig, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveconfigs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagevirtualdriveconfig := make(map[string]interface{})
		storagevirtualdriveconfig["access_policy"] = item.GetAccessPolicy()
		storagevirtualdriveconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagevirtualdriveconfig["boot_drive"] = item.GetBootDrive()
		storagevirtualdriveconfig["class_id"] = item.GetClassId()
		storagevirtualdriveconfig["disk_group_name"] = item.GetDiskGroupName()
		storagevirtualdriveconfig["disk_group_policy"] = item.GetDiskGroupPolicy()
		storagevirtualdriveconfig["drive_cache"] = item.GetDriveCache()
		storagevirtualdriveconfig["expand_to_available"] = item.GetExpandToAvailable()
		storagevirtualdriveconfig["io_policy"] = item.GetIoPolicy()
		storagevirtualdriveconfig["name"] = item.GetName()
		storagevirtualdriveconfig["object_type"] = item.GetObjectType()
		storagevirtualdriveconfig["read_policy"] = item.GetReadPolicy()
		storagevirtualdriveconfig["size"] = item.GetSize()
		storagevirtualdriveconfig["strip_size"] = item.GetStripSize()
		storagevirtualdriveconfig["vdid"] = item.GetVdid()
		storagevirtualdriveconfig["write_policy"] = item.GetWritePolicy()
		storagevirtualdriveconfigs = append(storagevirtualdriveconfigs, storagevirtualdriveconfig)
	}
	return storagevirtualdriveconfigs
}
func flattenListStorageVirtualDriveContainerRelationship(p []models.StorageVirtualDriveContainerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdrivecontainerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevirtualdrivecontainerrelationship := flattenMoMoRef(item)
		storagevirtualdrivecontainerrelationships = append(storagevirtualdrivecontainerrelationships, storagevirtualdrivecontainerrelationship)
	}
	return storagevirtualdrivecontainerrelationships
}
func flattenListStorageVirtualDriveExtensionRelationship(p []models.StorageVirtualDriveExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevirtualdriveextensionrelationship := flattenMoMoRef(item)
		storagevirtualdriveextensionrelationships = append(storagevirtualdriveextensionrelationships, storagevirtualdriveextensionrelationship)
	}
	return storagevirtualdriveextensionrelationships
}
func flattenListSyslogLocalClientBase(p []models.SyslogLocalClientBase, d *schema.ResourceData) []map[string]interface{} {
	var sysloglocalclientbases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sysloglocalclientbase := make(map[string]interface{})
		sysloglocalclientbase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sysloglocalclientbase["class_id"] = item.GetClassId()
		sysloglocalclientbase["min_severity"] = item.GetMinSeverity()
		sysloglocalclientbase["object_type"] = item.GetObjectType()
		sysloglocalclientbases = append(sysloglocalclientbases, sysloglocalclientbase)
	}
	return sysloglocalclientbases
}
func flattenListSyslogRemoteClientBase(p []models.SyslogRemoteClientBase, d *schema.ResourceData) []map[string]interface{} {
	var syslogremoteclientbases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		syslogremoteclientbase := make(map[string]interface{})
		syslogremoteclientbase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		syslogremoteclientbase["class_id"] = item.GetClassId()
		syslogremoteclientbase["enabled"] = item.GetEnabled()
		syslogremoteclientbase["hostname"] = item.GetHostname()
		syslogremoteclientbase["min_severity"] = item.GetMinSeverity()
		syslogremoteclientbase["object_type"] = item.GetObjectType()
		syslogremoteclientbase["port"] = item.GetPort()
		syslogremoteclientbase["protocol"] = item.GetProtocol()
		syslogremoteclientbases = append(syslogremoteclientbases, syslogremoteclientbase)
	}
	return syslogremoteclientbases
}
func flattenListTamAction(p []models.TamAction, d *schema.ResourceData) []map[string]interface{} {
	var tamactions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		tamaction := make(map[string]interface{})
		tamaction["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		tamaction["affected_object_type"] = item.GetAffectedObjectType()
		tamaction["alert_type"] = item.GetAlertType()
		tamaction["class_id"] = item.GetClassId()
		tamaction["identifiers"] = (func(p []models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var tamidentifierss []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamidentifiers := make(map[string]interface{})
				tamidentifiers["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamidentifiers["class_id"] = item.GetClassId()
				tamidentifiers["name"] = item.GetName()
				tamidentifiers["object_type"] = item.GetObjectType()
				tamidentifiers["value"] = item.GetValue()
				tamidentifierss = append(tamidentifierss, tamidentifiers)
			}
			return tamidentifierss
		})(item.GetIdentifiers(), d)
		tamaction["name"] = item.GetName()
		tamaction["object_type"] = item.GetObjectType()
		tamaction["operation_type"] = item.GetOperationType()
		tamaction["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.GetClassId()
				tamqueryentry["name"] = item.GetName()
				tamqueryentry["object_type"] = item.GetObjectType()
				tamqueryentry["priority"] = item.GetPriority()
				tamqueryentry["query"] = item.GetQuery()
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tamaction["type"] = item.GetType()
		tamactions = append(tamactions, tamaction)
	}
	return tamactions
}
func flattenListTamApiDataSource(p []models.TamApiDataSource, d *schema.ResourceData) []map[string]interface{} {
	var tamapidatasources []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		tamapidatasource := make(map[string]interface{})
		tamapidatasource["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		tamapidatasource["class_id"] = item.GetClassId()
		tamapidatasource["mo_type"] = item.GetMoType()
		tamapidatasource["name"] = item.GetName()
		tamapidatasource["object_type"] = item.GetObjectType()
		tamapidatasource["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.GetClassId()
				tamqueryentry["name"] = item.GetName()
				tamqueryentry["object_type"] = item.GetObjectType()
				tamqueryentry["priority"] = item.GetPriority()
				tamqueryentry["query"] = item.GetQuery()
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tamapidatasource["type"] = item.GetType()
		tamapidatasources = append(tamapidatasources, tamapidatasource)
	}
	return tamapidatasources
}
func flattenListTamS3DataSource(p []models.TamS3DataSource, d *schema.ResourceData) []map[string]interface{} {
	var tams3datasources []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		tams3datasource := make(map[string]interface{})
		tams3datasource["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		tams3datasource["class_id"] = item.GetClassId()
		tams3datasource["name"] = item.GetName()
		tams3datasource["object_type"] = item.GetObjectType()
		tams3datasource["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.GetClassId()
				tamqueryentry["name"] = item.GetName()
				tamqueryentry["object_type"] = item.GetObjectType()
				tamqueryentry["priority"] = item.GetPriority()
				tamqueryentry["query"] = item.GetQuery()
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tams3datasource["s3_path"] = item.GetS3Path()
		tams3datasource["type"] = item.GetType()
		tams3datasources = append(tams3datasources, tams3datasource)
	}
	return tams3datasources
}
func flattenListUcsdConnectorPack(p []models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var ucsdconnectorpacks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ucsdconnectorpack := make(map[string]interface{})
		ucsdconnectorpack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ucsdconnectorpack["class_id"] = item.GetClassId()
		ucsdconnectorpack["connector_feature"] = item.GetConnectorFeature()
		ucsdconnectorpack["dependency_names"] = item.GetDependencyNames()
		ucsdconnectorpack["downloaded_version"] = item.GetDownloadedVersion()
		ucsdconnectorpack["name"] = item.GetName()
		ucsdconnectorpack["object_type"] = item.GetObjectType()
		ucsdconnectorpack["services"] = item.GetServices()
		ucsdconnectorpack["state"] = item.GetState()
		ucsdconnectorpack["nr_version"] = item.GetVersion()
		ucsdconnectorpacks = append(ucsdconnectorpacks, ucsdconnectorpack)
	}
	return ucsdconnectorpacks
}
func flattenListUuidpoolBlockRelationship(p []models.UuidpoolBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		uuidpoolblockrelationship := flattenMoMoRef(item)
		uuidpoolblockrelationships = append(uuidpoolblockrelationships, uuidpoolblockrelationship)
	}
	return uuidpoolblockrelationships
}
func flattenListUuidpoolUuidBlock(p []models.UuidpoolUuidBlock, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluuidblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		uuidpooluuidblock := make(map[string]interface{})
		uuidpooluuidblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		uuidpooluuidblock["class_id"] = item.GetClassId()
		uuidpooluuidblock["from"] = item.GetFrom()
		uuidpooluuidblock["object_type"] = item.GetObjectType()
		uuidpooluuidblock["size"] = item.GetSize()
		uuidpooluuidblock["to"] = item.GetTo()
		uuidpooluuidblocks = append(uuidpooluuidblocks, uuidpooluuidblock)
	}
	return uuidpooluuidblocks
}
func flattenListVirtualizationBaseNetworkRelationship(p []models.VirtualizationBaseNetworkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbasenetworkrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationbasenetworkrelationship := flattenMoMoRef(item)
		virtualizationbasenetworkrelationships = append(virtualizationbasenetworkrelationships, virtualizationbasenetworkrelationship)
	}
	return virtualizationbasenetworkrelationships
}
func flattenListVirtualizationNetworkInterface(p []models.VirtualizationNetworkInterface, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationnetworkinterfaces []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		virtualizationnetworkinterface := make(map[string]interface{})
		virtualizationnetworkinterface["adaptor_type"] = item.GetAdaptorType()
		virtualizationnetworkinterface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		virtualizationnetworkinterface["bridge"] = item.GetBridge()
		virtualizationnetworkinterface["class_id"] = item.GetClassId()
		virtualizationnetworkinterface["connect_at_power_on"] = item.GetConnectAtPowerOn()
		virtualizationnetworkinterface["direct_path_io"] = item.GetDirectPathIo()
		virtualizationnetworkinterface["mac_address"] = item.GetMacAddress()
		virtualizationnetworkinterface["name"] = item.GetName()
		virtualizationnetworkinterface["object_type"] = item.GetObjectType()
		virtualizationnetworkinterfaces = append(virtualizationnetworkinterfaces, virtualizationnetworkinterface)
	}
	return virtualizationnetworkinterfaces
}
func flattenListVirtualizationVirtualMachineDisk(p []models.VirtualizationVirtualMachineDisk, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvirtualmachinedisks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		virtualizationvirtualmachinedisk := make(map[string]interface{})
		virtualizationvirtualmachinedisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		virtualizationvirtualmachinedisk["bus"] = item.GetBus()
		virtualizationvirtualmachinedisk["class_id"] = item.GetClassId()
		virtualizationvirtualmachinedisk["name"] = item.GetName()
		virtualizationvirtualmachinedisk["object_type"] = item.GetObjectType()
		virtualizationvirtualmachinedisk["order"] = item.GetOrder()
		virtualizationvirtualmachinedisk["type"] = item.GetType()
		virtualizationvirtualmachinedisk["virtual_disk"] = (func(p models.VirtualizationVirtualDiskConfig, d *schema.ResourceData) []map[string]interface{} {
			var virtualizationvirtualdiskconfigs []map[string]interface{}
			var ret models.VirtualizationVirtualDiskConfig
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			virtualizationvirtualdiskconfig := make(map[string]interface{})
			virtualizationvirtualdiskconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			virtualizationvirtualdiskconfig["capacity"] = item.GetCapacity()
			virtualizationvirtualdiskconfig["class_id"] = item.GetClassId()
			virtualizationvirtualdiskconfig["mode"] = item.GetMode()
			virtualizationvirtualdiskconfig["object_type"] = item.GetObjectType()
			virtualizationvirtualdiskconfig["source_certs"] = item.GetSourceCerts()
			virtualizationvirtualdiskconfig["source_disk_to_clone"] = item.GetSourceDiskToClone()
			virtualizationvirtualdiskconfig["source_file_path"] = item.GetSourceFilePath()

			virtualizationvirtualdiskconfigs = append(virtualizationvirtualdiskconfigs, virtualizationvirtualdiskconfig)
			return virtualizationvirtualdiskconfigs
		})(item.GetVirtualDisk(), d)
		virtualizationvirtualmachinedisk["virtual_disk_reference"] = item.GetVirtualDiskReference()
		virtualizationvirtualmachinedisks = append(virtualizationvirtualmachinedisks, virtualizationvirtualmachinedisk)
	}
	return virtualizationvirtualmachinedisks
}
func flattenListVirtualizationVmwareDatastoreRelationship(p []models.VirtualizationVmwareDatastoreRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatastorerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationvmwaredatastorerelationship := flattenMoMoRef(item)
		virtualizationvmwaredatastorerelationships = append(virtualizationvmwaredatastorerelationships, virtualizationvmwaredatastorerelationship)
	}
	return virtualizationvmwaredatastorerelationships
}
func flattenListVirtualizationVmwareDistributedNetworkRelationship(p []models.VirtualizationVmwareDistributedNetworkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredistributednetworkrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationvmwaredistributednetworkrelationship := flattenMoMoRef(item)
		virtualizationvmwaredistributednetworkrelationships = append(virtualizationvmwaredistributednetworkrelationships, virtualizationvmwaredistributednetworkrelationship)
	}
	return virtualizationvmwaredistributednetworkrelationships
}
func flattenListVirtualizationVmwareDistributedSwitchRelationship(p []models.VirtualizationVmwareDistributedSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredistributedswitchrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationvmwaredistributedswitchrelationship := flattenMoMoRef(item)
		virtualizationvmwaredistributedswitchrelationships = append(virtualizationvmwaredistributedswitchrelationships, virtualizationvmwaredistributedswitchrelationship)
	}
	return virtualizationvmwaredistributedswitchrelationships
}
func flattenListVirtualizationVmwareHostRelationship(p []models.VirtualizationVmwareHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationvmwarehostrelationship := flattenMoMoRef(item)
		virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	}
	return virtualizationvmwarehostrelationships
}
func flattenListVirtualizationVmwareVlanRange(p []models.VirtualizationVmwareVlanRange, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevlanranges []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		virtualizationvmwarevlanrange := make(map[string]interface{})
		virtualizationvmwarevlanrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		virtualizationvmwarevlanrange["class_id"] = item.GetClassId()
		virtualizationvmwarevlanrange["object_type"] = item.GetObjectType()
		virtualizationvmwarevlanrange["vlan_range_end"] = item.GetVlanRangeEnd()
		virtualizationvmwarevlanrange["vlan_range_start"] = item.GetVlanRangeStart()
		virtualizationvmwarevlanranges = append(virtualizationvmwarevlanranges, virtualizationvmwarevlanrange)
	}
	return virtualizationvmwarevlanranges
}
func flattenListVmediaMapping(p []models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var vmediamappings []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		vmediamapping := make(map[string]interface{})
		vmediamapping["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		vmediamapping["authentication_protocol"] = item.GetAuthenticationProtocol()
		vmediamapping["class_id"] = item.GetClassId()
		vmediamapping["device_type"] = item.GetDeviceType()
		vmediamapping["file_location"] = item.GetFileLocation()
		vmediamapping["host_name"] = item.GetHostName()
		vmediamapping["is_password_set"] = item.GetIsPasswordSet()
		vmediamapping["mount_options"] = item.GetMountOptions()
		vmediamapping["mount_protocol"] = item.GetMountProtocol()
		vmediamapping["object_type"] = item.GetObjectType()
		password_x, exists := d.GetOk("mappings")
		if exists && password_x != nil {
			vmediamapping["password"] = password_x.([]interface{})[len(vmediamappings)].(map[string]interface{})["password"]
		}
		vmediamapping["remote_file"] = item.GetRemoteFile()
		vmediamapping["remote_path"] = item.GetRemotePath()
		vmediamapping["sanitized_file_location"] = item.GetSanitizedFileLocation()
		vmediamapping["username"] = item.GetUsername()
		vmediamapping["volume_name"] = item.GetVolumeName()
		vmediamappings = append(vmediamappings, vmediamapping)
	}
	return vmediamappings
}
func flattenListVnicEthIfRelationship(p []models.VnicEthIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethifrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		vnicethifrelationship := flattenMoMoRef(item)
		vnicethifrelationships = append(vnicethifrelationships, vnicethifrelationship)
	}
	return vnicethifrelationships
}
func flattenListVnicEthNetworkPolicyRelationship(p []models.VnicEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethnetworkpolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		vnicethnetworkpolicyrelationship := flattenMoMoRef(item)
		vnicethnetworkpolicyrelationships = append(vnicethnetworkpolicyrelationships, vnicethnetworkpolicyrelationship)
	}
	return vnicethnetworkpolicyrelationships
}
func flattenListVnicFcIfRelationship(p []models.VnicFcIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcifrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		vnicfcifrelationship := flattenMoMoRef(item)
		vnicfcifrelationships = append(vnicfcifrelationships, vnicfcifrelationship)
	}
	return vnicfcifrelationships
}
func flattenListVnicVifStatus(p []models.VnicVifStatus, d *schema.ResourceData) []map[string]interface{} {
	var vnicvifstatuss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		vnicvifstatus := make(map[string]interface{})
		vnicvifstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		vnicvifstatus["class_id"] = item.GetClassId()
		vnicvifstatus["name"] = item.GetName()
		vnicvifstatus["object_type"] = item.GetObjectType()
		vnicvifstatus["reason"] = item.GetReason()
		vnicvifstatus["status"] = item.GetStatus()
		vnicvifstatuss = append(vnicvifstatuss, vnicvifstatus)
	}
	return vnicvifstatuss
}
func flattenListWorkflowApi(p []models.WorkflowApi, d *schema.ResourceData) []map[string]interface{} {
	var workflowapis []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowapi := make(map[string]interface{})
		workflowapi["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowapi["asset_target_moid"] = item.GetAssetTargetMoid()
		workflowapi["body"] = item.GetBody()
		workflowapi["class_id"] = item.GetClassId()
		workflowapi["content_type"] = item.GetContentType()
		workflowapi["description"] = item.GetDescription()
		workflowapi["error_content_type"] = item.GetErrorContentType()
		workflowapi["label"] = item.GetLabel()
		workflowapi["name"] = item.GetName()
		workflowapi["object_type"] = item.GetObjectType()
		workflowapi["outcomes"] = flattenAdditionalProperties(item.Outcomes)
		workflowapi["response_spec"] = flattenAdditionalProperties(item.ResponseSpec)
		workflowapi["skip_on_condition"] = item.GetSkipOnCondition()
		workflowapi["start_delay"] = item.GetStartDelay()
		workflowapi["timeout"] = item.GetTimeout()
		workflowapis = append(workflowapis, workflowapi)
	}
	return workflowapis
}
func flattenListWorkflowBaseDataType(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
	var workflowbasedatatypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowbasedatatype := make(map[string]interface{})
		workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowbasedatatype["class_id"] = item.GetClassId()
		workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
			var workflowdefaultvalues []map[string]interface{}
			var ret models.WorkflowDefaultValue
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowdefaultvalue := make(map[string]interface{})
			workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowdefaultvalue["class_id"] = item.GetClassId()
			workflowdefaultvalue["is_value_set"] = item.GetIsValueSet()
			workflowdefaultvalue["object_type"] = item.GetObjectType()
			workflowdefaultvalue["override"] = item.GetOverride()
			workflowdefaultvalue["value"] = flattenAdditionalProperties(item.Value)

			workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
			return workflowdefaultvalues
		})(item.GetDefault(), d)
		workflowbasedatatype["description"] = item.GetDescription()
		workflowbasedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
			var workflowdisplaymetas []map[string]interface{}
			var ret models.WorkflowDisplayMeta
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowdisplaymeta := make(map[string]interface{})
			workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowdisplaymeta["class_id"] = item.GetClassId()
			workflowdisplaymeta["inventory_selector"] = item.GetInventorySelector()
			workflowdisplaymeta["object_type"] = item.GetObjectType()
			workflowdisplaymeta["widget_type"] = item.GetWidgetType()

			workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
			return workflowdisplaymetas
		})(item.GetDisplayMeta(), d)
		workflowbasedatatype["input_parameters"] = flattenAdditionalProperties(item.InputParameters)
		workflowbasedatatype["label"] = item.GetLabel()
		workflowbasedatatype["name"] = item.GetName()
		workflowbasedatatype["object_type"] = item.GetObjectType()
		workflowbasedatatype["required"] = item.GetRequired()
		workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
	}
	return workflowbasedatatypes
}
func flattenListWorkflowDynamicWorkflowActionTaskList(p []models.WorkflowDynamicWorkflowActionTaskList, d *schema.ResourceData) []map[string]interface{} {
	var workflowdynamicworkflowactiontasklists []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowdynamicworkflowactiontasklist := make(map[string]interface{})
		workflowdynamicworkflowactiontasklist["action"] = item.GetAction()
		workflowdynamicworkflowactiontasklist["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowdynamicworkflowactiontasklist["class_id"] = item.GetClassId()
		workflowdynamicworkflowactiontasklist["object_type"] = item.GetObjectType()
		workflowdynamicworkflowactiontasklist["tasks"] = flattenAdditionalProperties(item.Tasks)
		workflowdynamicworkflowactiontasklists = append(workflowdynamicworkflowactiontasklists, workflowdynamicworkflowactiontasklist)
	}
	return workflowdynamicworkflowactiontasklists
}
func flattenListWorkflowMessage(p []models.WorkflowMessage, d *schema.ResourceData) []map[string]interface{} {
	var workflowmessages []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowmessage := make(map[string]interface{})
		workflowmessage["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowmessage["class_id"] = item.GetClassId()
		workflowmessage["message"] = item.GetMessage()
		workflowmessage["object_type"] = item.GetObjectType()
		workflowmessage["severity"] = item.GetSeverity()
		workflowmessages = append(workflowmessages, workflowmessage)
	}
	return workflowmessages
}
func flattenListWorkflowParameterSet(p []models.WorkflowParameterSet, d *schema.ResourceData) []map[string]interface{} {
	var workflowparametersets []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowparameterset := make(map[string]interface{})
		workflowparameterset["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowparameterset["class_id"] = item.GetClassId()
		workflowparameterset["condition"] = item.GetCondition()
		workflowparameterset["control_parameter"] = item.GetControlParameter()
		workflowparameterset["enable_parameters"] = item.GetEnableParameters()
		workflowparameterset["name"] = item.GetName()
		workflowparameterset["object_type"] = item.GetObjectType()
		workflowparameterset["value"] = item.GetValue()
		workflowparametersets = append(workflowparametersets, workflowparameterset)
	}
	return workflowparametersets
}
func flattenListWorkflowRollbackTask(p []models.WorkflowRollbackTask, d *schema.ResourceData) []map[string]interface{} {
	var workflowrollbacktasks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowrollbacktask := make(map[string]interface{})
		workflowrollbacktask["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowrollbacktask["catalog_moid"] = item.GetCatalogMoid()
		workflowrollbacktask["class_id"] = item.GetClassId()
		workflowrollbacktask["description"] = item.GetDescription()
		workflowrollbacktask["input_parameters"] = flattenAdditionalProperties(item.InputParameters)
		workflowrollbacktask["name"] = item.GetName()
		workflowrollbacktask["object_type"] = item.GetObjectType()
		workflowrollbacktask["task_moid"] = item.GetTaskMoid()
		workflowrollbacktask["nr_version"] = item.GetVersion()
		workflowrollbacktasks = append(workflowrollbacktasks, workflowrollbacktask)
	}
	return workflowrollbacktasks
}
func flattenListWorkflowRollbackWorkflowTask(p []models.WorkflowRollbackWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var workflowrollbackworkflowtasks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowrollbackworkflowtask := make(map[string]interface{})
		workflowrollbackworkflowtask["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowrollbackworkflowtask["class_id"] = item.GetClassId()
		workflowrollbackworkflowtask["description"] = item.GetDescription()
		workflowrollbackworkflowtask["name"] = item.GetName()
		workflowrollbackworkflowtask["object_type"] = item.GetObjectType()
		workflowrollbackworkflowtask["ref_name"] = item.GetRefName()
		workflowrollbackworkflowtask["rollback_completed"] = item.GetRollbackCompleted()
		workflowrollbackworkflowtask["rollback_task_name"] = item.GetRollbackTaskName()
		workflowrollbackworkflowtask["status"] = item.GetStatus()
		workflowrollbackworkflowtask["task_info_moid"] = item.GetTaskInfoMoid()
		workflowrollbackworkflowtask["task_path"] = item.GetTaskPath()
		workflowrollbackworkflowtasks = append(workflowrollbackworkflowtasks, workflowrollbackworkflowtask)
	}
	return workflowrollbackworkflowtasks
}
func flattenListWorkflowTaskDefinitionRelationship(p []models.WorkflowTaskDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		workflowtaskdefinitionrelationship := flattenMoMoRef(item)
		workflowtaskdefinitionrelationships = append(workflowtaskdefinitionrelationships, workflowtaskdefinitionrelationship)
	}
	return workflowtaskdefinitionrelationships
}
func flattenListWorkflowTaskInfoRelationship(p []models.WorkflowTaskInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		workflowtaskinforelationship := flattenMoMoRef(item)
		workflowtaskinforelationships = append(workflowtaskinforelationships, workflowtaskinforelationship)
	}
	return workflowtaskinforelationships
}
func flattenListWorkflowTaskRetryInfo(p []models.WorkflowTaskRetryInfo, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskretryinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowtaskretryinfo := make(map[string]interface{})
		workflowtaskretryinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowtaskretryinfo["class_id"] = item.GetClassId()
		workflowtaskretryinfo["object_type"] = item.GetObjectType()
		workflowtaskretryinfo["status"] = item.GetStatus()
		workflowtaskretryinfo["task_inst_id"] = item.GetTaskInstId()
		workflowtaskretryinfos = append(workflowtaskretryinfos, workflowtaskretryinfo)
	}
	return workflowtaskretryinfos
}
func flattenListWorkflowUiInputFilter(p []models.WorkflowUiInputFilter, d *schema.ResourceData) []map[string]interface{} {
	var workflowuiinputfilters []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowuiinputfilter := make(map[string]interface{})
		workflowuiinputfilter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowuiinputfilter["class_id"] = item.GetClassId()
		workflowuiinputfilter["filters"] = item.GetFilters()
		workflowuiinputfilter["name"] = item.GetName()
		workflowuiinputfilter["object_type"] = item.GetObjectType()
		workflowuiinputfilter["user_help_message"] = item.GetUserHelpMessage()
		workflowuiinputfilters = append(workflowuiinputfilters, workflowuiinputfilter)
	}
	return workflowuiinputfilters
}
func flattenListWorkflowWorkflowInfoRelationship(p []models.WorkflowWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		workflowworkflowinforelationship := flattenMoMoRef(item)
		workflowworkflowinforelationships = append(workflowworkflowinforelationships, workflowworkflowinforelationship)
	}
	return workflowworkflowinforelationships
}
func flattenListWorkflowWorkflowTask(p []models.WorkflowWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowtasks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowworkflowtask := make(map[string]interface{})
		workflowworkflowtask["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowworkflowtask["class_id"] = item.GetClassId()
		workflowworkflowtask["description"] = item.GetDescription()
		workflowworkflowtask["label"] = item.GetLabel()
		workflowworkflowtask["name"] = item.GetName()
		workflowworkflowtask["object_type"] = item.GetObjectType()
		workflowworkflowtasks = append(workflowworkflowtasks, workflowworkflowtask)
	}
	return workflowworkflowtasks
}
func flattenListX509Certificate(p []models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		x509certificate := make(map[string]interface{})
		x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		x509certificate["class_id"] = item.GetClassId()
		x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.GetClassId()
			pkixdistinguishedname["common_name"] = item.GetCommonName()
			pkixdistinguishedname["country"] = item.GetCountry()
			pkixdistinguishedname["locality"] = item.GetLocality()
			pkixdistinguishedname["object_type"] = item.GetObjectType()
			pkixdistinguishedname["organization"] = item.GetOrganization()
			pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
			pkixdistinguishedname["state"] = item.GetState()

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetIssuer(), d)
		x509certificate["not_after"] = item.GetNotAfter()
		x509certificate["not_before"] = item.GetNotBefore()
		x509certificate["object_type"] = item.GetObjectType()
		x509certificate["pem_certificate"] = item.GetPemCertificate()
		x509certificate["sha256_fingerprint"] = item.GetSha256Fingerprint()
		x509certificate["signature_algorithm"] = item.GetSignatureAlgorithm()
		x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.GetClassId()
			pkixdistinguishedname["common_name"] = item.GetCommonName()
			pkixdistinguishedname["country"] = item.GetCountry()
			pkixdistinguishedname["locality"] = item.GetLocality()
			pkixdistinguishedname["object_type"] = item.GetObjectType()
			pkixdistinguishedname["organization"] = item.GetOrganization()
			pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
			pkixdistinguishedname["state"] = item.GetState()

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetSubject(), d)
		x509certificates = append(x509certificates, x509certificate)
	}
	return x509certificates
}
func flattenMapAccessAddressType(p models.AccessAddressType, d *schema.ResourceData) []map[string]interface{} {
	var accessaddresstypes []map[string]interface{}
	var ret models.AccessAddressType
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	accessaddresstype := make(map[string]interface{})
	accessaddresstype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	accessaddresstype["class_id"] = item.GetClassId()
	accessaddresstype["enable_ip_v4"] = item.GetEnableIpV4()
	accessaddresstype["enable_ip_v6"] = item.GetEnableIpV6()
	accessaddresstype["object_type"] = item.GetObjectType()

	accessaddresstypes = append(accessaddresstypes, accessaddresstype)
	return accessaddresstypes
}
func flattenMapAdapterUnitRelationship(p models.AdapterUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrelationships []map[string]interface{}
	var ret models.AdapterUnitRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	adapterunitrelationship := make(map[string]interface{})
	adapterunitrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	adapterunitrelationship["class_id"] = item.GetClassId()
	adapterunitrelationship["moid"] = item.GetMoid()
	adapterunitrelationship["object_type"] = item.GetObjectType()
	adapterunitrelationship["selector"] = item.GetSelector()

	adapterunitrelationships = append(adapterunitrelationships, adapterunitrelationship)
	return adapterunitrelationships
}
func flattenMapAdapterUnitExpanderRelationship(p models.AdapterUnitExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitexpanderrelationships []map[string]interface{}
	var ret models.AdapterUnitExpanderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	adapterunitexpanderrelationship := make(map[string]interface{})
	adapterunitexpanderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	adapterunitexpanderrelationship["class_id"] = item.GetClassId()
	adapterunitexpanderrelationship["moid"] = item.GetMoid()
	adapterunitexpanderrelationship["object_type"] = item.GetObjectType()
	adapterunitexpanderrelationship["selector"] = item.GetSelector()

	adapterunitexpanderrelationships = append(adapterunitexpanderrelationships, adapterunitexpanderrelationship)
	return adapterunitexpanderrelationships
}
func flattenMapApplianceCertRenewalPhase(p models.ApplianceCertRenewalPhase, d *schema.ResourceData) []map[string]interface{} {
	var appliancecertrenewalphases []map[string]interface{}
	var ret models.ApplianceCertRenewalPhase
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	appliancecertrenewalphase := make(map[string]interface{})
	appliancecertrenewalphase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancecertrenewalphase["class_id"] = item.GetClassId()
	appliancecertrenewalphase["end_time"] = item.GetEndTime()
	appliancecertrenewalphase["failed"] = item.GetFailed()
	appliancecertrenewalphase["message"] = item.GetMessage()
	appliancecertrenewalphase["name"] = item.GetName()
	appliancecertrenewalphase["object_type"] = item.GetObjectType()
	appliancecertrenewalphase["start_time"] = item.GetStartTime()

	appliancecertrenewalphases = append(appliancecertrenewalphases, appliancecertrenewalphase)
	return appliancecertrenewalphases
}
func flattenMapApplianceDataExportPolicyRelationship(p models.ApplianceDataExportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrelationships []map[string]interface{}
	var ret models.ApplianceDataExportPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancedataexportpolicyrelationship := make(map[string]interface{})
	appliancedataexportpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancedataexportpolicyrelationship["class_id"] = item.GetClassId()
	appliancedataexportpolicyrelationship["moid"] = item.GetMoid()
	appliancedataexportpolicyrelationship["object_type"] = item.GetObjectType()
	appliancedataexportpolicyrelationship["selector"] = item.GetSelector()

	appliancedataexportpolicyrelationships = append(appliancedataexportpolicyrelationships, appliancedataexportpolicyrelationship)
	return appliancedataexportpolicyrelationships
}
func flattenMapApplianceGroupStatusRelationship(p models.ApplianceGroupStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancegroupstatusrelationships []map[string]interface{}
	var ret models.ApplianceGroupStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancegroupstatusrelationship := make(map[string]interface{})
	appliancegroupstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancegroupstatusrelationship["class_id"] = item.GetClassId()
	appliancegroupstatusrelationship["moid"] = item.GetMoid()
	appliancegroupstatusrelationship["object_type"] = item.GetObjectType()
	appliancegroupstatusrelationship["selector"] = item.GetSelector()

	appliancegroupstatusrelationships = append(appliancegroupstatusrelationships, appliancegroupstatusrelationship)
	return appliancegroupstatusrelationships
}
func flattenMapApplianceImageBundleRelationship(p models.ApplianceImageBundleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var applianceimagebundlerelationships []map[string]interface{}
	var ret models.ApplianceImageBundleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	applianceimagebundlerelationship := make(map[string]interface{})
	applianceimagebundlerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	applianceimagebundlerelationship["class_id"] = item.GetClassId()
	applianceimagebundlerelationship["moid"] = item.GetMoid()
	applianceimagebundlerelationship["object_type"] = item.GetObjectType()
	applianceimagebundlerelationship["selector"] = item.GetSelector()

	applianceimagebundlerelationships = append(applianceimagebundlerelationships, applianceimagebundlerelationship)
	return applianceimagebundlerelationships
}
func flattenMapApplianceNodeInfoRelationship(p models.ApplianceNodeInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancenodeinforelationships []map[string]interface{}
	var ret models.ApplianceNodeInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancenodeinforelationship := make(map[string]interface{})
	appliancenodeinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancenodeinforelationship["class_id"] = item.GetClassId()
	appliancenodeinforelationship["moid"] = item.GetMoid()
	appliancenodeinforelationship["object_type"] = item.GetObjectType()
	appliancenodeinforelationship["selector"] = item.GetSelector()

	appliancenodeinforelationships = append(appliancenodeinforelationships, appliancenodeinforelationship)
	return appliancenodeinforelationships
}
func flattenMapApplianceNodeStatusRelationship(p models.ApplianceNodeStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancenodestatusrelationships []map[string]interface{}
	var ret models.ApplianceNodeStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancenodestatusrelationship := make(map[string]interface{})
	appliancenodestatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancenodestatusrelationship["class_id"] = item.GetClassId()
	appliancenodestatusrelationship["moid"] = item.GetMoid()
	appliancenodestatusrelationship["object_type"] = item.GetObjectType()
	appliancenodestatusrelationship["selector"] = item.GetSelector()

	appliancenodestatusrelationships = append(appliancenodestatusrelationships, appliancenodestatusrelationship)
	return appliancenodestatusrelationships
}
func flattenMapApplianceSystemInfoRelationship(p models.ApplianceSystemInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancesysteminforelationships []map[string]interface{}
	var ret models.ApplianceSystemInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancesysteminforelationship := make(map[string]interface{})
	appliancesysteminforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancesysteminforelationship["class_id"] = item.GetClassId()
	appliancesysteminforelationship["moid"] = item.GetMoid()
	appliancesysteminforelationship["object_type"] = item.GetObjectType()
	appliancesysteminforelationship["selector"] = item.GetSelector()

	appliancesysteminforelationships = append(appliancesysteminforelationships, appliancesysteminforelationship)
	return appliancesysteminforelationships
}
func flattenMapApplianceSystemStatusRelationship(p models.ApplianceSystemStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancesystemstatusrelationships []map[string]interface{}
	var ret models.ApplianceSystemStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancesystemstatusrelationship := make(map[string]interface{})
	appliancesystemstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancesystemstatusrelationship["class_id"] = item.GetClassId()
	appliancesystemstatusrelationship["moid"] = item.GetMoid()
	appliancesystemstatusrelationship["object_type"] = item.GetObjectType()
	appliancesystemstatusrelationship["selector"] = item.GetSelector()

	appliancesystemstatusrelationships = append(appliancesystemstatusrelationships, appliancesystemstatusrelationship)
	return appliancesystemstatusrelationships
}
func flattenMapAssetClaimSignature(p models.AssetClaimSignature, d *schema.ResourceData) []map[string]interface{} {
	var assetclaimsignatures []map[string]interface{}
	var ret models.AssetClaimSignature
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetclaimsignature := make(map[string]interface{})
	assetclaimsignature["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetclaimsignature["class_id"] = item.GetClassId()
	assetclaimsignature["object_type"] = item.GetObjectType()
	assetclaimsignature["time_stamp"] = item.GetTimeStamp()

	assetclaimsignatures = append(assetclaimsignatures, assetclaimsignature)
	return assetclaimsignatures
}
func flattenMapAssetClusterMemberRelationship(p models.AssetClusterMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrelationships []map[string]interface{}
	var ret models.AssetClusterMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetclustermemberrelationship := make(map[string]interface{})
	assetclustermemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetclustermemberrelationship["class_id"] = item.GetClassId()
	assetclustermemberrelationship["moid"] = item.GetMoid()
	assetclustermemberrelationship["object_type"] = item.GetObjectType()
	assetclustermemberrelationship["selector"] = item.GetSelector()

	assetclustermemberrelationships = append(assetclustermemberrelationships, assetclustermemberrelationship)
	return assetclustermemberrelationships
}
func flattenMapAssetContractInformation(p models.AssetContractInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcontractinformations []map[string]interface{}
	var ret models.AssetContractInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetcontractinformation := make(map[string]interface{})
	assetcontractinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetcontractinformation["bill_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.GetAddress1()
		assetaddressinformation["address2"] = item.GetAddress2()
		assetaddressinformation["address3"] = item.GetAddress3()
		assetaddressinformation["city"] = item.GetCity()
		assetaddressinformation["class_id"] = item.GetClassId()
		assetaddressinformation["country"] = item.GetCountry()
		assetaddressinformation["county"] = item.GetCounty()
		assetaddressinformation["location"] = item.GetLocation()
		assetaddressinformation["name"] = item.GetName()
		assetaddressinformation["object_type"] = item.GetObjectType()
		assetaddressinformation["postal_code"] = item.GetPostalCode()
		assetaddressinformation["province"] = item.GetProvince()
		assetaddressinformation["state"] = item.GetState()

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetBillTo(), d)
	assetcontractinformation["bill_to_global_ultimate"] = (func(p models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
		var assetglobalultimates []map[string]interface{}
		var ret models.AssetGlobalUltimate
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetglobalultimate := make(map[string]interface{})
		assetglobalultimate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetglobalultimate["class_id"] = item.GetClassId()
		assetglobalultimate["id"] = item.GetId()
		assetglobalultimate["name"] = item.GetName()
		assetglobalultimate["object_type"] = item.GetObjectType()

		assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
		return assetglobalultimates
	})(item.GetBillToGlobalUltimate(), d)
	assetcontractinformation["class_id"] = item.GetClassId()
	assetcontractinformation["contract_number"] = item.GetContractNumber()
	assetcontractinformation["line_status"] = item.GetLineStatus()
	assetcontractinformation["object_type"] = item.GetObjectType()

	assetcontractinformations = append(assetcontractinformations, assetcontractinformation)
	return assetcontractinformations
}
func flattenMapAssetCustomerInformation(p models.AssetCustomerInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcustomerinformations []map[string]interface{}
	var ret models.AssetCustomerInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetcustomerinformation := make(map[string]interface{})
	assetcustomerinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetcustomerinformation["address"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.GetAddress1()
		assetaddressinformation["address2"] = item.GetAddress2()
		assetaddressinformation["address3"] = item.GetAddress3()
		assetaddressinformation["city"] = item.GetCity()
		assetaddressinformation["class_id"] = item.GetClassId()
		assetaddressinformation["country"] = item.GetCountry()
		assetaddressinformation["county"] = item.GetCounty()
		assetaddressinformation["location"] = item.GetLocation()
		assetaddressinformation["name"] = item.GetName()
		assetaddressinformation["object_type"] = item.GetObjectType()
		assetaddressinformation["postal_code"] = item.GetPostalCode()
		assetaddressinformation["province"] = item.GetProvince()
		assetaddressinformation["state"] = item.GetState()

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetAddress(), d)
	assetcustomerinformation["class_id"] = item.GetClassId()
	assetcustomerinformation["id"] = item.GetId()
	assetcustomerinformation["name"] = item.GetName()
	assetcustomerinformation["object_type"] = item.GetObjectType()

	assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
	return assetcustomerinformations
}
func flattenMapAssetDeploymentRelationship(p models.AssetDeploymentRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeploymentrelationships []map[string]interface{}
	var ret models.AssetDeploymentRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeploymentrelationship := make(map[string]interface{})
	assetdeploymentrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeploymentrelationship["class_id"] = item.GetClassId()
	assetdeploymentrelationship["moid"] = item.GetMoid()
	assetdeploymentrelationship["object_type"] = item.GetObjectType()
	assetdeploymentrelationship["selector"] = item.GetSelector()

	assetdeploymentrelationships = append(assetdeploymentrelationships, assetdeploymentrelationship)
	return assetdeploymentrelationships
}
func flattenMapAssetDeploymentDeviceInformation(p models.AssetDeploymentDeviceInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetdeploymentdeviceinformations []map[string]interface{}
	var ret models.AssetDeploymentDeviceInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetdeploymentdeviceinformation := make(map[string]interface{})
	assetdeploymentdeviceinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeploymentdeviceinformation["application_name"] = item.GetApplicationName()
	assetdeploymentdeviceinformation["class_id"] = item.GetClassId()
	assetdeploymentdeviceinformation["description"] = item.GetDescription()
	assetdeploymentdeviceinformation["device_transactions"] = (func(p []models.AssetDeviceTransaction, d *schema.ResourceData) []map[string]interface{} {
		var assetdevicetransactions []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			assetdevicetransaction := make(map[string]interface{})
			assetdevicetransaction["action"] = item.GetAction()
			assetdevicetransaction["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetdevicetransaction["class_id"] = item.GetClassId()
			assetdevicetransaction["object_type"] = item.GetObjectType()
			assetdevicetransaction["status_description"] = item.GetStatusDescription()
			assetdevicetransaction["status_id"] = item.GetStatusId()
			assetdevicetransaction["timestamp"] = item.GetTimestamp()
			assetdevicetransaction["transaction_batch_id"] = item.GetTransactionBatchId()
			assetdevicetransaction["transaction_date"] = item.GetTransactionDate()
			assetdevicetransaction["transaction_sequence"] = item.GetTransactionSequence()
			assetdevicetransactions = append(assetdevicetransactions, assetdevicetransaction)
		}
		return assetdevicetransactions
	})(item.GetDeviceTransactions(), d)
	assetdeploymentdeviceinformation["instance_id"] = item.GetInstanceId()
	assetdeploymentdeviceinformation["item_type"] = item.GetItemType()
	assetdeploymentdeviceinformation["mlb_product_id"] = item.GetMlbProductId()
	assetdeploymentdeviceinformation["mlb_product_name"] = item.GetMlbProductName()
	assetdeploymentdeviceinformation["object_type"] = item.GetObjectType()
	assetdeploymentdeviceinformation["old_device_id"] = item.GetOldDeviceId()
	assetdeploymentdeviceinformation["old_device_status_description"] = item.GetOldDeviceStatusDescription()
	assetdeploymentdeviceinformation["old_device_status_id"] = item.GetOldDeviceStatusId()
	assetdeploymentdeviceinformation["old_instance_id"] = item.GetOldInstanceId()

	assetdeploymentdeviceinformations = append(assetdeploymentdeviceinformations, assetdeploymentdeviceinformation)
	return assetdeploymentdeviceinformations
}
func flattenMapAssetDeviceClaimRelationship(p models.AssetDeviceClaimRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceclaimrelationships []map[string]interface{}
	var ret models.AssetDeviceClaimRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceclaimrelationship := make(map[string]interface{})
	assetdeviceclaimrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceclaimrelationship["class_id"] = item.GetClassId()
	assetdeviceclaimrelationship["moid"] = item.GetMoid()
	assetdeviceclaimrelationship["object_type"] = item.GetObjectType()
	assetdeviceclaimrelationship["selector"] = item.GetSelector()

	assetdeviceclaimrelationships = append(assetdeviceclaimrelationships, assetdeviceclaimrelationship)
	return assetdeviceclaimrelationships
}
func flattenMapAssetDeviceConfigurationRelationship(p models.AssetDeviceConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconfigurationrelationships []map[string]interface{}
	var ret models.AssetDeviceConfigurationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceconfigurationrelationship := make(map[string]interface{})
	assetdeviceconfigurationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceconfigurationrelationship["class_id"] = item.GetClassId()
	assetdeviceconfigurationrelationship["moid"] = item.GetMoid()
	assetdeviceconfigurationrelationship["object_type"] = item.GetObjectType()
	assetdeviceconfigurationrelationship["selector"] = item.GetSelector()

	assetdeviceconfigurationrelationships = append(assetdeviceconfigurationrelationships, assetdeviceconfigurationrelationship)
	return assetdeviceconfigurationrelationships
}
func flattenMapAssetDeviceConnectionRelationship(p models.AssetDeviceConnectionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconnectionrelationships []map[string]interface{}
	var ret models.AssetDeviceConnectionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceconnectionrelationship := make(map[string]interface{})
	assetdeviceconnectionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceconnectionrelationship["class_id"] = item.GetClassId()
	assetdeviceconnectionrelationship["moid"] = item.GetMoid()
	assetdeviceconnectionrelationship["object_type"] = item.GetObjectType()
	assetdeviceconnectionrelationship["selector"] = item.GetSelector()

	assetdeviceconnectionrelationships = append(assetdeviceconnectionrelationships, assetdeviceconnectionrelationship)
	return assetdeviceconnectionrelationships
}
func flattenMapAssetDeviceContractInformationRelationship(p models.AssetDeviceContractInformationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdevicecontractinformationrelationships []map[string]interface{}
	var ret models.AssetDeviceContractInformationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdevicecontractinformationrelationship := make(map[string]interface{})
	assetdevicecontractinformationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdevicecontractinformationrelationship["class_id"] = item.GetClassId()
	assetdevicecontractinformationrelationship["moid"] = item.GetMoid()
	assetdevicecontractinformationrelationship["object_type"] = item.GetObjectType()
	assetdevicecontractinformationrelationship["selector"] = item.GetSelector()

	assetdevicecontractinformationrelationships = append(assetdevicecontractinformationrelationships, assetdevicecontractinformationrelationship)
	return assetdevicecontractinformationrelationships
}
func flattenMapAssetDeviceInformation(p models.AssetDeviceInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceinformations []map[string]interface{}
	var ret models.AssetDeviceInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetdeviceinformation := make(map[string]interface{})
	assetdeviceinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceinformation["application_name"] = item.GetApplicationName()
	assetdeviceinformation["class_id"] = item.GetClassId()
	assetdeviceinformation["device_transactions"] = (func(p []models.AssetDeviceTransaction, d *schema.ResourceData) []map[string]interface{} {
		var assetdevicetransactions []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			assetdevicetransaction := make(map[string]interface{})
			assetdevicetransaction["action"] = item.GetAction()
			assetdevicetransaction["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetdevicetransaction["class_id"] = item.GetClassId()
			assetdevicetransaction["object_type"] = item.GetObjectType()
			assetdevicetransaction["status_description"] = item.GetStatusDescription()
			assetdevicetransaction["status_id"] = item.GetStatusId()
			assetdevicetransaction["timestamp"] = item.GetTimestamp()
			assetdevicetransaction["transaction_batch_id"] = item.GetTransactionBatchId()
			assetdevicetransaction["transaction_date"] = item.GetTransactionDate()
			assetdevicetransaction["transaction_sequence"] = item.GetTransactionSequence()
			assetdevicetransactions = append(assetdevicetransactions, assetdevicetransaction)
		}
		return assetdevicetransactions
	})(item.GetDeviceTransactions(), d)
	assetdeviceinformation["end_customer"] = (func(p models.AssetCustomerInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetcustomerinformations []map[string]interface{}
		var ret models.AssetCustomerInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetcustomerinformation := make(map[string]interface{})
		assetcustomerinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetcustomerinformation["address"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
			var assetaddressinformations []map[string]interface{}
			var ret models.AssetAddressInformation
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			assetaddressinformation := make(map[string]interface{})
			assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetaddressinformation["address1"] = item.GetAddress1()
			assetaddressinformation["address2"] = item.GetAddress2()
			assetaddressinformation["address3"] = item.GetAddress3()
			assetaddressinformation["city"] = item.GetCity()
			assetaddressinformation["class_id"] = item.GetClassId()
			assetaddressinformation["country"] = item.GetCountry()
			assetaddressinformation["county"] = item.GetCounty()
			assetaddressinformation["location"] = item.GetLocation()
			assetaddressinformation["name"] = item.GetName()
			assetaddressinformation["object_type"] = item.GetObjectType()
			assetaddressinformation["postal_code"] = item.GetPostalCode()
			assetaddressinformation["province"] = item.GetProvince()
			assetaddressinformation["state"] = item.GetState()

			assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
			return assetaddressinformations
		})(item.GetAddress(), d)
		assetcustomerinformation["class_id"] = item.GetClassId()
		assetcustomerinformation["id"] = item.GetId()
		assetcustomerinformation["name"] = item.GetName()
		assetcustomerinformation["object_type"] = item.GetObjectType()

		assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
		return assetcustomerinformations
	})(item.GetEndCustomer(), d)
	assetdeviceinformation["instance_id"] = item.GetInstanceId()
	assetdeviceinformation["item_type"] = item.GetItemType()
	assetdeviceinformation["mlb_offer_type"] = item.GetMlbOfferType()
	assetdeviceinformation["mlb_product_id"] = item.GetMlbProductId()
	assetdeviceinformation["mlb_product_name"] = item.GetMlbProductName()
	assetdeviceinformation["object_type"] = item.GetObjectType()
	assetdeviceinformation["old_device_id"] = item.GetOldDeviceId()
	assetdeviceinformation["old_device_status_description"] = item.GetOldDeviceStatusDescription()
	assetdeviceinformation["old_device_status_id"] = item.GetOldDeviceStatusId()
	assetdeviceinformation["old_instance_id"] = item.GetOldInstanceId()
	assetdeviceinformation["product_family"] = item.GetProductFamily()
	assetdeviceinformation["product_type"] = item.GetProductType()
	assetdeviceinformation["unit_of_measure"] = item.GetUnitOfMeasure()

	assetdeviceinformations = append(assetdeviceinformations, assetdeviceinformation)
	return assetdeviceinformations
}
func flattenMapAssetDeviceRegistrationRelationship(p models.AssetDeviceRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceregistrationrelationships []map[string]interface{}
	var ret models.AssetDeviceRegistrationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceregistrationrelationship := make(map[string]interface{})
	assetdeviceregistrationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceregistrationrelationship["class_id"] = item.GetClassId()
	assetdeviceregistrationrelationship["moid"] = item.GetMoid()
	assetdeviceregistrationrelationship["object_type"] = item.GetObjectType()
	assetdeviceregistrationrelationship["selector"] = item.GetSelector()

	assetdeviceregistrationrelationships = append(assetdeviceregistrationrelationships, assetdeviceregistrationrelationship)
	return assetdeviceregistrationrelationships
}
func flattenMapAssetDeviceStatistics(p models.AssetDeviceStatistics, d *schema.ResourceData) []map[string]interface{} {
	var assetdevicestatisticss []map[string]interface{}
	var ret models.AssetDeviceStatistics
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetdevicestatistics := make(map[string]interface{})
	assetdevicestatistics["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdevicestatistics["class_id"] = item.GetClassId()
	assetdevicestatistics["cluster_name"] = item.GetClusterName()
	assetdevicestatistics["connected"] = item.GetConnected()
	assetdevicestatistics["membership_ratio"] = item.GetMembershipRatio()
	assetdevicestatistics["memory_mirroring_factor"] = item.GetMemoryMirroringFactor()
	assetdevicestatistics["object_type"] = item.GetObjectType()
	assetdevicestatistics["vm_host"] = (func(p models.AssetVmHost, d *schema.ResourceData) []map[string]interface{} {
		var assetvmhosts []map[string]interface{}
		var ret models.AssetVmHost
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetvmhost := make(map[string]interface{})
		assetvmhost["account_moid"] = item.GetAccountMoid()
		assetvmhost["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetvmhost["class_id"] = item.GetClassId()
		assetvmhost["cluster_identity"] = item.GetClusterIdentity()
		assetvmhost["cluster_moid"] = item.GetClusterMoid()
		assetvmhost["cluster_name"] = item.GetClusterName()
		assetvmhost["object_type"] = item.GetObjectType()

		assetvmhosts = append(assetvmhosts, assetvmhost)
		return assetvmhosts
	})(item.GetVmHost(), d)

	assetdevicestatisticss = append(assetdevicestatisticss, assetdevicestatistics)
	return assetdevicestatisticss
}
func flattenMapAssetGlobalUltimate(p models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
	var assetglobalultimates []map[string]interface{}
	var ret models.AssetGlobalUltimate
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetglobalultimate := make(map[string]interface{})
	assetglobalultimate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetglobalultimate["class_id"] = item.GetClassId()
	assetglobalultimate["id"] = item.GetId()
	assetglobalultimate["name"] = item.GetName()
	assetglobalultimate["object_type"] = item.GetObjectType()

	assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
	return assetglobalultimates
}
func flattenMapAssetProductInformation(p models.AssetProductInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetproductinformations []map[string]interface{}
	var ret models.AssetProductInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetproductinformation := make(map[string]interface{})
	assetproductinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetproductinformation["bill_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.GetAddress1()
		assetaddressinformation["address2"] = item.GetAddress2()
		assetaddressinformation["address3"] = item.GetAddress3()
		assetaddressinformation["city"] = item.GetCity()
		assetaddressinformation["class_id"] = item.GetClassId()
		assetaddressinformation["country"] = item.GetCountry()
		assetaddressinformation["county"] = item.GetCounty()
		assetaddressinformation["location"] = item.GetLocation()
		assetaddressinformation["name"] = item.GetName()
		assetaddressinformation["object_type"] = item.GetObjectType()
		assetaddressinformation["postal_code"] = item.GetPostalCode()
		assetaddressinformation["province"] = item.GetProvince()
		assetaddressinformation["state"] = item.GetState()

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetBillTo(), d)
	assetproductinformation["class_id"] = item.GetClassId()
	assetproductinformation["description"] = item.GetDescription()
	assetproductinformation["family"] = item.GetFamily()
	assetproductinformation["group"] = item.GetGroup()
	assetproductinformation["number"] = item.GetNumber()
	assetproductinformation["object_type"] = item.GetObjectType()
	assetproductinformation["ship_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.GetAddress1()
		assetaddressinformation["address2"] = item.GetAddress2()
		assetaddressinformation["address3"] = item.GetAddress3()
		assetaddressinformation["city"] = item.GetCity()
		assetaddressinformation["class_id"] = item.GetClassId()
		assetaddressinformation["country"] = item.GetCountry()
		assetaddressinformation["county"] = item.GetCounty()
		assetaddressinformation["location"] = item.GetLocation()
		assetaddressinformation["name"] = item.GetName()
		assetaddressinformation["object_type"] = item.GetObjectType()
		assetaddressinformation["postal_code"] = item.GetPostalCode()
		assetaddressinformation["province"] = item.GetProvince()
		assetaddressinformation["state"] = item.GetState()

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetShipTo(), d)
	assetproductinformation["sub_type"] = item.GetSubType()

	assetproductinformations = append(assetproductinformations, assetproductinformation)
	return assetproductinformations
}
func flattenMapAssetSubscriptionRelationship(p models.AssetSubscriptionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetsubscriptionrelationships []map[string]interface{}
	var ret models.AssetSubscriptionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetsubscriptionrelationship := make(map[string]interface{})
	assetsubscriptionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetsubscriptionrelationship["class_id"] = item.GetClassId()
	assetsubscriptionrelationship["moid"] = item.GetMoid()
	assetsubscriptionrelationship["object_type"] = item.GetObjectType()
	assetsubscriptionrelationship["selector"] = item.GetSelector()

	assetsubscriptionrelationships = append(assetsubscriptionrelationships, assetsubscriptionrelationship)
	return assetsubscriptionrelationships
}
func flattenMapAssetSubscriptionAccountRelationship(p models.AssetSubscriptionAccountRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetsubscriptionaccountrelationships []map[string]interface{}
	var ret models.AssetSubscriptionAccountRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetsubscriptionaccountrelationship := make(map[string]interface{})
	assetsubscriptionaccountrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetsubscriptionaccountrelationship["class_id"] = item.GetClassId()
	assetsubscriptionaccountrelationship["moid"] = item.GetMoid()
	assetsubscriptionaccountrelationship["object_type"] = item.GetObjectType()
	assetsubscriptionaccountrelationship["selector"] = item.GetSelector()

	assetsubscriptionaccountrelationships = append(assetsubscriptionaccountrelationships, assetsubscriptionaccountrelationship)
	return assetsubscriptionaccountrelationships
}
func flattenMapAssetSudiInfo(p models.AssetSudiInfo, d *schema.ResourceData) []map[string]interface{} {
	var assetsudiinfos []map[string]interface{}
	var ret models.AssetSudiInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetsudiinfo := make(map[string]interface{})
	assetsudiinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetsudiinfo["class_id"] = item.GetClassId()
	assetsudiinfo["object_type"] = item.GetObjectType()
	assetsudiinfo["pid"] = item.GetPid()
	assetsudiinfo["serial_number"] = item.GetSerialNumber()
	assetsudiinfo["signature"] = item.GetSignature()
	assetsudiinfo["status"] = item.GetStatus()
	assetsudiinfo["sudi_certificate"] = (func(p models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
		var x509certificates []map[string]interface{}
		var ret models.X509Certificate
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		x509certificate := make(map[string]interface{})
		x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		x509certificate["class_id"] = item.GetClassId()
		x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.GetClassId()
			pkixdistinguishedname["common_name"] = item.GetCommonName()
			pkixdistinguishedname["country"] = item.GetCountry()
			pkixdistinguishedname["locality"] = item.GetLocality()
			pkixdistinguishedname["object_type"] = item.GetObjectType()
			pkixdistinguishedname["organization"] = item.GetOrganization()
			pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
			pkixdistinguishedname["state"] = item.GetState()

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetIssuer(), d)
		x509certificate["not_after"] = item.GetNotAfter()
		x509certificate["not_before"] = item.GetNotBefore()
		x509certificate["object_type"] = item.GetObjectType()
		x509certificate["pem_certificate"] = item.GetPemCertificate()
		x509certificate["sha256_fingerprint"] = item.GetSha256Fingerprint()
		x509certificate["signature_algorithm"] = item.GetSignatureAlgorithm()
		x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.GetClassId()
			pkixdistinguishedname["common_name"] = item.GetCommonName()
			pkixdistinguishedname["country"] = item.GetCountry()
			pkixdistinguishedname["locality"] = item.GetLocality()
			pkixdistinguishedname["object_type"] = item.GetObjectType()
			pkixdistinguishedname["organization"] = item.GetOrganization()
			pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
			pkixdistinguishedname["state"] = item.GetState()

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetSubject(), d)

		x509certificates = append(x509certificates, x509certificate)
		return x509certificates
	})(item.GetSudiCertificate(), d)

	assetsudiinfos = append(assetsudiinfos, assetsudiinfo)
	return assetsudiinfos
}
func flattenMapAssetTargetRelationship(p models.AssetTargetRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assettargetrelationships []map[string]interface{}
	var ret models.AssetTargetRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assettargetrelationship := make(map[string]interface{})
	assettargetrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assettargetrelationship["class_id"] = item.GetClassId()
	assettargetrelationship["moid"] = item.GetMoid()
	assettargetrelationship["object_type"] = item.GetObjectType()
	assettargetrelationship["selector"] = item.GetSelector()

	assettargetrelationships = append(assettargetrelationships, assettargetrelationship)
	return assettargetrelationships
}
func flattenMapBiosBootModeRelationship(p models.BiosBootModeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosbootmoderelationships []map[string]interface{}
	var ret models.BiosBootModeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biosbootmoderelationship := make(map[string]interface{})
	biosbootmoderelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biosbootmoderelationship["class_id"] = item.GetClassId()
	biosbootmoderelationship["moid"] = item.GetMoid()
	biosbootmoderelationship["object_type"] = item.GetObjectType()
	biosbootmoderelationship["selector"] = item.GetSelector()

	biosbootmoderelationships = append(biosbootmoderelationships, biosbootmoderelationship)
	return biosbootmoderelationships
}
func flattenMapBiosSystemBootOrderRelationship(p models.BiosSystemBootOrderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biossystembootorderrelationships []map[string]interface{}
	var ret models.BiosSystemBootOrderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biossystembootorderrelationship := make(map[string]interface{})
	biossystembootorderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biossystembootorderrelationship["class_id"] = item.GetClassId()
	biossystembootorderrelationship["moid"] = item.GetMoid()
	biossystembootorderrelationship["object_type"] = item.GetObjectType()
	biossystembootorderrelationship["selector"] = item.GetSelector()

	biossystembootorderrelationships = append(biossystembootorderrelationships, biossystembootorderrelationship)
	return biossystembootorderrelationships
}
func flattenMapBiosTokenSettingsRelationship(p models.BiosTokenSettingsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biostokensettingsrelationships []map[string]interface{}
	var ret models.BiosTokenSettingsRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biostokensettingsrelationship := make(map[string]interface{})
	biostokensettingsrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biostokensettingsrelationship["class_id"] = item.GetClassId()
	biostokensettingsrelationship["moid"] = item.GetMoid()
	biostokensettingsrelationship["object_type"] = item.GetObjectType()
	biostokensettingsrelationship["selector"] = item.GetSelector()

	biostokensettingsrelationships = append(biostokensettingsrelationships, biostokensettingsrelationship)
	return biostokensettingsrelationships
}
func flattenMapBiosUnitRelationship(p models.BiosUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrelationships []map[string]interface{}
	var ret models.BiosUnitRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biosunitrelationship := make(map[string]interface{})
	biosunitrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biosunitrelationship["class_id"] = item.GetClassId()
	biosunitrelationship["moid"] = item.GetMoid()
	biosunitrelationship["object_type"] = item.GetObjectType()
	biosunitrelationship["selector"] = item.GetSelector()

	biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	return biosunitrelationships
}
func flattenMapBiosVfSelectMemoryRasConfigurationRelationship(p models.BiosVfSelectMemoryRasConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosvfselectmemoryrasconfigurationrelationships []map[string]interface{}
	var ret models.BiosVfSelectMemoryRasConfigurationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biosvfselectmemoryrasconfigurationrelationship := make(map[string]interface{})
	biosvfselectmemoryrasconfigurationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biosvfselectmemoryrasconfigurationrelationship["class_id"] = item.GetClassId()
	biosvfselectmemoryrasconfigurationrelationship["moid"] = item.GetMoid()
	biosvfselectmemoryrasconfigurationrelationship["object_type"] = item.GetObjectType()
	biosvfselectmemoryrasconfigurationrelationship["selector"] = item.GetSelector()

	biosvfselectmemoryrasconfigurationrelationships = append(biosvfselectmemoryrasconfigurationrelationships, biosvfselectmemoryrasconfigurationrelationship)
	return biosvfselectmemoryrasconfigurationrelationships
}
func flattenMapBootDeviceBootModeRelationship(p models.BootDeviceBootModeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebootmoderelationships []map[string]interface{}
	var ret models.BootDeviceBootModeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	bootdevicebootmoderelationship := make(map[string]interface{})
	bootdevicebootmoderelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	bootdevicebootmoderelationship["class_id"] = item.GetClassId()
	bootdevicebootmoderelationship["moid"] = item.GetMoid()
	bootdevicebootmoderelationship["object_type"] = item.GetObjectType()
	bootdevicebootmoderelationship["selector"] = item.GetSelector()

	bootdevicebootmoderelationships = append(bootdevicebootmoderelationships, bootdevicebootmoderelationship)
	return bootdevicebootmoderelationships
}
func flattenMapBootDeviceBootSecurityRelationship(p models.BootDeviceBootSecurityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebootsecurityrelationships []map[string]interface{}
	var ret models.BootDeviceBootSecurityRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	bootdevicebootsecurityrelationship := make(map[string]interface{})
	bootdevicebootsecurityrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	bootdevicebootsecurityrelationship["class_id"] = item.GetClassId()
	bootdevicebootsecurityrelationship["moid"] = item.GetMoid()
	bootdevicebootsecurityrelationship["object_type"] = item.GetObjectType()
	bootdevicebootsecurityrelationship["selector"] = item.GetSelector()

	bootdevicebootsecurityrelationships = append(bootdevicebootsecurityrelationships, bootdevicebootsecurityrelationship)
	return bootdevicebootsecurityrelationships
}
func flattenMapCapabilitySwitchNetworkLimits(p models.CapabilitySwitchNetworkLimits, d *schema.ResourceData) []map[string]interface{} {
	var capabilityswitchnetworklimitss []map[string]interface{}
	var ret models.CapabilitySwitchNetworkLimits
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	capabilityswitchnetworklimits := make(map[string]interface{})
	capabilityswitchnetworklimits["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	capabilityswitchnetworklimits["class_id"] = item.GetClassId()
	capabilityswitchnetworklimits["max_compressed_port_vlan_count"] = item.GetMaxCompressedPortVlanCount()
	capabilityswitchnetworklimits["max_uncompressed_port_vlan_count"] = item.GetMaxUncompressedPortVlanCount()
	capabilityswitchnetworklimits["maximum_active_traffic_monitoring_sessions"] = item.GetMaximumActiveTrafficMonitoringSessions()
	capabilityswitchnetworklimits["maximum_ethernet_port_channels"] = item.GetMaximumEthernetPortChannels()
	capabilityswitchnetworklimits["maximum_ethernet_uplink_ports"] = item.GetMaximumEthernetUplinkPorts()
	capabilityswitchnetworklimits["maximum_fc_port_channel_members"] = item.GetMaximumFcPortChannelMembers()
	capabilityswitchnetworklimits["maximum_fc_port_channels"] = item.GetMaximumFcPortChannels()
	capabilityswitchnetworklimits["maximum_igmp_groups"] = item.GetMaximumIgmpGroups()
	capabilityswitchnetworklimits["maximum_port_channel_members"] = item.GetMaximumPortChannelMembers()
	capabilityswitchnetworklimits["maximum_primary_vlan"] = item.GetMaximumPrimaryVlan()
	capabilityswitchnetworklimits["maximum_secondary_vlan"] = item.GetMaximumSecondaryVlan()
	capabilityswitchnetworklimits["maximum_secondary_vlan_per_primary"] = item.GetMaximumSecondaryVlanPerPrimary()
	capabilityswitchnetworklimits["maximum_vifs"] = item.GetMaximumVifs()
	capabilityswitchnetworklimits["maximum_vlans"] = item.GetMaximumVlans()
	capabilityswitchnetworklimits["minimum_active_fans"] = item.GetMinimumActiveFans()
	capabilityswitchnetworklimits["object_type"] = item.GetObjectType()

	capabilityswitchnetworklimitss = append(capabilityswitchnetworklimitss, capabilityswitchnetworklimits)
	return capabilityswitchnetworklimitss
}
func flattenMapCapabilitySwitchStorageLimits(p models.CapabilitySwitchStorageLimits, d *schema.ResourceData) []map[string]interface{} {
	var capabilityswitchstoragelimitss []map[string]interface{}
	var ret models.CapabilitySwitchStorageLimits
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	capabilityswitchstoragelimits := make(map[string]interface{})
	capabilityswitchstoragelimits["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	capabilityswitchstoragelimits["class_id"] = item.GetClassId()
	capabilityswitchstoragelimits["maximum_user_zone_count"] = item.GetMaximumUserZoneCount()
	capabilityswitchstoragelimits["maximum_virtual_fc_interfaces"] = item.GetMaximumVirtualFcInterfaces()
	capabilityswitchstoragelimits["maximum_virtual_fc_interfaces_per_blade_server"] = item.GetMaximumVirtualFcInterfacesPerBladeServer()
	capabilityswitchstoragelimits["maximum_vsans"] = item.GetMaximumVsans()
	capabilityswitchstoragelimits["maximum_zone_count"] = item.GetMaximumZoneCount()
	capabilityswitchstoragelimits["object_type"] = item.GetObjectType()

	capabilityswitchstoragelimitss = append(capabilityswitchstoragelimitss, capabilityswitchstoragelimits)
	return capabilityswitchstoragelimitss
}
func flattenMapCapabilitySwitchSystemLimits(p models.CapabilitySwitchSystemLimits, d *schema.ResourceData) []map[string]interface{} {
	var capabilityswitchsystemlimitss []map[string]interface{}
	var ret models.CapabilitySwitchSystemLimits
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	capabilityswitchsystemlimits := make(map[string]interface{})
	capabilityswitchsystemlimits["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	capabilityswitchsystemlimits["class_id"] = item.GetClassId()
	capabilityswitchsystemlimits["maximum_chassis_count"] = item.GetMaximumChassisCount()
	capabilityswitchsystemlimits["maximum_fex_per_domain"] = item.GetMaximumFexPerDomain()
	capabilityswitchsystemlimits["maximum_servers_per_domain"] = item.GetMaximumServersPerDomain()
	capabilityswitchsystemlimits["object_type"] = item.GetObjectType()

	capabilityswitchsystemlimitss = append(capabilityswitchsystemlimitss, capabilityswitchsystemlimits)
	return capabilityswitchsystemlimitss
}
func flattenMapChassisConfigResultRelationship(p models.ChassisConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var chassisconfigresultrelationships []map[string]interface{}
	var ret models.ChassisConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	chassisconfigresultrelationship := make(map[string]interface{})
	chassisconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	chassisconfigresultrelationship["class_id"] = item.GetClassId()
	chassisconfigresultrelationship["moid"] = item.GetMoid()
	chassisconfigresultrelationship["object_type"] = item.GetObjectType()
	chassisconfigresultrelationship["selector"] = item.GetSelector()

	chassisconfigresultrelationships = append(chassisconfigresultrelationships, chassisconfigresultrelationship)
	return chassisconfigresultrelationships
}
func flattenMapChassisProfileRelationship(p models.ChassisProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var chassisprofilerelationships []map[string]interface{}
	var ret models.ChassisProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	chassisprofilerelationship := make(map[string]interface{})
	chassisprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	chassisprofilerelationship["class_id"] = item.GetClassId()
	chassisprofilerelationship["moid"] = item.GetMoid()
	chassisprofilerelationship["object_type"] = item.GetObjectType()
	chassisprofilerelationship["selector"] = item.GetSelector()

	chassisprofilerelationships = append(chassisprofilerelationships, chassisprofilerelationship)
	return chassisprofilerelationships
}
func flattenMapCommHttpProxyPolicyRelationship(p models.CommHttpProxyPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var commhttpproxypolicyrelationships []map[string]interface{}
	var ret models.CommHttpProxyPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	commhttpproxypolicyrelationship := make(map[string]interface{})
	commhttpproxypolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	commhttpproxypolicyrelationship["class_id"] = item.GetClassId()
	commhttpproxypolicyrelationship["moid"] = item.GetMoid()
	commhttpproxypolicyrelationship["object_type"] = item.GetObjectType()
	commhttpproxypolicyrelationship["selector"] = item.GetSelector()

	commhttpproxypolicyrelationships = append(commhttpproxypolicyrelationships, commhttpproxypolicyrelationship)
	return commhttpproxypolicyrelationships
}
func flattenMapCommIpV4Interface(p models.CommIpV4Interface, d *schema.ResourceData) []map[string]interface{} {
	var commipv4interfaces []map[string]interface{}
	var ret models.CommIpV4Interface
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	commipv4interface := make(map[string]interface{})
	commipv4interface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	commipv4interface["class_id"] = item.GetClassId()
	commipv4interface["gateway"] = item.GetGateway()
	commipv4interface["ip_address"] = item.GetIpAddress()
	commipv4interface["netmask"] = item.GetNetmask()
	commipv4interface["object_type"] = item.GetObjectType()

	commipv4interfaces = append(commipv4interfaces, commipv4interface)
	return commipv4interfaces
}
func flattenMapCommIpV6Interface(p models.CommIpV6Interface, d *schema.ResourceData) []map[string]interface{} {
	var commipv6interfaces []map[string]interface{}
	var ret models.CommIpV6Interface
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	commipv6interface := make(map[string]interface{})
	commipv6interface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	commipv6interface["class_id"] = item.GetClassId()
	commipv6interface["gateway"] = item.GetGateway()
	commipv6interface["ip_address"] = item.GetIpAddress()
	commipv6interface["object_type"] = item.GetObjectType()
	commipv6interface["prefix"] = item.GetPrefix()

	commipv6interfaces = append(commipv6interfaces, commipv6interface)
	return commipv6interfaces
}
func flattenMapComputeAlarmSummary(p models.ComputeAlarmSummary, d *schema.ResourceData) []map[string]interface{} {
	var computealarmsummarys []map[string]interface{}
	var ret models.ComputeAlarmSummary
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computealarmsummary := make(map[string]interface{})
	computealarmsummary["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computealarmsummary["class_id"] = item.GetClassId()
	computealarmsummary["critical"] = item.GetCritical()
	computealarmsummary["object_type"] = item.GetObjectType()
	computealarmsummary["warning"] = item.GetWarning()

	computealarmsummarys = append(computealarmsummarys, computealarmsummary)
	return computealarmsummarys
}
func flattenMapComputeBladeRelationship(p models.ComputeBladeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computebladerelationships []map[string]interface{}
	var ret models.ComputeBladeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computebladerelationship := make(map[string]interface{})
	computebladerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computebladerelationship["class_id"] = item.GetClassId()
	computebladerelationship["moid"] = item.GetMoid()
	computebladerelationship["object_type"] = item.GetObjectType()
	computebladerelationship["selector"] = item.GetSelector()

	computebladerelationships = append(computebladerelationships, computebladerelationship)
	return computebladerelationships
}
func flattenMapComputeBoardRelationship(p models.ComputeBoardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computeboardrelationships []map[string]interface{}
	var ret models.ComputeBoardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computeboardrelationship := make(map[string]interface{})
	computeboardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computeboardrelationship["class_id"] = item.GetClassId()
	computeboardrelationship["moid"] = item.GetMoid()
	computeboardrelationship["object_type"] = item.GetObjectType()
	computeboardrelationship["selector"] = item.GetSelector()

	computeboardrelationships = append(computeboardrelationships, computeboardrelationship)
	return computeboardrelationships
}
func flattenMapComputePersistentMemoryOperation(p models.ComputePersistentMemoryOperation, d *schema.ResourceData) []map[string]interface{} {
	var computepersistentmemoryoperations []map[string]interface{}
	var ret models.ComputePersistentMemoryOperation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computepersistentmemoryoperation := make(map[string]interface{})
	computepersistentmemoryoperation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computepersistentmemoryoperation["admin_action"] = item.GetAdminAction()
	computepersistentmemoryoperation["class_id"] = item.GetClassId()
	computepersistentmemoryoperation["is_secure_passphrase_set"] = item.GetIsSecurePassphraseSet()
	computepersistentmemoryoperation["modules"] = (func(p []models.ComputePersistentMemoryModule, d *schema.ResourceData) []map[string]interface{} {
		var computepersistentmemorymodules []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computepersistentmemorymodule := make(map[string]interface{})
			computepersistentmemorymodule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computepersistentmemorymodule["class_id"] = item.GetClassId()
			computepersistentmemorymodule["object_type"] = item.GetObjectType()
			computepersistentmemorymodule["socket_id"] = item.GetSocketId()
			computepersistentmemorymodule["socket_memory_id"] = item.GetSocketMemoryId()
			computepersistentmemorymodules = append(computepersistentmemorymodules, computepersistentmemorymodule)
		}
		return computepersistentmemorymodules
	})(item.GetModules(), d)
	computepersistentmemoryoperation["object_type"] = item.GetObjectType()
	secure_passphrase_x, exists := d.GetOk("persistent_memory_operation")
	if exists && secure_passphrase_x != nil {
		secure_passphrase_y := secure_passphrase_x.([]interface{})[0].(map[string]interface{})
		computepersistentmemoryoperation["secure_passphrase"] = secure_passphrase_y["secure_passphrase"]
	}

	computepersistentmemoryoperations = append(computepersistentmemoryoperations, computepersistentmemoryoperation)
	return computepersistentmemoryoperations
}
func flattenMapComputePhysicalRelationship(p models.ComputePhysicalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computephysicalrelationships []map[string]interface{}
	var ret models.ComputePhysicalRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computephysicalrelationship := make(map[string]interface{})
	computephysicalrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computephysicalrelationship["class_id"] = item.GetClassId()
	computephysicalrelationship["moid"] = item.GetMoid()
	computephysicalrelationship["object_type"] = item.GetObjectType()
	computephysicalrelationship["selector"] = item.GetSelector()

	computephysicalrelationships = append(computephysicalrelationships, computephysicalrelationship)
	return computephysicalrelationships
}
func flattenMapComputeRackUnitRelationship(p models.ComputeRackUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrelationships []map[string]interface{}
	var ret models.ComputeRackUnitRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computerackunitrelationship := make(map[string]interface{})
	computerackunitrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computerackunitrelationship["class_id"] = item.GetClassId()
	computerackunitrelationship["moid"] = item.GetMoid()
	computerackunitrelationship["object_type"] = item.GetObjectType()
	computerackunitrelationship["selector"] = item.GetSelector()

	computerackunitrelationships = append(computerackunitrelationships, computerackunitrelationship)
	return computerackunitrelationships
}
func flattenMapComputeServerConfig(p models.ComputeServerConfig, d *schema.ResourceData) []map[string]interface{} {
	var computeserverconfigs []map[string]interface{}
	var ret models.ComputeServerConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computeserverconfig := make(map[string]interface{})
	computeserverconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computeserverconfig["asset_tag"] = item.GetAssetTag()
	computeserverconfig["class_id"] = item.GetClassId()
	computeserverconfig["object_type"] = item.GetObjectType()
	computeserverconfig["user_label"] = item.GetUserLabel()

	computeserverconfigs = append(computeserverconfigs, computeserverconfig)
	return computeserverconfigs
}
func flattenMapComputeStorageControllerOperation(p models.ComputeStorageControllerOperation, d *schema.ResourceData) []map[string]interface{} {
	var computestoragecontrolleroperations []map[string]interface{}
	var ret models.ComputeStorageControllerOperation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computestoragecontrolleroperation := make(map[string]interface{})
	computestoragecontrolleroperation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computestoragecontrolleroperation["admin_action"] = item.GetAdminAction()
	computestoragecontrolleroperation["class_id"] = item.GetClassId()
	computestoragecontrolleroperation["controller_id"] = item.GetControllerId()
	computestoragecontrolleroperation["object_type"] = item.GetObjectType()

	computestoragecontrolleroperations = append(computestoragecontrolleroperations, computestoragecontrolleroperation)
	return computestoragecontrolleroperations
}
func flattenMapComputeStoragePhysicalDriveOperation(p models.ComputeStoragePhysicalDriveOperation, d *schema.ResourceData) []map[string]interface{} {
	var computestoragephysicaldriveoperations []map[string]interface{}
	var ret models.ComputeStoragePhysicalDriveOperation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computestoragephysicaldriveoperation := make(map[string]interface{})
	computestoragephysicaldriveoperation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computestoragephysicaldriveoperation["admin_action"] = item.GetAdminAction()
	computestoragephysicaldriveoperation["class_id"] = item.GetClassId()
	computestoragephysicaldriveoperation["controller_id"] = item.GetControllerId()
	computestoragephysicaldriveoperation["object_type"] = item.GetObjectType()
	computestoragephysicaldriveoperation["physical_drives"] = (func(p []models.ComputeStoragePhysicalDrive, d *schema.ResourceData) []map[string]interface{} {
		var computestoragephysicaldrives []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computestoragephysicaldrive := make(map[string]interface{})
			computestoragephysicaldrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computestoragephysicaldrive["class_id"] = item.GetClassId()
			computestoragephysicaldrive["object_type"] = item.GetObjectType()
			computestoragephysicaldrive["slot_number"] = item.GetSlotNumber()
			computestoragephysicaldrives = append(computestoragephysicaldrives, computestoragephysicaldrive)
		}
		return computestoragephysicaldrives
	})(item.GetPhysicalDrives(), d)

	computestoragephysicaldriveoperations = append(computestoragephysicaldriveoperations, computestoragephysicaldriveoperation)
	return computestoragephysicaldriveoperations
}
func flattenMapComputeStorageVirtualDriveOperation(p models.ComputeStorageVirtualDriveOperation, d *schema.ResourceData) []map[string]interface{} {
	var computestoragevirtualdriveoperations []map[string]interface{}
	var ret models.ComputeStorageVirtualDriveOperation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computestoragevirtualdriveoperation := make(map[string]interface{})
	computestoragevirtualdriveoperation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computestoragevirtualdriveoperation["admin_action"] = item.GetAdminAction()
	computestoragevirtualdriveoperation["class_id"] = item.GetClassId()
	computestoragevirtualdriveoperation["controller_id"] = item.GetControllerId()
	computestoragevirtualdriveoperation["object_type"] = item.GetObjectType()
	computestoragevirtualdriveoperation["virtual_drives"] = (func(p []models.ComputeStorageVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
		var computestoragevirtualdrives []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computestoragevirtualdrive := make(map[string]interface{})
			computestoragevirtualdrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computestoragevirtualdrive["class_id"] = item.GetClassId()
			computestoragevirtualdrive["id"] = item.GetId()
			computestoragevirtualdrive["object_type"] = item.GetObjectType()
			computestoragevirtualdrives = append(computestoragevirtualdrives, computestoragevirtualdrive)
		}
		return computestoragevirtualdrives
	})(item.GetVirtualDrives(), d)

	computestoragevirtualdriveoperations = append(computestoragevirtualdriveoperations, computestoragevirtualdriveoperation)
	return computestoragevirtualdriveoperations
}
func flattenMapComputeVmediaRelationship(p models.ComputeVmediaRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computevmediarelationships []map[string]interface{}
	var ret models.ComputeVmediaRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computevmediarelationship := make(map[string]interface{})
	computevmediarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computevmediarelationship["class_id"] = item.GetClassId()
	computevmediarelationship["moid"] = item.GetMoid()
	computevmediarelationship["object_type"] = item.GetObjectType()
	computevmediarelationship["selector"] = item.GetSelector()

	computevmediarelationships = append(computevmediarelationships, computevmediarelationship)
	return computevmediarelationships
}
func flattenMapCondAlarmSummary(p models.CondAlarmSummary, d *schema.ResourceData) []map[string]interface{} {
	var condalarmsummarys []map[string]interface{}
	var ret models.CondAlarmSummary
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	condalarmsummary := make(map[string]interface{})
	condalarmsummary["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	condalarmsummary["class_id"] = item.GetClassId()
	condalarmsummary["critical"] = item.GetCritical()
	condalarmsummary["object_type"] = item.GetObjectType()
	condalarmsummary["warning"] = item.GetWarning()

	condalarmsummarys = append(condalarmsummarys, condalarmsummary)
	return condalarmsummarys
}
func flattenMapCondHclStatusRelationship(p models.CondHclStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusrelationships []map[string]interface{}
	var ret models.CondHclStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	condhclstatusrelationship := make(map[string]interface{})
	condhclstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	condhclstatusrelationship["class_id"] = item.GetClassId()
	condhclstatusrelationship["moid"] = item.GetMoid()
	condhclstatusrelationship["object_type"] = item.GetObjectType()
	condhclstatusrelationship["selector"] = item.GetSelector()

	condhclstatusrelationships = append(condhclstatusrelationships, condhclstatusrelationship)
	return condhclstatusrelationships
}
func flattenMapConfigExportedItemRelationship(p models.ConfigExportedItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configexporteditemrelationships []map[string]interface{}
	var ret models.ConfigExportedItemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	configexporteditemrelationship := make(map[string]interface{})
	configexporteditemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configexporteditemrelationship["class_id"] = item.GetClassId()
	configexporteditemrelationship["moid"] = item.GetMoid()
	configexporteditemrelationship["object_type"] = item.GetObjectType()
	configexporteditemrelationship["selector"] = item.GetSelector()

	configexporteditemrelationships = append(configexporteditemrelationships, configexporteditemrelationship)
	return configexporteditemrelationships
}
func flattenMapConfigExporterRelationship(p models.ConfigExporterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configexporterrelationships []map[string]interface{}
	var ret models.ConfigExporterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	configexporterrelationship := make(map[string]interface{})
	configexporterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configexporterrelationship["class_id"] = item.GetClassId()
	configexporterrelationship["moid"] = item.GetMoid()
	configexporterrelationship["object_type"] = item.GetObjectType()
	configexporterrelationship["selector"] = item.GetSelector()

	configexporterrelationships = append(configexporterrelationships, configexporterrelationship)
	return configexporterrelationships
}
func flattenMapConfigImporterRelationship(p models.ConfigImporterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configimporterrelationships []map[string]interface{}
	var ret models.ConfigImporterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	configimporterrelationship := make(map[string]interface{})
	configimporterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configimporterrelationship["class_id"] = item.GetClassId()
	configimporterrelationship["moid"] = item.GetMoid()
	configimporterrelationship["object_type"] = item.GetObjectType()
	configimporterrelationship["selector"] = item.GetSelector()

	configimporterrelationships = append(configimporterrelationships, configimporterrelationship)
	return configimporterrelationships
}
func flattenMapConfigMoRef(p models.ConfigMoRef, d *schema.ResourceData) []map[string]interface{} {
	var configmorefs []map[string]interface{}
	var ret models.ConfigMoRef
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	configmoref := make(map[string]interface{})
	configmoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configmoref["class_id"] = item.GetClassId()
	configmoref["moid"] = item.GetMoid()
	configmoref["object_type"] = item.GetObjectType()

	configmorefs = append(configmorefs, configmoref)
	return configmorefs
}
func flattenMapConnectorFileChecksum(p models.ConnectorFileChecksum, d *schema.ResourceData) []map[string]interface{} {
	var connectorfilechecksums []map[string]interface{}
	var ret models.ConnectorFileChecksum
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	connectorfilechecksum := make(map[string]interface{})
	connectorfilechecksum["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	connectorfilechecksum["class_id"] = item.GetClassId()
	connectorfilechecksum["hash_algorithm"] = item.GetHashAlgorithm()
	connectorfilechecksum["object_type"] = item.GetObjectType()

	connectorfilechecksums = append(connectorfilechecksums, connectorfilechecksum)
	return connectorfilechecksums
}
func flattenMapConnectorPlatformParamBase(p models.ConnectorPlatformParamBase, d *schema.ResourceData) []map[string]interface{} {
	var connectorplatformparambases []map[string]interface{}
	var ret models.ConnectorPlatformParamBase
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	connectorplatformparambase := make(map[string]interface{})
	connectorplatformparambase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	connectorplatformparambase["class_id"] = item.GetClassId()
	connectorplatformparambase["object_type"] = item.GetObjectType()

	connectorplatformparambases = append(connectorplatformparambases, connectorplatformparambase)
	return connectorplatformparambases
}
func flattenMapEquipmentBaseRelationship(p models.EquipmentBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentbaserelationships []map[string]interface{}
	var ret models.EquipmentBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentbaserelationship := make(map[string]interface{})
	equipmentbaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentbaserelationship["class_id"] = item.GetClassId()
	equipmentbaserelationship["moid"] = item.GetMoid()
	equipmentbaserelationship["object_type"] = item.GetObjectType()
	equipmentbaserelationship["selector"] = item.GetSelector()

	equipmentbaserelationships = append(equipmentbaserelationships, equipmentbaserelationship)
	return equipmentbaserelationships
}
func flattenMapEquipmentChassisRelationship(p models.EquipmentChassisRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentchassisrelationships []map[string]interface{}
	var ret models.EquipmentChassisRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentchassisrelationship := make(map[string]interface{})
	equipmentchassisrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentchassisrelationship["class_id"] = item.GetClassId()
	equipmentchassisrelationship["moid"] = item.GetMoid()
	equipmentchassisrelationship["object_type"] = item.GetObjectType()
	equipmentchassisrelationship["selector"] = item.GetSelector()

	equipmentchassisrelationships = append(equipmentchassisrelationships, equipmentchassisrelationship)
	return equipmentchassisrelationships
}
func flattenMapEquipmentFanModuleRelationship(p models.EquipmentFanModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerelationships []map[string]interface{}
	var ret models.EquipmentFanModuleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentfanmodulerelationship := make(map[string]interface{})
	equipmentfanmodulerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentfanmodulerelationship["class_id"] = item.GetClassId()
	equipmentfanmodulerelationship["moid"] = item.GetMoid()
	equipmentfanmodulerelationship["object_type"] = item.GetObjectType()
	equipmentfanmodulerelationship["selector"] = item.GetSelector()

	equipmentfanmodulerelationships = append(equipmentfanmodulerelationships, equipmentfanmodulerelationship)
	return equipmentfanmodulerelationships
}
func flattenMapEquipmentFexRelationship(p models.EquipmentFexRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfexrelationships []map[string]interface{}
	var ret models.EquipmentFexRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentfexrelationship := make(map[string]interface{})
	equipmentfexrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentfexrelationship["class_id"] = item.GetClassId()
	equipmentfexrelationship["moid"] = item.GetMoid()
	equipmentfexrelationship["object_type"] = item.GetObjectType()
	equipmentfexrelationship["selector"] = item.GetSelector()

	equipmentfexrelationships = append(equipmentfexrelationships, equipmentfexrelationship)
	return equipmentfexrelationships
}
func flattenMapEquipmentFruRelationship(p models.EquipmentFruRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfrurelationships []map[string]interface{}
	var ret models.EquipmentFruRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentfrurelationship := make(map[string]interface{})
	equipmentfrurelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentfrurelationship["class_id"] = item.GetClassId()
	equipmentfrurelationship["moid"] = item.GetMoid()
	equipmentfrurelationship["object_type"] = item.GetObjectType()
	equipmentfrurelationship["selector"] = item.GetSelector()

	equipmentfrurelationships = append(equipmentfrurelationships, equipmentfrurelationship)
	return equipmentfrurelationships
}
func flattenMapEquipmentIoCardRelationship(p models.EquipmentIoCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardrelationships []map[string]interface{}
	var ret models.EquipmentIoCardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentiocardrelationship := make(map[string]interface{})
	equipmentiocardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentiocardrelationship["class_id"] = item.GetClassId()
	equipmentiocardrelationship["moid"] = item.GetMoid()
	equipmentiocardrelationship["object_type"] = item.GetObjectType()
	equipmentiocardrelationship["selector"] = item.GetSelector()

	equipmentiocardrelationships = append(equipmentiocardrelationships, equipmentiocardrelationship)
	return equipmentiocardrelationships
}
func flattenMapEquipmentIoCardBaseRelationship(p models.EquipmentIoCardBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardbaserelationships []map[string]interface{}
	var ret models.EquipmentIoCardBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentiocardbaserelationship := make(map[string]interface{})
	equipmentiocardbaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentiocardbaserelationship["class_id"] = item.GetClassId()
	equipmentiocardbaserelationship["moid"] = item.GetMoid()
	equipmentiocardbaserelationship["object_type"] = item.GetObjectType()
	equipmentiocardbaserelationship["selector"] = item.GetSelector()

	equipmentiocardbaserelationships = append(equipmentiocardbaserelationships, equipmentiocardbaserelationship)
	return equipmentiocardbaserelationships
}
func flattenMapEquipmentLocatorLedRelationship(p models.EquipmentLocatorLedRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentlocatorledrelationships []map[string]interface{}
	var ret models.EquipmentLocatorLedRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentlocatorledrelationship := make(map[string]interface{})
	equipmentlocatorledrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentlocatorledrelationship["class_id"] = item.GetClassId()
	equipmentlocatorledrelationship["moid"] = item.GetMoid()
	equipmentlocatorledrelationship["object_type"] = item.GetObjectType()
	equipmentlocatorledrelationship["selector"] = item.GetSelector()

	equipmentlocatorledrelationships = append(equipmentlocatorledrelationships, equipmentlocatorledrelationship)
	return equipmentlocatorledrelationships
}
func flattenMapEquipmentPhysicalIdentityRelationship(p models.EquipmentPhysicalIdentityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentphysicalidentityrelationships []map[string]interface{}
	var ret models.EquipmentPhysicalIdentityRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentphysicalidentityrelationship := make(map[string]interface{})
	equipmentphysicalidentityrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentphysicalidentityrelationship["class_id"] = item.GetClassId()
	equipmentphysicalidentityrelationship["moid"] = item.GetMoid()
	equipmentphysicalidentityrelationship["object_type"] = item.GetObjectType()
	equipmentphysicalidentityrelationship["selector"] = item.GetSelector()

	equipmentphysicalidentityrelationships = append(equipmentphysicalidentityrelationships, equipmentphysicalidentityrelationship)
	return equipmentphysicalidentityrelationships
}
func flattenMapEquipmentPsuControlRelationship(p models.EquipmentPsuControlRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentpsucontrolrelationships []map[string]interface{}
	var ret models.EquipmentPsuControlRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentpsucontrolrelationship := make(map[string]interface{})
	equipmentpsucontrolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentpsucontrolrelationship["class_id"] = item.GetClassId()
	equipmentpsucontrolrelationship["moid"] = item.GetMoid()
	equipmentpsucontrolrelationship["object_type"] = item.GetObjectType()
	equipmentpsucontrolrelationship["selector"] = item.GetSelector()

	equipmentpsucontrolrelationships = append(equipmentpsucontrolrelationships, equipmentpsucontrolrelationship)
	return equipmentpsucontrolrelationships
}
func flattenMapEquipmentRackEnclosureRelationship(p models.EquipmentRackEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosurerelationships []map[string]interface{}
	var ret models.EquipmentRackEnclosureRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentrackenclosurerelationship := make(map[string]interface{})
	equipmentrackenclosurerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentrackenclosurerelationship["class_id"] = item.GetClassId()
	equipmentrackenclosurerelationship["moid"] = item.GetMoid()
	equipmentrackenclosurerelationship["object_type"] = item.GetObjectType()
	equipmentrackenclosurerelationship["selector"] = item.GetSelector()

	equipmentrackenclosurerelationships = append(equipmentrackenclosurerelationships, equipmentrackenclosurerelationship)
	return equipmentrackenclosurerelationships
}
func flattenMapEquipmentRackEnclosureSlotRelationship(p models.EquipmentRackEnclosureSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrelationships []map[string]interface{}
	var ret models.EquipmentRackEnclosureSlotRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentrackenclosureslotrelationship := make(map[string]interface{})
	equipmentrackenclosureslotrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentrackenclosureslotrelationship["class_id"] = item.GetClassId()
	equipmentrackenclosureslotrelationship["moid"] = item.GetMoid()
	equipmentrackenclosureslotrelationship["object_type"] = item.GetObjectType()
	equipmentrackenclosureslotrelationship["selector"] = item.GetSelector()

	equipmentrackenclosureslotrelationships = append(equipmentrackenclosureslotrelationships, equipmentrackenclosureslotrelationship)
	return equipmentrackenclosureslotrelationships
}
func flattenMapEquipmentSharedIoModuleRelationship(p models.EquipmentSharedIoModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsharediomodulerelationships []map[string]interface{}
	var ret models.EquipmentSharedIoModuleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentsharediomodulerelationship := make(map[string]interface{})
	equipmentsharediomodulerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentsharediomodulerelationship["class_id"] = item.GetClassId()
	equipmentsharediomodulerelationship["moid"] = item.GetMoid()
	equipmentsharediomodulerelationship["object_type"] = item.GetObjectType()
	equipmentsharediomodulerelationship["selector"] = item.GetSelector()

	equipmentsharediomodulerelationships = append(equipmentsharediomodulerelationships, equipmentsharediomodulerelationship)
	return equipmentsharediomodulerelationships
}
func flattenMapEquipmentSwitchCardRelationship(p models.EquipmentSwitchCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrelationships []map[string]interface{}
	var ret models.EquipmentSwitchCardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentswitchcardrelationship := make(map[string]interface{})
	equipmentswitchcardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentswitchcardrelationship["class_id"] = item.GetClassId()
	equipmentswitchcardrelationship["moid"] = item.GetMoid()
	equipmentswitchcardrelationship["object_type"] = item.GetObjectType()
	equipmentswitchcardrelationship["selector"] = item.GetSelector()

	equipmentswitchcardrelationships = append(equipmentswitchcardrelationships, equipmentswitchcardrelationship)
	return equipmentswitchcardrelationships
}
func flattenMapEquipmentSystemIoControllerRelationship(p models.EquipmentSystemIoControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrelationships []map[string]interface{}
	var ret models.EquipmentSystemIoControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentsystemiocontrollerrelationship := make(map[string]interface{})
	equipmentsystemiocontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentsystemiocontrollerrelationship["class_id"] = item.GetClassId()
	equipmentsystemiocontrollerrelationship["moid"] = item.GetMoid()
	equipmentsystemiocontrollerrelationship["object_type"] = item.GetObjectType()
	equipmentsystemiocontrollerrelationship["selector"] = item.GetSelector()

	equipmentsystemiocontrollerrelationships = append(equipmentsystemiocontrollerrelationships, equipmentsystemiocontrollerrelationship)
	return equipmentsystemiocontrollerrelationships
}
func flattenMapEtherPhysicalPortRelationship(p models.EtherPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportrelationships []map[string]interface{}
	var ret models.EtherPhysicalPortRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	etherphysicalportrelationship := make(map[string]interface{})
	etherphysicalportrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	etherphysicalportrelationship["class_id"] = item.GetClassId()
	etherphysicalportrelationship["moid"] = item.GetMoid()
	etherphysicalportrelationship["object_type"] = item.GetObjectType()
	etherphysicalportrelationship["selector"] = item.GetSelector()

	etherphysicalportrelationships = append(etherphysicalportrelationships, etherphysicalportrelationship)
	return etherphysicalportrelationships
}
func flattenMapEtherPhysicalPortBaseRelationship(p models.EtherPhysicalPortBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportbaserelationships []map[string]interface{}
	var ret models.EtherPhysicalPortBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	etherphysicalportbaserelationship := make(map[string]interface{})
	etherphysicalportbaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	etherphysicalportbaserelationship["class_id"] = item.GetClassId()
	etherphysicalportbaserelationship["moid"] = item.GetMoid()
	etherphysicalportbaserelationship["object_type"] = item.GetObjectType()
	etherphysicalportbaserelationship["selector"] = item.GetSelector()

	etherphysicalportbaserelationships = append(etherphysicalportbaserelationships, etherphysicalportbaserelationship)
	return etherphysicalportbaserelationships
}
func flattenMapFabricConfigResultRelationship(p models.FabricConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricconfigresultrelationships []map[string]interface{}
	var ret models.FabricConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricconfigresultrelationship := make(map[string]interface{})
	fabricconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricconfigresultrelationship["class_id"] = item.GetClassId()
	fabricconfigresultrelationship["moid"] = item.GetMoid()
	fabricconfigresultrelationship["object_type"] = item.GetObjectType()
	fabricconfigresultrelationship["selector"] = item.GetSelector()

	fabricconfigresultrelationships = append(fabricconfigresultrelationships, fabricconfigresultrelationship)
	return fabricconfigresultrelationships
}
func flattenMapFabricEthNetworkControlPolicyRelationship(p models.FabricEthNetworkControlPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricethnetworkcontrolpolicyrelationships []map[string]interface{}
	var ret models.FabricEthNetworkControlPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricethnetworkcontrolpolicyrelationship := make(map[string]interface{})
	fabricethnetworkcontrolpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricethnetworkcontrolpolicyrelationship["class_id"] = item.GetClassId()
	fabricethnetworkcontrolpolicyrelationship["moid"] = item.GetMoid()
	fabricethnetworkcontrolpolicyrelationship["object_type"] = item.GetObjectType()
	fabricethnetworkcontrolpolicyrelationship["selector"] = item.GetSelector()

	fabricethnetworkcontrolpolicyrelationships = append(fabricethnetworkcontrolpolicyrelationships, fabricethnetworkcontrolpolicyrelationship)
	return fabricethnetworkcontrolpolicyrelationships
}
func flattenMapFabricEthNetworkGroupPolicyRelationship(p models.FabricEthNetworkGroupPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricethnetworkgrouppolicyrelationships []map[string]interface{}
	var ret models.FabricEthNetworkGroupPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricethnetworkgrouppolicyrelationship := make(map[string]interface{})
	fabricethnetworkgrouppolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricethnetworkgrouppolicyrelationship["class_id"] = item.GetClassId()
	fabricethnetworkgrouppolicyrelationship["moid"] = item.GetMoid()
	fabricethnetworkgrouppolicyrelationship["object_type"] = item.GetObjectType()
	fabricethnetworkgrouppolicyrelationship["selector"] = item.GetSelector()

	fabricethnetworkgrouppolicyrelationships = append(fabricethnetworkgrouppolicyrelationships, fabricethnetworkgrouppolicyrelationship)
	return fabricethnetworkgrouppolicyrelationships
}
func flattenMapFabricEthNetworkPolicyRelationship(p models.FabricEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricethnetworkpolicyrelationships []map[string]interface{}
	var ret models.FabricEthNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricethnetworkpolicyrelationship := make(map[string]interface{})
	fabricethnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricethnetworkpolicyrelationship["class_id"] = item.GetClassId()
	fabricethnetworkpolicyrelationship["moid"] = item.GetMoid()
	fabricethnetworkpolicyrelationship["object_type"] = item.GetObjectType()
	fabricethnetworkpolicyrelationship["selector"] = item.GetSelector()

	fabricethnetworkpolicyrelationships = append(fabricethnetworkpolicyrelationships, fabricethnetworkpolicyrelationship)
	return fabricethnetworkpolicyrelationships
}
func flattenMapFabricFcNetworkPolicyRelationship(p models.FabricFcNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricfcnetworkpolicyrelationships []map[string]interface{}
	var ret models.FabricFcNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricfcnetworkpolicyrelationship := make(map[string]interface{})
	fabricfcnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricfcnetworkpolicyrelationship["class_id"] = item.GetClassId()
	fabricfcnetworkpolicyrelationship["moid"] = item.GetMoid()
	fabricfcnetworkpolicyrelationship["object_type"] = item.GetObjectType()
	fabricfcnetworkpolicyrelationship["selector"] = item.GetSelector()

	fabricfcnetworkpolicyrelationships = append(fabricfcnetworkpolicyrelationships, fabricfcnetworkpolicyrelationship)
	return fabricfcnetworkpolicyrelationships
}
func flattenMapFabricFlowControlPolicyRelationship(p models.FabricFlowControlPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricflowcontrolpolicyrelationships []map[string]interface{}
	var ret models.FabricFlowControlPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricflowcontrolpolicyrelationship := make(map[string]interface{})
	fabricflowcontrolpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricflowcontrolpolicyrelationship["class_id"] = item.GetClassId()
	fabricflowcontrolpolicyrelationship["moid"] = item.GetMoid()
	fabricflowcontrolpolicyrelationship["object_type"] = item.GetObjectType()
	fabricflowcontrolpolicyrelationship["selector"] = item.GetSelector()

	fabricflowcontrolpolicyrelationships = append(fabricflowcontrolpolicyrelationships, fabricflowcontrolpolicyrelationship)
	return fabricflowcontrolpolicyrelationships
}
func flattenMapFabricLinkAggregationPolicyRelationship(p models.FabricLinkAggregationPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabriclinkaggregationpolicyrelationships []map[string]interface{}
	var ret models.FabricLinkAggregationPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabriclinkaggregationpolicyrelationship := make(map[string]interface{})
	fabriclinkaggregationpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabriclinkaggregationpolicyrelationship["class_id"] = item.GetClassId()
	fabriclinkaggregationpolicyrelationship["moid"] = item.GetMoid()
	fabriclinkaggregationpolicyrelationship["object_type"] = item.GetObjectType()
	fabriclinkaggregationpolicyrelationship["selector"] = item.GetSelector()

	fabriclinkaggregationpolicyrelationships = append(fabriclinkaggregationpolicyrelationships, fabriclinkaggregationpolicyrelationship)
	return fabriclinkaggregationpolicyrelationships
}
func flattenMapFabricLinkControlPolicyRelationship(p models.FabricLinkControlPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabriclinkcontrolpolicyrelationships []map[string]interface{}
	var ret models.FabricLinkControlPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabriclinkcontrolpolicyrelationship := make(map[string]interface{})
	fabriclinkcontrolpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabriclinkcontrolpolicyrelationship["class_id"] = item.GetClassId()
	fabriclinkcontrolpolicyrelationship["moid"] = item.GetMoid()
	fabriclinkcontrolpolicyrelationship["object_type"] = item.GetObjectType()
	fabriclinkcontrolpolicyrelationship["selector"] = item.GetSelector()

	fabriclinkcontrolpolicyrelationships = append(fabriclinkcontrolpolicyrelationships, fabriclinkcontrolpolicyrelationship)
	return fabriclinkcontrolpolicyrelationships
}
func flattenMapFabricLldpSettings(p models.FabricLldpSettings, d *schema.ResourceData) []map[string]interface{} {
	var fabriclldpsettingss []map[string]interface{}
	var ret models.FabricLldpSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fabriclldpsettings := make(map[string]interface{})
	fabriclldpsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabriclldpsettings["class_id"] = item.GetClassId()
	fabriclldpsettings["object_type"] = item.GetObjectType()
	fabriclldpsettings["receive_enabled"] = item.GetReceiveEnabled()
	fabriclldpsettings["transmit_enabled"] = item.GetTransmitEnabled()

	fabriclldpsettingss = append(fabriclldpsettingss, fabriclldpsettings)
	return fabriclldpsettingss
}
func flattenMapFabricMacAgingSettings(p models.FabricMacAgingSettings, d *schema.ResourceData) []map[string]interface{} {
	var fabricmacagingsettingss []map[string]interface{}
	var ret models.FabricMacAgingSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fabricmacagingsettings := make(map[string]interface{})
	fabricmacagingsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricmacagingsettings["class_id"] = item.GetClassId()
	fabricmacagingsettings["mac_aging_option"] = item.GetMacAgingOption()
	fabricmacagingsettings["mac_aging_time"] = item.GetMacAgingTime()
	fabricmacagingsettings["object_type"] = item.GetObjectType()

	fabricmacagingsettingss = append(fabricmacagingsettingss, fabricmacagingsettings)
	return fabricmacagingsettingss
}
func flattenMapFabricMulticastPolicyRelationship(p models.FabricMulticastPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricmulticastpolicyrelationships []map[string]interface{}
	var ret models.FabricMulticastPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricmulticastpolicyrelationship := make(map[string]interface{})
	fabricmulticastpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricmulticastpolicyrelationship["class_id"] = item.GetClassId()
	fabricmulticastpolicyrelationship["moid"] = item.GetMoid()
	fabricmulticastpolicyrelationship["object_type"] = item.GetObjectType()
	fabricmulticastpolicyrelationship["selector"] = item.GetSelector()

	fabricmulticastpolicyrelationships = append(fabricmulticastpolicyrelationships, fabricmulticastpolicyrelationship)
	return fabricmulticastpolicyrelationships
}
func flattenMapFabricPortPolicyRelationship(p models.FabricPortPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricportpolicyrelationships []map[string]interface{}
	var ret models.FabricPortPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricportpolicyrelationship := make(map[string]interface{})
	fabricportpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricportpolicyrelationship["class_id"] = item.GetClassId()
	fabricportpolicyrelationship["moid"] = item.GetMoid()
	fabricportpolicyrelationship["object_type"] = item.GetObjectType()
	fabricportpolicyrelationship["selector"] = item.GetSelector()

	fabricportpolicyrelationships = append(fabricportpolicyrelationships, fabricportpolicyrelationship)
	return fabricportpolicyrelationships
}
func flattenMapFabricSwitchClusterProfileRelationship(p models.FabricSwitchClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricswitchclusterprofilerelationships []map[string]interface{}
	var ret models.FabricSwitchClusterProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricswitchclusterprofilerelationship := make(map[string]interface{})
	fabricswitchclusterprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricswitchclusterprofilerelationship["class_id"] = item.GetClassId()
	fabricswitchclusterprofilerelationship["moid"] = item.GetMoid()
	fabricswitchclusterprofilerelationship["object_type"] = item.GetObjectType()
	fabricswitchclusterprofilerelationship["selector"] = item.GetSelector()

	fabricswitchclusterprofilerelationships = append(fabricswitchclusterprofilerelationships, fabricswitchclusterprofilerelationship)
	return fabricswitchclusterprofilerelationships
}
func flattenMapFabricSwitchProfileRelationship(p models.FabricSwitchProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricswitchprofilerelationships []map[string]interface{}
	var ret models.FabricSwitchProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricswitchprofilerelationship := make(map[string]interface{})
	fabricswitchprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricswitchprofilerelationship["class_id"] = item.GetClassId()
	fabricswitchprofilerelationship["moid"] = item.GetMoid()
	fabricswitchprofilerelationship["object_type"] = item.GetObjectType()
	fabricswitchprofilerelationship["selector"] = item.GetSelector()

	fabricswitchprofilerelationships = append(fabricswitchprofilerelationships, fabricswitchprofilerelationship)
	return fabricswitchprofilerelationships
}
func flattenMapFabricUdldGlobalSettings(p models.FabricUdldGlobalSettings, d *schema.ResourceData) []map[string]interface{} {
	var fabricudldglobalsettingss []map[string]interface{}
	var ret models.FabricUdldGlobalSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fabricudldglobalsettings := make(map[string]interface{})
	fabricudldglobalsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricudldglobalsettings["class_id"] = item.GetClassId()
	fabricudldglobalsettings["message_interval"] = item.GetMessageInterval()
	fabricudldglobalsettings["object_type"] = item.GetObjectType()
	fabricudldglobalsettings["recovery_action"] = item.GetRecoveryAction()

	fabricudldglobalsettingss = append(fabricudldglobalsettingss, fabricudldglobalsettings)
	return fabricudldglobalsettingss
}
func flattenMapFabricUdldSettings(p models.FabricUdldSettings, d *schema.ResourceData) []map[string]interface{} {
	var fabricudldsettingss []map[string]interface{}
	var ret models.FabricUdldSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fabricudldsettings := make(map[string]interface{})
	fabricudldsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricudldsettings["admin_state"] = item.GetAdminState()
	fabricudldsettings["class_id"] = item.GetClassId()
	fabricudldsettings["mode"] = item.GetMode()
	fabricudldsettings["object_type"] = item.GetObjectType()

	fabricudldsettingss = append(fabricudldsettingss, fabricudldsettings)
	return fabricudldsettingss
}
func flattenMapFabricVlanSettings(p models.FabricVlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var fabricvlansettingss []map[string]interface{}
	var ret models.FabricVlanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fabricvlansettings := make(map[string]interface{})
	fabricvlansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricvlansettings["allowed_vlans"] = item.GetAllowedVlans()
	fabricvlansettings["class_id"] = item.GetClassId()
	fabricvlansettings["native_vlan"] = item.GetNativeVlan()
	fabricvlansettings["object_type"] = item.GetObjectType()

	fabricvlansettingss = append(fabricvlansettingss, fabricvlansettings)
	return fabricvlansettingss
}
func flattenMapFcpoolBlock(p models.FcpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolblocks []map[string]interface{}
	var ret models.FcpoolBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fcpoolblock := make(map[string]interface{})
	fcpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolblock["class_id"] = item.GetClassId()
	fcpoolblock["from"] = item.GetFrom()
	fcpoolblock["object_type"] = item.GetObjectType()
	fcpoolblock["size"] = item.GetSize()
	fcpoolblock["to"] = item.GetTo()

	fcpoolblocks = append(fcpoolblocks, fcpoolblock)
	return fcpoolblocks
}
func flattenMapFcpoolFcBlockRelationship(p models.FcpoolFcBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolfcblockrelationships []map[string]interface{}
	var ret models.FcpoolFcBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolfcblockrelationship := make(map[string]interface{})
	fcpoolfcblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolfcblockrelationship["class_id"] = item.GetClassId()
	fcpoolfcblockrelationship["moid"] = item.GetMoid()
	fcpoolfcblockrelationship["object_type"] = item.GetObjectType()
	fcpoolfcblockrelationship["selector"] = item.GetSelector()

	fcpoolfcblockrelationships = append(fcpoolfcblockrelationships, fcpoolfcblockrelationship)
	return fcpoolfcblockrelationships
}
func flattenMapFcpoolLeaseRelationship(p models.FcpoolLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolleaserelationships []map[string]interface{}
	var ret models.FcpoolLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolleaserelationship := make(map[string]interface{})
	fcpoolleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolleaserelationship["class_id"] = item.GetClassId()
	fcpoolleaserelationship["moid"] = item.GetMoid()
	fcpoolleaserelationship["object_type"] = item.GetObjectType()
	fcpoolleaserelationship["selector"] = item.GetSelector()

	fcpoolleaserelationships = append(fcpoolleaserelationships, fcpoolleaserelationship)
	return fcpoolleaserelationships
}
func flattenMapFcpoolPoolRelationship(p models.FcpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolpoolrelationships []map[string]interface{}
	var ret models.FcpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolpoolrelationship := make(map[string]interface{})
	fcpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolpoolrelationship["class_id"] = item.GetClassId()
	fcpoolpoolrelationship["moid"] = item.GetMoid()
	fcpoolpoolrelationship["object_type"] = item.GetObjectType()
	fcpoolpoolrelationship["selector"] = item.GetSelector()

	fcpoolpoolrelationships = append(fcpoolpoolrelationships, fcpoolpoolrelationship)
	return fcpoolpoolrelationships
}
func flattenMapFcpoolPoolMemberRelationship(p models.FcpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolpoolmemberrelationships []map[string]interface{}
	var ret models.FcpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolpoolmemberrelationship := make(map[string]interface{})
	fcpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolpoolmemberrelationship["class_id"] = item.GetClassId()
	fcpoolpoolmemberrelationship["moid"] = item.GetMoid()
	fcpoolpoolmemberrelationship["object_type"] = item.GetObjectType()
	fcpoolpoolmemberrelationship["selector"] = item.GetSelector()

	fcpoolpoolmemberrelationships = append(fcpoolpoolmemberrelationships, fcpoolpoolmemberrelationship)
	return fcpoolpoolmemberrelationships
}
func flattenMapFcpoolUniverseRelationship(p models.FcpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpooluniverserelationships []map[string]interface{}
	var ret models.FcpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpooluniverserelationship := make(map[string]interface{})
	fcpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpooluniverserelationship["class_id"] = item.GetClassId()
	fcpooluniverserelationship["moid"] = item.GetMoid()
	fcpooluniverserelationship["object_type"] = item.GetObjectType()
	fcpooluniverserelationship["selector"] = item.GetSelector()

	fcpooluniverserelationships = append(fcpooluniverserelationships, fcpooluniverserelationship)
	return fcpooluniverserelationships
}
func flattenMapFirmwareBaseDistributableRelationship(p models.FirmwareBaseDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarebasedistributablerelationships []map[string]interface{}
	var ret models.FirmwareBaseDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwarebasedistributablerelationship := make(map[string]interface{})
	firmwarebasedistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwarebasedistributablerelationship["class_id"] = item.GetClassId()
	firmwarebasedistributablerelationship["moid"] = item.GetMoid()
	firmwarebasedistributablerelationship["object_type"] = item.GetObjectType()
	firmwarebasedistributablerelationship["selector"] = item.GetSelector()

	firmwarebasedistributablerelationships = append(firmwarebasedistributablerelationships, firmwarebasedistributablerelationship)
	return firmwarebasedistributablerelationships
}
func flattenMapFirmwareDirectDownload(p models.FirmwareDirectDownload, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredirectdownloads []map[string]interface{}
	var ret models.FirmwareDirectDownload
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	firmwaredirectdownload := make(map[string]interface{})
	firmwaredirectdownload["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwaredirectdownload["class_id"] = item.GetClassId()
	firmwaredirectdownload["http_server"] = (func(p models.FirmwareHttpServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		var ret models.FirmwareHttpServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarehttpserver := make(map[string]interface{})
		firmwarehttpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarehttpserver["class_id"] = item.GetClassId()
		firmwarehttpserver["location_link"] = item.GetLocationLink()
		firmwarehttpserver["mount_options"] = item.GetMountOptions()
		firmwarehttpserver["object_type"] = item.GetObjectType()

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.GetHttpServer(), d)
	firmwaredirectdownload["image_source"] = item.GetImageSource()
	firmwaredirectdownload["is_password_set"] = item.GetIsPasswordSet()
	firmwaredirectdownload["object_type"] = item.GetObjectType()
	password_x, exists := d.GetOk("direct_download")
	if exists && password_x != nil {
		password_y := password_x.([]interface{})[0].(map[string]interface{})
		firmwaredirectdownload["password"] = password_y["password"]
	}
	firmwaredirectdownload["upgradeoption"] = item.GetUpgradeoption()
	firmwaredirectdownload["username"] = item.GetUsername()

	firmwaredirectdownloads = append(firmwaredirectdownloads, firmwaredirectdownload)
	return firmwaredirectdownloads
}
func flattenMapFirmwareDistributableRelationship(p models.FirmwareDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredistributablerelationships []map[string]interface{}
	var ret models.FirmwareDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwaredistributablerelationship := make(map[string]interface{})
	firmwaredistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwaredistributablerelationship["class_id"] = item.GetClassId()
	firmwaredistributablerelationship["moid"] = item.GetMoid()
	firmwaredistributablerelationship["object_type"] = item.GetObjectType()
	firmwaredistributablerelationship["selector"] = item.GetSelector()

	firmwaredistributablerelationships = append(firmwaredistributablerelationships, firmwaredistributablerelationship)
	return firmwaredistributablerelationships
}
func flattenMapFirmwareNetworkShare(p models.FirmwareNetworkShare, d *schema.ResourceData) []map[string]interface{} {
	var firmwarenetworkshares []map[string]interface{}
	var ret models.FirmwareNetworkShare
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	firmwarenetworkshare := make(map[string]interface{})
	firmwarenetworkshare["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwarenetworkshare["cifs_server"] = (func(p models.FirmwareCifsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarecifsservers []map[string]interface{}
		var ret models.FirmwareCifsServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarecifsserver := make(map[string]interface{})
		firmwarecifsserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarecifsserver["class_id"] = item.GetClassId()
		firmwarecifsserver["file_location"] = item.GetFileLocation()
		firmwarecifsserver["mount_options"] = item.GetMountOptions()
		firmwarecifsserver["object_type"] = item.GetObjectType()
		firmwarecifsserver["remote_file"] = item.GetRemoteFile()
		firmwarecifsserver["remote_ip"] = item.GetRemoteIp()
		firmwarecifsserver["remote_share"] = item.GetRemoteShare()

		firmwarecifsservers = append(firmwarecifsservers, firmwarecifsserver)
		return firmwarecifsservers
	})(item.GetCifsServer(), d)
	firmwarenetworkshare["class_id"] = item.GetClassId()
	firmwarenetworkshare["http_server"] = (func(p models.FirmwareHttpServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		var ret models.FirmwareHttpServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarehttpserver := make(map[string]interface{})
		firmwarehttpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarehttpserver["class_id"] = item.GetClassId()
		firmwarehttpserver["location_link"] = item.GetLocationLink()
		firmwarehttpserver["mount_options"] = item.GetMountOptions()
		firmwarehttpserver["object_type"] = item.GetObjectType()

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.GetHttpServer(), d)
	firmwarenetworkshare["is_password_set"] = item.GetIsPasswordSet()
	firmwarenetworkshare["map_type"] = item.GetMapType()
	firmwarenetworkshare["nfs_server"] = (func(p models.FirmwareNfsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarenfsservers []map[string]interface{}
		var ret models.FirmwareNfsServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarenfsserver := make(map[string]interface{})
		firmwarenfsserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarenfsserver["class_id"] = item.GetClassId()
		firmwarenfsserver["file_location"] = item.GetFileLocation()
		firmwarenfsserver["mount_options"] = item.GetMountOptions()
		firmwarenfsserver["object_type"] = item.GetObjectType()
		firmwarenfsserver["remote_file"] = item.GetRemoteFile()
		firmwarenfsserver["remote_ip"] = item.GetRemoteIp()
		firmwarenfsserver["remote_share"] = item.GetRemoteShare()

		firmwarenfsservers = append(firmwarenfsservers, firmwarenfsserver)
		return firmwarenfsservers
	})(item.GetNfsServer(), d)
	firmwarenetworkshare["object_type"] = item.GetObjectType()
	password_x, exists := d.GetOk("network_share")
	if exists && password_x != nil {
		password_y := password_x.([]interface{})[0].(map[string]interface{})
		firmwarenetworkshare["password"] = password_y["password"]
	}
	firmwarenetworkshare["upgradeoption"] = item.GetUpgradeoption()
	firmwarenetworkshare["username"] = item.GetUsername()

	firmwarenetworkshares = append(firmwarenetworkshares, firmwarenetworkshare)
	return firmwarenetworkshares
}
func flattenMapFirmwareRunningFirmwareRelationship(p models.FirmwareRunningFirmwareRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerelationships []map[string]interface{}
	var ret models.FirmwareRunningFirmwareRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwarerunningfirmwarerelationship := make(map[string]interface{})
	firmwarerunningfirmwarerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwarerunningfirmwarerelationship["class_id"] = item.GetClassId()
	firmwarerunningfirmwarerelationship["moid"] = item.GetMoid()
	firmwarerunningfirmwarerelationship["object_type"] = item.GetObjectType()
	firmwarerunningfirmwarerelationship["selector"] = item.GetSelector()

	firmwarerunningfirmwarerelationships = append(firmwarerunningfirmwarerelationships, firmwarerunningfirmwarerelationship)
	return firmwarerunningfirmwarerelationships
}
func flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p models.FirmwareServerConfigurationUtilityDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareserverconfigurationutilitydistributablerelationships []map[string]interface{}
	var ret models.FirmwareServerConfigurationUtilityDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareserverconfigurationutilitydistributablerelationship := make(map[string]interface{})
	firmwareserverconfigurationutilitydistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareserverconfigurationutilitydistributablerelationship["class_id"] = item.GetClassId()
	firmwareserverconfigurationutilitydistributablerelationship["moid"] = item.GetMoid()
	firmwareserverconfigurationutilitydistributablerelationship["object_type"] = item.GetObjectType()
	firmwareserverconfigurationutilitydistributablerelationship["selector"] = item.GetSelector()

	firmwareserverconfigurationutilitydistributablerelationships = append(firmwareserverconfigurationutilitydistributablerelationships, firmwareserverconfigurationutilitydistributablerelationship)
	return firmwareserverconfigurationutilitydistributablerelationships
}
func flattenMapFirmwareUpgradeBaseRelationship(p models.FirmwareUpgradeBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradebaserelationships []map[string]interface{}
	var ret models.FirmwareUpgradeBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareupgradebaserelationship := make(map[string]interface{})
	firmwareupgradebaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareupgradebaserelationship["class_id"] = item.GetClassId()
	firmwareupgradebaserelationship["moid"] = item.GetMoid()
	firmwareupgradebaserelationship["object_type"] = item.GetObjectType()
	firmwareupgradebaserelationship["selector"] = item.GetSelector()

	firmwareupgradebaserelationships = append(firmwareupgradebaserelationships, firmwareupgradebaserelationship)
	return firmwareupgradebaserelationships
}
func flattenMapFirmwareUpgradeImpactStatusRelationship(p models.FirmwareUpgradeImpactStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradeimpactstatusrelationships []map[string]interface{}
	var ret models.FirmwareUpgradeImpactStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareupgradeimpactstatusrelationship := make(map[string]interface{})
	firmwareupgradeimpactstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareupgradeimpactstatusrelationship["class_id"] = item.GetClassId()
	firmwareupgradeimpactstatusrelationship["moid"] = item.GetMoid()
	firmwareupgradeimpactstatusrelationship["object_type"] = item.GetObjectType()
	firmwareupgradeimpactstatusrelationship["selector"] = item.GetSelector()

	firmwareupgradeimpactstatusrelationships = append(firmwareupgradeimpactstatusrelationships, firmwareupgradeimpactstatusrelationship)
	return firmwareupgradeimpactstatusrelationships
}
func flattenMapFirmwareUpgradeStatusRelationship(p models.FirmwareUpgradeStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradestatusrelationships []map[string]interface{}
	var ret models.FirmwareUpgradeStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareupgradestatusrelationship := make(map[string]interface{})
	firmwareupgradestatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareupgradestatusrelationship["class_id"] = item.GetClassId()
	firmwareupgradestatusrelationship["moid"] = item.GetMoid()
	firmwareupgradestatusrelationship["object_type"] = item.GetObjectType()
	firmwareupgradestatusrelationship["selector"] = item.GetSelector()

	firmwareupgradestatusrelationships = append(firmwareupgradestatusrelationships, firmwareupgradestatusrelationship)
	return firmwareupgradestatusrelationships
}
func flattenMapForecastCatalogRelationship(p models.ForecastCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastcatalogrelationships []map[string]interface{}
	var ret models.ForecastCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	forecastcatalogrelationship := make(map[string]interface{})
	forecastcatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastcatalogrelationship["class_id"] = item.GetClassId()
	forecastcatalogrelationship["moid"] = item.GetMoid()
	forecastcatalogrelationship["object_type"] = item.GetObjectType()
	forecastcatalogrelationship["selector"] = item.GetSelector()

	forecastcatalogrelationships = append(forecastcatalogrelationships, forecastcatalogrelationship)
	return forecastcatalogrelationships
}
func flattenMapForecastDefinitionRelationship(p models.ForecastDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrelationships []map[string]interface{}
	var ret models.ForecastDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	forecastdefinitionrelationship := make(map[string]interface{})
	forecastdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastdefinitionrelationship["class_id"] = item.GetClassId()
	forecastdefinitionrelationship["moid"] = item.GetMoid()
	forecastdefinitionrelationship["object_type"] = item.GetObjectType()
	forecastdefinitionrelationship["selector"] = item.GetSelector()

	forecastdefinitionrelationships = append(forecastdefinitionrelationships, forecastdefinitionrelationship)
	return forecastdefinitionrelationships
}
func flattenMapForecastInstanceRelationship(p models.ForecastInstanceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastinstancerelationships []map[string]interface{}
	var ret models.ForecastInstanceRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	forecastinstancerelationship := make(map[string]interface{})
	forecastinstancerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastinstancerelationship["class_id"] = item.GetClassId()
	forecastinstancerelationship["moid"] = item.GetMoid()
	forecastinstancerelationship["object_type"] = item.GetObjectType()
	forecastinstancerelationship["selector"] = item.GetSelector()

	forecastinstancerelationships = append(forecastinstancerelationships, forecastinstancerelationship)
	return forecastinstancerelationships
}
func flattenMapForecastModel(p models.ForecastModel, d *schema.ResourceData) []map[string]interface{} {
	var forecastmodels []map[string]interface{}
	var ret models.ForecastModel
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	forecastmodel := make(map[string]interface{})
	forecastmodel["accuracy"] = item.GetAccuracy()
	forecastmodel["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastmodel["class_id"] = item.GetClassId()
	forecastmodel["model_data"] = item.GetModelData()
	forecastmodel["model_type"] = item.GetModelType()
	forecastmodel["object_type"] = item.GetObjectType()

	forecastmodels = append(forecastmodels, forecastmodel)
	return forecastmodels
}
func flattenMapGraphicsCardRelationship(p models.GraphicsCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrelationships []map[string]interface{}
	var ret models.GraphicsCardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	graphicscardrelationship := make(map[string]interface{})
	graphicscardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	graphicscardrelationship["class_id"] = item.GetClassId()
	graphicscardrelationship["moid"] = item.GetMoid()
	graphicscardrelationship["object_type"] = item.GetObjectType()
	graphicscardrelationship["selector"] = item.GetSelector()

	graphicscardrelationships = append(graphicscardrelationships, graphicscardrelationship)
	return graphicscardrelationships
}
func flattenMapHclOperatingSystemRelationship(p models.HclOperatingSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrelationships []map[string]interface{}
	var ret models.HclOperatingSystemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hcloperatingsystemrelationship := make(map[string]interface{})
	hcloperatingsystemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hcloperatingsystemrelationship["class_id"] = item.GetClassId()
	hcloperatingsystemrelationship["moid"] = item.GetMoid()
	hcloperatingsystemrelationship["object_type"] = item.GetObjectType()
	hcloperatingsystemrelationship["selector"] = item.GetSelector()

	hcloperatingsystemrelationships = append(hcloperatingsystemrelationships, hcloperatingsystemrelationship)
	return hcloperatingsystemrelationships
}
func flattenMapHclOperatingSystemVendorRelationship(p models.HclOperatingSystemVendorRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemvendorrelationships []map[string]interface{}
	var ret models.HclOperatingSystemVendorRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hcloperatingsystemvendorrelationship := make(map[string]interface{})
	hcloperatingsystemvendorrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hcloperatingsystemvendorrelationship["class_id"] = item.GetClassId()
	hcloperatingsystemvendorrelationship["moid"] = item.GetMoid()
	hcloperatingsystemvendorrelationship["object_type"] = item.GetObjectType()
	hcloperatingsystemvendorrelationship["selector"] = item.GetSelector()

	hcloperatingsystemvendorrelationships = append(hcloperatingsystemvendorrelationships, hcloperatingsystemvendorrelationship)
	return hcloperatingsystemvendorrelationships
}
func flattenMapHyperflexAlarmSummary(p models.HyperflexAlarmSummary, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexalarmsummarys []map[string]interface{}
	var ret models.HyperflexAlarmSummary
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexalarmsummary := make(map[string]interface{})
	hyperflexalarmsummary["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexalarmsummary["class_id"] = item.GetClassId()
	hyperflexalarmsummary["critical"] = item.GetCritical()
	hyperflexalarmsummary["object_type"] = item.GetObjectType()
	hyperflexalarmsummary["warning"] = item.GetWarning()

	hyperflexalarmsummarys = append(hyperflexalarmsummarys, hyperflexalarmsummary)
	return hyperflexalarmsummarys
}
func flattenMapHyperflexAppCatalogRelationship(p models.HyperflexAppCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexappcatalogrelationships []map[string]interface{}
	var ret models.HyperflexAppCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexappcatalogrelationship := make(map[string]interface{})
	hyperflexappcatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexappcatalogrelationship["class_id"] = item.GetClassId()
	hyperflexappcatalogrelationship["moid"] = item.GetMoid()
	hyperflexappcatalogrelationship["object_type"] = item.GetObjectType()
	hyperflexappcatalogrelationship["selector"] = item.GetSelector()

	hyperflexappcatalogrelationships = append(hyperflexappcatalogrelationships, hyperflexappcatalogrelationship)
	return hyperflexappcatalogrelationships
}
func flattenMapHyperflexAppSettingConstraint(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexappsettingconstraints []map[string]interface{}
	var ret models.HyperflexAppSettingConstraint
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexappsettingconstraint := make(map[string]interface{})
	hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexappsettingconstraint["class_id"] = item.GetClassId()
	hyperflexappsettingconstraint["hxdp_version"] = item.GetHxdpVersion()
	hyperflexappsettingconstraint["hypervisor_type"] = item.GetHypervisorType()
	hyperflexappsettingconstraint["mgmt_platform"] = item.GetMgmtPlatform()
	hyperflexappsettingconstraint["object_type"] = item.GetObjectType()
	hyperflexappsettingconstraint["server_model"] = item.GetServerModel()

	hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
	return hyperflexappsettingconstraints
}
func flattenMapHyperflexAutoSupportPolicyRelationship(p models.HyperflexAutoSupportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexautosupportpolicyrelationships []map[string]interface{}
	var ret models.HyperflexAutoSupportPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexautosupportpolicyrelationship := make(map[string]interface{})
	hyperflexautosupportpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexautosupportpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexautosupportpolicyrelationship["moid"] = item.GetMoid()
	hyperflexautosupportpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexautosupportpolicyrelationship["selector"] = item.GetSelector()

	hyperflexautosupportpolicyrelationships = append(hyperflexautosupportpolicyrelationships, hyperflexautosupportpolicyrelationship)
	return hyperflexautosupportpolicyrelationships
}
func flattenMapHyperflexBackupClusterRelationship(p models.HyperflexBackupClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexbackupclusterrelationships []map[string]interface{}
	var ret models.HyperflexBackupClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexbackupclusterrelationship := make(map[string]interface{})
	hyperflexbackupclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexbackupclusterrelationship["class_id"] = item.GetClassId()
	hyperflexbackupclusterrelationship["moid"] = item.GetMoid()
	hyperflexbackupclusterrelationship["object_type"] = item.GetObjectType()
	hyperflexbackupclusterrelationship["selector"] = item.GetSelector()

	hyperflexbackupclusterrelationships = append(hyperflexbackupclusterrelationships, hyperflexbackupclusterrelationship)
	return hyperflexbackupclusterrelationships
}
func flattenMapHyperflexBaseClusterRelationship(p models.HyperflexBaseClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexbaseclusterrelationships []map[string]interface{}
	var ret models.HyperflexBaseClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexbaseclusterrelationship := make(map[string]interface{})
	hyperflexbaseclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexbaseclusterrelationship["class_id"] = item.GetClassId()
	hyperflexbaseclusterrelationship["moid"] = item.GetMoid()
	hyperflexbaseclusterrelationship["object_type"] = item.GetObjectType()
	hyperflexbaseclusterrelationship["selector"] = item.GetSelector()

	hyperflexbaseclusterrelationships = append(hyperflexbaseclusterrelationships, hyperflexbaseclusterrelationship)
	return hyperflexbaseclusterrelationships
}
func flattenMapHyperflexBondState(p models.HyperflexBondState, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexbondstates []map[string]interface{}
	var ret models.HyperflexBondState
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexbondstate := make(map[string]interface{})
	hyperflexbondstate["active_slave"] = item.GetActiveSlave()
	hyperflexbondstate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexbondstate["class_id"] = item.GetClassId()
	hyperflexbondstate["mode"] = item.GetMode()
	hyperflexbondstate["object_type"] = item.GetObjectType()
	hyperflexbondstate["slaves"] = item.GetSlaves()

	hyperflexbondstates = append(hyperflexbondstates, hyperflexbondstate)
	return hyperflexbondstates
}
func flattenMapHyperflexCiscoHypervisorManagerRelationship(p models.HyperflexCiscoHypervisorManagerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexciscohypervisormanagerrelationships []map[string]interface{}
	var ret models.HyperflexCiscoHypervisorManagerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexciscohypervisormanagerrelationship := make(map[string]interface{})
	hyperflexciscohypervisormanagerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexciscohypervisormanagerrelationship["class_id"] = item.GetClassId()
	hyperflexciscohypervisormanagerrelationship["moid"] = item.GetMoid()
	hyperflexciscohypervisormanagerrelationship["object_type"] = item.GetObjectType()
	hyperflexciscohypervisormanagerrelationship["selector"] = item.GetSelector()

	hyperflexciscohypervisormanagerrelationships = append(hyperflexciscohypervisormanagerrelationships, hyperflexciscohypervisormanagerrelationship)
	return hyperflexciscohypervisormanagerrelationships
}
func flattenMapHyperflexClusterRelationship(p models.HyperflexClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterrelationships []map[string]interface{}
	var ret models.HyperflexClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusterrelationship := make(map[string]interface{})
	hyperflexclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusterrelationship["class_id"] = item.GetClassId()
	hyperflexclusterrelationship["moid"] = item.GetMoid()
	hyperflexclusterrelationship["object_type"] = item.GetObjectType()
	hyperflexclusterrelationship["selector"] = item.GetSelector()

	hyperflexclusterrelationships = append(hyperflexclusterrelationships, hyperflexclusterrelationship)
	return hyperflexclusterrelationships
}
func flattenMapHyperflexClusterNetworkPolicyRelationship(p models.HyperflexClusterNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusternetworkpolicyrelationships []map[string]interface{}
	var ret models.HyperflexClusterNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusternetworkpolicyrelationship := make(map[string]interface{})
	hyperflexclusternetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusternetworkpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexclusternetworkpolicyrelationship["moid"] = item.GetMoid()
	hyperflexclusternetworkpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexclusternetworkpolicyrelationship["selector"] = item.GetSelector()

	hyperflexclusternetworkpolicyrelationships = append(hyperflexclusternetworkpolicyrelationships, hyperflexclusternetworkpolicyrelationship)
	return hyperflexclusternetworkpolicyrelationships
}
func flattenMapHyperflexClusterProfileRelationship(p models.HyperflexClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerelationships []map[string]interface{}
	var ret models.HyperflexClusterProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusterprofilerelationship := make(map[string]interface{})
	hyperflexclusterprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusterprofilerelationship["class_id"] = item.GetClassId()
	hyperflexclusterprofilerelationship["moid"] = item.GetMoid()
	hyperflexclusterprofilerelationship["object_type"] = item.GetObjectType()
	hyperflexclusterprofilerelationship["selector"] = item.GetSelector()

	hyperflexclusterprofilerelationships = append(hyperflexclusterprofilerelationships, hyperflexclusterprofilerelationship)
	return hyperflexclusterprofilerelationships
}
func flattenMapHyperflexClusterStoragePolicyRelationship(p models.HyperflexClusterStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterstoragepolicyrelationships []map[string]interface{}
	var ret models.HyperflexClusterStoragePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusterstoragepolicyrelationship := make(map[string]interface{})
	hyperflexclusterstoragepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusterstoragepolicyrelationship["class_id"] = item.GetClassId()
	hyperflexclusterstoragepolicyrelationship["moid"] = item.GetMoid()
	hyperflexclusterstoragepolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexclusterstoragepolicyrelationship["selector"] = item.GetSelector()

	hyperflexclusterstoragepolicyrelationships = append(hyperflexclusterstoragepolicyrelationships, hyperflexclusterstoragepolicyrelationship)
	return hyperflexclusterstoragepolicyrelationships
}
func flattenMapHyperflexConfigResultRelationship(p models.HyperflexConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultrelationships []map[string]interface{}
	var ret models.HyperflexConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexconfigresultrelationship := make(map[string]interface{})
	hyperflexconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexconfigresultrelationship["class_id"] = item.GetClassId()
	hyperflexconfigresultrelationship["moid"] = item.GetMoid()
	hyperflexconfigresultrelationship["object_type"] = item.GetObjectType()
	hyperflexconfigresultrelationship["selector"] = item.GetSelector()

	hyperflexconfigresultrelationships = append(hyperflexconfigresultrelationships, hyperflexconfigresultrelationship)
	return hyperflexconfigresultrelationships
}
func flattenMapHyperflexDataProtectionPeerRelationship(p models.HyperflexDataProtectionPeerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexdataprotectionpeerrelationships []map[string]interface{}
	var ret models.HyperflexDataProtectionPeerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexdataprotectionpeerrelationship := make(map[string]interface{})
	hyperflexdataprotectionpeerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexdataprotectionpeerrelationship["class_id"] = item.GetClassId()
	hyperflexdataprotectionpeerrelationship["moid"] = item.GetMoid()
	hyperflexdataprotectionpeerrelationship["object_type"] = item.GetObjectType()
	hyperflexdataprotectionpeerrelationship["selector"] = item.GetSelector()

	hyperflexdataprotectionpeerrelationships = append(hyperflexdataprotectionpeerrelationships, hyperflexdataprotectionpeerrelationship)
	return hyperflexdataprotectionpeerrelationships
}
func flattenMapHyperflexDiskStatus(p models.HyperflexDiskStatus, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexdiskstatuss []map[string]interface{}
	var ret models.HyperflexDiskStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexdiskstatus := make(map[string]interface{})
	hyperflexdiskstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexdiskstatus["class_id"] = item.GetClassId()
	hyperflexdiskstatus["download_percentage"] = item.GetDownloadPercentage()
	hyperflexdiskstatus["object_type"] = item.GetObjectType()
	hyperflexdiskstatus["state"] = item.GetState()
	hyperflexdiskstatus["volume_handle"] = item.GetVolumeHandle()
	hyperflexdiskstatus["volume_name"] = item.GetVolumeName()

	hyperflexdiskstatuss = append(hyperflexdiskstatuss, hyperflexdiskstatus)
	return hyperflexdiskstatuss
}
func flattenMapHyperflexEntityReference(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexentityreferences []map[string]interface{}
	var ret models.HyperflexEntityReference
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexentityreference := make(map[string]interface{})
	hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexentityreference["class_id"] = item.GetClassId()
	hyperflexentityreference["confignum"] = item.GetConfignum()
	hyperflexentityreference["id"] = item.GetId()
	hyperflexentityreference["idtype"] = item.GetIdtype()
	hyperflexentityreference["name"] = item.GetName()
	hyperflexentityreference["object_type"] = item.GetObjectType()
	hyperflexentityreference["type"] = item.GetType()

	hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
	return hyperflexentityreferences
}
func flattenMapHyperflexErrorStack(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexerrorstacks []map[string]interface{}
	var ret models.HyperflexErrorStack
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexerrorstack := make(map[string]interface{})
	hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexerrorstack["class_id"] = item.GetClassId()
	hyperflexerrorstack["message"] = item.GetMessage()
	hyperflexerrorstack["message_id"] = item.GetMessageId()
	hyperflexerrorstack["object_type"] = item.GetObjectType()

	hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
	return hyperflexerrorstacks
}
func flattenMapHyperflexExtFcStoragePolicyRelationship(p models.HyperflexExtFcStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextfcstoragepolicyrelationships []map[string]interface{}
	var ret models.HyperflexExtFcStoragePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexextfcstoragepolicyrelationship := make(map[string]interface{})
	hyperflexextfcstoragepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexextfcstoragepolicyrelationship["class_id"] = item.GetClassId()
	hyperflexextfcstoragepolicyrelationship["moid"] = item.GetMoid()
	hyperflexextfcstoragepolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexextfcstoragepolicyrelationship["selector"] = item.GetSelector()

	hyperflexextfcstoragepolicyrelationships = append(hyperflexextfcstoragepolicyrelationships, hyperflexextfcstoragepolicyrelationship)
	return hyperflexextfcstoragepolicyrelationships
}
func flattenMapHyperflexExtIscsiStoragePolicyRelationship(p models.HyperflexExtIscsiStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextiscsistoragepolicyrelationships []map[string]interface{}
	var ret models.HyperflexExtIscsiStoragePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexextiscsistoragepolicyrelationship := make(map[string]interface{})
	hyperflexextiscsistoragepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexextiscsistoragepolicyrelationship["class_id"] = item.GetClassId()
	hyperflexextiscsistoragepolicyrelationship["moid"] = item.GetMoid()
	hyperflexextiscsistoragepolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexextiscsistoragepolicyrelationship["selector"] = item.GetSelector()

	hyperflexextiscsistoragepolicyrelationships = append(hyperflexextiscsistoragepolicyrelationships, hyperflexextiscsistoragepolicyrelationship)
	return hyperflexextiscsistoragepolicyrelationships
}
func flattenMapHyperflexFeatureLimitExternalRelationship(p models.HyperflexFeatureLimitExternalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitexternalrelationships []map[string]interface{}
	var ret models.HyperflexFeatureLimitExternalRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexfeaturelimitexternalrelationship := make(map[string]interface{})
	hyperflexfeaturelimitexternalrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexfeaturelimitexternalrelationship["class_id"] = item.GetClassId()
	hyperflexfeaturelimitexternalrelationship["moid"] = item.GetMoid()
	hyperflexfeaturelimitexternalrelationship["object_type"] = item.GetObjectType()
	hyperflexfeaturelimitexternalrelationship["selector"] = item.GetSelector()

	hyperflexfeaturelimitexternalrelationships = append(hyperflexfeaturelimitexternalrelationships, hyperflexfeaturelimitexternalrelationship)
	return hyperflexfeaturelimitexternalrelationships
}
func flattenMapHyperflexFeatureLimitInternalRelationship(p models.HyperflexFeatureLimitInternalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitinternalrelationships []map[string]interface{}
	var ret models.HyperflexFeatureLimitInternalRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexfeaturelimitinternalrelationship := make(map[string]interface{})
	hyperflexfeaturelimitinternalrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexfeaturelimitinternalrelationship["class_id"] = item.GetClassId()
	hyperflexfeaturelimitinternalrelationship["moid"] = item.GetMoid()
	hyperflexfeaturelimitinternalrelationship["object_type"] = item.GetObjectType()
	hyperflexfeaturelimitinternalrelationship["selector"] = item.GetSelector()

	hyperflexfeaturelimitinternalrelationships = append(hyperflexfeaturelimitinternalrelationships, hyperflexfeaturelimitinternalrelationship)
	return hyperflexfeaturelimitinternalrelationships
}
func flattenMapHyperflexHealthRelationship(p models.HyperflexHealthRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthrelationships []map[string]interface{}
	var ret models.HyperflexHealthRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhealthrelationship := make(map[string]interface{})
	hyperflexhealthrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhealthrelationship["class_id"] = item.GetClassId()
	hyperflexhealthrelationship["moid"] = item.GetMoid()
	hyperflexhealthrelationship["object_type"] = item.GetObjectType()
	hyperflexhealthrelationship["selector"] = item.GetSelector()

	hyperflexhealthrelationships = append(hyperflexhealthrelationships, hyperflexhealthrelationship)
	return hyperflexhealthrelationships
}
func flattenMapHyperflexHealthCheckDefinitionRelationship(p models.HyperflexHealthCheckDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthcheckdefinitionrelationships []map[string]interface{}
	var ret models.HyperflexHealthCheckDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhealthcheckdefinitionrelationship := make(map[string]interface{})
	hyperflexhealthcheckdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhealthcheckdefinitionrelationship["class_id"] = item.GetClassId()
	hyperflexhealthcheckdefinitionrelationship["moid"] = item.GetMoid()
	hyperflexhealthcheckdefinitionrelationship["object_type"] = item.GetObjectType()
	hyperflexhealthcheckdefinitionrelationship["selector"] = item.GetSelector()

	hyperflexhealthcheckdefinitionrelationships = append(hyperflexhealthcheckdefinitionrelationships, hyperflexhealthcheckdefinitionrelationship)
	return hyperflexhealthcheckdefinitionrelationships
}
func flattenMapHyperflexHealthCheckScriptInfo(p models.HyperflexHealthCheckScriptInfo, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthcheckscriptinfos []map[string]interface{}
	var ret models.HyperflexHealthCheckScriptInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhealthcheckscriptinfo := make(map[string]interface{})
	hyperflexhealthcheckscriptinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhealthcheckscriptinfo["aggregate_script_name"] = item.GetAggregateScriptName()
	hyperflexhealthcheckscriptinfo["class_id"] = item.GetClassId()
	hyperflexhealthcheckscriptinfo["hyperflex_version"] = item.GetHyperflexVersion()
	hyperflexhealthcheckscriptinfo["object_type"] = item.GetObjectType()
	hyperflexhealthcheckscriptinfo["script_execute_location"] = item.GetScriptExecuteLocation()
	hyperflexhealthcheckscriptinfo["script_input"] = item.GetScriptInput()
	hyperflexhealthcheckscriptinfo["script_name"] = item.GetScriptName()
	hyperflexhealthcheckscriptinfo["nr_version"] = item.GetVersion()

	hyperflexhealthcheckscriptinfos = append(hyperflexhealthcheckscriptinfos, hyperflexhealthcheckscriptinfo)
	return hyperflexhealthcheckscriptinfos
}
func flattenMapHyperflexHxLicenseAuthorizationDetailsDt(p models.HyperflexHxLicenseAuthorizationDetailsDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxlicenseauthorizationdetailsdts []map[string]interface{}
	var ret models.HyperflexHxLicenseAuthorizationDetailsDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxlicenseauthorizationdetailsdt := make(map[string]interface{})
	hyperflexhxlicenseauthorizationdetailsdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxlicenseauthorizationdetailsdt["class_id"] = item.GetClassId()
	hyperflexhxlicenseauthorizationdetailsdt["communication_deadline_date"] = item.GetCommunicationDeadlineDate()
	hyperflexhxlicenseauthorizationdetailsdt["evaluation_period_expired_at"] = item.GetEvaluationPeriodExpiredAt()
	hyperflexhxlicenseauthorizationdetailsdt["evaluation_period_remaining"] = item.GetEvaluationPeriodRemaining()
	hyperflexhxlicenseauthorizationdetailsdt["last_communication_attempt_date"] = item.GetLastCommunicationAttemptDate()
	hyperflexhxlicenseauthorizationdetailsdt["next_communication_attempt_date"] = item.GetNextCommunicationAttemptDate()
	hyperflexhxlicenseauthorizationdetailsdt["object_type"] = item.GetObjectType()
	hyperflexhxlicenseauthorizationdetailsdt["status"] = item.GetStatus()

	hyperflexhxlicenseauthorizationdetailsdts = append(hyperflexhxlicenseauthorizationdetailsdts, hyperflexhxlicenseauthorizationdetailsdt)
	return hyperflexhxlicenseauthorizationdetailsdts
}
func flattenMapHyperflexHxNetworkAddressDt(p models.HyperflexHxNetworkAddressDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxnetworkaddressdts []map[string]interface{}
	var ret models.HyperflexHxNetworkAddressDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxnetworkaddressdt := make(map[string]interface{})
	hyperflexhxnetworkaddressdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxnetworkaddressdt["address"] = item.GetAddress()
	hyperflexhxnetworkaddressdt["class_id"] = item.GetClassId()
	hyperflexhxnetworkaddressdt["fqdn"] = item.GetFqdn()
	hyperflexhxnetworkaddressdt["ip"] = item.GetIp()
	hyperflexhxnetworkaddressdt["object_type"] = item.GetObjectType()

	hyperflexhxnetworkaddressdts = append(hyperflexhxnetworkaddressdts, hyperflexhxnetworkaddressdt)
	return hyperflexhxnetworkaddressdts
}
func flattenMapHyperflexHxPlatformDatastoreConfigDt(p models.HyperflexHxPlatformDatastoreConfigDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxplatformdatastoreconfigdts []map[string]interface{}
	var ret models.HyperflexHxPlatformDatastoreConfigDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxplatformdatastoreconfigdt := make(map[string]interface{})
	hyperflexhxplatformdatastoreconfigdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxplatformdatastoreconfigdt["class_id"] = item.GetClassId()
	hyperflexhxplatformdatastoreconfigdt["data_block_size"] = item.GetDataBlockSize()
	hyperflexhxplatformdatastoreconfigdt["name"] = item.GetName()
	hyperflexhxplatformdatastoreconfigdt["num_mirrors"] = item.GetNumMirrors()
	hyperflexhxplatformdatastoreconfigdt["num_stripes_for_large_files"] = item.GetNumStripesForLargeFiles()
	hyperflexhxplatformdatastoreconfigdt["object_type"] = item.GetObjectType()
	hyperflexhxplatformdatastoreconfigdt["provisioned_capacity"] = item.GetProvisionedCapacity()
	hyperflexhxplatformdatastoreconfigdt["system_datastore"] = item.GetSystemDatastore()

	hyperflexhxplatformdatastoreconfigdts = append(hyperflexhxplatformdatastoreconfigdts, hyperflexhxplatformdatastoreconfigdt)
	return hyperflexhxplatformdatastoreconfigdts
}
func flattenMapHyperflexHxRegistrationDetailsDt(p models.HyperflexHxRegistrationDetailsDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxregistrationdetailsdts []map[string]interface{}
	var ret models.HyperflexHxRegistrationDetailsDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxregistrationdetailsdt := make(map[string]interface{})
	hyperflexhxregistrationdetailsdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxregistrationdetailsdt["class_id"] = item.GetClassId()
	hyperflexhxregistrationdetailsdt["initial_registration_date"] = item.GetInitialRegistrationDate()
	hyperflexhxregistrationdetailsdt["last_renewal_attempt_date"] = item.GetLastRenewalAttemptDate()
	hyperflexhxregistrationdetailsdt["next_renewal_attempt_date"] = item.GetNextRenewalAttemptDate()
	hyperflexhxregistrationdetailsdt["object_type"] = item.GetObjectType()
	hyperflexhxregistrationdetailsdt["out_of_compliance_start_date"] = item.GetOutOfComplianceStartDate()
	hyperflexhxregistrationdetailsdt["registration_expire_date"] = item.GetRegistrationExpireDate()
	hyperflexhxregistrationdetailsdt["registration_failed_reason"] = item.GetRegistrationFailedReason()
	hyperflexhxregistrationdetailsdt["smart_account"] = item.GetSmartAccount()
	hyperflexhxregistrationdetailsdt["status"] = item.GetStatus()
	hyperflexhxregistrationdetailsdt["virtual_account"] = item.GetVirtualAccount()

	hyperflexhxregistrationdetailsdts = append(hyperflexhxregistrationdetailsdts, hyperflexhxregistrationdetailsdt)
	return hyperflexhxregistrationdetailsdts
}
func flattenMapHyperflexHxResiliencyInfoDt(p models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxresiliencyinfodts []map[string]interface{}
	var ret models.HyperflexHxResiliencyInfoDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxresiliencyinfodt := make(map[string]interface{})
	hyperflexhxresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxresiliencyinfodt["class_id"] = item.GetClassId()
	hyperflexhxresiliencyinfodt["data_replication_factor"] = item.GetDataReplicationFactor()
	hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.GetHddFailuresTolerable()
	hyperflexhxresiliencyinfodt["messages"] = item.GetMessages()
	hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.GetNodeFailuresTolerable()
	hyperflexhxresiliencyinfodt["object_type"] = item.GetObjectType()
	hyperflexhxresiliencyinfodt["policy_compliance"] = item.GetPolicyCompliance()
	hyperflexhxresiliencyinfodt["resiliency_state"] = item.GetResiliencyState()
	hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.GetSsdFailuresTolerable()

	hyperflexhxresiliencyinfodts = append(hyperflexhxresiliencyinfodts, hyperflexhxresiliencyinfodt)
	return hyperflexhxresiliencyinfodts
}
func flattenMapHyperflexHxSiteDt(p models.HyperflexHxSiteDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxsitedts []map[string]interface{}
	var ret models.HyperflexHxSiteDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxsitedt := make(map[string]interface{})
	hyperflexhxsitedt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxsitedt["class_id"] = item.GetClassId()
	hyperflexhxsitedt["name"] = item.GetName()
	hyperflexhxsitedt["object_type"] = item.GetObjectType()
	hyperflexhxsitedt["zone"] = (func(p models.HyperflexHxZoneInfoDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxzoneinfodts []map[string]interface{}
		var ret models.HyperflexHxZoneInfoDt
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexhxzoneinfodt := make(map[string]interface{})
		hyperflexhxzoneinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhxzoneinfodt["class_id"] = item.GetClassId()
		hyperflexhxzoneinfodt["num_nodes"] = item.GetNumNodes()
		hyperflexhxzoneinfodt["object_type"] = item.GetObjectType()
		hyperflexhxzoneinfodt["uuid"] = item.GetUuid()

		hyperflexhxzoneinfodts = append(hyperflexhxzoneinfodts, hyperflexhxzoneinfodt)
		return hyperflexhxzoneinfodts
	})(item.GetZone(), d)

	hyperflexhxsitedts = append(hyperflexhxsitedts, hyperflexhxsitedt)
	return hyperflexhxsitedts
}
func flattenMapHyperflexHxUuIdDt(p models.HyperflexHxUuIdDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxuuiddts []map[string]interface{}
	var ret models.HyperflexHxUuIdDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxuuiddt := make(map[string]interface{})
	hyperflexhxuuiddt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxuuiddt["class_id"] = item.GetClassId()
	hyperflexhxuuiddt["links"] = (func(p []models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxlinkdts []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			hyperflexhxlinkdt := make(map[string]interface{})
			hyperflexhxlinkdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexhxlinkdt["class_id"] = item.GetClassId()
			hyperflexhxlinkdt["comments"] = item.GetComments()
			hyperflexhxlinkdt["href"] = item.GetHref()
			hyperflexhxlinkdt["method"] = item.GetMethod()
			hyperflexhxlinkdt["object_type"] = item.GetObjectType()
			hyperflexhxlinkdt["rel"] = item.GetRel()
			hyperflexhxlinkdts = append(hyperflexhxlinkdts, hyperflexhxlinkdt)
		}
		return hyperflexhxlinkdts
	})(item.GetLinks(), d)
	hyperflexhxuuiddt["object_type"] = item.GetObjectType()
	hyperflexhxuuiddt["uuid"] = item.GetUuid()

	hyperflexhxuuiddts = append(hyperflexhxuuiddts, hyperflexhxuuiddt)
	return hyperflexhxuuiddts
}
func flattenMapHyperflexHxapClusterRelationship(p models.HyperflexHxapClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxapclusterrelationships []map[string]interface{}
	var ret models.HyperflexHxapClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhxapclusterrelationship := make(map[string]interface{})
	hyperflexhxapclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxapclusterrelationship["class_id"] = item.GetClassId()
	hyperflexhxapclusterrelationship["moid"] = item.GetMoid()
	hyperflexhxapclusterrelationship["object_type"] = item.GetObjectType()
	hyperflexhxapclusterrelationship["selector"] = item.GetSelector()

	hyperflexhxapclusterrelationships = append(hyperflexhxapclusterrelationships, hyperflexhxapclusterrelationship)
	return hyperflexhxapclusterrelationships
}
func flattenMapHyperflexHxapDvUplinkRelationship(p models.HyperflexHxapDvUplinkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxapdvuplinkrelationships []map[string]interface{}
	var ret models.HyperflexHxapDvUplinkRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhxapdvuplinkrelationship := make(map[string]interface{})
	hyperflexhxapdvuplinkrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxapdvuplinkrelationship["class_id"] = item.GetClassId()
	hyperflexhxapdvuplinkrelationship["moid"] = item.GetMoid()
	hyperflexhxapdvuplinkrelationship["object_type"] = item.GetObjectType()
	hyperflexhxapdvuplinkrelationship["selector"] = item.GetSelector()

	hyperflexhxapdvuplinkrelationships = append(hyperflexhxapdvuplinkrelationships, hyperflexhxapdvuplinkrelationship)
	return hyperflexhxapdvuplinkrelationships
}
func flattenMapHyperflexHxapDvswitchRelationship(p models.HyperflexHxapDvswitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxapdvswitchrelationships []map[string]interface{}
	var ret models.HyperflexHxapDvswitchRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhxapdvswitchrelationship := make(map[string]interface{})
	hyperflexhxapdvswitchrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxapdvswitchrelationship["class_id"] = item.GetClassId()
	hyperflexhxapdvswitchrelationship["moid"] = item.GetMoid()
	hyperflexhxapdvswitchrelationship["object_type"] = item.GetObjectType()
	hyperflexhxapdvswitchrelationship["selector"] = item.GetSelector()

	hyperflexhxapdvswitchrelationships = append(hyperflexhxapdvswitchrelationships, hyperflexhxapdvswitchrelationship)
	return hyperflexhxapdvswitchrelationships
}
func flattenMapHyperflexHxapHostRelationship(p models.HyperflexHxapHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxaphostrelationships []map[string]interface{}
	var ret models.HyperflexHxapHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhxaphostrelationship := make(map[string]interface{})
	hyperflexhxaphostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxaphostrelationship["class_id"] = item.GetClassId()
	hyperflexhxaphostrelationship["moid"] = item.GetMoid()
	hyperflexhxaphostrelationship["object_type"] = item.GetObjectType()
	hyperflexhxaphostrelationship["selector"] = item.GetSelector()

	hyperflexhxaphostrelationships = append(hyperflexhxaphostrelationships, hyperflexhxaphostrelationship)
	return hyperflexhxaphostrelationships
}
func flattenMapHyperflexHxapNetworkRelationship(p models.HyperflexHxapNetworkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxapnetworkrelationships []map[string]interface{}
	var ret models.HyperflexHxapNetworkRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhxapnetworkrelationship := make(map[string]interface{})
	hyperflexhxapnetworkrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxapnetworkrelationship["class_id"] = item.GetClassId()
	hyperflexhxapnetworkrelationship["moid"] = item.GetMoid()
	hyperflexhxapnetworkrelationship["object_type"] = item.GetObjectType()
	hyperflexhxapnetworkrelationship["selector"] = item.GetSelector()

	hyperflexhxapnetworkrelationships = append(hyperflexhxapnetworkrelationships, hyperflexhxapnetworkrelationship)
	return hyperflexhxapnetworkrelationships
}
func flattenMapHyperflexHxapVirtualMachineRelationship(p models.HyperflexHxapVirtualMachineRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxapvirtualmachinerelationships []map[string]interface{}
	var ret models.HyperflexHxapVirtualMachineRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhxapvirtualmachinerelationship := make(map[string]interface{})
	hyperflexhxapvirtualmachinerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxapvirtualmachinerelationship["class_id"] = item.GetClassId()
	hyperflexhxapvirtualmachinerelationship["moid"] = item.GetMoid()
	hyperflexhxapvirtualmachinerelationship["object_type"] = item.GetObjectType()
	hyperflexhxapvirtualmachinerelationship["selector"] = item.GetSelector()

	hyperflexhxapvirtualmachinerelationships = append(hyperflexhxapvirtualmachinerelationships, hyperflexhxapvirtualmachinerelationship)
	return hyperflexhxapvirtualmachinerelationships
}
func flattenMapHyperflexIpAddrRange(p models.HyperflexIpAddrRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexipaddrranges []map[string]interface{}
	var ret models.HyperflexIpAddrRange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexipaddrrange := make(map[string]interface{})
	hyperflexipaddrrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexipaddrrange["class_id"] = item.GetClassId()
	hyperflexipaddrrange["end_addr"] = item.GetEndAddr()
	hyperflexipaddrrange["gateway"] = item.GetGateway()
	hyperflexipaddrrange["netmask"] = item.GetNetmask()
	hyperflexipaddrrange["object_type"] = item.GetObjectType()
	hyperflexipaddrrange["start_addr"] = item.GetStartAddr()

	hyperflexipaddrranges = append(hyperflexipaddrranges, hyperflexipaddrrange)
	return hyperflexipaddrranges
}
func flattenMapHyperflexLicenseRelationship(p models.HyperflexLicenseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlicenserelationships []map[string]interface{}
	var ret models.HyperflexLicenseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexlicenserelationship := make(map[string]interface{})
	hyperflexlicenserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexlicenserelationship["class_id"] = item.GetClassId()
	hyperflexlicenserelationship["moid"] = item.GetMoid()
	hyperflexlicenserelationship["object_type"] = item.GetObjectType()
	hyperflexlicenserelationship["selector"] = item.GetSelector()

	hyperflexlicenserelationships = append(hyperflexlicenserelationships, hyperflexlicenserelationship)
	return hyperflexlicenserelationships
}
func flattenMapHyperflexLocalCredentialPolicyRelationship(p models.HyperflexLocalCredentialPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlocalcredentialpolicyrelationships []map[string]interface{}
	var ret models.HyperflexLocalCredentialPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexlocalcredentialpolicyrelationship := make(map[string]interface{})
	hyperflexlocalcredentialpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexlocalcredentialpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexlocalcredentialpolicyrelationship["moid"] = item.GetMoid()
	hyperflexlocalcredentialpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexlocalcredentialpolicyrelationship["selector"] = item.GetSelector()

	hyperflexlocalcredentialpolicyrelationships = append(hyperflexlocalcredentialpolicyrelationships, hyperflexlocalcredentialpolicyrelationship)
	return hyperflexlocalcredentialpolicyrelationships
}
func flattenMapHyperflexLogicalAvailabilityZone(p models.HyperflexLogicalAvailabilityZone, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlogicalavailabilityzones []map[string]interface{}
	var ret models.HyperflexLogicalAvailabilityZone
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexlogicalavailabilityzone := make(map[string]interface{})
	hyperflexlogicalavailabilityzone["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexlogicalavailabilityzone["auto_config"] = item.GetAutoConfig()
	hyperflexlogicalavailabilityzone["class_id"] = item.GetClassId()
	hyperflexlogicalavailabilityzone["object_type"] = item.GetObjectType()

	hyperflexlogicalavailabilityzones = append(hyperflexlogicalavailabilityzones, hyperflexlogicalavailabilityzone)
	return hyperflexlogicalavailabilityzones
}
func flattenMapHyperflexMacAddrPrefixRange(p models.HyperflexMacAddrPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexmacaddrprefixranges []map[string]interface{}
	var ret models.HyperflexMacAddrPrefixRange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexmacaddrprefixrange := make(map[string]interface{})
	hyperflexmacaddrprefixrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexmacaddrprefixrange["class_id"] = item.GetClassId()
	hyperflexmacaddrprefixrange["end_addr"] = item.GetEndAddr()
	hyperflexmacaddrprefixrange["object_type"] = item.GetObjectType()
	hyperflexmacaddrprefixrange["start_addr"] = item.GetStartAddr()

	hyperflexmacaddrprefixranges = append(hyperflexmacaddrprefixranges, hyperflexmacaddrprefixrange)
	return hyperflexmacaddrprefixranges
}
func flattenMapHyperflexNamedVlan(p models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	var ret models.HyperflexNamedVlan
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexnamedvlan := make(map[string]interface{})
	hyperflexnamedvlan["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexnamedvlan["class_id"] = item.GetClassId()
	hyperflexnamedvlan["name"] = item.GetName()
	hyperflexnamedvlan["object_type"] = item.GetObjectType()
	hyperflexnamedvlan["vlan_id"] = item.GetVlanId()

	hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	return hyperflexnamedvlans
}
func flattenMapHyperflexNamedVsan(p models.HyperflexNamedVsan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvsans []map[string]interface{}
	var ret models.HyperflexNamedVsan
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexnamedvsan := make(map[string]interface{})
	hyperflexnamedvsan["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexnamedvsan["class_id"] = item.GetClassId()
	hyperflexnamedvsan["name"] = item.GetName()
	hyperflexnamedvsan["object_type"] = item.GetObjectType()
	hyperflexnamedvsan["vsan_id"] = item.GetVsanId()

	hyperflexnamedvsans = append(hyperflexnamedvsans, hyperflexnamedvsan)
	return hyperflexnamedvsans
}
func flattenMapHyperflexNodeConfigPolicyRelationship(p models.HyperflexNodeConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexNodeConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexnodeconfigpolicyrelationship := make(map[string]interface{})
	hyperflexnodeconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexnodeconfigpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexnodeconfigpolicyrelationship["moid"] = item.GetMoid()
	hyperflexnodeconfigpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexnodeconfigpolicyrelationship["selector"] = item.GetSelector()

	hyperflexnodeconfigpolicyrelationships = append(hyperflexnodeconfigpolicyrelationships, hyperflexnodeconfigpolicyrelationship)
	return hyperflexnodeconfigpolicyrelationships
}
func flattenMapHyperflexProxySettingPolicyRelationship(p models.HyperflexProxySettingPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexproxysettingpolicyrelationships []map[string]interface{}
	var ret models.HyperflexProxySettingPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexproxysettingpolicyrelationship := make(map[string]interface{})
	hyperflexproxysettingpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexproxysettingpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexproxysettingpolicyrelationship["moid"] = item.GetMoid()
	hyperflexproxysettingpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexproxysettingpolicyrelationship["selector"] = item.GetSelector()

	hyperflexproxysettingpolicyrelationships = append(hyperflexproxysettingpolicyrelationships, hyperflexproxysettingpolicyrelationship)
	return hyperflexproxysettingpolicyrelationships
}
func flattenMapHyperflexReplicationPeerInfo(p models.HyperflexReplicationPeerInfo, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexreplicationpeerinfos []map[string]interface{}
	var ret models.HyperflexReplicationPeerInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexreplicationpeerinfo := make(map[string]interface{})
	hyperflexreplicationpeerinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexreplicationpeerinfo["class_id"] = item.GetClassId()
	hyperflexreplicationpeerinfo["datastores"] = (func(p []models.HyperflexReplicationPlatDatastorePair, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexreplicationplatdatastorepairs []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			hyperflexreplicationplatdatastorepair := make(map[string]interface{})
			hyperflexreplicationplatdatastorepair["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexreplicationplatdatastorepair["ads"] = (func(p models.HyperflexReplicationPlatDatastore, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationplatdatastores []map[string]interface{}
				var ret models.HyperflexReplicationPlatDatastore
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationplatdatastore := make(map[string]interface{})
				hyperflexreplicationplatdatastore["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationplatdatastore["class_id"] = item.GetClassId()
				hyperflexreplicationplatdatastore["cluster_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetClusterEr(), d)
				hyperflexreplicationplatdatastore["datastore_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetDatastoreEr(), d)
				hyperflexreplicationplatdatastore["object_type"] = item.GetObjectType()

				hyperflexreplicationplatdatastores = append(hyperflexreplicationplatdatastores, hyperflexreplicationplatdatastore)
				return hyperflexreplicationplatdatastores
			})(item.GetAds(), d)
			hyperflexreplicationplatdatastorepair["backup_only"] = item.GetBackupOnly()
			hyperflexreplicationplatdatastorepair["bds"] = (func(p models.HyperflexReplicationPlatDatastore, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationplatdatastores []map[string]interface{}
				var ret models.HyperflexReplicationPlatDatastore
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationplatdatastore := make(map[string]interface{})
				hyperflexreplicationplatdatastore["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationplatdatastore["class_id"] = item.GetClassId()
				hyperflexreplicationplatdatastore["cluster_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetClusterEr(), d)
				hyperflexreplicationplatdatastore["datastore_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetDatastoreEr(), d)
				hyperflexreplicationplatdatastore["object_type"] = item.GetObjectType()

				hyperflexreplicationplatdatastores = append(hyperflexreplicationplatdatastores, hyperflexreplicationplatdatastore)
				return hyperflexreplicationplatdatastores
			})(item.GetBds(), d)
			hyperflexreplicationplatdatastorepair["class_id"] = item.GetClassId()
			hyperflexreplicationplatdatastorepair["object_type"] = item.GetObjectType()
			hyperflexreplicationplatdatastorepair["quiesce"] = item.GetQuiesce()
			hyperflexreplicationplatdatastorepair["replication_interval_in_minutes"] = item.GetReplicationIntervalInMinutes()
			hyperflexreplicationplatdatastorepair["sourceds"] = (func(p models.HyperflexReplicationPlatDatastore, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationplatdatastores []map[string]interface{}
				var ret models.HyperflexReplicationPlatDatastore
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationplatdatastore := make(map[string]interface{})
				hyperflexreplicationplatdatastore["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationplatdatastore["class_id"] = item.GetClassId()
				hyperflexreplicationplatdatastore["cluster_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetClusterEr(), d)
				hyperflexreplicationplatdatastore["datastore_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.GetClassId()
					hyperflexentityreference["confignum"] = item.GetConfignum()
					hyperflexentityreference["id"] = item.GetId()
					hyperflexentityreference["idtype"] = item.GetIdtype()
					hyperflexentityreference["name"] = item.GetName()
					hyperflexentityreference["object_type"] = item.GetObjectType()
					hyperflexentityreference["type"] = item.GetType()

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetDatastoreEr(), d)
				hyperflexreplicationplatdatastore["object_type"] = item.GetObjectType()

				hyperflexreplicationplatdatastores = append(hyperflexreplicationplatdatastores, hyperflexreplicationplatdatastore)
				return hyperflexreplicationplatdatastores
			})(item.GetSourceds(), d)
			hyperflexreplicationplatdatastorepair["storage_only"] = item.GetStorageOnly()
			hyperflexreplicationplatdatastorepairs = append(hyperflexreplicationplatdatastorepairs, hyperflexreplicationplatdatastorepair)
		}
		return hyperflexreplicationplatdatastorepairs
	})(item.GetDatastores(), d)
	hyperflexreplicationpeerinfo["dcip"] = item.GetDcip()
	hyperflexreplicationpeerinfo["er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexentityreferences []map[string]interface{}
		var ret models.HyperflexEntityReference
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexentityreference := make(map[string]interface{})
		hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexentityreference["class_id"] = item.GetClassId()
		hyperflexentityreference["confignum"] = item.GetConfignum()
		hyperflexentityreference["id"] = item.GetId()
		hyperflexentityreference["idtype"] = item.GetIdtype()
		hyperflexentityreference["name"] = item.GetName()
		hyperflexentityreference["object_type"] = item.GetObjectType()
		hyperflexentityreference["type"] = item.GetType()

		hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
		return hyperflexentityreferences
	})(item.GetEr(), d)
	hyperflexreplicationpeerinfo["mcip"] = item.GetMcip()
	hyperflexreplicationpeerinfo["object_type"] = item.GetObjectType()
	hyperflexreplicationpeerinfo["ports"] = (func(p []models.HyperflexPortTypeToPortNumberMap, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexporttypetoportnumbermaps []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			hyperflexporttypetoportnumbermap := make(map[string]interface{})
			hyperflexporttypetoportnumbermap["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexporttypetoportnumbermap["class_id"] = item.GetClassId()
			hyperflexporttypetoportnumbermap["i16"] = item.GetI16()
			hyperflexporttypetoportnumbermap["object_type"] = item.GetObjectType()
			hyperflexporttypetoportnumbermap["string"] = item.GetString()
			hyperflexporttypetoportnumbermaps = append(hyperflexporttypetoportnumbermaps, hyperflexporttypetoportnumbermap)
		}
		return hyperflexporttypetoportnumbermaps
	})(item.GetPorts(), d)
	hyperflexreplicationpeerinfo["repl_cip"] = item.GetReplCip()
	hyperflexreplicationpeerinfo["status"] = item.GetStatus()
	hyperflexreplicationpeerinfo["status_details"] = item.GetStatusDetails()

	hyperflexreplicationpeerinfos = append(hyperflexreplicationpeerinfos, hyperflexreplicationpeerinfo)
	return hyperflexreplicationpeerinfos
}
func flattenMapHyperflexReplicationSchedule(p models.HyperflexReplicationSchedule, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexreplicationschedules []map[string]interface{}
	var ret models.HyperflexReplicationSchedule
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexreplicationschedule := make(map[string]interface{})
	hyperflexreplicationschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexreplicationschedule["backup_interval"] = item.GetBackupInterval()
	hyperflexreplicationschedule["class_id"] = item.GetClassId()
	hyperflexreplicationschedule["object_type"] = item.GetObjectType()

	hyperflexreplicationschedules = append(hyperflexreplicationschedules, hyperflexreplicationschedule)
	return hyperflexreplicationschedules
}
func flattenMapHyperflexServerFirmwareVersionRelationship(p models.HyperflexServerFirmwareVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionrelationships []map[string]interface{}
	var ret models.HyperflexServerFirmwareVersionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexserverfirmwareversionrelationship := make(map[string]interface{})
	hyperflexserverfirmwareversionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexserverfirmwareversionrelationship["class_id"] = item.GetClassId()
	hyperflexserverfirmwareversionrelationship["moid"] = item.GetMoid()
	hyperflexserverfirmwareversionrelationship["object_type"] = item.GetObjectType()
	hyperflexserverfirmwareversionrelationship["selector"] = item.GetSelector()

	hyperflexserverfirmwareversionrelationships = append(hyperflexserverfirmwareversionrelationships, hyperflexserverfirmwareversionrelationship)
	return hyperflexserverfirmwareversionrelationships
}
func flattenMapHyperflexServerModelRelationship(p models.HyperflexServerModelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelrelationships []map[string]interface{}
	var ret models.HyperflexServerModelRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexservermodelrelationship := make(map[string]interface{})
	hyperflexservermodelrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexservermodelrelationship["class_id"] = item.GetClassId()
	hyperflexservermodelrelationship["moid"] = item.GetMoid()
	hyperflexservermodelrelationship["object_type"] = item.GetObjectType()
	hyperflexservermodelrelationship["selector"] = item.GetSelector()

	hyperflexservermodelrelationships = append(hyperflexservermodelrelationships, hyperflexservermodelrelationship)
	return hyperflexservermodelrelationships
}
func flattenMapHyperflexSoftwareDistributionEntryRelationship(p models.HyperflexSoftwareDistributionEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwaredistributionentryrelationships []map[string]interface{}
	var ret models.HyperflexSoftwareDistributionEntryRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexsoftwaredistributionentryrelationship := make(map[string]interface{})
	hyperflexsoftwaredistributionentryrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsoftwaredistributionentryrelationship["class_id"] = item.GetClassId()
	hyperflexsoftwaredistributionentryrelationship["moid"] = item.GetMoid()
	hyperflexsoftwaredistributionentryrelationship["object_type"] = item.GetObjectType()
	hyperflexsoftwaredistributionentryrelationship["selector"] = item.GetSelector()

	hyperflexsoftwaredistributionentryrelationships = append(hyperflexsoftwaredistributionentryrelationships, hyperflexsoftwaredistributionentryrelationship)
	return hyperflexsoftwaredistributionentryrelationships
}
func flattenMapHyperflexSoftwareDistributionVersionRelationship(p models.HyperflexSoftwareDistributionVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwaredistributionversionrelationships []map[string]interface{}
	var ret models.HyperflexSoftwareDistributionVersionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexsoftwaredistributionversionrelationship := make(map[string]interface{})
	hyperflexsoftwaredistributionversionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsoftwaredistributionversionrelationship["class_id"] = item.GetClassId()
	hyperflexsoftwaredistributionversionrelationship["moid"] = item.GetMoid()
	hyperflexsoftwaredistributionversionrelationship["object_type"] = item.GetObjectType()
	hyperflexsoftwaredistributionversionrelationship["selector"] = item.GetSelector()

	hyperflexsoftwaredistributionversionrelationships = append(hyperflexsoftwaredistributionversionrelationships, hyperflexsoftwaredistributionversionrelationship)
	return hyperflexsoftwaredistributionversionrelationships
}
func flattenMapHyperflexSoftwareVersionPolicyRelationship(p models.HyperflexSoftwareVersionPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwareversionpolicyrelationships []map[string]interface{}
	var ret models.HyperflexSoftwareVersionPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexsoftwareversionpolicyrelationship := make(map[string]interface{})
	hyperflexsoftwareversionpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsoftwareversionpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexsoftwareversionpolicyrelationship["moid"] = item.GetMoid()
	hyperflexsoftwareversionpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexsoftwareversionpolicyrelationship["selector"] = item.GetSelector()

	hyperflexsoftwareversionpolicyrelationships = append(hyperflexsoftwareversionpolicyrelationships, hyperflexsoftwareversionpolicyrelationship)
	return hyperflexsoftwareversionpolicyrelationships
}
func flattenMapHyperflexStorageContainerRelationship(p models.HyperflexStorageContainerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexstoragecontainerrelationships []map[string]interface{}
	var ret models.HyperflexStorageContainerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexstoragecontainerrelationship := make(map[string]interface{})
	hyperflexstoragecontainerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexstoragecontainerrelationship["class_id"] = item.GetClassId()
	hyperflexstoragecontainerrelationship["moid"] = item.GetMoid()
	hyperflexstoragecontainerrelationship["object_type"] = item.GetObjectType()
	hyperflexstoragecontainerrelationship["selector"] = item.GetSelector()

	hyperflexstoragecontainerrelationships = append(hyperflexstoragecontainerrelationships, hyperflexstoragecontainerrelationship)
	return hyperflexstoragecontainerrelationships
}
func flattenMapHyperflexSummary(p models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsummarys []map[string]interface{}
	var ret models.HyperflexSummary
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexsummary := make(map[string]interface{})
	hyperflexsummary["active_nodes"] = item.GetActiveNodes()
	hyperflexsummary["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsummary["address"] = item.GetAddress()
	hyperflexsummary["boottime"] = item.GetBoottime()
	hyperflexsummary["class_id"] = item.GetClassId()
	hyperflexsummary["cluster_access_policy"] = item.GetClusterAccessPolicy()
	hyperflexsummary["compression_savings"] = item.GetCompressionSavings()
	hyperflexsummary["data_replication_compliance"] = item.GetDataReplicationCompliance()
	hyperflexsummary["data_replication_factor"] = item.GetDataReplicationFactor()
	hyperflexsummary["deduplication_savings"] = item.GetDeduplicationSavings()
	hyperflexsummary["downtime"] = item.GetDowntime()
	hyperflexsummary["free_capacity"] = item.GetFreeCapacity()
	hyperflexsummary["healing_info"] = (func(p models.HyperflexStPlatformClusterHealingInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterhealinginfos []map[string]interface{}
		var ret models.HyperflexStPlatformClusterHealingInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexstplatformclusterhealinginfo := make(map[string]interface{})
		hyperflexstplatformclusterhealinginfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexstplatformclusterhealinginfo["class_id"] = item.GetClassId()
		hyperflexstplatformclusterhealinginfo["estimated_completion_time_in_seconds"] = item.GetEstimatedCompletionTimeInSeconds()
		hyperflexstplatformclusterhealinginfo["in_progress"] = item.GetInProgress()
		hyperflexstplatformclusterhealinginfo["messages"] = item.GetMessages()
		hyperflexstplatformclusterhealinginfo["messages_iterator"] = flattenAdditionalProperties(item.MessagesIterator)
		hyperflexstplatformclusterhealinginfo["messages_size"] = item.GetMessagesSize()
		hyperflexstplatformclusterhealinginfo["object_type"] = item.GetObjectType()
		hyperflexstplatformclusterhealinginfo["percent_complete"] = item.GetPercentComplete()

		hyperflexstplatformclusterhealinginfos = append(hyperflexstplatformclusterhealinginfos, hyperflexstplatformclusterhealinginfo)
		return hyperflexstplatformclusterhealinginfos
	})(item.GetHealingInfo(), d)
	hyperflexsummary["name"] = item.GetName()
	hyperflexsummary["object_type"] = item.GetObjectType()
	hyperflexsummary["resiliency_details"] = flattenAdditionalProperties(item.ResiliencyDetails)
	hyperflexsummary["resiliency_details_size"] = item.GetResiliencyDetailsSize()
	hyperflexsummary["resiliency_info"] = (func(p models.HyperflexStPlatformClusterResiliencyInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterresiliencyinfos []map[string]interface{}
		var ret models.HyperflexStPlatformClusterResiliencyInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexstplatformclusterresiliencyinfo := make(map[string]interface{})
		hyperflexstplatformclusterresiliencyinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexstplatformclusterresiliencyinfo["class_id"] = item.GetClassId()
		hyperflexstplatformclusterresiliencyinfo["hdd_failures_tolerable"] = item.GetHddFailuresTolerable()
		hyperflexstplatformclusterresiliencyinfo["messages"] = item.GetMessages()
		hyperflexstplatformclusterresiliencyinfo["messages_iterator"] = flattenAdditionalProperties(item.MessagesIterator)
		hyperflexstplatformclusterresiliencyinfo["messages_size"] = item.GetMessagesSize()
		hyperflexstplatformclusterresiliencyinfo["node_failures_tolerable"] = item.GetNodeFailuresTolerable()
		hyperflexstplatformclusterresiliencyinfo["object_type"] = item.GetObjectType()
		hyperflexstplatformclusterresiliencyinfo["ssd_failures_tolerable"] = item.GetSsdFailuresTolerable()
		hyperflexstplatformclusterresiliencyinfo["state"] = item.GetState()

		hyperflexstplatformclusterresiliencyinfos = append(hyperflexstplatformclusterresiliencyinfos, hyperflexstplatformclusterresiliencyinfo)
		return hyperflexstplatformclusterresiliencyinfos
	})(item.GetResiliencyInfo(), d)
	hyperflexsummary["space_status"] = item.GetSpaceStatus()
	hyperflexsummary["state"] = item.GetState()
	hyperflexsummary["total_capacity"] = item.GetTotalCapacity()
	hyperflexsummary["total_savings"] = item.GetTotalSavings()
	hyperflexsummary["uptime"] = item.GetUptime()
	hyperflexsummary["used_capacity"] = item.GetUsedCapacity()

	hyperflexsummarys = append(hyperflexsummarys, hyperflexsummary)
	return hyperflexsummarys
}
func flattenMapHyperflexSysConfigPolicyRelationship(p models.HyperflexSysConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsysconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexSysConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexsysconfigpolicyrelationship := make(map[string]interface{})
	hyperflexsysconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsysconfigpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexsysconfigpolicyrelationship["moid"] = item.GetMoid()
	hyperflexsysconfigpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexsysconfigpolicyrelationship["selector"] = item.GetSelector()

	hyperflexsysconfigpolicyrelationships = append(hyperflexsysconfigpolicyrelationships, hyperflexsysconfigpolicyrelationship)
	return hyperflexsysconfigpolicyrelationships
}
func flattenMapHyperflexUcsmConfigPolicyRelationship(p models.HyperflexUcsmConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexucsmconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexUcsmConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexucsmconfigpolicyrelationship := make(map[string]interface{})
	hyperflexucsmconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexucsmconfigpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexucsmconfigpolicyrelationship["moid"] = item.GetMoid()
	hyperflexucsmconfigpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexucsmconfigpolicyrelationship["selector"] = item.GetSelector()

	hyperflexucsmconfigpolicyrelationships = append(hyperflexucsmconfigpolicyrelationships, hyperflexucsmconfigpolicyrelationship)
	return hyperflexucsmconfigpolicyrelationships
}
func flattenMapHyperflexVcenterConfigPolicyRelationship(p models.HyperflexVcenterConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvcenterconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexVcenterConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexvcenterconfigpolicyrelationship := make(map[string]interface{})
	hyperflexvcenterconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexvcenterconfigpolicyrelationship["class_id"] = item.GetClassId()
	hyperflexvcenterconfigpolicyrelationship["moid"] = item.GetMoid()
	hyperflexvcenterconfigpolicyrelationship["object_type"] = item.GetObjectType()
	hyperflexvcenterconfigpolicyrelationship["selector"] = item.GetSelector()

	hyperflexvcenterconfigpolicyrelationships = append(hyperflexvcenterconfigpolicyrelationships, hyperflexvcenterconfigpolicyrelationship)
	return hyperflexvcenterconfigpolicyrelationships
}
func flattenMapHyperflexVirtualMachine(p models.HyperflexVirtualMachine, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvirtualmachines []map[string]interface{}
	var ret models.HyperflexVirtualMachine
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexvirtualmachine := make(map[string]interface{})
	hyperflexvirtualmachine["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexvirtualmachine["class_id"] = item.GetClassId()
	hyperflexvirtualmachine["object_type"] = item.GetObjectType()
	hyperflexvirtualmachine["run_time_info"] = (func(p models.HyperflexVirtualMachineRuntimeInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexvirtualmachineruntimeinfos []map[string]interface{}
		var ret models.HyperflexVirtualMachineRuntimeInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexvirtualmachineruntimeinfo := make(map[string]interface{})
		hyperflexvirtualmachineruntimeinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexvirtualmachineruntimeinfo["bios_uuid"] = item.GetBiosUuid()
		hyperflexvirtualmachineruntimeinfo["class_id"] = item.GetClassId()
		hyperflexvirtualmachineruntimeinfo["connection_state"] = item.GetConnectionState()
		hyperflexvirtualmachineruntimeinfo["cpu_usage"] = item.GetCpuUsage()
		hyperflexvirtualmachineruntimeinfo["folder"] = item.GetFolder()
		hyperflexvirtualmachineruntimeinfo["guest_family"] = item.GetGuestFamily()
		hyperflexvirtualmachineruntimeinfo["guest_full_name"] = item.GetGuestFullName()
		hyperflexvirtualmachineruntimeinfo["guest_id"] = item.GetGuestId()
		hyperflexvirtualmachineruntimeinfo["guest_state"] = item.GetGuestState()
		hyperflexvirtualmachineruntimeinfo["host_name"] = item.GetHostName()
		hyperflexvirtualmachineruntimeinfo["instance_uuid"] = item.GetInstanceUuid()
		hyperflexvirtualmachineruntimeinfo["memory_mb"] = item.GetMemoryMb()
		hyperflexvirtualmachineruntimeinfo["memory_usage"] = item.GetMemoryUsage()
		hyperflexvirtualmachineruntimeinfo["moid"] = item.GetMoid()
		hyperflexvirtualmachineruntimeinfo["name"] = item.GetName()
		hyperflexvirtualmachineruntimeinfo["networks"] = item.GetNetworks()
		hyperflexvirtualmachineruntimeinfo["num_cpu"] = item.GetNumCpu()
		hyperflexvirtualmachineruntimeinfo["object_type"] = item.GetObjectType()
		hyperflexvirtualmachineruntimeinfo["power_state"] = item.GetPowerState()
		hyperflexvirtualmachineruntimeinfo["provisioned_size"] = item.GetProvisionedSize()
		hyperflexvirtualmachineruntimeinfo["resource_pool"] = item.GetResourcePool()
		hyperflexvirtualmachineruntimeinfo["used_size"] = item.GetUsedSize()
		hyperflexvirtualmachineruntimeinfo["nr_version"] = item.GetVersion()
		hyperflexvirtualmachineruntimeinfo["vmx_path"] = item.GetVmxPath()

		hyperflexvirtualmachineruntimeinfos = append(hyperflexvirtualmachineruntimeinfos, hyperflexvirtualmachineruntimeinfo)
		return hyperflexvirtualmachineruntimeinfos
	})(item.GetRunTimeInfo(), d)
	hyperflexvirtualmachine["status_code"] = item.GetStatusCode()
	hyperflexvirtualmachine["uuid"] = item.GetUuid()

	hyperflexvirtualmachines = append(hyperflexvirtualmachines, hyperflexvirtualmachine)
	return hyperflexvirtualmachines
}
func flattenMapHyperflexVmBackupInfoRelationship(p models.HyperflexVmBackupInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvmbackupinforelationships []map[string]interface{}
	var ret models.HyperflexVmBackupInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexvmbackupinforelationship := make(map[string]interface{})
	hyperflexvmbackupinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexvmbackupinforelationship["class_id"] = item.GetClassId()
	hyperflexvmbackupinforelationship["moid"] = item.GetMoid()
	hyperflexvmbackupinforelationship["object_type"] = item.GetObjectType()
	hyperflexvmbackupinforelationship["selector"] = item.GetSelector()

	hyperflexvmbackupinforelationships = append(hyperflexvmbackupinforelationships, hyperflexvmbackupinforelationship)
	return hyperflexvmbackupinforelationships
}
func flattenMapHyperflexVmSnapshotInfoRelationship(p models.HyperflexVmSnapshotInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvmsnapshotinforelationships []map[string]interface{}
	var ret models.HyperflexVmSnapshotInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexvmsnapshotinforelationship := make(map[string]interface{})
	hyperflexvmsnapshotinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexvmsnapshotinforelationship["class_id"] = item.GetClassId()
	hyperflexvmsnapshotinforelationship["moid"] = item.GetMoid()
	hyperflexvmsnapshotinforelationship["object_type"] = item.GetObjectType()
	hyperflexvmsnapshotinforelationship["selector"] = item.GetSelector()

	hyperflexvmsnapshotinforelationships = append(hyperflexvmsnapshotinforelationships, hyperflexvmsnapshotinforelationship)
	return hyperflexvmsnapshotinforelationships
}
func flattenMapHyperflexWwxnPrefixRange(p models.HyperflexWwxnPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexwwxnprefixranges []map[string]interface{}
	var ret models.HyperflexWwxnPrefixRange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexwwxnprefixrange := make(map[string]interface{})
	hyperflexwwxnprefixrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexwwxnprefixrange["class_id"] = item.GetClassId()
	hyperflexwwxnprefixrange["end_addr"] = item.GetEndAddr()
	hyperflexwwxnprefixrange["object_type"] = item.GetObjectType()
	hyperflexwwxnprefixrange["start_addr"] = item.GetStartAddr()

	hyperflexwwxnprefixranges = append(hyperflexwwxnprefixranges, hyperflexwwxnprefixrange)
	return hyperflexwwxnprefixranges
}
func flattenMapIaasLicenseInfoRelationship(p models.IaasLicenseInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseinforelationships []map[string]interface{}
	var ret models.IaasLicenseInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iaaslicenseinforelationship := make(map[string]interface{})
	iaaslicenseinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iaaslicenseinforelationship["class_id"] = item.GetClassId()
	iaaslicenseinforelationship["moid"] = item.GetMoid()
	iaaslicenseinforelationship["object_type"] = item.GetObjectType()
	iaaslicenseinforelationship["selector"] = item.GetSelector()

	iaaslicenseinforelationships = append(iaaslicenseinforelationships, iaaslicenseinforelationship)
	return iaaslicenseinforelationships
}
func flattenMapIaasUcsdInfoRelationship(p models.IaasUcsdInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdinforelationships []map[string]interface{}
	var ret models.IaasUcsdInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iaasucsdinforelationship := make(map[string]interface{})
	iaasucsdinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iaasucsdinforelationship["class_id"] = item.GetClassId()
	iaasucsdinforelationship["moid"] = item.GetMoid()
	iaasucsdinforelationship["object_type"] = item.GetObjectType()
	iaasucsdinforelationship["selector"] = item.GetSelector()

	iaasucsdinforelationships = append(iaasucsdinforelationships, iaasucsdinforelationship)
	return iaasucsdinforelationships
}
func flattenMapIaasUcsdManagedInfraRelationship(p models.IaasUcsdManagedInfraRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdmanagedinfrarelationships []map[string]interface{}
	var ret models.IaasUcsdManagedInfraRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iaasucsdmanagedinfrarelationship := make(map[string]interface{})
	iaasucsdmanagedinfrarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iaasucsdmanagedinfrarelationship["class_id"] = item.GetClassId()
	iaasucsdmanagedinfrarelationship["moid"] = item.GetMoid()
	iaasucsdmanagedinfrarelationship["object_type"] = item.GetObjectType()
	iaasucsdmanagedinfrarelationship["selector"] = item.GetSelector()

	iaasucsdmanagedinfrarelationships = append(iaasucsdmanagedinfrarelationships, iaasucsdmanagedinfrarelationship)
	return iaasucsdmanagedinfrarelationships
}
func flattenMapIamAccountRelationship(p models.IamAccountRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountrelationships []map[string]interface{}
	var ret models.IamAccountRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamaccountrelationship := make(map[string]interface{})
	iamaccountrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamaccountrelationship["class_id"] = item.GetClassId()
	iamaccountrelationship["moid"] = item.GetMoid()
	iamaccountrelationship["object_type"] = item.GetObjectType()
	iamaccountrelationship["selector"] = item.GetSelector()

	iamaccountrelationships = append(iamaccountrelationships, iamaccountrelationship)
	return iamaccountrelationships
}
func flattenMapIamAppRegistrationRelationship(p models.IamAppRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrelationships []map[string]interface{}
	var ret models.IamAppRegistrationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamappregistrationrelationship := make(map[string]interface{})
	iamappregistrationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamappregistrationrelationship["class_id"] = item.GetClassId()
	iamappregistrationrelationship["moid"] = item.GetMoid()
	iamappregistrationrelationship["object_type"] = item.GetObjectType()
	iamappregistrationrelationship["selector"] = item.GetSelector()

	iamappregistrationrelationships = append(iamappregistrationrelationships, iamappregistrationrelationship)
	return iamappregistrationrelationships
}
func flattenMapIamCertificateRelationship(p models.IamCertificateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterelationships []map[string]interface{}
	var ret models.IamCertificateRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamcertificaterelationship := make(map[string]interface{})
	iamcertificaterelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamcertificaterelationship["class_id"] = item.GetClassId()
	iamcertificaterelationship["moid"] = item.GetMoid()
	iamcertificaterelationship["object_type"] = item.GetObjectType()
	iamcertificaterelationship["selector"] = item.GetSelector()

	iamcertificaterelationships = append(iamcertificaterelationships, iamcertificaterelationship)
	return iamcertificaterelationships
}
func flattenMapIamCertificateRequestRelationship(p models.IamCertificateRequestRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterequestrelationships []map[string]interface{}
	var ret models.IamCertificateRequestRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamcertificaterequestrelationship := make(map[string]interface{})
	iamcertificaterequestrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamcertificaterequestrelationship["class_id"] = item.GetClassId()
	iamcertificaterequestrelationship["moid"] = item.GetMoid()
	iamcertificaterequestrelationship["object_type"] = item.GetObjectType()
	iamcertificaterequestrelationship["selector"] = item.GetSelector()

	iamcertificaterequestrelationships = append(iamcertificaterequestrelationships, iamcertificaterequestrelationship)
	return iamcertificaterequestrelationships
}
func flattenMapIamClientMeta(p models.IamClientMeta, d *schema.ResourceData) []map[string]interface{} {
	var iamclientmetas []map[string]interface{}
	var ret models.IamClientMeta
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamclientmeta := make(map[string]interface{})
	iamclientmeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamclientmeta["class_id"] = item.GetClassId()
	iamclientmeta["device_model"] = item.GetDeviceModel()
	iamclientmeta["object_type"] = item.GetObjectType()
	iamclientmeta["user_agent"] = item.GetUserAgent()

	iamclientmetas = append(iamclientmetas, iamclientmeta)
	return iamclientmetas
}
func flattenMapIamDomainGroupRelationship(p models.IamDomainGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprelationships []map[string]interface{}
	var ret models.IamDomainGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamdomaingrouprelationship := make(map[string]interface{})
	iamdomaingrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamdomaingrouprelationship["class_id"] = item.GetClassId()
	iamdomaingrouprelationship["moid"] = item.GetMoid()
	iamdomaingrouprelationship["object_type"] = item.GetObjectType()
	iamdomaingrouprelationship["selector"] = item.GetSelector()

	iamdomaingrouprelationships = append(iamdomaingrouprelationships, iamdomaingrouprelationship)
	return iamdomaingrouprelationships
}
func flattenMapIamEndPointPasswordProperties(p models.IamEndPointPasswordProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointpasswordpropertiess []map[string]interface{}
	var ret models.IamEndPointPasswordProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamendpointpasswordproperties := make(map[string]interface{})
	iamendpointpasswordproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamendpointpasswordproperties["class_id"] = item.GetClassId()
	iamendpointpasswordproperties["enable_password_expiry"] = item.GetEnablePasswordExpiry()
	iamendpointpasswordproperties["enforce_strong_password"] = item.GetEnforceStrongPassword()
	iamendpointpasswordproperties["force_send_password"] = item.GetForceSendPassword()
	iamendpointpasswordproperties["grace_period"] = item.GetGracePeriod()
	iamendpointpasswordproperties["notification_period"] = item.GetNotificationPeriod()
	iamendpointpasswordproperties["object_type"] = item.GetObjectType()
	iamendpointpasswordproperties["password_expiry_duration"] = item.GetPasswordExpiryDuration()
	iamendpointpasswordproperties["password_history"] = item.GetPasswordHistory()

	iamendpointpasswordpropertiess = append(iamendpointpasswordpropertiess, iamendpointpasswordproperties)
	return iamendpointpasswordpropertiess
}
func flattenMapIamEndPointUserRelationship(p models.IamEndPointUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrelationships []map[string]interface{}
	var ret models.IamEndPointUserRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamendpointuserrelationship := make(map[string]interface{})
	iamendpointuserrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamendpointuserrelationship["class_id"] = item.GetClassId()
	iamendpointuserrelationship["moid"] = item.GetMoid()
	iamendpointuserrelationship["object_type"] = item.GetObjectType()
	iamendpointuserrelationship["selector"] = item.GetSelector()

	iamendpointuserrelationships = append(iamendpointuserrelationships, iamendpointuserrelationship)
	return iamendpointuserrelationships
}
func flattenMapIamEndPointUserPolicyRelationship(p models.IamEndPointUserPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserpolicyrelationships []map[string]interface{}
	var ret models.IamEndPointUserPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamendpointuserpolicyrelationship := make(map[string]interface{})
	iamendpointuserpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamendpointuserpolicyrelationship["class_id"] = item.GetClassId()
	iamendpointuserpolicyrelationship["moid"] = item.GetMoid()
	iamendpointuserpolicyrelationship["object_type"] = item.GetObjectType()
	iamendpointuserpolicyrelationship["selector"] = item.GetSelector()

	iamendpointuserpolicyrelationships = append(iamendpointuserpolicyrelationships, iamendpointuserpolicyrelationship)
	return iamendpointuserpolicyrelationships
}
func flattenMapIamIdpRelationship(p models.IamIdpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidprelationships []map[string]interface{}
	var ret models.IamIdpRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamidprelationship := make(map[string]interface{})
	iamidprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamidprelationship["class_id"] = item.GetClassId()
	iamidprelationship["moid"] = item.GetMoid()
	iamidprelationship["object_type"] = item.GetObjectType()
	iamidprelationship["selector"] = item.GetSelector()

	iamidprelationships = append(iamidprelationships, iamidprelationship)
	return iamidprelationships
}
func flattenMapIamIdpReferenceRelationship(p models.IamIdpReferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerelationships []map[string]interface{}
	var ret models.IamIdpReferenceRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamidpreferencerelationship := make(map[string]interface{})
	iamidpreferencerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamidpreferencerelationship["class_id"] = item.GetClassId()
	iamidpreferencerelationship["moid"] = item.GetMoid()
	iamidpreferencerelationship["object_type"] = item.GetObjectType()
	iamidpreferencerelationship["selector"] = item.GetSelector()

	iamidpreferencerelationships = append(iamidpreferencerelationships, iamidpreferencerelationship)
	return iamidpreferencerelationships
}
func flattenMapIamIpAccessManagementRelationship(p models.IamIpAccessManagementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamipaccessmanagementrelationships []map[string]interface{}
	var ret models.IamIpAccessManagementRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamipaccessmanagementrelationship := make(map[string]interface{})
	iamipaccessmanagementrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamipaccessmanagementrelationship["class_id"] = item.GetClassId()
	iamipaccessmanagementrelationship["moid"] = item.GetMoid()
	iamipaccessmanagementrelationship["object_type"] = item.GetObjectType()
	iamipaccessmanagementrelationship["selector"] = item.GetSelector()

	iamipaccessmanagementrelationships = append(iamipaccessmanagementrelationships, iamipaccessmanagementrelationship)
	return iamipaccessmanagementrelationships
}
func flattenMapIamLdapBaseProperties(p models.IamLdapBaseProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamldapbasepropertiess []map[string]interface{}
	var ret models.IamLdapBaseProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamldapbaseproperties := make(map[string]interface{})
	iamldapbaseproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamldapbaseproperties["attribute"] = item.GetAttribute()
	iamldapbaseproperties["base_dn"] = item.GetBaseDn()
	iamldapbaseproperties["bind_dn"] = item.GetBindDn()
	iamldapbaseproperties["bind_method"] = item.GetBindMethod()
	iamldapbaseproperties["class_id"] = item.GetClassId()
	iamldapbaseproperties["domain"] = item.GetDomain()
	iamldapbaseproperties["enable_encryption"] = item.GetEnableEncryption()
	iamldapbaseproperties["enable_group_authorization"] = item.GetEnableGroupAuthorization()
	iamldapbaseproperties["filter"] = item.GetFilter()
	iamldapbaseproperties["group_attribute"] = item.GetGroupAttribute()
	iamldapbaseproperties["is_password_set"] = item.GetIsPasswordSet()
	iamldapbaseproperties["nested_group_search_depth"] = item.GetNestedGroupSearchDepth()
	iamldapbaseproperties["object_type"] = item.GetObjectType()
	password_x, exists := d.GetOk("base_properties")
	if exists && password_x != nil {
		password_y := password_x.([]interface{})[0].(map[string]interface{})
		iamldapbaseproperties["password"] = password_y["password"]
	}
	iamldapbaseproperties["timeout"] = item.GetTimeout()

	iamldapbasepropertiess = append(iamldapbasepropertiess, iamldapbaseproperties)
	return iamldapbasepropertiess
}
func flattenMapIamLdapDnsParameters(p models.IamLdapDnsParameters, d *schema.ResourceData) []map[string]interface{} {
	var iamldapdnsparameterss []map[string]interface{}
	var ret models.IamLdapDnsParameters
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamldapdnsparameters := make(map[string]interface{})
	iamldapdnsparameters["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamldapdnsparameters["class_id"] = item.GetClassId()
	iamldapdnsparameters["object_type"] = item.GetObjectType()
	iamldapdnsparameters["search_domain"] = item.GetSearchDomain()
	iamldapdnsparameters["search_forest"] = item.GetSearchForest()
	iamldapdnsparameters["nr_source"] = item.GetSource()

	iamldapdnsparameterss = append(iamldapdnsparameterss, iamldapdnsparameters)
	return iamldapdnsparameterss
}
func flattenMapIamLdapPolicyRelationship(p models.IamLdapPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldappolicyrelationships []map[string]interface{}
	var ret models.IamLdapPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamldappolicyrelationship := make(map[string]interface{})
	iamldappolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamldappolicyrelationship["class_id"] = item.GetClassId()
	iamldappolicyrelationship["moid"] = item.GetMoid()
	iamldappolicyrelationship["object_type"] = item.GetObjectType()
	iamldappolicyrelationship["selector"] = item.GetSelector()

	iamldappolicyrelationships = append(iamldappolicyrelationships, iamldappolicyrelationship)
	return iamldappolicyrelationships
}
func flattenMapIamLocalUserPasswordRelationship(p models.IamLocalUserPasswordRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamlocaluserpasswordrelationships []map[string]interface{}
	var ret models.IamLocalUserPasswordRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamlocaluserpasswordrelationship := make(map[string]interface{})
	iamlocaluserpasswordrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamlocaluserpasswordrelationship["class_id"] = item.GetClassId()
	iamlocaluserpasswordrelationship["moid"] = item.GetMoid()
	iamlocaluserpasswordrelationship["object_type"] = item.GetObjectType()
	iamlocaluserpasswordrelationship["selector"] = item.GetSelector()

	iamlocaluserpasswordrelationships = append(iamlocaluserpasswordrelationships, iamlocaluserpasswordrelationship)
	return iamlocaluserpasswordrelationships
}
func flattenMapIamPermissionRelationship(p models.IamPermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrelationships []map[string]interface{}
	var ret models.IamPermissionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iampermissionrelationship := make(map[string]interface{})
	iampermissionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iampermissionrelationship["class_id"] = item.GetClassId()
	iampermissionrelationship["moid"] = item.GetMoid()
	iampermissionrelationship["object_type"] = item.GetObjectType()
	iampermissionrelationship["selector"] = item.GetSelector()

	iampermissionrelationships = append(iampermissionrelationships, iampermissionrelationship)
	return iampermissionrelationships
}
func flattenMapIamPrivateKeySpecRelationship(p models.IamPrivateKeySpecRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivatekeyspecrelationships []map[string]interface{}
	var ret models.IamPrivateKeySpecRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamprivatekeyspecrelationship := make(map[string]interface{})
	iamprivatekeyspecrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamprivatekeyspecrelationship["class_id"] = item.GetClassId()
	iamprivatekeyspecrelationship["moid"] = item.GetMoid()
	iamprivatekeyspecrelationship["object_type"] = item.GetObjectType()
	iamprivatekeyspecrelationship["selector"] = item.GetSelector()

	iamprivatekeyspecrelationships = append(iamprivatekeyspecrelationships, iamprivatekeyspecrelationship)
	return iamprivatekeyspecrelationships
}
func flattenMapIamQualifierRelationship(p models.IamQualifierRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamqualifierrelationships []map[string]interface{}
	var ret models.IamQualifierRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamqualifierrelationship := make(map[string]interface{})
	iamqualifierrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamqualifierrelationship["class_id"] = item.GetClassId()
	iamqualifierrelationship["moid"] = item.GetMoid()
	iamqualifierrelationship["object_type"] = item.GetObjectType()
	iamqualifierrelationship["selector"] = item.GetSelector()

	iamqualifierrelationships = append(iamqualifierrelationships, iamqualifierrelationship)
	return iamqualifierrelationships
}
func flattenMapIamResourceLimitsRelationship(p models.IamResourceLimitsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcelimitsrelationships []map[string]interface{}
	var ret models.IamResourceLimitsRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamresourcelimitsrelationship := make(map[string]interface{})
	iamresourcelimitsrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamresourcelimitsrelationship["class_id"] = item.GetClassId()
	iamresourcelimitsrelationship["moid"] = item.GetMoid()
	iamresourcelimitsrelationship["object_type"] = item.GetObjectType()
	iamresourcelimitsrelationship["selector"] = item.GetSelector()

	iamresourcelimitsrelationships = append(iamresourcelimitsrelationships, iamresourcelimitsrelationship)
	return iamresourcelimitsrelationships
}
func flattenMapIamSecurityHolderRelationship(p models.IamSecurityHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsecurityholderrelationships []map[string]interface{}
	var ret models.IamSecurityHolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsecurityholderrelationship := make(map[string]interface{})
	iamsecurityholderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsecurityholderrelationship["class_id"] = item.GetClassId()
	iamsecurityholderrelationship["moid"] = item.GetMoid()
	iamsecurityholderrelationship["object_type"] = item.GetObjectType()
	iamsecurityholderrelationship["selector"] = item.GetSelector()

	iamsecurityholderrelationships = append(iamsecurityholderrelationships, iamsecurityholderrelationship)
	return iamsecurityholderrelationships
}
func flattenMapIamServiceProviderRelationship(p models.IamServiceProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamserviceproviderrelationships []map[string]interface{}
	var ret models.IamServiceProviderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamserviceproviderrelationship := make(map[string]interface{})
	iamserviceproviderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamserviceproviderrelationship["class_id"] = item.GetClassId()
	iamserviceproviderrelationship["moid"] = item.GetMoid()
	iamserviceproviderrelationship["object_type"] = item.GetObjectType()
	iamserviceproviderrelationship["selector"] = item.GetSelector()

	iamserviceproviderrelationships = append(iamserviceproviderrelationships, iamserviceproviderrelationship)
	return iamserviceproviderrelationships
}
func flattenMapIamSessionRelationship(p models.IamSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrelationships []map[string]interface{}
	var ret models.IamSessionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsessionrelationship := make(map[string]interface{})
	iamsessionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsessionrelationship["class_id"] = item.GetClassId()
	iamsessionrelationship["moid"] = item.GetMoid()
	iamsessionrelationship["object_type"] = item.GetObjectType()
	iamsessionrelationship["selector"] = item.GetSelector()

	iamsessionrelationships = append(iamsessionrelationships, iamsessionrelationship)
	return iamsessionrelationships
}
func flattenMapIamSessionLimitsRelationship(p models.IamSessionLimitsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionlimitsrelationships []map[string]interface{}
	var ret models.IamSessionLimitsRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsessionlimitsrelationship := make(map[string]interface{})
	iamsessionlimitsrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsessionlimitsrelationship["class_id"] = item.GetClassId()
	iamsessionlimitsrelationship["moid"] = item.GetMoid()
	iamsessionlimitsrelationship["object_type"] = item.GetObjectType()
	iamsessionlimitsrelationship["selector"] = item.GetSelector()

	iamsessionlimitsrelationships = append(iamsessionlimitsrelationships, iamsessionlimitsrelationship)
	return iamsessionlimitsrelationships
}
func flattenMapIamSystemRelationship(p models.IamSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsystemrelationships []map[string]interface{}
	var ret models.IamSystemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsystemrelationship := make(map[string]interface{})
	iamsystemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsystemrelationship["class_id"] = item.GetClassId()
	iamsystemrelationship["moid"] = item.GetMoid()
	iamsystemrelationship["object_type"] = item.GetObjectType()
	iamsystemrelationship["selector"] = item.GetSelector()

	iamsystemrelationships = append(iamsystemrelationships, iamsystemrelationship)
	return iamsystemrelationships
}
func flattenMapIamUserRelationship(p models.IamUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrelationships []map[string]interface{}
	var ret models.IamUserRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamuserrelationship := make(map[string]interface{})
	iamuserrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamuserrelationship["class_id"] = item.GetClassId()
	iamuserrelationship["moid"] = item.GetMoid()
	iamuserrelationship["object_type"] = item.GetObjectType()
	iamuserrelationship["selector"] = item.GetSelector()

	iamuserrelationships = append(iamuserrelationships, iamuserrelationship)
	return iamuserrelationships
}
func flattenMapIamUserGroupRelationship(p models.IamUserGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprelationships []map[string]interface{}
	var ret models.IamUserGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamusergrouprelationship := make(map[string]interface{})
	iamusergrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamusergrouprelationship["class_id"] = item.GetClassId()
	iamusergrouprelationship["moid"] = item.GetMoid()
	iamusergrouprelationship["object_type"] = item.GetObjectType()
	iamusergrouprelationship["selector"] = item.GetSelector()

	iamusergrouprelationships = append(iamusergrouprelationships, iamusergrouprelationship)
	return iamusergrouprelationships
}
func flattenMapInfraHardwareInfo(p models.InfraHardwareInfo, d *schema.ResourceData) []map[string]interface{} {
	var infrahardwareinfos []map[string]interface{}
	var ret models.InfraHardwareInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	infrahardwareinfo := make(map[string]interface{})
	infrahardwareinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	infrahardwareinfo["class_id"] = item.GetClassId()
	infrahardwareinfo["cpu_cores"] = item.GetCpuCores()
	infrahardwareinfo["cpu_speed"] = item.GetCpuSpeed()
	infrahardwareinfo["memory_size"] = item.GetMemorySize()
	infrahardwareinfo["object_type"] = item.GetObjectType()

	infrahardwareinfos = append(infrahardwareinfos, infrahardwareinfo)
	return infrahardwareinfos
}
func flattenMapInventoryBaseRelationship(p models.InventoryBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorybaserelationships []map[string]interface{}
	var ret models.InventoryBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventorybaserelationship := make(map[string]interface{})
	inventorybaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventorybaserelationship["class_id"] = item.GetClassId()
	inventorybaserelationship["moid"] = item.GetMoid()
	inventorybaserelationship["object_type"] = item.GetObjectType()
	inventorybaserelationship["selector"] = item.GetSelector()

	inventorybaserelationships = append(inventorybaserelationships, inventorybaserelationship)
	return inventorybaserelationships
}
func flattenMapInventoryDeviceInfoRelationship(p models.InventoryDeviceInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorydeviceinforelationships []map[string]interface{}
	var ret models.InventoryDeviceInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventorydeviceinforelationship := make(map[string]interface{})
	inventorydeviceinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventorydeviceinforelationship["class_id"] = item.GetClassId()
	inventorydeviceinforelationship["moid"] = item.GetMoid()
	inventorydeviceinforelationship["object_type"] = item.GetObjectType()
	inventorydeviceinforelationship["selector"] = item.GetSelector()

	inventorydeviceinforelationships = append(inventorydeviceinforelationships, inventorydeviceinforelationship)
	return inventorydeviceinforelationships
}
func flattenMapInventoryGenericInventoryHolderRelationship(p models.InventoryGenericInventoryHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrelationships []map[string]interface{}
	var ret models.InventoryGenericInventoryHolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventorygenericinventoryholderrelationship := make(map[string]interface{})
	inventorygenericinventoryholderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventorygenericinventoryholderrelationship["class_id"] = item.GetClassId()
	inventorygenericinventoryholderrelationship["moid"] = item.GetMoid()
	inventorygenericinventoryholderrelationship["object_type"] = item.GetObjectType()
	inventorygenericinventoryholderrelationship["selector"] = item.GetSelector()

	inventorygenericinventoryholderrelationships = append(inventorygenericinventoryholderrelationships, inventorygenericinventoryholderrelationship)
	return inventorygenericinventoryholderrelationships
}
func flattenMapIppoolBlockLeaseRelationship(p models.IppoolBlockLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolblockleaserelationships []map[string]interface{}
	var ret models.IppoolBlockLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolblockleaserelationship := make(map[string]interface{})
	ippoolblockleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolblockleaserelationship["class_id"] = item.GetClassId()
	ippoolblockleaserelationship["moid"] = item.GetMoid()
	ippoolblockleaserelationship["object_type"] = item.GetObjectType()
	ippoolblockleaserelationship["selector"] = item.GetSelector()

	ippoolblockleaserelationships = append(ippoolblockleaserelationships, ippoolblockleaserelationship)
	return ippoolblockleaserelationships
}
func flattenMapIppoolIpLeaseRelationship(p models.IppoolIpLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipleaserelationships []map[string]interface{}
	var ret models.IppoolIpLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolipleaserelationship := make(map[string]interface{})
	ippoolipleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipleaserelationship["class_id"] = item.GetClassId()
	ippoolipleaserelationship["moid"] = item.GetMoid()
	ippoolipleaserelationship["object_type"] = item.GetObjectType()
	ippoolipleaserelationship["selector"] = item.GetSelector()

	ippoolipleaserelationships = append(ippoolipleaserelationships, ippoolipleaserelationship)
	return ippoolipleaserelationships
}
func flattenMapIppoolIpV4Block(p models.IppoolIpV4Block, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv4blocks []map[string]interface{}
	var ret models.IppoolIpV4Block
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	ippoolipv4block := make(map[string]interface{})
	ippoolipv4block["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipv4block["class_id"] = item.GetClassId()
	ippoolipv4block["from"] = item.GetFrom()
	ippoolipv4block["object_type"] = item.GetObjectType()
	ippoolipv4block["size"] = item.GetSize()
	ippoolipv4block["to"] = item.GetTo()

	ippoolipv4blocks = append(ippoolipv4blocks, ippoolipv4block)
	return ippoolipv4blocks
}
func flattenMapIppoolIpV4Config(p models.IppoolIpV4Config, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv4configs []map[string]interface{}
	var ret models.IppoolIpV4Config
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	ippoolipv4config := make(map[string]interface{})
	ippoolipv4config["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipv4config["class_id"] = item.GetClassId()
	ippoolipv4config["gateway"] = item.GetGateway()
	ippoolipv4config["netmask"] = item.GetNetmask()
	ippoolipv4config["object_type"] = item.GetObjectType()
	ippoolipv4config["primary_dns"] = item.GetPrimaryDns()
	ippoolipv4config["secondary_dns"] = item.GetSecondaryDns()

	ippoolipv4configs = append(ippoolipv4configs, ippoolipv4config)
	return ippoolipv4configs
}
func flattenMapIppoolIpV6Block(p models.IppoolIpV6Block, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv6blocks []map[string]interface{}
	var ret models.IppoolIpV6Block
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	ippoolipv6block := make(map[string]interface{})
	ippoolipv6block["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipv6block["class_id"] = item.GetClassId()
	ippoolipv6block["from"] = item.GetFrom()
	ippoolipv6block["object_type"] = item.GetObjectType()
	ippoolipv6block["size"] = item.GetSize()
	ippoolipv6block["to"] = item.GetTo()

	ippoolipv6blocks = append(ippoolipv6blocks, ippoolipv6block)
	return ippoolipv6blocks
}
func flattenMapIppoolIpV6Config(p models.IppoolIpV6Config, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv6configs []map[string]interface{}
	var ret models.IppoolIpV6Config
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	ippoolipv6config := make(map[string]interface{})
	ippoolipv6config["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipv6config["class_id"] = item.GetClassId()
	ippoolipv6config["gateway"] = item.GetGateway()
	ippoolipv6config["object_type"] = item.GetObjectType()
	ippoolipv6config["prefix"] = item.GetPrefix()
	ippoolipv6config["primary_dns"] = item.GetPrimaryDns()
	ippoolipv6config["secondary_dns"] = item.GetSecondaryDns()

	ippoolipv6configs = append(ippoolipv6configs, ippoolipv6config)
	return ippoolipv6configs
}
func flattenMapIppoolPoolRelationship(p models.IppoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolpoolrelationships []map[string]interface{}
	var ret models.IppoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolpoolrelationship := make(map[string]interface{})
	ippoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolpoolrelationship["class_id"] = item.GetClassId()
	ippoolpoolrelationship["moid"] = item.GetMoid()
	ippoolpoolrelationship["object_type"] = item.GetObjectType()
	ippoolpoolrelationship["selector"] = item.GetSelector()

	ippoolpoolrelationships = append(ippoolpoolrelationships, ippoolpoolrelationship)
	return ippoolpoolrelationships
}
func flattenMapIppoolPoolMemberRelationship(p models.IppoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolpoolmemberrelationships []map[string]interface{}
	var ret models.IppoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolpoolmemberrelationship := make(map[string]interface{})
	ippoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolpoolmemberrelationship["class_id"] = item.GetClassId()
	ippoolpoolmemberrelationship["moid"] = item.GetMoid()
	ippoolpoolmemberrelationship["object_type"] = item.GetObjectType()
	ippoolpoolmemberrelationship["selector"] = item.GetSelector()

	ippoolpoolmemberrelationships = append(ippoolpoolmemberrelationships, ippoolpoolmemberrelationship)
	return ippoolpoolmemberrelationships
}
func flattenMapIppoolShadowBlockRelationship(p models.IppoolShadowBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowblockrelationships []map[string]interface{}
	var ret models.IppoolShadowBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolshadowblockrelationship := make(map[string]interface{})
	ippoolshadowblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolshadowblockrelationship["class_id"] = item.GetClassId()
	ippoolshadowblockrelationship["moid"] = item.GetMoid()
	ippoolshadowblockrelationship["object_type"] = item.GetObjectType()
	ippoolshadowblockrelationship["selector"] = item.GetSelector()

	ippoolshadowblockrelationships = append(ippoolshadowblockrelationships, ippoolshadowblockrelationship)
	return ippoolshadowblockrelationships
}
func flattenMapIppoolShadowPoolRelationship(p models.IppoolShadowPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowpoolrelationships []map[string]interface{}
	var ret models.IppoolShadowPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolshadowpoolrelationship := make(map[string]interface{})
	ippoolshadowpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolshadowpoolrelationship["class_id"] = item.GetClassId()
	ippoolshadowpoolrelationship["moid"] = item.GetMoid()
	ippoolshadowpoolrelationship["object_type"] = item.GetObjectType()
	ippoolshadowpoolrelationship["selector"] = item.GetSelector()

	ippoolshadowpoolrelationships = append(ippoolshadowpoolrelationships, ippoolshadowpoolrelationship)
	return ippoolshadowpoolrelationships
}
func flattenMapIppoolUniverseRelationship(p models.IppoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippooluniverserelationships []map[string]interface{}
	var ret models.IppoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippooluniverserelationship := make(map[string]interface{})
	ippooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippooluniverserelationship["class_id"] = item.GetClassId()
	ippooluniverserelationship["moid"] = item.GetMoid()
	ippooluniverserelationship["object_type"] = item.GetObjectType()
	ippooluniverserelationship["selector"] = item.GetSelector()

	ippooluniverserelationships = append(ippooluniverserelationships, ippooluniverserelationship)
	return ippooluniverserelationships
}
func flattenMapIqnpoolBlockRelationship(p models.IqnpoolBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iqnpoolblockrelationships []map[string]interface{}
	var ret models.IqnpoolBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iqnpoolblockrelationship := make(map[string]interface{})
	iqnpoolblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iqnpoolblockrelationship["class_id"] = item.GetClassId()
	iqnpoolblockrelationship["moid"] = item.GetMoid()
	iqnpoolblockrelationship["object_type"] = item.GetObjectType()
	iqnpoolblockrelationship["selector"] = item.GetSelector()

	iqnpoolblockrelationships = append(iqnpoolblockrelationships, iqnpoolblockrelationship)
	return iqnpoolblockrelationships
}
func flattenMapIqnpoolIqnSuffixBlock(p models.IqnpoolIqnSuffixBlock, d *schema.ResourceData) []map[string]interface{} {
	var iqnpooliqnsuffixblocks []map[string]interface{}
	var ret models.IqnpoolIqnSuffixBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iqnpooliqnsuffixblock := make(map[string]interface{})
	iqnpooliqnsuffixblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iqnpooliqnsuffixblock["class_id"] = item.GetClassId()
	iqnpooliqnsuffixblock["from"] = item.GetFrom()
	iqnpooliqnsuffixblock["object_type"] = item.GetObjectType()
	iqnpooliqnsuffixblock["size"] = item.GetSize()
	iqnpooliqnsuffixblock["suffix"] = item.GetSuffix()
	iqnpooliqnsuffixblock["to"] = item.GetTo()

	iqnpooliqnsuffixblocks = append(iqnpooliqnsuffixblocks, iqnpooliqnsuffixblock)
	return iqnpooliqnsuffixblocks
}
func flattenMapIqnpoolLeaseRelationship(p models.IqnpoolLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iqnpoolleaserelationships []map[string]interface{}
	var ret models.IqnpoolLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iqnpoolleaserelationship := make(map[string]interface{})
	iqnpoolleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iqnpoolleaserelationship["class_id"] = item.GetClassId()
	iqnpoolleaserelationship["moid"] = item.GetMoid()
	iqnpoolleaserelationship["object_type"] = item.GetObjectType()
	iqnpoolleaserelationship["selector"] = item.GetSelector()

	iqnpoolleaserelationships = append(iqnpoolleaserelationships, iqnpoolleaserelationship)
	return iqnpoolleaserelationships
}
func flattenMapIqnpoolPoolRelationship(p models.IqnpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iqnpoolpoolrelationships []map[string]interface{}
	var ret models.IqnpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iqnpoolpoolrelationship := make(map[string]interface{})
	iqnpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iqnpoolpoolrelationship["class_id"] = item.GetClassId()
	iqnpoolpoolrelationship["moid"] = item.GetMoid()
	iqnpoolpoolrelationship["object_type"] = item.GetObjectType()
	iqnpoolpoolrelationship["selector"] = item.GetSelector()

	iqnpoolpoolrelationships = append(iqnpoolpoolrelationships, iqnpoolpoolrelationship)
	return iqnpoolpoolrelationships
}
func flattenMapIqnpoolPoolMemberRelationship(p models.IqnpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iqnpoolpoolmemberrelationships []map[string]interface{}
	var ret models.IqnpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iqnpoolpoolmemberrelationship := make(map[string]interface{})
	iqnpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iqnpoolpoolmemberrelationship["class_id"] = item.GetClassId()
	iqnpoolpoolmemberrelationship["moid"] = item.GetMoid()
	iqnpoolpoolmemberrelationship["object_type"] = item.GetObjectType()
	iqnpoolpoolmemberrelationship["selector"] = item.GetSelector()

	iqnpoolpoolmemberrelationships = append(iqnpoolpoolmemberrelationships, iqnpoolpoolmemberrelationship)
	return iqnpoolpoolmemberrelationships
}
func flattenMapIqnpoolUniverseRelationship(p models.IqnpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iqnpooluniverserelationships []map[string]interface{}
	var ret models.IqnpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iqnpooluniverserelationship := make(map[string]interface{})
	iqnpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iqnpooluniverserelationship["class_id"] = item.GetClassId()
	iqnpooluniverserelationship["moid"] = item.GetMoid()
	iqnpooluniverserelationship["object_type"] = item.GetObjectType()
	iqnpooluniverserelationship["selector"] = item.GetSelector()

	iqnpooluniverserelationships = append(iqnpooluniverserelationships, iqnpooluniverserelationship)
	return iqnpooluniverserelationships
}
func flattenMapKubernetesAciCniProfileRelationship(p models.KubernetesAciCniProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesacicniprofilerelationships []map[string]interface{}
	var ret models.KubernetesAciCniProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesacicniprofilerelationship := make(map[string]interface{})
	kubernetesacicniprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesacicniprofilerelationship["class_id"] = item.GetClassId()
	kubernetesacicniprofilerelationship["moid"] = item.GetMoid()
	kubernetesacicniprofilerelationship["object_type"] = item.GetObjectType()
	kubernetesacicniprofilerelationship["selector"] = item.GetSelector()

	kubernetesacicniprofilerelationships = append(kubernetesacicniprofilerelationships, kubernetesacicniprofilerelationship)
	return kubernetesacicniprofilerelationships
}
func flattenMapKubernetesActionInfo(p models.KubernetesActionInfo, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesactioninfos []map[string]interface{}
	var ret models.KubernetesActionInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesactioninfo := make(map[string]interface{})
	kubernetesactioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesactioninfo["class_id"] = item.GetClassId()
	kubernetesactioninfo["failure_reason"] = item.GetFailureReason()
	kubernetesactioninfo["name"] = item.GetName()
	kubernetesactioninfo["object_type"] = item.GetObjectType()
	kubernetesactioninfo["status"] = item.GetStatus()

	kubernetesactioninfos = append(kubernetesactioninfos, kubernetesactioninfo)
	return kubernetesactioninfos
}
func flattenMapKubernetesAddonConfiguration(p models.KubernetesAddonConfiguration, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesaddonconfigurations []map[string]interface{}
	var ret models.KubernetesAddonConfiguration
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesaddonconfiguration := make(map[string]interface{})
	kubernetesaddonconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesaddonconfiguration["class_id"] = item.GetClassId()
	kubernetesaddonconfiguration["install_strategy"] = item.GetInstallStrategy()
	kubernetesaddonconfiguration["object_type"] = item.GetObjectType()
	kubernetesaddonconfiguration["override_sets"] = (func(p []models.KubernetesKeyValue, d *schema.ResourceData) []map[string]interface{} {
		var kuberneteskeyvalues []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			kuberneteskeyvalue := make(map[string]interface{})
			kuberneteskeyvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			kuberneteskeyvalue["class_id"] = item.GetClassId()
			kuberneteskeyvalue["key"] = item.GetKey()
			kuberneteskeyvalue["object_type"] = item.GetObjectType()
			kuberneteskeyvalue["value"] = item.GetValue()
			kuberneteskeyvalues = append(kuberneteskeyvalues, kuberneteskeyvalue)
		}
		return kuberneteskeyvalues
	})(item.GetOverrideSets(), d)
	kubernetesaddonconfiguration["overrides"] = item.GetOverrides()
	kubernetesaddonconfiguration["release_name"] = item.GetReleaseName()
	kubernetesaddonconfiguration["release_namespace"] = item.GetReleaseNamespace()
	kubernetesaddonconfiguration["upgrade_strategy"] = item.GetUpgradeStrategy()

	kubernetesaddonconfigurations = append(kubernetesaddonconfigurations, kubernetesaddonconfiguration)
	return kubernetesaddonconfigurations
}
func flattenMapKubernetesAddonDefinitionRelationship(p models.KubernetesAddonDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesaddondefinitionrelationships []map[string]interface{}
	var ret models.KubernetesAddonDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesaddondefinitionrelationship := make(map[string]interface{})
	kubernetesaddondefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesaddondefinitionrelationship["class_id"] = item.GetClassId()
	kubernetesaddondefinitionrelationship["moid"] = item.GetMoid()
	kubernetesaddondefinitionrelationship["object_type"] = item.GetObjectType()
	kubernetesaddondefinitionrelationship["selector"] = item.GetSelector()

	kubernetesaddondefinitionrelationships = append(kubernetesaddondefinitionrelationships, kubernetesaddondefinitionrelationship)
	return kubernetesaddondefinitionrelationships
}
func flattenMapKubernetesBaseInfrastructureProviderRelationship(p models.KubernetesBaseInfrastructureProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesbaseinfrastructureproviderrelationships []map[string]interface{}
	var ret models.KubernetesBaseInfrastructureProviderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesbaseinfrastructureproviderrelationship := make(map[string]interface{})
	kubernetesbaseinfrastructureproviderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesbaseinfrastructureproviderrelationship["class_id"] = item.GetClassId()
	kubernetesbaseinfrastructureproviderrelationship["moid"] = item.GetMoid()
	kubernetesbaseinfrastructureproviderrelationship["object_type"] = item.GetObjectType()
	kubernetesbaseinfrastructureproviderrelationship["selector"] = item.GetSelector()

	kubernetesbaseinfrastructureproviderrelationships = append(kubernetesbaseinfrastructureproviderrelationships, kubernetesbaseinfrastructureproviderrelationship)
	return kubernetesbaseinfrastructureproviderrelationships
}
func flattenMapKubernetesBaseVirtualMachineInfraConfig(p models.KubernetesBaseVirtualMachineInfraConfig, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesbasevirtualmachineinfraconfigs []map[string]interface{}
	var ret models.KubernetesBaseVirtualMachineInfraConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesbasevirtualmachineinfraconfig := make(map[string]interface{})
	kubernetesbasevirtualmachineinfraconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesbasevirtualmachineinfraconfig["class_id"] = item.GetClassId()
	kubernetesbasevirtualmachineinfraconfig["interfaces"] = item.GetInterfaces()
	kubernetesbasevirtualmachineinfraconfig["object_type"] = item.GetObjectType()

	kubernetesbasevirtualmachineinfraconfigs = append(kubernetesbasevirtualmachineinfraconfigs, kubernetesbasevirtualmachineinfraconfig)
	return kubernetesbasevirtualmachineinfraconfigs
}
func flattenMapKubernetesCatalogRelationship(p models.KubernetesCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetescatalogrelationships []map[string]interface{}
	var ret models.KubernetesCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetescatalogrelationship := make(map[string]interface{})
	kubernetescatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetescatalogrelationship["class_id"] = item.GetClassId()
	kubernetescatalogrelationship["moid"] = item.GetMoid()
	kubernetescatalogrelationship["object_type"] = item.GetObjectType()
	kubernetescatalogrelationship["selector"] = item.GetSelector()

	kubernetescatalogrelationships = append(kubernetescatalogrelationships, kubernetescatalogrelationship)
	return kubernetescatalogrelationships
}
func flattenMapKubernetesClusterRelationship(p models.KubernetesClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesclusterrelationships []map[string]interface{}
	var ret models.KubernetesClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesclusterrelationship := make(map[string]interface{})
	kubernetesclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesclusterrelationship["class_id"] = item.GetClassId()
	kubernetesclusterrelationship["moid"] = item.GetMoid()
	kubernetesclusterrelationship["object_type"] = item.GetObjectType()
	kubernetesclusterrelationship["selector"] = item.GetSelector()

	kubernetesclusterrelationships = append(kubernetesclusterrelationships, kubernetesclusterrelationship)
	return kubernetesclusterrelationships
}
func flattenMapKubernetesClusterAddonProfileRelationship(p models.KubernetesClusterAddonProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesclusteraddonprofilerelationships []map[string]interface{}
	var ret models.KubernetesClusterAddonProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesclusteraddonprofilerelationship := make(map[string]interface{})
	kubernetesclusteraddonprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesclusteraddonprofilerelationship["class_id"] = item.GetClassId()
	kubernetesclusteraddonprofilerelationship["moid"] = item.GetMoid()
	kubernetesclusteraddonprofilerelationship["object_type"] = item.GetObjectType()
	kubernetesclusteraddonprofilerelationship["selector"] = item.GetSelector()

	kubernetesclusteraddonprofilerelationships = append(kubernetesclusteraddonprofilerelationships, kubernetesclusteraddonprofilerelationship)
	return kubernetesclusteraddonprofilerelationships
}
func flattenMapKubernetesClusterCertificateConfiguration(p models.KubernetesClusterCertificateConfiguration, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesclustercertificateconfigurations []map[string]interface{}
	var ret models.KubernetesClusterCertificateConfiguration
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesclustercertificateconfiguration := make(map[string]interface{})
	kubernetesclustercertificateconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesclustercertificateconfiguration["ca_cert"] = item.GetCaCert()
	ca_key_x, exists := d.GetOk("cert_config")
	if exists && ca_key_x != nil {
		ca_key_y := ca_key_x.([]interface{})[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["ca_key"] = ca_key_y["ca_key"]
	}
	kubernetesclustercertificateconfiguration["class_id"] = item.GetClassId()
	kubernetesclustercertificateconfiguration["etcd_cert"] = item.GetEtcdCert()
	kubernetesclustercertificateconfiguration["etcd_encryption_key"] = item.GetEtcdEncryptionKey()
	etcd_key_x, exists := d.GetOk("cert_config")
	if exists && etcd_key_x != nil {
		etcd_key_y := etcd_key_x.([]interface{})[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["etcd_key"] = etcd_key_y["etcd_key"]
	}
	kubernetesclustercertificateconfiguration["front_proxy_cert"] = item.GetFrontProxyCert()
	front_proxy_key_x, exists := d.GetOk("cert_config")
	if exists && front_proxy_key_x != nil {
		front_proxy_key_y := front_proxy_key_x.([]interface{})[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["front_proxy_key"] = front_proxy_key_y["front_proxy_key"]
	}
	kubernetesclustercertificateconfiguration["object_type"] = item.GetObjectType()
	sa_private_key_x, exists := d.GetOk("cert_config")
	if exists && sa_private_key_x != nil {
		sa_private_key_y := sa_private_key_x.([]interface{})[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["sa_private_key"] = sa_private_key_y["sa_private_key"]
	}
	kubernetesclustercertificateconfiguration["sa_public_key"] = item.GetSaPublicKey()

	kubernetesclustercertificateconfigurations = append(kubernetesclustercertificateconfigurations, kubernetesclustercertificateconfiguration)
	return kubernetesclustercertificateconfigurations
}
func flattenMapKubernetesClusterManagementConfig(p models.KubernetesClusterManagementConfig, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesclustermanagementconfigs []map[string]interface{}
	var ret models.KubernetesClusterManagementConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesclustermanagementconfig := make(map[string]interface{})
	kubernetesclustermanagementconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesclustermanagementconfig["class_id"] = item.GetClassId()
	kubernetesclustermanagementconfig["load_balancer_count"] = item.GetLoadBalancerCount()
	kubernetesclustermanagementconfig["load_balancers"] = item.GetLoadBalancers()
	kubernetesclustermanagementconfig["master_vip"] = item.GetMasterVip()
	kubernetesclustermanagementconfig["object_type"] = item.GetObjectType()
	kubernetesclustermanagementconfig["ssh_keys"] = item.GetSshKeys()
	kubernetesclustermanagementconfig["ssh_user"] = item.GetSshUser()

	kubernetesclustermanagementconfigs = append(kubernetesclustermanagementconfigs, kubernetesclustermanagementconfig)
	return kubernetesclustermanagementconfigs
}
func flattenMapKubernetesClusterProfileRelationship(p models.KubernetesClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesclusterprofilerelationships []map[string]interface{}
	var ret models.KubernetesClusterProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesclusterprofilerelationship := make(map[string]interface{})
	kubernetesclusterprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesclusterprofilerelationship["class_id"] = item.GetClassId()
	kubernetesclusterprofilerelationship["moid"] = item.GetMoid()
	kubernetesclusterprofilerelationship["object_type"] = item.GetObjectType()
	kubernetesclusterprofilerelationship["selector"] = item.GetSelector()

	kubernetesclusterprofilerelationships = append(kubernetesclusterprofilerelationships, kubernetesclusterprofilerelationship)
	return kubernetesclusterprofilerelationships
}
func flattenMapKubernetesCniConfig(p models.KubernetesCniConfig, d *schema.ResourceData) []map[string]interface{} {
	var kubernetescniconfigs []map[string]interface{}
	var ret models.KubernetesCniConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetescniconfig := make(map[string]interface{})
	kubernetescniconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetescniconfig["class_id"] = item.GetClassId()
	kubernetescniconfig["object_type"] = item.GetObjectType()
	kubernetescniconfig["registry"] = item.GetRegistry()
	kubernetescniconfig["nr_version"] = item.GetVersion()

	kubernetescniconfigs = append(kubernetescniconfigs, kubernetescniconfig)
	return kubernetescniconfigs
}
func flattenMapKubernetesConfigResultRelationship(p models.KubernetesConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesconfigresultrelationships []map[string]interface{}
	var ret models.KubernetesConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesconfigresultrelationship := make(map[string]interface{})
	kubernetesconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesconfigresultrelationship["class_id"] = item.GetClassId()
	kubernetesconfigresultrelationship["moid"] = item.GetMoid()
	kubernetesconfigresultrelationship["object_type"] = item.GetObjectType()
	kubernetesconfigresultrelationship["selector"] = item.GetSelector()

	kubernetesconfigresultrelationships = append(kubernetesconfigresultrelationships, kubernetesconfigresultrelationship)
	return kubernetesconfigresultrelationships
}
func flattenMapKubernetesConfiguration(p models.KubernetesConfiguration, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesconfigurations []map[string]interface{}
	var ret models.KubernetesConfiguration
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesconfiguration := make(map[string]interface{})
	kubernetesconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesconfiguration["class_id"] = item.GetClassId()
	kubernetesconfiguration["kube_config"] = item.GetKubeConfig()
	kubernetesconfiguration["object_type"] = item.GetObjectType()

	kubernetesconfigurations = append(kubernetesconfigurations, kubernetesconfiguration)
	return kubernetesconfigurations
}
func flattenMapKubernetesContainerRuntimePolicyRelationship(p models.KubernetesContainerRuntimePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetescontainerruntimepolicyrelationships []map[string]interface{}
	var ret models.KubernetesContainerRuntimePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetescontainerruntimepolicyrelationship := make(map[string]interface{})
	kubernetescontainerruntimepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetescontainerruntimepolicyrelationship["class_id"] = item.GetClassId()
	kubernetescontainerruntimepolicyrelationship["moid"] = item.GetMoid()
	kubernetescontainerruntimepolicyrelationship["object_type"] = item.GetObjectType()
	kubernetescontainerruntimepolicyrelationship["selector"] = item.GetSelector()

	kubernetescontainerruntimepolicyrelationships = append(kubernetescontainerruntimepolicyrelationships, kubernetescontainerruntimepolicyrelationship)
	return kubernetescontainerruntimepolicyrelationships
}
func flattenMapKubernetesDaemonSetStatus(p models.KubernetesDaemonSetStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesdaemonsetstatuss []map[string]interface{}
	var ret models.KubernetesDaemonSetStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesdaemonsetstatus := make(map[string]interface{})
	kubernetesdaemonsetstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesdaemonsetstatus["class_id"] = item.GetClassId()
	kubernetesdaemonsetstatus["current_number_scheduled"] = item.GetCurrentNumberScheduled()
	kubernetesdaemonsetstatus["desired_number_scheduled"] = item.GetDesiredNumberScheduled()
	kubernetesdaemonsetstatus["number_available"] = item.GetNumberAvailable()
	kubernetesdaemonsetstatus["number_misscheduled"] = item.GetNumberMisscheduled()
	kubernetesdaemonsetstatus["number_ready"] = item.GetNumberReady()
	kubernetesdaemonsetstatus["object_type"] = item.GetObjectType()
	kubernetesdaemonsetstatus["observed_generation"] = item.GetObservedGeneration()
	kubernetesdaemonsetstatus["updated_number_scheduled"] = item.GetUpdatedNumberScheduled()

	kubernetesdaemonsetstatuss = append(kubernetesdaemonsetstatuss, kubernetesdaemonsetstatus)
	return kubernetesdaemonsetstatuss
}
func flattenMapKubernetesDeploymentStatus(p models.KubernetesDeploymentStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesdeploymentstatuss []map[string]interface{}
	var ret models.KubernetesDeploymentStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesdeploymentstatus := make(map[string]interface{})
	kubernetesdeploymentstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesdeploymentstatus["available_replicas"] = item.GetAvailableReplicas()
	kubernetesdeploymentstatus["class_id"] = item.GetClassId()
	kubernetesdeploymentstatus["object_type"] = item.GetObjectType()
	kubernetesdeploymentstatus["observed_generation"] = item.GetObservedGeneration()
	kubernetesdeploymentstatus["ready_replicas"] = item.GetReadyReplicas()
	kubernetesdeploymentstatus["replicas"] = item.GetReplicas()
	kubernetesdeploymentstatus["updated_replicas"] = item.GetUpdatedReplicas()

	kubernetesdeploymentstatuss = append(kubernetesdeploymentstatuss, kubernetesdeploymentstatus)
	return kubernetesdeploymentstatuss
}
func flattenMapKubernetesIngressStatus(p models.KubernetesIngressStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesingressstatuss []map[string]interface{}
	var ret models.KubernetesIngressStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesingressstatus := make(map[string]interface{})
	kubernetesingressstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesingressstatus["class_id"] = item.GetClassId()
	kubernetesingressstatus["load_balancer"] = (func(p models.KubernetesLoadBalancer, d *schema.ResourceData) []map[string]interface{} {
		var kubernetesloadbalancers []map[string]interface{}
		var ret models.KubernetesLoadBalancer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		kubernetesloadbalancer := make(map[string]interface{})
		kubernetesloadbalancer["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesloadbalancer["class_id"] = item.GetClassId()
		kubernetesloadbalancer["ip_addresses"] = item.GetIpAddresses()
		kubernetesloadbalancer["object_type"] = item.GetObjectType()

		kubernetesloadbalancers = append(kubernetesloadbalancers, kubernetesloadbalancer)
		return kubernetesloadbalancers
	})(item.GetLoadBalancer(), d)
	kubernetesingressstatus["object_type"] = item.GetObjectType()

	kubernetesingressstatuss = append(kubernetesingressstatuss, kubernetesingressstatus)
	return kubernetesingressstatuss
}
func flattenMapKubernetesNetworkPolicyRelationship(p models.KubernetesNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnetworkpolicyrelationships []map[string]interface{}
	var ret models.KubernetesNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesnetworkpolicyrelationship := make(map[string]interface{})
	kubernetesnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesnetworkpolicyrelationship["class_id"] = item.GetClassId()
	kubernetesnetworkpolicyrelationship["moid"] = item.GetMoid()
	kubernetesnetworkpolicyrelationship["object_type"] = item.GetObjectType()
	kubernetesnetworkpolicyrelationship["selector"] = item.GetSelector()

	kubernetesnetworkpolicyrelationships = append(kubernetesnetworkpolicyrelationships, kubernetesnetworkpolicyrelationship)
	return kubernetesnetworkpolicyrelationships
}
func flattenMapKubernetesNodeGroupProfileRelationship(p models.KubernetesNodeGroupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodegroupprofilerelationships []map[string]interface{}
	var ret models.KubernetesNodeGroupProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesnodegroupprofilerelationship := make(map[string]interface{})
	kubernetesnodegroupprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesnodegroupprofilerelationship["class_id"] = item.GetClassId()
	kubernetesnodegroupprofilerelationship["moid"] = item.GetMoid()
	kubernetesnodegroupprofilerelationship["object_type"] = item.GetObjectType()
	kubernetesnodegroupprofilerelationship["selector"] = item.GetSelector()

	kubernetesnodegroupprofilerelationships = append(kubernetesnodegroupprofilerelationships, kubernetesnodegroupprofilerelationship)
	return kubernetesnodegroupprofilerelationships
}
func flattenMapKubernetesNodeInfo(p models.KubernetesNodeInfo, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodeinfos []map[string]interface{}
	var ret models.KubernetesNodeInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesnodeinfo := make(map[string]interface{})
	kubernetesnodeinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesnodeinfo["architecture"] = item.GetArchitecture()
	kubernetesnodeinfo["boot_id"] = item.GetBootId()
	kubernetesnodeinfo["class_id"] = item.GetClassId()
	kubernetesnodeinfo["container_runtime_version"] = item.GetContainerRuntimeVersion()
	kubernetesnodeinfo["kernel_version"] = item.GetKernelVersion()
	kubernetesnodeinfo["kube_proxy_version"] = item.GetKubeProxyVersion()
	kubernetesnodeinfo["kubelet_version"] = item.GetKubeletVersion()
	kubernetesnodeinfo["machine_id"] = item.GetMachineId()
	kubernetesnodeinfo["object_type"] = item.GetObjectType()
	kubernetesnodeinfo["operating_system"] = item.GetOperatingSystem()
	kubernetesnodeinfo["os_image"] = item.GetOsImage()
	kubernetesnodeinfo["system_uuid"] = item.GetSystemUuid()

	kubernetesnodeinfos = append(kubernetesnodeinfos, kubernetesnodeinfo)
	return kubernetesnodeinfos
}
func flattenMapKubernetesNodeProfileRelationship(p models.KubernetesNodeProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodeprofilerelationships []map[string]interface{}
	var ret models.KubernetesNodeProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesnodeprofilerelationship := make(map[string]interface{})
	kubernetesnodeprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesnodeprofilerelationship["class_id"] = item.GetClassId()
	kubernetesnodeprofilerelationship["moid"] = item.GetMoid()
	kubernetesnodeprofilerelationship["object_type"] = item.GetObjectType()
	kubernetesnodeprofilerelationship["selector"] = item.GetSelector()

	kubernetesnodeprofilerelationships = append(kubernetesnodeprofilerelationships, kubernetesnodeprofilerelationship)
	return kubernetesnodeprofilerelationships
}
func flattenMapKubernetesNodeSpec(p models.KubernetesNodeSpec, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesnodespecs []map[string]interface{}
	var ret models.KubernetesNodeSpec
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesnodespec := make(map[string]interface{})
	kubernetesnodespec["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesnodespec["class_id"] = item.GetClassId()
	kubernetesnodespec["object_type"] = item.GetObjectType()
	kubernetesnodespec["pod_cidr"] = item.GetPodCidr()
	kubernetesnodespec["provider_id"] = item.GetProviderId()

	kubernetesnodespecs = append(kubernetesnodespecs, kubernetesnodespec)
	return kubernetesnodespecs
}
func flattenMapKubernetesObjectMeta(p models.KubernetesObjectMeta, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesobjectmetas []map[string]interface{}
	var ret models.KubernetesObjectMeta
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesobjectmeta := make(map[string]interface{})
	kubernetesobjectmeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesobjectmeta["class_id"] = item.GetClassId()
	kubernetesobjectmeta["creation_timestamp"] = item.GetCreationTimestamp()
	kubernetesobjectmeta["name"] = item.GetName()
	kubernetesobjectmeta["namespace"] = item.GetNamespace()
	kubernetesobjectmeta["object_type"] = item.GetObjectType()
	kubernetesobjectmeta["resource_version"] = item.GetResourceVersion()
	kubernetesobjectmeta["uuid"] = item.GetUuid()

	kubernetesobjectmetas = append(kubernetesobjectmetas, kubernetesobjectmeta)
	return kubernetesobjectmetas
}
func flattenMapKubernetesPodStatus(p models.KubernetesPodStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetespodstatuss []map[string]interface{}
	var ret models.KubernetesPodStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetespodstatus := make(map[string]interface{})
	kubernetespodstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetespodstatus["class_id"] = item.GetClassId()
	kubernetespodstatus["host_ip"] = item.GetHostIp()
	kubernetespodstatus["object_type"] = item.GetObjectType()
	kubernetespodstatus["phase"] = item.GetPhase()
	kubernetespodstatus["pod_ip"] = item.GetPodIp()
	kubernetespodstatus["qos_class"] = item.GetQosClass()
	kubernetespodstatus["start_time"] = item.GetStartTime()

	kubernetespodstatuss = append(kubernetespodstatuss, kubernetespodstatus)
	return kubernetespodstatuss
}
func flattenMapKubernetesProxyConfig(p models.KubernetesProxyConfig, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesproxyconfigs []map[string]interface{}
	var ret models.KubernetesProxyConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesproxyconfig := make(map[string]interface{})
	kubernetesproxyconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesproxyconfig["class_id"] = item.GetClassId()
	kubernetesproxyconfig["hostname"] = item.GetHostname()
	kubernetesproxyconfig["is_password_set"] = item.GetIsPasswordSet()
	kubernetesproxyconfig["object_type"] = item.GetObjectType()
	password_x, exists := d.GetOk("docker_http_proxy")
	if exists && password_x != nil {
		password_y := password_x.([]interface{})[0].(map[string]interface{})
		kubernetesproxyconfig["password"] = password_y["password"]
	}
	kubernetesproxyconfig["port"] = item.GetPort()
	kubernetesproxyconfig["protocol"] = item.GetProtocol()
	kubernetesproxyconfig["username"] = item.GetUsername()

	kubernetesproxyconfigs = append(kubernetesproxyconfigs, kubernetesproxyconfig)
	return kubernetesproxyconfigs
}
func flattenMapKubernetesServiceStatus(p models.KubernetesServiceStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesservicestatuss []map[string]interface{}
	var ret models.KubernetesServiceStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesservicestatus := make(map[string]interface{})
	kubernetesservicestatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesservicestatus["class_id"] = item.GetClassId()
	kubernetesservicestatus["load_balancer"] = (func(p models.KubernetesLoadBalancer, d *schema.ResourceData) []map[string]interface{} {
		var kubernetesloadbalancers []map[string]interface{}
		var ret models.KubernetesLoadBalancer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		kubernetesloadbalancer := make(map[string]interface{})
		kubernetesloadbalancer["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesloadbalancer["class_id"] = item.GetClassId()
		kubernetesloadbalancer["ip_addresses"] = item.GetIpAddresses()
		kubernetesloadbalancer["object_type"] = item.GetObjectType()

		kubernetesloadbalancers = append(kubernetesloadbalancers, kubernetesloadbalancer)
		return kubernetesloadbalancers
	})(item.GetLoadBalancer(), d)
	kubernetesservicestatus["object_type"] = item.GetObjectType()

	kubernetesservicestatuss = append(kubernetesservicestatuss, kubernetesservicestatus)
	return kubernetesservicestatuss
}
func flattenMapKubernetesStatefulSetStatus(p models.KubernetesStatefulSetStatus, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesstatefulsetstatuss []map[string]interface{}
	var ret models.KubernetesStatefulSetStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	kubernetesstatefulsetstatus := make(map[string]interface{})
	kubernetesstatefulsetstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesstatefulsetstatus["available_replicas"] = item.GetAvailableReplicas()
	kubernetesstatefulsetstatus["class_id"] = item.GetClassId()
	kubernetesstatefulsetstatus["collision_count"] = item.GetCollisionCount()
	kubernetesstatefulsetstatus["current_revision"] = item.GetCurrentRevision()
	kubernetesstatefulsetstatus["object_type"] = item.GetObjectType()
	kubernetesstatefulsetstatus["observed_generation"] = item.GetObservedGeneration()
	kubernetesstatefulsetstatus["ready_replicas"] = item.GetReadyReplicas()
	kubernetesstatefulsetstatus["replicas"] = item.GetReplicas()
	kubernetesstatefulsetstatus["update_revision"] = item.GetUpdateRevision()
	kubernetesstatefulsetstatus["updated_replicas"] = item.GetUpdatedReplicas()

	kubernetesstatefulsetstatuss = append(kubernetesstatefulsetstatuss, kubernetesstatefulsetstatus)
	return kubernetesstatefulsetstatuss
}
func flattenMapKubernetesSysConfigPolicyRelationship(p models.KubernetesSysConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetessysconfigpolicyrelationships []map[string]interface{}
	var ret models.KubernetesSysConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetessysconfigpolicyrelationship := make(map[string]interface{})
	kubernetessysconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetessysconfigpolicyrelationship["class_id"] = item.GetClassId()
	kubernetessysconfigpolicyrelationship["moid"] = item.GetMoid()
	kubernetessysconfigpolicyrelationship["object_type"] = item.GetObjectType()
	kubernetessysconfigpolicyrelationship["selector"] = item.GetSelector()

	kubernetessysconfigpolicyrelationships = append(kubernetessysconfigpolicyrelationships, kubernetessysconfigpolicyrelationship)
	return kubernetessysconfigpolicyrelationships
}
func flattenMapKubernetesTrustedRegistriesPolicyRelationship(p models.KubernetesTrustedRegistriesPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetestrustedregistriespolicyrelationships []map[string]interface{}
	var ret models.KubernetesTrustedRegistriesPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetestrustedregistriespolicyrelationship := make(map[string]interface{})
	kubernetestrustedregistriespolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetestrustedregistriespolicyrelationship["class_id"] = item.GetClassId()
	kubernetestrustedregistriespolicyrelationship["moid"] = item.GetMoid()
	kubernetestrustedregistriespolicyrelationship["object_type"] = item.GetObjectType()
	kubernetestrustedregistriespolicyrelationship["selector"] = item.GetSelector()

	kubernetestrustedregistriespolicyrelationships = append(kubernetestrustedregistriespolicyrelationships, kubernetestrustedregistriespolicyrelationship)
	return kubernetestrustedregistriespolicyrelationships
}
func flattenMapKubernetesVersionRelationship(p models.KubernetesVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesversionrelationships []map[string]interface{}
	var ret models.KubernetesVersionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesversionrelationship := make(map[string]interface{})
	kubernetesversionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesversionrelationship["class_id"] = item.GetClassId()
	kubernetesversionrelationship["moid"] = item.GetMoid()
	kubernetesversionrelationship["object_type"] = item.GetObjectType()
	kubernetesversionrelationship["selector"] = item.GetSelector()

	kubernetesversionrelationships = append(kubernetesversionrelationships, kubernetesversionrelationship)
	return kubernetesversionrelationships
}
func flattenMapKubernetesVersionPolicyRelationship(p models.KubernetesVersionPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesversionpolicyrelationships []map[string]interface{}
	var ret models.KubernetesVersionPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesversionpolicyrelationship := make(map[string]interface{})
	kubernetesversionpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesversionpolicyrelationship["class_id"] = item.GetClassId()
	kubernetesversionpolicyrelationship["moid"] = item.GetMoid()
	kubernetesversionpolicyrelationship["object_type"] = item.GetObjectType()
	kubernetesversionpolicyrelationship["selector"] = item.GetSelector()

	kubernetesversionpolicyrelationships = append(kubernetesversionpolicyrelationships, kubernetesversionpolicyrelationship)
	return kubernetesversionpolicyrelationships
}
func flattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(p models.KubernetesVirtualMachineInfraConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesvirtualmachineinfraconfigpolicyrelationships []map[string]interface{}
	var ret models.KubernetesVirtualMachineInfraConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesvirtualmachineinfraconfigpolicyrelationship := make(map[string]interface{})
	kubernetesvirtualmachineinfraconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesvirtualmachineinfraconfigpolicyrelationship["class_id"] = item.GetClassId()
	kubernetesvirtualmachineinfraconfigpolicyrelationship["moid"] = item.GetMoid()
	kubernetesvirtualmachineinfraconfigpolicyrelationship["object_type"] = item.GetObjectType()
	kubernetesvirtualmachineinfraconfigpolicyrelationship["selector"] = item.GetSelector()

	kubernetesvirtualmachineinfraconfigpolicyrelationships = append(kubernetesvirtualmachineinfraconfigpolicyrelationships, kubernetesvirtualmachineinfraconfigpolicyrelationship)
	return kubernetesvirtualmachineinfraconfigpolicyrelationships
}
func flattenMapKubernetesVirtualMachineInstanceTypeRelationship(p models.KubernetesVirtualMachineInstanceTypeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kubernetesvirtualmachineinstancetyperelationships []map[string]interface{}
	var ret models.KubernetesVirtualMachineInstanceTypeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kubernetesvirtualmachineinstancetyperelationship := make(map[string]interface{})
	kubernetesvirtualmachineinstancetyperelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kubernetesvirtualmachineinstancetyperelationship["class_id"] = item.GetClassId()
	kubernetesvirtualmachineinstancetyperelationship["moid"] = item.GetMoid()
	kubernetesvirtualmachineinstancetyperelationship["object_type"] = item.GetObjectType()
	kubernetesvirtualmachineinstancetyperelationship["selector"] = item.GetSelector()

	kubernetesvirtualmachineinstancetyperelationships = append(kubernetesvirtualmachineinstancetyperelationships, kubernetesvirtualmachineinstancetyperelationship)
	return kubernetesvirtualmachineinstancetyperelationships
}
func flattenMapKvmSessionRelationship(p models.KvmSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kvmsessionrelationships []map[string]interface{}
	var ret models.KvmSessionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kvmsessionrelationship := make(map[string]interface{})
	kvmsessionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kvmsessionrelationship["class_id"] = item.GetClassId()
	kvmsessionrelationship["moid"] = item.GetMoid()
	kvmsessionrelationship["object_type"] = item.GetObjectType()
	kvmsessionrelationship["selector"] = item.GetSelector()

	kvmsessionrelationships = append(kvmsessionrelationships, kvmsessionrelationship)
	return kvmsessionrelationships
}
func flattenMapKvmTunnelRelationship(p models.KvmTunnelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var kvmtunnelrelationships []map[string]interface{}
	var ret models.KvmTunnelRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	kvmtunnelrelationship := make(map[string]interface{})
	kvmtunnelrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	kvmtunnelrelationship["class_id"] = item.GetClassId()
	kvmtunnelrelationship["moid"] = item.GetMoid()
	kvmtunnelrelationship["object_type"] = item.GetObjectType()
	kvmtunnelrelationship["selector"] = item.GetSelector()

	kvmtunnelrelationships = append(kvmtunnelrelationships, kvmtunnelrelationship)
	return kvmtunnelrelationships
}
func flattenMapLicenseAccountLicenseDataRelationship(p models.LicenseAccountLicenseDataRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenseaccountlicensedatarelationships []map[string]interface{}
	var ret models.LicenseAccountLicenseDataRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licenseaccountlicensedatarelationship := make(map[string]interface{})
	licenseaccountlicensedatarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licenseaccountlicensedatarelationship["class_id"] = item.GetClassId()
	licenseaccountlicensedatarelationship["moid"] = item.GetMoid()
	licenseaccountlicensedatarelationship["object_type"] = item.GetObjectType()
	licenseaccountlicensedatarelationship["selector"] = item.GetSelector()

	licenseaccountlicensedatarelationships = append(licenseaccountlicensedatarelationships, licenseaccountlicensedatarelationship)
	return licenseaccountlicensedatarelationships
}
func flattenMapLicenseCustomerOpRelationship(p models.LicenseCustomerOpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licensecustomeroprelationships []map[string]interface{}
	var ret models.LicenseCustomerOpRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licensecustomeroprelationship := make(map[string]interface{})
	licensecustomeroprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licensecustomeroprelationship["class_id"] = item.GetClassId()
	licensecustomeroprelationship["moid"] = item.GetMoid()
	licensecustomeroprelationship["object_type"] = item.GetObjectType()
	licensecustomeroprelationship["selector"] = item.GetSelector()

	licensecustomeroprelationships = append(licensecustomeroprelationships, licensecustomeroprelationship)
	return licensecustomeroprelationships
}
func flattenMapLicenseIwoCustomerOpRelationship(p models.LicenseIwoCustomerOpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenseiwocustomeroprelationships []map[string]interface{}
	var ret models.LicenseIwoCustomerOpRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licenseiwocustomeroprelationship := make(map[string]interface{})
	licenseiwocustomeroprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licenseiwocustomeroprelationship["class_id"] = item.GetClassId()
	licenseiwocustomeroprelationship["moid"] = item.GetMoid()
	licenseiwocustomeroprelationship["object_type"] = item.GetObjectType()
	licenseiwocustomeroprelationship["selector"] = item.GetSelector()

	licenseiwocustomeroprelationships = append(licenseiwocustomeroprelationships, licenseiwocustomeroprelationship)
	return licenseiwocustomeroprelationships
}
func flattenMapLicenseIwoLicenseCountRelationship(p models.LicenseIwoLicenseCountRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenseiwolicensecountrelationships []map[string]interface{}
	var ret models.LicenseIwoLicenseCountRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licenseiwolicensecountrelationship := make(map[string]interface{})
	licenseiwolicensecountrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licenseiwolicensecountrelationship["class_id"] = item.GetClassId()
	licenseiwolicensecountrelationship["moid"] = item.GetMoid()
	licenseiwolicensecountrelationship["object_type"] = item.GetObjectType()
	licenseiwolicensecountrelationship["selector"] = item.GetSelector()

	licenseiwolicensecountrelationships = append(licenseiwolicensecountrelationships, licenseiwolicensecountrelationship)
	return licenseiwolicensecountrelationships
}
func flattenMapLicenseSmartlicenseTokenRelationship(p models.LicenseSmartlicenseTokenRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licensesmartlicensetokenrelationships []map[string]interface{}
	var ret models.LicenseSmartlicenseTokenRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licensesmartlicensetokenrelationship := make(map[string]interface{})
	licensesmartlicensetokenrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licensesmartlicensetokenrelationship["class_id"] = item.GetClassId()
	licensesmartlicensetokenrelationship["moid"] = item.GetMoid()
	licensesmartlicensetokenrelationship["object_type"] = item.GetObjectType()
	licensesmartlicensetokenrelationship["selector"] = item.GetSelector()

	licensesmartlicensetokenrelationships = append(licensesmartlicensetokenrelationships, licensesmartlicensetokenrelationship)
	return licensesmartlicensetokenrelationships
}
func flattenMapMacpoolBlock(p models.MacpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var macpoolblocks []map[string]interface{}
	var ret models.MacpoolBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	macpoolblock := make(map[string]interface{})
	macpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolblock["class_id"] = item.GetClassId()
	macpoolblock["from"] = item.GetFrom()
	macpoolblock["object_type"] = item.GetObjectType()
	macpoolblock["size"] = item.GetSize()
	macpoolblock["to"] = item.GetTo()

	macpoolblocks = append(macpoolblocks, macpoolblock)
	return macpoolblocks
}
func flattenMapMacpoolIdBlockRelationship(p models.MacpoolIdBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolidblockrelationships []map[string]interface{}
	var ret models.MacpoolIdBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolidblockrelationship := make(map[string]interface{})
	macpoolidblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolidblockrelationship["class_id"] = item.GetClassId()
	macpoolidblockrelationship["moid"] = item.GetMoid()
	macpoolidblockrelationship["object_type"] = item.GetObjectType()
	macpoolidblockrelationship["selector"] = item.GetSelector()

	macpoolidblockrelationships = append(macpoolidblockrelationships, macpoolidblockrelationship)
	return macpoolidblockrelationships
}
func flattenMapMacpoolLeaseRelationship(p models.MacpoolLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolleaserelationships []map[string]interface{}
	var ret models.MacpoolLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolleaserelationship := make(map[string]interface{})
	macpoolleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolleaserelationship["class_id"] = item.GetClassId()
	macpoolleaserelationship["moid"] = item.GetMoid()
	macpoolleaserelationship["object_type"] = item.GetObjectType()
	macpoolleaserelationship["selector"] = item.GetSelector()

	macpoolleaserelationships = append(macpoolleaserelationships, macpoolleaserelationship)
	return macpoolleaserelationships
}
func flattenMapMacpoolPoolRelationship(p models.MacpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolpoolrelationships []map[string]interface{}
	var ret models.MacpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolpoolrelationship := make(map[string]interface{})
	macpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolpoolrelationship["class_id"] = item.GetClassId()
	macpoolpoolrelationship["moid"] = item.GetMoid()
	macpoolpoolrelationship["object_type"] = item.GetObjectType()
	macpoolpoolrelationship["selector"] = item.GetSelector()

	macpoolpoolrelationships = append(macpoolpoolrelationships, macpoolpoolrelationship)
	return macpoolpoolrelationships
}
func flattenMapMacpoolPoolMemberRelationship(p models.MacpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolpoolmemberrelationships []map[string]interface{}
	var ret models.MacpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolpoolmemberrelationship := make(map[string]interface{})
	macpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolpoolmemberrelationship["class_id"] = item.GetClassId()
	macpoolpoolmemberrelationship["moid"] = item.GetMoid()
	macpoolpoolmemberrelationship["object_type"] = item.GetObjectType()
	macpoolpoolmemberrelationship["selector"] = item.GetSelector()

	macpoolpoolmemberrelationships = append(macpoolpoolmemberrelationships, macpoolpoolmemberrelationship)
	return macpoolpoolmemberrelationships
}
func flattenMapMacpoolUniverseRelationship(p models.MacpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpooluniverserelationships []map[string]interface{}
	var ret models.MacpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpooluniverserelationship := make(map[string]interface{})
	macpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpooluniverserelationship["class_id"] = item.GetClassId()
	macpooluniverserelationship["moid"] = item.GetMoid()
	macpooluniverserelationship["object_type"] = item.GetObjectType()
	macpooluniverserelationship["selector"] = item.GetSelector()

	macpooluniverserelationships = append(macpooluniverserelationships, macpooluniverserelationship)
	return macpooluniverserelationships
}
func flattenMapManagementControllerRelationship(p models.ManagementControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managementcontrollerrelationships []map[string]interface{}
	var ret models.ManagementControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	managementcontrollerrelationship := make(map[string]interface{})
	managementcontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	managementcontrollerrelationship["class_id"] = item.GetClassId()
	managementcontrollerrelationship["moid"] = item.GetMoid()
	managementcontrollerrelationship["object_type"] = item.GetObjectType()
	managementcontrollerrelationship["selector"] = item.GetSelector()

	managementcontrollerrelationships = append(managementcontrollerrelationships, managementcontrollerrelationship)
	return managementcontrollerrelationships
}
func flattenMapManagementEntityRelationship(p models.ManagementEntityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managemententityrelationships []map[string]interface{}
	var ret models.ManagementEntityRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	managemententityrelationship := make(map[string]interface{})
	managemententityrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	managemententityrelationship["class_id"] = item.GetClassId()
	managemententityrelationship["moid"] = item.GetMoid()
	managemententityrelationship["object_type"] = item.GetObjectType()
	managemententityrelationship["selector"] = item.GetSelector()

	managemententityrelationships = append(managemententityrelationships, managemententityrelationship)
	return managemententityrelationships
}
func flattenMapMemoryArrayRelationship(p models.MemoryArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrelationships []map[string]interface{}
	var ret models.MemoryArrayRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memoryarrayrelationship := make(map[string]interface{})
	memoryarrayrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memoryarrayrelationship["class_id"] = item.GetClassId()
	memoryarrayrelationship["moid"] = item.GetMoid()
	memoryarrayrelationship["object_type"] = item.GetObjectType()
	memoryarrayrelationship["selector"] = item.GetSelector()

	memoryarrayrelationships = append(memoryarrayrelationships, memoryarrayrelationship)
	return memoryarrayrelationships
}
func flattenMapMemoryPersistentMemoryConfigResultRelationship(p models.MemoryPersistentMemoryConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigresultrelationships []map[string]interface{}
	var ret models.MemoryPersistentMemoryConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memorypersistentmemoryconfigresultrelationship := make(map[string]interface{})
	memorypersistentmemoryconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemoryconfigresultrelationship["class_id"] = item.GetClassId()
	memorypersistentmemoryconfigresultrelationship["moid"] = item.GetMoid()
	memorypersistentmemoryconfigresultrelationship["object_type"] = item.GetObjectType()
	memorypersistentmemoryconfigresultrelationship["selector"] = item.GetSelector()

	memorypersistentmemoryconfigresultrelationships = append(memorypersistentmemoryconfigresultrelationships, memorypersistentmemoryconfigresultrelationship)
	return memorypersistentmemoryconfigresultrelationships
}
func flattenMapMemoryPersistentMemoryConfigurationRelationship(p models.MemoryPersistentMemoryConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigurationrelationships []map[string]interface{}
	var ret models.MemoryPersistentMemoryConfigurationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memorypersistentmemoryconfigurationrelationship := make(map[string]interface{})
	memorypersistentmemoryconfigurationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemoryconfigurationrelationship["class_id"] = item.GetClassId()
	memorypersistentmemoryconfigurationrelationship["moid"] = item.GetMoid()
	memorypersistentmemoryconfigurationrelationship["object_type"] = item.GetObjectType()
	memorypersistentmemoryconfigurationrelationship["selector"] = item.GetSelector()

	memorypersistentmemoryconfigurationrelationships = append(memorypersistentmemoryconfigurationrelationships, memorypersistentmemoryconfigurationrelationship)
	return memorypersistentmemoryconfigurationrelationships
}
func flattenMapMemoryPersistentMemoryLocalSecurity(p models.MemoryPersistentMemoryLocalSecurity, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylocalsecuritys []map[string]interface{}
	var ret models.MemoryPersistentMemoryLocalSecurity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	memorypersistentmemorylocalsecurity := make(map[string]interface{})
	memorypersistentmemorylocalsecurity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemorylocalsecurity["class_id"] = item.GetClassId()
	memorypersistentmemorylocalsecurity["enabled"] = item.GetEnabled()
	memorypersistentmemorylocalsecurity["is_secure_passphrase_set"] = item.GetIsSecurePassphraseSet()
	memorypersistentmemorylocalsecurity["object_type"] = item.GetObjectType()
	secure_passphrase_x, exists := d.GetOk("local_security")
	if exists && secure_passphrase_x != nil {
		secure_passphrase_y := secure_passphrase_x.([]interface{})[0].(map[string]interface{})
		memorypersistentmemorylocalsecurity["secure_passphrase"] = secure_passphrase_y["secure_passphrase"]
	}

	memorypersistentmemorylocalsecuritys = append(memorypersistentmemorylocalsecuritys, memorypersistentmemorylocalsecurity)
	return memorypersistentmemorylocalsecuritys
}
func flattenMapMemoryPersistentMemoryRegionRelationship(p models.MemoryPersistentMemoryRegionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrelationships []map[string]interface{}
	var ret models.MemoryPersistentMemoryRegionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memorypersistentmemoryregionrelationship := make(map[string]interface{})
	memorypersistentmemoryregionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemoryregionrelationship["class_id"] = item.GetClassId()
	memorypersistentmemoryregionrelationship["moid"] = item.GetMoid()
	memorypersistentmemoryregionrelationship["object_type"] = item.GetObjectType()
	memorypersistentmemoryregionrelationship["selector"] = item.GetSelector()

	memorypersistentmemoryregionrelationships = append(memorypersistentmemoryregionrelationships, memorypersistentmemoryregionrelationship)
	return memorypersistentmemoryregionrelationships
}
func flattenMapMoBaseMoRelationship(p models.MoBaseMoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorelationships []map[string]interface{}
	var ret models.MoBaseMoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	mobasemorelationship := make(map[string]interface{})
	mobasemorelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	mobasemorelationship["class_id"] = item.GetClassId()
	mobasemorelationship["moid"] = item.GetMoid()
	mobasemorelationship["object_type"] = item.GetObjectType()
	mobasemorelationship["selector"] = item.GetSelector()

	mobasemorelationships = append(mobasemorelationships, mobasemorelationship)
	return mobasemorelationships
}
func flattenMapMoVersionContext(p models.MoVersionContext, d *schema.ResourceData) []map[string]interface{} {
	var moversioncontexts []map[string]interface{}
	var ret models.MoVersionContext
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	moversioncontext := make(map[string]interface{})
	moversioncontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	moversioncontext["class_id"] = item.GetClassId()
	moversioncontext["interested_mos"] = (func(p []models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
		var momorefs []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.GetClassId()
			momoref["moid"] = item.GetMoid()
			momoref["object_type"] = item.GetObjectType()
			momoref["selector"] = item.GetSelector()
			momorefs = append(momorefs, momoref)
		}
		return momorefs
	})(item.GetInterestedMos(), d)
	moversioncontext["object_type"] = item.GetObjectType()
	moversioncontext["ref_mo"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
		var momorefs []map[string]interface{}
		var ret models.MoMoRef
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		momoref := make(map[string]interface{})
		momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		momoref["class_id"] = item.GetClassId()
		momoref["moid"] = item.GetMoid()
		momoref["object_type"] = item.GetObjectType()
		momoref["selector"] = item.GetSelector()

		momorefs = append(momorefs, momoref)
		return momorefs
	})(item.GetRefMo(), d)
	moversioncontext["timestamp"] = item.GetTimestamp()
	moversioncontext["nr_version"] = item.GetVersion()
	moversioncontext["version_type"] = item.GetVersionType()

	moversioncontexts = append(moversioncontexts, moversioncontext)
	return moversioncontexts
}
func flattenMapNetworkElementRelationship(p models.NetworkElementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrelationships []map[string]interface{}
	var ret models.NetworkElementRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	networkelementrelationship := make(map[string]interface{})
	networkelementrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	networkelementrelationship["class_id"] = item.GetClassId()
	networkelementrelationship["moid"] = item.GetMoid()
	networkelementrelationship["object_type"] = item.GetObjectType()
	networkelementrelationship["selector"] = item.GetSelector()

	networkelementrelationships = append(networkelementrelationships, networkelementrelationship)
	return networkelementrelationships
}
func flattenMapNetworkFcZoneInfoRelationship(p models.NetworkFcZoneInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkfczoneinforelationships []map[string]interface{}
	var ret models.NetworkFcZoneInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	networkfczoneinforelationship := make(map[string]interface{})
	networkfczoneinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	networkfczoneinforelationship["class_id"] = item.GetClassId()
	networkfczoneinforelationship["moid"] = item.GetMoid()
	networkfczoneinforelationship["object_type"] = item.GetObjectType()
	networkfczoneinforelationship["selector"] = item.GetSelector()

	networkfczoneinforelationships = append(networkfczoneinforelationships, networkfczoneinforelationship)
	return networkfczoneinforelationships
}
func flattenMapNetworkVlanPortInfoRelationship(p models.NetworkVlanPortInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkvlanportinforelationships []map[string]interface{}
	var ret models.NetworkVlanPortInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	networkvlanportinforelationship := make(map[string]interface{})
	networkvlanportinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	networkvlanportinforelationship["class_id"] = item.GetClassId()
	networkvlanportinforelationship["moid"] = item.GetMoid()
	networkvlanportinforelationship["object_type"] = item.GetObjectType()
	networkvlanportinforelationship["selector"] = item.GetSelector()

	networkvlanportinforelationships = append(networkvlanportinforelationships, networkvlanportinforelationship)
	return networkvlanportinforelationships
}
func flattenMapNiaapiNewReleaseDetail(p models.NiaapiNewReleaseDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapinewreleasedetails []map[string]interface{}
	var ret models.NiaapiNewReleaseDetail
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niaapinewreleasedetail := make(map[string]interface{})
	niaapinewreleasedetail["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niaapinewreleasedetail["class_id"] = item.GetClassId()
	niaapinewreleasedetail["description"] = item.GetDescription()
	niaapinewreleasedetail["link"] = item.GetLink()
	niaapinewreleasedetail["object_type"] = item.GetObjectType()
	niaapinewreleasedetail["release_note_link"] = item.GetReleaseNoteLink()
	niaapinewreleasedetail["release_note_link_title"] = item.GetReleaseNoteLinkTitle()
	niaapinewreleasedetail["software_download_link"] = item.GetSoftwareDownloadLink()
	niaapinewreleasedetail["software_download_link_title"] = item.GetSoftwareDownloadLinkTitle()
	niaapinewreleasedetail["title"] = item.GetTitle()
	niaapinewreleasedetail["nr_version"] = item.GetVersion()

	niaapinewreleasedetails = append(niaapinewreleasedetails, niaapinewreleasedetail)
	return niaapinewreleasedetails
}
func flattenMapNiaapiVersionRegexPlatform(p models.NiaapiVersionRegexPlatform, d *schema.ResourceData) []map[string]interface{} {
	var niaapiversionregexplatforms []map[string]interface{}
	var ret models.NiaapiVersionRegexPlatform
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niaapiversionregexplatform := make(map[string]interface{})
	niaapiversionregexplatform["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niaapiversionregexplatform["anyllregex"] = item.GetAnyllregex()
	niaapiversionregexplatform["class_id"] = item.GetClassId()
	niaapiversionregexplatform["currentlltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.GetClassId()
		niaapisoftwareregex["object_type"] = item.GetObjectType()
		niaapisoftwareregex["regex"] = item.GetRegex()
		niaapisoftwareregex["software_version"] = item.GetSoftwareVersion()

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetCurrentlltrain(), d)
	niaapiversionregexplatform["latestsltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.GetClassId()
		niaapisoftwareregex["object_type"] = item.GetObjectType()
		niaapisoftwareregex["regex"] = item.GetRegex()
		niaapisoftwareregex["software_version"] = item.GetSoftwareVersion()

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetLatestsltrain(), d)
	niaapiversionregexplatform["object_type"] = item.GetObjectType()
	niaapiversionregexplatform["sltrain"] = (func(p []models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			niaapisoftwareregex := make(map[string]interface{})
			niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			niaapisoftwareregex["class_id"] = item.GetClassId()
			niaapisoftwareregex["object_type"] = item.GetObjectType()
			niaapisoftwareregex["regex"] = item.GetRegex()
			niaapisoftwareregex["software_version"] = item.GetSoftwareVersion()
			niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		}
		return niaapisoftwareregexs
	})(item.GetSltrain(), d)
	niaapiversionregexplatform["upcominglltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.GetClassId()
		niaapisoftwareregex["object_type"] = item.GetObjectType()
		niaapisoftwareregex["regex"] = item.GetRegex()
		niaapisoftwareregex["software_version"] = item.GetSoftwareVersion()

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetUpcominglltrain(), d)

	niaapiversionregexplatforms = append(niaapiversionregexplatforms, niaapiversionregexplatform)
	return niaapiversionregexplatforms
}
func flattenMapNiatelemetryBootflashDetails(p models.NiatelemetryBootflashDetails, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrybootflashdetailss []map[string]interface{}
	var ret models.NiatelemetryBootflashDetails
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrybootflashdetails := make(map[string]interface{})
	niatelemetrybootflashdetails["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrybootflashdetails["class_id"] = item.GetClassId()
	niatelemetrybootflashdetails["fw_rev"] = item.GetFwRev()
	niatelemetrybootflashdetails["model_type"] = item.GetModelType()
	niatelemetrybootflashdetails["object_type"] = item.GetObjectType()
	niatelemetrybootflashdetails["serial"] = item.GetSerial()

	niatelemetrybootflashdetailss = append(niatelemetrybootflashdetailss, niatelemetrybootflashdetails)
	return niatelemetrybootflashdetailss
}
func flattenMapNiatelemetryDiskinfo(p models.NiatelemetryDiskinfo, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrydiskinfos []map[string]interface{}
	var ret models.NiatelemetryDiskinfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrydiskinfo := make(map[string]interface{})
	niatelemetrydiskinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrydiskinfo["class_id"] = item.GetClassId()
	niatelemetrydiskinfo["free"] = item.GetFree()
	niatelemetrydiskinfo["name"] = item.GetName()
	niatelemetrydiskinfo["object_type"] = item.GetObjectType()
	niatelemetrydiskinfo["total"] = item.GetTotal()
	niatelemetrydiskinfo["used"] = item.GetUsed()

	niatelemetrydiskinfos = append(niatelemetrydiskinfos, niatelemetrydiskinfo)
	return niatelemetrydiskinfos
}
func flattenMapNiatelemetryInterface(p models.NiatelemetryInterface, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetryinterfaces []map[string]interface{}
	var ret models.NiatelemetryInterface
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetryinterface := make(map[string]interface{})
	niatelemetryinterface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetryinterface["class_id"] = item.GetClassId()
	niatelemetryinterface["interface_down_count"] = item.GetInterfaceDownCount()
	niatelemetryinterface["interface_up_count"] = item.GetInterfaceUpCount()
	niatelemetryinterface["object_type"] = item.GetObjectType()

	niatelemetryinterfaces = append(niatelemetryinterfaces, niatelemetryinterface)
	return niatelemetryinterfaces
}
func flattenMapNiatelemetryNexusDashboardsRelationship(p models.NiatelemetryNexusDashboardsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynexusdashboardsrelationships []map[string]interface{}
	var ret models.NiatelemetryNexusDashboardsRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	niatelemetrynexusdashboardsrelationship := make(map[string]interface{})
	niatelemetrynexusdashboardsrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynexusdashboardsrelationship["class_id"] = item.GetClassId()
	niatelemetrynexusdashboardsrelationship["moid"] = item.GetMoid()
	niatelemetrynexusdashboardsrelationship["object_type"] = item.GetObjectType()
	niatelemetrynexusdashboardsrelationship["selector"] = item.GetSelector()

	niatelemetrynexusdashboardsrelationships = append(niatelemetrynexusdashboardsrelationships, niatelemetrynexusdashboardsrelationship)
	return niatelemetrynexusdashboardsrelationships
}
func flattenMapNiatelemetryNiaInventoryRelationship(p models.NiatelemetryNiaInventoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetryniainventoryrelationships []map[string]interface{}
	var ret models.NiatelemetryNiaInventoryRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	niatelemetryniainventoryrelationship := make(map[string]interface{})
	niatelemetryniainventoryrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetryniainventoryrelationship["class_id"] = item.GetClassId()
	niatelemetryniainventoryrelationship["moid"] = item.GetMoid()
	niatelemetryniainventoryrelationship["object_type"] = item.GetObjectType()
	niatelemetryniainventoryrelationship["selector"] = item.GetSelector()

	niatelemetryniainventoryrelationships = append(niatelemetryniainventoryrelationships, niatelemetryniainventoryrelationship)
	return niatelemetryniainventoryrelationships
}
func flattenMapNiatelemetryNiaLicenseStateRelationship(p models.NiatelemetryNiaLicenseStateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynialicensestaterelationships []map[string]interface{}
	var ret models.NiatelemetryNiaLicenseStateRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	niatelemetrynialicensestaterelationship := make(map[string]interface{})
	niatelemetrynialicensestaterelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynialicensestaterelationship["class_id"] = item.GetClassId()
	niatelemetrynialicensestaterelationship["moid"] = item.GetMoid()
	niatelemetrynialicensestaterelationship["object_type"] = item.GetObjectType()
	niatelemetrynialicensestaterelationship["selector"] = item.GetSelector()

	niatelemetrynialicensestaterelationships = append(niatelemetrynialicensestaterelationships, niatelemetrynialicensestaterelationship)
	return niatelemetrynialicensestaterelationships
}
func flattenMapNiatelemetryNvePacketCounters(p models.NiatelemetryNvePacketCounters, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynvepacketcounterss []map[string]interface{}
	var ret models.NiatelemetryNvePacketCounters
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrynvepacketcounters := make(map[string]interface{})
	niatelemetrynvepacketcounters["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynvepacketcounters["class_id"] = item.GetClassId()
	niatelemetrynvepacketcounters["mcast_inpkts"] = item.GetMcastInpkts()
	niatelemetrynvepacketcounters["mcast_outbytes"] = item.GetMcastOutbytes()
	niatelemetrynvepacketcounters["object_type"] = item.GetObjectType()
	niatelemetrynvepacketcounters["ucast_inpkts"] = item.GetUcastInpkts()
	niatelemetrynvepacketcounters["ucast_outpkts"] = item.GetUcastOutpkts()

	niatelemetrynvepacketcounterss = append(niatelemetrynvepacketcounterss, niatelemetrynvepacketcounters)
	return niatelemetrynvepacketcounterss
}
func flattenMapNiatelemetryNveVni(p models.NiatelemetryNveVni, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynvevnis []map[string]interface{}
	var ret models.NiatelemetryNveVni
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrynvevni := make(map[string]interface{})
	niatelemetrynvevni["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynvevni["class_id"] = item.GetClassId()
	niatelemetrynvevni["cp_vni_count"] = item.GetCpVniCount()
	niatelemetrynvevni["cp_vni_down"] = item.GetCpVniDown()
	niatelemetrynvevni["cp_vni_up"] = item.GetCpVniUp()
	niatelemetrynvevni["dp_vni_count"] = item.GetDpVniCount()
	niatelemetrynvevni["dp_vni_down"] = item.GetDpVniDown()
	niatelemetrynvevni["dp_vni_up"] = item.GetDpVniUp()
	niatelemetrynvevni["object_type"] = item.GetObjectType()

	niatelemetrynvevnis = append(niatelemetrynvevnis, niatelemetrynvevni)
	return niatelemetrynvevnis
}
func flattenMapNiatelemetryNxosBgpMvpn(p models.NiatelemetryNxosBgpMvpn, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynxosbgpmvpns []map[string]interface{}
	var ret models.NiatelemetryNxosBgpMvpn
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrynxosbgpmvpn := make(map[string]interface{})
	niatelemetrynxosbgpmvpn["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynxosbgpmvpn["capable_peers"] = item.GetCapablePeers()
	niatelemetrynxosbgpmvpn["class_id"] = item.GetClassId()
	niatelemetrynxosbgpmvpn["configured_peers"] = item.GetConfiguredPeers()
	niatelemetrynxosbgpmvpn["memory_used"] = item.GetMemoryUsed()
	niatelemetrynxosbgpmvpn["number_of_cluster_lists"] = item.GetNumberOfClusterLists()
	niatelemetrynxosbgpmvpn["number_of_communities"] = item.GetNumberOfCommunities()
	niatelemetrynxosbgpmvpn["object_type"] = item.GetObjectType()
	niatelemetrynxosbgpmvpn["table_version"] = item.GetTableVersion()
	niatelemetrynxosbgpmvpn["total_networks"] = item.GetTotalNetworks()
	niatelemetrynxosbgpmvpn["total_paths"] = item.GetTotalPaths()

	niatelemetrynxosbgpmvpns = append(niatelemetrynxosbgpmvpns, niatelemetrynxosbgpmvpn)
	return niatelemetrynxosbgpmvpns
}
func flattenMapNiatelemetryNxosVtp(p models.NiatelemetryNxosVtp, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynxosvtps []map[string]interface{}
	var ret models.NiatelemetryNxosVtp
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrynxosvtp := make(map[string]interface{})
	niatelemetrynxosvtp["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynxosvtp["class_id"] = item.GetClassId()
	niatelemetrynxosvtp["object_type"] = item.GetObjectType()
	niatelemetrynxosvtp["oper_mode"] = item.GetOperMode()
	niatelemetrynxosvtp["pruning_mode"] = item.GetPruningMode()
	niatelemetrynxosvtp["running_version"] = item.GetRunningVersion()
	niatelemetrynxosvtp["trap_enabled"] = item.GetTrapEnabled()
	niatelemetrynxosvtp["v2_mode"] = item.GetV2Mode()
	niatelemetrynxosvtp["nr_version"] = item.GetVersion()

	niatelemetrynxosvtps = append(niatelemetrynxosvtps, niatelemetrynxosvtp)
	return niatelemetrynxosvtps
}
func flattenMapNiatelemetrySmartLicense(p models.NiatelemetrySmartLicense, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrysmartlicenses []map[string]interface{}
	var ret models.NiatelemetrySmartLicense
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrysmartlicense := make(map[string]interface{})
	niatelemetrysmartlicense["active_mode"] = item.GetActiveMode()
	niatelemetrysmartlicense["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrysmartlicense["auth_status"] = item.GetAuthStatus()
	niatelemetrysmartlicense["class_id"] = item.GetClassId()
	niatelemetrysmartlicense["license_udi"] = item.GetLicenseUdi()
	niatelemetrysmartlicense["object_type"] = item.GetObjectType()
	niatelemetrysmartlicense["smart_account"] = item.GetSmartAccount()

	niatelemetrysmartlicenses = append(niatelemetrysmartlicenses, niatelemetrysmartlicense)
	return niatelemetrysmartlicenses
}
func flattenMapOnpremSchedule(p models.OnpremSchedule, d *schema.ResourceData) []map[string]interface{} {
	var onpremschedules []map[string]interface{}
	var ret models.OnpremSchedule
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	onpremschedule := make(map[string]interface{})
	onpremschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	onpremschedule["class_id"] = item.GetClassId()
	onpremschedule["day_of_month"] = item.GetDayOfMonth()
	onpremschedule["day_of_week"] = item.GetDayOfWeek()
	onpremschedule["month_of_year"] = item.GetMonthOfYear()
	onpremschedule["object_type"] = item.GetObjectType()
	onpremschedule["repeat_interval"] = item.GetRepeatInterval()
	onpremschedule["time_of_day"] = item.GetTimeOfDay()
	onpremschedule["time_zone"] = item.GetTimeZone()
	onpremschedule["week_of_month"] = item.GetWeekOfMonth()

	onpremschedules = append(onpremschedules, onpremschedule)
	return onpremschedules
}
func flattenMapOnpremUpgradePhase(p models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	var ret models.OnpremUpgradePhase
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	onpremupgradephase := make(map[string]interface{})
	onpremupgradephase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	onpremupgradephase["class_id"] = item.GetClassId()
	onpremupgradephase["elapsed_time"] = item.GetElapsedTime()
	onpremupgradephase["end_time"] = item.GetEndTime()
	onpremupgradephase["failed"] = item.GetFailed()
	onpremupgradephase["message"] = item.GetMessage()
	onpremupgradephase["name"] = item.GetName()
	onpremupgradephase["object_type"] = item.GetObjectType()
	onpremupgradephase["start_time"] = item.GetStartTime()

	onpremupgradephases = append(onpremupgradephases, onpremupgradephase)
	return onpremupgradephases
}
func flattenMapOrganizationOrganizationRelationship(p models.OrganizationOrganizationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrelationships []map[string]interface{}
	var ret models.OrganizationOrganizationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	organizationorganizationrelationship := make(map[string]interface{})
	organizationorganizationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	organizationorganizationrelationship["class_id"] = item.GetClassId()
	organizationorganizationrelationship["moid"] = item.GetMoid()
	organizationorganizationrelationship["object_type"] = item.GetObjectType()
	organizationorganizationrelationship["selector"] = item.GetSelector()

	organizationorganizationrelationships = append(organizationorganizationrelationships, organizationorganizationrelationship)
	return organizationorganizationrelationships
}
func flattenMapOsAnswers(p models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {
	var osanswerss []map[string]interface{}
	var ret models.OsAnswers
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	osanswers := make(map[string]interface{})
	osanswers["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	answer_file_x, exists := d.GetOk("answers")
	if exists && answer_file_x != nil {
		answer_file_y := answer_file_x.([]interface{})[0].(map[string]interface{})
		osanswers["answer_file"] = answer_file_y["answer_file"]
	}
	osanswers["class_id"] = item.GetClassId()
	osanswers["hostname"] = item.GetHostname()
	osanswers["ip_config_type"] = item.GetIpConfigType()
	osanswers["ip_configuration"] = (func(p models.OsIpConfiguration, d *schema.ResourceData) []map[string]interface{} {
		var osipconfigurations []map[string]interface{}
		var ret models.OsIpConfiguration
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		osipconfiguration := make(map[string]interface{})
		osipconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osipconfiguration["class_id"] = item.GetClassId()
		osipconfiguration["object_type"] = item.GetObjectType()

		osipconfigurations = append(osipconfigurations, osipconfiguration)
		return osipconfigurations
	})(item.GetIpConfiguration(), d)
	osanswers["is_answer_file_set"] = item.GetIsAnswerFileSet()
	osanswers["is_root_password_crypted"] = item.GetIsRootPasswordCrypted()
	osanswers["is_root_password_set"] = item.GetIsRootPasswordSet()
	osanswers["nameserver"] = item.GetNameserver()
	osanswers["object_type"] = item.GetObjectType()
	osanswers["product_key"] = item.GetProductKey()
	root_password_x, exists := d.GetOk("answers")
	if exists && root_password_x != nil {
		root_password_y := root_password_x.([]interface{})[0].(map[string]interface{})
		osanswers["root_password"] = root_password_y["root_password"]
	}
	osanswers["nr_source"] = item.GetSource()

	osanswerss = append(osanswerss, osanswers)
	return osanswerss
}
func flattenMapOsCatalogRelationship(p models.OsCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var oscatalogrelationships []map[string]interface{}
	var ret models.OsCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	oscatalogrelationship := make(map[string]interface{})
	oscatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	oscatalogrelationship["class_id"] = item.GetClassId()
	oscatalogrelationship["moid"] = item.GetMoid()
	oscatalogrelationship["object_type"] = item.GetObjectType()
	oscatalogrelationship["selector"] = item.GetSelector()

	oscatalogrelationships = append(oscatalogrelationships, oscatalogrelationship)
	return oscatalogrelationships
}
func flattenMapOsConfigurationFileRelationship(p models.OsConfigurationFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerelationships []map[string]interface{}
	var ret models.OsConfigurationFileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	osconfigurationfilerelationship := make(map[string]interface{})
	osconfigurationfilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	osconfigurationfilerelationship["class_id"] = item.GetClassId()
	osconfigurationfilerelationship["moid"] = item.GetMoid()
	osconfigurationfilerelationship["object_type"] = item.GetObjectType()
	osconfigurationfilerelationship["selector"] = item.GetSelector()

	osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	return osconfigurationfilerelationships
}
func flattenMapOsGlobalConfig(p models.OsGlobalConfig, d *schema.ResourceData) []map[string]interface{} {
	var osglobalconfigs []map[string]interface{}
	var ret models.OsGlobalConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	osglobalconfig := make(map[string]interface{})
	osglobalconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	osglobalconfig["class_id"] = item.GetClassId()
	osglobalconfig["configuration_file_name"] = item.GetConfigurationFileName()
	osglobalconfig["configuration_source"] = item.GetConfigurationSource()
	osglobalconfig["install_method"] = item.GetInstallMethod()
	osglobalconfig["install_target_type"] = item.GetInstallTargetType()
	osglobalconfig["object_type"] = item.GetObjectType()
	osglobalconfig["operating_system_parameters"] = (func(p models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {
		var osoperatingsystemparameterss []map[string]interface{}
		var ret models.OsOperatingSystemParameters
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		osoperatingsystemparameters := make(map[string]interface{})
		osoperatingsystemparameters["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osoperatingsystemparameters["class_id"] = item.GetClassId()
		osoperatingsystemparameters["object_type"] = item.GetObjectType()

		osoperatingsystemparameterss = append(osoperatingsystemparameterss, osoperatingsystemparameters)
		return osoperatingsystemparameterss
	})(item.GetOperatingSystemParameters(), d)
	osglobalconfig["os_image_name"] = item.GetOsImageName()
	osglobalconfig["scu_image_name"] = item.GetScuImageName()
	osglobalconfig["windows_edition"] = item.GetWindowsEdition()

	osglobalconfigs = append(osglobalconfigs, osglobalconfig)
	return osglobalconfigs
}
func flattenMapOsInstallTarget(p models.OsInstallTarget, d *schema.ResourceData) []map[string]interface{} {
	var osinstalltargets []map[string]interface{}
	var ret models.OsInstallTarget
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	osinstalltarget := make(map[string]interface{})
	osinstalltarget["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	osinstalltarget["class_id"] = item.GetClassId()
	osinstalltarget["object_type"] = item.GetObjectType()

	osinstalltargets = append(osinstalltargets, osinstalltarget)
	return osinstalltargets
}
func flattenMapOsOperatingSystemParameters(p models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {
	var osoperatingsystemparameterss []map[string]interface{}
	var ret models.OsOperatingSystemParameters
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	osoperatingsystemparameters := make(map[string]interface{})
	osoperatingsystemparameters["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	osoperatingsystemparameters["class_id"] = item.GetClassId()
	osoperatingsystemparameters["object_type"] = item.GetObjectType()

	osoperatingsystemparameterss = append(osoperatingsystemparameterss, osoperatingsystemparameters)
	return osoperatingsystemparameterss
}
func flattenMapPciSwitchRelationship(p models.PciSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrelationships []map[string]interface{}
	var ret models.PciSwitchRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	pciswitchrelationship := make(map[string]interface{})
	pciswitchrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pciswitchrelationship["class_id"] = item.GetClassId()
	pciswitchrelationship["moid"] = item.GetMoid()
	pciswitchrelationship["object_type"] = item.GetObjectType()
	pciswitchrelationship["selector"] = item.GetSelector()

	pciswitchrelationships = append(pciswitchrelationships, pciswitchrelationship)
	return pciswitchrelationships
}
func flattenMapPkixDistinguishedName(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
	var pkixdistinguishednames []map[string]interface{}
	var ret models.PkixDistinguishedName
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	pkixdistinguishedname := make(map[string]interface{})
	pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pkixdistinguishedname["class_id"] = item.GetClassId()
	pkixdistinguishedname["common_name"] = item.GetCommonName()
	pkixdistinguishedname["country"] = item.GetCountry()
	pkixdistinguishedname["locality"] = item.GetLocality()
	pkixdistinguishedname["object_type"] = item.GetObjectType()
	pkixdistinguishedname["organization"] = item.GetOrganization()
	pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
	pkixdistinguishedname["state"] = item.GetState()

	pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
	return pkixdistinguishednames
}
func flattenMapPkixKeyGenerationSpec(p models.PkixKeyGenerationSpec, d *schema.ResourceData) []map[string]interface{} {
	var pkixkeygenerationspecs []map[string]interface{}
	var ret models.PkixKeyGenerationSpec
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	pkixkeygenerationspec := make(map[string]interface{})
	pkixkeygenerationspec["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pkixkeygenerationspec["class_id"] = item.GetClassId()
	pkixkeygenerationspec["name"] = item.GetName()
	pkixkeygenerationspec["object_type"] = item.GetObjectType()

	pkixkeygenerationspecs = append(pkixkeygenerationspecs, pkixkeygenerationspec)
	return pkixkeygenerationspecs
}
func flattenMapPkixSubjectAlternateName(p models.PkixSubjectAlternateName, d *schema.ResourceData) []map[string]interface{} {
	var pkixsubjectalternatenames []map[string]interface{}
	var ret models.PkixSubjectAlternateName
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	pkixsubjectalternatename := make(map[string]interface{})
	pkixsubjectalternatename["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pkixsubjectalternatename["class_id"] = item.GetClassId()
	pkixsubjectalternatename["dns_name"] = item.GetDnsName()
	pkixsubjectalternatename["email_address"] = item.GetEmailAddress()
	pkixsubjectalternatename["ip_address"] = item.GetIpAddress()
	pkixsubjectalternatename["object_type"] = item.GetObjectType()
	pkixsubjectalternatename["uri"] = item.GetUri()

	pkixsubjectalternatenames = append(pkixsubjectalternatenames, pkixsubjectalternatename)
	return pkixsubjectalternatenames
}
func flattenMapPolicyAbstractConfigProfileRelationship(p models.PolicyAbstractConfigProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractconfigprofilerelationships []map[string]interface{}
	var ret models.PolicyAbstractConfigProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	policyabstractconfigprofilerelationship := make(map[string]interface{})
	policyabstractconfigprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyabstractconfigprofilerelationship["class_id"] = item.GetClassId()
	policyabstractconfigprofilerelationship["moid"] = item.GetMoid()
	policyabstractconfigprofilerelationship["object_type"] = item.GetObjectType()
	policyabstractconfigprofilerelationship["selector"] = item.GetSelector()

	policyabstractconfigprofilerelationships = append(policyabstractconfigprofilerelationships, policyabstractconfigprofilerelationship)
	return policyabstractconfigprofilerelationships
}
func flattenMapPolicyAbstractProfileRelationship(p models.PolicyAbstractProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractprofilerelationships []map[string]interface{}
	var ret models.PolicyAbstractProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	policyabstractprofilerelationship := make(map[string]interface{})
	policyabstractprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyabstractprofilerelationship["class_id"] = item.GetClassId()
	policyabstractprofilerelationship["moid"] = item.GetMoid()
	policyabstractprofilerelationship["object_type"] = item.GetObjectType()
	policyabstractprofilerelationship["selector"] = item.GetSelector()

	policyabstractprofilerelationships = append(policyabstractprofilerelationships, policyabstractprofilerelationship)
	return policyabstractprofilerelationships
}
func flattenMapPolicyConfigChange(p models.PolicyConfigChange, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigchanges []map[string]interface{}
	var ret models.PolicyConfigChange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigchange := make(map[string]interface{})
	policyconfigchange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigchange["changes"] = item.GetChanges()
	policyconfigchange["class_id"] = item.GetClassId()
	policyconfigchange["disruptions"] = item.GetDisruptions()
	policyconfigchange["object_type"] = item.GetObjectType()

	policyconfigchanges = append(policyconfigchanges, policyconfigchange)
	return policyconfigchanges
}
func flattenMapPolicyConfigChangeContext(p models.PolicyConfigChangeContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigchangecontexts []map[string]interface{}
	var ret models.PolicyConfigChangeContext
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigchangecontext := make(map[string]interface{})
	policyconfigchangecontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigchangecontext["class_id"] = item.GetClassId()
	policyconfigchangecontext["config_change_error"] = item.GetConfigChangeError()
	policyconfigchangecontext["config_change_state"] = item.GetConfigChangeState()
	policyconfigchangecontext["initial_config_context"] = (func(p models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {
		var policyconfigcontexts []map[string]interface{}
		var ret models.PolicyConfigContext
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		policyconfigcontext := make(map[string]interface{})
		policyconfigcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		policyconfigcontext["class_id"] = item.GetClassId()
		policyconfigcontext["config_state"] = item.GetConfigState()
		policyconfigcontext["control_action"] = item.GetControlAction()
		policyconfigcontext["error_state"] = item.GetErrorState()
		policyconfigcontext["object_type"] = item.GetObjectType()
		policyconfigcontext["oper_state"] = item.GetOperState()

		policyconfigcontexts = append(policyconfigcontexts, policyconfigcontext)
		return policyconfigcontexts
	})(item.GetInitialConfigContext(), d)
	policyconfigchangecontext["object_type"] = item.GetObjectType()

	policyconfigchangecontexts = append(policyconfigchangecontexts, policyconfigchangecontext)
	return policyconfigchangecontexts
}
func flattenMapPolicyConfigContext(p models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigcontexts []map[string]interface{}
	var ret models.PolicyConfigContext
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigcontext := make(map[string]interface{})
	policyconfigcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigcontext["class_id"] = item.GetClassId()
	policyconfigcontext["config_state"] = item.GetConfigState()
	policyconfigcontext["control_action"] = item.GetControlAction()
	policyconfigcontext["error_state"] = item.GetErrorState()
	policyconfigcontext["object_type"] = item.GetObjectType()
	policyconfigcontext["oper_state"] = item.GetOperState()

	policyconfigcontexts = append(policyconfigcontexts, policyconfigcontext)
	return policyconfigcontexts
}
func flattenMapPolicyConfigResultContext(p models.PolicyConfigResultContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigresultcontexts []map[string]interface{}
	var ret models.PolicyConfigResultContext
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigresultcontext := make(map[string]interface{})
	policyconfigresultcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigresultcontext["class_id"] = item.GetClassId()
	policyconfigresultcontext["entity_data"] = flattenAdditionalProperties(item.EntityData)
	policyconfigresultcontext["entity_moid"] = item.GetEntityMoid()
	policyconfigresultcontext["entity_name"] = item.GetEntityName()
	policyconfigresultcontext["entity_type"] = item.GetEntityType()
	policyconfigresultcontext["object_type"] = item.GetObjectType()

	policyconfigresultcontexts = append(policyconfigresultcontexts, policyconfigresultcontext)
	return policyconfigresultcontexts
}
func flattenMapPortGroupRelationship(p models.PortGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprelationships []map[string]interface{}
	var ret models.PortGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	portgrouprelationship := make(map[string]interface{})
	portgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	portgrouprelationship["class_id"] = item.GetClassId()
	portgrouprelationship["moid"] = item.GetMoid()
	portgrouprelationship["object_type"] = item.GetObjectType()
	portgrouprelationship["selector"] = item.GetSelector()

	portgrouprelationships = append(portgrouprelationships, portgrouprelationship)
	return portgrouprelationships
}
func flattenMapPortInterfaceBaseRelationship(p models.PortInterfaceBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portinterfacebaserelationships []map[string]interface{}
	var ret models.PortInterfaceBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	portinterfacebaserelationship := make(map[string]interface{})
	portinterfacebaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	portinterfacebaserelationship["class_id"] = item.GetClassId()
	portinterfacebaserelationship["moid"] = item.GetMoid()
	portinterfacebaserelationship["object_type"] = item.GetObjectType()
	portinterfacebaserelationship["selector"] = item.GetSelector()

	portinterfacebaserelationships = append(portinterfacebaserelationships, portinterfacebaserelationship)
	return portinterfacebaserelationships
}
func flattenMapPortSubGroupRelationship(p models.PortSubGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprelationships []map[string]interface{}
	var ret models.PortSubGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	portsubgrouprelationship := make(map[string]interface{})
	portsubgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	portsubgrouprelationship["class_id"] = item.GetClassId()
	portsubgrouprelationship["moid"] = item.GetMoid()
	portsubgrouprelationship["object_type"] = item.GetObjectType()
	portsubgrouprelationship["selector"] = item.GetSelector()

	portsubgrouprelationships = append(portsubgrouprelationships, portsubgrouprelationship)
	return portsubgrouprelationships
}
func flattenMapRecommendationCapacityRunwayRelationship(p models.RecommendationCapacityRunwayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recommendationcapacityrunwayrelationships []map[string]interface{}
	var ret models.RecommendationCapacityRunwayRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recommendationcapacityrunwayrelationship := make(map[string]interface{})
	recommendationcapacityrunwayrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recommendationcapacityrunwayrelationship["class_id"] = item.GetClassId()
	recommendationcapacityrunwayrelationship["moid"] = item.GetMoid()
	recommendationcapacityrunwayrelationship["object_type"] = item.GetObjectType()
	recommendationcapacityrunwayrelationship["selector"] = item.GetSelector()

	recommendationcapacityrunwayrelationships = append(recommendationcapacityrunwayrelationships, recommendationcapacityrunwayrelationship)
	return recommendationcapacityrunwayrelationships
}
func flattenMapRecoveryAbstractBackupInfoRelationship(p models.RecoveryAbstractBackupInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryabstractbackupinforelationships []map[string]interface{}
	var ret models.RecoveryAbstractBackupInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoveryabstractbackupinforelationship := make(map[string]interface{})
	recoveryabstractbackupinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryabstractbackupinforelationship["class_id"] = item.GetClassId()
	recoveryabstractbackupinforelationship["moid"] = item.GetMoid()
	recoveryabstractbackupinforelationship["object_type"] = item.GetObjectType()
	recoveryabstractbackupinforelationship["selector"] = item.GetSelector()

	recoveryabstractbackupinforelationships = append(recoveryabstractbackupinforelationships, recoveryabstractbackupinforelationship)
	return recoveryabstractbackupinforelationships
}
func flattenMapRecoveryBackupConfigPolicyRelationship(p models.RecoveryBackupConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupconfigpolicyrelationships []map[string]interface{}
	var ret models.RecoveryBackupConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoverybackupconfigpolicyrelationship := make(map[string]interface{})
	recoverybackupconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoverybackupconfigpolicyrelationship["class_id"] = item.GetClassId()
	recoverybackupconfigpolicyrelationship["moid"] = item.GetMoid()
	recoverybackupconfigpolicyrelationship["object_type"] = item.GetObjectType()
	recoverybackupconfigpolicyrelationship["selector"] = item.GetSelector()

	recoverybackupconfigpolicyrelationships = append(recoverybackupconfigpolicyrelationships, recoverybackupconfigpolicyrelationship)
	return recoverybackupconfigpolicyrelationships
}
func flattenMapRecoveryBackupProfileRelationship(p models.RecoveryBackupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerelationships []map[string]interface{}
	var ret models.RecoveryBackupProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoverybackupprofilerelationship := make(map[string]interface{})
	recoverybackupprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoverybackupprofilerelationship["class_id"] = item.GetClassId()
	recoverybackupprofilerelationship["moid"] = item.GetMoid()
	recoverybackupprofilerelationship["object_type"] = item.GetObjectType()
	recoverybackupprofilerelationship["selector"] = item.GetSelector()

	recoverybackupprofilerelationships = append(recoverybackupprofilerelationships, recoverybackupprofilerelationship)
	return recoverybackupprofilerelationships
}
func flattenMapRecoveryBackupSchedule(p models.RecoveryBackupSchedule, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupschedules []map[string]interface{}
	var ret models.RecoveryBackupSchedule
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	recoverybackupschedule := make(map[string]interface{})
	recoverybackupschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoverybackupschedule["class_id"] = item.GetClassId()
	recoverybackupschedule["execution_time"] = item.GetExecutionTime()
	recoverybackupschedule["frequency_unit"] = item.GetFrequencyUnit()
	recoverybackupschedule["hours"] = item.GetHours()
	recoverybackupschedule["object_type"] = item.GetObjectType()

	recoverybackupschedules = append(recoverybackupschedules, recoverybackupschedule)
	return recoverybackupschedules
}
func flattenMapRecoveryConfigParams(p models.RecoveryConfigParams, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigparamss []map[string]interface{}
	var ret models.RecoveryConfigParams
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	recoveryconfigparams := make(map[string]interface{})
	recoveryconfigparams["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryconfigparams["class_id"] = item.GetClassId()
	recoveryconfigparams["object_type"] = item.GetObjectType()

	recoveryconfigparamss = append(recoveryconfigparamss, recoveryconfigparams)
	return recoveryconfigparamss
}
func flattenMapRecoveryConfigResultRelationship(p models.RecoveryConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultrelationships []map[string]interface{}
	var ret models.RecoveryConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoveryconfigresultrelationship := make(map[string]interface{})
	recoveryconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryconfigresultrelationship["class_id"] = item.GetClassId()
	recoveryconfigresultrelationship["moid"] = item.GetMoid()
	recoveryconfigresultrelationship["object_type"] = item.GetObjectType()
	recoveryconfigresultrelationship["selector"] = item.GetSelector()

	recoveryconfigresultrelationships = append(recoveryconfigresultrelationships, recoveryconfigresultrelationship)
	return recoveryconfigresultrelationships
}
func flattenMapRecoveryScheduleConfigPolicyRelationship(p models.RecoveryScheduleConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryscheduleconfigpolicyrelationships []map[string]interface{}
	var ret models.RecoveryScheduleConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoveryscheduleconfigpolicyrelationship := make(map[string]interface{})
	recoveryscheduleconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryscheduleconfigpolicyrelationship["class_id"] = item.GetClassId()
	recoveryscheduleconfigpolicyrelationship["moid"] = item.GetMoid()
	recoveryscheduleconfigpolicyrelationship["object_type"] = item.GetObjectType()
	recoveryscheduleconfigpolicyrelationship["selector"] = item.GetSelector()

	recoveryscheduleconfigpolicyrelationships = append(recoveryscheduleconfigpolicyrelationships, recoveryscheduleconfigpolicyrelationship)
	return recoveryscheduleconfigpolicyrelationships
}
func flattenMapResourceGroupRelationship(p models.ResourceGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprelationships []map[string]interface{}
	var ret models.ResourceGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	resourcegrouprelationship := make(map[string]interface{})
	resourcegrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	resourcegrouprelationship["class_id"] = item.GetClassId()
	resourcegrouprelationship["moid"] = item.GetMoid()
	resourcegrouprelationship["object_type"] = item.GetObjectType()
	resourcegrouprelationship["selector"] = item.GetSelector()

	resourcegrouprelationships = append(resourcegrouprelationships, resourcegrouprelationship)
	return resourcegrouprelationships
}
func flattenMapResourceMembershipHolderRelationship(p models.ResourceMembershipHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcemembershipholderrelationships []map[string]interface{}
	var ret models.ResourceMembershipHolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	resourcemembershipholderrelationship := make(map[string]interface{})
	resourcemembershipholderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	resourcemembershipholderrelationship["class_id"] = item.GetClassId()
	resourcemembershipholderrelationship["moid"] = item.GetMoid()
	resourcemembershipholderrelationship["object_type"] = item.GetObjectType()
	resourcemembershipholderrelationship["selector"] = item.GetSelector()

	resourcemembershipholderrelationships = append(resourcemembershipholderrelationships, resourcemembershipholderrelationship)
	return resourcemembershipholderrelationships
}
func flattenMapSdwanProfileRelationship(p models.SdwanProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerelationships []map[string]interface{}
	var ret models.SdwanProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sdwanprofilerelationship := make(map[string]interface{})
	sdwanprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sdwanprofilerelationship["class_id"] = item.GetClassId()
	sdwanprofilerelationship["moid"] = item.GetMoid()
	sdwanprofilerelationship["object_type"] = item.GetObjectType()
	sdwanprofilerelationship["selector"] = item.GetSelector()

	sdwanprofilerelationships = append(sdwanprofilerelationships, sdwanprofilerelationship)
	return sdwanprofilerelationships
}
func flattenMapSdwanRouterPolicyRelationship(p models.SdwanRouterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouterpolicyrelationships []map[string]interface{}
	var ret models.SdwanRouterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sdwanrouterpolicyrelationship := make(map[string]interface{})
	sdwanrouterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sdwanrouterpolicyrelationship["class_id"] = item.GetClassId()
	sdwanrouterpolicyrelationship["moid"] = item.GetMoid()
	sdwanrouterpolicyrelationship["object_type"] = item.GetObjectType()
	sdwanrouterpolicyrelationship["selector"] = item.GetSelector()

	sdwanrouterpolicyrelationships = append(sdwanrouterpolicyrelationships, sdwanrouterpolicyrelationship)
	return sdwanrouterpolicyrelationships
}
func flattenMapSdwanVmanageAccountPolicyRelationship(p models.SdwanVmanageAccountPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanvmanageaccountpolicyrelationships []map[string]interface{}
	var ret models.SdwanVmanageAccountPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sdwanvmanageaccountpolicyrelationship := make(map[string]interface{})
	sdwanvmanageaccountpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sdwanvmanageaccountpolicyrelationship["class_id"] = item.GetClassId()
	sdwanvmanageaccountpolicyrelationship["moid"] = item.GetMoid()
	sdwanvmanageaccountpolicyrelationship["object_type"] = item.GetObjectType()
	sdwanvmanageaccountpolicyrelationship["selector"] = item.GetSelector()

	sdwanvmanageaccountpolicyrelationships = append(sdwanvmanageaccountpolicyrelationships, sdwanvmanageaccountpolicyrelationship)
	return sdwanvmanageaccountpolicyrelationships
}
func flattenMapServerConfigResultRelationship(p models.ServerConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultrelationships []map[string]interface{}
	var ret models.ServerConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	serverconfigresultrelationship := make(map[string]interface{})
	serverconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	serverconfigresultrelationship["class_id"] = item.GetClassId()
	serverconfigresultrelationship["moid"] = item.GetMoid()
	serverconfigresultrelationship["object_type"] = item.GetObjectType()
	serverconfigresultrelationship["selector"] = item.GetSelector()

	serverconfigresultrelationships = append(serverconfigresultrelationships, serverconfigresultrelationship)
	return serverconfigresultrelationships
}
func flattenMapServerProfileRelationship(p models.ServerProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverprofilerelationships []map[string]interface{}
	var ret models.ServerProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	serverprofilerelationship := make(map[string]interface{})
	serverprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	serverprofilerelationship["class_id"] = item.GetClassId()
	serverprofilerelationship["moid"] = item.GetMoid()
	serverprofilerelationship["object_type"] = item.GetObjectType()
	serverprofilerelationship["selector"] = item.GetSelector()

	serverprofilerelationships = append(serverprofilerelationships, serverprofilerelationship)
	return serverprofilerelationships
}
func flattenMapSessionAbstractSessionRelationship(p models.SessionAbstractSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sessionabstractsessionrelationships []map[string]interface{}
	var ret models.SessionAbstractSessionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sessionabstractsessionrelationship := make(map[string]interface{})
	sessionabstractsessionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sessionabstractsessionrelationship["class_id"] = item.GetClassId()
	sessionabstractsessionrelationship["moid"] = item.GetMoid()
	sessionabstractsessionrelationship["object_type"] = item.GetObjectType()
	sessionabstractsessionrelationship["selector"] = item.GetSelector()

	sessionabstractsessionrelationships = append(sessionabstractsessionrelationships, sessionabstractsessionrelationship)
	return sessionabstractsessionrelationships
}
func flattenMapSoftwareHyperflexDistributableRelationship(p models.SoftwareHyperflexDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarehyperflexdistributablerelationships []map[string]interface{}
	var ret models.SoftwareHyperflexDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarehyperflexdistributablerelationship := make(map[string]interface{})
	softwarehyperflexdistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarehyperflexdistributablerelationship["class_id"] = item.GetClassId()
	softwarehyperflexdistributablerelationship["moid"] = item.GetMoid()
	softwarehyperflexdistributablerelationship["object_type"] = item.GetObjectType()
	softwarehyperflexdistributablerelationship["selector"] = item.GetSelector()

	softwarehyperflexdistributablerelationships = append(softwarehyperflexdistributablerelationships, softwarehyperflexdistributablerelationship)
	return softwarehyperflexdistributablerelationships
}
func flattenMapSoftwareSolutionDistributableRelationship(p models.SoftwareSolutionDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwaresolutiondistributablerelationships []map[string]interface{}
	var ret models.SoftwareSolutionDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwaresolutiondistributablerelationship := make(map[string]interface{})
	softwaresolutiondistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwaresolutiondistributablerelationship["class_id"] = item.GetClassId()
	softwaresolutiondistributablerelationship["moid"] = item.GetMoid()
	softwaresolutiondistributablerelationship["object_type"] = item.GetObjectType()
	softwaresolutiondistributablerelationship["selector"] = item.GetSelector()

	softwaresolutiondistributablerelationships = append(softwaresolutiondistributablerelationships, softwaresolutiondistributablerelationship)
	return softwaresolutiondistributablerelationships
}
func flattenMapSoftwarerepositoryCatalogRelationship(p models.SoftwarerepositoryCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositorycatalogrelationships []map[string]interface{}
	var ret models.SoftwarerepositoryCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositorycatalogrelationship := make(map[string]interface{})
	softwarerepositorycatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositorycatalogrelationship["class_id"] = item.GetClassId()
	softwarerepositorycatalogrelationship["moid"] = item.GetMoid()
	softwarerepositorycatalogrelationship["object_type"] = item.GetObjectType()
	softwarerepositorycatalogrelationship["selector"] = item.GetSelector()

	softwarerepositorycatalogrelationships = append(softwarerepositorycatalogrelationships, softwarerepositorycatalogrelationship)
	return softwarerepositorycatalogrelationships
}
func flattenMapSoftwarerepositoryFileRelationship(p models.SoftwarerepositoryFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryfilerelationships []map[string]interface{}
	var ret models.SoftwarerepositoryFileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositoryfilerelationship := make(map[string]interface{})
	softwarerepositoryfilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryfilerelationship["class_id"] = item.GetClassId()
	softwarerepositoryfilerelationship["moid"] = item.GetMoid()
	softwarerepositoryfilerelationship["object_type"] = item.GetObjectType()
	softwarerepositoryfilerelationship["selector"] = item.GetSelector()

	softwarerepositoryfilerelationships = append(softwarerepositoryfilerelationships, softwarerepositoryfilerelationship)
	return softwarerepositoryfilerelationships
}
func flattenMapSoftwarerepositoryFileServer(p models.SoftwarerepositoryFileServer, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryfileservers []map[string]interface{}
	var ret models.SoftwarerepositoryFileServer
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	softwarerepositoryfileserver := make(map[string]interface{})
	softwarerepositoryfileserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryfileserver["class_id"] = item.GetClassId()
	softwarerepositoryfileserver["object_type"] = item.GetObjectType()

	softwarerepositoryfileservers = append(softwarerepositoryfileservers, softwarerepositoryfileserver)
	return softwarerepositoryfileservers
}
func flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p models.SoftwarerepositoryOperatingSystemFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryoperatingsystemfilerelationships []map[string]interface{}
	var ret models.SoftwarerepositoryOperatingSystemFileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositoryoperatingsystemfilerelationship := make(map[string]interface{})
	softwarerepositoryoperatingsystemfilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryoperatingsystemfilerelationship["class_id"] = item.GetClassId()
	softwarerepositoryoperatingsystemfilerelationship["moid"] = item.GetMoid()
	softwarerepositoryoperatingsystemfilerelationship["object_type"] = item.GetObjectType()
	softwarerepositoryoperatingsystemfilerelationship["selector"] = item.GetSelector()

	softwarerepositoryoperatingsystemfilerelationships = append(softwarerepositoryoperatingsystemfilerelationships, softwarerepositoryoperatingsystemfilerelationship)
	return softwarerepositoryoperatingsystemfilerelationships
}
func flattenMapSoftwarerepositoryReleaseRelationship(p models.SoftwarerepositoryReleaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryreleaserelationships []map[string]interface{}
	var ret models.SoftwarerepositoryReleaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositoryreleaserelationship := make(map[string]interface{})
	softwarerepositoryreleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryreleaserelationship["class_id"] = item.GetClassId()
	softwarerepositoryreleaserelationship["moid"] = item.GetMoid()
	softwarerepositoryreleaserelationship["object_type"] = item.GetObjectType()
	softwarerepositoryreleaserelationship["selector"] = item.GetSelector()

	softwarerepositoryreleaserelationships = append(softwarerepositoryreleaserelationships, softwarerepositoryreleaserelationship)
	return softwarerepositoryreleaserelationships
}
func flattenMapStorageBaseCapacity(p models.StorageBaseCapacity, d *schema.ResourceData) []map[string]interface{} {
	var storagebasecapacitys []map[string]interface{}
	var ret models.StorageBaseCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	storagebasecapacity := make(map[string]interface{})
	storagebasecapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagebasecapacity["available"] = item.GetAvailable()
	storagebasecapacity["capacity_utilization"] = item.GetCapacityUtilization()
	storagebasecapacity["class_id"] = item.GetClassId()
	storagebasecapacity["free"] = item.GetFree()
	storagebasecapacity["object_type"] = item.GetObjectType()
	storagebasecapacity["total"] = item.GetTotal()
	storagebasecapacity["used"] = item.GetUsed()

	storagebasecapacitys = append(storagebasecapacitys, storagebasecapacity)
	return storagebasecapacitys
}
func flattenMapStorageControllerRelationship(p models.StorageControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrelationships []map[string]interface{}
	var ret models.StorageControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagecontrollerrelationship := make(map[string]interface{})
	storagecontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagecontrollerrelationship["class_id"] = item.GetClassId()
	storagecontrollerrelationship["moid"] = item.GetMoid()
	storagecontrollerrelationship["object_type"] = item.GetObjectType()
	storagecontrollerrelationship["selector"] = item.GetSelector()

	storagecontrollerrelationships = append(storagecontrollerrelationships, storagecontrollerrelationship)
	return storagecontrollerrelationships
}
func flattenMapStorageDiskGroupRelationship(p models.StorageDiskGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouprelationships []map[string]interface{}
	var ret models.StorageDiskGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagediskgrouprelationship := make(map[string]interface{})
	storagediskgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagediskgrouprelationship["class_id"] = item.GetClassId()
	storagediskgrouprelationship["moid"] = item.GetMoid()
	storagediskgrouprelationship["object_type"] = item.GetObjectType()
	storagediskgrouprelationship["selector"] = item.GetSelector()

	storagediskgrouprelationships = append(storagediskgrouprelationships, storagediskgrouprelationship)
	return storagediskgrouprelationships
}
func flattenMapStorageDiskSlotRelationship(p models.StorageDiskSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskslotrelationships []map[string]interface{}
	var ret models.StorageDiskSlotRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagediskslotrelationship := make(map[string]interface{})
	storagediskslotrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagediskslotrelationship["class_id"] = item.GetClassId()
	storagediskslotrelationship["moid"] = item.GetMoid()
	storagediskslotrelationship["object_type"] = item.GetObjectType()
	storagediskslotrelationship["selector"] = item.GetSelector()

	storagediskslotrelationships = append(storagediskslotrelationships, storagediskslotrelationship)
	return storagediskslotrelationships
}
func flattenMapStorageEnclosureRelationship(p models.StorageEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerelationships []map[string]interface{}
	var ret models.StorageEnclosureRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storageenclosurerelationship := make(map[string]interface{})
	storageenclosurerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storageenclosurerelationship["class_id"] = item.GetClassId()
	storageenclosurerelationship["moid"] = item.GetMoid()
	storageenclosurerelationship["object_type"] = item.GetObjectType()
	storageenclosurerelationship["selector"] = item.GetSelector()

	storageenclosurerelationships = append(storageenclosurerelationships, storageenclosurerelationship)
	return storageenclosurerelationships
}
func flattenMapStorageFlexFlashControllerRelationship(p models.StorageFlexFlashControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrelationships []map[string]interface{}
	var ret models.StorageFlexFlashControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storageflexflashcontrollerrelationship := make(map[string]interface{})
	storageflexflashcontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storageflexflashcontrollerrelationship["class_id"] = item.GetClassId()
	storageflexflashcontrollerrelationship["moid"] = item.GetMoid()
	storageflexflashcontrollerrelationship["object_type"] = item.GetObjectType()
	storageflexflashcontrollerrelationship["selector"] = item.GetSelector()

	storageflexflashcontrollerrelationships = append(storageflexflashcontrollerrelationships, storageflexflashcontrollerrelationship)
	return storageflexflashcontrollerrelationships
}
func flattenMapStorageFlexUtilControllerRelationship(p models.StorageFlexUtilControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrelationships []map[string]interface{}
	var ret models.StorageFlexUtilControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storageflexutilcontrollerrelationship := make(map[string]interface{})
	storageflexutilcontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storageflexutilcontrollerrelationship["class_id"] = item.GetClassId()
	storageflexutilcontrollerrelationship["moid"] = item.GetMoid()
	storageflexutilcontrollerrelationship["object_type"] = item.GetObjectType()
	storageflexutilcontrollerrelationship["selector"] = item.GetSelector()

	storageflexutilcontrollerrelationships = append(storageflexutilcontrollerrelationships, storageflexutilcontrollerrelationship)
	return storageflexutilcontrollerrelationships
}
func flattenMapStorageHitachiArrayRelationship(p models.StorageHitachiArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehitachiarrayrelationships []map[string]interface{}
	var ret models.StorageHitachiArrayRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagehitachiarrayrelationship := make(map[string]interface{})
	storagehitachiarrayrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagehitachiarrayrelationship["class_id"] = item.GetClassId()
	storagehitachiarrayrelationship["moid"] = item.GetMoid()
	storagehitachiarrayrelationship["object_type"] = item.GetObjectType()
	storagehitachiarrayrelationship["selector"] = item.GetSelector()

	storagehitachiarrayrelationships = append(storagehitachiarrayrelationships, storagehitachiarrayrelationship)
	return storagehitachiarrayrelationships
}
func flattenMapStorageHitachiHostRelationship(p models.StorageHitachiHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehitachihostrelationships []map[string]interface{}
	var ret models.StorageHitachiHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagehitachihostrelationship := make(map[string]interface{})
	storagehitachihostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagehitachihostrelationship["class_id"] = item.GetClassId()
	storagehitachihostrelationship["moid"] = item.GetMoid()
	storagehitachihostrelationship["object_type"] = item.GetObjectType()
	storagehitachihostrelationship["selector"] = item.GetSelector()

	storagehitachihostrelationships = append(storagehitachihostrelationships, storagehitachihostrelationship)
	return storagehitachihostrelationships
}
func flattenMapStorageHitachiParityGroupRelationship(p models.StorageHitachiParityGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehitachiparitygrouprelationships []map[string]interface{}
	var ret models.StorageHitachiParityGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagehitachiparitygrouprelationship := make(map[string]interface{})
	storagehitachiparitygrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagehitachiparitygrouprelationship["class_id"] = item.GetClassId()
	storagehitachiparitygrouprelationship["moid"] = item.GetMoid()
	storagehitachiparitygrouprelationship["object_type"] = item.GetObjectType()
	storagehitachiparitygrouprelationship["selector"] = item.GetSelector()

	storagehitachiparitygrouprelationships = append(storagehitachiparitygrouprelationships, storagehitachiparitygrouprelationship)
	return storagehitachiparitygrouprelationships
}
func flattenMapStorageHitachiPoolRelationship(p models.StorageHitachiPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehitachipoolrelationships []map[string]interface{}
	var ret models.StorageHitachiPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagehitachipoolrelationship := make(map[string]interface{})
	storagehitachipoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagehitachipoolrelationship["class_id"] = item.GetClassId()
	storagehitachipoolrelationship["moid"] = item.GetMoid()
	storagehitachipoolrelationship["object_type"] = item.GetObjectType()
	storagehitachipoolrelationship["selector"] = item.GetSelector()

	storagehitachipoolrelationships = append(storagehitachipoolrelationships, storagehitachipoolrelationship)
	return storagehitachipoolrelationships
}
func flattenMapStorageHitachiVolumeRelationship(p models.StorageHitachiVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehitachivolumerelationships []map[string]interface{}
	var ret models.StorageHitachiVolumeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagehitachivolumerelationship := make(map[string]interface{})
	storagehitachivolumerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagehitachivolumerelationship["class_id"] = item.GetClassId()
	storagehitachivolumerelationship["moid"] = item.GetMoid()
	storagehitachivolumerelationship["object_type"] = item.GetObjectType()
	storagehitachivolumerelationship["selector"] = item.GetSelector()

	storagehitachivolumerelationships = append(storagehitachivolumerelationships, storagehitachivolumerelationship)
	return storagehitachivolumerelationships
}
func flattenMapStorageHyperFlexStorageContainerRelationship(p models.StorageHyperFlexStorageContainerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehyperflexstoragecontainerrelationships []map[string]interface{}
	var ret models.StorageHyperFlexStorageContainerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagehyperflexstoragecontainerrelationship := make(map[string]interface{})
	storagehyperflexstoragecontainerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagehyperflexstoragecontainerrelationship["class_id"] = item.GetClassId()
	storagehyperflexstoragecontainerrelationship["moid"] = item.GetMoid()
	storagehyperflexstoragecontainerrelationship["object_type"] = item.GetObjectType()
	storagehyperflexstoragecontainerrelationship["selector"] = item.GetSelector()

	storagehyperflexstoragecontainerrelationships = append(storagehyperflexstoragecontainerrelationships, storagehyperflexstoragecontainerrelationship)
	return storagehyperflexstoragecontainerrelationships
}
func flattenMapStorageNetAppClusterRelationship(p models.StorageNetAppClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappclusterrelationships []map[string]interface{}
	var ret models.StorageNetAppClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetappclusterrelationship := make(map[string]interface{})
	storagenetappclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetappclusterrelationship["class_id"] = item.GetClassId()
	storagenetappclusterrelationship["moid"] = item.GetMoid()
	storagenetappclusterrelationship["object_type"] = item.GetObjectType()
	storagenetappclusterrelationship["selector"] = item.GetSelector()

	storagenetappclusterrelationships = append(storagenetappclusterrelationships, storagenetappclusterrelationship)
	return storagenetappclusterrelationships
}
func flattenMapStorageNetAppEthernetPortRelationship(p models.StorageNetAppEthernetPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappethernetportrelationships []map[string]interface{}
	var ret models.StorageNetAppEthernetPortRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetappethernetportrelationship := make(map[string]interface{})
	storagenetappethernetportrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetappethernetportrelationship["class_id"] = item.GetClassId()
	storagenetappethernetportrelationship["moid"] = item.GetMoid()
	storagenetappethernetportrelationship["object_type"] = item.GetObjectType()
	storagenetappethernetportrelationship["selector"] = item.GetSelector()

	storagenetappethernetportrelationships = append(storagenetappethernetportrelationships, storagenetappethernetportrelationship)
	return storagenetappethernetportrelationships
}
func flattenMapStorageNetAppFcPortRelationship(p models.StorageNetAppFcPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappfcportrelationships []map[string]interface{}
	var ret models.StorageNetAppFcPortRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetappfcportrelationship := make(map[string]interface{})
	storagenetappfcportrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetappfcportrelationship["class_id"] = item.GetClassId()
	storagenetappfcportrelationship["moid"] = item.GetMoid()
	storagenetappfcportrelationship["object_type"] = item.GetObjectType()
	storagenetappfcportrelationship["selector"] = item.GetSelector()

	storagenetappfcportrelationships = append(storagenetappfcportrelationships, storagenetappfcportrelationship)
	return storagenetappfcportrelationships
}
func flattenMapStorageNetAppLunRelationship(p models.StorageNetAppLunRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetapplunrelationships []map[string]interface{}
	var ret models.StorageNetAppLunRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetapplunrelationship := make(map[string]interface{})
	storagenetapplunrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetapplunrelationship["class_id"] = item.GetClassId()
	storagenetapplunrelationship["moid"] = item.GetMoid()
	storagenetapplunrelationship["object_type"] = item.GetObjectType()
	storagenetapplunrelationship["selector"] = item.GetSelector()

	storagenetapplunrelationships = append(storagenetapplunrelationships, storagenetapplunrelationship)
	return storagenetapplunrelationships
}
func flattenMapStorageNetAppNodeRelationship(p models.StorageNetAppNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappnoderelationships []map[string]interface{}
	var ret models.StorageNetAppNodeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetappnoderelationship := make(map[string]interface{})
	storagenetappnoderelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetappnoderelationship["class_id"] = item.GetClassId()
	storagenetappnoderelationship["moid"] = item.GetMoid()
	storagenetappnoderelationship["object_type"] = item.GetObjectType()
	storagenetappnoderelationship["selector"] = item.GetSelector()

	storagenetappnoderelationships = append(storagenetappnoderelationships, storagenetappnoderelationship)
	return storagenetappnoderelationships
}
func flattenMapStorageNetAppStorageVmRelationship(p models.StorageNetAppStorageVmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappstoragevmrelationships []map[string]interface{}
	var ret models.StorageNetAppStorageVmRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetappstoragevmrelationship := make(map[string]interface{})
	storagenetappstoragevmrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetappstoragevmrelationship["class_id"] = item.GetClassId()
	storagenetappstoragevmrelationship["moid"] = item.GetMoid()
	storagenetappstoragevmrelationship["object_type"] = item.GetObjectType()
	storagenetappstoragevmrelationship["selector"] = item.GetSelector()

	storagenetappstoragevmrelationships = append(storagenetappstoragevmrelationships, storagenetappstoragevmrelationship)
	return storagenetappstoragevmrelationships
}
func flattenMapStorageNetAppVolumeRelationship(p models.StorageNetAppVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagenetappvolumerelationships []map[string]interface{}
	var ret models.StorageNetAppVolumeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagenetappvolumerelationship := make(map[string]interface{})
	storagenetappvolumerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagenetappvolumerelationship["class_id"] = item.GetClassId()
	storagenetappvolumerelationship["moid"] = item.GetMoid()
	storagenetappvolumerelationship["object_type"] = item.GetObjectType()
	storagenetappvolumerelationship["selector"] = item.GetSelector()

	storagenetappvolumerelationships = append(storagenetappvolumerelationships, storagenetappvolumerelationship)
	return storagenetappvolumerelationships
}
func flattenMapStoragePhysicalDiskRelationship(p models.StoragePhysicalDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrelationships []map[string]interface{}
	var ret models.StoragePhysicalDiskRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagephysicaldiskrelationship := make(map[string]interface{})
	storagephysicaldiskrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagephysicaldiskrelationship["class_id"] = item.GetClassId()
	storagephysicaldiskrelationship["moid"] = item.GetMoid()
	storagephysicaldiskrelationship["object_type"] = item.GetObjectType()
	storagephysicaldiskrelationship["selector"] = item.GetSelector()

	storagephysicaldiskrelationships = append(storagephysicaldiskrelationships, storagephysicaldiskrelationship)
	return storagephysicaldiskrelationships
}
func flattenMapStoragePureArrayRelationship(p models.StoragePureArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurearrayrelationships []map[string]interface{}
	var ret models.StoragePureArrayRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurearrayrelationship := make(map[string]interface{})
	storagepurearrayrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurearrayrelationship["class_id"] = item.GetClassId()
	storagepurearrayrelationship["moid"] = item.GetMoid()
	storagepurearrayrelationship["object_type"] = item.GetObjectType()
	storagepurearrayrelationship["selector"] = item.GetSelector()

	storagepurearrayrelationships = append(storagepurearrayrelationships, storagepurearrayrelationship)
	return storagepurearrayrelationships
}
func flattenMapStoragePureControllerRelationship(p models.StoragePureControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurecontrollerrelationships []map[string]interface{}
	var ret models.StoragePureControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurecontrollerrelationship := make(map[string]interface{})
	storagepurecontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurecontrollerrelationship["class_id"] = item.GetClassId()
	storagepurecontrollerrelationship["moid"] = item.GetMoid()
	storagepurecontrollerrelationship["object_type"] = item.GetObjectType()
	storagepurecontrollerrelationship["selector"] = item.GetSelector()

	storagepurecontrollerrelationships = append(storagepurecontrollerrelationships, storagepurecontrollerrelationship)
	return storagepurecontrollerrelationships
}
func flattenMapStoragePureHostRelationship(p models.StoragePureHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrelationships []map[string]interface{}
	var ret models.StoragePureHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurehostrelationship := make(map[string]interface{})
	storagepurehostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurehostrelationship["class_id"] = item.GetClassId()
	storagepurehostrelationship["moid"] = item.GetMoid()
	storagepurehostrelationship["object_type"] = item.GetObjectType()
	storagepurehostrelationship["selector"] = item.GetSelector()

	storagepurehostrelationships = append(storagepurehostrelationships, storagepurehostrelationship)
	return storagepurehostrelationships
}
func flattenMapStoragePureHostGroupRelationship(p models.StoragePureHostGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprelationships []map[string]interface{}
	var ret models.StoragePureHostGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurehostgrouprelationship := make(map[string]interface{})
	storagepurehostgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurehostgrouprelationship["class_id"] = item.GetClassId()
	storagepurehostgrouprelationship["moid"] = item.GetMoid()
	storagepurehostgrouprelationship["object_type"] = item.GetObjectType()
	storagepurehostgrouprelationship["selector"] = item.GetSelector()

	storagepurehostgrouprelationships = append(storagepurehostgrouprelationships, storagepurehostgrouprelationship)
	return storagepurehostgrouprelationships
}
func flattenMapStoragePureProtectionGroupRelationship(p models.StoragePureProtectionGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongrouprelationships []map[string]interface{}
	var ret models.StoragePureProtectionGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepureprotectiongrouprelationship := make(map[string]interface{})
	storagepureprotectiongrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepureprotectiongrouprelationship["class_id"] = item.GetClassId()
	storagepureprotectiongrouprelationship["moid"] = item.GetMoid()
	storagepureprotectiongrouprelationship["object_type"] = item.GetObjectType()
	storagepureprotectiongrouprelationship["selector"] = item.GetSelector()

	storagepureprotectiongrouprelationships = append(storagepureprotectiongrouprelationships, storagepureprotectiongrouprelationship)
	return storagepureprotectiongrouprelationships
}
func flattenMapStoragePureProtectionGroupSnapshotRelationship(p models.StoragePureProtectionGroupSnapshotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongroupsnapshotrelationships []map[string]interface{}
	var ret models.StoragePureProtectionGroupSnapshotRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepureprotectiongroupsnapshotrelationship := make(map[string]interface{})
	storagepureprotectiongroupsnapshotrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepureprotectiongroupsnapshotrelationship["class_id"] = item.GetClassId()
	storagepureprotectiongroupsnapshotrelationship["moid"] = item.GetMoid()
	storagepureprotectiongroupsnapshotrelationship["object_type"] = item.GetObjectType()
	storagepureprotectiongroupsnapshotrelationship["selector"] = item.GetSelector()

	storagepureprotectiongroupsnapshotrelationships = append(storagepureprotectiongroupsnapshotrelationships, storagepureprotectiongroupsnapshotrelationship)
	return storagepureprotectiongroupsnapshotrelationships
}
func flattenMapStoragePureVolumeRelationship(p models.StoragePureVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerelationships []map[string]interface{}
	var ret models.StoragePureVolumeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurevolumerelationship := make(map[string]interface{})
	storagepurevolumerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurevolumerelationship["class_id"] = item.GetClassId()
	storagepurevolumerelationship["moid"] = item.GetMoid()
	storagepurevolumerelationship["object_type"] = item.GetObjectType()
	storagepurevolumerelationship["selector"] = item.GetSelector()

	storagepurevolumerelationships = append(storagepurevolumerelationships, storagepurevolumerelationship)
	return storagepurevolumerelationships
}
func flattenMapStorageSasExpanderRelationship(p models.StorageSasExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrelationships []map[string]interface{}
	var ret models.StorageSasExpanderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagesasexpanderrelationship := make(map[string]interface{})
	storagesasexpanderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagesasexpanderrelationship["class_id"] = item.GetClassId()
	storagesasexpanderrelationship["moid"] = item.GetMoid()
	storagesasexpanderrelationship["object_type"] = item.GetObjectType()
	storagesasexpanderrelationship["selector"] = item.GetSelector()

	storagesasexpanderrelationships = append(storagesasexpanderrelationships, storagesasexpanderrelationship)
	return storagesasexpanderrelationships
}
func flattenMapStorageVirtualDriveRelationship(p models.StorageVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverelationships []map[string]interface{}
	var ret models.StorageVirtualDriveRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagevirtualdriverelationship := make(map[string]interface{})
	storagevirtualdriverelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagevirtualdriverelationship["class_id"] = item.GetClassId()
	storagevirtualdriverelationship["moid"] = item.GetMoid()
	storagevirtualdriverelationship["object_type"] = item.GetObjectType()
	storagevirtualdriverelationship["selector"] = item.GetSelector()

	storagevirtualdriverelationships = append(storagevirtualdriverelationships, storagevirtualdriverelationship)
	return storagevirtualdriverelationships
}
func flattenMapStorageVirtualDriveContainerRelationship(p models.StorageVirtualDriveContainerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdrivecontainerrelationships []map[string]interface{}
	var ret models.StorageVirtualDriveContainerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagevirtualdrivecontainerrelationship := make(map[string]interface{})
	storagevirtualdrivecontainerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagevirtualdrivecontainerrelationship["class_id"] = item.GetClassId()
	storagevirtualdrivecontainerrelationship["moid"] = item.GetMoid()
	storagevirtualdrivecontainerrelationship["object_type"] = item.GetObjectType()
	storagevirtualdrivecontainerrelationship["selector"] = item.GetSelector()

	storagevirtualdrivecontainerrelationships = append(storagevirtualdrivecontainerrelationships, storagevirtualdrivecontainerrelationship)
	return storagevirtualdrivecontainerrelationships
}
func flattenMapStorageVirtualDriveExtensionRelationship(p models.StorageVirtualDriveExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrelationships []map[string]interface{}
	var ret models.StorageVirtualDriveExtensionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagevirtualdriveextensionrelationship := make(map[string]interface{})
	storagevirtualdriveextensionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagevirtualdriveextensionrelationship["class_id"] = item.GetClassId()
	storagevirtualdriveextensionrelationship["moid"] = item.GetMoid()
	storagevirtualdriveextensionrelationship["object_type"] = item.GetObjectType()
	storagevirtualdriveextensionrelationship["selector"] = item.GetSelector()

	storagevirtualdriveextensionrelationships = append(storagevirtualdriveextensionrelationships, storagevirtualdriveextensionrelationship)
	return storagevirtualdriveextensionrelationships
}
func flattenMapTamBaseAdvisoryRelationship(p models.TamBaseAdvisoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var tambaseadvisoryrelationships []map[string]interface{}
	var ret models.TamBaseAdvisoryRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	tambaseadvisoryrelationship := make(map[string]interface{})
	tambaseadvisoryrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	tambaseadvisoryrelationship["class_id"] = item.GetClassId()
	tambaseadvisoryrelationship["moid"] = item.GetMoid()
	tambaseadvisoryrelationship["object_type"] = item.GetObjectType()
	tambaseadvisoryrelationship["selector"] = item.GetSelector()

	tambaseadvisoryrelationships = append(tambaseadvisoryrelationships, tambaseadvisoryrelationship)
	return tambaseadvisoryrelationships
}
func flattenMapTamBaseAdvisoryDetails(p models.TamBaseAdvisoryDetails, d *schema.ResourceData) []map[string]interface{} {
	var tambaseadvisorydetailss []map[string]interface{}
	var ret models.TamBaseAdvisoryDetails
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	tambaseadvisorydetails := make(map[string]interface{})
	tambaseadvisorydetails["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	tambaseadvisorydetails["class_id"] = item.GetClassId()
	tambaseadvisorydetails["description"] = item.GetDescription()
	tambaseadvisorydetails["object_type"] = item.GetObjectType()

	tambaseadvisorydetailss = append(tambaseadvisorydetailss, tambaseadvisorydetails)
	return tambaseadvisorydetailss
}
func flattenMapTamSeverity(p models.TamSeverity, d *schema.ResourceData) []map[string]interface{} {
	var tamseveritys []map[string]interface{}
	var ret models.TamSeverity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	tamseverity := make(map[string]interface{})
	tamseverity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	tamseverity["class_id"] = item.GetClassId()
	tamseverity["object_type"] = item.GetObjectType()

	tamseveritys = append(tamseveritys, tamseverity)
	return tamseveritys
}
func flattenMapTechsupportmanagementTechSupportBundleRelationship(p models.TechsupportmanagementTechSupportBundleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var techsupportmanagementtechsupportbundlerelationships []map[string]interface{}
	var ret models.TechsupportmanagementTechSupportBundleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	techsupportmanagementtechsupportbundlerelationship := make(map[string]interface{})
	techsupportmanagementtechsupportbundlerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	techsupportmanagementtechsupportbundlerelationship["class_id"] = item.GetClassId()
	techsupportmanagementtechsupportbundlerelationship["moid"] = item.GetMoid()
	techsupportmanagementtechsupportbundlerelationship["object_type"] = item.GetObjectType()
	techsupportmanagementtechsupportbundlerelationship["selector"] = item.GetSelector()

	techsupportmanagementtechsupportbundlerelationships = append(techsupportmanagementtechsupportbundlerelationships, techsupportmanagementtechsupportbundlerelationship)
	return techsupportmanagementtechsupportbundlerelationships
}
func flattenMapTechsupportmanagementTechSupportStatusRelationship(p models.TechsupportmanagementTechSupportStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var techsupportmanagementtechsupportstatusrelationships []map[string]interface{}
	var ret models.TechsupportmanagementTechSupportStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	techsupportmanagementtechsupportstatusrelationship := make(map[string]interface{})
	techsupportmanagementtechsupportstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	techsupportmanagementtechsupportstatusrelationship["class_id"] = item.GetClassId()
	techsupportmanagementtechsupportstatusrelationship["moid"] = item.GetMoid()
	techsupportmanagementtechsupportstatusrelationship["object_type"] = item.GetObjectType()
	techsupportmanagementtechsupportstatusrelationship["selector"] = item.GetSelector()

	techsupportmanagementtechsupportstatusrelationships = append(techsupportmanagementtechsupportstatusrelationships, techsupportmanagementtechsupportstatusrelationship)
	return techsupportmanagementtechsupportstatusrelationships
}
func flattenMapTopSystemRelationship(p models.TopSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var topsystemrelationships []map[string]interface{}
	var ret models.TopSystemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	topsystemrelationship := make(map[string]interface{})
	topsystemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	topsystemrelationship["class_id"] = item.GetClassId()
	topsystemrelationship["moid"] = item.GetMoid()
	topsystemrelationship["object_type"] = item.GetObjectType()
	topsystemrelationship["selector"] = item.GetSelector()

	topsystemrelationships = append(topsystemrelationships, topsystemrelationship)
	return topsystemrelationships
}
func flattenMapUuidpoolBlockRelationship(p models.UuidpoolBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolblockrelationships []map[string]interface{}
	var ret models.UuidpoolBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpoolblockrelationship := make(map[string]interface{})
	uuidpoolblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpoolblockrelationship["class_id"] = item.GetClassId()
	uuidpoolblockrelationship["moid"] = item.GetMoid()
	uuidpoolblockrelationship["object_type"] = item.GetObjectType()
	uuidpoolblockrelationship["selector"] = item.GetSelector()

	uuidpoolblockrelationships = append(uuidpoolblockrelationships, uuidpoolblockrelationship)
	return uuidpoolblockrelationships
}
func flattenMapUuidpoolPoolRelationship(p models.UuidpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolpoolrelationships []map[string]interface{}
	var ret models.UuidpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpoolpoolrelationship := make(map[string]interface{})
	uuidpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpoolpoolrelationship["class_id"] = item.GetClassId()
	uuidpoolpoolrelationship["moid"] = item.GetMoid()
	uuidpoolpoolrelationship["object_type"] = item.GetObjectType()
	uuidpoolpoolrelationship["selector"] = item.GetSelector()

	uuidpoolpoolrelationships = append(uuidpoolpoolrelationships, uuidpoolpoolrelationship)
	return uuidpoolpoolrelationships
}
func flattenMapUuidpoolPoolMemberRelationship(p models.UuidpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolpoolmemberrelationships []map[string]interface{}
	var ret models.UuidpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpoolpoolmemberrelationship := make(map[string]interface{})
	uuidpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpoolpoolmemberrelationship["class_id"] = item.GetClassId()
	uuidpoolpoolmemberrelationship["moid"] = item.GetMoid()
	uuidpoolpoolmemberrelationship["object_type"] = item.GetObjectType()
	uuidpoolpoolmemberrelationship["selector"] = item.GetSelector()

	uuidpoolpoolmemberrelationships = append(uuidpoolpoolmemberrelationships, uuidpoolpoolmemberrelationship)
	return uuidpoolpoolmemberrelationships
}
func flattenMapUuidpoolUniverseRelationship(p models.UuidpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluniverserelationships []map[string]interface{}
	var ret models.UuidpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpooluniverserelationship := make(map[string]interface{})
	uuidpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpooluniverserelationship["class_id"] = item.GetClassId()
	uuidpooluniverserelationship["moid"] = item.GetMoid()
	uuidpooluniverserelationship["object_type"] = item.GetObjectType()
	uuidpooluniverserelationship["selector"] = item.GetSelector()

	uuidpooluniverserelationships = append(uuidpooluniverserelationships, uuidpooluniverserelationship)
	return uuidpooluniverserelationships
}
func flattenMapUuidpoolUuidBlock(p models.UuidpoolUuidBlock, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluuidblocks []map[string]interface{}
	var ret models.UuidpoolUuidBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	uuidpooluuidblock := make(map[string]interface{})
	uuidpooluuidblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpooluuidblock["class_id"] = item.GetClassId()
	uuidpooluuidblock["from"] = item.GetFrom()
	uuidpooluuidblock["object_type"] = item.GetObjectType()
	uuidpooluuidblock["size"] = item.GetSize()
	uuidpooluuidblock["to"] = item.GetTo()

	uuidpooluuidblocks = append(uuidpooluuidblocks, uuidpooluuidblock)
	return uuidpooluuidblocks
}
func flattenMapUuidpoolUuidLeaseRelationship(p models.UuidpoolUuidLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluuidleaserelationships []map[string]interface{}
	var ret models.UuidpoolUuidLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpooluuidleaserelationship := make(map[string]interface{})
	uuidpooluuidleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpooluuidleaserelationship["class_id"] = item.GetClassId()
	uuidpooluuidleaserelationship["moid"] = item.GetMoid()
	uuidpooluuidleaserelationship["object_type"] = item.GetObjectType()
	uuidpooluuidleaserelationship["selector"] = item.GetSelector()

	uuidpooluuidleaserelationships = append(uuidpooluuidleaserelationships, uuidpooluuidleaserelationship)
	return uuidpooluuidleaserelationships
}
func flattenMapVirtualizationActionInfo(p models.VirtualizationActionInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationactioninfos []map[string]interface{}
	var ret models.VirtualizationActionInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationactioninfo := make(map[string]interface{})
	virtualizationactioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationactioninfo["class_id"] = item.GetClassId()
	virtualizationactioninfo["failure_reason"] = item.GetFailureReason()
	virtualizationactioninfo["name"] = item.GetName()
	virtualizationactioninfo["object_type"] = item.GetObjectType()
	virtualizationactioninfo["status"] = item.GetStatus()

	virtualizationactioninfos = append(virtualizationactioninfos, virtualizationactioninfo)
	return virtualizationactioninfos
}
func flattenMapVirtualizationBaseClusterRelationship(p models.VirtualizationBaseClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbaseclusterrelationships []map[string]interface{}
	var ret models.VirtualizationBaseClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationbaseclusterrelationship := make(map[string]interface{})
	virtualizationbaseclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationbaseclusterrelationship["class_id"] = item.GetClassId()
	virtualizationbaseclusterrelationship["moid"] = item.GetMoid()
	virtualizationbaseclusterrelationship["object_type"] = item.GetObjectType()
	virtualizationbaseclusterrelationship["selector"] = item.GetSelector()

	virtualizationbaseclusterrelationships = append(virtualizationbaseclusterrelationships, virtualizationbaseclusterrelationship)
	return virtualizationbaseclusterrelationships
}
func flattenMapVirtualizationBaseHostRelationship(p models.VirtualizationBaseHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbasehostrelationships []map[string]interface{}
	var ret models.VirtualizationBaseHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationbasehostrelationship := make(map[string]interface{})
	virtualizationbasehostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationbasehostrelationship["class_id"] = item.GetClassId()
	virtualizationbasehostrelationship["moid"] = item.GetMoid()
	virtualizationbasehostrelationship["object_type"] = item.GetObjectType()
	virtualizationbasehostrelationship["selector"] = item.GetSelector()

	virtualizationbasehostrelationships = append(virtualizationbasehostrelationships, virtualizationbasehostrelationship)
	return virtualizationbasehostrelationships
}
func flattenMapVirtualizationBaseNetworkRelationship(p models.VirtualizationBaseNetworkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbasenetworkrelationships []map[string]interface{}
	var ret models.VirtualizationBaseNetworkRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationbasenetworkrelationship := make(map[string]interface{})
	virtualizationbasenetworkrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationbasenetworkrelationship["class_id"] = item.GetClassId()
	virtualizationbasenetworkrelationship["moid"] = item.GetMoid()
	virtualizationbasenetworkrelationship["object_type"] = item.GetObjectType()
	virtualizationbasenetworkrelationship["selector"] = item.GetSelector()

	virtualizationbasenetworkrelationships = append(virtualizationbasenetworkrelationships, virtualizationbasenetworkrelationship)
	return virtualizationbasenetworkrelationships
}
func flattenMapVirtualizationBaseVirtualDiskRelationship(p models.VirtualizationBaseVirtualDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbasevirtualdiskrelationships []map[string]interface{}
	var ret models.VirtualizationBaseVirtualDiskRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationbasevirtualdiskrelationship := make(map[string]interface{})
	virtualizationbasevirtualdiskrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationbasevirtualdiskrelationship["class_id"] = item.GetClassId()
	virtualizationbasevirtualdiskrelationship["moid"] = item.GetMoid()
	virtualizationbasevirtualdiskrelationship["object_type"] = item.GetObjectType()
	virtualizationbasevirtualdiskrelationship["selector"] = item.GetSelector()

	virtualizationbasevirtualdiskrelationships = append(virtualizationbasevirtualdiskrelationships, virtualizationbasevirtualdiskrelationship)
	return virtualizationbasevirtualdiskrelationships
}
func flattenMapVirtualizationBaseVirtualMachineRelationship(p models.VirtualizationBaseVirtualMachineRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbasevirtualmachinerelationships []map[string]interface{}
	var ret models.VirtualizationBaseVirtualMachineRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationbasevirtualmachinerelationship := make(map[string]interface{})
	virtualizationbasevirtualmachinerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationbasevirtualmachinerelationship["class_id"] = item.GetClassId()
	virtualizationbasevirtualmachinerelationship["moid"] = item.GetMoid()
	virtualizationbasevirtualmachinerelationship["object_type"] = item.GetObjectType()
	virtualizationbasevirtualmachinerelationship["selector"] = item.GetSelector()

	virtualizationbasevirtualmachinerelationships = append(virtualizationbasevirtualmachinerelationships, virtualizationbasevirtualmachinerelationship)
	return virtualizationbasevirtualmachinerelationships
}
func flattenMapVirtualizationBaseVmConfiguration(p models.VirtualizationBaseVmConfiguration, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationbasevmconfigurations []map[string]interface{}
	var ret models.VirtualizationBaseVmConfiguration
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationbasevmconfiguration := make(map[string]interface{})
	virtualizationbasevmconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationbasevmconfiguration["class_id"] = item.GetClassId()
	virtualizationbasevmconfiguration["object_type"] = item.GetObjectType()

	virtualizationbasevmconfigurations = append(virtualizationbasevmconfigurations, virtualizationbasevmconfiguration)
	return virtualizationbasevmconfigurations
}
func flattenMapVirtualizationCloudInitConfig(p models.VirtualizationCloudInitConfig, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcloudinitconfigs []map[string]interface{}
	var ret models.VirtualizationCloudInitConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationcloudinitconfig := make(map[string]interface{})
	virtualizationcloudinitconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationcloudinitconfig["class_id"] = item.GetClassId()
	virtualizationcloudinitconfig["config_type"] = item.GetConfigType()
	virtualizationcloudinitconfig["network_data"] = item.GetNetworkData()
	virtualizationcloudinitconfig["network_data_base64_encoded"] = item.GetNetworkDataBase64Encoded()
	virtualizationcloudinitconfig["object_type"] = item.GetObjectType()
	virtualizationcloudinitconfig["user_data"] = item.GetUserData()
	virtualizationcloudinitconfig["user_data_base64_encoded"] = item.GetUserDataBase64Encoded()

	virtualizationcloudinitconfigs = append(virtualizationcloudinitconfigs, virtualizationcloudinitconfig)
	return virtualizationcloudinitconfigs
}
func flattenMapVirtualizationComputeCapacity(p models.VirtualizationComputeCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcomputecapacitys []map[string]interface{}
	var ret models.VirtualizationComputeCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationcomputecapacity := make(map[string]interface{})
	virtualizationcomputecapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationcomputecapacity["capacity"] = item.GetCapacity()
	virtualizationcomputecapacity["class_id"] = item.GetClassId()
	virtualizationcomputecapacity["free"] = item.GetFree()
	virtualizationcomputecapacity["object_type"] = item.GetObjectType()
	virtualizationcomputecapacity["used"] = item.GetUsed()

	virtualizationcomputecapacitys = append(virtualizationcomputecapacitys, virtualizationcomputecapacity)
	return virtualizationcomputecapacitys
}
func flattenMapVirtualizationCpuAllocation(p models.VirtualizationCpuAllocation, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcpuallocations []map[string]interface{}
	var ret models.VirtualizationCpuAllocation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationcpuallocation := make(map[string]interface{})
	virtualizationcpuallocation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationcpuallocation["class_id"] = item.GetClassId()
	virtualizationcpuallocation["free"] = item.GetFree()
	virtualizationcpuallocation["object_type"] = item.GetObjectType()
	virtualizationcpuallocation["reserved"] = item.GetReserved()
	virtualizationcpuallocation["total"] = item.GetTotal()
	virtualizationcpuallocation["used"] = item.GetUsed()

	virtualizationcpuallocations = append(virtualizationcpuallocations, virtualizationcpuallocation)
	return virtualizationcpuallocations
}
func flattenMapVirtualizationCpuInfo(p models.VirtualizationCpuInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcpuinfos []map[string]interface{}
	var ret models.VirtualizationCpuInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationcpuinfo := make(map[string]interface{})
	virtualizationcpuinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationcpuinfo["class_id"] = item.GetClassId()
	virtualizationcpuinfo["cores"] = item.GetCores()
	virtualizationcpuinfo["description"] = item.GetDescription()
	virtualizationcpuinfo["object_type"] = item.GetObjectType()
	virtualizationcpuinfo["sockets"] = item.GetSockets()
	virtualizationcpuinfo["speed"] = item.GetSpeed()
	virtualizationcpuinfo["vendor"] = item.GetVendor()

	virtualizationcpuinfos = append(virtualizationcpuinfos, virtualizationcpuinfo)
	return virtualizationcpuinfos
}
func flattenMapVirtualizationGuestInfo(p models.VirtualizationGuestInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationguestinfos []map[string]interface{}
	var ret models.VirtualizationGuestInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationguestinfo := make(map[string]interface{})
	virtualizationguestinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationguestinfo["class_id"] = item.GetClassId()
	virtualizationguestinfo["hostname"] = item.GetHostname()
	virtualizationguestinfo["ip_address"] = item.GetIpAddress()
	virtualizationguestinfo["name"] = item.GetName()
	virtualizationguestinfo["object_type"] = item.GetObjectType()
	virtualizationguestinfo["operating_system"] = item.GetOperatingSystem()

	virtualizationguestinfos = append(virtualizationguestinfos, virtualizationguestinfo)
	return virtualizationguestinfos
}
func flattenMapVirtualizationMemoryAllocation(p models.VirtualizationMemoryAllocation, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationmemoryallocations []map[string]interface{}
	var ret models.VirtualizationMemoryAllocation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationmemoryallocation := make(map[string]interface{})
	virtualizationmemoryallocation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationmemoryallocation["class_id"] = item.GetClassId()
	virtualizationmemoryallocation["free"] = item.GetFree()
	virtualizationmemoryallocation["object_type"] = item.GetObjectType()
	virtualizationmemoryallocation["reserved"] = item.GetReserved()
	virtualizationmemoryallocation["total"] = item.GetTotal()
	virtualizationmemoryallocation["used"] = item.GetUsed()

	virtualizationmemoryallocations = append(virtualizationmemoryallocations, virtualizationmemoryallocation)
	return virtualizationmemoryallocations
}
func flattenMapVirtualizationMemoryCapacity(p models.VirtualizationMemoryCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationmemorycapacitys []map[string]interface{}
	var ret models.VirtualizationMemoryCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationmemorycapacity := make(map[string]interface{})
	virtualizationmemorycapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationmemorycapacity["capacity"] = item.GetCapacity()
	virtualizationmemorycapacity["class_id"] = item.GetClassId()
	virtualizationmemorycapacity["free"] = item.GetFree()
	virtualizationmemorycapacity["object_type"] = item.GetObjectType()
	virtualizationmemorycapacity["used"] = item.GetUsed()

	virtualizationmemorycapacitys = append(virtualizationmemorycapacitys, virtualizationmemorycapacity)
	return virtualizationmemorycapacitys
}
func flattenMapVirtualizationProductInfo(p models.VirtualizationProductInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationproductinfos []map[string]interface{}
	var ret models.VirtualizationProductInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationproductinfo := make(map[string]interface{})
	virtualizationproductinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationproductinfo["build"] = item.GetBuild()
	virtualizationproductinfo["class_id"] = item.GetClassId()
	virtualizationproductinfo["object_type"] = item.GetObjectType()
	virtualizationproductinfo["product_name"] = item.GetProductName()
	virtualizationproductinfo["product_type"] = item.GetProductType()
	virtualizationproductinfo["product_vendor"] = item.GetProductVendor()
	virtualizationproductinfo["nr_version"] = item.GetVersion()

	virtualizationproductinfos = append(virtualizationproductinfos, virtualizationproductinfo)
	return virtualizationproductinfos
}
func flattenMapVirtualizationStorageCapacity(p models.VirtualizationStorageCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationstoragecapacitys []map[string]interface{}
	var ret models.VirtualizationStorageCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationstoragecapacity := make(map[string]interface{})
	virtualizationstoragecapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationstoragecapacity["capacity"] = item.GetCapacity()
	virtualizationstoragecapacity["class_id"] = item.GetClassId()
	virtualizationstoragecapacity["free"] = item.GetFree()
	virtualizationstoragecapacity["object_type"] = item.GetObjectType()
	virtualizationstoragecapacity["used"] = item.GetUsed()

	virtualizationstoragecapacitys = append(virtualizationstoragecapacitys, virtualizationstoragecapacity)
	return virtualizationstoragecapacitys
}
func flattenMapVirtualizationVirtualMachineRelationship(p models.VirtualizationVirtualMachineRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvirtualmachinerelationships []map[string]interface{}
	var ret models.VirtualizationVirtualMachineRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvirtualmachinerelationship := make(map[string]interface{})
	virtualizationvirtualmachinerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvirtualmachinerelationship["class_id"] = item.GetClassId()
	virtualizationvirtualmachinerelationship["moid"] = item.GetMoid()
	virtualizationvirtualmachinerelationship["object_type"] = item.GetObjectType()
	virtualizationvirtualmachinerelationship["selector"] = item.GetSelector()

	virtualizationvirtualmachinerelationships = append(virtualizationvirtualmachinerelationships, virtualizationvirtualmachinerelationship)
	return virtualizationvirtualmachinerelationships
}
func flattenMapVirtualizationVmwareClusterRelationship(p models.VirtualizationVmwareClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareclusterrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwareclusterrelationship := make(map[string]interface{})
	virtualizationvmwareclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareclusterrelationship["class_id"] = item.GetClassId()
	virtualizationvmwareclusterrelationship["moid"] = item.GetMoid()
	virtualizationvmwareclusterrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwareclusterrelationship["selector"] = item.GetSelector()

	virtualizationvmwareclusterrelationships = append(virtualizationvmwareclusterrelationships, virtualizationvmwareclusterrelationship)
	return virtualizationvmwareclusterrelationships
}
func flattenMapVirtualizationVmwareDatacenterRelationship(p models.VirtualizationVmwareDatacenterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatacenterrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareDatacenterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwaredatacenterrelationship := make(map[string]interface{})
	virtualizationvmwaredatacenterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwaredatacenterrelationship["class_id"] = item.GetClassId()
	virtualizationvmwaredatacenterrelationship["moid"] = item.GetMoid()
	virtualizationvmwaredatacenterrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwaredatacenterrelationship["selector"] = item.GetSelector()

	virtualizationvmwaredatacenterrelationships = append(virtualizationvmwaredatacenterrelationships, virtualizationvmwaredatacenterrelationship)
	return virtualizationvmwaredatacenterrelationships
}
func flattenMapVirtualizationVmwareDatastoreRelationship(p models.VirtualizationVmwareDatastoreRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatastorerelationships []map[string]interface{}
	var ret models.VirtualizationVmwareDatastoreRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwaredatastorerelationship := make(map[string]interface{})
	virtualizationvmwaredatastorerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwaredatastorerelationship["class_id"] = item.GetClassId()
	virtualizationvmwaredatastorerelationship["moid"] = item.GetMoid()
	virtualizationvmwaredatastorerelationship["object_type"] = item.GetObjectType()
	virtualizationvmwaredatastorerelationship["selector"] = item.GetSelector()

	virtualizationvmwaredatastorerelationships = append(virtualizationvmwaredatastorerelationships, virtualizationvmwaredatastorerelationship)
	return virtualizationvmwaredatastorerelationships
}
func flattenMapVirtualizationVmwareDistributedNetworkRelationship(p models.VirtualizationVmwareDistributedNetworkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredistributednetworkrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareDistributedNetworkRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwaredistributednetworkrelationship := make(map[string]interface{})
	virtualizationvmwaredistributednetworkrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwaredistributednetworkrelationship["class_id"] = item.GetClassId()
	virtualizationvmwaredistributednetworkrelationship["moid"] = item.GetMoid()
	virtualizationvmwaredistributednetworkrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwaredistributednetworkrelationship["selector"] = item.GetSelector()

	virtualizationvmwaredistributednetworkrelationships = append(virtualizationvmwaredistributednetworkrelationships, virtualizationvmwaredistributednetworkrelationship)
	return virtualizationvmwaredistributednetworkrelationships
}
func flattenMapVirtualizationVmwareDistributedSwitchRelationship(p models.VirtualizationVmwareDistributedSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredistributedswitchrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareDistributedSwitchRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwaredistributedswitchrelationship := make(map[string]interface{})
	virtualizationvmwaredistributedswitchrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwaredistributedswitchrelationship["class_id"] = item.GetClassId()
	virtualizationvmwaredistributedswitchrelationship["moid"] = item.GetMoid()
	virtualizationvmwaredistributedswitchrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwaredistributedswitchrelationship["selector"] = item.GetSelector()

	virtualizationvmwaredistributedswitchrelationships = append(virtualizationvmwaredistributedswitchrelationships, virtualizationvmwaredistributedswitchrelationship)
	return virtualizationvmwaredistributedswitchrelationships
}
func flattenMapVirtualizationVmwareFolderRelationship(p models.VirtualizationVmwareFolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarefolderrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareFolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarefolderrelationship := make(map[string]interface{})
	virtualizationvmwarefolderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarefolderrelationship["class_id"] = item.GetClassId()
	virtualizationvmwarefolderrelationship["moid"] = item.GetMoid()
	virtualizationvmwarefolderrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarefolderrelationship["selector"] = item.GetSelector()

	virtualizationvmwarefolderrelationships = append(virtualizationvmwarefolderrelationships, virtualizationvmwarefolderrelationship)
	return virtualizationvmwarefolderrelationships
}
func flattenMapVirtualizationVmwareHostRelationship(p models.VirtualizationVmwareHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarehostrelationship := make(map[string]interface{})
	virtualizationvmwarehostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarehostrelationship["class_id"] = item.GetClassId()
	virtualizationvmwarehostrelationship["moid"] = item.GetMoid()
	virtualizationvmwarehostrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarehostrelationship["selector"] = item.GetSelector()

	virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	return virtualizationvmwarehostrelationships
}
func flattenMapVirtualizationVmwareNetworkRelationship(p models.VirtualizationVmwareNetworkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarenetworkrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareNetworkRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarenetworkrelationship := make(map[string]interface{})
	virtualizationvmwarenetworkrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarenetworkrelationship["class_id"] = item.GetClassId()
	virtualizationvmwarenetworkrelationship["moid"] = item.GetMoid()
	virtualizationvmwarenetworkrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarenetworkrelationship["selector"] = item.GetSelector()

	virtualizationvmwarenetworkrelationships = append(virtualizationvmwarenetworkrelationships, virtualizationvmwarenetworkrelationship)
	return virtualizationvmwarenetworkrelationships
}
func flattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(p models.VirtualizationVmwarePhysicalNetworkInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarephysicalnetworkinterfacerelationships []map[string]interface{}
	var ret models.VirtualizationVmwarePhysicalNetworkInterfaceRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarephysicalnetworkinterfacerelationship := make(map[string]interface{})
	virtualizationvmwarephysicalnetworkinterfacerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarephysicalnetworkinterfacerelationship["class_id"] = item.GetClassId()
	virtualizationvmwarephysicalnetworkinterfacerelationship["moid"] = item.GetMoid()
	virtualizationvmwarephysicalnetworkinterfacerelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarephysicalnetworkinterfacerelationship["selector"] = item.GetSelector()

	virtualizationvmwarephysicalnetworkinterfacerelationships = append(virtualizationvmwarephysicalnetworkinterfacerelationships, virtualizationvmwarephysicalnetworkinterfacerelationship)
	return virtualizationvmwarephysicalnetworkinterfacerelationships
}
func flattenMapVirtualizationVmwareRemoteDisplayInfo(p models.VirtualizationVmwareRemoteDisplayInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareremotedisplayinfos []map[string]interface{}
	var ret models.VirtualizationVmwareRemoteDisplayInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwareremotedisplayinfo := make(map[string]interface{})
	virtualizationvmwareremotedisplayinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareremotedisplayinfo["class_id"] = item.GetClassId()
	virtualizationvmwareremotedisplayinfo["object_type"] = item.GetObjectType()
	virtualizationvmwareremotedisplayinfo["remote_display_password"] = item.GetRemoteDisplayPassword()
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_key"] = item.GetRemoteDisplayVncKey()
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_port"] = item.GetRemoteDisplayVncPort()

	virtualizationvmwareremotedisplayinfos = append(virtualizationvmwareremotedisplayinfos, virtualizationvmwareremotedisplayinfo)
	return virtualizationvmwareremotedisplayinfos
}
func flattenMapVirtualizationVmwareResourceConsumption(p models.VirtualizationVmwareResourceConsumption, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareresourceconsumptions []map[string]interface{}
	var ret models.VirtualizationVmwareResourceConsumption
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwareresourceconsumption := make(map[string]interface{})
	virtualizationvmwareresourceconsumption["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareresourceconsumption["class_id"] = item.GetClassId()
	virtualizationvmwareresourceconsumption["cpu_consumed"] = item.GetCpuConsumed()
	virtualizationvmwareresourceconsumption["memory_consumed"] = item.GetMemoryConsumed()
	virtualizationvmwareresourceconsumption["object_type"] = item.GetObjectType()

	virtualizationvmwareresourceconsumptions = append(virtualizationvmwareresourceconsumptions, virtualizationvmwareresourceconsumption)
	return virtualizationvmwareresourceconsumptions
}
func flattenMapVirtualizationVmwareSharesInfo(p models.VirtualizationVmwareSharesInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaresharesinfos []map[string]interface{}
	var ret models.VirtualizationVmwareSharesInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwaresharesinfo := make(map[string]interface{})
	virtualizationvmwaresharesinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwaresharesinfo["class_id"] = item.GetClassId()
	virtualizationvmwaresharesinfo["level"] = item.GetLevel()
	virtualizationvmwaresharesinfo["object_type"] = item.GetObjectType()
	virtualizationvmwaresharesinfo["shares"] = item.GetShares()

	virtualizationvmwaresharesinfos = append(virtualizationvmwaresharesinfos, virtualizationvmwaresharesinfo)
	return virtualizationvmwaresharesinfos
}
func flattenMapVirtualizationVmwareTeamingAndFailover(p models.VirtualizationVmwareTeamingAndFailover, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareteamingandfailovers []map[string]interface{}
	var ret models.VirtualizationVmwareTeamingAndFailover
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwareteamingandfailover := make(map[string]interface{})
	virtualizationvmwareteamingandfailover["active_adapters"] = item.GetActiveAdapters()
	virtualizationvmwareteamingandfailover["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareteamingandfailover["class_id"] = item.GetClassId()
	virtualizationvmwareteamingandfailover["failback"] = item.GetFailback()
	virtualizationvmwareteamingandfailover["load_balancing"] = item.GetLoadBalancing()
	virtualizationvmwareteamingandfailover["name"] = item.GetName()
	virtualizationvmwareteamingandfailover["network_failure_detection"] = item.GetNetworkFailureDetection()
	virtualizationvmwareteamingandfailover["notify_switches"] = item.GetNotifySwitches()
	virtualizationvmwareteamingandfailover["object_type"] = item.GetObjectType()
	virtualizationvmwareteamingandfailover["standby_adapters"] = item.GetStandbyAdapters()

	virtualizationvmwareteamingandfailovers = append(virtualizationvmwareteamingandfailovers, virtualizationvmwareteamingandfailover)
	return virtualizationvmwareteamingandfailovers
}
func flattenMapVirtualizationVmwareVcenterRelationship(p models.VirtualizationVmwareVcenterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevcenterrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareVcenterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarevcenterrelationship := make(map[string]interface{})
	virtualizationvmwarevcenterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevcenterrelationship["class_id"] = item.GetClassId()
	virtualizationvmwarevcenterrelationship["moid"] = item.GetMoid()
	virtualizationvmwarevcenterrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarevcenterrelationship["selector"] = item.GetSelector()

	virtualizationvmwarevcenterrelationships = append(virtualizationvmwarevcenterrelationships, virtualizationvmwarevcenterrelationship)
	return virtualizationvmwarevcenterrelationships
}
func flattenMapVirtualizationVmwareVirtualMachineRelationship(p models.VirtualizationVmwareVirtualMachineRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevirtualmachinerelationships []map[string]interface{}
	var ret models.VirtualizationVmwareVirtualMachineRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarevirtualmachinerelationship := make(map[string]interface{})
	virtualizationvmwarevirtualmachinerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevirtualmachinerelationship["class_id"] = item.GetClassId()
	virtualizationvmwarevirtualmachinerelationship["moid"] = item.GetMoid()
	virtualizationvmwarevirtualmachinerelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarevirtualmachinerelationship["selector"] = item.GetSelector()

	virtualizationvmwarevirtualmachinerelationships = append(virtualizationvmwarevirtualmachinerelationships, virtualizationvmwarevirtualmachinerelationship)
	return virtualizationvmwarevirtualmachinerelationships
}
func flattenMapVirtualizationVmwareVirtualSwitchRelationship(p models.VirtualizationVmwareVirtualSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevirtualswitchrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareVirtualSwitchRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarevirtualswitchrelationship := make(map[string]interface{})
	virtualizationvmwarevirtualswitchrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevirtualswitchrelationship["class_id"] = item.GetClassId()
	virtualizationvmwarevirtualswitchrelationship["moid"] = item.GetMoid()
	virtualizationvmwarevirtualswitchrelationship["object_type"] = item.GetObjectType()
	virtualizationvmwarevirtualswitchrelationship["selector"] = item.GetSelector()

	virtualizationvmwarevirtualswitchrelationships = append(virtualizationvmwarevirtualswitchrelationships, virtualizationvmwarevirtualswitchrelationship)
	return virtualizationvmwarevirtualswitchrelationships
}
func flattenMapVirtualizationVmwareVmCpuShareInfo(p models.VirtualizationVmwareVmCpuShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmcpushareinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmCpuShareInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmcpushareinfo := make(map[string]interface{})
	virtualizationvmwarevmcpushareinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmcpushareinfo["class_id"] = item.GetClassId()
	virtualizationvmwarevmcpushareinfo["cpu_limit"] = item.GetCpuLimit()
	virtualizationvmwarevmcpushareinfo["cpu_overhead_limit"] = item.GetCpuOverheadLimit()
	virtualizationvmwarevmcpushareinfo["cpu_reservation"] = item.GetCpuReservation()
	virtualizationvmwarevmcpushareinfo["cpu_shares"] = item.GetCpuShares()
	virtualizationvmwarevmcpushareinfo["object_type"] = item.GetObjectType()

	virtualizationvmwarevmcpushareinfos = append(virtualizationvmwarevmcpushareinfos, virtualizationvmwarevmcpushareinfo)
	return virtualizationvmwarevmcpushareinfos
}
func flattenMapVirtualizationVmwareVmCpuSocketInfo(p models.VirtualizationVmwareVmCpuSocketInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmcpusocketinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmCpuSocketInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmcpusocketinfo := make(map[string]interface{})
	virtualizationvmwarevmcpusocketinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmcpusocketinfo["class_id"] = item.GetClassId()
	virtualizationvmwarevmcpusocketinfo["cores_per_socket"] = item.GetCoresPerSocket()
	virtualizationvmwarevmcpusocketinfo["num_cpus"] = item.GetNumCpus()
	virtualizationvmwarevmcpusocketinfo["num_sockets"] = item.GetNumSockets()
	virtualizationvmwarevmcpusocketinfo["object_type"] = item.GetObjectType()

	virtualizationvmwarevmcpusocketinfos = append(virtualizationvmwarevmcpusocketinfos, virtualizationvmwarevmcpusocketinfo)
	return virtualizationvmwarevmcpusocketinfos
}
func flattenMapVirtualizationVmwareVmDiskCommitInfo(p models.VirtualizationVmwareVmDiskCommitInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmdiskcommitinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmDiskCommitInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmdiskcommitinfo := make(map[string]interface{})
	virtualizationvmwarevmdiskcommitinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmdiskcommitinfo["class_id"] = item.GetClassId()
	virtualizationvmwarevmdiskcommitinfo["committed_disk"] = item.GetCommittedDisk()
	virtualizationvmwarevmdiskcommitinfo["object_type"] = item.GetObjectType()
	virtualizationvmwarevmdiskcommitinfo["un_committed_disk"] = item.GetUnCommittedDisk()
	virtualizationvmwarevmdiskcommitinfo["unshared_disk"] = item.GetUnsharedDisk()

	virtualizationvmwarevmdiskcommitinfos = append(virtualizationvmwarevmdiskcommitinfos, virtualizationvmwarevmdiskcommitinfo)
	return virtualizationvmwarevmdiskcommitinfos
}
func flattenMapVirtualizationVmwareVmMemoryShareInfo(p models.VirtualizationVmwareVmMemoryShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmmemoryshareinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmMemoryShareInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmmemoryshareinfo := make(map[string]interface{})
	virtualizationvmwarevmmemoryshareinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmmemoryshareinfo["class_id"] = item.GetClassId()
	virtualizationvmwarevmmemoryshareinfo["mem_limit"] = item.GetMemLimit()
	virtualizationvmwarevmmemoryshareinfo["mem_overhead_limit"] = item.GetMemOverheadLimit()
	virtualizationvmwarevmmemoryshareinfo["mem_reservation"] = item.GetMemReservation()
	virtualizationvmwarevmmemoryshareinfo["mem_shares"] = item.GetMemShares()
	virtualizationvmwarevmmemoryshareinfo["object_type"] = item.GetObjectType()

	virtualizationvmwarevmmemoryshareinfos = append(virtualizationvmwarevmmemoryshareinfos, virtualizationvmwarevmmemoryshareinfo)
	return virtualizationvmwarevmmemoryshareinfos
}
func flattenMapVnicArfsSettings(p models.VnicArfsSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicarfssettingss []map[string]interface{}
	var ret models.VnicArfsSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicarfssettings := make(map[string]interface{})
	vnicarfssettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicarfssettings["class_id"] = item.GetClassId()
	vnicarfssettings["enabled"] = item.GetEnabled()
	vnicarfssettings["object_type"] = item.GetObjectType()

	vnicarfssettingss = append(vnicarfssettingss, vnicarfssettings)
	return vnicarfssettingss
}
func flattenMapVnicCdn(p models.VnicCdn, d *schema.ResourceData) []map[string]interface{} {
	var vniccdns []map[string]interface{}
	var ret models.VnicCdn
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vniccdn := make(map[string]interface{})
	vniccdn["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniccdn["class_id"] = item.GetClassId()
	vniccdn["object_type"] = item.GetObjectType()
	vniccdn["nr_source"] = item.GetSource()
	vniccdn["value"] = item.GetValue()

	vniccdns = append(vniccdns, vniccdn)
	return vniccdns
}
func flattenMapVnicCompletionQueueSettings(p models.VnicCompletionQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vniccompletionqueuesettingss []map[string]interface{}
	var ret models.VnicCompletionQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vniccompletionqueuesettings := make(map[string]interface{})
	vniccompletionqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniccompletionqueuesettings["class_id"] = item.GetClassId()
	vniccompletionqueuesettings["nr_count"] = item.GetCount()
	vniccompletionqueuesettings["object_type"] = item.GetObjectType()
	vniccompletionqueuesettings["ring_size"] = item.GetRingSize()

	vniccompletionqueuesettingss = append(vniccompletionqueuesettingss, vniccompletionqueuesettings)
	return vniccompletionqueuesettingss
}
func flattenMapVnicEthAdapterPolicyRelationship(p models.VnicEthAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethadapterpolicyrelationships []map[string]interface{}
	var ret models.VnicEthAdapterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethadapterpolicyrelationship := make(map[string]interface{})
	vnicethadapterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethadapterpolicyrelationship["class_id"] = item.GetClassId()
	vnicethadapterpolicyrelationship["moid"] = item.GetMoid()
	vnicethadapterpolicyrelationship["object_type"] = item.GetObjectType()
	vnicethadapterpolicyrelationship["selector"] = item.GetSelector()

	vnicethadapterpolicyrelationships = append(vnicethadapterpolicyrelationships, vnicethadapterpolicyrelationship)
	return vnicethadapterpolicyrelationships
}
func flattenMapVnicEthIfRelationship(p models.VnicEthIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethifrelationships []map[string]interface{}
	var ret models.VnicEthIfRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethifrelationship := make(map[string]interface{})
	vnicethifrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethifrelationship["class_id"] = item.GetClassId()
	vnicethifrelationship["moid"] = item.GetMoid()
	vnicethifrelationship["object_type"] = item.GetObjectType()
	vnicethifrelationship["selector"] = item.GetSelector()

	vnicethifrelationships = append(vnicethifrelationships, vnicethifrelationship)
	return vnicethifrelationships
}
func flattenMapVnicEthInterruptSettings(p models.VnicEthInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethinterruptsettingss []map[string]interface{}
	var ret models.VnicEthInterruptSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicethinterruptsettings := make(map[string]interface{})
	vnicethinterruptsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethinterruptsettings["class_id"] = item.GetClassId()
	vnicethinterruptsettings["coalescing_time"] = item.GetCoalescingTime()
	vnicethinterruptsettings["coalescing_type"] = item.GetCoalescingType()
	vnicethinterruptsettings["nr_count"] = item.GetCount()
	vnicethinterruptsettings["mode"] = item.GetMode()
	vnicethinterruptsettings["object_type"] = item.GetObjectType()

	vnicethinterruptsettingss = append(vnicethinterruptsettingss, vnicethinterruptsettings)
	return vnicethinterruptsettingss
}
func flattenMapVnicEthNetworkPolicyRelationship(p models.VnicEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethnetworkpolicyrelationships []map[string]interface{}
	var ret models.VnicEthNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethnetworkpolicyrelationship := make(map[string]interface{})
	vnicethnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethnetworkpolicyrelationship["class_id"] = item.GetClassId()
	vnicethnetworkpolicyrelationship["moid"] = item.GetMoid()
	vnicethnetworkpolicyrelationship["object_type"] = item.GetObjectType()
	vnicethnetworkpolicyrelationship["selector"] = item.GetSelector()

	vnicethnetworkpolicyrelationships = append(vnicethnetworkpolicyrelationships, vnicethnetworkpolicyrelationship)
	return vnicethnetworkpolicyrelationships
}
func flattenMapVnicEthQosPolicyRelationship(p models.VnicEthQosPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethqospolicyrelationships []map[string]interface{}
	var ret models.VnicEthQosPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethqospolicyrelationship := make(map[string]interface{})
	vnicethqospolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethqospolicyrelationship["class_id"] = item.GetClassId()
	vnicethqospolicyrelationship["moid"] = item.GetMoid()
	vnicethqospolicyrelationship["object_type"] = item.GetObjectType()
	vnicethqospolicyrelationship["selector"] = item.GetSelector()

	vnicethqospolicyrelationships = append(vnicethqospolicyrelationships, vnicethqospolicyrelationship)
	return vnicethqospolicyrelationships
}
func flattenMapVnicEthRxQueueSettings(p models.VnicEthRxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethrxqueuesettingss []map[string]interface{}
	var ret models.VnicEthRxQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicethrxqueuesettings := make(map[string]interface{})
	vnicethrxqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethrxqueuesettings["class_id"] = item.GetClassId()
	vnicethrxqueuesettings["nr_count"] = item.GetCount()
	vnicethrxqueuesettings["object_type"] = item.GetObjectType()
	vnicethrxqueuesettings["ring_size"] = item.GetRingSize()

	vnicethrxqueuesettingss = append(vnicethrxqueuesettingss, vnicethrxqueuesettings)
	return vnicethrxqueuesettingss
}
func flattenMapVnicEthTxQueueSettings(p models.VnicEthTxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethtxqueuesettingss []map[string]interface{}
	var ret models.VnicEthTxQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicethtxqueuesettings := make(map[string]interface{})
	vnicethtxqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethtxqueuesettings["class_id"] = item.GetClassId()
	vnicethtxqueuesettings["nr_count"] = item.GetCount()
	vnicethtxqueuesettings["object_type"] = item.GetObjectType()
	vnicethtxqueuesettings["ring_size"] = item.GetRingSize()

	vnicethtxqueuesettingss = append(vnicethtxqueuesettingss, vnicethtxqueuesettings)
	return vnicethtxqueuesettingss
}
func flattenMapVnicFcAdapterPolicyRelationship(p models.VnicFcAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcadapterpolicyrelationships []map[string]interface{}
	var ret models.VnicFcAdapterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcadapterpolicyrelationship := make(map[string]interface{})
	vnicfcadapterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcadapterpolicyrelationship["class_id"] = item.GetClassId()
	vnicfcadapterpolicyrelationship["moid"] = item.GetMoid()
	vnicfcadapterpolicyrelationship["object_type"] = item.GetObjectType()
	vnicfcadapterpolicyrelationship["selector"] = item.GetSelector()

	vnicfcadapterpolicyrelationships = append(vnicfcadapterpolicyrelationships, vnicfcadapterpolicyrelationship)
	return vnicfcadapterpolicyrelationships
}
func flattenMapVnicFcErrorRecoverySettings(p models.VnicFcErrorRecoverySettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcerrorrecoverysettingss []map[string]interface{}
	var ret models.VnicFcErrorRecoverySettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicfcerrorrecoverysettings := make(map[string]interface{})
	vnicfcerrorrecoverysettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcerrorrecoverysettings["class_id"] = item.GetClassId()
	vnicfcerrorrecoverysettings["enabled"] = item.GetEnabled()
	vnicfcerrorrecoverysettings["io_retry_count"] = item.GetIoRetryCount()
	vnicfcerrorrecoverysettings["io_retry_timeout"] = item.GetIoRetryTimeout()
	vnicfcerrorrecoverysettings["link_down_timeout"] = item.GetLinkDownTimeout()
	vnicfcerrorrecoverysettings["object_type"] = item.GetObjectType()
	vnicfcerrorrecoverysettings["port_down_timeout"] = item.GetPortDownTimeout()

	vnicfcerrorrecoverysettingss = append(vnicfcerrorrecoverysettingss, vnicfcerrorrecoverysettings)
	return vnicfcerrorrecoverysettingss
}
func flattenMapVnicFcIfRelationship(p models.VnicFcIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcifrelationships []map[string]interface{}
	var ret models.VnicFcIfRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcifrelationship := make(map[string]interface{})
	vnicfcifrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcifrelationship["class_id"] = item.GetClassId()
	vnicfcifrelationship["moid"] = item.GetMoid()
	vnicfcifrelationship["object_type"] = item.GetObjectType()
	vnicfcifrelationship["selector"] = item.GetSelector()

	vnicfcifrelationships = append(vnicfcifrelationships, vnicfcifrelationship)
	return vnicfcifrelationships
}
func flattenMapVnicFcInterruptSettings(p models.VnicFcInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcinterruptsettingss []map[string]interface{}
	var ret models.VnicFcInterruptSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicfcinterruptsettings := make(map[string]interface{})
	vnicfcinterruptsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcinterruptsettings["class_id"] = item.GetClassId()
	vnicfcinterruptsettings["mode"] = item.GetMode()
	vnicfcinterruptsettings["object_type"] = item.GetObjectType()

	vnicfcinterruptsettingss = append(vnicfcinterruptsettingss, vnicfcinterruptsettings)
	return vnicfcinterruptsettingss
}
func flattenMapVnicFcNetworkPolicyRelationship(p models.VnicFcNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcnetworkpolicyrelationships []map[string]interface{}
	var ret models.VnicFcNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcnetworkpolicyrelationship := make(map[string]interface{})
	vnicfcnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcnetworkpolicyrelationship["class_id"] = item.GetClassId()
	vnicfcnetworkpolicyrelationship["moid"] = item.GetMoid()
	vnicfcnetworkpolicyrelationship["object_type"] = item.GetObjectType()
	vnicfcnetworkpolicyrelationship["selector"] = item.GetSelector()

	vnicfcnetworkpolicyrelationships = append(vnicfcnetworkpolicyrelationships, vnicfcnetworkpolicyrelationship)
	return vnicfcnetworkpolicyrelationships
}
func flattenMapVnicFcQosPolicyRelationship(p models.VnicFcQosPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqospolicyrelationships []map[string]interface{}
	var ret models.VnicFcQosPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcqospolicyrelationship := make(map[string]interface{})
	vnicfcqospolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcqospolicyrelationship["class_id"] = item.GetClassId()
	vnicfcqospolicyrelationship["moid"] = item.GetMoid()
	vnicfcqospolicyrelationship["object_type"] = item.GetObjectType()
	vnicfcqospolicyrelationship["selector"] = item.GetSelector()

	vnicfcqospolicyrelationships = append(vnicfcqospolicyrelationships, vnicfcqospolicyrelationship)
	return vnicfcqospolicyrelationships
}
func flattenMapVnicFcQueueSettings(p models.VnicFcQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqueuesettingss []map[string]interface{}
	var ret models.VnicFcQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicfcqueuesettings := make(map[string]interface{})
	vnicfcqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcqueuesettings["class_id"] = item.GetClassId()
	vnicfcqueuesettings["nr_count"] = item.GetCount()
	vnicfcqueuesettings["object_type"] = item.GetObjectType()
	vnicfcqueuesettings["ring_size"] = item.GetRingSize()

	vnicfcqueuesettingss = append(vnicfcqueuesettingss, vnicfcqueuesettings)
	return vnicfcqueuesettingss
}
func flattenMapVnicFlogiSettings(p models.VnicFlogiSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicflogisettingss []map[string]interface{}
	var ret models.VnicFlogiSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicflogisettings := make(map[string]interface{})
	vnicflogisettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicflogisettings["class_id"] = item.GetClassId()
	vnicflogisettings["object_type"] = item.GetObjectType()
	vnicflogisettings["retries"] = item.GetRetries()
	vnicflogisettings["timeout"] = item.GetTimeout()

	vnicflogisettingss = append(vnicflogisettingss, vnicflogisettings)
	return vnicflogisettingss
}
func flattenMapVnicIscsiAdapterPolicyRelationship(p models.VnicIscsiAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vniciscsiadapterpolicyrelationships []map[string]interface{}
	var ret models.VnicIscsiAdapterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vniciscsiadapterpolicyrelationship := make(map[string]interface{})
	vniciscsiadapterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniciscsiadapterpolicyrelationship["class_id"] = item.GetClassId()
	vniciscsiadapterpolicyrelationship["moid"] = item.GetMoid()
	vniciscsiadapterpolicyrelationship["object_type"] = item.GetObjectType()
	vniciscsiadapterpolicyrelationship["selector"] = item.GetSelector()

	vniciscsiadapterpolicyrelationships = append(vniciscsiadapterpolicyrelationships, vniciscsiadapterpolicyrelationship)
	return vniciscsiadapterpolicyrelationships
}
func flattenMapVnicIscsiAuthProfile(p models.VnicIscsiAuthProfile, d *schema.ResourceData) []map[string]interface{} {
	var vniciscsiauthprofiles []map[string]interface{}
	var ret models.VnicIscsiAuthProfile
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vniciscsiauthprofile := make(map[string]interface{})
	vniciscsiauthprofile["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniciscsiauthprofile["class_id"] = item.GetClassId()
	vniciscsiauthprofile["is_password_set"] = item.GetIsPasswordSet()
	vniciscsiauthprofile["object_type"] = item.GetObjectType()
	password_x, exists := d.GetOk("chap")
	if exists && password_x != nil {
		password_y := password_x.([]interface{})[0].(map[string]interface{})
		vniciscsiauthprofile["password"] = password_y["password"]
	}
	vniciscsiauthprofile["user_id"] = item.GetUserId()

	vniciscsiauthprofiles = append(vniciscsiauthprofiles, vniciscsiauthprofile)
	return vniciscsiauthprofiles
}
func flattenMapVnicIscsiBootPolicyRelationship(p models.VnicIscsiBootPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vniciscsibootpolicyrelationships []map[string]interface{}
	var ret models.VnicIscsiBootPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vniciscsibootpolicyrelationship := make(map[string]interface{})
	vniciscsibootpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniciscsibootpolicyrelationship["class_id"] = item.GetClassId()
	vniciscsibootpolicyrelationship["moid"] = item.GetMoid()
	vniciscsibootpolicyrelationship["object_type"] = item.GetObjectType()
	vniciscsibootpolicyrelationship["selector"] = item.GetSelector()

	vniciscsibootpolicyrelationships = append(vniciscsibootpolicyrelationships, vniciscsibootpolicyrelationship)
	return vniciscsibootpolicyrelationships
}
func flattenMapVnicIscsiStaticTargetPolicyRelationship(p models.VnicIscsiStaticTargetPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vniciscsistatictargetpolicyrelationships []map[string]interface{}
	var ret models.VnicIscsiStaticTargetPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vniciscsistatictargetpolicyrelationship := make(map[string]interface{})
	vniciscsistatictargetpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniciscsistatictargetpolicyrelationship["class_id"] = item.GetClassId()
	vniciscsistatictargetpolicyrelationship["moid"] = item.GetMoid()
	vniciscsistatictargetpolicyrelationship["object_type"] = item.GetObjectType()
	vniciscsistatictargetpolicyrelationship["selector"] = item.GetSelector()

	vniciscsistatictargetpolicyrelationships = append(vniciscsistatictargetpolicyrelationships, vniciscsistatictargetpolicyrelationship)
	return vniciscsistatictargetpolicyrelationships
}
func flattenMapVnicLanConnectivityPolicyRelationship(p models.VnicLanConnectivityPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vniclanconnectivitypolicyrelationships []map[string]interface{}
	var ret models.VnicLanConnectivityPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vniclanconnectivitypolicyrelationship := make(map[string]interface{})
	vniclanconnectivitypolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniclanconnectivitypolicyrelationship["class_id"] = item.GetClassId()
	vniclanconnectivitypolicyrelationship["moid"] = item.GetMoid()
	vniclanconnectivitypolicyrelationship["object_type"] = item.GetObjectType()
	vniclanconnectivitypolicyrelationship["selector"] = item.GetSelector()

	vniclanconnectivitypolicyrelationships = append(vniclanconnectivitypolicyrelationships, vniclanconnectivitypolicyrelationship)
	return vniclanconnectivitypolicyrelationships
}
func flattenMapVnicLun(p models.VnicLun, d *schema.ResourceData) []map[string]interface{} {
	var vnicluns []map[string]interface{}
	var ret models.VnicLun
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vniclun := make(map[string]interface{})
	vniclun["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniclun["bootable"] = item.GetBootable()
	vniclun["class_id"] = item.GetClassId()
	vniclun["lun_id"] = item.GetLunId()
	vniclun["object_type"] = item.GetObjectType()

	vnicluns = append(vnicluns, vniclun)
	return vnicluns
}
func flattenMapVnicNvgreSettings(p models.VnicNvgreSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicnvgresettingss []map[string]interface{}
	var ret models.VnicNvgreSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicnvgresettings := make(map[string]interface{})
	vnicnvgresettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicnvgresettings["class_id"] = item.GetClassId()
	vnicnvgresettings["enabled"] = item.GetEnabled()
	vnicnvgresettings["object_type"] = item.GetObjectType()

	vnicnvgresettingss = append(vnicnvgresettingss, vnicnvgresettings)
	return vnicnvgresettingss
}
func flattenMapVnicPlacementSettings(p models.VnicPlacementSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicplacementsettingss []map[string]interface{}
	var ret models.VnicPlacementSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicplacementsettings := make(map[string]interface{})
	vnicplacementsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicplacementsettings["class_id"] = item.GetClassId()
	vnicplacementsettings["id"] = item.GetId()
	vnicplacementsettings["object_type"] = item.GetObjectType()
	vnicplacementsettings["pci_link"] = item.GetPciLink()
	vnicplacementsettings["switch_id"] = item.GetSwitchId()
	vnicplacementsettings["uplink"] = item.GetUplink()

	vnicplacementsettingss = append(vnicplacementsettingss, vnicplacementsettings)
	return vnicplacementsettingss
}
func flattenMapVnicPlogiSettings(p models.VnicPlogiSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicplogisettingss []map[string]interface{}
	var ret models.VnicPlogiSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicplogisettings := make(map[string]interface{})
	vnicplogisettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicplogisettings["class_id"] = item.GetClassId()
	vnicplogisettings["object_type"] = item.GetObjectType()
	vnicplogisettings["retries"] = item.GetRetries()
	vnicplogisettings["timeout"] = item.GetTimeout()

	vnicplogisettingss = append(vnicplogisettingss, vnicplogisettings)
	return vnicplogisettingss
}
func flattenMapVnicRoceSettings(p models.VnicRoceSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicrocesettingss []map[string]interface{}
	var ret models.VnicRoceSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicrocesettings := make(map[string]interface{})
	vnicrocesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicrocesettings["class_id"] = item.GetClassId()
	vnicrocesettings["class_of_service"] = item.GetClassOfService()
	vnicrocesettings["enabled"] = item.GetEnabled()
	vnicrocesettings["memory_regions"] = item.GetMemoryRegions()
	vnicrocesettings["object_type"] = item.GetObjectType()
	vnicrocesettings["queue_pairs"] = item.GetQueuePairs()
	vnicrocesettings["resource_groups"] = item.GetResourceGroups()
	vnicrocesettings["nr_version"] = item.GetVersion()

	vnicrocesettingss = append(vnicrocesettingss, vnicrocesettings)
	return vnicrocesettingss
}
func flattenMapVnicRssHashSettings(p models.VnicRssHashSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicrsshashsettingss []map[string]interface{}
	var ret models.VnicRssHashSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicrsshashsettings := make(map[string]interface{})
	vnicrsshashsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicrsshashsettings["class_id"] = item.GetClassId()
	vnicrsshashsettings["ipv4_hash"] = item.GetIpv4Hash()
	vnicrsshashsettings["ipv6_ext_hash"] = item.GetIpv6ExtHash()
	vnicrsshashsettings["ipv6_hash"] = item.GetIpv6Hash()
	vnicrsshashsettings["object_type"] = item.GetObjectType()
	vnicrsshashsettings["tcp_ipv4_hash"] = item.GetTcpIpv4Hash()
	vnicrsshashsettings["tcp_ipv6_ext_hash"] = item.GetTcpIpv6ExtHash()
	vnicrsshashsettings["tcp_ipv6_hash"] = item.GetTcpIpv6Hash()
	vnicrsshashsettings["udp_ipv4_hash"] = item.GetUdpIpv4Hash()
	vnicrsshashsettings["udp_ipv6_hash"] = item.GetUdpIpv6Hash()

	vnicrsshashsettingss = append(vnicrsshashsettingss, vnicrsshashsettings)
	return vnicrsshashsettingss
}
func flattenMapVnicSanConnectivityPolicyRelationship(p models.VnicSanConnectivityPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicsanconnectivitypolicyrelationships []map[string]interface{}
	var ret models.VnicSanConnectivityPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicsanconnectivitypolicyrelationship := make(map[string]interface{})
	vnicsanconnectivitypolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicsanconnectivitypolicyrelationship["class_id"] = item.GetClassId()
	vnicsanconnectivitypolicyrelationship["moid"] = item.GetMoid()
	vnicsanconnectivitypolicyrelationship["object_type"] = item.GetObjectType()
	vnicsanconnectivitypolicyrelationship["selector"] = item.GetSelector()

	vnicsanconnectivitypolicyrelationships = append(vnicsanconnectivitypolicyrelationships, vnicsanconnectivitypolicyrelationship)
	return vnicsanconnectivitypolicyrelationships
}
func flattenMapVnicScsiQueueSettings(p models.VnicScsiQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicscsiqueuesettingss []map[string]interface{}
	var ret models.VnicScsiQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicscsiqueuesettings := make(map[string]interface{})
	vnicscsiqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicscsiqueuesettings["class_id"] = item.GetClassId()
	vnicscsiqueuesettings["nr_count"] = item.GetCount()
	vnicscsiqueuesettings["object_type"] = item.GetObjectType()
	vnicscsiqueuesettings["ring_size"] = item.GetRingSize()

	vnicscsiqueuesettingss = append(vnicscsiqueuesettingss, vnicscsiqueuesettings)
	return vnicscsiqueuesettingss
}
func flattenMapVnicTcpOffloadSettings(p models.VnicTcpOffloadSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnictcpoffloadsettingss []map[string]interface{}
	var ret models.VnicTcpOffloadSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnictcpoffloadsettings := make(map[string]interface{})
	vnictcpoffloadsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnictcpoffloadsettings["class_id"] = item.GetClassId()
	vnictcpoffloadsettings["large_receive"] = item.GetLargeReceive()
	vnictcpoffloadsettings["large_send"] = item.GetLargeSend()
	vnictcpoffloadsettings["object_type"] = item.GetObjectType()
	vnictcpoffloadsettings["rx_checksum"] = item.GetRxChecksum()
	vnictcpoffloadsettings["tx_checksum"] = item.GetTxChecksum()

	vnictcpoffloadsettingss = append(vnictcpoffloadsettingss, vnictcpoffloadsettings)
	return vnictcpoffloadsettingss
}
func flattenMapVnicUsnicSettings(p models.VnicUsnicSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicusnicsettingss []map[string]interface{}
	var ret models.VnicUsnicSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicusnicsettings := make(map[string]interface{})
	vnicusnicsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicusnicsettings["class_id"] = item.GetClassId()
	vnicusnicsettings["cos"] = item.GetCos()
	vnicusnicsettings["nr_count"] = item.GetCount()
	vnicusnicsettings["object_type"] = item.GetObjectType()
	vnicusnicsettings["usnic_adapter_policy"] = item.GetUsnicAdapterPolicy()

	vnicusnicsettingss = append(vnicusnicsettingss, vnicusnicsettings)
	return vnicusnicsettingss
}
func flattenMapVnicVlanSettings(p models.VnicVlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvlansettingss []map[string]interface{}
	var ret models.VnicVlanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvlansettings := make(map[string]interface{})
	vnicvlansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvlansettings["allowed_vlans"] = item.GetAllowedVlans()
	vnicvlansettings["class_id"] = item.GetClassId()
	vnicvlansettings["default_vlan"] = item.GetDefaultVlan()
	vnicvlansettings["mode"] = item.GetMode()
	vnicvlansettings["object_type"] = item.GetObjectType()

	vnicvlansettingss = append(vnicvlansettingss, vnicvlansettings)
	return vnicvlansettingss
}
func flattenMapVnicVmqSettings(p models.VnicVmqSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvmqsettingss []map[string]interface{}
	var ret models.VnicVmqSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvmqsettings := make(map[string]interface{})
	vnicvmqsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvmqsettings["class_id"] = item.GetClassId()
	vnicvmqsettings["enabled"] = item.GetEnabled()
	vnicvmqsettings["multi_queue_support"] = item.GetMultiQueueSupport()
	vnicvmqsettings["num_interrupts"] = item.GetNumInterrupts()
	vnicvmqsettings["num_sub_vnics"] = item.GetNumSubVnics()
	vnicvmqsettings["num_vmqs"] = item.GetNumVmqs()
	vnicvmqsettings["object_type"] = item.GetObjectType()
	vnicvmqsettings["vmmq_adapter_policy"] = item.GetVmmqAdapterPolicy()

	vnicvmqsettingss = append(vnicvmqsettingss, vnicvmqsettings)
	return vnicvmqsettingss
}
func flattenMapVnicVsanSettings(p models.VnicVsanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvsansettingss []map[string]interface{}
	var ret models.VnicVsanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvsansettings := make(map[string]interface{})
	vnicvsansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvsansettings["class_id"] = item.GetClassId()
	vnicvsansettings["default_vlan_id"] = item.GetDefaultVlanId()
	vnicvsansettings["id"] = item.GetId()
	vnicvsansettings["object_type"] = item.GetObjectType()

	vnicvsansettingss = append(vnicvsansettingss, vnicvsansettings)
	return vnicvsansettingss
}
func flattenMapVnicVxlanSettings(p models.VnicVxlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvxlansettingss []map[string]interface{}
	var ret models.VnicVxlanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvxlansettings := make(map[string]interface{})
	vnicvxlansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvxlansettings["class_id"] = item.GetClassId()
	vnicvxlansettings["enabled"] = item.GetEnabled()
	vnicvxlansettings["object_type"] = item.GetObjectType()

	vnicvxlansettingss = append(vnicvxlansettingss, vnicvxlansettings)
	return vnicvxlansettingss
}
func flattenMapVrfVrfRelationship(p models.VrfVrfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vrfvrfrelationships []map[string]interface{}
	var ret models.VrfVrfRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vrfvrfrelationship := make(map[string]interface{})
	vrfvrfrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vrfvrfrelationship["class_id"] = item.GetClassId()
	vrfvrfrelationship["moid"] = item.GetMoid()
	vrfvrfrelationship["object_type"] = item.GetObjectType()
	vrfvrfrelationship["selector"] = item.GetSelector()

	vrfvrfrelationships = append(vrfvrfrelationships, vrfvrfrelationship)
	return vrfvrfrelationships
}
func flattenMapWorkflowCatalogRelationship(p models.WorkflowCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowcatalogrelationships []map[string]interface{}
	var ret models.WorkflowCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowcatalogrelationship := make(map[string]interface{})
	workflowcatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowcatalogrelationship["class_id"] = item.GetClassId()
	workflowcatalogrelationship["moid"] = item.GetMoid()
	workflowcatalogrelationship["object_type"] = item.GetObjectType()
	workflowcatalogrelationship["selector"] = item.GetSelector()

	workflowcatalogrelationships = append(workflowcatalogrelationships, workflowcatalogrelationship)
	return workflowcatalogrelationships
}
func flattenMapWorkflowComments(p models.WorkflowComments, d *schema.ResourceData) []map[string]interface{} {
	var workflowcommentss []map[string]interface{}
	var ret models.WorkflowComments
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowcomments := make(map[string]interface{})
	workflowcomments["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowcomments["class_id"] = item.GetClassId()
	workflowcomments["description"] = item.GetDescription()
	workflowcomments["examples"] = item.GetExamples()
	workflowcomments["object_type"] = item.GetObjectType()

	workflowcommentss = append(workflowcommentss, workflowcomments)
	return workflowcommentss
}
func flattenMapWorkflowCustomDataTypeProperties(p models.WorkflowCustomDataTypeProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowcustomdatatypepropertiess []map[string]interface{}
	var ret models.WorkflowCustomDataTypeProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowcustomdatatypeproperties := make(map[string]interface{})
	workflowcustomdatatypeproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowcustomdatatypeproperties["class_id"] = item.GetClassId()
	workflowcustomdatatypeproperties["external_meta"] = item.GetExternalMeta()
	workflowcustomdatatypeproperties["object_type"] = item.GetObjectType()

	workflowcustomdatatypepropertiess = append(workflowcustomdatatypepropertiess, workflowcustomdatatypeproperties)
	return workflowcustomdatatypepropertiess
}
func flattenMapWorkflowErrorResponseHandlerRelationship(p models.WorkflowErrorResponseHandlerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowerrorresponsehandlerrelationships []map[string]interface{}
	var ret models.WorkflowErrorResponseHandlerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowerrorresponsehandlerrelationship := make(map[string]interface{})
	workflowerrorresponsehandlerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowerrorresponsehandlerrelationship["class_id"] = item.GetClassId()
	workflowerrorresponsehandlerrelationship["moid"] = item.GetMoid()
	workflowerrorresponsehandlerrelationship["object_type"] = item.GetObjectType()
	workflowerrorresponsehandlerrelationship["selector"] = item.GetSelector()

	workflowerrorresponsehandlerrelationships = append(workflowerrorresponsehandlerrelationships, workflowerrorresponsehandlerrelationship)
	return workflowerrorresponsehandlerrelationships
}
func flattenMapWorkflowInternalProperties(p models.WorkflowInternalProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowinternalpropertiess []map[string]interface{}
	var ret models.WorkflowInternalProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowinternalproperties := make(map[string]interface{})
	workflowinternalproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowinternalproperties["base_task_type"] = item.GetBaseTaskType()
	workflowinternalproperties["class_id"] = item.GetClassId()
	workflowinternalproperties["constraints"] = (func(p models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
		var workflowtaskconstraintss []map[string]interface{}
		var ret models.WorkflowTaskConstraints
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		workflowtaskconstraints := make(map[string]interface{})
		workflowtaskconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowtaskconstraints["class_id"] = item.GetClassId()
		workflowtaskconstraints["object_type"] = item.GetObjectType()
		workflowtaskconstraints["target_data_type"] = flattenAdditionalProperties(item.TargetDataType)

		workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
		return workflowtaskconstraintss
	})(item.GetConstraints(), d)
	workflowinternalproperties["internal"] = item.GetInternal()
	workflowinternalproperties["object_type"] = item.GetObjectType()
	workflowinternalproperties["owner"] = item.GetOwner()

	workflowinternalpropertiess = append(workflowinternalpropertiess, workflowinternalproperties)
	return workflowinternalpropertiess
}
func flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p models.WorkflowPendingDynamicWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowpendingdynamicworkflowinforelationships []map[string]interface{}
	var ret models.WorkflowPendingDynamicWorkflowInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowpendingdynamicworkflowinforelationship := make(map[string]interface{})
	workflowpendingdynamicworkflowinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowpendingdynamicworkflowinforelationship["class_id"] = item.GetClassId()
	workflowpendingdynamicworkflowinforelationship["moid"] = item.GetMoid()
	workflowpendingdynamicworkflowinforelationship["object_type"] = item.GetObjectType()
	workflowpendingdynamicworkflowinforelationship["selector"] = item.GetSelector()

	workflowpendingdynamicworkflowinforelationships = append(workflowpendingdynamicworkflowinforelationships, workflowpendingdynamicworkflowinforelationship)
	return workflowpendingdynamicworkflowinforelationships
}
func flattenMapWorkflowProperties(p models.WorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowpropertiess []map[string]interface{}
	var ret models.WorkflowProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowproperties := make(map[string]interface{})
	workflowproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowproperties["class_id"] = item.GetClassId()
	workflowproperties["external_meta"] = item.GetExternalMeta()
	workflowproperties["input_definition"] = (func(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowbasedatatype["class_id"] = item.GetClassId()
			workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.GetClassId()
				workflowdefaultvalue["is_value_set"] = item.GetIsValueSet()
				workflowdefaultvalue["object_type"] = item.GetObjectType()
				workflowdefaultvalue["override"] = item.GetOverride()
				workflowdefaultvalue["value"] = flattenAdditionalProperties(item.Value)

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowbasedatatype["description"] = item.GetDescription()
			workflowbasedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
				var workflowdisplaymetas []map[string]interface{}
				var ret models.WorkflowDisplayMeta
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdisplaymeta := make(map[string]interface{})
				workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdisplaymeta["class_id"] = item.GetClassId()
				workflowdisplaymeta["inventory_selector"] = item.GetInventorySelector()
				workflowdisplaymeta["object_type"] = item.GetObjectType()
				workflowdisplaymeta["widget_type"] = item.GetWidgetType()

				workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
				return workflowdisplaymetas
			})(item.GetDisplayMeta(), d)
			workflowbasedatatype["input_parameters"] = flattenAdditionalProperties(item.InputParameters)
			workflowbasedatatype["label"] = item.GetLabel()
			workflowbasedatatype["name"] = item.GetName()
			workflowbasedatatype["object_type"] = item.GetObjectType()
			workflowbasedatatype["required"] = item.GetRequired()
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.GetInputDefinition(), d)
	workflowproperties["object_type"] = item.GetObjectType()
	workflowproperties["output_definition"] = (func(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowbasedatatype["class_id"] = item.GetClassId()
			workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.GetClassId()
				workflowdefaultvalue["is_value_set"] = item.GetIsValueSet()
				workflowdefaultvalue["object_type"] = item.GetObjectType()
				workflowdefaultvalue["override"] = item.GetOverride()
				workflowdefaultvalue["value"] = flattenAdditionalProperties(item.Value)

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowbasedatatype["description"] = item.GetDescription()
			workflowbasedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
				var workflowdisplaymetas []map[string]interface{}
				var ret models.WorkflowDisplayMeta
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdisplaymeta := make(map[string]interface{})
				workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdisplaymeta["class_id"] = item.GetClassId()
				workflowdisplaymeta["inventory_selector"] = item.GetInventorySelector()
				workflowdisplaymeta["object_type"] = item.GetObjectType()
				workflowdisplaymeta["widget_type"] = item.GetWidgetType()

				workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
				return workflowdisplaymetas
			})(item.GetDisplayMeta(), d)
			workflowbasedatatype["input_parameters"] = flattenAdditionalProperties(item.InputParameters)
			workflowbasedatatype["label"] = item.GetLabel()
			workflowbasedatatype["name"] = item.GetName()
			workflowbasedatatype["object_type"] = item.GetObjectType()
			workflowbasedatatype["required"] = item.GetRequired()
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.GetOutputDefinition(), d)
	workflowproperties["retry_count"] = item.GetRetryCount()
	workflowproperties["retry_delay"] = item.GetRetryDelay()
	workflowproperties["retry_policy"] = item.GetRetryPolicy()
	workflowproperties["support_status"] = item.GetSupportStatus()
	workflowproperties["timeout"] = item.GetTimeout()
	workflowproperties["timeout_policy"] = item.GetTimeoutPolicy()

	workflowpropertiess = append(workflowpropertiess, workflowproperties)
	return workflowpropertiess
}
func flattenMapWorkflowTaskConstraints(p models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskconstraintss []map[string]interface{}
	var ret models.WorkflowTaskConstraints
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowtaskconstraints := make(map[string]interface{})
	workflowtaskconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskconstraints["class_id"] = item.GetClassId()
	workflowtaskconstraints["object_type"] = item.GetObjectType()
	workflowtaskconstraints["target_data_type"] = flattenAdditionalProperties(item.TargetDataType)

	workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
	return workflowtaskconstraintss
}
func flattenMapWorkflowTaskDefinitionRelationship(p models.WorkflowTaskDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrelationships []map[string]interface{}
	var ret models.WorkflowTaskDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowtaskdefinitionrelationship := make(map[string]interface{})
	workflowtaskdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskdefinitionrelationship["class_id"] = item.GetClassId()
	workflowtaskdefinitionrelationship["moid"] = item.GetMoid()
	workflowtaskdefinitionrelationship["object_type"] = item.GetObjectType()
	workflowtaskdefinitionrelationship["selector"] = item.GetSelector()

	workflowtaskdefinitionrelationships = append(workflowtaskdefinitionrelationships, workflowtaskdefinitionrelationship)
	return workflowtaskdefinitionrelationships
}
func flattenMapWorkflowTaskInfoRelationship(p models.WorkflowTaskInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforelationships []map[string]interface{}
	var ret models.WorkflowTaskInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowtaskinforelationship := make(map[string]interface{})
	workflowtaskinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskinforelationship["class_id"] = item.GetClassId()
	workflowtaskinforelationship["moid"] = item.GetMoid()
	workflowtaskinforelationship["object_type"] = item.GetObjectType()
	workflowtaskinforelationship["selector"] = item.GetSelector()

	workflowtaskinforelationships = append(workflowtaskinforelationships, workflowtaskinforelationship)
	return workflowtaskinforelationships
}
func flattenMapWorkflowTaskMetadataRelationship(p models.WorkflowTaskMetadataRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskmetadatarelationships []map[string]interface{}
	var ret models.WorkflowTaskMetadataRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowtaskmetadatarelationship := make(map[string]interface{})
	workflowtaskmetadatarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskmetadatarelationship["class_id"] = item.GetClassId()
	workflowtaskmetadatarelationship["moid"] = item.GetMoid()
	workflowtaskmetadatarelationship["object_type"] = item.GetObjectType()
	workflowtaskmetadatarelationship["selector"] = item.GetSelector()

	workflowtaskmetadatarelationships = append(workflowtaskmetadatarelationships, workflowtaskmetadatarelationship)
	return workflowtaskmetadatarelationships
}
func flattenMapWorkflowValidationInformation(p models.WorkflowValidationInformation, d *schema.ResourceData) []map[string]interface{} {
	var workflowvalidationinformations []map[string]interface{}
	var ret models.WorkflowValidationInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowvalidationinformation := make(map[string]interface{})
	workflowvalidationinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowvalidationinformation["class_id"] = item.GetClassId()
	workflowvalidationinformation["object_type"] = item.GetObjectType()
	workflowvalidationinformation["state"] = item.GetState()
	workflowvalidationinformation["validation_error"] = (func(p []models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var workflowvalidationerrors []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowvalidationerror := make(map[string]interface{})
			workflowvalidationerror["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowvalidationerror["class_id"] = item.GetClassId()
			workflowvalidationerror["error_log"] = item.GetErrorLog()
			workflowvalidationerror["field"] = item.GetField()
			workflowvalidationerror["object_type"] = item.GetObjectType()
			workflowvalidationerror["task_name"] = item.GetTaskName()
			workflowvalidationerror["transition_name"] = item.GetTransitionName()
			workflowvalidationerrors = append(workflowvalidationerrors, workflowvalidationerror)
		}
		return workflowvalidationerrors
	})(item.GetValidationError(), d)

	workflowvalidationinformations = append(workflowvalidationinformations, workflowvalidationinformation)
	return workflowvalidationinformations
}
func flattenMapWorkflowWorkflowCtx(p models.WorkflowWorkflowCtx, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowctxs []map[string]interface{}
	var ret models.WorkflowWorkflowCtx
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowworkflowctx := make(map[string]interface{})
	workflowworkflowctx["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowctx["class_id"] = item.GetClassId()
	workflowworkflowctx["initiator_ctx"] = (func(p models.WorkflowInitiatorContext, d *schema.ResourceData) []map[string]interface{} {
		var workflowinitiatorcontexts []map[string]interface{}
		var ret models.WorkflowInitiatorContext
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		workflowinitiatorcontext := make(map[string]interface{})
		workflowinitiatorcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowinitiatorcontext["class_id"] = item.GetClassId()
		workflowinitiatorcontext["initiator_moid"] = item.GetInitiatorMoid()
		workflowinitiatorcontext["initiator_name"] = item.GetInitiatorName()
		workflowinitiatorcontext["initiator_type"] = item.GetInitiatorType()
		workflowinitiatorcontext["object_type"] = item.GetObjectType()

		workflowinitiatorcontexts = append(workflowinitiatorcontexts, workflowinitiatorcontext)
		return workflowinitiatorcontexts
	})(item.GetInitiatorCtx(), d)
	workflowworkflowctx["object_type"] = item.GetObjectType()
	workflowworkflowctx["target_ctx_list"] = (func(p []models.WorkflowTargetContext, d *schema.ResourceData) []map[string]interface{} {
		var workflowtargetcontexts []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowtargetcontext := make(map[string]interface{})
			workflowtargetcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowtargetcontext["class_id"] = item.GetClassId()
			workflowtargetcontext["object_type"] = item.GetObjectType()
			workflowtargetcontext["target_moid"] = item.GetTargetMoid()
			workflowtargetcontext["target_name"] = item.GetTargetName()
			workflowtargetcontext["target_type"] = item.GetTargetType()
			workflowtargetcontexts = append(workflowtargetcontexts, workflowtargetcontext)
		}
		return workflowtargetcontexts
	})(item.GetTargetCtxList(), d)
	workflowworkflowctx["workflow_meta_name"] = item.GetWorkflowMetaName()
	workflowworkflowctx["workflow_subtype"] = item.GetWorkflowSubtype()
	workflowworkflowctx["workflow_type"] = item.GetWorkflowType()

	workflowworkflowctxs = append(workflowworkflowctxs, workflowworkflowctx)
	return workflowworkflowctxs
}
func flattenMapWorkflowWorkflowDefinitionRelationship(p models.WorkflowWorkflowDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowdefinitionrelationships []map[string]interface{}
	var ret models.WorkflowWorkflowDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowworkflowdefinitionrelationship := make(map[string]interface{})
	workflowworkflowdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowdefinitionrelationship["class_id"] = item.GetClassId()
	workflowworkflowdefinitionrelationship["moid"] = item.GetMoid()
	workflowworkflowdefinitionrelationship["object_type"] = item.GetObjectType()
	workflowworkflowdefinitionrelationship["selector"] = item.GetSelector()

	workflowworkflowdefinitionrelationships = append(workflowworkflowdefinitionrelationships, workflowworkflowdefinitionrelationship)
	return workflowworkflowdefinitionrelationships
}
func flattenMapWorkflowWorkflowInfoRelationship(p models.WorkflowWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforelationships []map[string]interface{}
	var ret models.WorkflowWorkflowInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowworkflowinforelationship := make(map[string]interface{})
	workflowworkflowinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowinforelationship["class_id"] = item.GetClassId()
	workflowworkflowinforelationship["moid"] = item.GetMoid()
	workflowworkflowinforelationship["object_type"] = item.GetObjectType()
	workflowworkflowinforelationship["selector"] = item.GetSelector()

	workflowworkflowinforelationships = append(workflowworkflowinforelationships, workflowworkflowinforelationship)
	return workflowworkflowinforelationships
}
func flattenMapWorkflowWorkflowInfoProperties(p models.WorkflowWorkflowInfoProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinfopropertiess []map[string]interface{}
	var ret models.WorkflowWorkflowInfoProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowworkflowinfoproperties := make(map[string]interface{})
	workflowworkflowinfoproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowinfoproperties["class_id"] = item.GetClassId()
	workflowworkflowinfoproperties["object_type"] = item.GetObjectType()
	workflowworkflowinfoproperties["retryable"] = item.GetRetryable()
	workflowworkflowinfoproperties["rollback_action"] = item.GetRollbackAction()

	workflowworkflowinfopropertiess = append(workflowworkflowinfopropertiess, workflowworkflowinfoproperties)
	return workflowworkflowinfopropertiess
}
func flattenMapWorkflowWorkflowMetadataRelationship(p models.WorkflowWorkflowMetadataRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowmetadatarelationships []map[string]interface{}
	var ret models.WorkflowWorkflowMetadataRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowworkflowmetadatarelationship := make(map[string]interface{})
	workflowworkflowmetadatarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowmetadatarelationship["class_id"] = item.GetClassId()
	workflowworkflowmetadatarelationship["moid"] = item.GetMoid()
	workflowworkflowmetadatarelationship["object_type"] = item.GetObjectType()
	workflowworkflowmetadatarelationship["selector"] = item.GetSelector()

	workflowworkflowmetadatarelationships = append(workflowworkflowmetadatarelationships, workflowworkflowmetadatarelationship)
	return workflowworkflowmetadatarelationships
}
func flattenMapWorkflowWorkflowProperties(p models.WorkflowWorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowpropertiess []map[string]interface{}
	var ret models.WorkflowWorkflowProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowworkflowproperties := make(map[string]interface{})
	workflowworkflowproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowproperties["class_id"] = item.GetClassId()
	workflowworkflowproperties["cloneable"] = item.GetCloneable()
	workflowworkflowproperties["enable_debug"] = item.GetEnableDebug()
	workflowworkflowproperties["external_meta"] = item.GetExternalMeta()
	workflowworkflowproperties["object_type"] = item.GetObjectType()
	workflowworkflowproperties["retryable"] = item.GetRetryable()
	workflowworkflowproperties["support_status"] = item.GetSupportStatus()

	workflowworkflowpropertiess = append(workflowworkflowpropertiess, workflowworkflowproperties)
	return workflowworkflowpropertiess
}
func flattenMapX509Certificate(p models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	var ret models.X509Certificate
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	x509certificate := make(map[string]interface{})
	x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	x509certificate["class_id"] = item.GetClassId()
	x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		var ret models.PkixDistinguishedName
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		pkixdistinguishedname["class_id"] = item.GetClassId()
		pkixdistinguishedname["common_name"] = item.GetCommonName()
		pkixdistinguishedname["country"] = item.GetCountry()
		pkixdistinguishedname["locality"] = item.GetLocality()
		pkixdistinguishedname["object_type"] = item.GetObjectType()
		pkixdistinguishedname["organization"] = item.GetOrganization()
		pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
		pkixdistinguishedname["state"] = item.GetState()

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.GetIssuer(), d)
	x509certificate["not_after"] = item.GetNotAfter()
	x509certificate["not_before"] = item.GetNotBefore()
	x509certificate["object_type"] = item.GetObjectType()
	x509certificate["pem_certificate"] = item.GetPemCertificate()
	x509certificate["sha256_fingerprint"] = item.GetSha256Fingerprint()
	x509certificate["signature_algorithm"] = item.GetSignatureAlgorithm()
	x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		var ret models.PkixDistinguishedName
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		pkixdistinguishedname["class_id"] = item.GetClassId()
		pkixdistinguishedname["common_name"] = item.GetCommonName()
		pkixdistinguishedname["country"] = item.GetCountry()
		pkixdistinguishedname["locality"] = item.GetLocality()
		pkixdistinguishedname["object_type"] = item.GetObjectType()
		pkixdistinguishedname["organization"] = item.GetOrganization()
		pkixdistinguishedname["organizational_unit"] = item.GetOrganizationalUnit()
		pkixdistinguishedname["state"] = item.GetState()

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.GetSubject(), d)

	x509certificates = append(x509certificates, x509certificate)
	return x509certificates
}
