package aws

const (
	// AvailabilityZoneType is the type of regular zone placed on the region.
	AvailabilityZoneType = "availability-zone"
	// LocalZoneType is the type of Local zone placed on the metropolitan areas.
	LocalZoneType = "local-zone"
	// WavelengtyZoneType is the type of Wavelength zone placed in the Carrier infrastructure closer to areas.
	WavelengtyZoneType = "wavelength-zone"
	// ZoneOptInStatusOptedIn is the opt-in status of the zone.
	// For Availability Zones, this parameter always has the value of opt-in-not-required.
	// For Local Zones and Wavelength Zones, this parameter is the opt-in status.
	ZoneOptInStatusOptedIn = "opted-in"
)
