# EVChargerFinder

EVChargerFinder is a simple API endpoint that locates the nearest EV Chargers to a given coordinate.

Many of the available charger locators are not kept up to date with the latest chargers, or indicate the locations of only the chargers for their networks, or are otherwise unreliable. EVChargerFinder differentiates itself from other charger locators primarily through its data source. Its data comes from the National Renewable Energy Laboratory (NREL) database. As the NREL is dedicated towards renewable energies, it has no biases toward any one network and it has incentive to keep the list up to date.

## Compilation

EVChargerFinder requires Go 1.18 or higher.

```
$ git clone https://github.com/kevin-fwu/EVChargerFinder
$ cd EVChargerFinder/src
$ go build -o evchargerfinder
```

This creates a binary `evchargerfinder`

## Running the binary

EVChargerFinder requires a configuration file to operate. See `dist/evchargerfinder.json.example` for an example configuration file.

To run the server:

```
$ ./evchargerfinder -conf=evchargerfinder.json
```

One could also perform a one-time lookup:

```
$ ./evchargerfinder -conf=evchargerfinder.json -latitude=40.533 -longitude=-74.3727 -distance=10 -limit=10 | jq
```

`jq` is not required, but provides a cleaner output by formatting the output.

## Implementation Details

EVChargerFinder uses the NREL fuel station database as its data source. This list is updated once per day. EVChargerFinder processes the data, then builds a K-D Tree. This tree is the main station filter: a K-D tree enables quick filtering by locating stations which are approximately within a provided distance. After creating a list of approximate stations, EVChargerFinder calculates the distance between the given location and each station using the Haversine Formula. From there, it sorts and returns the final list.
