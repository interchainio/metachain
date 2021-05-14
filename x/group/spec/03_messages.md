<!--
order: 3
-->

# Msg Service

## Msg/CreateGroup

A new group can be created with the `MsgCreateGroupRequest`, which has an admin address, a list of members and some optional metadata bytes.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L53-L64

It's expecting to fail if metadata length is greater than some `MaxMetadataLength`.

## Msg/UpdateGroupMembers

Group members can be updated with the `UpdateGroupMembersRequest`.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L73-L85

In the list of `MemberUpdates`, an existing member can be removed by setting its weight to 0.

It's expecting to fail if the signer is not the admin of the group.

## Msg/UpdateGroupAdmin

The `UpdateGroupAdminRequest` can be used to update a group admin.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L90-L101

It's expecting to fail if the signer is not the admin of the group.

## Msg/UpdateGroupMetadata

The `UpdateGroupMetadataRequest` can be used to update a group metadata.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L106-L117

It's expecting to fail if:
- new metadata length is greater than some `MaxMetadataLength`.
- the signer is not the admin of the group.

## Msg/CreateGroupAccount

A new group account can be created with the `MsgCreateGroupAccountRequest`, which has an admin address, a group id, a decision policy and some optional metadata bytes.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L126-L141

It's expecting to fail if metadata length is greater than some `MaxMetadataLength`.

## Msg/UpdateGroupAccountAdmin

The `UpdateGroupAccountAdminRequest` can be used to update a group account admin.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L150-L161

It's expecting to fail if the signer is not the admin of the group account.

## Msg/UpdateGroupAccountMetadata

The `UpdateGroupAccountMetadataRequest` can be used to update a group account metadata.

+++ https://github.com/regen-network/regen-ledger/blob/24290a0f2b219ac831482006952cb7a11d173ecf/proto/regen/group/v1alpha1/tx.proto#L183-L194

It's expecting to fail if:
- new metadata length is greater than some `MaxMetadataLength`.
- the signer is not the admin of the group.
