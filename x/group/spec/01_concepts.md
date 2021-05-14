<!--
order: 1
-->

# Concepts

## Group

A group is simply an aggregation of accounts with associated weights. It is not
an account and doesn't have a balance. It doesn't in and of itself have any
sort of voting or decision weight. It does have an "administrator" which has
the ability to add, remove and update members in the group. An administrator of
a group could be a single dictatorial user, another group, a module such as
`dao` or some equivalent that offers governance capabilities. It could even be
an account connected with off-chain governance.

```proto
message Group {
    // group_id is the unique ID of this group
    uint64 group_id = 1;

    // admin is the account address of the admin
    string admin = 2;

    // the members within the group
    repeated Member members = 3;
}
```

```proto
message Member {
    // the member accounts address
    string address = 1;

    // the members weight in the group
    string weight = 2;
}
```

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