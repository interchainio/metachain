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

```proto
// GroupInfo represents the high-level on-chain information for a group.
message GroupInfo {

    // group_id is the unique ID of the group.
    uint64 group_id = 1;

    // admin is the account address of the group's admin.
    string admin = 2;
    
    // metadata is any arbitrary metadata to attached to the group.
    bytes metadata = 3;

    // slices keeps track of all the heights that snapshots of the membership
    // set are stored
    repeated uint64 slices = 4;

    // total_weight is the current sum of the group members' weights.
    string total_weight = 5;
}
```

```proto
message Member {
    // address is the member's account address.
    string address = 1;
    
    // weight is the member's voting weight that should be greater than 0.
    string weight = 2;
    
    // metadata is any arbitrary metadata to attached to the member.
    bytes metadata = 3;
}
```

## Differential Group

A differential group extends a regular group by using slices. It structures 
the group membership and membership changes in such a way as to efficiently 
persist multiple states. This is beneficial when requiring group snapshots 
which may be needed for proposals. One can view this like the following interface:

```golang
type DifferentialGroup interface {
    // Returns the group membership at that height if it exists.
    // Use a height of 0 to return the most current group
    GetSlice(height int64) *Group
    
    // Creates a new slice at the current height
    SetSlice() *Group

    // Returns all the slices that exist
    Slices() []int64

    // Deletes a slice when no other service requires it any longer
    DeleteSlice(height int64)
}
```

As an example, say the current height was 10 and we had a snapshot of the group
at height 5. We would save the slice as representing the changes made between 5
and 10.

Group at height 10:

`a` = 10
`b` = 15
`c` = 20

Group at height 5:

`a` = 10
`b` = 10
`d` = 10

Slice:

`b` = 10
`c` = 0
`d` = 10

With the group at height 10 and the slice we could work backwards and obtain the
group at height 5.

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
}
```