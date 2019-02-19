# Ethereum Voting Contracts
> Orbs voting is performed using the Ethereum platform as described in the [voting flow](../../flows/voting.md))

![alt text][ethereum_contracts] <br/><br/>

[ethereum_contracts]: ../../_img/ethereum_contracts.png "Orbs Ethereum contracts"

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
* addValidator(address _validator) public onlyOwner
* isValidator(address _validator) external view returns (bool)
* getValidators() external view returns (address[])
* leave() public


&nbsp;
## OrbsValidatorsRegistry
> ownership: no
> enables nodes to advertise their details and IP addresses

### register(string _name, bytes _ipvAddress, string _website) public 
> Registers a validator's data or update an existing one.

#### Checks:
* Check that the validator is a valid validator by calling `OrbsValidators.isValidator(sender)`
* _name is not ""
* _website is not ""
* _ipAddress is 4B (IPv4, tbd - IPv6)

#### Log the validator data
* Save in a map
  
#### leave() public 
> Remove the validator from the registry

#### getData(address _validator) public view return (string _name, bytes _ipvAddress, string _website)
> Read a validator's data


&nbsp;
### OrbsVoting
> ownership: none

#### vote(address[] nodes)
> Access:public
> Enable an account to cast a vote on the approved nodes.
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
> returns the current vote of an activist along with the last vote block_height.


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
### OrbsActivistsRegistry
> ownership: no
> enables activists to advertise their details

#### Register(string _name, string _website) public 
> Registers an activist's data or update an existing one.
* Checks:
  * _name is not ""
  * _website is not ""

#### Leave() public 
> Enable a activists to remove itself from the registry

#### GetData(address _activist) public view return (string _name, bytes _ipvAddress, string _website)
> Read a validator's data


&nbsp;
### OrbsRewardsDistribution
> ownership: no
> enables activists to advertise their details

#### Register(string _name, string _website) public 
> Registers an activist's data or update an existing one.
* Checks:
  * _name is not ""
  * _website is not ""

#### Leave() public 
> Enable a activists to remove itself from the registry

#### GetData(address _activist) public view return (string _name, bytes _ipvAddress, string _website)
> Read a validator's data



