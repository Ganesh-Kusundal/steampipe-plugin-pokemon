package pokemon

import (
        "context"

		"github.com/mtslzr/pokeapi-go"
		"github.com/mtslzr/pokeapi-go/structs"

        "github.com/turbot/steampipe-plugin-sdk/grpc/proto"
        "github.com/turbot/steampipe-plugin-sdk/plugin"
        "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePokemonShapes(ctx context.Context) *plugin.Table {
        return &plugin.Table{
                Name:        "pokemon_shape",
                Description: "Shapes used for sorting Pokémon in a Pokédex.",
                List: &plugin.ListConfig{
                        Hydrate: listPokemonShape,
                },
                Get: &plugin.GetConfig{
                        KeyColumns: plugin.AnyColumn([]string{"name"}),
                        Hydrate: getPokemonShape,
                        ShouldIgnoreError: isNotFoundError([]string{"invalid character 'N' looking for beginning of value"}),
                },
                Columns: []*plugin.Column{
                        {
                                Name:        "id",
                                Description: "The identifier for this resource.",
                                Type:        proto.ColumnType_INT,
                                Hydrate:     getPokemonShape,
                                Transform:   transform.FromGo(),
                        },
                        {
                                Name:        "name",
                                Description: "The name for this resource.",
                                Type:        proto.ColumnType_STRING,
                        },
                        {
                                Name:        "awesome_names",
                                Description: "The \"scientific\" name of this Pokémon shape listed in different languages.",
                                Type:        proto.ColumnType_JSON,
                                Hydrate:     getPokemonShape,
                        },
                        {
                                Name:        "names",
                                Description: "The name of this resource listed in different languages.",
                                Type:        proto.ColumnType_JSON,
                                Hydrate:     getPokemonShape,
                        },
                        {
                                Name:        "pokemon_species",
                                Description: "A list of the Pokémon species that have this shape.",
                                Type:        proto.ColumnType_JSON,
                                Hydrate:     getPokemonShape,
                        },
                        // Standard columns
                        {
                                Name:        "title",
                                Description: "Title of the resource.",
                                Type:        proto.ColumnType_STRING,
                                Transform:   transform.FromField("Name"),
                        },
                },
        }
}

func listPokemonShape(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("listPokemonShape")

	offset := 0

	for true {
		resources, err := pokeapi.Resource("pokemon-shape", offset)

		if err != nil {
			plugin.Logger(ctx).Error("pokemon_shapes.listPokemonShape", "query_error", err)
			return nil, err
		}

		for _, shape := range resources.Results {
			d.StreamListItem(ctx, shape)
		}

		// No next URL returned
		if len(resources.Next) == 0 {
			break
		}

		urlOffset, err := extractUrlOffset(resources.Next)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_shapes.listPokemonShape", "extract_url_offset_error", err)
			return nil, err
		}

		// Set next offset
		offset = urlOffset
	}

	return nil, nil
}

func getPokemonShape(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getPokemonShape")

	var name string

	if h.Item != nil {
		result := h.Item.(structs.Result)
		name = result.Name
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	logger.Debug("Name", name)

	pokemonShape, err := pokeapi.PokemonShape(name)

	if err != nil {
		plugin.Logger(ctx).Error("pokemon_shapes.getPokemonShape", "query_error", err)
		return nil, err
	}

	return pokemonShape, nil
}