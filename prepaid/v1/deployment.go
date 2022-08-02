package v1

import (
	data "github.com/arangodb-managed/apis/data/v1"
)

//CanDeploymentBeUpgraded returns whenever the deployment spec differs from PrepaidDeployment specs
func (pd *PrepaidDeployment) CanDeploymentBeUpgraded(depl *data.Deployment) bool {
	if pd == nil {
		return false
	}

	if depl == nil {
		return false
	}

	return pd.GetSupportPlanId() != depl.GetSupportPlanId() ||
		pd.GetModel().GetModel() != depl.GetModel().GetModel() ||
		pd.GetModel().GetNodeCount() != depl.GetModel().GetNodeCount() ||
		pd.GetModel().GetNodeDiskSize() != depl.GetModel().GetNodeDiskSize() ||
		pd.GetModel().GetNodeSizeId() != depl.GetModel().GetNodeSizeId() ||
		pd.GetDiskPerformanceId() != depl.GetDiskPerformanceId()
}

// IsAddonAvailable validates if the provided addon is part of the prepaid deployment
func (pd *PrepaidDeployment) IsAddonAvailable(addon string) bool {
	for _, a := range pd.GetAddons() {
		if a == data.AddonIDPrivateEndpointService {
			return true
		}
	}
	return false
}
