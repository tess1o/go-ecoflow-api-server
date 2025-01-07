package constants

const (
	ErrMandatoryHeaderMissing = "0001"
	ErrInvalidAuthHeader      = "0002"
	ErrInvalidJsonBody        = "0003"
	ErrInvalidParameters      = "0004"
	ErrRateLimitExceeded      = "0005"

	ErrGetDevicesList         = "0100"
	ErrGetAllDeviceParameters = "0101"
	ErrGetDeviceParameters    = "0102"

	ErrEnableCarOut = "0200"
	ErrEnableDcOut  = "0201"
	ErrEnableAcOut  = "0203"

	ErrPowerStationSetChargingSpeed = "0204"
	ErrPowerStationSetCarInput      = "0205"
	ErrPowerStationSetStandBy       = "0206"
)
