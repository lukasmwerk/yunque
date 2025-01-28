package server

import "github.com/lukasmwerk/yunque/libraries/types"

// FeatureFlags defines toggles for features that should be enabled or disabled.
// These features may be toggled off at runtime, either by the developer or
// automatically by the application due to a critical failure.
type FeatureFlags struct {
	RunMode types.RunState
	LogMode int

	AutomatedCoding bool
}

// newFeatureFlags returns a default config for features
func NewFeatureFlags() *FeatureFlags {
	return &FeatureFlags{AutomatedCoding: false}
}

// Reset resets the feature flags to the default config
func (ff *FeatureFlags) Reset() {

}

// IsValid checks if enabled features have their dependencies
// enabled. For immediate resolution of dependencies use MakeValid
func (ff *FeatureFlags) IsValid() bool {
	return true
}

// MakeValid disables features whose dependencies are already disabled
func (ff *FeatureFlags) MakeValid() {
}
