## Ethereum Voting Contracts
> Orbs voting is performed using the Etehreum platform as described in the [voting flow](../../flows/voting.md))

&nbsp;
### OrbsFederation => OrbsValidators 
> ownership: yes
> Manages the list of node candidates that underwwent due-diligence

#### Exisiting functions
**IF:**
* isMember(address _member) external view returns (bool)
* getMembers() external view returns (address[])
* getConsensusThreshold() external view returns (uint)
* getFederationRevision() external view returns (uint)
* isMemberByRevision(uint _federationRevision, address _member) external view returns (bool)
* getMembersByRevision(uint _federationRevision) external view returns (address[])
* getConsensusThresholdByRevision(uint _federationRevision) external view returns (uint)

**Owner / internal:**
* addMember(address _member) public onlyOwner
* removeMember(address _member) public onlyOwner
* removeMemberByIndex(uint _i) public
* upgradeSubscriptionManager(Upgradable _newSubscriptionManager) public onlyOwner
* findMemberIndex(address[] _members, address _member) private pure returns(uint, bool)
* isFedererationMembersListValid(address[] _members) private pure returns (bool)
* isMember(address[] _members, address _member) private pure returns (bool)
* getConsensusThresholdForMembers(address[] _members) private pure returns (uint)

#### New functions / updates
* getRevisionByBlockHeight(block_height) external view returns (uint)
> Retruns the active revision for the block_height.

* removeMemberByIndex(uint _i) internal
  * Bug fix - Internal instead of public.
  
* removeMember(address _member) public onlyOwner
  * Allow only before March 26 end of day.
  
* endMembership() public
> Allows a validator to remove itself from the list


&nbsp;
### OrbsVoting
> ownership: none

#### vote(address[] nodes)
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
* Increment the global vote_counter.
  * Used as a reference to indicate that all votes were counted.
* Store the vote for the activist in state along with its block_height
  * Note: for state storage efficiency, the federation contract may allocate unique (never reused) ids per node and use a mapping.


#### getCurrentVote(address activist) returns ([]address nodes, uint block_height)
> Access:public, view
> returns the current vote of an activist alomg with the last vote block_height.


#### delegate(address delegatee)
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


&nbsp;
### OrbsValidatorsRegistery
> ownership: no
> enables nodes to advertise their detalis and IP addreses

#### RegisterValidator(string _name, bytes _ipvAddress, string _website) public 
> Registers a validator's data or update an existing one.
* Checks:
  * _name is not ""
  * _website is not ""
  * _ipAddress is 4B (IPv4, tbd - IPv6)

#### RemoveValidator() public 
> Removes a validator from the registery


&nbsp;
### OrbsActivistsRegistery
> ownership: no
> enables activists to advertise their detalis

#### RegisterActivist(string _name, string _website) public 
> Registers a activist's data or update an existing one.
* Checks:
  * _name is not ""
  * _website is not ""

#### RemoveActivist() public 
> Removes a activist from the registery



