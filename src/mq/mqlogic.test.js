import chai from 'chai';
import MQLogic from './mqlogic.js';
const expect = chai.expect;

describe('MQ Retry Logic Unit Test', function () {
  
  it('should perform send properly', function(done){
    let logic = new MQLogic({id:'logic1', timeout:1000, totalTimeout:3000});
    let doneNow = false;
    logic.on('PerformSend', function (params) {
      expect(params.payload).to.equal('testPayload-1');
      expect(params.dest).to.equal('testDest-1');
      //expect(params.msgId).to.equal(1,' msg 1');
      
      if ( doneNow == false ) {
        doneNow = true;  
        expect(params.seqId).to.equal(1, 'seq 1');
        logic._cleanUp(params.msgId);
        done();
      }
    });
    logic.Send('testDest-1','testPayload-1');
  });

  it('should do clean up properly', function(done) {
    let logic = new MQLogic({timeout:1000, totalTimeout:3000});
    let msgId;
    logic.on('PerformSend', function (params) {
      expect(params.payload).to.equal('testPayload-2');
      expect(params.dest).to.equal('testDest-2');
      //expect(params.msgId).to.equal(1, 'first msg')
      //expect(params.seqId).to.equal(1, 'first sequence');
      msgId = params.msgId;
    });
    logic.on('PerformCleanUp', function(seqId){
      //expect(seqId).to.equal(1, 'check cleanup');
      done();
    });
    logic.Send('testDest-2','testPayload-2');
    logic.AckReceived(msgId);
  });

  it('should handle retry command properly', function(done){
    let count = 0;
    let logic2 = new MQLogic({timeout:200, totalTimeout:500});
    logic2.on('PerformSend', function (params) {
      count++;
      //expect(params.msgId).to.equal(1, 'msg id 1');
      expect(params.seqId).to.equal(count, 'seq id should increase');
    });

    logic2.on('PerformTotalTimeout', function(params){
      //expect(params.msgId).to.equal(1, 'timeout for id 1');
      done();
    });
    logic2.Send('testDest-3','testPayload-3');
    
  });
});

