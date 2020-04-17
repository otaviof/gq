#!/usr/bin/env bats

GQ="./gq"

TMPL='{{ .status.host }}'
EXAMPLE_YAML="test/data/file.yaml"
EXAMPLE_JSON="test/data/file.json"

@test "gq: using defults" {
    res=$(${GQ} "${TMPL}" ${EXAMPLE_YAML})
    [ "${res}" == "expected" ]
}

@test "gq: json file" {
    res=$(${GQ} --type=json "${TMPL}" ${EXAMPLE_JSON})
    [ "${res}" == "expected" ]
}

@test "gq: multiple files (YAML)" {
    res=$(${GQ} --type=yaml "${TMPL}" ${EXAMPLE_YAML} ${EXAMPLE_YAML})
    [ "${res}" == "expectedexpected" ]
}

@test "gq: multiple files (JSON)" {
    res=$(${GQ} --type=json "${TMPL}" ${EXAMPLE_JSON} ${EXAMPLE_JSON})
    [ "${res}" == "expectedexpected" ]
}

@test "gq: stdin (YAML)" {
    res="$(cat ${EXAMPLE_YAML} |${GQ} "${TMPL}")"
    [ "${res}" == "expected" ]
}