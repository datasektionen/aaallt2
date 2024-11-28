# Aaallt

Run locally with: `go run .`

The different versions of Aaallt, with different links, are loaded dynamically depending on the subdomain and loads the corresponding YAML file with the same name in [links/](links/).

[links/systems.yml](links/systems.yml) are the default links loaded for every subdomain.

To create a new version/link tree, just create a new YAML file in [links/](links/) with the same format and point a subdomain to this system.
