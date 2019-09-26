# Ethereum Voting Contracts
> Orbs voting is performed using the Ethereum platform as described in the [voting flow](./election-flow.md)

> For more defails on Orbs PoS architecture, see [Orbs PoS architecture](https://github.com/orbs-network/orbs-spec/tree/pos2/pos-architecture)

<!--
![alt text][ethereum_contracts] <br/><br/>

[ethereum_contracts]: behaviors/_img/ethereum_contracts.png "Orbs Ethereum contracts"
-->

&nbsp;
## OrbsValidators 
> ownership: yes
> upgradeable: yes
> Manages the list of node candidates that underwent due-diligence

#### Public variables / constants
* VERSION
* MAX_VALIDATORS 

#### Events
* event ValidatorAdded(address indexed validator);
* event ValidatorLeft(address indexed validator);

#### Functions
* addValidator(address validator) public onlyOwner
* removeValidator(address validator) public onlyOwner
* isValidator(address validator) public view returns (bool)
* getValidators() public view returns (address[])
* getValidatorLastUpdate(address validator) public view returns (uint block_height)
* leave() public

&nbsp;
## OrbsValidatorsRegistry
> ownership: no
> enables nodes to advertise their details and IP addresses

#### Events
* event ValidatorRegistered(address indexed validator);
* event ValidatorLeft(address indexed validator);

#### Functions
* register(string name, bytes ipvAddress, string website, bytes20 orbsAddress, bytes32 validatorDeclarationHash) public 
* getValidatorData(address validator) public view return (string name, bytes ipvAddress, string website, address OrbsAddress, uint registration_block_height, uint last_update_block_height)
* getOrbsAddress(address _validator) public view return (address _orbsAddress)
* isValidator(address validator) public view returns (bool)
* leave() public


### register(string name, bytes ipvAddress, string website, bytes20 orbsAddress, bytes32 validatorDeclarationHash) public 
> Registers a validator's data or update an existing one.

#### Checks:
* Check that the validator is a valid validator by calling `OrbsValidators.isValidator(sender)`
* _name is not "" and unique
* _website is not "" and unique
* _ipAddress is 4B (IPv4, tbd - IPv6) and unique

#### Log the validator data
* Save in a map
  * Include also: first registration_block_height, last_update_block_height
  
#### leave() public 
> Remove the validator from the registry

#### getValidatorData(address validator) public view return (string name, bytes ipvAddress, string website, address OrbsAddress, uint registration_block_height, uint last_update_block_height)
> Access:public, view
> Read a validator's data

#### getOrbsAddress(address _validator) public view return (address _orbsAddress)
> Access:public, view
> Read a validator's orbs address


&nbsp;
### OrbsVoting
> ownership: none

#### Events
* event DisapproveVote(address indexed voter, bytes20[] validators, uint voteCounter);

#### Functions
* disapprove_vote(address[] validtors) public
* getLastVote(address guardian) returns ([]address validators, uint block_height)

#### disapprove_vote(address[] validtors)
> Access:public
> Enable an account to cast a vote on the approved nodes.
* Checks:
  * non empty list.
  * non 0
* Emit event with the vote: 
  * `Vote(address voter, byte[] nodes_list, uint vote_counter)`
    * voter: sender
    * nodes list: concatenated addresses - i.e. length = 20 x voted nodes.
* Increment the global vote_counter.
  * Used as a reference to indicate that all votes were counted.
* Store the vote for the guardian in state along with its block_height
  * Note: for state storage efficiency, the federation contract may allocate unique (never reused) ids per node and use a mapping.


#### getLastVote(address guardian) returns ([]address nodes, uint block_height)
> Access:public, view
> returns the current vote of an _guardian along with the last vote block_height.


#### delegate(address delegatee)
> Access:public
> Delegates the account stake to another account.
* Checks: N/A
  * Note: delegation to `0` indicates stop delegating.
* Emit event with the vote: 
  * `Delegate(address stakeholder, address to, uint delegate_counter)`
    * stakeholder: sender
    * to: delegatee
* Increment the global delegate_counter.
  * Used as a reference to indicate that all delegations were counted.


&nbsp;
## OrbsGuardians 
> ownership: yes
> upgradeable: yes
> Manages the list of node candidates that underwent due-diligence

#### Public variables / constants
* VERSION
* MAX_VALIDATORS

#### Events
* event GuardianAdded(address indexed validator);
* event GuardianLeft(address indexed validator);

#### Functions
* addGuardian(address _guardian) public onlyOwner
* isGuardian(address _guardian) external view returns (bool)
* getGuardians() external view returns (address[])
* leave() public


&nbsp;
### OrbsGuardiansRegistry
> ownership: no
> enables guardians to advertise their details

#### Register(string _name, string _website) public 
> Registers an guardian's data or update an existing one.
* Checks:
  * _name is not ""
  * _website is not ""

#### Leave() public 
> Enable a guardian to remove itself from the registry

#### GetGuardianData(address _guardian) public view return (string _name, string _website)
> Read a guardian's data


