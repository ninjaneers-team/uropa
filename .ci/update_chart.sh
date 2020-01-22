#!/bin/bash -ex

echo "Upload chart $NAME"

curl -X PUT -H "Authorization: token $GITHUB_TOKEN" -H "Content-Type: application/json" \
  "https://api.github.com/repos/$TEAM/$REPOSITORY/contents/$NAME-$VERSION.tgz" \
  --data @<(printf '%s' "{
        \"message\":\"Chart update\",
        \"committer\": {
          \"name\": \"Chart Release Bot\",
          \"email\": \"charts@ninjaneers.de\"
        },
        \"content\":\"$(base64 $NAME-$VERSION.tgz | tr -d \\n)\"
      }") > /dev/null