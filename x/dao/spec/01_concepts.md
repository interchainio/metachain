# Concepts

DAO's are an increasing commonplace framework within blockchain ecosystems. They
allow for inclusive institutions and collective decision making. This section
goes over the main features of the DAO module.

## Proposal

The base of every governance system is a proposal, the ability for the group to
collectively take an action. Proposals are submitted by one account within the
context of a group account. They usually contain an array of messages that will
get executed when a proposal is accepted. Proposals look something like the
following: 

```proto
message Proposal {
    // the unique id of the proposal
    uint64 proposal_id = 1;

    // the account proposing the action
    string proposer = 2;

    // the group account that will execute the action, define 
    // the members eligible to vote as well as the proposal filter
    // and decision policy involved with the proposal
    string group = 3;

    // the deposit associated with the proposal
    repeated sdk.Coin deposit = 4;

    // usually a url representing the case of the proposal
    string content = 5; 

    // additional meta data that might be used as part of 
    // decision policy
    bytes metadata = 6; 

    // an array of messages to be executed upon acceptance
    repeated sdk.Msg messages = 10;
}
```

### Proposal Filter

Other than the basic validation of the proposal, proposals can be constrained
via a filter which can limit what a group account is capable of proposing. This
filter has a simple interface:

```golang
type ProposalFilter interface {
    Validate(group GroupAccount, proposal Proposal, otherProposals []Proposal) error
}
```

### Deposit

Each proposal also contauns a `deposit` field. Primarily this is used as a
deterrent for spam proposals by fining proposers who propose "bad" proposals.
When this is not needed it can be left empty.


## Decision Policy

The verdict of a proposal is reached through a decision policy which takes in votes
and returns the result. Decision policies are stateless, based on GroupAccounts
and can be altered. An example of such an interface is as follows:

```golang
type DecisionPolicy interface {
    // This is called when a decision policy is first added to a DAO to ensure that
    // all the parameters are correctly set
    ValidateBasic() error

    // Vote is called for every received vote and can manipulate the tally and
    // return the result of the proposal
    Vote(vote Vote, tally Tally, votingDuration time.Duration) (Result, Tally)

    // If the result from a Vote is both complete and has passed, Execute is 
    // called which allows a decision policy to modify the set of messages before
    // they are executed
    Execute(messages []sdk.Message, tally Tally, meta []byte) []sdk.Message
}

type Result struct {
    Completed bool
    Passed bool
}
```

The DAO module aims to provide abstractions that allow for a wide range of
sophisticated and nuanced decision policies such as quadratic voting, secret
ballots, and liquid democracy. 

### Tally

```proto
message Tally {
    uint64 proposal_id = 1;

    // count is an array of the totals for each choice
    repeated uint64 count = 2;

    bytes meta_data = 3;
}
```

## Vote

A vote represents a choice from a member. To keep things 

```proto
message Vote {
    // the member of the group that voted
    string voter = 1;

    // there choice
    uint64 choice = 2;

    // additional data
    bytes meta_data = 3; 
}
```



