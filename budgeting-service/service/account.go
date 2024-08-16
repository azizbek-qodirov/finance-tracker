package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "budget-service/genprotos"
	"budget-service/storage"

	"github.com/go-redis/redis/v8" // Make sure to install the Redis client: go get github.com/go-redis/redis/v8
)

type AccountService struct {
	storage storage.StorageI
	redis   *redis.Client
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(storage storage.StorageI, redis *redis.Client) *AccountService {
	return &AccountService{storage: storage, redis: redis}
}

func (s *AccountService) GetAccount(ctx context.Context, req *pb.ByUserID) (*pb.AccountGRes, error) {
	// 1. Try to get account from Redis
	cacheKey := fmt.Sprintf("account:%s", req.UserId)
	cachedData, err := s.redis.Get(ctx, cacheKey).Result()

	// 2. If found in Redis, return cached data
	if err == nil {
		var account pb.AccountGRes
		if err := json.Unmarshal([]byte(cachedData), &account); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached account: %v", err)
		}
		return &account, nil
	} else if err != redis.Nil {
		// Log the error only if it's not a "key not found" error
		log.Printf("Error checking Redis cache: %v", err)
	}

	// 3. If not found in Redis, get from MongoDB
	account, err := s.storage.Account().GetAccount(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get account from storage: %v", err)
	}

	// 4. Cache the account in Redis for 10 minutes
	accountJSON, err := json.Marshal(account)
	if err != nil {
		// Log the error, but don't fail the request
		log.Printf("Failed to marshal account for caching: %v", err)
		return account, nil // Return the account even if caching fails
	}

	if err := s.redis.Set(ctx, cacheKey, accountJSON, 10*time.Minute).Err(); err != nil {
		// Log the error, but don't fail the request
		log.Printf("Failed to cache account in Redis: %v", err)
	}

	return account, nil
}
func (s *AccountService) GetBalance(ctx context.Context, req *pb.ByUserID) (*pb.AccountBalanceGRes, error) {
	return s.storage.Account().GetBalance(req)
}

func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.AccountUReq) (*pb.Void, error) {
	return s.storage.Account().UpdateAccount(req)
}

func (s *AccountService) UpdateBalance(ctx context.Context, req *pb.AccountBalanceUReq) (*pb.Void, error) {
	return s.storage.Account().UpdateBalance(req)
}
