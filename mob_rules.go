package main
import ("fmt"
"math/rand"
)
func randInt(min, max int) int {
	return min+rand.Intn(max-min)
}

func main() {
  var baseac,numberofattacks,basedamage,baseattackbonus,d20needed,totaldamage,totalhits,attackersneeded int
  var roll1,roll2,roll3,roll4 int
  fmt.Print("Type your stats in the following order AC,NumberOf Attacks,AttackBonus,Damage:") 
  fmt.Print("") 
  fmt.Scanln(&baseac,&numberofattacks,&baseattackbonus,&basedamage)

  d20needed=baseac-baseattackbonus
if  d20needed >= 20{
	d20needed= 20
	attackersneeded = 20
}
if  d20needed == 19{
	attackersneeded = 10
}
if  d20needed == 17 ||d20needed == 18{
	attackersneeded = 10
}
if  d20needed == 15 ||d20needed == 16{
	attackersneeded = 5
}
if  d20needed == 13 || d20needed == 14{
	attackersneeded = 4
}
if  d20needed >=6 && d20needed <= 12{
	attackersneeded = 2
}
if  d20needed >=1 && d20needed <= 5{
	attackersneeded = 1
}

totalhits=numberofattacks/attackersneeded
totaldamage=basedamage*totalhits
roll1 = randInt(1,20)
roll2 = randInt(1,20)
roll3 = randInt(1,20)
roll4 = randInt(1,20)

fmt.Println("BaseAC:", baseac, "Number Of Attacks:", numberofattacks,"AttackBonus:",baseattackbonus,"BaseDamage:",basedamage)
fmt.Println("D20Needed:", d20needed, "attackers Needed",attackersneeded)
fmt.Println("Total Hits:",totalhits, "TotalDamage",totaldamage)
fmt.Printf("I am a random roll: %d %d %d %d\n",roll1, roll2,roll3,roll4)
}