package notification

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openmeterio/openmeter/api"
	"github.com/openmeterio/openmeter/internal/notification"
	"github.com/openmeterio/openmeter/internal/productcatalog"
	"github.com/openmeterio/openmeter/pkg/convert"
	"github.com/openmeterio/openmeter/pkg/errorsx"
	"github.com/openmeterio/openmeter/pkg/models"
)

func NewBalanceThresholdPayload() notification.EventPayload {
	return notification.EventPayload{
		EventPayloadMeta: notification.EventPayloadMeta{
			Type: notification.EventTypeBalanceThreshold,
		},
		BalanceThreshold: notification.BalanceThresholdPayload{
			Entitlement: api.EntitlementMetered{
				CreatedAt: convert.ToPointer(time.Now().Add(-10 * 24 * time.Hour).UTC()),
				CurrentUsagePeriod: api.Period{
					From: time.Now().Add(-24 * time.Hour).UTC(),
					To:   time.Now().UTC(),
				},
				DeletedAt:               nil,
				FeatureId:               "01J4VCZKH5QAF85GE501M8637W",
				FeatureKey:              "feature-1",
				Id:                      convert.ToPointer("01J4VCTKG06VJ0H78GD0MZBE49"),
				IsSoftLimit:             nil,
				IsUnlimited:             nil,
				IssueAfterReset:         nil,
				IssueAfterResetPriority: nil,
				LastReset:               time.Time{},
				MeasureUsageFrom:        time.Time{},
				Metadata:                nil,
				SubjectKey:              "customer-1",
				Type:                    "",
				UpdatedAt:               convert.ToPointer(time.Now().Add(-2 * time.Hour).UTC()),
				UsagePeriod: api.RecurringPeriod{
					Anchor:   time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC),
					Interval: "MONTH",
				},
			},
			Feature: api.Feature{
				ArchivedAt:          nil,
				CreatedAt:           convert.ToPointer(time.Now().Add(-10 * 24 * time.Hour).UTC()),
				DeletedAt:           nil,
				Id:                  convert.ToPointer("01J4VCZKH5QAF85GE501M8637W"),
				Key:                 "feature-1",
				Metadata:            nil,
				MeterGroupByFilters: nil,
				MeterSlug:           nil,
				Name:                "feature-1",
				UpdatedAt:           convert.ToPointer(time.Now().Add(-24 * time.Hour).UTC()),
			},
			Subject: api.Subject{
				CurrentPeriodEnd:   &time.Time{},
				CurrentPeriodStart: &time.Time{},
				DisplayName:        nil,
				Id:                 convert.ToPointer("01J4VD1XZH5HM705DCPB8XD5QD"),
				Key:                "customer-1",
				Metadata:           nil,
				StripeCustomerId:   nil,
			},
			Balance: api.EntitlementValue{
				Balance:   convert.ToPointer(10000.0),
				Config:    nil,
				HasAccess: convert.ToPointer(true),
				Overage:   convert.ToPointer(500.0),
				Usage:     convert.ToPointer(50000.0),
			},
			Threshold: api.NotificationRuleBalanceThresholdValue{
				Type:  notification.BalanceThresholdTypePercent,
				Value: 50,
			},
		},
	}
}

func NewCreateEventInput(t notification.EventType, ruleID string, payload notification.EventPayload) notification.CreateEventInput {
	return notification.CreateEventInput{
		NamespacedModel: models.NamespacedModel{
			Namespace: TestNamespace,
		},
		Type:    t,
		Payload: payload,
		RuleID:  ruleID,
	}
}

type EventTestSuite struct {
	Env TestEnv

	subjectKey  models.SubjectKey
	entitlement interface{}

	channel notification.Channel
	rule    notification.Rule
	feature productcatalog.Feature
}

func (s *EventTestSuite) Setup(ctx context.Context, t *testing.T) {
	t.Helper()

	meter, err := s.Env.Meter().GetMeterByIDOrSlug(ctx, TestNamespace, TestMeterSlug)
	require.NoError(t, err, "Getting meter must not return error")

	feature, err := s.Env.Feature().GetFeature(ctx, TestNamespace, TestFeatureKey, false)
	if _, ok := errorsx.ErrorAs[*productcatalog.FeatureNotFoundError](err); !ok {
		require.NoError(t, err, "Getting feature must not return error")
	}
	if feature != nil {
		s.feature = *feature
	} else {
		s.feature, err = s.Env.Feature().CreateFeature(ctx, productcatalog.CreateFeatureInputs{
			Name:                TestFeatureName,
			Key:                 TestFeatureKey,
			Namespace:           TestNamespace,
			MeterSlug:           convert.ToPointer(meter.Slug),
			MeterGroupByFilters: meter.GroupBy,
		})
	}
	require.NoError(t, err, "Creating feature must not return error")

	s.subjectKey = TestSubjectKey

	service := s.Env.Notification()

	channelIn := NewCreateChannelInput("NotificationEventTest")

	channel, err := service.CreateChannel(ctx, channelIn)
	require.NoError(t, err, "Creating channel must not return error")
	require.NotNil(t, channel, "Channel must not be nil")

	s.channel = *channel

	ruleIn := NewCreateRuleInput("NotificationEvent", s.channel.ID)

	rule, err := service.CreateRule(ctx, ruleIn)
	require.NoError(t, err, "Creating rule must not return error")
	require.NotNil(t, rule, "Rule must not be nil")

	s.rule = *rule
}

func (s *EventTestSuite) TestCreateEvent(ctx context.Context, t *testing.T) {
	service := s.Env.Notification()

	input := NewCreateEventInput(notification.EventTypeBalanceThreshold, s.rule.ID, NewBalanceThresholdPayload())

	event, err := service.CreateEvent(ctx, input)
	require.NoError(t, err, "Creating rule must not return error")
	require.NotNil(t, event, "Rule must not be nil")

	assert.Equal(t, float64(50), event.Payload.BalanceThreshold.Threshold.Value)
}

func (s *EventTestSuite) TestGetEvent(ctx context.Context, t *testing.T) {
	service := s.Env.Notification()

	input := NewCreateEventInput(notification.EventTypeBalanceThreshold, s.rule.ID, NewBalanceThresholdPayload())

	event, err := service.CreateEvent(ctx, input)
	require.NoError(t, err, "Creating rule must not return error")
	require.NotNil(t, event, "Rule must not be nil")

	event2, err := service.GetEvent(ctx, notification.GetEventInput{
		NamespacedID: models.NamespacedID{
			Namespace: event.Namespace,
			ID:        event.ID,
		},
	})
	require.NoError(t, err, "Creating rule must not return error")
	require.NotNil(t, event2, "Rule must not be nil")

}

func (s *EventTestSuite) TestListEvents(ctx context.Context, t *testing.T) {
	service := s.Env.Notification()

	createIn := NewCreateEventInput(notification.EventTypeBalanceThreshold, s.rule.ID, NewBalanceThresholdPayload())

	event, err := service.CreateEvent(ctx, createIn)
	require.NoError(t, err, "Creating notification event must not return error")
	require.NotNil(t, event, "Notification event must not be nil")

	listIn := notification.ListEventsInput{
		Namespaces: []string{
			event.Namespace,
		},
		Events: []string{event.ID},
		From:   event.CreatedAt.Add(-time.Minute),
		To:     event.CreatedAt.Add(time.Minute),
	}

	events, err := service.ListEvents(ctx, listIn)
	require.NoError(t, err, "Listing notification events must not return error")
	require.NotNil(t, event, "Notification events must not be nil")

	expectedList := []notification.Event{
		*event,
	}

	assert.EqualValues(t, expectedList, events.Items, "Unexpected items returned by listing events")
}

func (s *EventTestSuite) TestListDeliveryStatus(ctx context.Context, t *testing.T) {
	service := s.Env.Notification()

	createIn := NewCreateEventInput(notification.EventTypeBalanceThreshold, s.rule.ID, NewBalanceThresholdPayload())

	event, err := service.CreateEvent(ctx, createIn)
	require.NoError(t, err, "Creating notification event must not return error")
	require.NotNil(t, event, "Notification event must not be nil")

	listIn := notification.ListEventsDeliveryStatusInput{
		Namespaces: []string{
			event.Namespace,
		},
		Events: []string{event.ID},
	}

	statuses, err := service.ListEventsDeliveryStatus(ctx, listIn)
	require.NoError(t, err, "Listing notification event delivery statuses must not return error")
	require.NotNil(t, event, "Notification event delivery statuses must not be nil")

	assert.Equal(t, statuses.TotalCount, len(s.rule.Channels), "Unexpected number of delivery statuses returned by listing events")

	channelsIDs := func() []string {
		var channelIDs []string
		for _, channel := range s.rule.Channels {
			channelIDs = append(channelIDs, channel.ID)
		}

		return channelIDs
	}()

	for _, status := range statuses.Items {
		assert.Equal(t, status.EventID, event.ID, "Unexpected event ID returned by listing events")
		assert.Truef(t, slices.Contains(channelsIDs, status.ChannelID), "Unexpected channel ID")
	}
}

func (s *EventTestSuite) TestUpdateDeliveryStatus(ctx context.Context, t *testing.T) {
	service := s.Env.Notification()

	createIn := NewCreateEventInput(notification.EventTypeBalanceThreshold, s.rule.ID, NewBalanceThresholdPayload())

	event, err := service.CreateEvent(ctx, createIn)
	require.NoError(t, err, "Creating notification event must not return error")
	require.NotNil(t, event, "Notification event must not be nil")

	subTests := []struct {
		Name  string
		Input notification.UpdateEventDeliveryStatusInput
	}{
		{
			Name: "WithID",
			Input: notification.UpdateEventDeliveryStatusInput{
				NamespacedModel: models.NamespacedModel{
					Namespace: event.Namespace,
				},
				ID:    event.DeliveryStatus[0].ID,
				State: notification.EventDeliveryStatusStateFailed,
			},
		},
		{
			Name: "WithEventIDAndChannelID",
			Input: notification.UpdateEventDeliveryStatusInput{
				NamespacedModel: models.NamespacedModel{
					Namespace: event.Namespace,
				},
				EventID:   event.ID,
				ChannelID: event.Rule.Channels[0].ID,
				State:     notification.EventDeliveryStatusStateSuccess,
			},
		},
	}

	for _, test := range subTests {
		t.Run(test.Name, func(t *testing.T) {
			status, err := service.UpdateEventDeliveryStatus(ctx, test.Input)
			require.NoError(t, err, "Updating notification event delivery status must not return error")
			require.NotNil(t, status, "Notification event must not be nil")
		})
	}
}
