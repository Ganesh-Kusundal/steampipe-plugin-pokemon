# Table: pokemon_shape

Shapes used for sorting Pokémon in a Pokédex.

## Examples

### Basic info

```sql
select
  name,
  id,
  awesome_names,
  names,
  pokemon_species
from
  pokemon_shape
```

### List all Pokémon having shape as 'ball'

```sql
select
  name,
  id,
  pokemon_species
from
  pokemon_shape
where
  name = 'ball'
```  
  
### List Pokémon shape whose id is 6

```sql
select
  name,
  id,
  pokemon_species
from
  pokemon_shape
where
  id = 6
```