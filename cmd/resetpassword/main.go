package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
		utils.Log.Fatal("Failed to Connect to MongoDB", zap.Error(err))
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			utils.Log.Error("Failed to Disconnect MongoDB", zap.Error(err))
		}
	}()

	fmt.Println("=====================================")
	fmt.Println("==== ZeroAxiis Rest Password CLI ====")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		utils.Log.Fatal("Failed to Read Email", zap.Error(err))
	}

	fmt.Print("Enter New Password: ")
	password, err := reader.ReadString('\n')

	if err != nil {
		utils.Log.Fatal("Failed to Read Password", zap.Error(err))
	}

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	adminCollection := client.Database("zeroaxiiscom").Collection("admin")

	var admin models.Admin
	result := adminCollection.FindOne(
		context.Background(),
		bson.M{
			"email": email,
		},
	)
	if errors.Is(err, mongo.ErrNoDocuments) {
		utils.Log.Fatal("Admin not Found")
	}
	err = result.Decode(&admin)
	if err != nil {
		utils.Log.Fatal("Failed to find Admin", zap.Error(err))
	}
	utils.Log.Info("Admin Found Successfully...!")
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		utils.Log.Fatal("Failed to hash password")
	}

	updateResult, err := adminCollection.UpdateOne(
		context.Background(),
		bson.M{
			"email": email,
		},
		bson.M{
			"$set": bson.M{
				"password":  hashedPassword,
				"updatedAt": time.Now(),
			},
		},
	)

	if err != nil {
		utils.Log.Fatal("Failed to Update Password", zap.Error(err))
	}

	utils.Log.Info(
		"Password Reset Successfully",
		zap.Int64("Matched", updateResult.MatchedCount),
		zap.Int64("modified", updateResult.ModifiedCount),
	)

}
