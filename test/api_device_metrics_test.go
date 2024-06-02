/*
Torizon OTA

Testing DeviceMetricsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/commontorizon/torizon-openapi-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_DeviceMetricsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceMetricsAPIService GetDeviceDataDevicesDeviceuuidMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceUuid string

		resp, httpRes, err := apiClient.DeviceMetricsAPI.GetDeviceDataDevicesDeviceuuidMetrics(context.Background(), deviceUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceMetricsAPIService GetDeviceDataFleetsFleetidMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fleetId string

		resp, httpRes, err := apiClient.DeviceMetricsAPI.GetDeviceDataFleetsFleetidMetrics(context.Background(), fleetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceMetricsAPIService GetDeviceDataMetricNames", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeviceMetricsAPI.GetDeviceDataMetricNames(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
