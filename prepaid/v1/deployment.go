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
		pd.GetModel().GetNodeSizeId() != depl.GetModel().GetNodeSizeId()
}
