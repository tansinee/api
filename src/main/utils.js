import fetch from 'node-fetch';
import forge from 'node-forge';

let nonce = Date.now() % 10000;
const logicUrl =
  process.env.TENDERMINT_ADDRESS ||
  'http://localhost:' + (process.env.ROLE === 'rp' ? '45001' : '45000');

function retrieveResult(obj, isQuery) {
  if (obj.error) {
    console.error(obj.error);
    return obj.error;
  }
  if (isQuery) {
    let result = Buffer.from(obj.result.response.value, 'base64').toString();
    return JSON.parse(result);
  } else if (obj.result.deliver_tx.log === 'success') return true;
  else {
    console.error('Update chain failed:', obj);
    return false;
  }
}

export async function hash(stringToHash) {
  // TODO implement secure hashing
  // hash with SHA256 and encode hashed string with hex
  var md = forge.md.sha256.create();
  md.update(stringToHash);
  var hashedString = md.digest().toHex();
  return hashedString.toUpperCase();
}

export async function decryptAsymetricKey(key, message) {
  // TODO implement decryption
  // Assume that key format is ASN.1 encoded with base64 string
  if(key == null) return message.slice(message.indexOf('(') + 1, message.length - 1);
  // decode key string
  const derString = forge.util.decode64(key);
  // convert to ASN.1 format
  const asn = forge.asn1.fromDer(derString);
  const privateKey = forge.pki.privateKeyFromAsn1(asn);
  return privateKey.decrypt(message);
}

export async function encryptAsymetricKey(key, message) {
  // TODO implement encryption
  // encrypt with RSA public key, assume that key format is ASN.1 encoded with base64 string
  // note that this can only encrypt small size of data, if you want to encrypt file, please consider using gpg encryption
  if(key == null) return 'Encrypt_with_' + key + '(' + message + ')'; // dummy value
  // decode key string
  const derString = forge.util.decode64(key);
  // convert to ASN.1 format
  const asn = forge.asn1.fromDer(derString);
  const publicKey = forge.pki.publicKeyFromAsn1(asn);
  return forge.util.encode64(publicKey.encrypt(message));
}

export function generateIdentityProof(data) {
  // TODO
  return '<some-voodoo-happen-here>';
}

export async function createRequestId(privkey, data, nonce) {
  // TODO implement real request_id generating algorithm
  return await hash(
    'Concat_with_nonce_' +
      nonce +
      '(' +
      Buffer.from(JSON.stringify(data)).toString('base64') +
      ')'
  );
}

export function getNonce() {
  // TODO
  return (nonce++).toString();
}

export async function queryChain(fnName, data) {
  let encoded = Buffer.from(fnName + '|' + JSON.stringify(data)).toString(
    'base64'
  );

  let result = await fetch(logicUrl + '/abci_query?data="' + encoded + '"');
  return retrieveResult(await result.json(), true);
}

export async function updateChain(fnName, data, nonce) {
  let encoded = Buffer.from(
    fnName + '|' + JSON.stringify(data) + '|' + nonce
  ).toString('base64');

  let result = await fetch(
    logicUrl + '/broadcast_tx_commit?tx="' + encoded + '"'
  );
  return retrieveResult(await result.json());
}
