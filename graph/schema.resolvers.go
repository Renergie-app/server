package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"renergie-server/graph/generated"
	"renergie-server/graph/model"
)

type FacadeCalc struct {
	FacadeInput         *model.Facade
	KWC                 float64
	KWH                 float64
	AmountOfSolarPanels int
}

func (r *queryResolver) SolarPanel(ctx context.Context, input *model.SolarPanelInput) (*model.SolarPanelResponse, error) {
	var facadesResult []*model.FacadeResponse
	var facades []FacadeCalc

	totalAmountOfSolarPanels := 0
	totalKWH := 0.0
	totalKWC := 0.0
	stateFinancialHelp := 0.0
	totalProfit := 0.0
	totalCost := 0.0

	for _, facade := range input.Facades {
		amountOfSolarPanels := int(math.Floor(*facade.Surface / 1.7))
		kwc := float64(amountOfSolarPanels) * 0.3
		kwh := kwcToKwh(postalCodeToDepartment(input.PostalCode), kwc) * PercentageWithOrientationAndAngle(*facade.Orientation, *facade.Angle)
		totalAmountOfSolarPanels += amountOfSolarPanels
		totalKWH += kwh
		totalKWC += kwc
		facades = append(facades, FacadeCalc{
			FacadeInput:         facade,
			KWC:                 kwc,
			KWH:                 kwh,
			AmountOfSolarPanels: amountOfSolarPanels,
		})
	}

	for _, facade := range facades {
		profit := 0.0
		cost := 0.0
		//	PROFIT
		if totalKWC <= 3 {
			if input.SellEverything {
				profit = 0.18 * facade.KWH
			} else {
				profit = 0.10 * facade.KWH
			}
		} else if totalKWC <= 9 {
			if input.SellEverything {
				profit = 0.15 * facade.KWH
			} else {
				profit = 0.10 * facade.KWH
			}
		} else if totalKWC <= 36 {
			if input.SellEverything {
				profit = 0.11 * facade.KWH
			} else {
				profit = 0.06 * facade.KWH
			}
		} else {
			if input.SellEverything {
				profit = 0.09 * facade.KWH
			} else {
				profit = 0.06 * facade.KWH
			}
		}
		totalProfit += profit
		//	COST
		if totalKWC <= 3 {
			if input.IntegratedInBuilding {
				cost += 255 * facade.KWC
			} else {
				cost += 2850 * facade.KWC
			}

		} else if totalKWC <= 6 {
			if input.IntegratedInBuilding {
				cost += 2250 * facade.KWC
			} else {
				cost += 2500 * facade.KWC
			}

		} else if totalKWC <= 9 {
			if input.IntegratedInBuilding {
				cost += 2000 * facade.KWC
			} else {
				cost += 2200 * facade.KWC
			}

		} else {
			if input.IntegratedInBuilding {
				cost += 1950 * facade.KWC
			} else {
				cost += 2100 * facade.KWC
			}

		}
		totalCost += cost

		facadesResult = append(facadesResult, &model.FacadeResponse{
			PowerOutputKwh:      facade.KWH,
			Cost:                cost,
			Profit:              profit,
			AmountOfSolarPanels: facade.AmountOfSolarPanels,
			Orientation:         facade.FacadeInput.Orientation,
			Angle:               facade.FacadeInput.Angle,
		})

	}

	if totalKWC <= 3 {
		stateFinancialHelp += 380 * totalKWC
	} else if totalKWC <= 9 {
		stateFinancialHelp += 280 * totalKWC
	} else if totalKWC <= 36 {
		stateFinancialHelp += 160 * totalKWC
	} else {
		stateFinancialHelp += 80 * totalKWC
	}

	if totalCost > 0 {
		totalCost += 2000
	}

	return &model.SolarPanelResponse{
		TotalPowerOutputKwh:      totalKWH,
		TotalProfit:              totalProfit,
		TotalCost:                totalCost,
		TotalAmountOfSolarPanels: totalAmountOfSolarPanels,
		StateFinancialHelp:       stateFinancialHelp,
		PerFacadeDetails:         facadesResult,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func kwcToKwh(department string, kwc float64) float64 {
	m := map[string]int{
		"01": 1050,
		"02": 900,
		"03": 1150,
		"04": 1300,
		"05": 1300,
		"06": 1300,
		"07": 1300,
		"08": 900,
		"09": 1150,
		"10": 900,
		"11": 1300,
		"12": 1150,
		"13": 1300,
		"14": 900,
		"15": 1150,
		"16": 1150,
		"17": 1150,
		"18": 1050,
		"19": 1150,
		"20": 1300,
		"21": 900,
		"22": 1050,
		"23": 1050,
		"24": 1150,
		"25": 900,
		"26": 1300,
		"27": 900,
		"28": 900,
		"29": 1050,
		"30": 1300,
		"31": 1150,
		"32": 1150,
		"33": 1150,
		"34": 1300,
		"35": 1050,
		"36": 1050,
		"37": 1050,
		"38": 1150,
		"39": 900,
		"40": 1150,
		"41": 1050,
		"42": 1050,
		"43": 1150,
		"44": 1050,
		"45": 900,
		"46": 1150,
		"47": 1150,
		"48": 1150,
		"49": 1050,
		"50": 900,
		"51": 900,
		"52": 900,
		"53": 1050,
		"54": 900,
		"55": 900,
		"56": 1050,
		"57": 900,
		"58": 1050,
		"59": 900,
		"60": 900,
		"61": 900,
		"62": 900,
		"63": 1050,
		"64": 1150,
		"65": 1150,
		"66": 1300,
		"67": 900,
		"68": 900,
		"69": 1050,
		"70": 900,
		"71": 1050,
		"72": 1050,
		"73": 1150,
		"74": 1050,
		"75": 900,
		"76": 900,
		"77": 900,
		"78": 900,
		"79": 1050,
		"80": 900,
		"81": 1150,
		"82": 1150,
		"83": 1300,
		"84": 1300,
		"85": 1150,
		"86": 1050,
		"87": 1050,
		"88": 900,
		"89": 900,
		"90": 1150,
		"91": 900,
		"92": 900,
		"93": 900,
		"94": 900,
		"95": 900,
	}
	i := m[department]
	return float64(i) * kwc

}
func postalCodeToDepartment(postalCode string) string {
	if len(postalCode) == 5 {
		return postalCode[0:2]
	} else {
		return ""
	}
}
func PercentageWithOrientationAndAngle(orientation model.Orientation, angle int) float64 {
	switch orientation {
	case model.OrientationSouth:
		if angle <= 30 {
			return getPercentage(0, 30, 93, 100, angle)
		} else if angle <= 60 {
			return getPercentage(30, 60, 100, 91, angle)
		} else {
			return getPercentage(60, 90, 91, 68, angle)
		}
	case model.OrientationEast:
		if angle <= 30 {
			return getPercentage(0, 30, 93, 90, angle)
		} else if angle <= 60 {
			return getPercentage(30, 60, 90, 78, angle)
		} else {
			return getPercentage(60, 90, 78, 55, angle)
		}
	case model.OrientationSouthEast:
		if angle <= 30 {
			return getPercentage(0, 30, 93, 96, angle)
		} else if angle <= 60 {
			return getPercentage(30, 60, 96, 88, angle)
		} else {
			return getPercentage(60, 90, 88, 66, angle)
		}
	case model.OrientationSouthWest:
		if angle <= 30 {
			return getPercentage(0, 30, 93, 96, angle)
		} else if angle <= 60 {
			return getPercentage(30, 60, 96, 88, angle)
		} else {
			return getPercentage(60, 90, 88, 66, angle)
		}
	case model.OrientationWest:
		if angle <= 30 {
			return getPercentage(0, 30, 93, 90, angle)
		} else if angle <= 60 {
			return getPercentage(30, 60, 90, 78, angle)
		} else {
			return getPercentage(60, 90, 78, 55, angle)
		}
	default:
		return 93

	}
}
func getPercentage(alpha int, beta int, a int, b int, x int) float64 {
	if x/30 >= 1 && x%30 == 0 {
		return float64(a) + ((float64(30))*float64(b-a))/float64(beta-alpha)
	} else {
		return float64(a) + ((float64(x%30))*float64(b-a))/float64(beta-alpha)
	}

}
