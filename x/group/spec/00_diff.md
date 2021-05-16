# Diff

This page captures the changes made between the current
[groups](https://github.com/regen-network/regen-ledger/tree/v1.0.0/x/group/spec)
module and what is proposed.

## Transfer Decision Policy, Voting and Proposals across to the DAO module.

The types and messages that are concerned with the governance of the group:
`DecisionPolicy`, `Vote`, `Proposal` have been moved to the `dao` module. 

The idea behind the change is to scope the `group` module just to the notion of
a weighted group of members and sets of associated accounts. Thus `group` can
change set membership, the admin and create many accounts per group but it does
not concern itself with the actions that the group decides upon.

This decoupling means that groups wishing to use the `dao` module for governance
need to set the admin to the `dao` account and register the group with the dao
(along setting a decision policy and other governance parameters). This is seen as
the standard pairing but this is not always the case. The `admin` field sets who
has authority of the group. This allows for other cases such as off-chain
governance (where the account of the off-chain service is used instead). It can
also be the address of another group of another module or any other account for
that matter.

## Add Differential Groups

A different
