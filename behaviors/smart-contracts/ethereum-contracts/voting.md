## Ethereum Voting Contracts
> Orbs voting is performed using the Etehreum platform as described in the [voting flow](../../flows/voting.md))

&nbsp;
### OrbsNodesCandidates
> ownership: yes

#### addNode(address)
> Access: public onlyOwner
> Adds a node candidate to the list of due-diliganced ndoes
* Checks:
  * not duplicate
  * non 0
  * owner

#### getNodeList() returns (address[] nodes)
> Access: public, view
> Returns the candidate nodes list

#### getNodeIndex(address node) returns int index
> Access: public, view
> Returns the index of a candidate node, if the address is not in the list return -1.

#### getNodeAddressByIndex(int index) returns address node
> Access: public, view
> Returns the adddress of a node by its index.
* Check valid index

#### GetNodesIndices(address[] nodes) returns (byte[] nodes_mask)
> Access: public, view
> Returns a bitmask of the requested addresses, if an address is not in the list ignore the address.

&nbsp;
### OrbsVoting
> ownership: none

#### Vote(address[] nodes)
> Access:public
> Enable an acount to cast a vote on the approved nodes.
* Checks:
  * non empty list.
  * non duplicate
  * non 0
  * That the nodes are within the OrbsNodesCandidates list.
* Emit event with the vote: 
  * `Vote(address voter, byte[] nodes_list, uint vote_counter)`
    * voter: sender
    * nodes list: concatenated addresses - i.e. length = 20 x voted nodes.
* Record the vote in a map of:
  * votes[voter]
    * The nodes may be efficiently stored in a bitmask using OrbsNodesCandidates.GetNodesIndices(nodes)
  * vote_block_height[voter]
* Increment the global vote_counter
  * Used as a reference to indicate that all votes were counted.

#### getCurrentVote(address account) returns ([]address nodes, uint block_height)
> Access:public, view
> returns the current vote of an account alomg with the vote block_height.

#### DelegateVote(address delegatee)
> Access:public
> Delegates the account stake to another account.
* Checks: N/A
  * Note: delegation to `0` indicates stop delegating.
* Emit event with the vote: 
  * `Delegate(address delegator, address to, uint delegate_counter)`
    * delegator: sender
    * to: delegatee
* Increment the global delegate_counter.
  * Used as a reference to indicate that all delegations were counted.

