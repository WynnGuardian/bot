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
					Description: "Open a new criteria survey",
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
					Description: "Approve an currently open survey",
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
					Description: "Discard an currently waiting approval survey",
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
					Description: "Close a survey",
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
					Description: "Cancel a survey",
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
					Description: "Authenticate player item",
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
					Description: "Update an item rank",
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
	}
)
