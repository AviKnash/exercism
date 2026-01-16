package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
    	case "two":
        	return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
        case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
        	return 9
        case "ten":
        	return 10
        case "jack":
        	return 10
        case "queen":
        	return 10
        case "king":
        	return 10
        default:
            return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

		// 1. Always split aces
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	playerTotal := ParseCard(card1) + ParseCard(card2)

	// 2. Blackjack rules
	if playerTotal == 21 {
		switch dealerCard {
		case "ace", "king", "queen", "jack", "ten":
			return "S"
		default:
			return "W"
		}
	}

	dealerValue := ParseCard(dealerCard)

	// 3. Stand on 17–20
	if playerTotal >= 17 && playerTotal <= 20 {
		return "S"
	}

	// 4. 12–16 logic
	if playerTotal >= 12 && playerTotal <= 16 {
		if dealerValue >= 7 {
			return "H"
		}
		return "S"
	}

	// 5. Always hit on 11 or lower
	return "H"
}
