<!--
order: 1
-->

# Concepts

## Group

A group is simply an aggregation of accounts with associated weights. It is not
an account and doesn't have a balance. It doesn't in and of itself have any
sort of voting or decision weight. It does have an "administrator" which has
the ability to add, remove and update members in the group. An administrator of
a group could be a single user, another group, a module such as `dao` or some 
equivalent that offers governance capabilities. It could even be an account 
connected with off-chain governance.

Groups can have different implementations largely depending on the tradeoffs
between frequency of weight changes and membership changes, the size of the
group, and the complexity of the decision policies. To unify these tradeoffs
groups can be represented as the following interface:

```golang
type Group interface {
    // Get's the current memberset and height. This also signals to the 
    // group that a process requires the memberset and may need it at a later point.
    GetMemberSet() (MemberSet, uint64)

    // Get's a member at the current height
    GetMember(member string) (Member, uint64)

    // Returns the weight of a specific member at the specified height.
    GetMemberAtHeight(member string, height uint64) (Member, error)

    // Returns the member set at that height
    GetMemberSetAtHeight(height uint64) (MemberSet, error)

    // Returns all the heights that the group has knowledge of the MemberSet
    GetAllMemberSetHeights() []uint64

    // Signals that the MemberSet at that height is no longer needed by a process
    ReleaseMemberSet(height uint64) 
}

type MemberSet []Member

type Member struct {
    address string
    weight string
    meta []byte
}
```
Note that this interface only pertains to reading the `Group`. Writes to group
membership are subject to the individual concrete types.

### Static Group

A static group is the simplest implementation. It doesn't retain any temporal
knowledge rather only holds the memberset at the current height. This is
suitable for groups whose membership doesn't change too often or for groups that
employ a governance mechanism which only requires the memberset at a single
height (i.e. when votes are counted at the end of a proposal)

### Differential Group

A differential group extends a regular group by using slices. It structures 
the group membership and membership changes in such a way as to efficiently 
persist multiple states. This is beneficial when requiring group snapshots 
which may be needed for certain proposals.

As an example, say the current height was 10 and we had a snapshot of the group
at height 5. We would save the slice as representing the changes made between 5
and 10.

Group at height 10:

`a` = 10, `b` = 15, `c` = 20

Group at height 5:

`a` = 10, `b` = 10, `d` = 10

Slice:

`b` = 10, `c` = 0, `d` = 10

With the group at height 10 and the slice we could work backwards and obtain the
group at height 5. When a membership update happens, we thus update the group at
height 10 and sometimes the slice in between (Note: that all slices apart from the
most recent stay the same, hence this is quite efficient):

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


### Epoch Group

For groups that update extremely often (like token groups) and can have many
simulatenous proposals, computation in updating the slice and computation in
generating the group at a specified height can be too high. In this case we have
an epoch group. Epoch groups keep the current member set and one prior snapshot.
This snashot has the characteristic of being updated every x heights - hence
defining an epoch. Proposals within an epoch all use the snapshot made at the
beginning of the epoch as a reference for the member set. Users of such a group
should be wary of proposals that cross over multiple epochs and use an
appropriate governance mechanism.

## Group Account

A group account is a group as well as an account. Group accounts are abstracted 
from groups because a single group may have multiple accounts designed for 
different purposes and having different governance mechanisms. Managing group
membership separately from decision policies results in the least overhead
and keeps membership consistent across different policies. The pattern that
is recommended is to have a single master group account for a given group,
and then to create separate group accounts with different decision policies
and delegate the desired permissions from the master account to
those "sub-accounts" using the `x/authz` module. A group account has an
administrator which could be set up in the same ways as a group administrator

```proto
message GroupAccount {
    // the group account address
    string address = 1;

    // the address of the admin of the account
    string admin = 2; 

    // the group_id the account is associated with
    uint64 group_id = 3;

    // metadata is any arbitrary metadata attached to the group account.
    bytes metadata = 4;
}
```