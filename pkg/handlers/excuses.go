package handlers

import "math/rand"

func getExcuse() string {
	excuses := []string{"clock speed",
		"solar flares",
		"electromagnetic radiation from satellite debris",
		"static from nylon underwear",
		"static from plastic slide rules",
		"global warming",
		"poor power conditioning",
		"static buildup",
		"doppler effect",
		"hardware stress fractures",
		"magnetic interference from money/credit cards",
		"dry joints on cable plug",
		"we're waiting for [the phone company] to fix that line",
		"sounds like a Windows problem, try calling Microsoft support",
		"temporary routing anomaly",
		"somebody was calculating pi on the server",
		"fat electrons in the lines",
		"excess surge protection",
		"floating point processor overflow",
		"divide-by-zero error",
		"POSIX compliance problem",
		"monitor resolution too high",
		"improperly oriented keyboard",
		"network packets travelling uphill (use a carrier pigeon)",
		"Decreasing electron flux",
		"first Saturday after first full moon in Winter",
		"radiosity depletion",
		"CPU radiator broken",
		"It works the way the Wang did, what's the problem",
		"positron router malfunction",
		"cellular telephone interference",
		"techtonic stress"}
	excuseIndex := rand.Intn(len(excuses) - 1)
	return excuses[excuseIndex]
}
