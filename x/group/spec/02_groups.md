# Groups

## Static Group

A static group is the simplest implementation. It doesn't retain any temporal
knowledge rather only holds the memberset at the current height. This is
suitable for groups whose membership doesn't change too often or for groups that
employ a governance mechanism which only requires the memberset at a single
height (i.e. when votes are counted at the end of a proposal)

## Differential Group

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


## Epoch Group

For groups that update extremely often (like token groups) and can have many
simulatenous proposals, computation in updating the slice and computation in
generating the group at a specified height can be too high. In this case we have
an epoch group. Epoch groups keep the current member set and one prior snapshot.
This snashot has the characteristic of being updated every x heights - hence
defining an epoch. Proposals within an epoch all use the snapshot made at the
beginning of the epoch as a reference for the member set. Users of such a group
should be wary of proposals that cross over multiple epochs and use an
appropriate governance mechanism.
