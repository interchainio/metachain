# Concepts

DAO's are an increasing commonplace framework within blockchain ecosystems. They
allow for inclusive institutions and collective decision making. This section
goes over the main features of the DAO module.

## Polity

> NOTE: The terminology here is still subject to change.

A polity is a political entity, a group of members coupled with a governance
system and a communual account.

```proto
message Polity {
    // the account or address that collectively represents the polity
    string account = 1;

    // the group interface for reading member sets
    google.protobuf.Any group = 2 [(cosmos_proto.accepts_interface) = "Group"];

    // decision_policy is the updated group account decision policy.
    google.protobuf.Any decision_policy = 3 [(cosmos_proto.accepts_interface) = "DecisionPolicy"];
    
    // proposal_filter filters incoming proposals.
    google.protobuf.Any proposal_filter = 4 [(cosmos_proto.accepts_interface) = "ProposalFilter"];
    
    // the voting window for each proposal
    google.protobuf.Duration voting_period = 5;

    // indicates whether proposals can be received from anyone or members only
    bool public = 6;

    // indicates which account has control over the polity i.e. changing 
    // the decision policy or the proposal filter. This allows for nested polity's 
    // where a parent polity has partial governance over this polity. Leave it empty
    // to indicate that the polity is soverign 
    string parent = 7; 
}
```

The typical flow would be:
    1. Set up a group
    2. Set up a polity linking to the group

Optional extra steps
    3. Change the group admin to the polity account (this allows proposals to
    dictate membership changes)

    4. Setup another polity with the same group but with a different decision policy 
    or for a different category of policies within the same group

    5. Create another group and another polity linked to that group. Set the parent
    of the new polity as the old one and set the admin of the new group to the old 
    polity and you have an elected council.


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
    string group_address = 3;

    // the deposit associated with the proposal
    repeated sdk.Coin deposit = 4;

    // usually a url making the case for the proposal
    string content = 5; 

    // additional meta data that might be used as part of 
    // decision policy
    bytes metadata = 6; 

    // an array of messages to be executed upon acceptance
    repeated sdk.Msg messages = 7;
}
```

### Proposal Filter

Other than the basic validation of the proposal, proposals can be constrained
via a filter which can limit what a group account is capable of proposing,
amending or withdrawing. This filter has the following interface:

```golang
type ProposalFilter interface {
    // Validate assesses the validity of a new proposal within a polity
    Validate(group Group, proposal Proposal, otherProposals []Proposal) error

    // Amend assesses whether the old porposal can be ammended with the new proposal
    Amend(oldProposal, newProposal Proposal) error

    // Withdraw assesses whether the proposer is able to withdraw their proposal
    Withdraw(group Group, proposal Proposal) error
}
```

### Deposit

Each proposal also contains a `deposit` field. Primarily this is used as a
deterrent for spam proposals by fining proposers who propose "bad" proposals.
When this is not needed it can be left empty. The funds of the proposal go
directly to the account. It would thus be commong practice within the proposal
to also request to get your deposit back.

## Decision Policy

The verdict of a proposal is reached through a decision policy which takes in votes
and returns the result. Decision policies are stateless, based on GroupAccounts
and can be altered. An example of such an interface is as follows:

```golang
type DecisionPolicy interface {
    // This is called when a decision policy is first added to a DAO to ensure that
    // all the parameters within it are correctly set
    ValidateBasic() error

    // Vote is called for every received vote and can manipulate the tally and
    // return the result of the proposal
    Vote(vote Vote, tally Tally) (Result, Tally)

    // End is called when a proposal passes the voting window and a verdict must
    // be reached.
    End(tally Tally) Result

    // If the result from a Vote is both complete and has passed, Execute is 
    // called which allows a decision policy to modify the set of messages before
    // they are executed
    // TODO: We may want to consider running execute when a proposal is defeated 
    // to allow for fallback actions to be executed
    Execute(messages []sdk.Message, tally Tally, meta []byte) []sdk.Message
}

type Result struct {
    Completed bool
    Passed bool
}
```

Votes are counted based on the weights of the member at the moment the proposal
was created. 

The DAO module aims to provide abstractions that allow for a wide range of
sophisticated and nuanced decision policies such as quadratic voting, secret
ballots, and liquid democracy. 

### Tally

A tally is the underlying data struture that tracks the progress of a proposal
over it's lifecycle.

```proto
message Tally {
    // the proposal the tally is tracking
    uint64 proposal_id = 1;

    // count is an array of the totals for each choice
    repeated Choice votes = 2; 

    // relevant meta data. This is useful for caching infomation like the total 
    // accumulated weight for each choice
    bytes meta_data = 3;
}
```

More information on the composition of choices can be found below. 

## Vote

A vote represents a choice from a member for a proposal. Choices are
intentionaly left as abstract as possible. They could be a simple binary yes/no,
a set of multiple choices or an amount (in the case of a community spend).


```proto
message Vote {
    // the member of the group that voted
    string voter = 1;

    // the proposal that they are voting on
    uint64 proposal_id = 2;

    // their choice
    Choice choice = 3;

    // the weight behind their choice. This defaults to the entire weight 
    // of the member if none is described. This allows for split voting.
    string weight = 4;
}
```

> NOTE: We could look to make this non breaking if we reserve 3 and handle
> VoteOptions with a reasonable default

### Choices

Each vote contains a choice. What choices are really depends on the users of the
dao and how they set up the proposal, the proposal filter and the decision
policy. A choice is thus more a represenation of the index of possible choices. For
example 0 = hasn't voted, 1 = no, 2 = yes and so forth. What's important to note
is the degree of freedom of choices. Currently we represent a choice as a single
byte which means a proposal is constrained to 256 options. This should be
sufficient for most proposals but in the future we could set it as multiple
bytes (i.e. uint32)

```golang
type Choice byte
```

Choices have a default 0. This should usually be something likes "hasn't voted"
but could technically be anything such as default yes unless someone votes no.

### Delegation

A feature of governance that is important to port over to the dao
module is the ability to delegate votes. Delegation could be seen on a per
proposal basis or overall. Given its simplicity and that it is perhaps more 
commonplace to trust someone's judgement overall then to trust them specifically
for a proposal, we treat delegation (at least to begin with) as overall

```proto
message Delegate {
    // the account transferring their voting power
    string delegator = 1;

    // the proposal their referring to
    uint64 proposal_id = 2;

    // the receiving of the voting power
    string candidate = 3;

    // the weight they are transferring the delegation to
    string weight = 4; 
}
```

A member can equally undelegate. 



