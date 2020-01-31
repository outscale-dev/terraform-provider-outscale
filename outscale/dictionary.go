package outscale

import "github.com/hashicorp/terraform/helper/schema"

//Dictionary for the Outscale APIs maps the apis to their respective functions
type Dictionary map[string]ResourceMap

//ResourceMap maps a schema to their resource or datasource implementation
type ResourceMap map[string]SchemaFunc

//SchemaFunc maps a function that returns a schema
type SchemaFunc func() *schema.Resource

var resources Dictionary
var datasources Dictionary

func init() {
	resources = Dictionary{
		"oapi": ResourceMap{
			"outscale_vm":                      resourceOutscaleOApiVM,
			"outscale_firewall_rules_set":      resourceOutscaleOAPISecurityGroup,
			"outscale_security_group":          resourceOutscaleOAPISecurityGroup,
			"outscale_image":                   resourceOutscaleOAPIImage,
			"outscale_keypair":                 resourceOutscaleOAPIKeyPair,
			"outscale_public_ip":               resourceOutscaleOAPIPublicIP,
			"outscale_public_ip_link":          resourceOutscaleOAPIPublicIPLink,
			"outscale_volume":                  resourceOutscaleOAPIVolume,
			"outscale_volumes_link":            resourceOutscaleOAPIVolumeLink,
			"outscale_outbound_rule":           resourceOutscaleOAPIOutboundRule,
			"outscale_security_group_rule":     resourceOutscaleOAPIOutboundRule,
			"outscale_tag":                     resourceOutscaleOAPITags,
			"outscale_net_attributes":          resourceOutscaleOAPILinAttributes,
			"outscale_net":                     resourceOutscaleOAPINet,
			"outscale_internet_service":        resourceOutscaleOAPIInternetService,
			"outscale_internet_service_link":   resourceOutscaleOAPIInternetServiceLink,
			"outscale_nat_service":             resourceOutscaleOAPINatService,
			"outscale_subnet":                  resourceOutscaleOAPISubNet,
			"outscale_route":                   resourceOutscaleOAPIRoute,
			"outscale_route_table":             resourceOutscaleOAPIRouteTable,
			"outscale_route_table_link":        resourceOutscaleOAPILinkRouteTable,
			"outscale_snapshot":                resourceOutscaleOAPISnapshot,
			"outscale_keypair_importation":     resourceOutscaleOAPIKeyPairImportation,
			"outscale_image_launch_permission": resourceOutscaleOAPIImageLaunchPermission,
			"outscale_net_peering":             resourceOutscaleOAPILinPeeringConnection,
			"outscale_nic_private_ip":          resourceOutscaleOAPINetworkInterfacePrivateIP,
			"outscale_nic_link":                resourceOutscaleOAPINetworkInterfaceAttachment,
			"outscale_nic":                     resourceOutscaleOAPINic,
			"outscale_image_tasks":             resourceOutscaleOAPIImageTasks,
			"outscale_image_copy":              resourceOutscaleOAPIImageCopy,
			"outscale_snapshot_attributes":     resourcedOutscaleOAPISnapshotAttributes,
			"outscale_image_register":          resourceOutscaleOAPIImageRegister,
			"outscale_net_peering_acceptation": resourceOutscaleOAPILinPeeringConnectionAccepter,
		},
	}

	datasources = Dictionary{
		"oapi": ResourceMap{
			"outscale_vm":                  dataSourceOutscaleOAPIVM,
			"outscale_vms":                 datasourceOutscaleOApiVMS,
			"outscale_firewall_rules_sets": dataSourceOutscaleOAPISecurityGroups,
			"outscale_security_groups":     dataSourceOutscaleOAPISecurityGroups,
			"outscale_images":              dataSourceOutscaleOAPIImages,
			"outscale_firewall_rules_set":  dataSourceOutscaleOAPISecurityGroup,
			"outscale_security_group":      dataSourceOutscaleOAPISecurityGroup,
			"outscale_tag":                 dataSourceOutscaleOAPITag,
			"outscale_tags":                dataSourceOutscaleOAPITags,
			"outscale_volume":              datasourceOutscaleOAPIVolume,
			"outscale_volumes":             datasourceOutscaleOAPIVolumes,
			"outscale_keypair":             datasourceOutscaleOAPIKeyPair,
			"outscale_keypairs":            datasourceOutscaleOAPIKeyPairs,
			"outscale_internet_service":    datasourceOutscaleOAPIInternetService,
			"outscale_internet_services":   datasourceOutscaleOAPIInternetServices,
			"outscale_subnet":              dataSourceOutscaleOAPISubnet,
			"outscale_subnets":             dataSourceOutscaleOAPISubnets,
			"outscale_vm_state":            dataSourceOutscaleOAPIVMState,
			"outscale_vms_state":           dataSourceOutscaleOAPIVMSState,
			"outscale_net":                 dataSourceOutscaleOAPIVpc,
			"outscale_nets":                dataSourceOutscaleOAPIVpcs,
			"outscale_net_attributes":      dataSourceOutscaleOAPIVpcAttr,
			"outscale_route_table":         dataSourceOutscaleOAPIRouteTable,
			"outscale_route_tables":        dataSourceOutscaleOAPIRouteTables,
			"outscale_snapshot":            dataSourceOutscaleOAPISnapshot,
			"outscale_snapshots":           dataSourceOutscaleOAPISnapshots,
			"outscale_net_peering":         dataSourceOutscaleOAPILinPeeringConnection,
			"outscale_net_peerings":        dataSourceOutscaleOAPILinPeeringsConnection,
			"outscale_nic":                 dataSourceOutscaleOAPINic,
			"outscale_nics":                dataSourceOutscaleOAPINics,
			"outscale_image":               dataSourceOutscaleOAPIImage,
			"outscale_public_ip":           dataSourceOutscaleOAPIPublicIP,
			"outscale_public_ips":          dataSourceOutscaleOAPIPublicIPS,
			"outscale_nat_service":         dataSourceOutscaleOAPINatService,
			"outscale_nat_services":        dataSourceOutscaleOAPINatServices,
		},
	}
}

// GetResource ...
func GetResource(api, resource string) SchemaFunc {
	var a ResourceMap

	if _, ok := resources[api]; !ok {
		return nil
	}

	a = resources[api]

	if _, ok := a[resource]; !ok {
		return nil
	}
	return a[resource]
}

//GetDatasource receives the apu and the name of the datasource
//and returns the corrresponding
func GetDatasource(api, datasource string) SchemaFunc {
	var a ResourceMap
	if _, ok := datasources[api]; !ok {
		return nil
	}

	a = datasources[api]

	if _, ok := a[datasource]; !ok {
		return nil
	}
	return a[datasource]
}
