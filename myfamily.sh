curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id|tonumber)) | .connections.relative'
