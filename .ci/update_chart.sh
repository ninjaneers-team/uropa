#!/bin/bash -ex

SHA=$(echo $(cat $NAME-$VERSION.tgz) | git hash-object --stdin)
CONTENT=$(echo $(cat $NAME-$VERSION.tgz) | base64)

echo "Upload chart $NAME"

curl -X PUT -H "Authorization: token $GITHUB_TOKEN" -H "Content-Type: application/json" \
  "https://api.github.com/repos/$TEAM/$REPOSITORY/contents/$NAME-$VERSION.tgz" \
  -d "{
        \"message\":\"Chart update\",
        \"committer\": {
          \"name\": \"Chart Release Bot\",
          \"email\": \"charts@ninjaneers.de\"
        },
        \"content\":\"$CONTENT\",
        \"sha\":\"$SHA\"
      }"
