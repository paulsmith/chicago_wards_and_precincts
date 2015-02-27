FILES = chicago_2015_wards_topo.json chicago_2015_precincts_topo.json chicago_2015_wards.zip chicago_2015_precincts.zip chicago_2015_wards.kml chicago_2015_precincts.kml chicago_2015_wards_precincts.sql.gz

files: $(FILES)

all: import postprocess files

import:
	go run import.go

.PHONY : import

postprocess:
	psql -f process-precincts.sql -d chicago
	psql -f make-wards.sql -d chicago

.PHONY : postprocess

clean:
	-rm -f $(FILES)

chicago_2015_wards.shp:
	ogr2ogr -f ESRI\ Shapefile -overwrite $@ PG:dbname=chicago wards

zip: chicago_2015_wards.zip chicago_2015_precincts.zip

chicago_2015_wards.zip: chicago_2015_wards.shp
	mkdir chicago_2015_wards
	cp chicago_2015_wards.{shp,dbf,prj,shx} chicago_2015_wards
	zip -r $@ chicago_2015_wards
	rm -rf chicago_2015_wards

chicago_2015_precincts.shp:
	ogr2ogr -f ESRI\ Shapefile -overwrite $@ PG:dbname=chicago precincts

chicago_2015_precincts.zip: chicago_2015_precincts.shp
	mkdir chicago_2015_precincts
	cp chicago_2015_precincts.{shp,dbf,prj,shx} chicago_2015_precincts
	zip -r $@ chicago_2015_precincts
	rm -rf chicago_2015_precincts

chicago_2015_wards.json: chicago_2015_wards.shp
	ogr2ogr -f GeoJSON $@ $<

chicago_2015_precincts.json: chicago_2015_precincts.shp
	ogr2ogr -f GeoJSON $@ $<

topo: chicago_2015_wards_topo.json chicago_2015_precincts_topo.json

chicago_2015_wards_topo.json: chicago_2015_wards.json
	topojson --id-property ward -o $@ $<

chicago_2015_precincts_topo.json: chicago_2015_precincts.json
	topojson --id-property id -o $@ $<

chicago_2015_wards.kml: chicago_2015_wards.shp
	ogr2ogr -f KML $@ $<

chicago_2015_precincts.kml: chicago_2015_precincts.shp
	ogr2ogr -f KML $@ $<

chicago_2015_wards_precincts.sql:
	pg_dump chicago > $@

chicago_2015_wards_precincts.sql.gz: chicago_2015_wards_precincts.sql
	gzip < $< > $@
