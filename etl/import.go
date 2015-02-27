package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func importWard(ward int) error {
	shapefile := shapefiles[ward-1]
	args := []string{"-f", "PostgreSQL", "PG:dbname=chicago", shapefile, "-nln", "precincts", "-t_srs", "EPSG:4326"}
	if ward == 1 {
		args = append(args, "-lco", "OVERWRITE=yes", "-overwrite")
	} else {
		args = append(args, "-update", "-append")
	}
	cmd := exec.Command("ogr2ogr", args...)
	var out bytes.Buffer
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "-- %s\n", out.String())
		return err
	}
	return nil
}

func main() {
	for ward := 1; ward <= len(shapefiles); ward++ {
		if err := importWard(ward); err != nil {
			log.Fatalf("importing ward %d: %s", ward, err)
		}
		fmt.Fprintf(os.Stderr, "imported ward %d\n", ward)
	}
}

var shapefiles = []string{
	"./WD 1-25/Ward 01/P_Ward_01_Precincts_Final.shp",
	"./WD 1-25/Ward 02/P_Ward_02_Precincts_Final.shp",
	"./WD 1-25/Ward 03/P_Wd_03_PREC12_Final.shp",
	"./WD 1-25/Ward 04/P_Wd_04_PREC12_Final.shp",
	"./WD 1-25/Ward 05/P_Wd_05_PREC12_Final.shp",
	"./WD 1-25/Ward 06/P_Wd_06_PREC12_Final.shp",
	"./WD 1-25/Ward 07/P_Wd_07_PREC12_Final.shp",
	"./WD 1-25/Ward 08/P_Wd_08_PREC12_Final.shp",
	"./WD 1-25/Ward 09/P_Ward_09_Precincts_Final.shp",
	"./WD 1-25/Ward 10/P_Wd_10_PREC12_Final.shp",
	"./WD 1-25/Ward 11/P_Wd_11_PREC12_Final.shp",
	"./WD 1-25/Ward 12/P_Wd_12_PREC12_Final.shp",
	"./WD 1-25/Ward 13/P_Wd_13_Precincts_Final.shp",
	"./WD 1-25/Ward 14/P_Wd_14_PREC12_Final.shp",
	"./WD 1-25/Ward 15/P_Wd_15_PREC12_Final.shp",
	"./WD 1-25/Ward 16/P_Wd_16_PREC12_Final.shp",
	"./WD 1-25/Ward 17/P_Wd_17_PREC12_Final.shp",
	"./WD 1-25/Ward 18/P_Wd_18_PREC12_Final.shp",
	"./WD 1-25/Ward 19/P_Wd_19_Precincts_Final.shp",
	"./WD 1-25/Ward 20/P_Ward_20_Precincts_Final.shp",
	"./WD 1-25/Ward 21/P_Ward_21_Precincts_Final.shp",
	"./WD 1-25/Ward 22/P_Ward_22_Precincts_Final.shp",
	"./WD 1-25/Ward 23/P_Wd_23_Precincts_Final.shp",
	"./WD 1-25/Ward 24/P_Ward_24_Precincts_Final.shp",
	"./WD 1-25/Ward 25/P_Ward_25_Precincts_Final.shp",
	"./WD 26-50/Ward 26/P_Wd_26_Precincts_Final.shp",
	"./WD 26-50/Ward 27/P_Ward_27_Precincts_Final.shp",
	"./WD 26-50/Ward 28/P_Wd_28_PREC12_Final.shp",
	"./WD 26-50/Ward 29/P_Wd_29_PREC12_Final.shp",
	"./WD 26-50/Ward 30/P_Wd_30_PREC12_Final.shp",
	"./WD 26-50/Ward 31/P_Ward_31_Precincts_Final.shp",
	"./WD 26-50/Ward 32/P_Ward_32_Precincts_Final.shp",
	"./WD 26-50/Ward 33/P_Ward_33_Precincts_Final.shp",
	"./WD 26-50/Ward 34/P_Ward_34_Precincts_Final.shp",
	"./WD 26-50/Ward 35/P_Ward_35_Precincts_Final.shp",
	"./WD 26-50/Ward 36/P_Ward_36_Precincts_Final.shp",
	"./WD 26-50/Ward 37/P_Wd_37_PREC12_Final.shp",
	"./WD 26-50/Ward 38/P_Wd_38_PREC12_Final.shp",
	"./WD 26-50/Ward 39/P_Ward_39_Precincts_Final.shp",
	"./WD 26-50/Ward 40/P_Ward_40_Precincts_Final.shp",
	"./WD 26-50/Ward 41/P_Wd_41_PREC12_FinalFinal.shp",
	"./WD 26-50/Ward 42/P_Ward_42_Precincts_Final.shp",
	"./WD 26-50/Ward 43/P_Wd_43_PREC12_Final.shp",
	"./WD 26-50/Ward 44/P_Ward_44_Precincts_Final.shp",
	"./WD 26-50/Ward 45/P_Wd_45_PREC12_Final.shp",
	"./WD 26-50/Ward 46/P_Ward_46_Precincts_Final.shp",
	"./WD 26-50/Ward 47/P_Wd_47_PREC12_Final.shp",
	"./WD 26-50/Ward 48/P_Wd_48_PREC12_Final.shp",
	"./WD 26-50/Ward 49/P_Wd_49_PREC12_Final.shp",
	"./WD 26-50/Ward 50/P_Ward_50_Precincts_Final.shp",
}
