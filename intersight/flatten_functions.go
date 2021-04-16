package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenAdditionalProperties(m map[string]interface{}) string {
	var s string
	if len(m) > 0 {
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
	x["class_id"] = ref.ClassId
	x["moid"] = ref.Moid
	x["object_type"] = ref.ObjectType
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
		adapteradapterconfig["class_id"] = item.ClassId
		adapteradapterconfig["dce_interface_settings"] = (func(p []models.AdapterDceInterfaceSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterdceinterfacesettingss []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				adapterdceinterfacesettings := make(map[string]interface{})
				adapterdceinterfacesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				adapterdceinterfacesettings["class_id"] = item.ClassId
				adapterdceinterfacesettings["fec_mode"] = item.FecMode
				adapterdceinterfacesettings["interface_id"] = item.InterfaceId
				adapterdceinterfacesettings["object_type"] = item.ObjectType
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
			adapterethsettings["class_id"] = item.ClassId
			adapterethsettings["lldp_enabled"] = item.LldpEnabled
			adapterethsettings["object_type"] = item.ObjectType

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
			adapterfcsettings["class_id"] = item.ClassId
			adapterfcsettings["fip_enabled"] = item.FipEnabled
			adapterfcsettings["object_type"] = item.ObjectType

			adapterfcsettingss = append(adapterfcsettingss, adapterfcsettings)
			return adapterfcsettingss
		})(item.GetFcSettings(), d)
		adapteradapterconfig["object_type"] = item.ObjectType
		adapteradapterconfig["port_channel_settings"] = (func(p models.AdapterPortChannelSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterportchannelsettingss []map[string]interface{}
			var ret models.AdapterPortChannelSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterportchannelsettings := make(map[string]interface{})
			adapterportchannelsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterportchannelsettings["class_id"] = item.ClassId
			adapterportchannelsettings["enabled"] = item.Enabled
			adapterportchannelsettings["object_type"] = item.ObjectType

			adapterportchannelsettingss = append(adapterportchannelsettingss, adapterportchannelsettings)
			return adapterportchannelsettingss
		})(item.GetPortChannelSettings(), d)
		adapteradapterconfig["slot_id"] = item.SlotId
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
		applianceapistatus["class_id"] = item.ClassId
		applianceapistatus["elapsed_time"] = item.ElapsedTime
		applianceapistatus["object_name"] = item.ObjectName
		applianceapistatus["object_type"] = item.ObjectType
		applianceapistatus["reason"] = item.Reason
		applianceapistatus["response"] = item.Response
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
		appliancecertrenewalphase["class_id"] = item.ClassId
		appliancecertrenewalphase["end_time"] = item.EndTime
		appliancecertrenewalphase["failed"] = item.Failed
		appliancecertrenewalphase["message"] = item.Message
		appliancecertrenewalphase["name"] = item.Name
		appliancecertrenewalphase["object_type"] = item.ObjectType
		appliancecertrenewalphase["start_time"] = item.StartTime
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
		appliancekeyvaluepair["class_id"] = item.ClassId
		appliancekeyvaluepair["key"] = item.Key
		appliancekeyvaluepair["object_type"] = item.ObjectType
		appliancekeyvaluepair["value"] = item.Value
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
		appliancestatuscheck["class_id"] = item.ClassId
		appliancestatuscheck["code"] = item.Code
		appliancestatuscheck["object_type"] = item.ObjectType
		appliancestatuscheck["result"] = item.Result
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
		assetconnection["class_id"] = item.ClassId
		assetconnection["credential"] = (func(p models.AssetCredential, d *schema.ResourceData) []map[string]interface{} {
			var assetcredentials []map[string]interface{}
			var ret models.AssetCredential
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			assetcredential := make(map[string]interface{})
			assetcredential["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetcredential["class_id"] = item.ClassId
			assetcredential["object_type"] = item.ObjectType

			assetcredentials = append(assetcredentials, assetcredential)
			return assetcredentials
		})(item.GetCredential(), d)
		assetconnection["object_type"] = item.ObjectType
		assetconnections = append(assetconnections, assetconnection)
	}
	return assetconnections
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
func flattenListAssetService(p []models.AssetService, d *schema.ResourceData) []map[string]interface{} {
	var assetservices []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		assetservice := make(map[string]interface{})
		assetservice["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetservice["class_id"] = item.ClassId
		assetservice["object_type"] = item.ObjectType
		assetservice["options"] = (func(p models.AssetServiceOptions, d *schema.ResourceData) []map[string]interface{} {
			var assetserviceoptionss []map[string]interface{}
			var ret models.AssetServiceOptions
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			assetserviceoptions := make(map[string]interface{})
			assetserviceoptions["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetserviceoptions["class_id"] = item.ClassId
			assetserviceoptions["object_type"] = item.ObjectType

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
		bootdevicebase["class_id"] = item.ClassId
		bootdevicebase["enabled"] = item.Enabled
		bootdevicebase["name"] = item.Name
		bootdevicebase["object_type"] = item.ObjectType
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
		capabilityportrange["class_id"] = item.ClassId
		capabilityportrange["end_port_id"] = item.EndPortId
		capabilityportrange["end_slot_id"] = item.EndSlotId
		capabilityportrange["object_type"] = item.ObjectType
		capabilityportrange["start_port_id"] = item.StartPortId
		capabilityportrange["start_slot_id"] = item.StartSlotId
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
		capabilityswitchingmodecapability["class_id"] = item.ClassId
		capabilityswitchingmodecapability["object_type"] = item.ObjectType
		capabilityswitchingmodecapability["switching_mode"] = item.SwitchingMode
		capabilityswitchingmodecapability["vp_compression_supported"] = item.VpCompressionSupported
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
			x509certificate["class_id"] = item.ClassId
			x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
				var pkixdistinguishednames []map[string]interface{}
				var ret models.PkixDistinguishedName
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				pkixdistinguishedname := make(map[string]interface{})
				pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				pkixdistinguishedname["class_id"] = item.ClassId
				pkixdistinguishedname["common_name"] = item.CommonName
				pkixdistinguishedname["country"] = item.Country
				pkixdistinguishedname["locality"] = item.Locality
				pkixdistinguishedname["object_type"] = item.ObjectType
				pkixdistinguishedname["organization"] = item.Organization
				pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
				pkixdistinguishedname["state"] = item.State

				pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
				return pkixdistinguishednames
			})(item.GetIssuer(), d)
			x509certificate["not_after"] = item.NotAfter
			x509certificate["not_before"] = item.NotBefore
			x509certificate["object_type"] = item.ObjectType
			x509certificate["pem_certificate"] = item.PemCertificate
			x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
			x509certificate["signature_algorithm"] = item.SignatureAlgorithm
			x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
				var pkixdistinguishednames []map[string]interface{}
				var ret models.PkixDistinguishedName
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				pkixdistinguishedname := make(map[string]interface{})
				pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				pkixdistinguishedname["class_id"] = item.ClassId
				pkixdistinguishedname["common_name"] = item.CommonName
				pkixdistinguishedname["country"] = item.Country
				pkixdistinguishedname["locality"] = item.Locality
				pkixdistinguishedname["object_type"] = item.ObjectType
				pkixdistinguishedname["organization"] = item.Organization
				pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
				pkixdistinguishedname["state"] = item.State

				pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
				return pkixdistinguishednames
			})(item.GetSubject(), d)

			x509certificates = append(x509certificates, x509certificate)
			return x509certificates
		})(item.GetCertificate(), d)
		certificatemanagementcertificatebase["class_id"] = item.ClassId
		certificatemanagementcertificatebase["enabled"] = item.Enabled
		certificatemanagementcertificatebase["is_privatekey_set"] = item.IsPrivatekeySet
		certificatemanagementcertificatebase["object_type"] = item.ObjectType
		privatekey_x := d.Get("certificates").([]interface{})
		if len(privatekey_x) > 0 {
			privatekey_y := privatekey_x[0].(map[string]interface{})
			certificatemanagementcertificatebase["privatekey"] = privatekey_y["privatekey"]
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
		cloudcustomattributes["attribute_name"] = item.AttributeName
		cloudcustomattributes["attribute_type"] = item.AttributeType
		cloudcustomattributes["attribute_value"] = item.AttributeValue
		cloudcustomattributes["class_id"] = item.ClassId
		cloudcustomattributes["object_type"] = item.ObjectType
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
		computeipaddress["address"] = item.Address
		computeipaddress["category"] = item.Category
		computeipaddress["class_id"] = item.ClassId
		computeipaddress["default_gateway"] = item.DefaultGateway
		computeipaddress["dn"] = item.Dn
		computeipaddress["http_port"] = item.HttpPort
		computeipaddress["https_port"] = item.HttpsPort
		computeipaddress["kvm_port"] = item.KvmPort
		computeipaddress["kvm_vlan"] = item.KvmVlan
		computeipaddress["name"] = item.Name
		computeipaddress["object_type"] = item.ObjectType
		computeipaddress["subnet"] = item.Subnet
		computeipaddress["type"] = item.Type
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
		configmoref["class_id"] = item.ClassId
		configmoref["moid"] = item.Moid
		configmoref["object_type"] = item.ObjectType
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
		connectorpackconnectorpackupdate["class_id"] = item.ClassId
		connectorpackconnectorpackupdate["current_version"] = item.CurrentVersion
		connectorpackconnectorpackupdate["name"] = item.Name
		connectorpackconnectorpackupdate["new_version"] = item.NewVersion
		connectorpackconnectorpackupdate["object_type"] = item.ObjectType
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
		contentcomplextype["class_id"] = item.ClassId
		contentcomplextype["name"] = item.Name
		contentcomplextype["object_type"] = item.ObjectType
		contentcomplextype["parameters"] = (func(p []models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
			var contentbaseparameters []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				contentbaseparameter := make(map[string]interface{})
				contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
				contentbaseparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				contentbaseparameter["class_id"] = item.ClassId
				contentbaseparameter["complex_type"] = item.ComplexType
				contentbaseparameter["item_type"] = item.ItemType
				contentbaseparameter["name"] = item.Name
				contentbaseparameter["object_type"] = item.ObjectType
				contentbaseparameter["path"] = item.Path
				contentbaseparameter["type"] = item.Type
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
		contentparameter["accept_single_value"] = item.AcceptSingleValue
		contentparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		contentparameter["class_id"] = item.ClassId
		contentparameter["complex_type"] = item.ComplexType
		contentparameter["item_type"] = item.ItemType
		contentparameter["name"] = item.Name
		contentparameter["object_type"] = item.ObjectType
		contentparameter["path"] = item.Path
		contentparameter["type"] = item.Type
		contentparameters = append(contentparameters, contentparameter)
	}
	return contentparameters
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
		equipmentiocardidentity["class_id"] = item.ClassId
		equipmentiocardidentity["io_card_moid"] = item.IoCardMoid
		equipmentiocardidentity["module_id"] = item.ModuleId
		equipmentiocardidentity["network_element_moid"] = item.NetworkElementMoid
		equipmentiocardidentity["object_type"] = item.ObjectType
		equipmentiocardidentity["switch_id"] = item.SwitchId
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
		fabricportidentifier["aggregate_port_id"] = item.AggregatePortId
		fabricportidentifier["class_id"] = item.ClassId
		fabricportidentifier["object_type"] = item.ObjectType
		fabricportidentifier["port_id"] = item.PortId
		fabricportidentifier["slot_id"] = item.SlotId
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
		fabricqosclass["admin_state"] = item.AdminState
		fabricqosclass["bandwidth_percent"] = item.BandwidthPercent
		fabricqosclass["class_id"] = item.ClassId
		fabricqosclass["cos"] = item.Cos
		fabricqosclass["mtu"] = item.Mtu
		fabricqosclass["multicast_optimize"] = item.MulticastOptimize
		fabricqosclass["name"] = item.Name
		fabricqosclass["object_type"] = item.ObjectType
		fabricqosclass["packet_drop"] = item.PacketDrop
		fabricqosclass["weight"] = item.Weight
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
		fcpoolblock["class_id"] = item.ClassId
		fcpoolblock["from"] = item.From
		fcpoolblock["object_type"] = item.ObjectType
		fcpoolblock["size"] = item.Size
		fcpoolblock["to"] = item.To
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
		firmwarebaseimpact["class_id"] = item.ClassId
		firmwarebaseimpact["computation_error"] = item.ComputationError
		firmwarebaseimpact["computation_status_detail"] = item.ComputationStatusDetail
		firmwarebaseimpact["domain_name"] = item.DomainName
		firmwarebaseimpact["end_point"] = item.EndPoint
		firmwarebaseimpact["firmware_version"] = item.FirmwareVersion
		firmwarebaseimpact["impact_type"] = item.ImpactType
		firmwarebaseimpact["model"] = item.Model
		firmwarebaseimpact["object_type"] = item.ObjectType
		firmwarebaseimpact["target_firmware_version"] = item.TargetFirmwareVersion
		firmwarebaseimpact["version_impact"] = item.VersionImpact
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
		firmwarecomponentmeta["class_id"] = item.ClassId
		firmwarecomponentmeta["component_label"] = item.ComponentLabel
		firmwarecomponentmeta["component_type"] = item.ComponentType
		firmwarecomponentmeta["description"] = item.Description
		firmwarecomponentmeta["disruption"] = item.Disruption
		firmwarecomponentmeta["image_path"] = item.ImagePath
		firmwarecomponentmeta["is_oob_supported"] = item.IsOobSupported
		firmwarecomponentmeta["model"] = item.Model
		firmwarecomponentmeta["object_type"] = item.ObjectType
		firmwarecomponentmeta["oob_manageability"] = item.OobManageability
		firmwarecomponentmeta["packed_version"] = item.PackedVersion
		firmwarecomponentmeta["redfish_url"] = item.RedfishUrl
		firmwarecomponentmeta["vendor"] = item.Vendor
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
		firmwarefirmwareinventory["category"] = item.Category
		firmwarefirmwareinventory["class_id"] = item.ClassId
		firmwarefirmwareinventory["label"] = item.Label
		firmwarefirmwareinventory["model"] = item.Model
		firmwarefirmwareinventory["object_type"] = item.ObjectType
		firmwarefirmwareinventory["update_uri"] = item.UpdateUri
		firmwarefirmwareinventory["vendor"] = item.Vendor
		firmwarefirmwareinventory["nr_version"] = item.Version
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
		hclconstraint["class_id"] = item.ClassId
		hclconstraint["constraint_name"] = item.ConstraintName
		hclconstraint["constraint_value"] = item.ConstraintValue
		hclconstraint["object_type"] = item.ObjectType
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
		hyperflexfeaturelimitentry["class_id"] = item.ClassId
		hyperflexfeaturelimitentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.ClassId
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexfeaturelimitentry["name"] = item.Name
		hyperflexfeaturelimitentry["object_type"] = item.ObjectType
		hyperflexfeaturelimitentry["value"] = item.Value
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
		hyperflexhealthcheckscriptinfo["aggregate_script_name"] = item.AggregateScriptName
		hyperflexhealthcheckscriptinfo["class_id"] = item.ClassId
		hyperflexhealthcheckscriptinfo["hyperflex_version"] = item.HyperflexVersion
		hyperflexhealthcheckscriptinfo["object_type"] = item.ObjectType
		hyperflexhealthcheckscriptinfo["script_execute_location"] = item.ScriptExecuteLocation
		hyperflexhealthcheckscriptinfo["script_name"] = item.ScriptName
		hyperflexhealthcheckscriptinfo["nr_version"] = item.Version
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
		hyperflexhxhostmountstatusdt["accessibility"] = item.Accessibility
		hyperflexhxhostmountstatusdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhxhostmountstatusdt["class_id"] = item.ClassId
		hyperflexhxhostmountstatusdt["host_name"] = item.HostName
		hyperflexhxhostmountstatusdt["mounted"] = item.Mounted
		hyperflexhxhostmountstatusdt["object_type"] = item.ObjectType
		hyperflexhxhostmountstatusdt["reason"] = item.Reason
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
		hyperflexhxzoneresiliencyinfodt["class_id"] = item.ClassId
		hyperflexhxzoneresiliencyinfodt["name"] = item.Name
		hyperflexhxzoneresiliencyinfodt["object_type"] = item.ObjectType
		hyperflexhxzoneresiliencyinfodt["resiliency_info"] = (func(p models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexhxresiliencyinfodts []map[string]interface{}
			var ret models.HyperflexHxResiliencyInfoDt
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexhxresiliencyinfodt := make(map[string]interface{})
			hyperflexhxresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexhxresiliencyinfodt["class_id"] = item.ClassId
			hyperflexhxresiliencyinfodt["data_replication_factor"] = item.DataReplicationFactor
			hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.HddFailuresTolerable
			hyperflexhxresiliencyinfodt["messages"] = item.Messages
			hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.NodeFailuresTolerable
			hyperflexhxresiliencyinfodt["object_type"] = item.ObjectType
			hyperflexhxresiliencyinfodt["policy_compliance"] = item.PolicyCompliance
			hyperflexhxresiliencyinfodt["resiliency_state"] = item.ResiliencyState
			hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.SsdFailuresTolerable

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
		hyperflexipaddrrange["class_id"] = item.ClassId
		hyperflexipaddrrange["end_addr"] = item.EndAddr
		hyperflexipaddrrange["gateway"] = item.Gateway
		hyperflexipaddrrange["netmask"] = item.Netmask
		hyperflexipaddrrange["object_type"] = item.ObjectType
		hyperflexipaddrrange["start_addr"] = item.StartAddr
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
		hyperflexmapclusteridtoprotectioninfo["class_id"] = item.ClassId
		hyperflexmapclusteridtoprotectioninfo["cluster_id"] = item.ClusterId
		hyperflexmapclusteridtoprotectioninfo["object_type"] = item.ObjectType
		hyperflexmapclusteridtoprotectioninfo["protection_info"] = (func(p models.HyperflexProtectionInfo, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexprotectioninfos []map[string]interface{}
			var ret models.HyperflexProtectionInfo
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexprotectioninfo := make(map[string]interface{})
			hyperflexprotectioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexprotectioninfo["class_id"] = item.ClassId
			hyperflexprotectioninfo["object_type"] = item.ObjectType
			hyperflexprotectioninfo["vm_current_protection_info"] = (func(p models.HyperflexSnapshotInfoBrief, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexsnapshotinfobriefs []map[string]interface{}
				var ret models.HyperflexSnapshotInfoBrief
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexsnapshotinfobrief := make(map[string]interface{})
				hyperflexsnapshotinfobrief["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexsnapshotinfobrief["class_id"] = item.ClassId
				hyperflexsnapshotinfobrief["object_type"] = item.ObjectType
				hyperflexsnapshotinfobrief["replication_status"] = (func(p models.HyperflexReplicationStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexreplicationstatuss []map[string]interface{}
					var ret models.HyperflexReplicationStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexreplicationstatus := make(map[string]interface{})
					hyperflexreplicationstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexreplicationstatus["bytes_replicated"] = item.BytesReplicated
					hyperflexreplicationstatus["class_id"] = item.ClassId
					hyperflexreplicationstatus["end_time"] = item.EndTime
					hyperflexreplicationstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.ClassId
						hyperflexerrorstack["message"] = item.Message
						hyperflexerrorstack["message_id"] = item.MessageId
						hyperflexerrorstack["object_type"] = item.ObjectType

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexreplicationstatus["object_type"] = item.ObjectType
					hyperflexreplicationstatus["pack_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexentityreferences []map[string]interface{}
						var ret models.HyperflexEntityReference
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexentityreference := make(map[string]interface{})
						hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexentityreference["class_id"] = item.ClassId
						hyperflexentityreference["confignum"] = item.Confignum
						hyperflexentityreference["id"] = item.Id
						hyperflexentityreference["idtype"] = item.Idtype
						hyperflexentityreference["name"] = item.Name
						hyperflexentityreference["object_type"] = item.ObjectType
						hyperflexentityreference["type"] = item.Type

						hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
						return hyperflexentityreferences
					})(item.GetPackEntityReference(), d)
					hyperflexreplicationstatus["pct_complete"] = item.PctComplete
					hyperflexreplicationstatus["rpo_status"] = (func(p models.HyperflexRpoStatus, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexrpostatuss []map[string]interface{}
						var ret models.HyperflexRpoStatus
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexrpostatus := make(map[string]interface{})
						hyperflexrpostatus["actual"] = item.Actual
						hyperflexrpostatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexrpostatus["class_id"] = item.ClassId
						hyperflexrpostatus["expected"] = item.Expected
						hyperflexrpostatus["object_type"] = item.ObjectType
						hyperflexrpostatus["rpo_exceeded"] = item.RpoExceeded

						hyperflexrpostatuss = append(hyperflexrpostatuss, hyperflexrpostatus)
						return hyperflexrpostatuss
					})(item.GetRpoStatus(), d)
					hyperflexreplicationstatus["start_time"] = item.StartTime
					hyperflexreplicationstatus["status"] = item.Status

					hyperflexreplicationstatuss = append(hyperflexreplicationstatuss, hyperflexreplicationstatus)
					return hyperflexreplicationstatuss
				})(item.GetReplicationStatus(), d)
				hyperflexsnapshotinfobrief["site"] = item.Site
				hyperflexsnapshotinfobrief["snapshot_status"] = (func(p models.HyperflexSnapshotStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexsnapshotstatuss []map[string]interface{}
					var ret models.HyperflexSnapshotStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexsnapshotstatus := make(map[string]interface{})
					hyperflexsnapshotstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexsnapshotstatus["class_id"] = item.ClassId
					hyperflexsnapshotstatus["description"] = item.Description
					hyperflexsnapshotstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.ClassId
						hyperflexerrorstack["message"] = item.Message
						hyperflexerrorstack["message_id"] = item.MessageId
						hyperflexerrorstack["object_type"] = item.ObjectType

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexsnapshotstatus["object_type"] = item.ObjectType
					hyperflexsnapshotstatus["pct_complete"] = item.PctComplete
					hyperflexsnapshotstatus["status"] = item.Status
					hyperflexsnapshotstatus["timestamp"] = item.Timestamp
					hyperflexsnapshotstatus["used_space"] = item.UsedSpace

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
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

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
				hyperflexsnapshotinfobrief["class_id"] = item.ClassId
				hyperflexsnapshotinfobrief["object_type"] = item.ObjectType
				hyperflexsnapshotinfobrief["replication_status"] = (func(p models.HyperflexReplicationStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexreplicationstatuss []map[string]interface{}
					var ret models.HyperflexReplicationStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexreplicationstatus := make(map[string]interface{})
					hyperflexreplicationstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexreplicationstatus["bytes_replicated"] = item.BytesReplicated
					hyperflexreplicationstatus["class_id"] = item.ClassId
					hyperflexreplicationstatus["end_time"] = item.EndTime
					hyperflexreplicationstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.ClassId
						hyperflexerrorstack["message"] = item.Message
						hyperflexerrorstack["message_id"] = item.MessageId
						hyperflexerrorstack["object_type"] = item.ObjectType

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexreplicationstatus["object_type"] = item.ObjectType
					hyperflexreplicationstatus["pack_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexentityreferences []map[string]interface{}
						var ret models.HyperflexEntityReference
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexentityreference := make(map[string]interface{})
						hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexentityreference["class_id"] = item.ClassId
						hyperflexentityreference["confignum"] = item.Confignum
						hyperflexentityreference["id"] = item.Id
						hyperflexentityreference["idtype"] = item.Idtype
						hyperflexentityreference["name"] = item.Name
						hyperflexentityreference["object_type"] = item.ObjectType
						hyperflexentityreference["type"] = item.Type

						hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
						return hyperflexentityreferences
					})(item.GetPackEntityReference(), d)
					hyperflexreplicationstatus["pct_complete"] = item.PctComplete
					hyperflexreplicationstatus["rpo_status"] = (func(p models.HyperflexRpoStatus, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexrpostatuss []map[string]interface{}
						var ret models.HyperflexRpoStatus
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexrpostatus := make(map[string]interface{})
						hyperflexrpostatus["actual"] = item.Actual
						hyperflexrpostatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexrpostatus["class_id"] = item.ClassId
						hyperflexrpostatus["expected"] = item.Expected
						hyperflexrpostatus["object_type"] = item.ObjectType
						hyperflexrpostatus["rpo_exceeded"] = item.RpoExceeded

						hyperflexrpostatuss = append(hyperflexrpostatuss, hyperflexrpostatus)
						return hyperflexrpostatuss
					})(item.GetRpoStatus(), d)
					hyperflexreplicationstatus["start_time"] = item.StartTime
					hyperflexreplicationstatus["status"] = item.Status

					hyperflexreplicationstatuss = append(hyperflexreplicationstatuss, hyperflexreplicationstatus)
					return hyperflexreplicationstatuss
				})(item.GetReplicationStatus(), d)
				hyperflexsnapshotinfobrief["site"] = item.Site
				hyperflexsnapshotinfobrief["snapshot_status"] = (func(p models.HyperflexSnapshotStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexsnapshotstatuss []map[string]interface{}
					var ret models.HyperflexSnapshotStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexsnapshotstatus := make(map[string]interface{})
					hyperflexsnapshotstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexsnapshotstatus["class_id"] = item.ClassId
					hyperflexsnapshotstatus["description"] = item.Description
					hyperflexsnapshotstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
						var hyperflexerrorstacks []map[string]interface{}
						var ret models.HyperflexErrorStack
						if reflect.DeepEqual(ret, p) {
							return nil
						}
						item := p
						hyperflexerrorstack := make(map[string]interface{})
						hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexerrorstack["class_id"] = item.ClassId
						hyperflexerrorstack["message"] = item.Message
						hyperflexerrorstack["message_id"] = item.MessageId
						hyperflexerrorstack["object_type"] = item.ObjectType

						hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
						return hyperflexerrorstacks
					})(item.GetError(), d)
					hyperflexsnapshotstatus["object_type"] = item.ObjectType
					hyperflexsnapshotstatus["pct_complete"] = item.PctComplete
					hyperflexsnapshotstatus["status"] = item.Status
					hyperflexsnapshotstatus["timestamp"] = item.Timestamp
					hyperflexsnapshotstatus["used_space"] = item.UsedSpace

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
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

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
				hyperflexvmprotectionspaceusage["class_id"] = item.ClassId
				hyperflexvmprotectionspaceusage["object_type"] = item.ObjectType
				hyperflexvmprotectionspaceusage["space_usage"] = item.SpaceUsage

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
		hyperflexmapclusteridtostsnapshotpoint["class_id"] = item.ClassId
		hyperflexmapclusteridtostsnapshotpoint["cluster_id"] = item.ClusterId
		hyperflexmapclusteridtostsnapshotpoint["object_type"] = item.ObjectType
		hyperflexmapclusteridtostsnapshotpoint["snapshot_point"] = (func(p models.HyperflexSnapshotPoint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexsnapshotpoints []map[string]interface{}
			var ret models.HyperflexSnapshotPoint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexsnapshotpoint := make(map[string]interface{})
			hyperflexsnapshotpoint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexsnapshotpoint["class_id"] = item.ClassId
			hyperflexsnapshotpoint["cluster_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexentityreferences []map[string]interface{}
				var ret models.HyperflexEntityReference
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexentityreference := make(map[string]interface{})
				hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexentityreference["class_id"] = item.ClassId
				hyperflexentityreference["confignum"] = item.Confignum
				hyperflexentityreference["id"] = item.Id
				hyperflexentityreference["idtype"] = item.Idtype
				hyperflexentityreference["name"] = item.Name
				hyperflexentityreference["object_type"] = item.ObjectType
				hyperflexentityreference["type"] = item.Type

				hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
				return hyperflexentityreferences
			})(item.GetClusterEntityReference(), d)
			hyperflexsnapshotpoint["object_type"] = item.ObjectType
			hyperflexsnapshotpoint["replication_status"] = (func(p models.HyperflexReplicationStatus, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationstatuss []map[string]interface{}
				var ret models.HyperflexReplicationStatus
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationstatus := make(map[string]interface{})
				hyperflexreplicationstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationstatus["bytes_replicated"] = item.BytesReplicated
				hyperflexreplicationstatus["class_id"] = item.ClassId
				hyperflexreplicationstatus["end_time"] = item.EndTime
				hyperflexreplicationstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexerrorstacks []map[string]interface{}
					var ret models.HyperflexErrorStack
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexerrorstack := make(map[string]interface{})
					hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexerrorstack["class_id"] = item.ClassId
					hyperflexerrorstack["message"] = item.Message
					hyperflexerrorstack["message_id"] = item.MessageId
					hyperflexerrorstack["object_type"] = item.ObjectType

					hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
					return hyperflexerrorstacks
				})(item.GetError(), d)
				hyperflexreplicationstatus["object_type"] = item.ObjectType
				hyperflexreplicationstatus["pack_entity_reference"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetPackEntityReference(), d)
				hyperflexreplicationstatus["pct_complete"] = item.PctComplete
				hyperflexreplicationstatus["rpo_status"] = (func(p models.HyperflexRpoStatus, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexrpostatuss []map[string]interface{}
					var ret models.HyperflexRpoStatus
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexrpostatus := make(map[string]interface{})
					hyperflexrpostatus["actual"] = item.Actual
					hyperflexrpostatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexrpostatus["class_id"] = item.ClassId
					hyperflexrpostatus["expected"] = item.Expected
					hyperflexrpostatus["object_type"] = item.ObjectType
					hyperflexrpostatus["rpo_exceeded"] = item.RpoExceeded

					hyperflexrpostatuss = append(hyperflexrpostatuss, hyperflexrpostatus)
					return hyperflexrpostatuss
				})(item.GetRpoStatus(), d)
				hyperflexreplicationstatus["start_time"] = item.StartTime
				hyperflexreplicationstatus["status"] = item.Status

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
				hyperflexsnapshotfiles["class_id"] = item.ClassId
				hyperflexsnapshotfiles["name_tracked_files"] = (func(p []models.HyperflexFilePath, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexfilepaths []map[string]interface{}
					if len(p) == 0 {
						return nil
					}
					for _, item := range p {
						hyperflexfilepath := make(map[string]interface{})
						hyperflexfilepath["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexfilepath["class_id"] = item.ClassId
						hyperflexfilepath["ds_info"] = (func(p models.HyperflexDatastoreInfo, d *schema.ResourceData) []map[string]interface{} {
							var hyperflexdatastoreinfos []map[string]interface{}
							var ret models.HyperflexDatastoreInfo
							if reflect.DeepEqual(ret, p) {
								return nil
							}
							item := p
							hyperflexdatastoreinfo := make(map[string]interface{})
							hyperflexdatastoreinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							hyperflexdatastoreinfo["class_id"] = item.ClassId
							hyperflexdatastoreinfo["ds_backend_id"] = item.DsBackendId
							hyperflexdatastoreinfo["ds_frontend_id"] = item.DsFrontendId
							hyperflexdatastoreinfo["object_type"] = item.ObjectType

							hyperflexdatastoreinfos = append(hyperflexdatastoreinfos, hyperflexdatastoreinfo)
							return hyperflexdatastoreinfos
						})(item.GetDsInfo(), d)
						hyperflexfilepath["object_type"] = item.ObjectType
						hyperflexfilepath["relative_file_path"] = item.RelativeFilePath
						hyperflexfilepaths = append(hyperflexfilepaths, hyperflexfilepath)
					}
					return hyperflexfilepaths
				})(item.GetNameTrackedFiles(), d)
				hyperflexsnapshotfiles["object_type"] = item.ObjectType
				hyperflexsnapshotfiles["uuid_tracked_disks_map"] = (func(p []models.HyperflexMapUuidToTrackedDisk, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexmapuuidtotrackeddisks []map[string]interface{}
					if len(p) == 0 {
						return nil
					}
					for _, item := range p {
						hyperflexmapuuidtotrackeddisk := make(map[string]interface{})
						hyperflexmapuuidtotrackeddisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						hyperflexmapuuidtotrackeddisk["class_id"] = item.ClassId
						hyperflexmapuuidtotrackeddisk["object_type"] = item.ObjectType
						hyperflexmapuuidtotrackeddisk["tracked_disk"] = (func(p models.HyperflexTrackedDisk, d *schema.ResourceData) []map[string]interface{} {
							var hyperflextrackeddisks []map[string]interface{}
							var ret models.HyperflexTrackedDisk
							if reflect.DeepEqual(ret, p) {
								return nil
							}
							item := p
							hyperflextrackeddisk := make(map[string]interface{})
							hyperflextrackeddisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							hyperflextrackeddisk["class_id"] = item.ClassId
							hyperflextrackeddisk["disk_files"] = (func(p []models.HyperflexTrackedFile, d *schema.ResourceData) []map[string]interface{} {
								var hyperflextrackedfiles []map[string]interface{}
								if len(p) == 0 {
									return nil
								}
								for _, item := range p {
									hyperflextrackedfile := make(map[string]interface{})
									hyperflextrackedfile["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
									hyperflextrackedfile["class_id"] = item.ClassId
									hyperflextrackedfile["file_path"] = (func(p models.HyperflexFilePath, d *schema.ResourceData) []map[string]interface{} {
										var hyperflexfilepaths []map[string]interface{}
										var ret models.HyperflexFilePath
										if reflect.DeepEqual(ret, p) {
											return nil
										}
										item := p
										hyperflexfilepath := make(map[string]interface{})
										hyperflexfilepath["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
										hyperflexfilepath["class_id"] = item.ClassId
										hyperflexfilepath["ds_info"] = (func(p models.HyperflexDatastoreInfo, d *schema.ResourceData) []map[string]interface{} {
											var hyperflexdatastoreinfos []map[string]interface{}
											var ret models.HyperflexDatastoreInfo
											if reflect.DeepEqual(ret, p) {
												return nil
											}
											item := p
											hyperflexdatastoreinfo := make(map[string]interface{})
											hyperflexdatastoreinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
											hyperflexdatastoreinfo["class_id"] = item.ClassId
											hyperflexdatastoreinfo["ds_backend_id"] = item.DsBackendId
											hyperflexdatastoreinfo["ds_frontend_id"] = item.DsFrontendId
											hyperflexdatastoreinfo["object_type"] = item.ObjectType

											hyperflexdatastoreinfos = append(hyperflexdatastoreinfos, hyperflexdatastoreinfo)
											return hyperflexdatastoreinfos
										})(item.GetDsInfo(), d)
										hyperflexfilepath["object_type"] = item.ObjectType
										hyperflexfilepath["relative_file_path"] = item.RelativeFilePath

										hyperflexfilepaths = append(hyperflexfilepaths, hyperflexfilepath)
										return hyperflexfilepaths
									})(item.GetFilePath(), d)
									hyperflextrackedfile["file_type"] = item.FileType
									hyperflextrackedfile["object_type"] = item.ObjectType
									hyperflextrackedfiles = append(hyperflextrackedfiles, hyperflextrackedfile)
								}
								return hyperflextrackedfiles
							})(item.GetDiskFiles(), d)
							hyperflextrackeddisk["disk_type"] = item.DiskType
							hyperflextrackeddisk["object_type"] = item.ObjectType

							hyperflextrackeddisks = append(hyperflextrackeddisks, hyperflextrackeddisk)
							return hyperflextrackeddisks
						})(item.GetTrackedDisk(), d)
						hyperflexmapuuidtotrackeddisk["uuid"] = item.Uuid
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
				hyperflexentityreference["class_id"] = item.ClassId
				hyperflexentityreference["confignum"] = item.Confignum
				hyperflexentityreference["id"] = item.Id
				hyperflexentityreference["idtype"] = item.Idtype
				hyperflexentityreference["name"] = item.Name
				hyperflexentityreference["object_type"] = item.ObjectType
				hyperflexentityreference["type"] = item.Type

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
				hyperflexsnapshotstatus["class_id"] = item.ClassId
				hyperflexsnapshotstatus["description"] = item.Description
				hyperflexsnapshotstatus["error"] = (func(p models.HyperflexErrorStack, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexerrorstacks []map[string]interface{}
					var ret models.HyperflexErrorStack
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexerrorstack := make(map[string]interface{})
					hyperflexerrorstack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexerrorstack["class_id"] = item.ClassId
					hyperflexerrorstack["message"] = item.Message
					hyperflexerrorstack["message_id"] = item.MessageId
					hyperflexerrorstack["object_type"] = item.ObjectType

					hyperflexerrorstacks = append(hyperflexerrorstacks, hyperflexerrorstack)
					return hyperflexerrorstacks
				})(item.GetError(), d)
				hyperflexsnapshotstatus["object_type"] = item.ObjectType
				hyperflexsnapshotstatus["pct_complete"] = item.PctComplete
				hyperflexsnapshotstatus["status"] = item.Status
				hyperflexsnapshotstatus["timestamp"] = item.Timestamp
				hyperflexsnapshotstatus["used_space"] = item.UsedSpace

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
				hyperflexvirtualmachine["class_id"] = item.ClassId
				hyperflexvirtualmachine["object_type"] = item.ObjectType
				hyperflexvirtualmachine["run_time_info"] = (func(p models.HyperflexVirtualMachineRuntimeInfo, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexvirtualmachineruntimeinfos []map[string]interface{}
					var ret models.HyperflexVirtualMachineRuntimeInfo
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexvirtualmachineruntimeinfo := make(map[string]interface{})
					hyperflexvirtualmachineruntimeinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexvirtualmachineruntimeinfo["bios_uuid"] = item.BiosUuid
					hyperflexvirtualmachineruntimeinfo["class_id"] = item.ClassId
					hyperflexvirtualmachineruntimeinfo["connection_state"] = item.ConnectionState
					hyperflexvirtualmachineruntimeinfo["cpu_usage"] = item.CpuUsage
					hyperflexvirtualmachineruntimeinfo["folder"] = item.Folder
					hyperflexvirtualmachineruntimeinfo["guest_family"] = item.GuestFamily
					hyperflexvirtualmachineruntimeinfo["guest_full_name"] = item.GuestFullName
					hyperflexvirtualmachineruntimeinfo["guest_id"] = item.GuestId
					hyperflexvirtualmachineruntimeinfo["guest_state"] = item.GuestState
					hyperflexvirtualmachineruntimeinfo["host_name"] = item.HostName
					hyperflexvirtualmachineruntimeinfo["instance_uuid"] = item.InstanceUuid
					hyperflexvirtualmachineruntimeinfo["memory_mb"] = item.MemoryMb
					hyperflexvirtualmachineruntimeinfo["memory_usage"] = item.MemoryUsage
					hyperflexvirtualmachineruntimeinfo["moid"] = item.Moid
					hyperflexvirtualmachineruntimeinfo["name"] = item.Name
					hyperflexvirtualmachineruntimeinfo["networks"] = item.Networks
					hyperflexvirtualmachineruntimeinfo["num_cpu"] = item.NumCpu
					hyperflexvirtualmachineruntimeinfo["object_type"] = item.ObjectType
					hyperflexvirtualmachineruntimeinfo["power_state"] = item.PowerState
					hyperflexvirtualmachineruntimeinfo["provisioned_size"] = item.ProvisionedSize
					hyperflexvirtualmachineruntimeinfo["resource_pool"] = item.ResourcePool
					hyperflexvirtualmachineruntimeinfo["used_size"] = item.UsedSize
					hyperflexvirtualmachineruntimeinfo["nr_version"] = item.Version
					hyperflexvirtualmachineruntimeinfo["vmx_path"] = item.VmxPath

					hyperflexvirtualmachineruntimeinfos = append(hyperflexvirtualmachineruntimeinfos, hyperflexvirtualmachineruntimeinfo)
					return hyperflexvirtualmachineruntimeinfos
				})(item.GetRunTimeInfo(), d)
				hyperflexvirtualmachine["status_code"] = item.StatusCode
				hyperflexvirtualmachine["uuid"] = item.Uuid

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
		hyperflexnamedvlan["class_id"] = item.ClassId
		hyperflexnamedvlan["name"] = item.Name
		hyperflexnamedvlan["object_type"] = item.ObjectType
		hyperflexnamedvlan["vlan_id"] = item.VlanId
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
			hyperflexbondstate["active_slave"] = item.ActiveSlave
			hyperflexbondstate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexbondstate["class_id"] = item.ClassId
			hyperflexbondstate["mode"] = item.Mode
			hyperflexbondstate["object_type"] = item.ObjectType
			hyperflexbondstate["slaves"] = item.Slaves

			hyperflexbondstates = append(hyperflexbondstates, hyperflexbondstate)
			return hyperflexbondstates
		})(item.GetBondState(), d)
		hyperflexnetworkport["class_id"] = item.ClassId
		hyperflexnetworkport["name"] = item.Name
		hyperflexnetworkport["net_interfaces"] = item.NetInterfaces
		hyperflexnetworkport["object_type"] = item.ObjectType
		hyperflexnetworkport["port_type"] = item.PortType
		hyperflexnetworkport["vlans"] = item.Vlans
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
		hyperflexreplicationclusterreferencetoschedule["class_id"] = item.ClassId
		hyperflexreplicationclusterreferencetoschedule["object_type"] = item.ObjectType
		hyperflexreplicationclusterreferencetoschedule["schedule"] = (func(p models.HyperflexReplicationSchedule, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexreplicationschedules []map[string]interface{}
			var ret models.HyperflexReplicationSchedule
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexreplicationschedule := make(map[string]interface{})
			hyperflexreplicationschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexreplicationschedule["backup_interval"] = item.BackupInterval
			hyperflexreplicationschedule["class_id"] = item.ClassId
			hyperflexreplicationschedule["object_type"] = item.ObjectType

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
			hyperflexentityreference["class_id"] = item.ClassId
			hyperflexentityreference["confignum"] = item.Confignum
			hyperflexentityreference["id"] = item.Id
			hyperflexentityreference["idtype"] = item.Idtype
			hyperflexentityreference["name"] = item.Name
			hyperflexentityreference["object_type"] = item.ObjectType
			hyperflexentityreference["type"] = item.Type

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
		hyperflexserverfirmwareversioninfo["class_id"] = item.ClassId
		hyperflexserverfirmwareversioninfo["object_type"] = item.ObjectType
		hyperflexserverfirmwareversioninfo["server_platform"] = item.ServerPlatform
		hyperflexserverfirmwareversioninfo["nr_version"] = item.Version
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
		hyperflexservermodelentry["class_id"] = item.ClassId
		hyperflexservermodelentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.ClassId
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexservermodelentry["name"] = item.Name
		hyperflexservermodelentry["object_type"] = item.ObjectType
		hyperflexservermodelentry["value"] = item.Value
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
		hyperflexvmdisk["boot_order"] = item.BootOrder
		hyperflexvmdisk["bus"] = item.Bus
		hyperflexvmdisk["class_id"] = item.ClassId
		hyperflexvmdisk["name"] = item.Name
		hyperflexvmdisk["object_type"] = item.ObjectType
		hyperflexvmdisk["type"] = item.Type
		hyperflexvmdisk["virtual_disk"] = (func(p models.HyperflexVdiskConfig, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexvdiskconfigs []map[string]interface{}
			var ret models.HyperflexVdiskConfig
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexvdiskconfig := make(map[string]interface{})
			hyperflexvdiskconfig["access_mode"] = item.AccessMode
			hyperflexvdiskconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexvdiskconfig["capacity"] = item.Capacity
			hyperflexvdiskconfig["class_id"] = item.ClassId
			hyperflexvdiskconfig["mode"] = item.Mode
			hyperflexvdiskconfig["name"] = item.Name
			hyperflexvdiskconfig["object_type"] = item.ObjectType
			hyperflexvdiskconfig["source_file_path"] = item.SourceFilePath
			hyperflexvdiskconfig["source_virtual_disk"] = item.SourceVirtualDisk
			hyperflexvdiskconfig["status"] = (func(p models.HyperflexDiskStatus, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexdiskstatuss []map[string]interface{}
				var ret models.HyperflexDiskStatus
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexdiskstatus := make(map[string]interface{})
				hyperflexdiskstatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexdiskstatus["class_id"] = item.ClassId
				hyperflexdiskstatus["download_percentage"] = item.DownloadPercentage
				hyperflexdiskstatus["object_type"] = item.ObjectType
				hyperflexdiskstatus["state"] = item.State
				hyperflexdiskstatus["volume_handle"] = item.VolumeHandle
				hyperflexdiskstatus["volume_name"] = item.VolumeName

				hyperflexdiskstatuss = append(hyperflexdiskstatuss, hyperflexdiskstatus)
				return hyperflexdiskstatuss
			})(item.GetStatus(), d)
			hyperflexvdiskconfig["uuid"] = item.Uuid

			hyperflexvdiskconfigs = append(hyperflexvdiskconfigs, hyperflexvdiskconfig)
			return hyperflexvdiskconfigs
		})(item.GetVirtualDisk(), d)
		hyperflexvmdisk["virtual_disk_reference"] = item.VirtualDiskReference
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
		hyperflexvminterface["bridge"] = item.Bridge
		hyperflexvminterface["class_id"] = item.ClassId
		hyperflexvminterface["ip_address"] = item.IpAddress
		hyperflexvminterface["mac_address"] = item.MacAddress
		hyperflexvminterface["model"] = item.Model
		hyperflexvminterface["name"] = item.Name
		hyperflexvminterface["object_type"] = item.ObjectType
		hyperflexvminterfaces = append(hyperflexvminterfaces, hyperflexvminterface)
	}
	return hyperflexvminterfaces
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
		iaaslicensekeysinfo["class_id"] = item.ClassId
		iaaslicensekeysinfo["nr_count"] = item.Count
		iaaslicensekeysinfo["expiration_date"] = item.ExpirationDate
		iaaslicensekeysinfo["license_id"] = item.LicenseId
		iaaslicensekeysinfo["object_type"] = item.ObjectType
		iaaslicensekeysinfo["pid"] = item.Pid
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
		iaaslicenseutilizationinfo["actual_used"] = item.ActualUsed
		iaaslicenseutilizationinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iaaslicenseutilizationinfo["class_id"] = item.ClassId
		iaaslicenseutilizationinfo["label"] = item.Label
		iaaslicenseutilizationinfo["licensed_limit"] = item.LicensedLimit
		iaaslicenseutilizationinfo["object_type"] = item.ObjectType
		iaaslicenseutilizationinfo["sku"] = item.Sku
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
		iaasworkflowsteps["class_id"] = item.ClassId
		iaasworkflowsteps["completed_time"] = item.CompletedTime
		iaasworkflowsteps["name"] = item.Name
		iaasworkflowsteps["object_type"] = item.ObjectType
		iaasworkflowsteps["status"] = item.Status
		iaasworkflowsteps["status_message"] = item.StatusMessage
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
		iamaccountpermissions["account_identifier"] = item.AccountIdentifier
		iamaccountpermissions["account_name"] = item.AccountName
		iamaccountpermissions["account_status"] = item.AccountStatus
		iamaccountpermissions["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamaccountpermissions["class_id"] = item.ClassId
		iamaccountpermissions["object_type"] = item.ObjectType
		iamaccountpermissions["permissions"] = (func(p []models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var iampermissionreferences []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				iampermissionreference := make(map[string]interface{})
				iampermissionreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				iampermissionreference["class_id"] = item.ClassId
				iampermissionreference["object_type"] = item.ObjectType
				iampermissionreference["permission_identifier"] = item.PermissionIdentifier
				iampermissionreference["permission_name"] = item.PermissionName
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
		iamfeaturedefinition["class_id"] = item.ClassId
		iamfeaturedefinition["feature"] = item.Feature
		iamfeaturedefinition["object_type"] = item.ObjectType
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
		iamgrouppermissiontoroles["class_id"] = item.ClassId
		iamgrouppermissiontoroles["group"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			var ret models.MoMoRef
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.ClassId
			momoref["moid"] = item.Moid
			momoref["object_type"] = item.ObjectType
			momoref["selector"] = item.Selector

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetGroup(), d)
		iamgrouppermissiontoroles["object_type"] = item.ObjectType
		iamgrouppermissiontoroles["orgs"] = (func(p []models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				momoref := make(map[string]interface{})
				momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				momoref["class_id"] = item.ClassId
				momoref["moid"] = item.Moid
				momoref["object_type"] = item.ObjectType
				momoref["selector"] = item.Selector
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
		iampermissiontoroles["class_id"] = item.ClassId
		iampermissiontoroles["object_type"] = item.ObjectType
		iampermissiontoroles["permission"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
			var momorefs []map[string]interface{}
			var ret models.MoMoRef
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.ClassId
			momoref["moid"] = item.Moid
			momoref["object_type"] = item.ObjectType
			momoref["selector"] = item.Selector

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
				momoref["class_id"] = item.ClassId
				momoref["moid"] = item.Moid
				momoref["object_type"] = item.ObjectType
				momoref["selector"] = item.Selector
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
		inframetadata["class_id"] = item.ClassId
		inframetadata["name"] = item.Name
		inframetadata["object_type"] = item.ObjectType
		inframetadata["value"] = item.Value
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
		ippoolipv4block["class_id"] = item.ClassId
		ippoolipv4block["from"] = item.From
		ippoolipv4block["object_type"] = item.ObjectType
		ippoolipv4block["size"] = item.Size
		ippoolipv4block["to"] = item.To
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
		ippoolipv6block["class_id"] = item.ClassId
		ippoolipv6block["from"] = item.From
		ippoolipv6block["object_type"] = item.ObjectType
		ippoolipv6block["size"] = item.Size
		ippoolipv6block["to"] = item.To
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
		iqnpooliqnsuffixblock["class_id"] = item.ClassId
		iqnpooliqnsuffixblock["from"] = item.From
		iqnpooliqnsuffixblock["object_type"] = item.ObjectType
		iqnpooliqnsuffixblock["size"] = item.Size
		iqnpooliqnsuffixblock["suffix"] = item.Suffix
		iqnpooliqnsuffixblock["to"] = item.To
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
			kubernetesaddonconfiguration["class_id"] = item.ClassId
			kubernetesaddonconfiguration["install_strategy"] = item.InstallStrategy
			kubernetesaddonconfiguration["object_type"] = item.ObjectType
			kubernetesaddonconfiguration["override_sets"] = (func(p []models.KubernetesKeyValue, d *schema.ResourceData) []map[string]interface{} {
				var kuberneteskeyvalues []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					kuberneteskeyvalue := make(map[string]interface{})
					kuberneteskeyvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					kuberneteskeyvalue["class_id"] = item.ClassId
					kuberneteskeyvalue["key"] = item.Key
					kuberneteskeyvalue["object_type"] = item.ObjectType
					kuberneteskeyvalue["value"] = item.Value
					kuberneteskeyvalues = append(kuberneteskeyvalues, kuberneteskeyvalue)
				}
				return kuberneteskeyvalues
			})(item.GetOverrideSets(), d)
			kubernetesaddonconfiguration["overrides"] = item.Overrides
			kubernetesaddonconfiguration["release_name"] = item.ReleaseName
			kubernetesaddonconfiguration["release_namespace"] = item.ReleaseNamespace
			kubernetesaddonconfiguration["upgrade_strategy"] = item.UpgradeStrategy

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
			momoref["class_id"] = item.ClassId
			momoref["moid"] = item.Moid
			momoref["object_type"] = item.ObjectType
			momoref["selector"] = item.Selector

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetAddonPolicy(), d)
		kubernetesaddon["class_id"] = item.ClassId
		kubernetesaddon["name"] = item.Name
		kubernetesaddon["object_type"] = item.ObjectType
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
			kubernetesaddonconfiguration["class_id"] = item.ClassId
			kubernetesaddonconfiguration["install_strategy"] = item.InstallStrategy
			kubernetesaddonconfiguration["object_type"] = item.ObjectType
			kubernetesaddonconfiguration["override_sets"] = (func(p []models.KubernetesKeyValue, d *schema.ResourceData) []map[string]interface{} {
				var kuberneteskeyvalues []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					kuberneteskeyvalue := make(map[string]interface{})
					kuberneteskeyvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					kuberneteskeyvalue["class_id"] = item.ClassId
					kuberneteskeyvalue["key"] = item.Key
					kuberneteskeyvalue["object_type"] = item.ObjectType
					kuberneteskeyvalue["value"] = item.Value
					kuberneteskeyvalues = append(kuberneteskeyvalues, kuberneteskeyvalue)
				}
				return kuberneteskeyvalues
			})(item.GetOverrideSets(), d)
			kubernetesaddonconfiguration["overrides"] = item.Overrides
			kubernetesaddonconfiguration["release_name"] = item.ReleaseName
			kubernetesaddonconfiguration["release_namespace"] = item.ReleaseNamespace
			kubernetesaddonconfiguration["upgrade_strategy"] = item.UpgradeStrategy

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
			momoref["class_id"] = item.ClassId
			momoref["moid"] = item.Moid
			momoref["object_type"] = item.ObjectType
			momoref["selector"] = item.Selector

			momorefs = append(momorefs, momoref)
			return momorefs
		})(item.GetAddonDefinition(), d)
		kubernetesessentialaddon["class_id"] = item.ClassId
		kubernetesessentialaddon["name"] = item.Name
		kubernetesessentialaddon["object_type"] = item.ObjectType
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
		kubernetesnodeaddress["address"] = item.Address
		kubernetesnodeaddress["class_id"] = item.ClassId
		kubernetesnodeaddress["object_type"] = item.ObjectType
		kubernetesnodeaddress["type"] = item.Type
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
		kubernetesnodegrouplabel["class_id"] = item.ClassId
		kubernetesnodegrouplabel["key"] = item.Key
		kubernetesnodegrouplabel["object_type"] = item.ObjectType
		kubernetesnodegrouplabel["value"] = item.Value
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
		kubernetesnodegrouptaint["class_id"] = item.ClassId
		kubernetesnodegrouptaint["effect"] = item.Effect
		kubernetesnodegrouptaint["key"] = item.Key
		kubernetesnodegrouptaint["object_type"] = item.ObjectType
		kubernetesnodegrouptaint["value"] = item.Value
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
		kubernetesnodestatus["class_id"] = item.ClassId
		kubernetesnodestatus["object_type"] = item.ObjectType
		kubernetesnodestatus["status"] = item.Status
		kubernetesnodestatus["type"] = item.Type
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
		macpoolblock["class_id"] = item.ClassId
		macpoolblock["from"] = item.From
		macpoolblock["object_type"] = item.ObjectType
		macpoolblock["size"] = item.Size
		macpoolblock["to"] = item.To
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
		memorypersistentmemorygoal["class_id"] = item.ClassId
		memorypersistentmemorygoal["memory_mode_percentage"] = item.MemoryModePercentage
		memorypersistentmemorygoal["object_type"] = item.ObjectType
		memorypersistentmemorygoal["persistent_memory_type"] = item.PersistentMemoryType
		memorypersistentmemorygoal["socket_id"] = item.SocketId
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
		memorypersistentmemorylogicalnamespace["capacity"] = item.Capacity
		memorypersistentmemorylogicalnamespace["class_id"] = item.ClassId
		memorypersistentmemorylogicalnamespace["mode"] = item.Mode
		memorypersistentmemorylogicalnamespace["name"] = item.Name
		memorypersistentmemorylogicalnamespace["object_type"] = item.ObjectType
		memorypersistentmemorylogicalnamespace["socket_id"] = item.SocketId
		memorypersistentmemorylogicalnamespace["socket_memory_id"] = item.SocketMemoryId
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
		metaaccessprivilege["class_id"] = item.ClassId
		metaaccessprivilege["method"] = item.Method
		metaaccessprivilege["object_type"] = item.ObjectType
		metaaccessprivilege["privilege"] = item.Privilege
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
		metadisplaynamedefinition["class_id"] = item.ClassId
		metadisplaynamedefinition["format"] = item.Format
		metadisplaynamedefinition["include_ancestor"] = item.IncludeAncestor
		metadisplaynamedefinition["name"] = item.Name
		metadisplaynamedefinition["object_type"] = item.ObjectType
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
		metapropdefinition["api_access"] = item.ApiAccess
		metapropdefinition["class_id"] = item.ClassId
		metapropdefinition["name"] = item.Name
		metapropdefinition["object_type"] = item.ObjectType
		metapropdefinition["op_security"] = item.OpSecurity
		metapropdefinition["search_weight"] = item.SearchWeight
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
		metarelationshipdefinition["api_access"] = item.ApiAccess
		metarelationshipdefinition["class_id"] = item.ClassId
		metarelationshipdefinition["collection"] = item.Collection
		metarelationshipdefinition["export"] = item.Export
		metarelationshipdefinition["export_with_peer"] = item.ExportWithPeer
		metarelationshipdefinition["name"] = item.Name
		metarelationshipdefinition["object_type"] = item.ObjectType
		metarelationshipdefinition["peer_rel_name"] = item.PeerRelName
		metarelationshipdefinition["peer_sync"] = item.PeerSync
		metarelationshipdefinition["type"] = item.Type
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
		motag["key"] = item.Key
		motag["value"] = item.Value
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
		niaapidetail["chksum"] = item.Chksum
		niaapidetail["class_id"] = item.ClassId
		niaapidetail["filename"] = item.Filename
		niaapidetail["name"] = item.Name
		niaapidetail["object_type"] = item.ObjectType
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
		niaapirevisioninfo["class_id"] = item.ClassId
		niaapirevisioninfo["date_published"] = item.DatePublished
		niaapirevisioninfo["object_type"] = item.ObjectType
		niaapirevisioninfo["revision_comment"] = item.RevisionComment
		niaapirevisioninfo["revision_no"] = item.RevisionNo
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
		niatelemetrylogicallink["class_id"] = item.ClassId
		niatelemetrylogicallink["db_id"] = item.DbId
		niatelemetrylogicallink["is_present"] = item.IsPresent
		niatelemetrylogicallink["link_addr1"] = item.LinkAddr1
		niatelemetrylogicallink["link_addr2"] = item.LinkAddr2
		niatelemetrylogicallink["link_state"] = item.LinkState
		niatelemetrylogicallink["link_type"] = item.LinkType
		niatelemetrylogicallink["object_type"] = item.ObjectType
		niatelemetrylogicallink["uptime"] = item.Uptime
		niatelemetrylogicallinks = append(niatelemetrylogicallinks, niatelemetrylogicallink)
	}
	return niatelemetrylogicallinks
}
func flattenListOnpremImagePackage(p []models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var onpremimagepackages []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremimagepackage := make(map[string]interface{})
		onpremimagepackage["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremimagepackage["class_id"] = item.ClassId
		onpremimagepackage["file_path"] = item.FilePath
		onpremimagepackage["file_sha"] = item.FileSha
		onpremimagepackage["file_size"] = item.FileSize
		onpremimagepackage["file_time"] = item.FileTime
		onpremimagepackage["filename"] = item.Filename
		onpremimagepackage["name"] = item.Name
		onpremimagepackage["object_type"] = item.ObjectType
		onpremimagepackage["package_type"] = item.PackageType
		onpremimagepackage["nr_version"] = item.Version
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
		onpremupgradenote["class_id"] = item.ClassId
		onpremupgradenote["message"] = item.Message
		onpremupgradenote["object_type"] = item.ObjectType
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
		onpremupgradephase["class_id"] = item.ClassId
		onpremupgradephase["elapsed_time"] = item.ElapsedTime
		onpremupgradephase["end_time"] = item.EndTime
		onpremupgradephase["failed"] = item.Failed
		onpremupgradephase["message"] = item.Message
		onpremupgradephase["name"] = item.Name
		onpremupgradephase["object_type"] = item.ObjectType
		onpremupgradephase["start_time"] = item.StartTime
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
		oprskvpair["class_id"] = item.ClassId
		oprskvpair["key"] = item.Key
		oprskvpair["object_type"] = item.ObjectType
		oprskvpair["value"] = item.Value
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
		osplaceholder["class_id"] = item.ClassId
		osplaceholder["is_value_set"] = item.IsValueSet
		osplaceholder["object_type"] = item.ObjectType
		osplaceholder["type"] = (func(p models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {
			var workflowprimitivedatatypes []map[string]interface{}
			var ret models.WorkflowPrimitiveDataType
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowprimitivedatatype := make(map[string]interface{})
			workflowprimitivedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowprimitivedatatype["class_id"] = item.ClassId
			workflowprimitivedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.ClassId
				workflowdefaultvalue["is_value_set"] = item.IsValueSet
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowprimitivedatatype["description"] = item.Description
			workflowprimitivedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
				var workflowdisplaymetas []map[string]interface{}
				var ret models.WorkflowDisplayMeta
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdisplaymeta := make(map[string]interface{})
				workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdisplaymeta["class_id"] = item.ClassId
				workflowdisplaymeta["inventory_selector"] = item.InventorySelector
				workflowdisplaymeta["object_type"] = item.ObjectType
				workflowdisplaymeta["widget_type"] = item.WidgetType

				workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
				return workflowdisplaymetas
			})(item.GetDisplayMeta(), d)
			workflowprimitivedatatype["input_parameters"] = item.InputParameters
			workflowprimitivedatatype["label"] = item.Label
			workflowprimitivedatatype["name"] = item.Name
			workflowprimitivedatatype["object_type"] = item.ObjectType
			workflowprimitivedatatype["properties"] = (func(p models.WorkflowPrimitiveDataProperty, d *schema.ResourceData) []map[string]interface{} {
				var workflowprimitivedatapropertys []map[string]interface{}
				var ret models.WorkflowPrimitiveDataProperty
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowprimitivedataproperty := make(map[string]interface{})
				workflowprimitivedataproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowprimitivedataproperty["class_id"] = item.ClassId
				workflowprimitivedataproperty["constraints"] = (func(p models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {
					var workflowconstraintss []map[string]interface{}
					var ret models.WorkflowConstraints
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					workflowconstraints := make(map[string]interface{})
					workflowconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					workflowconstraints["class_id"] = item.ClassId
					workflowconstraints["enum_list"] = (func(p []models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var workflowenumentrys []map[string]interface{}
						if len(p) == 0 {
							return nil
						}
						for _, item := range p {
							workflowenumentry := make(map[string]interface{})
							workflowenumentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							workflowenumentry["class_id"] = item.ClassId
							workflowenumentry["label"] = item.Label
							workflowenumentry["object_type"] = item.ObjectType
							workflowenumentry["value"] = item.Value
							workflowenumentrys = append(workflowenumentrys, workflowenumentry)
						}
						return workflowenumentrys
					})(item.GetEnumList(), d)
					workflowconstraints["max"] = item.Max
					workflowconstraints["min"] = item.Min
					workflowconstraints["object_type"] = item.ObjectType
					workflowconstraints["regex"] = item.Regex

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
						workflowmoreferenceproperty["class_id"] = item.ClassId
						workflowmoreferenceproperty["display_attributes"] = item.DisplayAttributes
						workflowmoreferenceproperty["object_type"] = item.ObjectType
						workflowmoreferenceproperty["selector"] = item.Selector
						workflowmoreferenceproperty["value_attribute"] = item.ValueAttribute
						workflowmoreferencepropertys = append(workflowmoreferencepropertys, workflowmoreferenceproperty)
					}
					return workflowmoreferencepropertys
				})(item.GetInventorySelector(), d)
				workflowprimitivedataproperty["object_type"] = item.ObjectType
				workflowprimitivedataproperty["secure"] = item.Secure
				workflowprimitivedataproperty["type"] = item.Type

				workflowprimitivedatapropertys = append(workflowprimitivedatapropertys, workflowprimitivedataproperty)
				return workflowprimitivedatapropertys
			})(item.GetProperties(), d)
			workflowprimitivedatatype["required"] = item.Required

			workflowprimitivedatatypes = append(workflowprimitivedatatypes, workflowprimitivedatatype)
			return workflowprimitivedatatypes
		})(item.GetType(), d)
		osplaceholder["value"] = item.Value
		osplaceholders = append(osplaceholders, osplaceholder)
	}
	return osplaceholders
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
		policyinventoryjobinfo["class_id"] = item.ClassId
		policyinventoryjobinfo["execution_status"] = item.ExecutionStatus
		policyinventoryjobinfo["last_scheduled_time"] = item.LastScheduledTime
		policyinventoryjobinfo["object_type"] = item.ObjectType
		policyinventoryjobinfo["policy_id"] = item.PolicyId
		policyinventoryjobinfo["policy_name"] = item.PolicyName
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
		resourcepertypecombinedselector["class_id"] = item.ClassId
		resourcepertypecombinedselector["combined_selector"] = item.CombinedSelector
		resourcepertypecombinedselector["empty_filter"] = item.EmptyFilter
		resourcepertypecombinedselector["object_type"] = item.ObjectType
		resourcepertypecombinedselector["selector_object_type"] = item.SelectorObjectType
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
		resourceselector["class_id"] = item.ClassId
		resourceselector["object_type"] = item.ObjectType
		resourceselector["selector"] = item.Selector
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
		sdcardpartition["class_id"] = item.ClassId
		sdcardpartition["object_type"] = item.ObjectType
		sdcardpartition["type"] = item.Type
		sdcardpartition["virtual_drives"] = (func(p []models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var sdcardvirtualdrives []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				sdcardvirtualdrive := make(map[string]interface{})
				sdcardvirtualdrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				sdcardvirtualdrive["class_id"] = item.ClassId
				sdcardvirtualdrive["enable"] = item.Enable
				sdcardvirtualdrive["object_type"] = item.ObjectType
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
		sdwannetworkconfigurationtype["class_id"] = item.ClassId
		sdwannetworkconfigurationtype["network_type"] = item.NetworkType
		sdwannetworkconfigurationtype["object_type"] = item.ObjectType
		sdwannetworkconfigurationtype["port_group"] = item.PortGroup
		sdwannetworkconfigurationtype["vlan"] = item.Vlan
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
		sdwantemplateinputstype["class_id"] = item.ClassId
		sdwantemplateinputstype["editable"] = item.Editable
		sdwantemplateinputstype["key"] = item.Key
		sdwantemplateinputstype["object_type"] = item.ObjectType
		sdwantemplateinputstype["required"] = item.Required
		sdwantemplateinputstype["template"] = item.Template
		sdwantemplateinputstype["title"] = item.Title
		sdwantemplateinputstype["type"] = item.Type
		sdwantemplateinputstype["value"] = item.Value
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
		snmptrap["class_id"] = item.ClassId
		snmptrap["destination"] = item.Destination
		snmptrap["enabled"] = item.Enabled
		snmptrap["object_type"] = item.ObjectType
		snmptrap["port"] = item.Port
		snmptrap["type"] = item.Type
		snmptrap["user"] = item.User
		snmptrap["nr_version"] = item.Version
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
		auth_password_x := d.Get("snmp_users").([]interface{})
		if len(auth_password_x) > 0 {
			auth_password_y := auth_password_x[0].(map[string]interface{})
			snmpuser["auth_password"] = auth_password_y["auth_password"]
		}
		snmpuser["auth_type"] = item.AuthType
		snmpuser["class_id"] = item.ClassId
		snmpuser["is_auth_password_set"] = item.IsAuthPasswordSet
		snmpuser["is_privacy_password_set"] = item.IsPrivacyPasswordSet
		snmpuser["name"] = item.Name
		snmpuser["object_type"] = item.ObjectType
		privacy_password_x := d.Get("snmp_users").([]interface{})
		if len(privacy_password_x) > 0 {
			privacy_password_y := privacy_password_x[0].(map[string]interface{})
			snmpuser["privacy_password"] = privacy_password_y["privacy_password"]
		}
		snmpuser["privacy_type"] = item.PrivacyType
		snmpuser["security_level"] = item.SecurityLevel
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
		softwarerepositoryconstraintmodels["class_id"] = item.ClassId
		softwarerepositoryconstraintmodels["min_version"] = item.MinVersion
		softwarerepositoryconstraintmodels["name"] = item.Name
		softwarerepositoryconstraintmodels["object_type"] = item.ObjectType
		softwarerepositoryconstraintmodels["platform_regex"] = item.PlatformRegex
		softwarerepositoryconstraintmodels["supported_models"] = item.SupportedModels
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
		storagebaseinitiator["class_id"] = item.ClassId
		storagebaseinitiator["iqn"] = item.Iqn
		storagebaseinitiator["name"] = item.Name
		storagebaseinitiator["object_type"] = item.ObjectType
		storagebaseinitiator["type"] = item.Type
		storagebaseinitiator["wwn"] = item.Wwn
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
		storagelocaldisk["class_id"] = item.ClassId
		storagelocaldisk["object_type"] = item.ObjectType
		storagelocaldisk["slot_number"] = item.SlotNumber
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
		storagenetappexportpolicyrule["class_id"] = item.ClassId
		storagenetappexportpolicyrule["client_match"] = item.ClientMatch
		storagenetappexportpolicyrule["index"] = item.Index
		storagenetappexportpolicyrule["object_type"] = item.ObjectType
		storagenetappexportpolicyrule["ro_rule"] = item.RoRule
		storagenetappexportpolicyrule["rw_rule"] = item.RwRule
		storagenetappexportpolicyrule["super_user"] = item.SuperUser
		storagenetappexportpolicyrule["user"] = item.User
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
		storagepurereplicationblackout["class_id"] = item.ClassId
		storagepurereplicationblackout["end"] = item.End
		storagepurereplicationblackout["object_type"] = item.ObjectType
		storagepurereplicationblackout["start"] = item.Start
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
		storagespangroup["class_id"] = item.ClassId
		storagespangroup["disks"] = (func(p []models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var storagelocaldisks []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				storagelocaldisk := make(map[string]interface{})
				storagelocaldisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				storagelocaldisk["class_id"] = item.ClassId
				storagelocaldisk["object_type"] = item.ObjectType
				storagelocaldisk["slot_number"] = item.SlotNumber
				storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
			}
			return storagelocaldisks
		})(item.GetDisks(), d)
		storagespangroup["object_type"] = item.ObjectType
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
		storagevirtualdriveconfig["access_policy"] = item.AccessPolicy
		storagevirtualdriveconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagevirtualdriveconfig["boot_drive"] = item.BootDrive
		storagevirtualdriveconfig["class_id"] = item.ClassId
		storagevirtualdriveconfig["disk_group_name"] = item.DiskGroupName
		storagevirtualdriveconfig["disk_group_policy"] = item.DiskGroupPolicy
		storagevirtualdriveconfig["drive_cache"] = item.DriveCache
		storagevirtualdriveconfig["expand_to_available"] = item.ExpandToAvailable
		storagevirtualdriveconfig["io_policy"] = item.IoPolicy
		storagevirtualdriveconfig["name"] = item.Name
		storagevirtualdriveconfig["object_type"] = item.ObjectType
		storagevirtualdriveconfig["read_policy"] = item.ReadPolicy
		storagevirtualdriveconfig["size"] = item.Size
		storagevirtualdriveconfig["strip_size"] = item.StripSize
		storagevirtualdriveconfig["vdid"] = item.Vdid
		storagevirtualdriveconfig["write_policy"] = item.WritePolicy
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
		sysloglocalclientbase["class_id"] = item.ClassId
		sysloglocalclientbase["min_severity"] = item.MinSeverity
		sysloglocalclientbase["object_type"] = item.ObjectType
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
		syslogremoteclientbase["class_id"] = item.ClassId
		syslogremoteclientbase["enabled"] = item.Enabled
		syslogremoteclientbase["hostname"] = item.Hostname
		syslogremoteclientbase["min_severity"] = item.MinSeverity
		syslogremoteclientbase["object_type"] = item.ObjectType
		syslogremoteclientbase["port"] = item.Port
		syslogremoteclientbase["protocol"] = item.Protocol
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
		tamaction["affected_object_type"] = item.AffectedObjectType
		tamaction["alert_type"] = item.AlertType
		tamaction["class_id"] = item.ClassId
		tamaction["identifiers"] = (func(p []models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var tamidentifierss []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamidentifiers := make(map[string]interface{})
				tamidentifiers["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamidentifiers["class_id"] = item.ClassId
				tamidentifiers["name"] = item.Name
				tamidentifiers["object_type"] = item.ObjectType
				tamidentifiers["value"] = item.Value
				tamidentifierss = append(tamidentifierss, tamidentifiers)
			}
			return tamidentifierss
		})(item.GetIdentifiers(), d)
		tamaction["name"] = item.Name
		tamaction["object_type"] = item.ObjectType
		tamaction["operation_type"] = item.OperationType
		tamaction["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.ClassId
				tamqueryentry["name"] = item.Name
				tamqueryentry["object_type"] = item.ObjectType
				tamqueryentry["priority"] = item.Priority
				tamqueryentry["query"] = item.Query
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tamaction["type"] = item.Type
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
		tamapidatasource["class_id"] = item.ClassId
		tamapidatasource["mo_type"] = item.MoType
		tamapidatasource["name"] = item.Name
		tamapidatasource["object_type"] = item.ObjectType
		tamapidatasource["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.ClassId
				tamqueryentry["name"] = item.Name
				tamqueryentry["object_type"] = item.ObjectType
				tamqueryentry["priority"] = item.Priority
				tamqueryentry["query"] = item.Query
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tamapidatasource["type"] = item.Type
		tamapidatasources = append(tamapidatasources, tamapidatasource)
	}
	return tamapidatasources
}
func flattenListUcsdConnectorPack(p []models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var ucsdconnectorpacks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ucsdconnectorpack := make(map[string]interface{})
		ucsdconnectorpack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ucsdconnectorpack["class_id"] = item.ClassId
		ucsdconnectorpack["connector_feature"] = item.ConnectorFeature
		ucsdconnectorpack["dependency_names"] = item.DependencyNames
		ucsdconnectorpack["downloaded_version"] = item.DownloadedVersion
		ucsdconnectorpack["name"] = item.Name
		ucsdconnectorpack["object_type"] = item.ObjectType
		ucsdconnectorpack["services"] = item.Services
		ucsdconnectorpack["state"] = item.State
		ucsdconnectorpack["nr_version"] = item.Version
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
		uuidpooluuidblock["class_id"] = item.ClassId
		uuidpooluuidblock["from"] = item.From
		uuidpooluuidblock["object_type"] = item.ObjectType
		uuidpooluuidblock["size"] = item.Size
		uuidpooluuidblock["to"] = item.To
		uuidpooluuidblocks = append(uuidpooluuidblocks, uuidpooluuidblock)
	}
	return uuidpooluuidblocks
}
func flattenListVirtualizationNetworkInterface(p []models.VirtualizationNetworkInterface, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationnetworkinterfaces []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		virtualizationnetworkinterface := make(map[string]interface{})
		virtualizationnetworkinterface["adaptor_type"] = item.AdaptorType
		virtualizationnetworkinterface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		virtualizationnetworkinterface["bridge"] = item.Bridge
		virtualizationnetworkinterface["class_id"] = item.ClassId
		virtualizationnetworkinterface["connect_at_power_on"] = item.ConnectAtPowerOn
		virtualizationnetworkinterface["direct_path_io"] = item.DirectPathIo
		virtualizationnetworkinterface["mac_address"] = item.MacAddress
		virtualizationnetworkinterface["name"] = item.Name
		virtualizationnetworkinterface["object_type"] = item.ObjectType
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
		virtualizationvirtualmachinedisk["bus"] = item.Bus
		virtualizationvirtualmachinedisk["class_id"] = item.ClassId
		virtualizationvirtualmachinedisk["name"] = item.Name
		virtualizationvirtualmachinedisk["object_type"] = item.ObjectType
		virtualizationvirtualmachinedisk["order"] = item.Order
		virtualizationvirtualmachinedisk["type"] = item.Type
		virtualizationvirtualmachinedisk["virtual_disk"] = (func(p models.VirtualizationVirtualDiskConfig, d *schema.ResourceData) []map[string]interface{} {
			var virtualizationvirtualdiskconfigs []map[string]interface{}
			var ret models.VirtualizationVirtualDiskConfig
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			virtualizationvirtualdiskconfig := make(map[string]interface{})
			virtualizationvirtualdiskconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			virtualizationvirtualdiskconfig["capacity"] = item.Capacity
			virtualizationvirtualdiskconfig["class_id"] = item.ClassId
			virtualizationvirtualdiskconfig["mode"] = item.Mode
			virtualizationvirtualdiskconfig["object_type"] = item.ObjectType
			virtualizationvirtualdiskconfig["source_certs"] = item.SourceCerts
			virtualizationvirtualdiskconfig["source_disk_to_clone"] = item.SourceDiskToClone
			virtualizationvirtualdiskconfig["source_file_path"] = item.SourceFilePath

			virtualizationvirtualdiskconfigs = append(virtualizationvirtualdiskconfigs, virtualizationvirtualdiskconfig)
			return virtualizationvirtualdiskconfigs
		})(item.GetVirtualDisk(), d)
		virtualizationvirtualmachinedisk["virtual_disk_reference"] = item.VirtualDiskReference
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
func flattenListVmediaMapping(p []models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var vmediamappings []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		vmediamapping := make(map[string]interface{})
		vmediamapping["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		vmediamapping["authentication_protocol"] = item.AuthenticationProtocol
		vmediamapping["class_id"] = item.ClassId
		vmediamapping["device_type"] = item.DeviceType
		vmediamapping["file_location"] = item.FileLocation
		vmediamapping["host_name"] = item.HostName
		vmediamapping["is_password_set"] = item.IsPasswordSet
		vmediamapping["mount_options"] = item.MountOptions
		vmediamapping["mount_protocol"] = item.MountProtocol
		vmediamapping["object_type"] = item.ObjectType
		password_x := d.Get("mappings").([]interface{})
		if len(password_x) > 0 {
			password_y := password_x[0].(map[string]interface{})
			vmediamapping["password"] = password_y["password"]
		}
		vmediamapping["remote_file"] = item.RemoteFile
		vmediamapping["remote_path"] = item.RemotePath
		vmediamapping["sanitized_file_location"] = item.SanitizedFileLocation
		vmediamapping["username"] = item.Username
		vmediamapping["volume_name"] = item.VolumeName
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
		vnicvifstatus["class_id"] = item.ClassId
		vnicvifstatus["name"] = item.Name
		vnicvifstatus["object_type"] = item.ObjectType
		vnicvifstatus["reason"] = item.Reason
		vnicvifstatus["status"] = item.Status
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
		workflowapi["asset_target_moid"] = item.AssetTargetMoid
		workflowapi["body"] = item.Body
		workflowapi["class_id"] = item.ClassId
		workflowapi["content_type"] = item.ContentType
		workflowapi["description"] = item.Description
		workflowapi["error_content_type"] = item.ErrorContentType
		workflowapi["label"] = item.Label
		workflowapi["name"] = item.Name
		workflowapi["object_type"] = item.ObjectType
		workflowapi["outcomes"] = item.Outcomes
		workflowapi["response_spec"] = item.ResponseSpec
		workflowapi["skip_on_condition"] = item.SkipOnCondition
		workflowapi["start_delay"] = item.StartDelay
		workflowapi["timeout"] = item.Timeout
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
		workflowbasedatatype["class_id"] = item.ClassId
		workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
			var workflowdefaultvalues []map[string]interface{}
			var ret models.WorkflowDefaultValue
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowdefaultvalue := make(map[string]interface{})
			workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowdefaultvalue["class_id"] = item.ClassId
			workflowdefaultvalue["is_value_set"] = item.IsValueSet
			workflowdefaultvalue["object_type"] = item.ObjectType
			workflowdefaultvalue["override"] = item.Override
			workflowdefaultvalue["value"] = item.Value

			workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
			return workflowdefaultvalues
		})(item.GetDefault(), d)
		workflowbasedatatype["description"] = item.Description
		workflowbasedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
			var workflowdisplaymetas []map[string]interface{}
			var ret models.WorkflowDisplayMeta
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowdisplaymeta := make(map[string]interface{})
			workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowdisplaymeta["class_id"] = item.ClassId
			workflowdisplaymeta["inventory_selector"] = item.InventorySelector
			workflowdisplaymeta["object_type"] = item.ObjectType
			workflowdisplaymeta["widget_type"] = item.WidgetType

			workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
			return workflowdisplaymetas
		})(item.GetDisplayMeta(), d)
		workflowbasedatatype["input_parameters"] = item.InputParameters
		workflowbasedatatype["label"] = item.Label
		workflowbasedatatype["name"] = item.Name
		workflowbasedatatype["object_type"] = item.ObjectType
		workflowbasedatatype["required"] = item.Required
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
		workflowdynamicworkflowactiontasklist["action"] = item.Action
		workflowdynamicworkflowactiontasklist["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowdynamicworkflowactiontasklist["class_id"] = item.ClassId
		workflowdynamicworkflowactiontasklist["object_type"] = item.ObjectType
		workflowdynamicworkflowactiontasklist["tasks"] = item.Tasks
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
		workflowmessage["class_id"] = item.ClassId
		workflowmessage["message"] = item.Message
		workflowmessage["object_type"] = item.ObjectType
		workflowmessage["severity"] = item.Severity
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
		workflowparameterset["class_id"] = item.ClassId
		workflowparameterset["condition"] = item.Condition
		workflowparameterset["control_parameter"] = item.ControlParameter
		workflowparameterset["enable_parameters"] = item.EnableParameters
		workflowparameterset["name"] = item.Name
		workflowparameterset["object_type"] = item.ObjectType
		workflowparameterset["value"] = item.Value
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
		workflowrollbacktask["catalog_moid"] = item.CatalogMoid
		workflowrollbacktask["class_id"] = item.ClassId
		workflowrollbacktask["description"] = item.Description
		workflowrollbacktask["input_parameters"] = item.InputParameters
		workflowrollbacktask["name"] = item.Name
		workflowrollbacktask["object_type"] = item.ObjectType
		workflowrollbacktask["task_moid"] = item.TaskMoid
		workflowrollbacktask["nr_version"] = item.Version
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
		workflowrollbackworkflowtask["class_id"] = item.ClassId
		workflowrollbackworkflowtask["description"] = item.Description
		workflowrollbackworkflowtask["name"] = item.Name
		workflowrollbackworkflowtask["object_type"] = item.ObjectType
		workflowrollbackworkflowtask["ref_name"] = item.RefName
		workflowrollbackworkflowtask["rollback_completed"] = item.RollbackCompleted
		workflowrollbackworkflowtask["rollback_task_name"] = item.RollbackTaskName
		workflowrollbackworkflowtask["status"] = item.Status
		workflowrollbackworkflowtask["task_info_moid"] = item.TaskInfoMoid
		workflowrollbackworkflowtask["task_path"] = item.TaskPath
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
		workflowtaskretryinfo["class_id"] = item.ClassId
		workflowtaskretryinfo["object_type"] = item.ObjectType
		workflowtaskretryinfo["status"] = item.Status
		workflowtaskretryinfo["task_inst_id"] = item.TaskInstId
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
		workflowuiinputfilter["class_id"] = item.ClassId
		workflowuiinputfilter["filters"] = item.Filters
		workflowuiinputfilter["name"] = item.Name
		workflowuiinputfilter["object_type"] = item.ObjectType
		workflowuiinputfilter["user_help_message"] = item.UserHelpMessage
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
		workflowworkflowtask["class_id"] = item.ClassId
		workflowworkflowtask["description"] = item.Description
		workflowworkflowtask["label"] = item.Label
		workflowworkflowtask["name"] = item.Name
		workflowworkflowtask["object_type"] = item.ObjectType
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
		x509certificate["class_id"] = item.ClassId
		x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetIssuer(), d)
		x509certificate["not_after"] = item.NotAfter
		x509certificate["not_before"] = item.NotBefore
		x509certificate["object_type"] = item.ObjectType
		x509certificate["pem_certificate"] = item.PemCertificate
		x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
		x509certificate["signature_algorithm"] = item.SignatureAlgorithm
		x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

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
	accessaddresstype["class_id"] = item.ClassId
	accessaddresstype["enable_ip_v4"] = item.EnableIpV4
	accessaddresstype["enable_ip_v6"] = item.EnableIpV6
	accessaddresstype["object_type"] = item.ObjectType

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
	adapterunitrelationship["class_id"] = item.ClassId
	adapterunitrelationship["moid"] = item.Moid
	adapterunitrelationship["object_type"] = item.ObjectType
	adapterunitrelationship["selector"] = item.Selector

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
	adapterunitexpanderrelationship["class_id"] = item.ClassId
	adapterunitexpanderrelationship["moid"] = item.Moid
	adapterunitexpanderrelationship["object_type"] = item.ObjectType
	adapterunitexpanderrelationship["selector"] = item.Selector

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
	appliancecertrenewalphase["class_id"] = item.ClassId
	appliancecertrenewalphase["end_time"] = item.EndTime
	appliancecertrenewalphase["failed"] = item.Failed
	appliancecertrenewalphase["message"] = item.Message
	appliancecertrenewalphase["name"] = item.Name
	appliancecertrenewalphase["object_type"] = item.ObjectType
	appliancecertrenewalphase["start_time"] = item.StartTime

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
	appliancedataexportpolicyrelationship["class_id"] = item.ClassId
	appliancedataexportpolicyrelationship["moid"] = item.Moid
	appliancedataexportpolicyrelationship["object_type"] = item.ObjectType
	appliancedataexportpolicyrelationship["selector"] = item.Selector

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
	appliancegroupstatusrelationship["class_id"] = item.ClassId
	appliancegroupstatusrelationship["moid"] = item.Moid
	appliancegroupstatusrelationship["object_type"] = item.ObjectType
	appliancegroupstatusrelationship["selector"] = item.Selector

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
	applianceimagebundlerelationship["class_id"] = item.ClassId
	applianceimagebundlerelationship["moid"] = item.Moid
	applianceimagebundlerelationship["object_type"] = item.ObjectType
	applianceimagebundlerelationship["selector"] = item.Selector

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
	appliancenodeinforelationship["class_id"] = item.ClassId
	appliancenodeinforelationship["moid"] = item.Moid
	appliancenodeinforelationship["object_type"] = item.ObjectType
	appliancenodeinforelationship["selector"] = item.Selector

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
	appliancenodestatusrelationship["class_id"] = item.ClassId
	appliancenodestatusrelationship["moid"] = item.Moid
	appliancenodestatusrelationship["object_type"] = item.ObjectType
	appliancenodestatusrelationship["selector"] = item.Selector

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
	appliancesysteminforelationship["class_id"] = item.ClassId
	appliancesysteminforelationship["moid"] = item.Moid
	appliancesysteminforelationship["object_type"] = item.ObjectType
	appliancesysteminforelationship["selector"] = item.Selector

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
	appliancesystemstatusrelationship["class_id"] = item.ClassId
	appliancesystemstatusrelationship["moid"] = item.Moid
	appliancesystemstatusrelationship["object_type"] = item.ObjectType
	appliancesystemstatusrelationship["selector"] = item.Selector

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
	assetclaimsignature["class_id"] = item.ClassId
	assetclaimsignature["object_type"] = item.ObjectType
	assetclaimsignature["time_stamp"] = item.TimeStamp

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
	assetclustermemberrelationship["class_id"] = item.ClassId
	assetclustermemberrelationship["moid"] = item.Moid
	assetclustermemberrelationship["object_type"] = item.ObjectType
	assetclustermemberrelationship["selector"] = item.Selector

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
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["address3"] = item.Address3
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["county"] = item.County
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["province"] = item.Province
		assetaddressinformation["state"] = item.State

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
		assetglobalultimate["class_id"] = item.ClassId
		assetglobalultimate["id"] = item.Id
		assetglobalultimate["name"] = item.Name
		assetglobalultimate["object_type"] = item.ObjectType

		assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
		return assetglobalultimates
	})(item.GetBillToGlobalUltimate(), d)
	assetcontractinformation["class_id"] = item.ClassId
	assetcontractinformation["contract_number"] = item.ContractNumber
	assetcontractinformation["line_status"] = item.LineStatus
	assetcontractinformation["object_type"] = item.ObjectType

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
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["address3"] = item.Address3
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["county"] = item.County
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["province"] = item.Province
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetAddress(), d)
	assetcustomerinformation["class_id"] = item.ClassId
	assetcustomerinformation["id"] = item.Id
	assetcustomerinformation["name"] = item.Name
	assetcustomerinformation["object_type"] = item.ObjectType

	assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
	return assetcustomerinformations
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
	assetdeviceclaimrelationship["class_id"] = item.ClassId
	assetdeviceclaimrelationship["moid"] = item.Moid
	assetdeviceclaimrelationship["object_type"] = item.ObjectType
	assetdeviceclaimrelationship["selector"] = item.Selector

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
	assetdeviceconfigurationrelationship["class_id"] = item.ClassId
	assetdeviceconfigurationrelationship["moid"] = item.Moid
	assetdeviceconfigurationrelationship["object_type"] = item.ObjectType
	assetdeviceconfigurationrelationship["selector"] = item.Selector

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
	assetdeviceconnectionrelationship["class_id"] = item.ClassId
	assetdeviceconnectionrelationship["moid"] = item.Moid
	assetdeviceconnectionrelationship["object_type"] = item.ObjectType
	assetdeviceconnectionrelationship["selector"] = item.Selector

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
	assetdevicecontractinformationrelationship["class_id"] = item.ClassId
	assetdevicecontractinformationrelationship["moid"] = item.Moid
	assetdevicecontractinformationrelationship["object_type"] = item.ObjectType
	assetdevicecontractinformationrelationship["selector"] = item.Selector

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
	assetdeviceinformation["application_name"] = item.ApplicationName
	assetdeviceinformation["class_id"] = item.ClassId
	assetdeviceinformation["device_transactions"] = (func(p []models.AssetDeviceTransaction, d *schema.ResourceData) []map[string]interface{} {
		var assetdevicetransactions []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			assetdevicetransaction := make(map[string]interface{})
			assetdevicetransaction["action"] = item.Action
			assetdevicetransaction["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetdevicetransaction["class_id"] = item.ClassId
			assetdevicetransaction["object_type"] = item.ObjectType
			assetdevicetransaction["status_description"] = item.StatusDescription
			assetdevicetransaction["status_id"] = item.StatusId
			assetdevicetransaction["timestamp"] = item.Timestamp
			assetdevicetransaction["transaction_batch_id"] = item.TransactionBatchId
			assetdevicetransaction["transaction_date"] = item.TransactionDate
			assetdevicetransaction["transaction_sequence"] = item.TransactionSequence
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
			assetaddressinformation["address1"] = item.Address1
			assetaddressinformation["address2"] = item.Address2
			assetaddressinformation["address3"] = item.Address3
			assetaddressinformation["city"] = item.City
			assetaddressinformation["class_id"] = item.ClassId
			assetaddressinformation["country"] = item.Country
			assetaddressinformation["county"] = item.County
			assetaddressinformation["location"] = item.Location
			assetaddressinformation["name"] = item.Name
			assetaddressinformation["object_type"] = item.ObjectType
			assetaddressinformation["postal_code"] = item.PostalCode
			assetaddressinformation["province"] = item.Province
			assetaddressinformation["state"] = item.State

			assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
			return assetaddressinformations
		})(item.GetAddress(), d)
		assetcustomerinformation["class_id"] = item.ClassId
		assetcustomerinformation["id"] = item.Id
		assetcustomerinformation["name"] = item.Name
		assetcustomerinformation["object_type"] = item.ObjectType

		assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
		return assetcustomerinformations
	})(item.GetEndCustomer(), d)
	assetdeviceinformation["instance_id"] = item.InstanceId
	assetdeviceinformation["item_type"] = item.ItemType
	assetdeviceinformation["mlb_offer_type"] = item.MlbOfferType
	assetdeviceinformation["mlb_product_id"] = item.MlbProductId
	assetdeviceinformation["mlb_product_name"] = item.MlbProductName
	assetdeviceinformation["object_type"] = item.ObjectType
	assetdeviceinformation["old_device_id"] = item.OldDeviceId
	assetdeviceinformation["old_device_status_description"] = item.OldDeviceStatusDescription
	assetdeviceinformation["old_device_status_id"] = item.OldDeviceStatusId
	assetdeviceinformation["old_instance_id"] = item.OldInstanceId
	assetdeviceinformation["product_family"] = item.ProductFamily
	assetdeviceinformation["product_type"] = item.ProductType
	assetdeviceinformation["unit_of_measure"] = item.UnitOfMeasure

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
	assetdeviceregistrationrelationship["class_id"] = item.ClassId
	assetdeviceregistrationrelationship["moid"] = item.Moid
	assetdeviceregistrationrelationship["object_type"] = item.ObjectType
	assetdeviceregistrationrelationship["selector"] = item.Selector

	assetdeviceregistrationrelationships = append(assetdeviceregistrationrelationships, assetdeviceregistrationrelationship)
	return assetdeviceregistrationrelationships
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
	assetglobalultimate["class_id"] = item.ClassId
	assetglobalultimate["id"] = item.Id
	assetglobalultimate["name"] = item.Name
	assetglobalultimate["object_type"] = item.ObjectType

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
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["address3"] = item.Address3
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["county"] = item.County
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["province"] = item.Province
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetBillTo(), d)
	assetproductinformation["class_id"] = item.ClassId
	assetproductinformation["description"] = item.Description
	assetproductinformation["family"] = item.Family
	assetproductinformation["group"] = item.Group
	assetproductinformation["number"] = item.Number
	assetproductinformation["object_type"] = item.ObjectType
	assetproductinformation["ship_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["address3"] = item.Address3
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["county"] = item.County
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["province"] = item.Province
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetShipTo(), d)
	assetproductinformation["sub_type"] = item.SubType

	assetproductinformations = append(assetproductinformations, assetproductinformation)
	return assetproductinformations
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
	assetsudiinfo["class_id"] = item.ClassId
	assetsudiinfo["object_type"] = item.ObjectType
	assetsudiinfo["pid"] = item.Pid
	assetsudiinfo["serial_number"] = item.SerialNumber
	assetsudiinfo["signature"] = item.Signature
	assetsudiinfo["status"] = item.Status
	assetsudiinfo["sudi_certificate"] = (func(p models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
		var x509certificates []map[string]interface{}
		var ret models.X509Certificate
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		x509certificate := make(map[string]interface{})
		x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		x509certificate["class_id"] = item.ClassId
		x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetIssuer(), d)
		x509certificate["not_after"] = item.NotAfter
		x509certificate["not_before"] = item.NotBefore
		x509certificate["object_type"] = item.ObjectType
		x509certificate["pem_certificate"] = item.PemCertificate
		x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
		x509certificate["signature_algorithm"] = item.SignatureAlgorithm
		x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

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
	assettargetrelationship["class_id"] = item.ClassId
	assettargetrelationship["moid"] = item.Moid
	assettargetrelationship["object_type"] = item.ObjectType
	assettargetrelationship["selector"] = item.Selector

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
	biosbootmoderelationship["class_id"] = item.ClassId
	biosbootmoderelationship["moid"] = item.Moid
	biosbootmoderelationship["object_type"] = item.ObjectType
	biosbootmoderelationship["selector"] = item.Selector

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
	biossystembootorderrelationship["class_id"] = item.ClassId
	biossystembootorderrelationship["moid"] = item.Moid
	biossystembootorderrelationship["object_type"] = item.ObjectType
	biossystembootorderrelationship["selector"] = item.Selector

	biossystembootorderrelationships = append(biossystembootorderrelationships, biossystembootorderrelationship)
	return biossystembootorderrelationships
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
	biosunitrelationship["class_id"] = item.ClassId
	biosunitrelationship["moid"] = item.Moid
	biosunitrelationship["object_type"] = item.ObjectType
	biosunitrelationship["selector"] = item.Selector

	biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	return biosunitrelationships
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
	bootdevicebootmoderelationship["class_id"] = item.ClassId
	bootdevicebootmoderelationship["moid"] = item.Moid
	bootdevicebootmoderelationship["object_type"] = item.ObjectType
	bootdevicebootmoderelationship["selector"] = item.Selector

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
	bootdevicebootsecurityrelationship["class_id"] = item.ClassId
	bootdevicebootsecurityrelationship["moid"] = item.Moid
	bootdevicebootsecurityrelationship["object_type"] = item.ObjectType
	bootdevicebootsecurityrelationship["selector"] = item.Selector

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
	capabilityswitchnetworklimits["class_id"] = item.ClassId
	capabilityswitchnetworklimits["max_compressed_port_vlan_count"] = item.MaxCompressedPortVlanCount
	capabilityswitchnetworklimits["max_uncompressed_port_vlan_count"] = item.MaxUncompressedPortVlanCount
	capabilityswitchnetworklimits["maximum_active_traffic_monitoring_sessions"] = item.MaximumActiveTrafficMonitoringSessions
	capabilityswitchnetworklimits["maximum_ethernet_port_channels"] = item.MaximumEthernetPortChannels
	capabilityswitchnetworklimits["maximum_ethernet_uplink_ports"] = item.MaximumEthernetUplinkPorts
	capabilityswitchnetworklimits["maximum_fc_port_channel_members"] = item.MaximumFcPortChannelMembers
	capabilityswitchnetworklimits["maximum_fc_port_channels"] = item.MaximumFcPortChannels
	capabilityswitchnetworklimits["maximum_igmp_groups"] = item.MaximumIgmpGroups
	capabilityswitchnetworklimits["maximum_port_channel_members"] = item.MaximumPortChannelMembers
	capabilityswitchnetworklimits["maximum_primary_vlan"] = item.MaximumPrimaryVlan
	capabilityswitchnetworklimits["maximum_secondary_vlan"] = item.MaximumSecondaryVlan
	capabilityswitchnetworklimits["maximum_secondary_vlan_per_primary"] = item.MaximumSecondaryVlanPerPrimary
	capabilityswitchnetworklimits["maximum_vifs"] = item.MaximumVifs
	capabilityswitchnetworklimits["maximum_vlans"] = item.MaximumVlans
	capabilityswitchnetworklimits["minimum_active_fans"] = item.MinimumActiveFans
	capabilityswitchnetworklimits["object_type"] = item.ObjectType

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
	capabilityswitchstoragelimits["class_id"] = item.ClassId
	capabilityswitchstoragelimits["maximum_user_zone_count"] = item.MaximumUserZoneCount
	capabilityswitchstoragelimits["maximum_virtual_fc_interfaces"] = item.MaximumVirtualFcInterfaces
	capabilityswitchstoragelimits["maximum_virtual_fc_interfaces_per_blade_server"] = item.MaximumVirtualFcInterfacesPerBladeServer
	capabilityswitchstoragelimits["maximum_vsans"] = item.MaximumVsans
	capabilityswitchstoragelimits["maximum_zone_count"] = item.MaximumZoneCount
	capabilityswitchstoragelimits["object_type"] = item.ObjectType

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
	capabilityswitchsystemlimits["class_id"] = item.ClassId
	capabilityswitchsystemlimits["maximum_chassis_count"] = item.MaximumChassisCount
	capabilityswitchsystemlimits["maximum_fex_per_domain"] = item.MaximumFexPerDomain
	capabilityswitchsystemlimits["maximum_servers_per_domain"] = item.MaximumServersPerDomain
	capabilityswitchsystemlimits["object_type"] = item.ObjectType

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
	chassisconfigresultrelationship["class_id"] = item.ClassId
	chassisconfigresultrelationship["moid"] = item.Moid
	chassisconfigresultrelationship["object_type"] = item.ObjectType
	chassisconfigresultrelationship["selector"] = item.Selector

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
	chassisprofilerelationship["class_id"] = item.ClassId
	chassisprofilerelationship["moid"] = item.Moid
	chassisprofilerelationship["object_type"] = item.ObjectType
	chassisprofilerelationship["selector"] = item.Selector

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
	commhttpproxypolicyrelationship["class_id"] = item.ClassId
	commhttpproxypolicyrelationship["moid"] = item.Moid
	commhttpproxypolicyrelationship["object_type"] = item.ObjectType
	commhttpproxypolicyrelationship["selector"] = item.Selector

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
	commipv4interface["class_id"] = item.ClassId
	commipv4interface["gateway"] = item.Gateway
	commipv4interface["ip_address"] = item.IpAddress
	commipv4interface["netmask"] = item.Netmask
	commipv4interface["object_type"] = item.ObjectType

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
	commipv6interface["class_id"] = item.ClassId
	commipv6interface["gateway"] = item.Gateway
	commipv6interface["ip_address"] = item.IpAddress
	commipv6interface["object_type"] = item.ObjectType
	commipv6interface["prefix"] = item.Prefix

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
	computealarmsummary["class_id"] = item.ClassId
	computealarmsummary["critical"] = item.Critical
	computealarmsummary["object_type"] = item.ObjectType
	computealarmsummary["warning"] = item.Warning

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
	computebladerelationship["class_id"] = item.ClassId
	computebladerelationship["moid"] = item.Moid
	computebladerelationship["object_type"] = item.ObjectType
	computebladerelationship["selector"] = item.Selector

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
	computeboardrelationship["class_id"] = item.ClassId
	computeboardrelationship["moid"] = item.Moid
	computeboardrelationship["object_type"] = item.ObjectType
	computeboardrelationship["selector"] = item.Selector

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
	computepersistentmemoryoperation["admin_action"] = item.AdminAction
	computepersistentmemoryoperation["class_id"] = item.ClassId
	computepersistentmemoryoperation["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	computepersistentmemoryoperation["modules"] = (func(p []models.ComputePersistentMemoryModule, d *schema.ResourceData) []map[string]interface{} {
		var computepersistentmemorymodules []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computepersistentmemorymodule := make(map[string]interface{})
			computepersistentmemorymodule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computepersistentmemorymodule["class_id"] = item.ClassId
			computepersistentmemorymodule["object_type"] = item.ObjectType
			computepersistentmemorymodule["socket_id"] = item.SocketId
			computepersistentmemorymodule["socket_memory_id"] = item.SocketMemoryId
			computepersistentmemorymodules = append(computepersistentmemorymodules, computepersistentmemorymodule)
		}
		return computepersistentmemorymodules
	})(item.GetModules(), d)
	computepersistentmemoryoperation["object_type"] = item.ObjectType
	secure_passphrase_x := d.Get("persistent_memory_operation").([]interface{})
	if len(secure_passphrase_x) > 0 {
		secure_passphrase_y := secure_passphrase_x[0].(map[string]interface{})
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
	computephysicalrelationship["class_id"] = item.ClassId
	computephysicalrelationship["moid"] = item.Moid
	computephysicalrelationship["object_type"] = item.ObjectType
	computephysicalrelationship["selector"] = item.Selector

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
	computerackunitrelationship["class_id"] = item.ClassId
	computerackunitrelationship["moid"] = item.Moid
	computerackunitrelationship["object_type"] = item.ObjectType
	computerackunitrelationship["selector"] = item.Selector

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
	computeserverconfig["asset_tag"] = item.AssetTag
	computeserverconfig["class_id"] = item.ClassId
	computeserverconfig["object_type"] = item.ObjectType
	computeserverconfig["user_label"] = item.UserLabel

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
	computestoragecontrolleroperation["admin_action"] = item.AdminAction
	computestoragecontrolleroperation["class_id"] = item.ClassId
	computestoragecontrolleroperation["controller_id"] = item.ControllerId
	computestoragecontrolleroperation["object_type"] = item.ObjectType

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
	computestoragephysicaldriveoperation["admin_action"] = item.AdminAction
	computestoragephysicaldriveoperation["class_id"] = item.ClassId
	computestoragephysicaldriveoperation["controller_id"] = item.ControllerId
	computestoragephysicaldriveoperation["object_type"] = item.ObjectType
	computestoragephysicaldriveoperation["physical_drives"] = (func(p []models.ComputeStoragePhysicalDrive, d *schema.ResourceData) []map[string]interface{} {
		var computestoragephysicaldrives []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computestoragephysicaldrive := make(map[string]interface{})
			computestoragephysicaldrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computestoragephysicaldrive["class_id"] = item.ClassId
			computestoragephysicaldrive["object_type"] = item.ObjectType
			computestoragephysicaldrive["slot_number"] = item.SlotNumber
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
	computestoragevirtualdriveoperation["admin_action"] = item.AdminAction
	computestoragevirtualdriveoperation["class_id"] = item.ClassId
	computestoragevirtualdriveoperation["controller_id"] = item.ControllerId
	computestoragevirtualdriveoperation["object_type"] = item.ObjectType
	computestoragevirtualdriveoperation["virtual_drives"] = (func(p []models.ComputeStorageVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
		var computestoragevirtualdrives []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computestoragevirtualdrive := make(map[string]interface{})
			computestoragevirtualdrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computestoragevirtualdrive["class_id"] = item.ClassId
			computestoragevirtualdrive["id"] = item.Id
			computestoragevirtualdrive["object_type"] = item.ObjectType
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
	computevmediarelationship["class_id"] = item.ClassId
	computevmediarelationship["moid"] = item.Moid
	computevmediarelationship["object_type"] = item.ObjectType
	computevmediarelationship["selector"] = item.Selector

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
	condalarmsummary["class_id"] = item.ClassId
	condalarmsummary["critical"] = item.Critical
	condalarmsummary["object_type"] = item.ObjectType
	condalarmsummary["warning"] = item.Warning

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
	condhclstatusrelationship["class_id"] = item.ClassId
	condhclstatusrelationship["moid"] = item.Moid
	condhclstatusrelationship["object_type"] = item.ObjectType
	condhclstatusrelationship["selector"] = item.Selector

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
	configexporteditemrelationship["class_id"] = item.ClassId
	configexporteditemrelationship["moid"] = item.Moid
	configexporteditemrelationship["object_type"] = item.ObjectType
	configexporteditemrelationship["selector"] = item.Selector

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
	configexporterrelationship["class_id"] = item.ClassId
	configexporterrelationship["moid"] = item.Moid
	configexporterrelationship["object_type"] = item.ObjectType
	configexporterrelationship["selector"] = item.Selector

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
	configimporterrelationship["class_id"] = item.ClassId
	configimporterrelationship["moid"] = item.Moid
	configimporterrelationship["object_type"] = item.ObjectType
	configimporterrelationship["selector"] = item.Selector

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
	configmoref["class_id"] = item.ClassId
	configmoref["moid"] = item.Moid
	configmoref["object_type"] = item.ObjectType

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
	connectorfilechecksum["class_id"] = item.ClassId
	connectorfilechecksum["hash_algorithm"] = item.HashAlgorithm
	connectorfilechecksum["object_type"] = item.ObjectType

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
	connectorplatformparambase["class_id"] = item.ClassId
	connectorplatformparambase["object_type"] = item.ObjectType

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
	equipmentbaserelationship["class_id"] = item.ClassId
	equipmentbaserelationship["moid"] = item.Moid
	equipmentbaserelationship["object_type"] = item.ObjectType
	equipmentbaserelationship["selector"] = item.Selector

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
	equipmentchassisrelationship["class_id"] = item.ClassId
	equipmentchassisrelationship["moid"] = item.Moid
	equipmentchassisrelationship["object_type"] = item.ObjectType
	equipmentchassisrelationship["selector"] = item.Selector

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
	equipmentfanmodulerelationship["class_id"] = item.ClassId
	equipmentfanmodulerelationship["moid"] = item.Moid
	equipmentfanmodulerelationship["object_type"] = item.ObjectType
	equipmentfanmodulerelationship["selector"] = item.Selector

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
	equipmentfexrelationship["class_id"] = item.ClassId
	equipmentfexrelationship["moid"] = item.Moid
	equipmentfexrelationship["object_type"] = item.ObjectType
	equipmentfexrelationship["selector"] = item.Selector

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
	equipmentfrurelationship["class_id"] = item.ClassId
	equipmentfrurelationship["moid"] = item.Moid
	equipmentfrurelationship["object_type"] = item.ObjectType
	equipmentfrurelationship["selector"] = item.Selector

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
	equipmentiocardrelationship["class_id"] = item.ClassId
	equipmentiocardrelationship["moid"] = item.Moid
	equipmentiocardrelationship["object_type"] = item.ObjectType
	equipmentiocardrelationship["selector"] = item.Selector

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
	equipmentiocardbaserelationship["class_id"] = item.ClassId
	equipmentiocardbaserelationship["moid"] = item.Moid
	equipmentiocardbaserelationship["object_type"] = item.ObjectType
	equipmentiocardbaserelationship["selector"] = item.Selector

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
	equipmentlocatorledrelationship["class_id"] = item.ClassId
	equipmentlocatorledrelationship["moid"] = item.Moid
	equipmentlocatorledrelationship["object_type"] = item.ObjectType
	equipmentlocatorledrelationship["selector"] = item.Selector

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
	equipmentphysicalidentityrelationship["class_id"] = item.ClassId
	equipmentphysicalidentityrelationship["moid"] = item.Moid
	equipmentphysicalidentityrelationship["object_type"] = item.ObjectType
	equipmentphysicalidentityrelationship["selector"] = item.Selector

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
	equipmentpsucontrolrelationship["class_id"] = item.ClassId
	equipmentpsucontrolrelationship["moid"] = item.Moid
	equipmentpsucontrolrelationship["object_type"] = item.ObjectType
	equipmentpsucontrolrelationship["selector"] = item.Selector

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
	equipmentrackenclosurerelationship["class_id"] = item.ClassId
	equipmentrackenclosurerelationship["moid"] = item.Moid
	equipmentrackenclosurerelationship["object_type"] = item.ObjectType
	equipmentrackenclosurerelationship["selector"] = item.Selector

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
	equipmentrackenclosureslotrelationship["class_id"] = item.ClassId
	equipmentrackenclosureslotrelationship["moid"] = item.Moid
	equipmentrackenclosureslotrelationship["object_type"] = item.ObjectType
	equipmentrackenclosureslotrelationship["selector"] = item.Selector

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
	equipmentsharediomodulerelationship["class_id"] = item.ClassId
	equipmentsharediomodulerelationship["moid"] = item.Moid
	equipmentsharediomodulerelationship["object_type"] = item.ObjectType
	equipmentsharediomodulerelationship["selector"] = item.Selector

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
	equipmentswitchcardrelationship["class_id"] = item.ClassId
	equipmentswitchcardrelationship["moid"] = item.Moid
	equipmentswitchcardrelationship["object_type"] = item.ObjectType
	equipmentswitchcardrelationship["selector"] = item.Selector

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
	equipmentsystemiocontrollerrelationship["class_id"] = item.ClassId
	equipmentsystemiocontrollerrelationship["moid"] = item.Moid
	equipmentsystemiocontrollerrelationship["object_type"] = item.ObjectType
	equipmentsystemiocontrollerrelationship["selector"] = item.Selector

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
	etherphysicalportrelationship["class_id"] = item.ClassId
	etherphysicalportrelationship["moid"] = item.Moid
	etherphysicalportrelationship["object_type"] = item.ObjectType
	etherphysicalportrelationship["selector"] = item.Selector

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
	etherphysicalportbaserelationship["class_id"] = item.ClassId
	etherphysicalportbaserelationship["moid"] = item.Moid
	etherphysicalportbaserelationship["object_type"] = item.ObjectType
	etherphysicalportbaserelationship["selector"] = item.Selector

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
	fabricconfigresultrelationship["class_id"] = item.ClassId
	fabricconfigresultrelationship["moid"] = item.Moid
	fabricconfigresultrelationship["object_type"] = item.ObjectType
	fabricconfigresultrelationship["selector"] = item.Selector

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
	fabricethnetworkcontrolpolicyrelationship["class_id"] = item.ClassId
	fabricethnetworkcontrolpolicyrelationship["moid"] = item.Moid
	fabricethnetworkcontrolpolicyrelationship["object_type"] = item.ObjectType
	fabricethnetworkcontrolpolicyrelationship["selector"] = item.Selector

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
	fabricethnetworkgrouppolicyrelationship["class_id"] = item.ClassId
	fabricethnetworkgrouppolicyrelationship["moid"] = item.Moid
	fabricethnetworkgrouppolicyrelationship["object_type"] = item.ObjectType
	fabricethnetworkgrouppolicyrelationship["selector"] = item.Selector

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
	fabricethnetworkpolicyrelationship["class_id"] = item.ClassId
	fabricethnetworkpolicyrelationship["moid"] = item.Moid
	fabricethnetworkpolicyrelationship["object_type"] = item.ObjectType
	fabricethnetworkpolicyrelationship["selector"] = item.Selector

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
	fabricfcnetworkpolicyrelationship["class_id"] = item.ClassId
	fabricfcnetworkpolicyrelationship["moid"] = item.Moid
	fabricfcnetworkpolicyrelationship["object_type"] = item.ObjectType
	fabricfcnetworkpolicyrelationship["selector"] = item.Selector

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
	fabricflowcontrolpolicyrelationship["class_id"] = item.ClassId
	fabricflowcontrolpolicyrelationship["moid"] = item.Moid
	fabricflowcontrolpolicyrelationship["object_type"] = item.ObjectType
	fabricflowcontrolpolicyrelationship["selector"] = item.Selector

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
	fabriclinkaggregationpolicyrelationship["class_id"] = item.ClassId
	fabriclinkaggregationpolicyrelationship["moid"] = item.Moid
	fabriclinkaggregationpolicyrelationship["object_type"] = item.ObjectType
	fabriclinkaggregationpolicyrelationship["selector"] = item.Selector

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
	fabriclinkcontrolpolicyrelationship["class_id"] = item.ClassId
	fabriclinkcontrolpolicyrelationship["moid"] = item.Moid
	fabriclinkcontrolpolicyrelationship["object_type"] = item.ObjectType
	fabriclinkcontrolpolicyrelationship["selector"] = item.Selector

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
	fabriclldpsettings["class_id"] = item.ClassId
	fabriclldpsettings["object_type"] = item.ObjectType
	fabriclldpsettings["receive_enabled"] = item.ReceiveEnabled
	fabriclldpsettings["transmit_enabled"] = item.TransmitEnabled

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
	fabricmacagingsettings["class_id"] = item.ClassId
	fabricmacagingsettings["mac_aging_option"] = item.MacAgingOption
	fabricmacagingsettings["mac_aging_time"] = item.MacAgingTime
	fabricmacagingsettings["object_type"] = item.ObjectType

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
	fabricmulticastpolicyrelationship["class_id"] = item.ClassId
	fabricmulticastpolicyrelationship["moid"] = item.Moid
	fabricmulticastpolicyrelationship["object_type"] = item.ObjectType
	fabricmulticastpolicyrelationship["selector"] = item.Selector

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
	fabricportpolicyrelationship["class_id"] = item.ClassId
	fabricportpolicyrelationship["moid"] = item.Moid
	fabricportpolicyrelationship["object_type"] = item.ObjectType
	fabricportpolicyrelationship["selector"] = item.Selector

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
	fabricswitchclusterprofilerelationship["class_id"] = item.ClassId
	fabricswitchclusterprofilerelationship["moid"] = item.Moid
	fabricswitchclusterprofilerelationship["object_type"] = item.ObjectType
	fabricswitchclusterprofilerelationship["selector"] = item.Selector

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
	fabricswitchprofilerelationship["class_id"] = item.ClassId
	fabricswitchprofilerelationship["moid"] = item.Moid
	fabricswitchprofilerelationship["object_type"] = item.ObjectType
	fabricswitchprofilerelationship["selector"] = item.Selector

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
	fabricudldglobalsettings["class_id"] = item.ClassId
	fabricudldglobalsettings["message_interval"] = item.MessageInterval
	fabricudldglobalsettings["object_type"] = item.ObjectType
	fabricudldglobalsettings["recovery_action"] = item.RecoveryAction

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
	fabricudldsettings["admin_state"] = item.AdminState
	fabricudldsettings["class_id"] = item.ClassId
	fabricudldsettings["mode"] = item.Mode
	fabricudldsettings["object_type"] = item.ObjectType

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
	fabricvlansettings["allowed_vlans"] = item.AllowedVlans
	fabricvlansettings["class_id"] = item.ClassId
	fabricvlansettings["native_vlan"] = item.NativeVlan
	fabricvlansettings["object_type"] = item.ObjectType

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
	fcpoolblock["class_id"] = item.ClassId
	fcpoolblock["from"] = item.From
	fcpoolblock["object_type"] = item.ObjectType
	fcpoolblock["size"] = item.Size
	fcpoolblock["to"] = item.To

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
	fcpoolfcblockrelationship["class_id"] = item.ClassId
	fcpoolfcblockrelationship["moid"] = item.Moid
	fcpoolfcblockrelationship["object_type"] = item.ObjectType
	fcpoolfcblockrelationship["selector"] = item.Selector

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
	fcpoolleaserelationship["class_id"] = item.ClassId
	fcpoolleaserelationship["moid"] = item.Moid
	fcpoolleaserelationship["object_type"] = item.ObjectType
	fcpoolleaserelationship["selector"] = item.Selector

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
	fcpoolpoolrelationship["class_id"] = item.ClassId
	fcpoolpoolrelationship["moid"] = item.Moid
	fcpoolpoolrelationship["object_type"] = item.ObjectType
	fcpoolpoolrelationship["selector"] = item.Selector

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
	fcpoolpoolmemberrelationship["class_id"] = item.ClassId
	fcpoolpoolmemberrelationship["moid"] = item.Moid
	fcpoolpoolmemberrelationship["object_type"] = item.ObjectType
	fcpoolpoolmemberrelationship["selector"] = item.Selector

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
	fcpooluniverserelationship["class_id"] = item.ClassId
	fcpooluniverserelationship["moid"] = item.Moid
	fcpooluniverserelationship["object_type"] = item.ObjectType
	fcpooluniverserelationship["selector"] = item.Selector

	fcpooluniverserelationships = append(fcpooluniverserelationships, fcpooluniverserelationship)
	return fcpooluniverserelationships
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
	firmwaredirectdownload["class_id"] = item.ClassId
	firmwaredirectdownload["http_server"] = (func(p models.FirmwareHttpServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		var ret models.FirmwareHttpServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarehttpserver := make(map[string]interface{})
		firmwarehttpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarehttpserver["class_id"] = item.ClassId
		firmwarehttpserver["location_link"] = item.LocationLink
		firmwarehttpserver["mount_options"] = item.MountOptions
		firmwarehttpserver["object_type"] = item.ObjectType

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.GetHttpServer(), d)
	firmwaredirectdownload["image_source"] = item.ImageSource
	firmwaredirectdownload["is_password_set"] = item.IsPasswordSet
	firmwaredirectdownload["object_type"] = item.ObjectType
	password_x := d.Get("direct_download").([]interface{})
	if len(password_x) > 0 {
		password_y := password_x[0].(map[string]interface{})
		firmwaredirectdownload["password"] = password_y["password"]
	}
	firmwaredirectdownload["upgradeoption"] = item.Upgradeoption
	firmwaredirectdownload["username"] = item.Username

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
	firmwaredistributablerelationship["class_id"] = item.ClassId
	firmwaredistributablerelationship["moid"] = item.Moid
	firmwaredistributablerelationship["object_type"] = item.ObjectType
	firmwaredistributablerelationship["selector"] = item.Selector

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
		firmwarecifsserver["class_id"] = item.ClassId
		firmwarecifsserver["file_location"] = item.FileLocation
		firmwarecifsserver["mount_options"] = item.MountOptions
		firmwarecifsserver["object_type"] = item.ObjectType
		firmwarecifsserver["remote_file"] = item.RemoteFile
		firmwarecifsserver["remote_ip"] = item.RemoteIp
		firmwarecifsserver["remote_share"] = item.RemoteShare

		firmwarecifsservers = append(firmwarecifsservers, firmwarecifsserver)
		return firmwarecifsservers
	})(item.GetCifsServer(), d)
	firmwarenetworkshare["class_id"] = item.ClassId
	firmwarenetworkshare["http_server"] = (func(p models.FirmwareHttpServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		var ret models.FirmwareHttpServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarehttpserver := make(map[string]interface{})
		firmwarehttpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarehttpserver["class_id"] = item.ClassId
		firmwarehttpserver["location_link"] = item.LocationLink
		firmwarehttpserver["mount_options"] = item.MountOptions
		firmwarehttpserver["object_type"] = item.ObjectType

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.GetHttpServer(), d)
	firmwarenetworkshare["is_password_set"] = item.IsPasswordSet
	firmwarenetworkshare["map_type"] = item.MapType
	firmwarenetworkshare["nfs_server"] = (func(p models.FirmwareNfsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarenfsservers []map[string]interface{}
		var ret models.FirmwareNfsServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarenfsserver := make(map[string]interface{})
		firmwarenfsserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarenfsserver["class_id"] = item.ClassId
		firmwarenfsserver["file_location"] = item.FileLocation
		firmwarenfsserver["mount_options"] = item.MountOptions
		firmwarenfsserver["object_type"] = item.ObjectType
		firmwarenfsserver["remote_file"] = item.RemoteFile
		firmwarenfsserver["remote_ip"] = item.RemoteIp
		firmwarenfsserver["remote_share"] = item.RemoteShare

		firmwarenfsservers = append(firmwarenfsservers, firmwarenfsserver)
		return firmwarenfsservers
	})(item.GetNfsServer(), d)
	firmwarenetworkshare["object_type"] = item.ObjectType
	password_x := d.Get("network_share").([]interface{})
	if len(password_x) > 0 {
		password_y := password_x[0].(map[string]interface{})
		firmwarenetworkshare["password"] = password_y["password"]
	}
	firmwarenetworkshare["upgradeoption"] = item.Upgradeoption
	firmwarenetworkshare["username"] = item.Username

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
	firmwarerunningfirmwarerelationship["class_id"] = item.ClassId
	firmwarerunningfirmwarerelationship["moid"] = item.Moid
	firmwarerunningfirmwarerelationship["object_type"] = item.ObjectType
	firmwarerunningfirmwarerelationship["selector"] = item.Selector

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
	firmwareserverconfigurationutilitydistributablerelationship["class_id"] = item.ClassId
	firmwareserverconfigurationutilitydistributablerelationship["moid"] = item.Moid
	firmwareserverconfigurationutilitydistributablerelationship["object_type"] = item.ObjectType
	firmwareserverconfigurationutilitydistributablerelationship["selector"] = item.Selector

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
	firmwareupgradebaserelationship["class_id"] = item.ClassId
	firmwareupgradebaserelationship["moid"] = item.Moid
	firmwareupgradebaserelationship["object_type"] = item.ObjectType
	firmwareupgradebaserelationship["selector"] = item.Selector

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
	firmwareupgradeimpactstatusrelationship["class_id"] = item.ClassId
	firmwareupgradeimpactstatusrelationship["moid"] = item.Moid
	firmwareupgradeimpactstatusrelationship["object_type"] = item.ObjectType
	firmwareupgradeimpactstatusrelationship["selector"] = item.Selector

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
	firmwareupgradestatusrelationship["class_id"] = item.ClassId
	firmwareupgradestatusrelationship["moid"] = item.Moid
	firmwareupgradestatusrelationship["object_type"] = item.ObjectType
	firmwareupgradestatusrelationship["selector"] = item.Selector

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
	forecastcatalogrelationship["class_id"] = item.ClassId
	forecastcatalogrelationship["moid"] = item.Moid
	forecastcatalogrelationship["object_type"] = item.ObjectType
	forecastcatalogrelationship["selector"] = item.Selector

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
	forecastdefinitionrelationship["class_id"] = item.ClassId
	forecastdefinitionrelationship["moid"] = item.Moid
	forecastdefinitionrelationship["object_type"] = item.ObjectType
	forecastdefinitionrelationship["selector"] = item.Selector

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
	forecastinstancerelationship["class_id"] = item.ClassId
	forecastinstancerelationship["moid"] = item.Moid
	forecastinstancerelationship["object_type"] = item.ObjectType
	forecastinstancerelationship["selector"] = item.Selector

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
	forecastmodel["accuracy"] = item.Accuracy
	forecastmodel["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastmodel["class_id"] = item.ClassId
	forecastmodel["model_data"] = item.ModelData
	forecastmodel["model_type"] = item.ModelType
	forecastmodel["object_type"] = item.ObjectType

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
	graphicscardrelationship["class_id"] = item.ClassId
	graphicscardrelationship["moid"] = item.Moid
	graphicscardrelationship["object_type"] = item.ObjectType
	graphicscardrelationship["selector"] = item.Selector

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
	hcloperatingsystemrelationship["class_id"] = item.ClassId
	hcloperatingsystemrelationship["moid"] = item.Moid
	hcloperatingsystemrelationship["object_type"] = item.ObjectType
	hcloperatingsystemrelationship["selector"] = item.Selector

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
	hcloperatingsystemvendorrelationship["class_id"] = item.ClassId
	hcloperatingsystemvendorrelationship["moid"] = item.Moid
	hcloperatingsystemvendorrelationship["object_type"] = item.ObjectType
	hcloperatingsystemvendorrelationship["selector"] = item.Selector

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
	hyperflexalarmsummary["class_id"] = item.ClassId
	hyperflexalarmsummary["critical"] = item.Critical
	hyperflexalarmsummary["object_type"] = item.ObjectType
	hyperflexalarmsummary["warning"] = item.Warning

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
	hyperflexappcatalogrelationship["class_id"] = item.ClassId
	hyperflexappcatalogrelationship["moid"] = item.Moid
	hyperflexappcatalogrelationship["object_type"] = item.ObjectType
	hyperflexappcatalogrelationship["selector"] = item.Selector

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
	hyperflexappsettingconstraint["class_id"] = item.ClassId
	hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
	hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
	hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
	hyperflexappsettingconstraint["object_type"] = item.ObjectType
	hyperflexappsettingconstraint["server_model"] = item.ServerModel

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
	hyperflexautosupportpolicyrelationship["class_id"] = item.ClassId
	hyperflexautosupportpolicyrelationship["moid"] = item.Moid
	hyperflexautosupportpolicyrelationship["object_type"] = item.ObjectType
	hyperflexautosupportpolicyrelationship["selector"] = item.Selector

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
	hyperflexbackupclusterrelationship["class_id"] = item.ClassId
	hyperflexbackupclusterrelationship["moid"] = item.Moid
	hyperflexbackupclusterrelationship["object_type"] = item.ObjectType
	hyperflexbackupclusterrelationship["selector"] = item.Selector

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
	hyperflexbaseclusterrelationship["class_id"] = item.ClassId
	hyperflexbaseclusterrelationship["moid"] = item.Moid
	hyperflexbaseclusterrelationship["object_type"] = item.ObjectType
	hyperflexbaseclusterrelationship["selector"] = item.Selector

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
	hyperflexbondstate["active_slave"] = item.ActiveSlave
	hyperflexbondstate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexbondstate["class_id"] = item.ClassId
	hyperflexbondstate["mode"] = item.Mode
	hyperflexbondstate["object_type"] = item.ObjectType
	hyperflexbondstate["slaves"] = item.Slaves

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
	hyperflexciscohypervisormanagerrelationship["class_id"] = item.ClassId
	hyperflexciscohypervisormanagerrelationship["moid"] = item.Moid
	hyperflexciscohypervisormanagerrelationship["object_type"] = item.ObjectType
	hyperflexciscohypervisormanagerrelationship["selector"] = item.Selector

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
	hyperflexclusterrelationship["class_id"] = item.ClassId
	hyperflexclusterrelationship["moid"] = item.Moid
	hyperflexclusterrelationship["object_type"] = item.ObjectType
	hyperflexclusterrelationship["selector"] = item.Selector

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
	hyperflexclusternetworkpolicyrelationship["class_id"] = item.ClassId
	hyperflexclusternetworkpolicyrelationship["moid"] = item.Moid
	hyperflexclusternetworkpolicyrelationship["object_type"] = item.ObjectType
	hyperflexclusternetworkpolicyrelationship["selector"] = item.Selector

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
	hyperflexclusterprofilerelationship["class_id"] = item.ClassId
	hyperflexclusterprofilerelationship["moid"] = item.Moid
	hyperflexclusterprofilerelationship["object_type"] = item.ObjectType
	hyperflexclusterprofilerelationship["selector"] = item.Selector

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
	hyperflexclusterstoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexclusterstoragepolicyrelationship["moid"] = item.Moid
	hyperflexclusterstoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexclusterstoragepolicyrelationship["selector"] = item.Selector

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
	hyperflexconfigresultrelationship["class_id"] = item.ClassId
	hyperflexconfigresultrelationship["moid"] = item.Moid
	hyperflexconfigresultrelationship["object_type"] = item.ObjectType
	hyperflexconfigresultrelationship["selector"] = item.Selector

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
	hyperflexdataprotectionpeerrelationship["class_id"] = item.ClassId
	hyperflexdataprotectionpeerrelationship["moid"] = item.Moid
	hyperflexdataprotectionpeerrelationship["object_type"] = item.ObjectType
	hyperflexdataprotectionpeerrelationship["selector"] = item.Selector

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
	hyperflexdiskstatus["class_id"] = item.ClassId
	hyperflexdiskstatus["download_percentage"] = item.DownloadPercentage
	hyperflexdiskstatus["object_type"] = item.ObjectType
	hyperflexdiskstatus["state"] = item.State
	hyperflexdiskstatus["volume_handle"] = item.VolumeHandle
	hyperflexdiskstatus["volume_name"] = item.VolumeName

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
	hyperflexentityreference["class_id"] = item.ClassId
	hyperflexentityreference["confignum"] = item.Confignum
	hyperflexentityreference["id"] = item.Id
	hyperflexentityreference["idtype"] = item.Idtype
	hyperflexentityreference["name"] = item.Name
	hyperflexentityreference["object_type"] = item.ObjectType
	hyperflexentityreference["type"] = item.Type

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
	hyperflexerrorstack["class_id"] = item.ClassId
	hyperflexerrorstack["message"] = item.Message
	hyperflexerrorstack["message_id"] = item.MessageId
	hyperflexerrorstack["object_type"] = item.ObjectType

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
	hyperflexextfcstoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexextfcstoragepolicyrelationship["moid"] = item.Moid
	hyperflexextfcstoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexextfcstoragepolicyrelationship["selector"] = item.Selector

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
	hyperflexextiscsistoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexextiscsistoragepolicyrelationship["moid"] = item.Moid
	hyperflexextiscsistoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexextiscsistoragepolicyrelationship["selector"] = item.Selector

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
	hyperflexfeaturelimitexternalrelationship["class_id"] = item.ClassId
	hyperflexfeaturelimitexternalrelationship["moid"] = item.Moid
	hyperflexfeaturelimitexternalrelationship["object_type"] = item.ObjectType
	hyperflexfeaturelimitexternalrelationship["selector"] = item.Selector

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
	hyperflexfeaturelimitinternalrelationship["class_id"] = item.ClassId
	hyperflexfeaturelimitinternalrelationship["moid"] = item.Moid
	hyperflexfeaturelimitinternalrelationship["object_type"] = item.ObjectType
	hyperflexfeaturelimitinternalrelationship["selector"] = item.Selector

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
	hyperflexhealthrelationship["class_id"] = item.ClassId
	hyperflexhealthrelationship["moid"] = item.Moid
	hyperflexhealthrelationship["object_type"] = item.ObjectType
	hyperflexhealthrelationship["selector"] = item.Selector

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
	hyperflexhealthcheckdefinitionrelationship["class_id"] = item.ClassId
	hyperflexhealthcheckdefinitionrelationship["moid"] = item.Moid
	hyperflexhealthcheckdefinitionrelationship["object_type"] = item.ObjectType
	hyperflexhealthcheckdefinitionrelationship["selector"] = item.Selector

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
	hyperflexhealthcheckscriptinfo["aggregate_script_name"] = item.AggregateScriptName
	hyperflexhealthcheckscriptinfo["class_id"] = item.ClassId
	hyperflexhealthcheckscriptinfo["hyperflex_version"] = item.HyperflexVersion
	hyperflexhealthcheckscriptinfo["object_type"] = item.ObjectType
	hyperflexhealthcheckscriptinfo["script_execute_location"] = item.ScriptExecuteLocation
	hyperflexhealthcheckscriptinfo["script_name"] = item.ScriptName
	hyperflexhealthcheckscriptinfo["nr_version"] = item.Version

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
	hyperflexhxlicenseauthorizationdetailsdt["class_id"] = item.ClassId
	hyperflexhxlicenseauthorizationdetailsdt["communication_deadline_date"] = item.CommunicationDeadlineDate
	hyperflexhxlicenseauthorizationdetailsdt["evaluation_period_expired_at"] = item.EvaluationPeriodExpiredAt
	hyperflexhxlicenseauthorizationdetailsdt["evaluation_period_remaining"] = item.EvaluationPeriodRemaining
	hyperflexhxlicenseauthorizationdetailsdt["last_communication_attempt_date"] = item.LastCommunicationAttemptDate
	hyperflexhxlicenseauthorizationdetailsdt["next_communication_attempt_date"] = item.NextCommunicationAttemptDate
	hyperflexhxlicenseauthorizationdetailsdt["object_type"] = item.ObjectType
	hyperflexhxlicenseauthorizationdetailsdt["status"] = item.Status

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
	hyperflexhxnetworkaddressdt["address"] = item.Address
	hyperflexhxnetworkaddressdt["class_id"] = item.ClassId
	hyperflexhxnetworkaddressdt["fqdn"] = item.Fqdn
	hyperflexhxnetworkaddressdt["ip"] = item.Ip
	hyperflexhxnetworkaddressdt["object_type"] = item.ObjectType

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
	hyperflexhxplatformdatastoreconfigdt["class_id"] = item.ClassId
	hyperflexhxplatformdatastoreconfigdt["data_block_size"] = item.DataBlockSize
	hyperflexhxplatformdatastoreconfigdt["name"] = item.Name
	hyperflexhxplatformdatastoreconfigdt["num_mirrors"] = item.NumMirrors
	hyperflexhxplatformdatastoreconfigdt["num_stripes_for_large_files"] = item.NumStripesForLargeFiles
	hyperflexhxplatformdatastoreconfigdt["object_type"] = item.ObjectType
	hyperflexhxplatformdatastoreconfigdt["provisioned_capacity"] = item.ProvisionedCapacity
	hyperflexhxplatformdatastoreconfigdt["system_datastore"] = item.SystemDatastore

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
	hyperflexhxregistrationdetailsdt["class_id"] = item.ClassId
	hyperflexhxregistrationdetailsdt["initial_registration_date"] = item.InitialRegistrationDate
	hyperflexhxregistrationdetailsdt["last_renewal_attempt_date"] = item.LastRenewalAttemptDate
	hyperflexhxregistrationdetailsdt["next_renewal_attempt_date"] = item.NextRenewalAttemptDate
	hyperflexhxregistrationdetailsdt["object_type"] = item.ObjectType
	hyperflexhxregistrationdetailsdt["out_of_compliance_start_date"] = item.OutOfComplianceStartDate
	hyperflexhxregistrationdetailsdt["registration_expire_date"] = item.RegistrationExpireDate
	hyperflexhxregistrationdetailsdt["registration_failed_reason"] = item.RegistrationFailedReason
	hyperflexhxregistrationdetailsdt["smart_account"] = item.SmartAccount
	hyperflexhxregistrationdetailsdt["status"] = item.Status
	hyperflexhxregistrationdetailsdt["virtual_account"] = item.VirtualAccount

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
	hyperflexhxresiliencyinfodt["class_id"] = item.ClassId
	hyperflexhxresiliencyinfodt["data_replication_factor"] = item.DataReplicationFactor
	hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.HddFailuresTolerable
	hyperflexhxresiliencyinfodt["messages"] = item.Messages
	hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.NodeFailuresTolerable
	hyperflexhxresiliencyinfodt["object_type"] = item.ObjectType
	hyperflexhxresiliencyinfodt["policy_compliance"] = item.PolicyCompliance
	hyperflexhxresiliencyinfodt["resiliency_state"] = item.ResiliencyState
	hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.SsdFailuresTolerable

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
	hyperflexhxsitedt["class_id"] = item.ClassId
	hyperflexhxsitedt["name"] = item.Name
	hyperflexhxsitedt["object_type"] = item.ObjectType
	hyperflexhxsitedt["zone"] = (func(p models.HyperflexHxZoneInfoDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxzoneinfodts []map[string]interface{}
		var ret models.HyperflexHxZoneInfoDt
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexhxzoneinfodt := make(map[string]interface{})
		hyperflexhxzoneinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhxzoneinfodt["class_id"] = item.ClassId
		hyperflexhxzoneinfodt["num_nodes"] = item.NumNodes
		hyperflexhxzoneinfodt["object_type"] = item.ObjectType
		hyperflexhxzoneinfodt["uuid"] = item.Uuid

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
	hyperflexhxuuiddt["class_id"] = item.ClassId
	hyperflexhxuuiddt["links"] = (func(p []models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxlinkdts []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			hyperflexhxlinkdt := make(map[string]interface{})
			hyperflexhxlinkdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexhxlinkdt["class_id"] = item.ClassId
			hyperflexhxlinkdt["comments"] = item.Comments
			hyperflexhxlinkdt["href"] = item.Href
			hyperflexhxlinkdt["method"] = item.Method
			hyperflexhxlinkdt["object_type"] = item.ObjectType
			hyperflexhxlinkdt["rel"] = item.Rel
			hyperflexhxlinkdts = append(hyperflexhxlinkdts, hyperflexhxlinkdt)
		}
		return hyperflexhxlinkdts
	})(item.GetLinks(), d)
	hyperflexhxuuiddt["object_type"] = item.ObjectType
	hyperflexhxuuiddt["uuid"] = item.Uuid

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
	hyperflexhxapclusterrelationship["class_id"] = item.ClassId
	hyperflexhxapclusterrelationship["moid"] = item.Moid
	hyperflexhxapclusterrelationship["object_type"] = item.ObjectType
	hyperflexhxapclusterrelationship["selector"] = item.Selector

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
	hyperflexhxapdvuplinkrelationship["class_id"] = item.ClassId
	hyperflexhxapdvuplinkrelationship["moid"] = item.Moid
	hyperflexhxapdvuplinkrelationship["object_type"] = item.ObjectType
	hyperflexhxapdvuplinkrelationship["selector"] = item.Selector

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
	hyperflexhxapdvswitchrelationship["class_id"] = item.ClassId
	hyperflexhxapdvswitchrelationship["moid"] = item.Moid
	hyperflexhxapdvswitchrelationship["object_type"] = item.ObjectType
	hyperflexhxapdvswitchrelationship["selector"] = item.Selector

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
	hyperflexhxaphostrelationship["class_id"] = item.ClassId
	hyperflexhxaphostrelationship["moid"] = item.Moid
	hyperflexhxaphostrelationship["object_type"] = item.ObjectType
	hyperflexhxaphostrelationship["selector"] = item.Selector

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
	hyperflexhxapnetworkrelationship["class_id"] = item.ClassId
	hyperflexhxapnetworkrelationship["moid"] = item.Moid
	hyperflexhxapnetworkrelationship["object_type"] = item.ObjectType
	hyperflexhxapnetworkrelationship["selector"] = item.Selector

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
	hyperflexhxapvirtualmachinerelationship["class_id"] = item.ClassId
	hyperflexhxapvirtualmachinerelationship["moid"] = item.Moid
	hyperflexhxapvirtualmachinerelationship["object_type"] = item.ObjectType
	hyperflexhxapvirtualmachinerelationship["selector"] = item.Selector

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
	hyperflexipaddrrange["class_id"] = item.ClassId
	hyperflexipaddrrange["end_addr"] = item.EndAddr
	hyperflexipaddrrange["gateway"] = item.Gateway
	hyperflexipaddrrange["netmask"] = item.Netmask
	hyperflexipaddrrange["object_type"] = item.ObjectType
	hyperflexipaddrrange["start_addr"] = item.StartAddr

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
	hyperflexlicenserelationship["class_id"] = item.ClassId
	hyperflexlicenserelationship["moid"] = item.Moid
	hyperflexlicenserelationship["object_type"] = item.ObjectType
	hyperflexlicenserelationship["selector"] = item.Selector

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
	hyperflexlocalcredentialpolicyrelationship["class_id"] = item.ClassId
	hyperflexlocalcredentialpolicyrelationship["moid"] = item.Moid
	hyperflexlocalcredentialpolicyrelationship["object_type"] = item.ObjectType
	hyperflexlocalcredentialpolicyrelationship["selector"] = item.Selector

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
	hyperflexlogicalavailabilityzone["auto_config"] = item.AutoConfig
	hyperflexlogicalavailabilityzone["class_id"] = item.ClassId
	hyperflexlogicalavailabilityzone["object_type"] = item.ObjectType

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
	hyperflexmacaddrprefixrange["class_id"] = item.ClassId
	hyperflexmacaddrprefixrange["end_addr"] = item.EndAddr
	hyperflexmacaddrprefixrange["object_type"] = item.ObjectType
	hyperflexmacaddrprefixrange["start_addr"] = item.StartAddr

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
	hyperflexnamedvlan["class_id"] = item.ClassId
	hyperflexnamedvlan["name"] = item.Name
	hyperflexnamedvlan["object_type"] = item.ObjectType
	hyperflexnamedvlan["vlan_id"] = item.VlanId

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
	hyperflexnamedvsan["class_id"] = item.ClassId
	hyperflexnamedvsan["name"] = item.Name
	hyperflexnamedvsan["object_type"] = item.ObjectType
	hyperflexnamedvsan["vsan_id"] = item.VsanId

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
	hyperflexnodeconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexnodeconfigpolicyrelationship["moid"] = item.Moid
	hyperflexnodeconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexnodeconfigpolicyrelationship["selector"] = item.Selector

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
	hyperflexproxysettingpolicyrelationship["class_id"] = item.ClassId
	hyperflexproxysettingpolicyrelationship["moid"] = item.Moid
	hyperflexproxysettingpolicyrelationship["object_type"] = item.ObjectType
	hyperflexproxysettingpolicyrelationship["selector"] = item.Selector

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
	hyperflexreplicationpeerinfo["class_id"] = item.ClassId
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
				hyperflexreplicationplatdatastore["class_id"] = item.ClassId
				hyperflexreplicationplatdatastore["cluster_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

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
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetDatastoreEr(), d)
				hyperflexreplicationplatdatastore["object_type"] = item.ObjectType

				hyperflexreplicationplatdatastores = append(hyperflexreplicationplatdatastores, hyperflexreplicationplatdatastore)
				return hyperflexreplicationplatdatastores
			})(item.GetAds(), d)
			hyperflexreplicationplatdatastorepair["backup_only"] = item.BackupOnly
			hyperflexreplicationplatdatastorepair["bds"] = (func(p models.HyperflexReplicationPlatDatastore, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationplatdatastores []map[string]interface{}
				var ret models.HyperflexReplicationPlatDatastore
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationplatdatastore := make(map[string]interface{})
				hyperflexreplicationplatdatastore["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationplatdatastore["class_id"] = item.ClassId
				hyperflexreplicationplatdatastore["cluster_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

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
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetDatastoreEr(), d)
				hyperflexreplicationplatdatastore["object_type"] = item.ObjectType

				hyperflexreplicationplatdatastores = append(hyperflexreplicationplatdatastores, hyperflexreplicationplatdatastore)
				return hyperflexreplicationplatdatastores
			})(item.GetBds(), d)
			hyperflexreplicationplatdatastorepair["class_id"] = item.ClassId
			hyperflexreplicationplatdatastorepair["object_type"] = item.ObjectType
			hyperflexreplicationplatdatastorepair["quiesce"] = item.Quiesce
			hyperflexreplicationplatdatastorepair["replication_interval_in_minutes"] = item.ReplicationIntervalInMinutes
			hyperflexreplicationplatdatastorepair["sourceds"] = (func(p models.HyperflexReplicationPlatDatastore, d *schema.ResourceData) []map[string]interface{} {
				var hyperflexreplicationplatdatastores []map[string]interface{}
				var ret models.HyperflexReplicationPlatDatastore
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				hyperflexreplicationplatdatastore := make(map[string]interface{})
				hyperflexreplicationplatdatastore["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				hyperflexreplicationplatdatastore["class_id"] = item.ClassId
				hyperflexreplicationplatdatastore["cluster_er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
					var hyperflexentityreferences []map[string]interface{}
					var ret models.HyperflexEntityReference
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					hyperflexentityreference := make(map[string]interface{})
					hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

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
					hyperflexentityreference["class_id"] = item.ClassId
					hyperflexentityreference["confignum"] = item.Confignum
					hyperflexentityreference["id"] = item.Id
					hyperflexentityreference["idtype"] = item.Idtype
					hyperflexentityreference["name"] = item.Name
					hyperflexentityreference["object_type"] = item.ObjectType
					hyperflexentityreference["type"] = item.Type

					hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
					return hyperflexentityreferences
				})(item.GetDatastoreEr(), d)
				hyperflexreplicationplatdatastore["object_type"] = item.ObjectType

				hyperflexreplicationplatdatastores = append(hyperflexreplicationplatdatastores, hyperflexreplicationplatdatastore)
				return hyperflexreplicationplatdatastores
			})(item.GetSourceds(), d)
			hyperflexreplicationplatdatastorepair["storage_only"] = item.StorageOnly
			hyperflexreplicationplatdatastorepairs = append(hyperflexreplicationplatdatastorepairs, hyperflexreplicationplatdatastorepair)
		}
		return hyperflexreplicationplatdatastorepairs
	})(item.GetDatastores(), d)
	hyperflexreplicationpeerinfo["dcip"] = item.Dcip
	hyperflexreplicationpeerinfo["er"] = (func(p models.HyperflexEntityReference, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexentityreferences []map[string]interface{}
		var ret models.HyperflexEntityReference
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexentityreference := make(map[string]interface{})
		hyperflexentityreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexentityreference["class_id"] = item.ClassId
		hyperflexentityreference["confignum"] = item.Confignum
		hyperflexentityreference["id"] = item.Id
		hyperflexentityreference["idtype"] = item.Idtype
		hyperflexentityreference["name"] = item.Name
		hyperflexentityreference["object_type"] = item.ObjectType
		hyperflexentityreference["type"] = item.Type

		hyperflexentityreferences = append(hyperflexentityreferences, hyperflexentityreference)
		return hyperflexentityreferences
	})(item.GetEr(), d)
	hyperflexreplicationpeerinfo["mcip"] = item.Mcip
	hyperflexreplicationpeerinfo["object_type"] = item.ObjectType
	hyperflexreplicationpeerinfo["ports"] = (func(p []models.HyperflexPortTypeToPortNumberMap, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexporttypetoportnumbermaps []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			hyperflexporttypetoportnumbermap := make(map[string]interface{})
			hyperflexporttypetoportnumbermap["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexporttypetoportnumbermap["class_id"] = item.ClassId
			hyperflexporttypetoportnumbermap["i16"] = item.I16
			hyperflexporttypetoportnumbermap["object_type"] = item.ObjectType
			hyperflexporttypetoportnumbermap["string"] = item.String
			hyperflexporttypetoportnumbermaps = append(hyperflexporttypetoportnumbermaps, hyperflexporttypetoportnumbermap)
		}
		return hyperflexporttypetoportnumbermaps
	})(item.GetPorts(), d)
	hyperflexreplicationpeerinfo["repl_cip"] = item.ReplCip
	hyperflexreplicationpeerinfo["status"] = item.Status
	hyperflexreplicationpeerinfo["status_details"] = item.StatusDetails

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
	hyperflexreplicationschedule["backup_interval"] = item.BackupInterval
	hyperflexreplicationschedule["class_id"] = item.ClassId
	hyperflexreplicationschedule["object_type"] = item.ObjectType

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
	hyperflexserverfirmwareversionrelationship["class_id"] = item.ClassId
	hyperflexserverfirmwareversionrelationship["moid"] = item.Moid
	hyperflexserverfirmwareversionrelationship["object_type"] = item.ObjectType
	hyperflexserverfirmwareversionrelationship["selector"] = item.Selector

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
	hyperflexservermodelrelationship["class_id"] = item.ClassId
	hyperflexservermodelrelationship["moid"] = item.Moid
	hyperflexservermodelrelationship["object_type"] = item.ObjectType
	hyperflexservermodelrelationship["selector"] = item.Selector

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
	hyperflexsoftwaredistributionentryrelationship["class_id"] = item.ClassId
	hyperflexsoftwaredistributionentryrelationship["moid"] = item.Moid
	hyperflexsoftwaredistributionentryrelationship["object_type"] = item.ObjectType
	hyperflexsoftwaredistributionentryrelationship["selector"] = item.Selector

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
	hyperflexsoftwaredistributionversionrelationship["class_id"] = item.ClassId
	hyperflexsoftwaredistributionversionrelationship["moid"] = item.Moid
	hyperflexsoftwaredistributionversionrelationship["object_type"] = item.ObjectType
	hyperflexsoftwaredistributionversionrelationship["selector"] = item.Selector

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
	hyperflexsoftwareversionpolicyrelationship["class_id"] = item.ClassId
	hyperflexsoftwareversionpolicyrelationship["moid"] = item.Moid
	hyperflexsoftwareversionpolicyrelationship["object_type"] = item.ObjectType
	hyperflexsoftwareversionpolicyrelationship["selector"] = item.Selector

	hyperflexsoftwareversionpolicyrelationships = append(hyperflexsoftwareversionpolicyrelationships, hyperflexsoftwareversionpolicyrelationship)
	return hyperflexsoftwareversionpolicyrelationships
}
func flattenMapHyperflexSummary(p models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsummarys []map[string]interface{}
	var ret models.HyperflexSummary
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexsummary := make(map[string]interface{})
	hyperflexsummary["active_nodes"] = item.ActiveNodes
	hyperflexsummary["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsummary["address"] = item.Address
	hyperflexsummary["boottime"] = item.Boottime
	hyperflexsummary["class_id"] = item.ClassId
	hyperflexsummary["cluster_access_policy"] = item.ClusterAccessPolicy
	hyperflexsummary["compression_savings"] = item.CompressionSavings
	hyperflexsummary["data_replication_compliance"] = item.DataReplicationCompliance
	hyperflexsummary["data_replication_factor"] = item.DataReplicationFactor
	hyperflexsummary["deduplication_savings"] = item.DeduplicationSavings
	hyperflexsummary["downtime"] = item.Downtime
	hyperflexsummary["free_capacity"] = item.FreeCapacity
	hyperflexsummary["healing_info"] = (func(p models.HyperflexStPlatformClusterHealingInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterhealinginfos []map[string]interface{}
		var ret models.HyperflexStPlatformClusterHealingInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexstplatformclusterhealinginfo := make(map[string]interface{})
		hyperflexstplatformclusterhealinginfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexstplatformclusterhealinginfo["class_id"] = item.ClassId
		hyperflexstplatformclusterhealinginfo["estimated_completion_time_in_seconds"] = item.EstimatedCompletionTimeInSeconds
		hyperflexstplatformclusterhealinginfo["in_progress"] = item.InProgress
		hyperflexstplatformclusterhealinginfo["messages"] = item.Messages
		hyperflexstplatformclusterhealinginfo["messages_iterator"] = item.MessagesIterator
		hyperflexstplatformclusterhealinginfo["messages_size"] = item.MessagesSize
		hyperflexstplatformclusterhealinginfo["object_type"] = item.ObjectType
		hyperflexstplatformclusterhealinginfo["percent_complete"] = item.PercentComplete

		hyperflexstplatformclusterhealinginfos = append(hyperflexstplatformclusterhealinginfos, hyperflexstplatformclusterhealinginfo)
		return hyperflexstplatformclusterhealinginfos
	})(item.GetHealingInfo(), d)
	hyperflexsummary["name"] = item.Name
	hyperflexsummary["object_type"] = item.ObjectType
	hyperflexsummary["resiliency_details"] = item.ResiliencyDetails
	hyperflexsummary["resiliency_details_size"] = item.ResiliencyDetailsSize
	hyperflexsummary["resiliency_info"] = (func(p models.HyperflexStPlatformClusterResiliencyInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterresiliencyinfos []map[string]interface{}
		var ret models.HyperflexStPlatformClusterResiliencyInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexstplatformclusterresiliencyinfo := make(map[string]interface{})
		hyperflexstplatformclusterresiliencyinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexstplatformclusterresiliencyinfo["class_id"] = item.ClassId
		hyperflexstplatformclusterresiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["messages"] = item.Messages
		hyperflexstplatformclusterresiliencyinfo["messages_iterator"] = item.MessagesIterator
		hyperflexstplatformclusterresiliencyinfo["messages_size"] = item.MessagesSize
		hyperflexstplatformclusterresiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["object_type"] = item.ObjectType
		hyperflexstplatformclusterresiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["state"] = item.State

		hyperflexstplatformclusterresiliencyinfos = append(hyperflexstplatformclusterresiliencyinfos, hyperflexstplatformclusterresiliencyinfo)
		return hyperflexstplatformclusterresiliencyinfos
	})(item.GetResiliencyInfo(), d)
	hyperflexsummary["space_status"] = item.SpaceStatus
	hyperflexsummary["state"] = item.State
	hyperflexsummary["total_capacity"] = item.TotalCapacity
	hyperflexsummary["total_savings"] = item.TotalSavings
	hyperflexsummary["uptime"] = item.Uptime
	hyperflexsummary["used_capacity"] = item.UsedCapacity

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
	hyperflexsysconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexsysconfigpolicyrelationship["moid"] = item.Moid
	hyperflexsysconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexsysconfigpolicyrelationship["selector"] = item.Selector

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
	hyperflexucsmconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexucsmconfigpolicyrelationship["moid"] = item.Moid
	hyperflexucsmconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexucsmconfigpolicyrelationship["selector"] = item.Selector

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
	hyperflexvcenterconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexvcenterconfigpolicyrelationship["moid"] = item.Moid
	hyperflexvcenterconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexvcenterconfigpolicyrelationship["selector"] = item.Selector

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
	hyperflexvirtualmachine["class_id"] = item.ClassId
	hyperflexvirtualmachine["object_type"] = item.ObjectType
	hyperflexvirtualmachine["run_time_info"] = (func(p models.HyperflexVirtualMachineRuntimeInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexvirtualmachineruntimeinfos []map[string]interface{}
		var ret models.HyperflexVirtualMachineRuntimeInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexvirtualmachineruntimeinfo := make(map[string]interface{})
		hyperflexvirtualmachineruntimeinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexvirtualmachineruntimeinfo["bios_uuid"] = item.BiosUuid
		hyperflexvirtualmachineruntimeinfo["class_id"] = item.ClassId
		hyperflexvirtualmachineruntimeinfo["connection_state"] = item.ConnectionState
		hyperflexvirtualmachineruntimeinfo["cpu_usage"] = item.CpuUsage
		hyperflexvirtualmachineruntimeinfo["folder"] = item.Folder
		hyperflexvirtualmachineruntimeinfo["guest_family"] = item.GuestFamily
		hyperflexvirtualmachineruntimeinfo["guest_full_name"] = item.GuestFullName
		hyperflexvirtualmachineruntimeinfo["guest_id"] = item.GuestId
		hyperflexvirtualmachineruntimeinfo["guest_state"] = item.GuestState
		hyperflexvirtualmachineruntimeinfo["host_name"] = item.HostName
		hyperflexvirtualmachineruntimeinfo["instance_uuid"] = item.InstanceUuid
		hyperflexvirtualmachineruntimeinfo["memory_mb"] = item.MemoryMb
		hyperflexvirtualmachineruntimeinfo["memory_usage"] = item.MemoryUsage
		hyperflexvirtualmachineruntimeinfo["moid"] = item.Moid
		hyperflexvirtualmachineruntimeinfo["name"] = item.Name
		hyperflexvirtualmachineruntimeinfo["networks"] = item.Networks
		hyperflexvirtualmachineruntimeinfo["num_cpu"] = item.NumCpu
		hyperflexvirtualmachineruntimeinfo["object_type"] = item.ObjectType
		hyperflexvirtualmachineruntimeinfo["power_state"] = item.PowerState
		hyperflexvirtualmachineruntimeinfo["provisioned_size"] = item.ProvisionedSize
		hyperflexvirtualmachineruntimeinfo["resource_pool"] = item.ResourcePool
		hyperflexvirtualmachineruntimeinfo["used_size"] = item.UsedSize
		hyperflexvirtualmachineruntimeinfo["nr_version"] = item.Version
		hyperflexvirtualmachineruntimeinfo["vmx_path"] = item.VmxPath

		hyperflexvirtualmachineruntimeinfos = append(hyperflexvirtualmachineruntimeinfos, hyperflexvirtualmachineruntimeinfo)
		return hyperflexvirtualmachineruntimeinfos
	})(item.GetRunTimeInfo(), d)
	hyperflexvirtualmachine["status_code"] = item.StatusCode
	hyperflexvirtualmachine["uuid"] = item.Uuid

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
	hyperflexvmbackupinforelationship["class_id"] = item.ClassId
	hyperflexvmbackupinforelationship["moid"] = item.Moid
	hyperflexvmbackupinforelationship["object_type"] = item.ObjectType
	hyperflexvmbackupinforelationship["selector"] = item.Selector

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
	hyperflexvmsnapshotinforelationship["class_id"] = item.ClassId
	hyperflexvmsnapshotinforelationship["moid"] = item.Moid
	hyperflexvmsnapshotinforelationship["object_type"] = item.ObjectType
	hyperflexvmsnapshotinforelationship["selector"] = item.Selector

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
	hyperflexwwxnprefixrange["class_id"] = item.ClassId
	hyperflexwwxnprefixrange["end_addr"] = item.EndAddr
	hyperflexwwxnprefixrange["object_type"] = item.ObjectType
	hyperflexwwxnprefixrange["start_addr"] = item.StartAddr

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
	iaaslicenseinforelationship["class_id"] = item.ClassId
	iaaslicenseinforelationship["moid"] = item.Moid
	iaaslicenseinforelationship["object_type"] = item.ObjectType
	iaaslicenseinforelationship["selector"] = item.Selector

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
	iaasucsdinforelationship["class_id"] = item.ClassId
	iaasucsdinforelationship["moid"] = item.Moid
	iaasucsdinforelationship["object_type"] = item.ObjectType
	iaasucsdinforelationship["selector"] = item.Selector

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
	iaasucsdmanagedinfrarelationship["class_id"] = item.ClassId
	iaasucsdmanagedinfrarelationship["moid"] = item.Moid
	iaasucsdmanagedinfrarelationship["object_type"] = item.ObjectType
	iaasucsdmanagedinfrarelationship["selector"] = item.Selector

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
	iamaccountrelationship["class_id"] = item.ClassId
	iamaccountrelationship["moid"] = item.Moid
	iamaccountrelationship["object_type"] = item.ObjectType
	iamaccountrelationship["selector"] = item.Selector

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
	iamappregistrationrelationship["class_id"] = item.ClassId
	iamappregistrationrelationship["moid"] = item.Moid
	iamappregistrationrelationship["object_type"] = item.ObjectType
	iamappregistrationrelationship["selector"] = item.Selector

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
	iamcertificaterelationship["class_id"] = item.ClassId
	iamcertificaterelationship["moid"] = item.Moid
	iamcertificaterelationship["object_type"] = item.ObjectType
	iamcertificaterelationship["selector"] = item.Selector

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
	iamcertificaterequestrelationship["class_id"] = item.ClassId
	iamcertificaterequestrelationship["moid"] = item.Moid
	iamcertificaterequestrelationship["object_type"] = item.ObjectType
	iamcertificaterequestrelationship["selector"] = item.Selector

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
	iamclientmeta["class_id"] = item.ClassId
	iamclientmeta["device_model"] = item.DeviceModel
	iamclientmeta["object_type"] = item.ObjectType
	iamclientmeta["user_agent"] = item.UserAgent

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
	iamdomaingrouprelationship["class_id"] = item.ClassId
	iamdomaingrouprelationship["moid"] = item.Moid
	iamdomaingrouprelationship["object_type"] = item.ObjectType
	iamdomaingrouprelationship["selector"] = item.Selector

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
	iamendpointpasswordproperties["class_id"] = item.ClassId
	iamendpointpasswordproperties["enable_password_expiry"] = item.EnablePasswordExpiry
	iamendpointpasswordproperties["enforce_strong_password"] = item.EnforceStrongPassword
	iamendpointpasswordproperties["force_send_password"] = item.ForceSendPassword
	iamendpointpasswordproperties["grace_period"] = item.GracePeriod
	iamendpointpasswordproperties["notification_period"] = item.NotificationPeriod
	iamendpointpasswordproperties["object_type"] = item.ObjectType
	iamendpointpasswordproperties["password_expiry_duration"] = item.PasswordExpiryDuration
	iamendpointpasswordproperties["password_history"] = item.PasswordHistory

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
	iamendpointuserrelationship["class_id"] = item.ClassId
	iamendpointuserrelationship["moid"] = item.Moid
	iamendpointuserrelationship["object_type"] = item.ObjectType
	iamendpointuserrelationship["selector"] = item.Selector

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
	iamendpointuserpolicyrelationship["class_id"] = item.ClassId
	iamendpointuserpolicyrelationship["moid"] = item.Moid
	iamendpointuserpolicyrelationship["object_type"] = item.ObjectType
	iamendpointuserpolicyrelationship["selector"] = item.Selector

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
	iamidprelationship["class_id"] = item.ClassId
	iamidprelationship["moid"] = item.Moid
	iamidprelationship["object_type"] = item.ObjectType
	iamidprelationship["selector"] = item.Selector

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
	iamidpreferencerelationship["class_id"] = item.ClassId
	iamidpreferencerelationship["moid"] = item.Moid
	iamidpreferencerelationship["object_type"] = item.ObjectType
	iamidpreferencerelationship["selector"] = item.Selector

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
	iamipaccessmanagementrelationship["class_id"] = item.ClassId
	iamipaccessmanagementrelationship["moid"] = item.Moid
	iamipaccessmanagementrelationship["object_type"] = item.ObjectType
	iamipaccessmanagementrelationship["selector"] = item.Selector

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
	iamldapbaseproperties["attribute"] = item.Attribute
	iamldapbaseproperties["base_dn"] = item.BaseDn
	iamldapbaseproperties["bind_dn"] = item.BindDn
	iamldapbaseproperties["bind_method"] = item.BindMethod
	iamldapbaseproperties["class_id"] = item.ClassId
	iamldapbaseproperties["domain"] = item.Domain
	iamldapbaseproperties["enable_encryption"] = item.EnableEncryption
	iamldapbaseproperties["enable_group_authorization"] = item.EnableGroupAuthorization
	iamldapbaseproperties["filter"] = item.Filter
	iamldapbaseproperties["group_attribute"] = item.GroupAttribute
	iamldapbaseproperties["is_password_set"] = item.IsPasswordSet
	iamldapbaseproperties["nested_group_search_depth"] = item.NestedGroupSearchDepth
	iamldapbaseproperties["object_type"] = item.ObjectType
	password_x := d.Get("base_properties").([]interface{})
	if len(password_x) > 0 {
		password_y := password_x[0].(map[string]interface{})
		iamldapbaseproperties["password"] = password_y["password"]
	}
	iamldapbaseproperties["timeout"] = item.Timeout

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
	iamldapdnsparameters["class_id"] = item.ClassId
	iamldapdnsparameters["object_type"] = item.ObjectType
	iamldapdnsparameters["search_domain"] = item.SearchDomain
	iamldapdnsparameters["search_forest"] = item.SearchForest
	iamldapdnsparameters["nr_source"] = item.Source

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
	iamldappolicyrelationship["class_id"] = item.ClassId
	iamldappolicyrelationship["moid"] = item.Moid
	iamldappolicyrelationship["object_type"] = item.ObjectType
	iamldappolicyrelationship["selector"] = item.Selector

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
	iamlocaluserpasswordrelationship["class_id"] = item.ClassId
	iamlocaluserpasswordrelationship["moid"] = item.Moid
	iamlocaluserpasswordrelationship["object_type"] = item.ObjectType
	iamlocaluserpasswordrelationship["selector"] = item.Selector

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
	iampermissionrelationship["class_id"] = item.ClassId
	iampermissionrelationship["moid"] = item.Moid
	iampermissionrelationship["object_type"] = item.ObjectType
	iampermissionrelationship["selector"] = item.Selector

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
	iamprivatekeyspecrelationship["class_id"] = item.ClassId
	iamprivatekeyspecrelationship["moid"] = item.Moid
	iamprivatekeyspecrelationship["object_type"] = item.ObjectType
	iamprivatekeyspecrelationship["selector"] = item.Selector

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
	iamqualifierrelationship["class_id"] = item.ClassId
	iamqualifierrelationship["moid"] = item.Moid
	iamqualifierrelationship["object_type"] = item.ObjectType
	iamqualifierrelationship["selector"] = item.Selector

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
	iamresourcelimitsrelationship["class_id"] = item.ClassId
	iamresourcelimitsrelationship["moid"] = item.Moid
	iamresourcelimitsrelationship["object_type"] = item.ObjectType
	iamresourcelimitsrelationship["selector"] = item.Selector

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
	iamsecurityholderrelationship["class_id"] = item.ClassId
	iamsecurityholderrelationship["moid"] = item.Moid
	iamsecurityholderrelationship["object_type"] = item.ObjectType
	iamsecurityholderrelationship["selector"] = item.Selector

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
	iamserviceproviderrelationship["class_id"] = item.ClassId
	iamserviceproviderrelationship["moid"] = item.Moid
	iamserviceproviderrelationship["object_type"] = item.ObjectType
	iamserviceproviderrelationship["selector"] = item.Selector

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
	iamsessionrelationship["class_id"] = item.ClassId
	iamsessionrelationship["moid"] = item.Moid
	iamsessionrelationship["object_type"] = item.ObjectType
	iamsessionrelationship["selector"] = item.Selector

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
	iamsessionlimitsrelationship["class_id"] = item.ClassId
	iamsessionlimitsrelationship["moid"] = item.Moid
	iamsessionlimitsrelationship["object_type"] = item.ObjectType
	iamsessionlimitsrelationship["selector"] = item.Selector

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
	iamsystemrelationship["class_id"] = item.ClassId
	iamsystemrelationship["moid"] = item.Moid
	iamsystemrelationship["object_type"] = item.ObjectType
	iamsystemrelationship["selector"] = item.Selector

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
	iamuserrelationship["class_id"] = item.ClassId
	iamuserrelationship["moid"] = item.Moid
	iamuserrelationship["object_type"] = item.ObjectType
	iamuserrelationship["selector"] = item.Selector

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
	iamusergrouprelationship["class_id"] = item.ClassId
	iamusergrouprelationship["moid"] = item.Moid
	iamusergrouprelationship["object_type"] = item.ObjectType
	iamusergrouprelationship["selector"] = item.Selector

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
	infrahardwareinfo["class_id"] = item.ClassId
	infrahardwareinfo["cpu_cores"] = item.CpuCores
	infrahardwareinfo["cpu_speed"] = item.CpuSpeed
	infrahardwareinfo["memory_size"] = item.MemorySize
	infrahardwareinfo["object_type"] = item.ObjectType

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
	inventorybaserelationship["class_id"] = item.ClassId
	inventorybaserelationship["moid"] = item.Moid
	inventorybaserelationship["object_type"] = item.ObjectType
	inventorybaserelationship["selector"] = item.Selector

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
	inventorydeviceinforelationship["class_id"] = item.ClassId
	inventorydeviceinforelationship["moid"] = item.Moid
	inventorydeviceinforelationship["object_type"] = item.ObjectType
	inventorydeviceinforelationship["selector"] = item.Selector

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
	inventorygenericinventoryholderrelationship["class_id"] = item.ClassId
	inventorygenericinventoryholderrelationship["moid"] = item.Moid
	inventorygenericinventoryholderrelationship["object_type"] = item.ObjectType
	inventorygenericinventoryholderrelationship["selector"] = item.Selector

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
	ippoolblockleaserelationship["class_id"] = item.ClassId
	ippoolblockleaserelationship["moid"] = item.Moid
	ippoolblockleaserelationship["object_type"] = item.ObjectType
	ippoolblockleaserelationship["selector"] = item.Selector

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
	ippoolipleaserelationship["class_id"] = item.ClassId
	ippoolipleaserelationship["moid"] = item.Moid
	ippoolipleaserelationship["object_type"] = item.ObjectType
	ippoolipleaserelationship["selector"] = item.Selector

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
	ippoolipv4block["class_id"] = item.ClassId
	ippoolipv4block["from"] = item.From
	ippoolipv4block["object_type"] = item.ObjectType
	ippoolipv4block["size"] = item.Size
	ippoolipv4block["to"] = item.To

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
	ippoolipv4config["class_id"] = item.ClassId
	ippoolipv4config["gateway"] = item.Gateway
	ippoolipv4config["netmask"] = item.Netmask
	ippoolipv4config["object_type"] = item.ObjectType
	ippoolipv4config["primary_dns"] = item.PrimaryDns
	ippoolipv4config["secondary_dns"] = item.SecondaryDns

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
	ippoolipv6block["class_id"] = item.ClassId
	ippoolipv6block["from"] = item.From
	ippoolipv6block["object_type"] = item.ObjectType
	ippoolipv6block["size"] = item.Size
	ippoolipv6block["to"] = item.To

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
	ippoolipv6config["class_id"] = item.ClassId
	ippoolipv6config["gateway"] = item.Gateway
	ippoolipv6config["object_type"] = item.ObjectType
	ippoolipv6config["prefix"] = item.Prefix
	ippoolipv6config["primary_dns"] = item.PrimaryDns
	ippoolipv6config["secondary_dns"] = item.SecondaryDns

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
	ippoolpoolrelationship["class_id"] = item.ClassId
	ippoolpoolrelationship["moid"] = item.Moid
	ippoolpoolrelationship["object_type"] = item.ObjectType
	ippoolpoolrelationship["selector"] = item.Selector

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
	ippoolpoolmemberrelationship["class_id"] = item.ClassId
	ippoolpoolmemberrelationship["moid"] = item.Moid
	ippoolpoolmemberrelationship["object_type"] = item.ObjectType
	ippoolpoolmemberrelationship["selector"] = item.Selector

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
	ippoolshadowblockrelationship["class_id"] = item.ClassId
	ippoolshadowblockrelationship["moid"] = item.Moid
	ippoolshadowblockrelationship["object_type"] = item.ObjectType
	ippoolshadowblockrelationship["selector"] = item.Selector

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
	ippoolshadowpoolrelationship["class_id"] = item.ClassId
	ippoolshadowpoolrelationship["moid"] = item.Moid
	ippoolshadowpoolrelationship["object_type"] = item.ObjectType
	ippoolshadowpoolrelationship["selector"] = item.Selector

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
	ippooluniverserelationship["class_id"] = item.ClassId
	ippooluniverserelationship["moid"] = item.Moid
	ippooluniverserelationship["object_type"] = item.ObjectType
	ippooluniverserelationship["selector"] = item.Selector

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
	iqnpoolblockrelationship["class_id"] = item.ClassId
	iqnpoolblockrelationship["moid"] = item.Moid
	iqnpoolblockrelationship["object_type"] = item.ObjectType
	iqnpoolblockrelationship["selector"] = item.Selector

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
	iqnpooliqnsuffixblock["class_id"] = item.ClassId
	iqnpooliqnsuffixblock["from"] = item.From
	iqnpooliqnsuffixblock["object_type"] = item.ObjectType
	iqnpooliqnsuffixblock["size"] = item.Size
	iqnpooliqnsuffixblock["suffix"] = item.Suffix
	iqnpooliqnsuffixblock["to"] = item.To

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
	iqnpoolleaserelationship["class_id"] = item.ClassId
	iqnpoolleaserelationship["moid"] = item.Moid
	iqnpoolleaserelationship["object_type"] = item.ObjectType
	iqnpoolleaserelationship["selector"] = item.Selector

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
	iqnpoolpoolrelationship["class_id"] = item.ClassId
	iqnpoolpoolrelationship["moid"] = item.Moid
	iqnpoolpoolrelationship["object_type"] = item.ObjectType
	iqnpoolpoolrelationship["selector"] = item.Selector

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
	iqnpoolpoolmemberrelationship["class_id"] = item.ClassId
	iqnpoolpoolmemberrelationship["moid"] = item.Moid
	iqnpoolpoolmemberrelationship["object_type"] = item.ObjectType
	iqnpoolpoolmemberrelationship["selector"] = item.Selector

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
	iqnpooluniverserelationship["class_id"] = item.ClassId
	iqnpooluniverserelationship["moid"] = item.Moid
	iqnpooluniverserelationship["object_type"] = item.ObjectType
	iqnpooluniverserelationship["selector"] = item.Selector

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
	kubernetesacicniprofilerelationship["class_id"] = item.ClassId
	kubernetesacicniprofilerelationship["moid"] = item.Moid
	kubernetesacicniprofilerelationship["object_type"] = item.ObjectType
	kubernetesacicniprofilerelationship["selector"] = item.Selector

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
	kubernetesactioninfo["class_id"] = item.ClassId
	kubernetesactioninfo["failure_reason"] = item.FailureReason
	kubernetesactioninfo["name"] = item.Name
	kubernetesactioninfo["object_type"] = item.ObjectType
	kubernetesactioninfo["status"] = item.Status

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
	kubernetesaddonconfiguration["class_id"] = item.ClassId
	kubernetesaddonconfiguration["install_strategy"] = item.InstallStrategy
	kubernetesaddonconfiguration["object_type"] = item.ObjectType
	kubernetesaddonconfiguration["override_sets"] = (func(p []models.KubernetesKeyValue, d *schema.ResourceData) []map[string]interface{} {
		var kuberneteskeyvalues []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			kuberneteskeyvalue := make(map[string]interface{})
			kuberneteskeyvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			kuberneteskeyvalue["class_id"] = item.ClassId
			kuberneteskeyvalue["key"] = item.Key
			kuberneteskeyvalue["object_type"] = item.ObjectType
			kuberneteskeyvalue["value"] = item.Value
			kuberneteskeyvalues = append(kuberneteskeyvalues, kuberneteskeyvalue)
		}
		return kuberneteskeyvalues
	})(item.GetOverrideSets(), d)
	kubernetesaddonconfiguration["overrides"] = item.Overrides
	kubernetesaddonconfiguration["release_name"] = item.ReleaseName
	kubernetesaddonconfiguration["release_namespace"] = item.ReleaseNamespace
	kubernetesaddonconfiguration["upgrade_strategy"] = item.UpgradeStrategy

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
	kubernetesaddondefinitionrelationship["class_id"] = item.ClassId
	kubernetesaddondefinitionrelationship["moid"] = item.Moid
	kubernetesaddondefinitionrelationship["object_type"] = item.ObjectType
	kubernetesaddondefinitionrelationship["selector"] = item.Selector

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
	kubernetesbaseinfrastructureproviderrelationship["class_id"] = item.ClassId
	kubernetesbaseinfrastructureproviderrelationship["moid"] = item.Moid
	kubernetesbaseinfrastructureproviderrelationship["object_type"] = item.ObjectType
	kubernetesbaseinfrastructureproviderrelationship["selector"] = item.Selector

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
	kubernetesbasevirtualmachineinfraconfig["class_id"] = item.ClassId
	kubernetesbasevirtualmachineinfraconfig["interfaces"] = item.Interfaces
	kubernetesbasevirtualmachineinfraconfig["object_type"] = item.ObjectType

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
	kubernetescatalogrelationship["class_id"] = item.ClassId
	kubernetescatalogrelationship["moid"] = item.Moid
	kubernetescatalogrelationship["object_type"] = item.ObjectType
	kubernetescatalogrelationship["selector"] = item.Selector

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
	kubernetesclusterrelationship["class_id"] = item.ClassId
	kubernetesclusterrelationship["moid"] = item.Moid
	kubernetesclusterrelationship["object_type"] = item.ObjectType
	kubernetesclusterrelationship["selector"] = item.Selector

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
	kubernetesclusteraddonprofilerelationship["class_id"] = item.ClassId
	kubernetesclusteraddonprofilerelationship["moid"] = item.Moid
	kubernetesclusteraddonprofilerelationship["object_type"] = item.ObjectType
	kubernetesclusteraddonprofilerelationship["selector"] = item.Selector

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
	kubernetesclustercertificateconfiguration["ca_cert"] = item.CaCert
	ca_key_x := d.Get("cert_config").([]interface{})
	if len(ca_key_x) > 0 {
		ca_key_y := ca_key_x[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["ca_key"] = ca_key_y["ca_key"]
	}
	kubernetesclustercertificateconfiguration["class_id"] = item.ClassId
	kubernetesclustercertificateconfiguration["etcd_cert"] = item.EtcdCert
	kubernetesclustercertificateconfiguration["etcd_encryption_key"] = item.EtcdEncryptionKey
	etcd_key_x := d.Get("cert_config").([]interface{})
	if len(etcd_key_x) > 0 {
		etcd_key_y := etcd_key_x[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["etcd_key"] = etcd_key_y["etcd_key"]
	}
	kubernetesclustercertificateconfiguration["front_proxy_cert"] = item.FrontProxyCert
	front_proxy_key_x := d.Get("cert_config").([]interface{})
	if len(front_proxy_key_x) > 0 {
		front_proxy_key_y := front_proxy_key_x[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["front_proxy_key"] = front_proxy_key_y["front_proxy_key"]
	}
	kubernetesclustercertificateconfiguration["object_type"] = item.ObjectType
	sa_private_key_x := d.Get("cert_config").([]interface{})
	if len(sa_private_key_x) > 0 {
		sa_private_key_y := sa_private_key_x[0].(map[string]interface{})
		kubernetesclustercertificateconfiguration["sa_private_key"] = sa_private_key_y["sa_private_key"]
	}
	kubernetesclustercertificateconfiguration["sa_public_key"] = item.SaPublicKey

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
	kubernetesclustermanagementconfig["class_id"] = item.ClassId
	kubernetesclustermanagementconfig["encrypted_etcd"] = item.EncryptedEtcd
	kubernetesclustermanagementconfig["load_balancer_count"] = item.LoadBalancerCount
	kubernetesclustermanagementconfig["load_balancers"] = item.LoadBalancers
	kubernetesclustermanagementconfig["master_vip"] = item.MasterVip
	kubernetesclustermanagementconfig["object_type"] = item.ObjectType
	kubernetesclustermanagementconfig["ssh_keys"] = item.SshKeys
	kubernetesclustermanagementconfig["ssh_user"] = item.SshUser

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
	kubernetesclusterprofilerelationship["class_id"] = item.ClassId
	kubernetesclusterprofilerelationship["moid"] = item.Moid
	kubernetesclusterprofilerelationship["object_type"] = item.ObjectType
	kubernetesclusterprofilerelationship["selector"] = item.Selector

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
	kubernetescniconfig["class_id"] = item.ClassId
	kubernetescniconfig["object_type"] = item.ObjectType
	kubernetescniconfig["registry"] = item.Registry
	kubernetescniconfig["nr_version"] = item.Version

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
	kubernetesconfigresultrelationship["class_id"] = item.ClassId
	kubernetesconfigresultrelationship["moid"] = item.Moid
	kubernetesconfigresultrelationship["object_type"] = item.ObjectType
	kubernetesconfigresultrelationship["selector"] = item.Selector

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
	kubernetesconfiguration["class_id"] = item.ClassId
	kubernetesconfiguration["kube_config"] = item.KubeConfig
	kubernetesconfiguration["object_type"] = item.ObjectType

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
	kubernetescontainerruntimepolicyrelationship["class_id"] = item.ClassId
	kubernetescontainerruntimepolicyrelationship["moid"] = item.Moid
	kubernetescontainerruntimepolicyrelationship["object_type"] = item.ObjectType
	kubernetescontainerruntimepolicyrelationship["selector"] = item.Selector

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
	kubernetesdaemonsetstatus["class_id"] = item.ClassId
	kubernetesdaemonsetstatus["current_number_scheduled"] = item.CurrentNumberScheduled
	kubernetesdaemonsetstatus["desired_number_scheduled"] = item.DesiredNumberScheduled
	kubernetesdaemonsetstatus["number_available"] = item.NumberAvailable
	kubernetesdaemonsetstatus["number_misscheduled"] = item.NumberMisscheduled
	kubernetesdaemonsetstatus["number_ready"] = item.NumberReady
	kubernetesdaemonsetstatus["object_type"] = item.ObjectType
	kubernetesdaemonsetstatus["observed_generation"] = item.ObservedGeneration
	kubernetesdaemonsetstatus["updated_number_scheduled"] = item.UpdatedNumberScheduled

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
	kubernetesdeploymentstatus["available_replicas"] = item.AvailableReplicas
	kubernetesdeploymentstatus["class_id"] = item.ClassId
	kubernetesdeploymentstatus["object_type"] = item.ObjectType
	kubernetesdeploymentstatus["observed_generation"] = item.ObservedGeneration
	kubernetesdeploymentstatus["ready_replicas"] = item.ReadyReplicas
	kubernetesdeploymentstatus["replicas"] = item.Replicas
	kubernetesdeploymentstatus["updated_replicas"] = item.UpdatedReplicas

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
	kubernetesingressstatus["class_id"] = item.ClassId
	kubernetesingressstatus["load_balancer"] = (func(p models.KubernetesLoadBalancer, d *schema.ResourceData) []map[string]interface{} {
		var kubernetesloadbalancers []map[string]interface{}
		var ret models.KubernetesLoadBalancer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		kubernetesloadbalancer := make(map[string]interface{})
		kubernetesloadbalancer["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesloadbalancer["class_id"] = item.ClassId
		kubernetesloadbalancer["ip_addresses"] = item.IpAddresses
		kubernetesloadbalancer["object_type"] = item.ObjectType

		kubernetesloadbalancers = append(kubernetesloadbalancers, kubernetesloadbalancer)
		return kubernetesloadbalancers
	})(item.GetLoadBalancer(), d)
	kubernetesingressstatus["object_type"] = item.ObjectType

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
	kubernetesnetworkpolicyrelationship["class_id"] = item.ClassId
	kubernetesnetworkpolicyrelationship["moid"] = item.Moid
	kubernetesnetworkpolicyrelationship["object_type"] = item.ObjectType
	kubernetesnetworkpolicyrelationship["selector"] = item.Selector

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
	kubernetesnodegroupprofilerelationship["class_id"] = item.ClassId
	kubernetesnodegroupprofilerelationship["moid"] = item.Moid
	kubernetesnodegroupprofilerelationship["object_type"] = item.ObjectType
	kubernetesnodegroupprofilerelationship["selector"] = item.Selector

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
	kubernetesnodeinfo["architecture"] = item.Architecture
	kubernetesnodeinfo["boot_id"] = item.BootId
	kubernetesnodeinfo["class_id"] = item.ClassId
	kubernetesnodeinfo["container_runtime_version"] = item.ContainerRuntimeVersion
	kubernetesnodeinfo["kernel_version"] = item.KernelVersion
	kubernetesnodeinfo["kube_proxy_version"] = item.KubeProxyVersion
	kubernetesnodeinfo["kubelet_version"] = item.KubeletVersion
	kubernetesnodeinfo["machine_id"] = item.MachineId
	kubernetesnodeinfo["object_type"] = item.ObjectType
	kubernetesnodeinfo["operating_system"] = item.OperatingSystem
	kubernetesnodeinfo["os_image"] = item.OsImage
	kubernetesnodeinfo["system_uuid"] = item.SystemUuid

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
	kubernetesnodeprofilerelationship["class_id"] = item.ClassId
	kubernetesnodeprofilerelationship["moid"] = item.Moid
	kubernetesnodeprofilerelationship["object_type"] = item.ObjectType
	kubernetesnodeprofilerelationship["selector"] = item.Selector

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
	kubernetesnodespec["class_id"] = item.ClassId
	kubernetesnodespec["object_type"] = item.ObjectType
	kubernetesnodespec["pod_cidr"] = item.PodCidr
	kubernetesnodespec["provider_id"] = item.ProviderId

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
	kubernetesobjectmeta["class_id"] = item.ClassId
	kubernetesobjectmeta["creation_timestamp"] = item.CreationTimestamp
	kubernetesobjectmeta["name"] = item.Name
	kubernetesobjectmeta["namespace"] = item.Namespace
	kubernetesobjectmeta["object_type"] = item.ObjectType
	kubernetesobjectmeta["resource_version"] = item.ResourceVersion
	kubernetesobjectmeta["uuid"] = item.Uuid

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
	kubernetespodstatus["class_id"] = item.ClassId
	kubernetespodstatus["host_ip"] = item.HostIp
	kubernetespodstatus["object_type"] = item.ObjectType
	kubernetespodstatus["phase"] = item.Phase
	kubernetespodstatus["pod_ip"] = item.PodIp
	kubernetespodstatus["qos_class"] = item.QosClass
	kubernetespodstatus["start_time"] = item.StartTime

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
	kubernetesproxyconfig["class_id"] = item.ClassId
	kubernetesproxyconfig["hostname"] = item.Hostname
	kubernetesproxyconfig["is_password_set"] = item.IsPasswordSet
	kubernetesproxyconfig["object_type"] = item.ObjectType
	password_x := d.Get("docker_http_proxy").([]interface{})
	if len(password_x) > 0 {
		password_y := password_x[0].(map[string]interface{})
		kubernetesproxyconfig["password"] = password_y["password"]
	}
	kubernetesproxyconfig["port"] = item.Port
	kubernetesproxyconfig["protocol"] = item.Protocol
	kubernetesproxyconfig["username"] = item.Username

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
	kubernetesservicestatus["class_id"] = item.ClassId
	kubernetesservicestatus["load_balancer"] = (func(p models.KubernetesLoadBalancer, d *schema.ResourceData) []map[string]interface{} {
		var kubernetesloadbalancers []map[string]interface{}
		var ret models.KubernetesLoadBalancer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		kubernetesloadbalancer := make(map[string]interface{})
		kubernetesloadbalancer["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		kubernetesloadbalancer["class_id"] = item.ClassId
		kubernetesloadbalancer["ip_addresses"] = item.IpAddresses
		kubernetesloadbalancer["object_type"] = item.ObjectType

		kubernetesloadbalancers = append(kubernetesloadbalancers, kubernetesloadbalancer)
		return kubernetesloadbalancers
	})(item.GetLoadBalancer(), d)
	kubernetesservicestatus["object_type"] = item.ObjectType

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
	kubernetesstatefulsetstatus["available_replicas"] = item.AvailableReplicas
	kubernetesstatefulsetstatus["class_id"] = item.ClassId
	kubernetesstatefulsetstatus["collision_count"] = item.CollisionCount
	kubernetesstatefulsetstatus["current_revision"] = item.CurrentRevision
	kubernetesstatefulsetstatus["object_type"] = item.ObjectType
	kubernetesstatefulsetstatus["observed_generation"] = item.ObservedGeneration
	kubernetesstatefulsetstatus["ready_replicas"] = item.ReadyReplicas
	kubernetesstatefulsetstatus["replicas"] = item.Replicas
	kubernetesstatefulsetstatus["update_revision"] = item.UpdateRevision
	kubernetesstatefulsetstatus["updated_replicas"] = item.UpdatedReplicas

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
	kubernetessysconfigpolicyrelationship["class_id"] = item.ClassId
	kubernetessysconfigpolicyrelationship["moid"] = item.Moid
	kubernetessysconfigpolicyrelationship["object_type"] = item.ObjectType
	kubernetessysconfigpolicyrelationship["selector"] = item.Selector

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
	kubernetestrustedregistriespolicyrelationship["class_id"] = item.ClassId
	kubernetestrustedregistriespolicyrelationship["moid"] = item.Moid
	kubernetestrustedregistriespolicyrelationship["object_type"] = item.ObjectType
	kubernetestrustedregistriespolicyrelationship["selector"] = item.Selector

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
	kubernetesversionrelationship["class_id"] = item.ClassId
	kubernetesversionrelationship["moid"] = item.Moid
	kubernetesversionrelationship["object_type"] = item.ObjectType
	kubernetesversionrelationship["selector"] = item.Selector

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
	kubernetesversionpolicyrelationship["class_id"] = item.ClassId
	kubernetesversionpolicyrelationship["moid"] = item.Moid
	kubernetesversionpolicyrelationship["object_type"] = item.ObjectType
	kubernetesversionpolicyrelationship["selector"] = item.Selector

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
	kubernetesvirtualmachineinfraconfigpolicyrelationship["class_id"] = item.ClassId
	kubernetesvirtualmachineinfraconfigpolicyrelationship["moid"] = item.Moid
	kubernetesvirtualmachineinfraconfigpolicyrelationship["object_type"] = item.ObjectType
	kubernetesvirtualmachineinfraconfigpolicyrelationship["selector"] = item.Selector

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
	kubernetesvirtualmachineinstancetyperelationship["class_id"] = item.ClassId
	kubernetesvirtualmachineinstancetyperelationship["moid"] = item.Moid
	kubernetesvirtualmachineinstancetyperelationship["object_type"] = item.ObjectType
	kubernetesvirtualmachineinstancetyperelationship["selector"] = item.Selector

	kubernetesvirtualmachineinstancetyperelationships = append(kubernetesvirtualmachineinstancetyperelationships, kubernetesvirtualmachineinstancetyperelationship)
	return kubernetesvirtualmachineinstancetyperelationships
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
	kvmtunnelrelationship["class_id"] = item.ClassId
	kvmtunnelrelationship["moid"] = item.Moid
	kvmtunnelrelationship["object_type"] = item.ObjectType
	kvmtunnelrelationship["selector"] = item.Selector

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
	licenseaccountlicensedatarelationship["class_id"] = item.ClassId
	licenseaccountlicensedatarelationship["moid"] = item.Moid
	licenseaccountlicensedatarelationship["object_type"] = item.ObjectType
	licenseaccountlicensedatarelationship["selector"] = item.Selector

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
	licensecustomeroprelationship["class_id"] = item.ClassId
	licensecustomeroprelationship["moid"] = item.Moid
	licensecustomeroprelationship["object_type"] = item.ObjectType
	licensecustomeroprelationship["selector"] = item.Selector

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
	licenseiwocustomeroprelationship["class_id"] = item.ClassId
	licenseiwocustomeroprelationship["moid"] = item.Moid
	licenseiwocustomeroprelationship["object_type"] = item.ObjectType
	licenseiwocustomeroprelationship["selector"] = item.Selector

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
	licenseiwolicensecountrelationship["class_id"] = item.ClassId
	licenseiwolicensecountrelationship["moid"] = item.Moid
	licenseiwolicensecountrelationship["object_type"] = item.ObjectType
	licenseiwolicensecountrelationship["selector"] = item.Selector

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
	licensesmartlicensetokenrelationship["class_id"] = item.ClassId
	licensesmartlicensetokenrelationship["moid"] = item.Moid
	licensesmartlicensetokenrelationship["object_type"] = item.ObjectType
	licensesmartlicensetokenrelationship["selector"] = item.Selector

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
	macpoolblock["class_id"] = item.ClassId
	macpoolblock["from"] = item.From
	macpoolblock["object_type"] = item.ObjectType
	macpoolblock["size"] = item.Size
	macpoolblock["to"] = item.To

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
	macpoolidblockrelationship["class_id"] = item.ClassId
	macpoolidblockrelationship["moid"] = item.Moid
	macpoolidblockrelationship["object_type"] = item.ObjectType
	macpoolidblockrelationship["selector"] = item.Selector

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
	macpoolleaserelationship["class_id"] = item.ClassId
	macpoolleaserelationship["moid"] = item.Moid
	macpoolleaserelationship["object_type"] = item.ObjectType
	macpoolleaserelationship["selector"] = item.Selector

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
	macpoolpoolrelationship["class_id"] = item.ClassId
	macpoolpoolrelationship["moid"] = item.Moid
	macpoolpoolrelationship["object_type"] = item.ObjectType
	macpoolpoolrelationship["selector"] = item.Selector

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
	macpoolpoolmemberrelationship["class_id"] = item.ClassId
	macpoolpoolmemberrelationship["moid"] = item.Moid
	macpoolpoolmemberrelationship["object_type"] = item.ObjectType
	macpoolpoolmemberrelationship["selector"] = item.Selector

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
	macpooluniverserelationship["class_id"] = item.ClassId
	macpooluniverserelationship["moid"] = item.Moid
	macpooluniverserelationship["object_type"] = item.ObjectType
	macpooluniverserelationship["selector"] = item.Selector

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
	managementcontrollerrelationship["class_id"] = item.ClassId
	managementcontrollerrelationship["moid"] = item.Moid
	managementcontrollerrelationship["object_type"] = item.ObjectType
	managementcontrollerrelationship["selector"] = item.Selector

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
	managemententityrelationship["class_id"] = item.ClassId
	managemententityrelationship["moid"] = item.Moid
	managemententityrelationship["object_type"] = item.ObjectType
	managemententityrelationship["selector"] = item.Selector

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
	memoryarrayrelationship["class_id"] = item.ClassId
	memoryarrayrelationship["moid"] = item.Moid
	memoryarrayrelationship["object_type"] = item.ObjectType
	memoryarrayrelationship["selector"] = item.Selector

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
	memorypersistentmemoryconfigresultrelationship["class_id"] = item.ClassId
	memorypersistentmemoryconfigresultrelationship["moid"] = item.Moid
	memorypersistentmemoryconfigresultrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryconfigresultrelationship["selector"] = item.Selector

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
	memorypersistentmemoryconfigurationrelationship["class_id"] = item.ClassId
	memorypersistentmemoryconfigurationrelationship["moid"] = item.Moid
	memorypersistentmemoryconfigurationrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryconfigurationrelationship["selector"] = item.Selector

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
	memorypersistentmemorylocalsecurity["class_id"] = item.ClassId
	memorypersistentmemorylocalsecurity["enabled"] = item.Enabled
	memorypersistentmemorylocalsecurity["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	memorypersistentmemorylocalsecurity["object_type"] = item.ObjectType
	secure_passphrase_x := d.Get("local_security").([]interface{})
	if len(secure_passphrase_x) > 0 {
		secure_passphrase_y := secure_passphrase_x[0].(map[string]interface{})
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
	memorypersistentmemoryregionrelationship["class_id"] = item.ClassId
	memorypersistentmemoryregionrelationship["moid"] = item.Moid
	memorypersistentmemoryregionrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryregionrelationship["selector"] = item.Selector

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
	mobasemorelationship["class_id"] = item.ClassId
	mobasemorelationship["moid"] = item.Moid
	mobasemorelationship["object_type"] = item.ObjectType
	mobasemorelationship["selector"] = item.Selector

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
	moversioncontext["class_id"] = item.ClassId
	moversioncontext["interested_mos"] = (func(p []models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
		var momorefs []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			momoref := make(map[string]interface{})
			momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			momoref["class_id"] = item.ClassId
			momoref["moid"] = item.Moid
			momoref["object_type"] = item.ObjectType
			momoref["selector"] = item.Selector
			momorefs = append(momorefs, momoref)
		}
		return momorefs
	})(item.GetInterestedMos(), d)
	moversioncontext["object_type"] = item.ObjectType
	moversioncontext["ref_mo"] = (func(p models.MoMoRef, d *schema.ResourceData) []map[string]interface{} {
		var momorefs []map[string]interface{}
		var ret models.MoMoRef
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		momoref := make(map[string]interface{})
		momoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		momoref["class_id"] = item.ClassId
		momoref["moid"] = item.Moid
		momoref["object_type"] = item.ObjectType
		momoref["selector"] = item.Selector

		momorefs = append(momorefs, momoref)
		return momorefs
	})(item.GetRefMo(), d)
	moversioncontext["timestamp"] = item.Timestamp
	moversioncontext["nr_version"] = item.Version
	moversioncontext["version_type"] = item.VersionType

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
	networkelementrelationship["class_id"] = item.ClassId
	networkelementrelationship["moid"] = item.Moid
	networkelementrelationship["object_type"] = item.ObjectType
	networkelementrelationship["selector"] = item.Selector

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
	networkfczoneinforelationship["class_id"] = item.ClassId
	networkfczoneinforelationship["moid"] = item.Moid
	networkfczoneinforelationship["object_type"] = item.ObjectType
	networkfczoneinforelationship["selector"] = item.Selector

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
	networkvlanportinforelationship["class_id"] = item.ClassId
	networkvlanportinforelationship["moid"] = item.Moid
	networkvlanportinforelationship["object_type"] = item.ObjectType
	networkvlanportinforelationship["selector"] = item.Selector

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
	niaapinewreleasedetail["class_id"] = item.ClassId
	niaapinewreleasedetail["description"] = item.Description
	niaapinewreleasedetail["link"] = item.Link
	niaapinewreleasedetail["object_type"] = item.ObjectType
	niaapinewreleasedetail["release_note_link"] = item.ReleaseNoteLink
	niaapinewreleasedetail["release_note_link_title"] = item.ReleaseNoteLinkTitle
	niaapinewreleasedetail["software_download_link"] = item.SoftwareDownloadLink
	niaapinewreleasedetail["software_download_link_title"] = item.SoftwareDownloadLinkTitle
	niaapinewreleasedetail["title"] = item.Title
	niaapinewreleasedetail["nr_version"] = item.Version

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
	niaapiversionregexplatform["anyllregex"] = item.Anyllregex
	niaapiversionregexplatform["class_id"] = item.ClassId
	niaapiversionregexplatform["currentlltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

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
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetLatestsltrain(), d)
	niaapiversionregexplatform["object_type"] = item.ObjectType
	niaapiversionregexplatform["sltrain"] = (func(p []models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			niaapisoftwareregex := make(map[string]interface{})
			niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			niaapisoftwareregex["class_id"] = item.ClassId
			niaapisoftwareregex["object_type"] = item.ObjectType
			niaapisoftwareregex["regex"] = item.Regex
			niaapisoftwareregex["software_version"] = item.SoftwareVersion
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
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetUpcominglltrain(), d)

	niaapiversionregexplatforms = append(niaapiversionregexplatforms, niaapiversionregexplatform)
	return niaapiversionregexplatforms
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
	niatelemetrydiskinfo["class_id"] = item.ClassId
	niatelemetrydiskinfo["free"] = item.Free
	niatelemetrydiskinfo["name"] = item.Name
	niatelemetrydiskinfo["object_type"] = item.ObjectType
	niatelemetrydiskinfo["total"] = item.Total
	niatelemetrydiskinfo["used"] = item.Used

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
	niatelemetryinterface["class_id"] = item.ClassId
	niatelemetryinterface["interface_down_count"] = item.InterfaceDownCount
	niatelemetryinterface["interface_up_count"] = item.InterfaceUpCount
	niatelemetryinterface["object_type"] = item.ObjectType

	niatelemetryinterfaces = append(niatelemetryinterfaces, niatelemetryinterface)
	return niatelemetryinterfaces
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
	niatelemetryniainventoryrelationship["class_id"] = item.ClassId
	niatelemetryniainventoryrelationship["moid"] = item.Moid
	niatelemetryniainventoryrelationship["object_type"] = item.ObjectType
	niatelemetryniainventoryrelationship["selector"] = item.Selector

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
	niatelemetrynialicensestaterelationship["class_id"] = item.ClassId
	niatelemetrynialicensestaterelationship["moid"] = item.Moid
	niatelemetrynialicensestaterelationship["object_type"] = item.ObjectType
	niatelemetrynialicensestaterelationship["selector"] = item.Selector

	niatelemetrynialicensestaterelationships = append(niatelemetrynialicensestaterelationships, niatelemetrynialicensestaterelationship)
	return niatelemetrynialicensestaterelationships
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
	niatelemetrynvevni["class_id"] = item.ClassId
	niatelemetrynvevni["cp_vni_count"] = item.CpVniCount
	niatelemetrynvevni["cp_vni_down"] = item.CpVniDown
	niatelemetrynvevni["cp_vni_up"] = item.CpVniUp
	niatelemetrynvevni["dp_vni_count"] = item.DpVniCount
	niatelemetrynvevni["dp_vni_down"] = item.DpVniDown
	niatelemetrynvevni["dp_vni_up"] = item.DpVniUp
	niatelemetrynvevni["object_type"] = item.ObjectType

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
	niatelemetrynxosbgpmvpn["capable_peers"] = item.CapablePeers
	niatelemetrynxosbgpmvpn["class_id"] = item.ClassId
	niatelemetrynxosbgpmvpn["configured_peers"] = item.ConfiguredPeers
	niatelemetrynxosbgpmvpn["memory_used"] = item.MemoryUsed
	niatelemetrynxosbgpmvpn["number_of_cluster_lists"] = item.NumberOfClusterLists
	niatelemetrynxosbgpmvpn["number_of_communities"] = item.NumberOfCommunities
	niatelemetrynxosbgpmvpn["object_type"] = item.ObjectType
	niatelemetrynxosbgpmvpn["table_version"] = item.TableVersion
	niatelemetrynxosbgpmvpn["total_networks"] = item.TotalNetworks
	niatelemetrynxosbgpmvpn["total_paths"] = item.TotalPaths

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
	niatelemetrynxosvtp["class_id"] = item.ClassId
	niatelemetrynxosvtp["object_type"] = item.ObjectType
	niatelemetrynxosvtp["oper_mode"] = item.OperMode
	niatelemetrynxosvtp["pruning_mode"] = item.PruningMode
	niatelemetrynxosvtp["running_version"] = item.RunningVersion
	niatelemetrynxosvtp["trap_enabled"] = item.TrapEnabled
	niatelemetrynxosvtp["v2_mode"] = item.V2Mode
	niatelemetrynxosvtp["nr_version"] = item.Version

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
	niatelemetrysmartlicense["active_mode"] = item.ActiveMode
	niatelemetrysmartlicense["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrysmartlicense["auth_status"] = item.AuthStatus
	niatelemetrysmartlicense["class_id"] = item.ClassId
	niatelemetrysmartlicense["license_udi"] = item.LicenseUdi
	niatelemetrysmartlicense["object_type"] = item.ObjectType
	niatelemetrysmartlicense["smart_account"] = item.SmartAccount

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
	onpremschedule["class_id"] = item.ClassId
	onpremschedule["day_of_month"] = item.DayOfMonth
	onpremschedule["day_of_week"] = item.DayOfWeek
	onpremschedule["month_of_year"] = item.MonthOfYear
	onpremschedule["object_type"] = item.ObjectType
	onpremschedule["repeat_interval"] = item.RepeatInterval
	onpremschedule["time_of_day"] = item.TimeOfDay
	onpremschedule["time_zone"] = item.TimeZone
	onpremschedule["week_of_month"] = item.WeekOfMonth

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
	onpremupgradephase["class_id"] = item.ClassId
	onpremupgradephase["elapsed_time"] = item.ElapsedTime
	onpremupgradephase["end_time"] = item.EndTime
	onpremupgradephase["failed"] = item.Failed
	onpremupgradephase["message"] = item.Message
	onpremupgradephase["name"] = item.Name
	onpremupgradephase["object_type"] = item.ObjectType
	onpremupgradephase["start_time"] = item.StartTime

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
	organizationorganizationrelationship["class_id"] = item.ClassId
	organizationorganizationrelationship["moid"] = item.Moid
	organizationorganizationrelationship["object_type"] = item.ObjectType
	organizationorganizationrelationship["selector"] = item.Selector

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
	answer_file_x := d.Get("answers").([]interface{})
	if len(answer_file_x) > 0 {
		answer_file_y := answer_file_x[0].(map[string]interface{})
		osanswers["answer_file"] = answer_file_y["answer_file"]
	}
	osanswers["class_id"] = item.ClassId
	osanswers["hostname"] = item.Hostname
	osanswers["ip_config_type"] = item.IpConfigType
	osanswers["ip_configuration"] = (func(p models.OsIpConfiguration, d *schema.ResourceData) []map[string]interface{} {
		var osipconfigurations []map[string]interface{}
		var ret models.OsIpConfiguration
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		osipconfiguration := make(map[string]interface{})
		osipconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osipconfiguration["class_id"] = item.ClassId
		osipconfiguration["object_type"] = item.ObjectType

		osipconfigurations = append(osipconfigurations, osipconfiguration)
		return osipconfigurations
	})(item.GetIpConfiguration(), d)
	osanswers["is_answer_file_set"] = item.IsAnswerFileSet
	osanswers["is_root_password_crypted"] = item.IsRootPasswordCrypted
	osanswers["is_root_password_set"] = item.IsRootPasswordSet
	osanswers["nameserver"] = item.Nameserver
	osanswers["object_type"] = item.ObjectType
	osanswers["product_key"] = item.ProductKey
	root_password_x := d.Get("answers").([]interface{})
	if len(root_password_x) > 0 {
		root_password_y := root_password_x[0].(map[string]interface{})
		osanswers["root_password"] = root_password_y["root_password"]
	}
	osanswers["nr_source"] = item.Source

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
	oscatalogrelationship["class_id"] = item.ClassId
	oscatalogrelationship["moid"] = item.Moid
	oscatalogrelationship["object_type"] = item.ObjectType
	oscatalogrelationship["selector"] = item.Selector

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
	osconfigurationfilerelationship["class_id"] = item.ClassId
	osconfigurationfilerelationship["moid"] = item.Moid
	osconfigurationfilerelationship["object_type"] = item.ObjectType
	osconfigurationfilerelationship["selector"] = item.Selector

	osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	return osconfigurationfilerelationships
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
	osinstalltarget["class_id"] = item.ClassId
	osinstalltarget["object_type"] = item.ObjectType

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
	osoperatingsystemparameters["class_id"] = item.ClassId
	osoperatingsystemparameters["object_type"] = item.ObjectType

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
	pciswitchrelationship["class_id"] = item.ClassId
	pciswitchrelationship["moid"] = item.Moid
	pciswitchrelationship["object_type"] = item.ObjectType
	pciswitchrelationship["selector"] = item.Selector

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
	pkixdistinguishedname["class_id"] = item.ClassId
	pkixdistinguishedname["common_name"] = item.CommonName
	pkixdistinguishedname["country"] = item.Country
	pkixdistinguishedname["locality"] = item.Locality
	pkixdistinguishedname["object_type"] = item.ObjectType
	pkixdistinguishedname["organization"] = item.Organization
	pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
	pkixdistinguishedname["state"] = item.State

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
	pkixkeygenerationspec["class_id"] = item.ClassId
	pkixkeygenerationspec["name"] = item.Name
	pkixkeygenerationspec["object_type"] = item.ObjectType

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
	pkixsubjectalternatename["class_id"] = item.ClassId
	pkixsubjectalternatename["dns_name"] = item.DnsName
	pkixsubjectalternatename["email_address"] = item.EmailAddress
	pkixsubjectalternatename["ip_address"] = item.IpAddress
	pkixsubjectalternatename["object_type"] = item.ObjectType
	pkixsubjectalternatename["uri"] = item.Uri

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
	policyabstractconfigprofilerelationship["class_id"] = item.ClassId
	policyabstractconfigprofilerelationship["moid"] = item.Moid
	policyabstractconfigprofilerelationship["object_type"] = item.ObjectType
	policyabstractconfigprofilerelationship["selector"] = item.Selector

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
	policyabstractprofilerelationship["class_id"] = item.ClassId
	policyabstractprofilerelationship["moid"] = item.Moid
	policyabstractprofilerelationship["object_type"] = item.ObjectType
	policyabstractprofilerelationship["selector"] = item.Selector

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
	policyconfigchange["changes"] = item.Changes
	policyconfigchange["class_id"] = item.ClassId
	policyconfigchange["disruptions"] = item.Disruptions
	policyconfigchange["object_type"] = item.ObjectType

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
	policyconfigchangecontext["class_id"] = item.ClassId
	policyconfigchangecontext["config_change_error"] = item.ConfigChangeError
	policyconfigchangecontext["config_change_state"] = item.ConfigChangeState
	policyconfigchangecontext["initial_config_context"] = (func(p models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {
		var policyconfigcontexts []map[string]interface{}
		var ret models.PolicyConfigContext
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		policyconfigcontext := make(map[string]interface{})
		policyconfigcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		policyconfigcontext["class_id"] = item.ClassId
		policyconfigcontext["config_state"] = item.ConfigState
		policyconfigcontext["control_action"] = item.ControlAction
		policyconfigcontext["error_state"] = item.ErrorState
		policyconfigcontext["object_type"] = item.ObjectType
		policyconfigcontext["oper_state"] = item.OperState

		policyconfigcontexts = append(policyconfigcontexts, policyconfigcontext)
		return policyconfigcontexts
	})(item.GetInitialConfigContext(), d)
	policyconfigchangecontext["object_type"] = item.ObjectType

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
	policyconfigcontext["class_id"] = item.ClassId
	policyconfigcontext["config_state"] = item.ConfigState
	policyconfigcontext["control_action"] = item.ControlAction
	policyconfigcontext["error_state"] = item.ErrorState
	policyconfigcontext["object_type"] = item.ObjectType
	policyconfigcontext["oper_state"] = item.OperState

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
	policyconfigresultcontext["class_id"] = item.ClassId
	policyconfigresultcontext["entity_data"] = item.EntityData
	policyconfigresultcontext["entity_moid"] = item.EntityMoid
	policyconfigresultcontext["entity_name"] = item.EntityName
	policyconfigresultcontext["entity_type"] = item.EntityType
	policyconfigresultcontext["object_type"] = item.ObjectType

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
	portgrouprelationship["class_id"] = item.ClassId
	portgrouprelationship["moid"] = item.Moid
	portgrouprelationship["object_type"] = item.ObjectType
	portgrouprelationship["selector"] = item.Selector

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
	portinterfacebaserelationship["class_id"] = item.ClassId
	portinterfacebaserelationship["moid"] = item.Moid
	portinterfacebaserelationship["object_type"] = item.ObjectType
	portinterfacebaserelationship["selector"] = item.Selector

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
	portsubgrouprelationship["class_id"] = item.ClassId
	portsubgrouprelationship["moid"] = item.Moid
	portsubgrouprelationship["object_type"] = item.ObjectType
	portsubgrouprelationship["selector"] = item.Selector

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
	recommendationcapacityrunwayrelationship["class_id"] = item.ClassId
	recommendationcapacityrunwayrelationship["moid"] = item.Moid
	recommendationcapacityrunwayrelationship["object_type"] = item.ObjectType
	recommendationcapacityrunwayrelationship["selector"] = item.Selector

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
	recoveryabstractbackupinforelationship["class_id"] = item.ClassId
	recoveryabstractbackupinforelationship["moid"] = item.Moid
	recoveryabstractbackupinforelationship["object_type"] = item.ObjectType
	recoveryabstractbackupinforelationship["selector"] = item.Selector

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
	recoverybackupconfigpolicyrelationship["class_id"] = item.ClassId
	recoverybackupconfigpolicyrelationship["moid"] = item.Moid
	recoverybackupconfigpolicyrelationship["object_type"] = item.ObjectType
	recoverybackupconfigpolicyrelationship["selector"] = item.Selector

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
	recoverybackupprofilerelationship["class_id"] = item.ClassId
	recoverybackupprofilerelationship["moid"] = item.Moid
	recoverybackupprofilerelationship["object_type"] = item.ObjectType
	recoverybackupprofilerelationship["selector"] = item.Selector

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
	recoverybackupschedule["class_id"] = item.ClassId
	recoverybackupschedule["execution_time"] = item.ExecutionTime
	recoverybackupschedule["frequency_unit"] = item.FrequencyUnit
	recoverybackupschedule["hours"] = item.Hours
	recoverybackupschedule["object_type"] = item.ObjectType

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
	recoveryconfigparams["class_id"] = item.ClassId
	recoveryconfigparams["object_type"] = item.ObjectType

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
	recoveryconfigresultrelationship["class_id"] = item.ClassId
	recoveryconfigresultrelationship["moid"] = item.Moid
	recoveryconfigresultrelationship["object_type"] = item.ObjectType
	recoveryconfigresultrelationship["selector"] = item.Selector

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
	recoveryscheduleconfigpolicyrelationship["class_id"] = item.ClassId
	recoveryscheduleconfigpolicyrelationship["moid"] = item.Moid
	recoveryscheduleconfigpolicyrelationship["object_type"] = item.ObjectType
	recoveryscheduleconfigpolicyrelationship["selector"] = item.Selector

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
	resourcegrouprelationship["class_id"] = item.ClassId
	resourcegrouprelationship["moid"] = item.Moid
	resourcegrouprelationship["object_type"] = item.ObjectType
	resourcegrouprelationship["selector"] = item.Selector

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
	resourcemembershipholderrelationship["class_id"] = item.ClassId
	resourcemembershipholderrelationship["moid"] = item.Moid
	resourcemembershipholderrelationship["object_type"] = item.ObjectType
	resourcemembershipholderrelationship["selector"] = item.Selector

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
	sdwanprofilerelationship["class_id"] = item.ClassId
	sdwanprofilerelationship["moid"] = item.Moid
	sdwanprofilerelationship["object_type"] = item.ObjectType
	sdwanprofilerelationship["selector"] = item.Selector

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
	sdwanrouterpolicyrelationship["class_id"] = item.ClassId
	sdwanrouterpolicyrelationship["moid"] = item.Moid
	sdwanrouterpolicyrelationship["object_type"] = item.ObjectType
	sdwanrouterpolicyrelationship["selector"] = item.Selector

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
	sdwanvmanageaccountpolicyrelationship["class_id"] = item.ClassId
	sdwanvmanageaccountpolicyrelationship["moid"] = item.Moid
	sdwanvmanageaccountpolicyrelationship["object_type"] = item.ObjectType
	sdwanvmanageaccountpolicyrelationship["selector"] = item.Selector

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
	serverconfigresultrelationship["class_id"] = item.ClassId
	serverconfigresultrelationship["moid"] = item.Moid
	serverconfigresultrelationship["object_type"] = item.ObjectType
	serverconfigresultrelationship["selector"] = item.Selector

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
	serverprofilerelationship["class_id"] = item.ClassId
	serverprofilerelationship["moid"] = item.Moid
	serverprofilerelationship["object_type"] = item.ObjectType
	serverprofilerelationship["selector"] = item.Selector

	serverprofilerelationships = append(serverprofilerelationships, serverprofilerelationship)
	return serverprofilerelationships
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
	softwarehyperflexdistributablerelationship["class_id"] = item.ClassId
	softwarehyperflexdistributablerelationship["moid"] = item.Moid
	softwarehyperflexdistributablerelationship["object_type"] = item.ObjectType
	softwarehyperflexdistributablerelationship["selector"] = item.Selector

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
	softwaresolutiondistributablerelationship["class_id"] = item.ClassId
	softwaresolutiondistributablerelationship["moid"] = item.Moid
	softwaresolutiondistributablerelationship["object_type"] = item.ObjectType
	softwaresolutiondistributablerelationship["selector"] = item.Selector

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
	softwarerepositorycatalogrelationship["class_id"] = item.ClassId
	softwarerepositorycatalogrelationship["moid"] = item.Moid
	softwarerepositorycatalogrelationship["object_type"] = item.ObjectType
	softwarerepositorycatalogrelationship["selector"] = item.Selector

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
	softwarerepositoryfilerelationship["class_id"] = item.ClassId
	softwarerepositoryfilerelationship["moid"] = item.Moid
	softwarerepositoryfilerelationship["object_type"] = item.ObjectType
	softwarerepositoryfilerelationship["selector"] = item.Selector

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
	softwarerepositoryfileserver["class_id"] = item.ClassId
	softwarerepositoryfileserver["object_type"] = item.ObjectType

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
	softwarerepositoryoperatingsystemfilerelationship["class_id"] = item.ClassId
	softwarerepositoryoperatingsystemfilerelationship["moid"] = item.Moid
	softwarerepositoryoperatingsystemfilerelationship["object_type"] = item.ObjectType
	softwarerepositoryoperatingsystemfilerelationship["selector"] = item.Selector

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
	softwarerepositoryreleaserelationship["class_id"] = item.ClassId
	softwarerepositoryreleaserelationship["moid"] = item.Moid
	softwarerepositoryreleaserelationship["object_type"] = item.ObjectType
	softwarerepositoryreleaserelationship["selector"] = item.Selector

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
	storagebasecapacity["available"] = item.Available
	storagebasecapacity["capacity_utilization"] = item.CapacityUtilization
	storagebasecapacity["class_id"] = item.ClassId
	storagebasecapacity["free"] = item.Free
	storagebasecapacity["object_type"] = item.ObjectType
	storagebasecapacity["total"] = item.Total
	storagebasecapacity["used"] = item.Used

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
	storagecontrollerrelationship["class_id"] = item.ClassId
	storagecontrollerrelationship["moid"] = item.Moid
	storagecontrollerrelationship["object_type"] = item.ObjectType
	storagecontrollerrelationship["selector"] = item.Selector

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
	storagediskgrouprelationship["class_id"] = item.ClassId
	storagediskgrouprelationship["moid"] = item.Moid
	storagediskgrouprelationship["object_type"] = item.ObjectType
	storagediskgrouprelationship["selector"] = item.Selector

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
	storagediskslotrelationship["class_id"] = item.ClassId
	storagediskslotrelationship["moid"] = item.Moid
	storagediskslotrelationship["object_type"] = item.ObjectType
	storagediskslotrelationship["selector"] = item.Selector

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
	storageenclosurerelationship["class_id"] = item.ClassId
	storageenclosurerelationship["moid"] = item.Moid
	storageenclosurerelationship["object_type"] = item.ObjectType
	storageenclosurerelationship["selector"] = item.Selector

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
	storageflexflashcontrollerrelationship["class_id"] = item.ClassId
	storageflexflashcontrollerrelationship["moid"] = item.Moid
	storageflexflashcontrollerrelationship["object_type"] = item.ObjectType
	storageflexflashcontrollerrelationship["selector"] = item.Selector

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
	storageflexutilcontrollerrelationship["class_id"] = item.ClassId
	storageflexutilcontrollerrelationship["moid"] = item.Moid
	storageflexutilcontrollerrelationship["object_type"] = item.ObjectType
	storageflexutilcontrollerrelationship["selector"] = item.Selector

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
	storagehitachiarrayrelationship["class_id"] = item.ClassId
	storagehitachiarrayrelationship["moid"] = item.Moid
	storagehitachiarrayrelationship["object_type"] = item.ObjectType
	storagehitachiarrayrelationship["selector"] = item.Selector

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
	storagehitachihostrelationship["class_id"] = item.ClassId
	storagehitachihostrelationship["moid"] = item.Moid
	storagehitachihostrelationship["object_type"] = item.ObjectType
	storagehitachihostrelationship["selector"] = item.Selector

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
	storagehitachiparitygrouprelationship["class_id"] = item.ClassId
	storagehitachiparitygrouprelationship["moid"] = item.Moid
	storagehitachiparitygrouprelationship["object_type"] = item.ObjectType
	storagehitachiparitygrouprelationship["selector"] = item.Selector

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
	storagehitachipoolrelationship["class_id"] = item.ClassId
	storagehitachipoolrelationship["moid"] = item.Moid
	storagehitachipoolrelationship["object_type"] = item.ObjectType
	storagehitachipoolrelationship["selector"] = item.Selector

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
	storagehitachivolumerelationship["class_id"] = item.ClassId
	storagehitachivolumerelationship["moid"] = item.Moid
	storagehitachivolumerelationship["object_type"] = item.ObjectType
	storagehitachivolumerelationship["selector"] = item.Selector

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
	storagehyperflexstoragecontainerrelationship["class_id"] = item.ClassId
	storagehyperflexstoragecontainerrelationship["moid"] = item.Moid
	storagehyperflexstoragecontainerrelationship["object_type"] = item.ObjectType
	storagehyperflexstoragecontainerrelationship["selector"] = item.Selector

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
	storagenetappclusterrelationship["class_id"] = item.ClassId
	storagenetappclusterrelationship["moid"] = item.Moid
	storagenetappclusterrelationship["object_type"] = item.ObjectType
	storagenetappclusterrelationship["selector"] = item.Selector

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
	storagenetappethernetportrelationship["class_id"] = item.ClassId
	storagenetappethernetportrelationship["moid"] = item.Moid
	storagenetappethernetportrelationship["object_type"] = item.ObjectType
	storagenetappethernetportrelationship["selector"] = item.Selector

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
	storagenetappfcportrelationship["class_id"] = item.ClassId
	storagenetappfcportrelationship["moid"] = item.Moid
	storagenetappfcportrelationship["object_type"] = item.ObjectType
	storagenetappfcportrelationship["selector"] = item.Selector

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
	storagenetapplunrelationship["class_id"] = item.ClassId
	storagenetapplunrelationship["moid"] = item.Moid
	storagenetapplunrelationship["object_type"] = item.ObjectType
	storagenetapplunrelationship["selector"] = item.Selector

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
	storagenetappnoderelationship["class_id"] = item.ClassId
	storagenetappnoderelationship["moid"] = item.Moid
	storagenetappnoderelationship["object_type"] = item.ObjectType
	storagenetappnoderelationship["selector"] = item.Selector

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
	storagenetappstoragevmrelationship["class_id"] = item.ClassId
	storagenetappstoragevmrelationship["moid"] = item.Moid
	storagenetappstoragevmrelationship["object_type"] = item.ObjectType
	storagenetappstoragevmrelationship["selector"] = item.Selector

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
	storagenetappvolumerelationship["class_id"] = item.ClassId
	storagenetappvolumerelationship["moid"] = item.Moid
	storagenetappvolumerelationship["object_type"] = item.ObjectType
	storagenetappvolumerelationship["selector"] = item.Selector

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
	storagephysicaldiskrelationship["class_id"] = item.ClassId
	storagephysicaldiskrelationship["moid"] = item.Moid
	storagephysicaldiskrelationship["object_type"] = item.ObjectType
	storagephysicaldiskrelationship["selector"] = item.Selector

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
	storagepurearrayrelationship["class_id"] = item.ClassId
	storagepurearrayrelationship["moid"] = item.Moid
	storagepurearrayrelationship["object_type"] = item.ObjectType
	storagepurearrayrelationship["selector"] = item.Selector

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
	storagepurecontrollerrelationship["class_id"] = item.ClassId
	storagepurecontrollerrelationship["moid"] = item.Moid
	storagepurecontrollerrelationship["object_type"] = item.ObjectType
	storagepurecontrollerrelationship["selector"] = item.Selector

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
	storagepurehostrelationship["class_id"] = item.ClassId
	storagepurehostrelationship["moid"] = item.Moid
	storagepurehostrelationship["object_type"] = item.ObjectType
	storagepurehostrelationship["selector"] = item.Selector

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
	storagepurehostgrouprelationship["class_id"] = item.ClassId
	storagepurehostgrouprelationship["moid"] = item.Moid
	storagepurehostgrouprelationship["object_type"] = item.ObjectType
	storagepurehostgrouprelationship["selector"] = item.Selector

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
	storagepureprotectiongrouprelationship["class_id"] = item.ClassId
	storagepureprotectiongrouprelationship["moid"] = item.Moid
	storagepureprotectiongrouprelationship["object_type"] = item.ObjectType
	storagepureprotectiongrouprelationship["selector"] = item.Selector

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
	storagepureprotectiongroupsnapshotrelationship["class_id"] = item.ClassId
	storagepureprotectiongroupsnapshotrelationship["moid"] = item.Moid
	storagepureprotectiongroupsnapshotrelationship["object_type"] = item.ObjectType
	storagepureprotectiongroupsnapshotrelationship["selector"] = item.Selector

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
	storagepurevolumerelationship["class_id"] = item.ClassId
	storagepurevolumerelationship["moid"] = item.Moid
	storagepurevolumerelationship["object_type"] = item.ObjectType
	storagepurevolumerelationship["selector"] = item.Selector

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
	storagesasexpanderrelationship["class_id"] = item.ClassId
	storagesasexpanderrelationship["moid"] = item.Moid
	storagesasexpanderrelationship["object_type"] = item.ObjectType
	storagesasexpanderrelationship["selector"] = item.Selector

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
	storagevirtualdriverelationship["class_id"] = item.ClassId
	storagevirtualdriverelationship["moid"] = item.Moid
	storagevirtualdriverelationship["object_type"] = item.ObjectType
	storagevirtualdriverelationship["selector"] = item.Selector

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
	storagevirtualdrivecontainerrelationship["class_id"] = item.ClassId
	storagevirtualdrivecontainerrelationship["moid"] = item.Moid
	storagevirtualdrivecontainerrelationship["object_type"] = item.ObjectType
	storagevirtualdrivecontainerrelationship["selector"] = item.Selector

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
	storagevirtualdriveextensionrelationship["class_id"] = item.ClassId
	storagevirtualdriveextensionrelationship["moid"] = item.Moid
	storagevirtualdriveextensionrelationship["object_type"] = item.ObjectType
	storagevirtualdriveextensionrelationship["selector"] = item.Selector

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
	tambaseadvisoryrelationship["class_id"] = item.ClassId
	tambaseadvisoryrelationship["moid"] = item.Moid
	tambaseadvisoryrelationship["object_type"] = item.ObjectType
	tambaseadvisoryrelationship["selector"] = item.Selector

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
	tambaseadvisorydetails["class_id"] = item.ClassId
	tambaseadvisorydetails["description"] = item.Description
	tambaseadvisorydetails["object_type"] = item.ObjectType

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
	tamseverity["class_id"] = item.ClassId
	tamseverity["object_type"] = item.ObjectType

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
	techsupportmanagementtechsupportbundlerelationship["class_id"] = item.ClassId
	techsupportmanagementtechsupportbundlerelationship["moid"] = item.Moid
	techsupportmanagementtechsupportbundlerelationship["object_type"] = item.ObjectType
	techsupportmanagementtechsupportbundlerelationship["selector"] = item.Selector

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
	techsupportmanagementtechsupportstatusrelationship["class_id"] = item.ClassId
	techsupportmanagementtechsupportstatusrelationship["moid"] = item.Moid
	techsupportmanagementtechsupportstatusrelationship["object_type"] = item.ObjectType
	techsupportmanagementtechsupportstatusrelationship["selector"] = item.Selector

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
	topsystemrelationship["class_id"] = item.ClassId
	topsystemrelationship["moid"] = item.Moid
	topsystemrelationship["object_type"] = item.ObjectType
	topsystemrelationship["selector"] = item.Selector

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
	uuidpoolblockrelationship["class_id"] = item.ClassId
	uuidpoolblockrelationship["moid"] = item.Moid
	uuidpoolblockrelationship["object_type"] = item.ObjectType
	uuidpoolblockrelationship["selector"] = item.Selector

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
	uuidpoolpoolrelationship["class_id"] = item.ClassId
	uuidpoolpoolrelationship["moid"] = item.Moid
	uuidpoolpoolrelationship["object_type"] = item.ObjectType
	uuidpoolpoolrelationship["selector"] = item.Selector

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
	uuidpoolpoolmemberrelationship["class_id"] = item.ClassId
	uuidpoolpoolmemberrelationship["moid"] = item.Moid
	uuidpoolpoolmemberrelationship["object_type"] = item.ObjectType
	uuidpoolpoolmemberrelationship["selector"] = item.Selector

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
	uuidpooluniverserelationship["class_id"] = item.ClassId
	uuidpooluniverserelationship["moid"] = item.Moid
	uuidpooluniverserelationship["object_type"] = item.ObjectType
	uuidpooluniverserelationship["selector"] = item.Selector

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
	uuidpooluuidblock["class_id"] = item.ClassId
	uuidpooluuidblock["from"] = item.From
	uuidpooluuidblock["object_type"] = item.ObjectType
	uuidpooluuidblock["size"] = item.Size
	uuidpooluuidblock["to"] = item.To

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
	uuidpooluuidleaserelationship["class_id"] = item.ClassId
	uuidpooluuidleaserelationship["moid"] = item.Moid
	uuidpooluuidleaserelationship["object_type"] = item.ObjectType
	uuidpooluuidleaserelationship["selector"] = item.Selector

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
	virtualizationactioninfo["class_id"] = item.ClassId
	virtualizationactioninfo["failure_reason"] = item.FailureReason
	virtualizationactioninfo["name"] = item.Name
	virtualizationactioninfo["object_type"] = item.ObjectType
	virtualizationactioninfo["status"] = item.Status

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
	virtualizationbaseclusterrelationship["class_id"] = item.ClassId
	virtualizationbaseclusterrelationship["moid"] = item.Moid
	virtualizationbaseclusterrelationship["object_type"] = item.ObjectType
	virtualizationbaseclusterrelationship["selector"] = item.Selector

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
	virtualizationbasehostrelationship["class_id"] = item.ClassId
	virtualizationbasehostrelationship["moid"] = item.Moid
	virtualizationbasehostrelationship["object_type"] = item.ObjectType
	virtualizationbasehostrelationship["selector"] = item.Selector

	virtualizationbasehostrelationships = append(virtualizationbasehostrelationships, virtualizationbasehostrelationship)
	return virtualizationbasehostrelationships
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
	virtualizationbasevirtualdiskrelationship["class_id"] = item.ClassId
	virtualizationbasevirtualdiskrelationship["moid"] = item.Moid
	virtualizationbasevirtualdiskrelationship["object_type"] = item.ObjectType
	virtualizationbasevirtualdiskrelationship["selector"] = item.Selector

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
	virtualizationbasevirtualmachinerelationship["class_id"] = item.ClassId
	virtualizationbasevirtualmachinerelationship["moid"] = item.Moid
	virtualizationbasevirtualmachinerelationship["object_type"] = item.ObjectType
	virtualizationbasevirtualmachinerelationship["selector"] = item.Selector

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
	virtualizationbasevmconfiguration["class_id"] = item.ClassId
	virtualizationbasevmconfiguration["object_type"] = item.ObjectType

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
	virtualizationcloudinitconfig["class_id"] = item.ClassId
	virtualizationcloudinitconfig["config_type"] = item.ConfigType
	virtualizationcloudinitconfig["network_data"] = item.NetworkData
	virtualizationcloudinitconfig["network_data_base64_encoded"] = item.NetworkDataBase64Encoded
	virtualizationcloudinitconfig["object_type"] = item.ObjectType
	virtualizationcloudinitconfig["user_data"] = item.UserData
	virtualizationcloudinitconfig["user_data_base64_encoded"] = item.UserDataBase64Encoded

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
	virtualizationcomputecapacity["capacity"] = item.Capacity
	virtualizationcomputecapacity["class_id"] = item.ClassId
	virtualizationcomputecapacity["free"] = item.Free
	virtualizationcomputecapacity["object_type"] = item.ObjectType
	virtualizationcomputecapacity["used"] = item.Used

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
	virtualizationcpuallocation["class_id"] = item.ClassId
	virtualizationcpuallocation["free"] = item.Free
	virtualizationcpuallocation["object_type"] = item.ObjectType
	virtualizationcpuallocation["reserved"] = item.Reserved
	virtualizationcpuallocation["total"] = item.Total
	virtualizationcpuallocation["used"] = item.Used

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
	virtualizationcpuinfo["class_id"] = item.ClassId
	virtualizationcpuinfo["cores"] = item.Cores
	virtualizationcpuinfo["description"] = item.Description
	virtualizationcpuinfo["object_type"] = item.ObjectType
	virtualizationcpuinfo["sockets"] = item.Sockets
	virtualizationcpuinfo["speed"] = item.Speed
	virtualizationcpuinfo["vendor"] = item.Vendor

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
	virtualizationguestinfo["class_id"] = item.ClassId
	virtualizationguestinfo["hostname"] = item.Hostname
	virtualizationguestinfo["ip_address"] = item.IpAddress
	virtualizationguestinfo["name"] = item.Name
	virtualizationguestinfo["object_type"] = item.ObjectType
	virtualizationguestinfo["operating_system"] = item.OperatingSystem

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
	virtualizationmemoryallocation["class_id"] = item.ClassId
	virtualizationmemoryallocation["free"] = item.Free
	virtualizationmemoryallocation["object_type"] = item.ObjectType
	virtualizationmemoryallocation["reserved"] = item.Reserved
	virtualizationmemoryallocation["total"] = item.Total
	virtualizationmemoryallocation["used"] = item.Used

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
	virtualizationmemorycapacity["capacity"] = item.Capacity
	virtualizationmemorycapacity["class_id"] = item.ClassId
	virtualizationmemorycapacity["free"] = item.Free
	virtualizationmemorycapacity["object_type"] = item.ObjectType
	virtualizationmemorycapacity["used"] = item.Used

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
	virtualizationproductinfo["class_id"] = item.ClassId
	virtualizationproductinfo["object_type"] = item.ObjectType
	virtualizationproductinfo["product_name"] = item.ProductName
	virtualizationproductinfo["product_type"] = item.ProductType
	virtualizationproductinfo["product_vendor"] = item.ProductVendor
	virtualizationproductinfo["nr_version"] = item.Version

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
	virtualizationstoragecapacity["capacity"] = item.Capacity
	virtualizationstoragecapacity["class_id"] = item.ClassId
	virtualizationstoragecapacity["free"] = item.Free
	virtualizationstoragecapacity["object_type"] = item.ObjectType
	virtualizationstoragecapacity["used"] = item.Used

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
	virtualizationvirtualmachinerelationship["class_id"] = item.ClassId
	virtualizationvirtualmachinerelationship["moid"] = item.Moid
	virtualizationvirtualmachinerelationship["object_type"] = item.ObjectType
	virtualizationvirtualmachinerelationship["selector"] = item.Selector

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
	virtualizationvmwareclusterrelationship["class_id"] = item.ClassId
	virtualizationvmwareclusterrelationship["moid"] = item.Moid
	virtualizationvmwareclusterrelationship["object_type"] = item.ObjectType
	virtualizationvmwareclusterrelationship["selector"] = item.Selector

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
	virtualizationvmwaredatacenterrelationship["class_id"] = item.ClassId
	virtualizationvmwaredatacenterrelationship["moid"] = item.Moid
	virtualizationvmwaredatacenterrelationship["object_type"] = item.ObjectType
	virtualizationvmwaredatacenterrelationship["selector"] = item.Selector

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
	virtualizationvmwaredatastorerelationship["class_id"] = item.ClassId
	virtualizationvmwaredatastorerelationship["moid"] = item.Moid
	virtualizationvmwaredatastorerelationship["object_type"] = item.ObjectType
	virtualizationvmwaredatastorerelationship["selector"] = item.Selector

	virtualizationvmwaredatastorerelationships = append(virtualizationvmwaredatastorerelationships, virtualizationvmwaredatastorerelationship)
	return virtualizationvmwaredatastorerelationships
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
	virtualizationvmwarefolderrelationship["class_id"] = item.ClassId
	virtualizationvmwarefolderrelationship["moid"] = item.Moid
	virtualizationvmwarefolderrelationship["object_type"] = item.ObjectType
	virtualizationvmwarefolderrelationship["selector"] = item.Selector

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
	virtualizationvmwarehostrelationship["class_id"] = item.ClassId
	virtualizationvmwarehostrelationship["moid"] = item.Moid
	virtualizationvmwarehostrelationship["object_type"] = item.ObjectType
	virtualizationvmwarehostrelationship["selector"] = item.Selector

	virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	return virtualizationvmwarehostrelationships
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
	virtualizationvmwareremotedisplayinfo["class_id"] = item.ClassId
	virtualizationvmwareremotedisplayinfo["object_type"] = item.ObjectType
	virtualizationvmwareremotedisplayinfo["remote_display_password"] = item.RemoteDisplayPassword
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_key"] = item.RemoteDisplayVncKey
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_port"] = item.RemoteDisplayVncPort

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
	virtualizationvmwareresourceconsumption["class_id"] = item.ClassId
	virtualizationvmwareresourceconsumption["cpu_consumed"] = item.CpuConsumed
	virtualizationvmwareresourceconsumption["memory_consumed"] = item.MemoryConsumed
	virtualizationvmwareresourceconsumption["object_type"] = item.ObjectType

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
	virtualizationvmwaresharesinfo["class_id"] = item.ClassId
	virtualizationvmwaresharesinfo["level"] = item.Level
	virtualizationvmwaresharesinfo["object_type"] = item.ObjectType
	virtualizationvmwaresharesinfo["shares"] = item.Shares

	virtualizationvmwaresharesinfos = append(virtualizationvmwaresharesinfos, virtualizationvmwaresharesinfo)
	return virtualizationvmwaresharesinfos
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
	virtualizationvmwarevcenterrelationship["class_id"] = item.ClassId
	virtualizationvmwarevcenterrelationship["moid"] = item.Moid
	virtualizationvmwarevcenterrelationship["object_type"] = item.ObjectType
	virtualizationvmwarevcenterrelationship["selector"] = item.Selector

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
	virtualizationvmwarevirtualmachinerelationship["class_id"] = item.ClassId
	virtualizationvmwarevirtualmachinerelationship["moid"] = item.Moid
	virtualizationvmwarevirtualmachinerelationship["object_type"] = item.ObjectType
	virtualizationvmwarevirtualmachinerelationship["selector"] = item.Selector

	virtualizationvmwarevirtualmachinerelationships = append(virtualizationvmwarevirtualmachinerelationships, virtualizationvmwarevirtualmachinerelationship)
	return virtualizationvmwarevirtualmachinerelationships
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
	virtualizationvmwarevmcpushareinfo["class_id"] = item.ClassId
	virtualizationvmwarevmcpushareinfo["cpu_limit"] = item.CpuLimit
	virtualizationvmwarevmcpushareinfo["cpu_overhead_limit"] = item.CpuOverheadLimit
	virtualizationvmwarevmcpushareinfo["cpu_reservation"] = item.CpuReservation
	virtualizationvmwarevmcpushareinfo["cpu_shares"] = item.CpuShares
	virtualizationvmwarevmcpushareinfo["object_type"] = item.ObjectType

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
	virtualizationvmwarevmcpusocketinfo["class_id"] = item.ClassId
	virtualizationvmwarevmcpusocketinfo["cores_per_socket"] = item.CoresPerSocket
	virtualizationvmwarevmcpusocketinfo["num_cpus"] = item.NumCpus
	virtualizationvmwarevmcpusocketinfo["num_sockets"] = item.NumSockets
	virtualizationvmwarevmcpusocketinfo["object_type"] = item.ObjectType

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
	virtualizationvmwarevmdiskcommitinfo["class_id"] = item.ClassId
	virtualizationvmwarevmdiskcommitinfo["committed_disk"] = item.CommittedDisk
	virtualizationvmwarevmdiskcommitinfo["object_type"] = item.ObjectType
	virtualizationvmwarevmdiskcommitinfo["un_committed_disk"] = item.UnCommittedDisk
	virtualizationvmwarevmdiskcommitinfo["unshared_disk"] = item.UnsharedDisk

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
	virtualizationvmwarevmmemoryshareinfo["class_id"] = item.ClassId
	virtualizationvmwarevmmemoryshareinfo["mem_limit"] = item.MemLimit
	virtualizationvmwarevmmemoryshareinfo["mem_overhead_limit"] = item.MemOverheadLimit
	virtualizationvmwarevmmemoryshareinfo["mem_reservation"] = item.MemReservation
	virtualizationvmwarevmmemoryshareinfo["mem_shares"] = item.MemShares
	virtualizationvmwarevmmemoryshareinfo["object_type"] = item.ObjectType

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
	vnicarfssettings["class_id"] = item.ClassId
	vnicarfssettings["enabled"] = item.Enabled
	vnicarfssettings["object_type"] = item.ObjectType

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
	vniccdn["class_id"] = item.ClassId
	vniccdn["object_type"] = item.ObjectType
	vniccdn["nr_source"] = item.Source
	vniccdn["value"] = item.Value

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
	vniccompletionqueuesettings["class_id"] = item.ClassId
	vniccompletionqueuesettings["nr_count"] = item.Count
	vniccompletionqueuesettings["object_type"] = item.ObjectType
	vniccompletionqueuesettings["ring_size"] = item.RingSize

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
	vnicethadapterpolicyrelationship["class_id"] = item.ClassId
	vnicethadapterpolicyrelationship["moid"] = item.Moid
	vnicethadapterpolicyrelationship["object_type"] = item.ObjectType
	vnicethadapterpolicyrelationship["selector"] = item.Selector

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
	vnicethifrelationship["class_id"] = item.ClassId
	vnicethifrelationship["moid"] = item.Moid
	vnicethifrelationship["object_type"] = item.ObjectType
	vnicethifrelationship["selector"] = item.Selector

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
	vnicethinterruptsettings["class_id"] = item.ClassId
	vnicethinterruptsettings["coalescing_time"] = item.CoalescingTime
	vnicethinterruptsettings["coalescing_type"] = item.CoalescingType
	vnicethinterruptsettings["nr_count"] = item.Count
	vnicethinterruptsettings["mode"] = item.Mode
	vnicethinterruptsettings["object_type"] = item.ObjectType

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
	vnicethnetworkpolicyrelationship["class_id"] = item.ClassId
	vnicethnetworkpolicyrelationship["moid"] = item.Moid
	vnicethnetworkpolicyrelationship["object_type"] = item.ObjectType
	vnicethnetworkpolicyrelationship["selector"] = item.Selector

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
	vnicethqospolicyrelationship["class_id"] = item.ClassId
	vnicethqospolicyrelationship["moid"] = item.Moid
	vnicethqospolicyrelationship["object_type"] = item.ObjectType
	vnicethqospolicyrelationship["selector"] = item.Selector

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
	vnicethrxqueuesettings["class_id"] = item.ClassId
	vnicethrxqueuesettings["nr_count"] = item.Count
	vnicethrxqueuesettings["object_type"] = item.ObjectType
	vnicethrxqueuesettings["ring_size"] = item.RingSize

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
	vnicethtxqueuesettings["class_id"] = item.ClassId
	vnicethtxqueuesettings["nr_count"] = item.Count
	vnicethtxqueuesettings["object_type"] = item.ObjectType
	vnicethtxqueuesettings["ring_size"] = item.RingSize

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
	vnicfcadapterpolicyrelationship["class_id"] = item.ClassId
	vnicfcadapterpolicyrelationship["moid"] = item.Moid
	vnicfcadapterpolicyrelationship["object_type"] = item.ObjectType
	vnicfcadapterpolicyrelationship["selector"] = item.Selector

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
	vnicfcerrorrecoverysettings["class_id"] = item.ClassId
	vnicfcerrorrecoverysettings["enabled"] = item.Enabled
	vnicfcerrorrecoverysettings["io_retry_count"] = item.IoRetryCount
	vnicfcerrorrecoverysettings["io_retry_timeout"] = item.IoRetryTimeout
	vnicfcerrorrecoverysettings["link_down_timeout"] = item.LinkDownTimeout
	vnicfcerrorrecoverysettings["object_type"] = item.ObjectType
	vnicfcerrorrecoverysettings["port_down_timeout"] = item.PortDownTimeout

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
	vnicfcifrelationship["class_id"] = item.ClassId
	vnicfcifrelationship["moid"] = item.Moid
	vnicfcifrelationship["object_type"] = item.ObjectType
	vnicfcifrelationship["selector"] = item.Selector

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
	vnicfcinterruptsettings["class_id"] = item.ClassId
	vnicfcinterruptsettings["mode"] = item.Mode
	vnicfcinterruptsettings["object_type"] = item.ObjectType

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
	vnicfcnetworkpolicyrelationship["class_id"] = item.ClassId
	vnicfcnetworkpolicyrelationship["moid"] = item.Moid
	vnicfcnetworkpolicyrelationship["object_type"] = item.ObjectType
	vnicfcnetworkpolicyrelationship["selector"] = item.Selector

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
	vnicfcqospolicyrelationship["class_id"] = item.ClassId
	vnicfcqospolicyrelationship["moid"] = item.Moid
	vnicfcqospolicyrelationship["object_type"] = item.ObjectType
	vnicfcqospolicyrelationship["selector"] = item.Selector

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
	vnicfcqueuesettings["class_id"] = item.ClassId
	vnicfcqueuesettings["nr_count"] = item.Count
	vnicfcqueuesettings["object_type"] = item.ObjectType
	vnicfcqueuesettings["ring_size"] = item.RingSize

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
	vnicflogisettings["class_id"] = item.ClassId
	vnicflogisettings["object_type"] = item.ObjectType
	vnicflogisettings["retries"] = item.Retries
	vnicflogisettings["timeout"] = item.Timeout

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
	vniciscsiadapterpolicyrelationship["class_id"] = item.ClassId
	vniciscsiadapterpolicyrelationship["moid"] = item.Moid
	vniciscsiadapterpolicyrelationship["object_type"] = item.ObjectType
	vniciscsiadapterpolicyrelationship["selector"] = item.Selector

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
	vniciscsiauthprofile["class_id"] = item.ClassId
	vniciscsiauthprofile["is_password_set"] = item.IsPasswordSet
	vniciscsiauthprofile["object_type"] = item.ObjectType
	password_x := d.Get("chap").([]interface{})
	if len(password_x) > 0 {
		password_y := password_x[0].(map[string]interface{})
		vniciscsiauthprofile["password"] = password_y["password"]
	}
	vniciscsiauthprofile["user_id"] = item.UserId

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
	vniciscsibootpolicyrelationship["class_id"] = item.ClassId
	vniciscsibootpolicyrelationship["moid"] = item.Moid
	vniciscsibootpolicyrelationship["object_type"] = item.ObjectType
	vniciscsibootpolicyrelationship["selector"] = item.Selector

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
	vniciscsistatictargetpolicyrelationship["class_id"] = item.ClassId
	vniciscsistatictargetpolicyrelationship["moid"] = item.Moid
	vniciscsistatictargetpolicyrelationship["object_type"] = item.ObjectType
	vniciscsistatictargetpolicyrelationship["selector"] = item.Selector

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
	vniclanconnectivitypolicyrelationship["class_id"] = item.ClassId
	vniclanconnectivitypolicyrelationship["moid"] = item.Moid
	vniclanconnectivitypolicyrelationship["object_type"] = item.ObjectType
	vniclanconnectivitypolicyrelationship["selector"] = item.Selector

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
	vniclun["bootable"] = item.Bootable
	vniclun["class_id"] = item.ClassId
	vniclun["lun_id"] = item.LunId
	vniclun["object_type"] = item.ObjectType

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
	vnicnvgresettings["class_id"] = item.ClassId
	vnicnvgresettings["enabled"] = item.Enabled
	vnicnvgresettings["object_type"] = item.ObjectType

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
	vnicplacementsettings["class_id"] = item.ClassId
	vnicplacementsettings["id"] = item.Id
	vnicplacementsettings["object_type"] = item.ObjectType
	vnicplacementsettings["pci_link"] = item.PciLink
	vnicplacementsettings["switch_id"] = item.SwitchId
	vnicplacementsettings["uplink"] = item.Uplink

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
	vnicplogisettings["class_id"] = item.ClassId
	vnicplogisettings["object_type"] = item.ObjectType
	vnicplogisettings["retries"] = item.Retries
	vnicplogisettings["timeout"] = item.Timeout

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
	vnicrocesettings["class_id"] = item.ClassId
	vnicrocesettings["class_of_service"] = item.ClassOfService
	vnicrocesettings["enabled"] = item.Enabled
	vnicrocesettings["memory_regions"] = item.MemoryRegions
	vnicrocesettings["object_type"] = item.ObjectType
	vnicrocesettings["queue_pairs"] = item.QueuePairs
	vnicrocesettings["resource_groups"] = item.ResourceGroups
	vnicrocesettings["nr_version"] = item.Version

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
	vnicrsshashsettings["class_id"] = item.ClassId
	vnicrsshashsettings["ipv4_hash"] = item.Ipv4Hash
	vnicrsshashsettings["ipv6_ext_hash"] = item.Ipv6ExtHash
	vnicrsshashsettings["ipv6_hash"] = item.Ipv6Hash
	vnicrsshashsettings["object_type"] = item.ObjectType
	vnicrsshashsettings["tcp_ipv4_hash"] = item.TcpIpv4Hash
	vnicrsshashsettings["tcp_ipv6_ext_hash"] = item.TcpIpv6ExtHash
	vnicrsshashsettings["tcp_ipv6_hash"] = item.TcpIpv6Hash
	vnicrsshashsettings["udp_ipv4_hash"] = item.UdpIpv4Hash
	vnicrsshashsettings["udp_ipv6_hash"] = item.UdpIpv6Hash

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
	vnicsanconnectivitypolicyrelationship["class_id"] = item.ClassId
	vnicsanconnectivitypolicyrelationship["moid"] = item.Moid
	vnicsanconnectivitypolicyrelationship["object_type"] = item.ObjectType
	vnicsanconnectivitypolicyrelationship["selector"] = item.Selector

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
	vnicscsiqueuesettings["class_id"] = item.ClassId
	vnicscsiqueuesettings["nr_count"] = item.Count
	vnicscsiqueuesettings["object_type"] = item.ObjectType
	vnicscsiqueuesettings["ring_size"] = item.RingSize

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
	vnictcpoffloadsettings["class_id"] = item.ClassId
	vnictcpoffloadsettings["large_receive"] = item.LargeReceive
	vnictcpoffloadsettings["large_send"] = item.LargeSend
	vnictcpoffloadsettings["object_type"] = item.ObjectType
	vnictcpoffloadsettings["rx_checksum"] = item.RxChecksum
	vnictcpoffloadsettings["tx_checksum"] = item.TxChecksum

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
	vnicusnicsettings["class_id"] = item.ClassId
	vnicusnicsettings["cos"] = item.Cos
	vnicusnicsettings["nr_count"] = item.Count
	vnicusnicsettings["object_type"] = item.ObjectType
	vnicusnicsettings["usnic_adapter_policy"] = item.UsnicAdapterPolicy

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
	vnicvlansettings["allowed_vlans"] = item.AllowedVlans
	vnicvlansettings["class_id"] = item.ClassId
	vnicvlansettings["default_vlan"] = item.DefaultVlan
	vnicvlansettings["mode"] = item.Mode
	vnicvlansettings["object_type"] = item.ObjectType

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
	vnicvmqsettings["class_id"] = item.ClassId
	vnicvmqsettings["enabled"] = item.Enabled
	vnicvmqsettings["multi_queue_support"] = item.MultiQueueSupport
	vnicvmqsettings["num_interrupts"] = item.NumInterrupts
	vnicvmqsettings["num_sub_vnics"] = item.NumSubVnics
	vnicvmqsettings["num_vmqs"] = item.NumVmqs
	vnicvmqsettings["object_type"] = item.ObjectType
	vnicvmqsettings["vmmq_adapter_policy"] = item.VmmqAdapterPolicy

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
	vnicvsansettings["class_id"] = item.ClassId
	vnicvsansettings["default_vlan_id"] = item.DefaultVlanId
	vnicvsansettings["id"] = item.Id
	vnicvsansettings["object_type"] = item.ObjectType

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
	vnicvxlansettings["class_id"] = item.ClassId
	vnicvxlansettings["enabled"] = item.Enabled
	vnicvxlansettings["object_type"] = item.ObjectType

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
	vrfvrfrelationship["class_id"] = item.ClassId
	vrfvrfrelationship["moid"] = item.Moid
	vrfvrfrelationship["object_type"] = item.ObjectType
	vrfvrfrelationship["selector"] = item.Selector

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
	workflowcatalogrelationship["class_id"] = item.ClassId
	workflowcatalogrelationship["moid"] = item.Moid
	workflowcatalogrelationship["object_type"] = item.ObjectType
	workflowcatalogrelationship["selector"] = item.Selector

	workflowcatalogrelationships = append(workflowcatalogrelationships, workflowcatalogrelationship)
	return workflowcatalogrelationships
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
	workflowerrorresponsehandlerrelationship["class_id"] = item.ClassId
	workflowerrorresponsehandlerrelationship["moid"] = item.Moid
	workflowerrorresponsehandlerrelationship["object_type"] = item.ObjectType
	workflowerrorresponsehandlerrelationship["selector"] = item.Selector

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
	workflowinternalproperties["base_task_type"] = item.BaseTaskType
	workflowinternalproperties["class_id"] = item.ClassId
	workflowinternalproperties["constraints"] = (func(p models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
		var workflowtaskconstraintss []map[string]interface{}
		var ret models.WorkflowTaskConstraints
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		workflowtaskconstraints := make(map[string]interface{})
		workflowtaskconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowtaskconstraints["class_id"] = item.ClassId
		workflowtaskconstraints["object_type"] = item.ObjectType
		workflowtaskconstraints["target_data_type"] = item.TargetDataType

		workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
		return workflowtaskconstraintss
	})(item.GetConstraints(), d)
	workflowinternalproperties["internal"] = item.Internal
	workflowinternalproperties["object_type"] = item.ObjectType
	workflowinternalproperties["owner"] = item.Owner

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
	workflowpendingdynamicworkflowinforelationship["class_id"] = item.ClassId
	workflowpendingdynamicworkflowinforelationship["moid"] = item.Moid
	workflowpendingdynamicworkflowinforelationship["object_type"] = item.ObjectType
	workflowpendingdynamicworkflowinforelationship["selector"] = item.Selector

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
	workflowproperties["class_id"] = item.ClassId
	workflowproperties["external_meta"] = item.ExternalMeta
	workflowproperties["input_definition"] = (func(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowbasedatatype["class_id"] = item.ClassId
			workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.ClassId
				workflowdefaultvalue["is_value_set"] = item.IsValueSet
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowbasedatatype["description"] = item.Description
			workflowbasedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
				var workflowdisplaymetas []map[string]interface{}
				var ret models.WorkflowDisplayMeta
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdisplaymeta := make(map[string]interface{})
				workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdisplaymeta["class_id"] = item.ClassId
				workflowdisplaymeta["inventory_selector"] = item.InventorySelector
				workflowdisplaymeta["object_type"] = item.ObjectType
				workflowdisplaymeta["widget_type"] = item.WidgetType

				workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
				return workflowdisplaymetas
			})(item.GetDisplayMeta(), d)
			workflowbasedatatype["input_parameters"] = item.InputParameters
			workflowbasedatatype["label"] = item.Label
			workflowbasedatatype["name"] = item.Name
			workflowbasedatatype["object_type"] = item.ObjectType
			workflowbasedatatype["required"] = item.Required
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.GetInputDefinition(), d)
	workflowproperties["object_type"] = item.ObjectType
	workflowproperties["output_definition"] = (func(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowbasedatatype["class_id"] = item.ClassId
			workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.ClassId
				workflowdefaultvalue["is_value_set"] = item.IsValueSet
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowbasedatatype["description"] = item.Description
			workflowbasedatatype["display_meta"] = (func(p models.WorkflowDisplayMeta, d *schema.ResourceData) []map[string]interface{} {
				var workflowdisplaymetas []map[string]interface{}
				var ret models.WorkflowDisplayMeta
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdisplaymeta := make(map[string]interface{})
				workflowdisplaymeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdisplaymeta["class_id"] = item.ClassId
				workflowdisplaymeta["inventory_selector"] = item.InventorySelector
				workflowdisplaymeta["object_type"] = item.ObjectType
				workflowdisplaymeta["widget_type"] = item.WidgetType

				workflowdisplaymetas = append(workflowdisplaymetas, workflowdisplaymeta)
				return workflowdisplaymetas
			})(item.GetDisplayMeta(), d)
			workflowbasedatatype["input_parameters"] = item.InputParameters
			workflowbasedatatype["label"] = item.Label
			workflowbasedatatype["name"] = item.Name
			workflowbasedatatype["object_type"] = item.ObjectType
			workflowbasedatatype["required"] = item.Required
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.GetOutputDefinition(), d)
	workflowproperties["retry_count"] = item.RetryCount
	workflowproperties["retry_delay"] = item.RetryDelay
	workflowproperties["retry_policy"] = item.RetryPolicy
	workflowproperties["support_status"] = item.SupportStatus
	workflowproperties["timeout"] = item.Timeout
	workflowproperties["timeout_policy"] = item.TimeoutPolicy

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
	workflowtaskconstraints["class_id"] = item.ClassId
	workflowtaskconstraints["object_type"] = item.ObjectType
	workflowtaskconstraints["target_data_type"] = item.TargetDataType

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
	workflowtaskdefinitionrelationship["class_id"] = item.ClassId
	workflowtaskdefinitionrelationship["moid"] = item.Moid
	workflowtaskdefinitionrelationship["object_type"] = item.ObjectType
	workflowtaskdefinitionrelationship["selector"] = item.Selector

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
	workflowtaskinforelationship["class_id"] = item.ClassId
	workflowtaskinforelationship["moid"] = item.Moid
	workflowtaskinforelationship["object_type"] = item.ObjectType
	workflowtaskinforelationship["selector"] = item.Selector

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
	workflowtaskmetadatarelationship["class_id"] = item.ClassId
	workflowtaskmetadatarelationship["moid"] = item.Moid
	workflowtaskmetadatarelationship["object_type"] = item.ObjectType
	workflowtaskmetadatarelationship["selector"] = item.Selector

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
	workflowvalidationinformation["class_id"] = item.ClassId
	workflowvalidationinformation["object_type"] = item.ObjectType
	workflowvalidationinformation["state"] = item.State
	workflowvalidationinformation["validation_error"] = (func(p []models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var workflowvalidationerrors []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowvalidationerror := make(map[string]interface{})
			workflowvalidationerror["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowvalidationerror["class_id"] = item.ClassId
			workflowvalidationerror["error_log"] = item.ErrorLog
			workflowvalidationerror["field"] = item.Field
			workflowvalidationerror["object_type"] = item.ObjectType
			workflowvalidationerror["task_name"] = item.TaskName
			workflowvalidationerror["transition_name"] = item.TransitionName
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
	workflowworkflowctx["class_id"] = item.ClassId
	workflowworkflowctx["initiator_ctx"] = (func(p models.WorkflowInitiatorContext, d *schema.ResourceData) []map[string]interface{} {
		var workflowinitiatorcontexts []map[string]interface{}
		var ret models.WorkflowInitiatorContext
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		workflowinitiatorcontext := make(map[string]interface{})
		workflowinitiatorcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowinitiatorcontext["class_id"] = item.ClassId
		workflowinitiatorcontext["initiator_moid"] = item.InitiatorMoid
		workflowinitiatorcontext["initiator_name"] = item.InitiatorName
		workflowinitiatorcontext["initiator_type"] = item.InitiatorType
		workflowinitiatorcontext["object_type"] = item.ObjectType

		workflowinitiatorcontexts = append(workflowinitiatorcontexts, workflowinitiatorcontext)
		return workflowinitiatorcontexts
	})(item.GetInitiatorCtx(), d)
	workflowworkflowctx["object_type"] = item.ObjectType
	workflowworkflowctx["target_ctx_list"] = (func(p []models.WorkflowTargetContext, d *schema.ResourceData) []map[string]interface{} {
		var workflowtargetcontexts []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowtargetcontext := make(map[string]interface{})
			workflowtargetcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowtargetcontext["class_id"] = item.ClassId
			workflowtargetcontext["object_type"] = item.ObjectType
			workflowtargetcontext["target_moid"] = item.TargetMoid
			workflowtargetcontext["target_name"] = item.TargetName
			workflowtargetcontext["target_type"] = item.TargetType
			workflowtargetcontexts = append(workflowtargetcontexts, workflowtargetcontext)
		}
		return workflowtargetcontexts
	})(item.GetTargetCtxList(), d)
	workflowworkflowctx["workflow_meta_name"] = item.WorkflowMetaName
	workflowworkflowctx["workflow_subtype"] = item.WorkflowSubtype
	workflowworkflowctx["workflow_type"] = item.WorkflowType

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
	workflowworkflowdefinitionrelationship["class_id"] = item.ClassId
	workflowworkflowdefinitionrelationship["moid"] = item.Moid
	workflowworkflowdefinitionrelationship["object_type"] = item.ObjectType
	workflowworkflowdefinitionrelationship["selector"] = item.Selector

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
	workflowworkflowinforelationship["class_id"] = item.ClassId
	workflowworkflowinforelationship["moid"] = item.Moid
	workflowworkflowinforelationship["object_type"] = item.ObjectType
	workflowworkflowinforelationship["selector"] = item.Selector

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
	workflowworkflowinfoproperties["class_id"] = item.ClassId
	workflowworkflowinfoproperties["object_type"] = item.ObjectType
	workflowworkflowinfoproperties["retryable"] = item.Retryable
	workflowworkflowinfoproperties["rollback_action"] = item.RollbackAction

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
	workflowworkflowmetadatarelationship["class_id"] = item.ClassId
	workflowworkflowmetadatarelationship["moid"] = item.Moid
	workflowworkflowmetadatarelationship["object_type"] = item.ObjectType
	workflowworkflowmetadatarelationship["selector"] = item.Selector

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
	workflowworkflowproperties["class_id"] = item.ClassId
	workflowworkflowproperties["external_meta"] = item.ExternalMeta
	workflowworkflowproperties["object_type"] = item.ObjectType
	workflowworkflowproperties["retryable"] = item.Retryable
	workflowworkflowproperties["support_status"] = item.SupportStatus

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
	x509certificate["class_id"] = item.ClassId
	x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		var ret models.PkixDistinguishedName
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		pkixdistinguishedname["class_id"] = item.ClassId
		pkixdistinguishedname["common_name"] = item.CommonName
		pkixdistinguishedname["country"] = item.Country
		pkixdistinguishedname["locality"] = item.Locality
		pkixdistinguishedname["object_type"] = item.ObjectType
		pkixdistinguishedname["organization"] = item.Organization
		pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
		pkixdistinguishedname["state"] = item.State

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.GetIssuer(), d)
	x509certificate["not_after"] = item.NotAfter
	x509certificate["not_before"] = item.NotBefore
	x509certificate["object_type"] = item.ObjectType
	x509certificate["pem_certificate"] = item.PemCertificate
	x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
	x509certificate["signature_algorithm"] = item.SignatureAlgorithm
	x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		var ret models.PkixDistinguishedName
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		pkixdistinguishedname["class_id"] = item.ClassId
		pkixdistinguishedname["common_name"] = item.CommonName
		pkixdistinguishedname["country"] = item.Country
		pkixdistinguishedname["locality"] = item.Locality
		pkixdistinguishedname["object_type"] = item.ObjectType
		pkixdistinguishedname["organization"] = item.Organization
		pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
		pkixdistinguishedname["state"] = item.State

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.GetSubject(), d)

	x509certificates = append(x509certificates, x509certificate)
	return x509certificates
}
