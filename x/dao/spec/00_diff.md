# Diff

This page tries to capture the difference between the existing
[Gov](https://github.com/cosmos/cosmos-sdk/tree/v0.42.4/x/gov/spec) module and
this DAO one.

Note that there is a considerable change in what the DAO module tries to achieve
from the outset. Gov may be considered the governance of the entire chain as a
whole only whereas DAO represents the toolkit for groups to govern one another
where one group most likely represents the chain as a whole (the staking group).
Some of the diff described here may be more in relation to the groups module
which also housed data structures such as `Vote` and `Proposal`

## Proposal Filter

This is a relatively novel feature (taken from
[PolicyKit](https://policykit.readthedocs.io/en/latest/policy_model.html#filter))
which allows for several features. Predominantly this allows for categorization
of proposals and thus a divide and conquer form of governannce or
representational forms of governance. A standard case would 