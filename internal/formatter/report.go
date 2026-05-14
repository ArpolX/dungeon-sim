package formatter

import "fmt"

func Registered(playerID int) string {
	return fmt.Sprintf("Player [%d] registered", playerID)
}

func EnteredDungeon(playerID int) string {
	return fmt.Sprintf("Player [%d] entered the dungeon", playerID)
}

func KilledMonster(playerID int) string {
	return fmt.Sprintf("Player [%d] killed the monster", playerID)
}

func NextFloor(playerID int) string {
	return fmt.Sprintf("Player [%d] went to the next floor", playerID)
}

func PreviousFloor(playerID int) string {
	return fmt.Sprintf("Player [%d] went to the previous floor", playerID)
}

func EnteredBossFloor(playerID int) string {
	return fmt.Sprintf("Player [%d] entered the boss's floor", playerID)
}

func KilledBoss(playerID int) string {
	return fmt.Sprintf("Player [%d] killed the boss", playerID)
}

func LeftDungeon(playerID int) string {
	return fmt.Sprintf("Player [%d] left the dungeon", playerID)
}

func RestoredHealth(playerID, hp int) string {
	return fmt.Sprintf("Player [%d] has restored [%d] of health", playerID, hp)
}

func ReceivedDamage(playerID, hp int) string {
	return fmt.Sprintf("Player [%d] received [%d] of damage", playerID, hp)
}

func PlayerDead(playerID int) string {
	return fmt.Sprintf("Player [%d] is dead", playerID)
}

func CannotContinue(playerID int, reason string) string {
	return fmt.Sprintf("Player [%d] cannot continue due to [%s]", playerID, reason)
}

func Disqualified(playerID int) string {
	return fmt.Sprintf("Player [%d] is disqualified", playerID)
}

func ImpossibleMove(playerID int, eventID int) string {
	return fmt.Sprintf(
		"Player [%d] makes impossible move [%d]",
		playerID,
		eventID,
	)
}
