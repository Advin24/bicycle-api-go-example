# Chart Naming:

## Coming from Templates

Notice the use of `"name"` instead of `"fullname"` within these charts
`"name"` and `"fullname"` declared/set from templates/_helpers.tpl

_ex_: `name: {{ template "name" . }}`

## Override the Templates

Set the `.Values.nameOverride`

Notice, this corresponds to the below:
```
    cat values.yml
      # random stuff
      nameOverride: boop-boop-beep-beep-app
      # more random stuff
```

# Chart Labeling:

## Coming from Templates

Notice, usage of the function `defined.labels` will generate default labels for your application. This keeps it easy to set over multiple files and delete/update at once.

_ex_: `{{ include "defined.labels" . | indent 4 }} //Notice, the indent is a seperate function for YAML parsing`

You can also add to these labels by adding to the `metadata.labels` field in values.yml

_ex_:
```
    cat values.yml
      # random stuff
      metadata:
        labels:
            beep: beep
            boop: boop
            iam: "a-computer"
      # more random stuff
```