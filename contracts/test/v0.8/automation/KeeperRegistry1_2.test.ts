import { ethers } from 'hardhat'
import { assert } from 'chai'
import { KeeperRegistry1_2__factory as KeeperRegistryFactory } from '../../../typechain/factories/KeeperRegistry1_2__factory'

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** REGISTRY v1.2 IS FROZEN ************************************/

const BYTECODE = KeeperRegistryFactory.bytecode
const BYTECODE_CHECKSUM =
  '0x4a23953416a64a0fa4c943954d9a92059f862257440f2cbcf5f238314b89f416'

describe('KeeperRegistry1_2 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(ethers.utils.id(BYTECODE), BYTECODE_CHECKSUM)
  })
})
