resource "intersight_kubernetes_virtual_machine_infra_config_policy" "vm_infra_config_policy" {
  name        = "TFC-Test-1"
  description = "VM Infrastructure Config Policy used for Kubernetes deployment to vSphere"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results[0].moid
  }
  vm_config {
    object_type = "kubernetes.EsxiVirtualMachineInfraConfig"
	interfaces = ["8", "4"]
    additional_properties = jsonencode({
      Cluster    =  "abc"
      Datastore  = "datastore"
      Passphrase =  "password123"
    })
  }
}
