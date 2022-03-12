terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "1.0.22"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}

data "intersight_organization_organization" "default" {
  name = "default"
}


resource "intersight_smtp_policy" "smtp1" {
  enabled      = false
  name         = "smtp1"
  description  = "testing smtp policy"
  smtp_port    = 32
  min_severity = "critical"
  smtp_server  = "10.10.10.1"
  smtp_recipients = [
    "aw@cisco.com",
    "cy@cisco.com",
  "dz@cisco.com"]
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_ntp_policy" "ntp1" {
  name        = "ntp1"
  description = "test policy"
  enabled     = true
  ntp_servers = [
    "ntp.esl.cisco.com",
    "time-a-g.nist.gov",
    "time-b-g.nist.gov"
  ]
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_boot_precision_policy" "boot_precision1" {
  name                     = "boot_precision1"
  description              = "test policy"
  configured_boot_mode     = "Legacy"
  enforce_uefi_secure_boot = false
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  boot_devices {
    enabled     = true
    name        = "scu-device-hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
  boot_devices {
    enabled     = true
    name        = "NIIODCIMCDVD"
    object_type = "boot.VirtualMedia"
    additional_properties = jsonencode({
      Subtype = "cimc-mapped-dvd"
    })
  }
  boot_devices {
    enabled     = true
    name        = "hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
}

resource "intersight_ssh_policy" "ssh_policy1" {
  name        = "ssh_policy1"
  description = "ssh policy"
  enabled     = true
  port        = 22
  timeout     = 1800
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_bios_policy" "bios_policy1" {
  name                                  = "TEST_BIOS_POLICY"
  description                           = "Bios policy"
  acs_control_gpu1state                 = "disabled"
  acs_control_gpu2state                 = "disabled"
  acs_control_gpu3state                 = "disabled"
  acs_control_gpu4state                 = "disabled"
  acs_control_gpu5state                 = "disabled"
  acs_control_gpu6state                 = "disabled"
  acs_control_gpu7state                 = "disabled"
  acs_control_gpu8state                 = "disabled"
  acs_control_slot11state               = "disabled"
  acs_control_slot12state               = "disabled"
  acs_control_slot13state               = "disabled"
  acs_control_slot14state               = "disabled"
  cdn_support                           = "disabled"
  lom_port0state                        = "disabled"
  lom_port1state                        = "disabled"
  lom_port2state                        = "disabled"
  lom_port3state                        = "disabled"
  lom_ports_all_state                   = "disabled"
  pci_option_ro_ms                      = "disabled"
  pci_rom_clp                           = "disabled"
  slot10link_speed                      = "Auto"
  slot10state                           = "disabled"
  slot11link_speed                      = "Auto"
  slot11state                           = "disabled"
  slot12link_speed                      = "Auto"
  slot12state                           = "disabled"
  slot13state                           = "disabled"
  slot14state                           = "disabled"
  slot1link_speed                       = "Auto"
  slot1state                            = "disabled"
  slot2link_speed                       = "Auto"
  slot2state                            = "disabled"
  slot3link_speed                       = "GEN3"
  slot3state                            = "disabled"
  slot4link_speed                       = "Auto"
  slot4state                            = "disabled"
  slot5link_speed                       = "GEN2"
  slot5state                            = "disabled"
  slot6link_speed                       = "Auto"
  slot6state                            = "disabled"
  slot7link_speed                       = "GEN1"
  slot7state                            = "disabled"
  slot8link_speed                       = "Auto"
  slot8state                            = "disabled"
  slot9link_speed                       = "Auto"
  slot9state                            = "disabled"
  slot_flom_link_speed                  = "Auto"
  slot_front_nvme1link_speed            = "Auto"
  slot_front_nvme2link_speed            = "Auto"
  slot_front_slot5link_speed            = "Auto"
  slot_front_slot6link_speed            = "Auto"
  slot_gpu1state                        = "disabled"
  slot_gpu2state                        = "disabled"
  slot_gpu3state                        = "disabled"
  slot_gpu4state                        = "disabled"
  slot_gpu5state                        = "disabled"
  slot_gpu6state                        = "disabled"
  slot_gpu7state                        = "disabled"
  slot_gpu8state                        = "disabled"
  slot_hba_link_speed                   = "Auto"
  slot_hba_state                        = "disabled"
  slot_lom1link                         = "disabled"
  slot_lom2link                         = "disabled"
  slot_mezz_state                       = "disabled"
  slot_mlom_link_speed                  = "Auto"
  slot_mlom_state                       = "disabled"
  slot_mraid_link_speed                 = "Auto"
  slot_mraid_state                      = "disabled"
  slot_n10state                         = "disabled"
  slot_n11state                         = "disabled"
  slot_n12state                         = "disabled"
  slot_n13state                         = "disabled"
  slot_n14state                         = "disabled"
  slot_n15state                         = "disabled"
  slot_n16state                         = "disabled"
  slot_n17state                         = "disabled"
  slot_n18state                         = "disabled"
  slot_n19state                         = "disabled"
  slot_n1state                          = "disabled"
  slot_n20state                         = "disabled"
  slot_n21state                         = "disabled"
  slot_n22state                         = "disabled"
  slot_n23state                         = "disabled"
  slot_n24state                         = "disabled"
  slot_n2state                          = "disabled"
  slot_n3state                          = "disabled"
  slot_n4state                          = "disabled"
  slot_n5state                          = "disabled"
  slot_n6state                          = "disabled"
  slot_n7state                          = "disabled"
  slot_n8state                          = "disabled"
  slot_n9state                          = "disabled"
  slot_raid_link_speed                  = "Auto"
  slot_raid_state                       = "disabled"
  slot_rear_nvme1link_speed             = "Auto"
  slot_rear_nvme1state                  = "disabled"
  slot_rear_nvme2link_speed             = "Auto"
  slot_rear_nvme2state                  = "disabled"
  slot_rear_nvme3state                  = "disabled"
  slot_rear_nvme4state                  = "disabled"
  slot_rear_nvme5state                  = "disabled"
  slot_rear_nvme6state                  = "disabled"
  slot_rear_nvme7state                  = "disabled"
  slot_rear_nvme8state                  = "disabled"
  slot_riser1link_speed                 = "Auto"
  slot_riser1slot1link_speed            = "Auto"
  slot_riser1slot2link_speed            = "Auto"
  slot_riser1slot3link_speed            = "Auto"
  slot_riser2link_speed                 = "Auto"
  slot_riser2slot4link_speed            = "Auto"
  slot_riser2slot5link_speed            = "Auto"
  slot_riser2slot6link_speed            = "Auto"
  slot_sas_state                        = "disabled"
  slot_ssd_slot1link_speed              = "Auto"
  slot_ssd_slot2link_speed              = "Auto"
  adjacent_cache_line_prefetch          = "disabled"
  altitude                              = "300-m"
  auto_cc_state                         = "disabled"
  autonumous_cstate_enable              = "disabled"
  boot_performance_mode                 = "Max Performance"
  cbs_cmn_cpu_gen_downcore_ctrl         = "FOUR (2 + 2)"
  channel_inter_leave                   = "auto"
  closed_loop_therm_throtl              = "disabled"
  cmci_enable                           = "disabled"
  config_tdp                            = "disabled"
  core_multi_processing                 = "2"
  cpu_energy_performance                = "performance"
  cpu_frequency_floor                   = "disabled"
  cpu_performance                       = "custom"
  cpu_power_management                  = "custom"
  demand_scrub                          = "disabled"
  direct_cache_access                   = "disabled"
  dram_clock_throttling                 = "Auto"
  energy_efficient_turbo                = "disabled"
  eng_perf_tuning                       = "OS"
  enhanced_intel_speed_step_tech        = "disabled"
  epp_profile                           = "Power"
  execute_disable_bit                   = "disabled"
  extended_apic                         = "disabled"
  hardware_prefetch                     = "disabled"
  hwpm_enable                           = "HWPM Native Mode"
  imc_interleave                        = "1-way Interleave"
  intel_hyper_threading_tech            = "disabled"
  intel_speed_select                    = "Base"
  intel_turbo_boost_tech                = "disabled"
  intel_virtualization_technology       = "disabled"
  ioh_error_enable                      = "Yes"
  ip_prefetch                           = "disabled"
  kti_prefetch                          = "disabled"
  llc_prefetch                          = "disabled"
  memory_inter_leave                    = "disabled"
  package_cstate_limit                  = "Auto"
  patrol_scrub                          = "disabled"
  patrol_scrub_duration                 = "platform-default"
  pc_ie_ssd_hot_plug_support            = "disabled"
  processor_c1e                         = "disabled"
  processor_c3report                    = "disabled"
  processor_c6report                    = "disabled"
  processor_cstate                      = "disabled"
  pstate_coord_type                     = "HW ALL"
  pwr_perf_tuning                       = "bios"
  rank_inter_leave                      = "auto"
  single_pctl_enable                    = "Yes"
  smt_mode                              = "Auto"
  snc                                   = "disabled"
  streamer_prefetch                     = "disabled"
  svm_mode                              = "disabled"
  work_load_config                      = "Balanced"
  xpt_prefetch                          = "Auto"
  all_usb_devices                       = "disabled"
  legacy_usb_support                    = "disabled"
  make_device_non_bootable              = "disabled"
  pch_usb30mode                         = "disabled"
  usb_emul6064                          = "disabled"
  usb_port_front                        = "disabled"
  usb_port_internal                     = "disabled"
  usb_port_kvm                          = "disabled"
  usb_port_rear                         = "disabled"
  usb_port_sd_card                      = "disabled"
  usb_port_vmedia                       = "disabled"
  usb_xhci_support                      = "disabled"
  aspm_support                          = "Auto"
  ioh_resource                          = "IOH0 24k IOH1 40k"
  memory_mapped_io_above4gb             = "disabled"
  mmcfg_base                            = "2 GB"
  onboard10gbit_lom                     = "disabled"
  onboard_gbit_lom                      = "disabled"
  sr_iov                                = "disabled"
  vga_priority                          = "Onboard"
  assert_nmi_on_perr                    = "disabled"
  assert_nmi_on_serr                    = "disabled"
  baud_rate                             = "115200"
  cdn_enable                            = "disabled"
  cisco_adaptive_mem_training           = "disabled"
  cisco_debug_level                     = "Minimum"
  cisco_oprom_launch_optimization       = "disabled"
  console_redirection                   = "disabled"
  flow_control                          = "rts-cts"
  frb2enable                            = "disabled"
  legacy_os_redirection                 = "disabled"
  os_boot_watchdog_timer                = "disabled"
  os_boot_watchdog_timer_policy         = "power-off"
  os_boot_watchdog_timer_timeout        = "15-minutes"
  out_of_band_mgmt_port                 = "disabled"
  putty_key_pad                         = "LINUX"
  redirection_after_post                = "Always Enable"
  terminal_type                         = "vt100"
  ucsm_boot_order_rule                  = "Loose"
  bme_dma_mitigation                    = "disabled"
  cbs_cmn_gnb_nb_iommu                  = "disabled"
  cbs_cmn_mem_ctrl_bank_group_swap_ddr4 = "disabled"
  cbs_cmn_mem_map_bank_interleave_ddr4  = "Auto"
  cbs_df_cmn_mem_intlv                  = "Auto"
  cbs_df_cmn_mem_intlv_size             = "Auto"
  dcpmm_firmware_downgrade              = "disabled"
  smee                                  = "disabled"
  boot_option_num_retry                 = "13"
  boot_option_re_cool_down              = "90"
  boot_option_retry                     = "disabled"
  ipv6pxe                               = "disabled"
  onboard_scu_storage_support           = "disabled"
  onboard_scu_storage_sw_stack          = "Intel RSTe"
  pop_support                           = "disabled"
  psata                                 = "AHCI"
  sata_mode_select                      = "AHCI"
  vmd_enable                            = "disabled"
  cbs_cmn_cpu_cpb                       = "Auto"
  cbs_cmn_cpu_global_cstate_ctrl        = "Auto"
  cbs_cmn_cpu_l1stream_hw_prefetcher    = "Auto"
  cbs_cmn_cpu_l2stream_hw_prefetcher    = "Auto"
  cbs_cmn_determinism_slider            = "Auto"
  cbs_cmnc_tdp_ctl                      = "Auto"
  cke_low_policy                        = "auto"
  dram_refresh_rate                     = "2x"
  lv_ddr_mode                           = "auto"
  mirroring_mode                        = "inter-socket"
  numa_optimized                        = "disabled"
  select_memory_ras_configuration       = "mirror-mode-1lm"
  sparing_mode                          = "dimm-sparing"
  intel_vt_for_directed_io              = "disabled"
  intel_vtd_coherency_support           = "disabled"
  intel_vtd_interrupt_remapping         = "disabled"
  intel_vtd_pass_through_dma_support    = "disabled"
  intel_vtdats_support                  = "disabled"
  post_error_pause                      = "disabled"
  tpm_support                           = "disabled"
  qpi_link_frequency                    = "7.2-gt/s"
  qpi_snoop_mode                        = "auto"
  serial_port_aenable                   = "disabled"
  tpm_control                           = "disabled"
  txt_support                           = "disabled"
}

resource "intersight_vmedia_policy" "vmedia1" {
  name          = "vmedia1"
  description   = "demo vmedia policy"
  enabled       = true
  encryption    = true
  low_power_usb = true
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  mappings = [{
    additional_properties   = ""
    authentication_protocol = "none"
    class_id                = "vmedia.Mapping"
    device_type             = "cdd"
    file_location           = "infra-chx.auslab.cisco.com/software/linux/ubuntu-18.04.5-server-amd64.iso"
    host_name               = "infra-chx.auslab.cisco.com"
    is_password_set         = false
    mount_options           = "RO"
    mount_protocol          = "nfs"
    object_type             = "vmedia.Mapping"
    password                = ""
    remote_file             = "ubuntu-18.04.5-server-amd64.iso"
    remote_path             = "/iso/software/linux"
    sanitized_file_location = "infra-chx.auslab.cisco.com/software/linux/ubuntu-18.04.5-server-amd64.iso"
    username                = ""
    volume_name             = "IMC_DVD"
  }]
}

resource "intersight_deviceconnector_policy" "dcp1" {
  name            = "device_con1"
  description     = "demo device connector policy"
  lockout_enabled = true
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_ipmioverlan_policy" "ipmi1" {
  name        = "ipmi1"
  description = "demo ipmi policy"
  enabled     = true
  privilege   = "admin"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_server_profile" "test" {
  name = "testing-server-profile"
  target_platform   = "Standalone"
  server_assignment_mode = "Static"
  action = "Deploy"
  assigned_server {
    moid = "620e269e76752d32344f45aa"
    object_type = "compute.RackUnit"
  }
  policy_bucket {
    moid        = intersight_smtp_policy.smtp1.moid
    object_type = intersight_smtp_policy.smtp1.object_type
  }

  policy_bucket {
        moid = intersight_ntp_policy.ntp1.moid
        object_type = intersight_ntp_policy.ntp1.object_type
  }
    policy_bucket {
    moid        = intersight_boot_precision_policy.boot_precision1.moid
    object_type = intersight_boot_precision_policy.boot_precision1.object_type
  }
      policy_bucket {
    moid        = intersight_ssh_policy.ssh_policy1.moid
    object_type = intersight_ssh_policy.ssh_policy1.object_type
  }
  policy_bucket {
        moid = intersight_bios_policy.bios_policy1.moid
        object_type = intersight_bios_policy.bios_policy1.object_type
  }
  policy_bucket {
        moid = intersight_memory_persistent_memory_policy.configured_from_os1.moid
        object_type = intersight_memory_persistent_memory_policy.configured_from_os1.object_type
  }
  policy_bucket {
        moid = intersight_vmedia_policy.vmedia1.moid
        object_type = intersight_vmedia_policy.vmedia1.object_type
  }
    policy_bucket {
        moid = intersight_deviceconnector_policy.dcp1.moid
        object_type = intersight_deviceconnector_policy.dcp1.object_type
  }
   policy_bucket {
        moid = intersight_ipmioverlan_policy.ipmi1.moid
        object_type = intersight_ipmioverlan_policy.ipmi1.object_type
  }
     policy_bucket {
        moid = intersight_iam_ldap_policy.ldap1.moid
        object_type = intersight_iam_ldap_policy.ldap1.object_type
  }
   policy_bucket {
        moid = intersight_iam_end_point_user_policy.user_policy1.moid
        object_type = intersight_iam_end_point_user_policy.user_policy1.object_type
  }
  policy_bucket {
        moid = intersight_networkconfig_policy.network_config1.moid
        object_type = intersight_networkconfig_policy.network_config1.object_type
  }
  policy_bucket {
        moid = intersight_sol_policy.sol1.moid
        object_type = intersight_sol_policy.sol1.object_type
  }
  policy_bucket {
        moid = intersight_snmp_policy.disabled.moid
        object_type = intersight_snmp_policy.disabled.object_type
  }
  policy_bucket {
        moid = intersight_syslog_policy.syslog1.moid
        object_type = intersight_syslog_policy.syslog1.object_type
  }
  policy_bucket {
        moid = intersight_kvm_policy.kvm1.moid
        object_type = intersight_kvm_policy.kvm1.object_type
  }
  policy_bucket {
        moid = intersight_sdcard_policy.sdcard1.moid
        object_type = intersight_sdcard_policy.sdcard1.object_type
  }
  policy_bucket {
        moid = intersight_storage_storage_policy.storage_storage1.moid
        object_type = intersight_storage_storage_policy.storage_storage1.object_type
  }
  policy_bucket {
        moid = intersight_adapter_config_policy.adapter_config1.moid
        object_type = intersight_adapter_config_policy.adapter_config1.object_type
  }
  policy_bucket {
        moid = intersight_vnic_lan_connectivity_policy.vnic_lan1.moid
        object_type = intersight_vnic_lan_connectivity_policy.vnic_lan1.object_type
  }
  policy_bucket {
        moid = intersight_vnic_san_connectivity_policy.vnic_san1.moid
        object_type = intersight_vnic_san_connectivity_policy.vnic_san1.object_type
  }
}


resource "intersight_networkconfig_policy" "network_config1" {
  name                     = "network_config1"
  description              = "demo network configuration policy"
  enable_dynamic_dns       = false
  preferred_ipv6dns_server = null
  enable_ipv6              = false
  enable_ipv6dns_from_dhcp = false
  preferred_ipv4dns_server = "10.10.10.1"
  alternate_ipv4dns_server = "10.10.10.1"
  alternate_ipv6dns_server = null
  dynamic_dns_domain       = ""
  enable_ipv4dns_from_dhcp = false
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_vnic_san_connectivity_policy" "vnic_san1" {
  name = "vnic_san1"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_vnic_lan_connectivity_policy" "vnic_lan1" {
  name                = "vnic_lan1"
  description         = "demo vnic lan connectivity policy"
  iqn_allocation_type = "None"
  placement_mode      = "auto"
  target_platform     = "FIAttached"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_adapter_config_policy" "adapter_config1" {
  name        = "adapter_config1"
  description = "test policy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  settings {
    object_type = "adapter.AdapterConfig"
    slot_id     = "1"
    eth_settings {
      lldp_enabled = true
      object_type  = "adapter.EthSettings"
    }
    fc_settings {
      object_type = "adapter.FcSettings"
      fip_enabled = true
    }
  }
  settings {
    object_type = "adapter.AdapterConfig"
    slot_id     = "MLOM"
    eth_settings {
      object_type  = "adapter.EthSettings"
      lldp_enabled = true
    }
    fc_settings {
      object_type = "adapter.FcSettings"
      fip_enabled = true
    }
  }
}

resource "intersight_storage_storage_policy" "storage_storage1" {
  name               = "storage_storage_policy1"
  description        = "storage policy test"
  unused_disks_state = "UnconfiguredGood"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results[0].moid
  }
  use_jbod_for_vd_creation = false
  m2_virtual_drive {
    enable      = false
    object_type = "storage.M2VirtualDriveConfig"
  }
  raid0_drive {
    enable      = true
    drive_slots = "2"
    object_type = "storage.R0Drive"
    virtual_drive_policy {
      strip_size    = 64
      access_policy = "Default"
      read_policy   = "Default"
      write_policy  = "Default"
      drive_cache   = "Default"
      object_type   = "storage.VirtualDrivePolicy"
    }
  }
}

resource "intersight_sdcard_policy" "sdcard1" {
  name        = "sdcard1"
  description = "demo sd card policy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  partitions {
    type        = "OS"
    object_type = "sdcard.Partition"

    virtual_drives {
      enable      = true
      object_type = "sdcard.OperatingSystem"
      additional_properties = jsonencode({
        Name = "Hypervisor"
      })
    }
  }
}

resource "intersight_kvm_policy" "kvm1" {
  name                      = "kvm1"
  description               = "demo kvm policy"
  enabled                   = true
  maximum_sessions          = 3
  remote_port               = 2069
  enable_video_encryption   = true
  enable_local_server_video = true
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_syslog_policy" "syslog1" {
  name        = "syslog1"
  description = "demo syslog policy"
  local_clients {
    min_severity = "emergency"
    object_type  = "syslog.LocalFileLoggingClient"
  }
  remote_clients {
    enabled      = true
    hostname     = "10.10.10.10"
    port         = 514
    protocol     = "tcp"
    min_severity = "emergency"
    object_type  = "syslog.RemoteLoggingClient"
  }
  remote_clients {
    enabled      = true
    hostname     = "2001:0db8:0a0b:12f0:0000:0000:0000:0004"
    port         = 64000
    protocol     = "udp"
    min_severity = "emergency"
    object_type  = "syslog.RemoteLoggingClient"
  }
  organization {
    object_type = "organization.Organization"
	selector = "$filter=Name eq 'default'"
	}
}

resource "intersight_snmp_policy" "disabled" {
  name = "disabled"
  enabled    = false
  v2_enabled = false
  v3_enabled = false
}

resource "intersight_sol_policy" "sol1" {
  name        = "sol2"
  description = "demo serial over lan policy"
  enabled     = false
  baud_rate   = 9600
  com_port    = "com1"
  ssh_port    = 1096
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_iam_end_point_user_policy" "user_policy1" {
  name        = "user_policy1"
  description = "test policy"

  password_properties {
    enforce_strong_password  = true
    enable_password_expiry   = true
    password_expiry_duration = 50
    password_history         = 5
    notification_period      = 1
    grace_period             = 2
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_iam_ldap_policy" "ldap1" {
  name                   = "ldap1"
  description            = "test policy"
  enabled                = true
  enable_dns             = true
  user_search_precedence = "LocalUserDb"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  base_properties {
    attribute                  = "CiscoAvPair"
    base_dn                    = "DC=QATCSLABTPI02DC=ciscoDC=com"
    bind_dn                    = "CN=administratorCN=UsersDC=QATCSLABTPI02DC=ciscoDC=com"
    bind_method                = "Anonymous"
    domain                     = "QATCSLABTPI02.cisco.com"
    enable_encryption          = true
    enable_group_authorization = true
    filter                     = "sAMAccountName"
    group_attribute            = "memberOf"
    nested_group_search_depth  = 128
    timeout                    = 180
  }
  dns_parameters {
    nr_source     = "Extracted"
    search_forest = "xyz"
    search_domain = "abc"
  }
}
resource "intersight_memory_persistent_memory_policy" "configured_from_os1" {
  name = "configured_from_os1"
  management_mode = "configured-from-operating-system"
}

