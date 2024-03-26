import { ethers } from 'hardhat'
import { assert } from 'chai'
import { KeeperRegistry1_3__factory as KeeperRegistryFactory } from '../../../typechain/factories/KeeperRegistry1_3__factory'

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** REGISTRY v1.3 IS FROZEN ************************************/

const BYTECODE = KeeperRegistryFactory.bytecode
const BYTECODE_CHECKSUM =
  '0x7e831ebc4e043fc2946449e11f0d170ba5b6085b213591973c437bc5109b1582'

describe('KeeperRegistry1_3 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(ethers.utils.id(BYTECODE), BYTECODE_CHECKSUM)
  })
})
