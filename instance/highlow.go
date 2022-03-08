// Copyright (C) 2022 | Mr.Gibson (Yaimsputnik5) | Team Bête Noire
//
// This source code has been released under the GNU Affero General Public
// License v3.0. A copy of this license is available at
// https://www.gnu.org/licenses/agpl-3.0.en.html

package instance

import (

	"github.com/Yaimsputnik5/MemerBreaker/discord"
	"github.com/Yaimsputnik5/MemerBreaker/instance/scheduler"
)

func (in *Instance) hl(msg discord.Message) {
	hint := exp.hl.FindStringSubmatch(msg.Embeds[0].Description)[1]
	if hint[0] > 50 {
		in.sdlr.ResumeWithCommand(&scheduler.Command{
			Actionrow: 1,
			Button: 1,
			Message: msg,
			Log: "Responding with Low",
		})
	}
	if hint[0] < 50 {
		in.sdlr.ResumeWithCommand(&scheduler.Command{
			Actionrow: 1,
			Button: 3,
			Message: msg,
			Log: "Responding with High",
		})
	}
	if hint[0] == 50 {
		in.sdlr.ResumeWithCommand(&scheduler.Command{
			Actionrow: 1,
			Button: 3,
			Message: msg,
			Log: "Responding with Jackpot",
		})
	}
}
