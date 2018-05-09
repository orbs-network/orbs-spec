# Block Storage

Holds the long term journal of all confirmed blocks, used to reconstruct the state.

&nbsp;
## `Data Structures`

#### Block database
* Holds all blocks with ability to query efficiently by block height.
* Able to return efficiently the last block height in the database.
* Needs to be persistent.
* When empty, initialized with the genesis block (empty block with no transactions).
* Data is not sharded, all blocks are stored locally on all machines.
* Recommended implementation: LevelDB
  * Fast key-value storage library written at Google.
  * Provides an ordered mapping from string keys to string values.

&nbsp;
## `Block Synchronization`

* The synchronization process occurs periodically with a configurable interval.
* Ask all block storage nodes if they have newer blocks than mine by calling `Gossip.BroadcastMessage` with "HasNewBlocksMessage" message.
* Choose one of the nodes that replies with new blocks and start a sync process to get all of its new blocks:
  * "HasNewBlocks" message to signify that there are new blocks.
  * "HasNewBlocksResponse" message to signify which node was chosen for sync.
  * "SendNewBlocks" message to request the batch of new blocks.
  * "SendNewBlocksResponse" message to contain the batch of blocks as a response.

&nbsp;
## `AddBlock` (method)
> Add a new block to the database

#### Verify block before adding to database
* Handle the edge case of adding a block while node sync is active.
* Correct protocol version.
* Block height should match the next available block height in the database.
* Field prevBlockHash is equal to the hash over the last block.

#### Add block to storage
* Store block in database indexed by block height.
* Update any indication of the last block height in the database.
* Mark all transactions in the block as committed by calling `TransactionPool.MarkCommittedTransactions`.

&nbsp;
## `GetBlocks` (method)
> Return an array of blocks starting from a specific block height and up to the last block

* Read all relevant blocks consecutively from the database into an array and return them.

&nbsp;
## `GetLastBlock` (method)
> Return the latest block height stored in the database

* Return the last block in the database.

&nbsp;
## `HasNewBlocks` (method)
> Given a block height return whether the database has newer blocks

* Compare the given block height to last block in the database.

&nbsp;
## `GossipMessageReceived` (method)
> Handle a gossip message from another node

#### On "HasNewBlocksMessage" message
* Call `HasNewBlocks` to generate a response.
* Respond by calling `Gossip.UnicastMessage` with "HasNewBlocksResponse" message.

#### On "HasNewBlocksResponse" message
* Start a sync process with one of the nodes that has new blocks (only one concurrent process allowed).
* When a new sync process starts, send the node "SendNewBlocks" message by calling `Gossip.UnicastMessage`.

#### On "SendNewBlocks" message
* Call `GetBlocks` to generate a response.
* Respond by calling `Gossip.UnicastMessage` with "SendNewBlocksResponse" message.

#### On "SendNewBlocksResponse" message
* Make sure sync process with the sending node is active.
* Add blocks to database sorted by block height by calling `AddBlock`.
