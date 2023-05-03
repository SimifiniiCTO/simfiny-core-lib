package feeds

import (
	"github.com/GetStream/stream-go2/v7"
	"github.com/SimifiniiCTO/simfiny-social-service/pkg/database/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func generateRandomUserIds(numUserIdsToGenerate int) []string {
	userIds := make([]string, 0, numUserIdsToGenerate)
	for i := 0; i < numUserIdsToGenerate; i++ {
		userIds = append(userIds, utils.GenerateRandomString(10))
	}

	return userIds
}

// generateManyFlatFeeds  generates a set of flat feeds for target feed
func generateManyFlatFeeds(t *testing.T, userIds []string) []stream.Feed {
	flatFeeds := make([]stream.Feed, 0, len(userIds))
	for _, id := range userIds {
		// create flat feed
		singularFlatFeed, err := newFlatFeedWithUserID(testClient.client, id)
		require.NoError(t, err)

		flatFeeds = append(flatFeeds, singularFlatFeed)
	}

	return flatFeeds
}

// generateManyAggregatedFeeds generates a set of aggregated feeds for target feed
func generateManyAggregatedFeeds(t *testing.T, userIds []string) []*stream.AggregatedFeed {
	aggregatedFeeds := make([]*stream.AggregatedFeed, 0, len(userIds))
	for _, id := range userIds {
		// create aggregated flat feed
		singularAggregatedFeed, err := newAggregatedFeedWithUserID(testClient.client, id)
		require.NoError(t, err)

		aggregatedFeeds = append(aggregatedFeeds, singularAggregatedFeed)
	}

	return aggregatedFeeds
}

// generates a random flat feed used for testing
func generateFlatFeed(t *testing.T, userId string) *stream.FlatFeed {
	flatfeed, err := newFlatFeedWithUserID(testClient.client, userId)
	require.NoError(t, err)

	return flatfeed
}

// generateNotificationFeed generates a notification flat feed used for testing
func generateNotificationFeed(t *testing.T, userId string) *stream.NotificationFeed {
	notificationFeed, err := newNotificationFeedWithUserID(testClient.client, userId)
	require.NoError(t, err)

	return notificationFeed
}

// generateRandomActivities generates a random set of activities
func generateRandomActivities(numActivities int) []*stream.Activity {
	activitySet := make([]*stream.Activity, 0, numActivities)
	for i := 0; i < numActivities; i++ {
		activitySet = append(activitySet, generateRandomActivity())
	}

	return activitySet
}

// generates a random activity
func generateRandomActivity() *stream.Activity {
	return &stream.Activity{
		ID:        utils.GenerateRandomString(20),
		Actor:     utils.GenerateRandomString(20),
		Verb:      utils.GenerateRandomString(20),
		Object:    utils.GenerateRandomString(20),
		ForeignID: utils.GenerateRandomString(20),
		Target:    utils.GenerateRandomString(20),
		Time: stream.Time{
			Time: time.Now(),
		},
		Origin: utils.GenerateRandomString(20),
		To: []string{
			utils.GenerateRandomString(10),
			utils.GenerateRandomString(10),
			utils.GenerateRandomString(10),
		},
		Score: float64(utils.GenerateRandomId(10, 100)),
		Extra: map[string]interface{}{
			"test-detail-1": utils.GenerateRandomString(10),
			"test-detail-2": utils.GenerateRandomString(10),
			"test-detail-3": utils.GenerateRandomString(10),
			"test-detail-4": utils.GenerateRandomString(10),
			"test-detail-5": utils.GenerateRandomString(10),
			"test-detail-6": utils.GenerateRandomString(10),
		},
	}
}
