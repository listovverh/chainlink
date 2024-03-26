import { ethers } from 'hardhat'
import { assert } from 'chai'
import { KeeperRegistry1_3__factory as KeeperRegistryFactory } from '../../../typechain/factories/KeeperRegistry1_3__factory'

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** REGISTRY v1.3 IS FROZEN ************************************/

// All tests are disabled for this contract, as we expect it to never change in the future.
// Instead, we test that the bytecode for the contract has not changed.
// If this test ever fails, you should remove it and then re-run the original test suite.

const BYTECODE = KeeperRegistryFactory.bytecode
const BYTECODE_CHECKSUM =
  '0x7e831ebc4e043fc2946449e11f0d170ba5b6085b213591973c437bc5109b1582'

describe('KeeperRegistry1_3 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(ethers.utils.id(BYTECODE), BYTECODE_CHECKSUM)
  })
})
