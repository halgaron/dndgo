package main
import ("fmt"
"math/rand"
)
func randInt(min, max int) int {
	return min+rand.Intn(max-min)
}

func main() {
  var baseac,numberofattacks,baseaveragedamage,baseattackbonus,d20needed,totaldamage,totalhits int
//   var roll1,roll2,roll3,roll4 int
var damagesummary,damagerolltotal,damageroll int
var damagetotal int
  var hitsummary,DMhitsummary string
  var damagedicenumber,damagedice,damagedicemodifier int
  var totalattack int
  var roll int
  var targetnamenumber int
  var Azathalia,Keldan,Banian,Evangeline,Tim int
var targetplayerac = [6]int{0,23,24,18,19,15}
var targetname = [6]string{"Custom","Azathalia","Keldan","Banian","Evangeline","Tim"}
// Anything above 6 is considered to be the target's CUSTOM AC.
  fmt.Print("Number of attacks vs Azathalia, Keldan, Banian, Evangeline, Tim \n") 
  fmt.Scanln(&Azathalia,&Keldan,&Banian,&Evangeline,&Tim)
  fmt.Print("Input base attack bonus, base average damage: Or alternatively add number of damage dice + dice number modifier") 
  fmt.Scanln(&baseattackbonus,&baseaveragedamage,&damagedicenumber,&damagedice,&damagedicemodifier)
 
  var Targetonehitsummary,Targettwohitsummary,Targetthreehitsummary,Targetfourhitsummary,Targetfivehitsummary string
  var TargetoneDMhitsummary,TargettwoDMhitsummary,TargetthreeDMhitsummary,TargetfourDMhitsummary,TargetfiveDMhitsummary string
  Targetonehitsummary = ""
  Targettwohitsummary = ""
  Targetthreehitsummary = ""
  Targetfourhitsummary = ""
  Targetfivehitsummary = ""

  TargetoneDMhitsummary = ""
  TargettwoDMhitsummary = ""
  TargetthreeDMhitsummary = ""
  TargetfourDMhitsummary = ""
  TargetfiveDMhitsummary = ""

  var TargetoneDamageSummary,TargettwoDamageSummary,TargetthreeDamageSummary,TargetfourDamageSummary,TargetfiveDamageSummary int
  var TargetOneTotalHits,TargetTwoTotalHits,TargetThreeTotalHits,TargetFourTotalHits,TargetFiveTotalHits int
  TargetoneDamageSummary=0

  damagetotal=0
//   Azathalia = 0 
//   Keldan = 0
//   Banian = 0 
//   Evangeline = 0 
//   Tim = 0
// if specificplayer > 0 && specificplayer < 5{
// 	baseac = targetplayerac[specificplayer]
// 	targetnamenumber=specificplayer
// }else{
// 	baseac = specificplayer
// 	targetnamenumber=0
// }






totalhits=0

// roll1 = randInt(1,20)
// if  roll1 >= d20needed {
// 	totalhits = totalhits+1
// }
// roll2 = randInt(1,20)
// if  roll2 >= d20needed {
// 	totalhits = totalhits+1
// }
// roll3 = randInt(1,20)
// if  roll3 >= d20needed {
// 	totalhits = totalhits+1
// }
// roll4 = randInt(1,20)
// if  roll4 >= d20needed {
// 	totalhits = totalhits+1
// }





	if Azathalia > 0 {
		baseac = targetplayerac[1]
		d20needed=baseac-baseattackbonus
		targetnamenumber=1
		numberofattacks=Azathalia
		totalhits=0
        hitsummary=""
        DMhitsummary =""
		damageroll=0
		damagerolltotal=0
		damagesummary=0
	for i:=0; i<Azathalia; i++{
		roll = randInt(1,21)
		totalattack=roll+baseattackbonus
		hitsummary=Targetonehitsummary
		DMhitsummary=TargetoneDMhitsummary
		if  roll == 20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(Critical HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber*2,"d",damagedice, "+",damagedicemodifier)
			for i:=0; i<damagedicenumber*2; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
				damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Crit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll >= d20needed && roll<20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber,"d",damagedice, "+",damagedicemodifier)		
			for i:=0; i<damagedicenumber; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
			damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Hit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll < d20needed && roll==1 {
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(Critical MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Crit Miss,"
			damagesummary=damagesummary+0
		}
		if  roll < d20needed && roll > 1 {
			fmt.Println(i,"Attack Roll vs :",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Miss,"
			damagesummary=damagesummary+0
		}

		Targetonehitsummary=hitsummary
		TargetoneDMhitsummary=DMhitsummary
		TargetoneDamageSummary=damagesummary
	}
	TargetOneTotalHits=totalhits
	totaldamage=baseaveragedamage*TargetOneTotalHits
	fmt.Println("Totals VS:", targetname[targetnamenumber])
	fmt.Println("BaseAC:", baseac, "Number Of Attacks:", numberofattacks)
	//fmt.Println("D20Needed:", d20needed)
	fmt.Println("Total Hits:",totalhits, "Total Average Damage",totaldamage,"\n")
	}

	if Keldan > 0 {
		baseac = targetplayerac[2]
		d20needed=baseac-baseattackbonus
		targetnamenumber=2
		numberofattacks=Keldan
		totalhits=0
        hitsummary=""
        DMhitsummary =""
		damageroll=0
		damagerolltotal=0
		damagesummary=0
	

	for i:=0; i<Keldan; i++{
		roll = randInt(1,21)
		totalattack=roll+baseattackbonus
		hitsummary=Targettwohitsummary
		DMhitsummary=TargettwoDMhitsummary
			if  roll == 20 {
				totalhits = totalhits+1
				damagerolltotal=0
				fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(Critical HIT!)")
				fmt.Println("Damage Roll:", damagedicenumber*2,"d",damagedice, "+",damagedicemodifier)
				for i:=0; i<damagedicenumber*2; i++{
					damageroll=randInt(1,damagedice+1)
					damagerolltotal=damagerolltotal+damageroll
					fmt.Println("Damage Roll:",damageroll)
					}
					damagetotal=damagerolltotal+damagedicemodifier
					fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
				hitsummary=hitsummary+"Hit,"
				DMhitsummary=DMhitsummary+"Crit,"
				damagesummary=damagesummary+damagetotal
			}
			if  roll >= d20needed && roll<20 {
				totalhits = totalhits+1
				damagerolltotal=0
				fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(HIT!)")
				fmt.Println("Damage Roll:", damagedicenumber,"d",damagedice, "+",damagedicemodifier)		
				for i:=0; i<damagedicenumber; i++{
					damageroll=randInt(1,damagedice+1)
					damagerolltotal=damagerolltotal+damageroll
					fmt.Println("Damage Roll:",damageroll)
					}
				damagetotal=damagerolltotal+damagedicemodifier
					fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
				hitsummary=hitsummary+"Hit,"
				DMhitsummary=DMhitsummary+"Hit,"
				damagesummary=damagesummary+damagetotal
			}
			if  roll < d20needed && roll==1 {
				fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(Critical MISS!)")
				hitsummary=hitsummary+"Miss,"
				DMhitsummary=DMhitsummary+"Crit Miss,"
				damagesummary=damagesummary+0
			}
			if  roll < d20needed && roll > 1 {
				fmt.Println(i,"Attack Roll vs :",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(MISS!)")
				hitsummary=hitsummary+"Miss,"
				DMhitsummary=DMhitsummary+"Miss,"
				damagesummary=damagesummary+0
			}
		
	
		Targettwohitsummary=hitsummary
		TargettwoDMhitsummary=DMhitsummary
		TargettwoDamageSummary=damagesummary
	}
	TargetTwoTotalHits=totalhits
	totaldamage=baseaveragedamage*TargetTwoTotalHits
	fmt.Println("Totals VS:", targetname[targetnamenumber])
	fmt.Println("BaseAC:", baseac, "Number Of Attacks:", numberofattacks)
	//fmt.Println("D20Needed:", d20needed)
	fmt.Println("Total Hits:",totalhits, "Total Average Damage",totaldamage,"\n")
	}

	if Banian > 0 {
		baseac = targetplayerac[3]
		d20needed=baseac-baseattackbonus
		targetnamenumber=3
		numberofattacks=Banian
		totalhits=0
        hitsummary=""
        DMhitsummary =""
		damageroll=0
		damagerolltotal=0
		damagesummary=0

	for i:=0; i<Banian; i++{
		roll = randInt(1,21)
		totalattack=roll+baseattackbonus
		hitsummary=Targetthreehitsummary
		DMhitsummary=TargetthreeDMhitsummary
		if  roll == 20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(Critical HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber*2,"d",damagedice, "+",damagedicemodifier)
			for i:=0; i<damagedicenumber*2; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
				damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Crit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll >= d20needed && roll<20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber,"d",damagedice, "+",damagedicemodifier)		
			for i:=0; i<damagedicenumber; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
			damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Hit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll < d20needed && roll==1 {
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(Critical MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Crit Miss,"
			damagesummary=damagesummary+0
		}
		if  roll < d20needed && roll > 1 {
			fmt.Println(i,"Attack Roll vs :",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Miss,"
			damagesummary=damagesummary+0
		}
		
		Targetthreehitsummary=hitsummary
		TargetthreeDMhitsummary=DMhitsummary
		TargetthreeDamageSummary=damagesummary
	}
	TargetThreeTotalHits=totalhits
	totaldamage=baseaveragedamage*TargetThreeTotalHits
	fmt.Println("Totals VS:", targetname[targetnamenumber])
	fmt.Println("BaseAC:", baseac, "Number Of Attacks:", numberofattacks)
	//fmt.Println("D20Needed:", d20needed)
	fmt.Println("Total Hits:",totalhits, "Total Average Damage",totaldamage,"\n")
	}

	if Evangeline > 0 {
		baseac = targetplayerac[4]
		d20needed=baseac-baseattackbonus
		targetnamenumber=4
		numberofattacks=Evangeline
		totalhits=0
        hitsummary=""
        DMhitsummary =""
		damageroll=0
		damagerolltotal=0
		damagesummary=0

	for i:=0; i<Evangeline; i++{
		roll = randInt(1,21)
		totalattack=roll+baseattackbonus
		hitsummary=Targetfourhitsummary
		DMhitsummary=TargetfourDMhitsummary
		if  roll == 20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(Critical HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber*2,"d",damagedice, "+",damagedicemodifier)
			for i:=0; i<damagedicenumber*2; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
				damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Crit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll >= d20needed && roll<20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber,"d",damagedice, "+",damagedicemodifier)		
			for i:=0; i<damagedicenumber; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
			damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Hit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll < d20needed && roll==1 {
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(Critical MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Crit Miss,"
			damagesummary=damagesummary+0
		}
		if  roll < d20needed && roll > 1 {
			fmt.Println(i,"Attack Roll vs :",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Miss,"
			damagesummary=damagesummary+0
		}
		
		Targetfourhitsummary=hitsummary
		TargetfourDMhitsummary=DMhitsummary
		TargetfourDamageSummary=damagesummary
	}
	TargetFourTotalHits=totalhits
	totaldamage=baseaveragedamage*TargetFourTotalHits
	fmt.Println("Totals VS:", targetname[targetnamenumber])
	fmt.Println("BaseAC:", baseac, "Number Of Attacks:", numberofattacks)
	//fmt.Println("D20Needed:", d20needed)
	fmt.Println("Total Hits:",totalhits, "Total Average Damage",totaldamage,"\n")
	}

	if Tim > 0 {
		baseac = targetplayerac[5]
		d20needed=baseac-baseattackbonus
		targetnamenumber=5
		numberofattacks=Tim
		totalhits=0
        hitsummary=""
        DMhitsummary =""
		damageroll=0
		damagerolltotal=0
		damagesummary=0

	for i:=0; i<Tim; i++{
		roll = randInt(1,21)
		totalattack=roll+baseattackbonus
		hitsummary = Targetfivehitsummary
		DMhitsummary=TargetfiveDMhitsummary
		if  roll == 20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(Critical HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber*2,"d",damagedice, "+",damagedicemodifier)
			for i:=0; i<damagedicenumber*2; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
				damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Crit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll >= d20needed && roll<20 {
			totalhits = totalhits+1
			damagerolltotal=0
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack, "(HIT!)")
			fmt.Println("Damage Roll:", damagedicenumber,"d",damagedice, "+",damagedicemodifier)		
			for i:=0; i<damagedicenumber; i++{
				damageroll=randInt(1,damagedice+1)
				damagerolltotal=damagerolltotal+damageroll
				fmt.Println("Damage Roll:",damageroll)
				}
			damagetotal=damagerolltotal+damagedicemodifier
				fmt.Println("Damage Roll Totals:",damagerolltotal,"+",damagedicemodifier,"=", damagetotal, "\n")
			hitsummary=hitsummary+"Hit,"
			DMhitsummary=DMhitsummary+"Hit,"
			damagesummary=damagesummary+damagetotal
		}
		if  roll < d20needed && roll==1 {
			fmt.Println(i,"Attack Roll vs:",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(Critical MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Crit Miss,"
			damagesummary=damagesummary+0
		}
		if  roll < d20needed && roll > 1 {
			fmt.Println(i,"Attack Roll vs :",targetname[targetnamenumber], roll, "+",baseattackbonus,"=", totalattack,"(MISS!)")
			hitsummary=hitsummary+"Miss,"
			DMhitsummary=DMhitsummary+"Miss,"
			damagesummary=damagesummary+0
		}
	
		Targetfivehitsummary=hitsummary
		TargetfiveDMhitsummary=DMhitsummary
		TargetfiveDamageSummary=damagesummary
	}
	TargetFiveTotalHits=totalhits
	totaldamage=baseaveragedamage*TargetFiveTotalHits
	fmt.Println("Totals VS:", targetname[targetnamenumber])
	fmt.Println("BaseAC:", baseac, "Number Of Attacks:", numberofattacks)
	//fmt.Println("D20Needed:", d20needed)
	fmt.Println("Total Hits:",totalhits, "Total Average Damage",totaldamage,"\n")
	}

	if Targetonehitsummary == "" {
		Targetonehitsummary = "No Attacks"
	}
	if Targettwohitsummary == "" {
		Targettwohitsummary = "No Attacks"
	}
	if Targetthreehitsummary == "" {
		Targetthreehitsummary = "No Attacks"
	}
	if Targetfourhitsummary == "" {
		Targetfourhitsummary = "No Attacks"
	}
	if Targetfivehitsummary == "" {
		Targetfivehitsummary = "No Attacks"
	}
	fmt.Println("Summary VS:", targetname[1],Targetonehitsummary)
	fmt.Println("Summary VS:", targetname[2],Targettwohitsummary)
	fmt.Println("Summary VS:", targetname[3],Targetthreehitsummary)
	fmt.Println("Summary VS:", targetname[4],Targetfourhitsummary)
	fmt.Println("Summary VS:", targetname[5],Targetfivehitsummary, "\n")

	fmt.Println("DM Summary VS:", targetname[1],Targetonehitsummary,"Average Damage:", baseaveragedamage*TargetOneTotalHits, "Rolled Damage Total:",TargetoneDamageSummary)
	fmt.Println("DM Summary VS:", targetname[2],Targettwohitsummary,"Average Damage:", baseaveragedamage*TargetTwoTotalHits, "Rolled Damage Total:",TargettwoDamageSummary)
	fmt.Println("DM Summary VS:", targetname[3],Targetthreehitsummary,"Average Damage:", baseaveragedamage*TargetThreeTotalHits, "Rolled Damage Total:",TargetthreeDamageSummary)
	fmt.Println("DM Summary VS:", targetname[4],Targetfourhitsummary,"Average Damage:", baseaveragedamage*TargetFourTotalHits, "Rolled Damage Total:",TargetfourDamageSummary)
	fmt.Println("DM Summary VS:", targetname[5],Targetfivehitsummary,"Average Damage:", baseaveragedamage*TargetFiveTotalHits, "Rolled Damage Total:",TargetfiveDamageSummary)
}

