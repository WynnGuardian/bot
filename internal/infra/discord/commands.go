package discord

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "survey",
			Description: "Survey commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "list",
					Description: "List all open surveys",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "open",
					Description: "*(ADMIN ONLY)* Open a new criteria survey",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "duration-days",
							Description: "Duration in days",
							Required:    true,
						},
					},
				},
				{
					Name:        "approve",
					Description: "*(ADMIN ONLY)* Approve a currently open survey",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "survey",
							Description: "Target survey",
							Required:    true,
						},
					},
				},
				{
					Name:        "discard",
					Description: "*(ADMIN ONLY)* Discard a currently waiting approval survey.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "survey",
							Description: "Target survey",
							Required:    true,
						},
					},
				},
				{
					Name:        "close",
					Description: "*(ADMIN ONLY)* Close a currently open survey, so it can receive no more votes.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Survey's item name",
							Required:    true,
						},
					},
				},
				{
					Name:        "cancel",
					Description: "*(ADMIN ONLY)* Cancel a currently open survey",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Survey's item name",
							Required:    true,
						},
					},
				},
				{
					Name:        "fill",
					Description: "Fill a survey",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:        "item",
			Description: "Item commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "weight",
					Description: "Item weight command. Item must be encoded in UTF16.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item",
							Description: "UTF16 encoded Item to be weigh",
							Required:    true,
						},
					},
				},
				{
					Name:        "track",
					Description: "Find a specific item in the database with a tracking code",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "tracking-code",
							Description: "Item tracking code",
							Required:    true,
						},
					},
				},
				{
					Name:        "authenticate",
					Description: "*(ADMIN ONLY)* Authenticate the item and make it trackable/rankable.",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item",
							Description: "UTF16 encoded item",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "mc-uuid",
							Description: "Player minecraft UUID",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "dc-id",
							Description: "Player Discord ID",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionBoolean,
							Name:        "public-owner",
							Description: "Make item owner public",
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:        "rank",
			Description: "Rank commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "update",
					Description: "*(ADMIN ONLY)* Force update an item rank",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
					},
				},
				{
					Name:        "view",
					Description: "View an item rank",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
					},
				},
			},
		},
		{
			Name:        "criteria",
			Description: "Criteria commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "create",
					Description: "*(ADMIN ONLY)* Create a criteria for an item",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "id-name",
							Description: "ID Name",
							Required:    true,
						},
					},
				},
				{
					Name:        "delete",
					Description: "*(ADMIN ONLY)* Delete item criteria",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "id-name",
							Description: "ID name",
							Required:    true,
						},
					},
				},
				{
					Name:        "view",
					Description: "View item criteria",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
					},
				},
				{
					Name:        "update",
					Description: "Update item criteria",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "item-name",
							Description: "Target item",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "criteria-id",
							Description: "Stat id",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "value",
							Description: "Value",
							Required:    true,
						},
					},
				},
			},
		},
	}
)
