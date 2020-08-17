package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListApps(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListApps(ctx, &ListAppsQuery{})
	})
}

func TestGetApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetApp(ctx, "10", &GetAppQuery{})
	})
}

func TestGetAppIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[
		{"type":"betaGroups"},{"type":"appStoreVersions"},{"type":"preReleaseVersions"},
		{"type":"betaAppLocalizations"},{"type":"builds"},{"type":"betaLicenseAgreements"},
		{"type":"betaAppReviewDetails"},{"type":"appInfos"},{"type":"endUserLicenseAgreements"},
		{"type":"appPreOrders"},{"type":"appPrices"},{"type":"territories"},{"type":"inAppPurchases"},
		{"type":"gameCenterEnabledVersions"},{"type":"perfPowerMetrics"}
		]}`, func(ctx context.Context, client *Client) {
		app, _, err := client.Apps.GetApp(ctx, "10", &GetAppQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, app.Included)

		assert.NotNil(t, app.Included[0].BetaGroup())
		assert.NotNil(t, app.Included[1].AppStoreVersion())
		assert.NotNil(t, app.Included[2].PrereleaseVersion())
		assert.NotNil(t, app.Included[3].BetaAppLocalization())
		assert.NotNil(t, app.Included[4].Build())
		assert.NotNil(t, app.Included[5].BetaLicenseAgreement())
		assert.NotNil(t, app.Included[6].BetaAppReviewDetail())
		assert.NotNil(t, app.Included[7].AppInfo())
		assert.NotNil(t, app.Included[8].EndUserLicenseAgreement())
		assert.NotNil(t, app.Included[9].AppPreOrder())
		assert.NotNil(t, app.Included[10].AppPrice())
		assert.NotNil(t, app.Included[11].Territory())
		assert.NotNil(t, app.Included[12].InAppPurchase())
		assert.NotNil(t, app.Included[13].GameCenterEnabledVersion())
		assert.NotNil(t, app.Included[14].PerfPowerMetric())

		assert.Nil(t, app.Included[0].AppStoreVersion())
		assert.Nil(t, app.Included[0].PrereleaseVersion())
		assert.Nil(t, app.Included[0].BetaAppLocalization())
		assert.Nil(t, app.Included[0].Build())
		assert.Nil(t, app.Included[0].BetaLicenseAgreement())
		assert.Nil(t, app.Included[0].BetaAppReviewDetail())
		assert.Nil(t, app.Included[0].AppInfo())
		assert.Nil(t, app.Included[0].EndUserLicenseAgreement())
		assert.Nil(t, app.Included[0].AppPreOrder())
		assert.Nil(t, app.Included[0].AppPrice())
		assert.Nil(t, app.Included[0].Territory())
		assert.Nil(t, app.Included[0].InAppPurchase())
		assert.Nil(t, app.Included[0].GameCenterEnabledVersion())
		assert.Nil(t, app.Included[0].PerfPowerMetric())
	})
}

func TestUpdateApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.UpdateApp(ctx, "10", &AppUpdateRequestAttributes{}, []string{"10"}, []string{"10"})
	})
}

func TestRemoveBetaTestersFromApp(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.RemoveBetaTestersFromApp(ctx, "10", []string{"10"})
	})
}

func TestListInAppPurchasesForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &InAppPurchasesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListInAppPurchasesForApp(ctx, "10", &ListInAppPurchasesQuery{})
	})
}

func TestGetInAppPurchase(t *testing.T) {
	testEndpointWithResponse(t, "{}", &InAppPurchaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.GetInAppPurchase(ctx, "10", &GetInAppPurchaseQuery{})
	})
}
