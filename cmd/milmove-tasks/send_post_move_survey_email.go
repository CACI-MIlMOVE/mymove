package main

import (
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/cli"
	"github.com/transcom/mymove/pkg/logging"
	"github.com/transcom/mymove/pkg/notifications"
)

const (
	offsetFlag string = "offset-days"
)

func checkPostMoveSurveyConfig(v *viper.Viper, logger *zap.Logger) error {

	logger.Debug("checking config")

	err := cli.CheckDatabase(v, logger)
	if err != nil {
		return err
	}

	return cli.CheckEmail(v)
}

func initPostMoveSurveyFlags(flag *pflag.FlagSet) {

	// DB Config
	cli.InitDatabaseFlags(flag)

	// Logging Levels
	cli.InitLoggingFlags(flag)

	// Email
	cli.InitEmailFlags(flag)

	flag.Int(offsetFlag, 15, "Number of days ago moves had their payment request reviewed")

	// Don't sort flags
	flag.SortFlags = false
}

// Command: go run github.com/transcom/mymove/cmd/send_post_move_survey_email
func sendPostMoveSurvey(cmd *cobra.Command, args []string) error {

	err := cmd.ParseFlags(args)
	if err != nil {
		return errors.Wrap(err, "Could not parse args")
	}
	flags := cmd.Flags()
	v := viper.New()
	err = v.BindPFlags(flags)
	if err != nil {
		return errors.Wrap(err, "Could not bind flags")
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	dbEnv := v.GetString(cli.DbEnvFlag)
	offsetDays := v.GetInt(offsetFlag)

	logger, _, err := logging.Config(
		logging.WithEnvironment(dbEnv),
		logging.WithLoggingLevel(v.GetString(cli.LoggingLevelFlag)),
		logging.WithStacktraceLength(v.GetInt(cli.StacktraceLengthFlag)),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Zap logging due to %v", err)
	}
	zap.ReplaceGlobals(logger)

	err = checkPostMoveSurveyConfig(v, logger)
	if err != nil {
		logger.Fatal("invalid configuration", zap.Error(err))
	}

	// Create a connection to the DB
	dbConnection, err := cli.InitDatabase(v, logger)
	if err != nil {
		logger.Fatal("Connecting to DB", zap.Error(err))
	}

	appCtx := appcontext.NewAppContext(dbConnection, logger, nil)

	targetDate := time.Now().AddDate(0, 0, -offsetDays)
	notificationSender, notificationSenderErr := notifications.InitEmail(v, logger)
	if notificationSenderErr != nil {
		logger.Fatal("notification sender sending not enabled", zap.Error(notificationSenderErr))
	}

	moveReviewedNotifier, err := notifications.NewMoveReviewed(targetDate)
	if err != nil {
		logger.Fatal("initializing MoveReviewed", zap.Error(err))
	}
	err = notificationSender.SendNotification(appCtx, moveReviewedNotifier)
	if err != nil {
		logger.Fatal("Emails failed to send", zap.Error(err))
	}
	return nil
}
