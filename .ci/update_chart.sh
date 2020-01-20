#!/bin/bash -ex

CONTENT=$(base64 $NAME-$VERSION.tgz | tr -d \\n)

echo "Upload chart $NAME"

curl -X PUT -H "Authorization: token $GITHUB_TOKEN" -H "Content-Type: application/json" \
  "https://api.github.com/repos/$TEAM/$REPOSITORY/contents/$NAME-$VERSION.tgz" \
  -d "{
        \"message\":\"Chart update\",
        \"committer\": {
          \"name\": \"Chart Release Bot\",
          \"email\": \"charts@ninjaneers.de\"
        },
        \"content\":\"$CONTENT\"
      }"
