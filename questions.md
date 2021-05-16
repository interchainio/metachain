# Questions and Unknowns

This section is designed to capture some of the outstanding questions and
unknowns with this project.

1. I see decision policies as being categorized into three groups based on
   what point in time the weight of the voter is counted. This could be a) Upon
   the start of the propsal, b) Upon receiving the vote and c) Upon the end of
   the proposal period. Each of these require different tally systems to
   accompany the decision policy and thus could be a lot of overhead. Do we want
   to allow such flexibility or should the `dao` module be fixed on a single
   style (in this case, my preference would be upon the start of the proposal)?