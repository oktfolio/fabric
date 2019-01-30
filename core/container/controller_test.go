/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package container_test

import (
	"testing"

	"github.com/oktfolio/hyperledger-fabric-gm/common/util"
	"github.com/oktfolio/hyperledger-fabric-gm/core/chaincode/platforms"
	"github.com/oktfolio/hyperledger-fabric-gm/core/chaincode/platforms/golang"
	"github.com/oktfolio/hyperledger-fabric-gm/core/container"
	pb "github.com/oktfolio/hyperledger-fabric-gm/protos/peer"
	"github.com/stretchr/testify/assert"
)

func TestVM_GetChaincodePackageBytes(t *testing.T) {
	_, err := container.GetChaincodePackageBytes(nil, nil)
	assert.Error(t, err,
		"GetChaincodePackageBytes did not return error when chaincode spec is nil")
	spec := &pb.ChaincodeSpec{ChaincodeId: nil}
	_, err = container.GetChaincodePackageBytes(nil, spec)
	assert.Error(t, err, "Error expected when GetChaincodePackageBytes is called with nil chaincode ID")
	assert.Contains(t, err.Error(), "invalid chaincode spec")
	spec = &pb.ChaincodeSpec{Type: pb.ChaincodeSpec_GOLANG,
		ChaincodeId: nil,
		Input:       &pb.ChaincodeInput{Args: util.ToChaincodeArgs("f")}}
	_, err = container.GetChaincodePackageBytes(platforms.NewRegistry(&golang.Platform{}), spec)
	assert.Error(t, err,
		"GetChaincodePackageBytes did not return error when chaincode ID is nil")
}
