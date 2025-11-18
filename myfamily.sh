curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq --argjson id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives' | sed 's/"//g'
