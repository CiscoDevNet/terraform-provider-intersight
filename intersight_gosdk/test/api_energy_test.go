/*
Cisco Intersight

Testing EnergyApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package intersight

import (
	"context"
	"testing"

	openapiclient "github.com/CiscoDevNet/intersight-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_intersight_EnergyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnergyApiService CreateEnergyDailyMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnergyApi.CreateEnergyDailyMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnergyApiService CreateEnergyMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnergyApi.CreateEnergyMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
