/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	configtxtest "github.com/oktfolio/hyperledger-fabric-gm/common/configtx/test"
	mockchannelconfig "github.com/oktfolio/hyperledger-fabric-gm/common/mocks/config"
	mockconfigtx "github.com/oktfolio/hyperledger-fabric-gm/common/mocks/configtx"
	mockpolicies "github.com/oktfolio/hyperledger-fabric-gm/common/mocks/policies"
	"github.com/oktfolio/hyperledger-fabric-gm/core/ledger"
	"github.com/oktfolio/hyperledger-fabric-gm/core/ledger/ledgermgmt"
)

//MockInitialize resets chains for test env
func MockInitialize() {
	ledgermgmt.InitializeTestEnvWithInitializer(
		&ledgermgmt.Initializer{
			CustomTxProcessors: ConfigTxProcessors,
		},
	)
	chains.list = make(map[string]*chain)
	chainInitializer = func(string) { return }
}

// MockCreateChain used for creating a ledger for a chain for tests
// without having to join
func MockCreateChain(cid string) error {
	var ledger ledger.PeerLedger
	var err error

	if ledger = GetLedger(cid); ledger == nil {
		gb, _ := configtxtest.MakeGenesisBlock(cid)
		if ledger, err = ledgermgmt.CreateLedger(gb); err != nil {
			return err
		}
	}

	chains.Lock()
	defer chains.Unlock()

	chains.list[cid] = &chain{
		cs: &chainSupport{
			Resources: &mockchannelconfig.Resources{
				PolicyManagerVal: &mockpolicies.Manager{
					Policy: &mockpolicies.Policy{},
				},
				ConfigtxValidatorVal: &mockconfigtx.Validator{},
				ApplicationConfigVal: &mockchannelconfig.MockApplication{CapabilitiesRv: &mockchannelconfig.MockApplicationCapabilities{}},
			},

			ledger: ledger,
		},
	}

	return nil
}
