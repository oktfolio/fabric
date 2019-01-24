/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package server

import (
	"github.com/oktfolio/hyperledger-fabric-gm/common/crypto"
	"github.com/oktfolio/hyperledger-fabric-gm/common/flogging"
	"github.com/oktfolio/hyperledger-fabric-gm/common/ledger/blockledger"
	"github.com/oktfolio/hyperledger-fabric-gm/core/comm"
	"github.com/oktfolio/hyperledger-fabric-gm/orderer/common/cluster"
	"github.com/oktfolio/hyperledger-fabric-gm/orderer/common/localconfig"
	"github.com/oktfolio/hyperledger-fabric-gm/orderer/consensus/etcdraft"
	"github.com/oktfolio/hyperledger-fabric-gm/protos/common"
	"github.com/oktfolio/hyperledger-fabric-gm/protos/utils"
)

type replicationInitiator struct {
	logger         *flogging.FabricLogger
	secOpts        *comm.SecureOptions
	conf           *localconfig.TopLevel
	bootstrapBlock *common.Block
	lf             cluster.LedgerFactory
	signer         crypto.LocalSigner
}

func (ri *replicationInitiator) replicateIfNeeded() {
	if ri.bootstrapBlock.Header.Number == 0 {
		ri.logger.Debug("Booted with a genesis block, replication isn't an option")
		return
	}

	consenterCert := etcdraft.ConsenterCertificate(ri.secOpts.Certificate)
	systemChannelName, err := utils.GetChainIDFromBlock(ri.bootstrapBlock)
	if err != nil {
		ri.logger.Panicf("Failed extracting system channel name from bootstrap block: %v", err)
	}
	pullerConfig := cluster.PullerConfigFromTopLevelConfig(systemChannelName, ri.conf, ri.secOpts.Key, ri.secOpts.Certificate, ri.signer)
	puller, err := cluster.BlockPullerFromConfigBlock(pullerConfig, ri.bootstrapBlock)
	if err != nil {
		ri.logger.Panicf("Failed creating puller config from bootstrap block: %v", err)
	}

	pullerLogger := flogging.MustGetLogger("orderer.common.cluster")

	replicator := &cluster.Replicator{
		LedgerFactory:    ri.lf,
		SystemChannel:    systemChannelName,
		BootBlock:        ri.bootstrapBlock,
		Logger:           pullerLogger,
		AmIPartOfChannel: consenterCert.IsConsenterOfChannel,
		Puller:           puller,
		ChannelLister: &cluster.ChainInspector{
			Logger:          pullerLogger,
			Puller:          puller,
			LastConfigBlock: ri.bootstrapBlock,
		},
	}

	replicationNeeded, err := replicator.IsReplicationNeeded()
	if err != nil {
		ri.logger.Panicf("Failed determining whether replication is needed: %v", err)
	}

	if !replicationNeeded {
		ri.logger.Info("Replication isn't needed")
		return
	}

	ri.logger.Info("Will now replicate chains")
	replicator.ReplicateChains()
}

type ledgerFactory struct {
	blockledger.Factory
}

func (lf *ledgerFactory) GetOrCreate(chainID string) (cluster.LedgerWriter, error) {
	return lf.Factory.GetOrCreate(chainID)
}
