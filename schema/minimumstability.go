package schema

// The minimum stability the packages must have to be install-able. Possible values are:
// dev, alpha, beta, RC, stable.
type MinimumStability string
const (
	Alpha MinimumStability = "alpha"
	Beta MinimumStability = "beta"
	Dev MinimumStability = "dev"
	MinimumStabilityRC MinimumStability = "RC"
	RC MinimumStability = "rc"
	Stable MinimumStability = "stable"
)
