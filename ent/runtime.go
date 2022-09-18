// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/lrstanley/ent-relay-conn-bug/ent/guild"
	"github.com/lrstanley/ent-relay-conn-bug/ent/guildsettings"
	"github.com/lrstanley/ent-relay-conn-bug/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	guildMixin := schema.Guild{}.Mixin()
	guildMixinFields0 := guildMixin[0].Fields()
	_ = guildMixinFields0
	guildFields := schema.Guild{}.Fields()
	_ = guildFields
	// guildDescCreateTime is the schema descriptor for create_time field.
	guildDescCreateTime := guildMixinFields0[0].Descriptor()
	// guild.DefaultCreateTime holds the default value on creation for the create_time field.
	guild.DefaultCreateTime = guildDescCreateTime.Default.(func() time.Time)
	// guildDescUpdateTime is the schema descriptor for update_time field.
	guildDescUpdateTime := guildMixinFields0[1].Descriptor()
	// guild.DefaultUpdateTime holds the default value on creation for the update_time field.
	guild.DefaultUpdateTime = guildDescUpdateTime.Default.(func() time.Time)
	// guild.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	guild.UpdateDefaultUpdateTime = guildDescUpdateTime.UpdateDefault.(func() time.Time)
	// guildDescName is the schema descriptor for name field.
	guildDescName := guildFields[1].Descriptor()
	// guild.NameValidator is a validator for the "name" field. It is called by the builders before save.
	guild.NameValidator = func() func(string) error {
		validators := guildDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// guildDescFeatures is the schema descriptor for features field.
	guildDescFeatures := guildFields[2].Descriptor()
	// guild.DefaultFeatures holds the default value on creation for the features field.
	guild.DefaultFeatures = guildDescFeatures.Default.([]string)
	// guildDescIconHash is the schema descriptor for icon_hash field.
	guildDescIconHash := guildFields[3].Descriptor()
	// guild.IconHashValidator is a validator for the "icon_hash" field. It is called by the builders before save.
	guild.IconHashValidator = guildDescIconHash.Validators[0].(func(string) error)
	// guildDescIconURL is the schema descriptor for icon_url field.
	guildDescIconURL := guildFields[4].Descriptor()
	// guild.IconURLValidator is a validator for the "icon_url" field. It is called by the builders before save.
	guild.IconURLValidator = guildDescIconURL.Validators[0].(func(string) error)
	// guildDescLarge is the schema descriptor for large field.
	guildDescLarge := guildFields[6].Descriptor()
	// guild.DefaultLarge holds the default value on creation for the large field.
	guild.DefaultLarge = guildDescLarge.Default.(bool)
	// guildDescMemberCount is the schema descriptor for member_count field.
	guildDescMemberCount := guildFields[7].Descriptor()
	// guild.DefaultMemberCount holds the default value on creation for the member_count field.
	guild.DefaultMemberCount = guildDescMemberCount.Default.(int)
	// guildDescPermissions is the schema descriptor for permissions field.
	guildDescPermissions := guildFields[9].Descriptor()
	// guild.DefaultPermissions holds the default value on creation for the permissions field.
	guild.DefaultPermissions = guildDescPermissions.Default.(uint64)
	// guildDescRegion is the schema descriptor for region field.
	guildDescRegion := guildFields[10].Descriptor()
	// guild.RegionValidator is a validator for the "region" field. It is called by the builders before save.
	guild.RegionValidator = guildDescRegion.Validators[0].(func(string) error)
	// guildDescSystemChannelFlags is the schema descriptor for system_channel_flags field.
	guildDescSystemChannelFlags := guildFields[11].Descriptor()
	// guild.SystemChannelFlagsValidator is a validator for the "system_channel_flags" field. It is called by the builders before save.
	guild.SystemChannelFlagsValidator = guildDescSystemChannelFlags.Validators[0].(func(string) error)
	guildsettingsMixin := schema.GuildSettings{}.Mixin()
	guildsettingsMixinFields0 := guildsettingsMixin[0].Fields()
	_ = guildsettingsMixinFields0
	guildsettingsFields := schema.GuildSettings{}.Fields()
	_ = guildsettingsFields
	// guildsettingsDescCreateTime is the schema descriptor for create_time field.
	guildsettingsDescCreateTime := guildsettingsMixinFields0[0].Descriptor()
	// guildsettings.DefaultCreateTime holds the default value on creation for the create_time field.
	guildsettings.DefaultCreateTime = guildsettingsDescCreateTime.Default.(func() time.Time)
	// guildsettingsDescUpdateTime is the schema descriptor for update_time field.
	guildsettingsDescUpdateTime := guildsettingsMixinFields0[1].Descriptor()
	// guildsettings.DefaultUpdateTime holds the default value on creation for the update_time field.
	guildsettings.DefaultUpdateTime = guildsettingsDescUpdateTime.Default.(func() time.Time)
	// guildsettings.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	guildsettings.UpdateDefaultUpdateTime = guildsettingsDescUpdateTime.UpdateDefault.(func() time.Time)
	// guildsettingsDescDefaultMaxClones is the schema descriptor for default_max_clones field.
	guildsettingsDescDefaultMaxClones := guildsettingsFields[1].Descriptor()
	// guildsettings.DefaultDefaultMaxClones holds the default value on creation for the default_max_clones field.
	guildsettings.DefaultDefaultMaxClones = guildsettingsDescDefaultMaxClones.Default.(int)
	// guildsettingsDescRegexMatch is the schema descriptor for regex_match field.
	guildsettingsDescRegexMatch := guildsettingsFields[2].Descriptor()
	// guildsettings.DefaultRegexMatch holds the default value on creation for the regex_match field.
	guildsettings.DefaultRegexMatch = guildsettingsDescRegexMatch.Default.(string)
	// guildsettings.RegexMatchValidator is a validator for the "regex_match" field. It is called by the builders before save.
	guildsettings.RegexMatchValidator = func() func(string) error {
		validators := guildsettingsDescRegexMatch.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(regex_match string) error {
			for _, fn := range fns {
				if err := fn(regex_match); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// guildsettingsDescContactEmail is the schema descriptor for contact_email field.
	guildsettingsDescContactEmail := guildsettingsFields[3].Descriptor()
	// guildsettings.DefaultContactEmail holds the default value on creation for the contact_email field.
	guildsettings.DefaultContactEmail = guildsettingsDescContactEmail.Default.(string)
	// guildsettings.ContactEmailValidator is a validator for the "contact_email" field. It is called by the builders before save.
	guildsettings.ContactEmailValidator = guildsettingsDescContactEmail.Validators[0].(func(string) error)
}
