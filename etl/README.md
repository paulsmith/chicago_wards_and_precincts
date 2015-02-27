Post-processing source precinct data from Chicago Board of Elections
====================================================================

The rough steps are:

* Import into a PostgreSQL database with PostGIS extension using `ogr2ogr`.
* Clean precincts table to split out ward and precinct IDs.
* Create new wards table and insert wards by grouping precincts by ward number
  and unioning (`ST_Union`) their geometries.
* Export from database to various formats using `ogr2ogr` and `topojson`.

See `Makefile` for details.
