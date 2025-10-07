package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph"
	recordmapper "github.com/MamangRust/paymentgatewaygraphql/internal/mapper/record"
	"github.com/MamangRust/paymentgatewaygraphql/internal/mapper/response/graphql"
	responseservice "github.com/MamangRust/paymentgatewaygraphql/internal/mapper/response/service"
	"github.com/MamangRust/paymentgatewaygraphql/internal/middlewares"
	"github.com/MamangRust/paymentgatewaygraphql/internal/permission"
	"github.com/MamangRust/paymentgatewaygraphql/internal/repository"
	"github.com/MamangRust/paymentgatewaygraphql/internal/service"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/auth"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/database"
	db "github.com/MamangRust/paymentgatewaygraphql/pkg/database/schema"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/database/seeder"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/dotenv"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/hash"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/logger"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
)

const defaultPort = "8080"

type Server struct {
	Logger       logger.LoggerInterface
	DB           *db.Queries
	TokenManager *auth.Manager
	Services     *service.Service
	Resolver     *graph.Resolver
	Ctx          context.Context
	Port         string
}

func NewServer() (*Server, error) {
	lg, err := logger.NewLogger()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	if err := dotenv.Viper(); err != nil {
		lg.Fatal("Failed to load .env file", zap.Error(err))
	}

	tokenManager, err := auth.NewManager(viper.GetString("SECRET_KEY"), lg)
	if err != nil {
		lg.Fatal("Failed to create token manager", zap.Error(err))
	}

	conn, err := database.NewClient(lg)
	if err != nil {
		lg.Fatal("Failed to connect to database", zap.Error(err))
	}
	DB := db.New(conn)

	ctx := context.Background()

	hashing := hash.NewHashingPassword()
	mapperRecord := recordmapper.NewRecordMapper()
	mapperResponse := responseservice.NewResponseServiceMapper()
	mapperGraphql := graphql.NewGraphqlMapper()

	repos := repository.NewRepositories(repository.Deps{
		DB:           DB,
		Ctx:          ctx,
		MapperRecord: mapperRecord,
	})

	services := service.NewService(service.Deps{
		Repositories: repos,
		Hash:         hashing,
		Token:        tokenManager,
		Logger:       lg,
		Mapper:       *mapperResponse,
	})

	permission := permission.NewPermission(services.Role, services.Merchant)

	resolver := graph.NewResolver(
		services.Auth,
		services.Role,
		services.User,
		services.Card,
		services.Merchant,
		services.Saldo,
		services.Topup,
		services.Transaction,
		services.Transfer,
		services.Withdraw,
		mapperGraphql,
		permission,
	)

	if viper.GetBool("DB_SEEDER") {
		lg.Debug("Running database seeder")

		s := seeder.NewSeeder(seeder.Deps{
			DB:     DB,
			Ctx:    ctx,
			Logger: lg,
			Hash:   hashing,
		})

		if err := s.Run(); err != nil {
			lg.Fatal("Failed to run seeder", zap.Error(err))
		}
	}

	port := viper.GetString("PORT")
	if port == "" {
		port = defaultPort
	}

	return &Server{
		Logger:       lg,
		DB:           DB,
		TokenManager: tokenManager,
		Services:     services,
		Ctx:          ctx,
		Port:         port,
		Resolver:     resolver,
	}, nil
}

func (s *Server) Run() error {
	s.Logger.Debug("Starting GraphQL server", zap.Any("port", s.Port))

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: s.Resolver,
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", middlewares.AuthMiddleware(s.TokenManager, s.Logger)(srv))

	s.Logger.Debug("GraphQL Playground running at", zap.String("url", "http://localhost:"+s.Port))
	return http.ListenAndServe(":"+s.Port, nil)
}
