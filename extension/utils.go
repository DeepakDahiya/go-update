package extension

import (
	"os"
)

var torClientMacExtensionID = "cldoidikboihgcjfkhdeidbpclkineef"
var torClientWindowsExtensionID = "cpoalefficncklhjfpglfiplenlpccdb"
var torClientLinuxExtensionID = "biahpgbdmdkfgndcmfiipgcebobojjkp"
var torClientLinuxArm64ExtensionID = "monolafkoghdlanndjfeebmdfkbklejg"

var torPluggableTransportsMacExtensionID = "einfndjnccmoohcngmlldpmellegjjnk"
var torPluggableTransportsWindowsExtensionID = "dnkcahhmfcanmkjhnjejoomdihffoefm"
var torPluggableTransportsLinuxExtensionID = "apfggiafobakjahnkchiecbomjgigkkn"

// TorClientExtensionIDs is used to add an exception to return the dedicated
// proxy url for downloading the tor client crx
var TorClientExtensionIDs = []string{torClientMacExtensionID, torClientWindowsExtensionID, torClientLinuxExtensionID, torClientLinuxArm64ExtensionID}
var TorPluggableTransportsExtensionIDs = []string{torPluggableTransportsMacExtensionID, torPluggableTransportsWindowsExtensionID, torPluggableTransportsLinuxExtensionID}

var ipfsClientMacExtensionID = "nljcddpbnaianmglkpkneakjaapinabi"
var ipfsClientWindowsExtensionID = "lnbclahgobmjphilkalbhebakmblnbij"
var ipfsClientLinuxExtensionID = "oecghfpdmkjlhnfpmmjegjacfimiafjp"
var ipfsClientMacArm64ExtensionID = "lejaflgbgglfaomemffoaappaihfligf"
var ipfsClientLinuxArm64ExtensionID = "fmmldihckdnognaabhligdpckkeancng"
var ipfsClientWindowsArm64ExtensionID = "nbhgdecjnlofkpafjpcjfajdceeppaih"

// ipfsClientExtensionIDs is used to add an exception to return the dedicated
// proxy url for downloading the ipfs crx
var ipfsClientExtensionIDs = []string{ipfsClientMacExtensionID, ipfsClientWindowsExtensionID, ipfsClientLinuxExtensionID, ipfsClientMacArm64ExtensionID, ipfsClientLinuxArm64ExtensionID, ipfsClientWindowsArm64ExtensionID}

func isTorExtension(id string) bool {
	for _, torID := range TorClientExtensionIDs {
		if torID == id {
			return true
		}
	}
	for _, torPtID := range TorPluggableTransportsExtensionIDs {
		if torPtID == id {
			return true
		}
	}
	return false
}

func isIPFSExtension(id string) bool {
	for _, ipfsID := range ipfsClientExtensionIDs {
		if ipfsID == id {
			return true
		}
	}
	return false
}

func lookupEnvFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetS3ExtensionBucketHost returns the url to use for accessing crx files
func GetS3ExtensionBucketHost(id string) string {
	if isTorExtension(id) {
		return GetS3TorExtensionBucketHost()
	}

	if isIPFSExtension(id) {
		return GetS3IPFSExtensionBucketHost()
	}

	return lookupEnvFallback("S3_EXTENSIONS_BUCKET_HOST", "express-core-ext.s3.us-west-2.amazonaws.com")
}

// GetS3TorExtensionBucketHost returns the url to use for accessing tor client crx
func GetS3TorExtensionBucketHost() string {
	return lookupEnvFallback("S3_EXTENSIONS_BUCKET_HOST_TOR", "tor.bravesoftware.com")
}

// GetS3IPFSExtensionBucketHost returns the url to use for accessing go-ipfs client crx
func GetS3IPFSExtensionBucketHost() string {
	return lookupEnvFallback("S3_EXTENSIONS_BUCKET_HOST_IPFS", "ipfs.bravesoftware.com")
}

// GetUpdateStatus returns the status of an update response for an extension
func GetUpdateStatus(extension Extension) string {
	if extension.Status == "" {
		return "ok"
	}
	return extension.Status
}

// GetComponentUpdaterHost returns the url to use for extension updates
func GetComponentUpdaterHost() string {
	return lookupEnvFallback("COMPONENT_UPDATER_HOST", "componentupdater.brave.com")
}
