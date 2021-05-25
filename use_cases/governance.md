# Chain-wide Governance

The almost archetypal use case for the new dao system is to obtain feature
parity with the current governance module.

# User story

A community member, Alice, wants to submit an on-chain proposal to change a parameter, the average number of blocks per year, which is used to calculate the inflation rate for the chain. To do this Alice first asks in a chat forum discord for instance whether this is a good idea and something the community would like to see happen. There is some initial discussion to confirm that this is in fact something the community wants. Another community member, Bob, also offers to collaborate on the proposal.

Alice and Bob have a zoom call and start working in a google doc to draft the proposal synchronously, after which Alice finishes the draft and Bob reviews her work. Alice then opens a pull request on the governance repo that includes the text document as well as the json message required to make the parameter proposal on chain.

Alice solicits community feedback on the PR, sharing it to the Discord and among validators, and is asked to make some minor changes, which are completed before the PR is finalized and merged by the governance repo owner.
Once the proposal has been finalized an IPFS hash of the README.md is added to the json.
The proposal is then submitted on chain through the CLI and a Cosmos forum post is made to notify the community that the proposal has been submitted. Links to the forum post are then shared in various community channels and on twitter. The merits of the proposal are discussed in these respective channels and validators / ATOM holders vote.


## Basic Implementation 

In this situation, the staking account creates and manages a group containing
validators and delegators with weights tracking the amount of tokens staked. The
`dao` module can create a gov `Polity` in which members can make proposals to
change params and spend community funds. Proposals contain abstract messages so
it is up to the other modules, `params` and `distribution` to grant authority
and safely update authority when it changes [NOTE: I'm not sure how the `authz`
comes into this].

The `ProposalFilter` is used to ensure a minimum deposit is fulfilled. A
proposal should include a transfer of the deposit back to the proposer.

The `DecisionPolicy` should use a `ThresholdPolicy`

## Further Options

1. **Creating a council** - the chain may want to delegate responsibility for
   certain parts of the network to be managed to a subgroup. To do this you
   would: 
   - Create a new proposal with the messages to create a new group with the
     members and weights of the council members. Set the admin as the parent
     `Polity` so that they can update the member set when needed. Add the create
     polity message setting the parent chain as 