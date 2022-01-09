package variables

import (
	"testing"

	dashboardv1 "github.com/kobsio/kobs/pkg/api/apis/dashboard/v1"

	"github.com/stretchr/testify/require"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestGetVariables(t *testing.T) {
	t.Run("variables are nil", func(t *testing.T) {
		require.Equal(t, []dashboardv1.Variable{
			{
				Name:  "__cluster",
				Label: "__cluster",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["cluster1"]}`)},
				},
			},
			{
				Name:  "__namespace",
				Label: "__namespace",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["namespace1"]}`)},
				},
			},
			{
				Name:  "placeholder1",
				Label: "placeholder1",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["value1"]}`)},
				},
			},
		}, GetVariables(nil, "cluster1", "namespace1", map[string]string{"placeholder1": "value1"}))
	})

	t.Run("placeholders are nil", func(t *testing.T) {
		require.Equal(t, []dashboardv1.Variable{
			{
				Name:  "__cluster",
				Label: "__cluster",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["cluster1"]}`)},
				},
			},
			{
				Name:  "__namespace",
				Label: "__namespace",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["namespace1"]}`)},
				},
			},
			{
				Name:  "variable1",
				Label: "Variable 1",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["value1"]}`)},
				},
			},
		}, GetVariables([]dashboardv1.Variable{
			{
				Name:  "variable1",
				Label: "Variable 1",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["value1"]}`)},
				},
			},
		}, "cluster1", "namespace1", nil))
	})

	t.Run("cluster and namespace are empty", func(t *testing.T) {
		require.Equal(t, []dashboardv1.Variable{
			{
				Name:  "placeholder1",
				Label: "placeholder1",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["value1"]}`)},
				},
			},
			{
				Name:  "variable1",
				Label: "Variable 1",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["value1"]}`)},
				},
			},
		}, GetVariables([]dashboardv1.Variable{
			{
				Name:  "variable1",
				Label: "Variable 1",
				Hide:  true,
				Plugin: dashboardv1.Plugin{
					Name:    "core",
					Options: &apiextensionsv1.JSON{Raw: []byte(`{"type": "static", "items": ["value1"]}`)},
				},
			},
		}, "", "", map[string]string{"placeholder1": "value1"}))
	})
}