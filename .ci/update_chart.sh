#!/bin/bash -ex

NAME=unknown
TEAM=unknown
REPOSITORY=unknown
SHA=$(cat uropa-$TRAVIS_TAG.tgz | git hash-object --stdin)
CONTENT=$(cat uropa-$TRAVIS_TAG.tgz | base64)

curl -X PUT -H "Authorization: token $GITHUB_TOKEN" -H "Content-Type: application/json" \
  "https://api.github.com/repos/$TEAM/$REPOSITORY/contents/$NAME-$TRAVIS_TAG.tgz" \
  -d '{
        "message":"Chart update",
        "committer": {
          "name": "Chart Release Bot",
          "email": "charts@ninjaneers.de"
        },
        "content":"'$CONTENT'",
        "sha":"'$SHA'"
      }'
