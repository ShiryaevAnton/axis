package app

import (
	"os"

	userservice "github.com/ShiryaevAnton/axis/axis_api_users/internal/domain/services/userservices"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/infrastructure/logger"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/infrastructure/persistence/userpersistence"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/infrastructure/usergrpc"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/interface/apigrpc/userserver"
	"github.com/ShiryaevAnton/axis/axis_api_users/internal/usecases/userusecases"
	"github.com/ShiryaevAnton/axis/axis_api_utils/apistore"
)

//StartApp ...
func StartApp() {

	username := getEnvWithDefault("user_pq_username", "postgres")
	password := getEnvWithDefault("user_pq_password", "1234")
	hostDB := getEnvWithDefault("user_pq_host", "localhost")
	portDB := getEnvWithDefault("user_pq_port", "5432")
	dbname := getEnvWithDefault("user_pq_db", "users")
	logLevel := getEnvWithDefault("user_log_Level", "info")
	logOutput := getEnvWithDefault("user_log_output", "stdout")
	hostGRPC := getEnvWithDefault("user_grpc_host", "localhost:50052")

	if err := logger.NewLogger(logLevel, logOutput); err != nil {
		logger.GetLogger().Fatal("logger fatal error", err)
	}

	logger.GetLogger().Info("Logger is connected")

	//TODO: change to interface
	db, err := apistore.InitDatabase(username, password, hostDB, portDB, dbname)
	if err != nil {
		logger.GetLogger().Fatal("database fatal error", err)
	}

	logger.GetLogger().Info("DB is connected")

	userRepo := userpersistence.NewUserRepository(db)
	userService := userservice.NewUserService()
	userUseCase := userusecases.NewUserUseCase(userRepo, userService)
	userServer := userserver.NewUserServer(userUseCase)

	usergrpc.NewGRPCLogger(logger.GetLogger())

	if err := usergrpc.StartGRPC(hostGRPC, userServer); err != nil {
		logger.GetLogger().Fatal("GRPC fatal error", err)
	}

}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
