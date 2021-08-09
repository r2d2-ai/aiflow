package config

import (
	"os"
	"strconv"
)

const (
	EnvMappingIgnoreError        = "AIflow_MAPPING_IGNORE_ERRORS"
	EnvMappingIgnoreErrorDefault = false

	EnvMappingSkipMissing        = "AIflow_MAPPING_SKIP_MISSING"
	EnvMappingSkipMissingDefault = false
)

func IsMappingIgnoreErrorsOn() bool {
	ignoreEror := os.Getenv(EnvMappingIgnoreError)
	if len(ignoreEror) <= 0 {
		return EnvMappingIgnoreErrorDefault
	}
	b, _ := strconv.ParseBool(ignoreEror)
	return b
}

func IsMappingSkipMissingOn() bool {
	skip := os.Getenv(EnvMappingSkipMissing)
	if len(skip) <= 0 {
		return EnvMappingSkipMissingDefault
	}
	b, _ := strconv.ParseBool(skip)
	return b
}
