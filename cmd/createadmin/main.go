package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zeroaxiis/ZeroAxiis-Services/internal/config"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/models"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()

	err := utils.Init(cfg.AppEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Log.Sync()

	client, err := database.ConnectMongo(cfg.MongoURI)
	if err != nil {
		utils.Log.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			utils.Log.Error("Failed to disconnect MongoDB", zap.Error(err))
		}
	}()

	fmt.Println("=====================================")
	fmt.Println("==== ZeroAxiis Create Admin CLI ====")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		utils.Log.Fatal("Failed to read email", zap.Error(err))
	}

	fmt.Print("Enter Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		utils.Log.Fatal("Failed to read name", zap.Error(err))
	}

	fmt.Print("Enter Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		utils.Log.Fatal("Failed to read password", zap.Error(err))
	}

	email = strings.TrimSpace(email)
	name = strings.TrimSpace(name)
	password = strings.TrimSpace(password)

	adminCollection := client.Database("zeroaxiiscom").Collection("admin")

	var admin models.Admin

	result := adminCollection.FindOne(
		context.Background(),
		bson.M{
			"email": email,
		},
	)

	err = result.Decode(&admin)

	if err == nil {
		utils.Log.Fatal("Admin already exists")
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		utils.Log.Info("No existing admin found")
	} else {
		utils.Log.Fatal("Failed to check existing admin", zap.Error(err))
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		utils.Log.Fatal("Failed to hash password", zap.Error(err))
	}

	admin = models.Admin{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	finalResult, err := adminCollection.InsertOne(
		context.Background(),
		admin,
	)
	if err != nil {
		utils.Log.Fatal("Failed to create admin", zap.Error(err))
	}

	utils.Log.Info(
		"Admin Created Successfully...!",
		zap.Any("id", finalResult.InsertedID),
	)
}