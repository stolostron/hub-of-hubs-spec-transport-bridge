package controller

import (
	"fmt"
	"time"

	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/controller/dbsyncer"
	statuswatcher "github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/controller/status-watcher"
	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/db"
	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/transport"
	ctrl "sigs.k8s.io/controller-runtime"
)

// AddDBToTransportSyncers adds the controllers that send info from DB to transport layer to the Manager.
func AddDBToTransportSyncers(mgr ctrl.Manager, specDB db.SpecDB, transportObj transport.Transport,
	syncInterval time.Duration) error {
	addDBSyncerFunctions := []func(ctrl.Manager, db.SpecDB, transport.Transport, time.Duration) error{
		dbsyncer.AddHoHConfigDBToTransportSyncer,
		dbsyncer.AddPoliciesDBToTransportSyncer,
		dbsyncer.AddPlacementRulesDBToTransportSyncer,
		dbsyncer.AddPlacementBindingsDBToTransportSyncer,
		dbsyncer.AddApplicationsDBToTransportSyncer,
		dbsyncer.AddSubscriptionsDBToTransportSyncer,
		dbsyncer.AddChannelsDBToTransportSyncer,
		dbsyncer.AddManagedClusterLabelsDBToTransportSyncer,
		dbsyncer.AddPlacementsDBToTransportSyncer,
		dbsyncer.AddManagedClusterSetsDBToTransportSyncer,
		dbsyncer.AddManagedClusterSetBindingsDBToTransportSyncer,
	}
	for _, addDBSyncerFunction := range addDBSyncerFunctions {
		if err := addDBSyncerFunction(mgr, specDB, transportObj, syncInterval); err != nil {
			return fmt.Errorf("failed to add DB Syncer: %w", err)
		}
	}

	return nil
}

// AddStatusDBWatchers adds the controllers that watch the status DB to update the spec DB to the Manager.
func AddStatusDBWatchers(mgr ctrl.Manager, specDB db.SpecDB, statusDB db.StatusDB) error {
	if err := statuswatcher.AddManagedClusterLabelsStatusWatcher(mgr, specDB, statusDB); err != nil {
		return fmt.Errorf("failed to add status watcher: %w", err)
	}

	return nil
}
