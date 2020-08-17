package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBetaTester(t *testing.T) {
	email := Email("me@email.com")
	want := &BetaTesterResponse{
		Data: BetaTester{
			Attributes: &BetaTesterAttributes{
				Email: &email,
			},
		},
	}
	testEndpointWithResponse(t, `{"data":{"attributes":{"email":"me@email.com"}}}`, want, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.CreateBetaTester(ctx, BetaTesterCreateRequestAttributes{Email: Email("me@email.com")}, []string{"10"}, []string{"10"})
	})
}

func TestDeleteBetaTester(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.DeleteBetaTester(ctx, "10")
	})
}

func TestListBetaTesters(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTestersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaTesters(ctx, &ListBetaTestersQuery{})
	})
}

func TestGetBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTesterResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.GetBetaTester(ctx, "10", &GetBetaTesterQuery{})
	})
}

func TestGetBetaTesterIncludeds(t *testing.T) {
	testEndpointCustomBehavior(`{"included":[{"type":"apps"},{"type":"betaGroups"},{"type":"builds"}]}`, func(ctx context.Context, client *Client) {
		tester, _, err := client.TestFlight.GetBetaTester(ctx, "10", &GetBetaTesterQuery{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tester.Included)

		assert.NotNil(t, tester.Included[0].App())
		assert.NotNil(t, tester.Included[1].BetaGroup())
		assert.NotNil(t, tester.Included[2].Build())

		assert.Nil(t, tester.Included[0].BetaGroup())
		assert.Nil(t, tester.Included[0].Build())
	})
}

func TestAddBetaTesterToBetaGroups(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.AddBetaTesterToBetaGroups(ctx, "10", []string{"10"})
	})
}

func TestRemoveBetaTesterFromBetaGroups(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.RemoveBetaTesterFromBetaGroups(ctx, "10", []string{"10"})
	})
}

func TestAssignSingleBetaTesterToBuilds(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.AssignSingleBetaTesterToBuilds(ctx, "10", []string{"10"})
	})
}

func TestUnassignSingleBetaTesterFromBuilds(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.UnassignSingleBetaTesterFromBuilds(ctx, "10", []string{"10"})
	})
}

func TestRemoveSingleBetaTesterAccessApps(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.TestFlight.RemoveSingleBetaTesterAccessApps(ctx, "10", []string{"10"})
	})
}

func TestListAppsForBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListAppsForBetaTester(ctx, "10", &ListAppsForBetaTesterQuery{})
	})
}

func TestListAppIDsForBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTesterAppsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListAppIDsForBetaTester(ctx, "10", &ListAppIDsForBetaTesterQuery{})
	})
}

func TestListBuildsIndividuallyAssignedToBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BuildsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildsIndividuallyAssignedToBetaTester(ctx, "10", &ListBuildsIndividuallyAssignedToBetaTesterQuery{})
	})
}

func TestListBuildIDsIndividuallyAssignedToBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTesterBuildsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBuildIDsIndividuallyAssignedToBetaTester(ctx, "10", &ListBuildIDsIndividuallyAssignedToBetaTesterQuery{})
	})
}

func TestListIndividualTestersForBuild(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTestersResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListIndividualTestersForBuild(ctx, "10", &ListIndividualTestersForBuildQuery{})
	})
}

func TestListBetaGroupsForBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaGroupsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaGroupsForBetaTester(ctx, "10", &ListBetaGroupsForBetaTesterQuery{})
	})
}

func TestListBetaGroupIDsForBetaTester(t *testing.T) {
	testEndpointWithResponse(t, "{}", &BetaTesterBetaGroupsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.TestFlight.ListBetaGroupIDsForBetaTester(ctx, "10", &ListBetaGroupIDsForBetaTesterQuery{})
	})
}
