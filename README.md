Chicago wards and precincts shapefiles
======================================

* [Download](https://paulsmith.github.io/chicago_wards_and_precincts/)

These Chicago wards and precincts boundary shapefiles are effective as of early
2015. As of Thursday, February 26, 2015, the shapefiles on the City of Chicago’s
data portal are stale—they are the boundaries from before the city council
voted to redraw them in 2012.

Source
------

The original source material was provided by the Chicago Board of
Elections in response to a FOIA request, in the form of 50 shapefiles
corresponding to precinct boundaries and attributes in each Chicago ward.

I applied minimal post-processing to produce various formats, and to generate
ward boundaries by aggregating precincts and unioning their geometries. See the
[`etl`](./tree/master/etl)
directory for the programs that perform the processing.
