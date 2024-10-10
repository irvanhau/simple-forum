package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"situs-forum/internal/configs"
	"situs-forum/internal/handlers/memberships"
	"situs-forum/internal/handlers/posts"
	membershipsRepo "situs-forum/internal/repository/memberships"
	postsRepo "situs-forum/internal/repository/posts"
	membershipsSvc "situs-forum/internal/service/memberships"
	postsSvc "situs-forum/internal/service/posts"
	"situs-forum/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFoler([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config: ", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipsRepo.NewRepository(db)
	postRepo := postsRepo.NewRepository(db)

	membershipService := membershipsSvc.NewService(membershipRepo, cfg)
	postService := postsSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	postHandler := posts.NewHandler(postService, r)

	membershipHandler.RegisterRoute()
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
