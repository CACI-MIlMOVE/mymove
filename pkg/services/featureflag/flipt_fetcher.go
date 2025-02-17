package featureflag

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"go.flipt.io/flipt/rpc/flipt"
	sdk "go.flipt.io/flipt/sdk/go"
	sdkhttp "go.flipt.io/flipt/sdk/go/http"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/cli"
	"github.com/transcom/mymove/pkg/services"
)

const (
	applicationName = "applicationName"
	isAdminUser     = "isAdminUser"
	isOfficeUser    = "isOfficeUser"
	isServiceMember = "isServiceMember"
	email           = "email"
)

type FliptFetcher struct {
	client sdk.SDK
	config cli.FeatureFlagConfig
}

func NewFliptFetcher(config cli.FeatureFlagConfig) (*FliptFetcher, error) {
	// For reasons I do not fully understand, trying to resolve the
	// service name in AWS ECS results in a panic(!!!) in the default
	// go resolver. Setting up this custom http client with the custom
	// resolver settings seems to work. ¯\_(ツ)_/¯
	//
	// ahobson - 2023-08-01
	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			// Inspect the network connection type
			DialContext: (&net.Dialer{
				Resolver: &net.Resolver{
					PreferGo:     true,
					StrictErrors: false,
				},
			}).DialContext,
		},
	}

	return NewFliptFetcherWithClient(config, client)
}

func NewFliptFetcherWithClient(config cli.FeatureFlagConfig, httpClient *http.Client) (*FliptFetcher, error) {
	if config.URL == "" {
		return nil, errors.New("FliptFetcher needs a non-empty Endpoint")
	}
	sdkOptions := []sdk.Option{}
	if config.Token != "" {
		// if flipt is not exposed to the internet, we can run it
		// without authentication
		provider := sdk.StaticClientTokenProvider(config.Token)
		sdkOptions = append(sdkOptions, sdk.WithClientTokenProvider(provider))
	}
	transportOptions := []sdkhttp.Option{}
	if httpClient != nil {
		transportOptions = append(transportOptions, sdkhttp.WithHTTPClient(httpClient))
	}
	transport := sdkhttp.NewTransport(config.URL, transportOptions...)
	client := sdk.New(transport, sdkOptions...)
	return &FliptFetcher{
		client: client,
		config: config,
	}, nil
}

func (ff *FliptFetcher) GetFlagForUser(ctx context.Context, appCtx appcontext.AppContext, key string, flagContext map[string]string) (services.FeatureFlag, error) {
	if nil == appCtx.Session() {
		featureFlag := services.FeatureFlag{}
		// if getting a flag for a user, a session must exist
		return featureFlag, errors.New("Nil session when calling GetFlagForUser")
	}

	entityID := appCtx.Session().UserID.String()

	// automatically set the context
	featureFlagContext := flagContext
	featureFlagContext[email] = appCtx.Session().Email
	featureFlagContext[applicationName] = string(appCtx.Session().ApplicationName)

	featureFlagContext[isAdminUser] = strconv.FormatBool(appCtx.Session().IsAdminUser())
	featureFlagContext[isOfficeUser] = strconv.FormatBool(appCtx.Session().IsOfficeUser())
	featureFlagContext[isServiceMember] = strconv.FormatBool(appCtx.Session().IsServiceMember())

	// instead of sending roles, send permissions as that is more
	// granular and flexible
	permissions := appCtx.Session().Permissions
	for i := range permissions {
		featureFlagContext["permissions."+permissions[i]] = strconv.FormatBool(true)
	}

	return ff.GetFlag(ctx, appCtx.Logger(), entityID, key, flagContext)
}

func (ff *FliptFetcher) GetFlag(ctx context.Context, logger *zap.Logger, entityID string, key string, flagContext map[string]string) (services.FeatureFlag, error) {

	// defaults in case the flag is not found
	featureFlag := services.FeatureFlag{
		Entity:    entityID,
		Key:       key,
		Match:     false,
		Namespace: ff.config.Namespace,
	}
	req := &flipt.EvaluationRequest{
		RequestId:    uuid.Must(uuid.NewV4()).String(),
		NamespaceKey: ff.config.Namespace,
		FlagKey:      key,
		EntityId:     entityID,
		Context:      flagContext,
	}
	logger.Debug("flipt evaluation request", zap.Any("req", req))
	result, err := ff.client.Flipt().Evaluate(ctx, req)

	if err != nil {
		logger.Warn("Flipt error", zap.Error(err))
		if s, ok := status.FromError(err); ok && s.Code() == codes.NotFound {
			// treat a missing feature flag as a disabled one
			logger.Warn("Feature flag not found",
				zap.String("key", key),
				zap.String("namespace", ff.config.Namespace))
			return featureFlag, nil
		}
		return featureFlag, err
	}

	featureFlag.Match = result.Match
	featureFlag.Value = result.Value

	return featureFlag, nil
}

func (ff *FliptFetcher) GetConfig() cli.FeatureFlagConfig {
	return ff.config
}
