import { ethers } from 'hardhat'
import { assert } from 'chai'
import { KeeperRegistry2_1__factory as KeeperRegistryFactory } from '../../../typechain/factories/KeeperRegistry2_1__factory'
import { KeeperRegistryLogicA2_1__factory as KeeperRegistryLogicAFactory } from '../../../typechain/factories/KeeperRegistryLogicA2_1__factory'
import { KeeperRegistryLogicB2_1__factory as KeeperRegistryLogicBFactory } from '../../../typechain/factories/KeeperRegistryLogicB2_1__factory'
import { AutomationForwarderLogic__factory as AutomationForwarderLogicFactory } from '../../../typechain/factories/AutomationForwarderLogic__factory'
//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** REGISTRY v2.1 IS FROZEN ************************************/

// We are leaving the original tests enabled, however as 2.1 is still actively being deployed

describe('KeeperRegistry2_1 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(
      ethers.utils.id(KeeperRegistryFactory.bytecode),
      '0x05aaa1024d7400e9c4824dde093b96edf5888fa6e6be2c2fc4dca7ae47cc9de9',
      'KeeperRegistry bytecode has changed',
    )
    assert.equal(
      ethers.utils.id(KeeperRegistryLogicAFactory.bytecode),
      '0xdcc8805e88c550b2a25b972bee9f4e4c3649f01e26f8dda6b25d7a9c5da8ab2f',
      'KeeperRegistryLogicA bytecode has changed',
    )
    assert.equal(
      ethers.utils.id(KeeperRegistryLogicBFactory.bytecode),
      '0x891c26ba35b9b13afc9400fac5471d15842828ab717cbdc70ee263210c542563',
      'KeeperRegistryLogicB bytecode has changed',
    )
    assert.equal(
      ethers.utils.id(AutomationForwarderLogicFactory.bytecode),
      '0x6b89065111e9236407329fae3d68b33c311b7d3b6c2ae3dd15c1691a28b1aca7',
      'AutomationForwarderLogic bytecode has changed',
    )
  })
})
