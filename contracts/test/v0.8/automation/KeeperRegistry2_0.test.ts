import { ethers } from 'hardhat'
import { assert } from 'chai'
import { KeeperRegistry2_0__factory as KeeperRegistryFactory } from '../../../typechain/factories/KeeperRegistry2_0__factory'

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** REGISTRY v2.0 IS FROZEN ************************************/

// All tests are disabled for this contract, as we expect it to never change in the future.
// Instead, we test that the bytecode for the contract has not changed.
// If this test ever fails, you should remove it and then re-run the original test suite.

const BYTECODE = KeeperRegistryFactory.bytecode
const BYTECODE_CHECKSUM =
  '0x60660453a335cdcd42b5aa64e58a8c04517e8a8645d2618b51a7552df6e2973b'

describe('KeeperRegistry2_0 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(ethers.utils.id(BYTECODE), BYTECODE_CHECKSUM)
  })
})
