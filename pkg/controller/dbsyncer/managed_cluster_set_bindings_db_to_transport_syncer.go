package dbsyncer

import (
	"context"
	"fmt"
	"time"

	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/bundle"
	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/db"
	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/intervalpolicy"
	"github.com/stolostron/hub-of-hubs-spec-transport-bridge/pkg/transport"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1beta1 "open-cluster-management.io/api/cluster/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	managedClusterSetBindingsTableName = "managedclustersetbindings"
	managedClusterSetBindingsMsgKey    = "ManagedClusterSetBindings"
)

// AddManagedClusterSetBindingsDBToTransportSyncer adds managed-cluster-set-bindings db to transport syncer to the
// manager.
func AddManagedClusterSetBindingsDBToTransportSyncer(mgr ctrl.Manager, specDB db.SpecDB,
	transportObj transport.Transport, syncInterval time.Duration) error {
	createObjFunc := func() metav1.Object { return &clusterv1beta1.ManagedClusterSetBinding{} }
	lastSyncTimestampPtr := &time.Time{}

	if err := mgr.Add(&genericDBToTransportSyncer{
		log:            ctrl.Log.WithName("managed-cluster-set-bindings-db-to-transport-syncer"),
		intervalPolicy: intervalpolicy.NewExponentialBackoffPolicy(syncInterval),
		syncBundleFunc: func(ctx context.Context) (bool, error) {
			return syncObjectsBundle(ctx, transportObj, managedClusterSetBindingsMsgKey, specDB,
				managedClusterSetBindingsTableName, createObjFunc, bundle.NewBaseObjectsBundle, lastSyncTimestampPtr)
		},
	}); err != nil {
		return fmt.Errorf("failed to add managed-cluster-set-bindings db to transport syncer - %w", err)
	}

	return nil
}
