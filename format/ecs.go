package format

// taking the full ecs yaml definitions from elastic and converting them automatically to golang struct and replacing yaml to json
// https://github.com/elastic/ecs/blob/main/generated/ecs/ecs_nested.yml
// https: //zhwt.github.io/yaml-to-go/

/* TODO
The inner struct is to cleanup the JSON, otherwise it would be 23000 lines of mostly repetative code. I'm using Name exclusively anyway, but for reusability for other projects, may need to just fold these into a generic struct that can be used by the ECS struct
*/

type inner struct {
	DashedName  string `json:"dashed_name"`
	Description string `json:"description"`
	Example     string `json:"example"`
	FlatName    string `json:"flat_name"`
	IgnoreAbove int    `json:"ignore_above"`
	Level       string `json:"level"`
	MultiFields []struct {
		FlatName string `json:"flat_name"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	} `json:"multi_fields"`
	Name             string        `json:"name"`
	Normalize        []interface{} `json:"normalize"`
	Pattern          string        `json:"pattern"`
	OutputFormat     string        `json:"output_format"`
	OutputPrecision  int           `json:"output_precision"`
	OriginalFieldset string        `json:"original_fieldset"`
	Short            string        `json:"short"`
	Type             string        `json:"type"`
}

type ecs struct {
	Agent struct {
		Description string `json:"description"`
		Fields      struct {
			AgentBuildOriginal struct {
				inner
			} `json:"agent.build.original"`
			AgentEphemeralID struct {
				inner
			} `json:"agent.ephemeral_id"`
			AgentID struct {
				inner
			} `json:"agent.id"`
			AgentName struct {
				inner
			} `json:"agent.name"`
			AgentType struct {
				inner
			} `json:"agent.type"`
			AgentVersion struct {
				inner
			} `json:"agent.version"`
		} `json:"fields"`
		Footnote string `json:"footnote"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Short    string `json:"short"`
		Title    string `json:"title"`
		Type     string `json:"type"`
	} `json:"agent"`
	As struct {
		Description string `json:"description"`
		Fields      struct {
			AsNumber struct {
				inner
			} `json:"as.number"`
			AsOrganizationName struct {
				inner
			} `json:"as.organization.name"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"as"`
	Base struct {
		Description string `json:"description"`
		Fields      struct {
			Timestamp struct {
				inner
			} `json:"@timestamp"`
			Labels struct {
				inner
			} `json:"labels"`
			Message struct {
				inner
			} `json:"message"`
			Tags struct {
				inner
			} `json:"tags"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Root   bool   `json:"root"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"base"`
	Client struct {
		Description string `json:"description"`
		Fields      struct {
			ClientAddress struct {
				inner
			} `json:"client.address"`
			ClientAsNumber struct {
				inner
			} `json:"client.as.number"`
			ClientAsOrganizationName struct {
				inner
			} `json:"client.as.organization.name"`
			ClientBytes struct {
				inner
			} `json:"client.bytes"`
			ClientDomain struct {
				inner
			} `json:"client.domain"`
			ClientGeoCityName struct {
				inner
			} `json:"client.geo.city_name"`
			ClientGeoContinentCode struct {
				inner
			} `json:"client.geo.continent_code"`
			ClientGeoContinentName struct {
				inner
			} `json:"client.geo.continent_name"`
			ClientGeoCountryIsoCode struct {
				inner
			} `json:"client.geo.country_iso_code"`
			ClientGeoCountryName struct {
				inner
			} `json:"client.geo.country_name"`
			ClientGeoLocation struct {
				inner
			} `json:"client.geo.location"`
			ClientGeoName struct {
				inner
			} `json:"client.geo.name"`
			ClientGeoPostalCode struct {
				inner
			} `json:"client.geo.postal_code"`
			ClientGeoRegionIsoCode struct {
				inner
			} `json:"client.geo.region_iso_code"`
			ClientGeoRegionName struct {
				inner
			} `json:"client.geo.region_name"`
			ClientGeoTimezone struct {
				inner
			} `json:"client.geo.timezone"`
			ClientIP struct {
				inner
			} `json:"client.ip"`
			ClientMac struct {
				inner
			} `json:"client.mac"`
			ClientNatIP struct {
				inner
			} `json:"client.nat.ip"`
			ClientNatPort struct {
				inner
			} `json:"client.nat.port"`
			ClientPackets struct {
				inner
			} `json:"client.packets"`
			ClientPort struct {
				inner
			} `json:"client.port"`
			ClientRegisteredDomain struct {
				inner
			} `json:"client.registered_domain"`
			ClientSubdomain struct {
				inner
			} `json:"client.subdomain"`
			ClientTopLevelDomain struct {
				inner
			} `json:"client.top_level_domain"`
			ClientUserDomain struct {
				inner
			} `json:"client.user.domain"`
			ClientUserEmail struct {
				inner
			} `json:"client.user.email"`
			ClientUserFullName struct {
				inner
			} `json:"client.user.full_name"`
			ClientUserGroupDomain struct {
				inner
			} `json:"client.user.group.domain"`
			ClientUserGroupID struct {
				inner
			} `json:"client.user.group.id"`
			ClientUserGroupName struct {
				inner
			} `json:"client.user.group.name"`
			ClientUserHash struct {
				inner
			} `json:"client.user.hash"`
			ClientUserID struct {
				inner
			} `json:"client.user.id"`
			ClientUserName struct {
				inner
			} `json:"client.user.name"`
			ClientUserRoles struct {
				inner
			} `json:"client.user.roles"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"client"`
	Cloud struct {
		Description string `json:"description"`
		Fields      struct {
			CloudAccountID struct {
				inner
			} `json:"cloud.account.id"`
			CloudAccountName struct {
				inner
			} `json:"cloud.account.name"`
			CloudAvailabilityZone struct {
				inner
			} `json:"cloud.availability_zone"`
			CloudInstanceID struct {
				inner
			} `json:"cloud.instance.id"`
			CloudInstanceName struct {
				inner
			} `json:"cloud.instance.name"`
			CloudMachineType struct {
				inner
			} `json:"cloud.machine.type"`
			CloudOriginAccountID struct {
				inner
			} `json:"cloud.origin.account.id"`
			CloudOriginAccountName struct {
				inner
			} `json:"cloud.origin.account.name"`
			CloudOriginAvailabilityZone struct {
				inner
			} `json:"cloud.origin.availability_zone"`
			CloudOriginInstanceID struct {
				inner
			} `json:"cloud.origin.instance.id"`
			CloudOriginInstanceName struct {
				inner
			} `json:"cloud.origin.instance.name"`
			CloudOriginMachineType struct {
				inner
			} `json:"cloud.origin.machine.type"`
			CloudOriginProjectID struct {
				inner
			} `json:"cloud.origin.project.id"`
			CloudOriginProjectName struct {
				inner
			} `json:"cloud.origin.project.name"`
			CloudOriginProvider struct {
				inner
			} `json:"cloud.origin.provider"`
			CloudOriginRegion struct {
				inner
			} `json:"cloud.origin.region"`
			CloudOriginServiceName struct {
				inner
			} `json:"cloud.origin.service.name"`
			CloudProjectID struct {
				inner
			} `json:"cloud.project.id"`
			CloudProjectName struct {
				inner
			} `json:"cloud.project.name"`
			CloudProvider struct {
				inner
			} `json:"cloud.provider"`
			CloudRegion struct {
				inner
			} `json:"cloud.region"`
			CloudServiceName struct {
				inner
			} `json:"cloud.service.name"`
			CloudTargetAccountID struct {
				inner
			} `json:"cloud.target.account.id"`
			CloudTargetAccountName struct {
				inner
			} `json:"cloud.target.account.name"`
			CloudTargetAvailabilityZone struct {
				inner
			} `json:"cloud.target.availability_zone"`
			CloudTargetInstanceID struct {
				inner
			} `json:"cloud.target.instance.id"`
			CloudTargetInstanceName struct {
				inner
			} `json:"cloud.target.instance.name"`
			CloudTargetMachineType struct {
				inner
			} `json:"cloud.target.machine.type"`
			CloudTargetProjectID struct {
				inner
			} `json:"cloud.target.project.id"`
			CloudTargetProjectName struct {
				inner
			} `json:"cloud.target.project.name"`
			CloudTargetProvider struct {
				inner
			} `json:"cloud.target.provider"`
			CloudTargetRegion struct {
				inner
			} `json:"cloud.target.region"`
			CloudTargetServiceName struct {
				inner
			} `json:"cloud.target.service.name"`
		} `json:"fields"`
		Footnote string   `json:"footnote"`
		Group    int      `json:"group"`
		Name     string   `json:"name"`
		Nestings []string `json:"nestings"`
		Prefix   string   `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As            string `json:"as"`
				At            string `json:"at"`
				Beta          string `json:"beta"`
				Full          string `json:"full"`
				ShortOverride string `json:"short_override"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		ReusedHere []struct {
			Beta       string `json:"beta"`
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"cloud"`
	CodeSignature struct {
		Description string `json:"description"`
		Fields      struct {
			CodeSignatureDigestAlgorithm struct {
				inner
			} `json:"code_signature.digest_algorithm"`
			CodeSignatureExists struct {
				inner
			} `json:"code_signature.exists"`
			CodeSignatureSigningID struct {
				inner
			} `json:"code_signature.signing_id"`
			CodeSignatureStatus struct {
				inner
			} `json:"code_signature.status"`
			CodeSignatureSubjectName struct {
				inner
			} `json:"code_signature.subject_name"`
			CodeSignatureTeamID struct {
				inner
			} `json:"code_signature.team_id"`
			CodeSignatureTimestamp struct {
				inner
			} `json:"code_signature.timestamp"`
			CodeSignatureTrusted struct {
				inner
			} `json:"code_signature.trusted"`
			CodeSignatureValid struct {
				inner
			} `json:"code_signature.valid"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"code_signature"`
	Container struct {
		Description string `json:"description"`
		Fields      struct {
			ContainerCPUUsage struct {
				inner
			} `json:"container.cpu.usage"`
			ContainerDiskReadBytes struct {
				inner
			} `json:"container.disk.read.bytes"`
			ContainerDiskWriteBytes struct {
				inner
			} `json:"container.disk.write.bytes"`
			ContainerID struct {
				inner
			} `json:"container.id"`
			ContainerImageHashAll struct {
				inner
			} `json:"container.image.hash.all"`
			ContainerImageName struct {
				inner
			} `json:"container.image.name"`
			ContainerImageTag struct {
				inner
			} `json:"container.image.tag"`
			ContainerLabels struct {
				inner
			} `json:"container.labels"`
			ContainerMemoryUsage struct {
				inner
			} `json:"container.memory.usage"`
			ContainerName struct {
				inner
			} `json:"container.name"`
			ContainerNetworkEgressBytes struct {
				inner
			} `json:"container.network.egress.bytes"`
			ContainerNetworkIngressBytes struct {
				inner
			} `json:"container.network.ingress.bytes"`
			ContainerRuntime struct {
				inner
			} `json:"container.runtime"`
			ContainerSecurityContextPrivileged struct {
				inner
			} `json:"container.security_context.privileged"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"container"`
	DataStream struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			DataStreamDataset struct {
				inner
			} `json:"data_stream.dataset"`
			DataStreamNamespace struct {
				inner
			} `json:"data_stream.namespace"`
			DataStreamType struct {
				inner
			} `json:"data_stream.type"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"data_stream"`
	Destination struct {
		Description string `json:"description"`
		Fields      struct {
			DestinationAddress struct {
				inner
			} `json:"destination.address"`
			DestinationAsNumber struct {
				inner
			} `json:"destination.as.number"`
			DestinationAsOrganizationName struct {
				inner
			} `json:"destination.as.organization.name"`
			DestinationBytes struct {
				inner
			} `json:"destination.bytes"`
			DestinationDomain struct {
				inner
			} `json:"destination.domain"`
			DestinationGeoCityName struct {
				inner
			} `json:"destination.geo.city_name"`
			DestinationGeoContinentCode struct {
				inner
			} `json:"destination.geo.continent_code"`
			DestinationGeoContinentName struct {
				inner
			} `json:"destination.geo.continent_name"`
			DestinationGeoCountryIsoCode struct {
				inner
			} `json:"destination.geo.country_iso_code"`
			DestinationGeoCountryName struct {
				inner
			} `json:"destination.geo.country_name"`
			DestinationGeoLocation struct {
				inner
			} `json:"destination.geo.location"`
			DestinationGeoName struct {
				inner
			} `json:"destination.geo.name"`
			DestinationGeoPostalCode struct {
				inner
			} `json:"destination.geo.postal_code"`
			DestinationGeoRegionIsoCode struct {
				inner
			} `json:"destination.geo.region_iso_code"`
			DestinationGeoRegionName struct {
				inner
			} `json:"destination.geo.region_name"`
			DestinationGeoTimezone struct {
				inner
			} `json:"destination.geo.timezone"`
			DestinationIP struct {
				inner
			} `json:"destination.ip"`
			DestinationMac struct {
				inner
			} `json:"destination.mac"`
			DestinationNatIP struct {
				inner
			} `json:"destination.nat.ip"`
			DestinationNatPort struct {
				inner
			} `json:"destination.nat.port"`
			DestinationPackets struct {
				inner
			} `json:"destination.packets"`
			DestinationPort struct {
				inner
			} `json:"destination.port"`
			DestinationRegisteredDomain struct {
				inner
			} `json:"destination.registered_domain"`
			DestinationSubdomain struct {
				inner
			} `json:"destination.subdomain"`
			DestinationTopLevelDomain struct {
				inner
			} `json:"destination.top_level_domain"`
			DestinationUserDomain struct {
				inner
			} `json:"destination.user.domain"`
			DestinationUserEmail struct {
				inner
			} `json:"destination.user.email"`
			DestinationUserFullName struct {
				inner
			} `json:"destination.user.full_name"`
			DestinationUserGroupDomain struct {
				inner
			} `json:"destination.user.group.domain"`
			DestinationUserGroupID struct {
				inner
			} `json:"destination.user.group.id"`
			DestinationUserGroupName struct {
				inner
			} `json:"destination.user.group.name"`
			DestinationUserHash struct {
				inner
			} `json:"destination.user.hash"`
			DestinationUserID struct {
				inner
			} `json:"destination.user.id"`
			DestinationUserName struct {
				inner
			} `json:"destination.user.name"`
			DestinationUserRoles struct {
				inner
			} `json:"destination.user.roles"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"destination"`
	Device struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			DeviceID struct {
				inner
			} `json:"device.id"`
			DeviceManufacturer struct {
				inner
			} `json:"device.manufacturer"`
			DeviceModelIdentifier struct {
				inner
			} `json:"device.model.identifier"`
			DeviceModelName struct {
				inner
			} `json:"device.model.name"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"device"`
	Dll struct {
		Description string `json:"description"`
		Fields      struct {
			DllCodeSignatureDigestAlgorithm struct {
				inner
			} `json:"dll.code_signature.digest_algorithm"`
			DllCodeSignatureExists struct {
				inner
			} `json:"dll.code_signature.exists"`
			DllCodeSignatureSigningID struct {
				inner
			} `json:"dll.code_signature.signing_id"`
			DllCodeSignatureStatus struct {
				inner
			} `json:"dll.code_signature.status"`
			DllCodeSignatureSubjectName struct {
				inner
			} `json:"dll.code_signature.subject_name"`
			DllCodeSignatureTeamID struct {
				inner
			} `json:"dll.code_signature.team_id"`
			DllCodeSignatureTimestamp struct {
				inner
			} `json:"dll.code_signature.timestamp"`
			DllCodeSignatureTrusted struct {
				inner
			} `json:"dll.code_signature.trusted"`
			DllCodeSignatureValid struct {
				inner
			} `json:"dll.code_signature.valid"`
			DllHashMd5 struct {
				inner
			} `json:"dll.hash.md5"`
			DllHashSha1 struct {
				inner
			} `json:"dll.hash.sha1"`
			DllHashSha256 struct {
				inner
			} `json:"dll.hash.sha256"`
			DllHashSha384 struct {
				inner
			} `json:"dll.hash.sha384"`
			DllHashSha512 struct {
				inner
			} `json:"dll.hash.sha512"`
			DllHashSsdeep struct {
				inner
			} `json:"dll.hash.ssdeep"`
			DllHashTlsh struct {
				inner
			} `json:"dll.hash.tlsh"`
			DllName struct {
				inner
			} `json:"dll.name"`
			DllPath struct {
				inner
			} `json:"dll.path"`
			DllPeArchitecture struct {
				inner
			} `json:"dll.pe.architecture"`
			DllPeCompany struct {
				inner
			} `json:"dll.pe.company"`
			DllPeDescription struct {
				inner
			} `json:"dll.pe.description"`
			DllPeFileVersion struct {
				inner
			} `json:"dll.pe.file_version"`
			DllPeGoImportHash struct {
				inner
			} `json:"dll.pe.go_import_hash"`
			DllPeGoImports struct {
				inner
			} `json:"dll.pe.go_imports"`
			DllPeGoImportsNamesEntropy struct {
				inner
			} `json:"dll.pe.go_imports_names_entropy"`
			DllPeGoImportsNamesVarEntropy struct {
				inner
			} `json:"dll.pe.go_imports_names_var_entropy"`
			DllPeGoStripped struct {
				inner
			} `json:"dll.pe.go_stripped"`
			DllPeImphash struct {
				inner
			} `json:"dll.pe.imphash"`
			DllPeImportHash struct {
				inner
			} `json:"dll.pe.import_hash"`
			DllPeImports struct {
				inner
			} `json:"dll.pe.imports"`
			DllPeImportsNamesEntropy struct {
				inner
			} `json:"dll.pe.imports_names_entropy"`
			DllPeImportsNamesVarEntropy struct {
				inner
			} `json:"dll.pe.imports_names_var_entropy"`
			DllPeOriginalFileName struct {
				inner
			} `json:"dll.pe.original_file_name"`
			DllPePehash struct {
				inner
			} `json:"dll.pe.pehash"`
			DllPeProduct struct {
				inner
			} `json:"dll.pe.product"`
			DllPeSections struct {
				inner
			} `json:"dll.pe.sections"`
			DllPeSectionsEntropy struct {
				inner
			} `json:"dll.pe.sections.entropy"`
			DllPeSectionsName struct {
				inner
			} `json:"dll.pe.sections.name"`
			DllPeSectionsPhysicalSize struct {
				inner
			} `json:"dll.pe.sections.physical_size"`
			DllPeSectionsVarEntropy struct {
				inner
			} `json:"dll.pe.sections.var_entropy"`
			DllPeSectionsVirtualSize struct {
				inner
			} `json:"dll.pe.sections.virtual_size"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"dll"`
	DNS struct {
		Description string `json:"description"`
		Fields      struct {
			DNSAnswers struct {
				inner
			} `json:"dns.answers"`
			DNSAnswersClass struct {
				inner
			} `json:"dns.answers.class"`
			DNSAnswersData struct {
				inner
			} `json:"dns.answers.data"`
			DNSAnswersName struct {
				inner
			} `json:"dns.answers.name"`
			DNSAnswersTTL struct {
				inner
			} `json:"dns.answers.ttl"`
			DNSAnswersType struct {
				inner
			} `json:"dns.answers.type"`
			DNSHeaderFlags struct {
				inner
			} `json:"dns.header_flags"`
			DNSID struct {
				inner
			} `json:"dns.id"`
			DNSOpCode struct {
				inner
			} `json:"dns.op_code"`
			DNSQuestionClass struct {
				inner
			} `json:"dns.question.class"`
			DNSQuestionName struct {
				inner
			} `json:"dns.question.name"`
			DNSQuestionRegisteredDomain struct {
				inner
			} `json:"dns.question.registered_domain"`
			DNSQuestionSubdomain struct {
				inner
			} `json:"dns.question.subdomain"`
			DNSQuestionTopLevelDomain struct {
				inner
			} `json:"dns.question.top_level_domain"`
			DNSQuestionType struct {
				inner
			} `json:"dns.question.type"`
			DNSResolvedIP struct {
				inner
			} `json:"dns.resolved_ip"`
			DNSResponseCode struct {
				inner
			} `json:"dns.response_code"`
			DNSType struct {
				inner
			} `json:"dns.type"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"dns"`
	Ecs struct {
		Description string `json:"description"`
		Fields      struct {
			EcsVersion struct {
				inner
			} `json:"ecs.version"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"ecs"`
	Elf struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			ElfArchitecture struct {
				inner
			} `json:"elf.architecture"`
			ElfByteOrder struct {
				inner
			} `json:"elf.byte_order"`
			ElfCPUType struct {
				inner
			} `json:"elf.cpu_type"`
			ElfCreationDate struct {
				inner
			} `json:"elf.creation_date"`
			ElfExports struct {
				inner
			} `json:"elf.exports"`
			ElfGoImportHash struct {
				inner
			} `json:"elf.go_import_hash"`
			ElfGoImports struct {
				inner
			} `json:"elf.go_imports"`
			ElfGoImportsNamesEntropy struct {
				inner
			} `json:"elf.go_imports_names_entropy"`
			ElfGoImportsNamesVarEntropy struct {
				inner
			} `json:"elf.go_imports_names_var_entropy"`
			ElfGoStripped struct {
				inner
			} `json:"elf.go_stripped"`
			ElfHeaderAbiVersion struct {
				inner
			} `json:"elf.header.abi_version"`
			ElfHeaderClass struct {
				inner
			} `json:"elf.header.class"`
			ElfHeaderData struct {
				inner
			} `json:"elf.header.data"`
			ElfHeaderEntrypoint struct {
				inner
			} `json:"elf.header.entrypoint"`
			ElfHeaderObjectVersion struct {
				inner
			} `json:"elf.header.object_version"`
			ElfHeaderOsAbi struct {
				inner
			} `json:"elf.header.os_abi"`
			ElfHeaderType struct {
				inner
			} `json:"elf.header.type"`
			ElfHeaderVersion struct {
				inner
			} `json:"elf.header.version"`
			ElfImportHash struct {
				inner
			} `json:"elf.import_hash"`
			ElfImports struct {
				inner
			} `json:"elf.imports"`
			ElfImportsNamesEntropy struct {
				inner
			} `json:"elf.imports_names_entropy"`
			ElfImportsNamesVarEntropy struct {
				inner
			} `json:"elf.imports_names_var_entropy"`
			ElfSections struct {
				inner
			} `json:"elf.sections"`
			ElfSectionsChi2 struct {
				inner
			} `json:"elf.sections.chi2"`
			ElfSectionsEntropy struct {
				inner
			} `json:"elf.sections.entropy"`
			ElfSectionsFlags struct {
				inner
			} `json:"elf.sections.flags"`
			ElfSectionsName struct {
				inner
			} `json:"elf.sections.name"`
			ElfSectionsPhysicalOffset struct {
				inner
			} `json:"elf.sections.physical_offset"`
			ElfSectionsPhysicalSize struct {
				inner
			} `json:"elf.sections.physical_size"`
			ElfSectionsType struct {
				inner
			} `json:"elf.sections.type"`
			ElfSectionsVarEntropy struct {
				inner
			} `json:"elf.sections.var_entropy"`
			ElfSectionsVirtualAddress struct {
				inner
			} `json:"elf.sections.virtual_address"`
			ElfSectionsVirtualSize struct {
				inner
			} `json:"elf.sections.virtual_size"`
			ElfSegments struct {
				inner
			} `json:"elf.segments"`
			ElfSegmentsSections struct {
				inner
			} `json:"elf.segments.sections"`
			ElfSegmentsType struct {
				inner
			} `json:"elf.segments.type"`
			ElfSharedLibraries struct {
				inner
			} `json:"elf.shared_libraries"`
			ElfTelfhash struct {
				inner
			} `json:"elf.telfhash"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Beta string `json:"beta"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"elf"`
	Email struct {
		Description string `json:"description"`
		Fields      struct {
			EmailAttachments struct {
				inner
			} `json:"email.attachments"`
			EmailAttachmentsFileExtension struct {
				inner
			} `json:"email.attachments.file.extension"`
			EmailAttachmentsFileHashMd5 struct {
				inner
			} `json:"email.attachments.file.hash.md5"`
			EmailAttachmentsFileHashSha1 struct {
				inner
			} `json:"email.attachments.file.hash.sha1"`
			EmailAttachmentsFileHashSha256 struct {
				inner
			} `json:"email.attachments.file.hash.sha256"`
			EmailAttachmentsFileHashSha384 struct {
				inner
			} `json:"email.attachments.file.hash.sha384"`
			EmailAttachmentsFileHashSha512 struct {
				inner
			} `json:"email.attachments.file.hash.sha512"`
			EmailAttachmentsFileHashSsdeep struct {
				inner
			} `json:"email.attachments.file.hash.ssdeep"`
			EmailAttachmentsFileHashTlsh struct {
				inner
			} `json:"email.attachments.file.hash.tlsh"`
			EmailAttachmentsFileMimeType struct {
				inner
			} `json:"email.attachments.file.mime_type"`
			EmailAttachmentsFileName struct {
				inner
			} `json:"email.attachments.file.name"`
			EmailAttachmentsFileSize struct {
				inner
			} `json:"email.attachments.file.size"`
			EmailBccAddress struct {
				inner
			} `json:"email.bcc.address"`
			EmailCcAddress struct {
				inner
			} `json:"email.cc.address"`
			EmailContentType struct {
				inner
			} `json:"email.content_type"`
			EmailDeliveryTimestamp struct {
				inner
			} `json:"email.delivery_timestamp"`
			EmailDirection struct {
				inner
			} `json:"email.direction"`
			EmailFromAddress struct {
				inner
			} `json:"email.from.address"`
			EmailLocalID struct {
				inner
			} `json:"email.local_id"`
			EmailMessageID struct {
				inner
			} `json:"email.message_id"`
			EmailOriginationTimestamp struct {
				inner
			} `json:"email.origination_timestamp"`
			EmailReplyToAddress struct {
				inner
			} `json:"email.reply_to.address"`
			EmailSenderAddress struct {
				inner
			} `json:"email.sender.address"`
			EmailSubject struct {
				inner
			} `json:"email.subject"`
			EmailToAddress struct {
				inner
			} `json:"email.to.address"`
			EmailXMailer struct {
				inner
			} `json:"email.x_mailer"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"email"`
	Error struct {
		Description string `json:"description"`
		Fields      struct {
			ErrorCode struct {
				inner
			} `json:"error.code"`
			ErrorID struct {
				inner
			} `json:"error.id"`
			ErrorMessage struct {
				inner
			} `json:"error.message"`
			ErrorStackTrace struct {
				inner
			} `json:"error.stack_trace"`
			ErrorType struct {
				inner
			} `json:"error.type"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"error"`
	Event struct {
		Description string `json:"description"`
		Fields      struct {
			EventAction struct {
				inner
			} `json:"event.action"`
			EventAgentIDStatus struct {
				inner
			} `json:"event.agent_id_status"`
			EventCategory struct {
				AllowedValues []struct {
					Description        string   `json:"description"`
					ExpectedEventTypes []string `json:"expected_event_types"`
					Name               string   `json:"name"`
				} `json:"allowed_values"`
				inner
			} `json:"event.category"`
			EventCode struct {
				inner
			} `json:"event.code"`
			EventCreated struct {
				inner
			} `json:"event.created"`
			EventDataset struct {
				inner
			} `json:"event.dataset"`
			EventDuration struct {
				inner
			} `json:"event.duration"`
			EventEnd struct {
				inner
			} `json:"event.end"`
			EventHash struct {
				inner
			} `json:"event.hash"`
			EventID struct {
				inner
			} `json:"event.id"`
			EventIngested struct {
				inner
			} `json:"event.ingested"`
			EventKind struct {
				AllowedValues []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
					Beta        string `json:"beta,omitempty"`
				} `json:"allowed_values"`
				inner
			} `json:"event.kind"`
			EventModule struct {
				inner
			} `json:"event.module"`
			EventOriginal struct {
				inner
			} `json:"event.original"`
			EventOutcome struct {
				AllowedValues []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"allowed_values"`
				inner
			} `json:"event.outcome"`
			EventProvider struct {
				inner
			} `json:"event.provider"`
			EventReason struct {
				inner
			} `json:"event.reason"`
			EventReference struct {
				inner
			} `json:"event.reference"`
			EventRiskScore struct {
				inner
			} `json:"event.risk_score"`
			EventRiskScoreNorm struct {
				inner
			} `json:"event.risk_score_norm"`
			EventSequence struct {
				inner
			} `json:"event.sequence"`
			EventSeverity struct {
				inner
			} `json:"event.severity"`
			EventStart struct {
				inner
			} `json:"event.start"`
			EventTimezone struct {
				inner
			} `json:"event.timezone"`
			EventType struct {
				AllowedValues []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"allowed_values"`
				inner
			} `json:"event.type"`
			EventURL struct {
				inner
			} `json:"event.url"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"event"`
	Faas struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			FaasColdstart struct {
				inner
			} `json:"faas.coldstart"`
			FaasExecution struct {
				inner
			} `json:"faas.execution"`
			FaasID struct {
				inner
			} `json:"faas.id"`
			FaasName struct {
				inner
			} `json:"faas.name"`
			FaasTriggerRequestID struct {
				inner
			} `json:"faas.trigger.request_id"`
			FaasTriggerType struct {
				inner
			} `json:"faas.trigger.type"`
			FaasVersion struct {
				inner
			} `json:"faas.version"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"faas"`
	File struct {
		Description string `json:"description"`
		Fields      struct {
			FileAccessed struct {
				inner
			} `json:"file.accessed"`
			FileAttributes struct {
				inner
			} `json:"file.attributes"`
			FileCodeSignatureDigestAlgorithm struct {
				inner
			} `json:"file.code_signature.digest_algorithm"`
			FileCodeSignatureExists struct {
				inner
			} `json:"file.code_signature.exists"`
			FileCodeSignatureSigningID struct {
				inner
			} `json:"file.code_signature.signing_id"`
			FileCodeSignatureStatus struct {
				inner
			} `json:"file.code_signature.status"`
			FileCodeSignatureSubjectName struct {
				inner
			} `json:"file.code_signature.subject_name"`
			FileCodeSignatureTeamID struct {
				inner
			} `json:"file.code_signature.team_id"`
			FileCodeSignatureTimestamp struct {
				inner
			} `json:"file.code_signature.timestamp"`
			FileCodeSignatureTrusted struct {
				inner
			} `json:"file.code_signature.trusted"`
			FileCodeSignatureValid struct {
				inner
			} `json:"file.code_signature.valid"`
			FileCreated struct {
				inner
			} `json:"file.created"`
			FileCtime struct {
				inner
			} `json:"file.ctime"`
			FileDevice struct {
				inner
			} `json:"file.device"`
			FileDirectory struct {
				inner
			} `json:"file.directory"`
			FileDriveLetter struct {
				inner
			} `json:"file.drive_letter"`
			FileElfArchitecture struct {
				inner
			} `json:"file.elf.architecture"`
			FileElfByteOrder struct {
				inner
			} `json:"file.elf.byte_order"`
			FileElfCPUType struct {
				inner
			} `json:"file.elf.cpu_type"`
			FileElfCreationDate struct {
				inner
			} `json:"file.elf.creation_date"`
			FileElfExports struct {
				inner
			} `json:"file.elf.exports"`
			FileElfGoImportHash struct {
				inner
			} `json:"file.elf.go_import_hash"`
			FileElfGoImports struct {
				inner
			} `json:"file.elf.go_imports"`
			FileElfGoImportsNamesEntropy struct {
				inner
			} `json:"file.elf.go_imports_names_entropy"`
			FileElfGoImportsNamesVarEntropy struct {
				inner
			} `json:"file.elf.go_imports_names_var_entropy"`
			FileElfGoStripped struct {
				inner
			} `json:"file.elf.go_stripped"`
			FileElfHeaderAbiVersion struct {
				inner
			} `json:"file.elf.header.abi_version"`
			FileElfHeaderClass struct {
				inner
			} `json:"file.elf.header.class"`
			FileElfHeaderData struct {
				inner
			} `json:"file.elf.header.data"`
			FileElfHeaderEntrypoint struct {
				inner
			} `json:"file.elf.header.entrypoint"`
			FileElfHeaderObjectVersion struct {
				inner
			} `json:"file.elf.header.object_version"`
			FileElfHeaderOsAbi struct {
				inner
			} `json:"file.elf.header.os_abi"`
			FileElfHeaderType struct {
				inner
			} `json:"file.elf.header.type"`
			FileElfHeaderVersion struct {
				inner
			} `json:"file.elf.header.version"`
			FileElfImportHash struct {
				inner
			} `json:"file.elf.import_hash"`
			FileElfImports struct {
				inner
			} `json:"file.elf.imports"`
			FileElfImportsNamesEntropy struct {
				inner
			} `json:"file.elf.imports_names_entropy"`
			FileElfImportsNamesVarEntropy struct {
				inner
			} `json:"file.elf.imports_names_var_entropy"`
			FileElfSections struct {
				inner
			} `json:"file.elf.sections"`
			FileElfSectionsChi2 struct {
				inner
			} `json:"file.elf.sections.chi2"`
			FileElfSectionsEntropy struct {
				inner
			} `json:"file.elf.sections.entropy"`
			FileElfSectionsFlags struct {
				inner
			} `json:"file.elf.sections.flags"`
			FileElfSectionsName struct {
				inner
			} `json:"file.elf.sections.name"`
			FileElfSectionsPhysicalOffset struct {
				inner
			} `json:"file.elf.sections.physical_offset"`
			FileElfSectionsPhysicalSize struct {
				inner
			} `json:"file.elf.sections.physical_size"`
			FileElfSectionsType struct {
				inner
			} `json:"file.elf.sections.type"`
			FileElfSectionsVarEntropy struct {
				inner
			} `json:"file.elf.sections.var_entropy"`
			FileElfSectionsVirtualAddress struct {
				inner
			} `json:"file.elf.sections.virtual_address"`
			FileElfSectionsVirtualSize struct {
				inner
			} `json:"file.elf.sections.virtual_size"`
			FileElfSegments struct {
				inner
			} `json:"file.elf.segments"`
			FileElfSegmentsSections struct {
				inner
			} `json:"file.elf.segments.sections"`
			FileElfSegmentsType struct {
				inner
			} `json:"file.elf.segments.type"`
			FileElfSharedLibraries struct {
				inner
			} `json:"file.elf.shared_libraries"`
			FileElfTelfhash struct {
				inner
			} `json:"file.elf.telfhash"`
			FileExtension struct {
				inner
			} `json:"file.extension"`
			FileForkName struct {
				inner
			} `json:"file.fork_name"`
			FileGid struct {
				inner
			} `json:"file.gid"`
			FileGroup struct {
				inner
			} `json:"file.group"`
			FileHashMd5 struct {
				inner
			} `json:"file.hash.md5"`
			FileHashSha1 struct {
				inner
			} `json:"file.hash.sha1"`
			FileHashSha256 struct {
				inner
			} `json:"file.hash.sha256"`
			FileHashSha384 struct {
				inner
			} `json:"file.hash.sha384"`
			FileHashSha512 struct {
				inner
			} `json:"file.hash.sha512"`
			FileHashSsdeep struct {
				inner
			} `json:"file.hash.ssdeep"`
			FileHashTlsh struct {
				inner
			} `json:"file.hash.tlsh"`
			FileInode struct {
				inner
			} `json:"file.inode"`
			FileMachoGoImportHash struct {
				inner
			} `json:"file.macho.go_import_hash"`
			FileMachoGoImports struct {
				inner
			} `json:"file.macho.go_imports"`
			FileMachoGoImportsNamesEntropy struct {
				inner
			} `json:"file.macho.go_imports_names_entropy"`
			FileMachoGoImportsNamesVarEntropy struct {
				inner
			} `json:"file.macho.go_imports_names_var_entropy"`
			FileMachoGoStripped struct {
				inner
			} `json:"file.macho.go_stripped"`
			FileMachoImportHash struct {
				inner
			} `json:"file.macho.import_hash"`
			FileMachoImports struct {
				inner
			} `json:"file.macho.imports"`
			FileMachoImportsNamesEntropy struct {
				inner
			} `json:"file.macho.imports_names_entropy"`
			FileMachoImportsNamesVarEntropy struct {
				inner
			} `json:"file.macho.imports_names_var_entropy"`
			FileMachoSections struct {
				inner
			} `json:"file.macho.sections"`
			FileMachoSectionsEntropy struct {
				inner
			} `json:"file.macho.sections.entropy"`
			FileMachoSectionsName struct {
				inner
			} `json:"file.macho.sections.name"`
			FileMachoSectionsPhysicalSize struct {
				inner
			} `json:"file.macho.sections.physical_size"`
			FileMachoSectionsVarEntropy struct {
				inner
			} `json:"file.macho.sections.var_entropy"`
			FileMachoSectionsVirtualSize struct {
				inner
			} `json:"file.macho.sections.virtual_size"`
			FileMachoSymhash struct {
				inner
			} `json:"file.macho.symhash"`
			FileMimeType struct {
				inner
			} `json:"file.mime_type"`
			FileMode struct {
				inner
			} `json:"file.mode"`
			FileMtime struct {
				inner
			} `json:"file.mtime"`
			FileName struct {
				inner
			} `json:"file.name"`
			FileOwner struct {
				inner
			} `json:"file.owner"`
			FilePath struct {
				inner
			} `json:"file.path"`
			FilePeArchitecture struct {
				inner
			} `json:"file.pe.architecture"`
			FilePeCompany struct {
				inner
			} `json:"file.pe.company"`
			FilePeDescription struct {
				inner
			} `json:"file.pe.description"`
			FilePeFileVersion struct {
				inner
			} `json:"file.pe.file_version"`
			FilePeGoImportHash struct {
				inner
			} `json:"file.pe.go_import_hash"`
			FilePeGoImports struct {
				inner
			} `json:"file.pe.go_imports"`
			FilePeGoImportsNamesEntropy struct {
				inner
			} `json:"file.pe.go_imports_names_entropy"`
			FilePeGoImportsNamesVarEntropy struct {
				inner
			} `json:"file.pe.go_imports_names_var_entropy"`
			FilePeGoStripped struct {
				inner
			} `json:"file.pe.go_stripped"`
			FilePeImphash struct {
				inner
			} `json:"file.pe.imphash"`
			FilePeImportHash struct {
				inner
			} `json:"file.pe.import_hash"`
			FilePeImports struct {
				inner
			} `json:"file.pe.imports"`
			FilePeImportsNamesEntropy struct {
				inner
			} `json:"file.pe.imports_names_entropy"`
			FilePeImportsNamesVarEntropy struct {
				inner
			} `json:"file.pe.imports_names_var_entropy"`
			FilePeOriginalFileName struct {
				inner
			} `json:"file.pe.original_file_name"`
			FilePePehash struct {
				inner
			} `json:"file.pe.pehash"`
			FilePeProduct struct {
				inner
			} `json:"file.pe.product"`
			FilePeSections struct {
				inner
			} `json:"file.pe.sections"`
			FilePeSectionsEntropy struct {
				inner
			} `json:"file.pe.sections.entropy"`
			FilePeSectionsName struct {
				inner
			} `json:"file.pe.sections.name"`
			FilePeSectionsPhysicalSize struct {
				inner
			} `json:"file.pe.sections.physical_size"`
			FilePeSectionsVarEntropy struct {
				inner
			} `json:"file.pe.sections.var_entropy"`
			FilePeSectionsVirtualSize struct {
				inner
			} `json:"file.pe.sections.virtual_size"`
			FileSize struct {
				inner
			} `json:"file.size"`
			FileTargetPath struct {
				inner
			} `json:"file.target_path"`
			FileType struct {
				inner
			} `json:"file.type"`
			FileUID struct {
				inner
			} `json:"file.uid"`
			FileX509AlternativeNames struct {
				inner
			} `json:"file.x509.alternative_names"`
			FileX509IssuerCommonName struct {
				inner
			} `json:"file.x509.issuer.common_name"`
			FileX509IssuerCountry struct {
				inner
			} `json:"file.x509.issuer.country"`
			FileX509IssuerDistinguishedName struct {
				inner
			} `json:"file.x509.issuer.distinguished_name"`
			FileX509IssuerLocality struct {
				inner
			} `json:"file.x509.issuer.locality"`
			FileX509IssuerOrganization struct {
				inner
			} `json:"file.x509.issuer.organization"`
			FileX509IssuerOrganizationalUnit struct {
				inner
			} `json:"file.x509.issuer.organizational_unit"`
			FileX509IssuerStateOrProvince struct {
				inner
			} `json:"file.x509.issuer.state_or_province"`
			FileX509NotAfter struct {
				inner
			} `json:"file.x509.not_after"`
			FileX509NotBefore struct {
				inner
			} `json:"file.x509.not_before"`
			FileX509PublicKeyAlgorithm struct {
				inner
			} `json:"file.x509.public_key_algorithm"`
			FileX509PublicKeyCurve struct {
				inner
			} `json:"file.x509.public_key_curve"`
			FileX509PublicKeyExponent struct {
				inner
			} `json:"file.x509.public_key_exponent"`
			FileX509PublicKeySize struct {
				inner
			} `json:"file.x509.public_key_size"`
			FileX509SerialNumber struct {
				inner
			} `json:"file.x509.serial_number"`
			FileX509SignatureAlgorithm struct {
				inner
			} `json:"file.x509.signature_algorithm"`
			FileX509SubjectCommonName struct {
				inner
			} `json:"file.x509.subject.common_name"`
			FileX509SubjectCountry struct {
				inner
			} `json:"file.x509.subject.country"`
			FileX509SubjectDistinguishedName struct {
				inner
			} `json:"file.x509.subject.distinguished_name"`
			FileX509SubjectLocality struct {
				inner
			} `json:"file.x509.subject.locality"`
			FileX509SubjectOrganization struct {
				inner
			} `json:"file.x509.subject.organization"`
			FileX509SubjectOrganizationalUnit struct {
				inner
			} `json:"file.x509.subject.organizational_unit"`
			FileX509SubjectStateOrProvince struct {
				inner
			} `json:"file.x509.subject.state_or_province"`
			FileX509VersionNumber struct {
				inner
			} `json:"file.x509.version_number"`
		} `json:"fields"`
		Group    int      `json:"group"`
		Name     string   `json:"name"`
		Nestings []string `json:"nestings"`
		Prefix   string   `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
			Beta       string `json:"beta,omitempty"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"file"`
	Geo struct {
		Description string `json:"description"`
		Fields      struct {
			GeoCityName struct {
				inner
			} `json:"geo.city_name"`
			GeoContinentCode struct {
				inner
			} `json:"geo.continent_code"`
			GeoContinentName struct {
				inner
			} `json:"geo.continent_name"`
			GeoCountryIsoCode struct {
				inner
			} `json:"geo.country_iso_code"`
			GeoCountryName struct {
				inner
			} `json:"geo.country_name"`
			GeoLocation struct {
				inner
			} `json:"geo.location"`
			GeoName struct {
				inner
			} `json:"geo.name"`
			GeoPostalCode struct {
				inner
			} `json:"geo.postal_code"`
			GeoRegionIsoCode struct {
				inner
			} `json:"geo.region_iso_code"`
			GeoRegionName struct {
				inner
			} `json:"geo.region_name"`
			GeoTimezone struct {
				inner
			} `json:"geo.timezone"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"geo"`
	Group struct {
		Description string `json:"description"`
		Fields      struct {
			GroupDomain struct {
				inner
			} `json:"group.domain"`
			GroupID struct {
				inner
			} `json:"group.id"`
			GroupName struct {
				inner
			} `json:"group.name"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As            string   `json:"as"`
				At            string   `json:"at"`
				Full          string   `json:"full"`
				ShortOverride string   `json:"short_override,omitempty"`
				Normalize     []string `json:"normalize,omitempty"`
				Beta          string   `json:"beta,omitempty"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"group"`
	Hash struct {
		Description string `json:"description"`
		Fields      struct {
			HashMd5 struct {
				inner
			} `json:"hash.md5"`
			HashSha1 struct {
				inner
			} `json:"hash.sha1"`
			HashSha256 struct {
				inner
			} `json:"hash.sha256"`
			HashSha384 struct {
				inner
			} `json:"hash.sha384"`
			HashSha512 struct {
				inner
			} `json:"hash.sha512"`
			HashSsdeep struct {
				inner
			} `json:"hash.ssdeep"`
			HashTlsh struct {
				inner
			} `json:"hash.tlsh"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"hash"`
	Host struct {
		Description string `json:"description"`
		Fields      struct {
			HostArchitecture struct {
				inner
			} `json:"host.architecture"`
			HostBootID struct {
				Beta string `json:"beta"`
				inner
			} `json:"host.boot.id"`
			HostCPUUsage struct {
				inner
			} `json:"host.cpu.usage"`
			HostDiskReadBytes struct {
				inner
			} `json:"host.disk.read.bytes"`
			HostDiskWriteBytes struct {
				inner
			} `json:"host.disk.write.bytes"`
			HostDomain struct {
				inner
			} `json:"host.domain"`
			HostGeoCityName struct {
				inner
			} `json:"host.geo.city_name"`
			HostGeoContinentCode struct {
				inner
			} `json:"host.geo.continent_code"`
			HostGeoContinentName struct {
				inner
			} `json:"host.geo.continent_name"`
			HostGeoCountryIsoCode struct {
				inner
			} `json:"host.geo.country_iso_code"`
			HostGeoCountryName struct {
				inner
			} `json:"host.geo.country_name"`
			HostGeoLocation struct {
				inner
			} `json:"host.geo.location"`
			HostGeoName struct {
				inner
			} `json:"host.geo.name"`
			HostGeoPostalCode struct {
				inner
			} `json:"host.geo.postal_code"`
			HostGeoRegionIsoCode struct {
				inner
			} `json:"host.geo.region_iso_code"`
			HostGeoRegionName struct {
				inner
			} `json:"host.geo.region_name"`
			HostGeoTimezone struct {
				inner
			} `json:"host.geo.timezone"`
			HostHostname struct {
				inner
			} `json:"host.hostname"`
			HostID struct {
				inner
			} `json:"host.id"`
			HostIP struct {
				inner
			} `json:"host.ip"`
			HostMac struct {
				inner
			} `json:"host.mac"`
			HostName struct {
				inner
			} `json:"host.name"`
			HostNetworkEgressBytes struct {
				inner
			} `json:"host.network.egress.bytes"`
			HostNetworkEgressPackets struct {
				inner
			} `json:"host.network.egress.packets"`
			HostNetworkIngressBytes struct {
				inner
			} `json:"host.network.ingress.bytes"`
			HostNetworkIngressPackets struct {
				inner
			} `json:"host.network.ingress.packets"`
			HostOsFamily struct {
				inner
			} `json:"host.os.family"`
			HostOsFull struct {
				inner
			} `json:"host.os.full"`
			HostOsKernel struct {
				inner
			} `json:"host.os.kernel"`
			HostOsName struct {
				inner
			} `json:"host.os.name"`
			HostOsPlatform struct {
				inner
			} `json:"host.os.platform"`
			HostOsType struct {
				inner
			} `json:"host.os.type"`
			HostOsVersion struct {
				inner
			} `json:"host.os.version"`
			HostPidNsIno struct {
				Beta string `json:"beta"`
				inner
			} `json:"host.pid_ns_ino"`
			HostRiskCalculatedLevel struct {
				inner
			} `json:"host.risk.calculated_level"`
			HostRiskCalculatedScore struct {
				inner
			} `json:"host.risk.calculated_score"`
			HostRiskCalculatedScoreNorm struct {
				inner
			} `json:"host.risk.calculated_score_norm"`
			HostRiskStaticLevel struct {
				inner
			} `json:"host.risk.static_level"`
			HostRiskStaticScore struct {
				inner
			} `json:"host.risk.static_score"`
			HostRiskStaticScoreNorm struct {
				inner
			} `json:"host.risk.static_score_norm"`
			HostType struct {
				inner
			} `json:"host.type"`
			HostUptime struct {
				inner
			} `json:"host.uptime"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"host"`
	HTTP struct {
		Description string `json:"description"`
		Fields      struct {
			HTTPRequestBodyBytes struct {
				inner
			} `json:"http.request.body.bytes"`
			HTTPRequestBodyContent struct {
				inner
			} `json:"http.request.body.content"`
			HTTPRequestBytes struct {
				inner
			} `json:"http.request.bytes"`
			HTTPRequestID struct {
				inner
			} `json:"http.request.id"`
			HTTPRequestMethod struct {
				inner
			} `json:"http.request.method"`
			HTTPRequestMimeType struct {
				inner
			} `json:"http.request.mime_type"`
			HTTPRequestReferrer struct {
				inner
			} `json:"http.request.referrer"`
			HTTPResponseBodyBytes struct {
				inner
			} `json:"http.response.body.bytes"`
			HTTPResponseBodyContent struct {
				inner
			} `json:"http.response.body.content"`
			HTTPResponseBytes struct {
				inner
			} `json:"http.response.bytes"`
			HTTPResponseMimeType struct {
				inner
			} `json:"http.response.mime_type"`
			HTTPResponseStatusCode struct {
				inner
			} `json:"http.response.status_code"`
			HTTPVersion struct {
				inner
			} `json:"http.version"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"http"`
	Interface struct {
		Description string `json:"description"`
		Fields      struct {
			InterfaceAlias struct {
				inner
			} `json:"interface.alias"`
			InterfaceID struct {
				inner
			} `json:"interface.id"`
			InterfaceName struct {
				inner
			} `json:"interface.name"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"interface"`
	Log struct {
		Description string `json:"description"`
		Fields      struct {
			LogFilePath struct {
				inner
			} `json:"log.file.path"`
			LogLevel struct {
				inner
			} `json:"log.level"`
			LogLogger struct {
				inner
			} `json:"log.logger"`
			LogOriginFileLine struct {
				inner
			} `json:"log.origin.file.line"`
			LogOriginFileName struct {
				inner
			} `json:"log.origin.file.name"`
			LogOriginFunction struct {
				inner
			} `json:"log.origin.function"`
			LogSyslog struct {
				inner
			} `json:"log.syslog"`
			LogSyslogAppname struct {
				inner
			} `json:"log.syslog.appname"`
			LogSyslogFacilityCode struct {
				inner
			} `json:"log.syslog.facility.code"`
			LogSyslogFacilityName struct {
				inner
			} `json:"log.syslog.facility.name"`
			LogSyslogHostname struct {
				inner
			} `json:"log.syslog.hostname"`
			LogSyslogMsgid struct {
				inner
			} `json:"log.syslog.msgid"`
			LogSyslogPriority struct {
				inner
			} `json:"log.syslog.priority"`
			LogSyslogProcid struct {
				inner
			} `json:"log.syslog.procid"`
			LogSyslogSeverityCode struct {
				inner
			} `json:"log.syslog.severity.code"`
			LogSyslogSeverityName struct {
				inner
			} `json:"log.syslog.severity.name"`
			LogSyslogStructuredData struct {
				inner
			} `json:"log.syslog.structured_data"`
			LogSyslogVersion struct {
				inner
			} `json:"log.syslog.version"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"log"`
	Macho struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			MachoGoImportHash struct {
				inner
			} `json:"macho.go_import_hash"`
			MachoGoImports struct {
				inner
			} `json:"macho.go_imports"`
			MachoGoImportsNamesEntropy struct {
				inner
			} `json:"macho.go_imports_names_entropy"`
			MachoGoImportsNamesVarEntropy struct {
				inner
			} `json:"macho.go_imports_names_var_entropy"`
			MachoGoStripped struct {
				inner
			} `json:"macho.go_stripped"`
			MachoImportHash struct {
				inner
			} `json:"macho.import_hash"`
			MachoImports struct {
				inner
			} `json:"macho.imports"`
			MachoImportsNamesEntropy struct {
				inner
			} `json:"macho.imports_names_entropy"`
			MachoImportsNamesVarEntropy struct {
				inner
			} `json:"macho.imports_names_var_entropy"`
			MachoSections struct {
				inner
			} `json:"macho.sections"`
			MachoSectionsEntropy struct {
				inner
			} `json:"macho.sections.entropy"`
			MachoSectionsName struct {
				inner
			} `json:"macho.sections.name"`
			MachoSectionsPhysicalSize struct {
				inner
			} `json:"macho.sections.physical_size"`
			MachoSectionsVarEntropy struct {
				inner
			} `json:"macho.sections.var_entropy"`
			MachoSectionsVirtualSize struct {
				inner
			} `json:"macho.sections.virtual_size"`
			MachoSymhash struct {
				inner
			} `json:"macho.symhash"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Beta string `json:"beta"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"macho"`
	Network struct {
		Description string `json:"description"`
		Fields      struct {
			NetworkApplication struct {
				inner
			} `json:"network.application"`
			NetworkBytes struct {
				inner
			} `json:"network.bytes"`
			NetworkCommunityID struct {
				inner
			} `json:"network.community_id"`
			NetworkDirection struct {
				inner
			} `json:"network.direction"`
			NetworkForwardedIP struct {
				inner
			} `json:"network.forwarded_ip"`
			NetworkIanaNumber struct {
				inner
			} `json:"network.iana_number"`
			NetworkInner struct {
				inner
			} `json:"network.inner"`
			NetworkInnerVlanID struct {
				inner
			} `json:"network.inner.vlan.id"`
			NetworkInnerVlanName struct {
				inner
			} `json:"network.inner.vlan.name"`
			NetworkName struct {
				inner
			} `json:"network.name"`
			NetworkPackets struct {
				inner
			} `json:"network.packets"`
			NetworkProtocol struct {
				inner
			} `json:"network.protocol"`
			NetworkTransport struct {
				inner
			} `json:"network.transport"`
			NetworkType struct {
				inner
			} `json:"network.type"`
			NetworkVlanID struct {
				inner
			} `json:"network.vlan.id"`
			NetworkVlanName struct {
				inner
			} `json:"network.vlan.name"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"network"`
	Observer struct {
		Description string `json:"description"`
		Fields      struct {
			ObserverEgress struct {
				inner
			} `json:"observer.egress"`
			ObserverEgressInterfaceAlias struct {
				inner
			} `json:"observer.egress.interface.alias"`
			ObserverEgressInterfaceID struct {
				inner
			} `json:"observer.egress.interface.id"`
			ObserverEgressInterfaceName struct {
				inner
			} `json:"observer.egress.interface.name"`
			ObserverEgressVlanID struct {
				inner
			} `json:"observer.egress.vlan.id"`
			ObserverEgressVlanName struct {
				inner
			} `json:"observer.egress.vlan.name"`
			ObserverEgressZone struct {
				inner
			} `json:"observer.egress.zone"`
			ObserverGeoCityName struct {
				inner
			} `json:"observer.geo.city_name"`
			ObserverGeoContinentCode struct {
				inner
			} `json:"observer.geo.continent_code"`
			ObserverGeoContinentName struct {
				inner
			} `json:"observer.geo.continent_name"`
			ObserverGeoCountryIsoCode struct {
				inner
			} `json:"observer.geo.country_iso_code"`
			ObserverGeoCountryName struct {
				inner
			} `json:"observer.geo.country_name"`
			ObserverGeoLocation struct {
				inner
			} `json:"observer.geo.location"`
			ObserverGeoName struct {
				inner
			} `json:"observer.geo.name"`
			ObserverGeoPostalCode struct {
				inner
			} `json:"observer.geo.postal_code"`
			ObserverGeoRegionIsoCode struct {
				inner
			} `json:"observer.geo.region_iso_code"`
			ObserverGeoRegionName struct {
				inner
			} `json:"observer.geo.region_name"`
			ObserverGeoTimezone struct {
				inner
			} `json:"observer.geo.timezone"`
			ObserverHostname struct {
				inner
			} `json:"observer.hostname"`
			ObserverIngress struct {
				inner
			} `json:"observer.ingress"`
			ObserverIngressInterfaceAlias struct {
				inner
			} `json:"observer.ingress.interface.alias"`
			ObserverIngressInterfaceID struct {
				inner
			} `json:"observer.ingress.interface.id"`
			ObserverIngressInterfaceName struct {
				inner
			} `json:"observer.ingress.interface.name"`
			ObserverIngressVlanID struct {
				inner
			} `json:"observer.ingress.vlan.id"`
			ObserverIngressVlanName struct {
				inner
			} `json:"observer.ingress.vlan.name"`
			ObserverIngressZone struct {
				inner
			} `json:"observer.ingress.zone"`
			ObserverIP struct {
				inner
			} `json:"observer.ip"`
			ObserverMac struct {
				inner
			} `json:"observer.mac"`
			ObserverName struct {
				inner
			} `json:"observer.name"`
			ObserverOsFamily struct {
				inner
			} `json:"observer.os.family"`
			ObserverOsFull struct {
				inner
			} `json:"observer.os.full"`
			ObserverOsKernel struct {
				inner
			} `json:"observer.os.kernel"`
			ObserverOsName struct {
				inner
			} `json:"observer.os.name"`
			ObserverOsPlatform struct {
				inner
			} `json:"observer.os.platform"`
			ObserverOsType struct {
				inner
			} `json:"observer.os.type"`
			ObserverOsVersion struct {
				inner
			} `json:"observer.os.version"`
			ObserverProduct struct {
				inner
			} `json:"observer.product"`
			ObserverSerialNumber struct {
				inner
			} `json:"observer.serial_number"`
			ObserverType struct {
				inner
			} `json:"observer.type"`
			ObserverVendor struct {
				inner
			} `json:"observer.vendor"`
			ObserverVersion struct {
				inner
			} `json:"observer.version"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"observer"`
	Orchestrator struct {
		Description string `json:"description"`
		Fields      struct {
			OrchestratorAPIVersion struct {
				inner
			} `json:"orchestrator.api_version"`
			OrchestratorClusterID struct {
				inner
			} `json:"orchestrator.cluster.id"`
			OrchestratorClusterName struct {
				inner
			} `json:"orchestrator.cluster.name"`
			OrchestratorClusterURL struct {
				inner
			} `json:"orchestrator.cluster.url"`
			OrchestratorClusterVersion struct {
				inner
			} `json:"orchestrator.cluster.version"`
			OrchestratorNamespace struct {
				inner
			} `json:"orchestrator.namespace"`
			OrchestratorOrganization struct {
				inner
			} `json:"orchestrator.organization"`
			OrchestratorResourceAnnotation struct {
				inner
			} `json:"orchestrator.resource.annotation"`
			OrchestratorResourceID struct {
				inner
			} `json:"orchestrator.resource.id"`
			OrchestratorResourceIP struct {
				inner
			} `json:"orchestrator.resource.ip"`
			OrchestratorResourceLabel struct {
				inner
			} `json:"orchestrator.resource.label"`
			OrchestratorResourceName struct {
				inner
			} `json:"orchestrator.resource.name"`
			OrchestratorResourceParentType struct {
				inner
			} `json:"orchestrator.resource.parent.type"`
			OrchestratorResourceType struct {
				inner
			} `json:"orchestrator.resource.type"`
			OrchestratorType struct {
				inner
			} `json:"orchestrator.type"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"orchestrator"`
	Organization struct {
		Description string `json:"description"`
		Fields      struct {
			OrganizationID struct {
				inner
			} `json:"organization.id"`
			OrganizationName struct {
				inner
			} `json:"organization.name"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"organization"`
	Os struct {
		Description string `json:"description"`
		Fields      struct {
			OsFamily struct {
				inner
			} `json:"os.family"`
			OsFull struct {
				inner
			} `json:"os.full"`
			OsKernel struct {
				inner
			} `json:"os.kernel"`
			OsName struct {
				inner
			} `json:"os.name"`
			OsPlatform struct {
				inner
			} `json:"os.platform"`
			OsType struct {
				inner
			} `json:"os.type"`
			OsVersion struct {
				inner
			} `json:"os.version"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"os"`
	Package struct {
		Description string `json:"description"`
		Fields      struct {
			PackageArchitecture struct {
				inner
			} `json:"package.architecture"`
			PackageBuildVersion struct {
				inner
			} `json:"package.build_version"`
			PackageChecksum struct {
				inner
			} `json:"package.checksum"`
			PackageDescription struct {
				inner
			} `json:"package.description"`
			PackageInstallScope struct {
				inner
			} `json:"package.install_scope"`
			PackageInstalled struct {
				inner
			} `json:"package.installed"`
			PackageLicense struct {
				inner
			} `json:"package.license"`
			PackageName struct {
				inner
			} `json:"package.name"`
			PackagePath struct {
				inner
			} `json:"package.path"`
			PackageReference struct {
				inner
			} `json:"package.reference"`
			PackageSize struct {
				inner
			} `json:"package.size"`
			PackageType struct {
				inner
			} `json:"package.type"`
			PackageVersion struct {
				inner
			} `json:"package.version"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"package"`
	Pe struct {
		Description string `json:"description"`
		Fields      struct {
			PeArchitecture struct {
				inner
			} `json:"pe.architecture"`
			PeCompany struct {
				inner
			} `json:"pe.company"`
			PeDescription struct {
				inner
			} `json:"pe.description"`
			PeFileVersion struct {
				inner
			} `json:"pe.file_version"`
			PeGoImportHash struct {
				inner
			} `json:"pe.go_import_hash"`
			PeGoImports struct {
				inner
			} `json:"pe.go_imports"`
			PeGoImportsNamesEntropy struct {
				inner
			} `json:"pe.go_imports_names_entropy"`
			PeGoImportsNamesVarEntropy struct {
				inner
			} `json:"pe.go_imports_names_var_entropy"`
			PeGoStripped struct {
				inner
			} `json:"pe.go_stripped"`
			PeImphash struct {
				inner
			} `json:"pe.imphash"`
			PeImportHash struct {
				inner
			} `json:"pe.import_hash"`
			PeImports struct {
				inner
			} `json:"pe.imports"`
			PeImportsNamesEntropy struct {
				inner
			} `json:"pe.imports_names_entropy"`
			PeImportsNamesVarEntropy struct {
				inner
			} `json:"pe.imports_names_var_entropy"`
			PeOriginalFileName struct {
				inner
			} `json:"pe.original_file_name"`
			PePehash struct {
				inner
			} `json:"pe.pehash"`
			PeProduct struct {
				inner
			} `json:"pe.product"`
			PeSections struct {
				inner
			} `json:"pe.sections"`
			PeSectionsEntropy struct {
				inner
			} `json:"pe.sections.entropy"`
			PeSectionsName struct {
				inner
			} `json:"pe.sections.name"`
			PeSectionsPhysicalSize struct {
				inner
			} `json:"pe.sections.physical_size"`
			PeSectionsVarEntropy struct {
				inner
			} `json:"pe.sections.var_entropy"`
			PeSectionsVirtualSize struct {
				inner
			} `json:"pe.sections.virtual_size"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"pe"`
	Process struct {
		Description string `json:"description"`
		Fields      struct {
			ProcessArgs struct {
				inner
			} `json:"process.args"`
			ProcessArgsCount struct {
				inner
			} `json:"process.args_count"`
			ProcessCodeSignatureDigestAlgorithm struct {
				inner
			} `json:"process.code_signature.digest_algorithm"`
			ProcessCodeSignatureExists struct {
				inner
			} `json:"process.code_signature.exists"`
			ProcessCodeSignatureSigningID struct {
				inner
			} `json:"process.code_signature.signing_id"`
			ProcessCodeSignatureStatus struct {
				inner
			} `json:"process.code_signature.status"`
			ProcessCodeSignatureSubjectName struct {
				inner
			} `json:"process.code_signature.subject_name"`
			ProcessCodeSignatureTeamID struct {
				inner
			} `json:"process.code_signature.team_id"`
			ProcessCodeSignatureTimestamp struct {
				inner
			} `json:"process.code_signature.timestamp"`
			ProcessCodeSignatureTrusted struct {
				inner
			} `json:"process.code_signature.trusted"`
			ProcessCodeSignatureValid struct {
				inner
			} `json:"process.code_signature.valid"`
			ProcessCommandLine struct {
				inner
			} `json:"process.command_line"`
			ProcessElfArchitecture struct {
				inner
			} `json:"process.elf.architecture"`
			ProcessElfByteOrder struct {
				inner
			} `json:"process.elf.byte_order"`
			ProcessElfCPUType struct {
				inner
			} `json:"process.elf.cpu_type"`
			ProcessElfCreationDate struct {
				inner
			} `json:"process.elf.creation_date"`
			ProcessElfExports struct {
				inner
			} `json:"process.elf.exports"`
			ProcessElfGoImportHash struct {
				inner
			} `json:"process.elf.go_import_hash"`
			ProcessElfGoImports struct {
				inner
			} `json:"process.elf.go_imports"`
			ProcessElfGoImportsNamesEntropy struct {
				inner
			} `json:"process.elf.go_imports_names_entropy"`
			ProcessElfGoImportsNamesVarEntropy struct {
				inner
			} `json:"process.elf.go_imports_names_var_entropy"`
			ProcessElfGoStripped struct {
				inner
			} `json:"process.elf.go_stripped"`
			ProcessElfHeaderAbiVersion struct {
				inner
			} `json:"process.elf.header.abi_version"`
			ProcessElfHeaderClass struct {
				inner
			} `json:"process.elf.header.class"`
			ProcessElfHeaderData struct {
				inner
			} `json:"process.elf.header.data"`
			ProcessElfHeaderEntrypoint struct {
				inner
			} `json:"process.elf.header.entrypoint"`
			ProcessElfHeaderObjectVersion struct {
				inner
			} `json:"process.elf.header.object_version"`
			ProcessElfHeaderOsAbi struct {
				inner
			} `json:"process.elf.header.os_abi"`
			ProcessElfHeaderType struct {
				inner
			} `json:"process.elf.header.type"`
			ProcessElfHeaderVersion struct {
				inner
			} `json:"process.elf.header.version"`
			ProcessElfImportHash struct {
				inner
			} `json:"process.elf.import_hash"`
			ProcessElfImports struct {
				inner
			} `json:"process.elf.imports"`
			ProcessElfImportsNamesEntropy struct {
				inner
			} `json:"process.elf.imports_names_entropy"`
			ProcessElfImportsNamesVarEntropy struct {
				inner
			} `json:"process.elf.imports_names_var_entropy"`
			ProcessElfSections struct {
				inner
			} `json:"process.elf.sections"`
			ProcessElfSectionsChi2 struct {
				inner
			} `json:"process.elf.sections.chi2"`
			ProcessElfSectionsEntropy struct {
				inner
			} `json:"process.elf.sections.entropy"`
			ProcessElfSectionsFlags struct {
				inner
			} `json:"process.elf.sections.flags"`
			ProcessElfSectionsName struct {
				inner
			} `json:"process.elf.sections.name"`
			ProcessElfSectionsPhysicalOffset struct {
				inner
			} `json:"process.elf.sections.physical_offset"`
			ProcessElfSectionsPhysicalSize struct {
				inner
			} `json:"process.elf.sections.physical_size"`
			ProcessElfSectionsType struct {
				inner
			} `json:"process.elf.sections.type"`
			ProcessElfSectionsVarEntropy struct {
				inner
			} `json:"process.elf.sections.var_entropy"`
			ProcessElfSectionsVirtualAddress struct {
				inner
			} `json:"process.elf.sections.virtual_address"`
			ProcessElfSectionsVirtualSize struct {
				inner
			} `json:"process.elf.sections.virtual_size"`
			ProcessElfSegments struct {
				inner
			} `json:"process.elf.segments"`
			ProcessElfSegmentsSections struct {
				inner
			} `json:"process.elf.segments.sections"`
			ProcessElfSegmentsType struct {
				inner
			} `json:"process.elf.segments.type"`
			ProcessElfSharedLibraries struct {
				inner
			} `json:"process.elf.shared_libraries"`
			ProcessElfTelfhash struct {
				inner
			} `json:"process.elf.telfhash"`
			ProcessEnd struct {
				inner
			} `json:"process.end"`
			ProcessEntityID struct {
				inner
			} `json:"process.entity_id"`
			ProcessEntryLeaderArgs struct {
				inner
			} `json:"process.entry_leader.args"`
			ProcessEntryLeaderArgsCount struct {
				inner
			} `json:"process.entry_leader.args_count"`
			ProcessEntryLeaderAttestedGroupsName struct {
				inner
			} `json:"process.entry_leader.attested_groups.name"`
			ProcessEntryLeaderAttestedUserID struct {
				inner
			} `json:"process.entry_leader.attested_user.id"`
			ProcessEntryLeaderAttestedUserName struct {
				inner
			} `json:"process.entry_leader.attested_user.name"`
			ProcessEntryLeaderCommandLine struct {
				inner
			} `json:"process.entry_leader.command_line"`
			ProcessEntryLeaderEntityID struct {
				inner
			} `json:"process.entry_leader.entity_id"`
			ProcessEntryLeaderEntryMetaSourceIP struct {
				inner
			} `json:"process.entry_leader.entry_meta.source.ip"`
			ProcessEntryLeaderEntryMetaType struct {
				inner
			} `json:"process.entry_leader.entry_meta.type"`
			ProcessEntryLeaderExecutable struct {
				inner
			} `json:"process.entry_leader.executable"`
			ProcessEntryLeaderGroupID struct {
				inner
			} `json:"process.entry_leader.group.id"`
			ProcessEntryLeaderGroupName struct {
				inner
			} `json:"process.entry_leader.group.name"`
			ProcessEntryLeaderInteractive struct {
				inner
			} `json:"process.entry_leader.interactive"`
			ProcessEntryLeaderName struct {
				inner
			} `json:"process.entry_leader.name"`
			ProcessEntryLeaderParentEntityID struct {
				inner
			} `json:"process.entry_leader.parent.entity_id"`
			ProcessEntryLeaderParentPid struct {
				inner
			} `json:"process.entry_leader.parent.pid"`
			ProcessEntryLeaderParentSessionLeaderEntityID struct {
				inner
			} `json:"process.entry_leader.parent.session_leader.entity_id"`
			ProcessEntryLeaderParentSessionLeaderPid struct {
				inner
			} `json:"process.entry_leader.parent.session_leader.pid"`
			ProcessEntryLeaderParentSessionLeaderStart struct {
				inner
			} `json:"process.entry_leader.parent.session_leader.start"`
			ProcessEntryLeaderParentSessionLeaderVpid struct {
				inner
			} `json:"process.entry_leader.parent.session_leader.vpid"`
			ProcessEntryLeaderParentStart struct {
				inner
			} `json:"process.entry_leader.parent.start"`
			ProcessEntryLeaderParentVpid struct {
				inner
			} `json:"process.entry_leader.parent.vpid"`
			ProcessEntryLeaderPid struct {
				inner
			} `json:"process.entry_leader.pid"`
			ProcessEntryLeaderRealGroupID struct {
				inner
			} `json:"process.entry_leader.real_group.id"`
			ProcessEntryLeaderRealGroupName struct {
				inner
			} `json:"process.entry_leader.real_group.name"`
			ProcessEntryLeaderRealUserID struct {
				inner
			} `json:"process.entry_leader.real_user.id"`
			ProcessEntryLeaderRealUserName struct {
				inner
			} `json:"process.entry_leader.real_user.name"`
			ProcessEntryLeaderSameAsProcess struct {
				inner
			} `json:"process.entry_leader.same_as_process"`
			ProcessEntryLeaderSavedGroupID struct {
				inner
			} `json:"process.entry_leader.saved_group.id"`
			ProcessEntryLeaderSavedGroupName struct {
				inner
			} `json:"process.entry_leader.saved_group.name"`
			ProcessEntryLeaderSavedUserID struct {
				inner
			} `json:"process.entry_leader.saved_user.id"`
			ProcessEntryLeaderSavedUserName struct {
				inner
			} `json:"process.entry_leader.saved_user.name"`
			ProcessEntryLeaderStart struct {
				inner
			} `json:"process.entry_leader.start"`
			ProcessEntryLeaderSupplementalGroupsID struct {
				inner
			} `json:"process.entry_leader.supplemental_groups.id"`
			ProcessEntryLeaderSupplementalGroupsName struct {
				inner
			} `json:"process.entry_leader.supplemental_groups.name"`
			ProcessEntryLeaderTty struct {
				inner
			} `json:"process.entry_leader.tty"`
			ProcessEntryLeaderTtyCharDeviceMajor struct {
				inner
			} `json:"process.entry_leader.tty.char_device.major"`
			ProcessEntryLeaderTtyCharDeviceMinor struct {
				inner
			} `json:"process.entry_leader.tty.char_device.minor"`
			ProcessEntryLeaderUserID struct {
				inner
			} `json:"process.entry_leader.user.id"`
			ProcessEntryLeaderUserName struct {
				inner
			} `json:"process.entry_leader.user.name"`
			ProcessEntryLeaderVpid struct {
				inner
			} `json:"process.entry_leader.vpid"`
			ProcessEntryLeaderWorkingDirectory struct {
				inner
			} `json:"process.entry_leader.working_directory"`
			ProcessEnvVars struct {
				inner
			} `json:"process.env_vars"`
			ProcessExecutable struct {
				inner
			} `json:"process.executable"`
			ProcessExitCode struct {
				inner
			} `json:"process.exit_code"`
			ProcessGroupLeaderArgs struct {
				inner
			} `json:"process.group_leader.args"`
			ProcessGroupLeaderArgsCount struct {
				inner
			} `json:"process.group_leader.args_count"`
			ProcessGroupLeaderCommandLine struct {
				inner
			} `json:"process.group_leader.command_line"`
			ProcessGroupLeaderEntityID struct {
				inner
			} `json:"process.group_leader.entity_id"`
			ProcessGroupLeaderExecutable struct {
				inner
			} `json:"process.group_leader.executable"`
			ProcessGroupLeaderGroupID struct {
				inner
			} `json:"process.group_leader.group.id"`
			ProcessGroupLeaderGroupName struct {
				inner
			} `json:"process.group_leader.group.name"`
			ProcessGroupLeaderInteractive struct {
				inner
			} `json:"process.group_leader.interactive"`
			ProcessGroupLeaderName struct {
				inner
			} `json:"process.group_leader.name"`
			ProcessGroupLeaderPid struct {
				inner
			} `json:"process.group_leader.pid"`
			ProcessGroupLeaderRealGroupID struct {
				inner
			} `json:"process.group_leader.real_group.id"`
			ProcessGroupLeaderRealGroupName struct {
				inner
			} `json:"process.group_leader.real_group.name"`
			ProcessGroupLeaderRealUserID struct {
				inner
			} `json:"process.group_leader.real_user.id"`
			ProcessGroupLeaderRealUserName struct {
				inner
			} `json:"process.group_leader.real_user.name"`
			ProcessGroupLeaderSameAsProcess struct {
				inner
			} `json:"process.group_leader.same_as_process"`
			ProcessGroupLeaderSavedGroupID struct {
				inner
			} `json:"process.group_leader.saved_group.id"`
			ProcessGroupLeaderSavedGroupName struct {
				inner
			} `json:"process.group_leader.saved_group.name"`
			ProcessGroupLeaderSavedUserID struct {
				inner
			} `json:"process.group_leader.saved_user.id"`
			ProcessGroupLeaderSavedUserName struct {
				inner
			} `json:"process.group_leader.saved_user.name"`
			ProcessGroupLeaderStart struct {
				inner
			} `json:"process.group_leader.start"`
			ProcessGroupLeaderSupplementalGroupsID struct {
				inner
			} `json:"process.group_leader.supplemental_groups.id"`
			ProcessGroupLeaderSupplementalGroupsName struct {
				inner
			} `json:"process.group_leader.supplemental_groups.name"`
			ProcessGroupLeaderTty struct {
				inner
			} `json:"process.group_leader.tty"`
			ProcessGroupLeaderTtyCharDeviceMajor struct {
				inner
			} `json:"process.group_leader.tty.char_device.major"`
			ProcessGroupLeaderTtyCharDeviceMinor struct {
				inner
			} `json:"process.group_leader.tty.char_device.minor"`
			ProcessGroupLeaderUserID struct {
				inner
			} `json:"process.group_leader.user.id"`
			ProcessGroupLeaderUserName struct {
				inner
			} `json:"process.group_leader.user.name"`
			ProcessGroupLeaderVpid struct {
				inner
			} `json:"process.group_leader.vpid"`
			ProcessGroupLeaderWorkingDirectory struct {
				inner
			} `json:"process.group_leader.working_directory"`
			ProcessHashMd5 struct {
				inner
			} `json:"process.hash.md5"`
			ProcessHashSha1 struct {
				inner
			} `json:"process.hash.sha1"`
			ProcessHashSha256 struct {
				inner
			} `json:"process.hash.sha256"`
			ProcessHashSha384 struct {
				inner
			} `json:"process.hash.sha384"`
			ProcessHashSha512 struct {
				inner
			} `json:"process.hash.sha512"`
			ProcessHashSsdeep struct {
				inner
			} `json:"process.hash.ssdeep"`
			ProcessHashTlsh struct {
				inner
			} `json:"process.hash.tlsh"`
			ProcessInteractive struct {
				inner
			} `json:"process.interactive"`
			ProcessIo struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io"`
			ProcessIoBytesSkipped struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.bytes_skipped"`
			ProcessIoBytesSkippedLength struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.bytes_skipped.length"`
			ProcessIoBytesSkippedOffset struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.bytes_skipped.offset"`
			ProcessIoMaxBytesPerProcessExceeded struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.max_bytes_per_process_exceeded"`
			ProcessIoText struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.text"`
			ProcessIoTotalBytesCaptured struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.total_bytes_captured"`
			ProcessIoTotalBytesSkipped struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.total_bytes_skipped"`
			ProcessIoType struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.io.type"`
			ProcessMachoGoImportHash struct {
				inner
			} `json:"process.macho.go_import_hash"`
			ProcessMachoGoImports struct {
				inner
			} `json:"process.macho.go_imports"`
			ProcessMachoGoImportsNamesEntropy struct {
				inner
			} `json:"process.macho.go_imports_names_entropy"`
			ProcessMachoGoImportsNamesVarEntropy struct {
				inner
			} `json:"process.macho.go_imports_names_var_entropy"`
			ProcessMachoGoStripped struct {
				inner
			} `json:"process.macho.go_stripped"`
			ProcessMachoImportHash struct {
				inner
			} `json:"process.macho.import_hash"`
			ProcessMachoImports struct {
				inner
			} `json:"process.macho.imports"`
			ProcessMachoImportsNamesEntropy struct {
				inner
			} `json:"process.macho.imports_names_entropy"`
			ProcessMachoImportsNamesVarEntropy struct {
				inner
			} `json:"process.macho.imports_names_var_entropy"`
			ProcessMachoSections struct {
				inner
			} `json:"process.macho.sections"`
			ProcessMachoSectionsEntropy struct {
				inner
			} `json:"process.macho.sections.entropy"`
			ProcessMachoSectionsName struct {
				inner
			} `json:"process.macho.sections.name"`
			ProcessMachoSectionsPhysicalSize struct {
				inner
			} `json:"process.macho.sections.physical_size"`
			ProcessMachoSectionsVarEntropy struct {
				inner
			} `json:"process.macho.sections.var_entropy"`
			ProcessMachoSectionsVirtualSize struct {
				inner
			} `json:"process.macho.sections.virtual_size"`
			ProcessMachoSymhash struct {
				inner
			} `json:"process.macho.symhash"`
			ProcessName struct {
				inner
			} `json:"process.name"`
			ProcessParentArgs struct {
				inner
			} `json:"process.parent.args"`
			ProcessParentArgsCount struct {
				inner
			} `json:"process.parent.args_count"`
			ProcessParentCodeSignatureDigestAlgorithm struct {
				inner
			} `json:"process.parent.code_signature.digest_algorithm"`
			ProcessParentCodeSignatureExists struct {
				inner
			} `json:"process.parent.code_signature.exists"`
			ProcessParentCodeSignatureSigningID struct {
				inner
			} `json:"process.parent.code_signature.signing_id"`
			ProcessParentCodeSignatureStatus struct {
				inner
			} `json:"process.parent.code_signature.status"`
			ProcessParentCodeSignatureSubjectName struct {
				inner
			} `json:"process.parent.code_signature.subject_name"`
			ProcessParentCodeSignatureTeamID struct {
				inner
			} `json:"process.parent.code_signature.team_id"`
			ProcessParentCodeSignatureTimestamp struct {
				inner
			} `json:"process.parent.code_signature.timestamp"`
			ProcessParentCodeSignatureTrusted struct {
				inner
			} `json:"process.parent.code_signature.trusted"`
			ProcessParentCodeSignatureValid struct {
				inner
			} `json:"process.parent.code_signature.valid"`
			ProcessParentCommandLine struct {
				inner
			} `json:"process.parent.command_line"`
			ProcessParentElfArchitecture struct {
				inner
			} `json:"process.parent.elf.architecture"`
			ProcessParentElfByteOrder struct {
				inner
			} `json:"process.parent.elf.byte_order"`
			ProcessParentElfCPUType struct {
				inner
			} `json:"process.parent.elf.cpu_type"`
			ProcessParentElfCreationDate struct {
				inner
			} `json:"process.parent.elf.creation_date"`
			ProcessParentElfExports struct {
				inner
			} `json:"process.parent.elf.exports"`
			ProcessParentElfGoImportHash struct {
				inner
			} `json:"process.parent.elf.go_import_hash"`
			ProcessParentElfGoImports struct {
				inner
			} `json:"process.parent.elf.go_imports"`
			ProcessParentElfGoImportsNamesEntropy struct {
				inner
			} `json:"process.parent.elf.go_imports_names_entropy"`
			ProcessParentElfGoImportsNamesVarEntropy struct {
				inner
			} `json:"process.parent.elf.go_imports_names_var_entropy"`
			ProcessParentElfGoStripped struct {
				inner
			} `json:"process.parent.elf.go_stripped"`
			ProcessParentElfHeaderAbiVersion struct {
				inner
			} `json:"process.parent.elf.header.abi_version"`
			ProcessParentElfHeaderClass struct {
				inner
			} `json:"process.parent.elf.header.class"`
			ProcessParentElfHeaderData struct {
				inner
			} `json:"process.parent.elf.header.data"`
			ProcessParentElfHeaderEntrypoint struct {
				inner
			} `json:"process.parent.elf.header.entrypoint"`
			ProcessParentElfHeaderObjectVersion struct {
				inner
			} `json:"process.parent.elf.header.object_version"`
			ProcessParentElfHeaderOsAbi struct {
				inner
			} `json:"process.parent.elf.header.os_abi"`
			ProcessParentElfHeaderType struct {
				inner
			} `json:"process.parent.elf.header.type"`
			ProcessParentElfHeaderVersion struct {
				inner
			} `json:"process.parent.elf.header.version"`
			ProcessParentElfImportHash struct {
				inner
			} `json:"process.parent.elf.import_hash"`
			ProcessParentElfImports struct {
				inner
			} `json:"process.parent.elf.imports"`
			ProcessParentElfImportsNamesEntropy struct {
				inner
			} `json:"process.parent.elf.imports_names_entropy"`
			ProcessParentElfImportsNamesVarEntropy struct {
				inner
			} `json:"process.parent.elf.imports_names_var_entropy"`
			ProcessParentElfSections struct {
				inner
			} `json:"process.parent.elf.sections"`
			ProcessParentElfSectionsChi2 struct {
				inner
			} `json:"process.parent.elf.sections.chi2"`
			ProcessParentElfSectionsEntropy struct {
				inner
			} `json:"process.parent.elf.sections.entropy"`
			ProcessParentElfSectionsFlags struct {
				inner
			} `json:"process.parent.elf.sections.flags"`
			ProcessParentElfSectionsName struct {
				inner
			} `json:"process.parent.elf.sections.name"`
			ProcessParentElfSectionsPhysicalOffset struct {
				inner
			} `json:"process.parent.elf.sections.physical_offset"`
			ProcessParentElfSectionsPhysicalSize struct {
				inner
			} `json:"process.parent.elf.sections.physical_size"`
			ProcessParentElfSectionsType struct {
				inner
			} `json:"process.parent.elf.sections.type"`
			ProcessParentElfSectionsVarEntropy struct {
				inner
			} `json:"process.parent.elf.sections.var_entropy"`
			ProcessParentElfSectionsVirtualAddress struct {
				inner
			} `json:"process.parent.elf.sections.virtual_address"`
			ProcessParentElfSectionsVirtualSize struct {
				inner
			} `json:"process.parent.elf.sections.virtual_size"`
			ProcessParentElfSegments struct {
				inner
			} `json:"process.parent.elf.segments"`
			ProcessParentElfSegmentsSections struct {
				inner
			} `json:"process.parent.elf.segments.sections"`
			ProcessParentElfSegmentsType struct {
				inner
			} `json:"process.parent.elf.segments.type"`
			ProcessParentElfSharedLibraries struct {
				inner
			} `json:"process.parent.elf.shared_libraries"`
			ProcessParentElfTelfhash struct {
				inner
			} `json:"process.parent.elf.telfhash"`
			ProcessParentEnd struct {
				inner
			} `json:"process.parent.end"`
			ProcessParentEntityID struct {
				inner
			} `json:"process.parent.entity_id"`
			ProcessParentExecutable struct {
				inner
			} `json:"process.parent.executable"`
			ProcessParentExitCode struct {
				inner
			} `json:"process.parent.exit_code"`
			ProcessParentGroupID struct {
				inner
			} `json:"process.parent.group.id"`
			ProcessParentGroupName struct {
				inner
			} `json:"process.parent.group.name"`
			ProcessParentGroupLeaderEntityID struct {
				inner
			} `json:"process.parent.group_leader.entity_id"`
			ProcessParentGroupLeaderPid struct {
				inner
			} `json:"process.parent.group_leader.pid"`
			ProcessParentGroupLeaderStart struct {
				inner
			} `json:"process.parent.group_leader.start"`
			ProcessParentGroupLeaderVpid struct {
				inner
			} `json:"process.parent.group_leader.vpid"`
			ProcessParentHashMd5 struct {
				inner
			} `json:"process.parent.hash.md5"`
			ProcessParentHashSha1 struct {
				inner
			} `json:"process.parent.hash.sha1"`
			ProcessParentHashSha256 struct {
				inner
			} `json:"process.parent.hash.sha256"`
			ProcessParentHashSha384 struct {
				inner
			} `json:"process.parent.hash.sha384"`
			ProcessParentHashSha512 struct {
				inner
			} `json:"process.parent.hash.sha512"`
			ProcessParentHashSsdeep struct {
				inner
			} `json:"process.parent.hash.ssdeep"`
			ProcessParentHashTlsh struct {
				inner
			} `json:"process.parent.hash.tlsh"`
			ProcessParentInteractive struct {
				inner
			} `json:"process.parent.interactive"`
			ProcessParentMachoGoImportHash struct {
				inner
			} `json:"process.parent.macho.go_import_hash"`
			ProcessParentMachoGoImports struct {
				inner
			} `json:"process.parent.macho.go_imports"`
			ProcessParentMachoGoImportsNamesEntropy struct {
				inner
			} `json:"process.parent.macho.go_imports_names_entropy"`
			ProcessParentMachoGoImportsNamesVarEntropy struct {
				inner
			} `json:"process.parent.macho.go_imports_names_var_entropy"`
			ProcessParentMachoGoStripped struct {
				inner
			} `json:"process.parent.macho.go_stripped"`
			ProcessParentMachoImportHash struct {
				inner
			} `json:"process.parent.macho.import_hash"`
			ProcessParentMachoImports struct {
				inner
			} `json:"process.parent.macho.imports"`
			ProcessParentMachoImportsNamesEntropy struct {
				inner
			} `json:"process.parent.macho.imports_names_entropy"`
			ProcessParentMachoImportsNamesVarEntropy struct {
				inner
			} `json:"process.parent.macho.imports_names_var_entropy"`
			ProcessParentMachoSections struct {
				inner
			} `json:"process.parent.macho.sections"`
			ProcessParentMachoSectionsEntropy struct {
				inner
			} `json:"process.parent.macho.sections.entropy"`
			ProcessParentMachoSectionsName struct {
				inner
			} `json:"process.parent.macho.sections.name"`
			ProcessParentMachoSectionsPhysicalSize struct {
				inner
			} `json:"process.parent.macho.sections.physical_size"`
			ProcessParentMachoSectionsVarEntropy struct {
				inner
			} `json:"process.parent.macho.sections.var_entropy"`
			ProcessParentMachoSectionsVirtualSize struct {
				inner
			} `json:"process.parent.macho.sections.virtual_size"`
			ProcessParentMachoSymhash struct {
				inner
			} `json:"process.parent.macho.symhash"`
			ProcessParentName struct {
				inner
			} `json:"process.parent.name"`
			ProcessParentPeArchitecture struct {
				inner
			} `json:"process.parent.pe.architecture"`
			ProcessParentPeCompany struct {
				inner
			} `json:"process.parent.pe.company"`
			ProcessParentPeDescription struct {
				inner
			} `json:"process.parent.pe.description"`
			ProcessParentPeFileVersion struct {
				inner
			} `json:"process.parent.pe.file_version"`
			ProcessParentPeGoImportHash struct {
				inner
			} `json:"process.parent.pe.go_import_hash"`
			ProcessParentPeGoImports struct {
				inner
			} `json:"process.parent.pe.go_imports"`
			ProcessParentPeGoImportsNamesEntropy struct {
				inner
			} `json:"process.parent.pe.go_imports_names_entropy"`
			ProcessParentPeGoImportsNamesVarEntropy struct {
				inner
			} `json:"process.parent.pe.go_imports_names_var_entropy"`
			ProcessParentPeGoStripped struct {
				inner
			} `json:"process.parent.pe.go_stripped"`
			ProcessParentPeImphash struct {
				inner
			} `json:"process.parent.pe.imphash"`
			ProcessParentPeImportHash struct {
				inner
			} `json:"process.parent.pe.import_hash"`
			ProcessParentPeImports struct {
				inner
			} `json:"process.parent.pe.imports"`
			ProcessParentPeImportsNamesEntropy struct {
				inner
			} `json:"process.parent.pe.imports_names_entropy"`
			ProcessParentPeImportsNamesVarEntropy struct {
				inner
			} `json:"process.parent.pe.imports_names_var_entropy"`
			ProcessParentPeOriginalFileName struct {
				inner
			} `json:"process.parent.pe.original_file_name"`
			ProcessParentPePehash struct {
				inner
			} `json:"process.parent.pe.pehash"`
			ProcessParentPeProduct struct {
				inner
			} `json:"process.parent.pe.product"`
			ProcessParentPeSections struct {
				inner
			} `json:"process.parent.pe.sections"`
			ProcessParentPeSectionsEntropy struct {
				inner
			} `json:"process.parent.pe.sections.entropy"`
			ProcessParentPeSectionsName struct {
				inner
			} `json:"process.parent.pe.sections.name"`
			ProcessParentPeSectionsPhysicalSize struct {
				inner
			} `json:"process.parent.pe.sections.physical_size"`
			ProcessParentPeSectionsVarEntropy struct {
				inner
			} `json:"process.parent.pe.sections.var_entropy"`
			ProcessParentPeSectionsVirtualSize struct {
				inner
			} `json:"process.parent.pe.sections.virtual_size"`
			ProcessParentPgid struct {
				inner
			} `json:"process.parent.pgid"`
			ProcessParentPid struct {
				inner
			} `json:"process.parent.pid"`
			ProcessParentRealGroupID struct {
				inner
			} `json:"process.parent.real_group.id"`
			ProcessParentRealGroupName struct {
				inner
			} `json:"process.parent.real_group.name"`
			ProcessParentRealUserID struct {
				inner
			} `json:"process.parent.real_user.id"`
			ProcessParentRealUserName struct {
				inner
			} `json:"process.parent.real_user.name"`
			ProcessParentSavedGroupID struct {
				inner
			} `json:"process.parent.saved_group.id"`
			ProcessParentSavedGroupName struct {
				inner
			} `json:"process.parent.saved_group.name"`
			ProcessParentSavedUserID struct {
				inner
			} `json:"process.parent.saved_user.id"`
			ProcessParentSavedUserName struct {
				inner
			} `json:"process.parent.saved_user.name"`
			ProcessParentStart struct {
				inner
			} `json:"process.parent.start"`
			ProcessParentSupplementalGroupsID struct {
				inner
			} `json:"process.parent.supplemental_groups.id"`
			ProcessParentSupplementalGroupsName struct {
				inner
			} `json:"process.parent.supplemental_groups.name"`
			ProcessParentThreadCapabilitiesEffective struct {
				inner
			} `json:"process.parent.thread.capabilities.effective"`
			ProcessParentThreadCapabilitiesPermitted struct {
				inner
			} `json:"process.parent.thread.capabilities.permitted"`
			ProcessParentThreadID struct {
				inner
			} `json:"process.parent.thread.id"`
			ProcessParentThreadName struct {
				inner
			} `json:"process.parent.thread.name"`
			ProcessParentTitle struct {
				inner
			} `json:"process.parent.title"`
			ProcessParentTty struct {
				inner
			} `json:"process.parent.tty"`
			ProcessParentTtyCharDeviceMajor struct {
				inner
			} `json:"process.parent.tty.char_device.major"`
			ProcessParentTtyCharDeviceMinor struct {
				inner
			} `json:"process.parent.tty.char_device.minor"`
			ProcessParentUptime struct {
				inner
			} `json:"process.parent.uptime"`
			ProcessParentUserID struct {
				inner
			} `json:"process.parent.user.id"`
			ProcessParentUserName struct {
				inner
			} `json:"process.parent.user.name"`
			ProcessParentVpid struct {
				inner
			} `json:"process.parent.vpid"`
			ProcessParentWorkingDirectory struct {
				inner
			} `json:"process.parent.working_directory"`
			ProcessPeArchitecture struct {
				inner
			} `json:"process.pe.architecture"`
			ProcessPeCompany struct {
				inner
			} `json:"process.pe.company"`
			ProcessPeDescription struct {
				inner
			} `json:"process.pe.description"`
			ProcessPeFileVersion struct {
				inner
			} `json:"process.pe.file_version"`
			ProcessPeGoImportHash struct {
				inner
			} `json:"process.pe.go_import_hash"`
			ProcessPeGoImports struct {
				inner
			} `json:"process.pe.go_imports"`
			ProcessPeGoImportsNamesEntropy struct {
				inner
			} `json:"process.pe.go_imports_names_entropy"`
			ProcessPeGoImportsNamesVarEntropy struct {
				inner
			} `json:"process.pe.go_imports_names_var_entropy"`
			ProcessPeGoStripped struct {
				inner
			} `json:"process.pe.go_stripped"`
			ProcessPeImphash struct {
				inner
			} `json:"process.pe.imphash"`
			ProcessPeImportHash struct {
				inner
			} `json:"process.pe.import_hash"`
			ProcessPeImports struct {
				inner
			} `json:"process.pe.imports"`
			ProcessPeImportsNamesEntropy struct {
				inner
			} `json:"process.pe.imports_names_entropy"`
			ProcessPeImportsNamesVarEntropy struct {
				inner
			} `json:"process.pe.imports_names_var_entropy"`
			ProcessPeOriginalFileName struct {
				inner
			} `json:"process.pe.original_file_name"`
			ProcessPePehash struct {
				inner
			} `json:"process.pe.pehash"`
			ProcessPeProduct struct {
				inner
			} `json:"process.pe.product"`
			ProcessPeSections struct {
				inner
			} `json:"process.pe.sections"`
			ProcessPeSectionsEntropy struct {
				inner
			} `json:"process.pe.sections.entropy"`
			ProcessPeSectionsName struct {
				inner
			} `json:"process.pe.sections.name"`
			ProcessPeSectionsPhysicalSize struct {
				inner
			} `json:"process.pe.sections.physical_size"`
			ProcessPeSectionsVarEntropy struct {
				inner
			} `json:"process.pe.sections.var_entropy"`
			ProcessPeSectionsVirtualSize struct {
				inner
			} `json:"process.pe.sections.virtual_size"`
			ProcessPgid struct {
				inner
			} `json:"process.pgid"`
			ProcessPid struct {
				inner
			} `json:"process.pid"`
			ProcessPreviousArgs struct {
				inner
			} `json:"process.previous.args"`
			ProcessPreviousArgsCount struct {
				inner
			} `json:"process.previous.args_count"`
			ProcessPreviousExecutable struct {
				inner
			} `json:"process.previous.executable"`
			ProcessRealGroupID struct {
				inner
			} `json:"process.real_group.id"`
			ProcessRealGroupName struct {
				inner
			} `json:"process.real_group.name"`
			ProcessRealUserID struct {
				inner
			} `json:"process.real_user.id"`
			ProcessRealUserName struct {
				inner
			} `json:"process.real_user.name"`
			ProcessSavedGroupID struct {
				inner
			} `json:"process.saved_group.id"`
			ProcessSavedGroupName struct {
				inner
			} `json:"process.saved_group.name"`
			ProcessSavedUserID struct {
				inner
			} `json:"process.saved_user.id"`
			ProcessSavedUserName struct {
				inner
			} `json:"process.saved_user.name"`
			ProcessSessionLeaderArgs struct {
				inner
			} `json:"process.session_leader.args"`
			ProcessSessionLeaderArgsCount struct {
				inner
			} `json:"process.session_leader.args_count"`
			ProcessSessionLeaderCommandLine struct {
				inner
			} `json:"process.session_leader.command_line"`
			ProcessSessionLeaderEntityID struct {
				inner
			} `json:"process.session_leader.entity_id"`
			ProcessSessionLeaderExecutable struct {
				inner
			} `json:"process.session_leader.executable"`
			ProcessSessionLeaderGroupID struct {
				inner
			} `json:"process.session_leader.group.id"`
			ProcessSessionLeaderGroupName struct {
				inner
			} `json:"process.session_leader.group.name"`
			ProcessSessionLeaderInteractive struct {
				inner
			} `json:"process.session_leader.interactive"`
			ProcessSessionLeaderName struct {
				inner
			} `json:"process.session_leader.name"`
			ProcessSessionLeaderParentEntityID struct {
				inner
			} `json:"process.session_leader.parent.entity_id"`
			ProcessSessionLeaderParentPid struct {
				inner
			} `json:"process.session_leader.parent.pid"`
			ProcessSessionLeaderParentSessionLeaderEntityID struct {
				inner
			} `json:"process.session_leader.parent.session_leader.entity_id"`
			ProcessSessionLeaderParentSessionLeaderPid struct {
				inner
			} `json:"process.session_leader.parent.session_leader.pid"`
			ProcessSessionLeaderParentSessionLeaderStart struct {
				inner
			} `json:"process.session_leader.parent.session_leader.start"`
			ProcessSessionLeaderParentSessionLeaderVpid struct {
				inner
			} `json:"process.session_leader.parent.session_leader.vpid"`
			ProcessSessionLeaderParentStart struct {
				inner
			} `json:"process.session_leader.parent.start"`
			ProcessSessionLeaderParentVpid struct {
				inner
			} `json:"process.session_leader.parent.vpid"`
			ProcessSessionLeaderPid struct {
				inner
			} `json:"process.session_leader.pid"`
			ProcessSessionLeaderRealGroupID struct {
				inner
			} `json:"process.session_leader.real_group.id"`
			ProcessSessionLeaderRealGroupName struct {
				inner
			} `json:"process.session_leader.real_group.name"`
			ProcessSessionLeaderRealUserID struct {
				inner
			} `json:"process.session_leader.real_user.id"`
			ProcessSessionLeaderRealUserName struct {
				inner
			} `json:"process.session_leader.real_user.name"`
			ProcessSessionLeaderSameAsProcess struct {
				inner
			} `json:"process.session_leader.same_as_process"`
			ProcessSessionLeaderSavedGroupID struct {
				inner
			} `json:"process.session_leader.saved_group.id"`
			ProcessSessionLeaderSavedGroupName struct {
				inner
			} `json:"process.session_leader.saved_group.name"`
			ProcessSessionLeaderSavedUserID struct {
				inner
			} `json:"process.session_leader.saved_user.id"`
			ProcessSessionLeaderSavedUserName struct {
				inner
			} `json:"process.session_leader.saved_user.name"`
			ProcessSessionLeaderStart struct {
				inner
			} `json:"process.session_leader.start"`
			ProcessSessionLeaderSupplementalGroupsID struct {
				inner
			} `json:"process.session_leader.supplemental_groups.id"`
			ProcessSessionLeaderSupplementalGroupsName struct {
				inner
			} `json:"process.session_leader.supplemental_groups.name"`
			ProcessSessionLeaderTty struct {
				inner
			} `json:"process.session_leader.tty"`
			ProcessSessionLeaderTtyCharDeviceMajor struct {
				inner
			} `json:"process.session_leader.tty.char_device.major"`
			ProcessSessionLeaderTtyCharDeviceMinor struct {
				inner
			} `json:"process.session_leader.tty.char_device.minor"`
			ProcessSessionLeaderUserID struct {
				inner
			} `json:"process.session_leader.user.id"`
			ProcessSessionLeaderUserName struct {
				inner
			} `json:"process.session_leader.user.name"`
			ProcessSessionLeaderVpid struct {
				inner
			} `json:"process.session_leader.vpid"`
			ProcessSessionLeaderWorkingDirectory struct {
				inner
			} `json:"process.session_leader.working_directory"`
			ProcessStart struct {
				inner
			} `json:"process.start"`
			ProcessSupplementalGroupsID struct {
				inner
			} `json:"process.supplemental_groups.id"`
			ProcessSupplementalGroupsName struct {
				inner
			} `json:"process.supplemental_groups.name"`
			ProcessThreadCapabilitiesEffective struct {
				inner
			} `json:"process.thread.capabilities.effective"`
			ProcessThreadCapabilitiesPermitted struct {
				inner
			} `json:"process.thread.capabilities.permitted"`
			ProcessThreadID struct {
				inner
			} `json:"process.thread.id"`
			ProcessThreadName struct {
				inner
			} `json:"process.thread.name"`
			ProcessTitle struct {
				inner
			} `json:"process.title"`
			ProcessTty struct {
				inner
			} `json:"process.tty"`
			ProcessTtyCharDeviceMajor struct {
				inner
			} `json:"process.tty.char_device.major"`
			ProcessTtyCharDeviceMinor struct {
				inner
			} `json:"process.tty.char_device.minor"`
			ProcessTtyColumns struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.tty.columns"`
			ProcessTtyRows struct {
				Beta string `json:"beta"`
				inner
			} `json:"process.tty.rows"`
			ProcessUptime struct {
				inner
			} `json:"process.uptime"`
			ProcessUserID struct {
				inner
			} `json:"process.user.id"`
			ProcessUserName struct {
				inner
			} `json:"process.user.name"`
			ProcessVpid struct {
				inner
			} `json:"process.vpid"`
			ProcessWorkingDirectory struct {
				inner
			} `json:"process.working_directory"`
		} `json:"fields"`
		Group    int      `json:"group"`
		Name     string   `json:"name"`
		Nestings []string `json:"nestings"`
		Prefix   string   `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As            string   `json:"as"`
				At            string   `json:"at"`
				Full          string   `json:"full"`
				ShortOverride string   `json:"short_override"`
				Normalize     []string `json:"normalize,omitempty"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		ReusedHere []struct {
			Full       string   `json:"full"`
			SchemaName string   `json:"schema_name"`
			Short      string   `json:"short"`
			Normalize  []string `json:"normalize,omitempty"`
			Beta       string   `json:"beta,omitempty"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"process"`
	Registry struct {
		Description string `json:"description"`
		Fields      struct {
			RegistryDataBytes struct {
				inner
			} `json:"registry.data.bytes"`
			RegistryDataStrings struct {
				inner
			} `json:"registry.data.strings"`
			RegistryDataType struct {
				inner
			} `json:"registry.data.type"`
			RegistryHive struct {
				inner
			} `json:"registry.hive"`
			RegistryKey struct {
				inner
			} `json:"registry.key"`
			RegistryPath struct {
				inner
			} `json:"registry.path"`
			RegistryValue struct {
				inner
			} `json:"registry.value"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"registry"`
	Related struct {
		Description string `json:"description"`
		Fields      struct {
			RelatedHash struct {
				inner
			} `json:"related.hash"`
			RelatedHosts struct {
				inner
			} `json:"related.hosts"`
			RelatedIP struct {
				inner
			} `json:"related.ip"`
			RelatedUser struct {
				inner
			} `json:"related.user"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"related"`
	Risk struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			RiskCalculatedLevel struct {
				inner
			} `json:"risk.calculated_level"`
			RiskCalculatedScore struct {
				inner
			} `json:"risk.calculated_score"`
			RiskCalculatedScoreNorm struct {
				inner
			} `json:"risk.calculated_score_norm"`
			RiskStaticLevel struct {
				inner
			} `json:"risk.static_level"`
			RiskStaticScore struct {
				inner
			} `json:"risk.static_score"`
			RiskStaticScoreNorm struct {
				inner
			} `json:"risk.static_score_norm"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"risk"`
	Rule struct {
		Description string `json:"description"`
		Fields      struct {
			RuleAuthor struct {
				inner
			} `json:"rule.author"`
			RuleCategory struct {
				inner
			} `json:"rule.category"`
			RuleDescription struct {
				inner
			} `json:"rule.description"`
			RuleID struct {
				inner
			} `json:"rule.id"`
			RuleLicense struct {
				inner
			} `json:"rule.license"`
			RuleName struct {
				inner
			} `json:"rule.name"`
			RuleReference struct {
				inner
			} `json:"rule.reference"`
			RuleRuleset struct {
				inner
			} `json:"rule.ruleset"`
			RuleUUID struct {
				inner
			} `json:"rule.uuid"`
			RuleVersion struct {
				inner
			} `json:"rule.version"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"rule"`
	Server struct {
		Description string `json:"description"`
		Fields      struct {
			ServerAddress struct {
				inner
			} `json:"server.address"`
			ServerAsNumber struct {
				inner
			} `json:"server.as.number"`
			ServerAsOrganizationName struct {
				inner
			} `json:"server.as.organization.name"`
			ServerBytes struct {
				inner
			} `json:"server.bytes"`
			ServerDomain struct {
				inner
			} `json:"server.domain"`
			ServerGeoCityName struct {
				inner
			} `json:"server.geo.city_name"`
			ServerGeoContinentCode struct {
				inner
			} `json:"server.geo.continent_code"`
			ServerGeoContinentName struct {
				inner
			} `json:"server.geo.continent_name"`
			ServerGeoCountryIsoCode struct {
				inner
			} `json:"server.geo.country_iso_code"`
			ServerGeoCountryName struct {
				inner
			} `json:"server.geo.country_name"`
			ServerGeoLocation struct {
				inner
			} `json:"server.geo.location"`
			ServerGeoName struct {
				inner
			} `json:"server.geo.name"`
			ServerGeoPostalCode struct {
				inner
			} `json:"server.geo.postal_code"`
			ServerGeoRegionIsoCode struct {
				inner
			} `json:"server.geo.region_iso_code"`
			ServerGeoRegionName struct {
				inner
			} `json:"server.geo.region_name"`
			ServerGeoTimezone struct {
				inner
			} `json:"server.geo.timezone"`
			ServerIP struct {
				inner
			} `json:"server.ip"`
			ServerMac struct {
				inner
			} `json:"server.mac"`
			ServerNatIP struct {
				inner
			} `json:"server.nat.ip"`
			ServerNatPort struct {
				inner
			} `json:"server.nat.port"`
			ServerPackets struct {
				inner
			} `json:"server.packets"`
			ServerPort struct {
				inner
			} `json:"server.port"`
			ServerRegisteredDomain struct {
				inner
			} `json:"server.registered_domain"`
			ServerSubdomain struct {
				inner
			} `json:"server.subdomain"`
			ServerTopLevelDomain struct {
				inner
			} `json:"server.top_level_domain"`
			ServerUserDomain struct {
				inner
			} `json:"server.user.domain"`
			ServerUserEmail struct {
				inner
			} `json:"server.user.email"`
			ServerUserFullName struct {
				inner
			} `json:"server.user.full_name"`
			ServerUserGroupDomain struct {
				inner
			} `json:"server.user.group.domain"`
			ServerUserGroupID struct {
				inner
			} `json:"server.user.group.id"`
			ServerUserGroupName struct {
				inner
			} `json:"server.user.group.name"`
			ServerUserHash struct {
				inner
			} `json:"server.user.hash"`
			ServerUserID struct {
				inner
			} `json:"server.user.id"`
			ServerUserName struct {
				inner
			} `json:"server.user.name"`
			ServerUserRoles struct {
				inner
			} `json:"server.user.roles"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"server"`
	Service struct {
		Description string `json:"description"`
		Fields      struct {
			ServiceAddress struct {
				inner
			} `json:"service.address"`
			ServiceEnvironment struct {
				Beta string `json:"beta"`
				inner
			} `json:"service.environment"`
			ServiceEphemeralID struct {
				inner
			} `json:"service.ephemeral_id"`
			ServiceID struct {
				inner
			} `json:"service.id"`
			ServiceName struct {
				inner
			} `json:"service.name"`
			ServiceNodeName struct {
				inner
			} `json:"service.node.name"`
			ServiceNodeRole struct {
				inner
			} `json:"service.node.role"`
			ServiceNodeRoles struct {
				inner
			} `json:"service.node.roles"`
			ServiceOriginAddress struct {
				inner
			} `json:"service.origin.address"`
			ServiceOriginEnvironment struct {
				Beta string `json:"beta"`
				inner
			} `json:"service.origin.environment"`
			ServiceOriginEphemeralID struct {
				inner
			} `json:"service.origin.ephemeral_id"`
			ServiceOriginID struct {
				inner
			} `json:"service.origin.id"`
			ServiceOriginName struct {
				inner
			} `json:"service.origin.name"`
			ServiceOriginNodeName struct {
				inner
			} `json:"service.origin.node.name"`
			ServiceOriginNodeRole struct {
				inner
			} `json:"service.origin.node.role"`
			ServiceOriginNodeRoles struct {
				inner
			} `json:"service.origin.node.roles"`
			ServiceOriginState struct {
				inner
			} `json:"service.origin.state"`
			ServiceOriginType struct {
				inner
			} `json:"service.origin.type"`
			ServiceOriginVersion struct {
				inner
			} `json:"service.origin.version"`
			ServiceState struct {
				inner
			} `json:"service.state"`
			ServiceTargetAddress struct {
				inner
			} `json:"service.target.address"`
			ServiceTargetEnvironment struct {
				Beta string `json:"beta"`
				inner
			} `json:"service.target.environment"`
			ServiceTargetEphemeralID struct {
				inner
			} `json:"service.target.ephemeral_id"`
			ServiceTargetID struct {
				inner
			} `json:"service.target.id"`
			ServiceTargetName struct {
				inner
			} `json:"service.target.name"`
			ServiceTargetNodeName struct {
				inner
			} `json:"service.target.node.name"`
			ServiceTargetNodeRole struct {
				inner
			} `json:"service.target.node.role"`
			ServiceTargetNodeRoles struct {
				inner
			} `json:"service.target.node.roles"`
			ServiceTargetState struct {
				inner
			} `json:"service.target.state"`
			ServiceTargetType struct {
				inner
			} `json:"service.target.type"`
			ServiceTargetVersion struct {
				inner
			} `json:"service.target.version"`
			ServiceType struct {
				inner
			} `json:"service.type"`
			ServiceVersion struct {
				inner
			} `json:"service.version"`
		} `json:"fields"`
		Footnote string   `json:"footnote"`
		Group    int      `json:"group"`
		Name     string   `json:"name"`
		Nestings []string `json:"nestings"`
		Prefix   string   `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As            string `json:"as"`
				At            string `json:"at"`
				Beta          string `json:"beta"`
				Full          string `json:"full"`
				ShortOverride string `json:"short_override"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		ReusedHere []struct {
			Beta       string `json:"beta"`
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"service"`
	Source struct {
		Description string `json:"description"`
		Fields      struct {
			SourceAddress struct {
				inner
			} `json:"source.address"`
			SourceAsNumber struct {
				inner
			} `json:"source.as.number"`
			SourceAsOrganizationName struct {
				inner
			} `json:"source.as.organization.name"`
			SourceBytes struct {
				inner
			} `json:"source.bytes"`
			SourceDomain struct {
				inner
			} `json:"source.domain"`
			SourceGeoCityName struct {
				inner
			} `json:"source.geo.city_name"`
			SourceGeoContinentCode struct {
				inner
			} `json:"source.geo.continent_code"`
			SourceGeoContinentName struct {
				inner
			} `json:"source.geo.continent_name"`
			SourceGeoCountryIsoCode struct {
				inner
			} `json:"source.geo.country_iso_code"`
			SourceGeoCountryName struct {
				inner
			} `json:"source.geo.country_name"`
			SourceGeoLocation struct {
				inner
			} `json:"source.geo.location"`
			SourceGeoName struct {
				inner
			} `json:"source.geo.name"`
			SourceGeoPostalCode struct {
				inner
			} `json:"source.geo.postal_code"`
			SourceGeoRegionIsoCode struct {
				inner
			} `json:"source.geo.region_iso_code"`
			SourceGeoRegionName struct {
				inner
			} `json:"source.geo.region_name"`
			SourceGeoTimezone struct {
				inner
			} `json:"source.geo.timezone"`
			SourceIP struct {
				inner
			} `json:"source.ip"`
			SourceMac struct {
				inner
			} `json:"source.mac"`
			SourceNatIP struct {
				inner
			} `json:"source.nat.ip"`
			SourceNatPort struct {
				inner
			} `json:"source.nat.port"`
			SourcePackets struct {
				inner
			} `json:"source.packets"`
			SourcePort struct {
				inner
			} `json:"source.port"`
			SourceRegisteredDomain struct {
				inner
			} `json:"source.registered_domain"`
			SourceSubdomain struct {
				inner
			} `json:"source.subdomain"`
			SourceTopLevelDomain struct {
				inner
			} `json:"source.top_level_domain"`
			SourceUserDomain struct {
				inner
			} `json:"source.user.domain"`
			SourceUserEmail struct {
				inner
			} `json:"source.user.email"`
			SourceUserFullName struct {
				inner
			} `json:"source.user.full_name"`
			SourceUserGroupDomain struct {
				inner
			} `json:"source.user.group.domain"`
			SourceUserGroupID struct {
				inner
			} `json:"source.user.group.id"`
			SourceUserGroupName struct {
				inner
			} `json:"source.user.group.name"`
			SourceUserHash struct {
				inner
			} `json:"source.user.hash"`
			SourceUserID struct {
				inner
			} `json:"source.user.id"`
			SourceUserName struct {
				inner
			} `json:"source.user.name"`
			SourceUserRoles struct {
				inner
			} `json:"source.user.roles"`
		} `json:"fields"`
		Group    int      `json:"group"`
		Name     string   `json:"name"`
		Nestings []string `json:"nestings"`
		Prefix   string   `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As            string `json:"as"`
				At            string `json:"at"`
				Full          string `json:"full"`
				ShortOverride string `json:"short_override"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"source"`
	Threat struct {
		Description string `json:"description"`
		Fields      struct {
			ThreatEnrichments struct {
				inner
			} `json:"threat.enrichments"`
			ThreatEnrichmentsIndicator struct {
				inner
			} `json:"threat.enrichments.indicator"`
			ThreatEnrichmentsIndicatorAsNumber struct {
				inner
			} `json:"threat.enrichments.indicator.as.number"`
			ThreatEnrichmentsIndicatorAsOrganizationName struct {
				inner
			} `json:"threat.enrichments.indicator.as.organization.name"`
			ThreatEnrichmentsIndicatorConfidence struct {
				inner
			} `json:"threat.enrichments.indicator.confidence"`
			ThreatEnrichmentsIndicatorDescription struct {
				inner
			} `json:"threat.enrichments.indicator.description"`
			ThreatEnrichmentsIndicatorEmailAddress struct {
				inner
			} `json:"threat.enrichments.indicator.email.address"`
			ThreatEnrichmentsIndicatorFileAccessed struct {
				inner
			} `json:"threat.enrichments.indicator.file.accessed"`
			ThreatEnrichmentsIndicatorFileAttributes struct {
				inner
			} `json:"threat.enrichments.indicator.file.attributes"`
			ThreatEnrichmentsIndicatorFileCodeSignatureDigestAlgorithm struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.digest_algorithm"`
			ThreatEnrichmentsIndicatorFileCodeSignatureExists struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.exists"`
			ThreatEnrichmentsIndicatorFileCodeSignatureSigningID struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.signing_id"`
			ThreatEnrichmentsIndicatorFileCodeSignatureStatus struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.status"`
			ThreatEnrichmentsIndicatorFileCodeSignatureSubjectName struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.subject_name"`
			ThreatEnrichmentsIndicatorFileCodeSignatureTeamID struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.team_id"`
			ThreatEnrichmentsIndicatorFileCodeSignatureTimestamp struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.timestamp"`
			ThreatEnrichmentsIndicatorFileCodeSignatureTrusted struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.trusted"`
			ThreatEnrichmentsIndicatorFileCodeSignatureValid struct {
				inner
			} `json:"threat.enrichments.indicator.file.code_signature.valid"`
			ThreatEnrichmentsIndicatorFileCreated struct {
				inner
			} `json:"threat.enrichments.indicator.file.created"`
			ThreatEnrichmentsIndicatorFileCtime struct {
				inner
			} `json:"threat.enrichments.indicator.file.ctime"`
			ThreatEnrichmentsIndicatorFileDevice struct {
				inner
			} `json:"threat.enrichments.indicator.file.device"`
			ThreatEnrichmentsIndicatorFileDirectory struct {
				inner
			} `json:"threat.enrichments.indicator.file.directory"`
			ThreatEnrichmentsIndicatorFileDriveLetter struct {
				inner
			} `json:"threat.enrichments.indicator.file.drive_letter"`
			ThreatEnrichmentsIndicatorFileElfArchitecture struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.architecture"`
			ThreatEnrichmentsIndicatorFileElfByteOrder struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.byte_order"`
			ThreatEnrichmentsIndicatorFileElfCPUType struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.cpu_type"`
			ThreatEnrichmentsIndicatorFileElfCreationDate struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.creation_date"`
			ThreatEnrichmentsIndicatorFileElfExports struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.exports"`
			ThreatEnrichmentsIndicatorFileElfGoImportHash struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.go_import_hash"`
			ThreatEnrichmentsIndicatorFileElfGoImports struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.go_imports"`
			ThreatEnrichmentsIndicatorFileElfGoImportsNamesEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.go_imports_names_entropy"`
			ThreatEnrichmentsIndicatorFileElfGoImportsNamesVarEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.go_imports_names_var_entropy"`
			ThreatEnrichmentsIndicatorFileElfGoStripped struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.go_stripped"`
			ThreatEnrichmentsIndicatorFileElfHeaderAbiVersion struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.abi_version"`
			ThreatEnrichmentsIndicatorFileElfHeaderClass struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.class"`
			ThreatEnrichmentsIndicatorFileElfHeaderData struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.data"`
			ThreatEnrichmentsIndicatorFileElfHeaderEntrypoint struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.entrypoint"`
			ThreatEnrichmentsIndicatorFileElfHeaderObjectVersion struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.object_version"`
			ThreatEnrichmentsIndicatorFileElfHeaderOsAbi struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.os_abi"`
			ThreatEnrichmentsIndicatorFileElfHeaderType struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.type"`
			ThreatEnrichmentsIndicatorFileElfHeaderVersion struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.header.version"`
			ThreatEnrichmentsIndicatorFileElfImportHash struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.import_hash"`
			ThreatEnrichmentsIndicatorFileElfImports struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.imports"`
			ThreatEnrichmentsIndicatorFileElfImportsNamesEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.imports_names_entropy"`
			ThreatEnrichmentsIndicatorFileElfImportsNamesVarEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.imports_names_var_entropy"`
			ThreatEnrichmentsIndicatorFileElfSections struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections"`
			ThreatEnrichmentsIndicatorFileElfSectionsChi2 struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.chi2"`
			ThreatEnrichmentsIndicatorFileElfSectionsEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.entropy"`
			ThreatEnrichmentsIndicatorFileElfSectionsFlags struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.flags"`
			ThreatEnrichmentsIndicatorFileElfSectionsName struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.name"`
			ThreatEnrichmentsIndicatorFileElfSectionsPhysicalOffset struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.physical_offset"`
			ThreatEnrichmentsIndicatorFileElfSectionsPhysicalSize struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.physical_size"`
			ThreatEnrichmentsIndicatorFileElfSectionsType struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.type"`
			ThreatEnrichmentsIndicatorFileElfSectionsVarEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.var_entropy"`
			ThreatEnrichmentsIndicatorFileElfSectionsVirtualAddress struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.virtual_address"`
			ThreatEnrichmentsIndicatorFileElfSectionsVirtualSize struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.sections.virtual_size"`
			ThreatEnrichmentsIndicatorFileElfSegments struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.segments"`
			ThreatEnrichmentsIndicatorFileElfSegmentsSections struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.segments.sections"`
			ThreatEnrichmentsIndicatorFileElfSegmentsType struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.segments.type"`
			ThreatEnrichmentsIndicatorFileElfSharedLibraries struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.shared_libraries"`
			ThreatEnrichmentsIndicatorFileElfTelfhash struct {
				inner
			} `json:"threat.enrichments.indicator.file.elf.telfhash"`
			ThreatEnrichmentsIndicatorFileExtension struct {
				inner
			} `json:"threat.enrichments.indicator.file.extension"`
			ThreatEnrichmentsIndicatorFileForkName struct {
				inner
			} `json:"threat.enrichments.indicator.file.fork_name"`
			ThreatEnrichmentsIndicatorFileGid struct {
				inner
			} `json:"threat.enrichments.indicator.file.gid"`
			ThreatEnrichmentsIndicatorFileGroup struct {
				inner
			} `json:"threat.enrichments.indicator.file.group"`
			ThreatEnrichmentsIndicatorFileHashMd5 struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.md5"`
			ThreatEnrichmentsIndicatorFileHashSha1 struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.sha1"`
			ThreatEnrichmentsIndicatorFileHashSha256 struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.sha256"`
			ThreatEnrichmentsIndicatorFileHashSha384 struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.sha384"`
			ThreatEnrichmentsIndicatorFileHashSha512 struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.sha512"`
			ThreatEnrichmentsIndicatorFileHashSsdeep struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.ssdeep"`
			ThreatEnrichmentsIndicatorFileHashTlsh struct {
				inner
			} `json:"threat.enrichments.indicator.file.hash.tlsh"`
			ThreatEnrichmentsIndicatorFileInode struct {
				inner
			} `json:"threat.enrichments.indicator.file.inode"`
			ThreatEnrichmentsIndicatorFileMimeType struct {
				inner
			} `json:"threat.enrichments.indicator.file.mime_type"`
			ThreatEnrichmentsIndicatorFileMode struct {
				inner
			} `json:"threat.enrichments.indicator.file.mode"`
			ThreatEnrichmentsIndicatorFileMtime struct {
				inner
			} `json:"threat.enrichments.indicator.file.mtime"`
			ThreatEnrichmentsIndicatorFileName struct {
				inner
			} `json:"threat.enrichments.indicator.file.name"`
			ThreatEnrichmentsIndicatorFileOwner struct {
				inner
			} `json:"threat.enrichments.indicator.file.owner"`
			ThreatEnrichmentsIndicatorFilePath struct {
				inner
			} `json:"threat.enrichments.indicator.file.path"`
			ThreatEnrichmentsIndicatorFilePeArchitecture struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.architecture"`
			ThreatEnrichmentsIndicatorFilePeCompany struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.company"`
			ThreatEnrichmentsIndicatorFilePeDescription struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.description"`
			ThreatEnrichmentsIndicatorFilePeFileVersion struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.file_version"`
			ThreatEnrichmentsIndicatorFilePeGoImportHash struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.go_import_hash"`
			ThreatEnrichmentsIndicatorFilePeGoImports struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.go_imports"`
			ThreatEnrichmentsIndicatorFilePeGoImportsNamesEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.go_imports_names_entropy"`
			ThreatEnrichmentsIndicatorFilePeGoImportsNamesVarEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.go_imports_names_var_entropy"`
			ThreatEnrichmentsIndicatorFilePeGoStripped struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.go_stripped"`
			ThreatEnrichmentsIndicatorFilePeImphash struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.imphash"`
			ThreatEnrichmentsIndicatorFilePeImportHash struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.import_hash"`
			ThreatEnrichmentsIndicatorFilePeImports struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.imports"`
			ThreatEnrichmentsIndicatorFilePeImportsNamesEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.imports_names_entropy"`
			ThreatEnrichmentsIndicatorFilePeImportsNamesVarEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.imports_names_var_entropy"`
			ThreatEnrichmentsIndicatorFilePeOriginalFileName struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.original_file_name"`
			ThreatEnrichmentsIndicatorFilePePehash struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.pehash"`
			ThreatEnrichmentsIndicatorFilePeProduct struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.product"`
			ThreatEnrichmentsIndicatorFilePeSections struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.sections"`
			ThreatEnrichmentsIndicatorFilePeSectionsEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.sections.entropy"`
			ThreatEnrichmentsIndicatorFilePeSectionsName struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.sections.name"`
			ThreatEnrichmentsIndicatorFilePeSectionsPhysicalSize struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.sections.physical_size"`
			ThreatEnrichmentsIndicatorFilePeSectionsVarEntropy struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.sections.var_entropy"`
			ThreatEnrichmentsIndicatorFilePeSectionsVirtualSize struct {
				inner
			} `json:"threat.enrichments.indicator.file.pe.sections.virtual_size"`
			ThreatEnrichmentsIndicatorFileSize struct {
				inner
			} `json:"threat.enrichments.indicator.file.size"`
			ThreatEnrichmentsIndicatorFileTargetPath struct {
				inner
			} `json:"threat.enrichments.indicator.file.target_path"`
			ThreatEnrichmentsIndicatorFileType struct {
				inner
			} `json:"threat.enrichments.indicator.file.type"`
			ThreatEnrichmentsIndicatorFileUID struct {
				inner
			} `json:"threat.enrichments.indicator.file.uid"`
			ThreatEnrichmentsIndicatorFileX509AlternativeNames struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.alternative_names"`
			ThreatEnrichmentsIndicatorFileX509IssuerCommonName struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.common_name"`
			ThreatEnrichmentsIndicatorFileX509IssuerCountry struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.country"`
			ThreatEnrichmentsIndicatorFileX509IssuerDistinguishedName struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.distinguished_name"`
			ThreatEnrichmentsIndicatorFileX509IssuerLocality struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.locality"`
			ThreatEnrichmentsIndicatorFileX509IssuerOrganization struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.organization"`
			ThreatEnrichmentsIndicatorFileX509IssuerOrganizationalUnit struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.organizational_unit"`
			ThreatEnrichmentsIndicatorFileX509IssuerStateOrProvince struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.issuer.state_or_province"`
			ThreatEnrichmentsIndicatorFileX509NotAfter struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.not_after"`
			ThreatEnrichmentsIndicatorFileX509NotBefore struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.not_before"`
			ThreatEnrichmentsIndicatorFileX509PublicKeyAlgorithm struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.public_key_algorithm"`
			ThreatEnrichmentsIndicatorFileX509PublicKeyCurve struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.public_key_curve"`
			ThreatEnrichmentsIndicatorFileX509PublicKeyExponent struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.public_key_exponent"`
			ThreatEnrichmentsIndicatorFileX509PublicKeySize struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.public_key_size"`
			ThreatEnrichmentsIndicatorFileX509SerialNumber struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.serial_number"`
			ThreatEnrichmentsIndicatorFileX509SignatureAlgorithm struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.signature_algorithm"`
			ThreatEnrichmentsIndicatorFileX509SubjectCommonName struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.common_name"`
			ThreatEnrichmentsIndicatorFileX509SubjectCountry struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.country"`
			ThreatEnrichmentsIndicatorFileX509SubjectDistinguishedName struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.distinguished_name"`
			ThreatEnrichmentsIndicatorFileX509SubjectLocality struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.locality"`
			ThreatEnrichmentsIndicatorFileX509SubjectOrganization struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.organization"`
			ThreatEnrichmentsIndicatorFileX509SubjectOrganizationalUnit struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.organizational_unit"`
			ThreatEnrichmentsIndicatorFileX509SubjectStateOrProvince struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.subject.state_or_province"`
			ThreatEnrichmentsIndicatorFileX509VersionNumber struct {
				inner
			} `json:"threat.enrichments.indicator.file.x509.version_number"`
			ThreatEnrichmentsIndicatorFirstSeen struct {
				inner
			} `json:"threat.enrichments.indicator.first_seen"`
			ThreatEnrichmentsIndicatorGeoCityName struct {
				inner
			} `json:"threat.enrichments.indicator.geo.city_name"`
			ThreatEnrichmentsIndicatorGeoContinentCode struct {
				inner
			} `json:"threat.enrichments.indicator.geo.continent_code"`
			ThreatEnrichmentsIndicatorGeoContinentName struct {
				inner
			} `json:"threat.enrichments.indicator.geo.continent_name"`
			ThreatEnrichmentsIndicatorGeoCountryIsoCode struct {
				inner
			} `json:"threat.enrichments.indicator.geo.country_iso_code"`
			ThreatEnrichmentsIndicatorGeoCountryName struct {
				inner
			} `json:"threat.enrichments.indicator.geo.country_name"`
			ThreatEnrichmentsIndicatorGeoLocation struct {
				inner
			} `json:"threat.enrichments.indicator.geo.location"`
			ThreatEnrichmentsIndicatorGeoName struct {
				inner
			} `json:"threat.enrichments.indicator.geo.name"`
			ThreatEnrichmentsIndicatorGeoPostalCode struct {
				inner
			} `json:"threat.enrichments.indicator.geo.postal_code"`
			ThreatEnrichmentsIndicatorGeoRegionIsoCode struct {
				inner
			} `json:"threat.enrichments.indicator.geo.region_iso_code"`
			ThreatEnrichmentsIndicatorGeoRegionName struct {
				inner
			} `json:"threat.enrichments.indicator.geo.region_name"`
			ThreatEnrichmentsIndicatorGeoTimezone struct {
				inner
			} `json:"threat.enrichments.indicator.geo.timezone"`
			ThreatEnrichmentsIndicatorIP struct {
				inner
			} `json:"threat.enrichments.indicator.ip"`
			ThreatEnrichmentsIndicatorLastSeen struct {
				inner
			} `json:"threat.enrichments.indicator.last_seen"`
			ThreatEnrichmentsIndicatorMarkingTlp struct {
				inner
			} `json:"threat.enrichments.indicator.marking.tlp"`
			ThreatEnrichmentsIndicatorMarkingTlpVersion struct {
				inner
			} `json:"threat.enrichments.indicator.marking.tlp_version"`
			ThreatEnrichmentsIndicatorModifiedAt struct {
				inner
			} `json:"threat.enrichments.indicator.modified_at"`
			ThreatEnrichmentsIndicatorName struct {
				inner
			} `json:"threat.enrichments.indicator.name"`
			ThreatEnrichmentsIndicatorPort struct {
				inner
			} `json:"threat.enrichments.indicator.port"`
			ThreatEnrichmentsIndicatorProvider struct {
				inner
			} `json:"threat.enrichments.indicator.provider"`
			ThreatEnrichmentsIndicatorReference struct {
				inner
			} `json:"threat.enrichments.indicator.reference"`
			ThreatEnrichmentsIndicatorRegistryDataBytes struct {
				inner
			} `json:"threat.enrichments.indicator.registry.data.bytes"`
			ThreatEnrichmentsIndicatorRegistryDataStrings struct {
				inner
			} `json:"threat.enrichments.indicator.registry.data.strings"`
			ThreatEnrichmentsIndicatorRegistryDataType struct {
				inner
			} `json:"threat.enrichments.indicator.registry.data.type"`
			ThreatEnrichmentsIndicatorRegistryHive struct {
				inner
			} `json:"threat.enrichments.indicator.registry.hive"`
			ThreatEnrichmentsIndicatorRegistryKey struct {
				inner
			} `json:"threat.enrichments.indicator.registry.key"`
			ThreatEnrichmentsIndicatorRegistryPath struct {
				inner
			} `json:"threat.enrichments.indicator.registry.path"`
			ThreatEnrichmentsIndicatorRegistryValue struct {
				inner
			} `json:"threat.enrichments.indicator.registry.value"`
			ThreatEnrichmentsIndicatorScannerStats struct {
				inner
			} `json:"threat.enrichments.indicator.scanner_stats"`
			ThreatEnrichmentsIndicatorSightings struct {
				inner
			} `json:"threat.enrichments.indicator.sightings"`
			ThreatEnrichmentsIndicatorType struct {
				inner
			} `json:"threat.enrichments.indicator.type"`
			ThreatEnrichmentsIndicatorURLDomain struct {
				inner
			} `json:"threat.enrichments.indicator.url.domain"`
			ThreatEnrichmentsIndicatorURLExtension struct {
				inner
			} `json:"threat.enrichments.indicator.url.extension"`
			ThreatEnrichmentsIndicatorURLFragment struct {
				inner
			} `json:"threat.enrichments.indicator.url.fragment"`
			ThreatEnrichmentsIndicatorURLFull struct {
				inner
			} `json:"threat.enrichments.indicator.url.full"`
			ThreatEnrichmentsIndicatorURLOriginal struct {
				inner
			} `json:"threat.enrichments.indicator.url.original"`
			ThreatEnrichmentsIndicatorURLPassword struct {
				inner
			} `json:"threat.enrichments.indicator.url.password"`
			ThreatEnrichmentsIndicatorURLPath struct {
				inner
			} `json:"threat.enrichments.indicator.url.path"`
			ThreatEnrichmentsIndicatorURLPort struct {
				inner
			} `json:"threat.enrichments.indicator.url.port"`
			ThreatEnrichmentsIndicatorURLQuery struct {
				inner
			} `json:"threat.enrichments.indicator.url.query"`
			ThreatEnrichmentsIndicatorURLRegisteredDomain struct {
				inner
			} `json:"threat.enrichments.indicator.url.registered_domain"`
			ThreatEnrichmentsIndicatorURLScheme struct {
				inner
			} `json:"threat.enrichments.indicator.url.scheme"`
			ThreatEnrichmentsIndicatorURLSubdomain struct {
				inner
			} `json:"threat.enrichments.indicator.url.subdomain"`
			ThreatEnrichmentsIndicatorURLTopLevelDomain struct {
				inner
			} `json:"threat.enrichments.indicator.url.top_level_domain"`
			ThreatEnrichmentsIndicatorURLUsername struct {
				inner
			} `json:"threat.enrichments.indicator.url.username"`
			ThreatEnrichmentsIndicatorX509AlternativeNames struct {
				inner
			} `json:"threat.enrichments.indicator.x509.alternative_names"`
			ThreatEnrichmentsIndicatorX509IssuerCommonName struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.common_name"`
			ThreatEnrichmentsIndicatorX509IssuerCountry struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.country"`
			ThreatEnrichmentsIndicatorX509IssuerDistinguishedName struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.distinguished_name"`
			ThreatEnrichmentsIndicatorX509IssuerLocality struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.locality"`
			ThreatEnrichmentsIndicatorX509IssuerOrganization struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.organization"`
			ThreatEnrichmentsIndicatorX509IssuerOrganizationalUnit struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.organizational_unit"`
			ThreatEnrichmentsIndicatorX509IssuerStateOrProvince struct {
				inner
			} `json:"threat.enrichments.indicator.x509.issuer.state_or_province"`
			ThreatEnrichmentsIndicatorX509NotAfter struct {
				inner
			} `json:"threat.enrichments.indicator.x509.not_after"`
			ThreatEnrichmentsIndicatorX509NotBefore struct {
				inner
			} `json:"threat.enrichments.indicator.x509.not_before"`
			ThreatEnrichmentsIndicatorX509PublicKeyAlgorithm struct {
				inner
			} `json:"threat.enrichments.indicator.x509.public_key_algorithm"`
			ThreatEnrichmentsIndicatorX509PublicKeyCurve struct {
				inner
			} `json:"threat.enrichments.indicator.x509.public_key_curve"`
			ThreatEnrichmentsIndicatorX509PublicKeyExponent struct {
				inner
			} `json:"threat.enrichments.indicator.x509.public_key_exponent"`
			ThreatEnrichmentsIndicatorX509PublicKeySize struct {
				inner
			} `json:"threat.enrichments.indicator.x509.public_key_size"`
			ThreatEnrichmentsIndicatorX509SerialNumber struct {
				inner
			} `json:"threat.enrichments.indicator.x509.serial_number"`
			ThreatEnrichmentsIndicatorX509SignatureAlgorithm struct {
				inner
			} `json:"threat.enrichments.indicator.x509.signature_algorithm"`
			ThreatEnrichmentsIndicatorX509SubjectCommonName struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.common_name"`
			ThreatEnrichmentsIndicatorX509SubjectCountry struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.country"`
			ThreatEnrichmentsIndicatorX509SubjectDistinguishedName struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.distinguished_name"`
			ThreatEnrichmentsIndicatorX509SubjectLocality struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.locality"`
			ThreatEnrichmentsIndicatorX509SubjectOrganization struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.organization"`
			ThreatEnrichmentsIndicatorX509SubjectOrganizationalUnit struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.organizational_unit"`
			ThreatEnrichmentsIndicatorX509SubjectStateOrProvince struct {
				inner
			} `json:"threat.enrichments.indicator.x509.subject.state_or_province"`
			ThreatEnrichmentsIndicatorX509VersionNumber struct {
				inner
			} `json:"threat.enrichments.indicator.x509.version_number"`
			ThreatEnrichmentsMatchedAtomic struct {
				inner
			} `json:"threat.enrichments.matched.atomic"`
			ThreatEnrichmentsMatchedField struct {
				inner
			} `json:"threat.enrichments.matched.field"`
			ThreatEnrichmentsMatchedID struct {
				inner
			} `json:"threat.enrichments.matched.id"`
			ThreatEnrichmentsMatchedIndex struct {
				inner
			} `json:"threat.enrichments.matched.index"`
			ThreatEnrichmentsMatchedOccurred struct {
				inner
			} `json:"threat.enrichments.matched.occurred"`
			ThreatEnrichmentsMatchedType struct {
				inner
			} `json:"threat.enrichments.matched.type"`
			ThreatFeedDashboardID struct {
				inner
			} `json:"threat.feed.dashboard_id"`
			ThreatFeedDescription struct {
				inner
			} `json:"threat.feed.description"`
			ThreatFeedName struct {
				inner
			} `json:"threat.feed.name"`
			ThreatFeedReference struct {
				inner
			} `json:"threat.feed.reference"`
			ThreatFramework struct {
				inner
			} `json:"threat.framework"`
			ThreatGroupAlias struct {
				inner
			} `json:"threat.group.alias"`
			ThreatGroupID struct {
				inner
			} `json:"threat.group.id"`
			ThreatGroupName struct {
				inner
			} `json:"threat.group.name"`
			ThreatGroupReference struct {
				inner
			} `json:"threat.group.reference"`
			ThreatIndicatorAsNumber struct {
				inner
			} `json:"threat.indicator.as.number"`
			ThreatIndicatorAsOrganizationName struct {
				inner
			} `json:"threat.indicator.as.organization.name"`
			ThreatIndicatorConfidence struct {
				inner
			} `json:"threat.indicator.confidence"`
			ThreatIndicatorDescription struct {
				inner
			} `json:"threat.indicator.description"`
			ThreatIndicatorEmailAddress struct {
				inner
			} `json:"threat.indicator.email.address"`
			ThreatIndicatorFileAccessed struct {
				inner
			} `json:"threat.indicator.file.accessed"`
			ThreatIndicatorFileAttributes struct {
				inner
			} `json:"threat.indicator.file.attributes"`
			ThreatIndicatorFileCodeSignatureDigestAlgorithm struct {
				inner
			} `json:"threat.indicator.file.code_signature.digest_algorithm"`
			ThreatIndicatorFileCodeSignatureExists struct {
				inner
			} `json:"threat.indicator.file.code_signature.exists"`
			ThreatIndicatorFileCodeSignatureSigningID struct {
				inner
			} `json:"threat.indicator.file.code_signature.signing_id"`
			ThreatIndicatorFileCodeSignatureStatus struct {
				inner
			} `json:"threat.indicator.file.code_signature.status"`
			ThreatIndicatorFileCodeSignatureSubjectName struct {
				inner
			} `json:"threat.indicator.file.code_signature.subject_name"`
			ThreatIndicatorFileCodeSignatureTeamID struct {
				inner
			} `json:"threat.indicator.file.code_signature.team_id"`
			ThreatIndicatorFileCodeSignatureTimestamp struct {
				inner
			} `json:"threat.indicator.file.code_signature.timestamp"`
			ThreatIndicatorFileCodeSignatureTrusted struct {
				inner
			} `json:"threat.indicator.file.code_signature.trusted"`
			ThreatIndicatorFileCodeSignatureValid struct {
				inner
			} `json:"threat.indicator.file.code_signature.valid"`
			ThreatIndicatorFileCreated struct {
				inner
			} `json:"threat.indicator.file.created"`
			ThreatIndicatorFileCtime struct {
				inner
			} `json:"threat.indicator.file.ctime"`
			ThreatIndicatorFileDevice struct {
				inner
			} `json:"threat.indicator.file.device"`
			ThreatIndicatorFileDirectory struct {
				inner
			} `json:"threat.indicator.file.directory"`
			ThreatIndicatorFileDriveLetter struct {
				inner
			} `json:"threat.indicator.file.drive_letter"`
			ThreatIndicatorFileElfArchitecture struct {
				inner
			} `json:"threat.indicator.file.elf.architecture"`
			ThreatIndicatorFileElfByteOrder struct {
				inner
			} `json:"threat.indicator.file.elf.byte_order"`
			ThreatIndicatorFileElfCPUType struct {
				inner
			} `json:"threat.indicator.file.elf.cpu_type"`
			ThreatIndicatorFileElfCreationDate struct {
				inner
			} `json:"threat.indicator.file.elf.creation_date"`
			ThreatIndicatorFileElfExports struct {
				inner
			} `json:"threat.indicator.file.elf.exports"`
			ThreatIndicatorFileElfGoImportHash struct {
				inner
			} `json:"threat.indicator.file.elf.go_import_hash"`
			ThreatIndicatorFileElfGoImports struct {
				inner
			} `json:"threat.indicator.file.elf.go_imports"`
			ThreatIndicatorFileElfGoImportsNamesEntropy struct {
				inner
			} `json:"threat.indicator.file.elf.go_imports_names_entropy"`
			ThreatIndicatorFileElfGoImportsNamesVarEntropy struct {
				inner
			} `json:"threat.indicator.file.elf.go_imports_names_var_entropy"`
			ThreatIndicatorFileElfGoStripped struct {
				inner
			} `json:"threat.indicator.file.elf.go_stripped"`
			ThreatIndicatorFileElfHeaderAbiVersion struct {
				inner
			} `json:"threat.indicator.file.elf.header.abi_version"`
			ThreatIndicatorFileElfHeaderClass struct {
				inner
			} `json:"threat.indicator.file.elf.header.class"`
			ThreatIndicatorFileElfHeaderData struct {
				inner
			} `json:"threat.indicator.file.elf.header.data"`
			ThreatIndicatorFileElfHeaderEntrypoint struct {
				inner
			} `json:"threat.indicator.file.elf.header.entrypoint"`
			ThreatIndicatorFileElfHeaderObjectVersion struct {
				inner
			} `json:"threat.indicator.file.elf.header.object_version"`
			ThreatIndicatorFileElfHeaderOsAbi struct {
				inner
			} `json:"threat.indicator.file.elf.header.os_abi"`
			ThreatIndicatorFileElfHeaderType struct {
				inner
			} `json:"threat.indicator.file.elf.header.type"`
			ThreatIndicatorFileElfHeaderVersion struct {
				inner
			} `json:"threat.indicator.file.elf.header.version"`
			ThreatIndicatorFileElfImportHash struct {
				inner
			} `json:"threat.indicator.file.elf.import_hash"`
			ThreatIndicatorFileElfImports struct {
				inner
			} `json:"threat.indicator.file.elf.imports"`
			ThreatIndicatorFileElfImportsNamesEntropy struct {
				inner
			} `json:"threat.indicator.file.elf.imports_names_entropy"`
			ThreatIndicatorFileElfImportsNamesVarEntropy struct {
				inner
			} `json:"threat.indicator.file.elf.imports_names_var_entropy"`
			ThreatIndicatorFileElfSections struct {
				inner
			} `json:"threat.indicator.file.elf.sections"`
			ThreatIndicatorFileElfSectionsChi2 struct {
				inner
			} `json:"threat.indicator.file.elf.sections.chi2"`
			ThreatIndicatorFileElfSectionsEntropy struct {
				inner
			} `json:"threat.indicator.file.elf.sections.entropy"`
			ThreatIndicatorFileElfSectionsFlags struct {
				inner
			} `json:"threat.indicator.file.elf.sections.flags"`
			ThreatIndicatorFileElfSectionsName struct {
				inner
			} `json:"threat.indicator.file.elf.sections.name"`
			ThreatIndicatorFileElfSectionsPhysicalOffset struct {
				inner
			} `json:"threat.indicator.file.elf.sections.physical_offset"`
			ThreatIndicatorFileElfSectionsPhysicalSize struct {
				inner
			} `json:"threat.indicator.file.elf.sections.physical_size"`
			ThreatIndicatorFileElfSectionsType struct {
				inner
			} `json:"threat.indicator.file.elf.sections.type"`
			ThreatIndicatorFileElfSectionsVarEntropy struct {
				inner
			} `json:"threat.indicator.file.elf.sections.var_entropy"`
			ThreatIndicatorFileElfSectionsVirtualAddress struct {
				inner
			} `json:"threat.indicator.file.elf.sections.virtual_address"`
			ThreatIndicatorFileElfSectionsVirtualSize struct {
				inner
			} `json:"threat.indicator.file.elf.sections.virtual_size"`
			ThreatIndicatorFileElfSegments struct {
				inner
			} `json:"threat.indicator.file.elf.segments"`
			ThreatIndicatorFileElfSegmentsSections struct {
				inner
			} `json:"threat.indicator.file.elf.segments.sections"`
			ThreatIndicatorFileElfSegmentsType struct {
				inner
			} `json:"threat.indicator.file.elf.segments.type"`
			ThreatIndicatorFileElfSharedLibraries struct {
				inner
			} `json:"threat.indicator.file.elf.shared_libraries"`
			ThreatIndicatorFileElfTelfhash struct {
				inner
			} `json:"threat.indicator.file.elf.telfhash"`
			ThreatIndicatorFileExtension struct {
				inner
			} `json:"threat.indicator.file.extension"`
			ThreatIndicatorFileForkName struct {
				inner
			} `json:"threat.indicator.file.fork_name"`
			ThreatIndicatorFileGid struct {
				inner
			} `json:"threat.indicator.file.gid"`
			ThreatIndicatorFileGroup struct {
				inner
			} `json:"threat.indicator.file.group"`
			ThreatIndicatorFileHashMd5 struct {
				inner
			} `json:"threat.indicator.file.hash.md5"`
			ThreatIndicatorFileHashSha1 struct {
				inner
			} `json:"threat.indicator.file.hash.sha1"`
			ThreatIndicatorFileHashSha256 struct {
				inner
			} `json:"threat.indicator.file.hash.sha256"`
			ThreatIndicatorFileHashSha384 struct {
				inner
			} `json:"threat.indicator.file.hash.sha384"`
			ThreatIndicatorFileHashSha512 struct {
				inner
			} `json:"threat.indicator.file.hash.sha512"`
			ThreatIndicatorFileHashSsdeep struct {
				inner
			} `json:"threat.indicator.file.hash.ssdeep"`
			ThreatIndicatorFileHashTlsh struct {
				inner
			} `json:"threat.indicator.file.hash.tlsh"`
			ThreatIndicatorFileInode struct {
				inner
			} `json:"threat.indicator.file.inode"`
			ThreatIndicatorFileMimeType struct {
				inner
			} `json:"threat.indicator.file.mime_type"`
			ThreatIndicatorFileMode struct {
				inner
			} `json:"threat.indicator.file.mode"`
			ThreatIndicatorFileMtime struct {
				inner
			} `json:"threat.indicator.file.mtime"`
			ThreatIndicatorFileName struct {
				inner
			} `json:"threat.indicator.file.name"`
			ThreatIndicatorFileOwner struct {
				inner
			} `json:"threat.indicator.file.owner"`
			ThreatIndicatorFilePath struct {
				inner
			} `json:"threat.indicator.file.path"`
			ThreatIndicatorFilePeArchitecture struct {
				inner
			} `json:"threat.indicator.file.pe.architecture"`
			ThreatIndicatorFilePeCompany struct {
				inner
			} `json:"threat.indicator.file.pe.company"`
			ThreatIndicatorFilePeDescription struct {
				inner
			} `json:"threat.indicator.file.pe.description"`
			ThreatIndicatorFilePeFileVersion struct {
				inner
			} `json:"threat.indicator.file.pe.file_version"`
			ThreatIndicatorFilePeGoImportHash struct {
				inner
			} `json:"threat.indicator.file.pe.go_import_hash"`
			ThreatIndicatorFilePeGoImports struct {
				inner
			} `json:"threat.indicator.file.pe.go_imports"`
			ThreatIndicatorFilePeGoImportsNamesEntropy struct {
				inner
			} `json:"threat.indicator.file.pe.go_imports_names_entropy"`
			ThreatIndicatorFilePeGoImportsNamesVarEntropy struct {
				inner
			} `json:"threat.indicator.file.pe.go_imports_names_var_entropy"`
			ThreatIndicatorFilePeGoStripped struct {
				inner
			} `json:"threat.indicator.file.pe.go_stripped"`
			ThreatIndicatorFilePeImphash struct {
				inner
			} `json:"threat.indicator.file.pe.imphash"`
			ThreatIndicatorFilePeImportHash struct {
				inner
			} `json:"threat.indicator.file.pe.import_hash"`
			ThreatIndicatorFilePeImports struct {
				inner
			} `json:"threat.indicator.file.pe.imports"`
			ThreatIndicatorFilePeImportsNamesEntropy struct {
				inner
			} `json:"threat.indicator.file.pe.imports_names_entropy"`
			ThreatIndicatorFilePeImportsNamesVarEntropy struct {
				inner
			} `json:"threat.indicator.file.pe.imports_names_var_entropy"`
			ThreatIndicatorFilePeOriginalFileName struct {
				inner
			} `json:"threat.indicator.file.pe.original_file_name"`
			ThreatIndicatorFilePePehash struct {
				inner
			} `json:"threat.indicator.file.pe.pehash"`
			ThreatIndicatorFilePeProduct struct {
				inner
			} `json:"threat.indicator.file.pe.product"`
			ThreatIndicatorFilePeSections struct {
				inner
			} `json:"threat.indicator.file.pe.sections"`
			ThreatIndicatorFilePeSectionsEntropy struct {
				inner
			} `json:"threat.indicator.file.pe.sections.entropy"`
			ThreatIndicatorFilePeSectionsName struct {
				inner
			} `json:"threat.indicator.file.pe.sections.name"`
			ThreatIndicatorFilePeSectionsPhysicalSize struct {
				inner
			} `json:"threat.indicator.file.pe.sections.physical_size"`
			ThreatIndicatorFilePeSectionsVarEntropy struct {
				inner
			} `json:"threat.indicator.file.pe.sections.var_entropy"`
			ThreatIndicatorFilePeSectionsVirtualSize struct {
				inner
			} `json:"threat.indicator.file.pe.sections.virtual_size"`
			ThreatIndicatorFileSize struct {
				inner
			} `json:"threat.indicator.file.size"`
			ThreatIndicatorFileTargetPath struct {
				inner
			} `json:"threat.indicator.file.target_path"`
			ThreatIndicatorFileType struct {
				inner
			} `json:"threat.indicator.file.type"`
			ThreatIndicatorFileUID struct {
				inner
			} `json:"threat.indicator.file.uid"`
			ThreatIndicatorFileX509AlternativeNames struct {
				inner
			} `json:"threat.indicator.file.x509.alternative_names"`
			ThreatIndicatorFileX509IssuerCommonName struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.common_name"`
			ThreatIndicatorFileX509IssuerCountry struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.country"`
			ThreatIndicatorFileX509IssuerDistinguishedName struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.distinguished_name"`
			ThreatIndicatorFileX509IssuerLocality struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.locality"`
			ThreatIndicatorFileX509IssuerOrganization struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.organization"`
			ThreatIndicatorFileX509IssuerOrganizationalUnit struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.organizational_unit"`
			ThreatIndicatorFileX509IssuerStateOrProvince struct {
				inner
			} `json:"threat.indicator.file.x509.issuer.state_or_province"`
			ThreatIndicatorFileX509NotAfter struct {
				inner
			} `json:"threat.indicator.file.x509.not_after"`
			ThreatIndicatorFileX509NotBefore struct {
				inner
			} `json:"threat.indicator.file.x509.not_before"`
			ThreatIndicatorFileX509PublicKeyAlgorithm struct {
				inner
			} `json:"threat.indicator.file.x509.public_key_algorithm"`
			ThreatIndicatorFileX509PublicKeyCurve struct {
				inner
			} `json:"threat.indicator.file.x509.public_key_curve"`
			ThreatIndicatorFileX509PublicKeyExponent struct {
				inner
			} `json:"threat.indicator.file.x509.public_key_exponent"`
			ThreatIndicatorFileX509PublicKeySize struct {
				inner
			} `json:"threat.indicator.file.x509.public_key_size"`
			ThreatIndicatorFileX509SerialNumber struct {
				inner
			} `json:"threat.indicator.file.x509.serial_number"`
			ThreatIndicatorFileX509SignatureAlgorithm struct {
				inner
			} `json:"threat.indicator.file.x509.signature_algorithm"`
			ThreatIndicatorFileX509SubjectCommonName struct {
				inner
			} `json:"threat.indicator.file.x509.subject.common_name"`
			ThreatIndicatorFileX509SubjectCountry struct {
				inner
			} `json:"threat.indicator.file.x509.subject.country"`
			ThreatIndicatorFileX509SubjectDistinguishedName struct {
				inner
			} `json:"threat.indicator.file.x509.subject.distinguished_name"`
			ThreatIndicatorFileX509SubjectLocality struct {
				inner
			} `json:"threat.indicator.file.x509.subject.locality"`
			ThreatIndicatorFileX509SubjectOrganization struct {
				inner
			} `json:"threat.indicator.file.x509.subject.organization"`
			ThreatIndicatorFileX509SubjectOrganizationalUnit struct {
				inner
			} `json:"threat.indicator.file.x509.subject.organizational_unit"`
			ThreatIndicatorFileX509SubjectStateOrProvince struct {
				inner
			} `json:"threat.indicator.file.x509.subject.state_or_province"`
			ThreatIndicatorFileX509VersionNumber struct {
				inner
			} `json:"threat.indicator.file.x509.version_number"`
			ThreatIndicatorFirstSeen struct {
				inner
			} `json:"threat.indicator.first_seen"`
			ThreatIndicatorGeoCityName struct {
				inner
			} `json:"threat.indicator.geo.city_name"`
			ThreatIndicatorGeoContinentCode struct {
				inner
			} `json:"threat.indicator.geo.continent_code"`
			ThreatIndicatorGeoContinentName struct {
				inner
			} `json:"threat.indicator.geo.continent_name"`
			ThreatIndicatorGeoCountryIsoCode struct {
				inner
			} `json:"threat.indicator.geo.country_iso_code"`
			ThreatIndicatorGeoCountryName struct {
				inner
			} `json:"threat.indicator.geo.country_name"`
			ThreatIndicatorGeoLocation struct {
				inner
			} `json:"threat.indicator.geo.location"`
			ThreatIndicatorGeoName struct {
				inner
			} `json:"threat.indicator.geo.name"`
			ThreatIndicatorGeoPostalCode struct {
				inner
			} `json:"threat.indicator.geo.postal_code"`
			ThreatIndicatorGeoRegionIsoCode struct {
				inner
			} `json:"threat.indicator.geo.region_iso_code"`
			ThreatIndicatorGeoRegionName struct {
				inner
			} `json:"threat.indicator.geo.region_name"`
			ThreatIndicatorGeoTimezone struct {
				inner
			} `json:"threat.indicator.geo.timezone"`
			ThreatIndicatorIP struct {
				inner
			} `json:"threat.indicator.ip"`
			ThreatIndicatorLastSeen struct {
				inner
			} `json:"threat.indicator.last_seen"`
			ThreatIndicatorMarkingTlp struct {
				inner
			} `json:"threat.indicator.marking.tlp"`
			ThreatIndicatorMarkingTlpVersion struct {
				inner
			} `json:"threat.indicator.marking.tlp_version"`
			ThreatIndicatorModifiedAt struct {
				inner
			} `json:"threat.indicator.modified_at"`
			ThreatIndicatorName struct {
				inner
			} `json:"threat.indicator.name"`
			ThreatIndicatorPort struct {
				inner
			} `json:"threat.indicator.port"`
			ThreatIndicatorProvider struct {
				inner
			} `json:"threat.indicator.provider"`
			ThreatIndicatorReference struct {
				inner
			} `json:"threat.indicator.reference"`
			ThreatIndicatorRegistryDataBytes struct {
				inner
			} `json:"threat.indicator.registry.data.bytes"`
			ThreatIndicatorRegistryDataStrings struct {
				inner
			} `json:"threat.indicator.registry.data.strings"`
			ThreatIndicatorRegistryDataType struct {
				inner
			} `json:"threat.indicator.registry.data.type"`
			ThreatIndicatorRegistryHive struct {
				inner
			} `json:"threat.indicator.registry.hive"`
			ThreatIndicatorRegistryKey struct {
				inner
			} `json:"threat.indicator.registry.key"`
			ThreatIndicatorRegistryPath struct {
				inner
			} `json:"threat.indicator.registry.path"`
			ThreatIndicatorRegistryValue struct {
				inner
			} `json:"threat.indicator.registry.value"`
			ThreatIndicatorScannerStats struct {
				inner
			} `json:"threat.indicator.scanner_stats"`
			ThreatIndicatorSightings struct {
				inner
			} `json:"threat.indicator.sightings"`
			ThreatIndicatorType struct {
				inner
			} `json:"threat.indicator.type"`
			ThreatIndicatorURLDomain struct {
				inner
			} `json:"threat.indicator.url.domain"`
			ThreatIndicatorURLExtension struct {
				inner
			} `json:"threat.indicator.url.extension"`
			ThreatIndicatorURLFragment struct {
				inner
			} `json:"threat.indicator.url.fragment"`
			ThreatIndicatorURLFull struct {
				inner
			} `json:"threat.indicator.url.full"`
			ThreatIndicatorURLOriginal struct {
				inner
			} `json:"threat.indicator.url.original"`
			ThreatIndicatorURLPassword struct {
				inner
			} `json:"threat.indicator.url.password"`
			ThreatIndicatorURLPath struct {
				inner
			} `json:"threat.indicator.url.path"`
			ThreatIndicatorURLPort struct {
				inner
			} `json:"threat.indicator.url.port"`
			ThreatIndicatorURLQuery struct {
				inner
			} `json:"threat.indicator.url.query"`
			ThreatIndicatorURLRegisteredDomain struct {
				inner
			} `json:"threat.indicator.url.registered_domain"`
			ThreatIndicatorURLScheme struct {
				inner
			} `json:"threat.indicator.url.scheme"`
			ThreatIndicatorURLSubdomain struct {
				inner
			} `json:"threat.indicator.url.subdomain"`
			ThreatIndicatorURLTopLevelDomain struct {
				inner
			} `json:"threat.indicator.url.top_level_domain"`
			ThreatIndicatorURLUsername struct {
				inner
			} `json:"threat.indicator.url.username"`
			ThreatIndicatorX509AlternativeNames struct {
				inner
			} `json:"threat.indicator.x509.alternative_names"`
			ThreatIndicatorX509IssuerCommonName struct {
				inner
			} `json:"threat.indicator.x509.issuer.common_name"`
			ThreatIndicatorX509IssuerCountry struct {
				inner
			} `json:"threat.indicator.x509.issuer.country"`
			ThreatIndicatorX509IssuerDistinguishedName struct {
				inner
			} `json:"threat.indicator.x509.issuer.distinguished_name"`
			ThreatIndicatorX509IssuerLocality struct {
				inner
			} `json:"threat.indicator.x509.issuer.locality"`
			ThreatIndicatorX509IssuerOrganization struct {
				inner
			} `json:"threat.indicator.x509.issuer.organization"`
			ThreatIndicatorX509IssuerOrganizationalUnit struct {
				inner
			} `json:"threat.indicator.x509.issuer.organizational_unit"`
			ThreatIndicatorX509IssuerStateOrProvince struct {
				inner
			} `json:"threat.indicator.x509.issuer.state_or_province"`
			ThreatIndicatorX509NotAfter struct {
				inner
			} `json:"threat.indicator.x509.not_after"`
			ThreatIndicatorX509NotBefore struct {
				inner
			} `json:"threat.indicator.x509.not_before"`
			ThreatIndicatorX509PublicKeyAlgorithm struct {
				inner
			} `json:"threat.indicator.x509.public_key_algorithm"`
			ThreatIndicatorX509PublicKeyCurve struct {
				inner
			} `json:"threat.indicator.x509.public_key_curve"`
			ThreatIndicatorX509PublicKeyExponent struct {
				inner
			} `json:"threat.indicator.x509.public_key_exponent"`
			ThreatIndicatorX509PublicKeySize struct {
				inner
			} `json:"threat.indicator.x509.public_key_size"`
			ThreatIndicatorX509SerialNumber struct {
				inner
			} `json:"threat.indicator.x509.serial_number"`
			ThreatIndicatorX509SignatureAlgorithm struct {
				inner
			} `json:"threat.indicator.x509.signature_algorithm"`
			ThreatIndicatorX509SubjectCommonName struct {
				inner
			} `json:"threat.indicator.x509.subject.common_name"`
			ThreatIndicatorX509SubjectCountry struct {
				inner
			} `json:"threat.indicator.x509.subject.country"`
			ThreatIndicatorX509SubjectDistinguishedName struct {
				inner
			} `json:"threat.indicator.x509.subject.distinguished_name"`
			ThreatIndicatorX509SubjectLocality struct {
				inner
			} `json:"threat.indicator.x509.subject.locality"`
			ThreatIndicatorX509SubjectOrganization struct {
				inner
			} `json:"threat.indicator.x509.subject.organization"`
			ThreatIndicatorX509SubjectOrganizationalUnit struct {
				inner
			} `json:"threat.indicator.x509.subject.organizational_unit"`
			ThreatIndicatorX509SubjectStateOrProvince struct {
				inner
			} `json:"threat.indicator.x509.subject.state_or_province"`
			ThreatIndicatorX509VersionNumber struct {
				inner
			} `json:"threat.indicator.x509.version_number"`
			ThreatSoftwareAlias struct {
				inner
			} `json:"threat.software.alias"`
			ThreatSoftwareID struct {
				inner
			} `json:"threat.software.id"`
			ThreatSoftwareName struct {
				inner
			} `json:"threat.software.name"`
			ThreatSoftwarePlatforms struct {
				inner
			} `json:"threat.software.platforms"`
			ThreatSoftwareReference struct {
				inner
			} `json:"threat.software.reference"`
			ThreatSoftwareType struct {
				inner
			} `json:"threat.software.type"`
			ThreatTacticID struct {
				inner
			} `json:"threat.tactic.id"`
			ThreatTacticName struct {
				inner
			} `json:"threat.tactic.name"`
			ThreatTacticReference struct {
				inner
			} `json:"threat.tactic.reference"`
			ThreatTechniqueID struct {
				inner
			} `json:"threat.technique.id"`
			ThreatTechniqueName struct {
				inner
			} `json:"threat.technique.name"`
			ThreatTechniqueReference struct {
				inner
			} `json:"threat.technique.reference"`
			ThreatTechniqueSubtechniqueID struct {
				inner
			} `json:"threat.technique.subtechnique.id"`
			ThreatTechniqueSubtechniqueName struct {
				inner
			} `json:"threat.technique.subtechnique.name"`
			ThreatTechniqueSubtechniqueReference struct {
				inner
			} `json:"threat.technique.subtechnique.reference"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"threat"`
	TLS struct {
		Description string `json:"description"`
		Fields      struct {
			TLSCipher struct {
				inner
			} `json:"tls.cipher"`
			TLSClientCertificate struct {
				inner
			} `json:"tls.client.certificate"`
			TLSClientCertificateChain struct {
				inner
			} `json:"tls.client.certificate_chain"`
			TLSClientHashMd5 struct {
				inner
			} `json:"tls.client.hash.md5"`
			TLSClientHashSha1 struct {
				inner
			} `json:"tls.client.hash.sha1"`
			TLSClientHashSha256 struct {
				inner
			} `json:"tls.client.hash.sha256"`
			TLSClientIssuer struct {
				inner
			} `json:"tls.client.issuer"`
			TLSClientJa3 struct {
				inner
			} `json:"tls.client.ja3"`
			TLSClientNotAfter struct {
				inner
			} `json:"tls.client.not_after"`
			TLSClientNotBefore struct {
				inner
			} `json:"tls.client.not_before"`
			TLSClientServerName struct {
				inner
			} `json:"tls.client.server_name"`
			TLSClientSubject struct {
				inner
			} `json:"tls.client.subject"`
			TLSClientSupportedCiphers struct {
				inner
			} `json:"tls.client.supported_ciphers"`
			TLSClientX509AlternativeNames struct {
				inner
			} `json:"tls.client.x509.alternative_names"`
			TLSClientX509IssuerCommonName struct {
				inner
			} `json:"tls.client.x509.issuer.common_name"`
			TLSClientX509IssuerCountry struct {
				inner
			} `json:"tls.client.x509.issuer.country"`
			TLSClientX509IssuerDistinguishedName struct {
				inner
			} `json:"tls.client.x509.issuer.distinguished_name"`
			TLSClientX509IssuerLocality struct {
				inner
			} `json:"tls.client.x509.issuer.locality"`
			TLSClientX509IssuerOrganization struct {
				inner
			} `json:"tls.client.x509.issuer.organization"`
			TLSClientX509IssuerOrganizationalUnit struct {
				inner
			} `json:"tls.client.x509.issuer.organizational_unit"`
			TLSClientX509IssuerStateOrProvince struct {
				inner
			} `json:"tls.client.x509.issuer.state_or_province"`
			TLSClientX509NotAfter struct {
				inner
			} `json:"tls.client.x509.not_after"`
			TLSClientX509NotBefore struct {
				inner
			} `json:"tls.client.x509.not_before"`
			TLSClientX509PublicKeyAlgorithm struct {
				inner
			} `json:"tls.client.x509.public_key_algorithm"`
			TLSClientX509PublicKeyCurve struct {
				inner
			} `json:"tls.client.x509.public_key_curve"`
			TLSClientX509PublicKeyExponent struct {
				inner
			} `json:"tls.client.x509.public_key_exponent"`
			TLSClientX509PublicKeySize struct {
				inner
			} `json:"tls.client.x509.public_key_size"`
			TLSClientX509SerialNumber struct {
				inner
			} `json:"tls.client.x509.serial_number"`
			TLSClientX509SignatureAlgorithm struct {
				inner
			} `json:"tls.client.x509.signature_algorithm"`
			TLSClientX509SubjectCommonName struct {
				inner
			} `json:"tls.client.x509.subject.common_name"`
			TLSClientX509SubjectCountry struct {
				inner
			} `json:"tls.client.x509.subject.country"`
			TLSClientX509SubjectDistinguishedName struct {
				inner
			} `json:"tls.client.x509.subject.distinguished_name"`
			TLSClientX509SubjectLocality struct {
				inner
			} `json:"tls.client.x509.subject.locality"`
			TLSClientX509SubjectOrganization struct {
				inner
			} `json:"tls.client.x509.subject.organization"`
			TLSClientX509SubjectOrganizationalUnit struct {
				inner
			} `json:"tls.client.x509.subject.organizational_unit"`
			TLSClientX509SubjectStateOrProvince struct {
				inner
			} `json:"tls.client.x509.subject.state_or_province"`
			TLSClientX509VersionNumber struct {
				inner
			} `json:"tls.client.x509.version_number"`
			TLSCurve struct {
				inner
			} `json:"tls.curve"`
			TLSEstablished struct {
				inner
			} `json:"tls.established"`
			TLSNextProtocol struct {
				inner
			} `json:"tls.next_protocol"`
			TLSResumed struct {
				inner
			} `json:"tls.resumed"`
			TLSServerCertificate struct {
				inner
			} `json:"tls.server.certificate"`
			TLSServerCertificateChain struct {
				inner
			} `json:"tls.server.certificate_chain"`
			TLSServerHashMd5 struct {
				inner
			} `json:"tls.server.hash.md5"`
			TLSServerHashSha1 struct {
				inner
			} `json:"tls.server.hash.sha1"`
			TLSServerHashSha256 struct {
				inner
			} `json:"tls.server.hash.sha256"`
			TLSServerIssuer struct {
				inner
			} `json:"tls.server.issuer"`
			TLSServerJa3S struct {
				inner
			} `json:"tls.server.ja3s"`
			TLSServerNotAfter struct {
				inner
			} `json:"tls.server.not_after"`
			TLSServerNotBefore struct {
				inner
			} `json:"tls.server.not_before"`
			TLSServerSubject struct {
				inner
			} `json:"tls.server.subject"`
			TLSServerX509AlternativeNames struct {
				inner
			} `json:"tls.server.x509.alternative_names"`
			TLSServerX509IssuerCommonName struct {
				inner
			} `json:"tls.server.x509.issuer.common_name"`
			TLSServerX509IssuerCountry struct {
				inner
			} `json:"tls.server.x509.issuer.country"`
			TLSServerX509IssuerDistinguishedName struct {
				inner
			} `json:"tls.server.x509.issuer.distinguished_name"`
			TLSServerX509IssuerLocality struct {
				inner
			} `json:"tls.server.x509.issuer.locality"`
			TLSServerX509IssuerOrganization struct {
				inner
			} `json:"tls.server.x509.issuer.organization"`
			TLSServerX509IssuerOrganizationalUnit struct {
				inner
			} `json:"tls.server.x509.issuer.organizational_unit"`
			TLSServerX509IssuerStateOrProvince struct {
				inner
			} `json:"tls.server.x509.issuer.state_or_province"`
			TLSServerX509NotAfter struct {
				inner
			} `json:"tls.server.x509.not_after"`
			TLSServerX509NotBefore struct {
				inner
			} `json:"tls.server.x509.not_before"`
			TLSServerX509PublicKeyAlgorithm struct {
				inner
			} `json:"tls.server.x509.public_key_algorithm"`
			TLSServerX509PublicKeyCurve struct {
				inner
			} `json:"tls.server.x509.public_key_curve"`
			TLSServerX509PublicKeyExponent struct {
				inner
			} `json:"tls.server.x509.public_key_exponent"`
			TLSServerX509PublicKeySize struct {
				inner
			} `json:"tls.server.x509.public_key_size"`
			TLSServerX509SerialNumber struct {
				inner
			} `json:"tls.server.x509.serial_number"`
			TLSServerX509SignatureAlgorithm struct {
				inner
			} `json:"tls.server.x509.signature_algorithm"`
			TLSServerX509SubjectCommonName struct {
				inner
			} `json:"tls.server.x509.subject.common_name"`
			TLSServerX509SubjectCountry struct {
				inner
			} `json:"tls.server.x509.subject.country"`
			TLSServerX509SubjectDistinguishedName struct {
				inner
			} `json:"tls.server.x509.subject.distinguished_name"`
			TLSServerX509SubjectLocality struct {
				inner
			} `json:"tls.server.x509.subject.locality"`
			TLSServerX509SubjectOrganization struct {
				inner
			} `json:"tls.server.x509.subject.organization"`
			TLSServerX509SubjectOrganizationalUnit struct {
				inner
			} `json:"tls.server.x509.subject.organizational_unit"`
			TLSServerX509SubjectStateOrProvince struct {
				inner
			} `json:"tls.server.x509.subject.state_or_province"`
			TLSServerX509VersionNumber struct {
				inner
			} `json:"tls.server.x509.version_number"`
			TLSVersion struct {
				inner
			} `json:"tls.version"`
			TLSVersionProtocol struct {
				inner
			} `json:"tls.version_protocol"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"tls"`
	Tracing struct {
		Description string `json:"description"`
		Fields      struct {
			SpanID struct {
				inner
			} `json:"span.id"`
			TraceID struct {
				inner
			} `json:"trace.id"`
			TransactionID struct {
				inner
			} `json:"transaction.id"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Root   bool   `json:"root"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"tracing"`
	URL struct {
		Description string `json:"description"`
		Fields      struct {
			URLDomain struct {
				inner
			} `json:"url.domain"`
			URLExtension struct {
				inner
			} `json:"url.extension"`
			URLFragment struct {
				inner
			} `json:"url.fragment"`
			URLFull struct {
				inner
			} `json:"url.full"`
			URLOriginal struct {
				inner
			} `json:"url.original"`
			URLPassword struct {
				inner
			} `json:"url.password"`
			URLPath struct {
				inner
			} `json:"url.path"`
			URLPort struct {
				inner
			} `json:"url.port"`
			URLQuery struct {
				inner
			} `json:"url.query"`
			URLRegisteredDomain struct {
				inner
			} `json:"url.registered_domain"`
			URLScheme struct {
				inner
			} `json:"url.scheme"`
			URLSubdomain struct {
				inner
			} `json:"url.subdomain"`
			URLTopLevelDomain struct {
				inner
			} `json:"url.top_level_domain"`
			URLUsername struct {
				inner
			} `json:"url.username"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"url"`
	User struct {
		Description string `json:"description"`
		Fields      struct {
			UserChangesDomain struct {
				inner
			} `json:"user.changes.domain"`
			UserChangesEmail struct {
				inner
			} `json:"user.changes.email"`
			UserChangesFullName struct {
				inner
			} `json:"user.changes.full_name"`
			UserChangesGroupDomain struct {
				inner
			} `json:"user.changes.group.domain"`
			UserChangesGroupID struct {
				inner
			} `json:"user.changes.group.id"`
			UserChangesGroupName struct {
				inner
			} `json:"user.changes.group.name"`
			UserChangesHash struct {
				inner
			} `json:"user.changes.hash"`
			UserChangesID struct {
				inner
			} `json:"user.changes.id"`
			UserChangesName struct {
				inner
			} `json:"user.changes.name"`
			UserChangesRoles struct {
				inner
			} `json:"user.changes.roles"`
			UserDomain struct {
				inner
			} `json:"user.domain"`
			UserEffectiveDomain struct {
				inner
			} `json:"user.effective.domain"`
			UserEffectiveEmail struct {
				inner
			} `json:"user.effective.email"`
			UserEffectiveFullName struct {
				inner
			} `json:"user.effective.full_name"`
			UserEffectiveGroupDomain struct {
				inner
			} `json:"user.effective.group.domain"`
			UserEffectiveGroupID struct {
				inner
			} `json:"user.effective.group.id"`
			UserEffectiveGroupName struct {
				inner
			} `json:"user.effective.group.name"`
			UserEffectiveHash struct {
				inner
			} `json:"user.effective.hash"`
			UserEffectiveID struct {
				inner
			} `json:"user.effective.id"`
			UserEffectiveName struct {
				inner
			} `json:"user.effective.name"`
			UserEffectiveRoles struct {
				inner
			} `json:"user.effective.roles"`
			UserEmail struct {
				inner
			} `json:"user.email"`
			UserFullName struct {
				inner
			} `json:"user.full_name"`
			UserGroupDomain struct {
				inner
			} `json:"user.group.domain"`
			UserGroupID struct {
				inner
			} `json:"user.group.id"`
			UserGroupName struct {
				inner
			} `json:"user.group.name"`
			UserHash struct {
				inner
			} `json:"user.hash"`
			UserID struct {
				inner
			} `json:"user.id"`
			UserName struct {
				inner
			} `json:"user.name"`
			UserRiskCalculatedLevel struct {
				inner
			} `json:"user.risk.calculated_level"`
			UserRiskCalculatedScore struct {
				inner
			} `json:"user.risk.calculated_score"`
			UserRiskCalculatedScoreNorm struct {
				inner
			} `json:"user.risk.calculated_score_norm"`
			UserRiskStaticLevel struct {
				inner
			} `json:"user.risk.static_level"`
			UserRiskStaticScore struct {
				inner
			} `json:"user.risk.static_score"`
			UserRiskStaticScoreNorm struct {
				inner
			} `json:"user.risk.static_score_norm"`
			UserRoles struct {
				inner
			} `json:"user.roles"`
			UserTargetDomain struct {
				inner
			} `json:"user.target.domain"`
			UserTargetEmail struct {
				inner
			} `json:"user.target.email"`
			UserTargetFullName struct {
				inner
			} `json:"user.target.full_name"`
			UserTargetGroupDomain struct {
				inner
			} `json:"user.target.group.domain"`
			UserTargetGroupID struct {
				inner
			} `json:"user.target.group.id"`
			UserTargetGroupName struct {
				inner
			} `json:"user.target.group.name"`
			UserTargetHash struct {
				inner
			} `json:"user.target.hash"`
			UserTargetID struct {
				inner
			} `json:"user.target.id"`
			UserTargetName struct {
				inner
			} `json:"user.target.name"`
			UserTargetRoles struct {
				inner
			} `json:"user.target.roles"`
		} `json:"fields"`
		Group    int      `json:"group"`
		Name     string   `json:"name"`
		Nestings []string `json:"nestings"`
		Prefix   string   `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As            string `json:"as"`
				At            string `json:"at"`
				Full          string `json:"full"`
				ShortOverride string `json:"short_override,omitempty"`
				Beta          string `json:"beta,omitempty"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"user"`
	UserAgent struct {
		Description string `json:"description"`
		Fields      struct {
			UserAgentDeviceName struct {
				inner
			} `json:"user_agent.device.name"`
			UserAgentName struct {
				inner
			} `json:"user_agent.name"`
			UserAgentOriginal struct {
				inner
			} `json:"user_agent.original"`
			UserAgentOsFamily struct {
				inner
			} `json:"user_agent.os.family"`
			UserAgentOsFull struct {
				inner
			} `json:"user_agent.os.full"`
			UserAgentOsKernel struct {
				inner
			} `json:"user_agent.os.kernel"`
			UserAgentOsName struct {
				inner
			} `json:"user_agent.os.name"`
			UserAgentOsPlatform struct {
				inner
			} `json:"user_agent.os.platform"`
			UserAgentOsType struct {
				inner
			} `json:"user_agent.os.type"`
			UserAgentOsVersion struct {
				inner
			} `json:"user_agent.os.version"`
			UserAgentVersion struct {
				inner
			} `json:"user_agent.version"`
		} `json:"fields"`
		Group      int      `json:"group"`
		Name       string   `json:"name"`
		Nestings   []string `json:"nestings"`
		Prefix     string   `json:"prefix"`
		ReusedHere []struct {
			Full       string `json:"full"`
			SchemaName string `json:"schema_name"`
			Short      string `json:"short"`
		} `json:"reused_here"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"user_agent"`
	Vlan struct {
		Description string `json:"description"`
		Fields      struct {
			VlanID struct {
				inner
			} `json:"vlan.id"`
			VlanName struct {
				inner
			} `json:"vlan.name"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"vlan"`
	Volume struct {
		Beta        string `json:"beta"`
		Description string `json:"description"`
		Fields      struct {
			VolumeBusType struct {
				inner
			} `json:"volume.bus_type"`
			VolumeDefaultAccess struct {
				inner
			} `json:"volume.default_access"`
			VolumeDeviceName struct {
				inner
			} `json:"volume.device_name"`
			VolumeDeviceType struct {
				inner
			} `json:"volume.device_type"`
			VolumeDosName struct {
				inner
			} `json:"volume.dos_name"`
			VolumeFileSystemType struct {
				inner
			} `json:"volume.file_system_type"`
			VolumeMountName struct {
				inner
			} `json:"volume.mount_name"`
			VolumeNtName struct {
				inner
			} `json:"volume.nt_name"`
			VolumeProductID struct {
				inner
			} `json:"volume.product_id"`
			VolumeProductName struct {
				inner
			} `json:"volume.product_name"`
			VolumeRemovable struct {
				inner
			} `json:"volume.removable"`
			VolumeSerialNumber struct {
				inner
			} `json:"volume.serial_number"`
			VolumeSize struct {
				inner
			} `json:"volume.size"`
			VolumeVendorID struct {
				inner
			} `json:"volume.vendor_id"`
			VolumeVendorName struct {
				inner
			} `json:"volume.vendor_name"`
			VolumeWritable struct {
				inner
			} `json:"volume.writable"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"volume"`
	Vulnerability struct {
		Description string `json:"description"`
		Fields      struct {
			VulnerabilityCategory struct {
				inner
			} `json:"vulnerability.category"`
			VulnerabilityClassification struct {
				inner
			} `json:"vulnerability.classification"`
			VulnerabilityDescription struct {
				inner
			} `json:"vulnerability.description"`
			VulnerabilityEnumeration struct {
				inner
			} `json:"vulnerability.enumeration"`
			VulnerabilityID struct {
				inner
			} `json:"vulnerability.id"`
			VulnerabilityReference struct {
				inner
			} `json:"vulnerability.reference"`
			VulnerabilityReportID struct {
				inner
			} `json:"vulnerability.report_id"`
			VulnerabilityScannerVendor struct {
				inner
			} `json:"vulnerability.scanner.vendor"`
			VulnerabilityScoreBase struct {
				inner
			} `json:"vulnerability.score.base"`
			VulnerabilityScoreEnvironmental struct {
				inner
			} `json:"vulnerability.score.environmental"`
			VulnerabilityScoreTemporal struct {
				inner
			} `json:"vulnerability.score.temporal"`
			VulnerabilityScoreVersion struct {
				inner
			} `json:"vulnerability.score.version"`
			VulnerabilitySeverity struct {
				inner
			} `json:"vulnerability.severity"`
		} `json:"fields"`
		Group  int    `json:"group"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
		Short  string `json:"short"`
		Title  string `json:"title"`
		Type   string `json:"type"`
	} `json:"vulnerability"`
	X509 struct {
		Description string `json:"description"`
		Fields      struct {
			X509AlternativeNames struct {
				inner
			} `json:"x509.alternative_names"`
			X509IssuerCommonName struct {
				inner
			} `json:"x509.issuer.common_name"`
			X509IssuerCountry struct {
				inner
			} `json:"x509.issuer.country"`
			X509IssuerDistinguishedName struct {
				inner
			} `json:"x509.issuer.distinguished_name"`
			X509IssuerLocality struct {
				inner
			} `json:"x509.issuer.locality"`
			X509IssuerOrganization struct {
				inner
			} `json:"x509.issuer.organization"`
			X509IssuerOrganizationalUnit struct {
				inner
			} `json:"x509.issuer.organizational_unit"`
			X509IssuerStateOrProvince struct {
				inner
			} `json:"x509.issuer.state_or_province"`
			X509NotAfter struct {
				inner
			} `json:"x509.not_after"`
			X509NotBefore struct {
				inner
			} `json:"x509.not_before"`
			X509PublicKeyAlgorithm struct {
				inner
			} `json:"x509.public_key_algorithm"`
			X509PublicKeyCurve struct {
				inner
			} `json:"x509.public_key_curve"`
			X509PublicKeyExponent struct {
				inner
			} `json:"x509.public_key_exponent"`
			X509PublicKeySize struct {
				inner
			} `json:"x509.public_key_size"`
			X509SerialNumber struct {
				inner
			} `json:"x509.serial_number"`
			X509SignatureAlgorithm struct {
				inner
			} `json:"x509.signature_algorithm"`
			X509SubjectCommonName struct {
				inner
			} `json:"x509.subject.common_name"`
			X509SubjectCountry struct {
				inner
			} `json:"x509.subject.country"`
			X509SubjectDistinguishedName struct {
				inner
			} `json:"x509.subject.distinguished_name"`
			X509SubjectLocality struct {
				inner
			} `json:"x509.subject.locality"`
			X509SubjectOrganization struct {
				inner
			} `json:"x509.subject.organization"`
			X509SubjectOrganizationalUnit struct {
				inner
			} `json:"x509.subject.organizational_unit"`
			X509SubjectStateOrProvince struct {
				inner
			} `json:"x509.subject.state_or_province"`
			X509VersionNumber struct {
				inner
			} `json:"x509.version_number"`
		} `json:"fields"`
		Group    int    `json:"group"`
		Name     string `json:"name"`
		Prefix   string `json:"prefix"`
		Reusable struct {
			Expected []struct {
				As   string `json:"as"`
				At   string `json:"at"`
				Full string `json:"full"`
			} `json:"expected"`
			TopLevel bool `json:"top_level"`
		} `json:"reusable"`
		Short string `json:"short"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"x509"`
}
