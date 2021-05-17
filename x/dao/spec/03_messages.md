# Messages



# MsgCreatePolity

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

The signer of the message must be the admin of the group associated with
`group_address`

`voting_period` must be less than `max_voting_period` (default: 3 weeks)

This will update the admin of the group account to the address of the `dao`
module itself

# MsgUpdatePolity

Update polity duplicates `MsgCreatePolity` in form (since no fields in polity
except address which is used as an identifier are immutable). We have different messages to
emphasise the semantic difference. Update polity also requires the polity to
already exist. It can only be called by the `dao` module, thus this message can
only be used within a proposal

# MsgCreateProposal

```proto
// MsgCreateProposalRequest is the Msg/CreateProposal request type.
message MsgCreateProposalRequest {
    option (gogoproto.goproto_getters) = false;

    // group address is the group account address.
    string group_address = 1;
    
    // proposer is the account address of the proposers.
    string proposer = 2;

    // the deposit associated with the proposal
    repeated cosmos.base.v1beta1.Coin deposit = 3;

    // content represents some endpoint which can lead 
    // a discussion / forum or simple state the case for // the proposal
    string content = 5;
    
    // metadata is any arbitrary metadata to attached to the proposal.
    bytes metadata = 3;

    // messages is a list of Msgs that will be executed if the proposal passes.
    repeated google.protobuf.Any messages = 4;
}
```

# MsgAmendProposal

# MsgVote

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

# MsgDelegateVote

```proto
message MsgDelegateVoteRequest { 
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