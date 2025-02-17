package cli

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// EnvironmentFlag is the Environment Flag
	EnvironmentFlag string = "environment"

	// ReviewBaseDomainFlag is the base domain name for review apps
	ReviewBaseDomainFlag = "review-base-domain"
	// ReviewBaseDomainDefault is the default base domain for review apps
	ReviewBaseDomainDefault = "review.localhost"

	// EnvironmentExperimental is the Experimental Environment name
	EnvironmentExperimental string = "experimental"
	// EnvironmentTest is the Test Environment name
	EnvironmentTest string = "test"
	// EnvironmentDevelopment is the Development Environment name
	EnvironmentDevelopment string = "development"
	// EnvironmentDemo is the GovCloud demo Environment name
	EnvironmentDemo string = "demo"
	// EnvironmentLoadtest is the GovCloud loadtest Environment name
	EnvironmentLoadtest string = "loadtest"
	// EnvironmentExp is the GovCloud exp Environment name
	EnvironmentExp string = "exp"
	// EnvironmentStg is the GovCloud stg Environment name
	EnvironmentStg string = "stg"
	// EnvironmentPrd is the GovCloud prd Environment name
	EnvironmentPrd string = "prd"
	// EnvironmentReview is a reviewapp
	EnvironmentReview string = "review"
)

var environments = []string{
	EnvironmentExperimental,
	EnvironmentTest,
	EnvironmentDevelopment,
	EnvironmentDemo,
	EnvironmentLoadtest,
	EnvironmentExp,
	EnvironmentStg,
	EnvironmentPrd,
	EnvironmentReview,
}

// InitEnvironmentFlags initializes the Environment command line flags
func InitEnvironmentFlags(flag *pflag.FlagSet) {
	flag.StringP(EnvironmentFlag, "e", EnvironmentDevelopment, fmt.Sprintf("The environment name, one of %v", environments))
	flag.String(ReviewBaseDomainFlag, ReviewBaseDomainDefault, fmt.Sprintf("The base domain name for review apps, defaults to %s", ReviewBaseDomainDefault))
}

// CheckEnvironment validates the Environment command line flags
func CheckEnvironment(v *viper.Viper) error {
	if environment := v.GetString(EnvironmentFlag); !stringSliceContains(environments, environment) {
		return fmt.Errorf("invalid environment %s, expecting one of %q", environment, environments)
	}
	return nil
}
