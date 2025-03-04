### Overview

`docs/bundle_descriptions.json` contains both autogenerated as well as manually written
descriptions for the json schema. Specifically
1. `resources` : almost all descriptions are autogenerated from the OpenAPI spec
2. `environments` : almost all descriptions are copied over from root level entities (eg: `bundle`, `artifacts`)
3. `bundle` : manually editted
4. `include` : manually editted
5. `workspace` : manually editted
6. `artifacts` : manually editted

These descriptions are rendered in the inline documentation in an IDE

### SOP: Add schema descriptions for new fields in bundle config

1. You can autogenerate empty descriptions for the new fields by running
`databricks bundle schema --only-docs > ~/databricks/bundle/schema/docs/bundle_descriptions.json`
2. Manually edit bundle_descriptions.json to add your descriptions
3. Build again to embed the new `bundle_descriptions.json` into the binary (`go build`)
4. Again run `databricks bundle schema --only-docs > ~/databricks/bundle/schema/docs/bundle_descriptions.json` to copy over any applicable descriptions to `environments`
5. push to repo


### SOP: Update descriptions in resources from a newer openapi spec

1. Run `databricks bundle schema --only-docs --openapi PATH_TO_SPEC > ~/databricks/bundle/schema/docs/bundle_descriptions.json`
2. push to repo
