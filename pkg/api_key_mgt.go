package pkg

import (
	"context"
	"github.com/lbrictson/cogs/ent"
	"sync"
)

// API usage can be quite heavy, to prevent the database from being hammered we keep the API key in memory
var apiKeyCache = make(map[string]UserModel)
var apiKeyCacheLock sync.Mutex

func seedAPIKeyCache(ctx context.Context, db *ent.Client) error {
	users, err := getUsers(ctx, db)
	if err != nil {
		return err
	}
	for _, user := range users {
		addAPIKeyToCache(user.APIKey, user)
	}
	return nil
}

// validateAPIKey will return true of the api key is found in the cache along with the user it belongs to
func validateAPIKey(key string) (bool, UserModel) {
	if user, ok := apiKeyCache[key]; ok {
		return true, user
	}
	return false, UserModel{}
}

// addAPIKeyToCache will add a key and user to the cache, if the key already exists it will be overwritten with the
// new data
func addAPIKeyToCache(key string, user UserModel) {
	apiKeyCacheLock.Lock()
	defer apiKeyCacheLock.Unlock()
	apiKeyCache[key] = user
}

// removeAPIKeyFromCache will remove a key from the cache
func removeAPIKeyFromCache(key string) {
	apiKeyCacheLock.Lock()
	defer apiKeyCacheLock.Unlock()
	delete(apiKeyCache, key)
}

// generateAPIKey will generate a new API key for the user and remove the users original key from the cache
func regenerateUsersAPIKey(ctx context.Context, db *ent.Client, userID int) (string, error) {
	newAPIKey := generateAPIKey()
	user, err := getUserByID(ctx, db, userID)
	if err != nil {
		return "", err
	}
	_, err = updateUser(ctx, db, user.ID, UpdateUserInput{
		APIKey: &newAPIKey,
	})
	if err != nil {
		return "", err
	}
	// Remove the old API key from the cache and then add the new one
	removeAPIKeyFromCache(user.APIKey)
	addAPIKeyToCache(newAPIKey, user)
	return newAPIKey, nil
}
