curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq --argjson '.[] | select(.id == '$HERO_ID') | .connections.relatives' | sed 's/"//g'
