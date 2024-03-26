import { ethers } from 'hardhat'
import { assert } from 'chai'
import { UpkeepTranscoder4_0__factory as UpkeepTranscoderFactory } from '../../../typechain/factories/UpkeepTranscoder4_0__factory'

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

/*********************************** TRANSCODER v4.0 IS FROZEN ************************************/

describe('UpkeepTranscoder v4.0 - Frozen [ @skip-coverage ]', () => {
  it('has not changed', () => {
    assert.equal(
      ethers.utils.id(UpkeepTranscoderFactory.bytecode),
      '0xf22c4701b0088e6e69c389a34a22041a69f00890a89246e3c2a6d38172222dae',
      'UpkeepTranscoder bytecode has changed',
    )
  })
})
