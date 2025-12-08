Chris Pickett had a way better solution for this.

He calculated the distance between all pairs of points just once, then sorted the resulting list of distances.
From there, he was able to just iterate through the distances and connect up each endpoint, while simultaneously
adding to (or merging) circuits.

https://github.com/parnic/advent-of-code-2025/blob/29fac32b011a6c795e989ec822cfcb1cde406a33/days/08.go#L37-L107
