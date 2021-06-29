# ADR 043: Governance and Groups Redesign

## Changelog

- 2020/06/16: Initial Draft

## Status

DRAFT Not Implemented

## Abstract

This ADR defines a set of changes made to `x/group` and `x/gov` to support a broader set of governance systems from small multisig accounts and social recovery systems, to asset management groups, DAO's and at the largest scale, chain-wide governance. 

## Context

The `x/group` module was formed predominantly from the desire to have on-chain multisig functionality that allowed for easy membership changes. However, groups is actually a very core module to cosmos chains! (think of groups in Linux). We don't just invoke the notion of groups for multisig accounts but it has been around since the beginning with the staking group, delegation system and chain-wide governance. What's required is a means of unifying our understanding of collections of users within an ecosystem that enables for versatile usage. Just as an address is associated with an account and permissions, so too are groups comprised of a bundle of accounts and permissions. What distinguishes groups are the rules that dictate how membership changes and the rules that dictate how an action is made. The first iteration of the groups module introduced such concepts. This ADR looks to further build on these.

## Decision

As covered in the context section, we can look at governance or any form of collective decision making as having two parts: who is involved in the decision and how is the decision reached. This abstraction is the core of how the cosmos ecosystem wil provide governance. More concretely, we define the groups module as fulfilling the role of membership management: who is part of the group and in what capacity and the governance module as fulfilling the role of a consensus system: how does the group propose an action and decide upon one. The interface between the two occurs in two places: 1) The governance module queries an aspect of membership within a group, 2) The governance module updates the membership of a group.

**Querying a Group**

The group module allows for various implementations to be fulfilled depending on the complexity required. Each group is unified by a single interface that can be used by `x/gov` or any other module

```golang
type Group interface {
    // Get's the current memberset and height
    GetMemberSet() []Member

    // Get's a member at the current height
    GetMemberWeight(member string) string
  
  	// -- Temporal Methods --
    // Returns the weight of a specific member at the specified height
    GetMemberWeightAtHeight(member string, height uint64) (string, error)

    // Returns the member set at that height
    GetMemberSetAtHeight(height uint64) (MemberSet, error)

    // Returns all the heights that the group has knowledge of the MemberSet
    GetAllMemberSetHeights() []uint64
  
  	// Allocates a snapshot of the member set for the current height by a process
  	AllocateMemberSet() []Member

    // Signals that the MemberSet at that height is no longer needed by a process
    ReleaseMemberSet(height uint64)
}

type Member struct {
    address string
    weight string
}
```

**Updating a Group's membership**

This can be viewed as a method that the `x/gov` calls to `x/group` like so. 

```golang
func UpdateMembership(groupId uint64, memberUpdates []Member)
```

Setting a weight to 0 effectively removes the member.

Each group has an `admin` field that it uses to authorize membership changes. `x/gov` has a root module account which can be set by any group and will essentially hand control to the governance module. What this also allows is for any arbitrary module, group or account to change membership, accomodating for multiple implementations of governance. By leaving admin empty, any member within that group has full control of membership and the group accounts that might be associated with it. It is expected that an organization will create a group, create a dao account via `x/gov` then add the `x/gov` module account to the group. 

### Types of Groups

This section will explore some initially supported group structures.

- **Static Group:** A static group is the simplest implementation. It doesn't retain any temporal knowledge rather only holds the memberset at the current height. This is suitable for groups whose membership doesn't change too often or for groups that employ a governance mechanism which only requires the memberset at a single height (i.e. when votes are counted at the end of a proposal)
- **Differential Group**: A differential group extends a regular group by using slices. It structures the group membership and membership changes in such a way as to efficiently persist multiple states. This is beneficial when requiring group snapshots which may be needed for certain forms of governance.
- **Epoch Group**: An epoch group also allows for snapshotting however It can only be changed every predefined epoch period. Proposals should not be allowed to flow over multiple epochs
- **Staking Group**: A staking group is an extension of the differential group which is kept within `x/staking`. The weights of members are proportionate to the value of tokens they stake. With a staking group, the module controls how membership is updated and only exposes methods for accounts to `stake` and `unstake`.



### Governance changes

So far, the only changes mentioned to `x/gov` are the addition of a root module account and the reading of group membership in the groups module over storing the members directly. Whilst the scope of changes to governance in this ADR remains limited leaving the bulk to a later ADR, there are a few other changes that must be addressed. 

**Organizations**

Akin, to the `GroupAccount`, `x/gov` allows for the instantiation of  `Organization`'s whereby a group can control a token account.

```protobuf
message Organization {
    // the account or address that collectively represents the polity
    string account = 1;

    // the group interface for reading member sets
    google.protobuf.Any group = 2 [(cosmos_proto.accepts_interface) = "Group"];

    // minimum amount of votes required to make a decision
    uint64 quorum = 3;
    
    // threshold of weighted votes required for a proposal to pass
    uint64 threshold = 4;
    
    // amount required to 
    uint64 deposit = 5; 

    // the voting window for each proposal
    google.protobuf.Duration voting_period = 6;

    // similar to a group, this indicated who has authority to change these parameters
    string admin = 7;
}
```

An `Organization` aims to have feature parity with the existing governance mechanism. There are many ways that this can be built upon (i.e. `DecisionPolicy`s) but as mentioned earlier this the subject of another ADR.

**Proposals**

The structure of a proposal has changed. Instead of implementing a handler for executing different types of proposals, a proposal consists of an array of arbitrary messages that is executed upon success. Thus a proposal has the following form:

```protobuf
message Proposal {
    // the unique id of the proposal
    uint64 proposal_id = 1;

    // the account proposing the action
    string proposer = 2;

    // the organization that will execute the action
    uint64 organization = 3;

    // the deposit associated with the proposal
    repeated sdk.Coin deposit = 4;

    // usually a url making the case for the proposal
    string content = 5;

    // additional meta data that might be used as part of
    // decision policy
    bytes metadata = 6;

    // an array of messages to be executed upon acceptance
    repeated sdk.Msg messages = 7;

    // When the proposal finishes
    google.protobuf.Time deadline = 8;

    // The current tally of the vote
    Tally tally = 9;

    // The final outcome of the proposal
    Result result = 10;
}
```

 

## Consequences

*to be completed*

> This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future.

### Backwards Compatibility

> All ADRs that introduce backwards incompatibilities must include a section describing these incompatibilities and their severity. The ADR must explain how the author proposes to deal with these incompatibilities. ADR submissions without a sufficient backwards compatibility treatise may be rejected outright.

### Positive

{positive consequences}

### Negative

{negative consequences}

### Neutral

{neutral consequences}

## Further Discussions

There are a few outstanding points that I have questions surrounding:

- Should groups still have a `GroupAccount` implementation or rely instead on other modules to implement it?
- Should the weight of each member always be represented as a fraction of the total?
- How should an `Organization` keep a record of a group? Should it just keep a reference or should it make a cross module call to the group when it needs the information? How do 

## Test Cases

*to be completed*

## References

- {reference link}

## Appendix

**Differential group example:**

A differential group consists of an array of members (representing the current height) just like a static group and slices that contain the information needed to go from the current height to the previous snapshot

As an example, say the current height was 10 and we had a snapshot of the group at height 5. We would save the slice as representing the changes made between 5 and 10.

Group at height 10:

`a` = 10, `b` = 15, `c` = 20

Group at height 5:

`a` = 10, `b` = 10, `d` = 10

Slice:

`b` = 10, `c` = 0, `d` = 10

With the group at height 10 and the slice we could work backwards and obtain the group at height 5. When a membership update happens, we thus update the group at height 10 and sometimes the slice in between (Note: that all slices apart from the most recent stay the same, hence this is quite efficient):

I.e. for a change in `b`'s weight from 15 -> 20

Group at height 11:

`a` = 10, `b` = 20, `c` = 20

Slice (stays the same):

`b` = 10, `c` = 0, `d` = 10

For a new member `e` = 10

Group at height 12:

`a` = 10, `b` = 20, `c` = 20, `e` = 10

Slice (adds `e`):

`b` = 10, `c` = 0, `d` = 10, `e` = 0,

