# Messages

## MsgCreatePolity

```proto
message MsgCreatePolityRequest {
    string group_address = 1;

    // decision_policy is the updated group account decision policy.
    google.protobuf.Any decision_policy = 2 [(cosmos_proto.accepts_interface) = "DecisionPolicy"];
    
    // proposal_filter filters incoming proposals.
    google.protobuf.Any proposal_filter = 3 [(cosmos_proto.accepts_interface) = "ProposalFilter"];
    
    // the voting window for each proposal
    google.protobuf.Duration voting_period = 4;

    // indicates whether proposals can be received from anyone or members only
    bool public = 5;
}
```

`voting_period` must be less than `max_voting_period` (default: 3 weeks)

This will update the admin of the group account to the address of the `dao`
module itself

## MsgUpdatePolity

```proto
message MsgUpdatePolityRequest {
    string address = 1;

    // decision_policy is the updated group account decision policy.
    google.protobuf.Any decision_policy = 2 [(cosmos_proto.accepts_interface) = "DecisionPolicy"];
    
    // proposal_filter filters incoming proposals.
    google.protobuf.Any proposal_filter = 3 [(cosmos_proto.accepts_interface) = "ProposalFilter"];
    
    // the voting window for each proposal
    google.protobuf.Duration voting_period = 4;

    // indicates whether proposals can be received from anyone or members only
    bool public = 5; 

    // an optional parent polity which has administative power over the polity
    string parent_address = 6;
}
```

## MsgDestroyPolity

```proto
message MsgDestroyPolityRequest{
    string address = 1;
}
```

## MsgCreateProposal

```proto
message MsgCreateProposalRequest {
    option (gogoproto.goproto_getters) = false;

    // group address is the group account address.
    string address = 1;
    
    // proposer is the account address of the proposers.
    string proposer = 2;

    // the deposit associated with the proposal
    repeated cosmos.base.v1beta1.Coin deposit = 3;

    // content represents some endpoint which can lead 
    // a discussion / forum or simple state the case for // the proposal
    string content = 5;
    
    // metadata is any arbitrary metadata to attached to the proposal.
    bytes metadata = 6;

    // messages is a list of Msgs that will be executed if the proposal passes.
    repeated google.protobuf.Any messages = 7;
}
```

## MsgAmendProposal

```proto
message MsgAmendProposalRequest {
    // unique id of the proposal
    uint64 proposal_id = 1;

    // content represents some endpoint which can lead 
    // a discussion / forum or simple state the case for // the proposal
    string content = 2;

    // metadata is any arbitrary metadata to attached to the proposal.
    bytes metadata = 3;

    // messages is a list of Msgs that will be executed if the proposal passes.
    repeated google.protobuf.Any messages = 4;
}
```

The signer of the message must be the proposer of that proposal. Any field left
empty leaves that part of the proposal untouched. The amendment must pass
`ProposalFilter.Amend` function without error.

## MsgWithdrawProposal

```proto
message MsgWithdrawProposalRequest {
    uint64 proposal_id = 1;
}
```

The signer of the message must be the proposer of that proposal. In addition the
proposal must pass the `ProposalFilter.Withdraw()` function without an error. 

## MsgVote

```proto
// MsgVoteRequest is the Msg/Vote request type.
message MsgVoteRequest {
    // proposal is the unique ID of the proposal.
    uint64 proposal_id = 1;

    // voter is the voter's address.
    string voter = 2;

    // choice is the voter's choice on the proposal.
    Choice choice = 3;

    // the weight behind their choice. This defaults to the entire weight 
    // of the member if none is described. This allows for split voting.
    string weight = 4;
}
```

The `proposal_id` must be for a valid and open proposal, the voter must be a
member of the group corresponding to the proposal. A weight of 0 corresponds to
the entire weight of a member.

> NOTE: There's nothing stopping a member from voting multiple times. The
> decision policy is in charge of correctly handling multiple votes. In the
> future we may want to support secret ballots. They most often involve two
> votes, one for each phase.

## MsgDelegate

```proto
message MsgDelegateVoteRequest { 
    // the account transferring their voting power
    string delegator = 1;

    // the receiving of the voting power
    string candidate = 2;

    // the weight they are transferring the delegation to
    string weight = 3; 
}
```

To delegate, both the `delegator` and `candidate` must be part of the group and
the `weight` must be less than or equal to the `delegator`'s total weight. A
weight of 0 implies the `delegator`'s total weight.

> NOTE: We may want to introduce rules to prevent delegation circle i.e. I
> delegate to you and you delegate to me.

## MsgUndelegate

```proto
message MsgUndelegateRequest { 
    // the delegator
    string delegator = 1;

    // the address being removed of the delegation
    string candidate = 2;

    // the weight they are transferring the delegation to
    string weight = 3; 
}
```

To undelegate, a delegator must have already delegated greater than or equal to
`weight` of their weight.

# MsgExec


```proto
// MsgExecRequest is the Msg/Exec request type.
message MsgExecRequest {

    // proposal is the unique ID of the proposal.
    uint64 proposal_id = 1;
    
    // signer is the account address used to execute the proposal.
    string signer = 2;
}
```