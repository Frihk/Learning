curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json \
| jq -r --argjson id "$HERO_ID" '.[] | select(.id == $id) | .connections.relative'
