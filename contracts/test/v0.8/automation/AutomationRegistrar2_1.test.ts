import { ethers } from 'hardhat'
import { assert } from 'chai'
import { AutomationRegistrar2_1__factory as AutomationRegistrarFactory } from '../../../typechain/factories/AutomationRegistrar2_1__factory'

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** REGISTRAR v2.1 IS FROZEN ************************************/

// We are leaving the original tests enabled, however as 2.1 is still actively being deployed

describe('AutomationRegistrar2_1 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(
      ethers.utils.id(AutomationRegistrarFactory.bytecode),
      '0x9633058bd81e8479f88baaee9bda533406295c80ccbc43d4509701001bbea6e3',
      'KeeperRegistry bytecode has changed',
    )
  })
})
