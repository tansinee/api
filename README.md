[![CircleCI](https://circleci.com/gh/ndidplatform/api.svg?style=svg)](https://circleci.com/gh/ndidplatform/api)

# NDID API (Node.js)

## Prerequisites

- Node.js 8.9 or later
- npm 5.6.0 or later

## Getting started

1.  Install dependencies

    ```sh
    npm install
    ```

2.  Run smart contract (tendermint ABCI app) server in `smart-contract` repository and wait for first commit to show up in an output.

3.  Add development keys to the system (for development mode only)

    ```sh
    TENDERMINT_IP=$IP \
    TENDERMINT_PORT=$PORT \
    NODE_ID=ndid1 \
    npm run initDevKey
    ```

4.  Run a server

    ```sh
    ROLE=$ROLE \
    NODE_ID=$NODE_ID \
    npm start
    ```

**Environment variable options**

| Environment variable                          | Allowed Values                                       | Default                                                                    | Required                                                                       | Description                                                                                                                                                       |
| --------------------------------------------- | ---------------------------------------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ROLE                                          | `idp`, `rp`, `as`, `ndid`                            | -                                                                          | Yes                                                                            |                                                                                                                                                                   |
| TENDERMINT_IP                                 | IP address                                           | `localhost`                                                                | No                                                                             | IP Address to contact tendermint RPC.                                                                                                                             |
| TENDERMINT_PORT                               | Port number                                          | `45000` for IDP, `45001` for RP, and `45002` for AS                        | No                                                                             | Port to contact tendermint RPC.                                                                                                                                   |
| MQ_CONTACT_IP                                 | IP address                                           | -                                                                          | Yes                                                                            | An IP address where node's message queue can be contacted.                                                                                                        |
| MQ_BINDING_PORT                               | Port number                                          | `5555`                                                                     | No                                                                             | A port to bind message queue.                                                                                                                                     |
| SERVER_PORT                                   | Port number                                          | `8080`                                                                     | No                                                                             | API server port                                                                                                                                                   |
| PRIVATE_KEY_PATH                              | File path                                            | Development: pre-generated key path; Production: -                         | Development: No; Production: Yes when `USE_EXTERNAL_CRYPTO_SERVICE` is `false` | Path to node's private key.                                                                                                                                       |
| PRIVATE_KEY_PASSPHRASE                        | Any                                                  | -                                                                          | No                                                                             | Passphrase for node's private key.                                                                                                                                |
| MASTER_PRIVATE_KEY_PATH                       | File path                                            | Development: pre-generated key path; Production: -                         | Development: No; Production: Yes when `USE_EXTERNAL_CRYPTO_SERVICE` is `false` | Path to node's master private key.                                                                                                                                |
| MASTER_PRIVATE_KEY_PASSPHRASE                 | Any                                                  | -                                                                          | No                                                                             | Passphrase for node's master private key.                                                                                                                         |
| USE_EXTERNAL_CRYPTO_SERVICE                   | `true`, `false`                                      | `false`                                                                    | No                                                                             | Use external service for decrypting and signing with node's key (e.g. HSM).                                                                                       |
| NODE_ID                                       | Any                                                  | -                                                                          | Yes                                                                            | Node ID. This ID is tied to public key. In development, we have `rp1`, `rp2`, `rp3`, `idp1`, `idp2`, `idp3`, `as1`, `as2`, and `as3`.                             |
| DATA_DIRECTORY_PATH                           | Directory path                                       | `__dirname/../data` (`data` directory in repository's directory)           | No                                                                             | Directory path for persistence data files.                                                                                                                        |
| LOG_LEVEL                                     | `error`, `warn`, `info`, `verbose`, `debug`, `silly` | Development: `debug`; Production: `info`                                   | No                                                                             | Log level                                                                                                                                                         |
| LOG_TARGET                                    | `console`, `file`                                    | Development: `console`; Production: `file`                                 | No                                                                             | Where should logger writes logs to.                                                                                                                               |
| LOG_COLOR                                     | `true`, `false`                                      | `true` when `LOG_TARGET` is `console`, `false` when `LOG_TARGET` is `file` | No                                                                             | Log highlight color                                                                                                                                               |
| LOG_DIRECTORY_PATH                            | Directory path                                       | `__dirname/../log` (`log` directory in repository's directory)             | No                                                                             | Directory path for log files.                                                                                                                                     |
| CLIENT_HTTP_ERROR_CODE                        | Number                                               | `400`                                                                      | No                                                                             | HTTP error code when responding a client error.                                                                                                                   |
| SERVER_HTTP_ERROR_CODE                        | Number                                               | `500`                                                                      | No                                                                             | HTTP error code when responding a server error.                                                                                                                   |
| HTTPS                                         | `true`, `false`                                      | `false`                                                                    | No                                                                             | Use HTTPS server.                                                                                                                                                 |
| HTTPS_KEY_PATH                                | File path                                            | Development: pre-generated key path; Production: -                         | Development: No; Production: Yes when `HTTPS` is `true`                        | HTTPS private key file path.                                                                                                                                      |
| HTTPS_CERT_PATH                               | File path                                            | Development: pre-generated cert path; Production: -                        | Development: No; Production: Yes when `HTTPS` is `true`                        | HTTPS certificate file path.                                                                                                                                      |
| CREATE_IDENTITY_REQUEST_MESSAGE_TEMPLATE_PATH | File path                                            | `../request_message_templates/create_identity.mustache`                    | Development: No; Production: Yes;                                              | Request message template in mustache format filepath to use in consent request when creating identity.                                                            |
| ADD_ACCESSOR_REQUEST_MESSAGE_TEMPLATE_PATH    | File path                                            | `../request_message_templates/add_accessor.mustache`                       | Development: No; Production: Yes;                                              | Request message template in mustache format filepath to use in consent request when adding new accessor.                                                          |
| CALLBACK_RETRY_TIMEOUT                        | Number                                               | `600`                                                                      | No                                                                             | Callback retry timeout in seconds. Only applies to some callbacks (that do not have shouldRetry function check e.g. request status update callback to RP client). |

- `ROLE`: Allowed values are `idp`, `rp`, `as`, and `ndid`. [Required]
- `TENDERMINT_IP`: IP Address to contact tendermint RPC. [Default: `localhost`]
- `TENDERMINT_PORT`: Port to contact tendermint RPC. [Default: `45000` for IDP, `45001` for RP, and `45002` for AS]
- `MQ_CONTACT_IP`: An IP address where node's message queue can be contacted. [Required]
- `MQ_BINDING_PORT`: A port to bind message queue. [Default: `5555`]
- `SERVER_PORT`: API server port. [Default: `8080`]
- `PRIVATE_KEY_PATH`: Path to node's private key. [Default: using pre-generated development key][required in production if `use_external_crypto_service` is set to false]
- `PRIVATE_KEY_PASSPHRASE`: Passphrase for node's private key.
- `MASTER_PRIVATE_KEY_PATH`: Path to node's master private key. [Default: using pre-generated development key][required in production if `use_external_crypto_service` is set to false]
- `MASTER_PRIVATE_KEY_PASSPHRASE`: Passphrase for node's master private key.
- `USE_EXTERNAL_CRYPTO_SERVICE`: Use external service for decrypting and signing with node's key (e.g. HSM). [Default: `false`]
- `NODE_ID`: Node ID. This ID is tied to public key. In development, we have `rp1`, `rp2`, `rp3`, `idp1`, `idp2`, `idp3`, `as1`, `as2`, and `as3`. [Required]
- `DATA_DIRECTORY_PATH`: Directory path for persistence data files. [Default: `__dirname/../data` (`data` directory in repository's directory)]
- `LOG_LEVEL`: Log level. Allowed values are `error`, `warn`, `info`, `verbose`, `debug` and `silly`. [Default: `debug` in development, `info` in production]
- `LOG_TARGET`: Where should logger writes logs to. Allowed values are `console` and `file`. [Default: `console` in development, `file` in production]
- `LOG_COLOR`: Log highlight color. [Default: `true` when log target is console, `false` otherwise]
- `LOG_DIRECTORY_PATH`: Directory path for log files. (only in `production` environment) [Default: `__dirname/../log` (`log` directory in repository's directory)]
- `CLIENT_HTTP_ERROR_CODE`: HTTP error code when responding a client error. [Default: `400`]
- `SERVER_HTTP_ERROR_CODE`: HTTP error code when responding a server error. [Default: `500`]
- `HTTPS`: Use HTTPS server. [Default: `false`]
- `HTTPS_KEY_PATH`: HTTPS private key file path. [Default: pre-generated development key][required in production if `https` is set to true]
- `HTTPS_CERT_PATH`: HTTPS certificate file path. [Default: pre-generated development cert][required in production if `https` is set to true]
- `CREATE_IDENTITY_REQUEST_MESSAGE_TEMPLATE_PATH`: Request message template in mustache format filepath to use in consent request when creating identity. [Default: `../request_message_templates/create_identity.mustache`][required in production]
- `ADD_ACCESSOR_REQUEST_MESSAGE_TEMPLATE_PATH`: Request message template in mustache format filepath to use in consent request when adding new accessor. [Default: `../request_message_templates/add_accessor.mustache`][required in production]
- `CALLBACK_RETRY_TIMEOUT`: Callback retry timeout in seconds. Only applies to some callbacks (that do not have shouldRetry function check e.g. request status update callback to RP client). [Default: `600`]

**_Examples_**

- Run a server as an IDP

  ```sh
  ROLE=idp \
  TENDERMINT_IP=127.0.0.1 \
  TENDERMINT_PORT=45000 \
  MQ_CONTACT_IP=127.0.0.1 \
  MQ_BINDING_PORT=5555 \
  SERVER_PORT=8080 \
  NODE_ID=idp1 \
  npm start
  ```

- Run a server as a RP

  ```sh
  ROLE=rp \
  TENDERMINT_IP=127.0.0.1 \
  TENDERMINT_PORT=45001 \
  MQ_CONTACT_IP=127.0.0.1 \
  MQ_BINDING_PORT=5556 \
  SERVER_PORT=8081 \
  NODE_ID=rp1 \
  npm start
  ```

- Run a server as a AS

  ```sh
  ROLE=as \
  TENDERMINT_IP=127.0.0.1 \
  TENDERMINT_PORT=45000 \
  MQ_CONTACT_IP=127.0.0.1 \
  MQ_BINDING_PORT=5557 \
  SERVER_PORT=8082 \
  NODE_ID=as1 \
  npm start
  ```

Don't forget to

1.  Set `SERVER_PORT` when running on the same machine to avoid port collision.
2.  Set `TENDERMINT_IP` and/or `TENDERMINT_PORT` when running `smart-contract`/`tendermint` on another machine.
3.  Set `NODE_ID` when there are more than one node per role in the system.

## Run in Docker

Required

- Docker CE 17.06+ [Install docker](https://docs.docker.com/install/)
- docker-compose 1.14.0+ [Install docker-compose](https://docs.docker.com/compose/install/)

### Run

```
npm run docker-up
```

or

```
docker-compose -f docker/docker-compose.yml up
```

### Build

```
npm run docker-build
```

or

```
./docker/build.sh
```

### Note

- To run docker container without building image, run command show in **Run** section (no building required). It will run docker container with image from Dockerhub (https://hub.docker.com/r/ndidplatform/api/).
- To pull latest image from Dockerhub, run `docker pull ndidplatform/api`

## Note

- When working in development, if you clear/delete the blockchain, you need to delete DB files and latest block height files by running `npm run reset-data-for-dev`. (Automatically run when running `npm run initDevKey`)

- Run `npm run delete-local-db-cache` to delete local DB used for caching. Local DB file name is `db-api-` following by node ID (env: `NODE_ID`) set on server start (e.g. `db-api-idp1` when node ID is set to `idp1`).

- Run `npm run delete-latest-block-height` to delete persistent latest block height saved by the server. File name is `latest-block-height-` following by node ID (env: `NODE_ID`).
