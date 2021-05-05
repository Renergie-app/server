package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"renergie-server/graph/generated"
	"renergie-server/graph/model"
)

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
		kwh := kwcToKwh(postalCodeToDepartment(input.PostalCode), kwc) * PercentageWithOrientationAndAngle(*facade.Orientation, *facade.Angle) / 100

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
				cost += 2550 * facade.KWC
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

func (r *queryResolver) WindTurbine(ctx context.Context, input *model.WindTurbineInput) (*model.WindTurbineResponse, error) {

	var windDataSpeed float64 = windSpeed(postalCodeToDepartment(input.PostalCode))

	kwh := 1000*windDataSpeed - 10000
	if input.Type == model.WindTurbineTypeVertical {
		kwh *= 2
	}
	var cost int
	if input.Type == model.WindTurbineTypeHorizontal {
		cost = 15850 * input.Amount
	} else {
		cost = 18950 * input.Amount
	}
	return &model.WindTurbineResponse{
		Cost:           cost,
		PowerOutputKwh: kwh * float64(input.Amount),
		Profit:         0.082 * kwh * float64(input.Amount),
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
type FacadeCalc struct {
	FacadeInput         *model.Facade
	KWC                 float64
	KWH                 float64
	AmountOfSolarPanels int
}

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
func windSpeed(department string) float64 {
	m := map[string]int{
		"01": 15,
		"02": 21,
		"03": 15,
		"04": 21,
		"05": 18,
		"06": 18,
		"07": 18,
		"08": 21,
		"09": 15,
		"10": 18,
		"11": 25,
		"12": 25,
		"13": 30,
		"14": 25,
		"15": 15,
		"16": 18,
		"17": 21,
		"18": 18,
		"19": 18,
		"20": 18,
		"21": 15,
		"22": 25,
		"23": 18,
		"24": 18,
		"25": 15,
		"26": 18,
		"27": 25,
		"28": 21,
		"29": 25,
		"30": 25,
		"31": 15,
		"32": 15,
		"33": 18,
		"34": 30,
		"35": 25,
		"36": 18,
		"37": 18,
		"38": 18,
		"39": 15,
		"40": 15,
		"41": 18,
		"42": 15,
		"43": 15,
		"44": 25,
		"45": 18,
		"46": 18,
		"47": 15,
		"48": 25,
		"49": 21,
		"50": 25,
		"51": 25,
		"52": 15,
		"53": 21,
		"54": 15,
		"55": 18,
		"56": 25,
		"57": 18,
		"58": 15,
		"59": 21,
		"60": 21,
		"61": 18,
		"62": 25,
		"63": 15,
		"64": 15,
		"65": 15,
		"66": 18,
		"67": 15,
		"68": 15,
		"69": 15,
		"70": 15,
		"71": 15,
		"72": 21,
		"73": 15,
		"74": 15,
		"75": 21,
		"76": 25,
		"77": 18,
		"78": 21,
		"79": 21,
		"80": 25,
		"81": 25,
		"82": 18,
		"83": 25,
		"84": 25,
		"85": 25,
		"86": 18,
		"87": 18,
		"88": 15,
		"89": 18,
		"90": 15,
		"91": 18,
		"92": 21,
		"93": 21,
		"94": 21,
		"95": 21,
	}
	return float64(m[department])
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
