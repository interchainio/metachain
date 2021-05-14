<!--
order: 4
-->

# Events

The group module emits the following events:

## EventCreateGroup

| Type                                  | Attribute Key | Attribute Value                       |
|---------------------------------------|---------------|---------------------------------------|
| message                               | action        | /regen.group.v1alpha1.Msg/CreateGroup |
| regen.group.v1alpha1.EventCreateGroup | group_id      | {groupId}                             |

## EventUpdateGroup

| Type                                  | Attribute Key | Attribute Value                                                 |
|---------------------------------------|---------------|-----------------------------------------------------------------|
| message                               | action        | /regen.group.v1alpha1.Msg/UpdateGroup{Admin\|Metadata\|Members} |
| regen.group.v1alpha1.EventUpdateGroup | group_id      | {groupId}                                                       |

## EventCreateGroupAccount

| Type                                         | Attribute Key | Attribute Value                              |
|----------------------------------------------|---------------|----------------------------------------------|
| message                                      | action        | /regen.group.v1alpha1.Msg/CreateGroupAccount |
| regen.group.v1alpha1.EventCreateGroupAccount | address       | {groupAccountAddress}                        |

## EventUpdateGroupAccount

| Type                                         | Attribute Key | Attribute Value                                                               |
|----------------------------------------------|---------------|-------------------------------------------------------------------------------|
| message                                      | action        | /regen.group.v1alpha1.Msg/UpdateGroupAccount{Admin\|Metadata\|DecisionPolicy} |
| regen.group.v1alpha1.EventUpdateGroupAccount | address       | {groupAccountAddress}                                                         |